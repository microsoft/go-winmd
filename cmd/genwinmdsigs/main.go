// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"debug/pe"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/microsoft/go-winmd"
	"github.com/microsoft/go-winmd/coded"
	"github.com/microsoft/go-winmd/flags"
	"github.com/microsoft/go-winmd/genwinsyscallproto"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	source := flag.String("source", "", "The win32metadata file to parse and generate signatures for.")
	output := flag.String("o", "", "Output file name (prints to stdout if omitted).")
	filter := flag.String("filter", "", "If specified, only generate signatures for methods 'namespace::method' that match this regex.")
	templatePath := flag.String("template", "", "If specified, a text/template style template file to use. Use '{{.SysCalls}}' for the '//sys' content.")

	flag.Parse()

	if *source == "" {
		return errors.New("no source passed: nothing to do")
	}
	var filterRegexp *regexp.Regexp
	if *filter != "" {
		var err error
		filterRegexp, err = regexp.Compile(*filter)
		if err != nil {
			return err
		}
	}

	var contentTemplate *template.Template
	if *templatePath != "" {
		var err error
		contentTemplate, err = template.ParseFiles(*templatePath)
		if err != nil {
			return err
		}
	}

	start := time.Now()

	pefile, err := pe.Open(*source)
	if err != nil {
		return err
	}
	defer pefile.Close()
	f, err := winmd.New(pefile)
	if err != nil {
		return err
	}

	var b strings.Builder

	if err := writePrototypes(&b, f, filterRegexp); err != nil {
		return err
	}

	content := b.String()
	var finalContent string

	if contentTemplate != nil {
		var templateBuilder strings.Builder
		if err := contentTemplate.Execute(&templateBuilder, struct{ SysCalls string }{content}); err != nil {
			return err
		}
		finalContent = templateBuilder.String()
	} else {
		finalContent = content
	}

	end := time.Now()
	log.Printf("Time elapsed to produce sys signatures: %v\n", end.Sub(start))

	if *output != "" {
		return os.WriteFile(*output, []byte(finalContent), 0666)
	}
	log.Println("Printing signature results because no output path was specified:")
	log.Println("---")
	log.Println(finalContent)
	return nil
}

func writePrototypes(b *strings.Builder, f *winmd.Metadata, filterRegexp *regexp.Regexp) error {
	context, err := genwinsyscallproto.NewContext(f)
	if err != nil {
		return err
	}

	// Track all typedefs by namespace + "::" + name for typeref resolution later.
	// TODO: Determine if disambiguating types in winmd files with multiple modules is necessary.
	allTypeDefs := make(map[string]winmd.Index)

	// Track all typerefs that are used in parameters. These will be resolved later if possible.
	usedTypeRefs := make(map[winmd.Index]struct{})
	// Track all explicitly used typedefs.
	// TODO: Add entrypoint for explicitly using a typedef.
	usedTypeDefs := make(map[winmd.Index]struct{})

	firstType := true
	for i := uint32(0); i < f.Tables.TypeDef.Len; i++ {
		r, err := f.Tables.TypeDef.Record(winmd.Index(i))
		if err != nil {
			return err
		}

		allTypeDefs[r.Namespace.String()+"::"+r.Name.String()] = winmd.Index(i)

		if !strings.Contains(r.Name.String(), "Apis") {
			continue
		}

		firstMethod := true
		for j := r.MethodList.Start; j < r.MethodList.End; j++ {
			md, err := f.Tables.MethodDef.Record(j)
			if err != nil {
				return err
			}
			spec := fmt.Sprintf("%v::%v.%v", r.Namespace, r.Name, md.Name)
			if filterRegexp != nil && !filterRegexp.MatchString(spec) {
				continue
			}

			// Find the DllImport pseudo-custom attribute (§II.21.2.1) for the module name.
			var moduleName string
			if implMap, ok := context.MethodDefImplMap[j]; ok {
				// TODO: Map of parsed module refs?
				mr, err := f.Tables.ModuleRef.Record(implMap.ImportScope)
				if err != nil {
					return err
				}
				moduleName = strings.ToLower(mr.Name.String())
				if moduleName == "kernel32" {
					moduleName = ""
				}
			}
			methodName := md.Name.String()

			// Handle newlines and section headers between chunks of methods from the same typedef.
			if firstMethod {
				if firstType {
					firstType = false
				} else {
					b.WriteString("\n\n")
				}
				firstMethod = false
				b.WriteString("// APIs for ")
				b.WriteString(r.Namespace.String())
				b.WriteString("\n")
			} else {
				b.WriteString("\n")
			}

			sig, err := f.MethodDefSignature(md.Signature)
			if err != nil {
				return err
			}
			for _, p := range sig.Param {
				switch p.Kind {
				case winmd.SigParamKind_ByValue:
					if err := addUsedTypeRefs(usedTypeRefs, f, &p.Type); err != nil {
						return err
					}
				default:
					return fmt.Errorf("unexpected SigParamKind for %v param %#v", spec, p)
				}
			}

			if err := context.WriteMethod(b, md, &sig, moduleName, methodName); err != nil {
				// Include context in the error for diag purposes.
				// writeSys may have partially written into b. This is actually convenient for diag.
				lines := strings.Split(b.String(), "\n")
				if len(lines) > 5 {
					lines = lines[len(lines)-5:]
				}

				return fmt.Errorf(
					"error context: \n---\n%v\n---\nfailed to write sys line for %v.Apis method %v: %v",
					strings.Join(lines, "\n"), r.Namespace, methodName, err)
			}
		}
	}

	for i := range usedTypeRefs {
		ref, err := f.Tables.TypeRef.Record(i)
		if err != nil {
			return err
		}
		switch ref.ResolutionScope.Tag {
		case coded.ResolutionScope_Module:
			// This TypeRef refers to the current module. We don't need to look at the Index.
		default:
			log.Printf("Skipping %v::%v: not defined in current module\n", ref.Namespace, ref.Name)
			continue
		}
		if def, ok := allTypeDefs[ref.Namespace.String()+"::"+ref.Name.String()]; ok {
			usedTypeDefs[def] = struct{}{}
		} else {
			return fmt.Errorf("TypeRef %v::%v has Module resolution scope, but is not found in the module", ref.Namespace, ref.Name)
		}
	}

	if len(usedTypeDefs) > 0 {
		b.WriteString("\n\n// Types used in generated APIs\n\n")
		// Sort indices to make output stable.
		usedTypeDefIndices := make([]winmd.Index, 0, len(usedTypeDefs))
		for index := range usedTypeDefs {
			usedTypeDefIndices = append(usedTypeDefIndices, index)
		}
		sort.Slice(usedTypeDefIndices, func(i, j int) bool {
			return usedTypeDefIndices[i] < usedTypeDefIndices[j]
		})
		for _, index := range usedTypeDefIndices {
			record, err := f.Tables.TypeDef.Record(index)
			if err != nil {
				return err
			}
			if err := context.WriteTypeDef(b, record); err != nil {
				return err
			}
			b.WriteString("\n")
		}
	}
	return nil
}

func addUsedTypeRefs(used map[winmd.Index]struct{}, f *winmd.Metadata, t *winmd.SigType) error {
	switch t.Kind {
	case flags.ElementType_PTR:
		innerType := t.Value.(winmd.SigType)
		if err := addUsedTypeRefs(used, f, &innerType); err != nil {
			return err
		}
	case flags.ElementType_VALUETYPE:
		switch v := t.Value.(type) {
		case winmd.CodedIndex:
			switch v.Tag {
			case coded.TypeDefOrRefOrSpec_TypeRef:
				used[v.Index] = struct{}{}
			default:
				return fmt.Errorf("unexpected CodedIndex tag: %v", v.Tag)
			}
		default:
			return fmt.Errorf("unexpected type for VALUETYPE value %#v", v)
		}
	case flags.ElementType_ARRAY:
		return fmt.Errorf("not implemented: array parameter of type %#v", t)
	}
	return nil
}
