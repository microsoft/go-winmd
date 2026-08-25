// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/microsoft/go-winmd/cmd/gowinmd/internal/gowinmd"
	"github.com/microsoft/go-winmd/winmd"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	source := flag.String("source", "", "The win32metadata file to parse and generate signatures for.")
	output := flag.String("output", "", "Output file name (prints to stdout if omitted).")
	formatFlag := flag.String("format", "", "Output format. Required. Supported values: mkwinsyscall.")
	projectionFlag := flag.String("projection", "raw", "Signature projection. Supported values: raw, idiomatic.")

	flag.Parse()

	if *source == "" {
		return errors.New("source is required: pass the path to a win32metadata file using the -source flag")
	}
	if *formatFlag != "mkwinsyscall" {
		return fmt.Errorf("format is required: pass -format mkwinsyscall (got %q)", *formatFlag)
	}
	var projection gowinmd.Projection
	switch *projectionFlag {
	case "raw":
		projection = gowinmd.ProjectionRaw
	case "idiomatic":
		projection = gowinmd.ProjectionIdiomatic
	default:
		return fmt.Errorf("unsupported projection %q: use raw or idiomatic", *projectionFlag)
	}

	inputFiles := flag.Args()
	if len(inputFiles) == 0 {
		return errors.New("no input files provided: pass Go files containing //winmd directives as non-flag arguments")
	}

	// Parse //winmd directives and infer package name from the input files.
	filter, selectedTypes, pkg, err := parseInputFiles(inputFiles)
	if err != nil {
		return err
	}
	_ = selectedTypes

	start := time.Now()

	f, err := winmd.Open(*source)
	if err != nil {
		return err
	}

	b := map[gowinmd.Arch]*strings.Builder{
		gowinmd.Arch386:   {},
		gowinmd.ArchAMD64: {},
		gowinmd.ArchARM64: {},
		gowinmd.ArchAll:   {},
		gowinmd.ArchNone:  {},
	}

	if err := writeSelectionsWithProjection(b, f, filter, selectedTypes, projection); err != nil {
		return err
	}

	for arch, w := range b {
		if w.Len() == 0 {
			continue
		}
		content := w.String()
		finalContent := generateFileContent(content, pkg)

		formattedContent, err := format.Source([]byte(finalContent))
		if err != nil {
			log.Printf("Unable to format generated code, writing unformatted code instead. Error: %v", err)
			formattedContent = []byte(finalContent)
		}

		end := time.Now()
		log.Printf("Time elapsed to produce sys signatures: %v\n", end.Sub(start))

		if *output != "" {
			target := *output
			if arch != gowinmd.ArchAll {
				target = strings.TrimSuffix(target, ".go") + "_" + arch.String() + ".go"
			}
			if err := os.WriteFile(target, formattedContent, 0666); err != nil {
				return fmt.Errorf("failed to write output file %s: %w", target, err)
			}
		} else {
			log.Printf("Printing signature results for %s because no output path was specified:\n", arch)
			log.Println("---")
			log.Println(finalContent)
		}
	}
	return nil
}

// parseInputFiles reads Go source files to extract //winmd directives and infer the package name.
// Each //winmd directive selects either a function or a fully qualified WinMD type.
func parseInputFiles(files []string) (methodFilter, typeFilter, string, error) {
	methods := make(methodFilter)
	types := make(typeFilter)
	var pkg string
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return nil, nil, "", err
		}

		s := bufio.NewScanner(file)
		line := 0
		for s.Scan() {
			line++
			t := strings.TrimSpace(s.Text())
			if strings.HasPrefix(t, "//winmd:type") {
				body := t[len("//winmd:type"):]
				if len(body) == 0 || (body[0] != ' ' && body[0] != '\t') {
					file.Close()
					return nil, nil, "", fmt.Errorf("%s:%d: malformed //winmd:type directive: expected //winmd:type <namespace>.<type> [-name <GoName>]", path, line)
				}
				selection, err := parseTypeDirective(strings.TrimSpace(body))
				if err != nil {
					file.Close()
					return nil, nil, "", fmt.Errorf("%s:%d: %w", path, line, err)
				}
				qualifiedName := selection.Namespace + "." + selection.Name
				if existing, ok := types[qualifiedName]; ok {
					if existing.GoName != "" && selection.GoName != "" && existing.GoName != selection.GoName {
						file.Close()
						return nil, nil, "", fmt.Errorf("%s:%d: conflicting Go names %q and %q for WinMD type %s", path, line, existing.GoName, selection.GoName, qualifiedName)
					}
					if existing.GoName == "" && selection.GoName != "" {
						types[qualifiedName] = selection
					}
				} else {
					types[qualifiedName] = selection
				}
				continue
			}
			if !strings.HasPrefix(t, "//winmd:func") {
				continue
			}
			t = t[len("//winmd:func"):]
			if len(t) == 0 || (t[0] != ' ' && t[0] != '\t') {
				continue
			}
			t = strings.TrimSpace(t)
			if t == "" {
				continue
			}
			// Parse: module.method [-name GoName]
			// If only a method name is given (no dot), default to kernel32.
			ref, goName := parseDirective(t)
			if !strings.Contains(ref, ".") {
				// Method on default module, e.g. "CreateFileW" -> "kernel32.CreateFileW".
				ref = "kernel32." + ref
			}
			methods[strings.ToLower(ref)] = goName
		}
		if err := s.Err(); err != nil {
			file.Close()
			return nil, nil, "", err
		}

		// Infer package name from the Go source file.
		if pkg == "" {
			if _, err := file.Seek(0, 0); err != nil {
				file.Close()
				return nil, nil, "", err
			}
			fset := token.NewFileSet()
			parsed, err := parser.ParseFile(fset, path, file, parser.PackageClauseOnly)
			if err != nil {
				file.Close()
				return nil, nil, "", fmt.Errorf("failed to parse package name from %s: %w", path, err)
			}
			pkg = parsed.Name.Name
		}

		file.Close()
	}
	if pkg == "" {
		return nil, nil, "", errors.New("could not determine package name from input files")
	}
	if len(methods) == 0 && len(types) == 0 {
		return nil, nil, "", errors.New("no //winmd:func or //winmd:type directives found in input files")
	}
	return methods, types, pkg, nil
}

type typeSelection struct {
	Namespace string
	Name      string
	GoName    string
}

type typeFilter map[string]typeSelection

func parseTypeDirective(s string) (typeSelection, error) {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return typeSelection{}, errors.New("malformed //winmd:type directive: missing fully qualified type name")
	}

	qualifiedName := fields[0]
	separator := strings.LastIndexByte(qualifiedName, '.')
	if separator <= 0 || separator == len(qualifiedName)-1 || strings.Contains(qualifiedName, "..") {
		return typeSelection{}, fmt.Errorf("WinMD type %q is not fully qualified; expected <namespace>.<type>", qualifiedName)
	}
	selection := typeSelection{
		Namespace: qualifiedName[:separator],
		Name:      qualifiedName[separator+1:],
	}

	switch len(fields) {
	case 1:
		return selection, nil
	case 3:
		if fields[1] != "-name" {
			return typeSelection{}, fmt.Errorf("malformed //winmd:type directive: unexpected option %q", fields[1])
		}
		if !token.IsIdentifier(fields[2]) || fields[2] == "_" {
			return typeSelection{}, fmt.Errorf("invalid Go type name %q: expected a non-blank Go identifier", fields[2])
		}
		selection.GoName = fields[2]
		return selection, nil
	default:
		return typeSelection{}, errors.New("malformed //winmd:type directive: expected <namespace>.<type> [-name <GoName>]")
	}
}

// parseDirective parses a //winmd directive body into the ref (module.method or module)
// and an optional Go function name from -name.
func parseDirective(s string) (ref, goName string) {
	fields := strings.Fields(s)
	ref = fields[0]
	for i := 1; i < len(fields)-1; i++ {
		if fields[i] == "-name" {
			goName = fields[i+1]
			break
		}
	}
	return
}

// methodFilter maps "module.method" keys to an optional Go function name override.
// An empty string value means use the default name. A nil filter means all methods are included.
type methodFilter map[string]string

func writePrototypes(b map[gowinmd.Arch]*strings.Builder, f *winmd.Metadata, filter methodFilter) error {
	return writePrototypesWithProjection(b, f, filter, gowinmd.ProjectionRaw)
}

func writePrototypesWithProjection(b map[gowinmd.Arch]*strings.Builder, f *winmd.Metadata, filter methodFilter, projection gowinmd.Projection) error {
	return writeSelectionsWithProjection(b, f, filter, nil, projection)
}

func writeSelectionsWithProjection(b map[gowinmd.Arch]*strings.Builder, f *winmd.Metadata, filter methodFilter, selectedTypes typeFilter, projection gowinmd.Projection) error {
	context, err := gowinmd.NewContext(f)
	if err != nil {
		return err
	}
	qualifiedNames := make([]string, 0, len(selectedTypes))
	for qualifiedName := range selectedTypes {
		qualifiedNames = append(qualifiedNames, qualifiedName)
	}
	sort.Strings(qualifiedNames)
	for _, qualifiedName := range qualifiedNames {
		selection := selectedTypes[qualifiedName]
		if err := context.SelectTypeDef(selection.Namespace, selection.Name, selection.GoName); err != nil {
			return err
		}
	}

	for idx := range f.Tables.TypeDef.Indices() {
		r, err := f.Tables.TypeDef.At(idx)
		if err != nil {
			return err
		}

		archSeen := make(map[gowinmd.Arch]bool)
		for j := range r.MethodList.All() {
			md, err := f.Tables.MethodDef.At(j)
			if err != nil {
				return err
			}

			var override string
			if filter != nil {
				// Build the "module.method" key for this method using the ImplMap/ModuleRef lookup.
				moduleName, key := methodFilterKey(context, j, md)
				// Match if the exact "module.method" is listed, or "module.*" for all methods in a module.
				goName, exactMatch := filter[key]
				_, moduleMatch := filter[moduleName+".*"]
				if !exactMatch && !moduleMatch {
					continue
				}
				// Pass rename to WriteMethod if specified on the exact match.
				override = goName
			} else {
				// When no filter is applied, skip methods without an ImplMap entry
				// (e.g. .ctor) since they aren't P/Invoke methods.
				if context.MethodModuleName(j) == "" {
					continue
				}
			}

			supportedArches := context.MethodDefSupportedArch(j)
			for _, arch := range supportedArches.Unique() {
				w := b[arch]

				// Write a comment describing this chunk of methods.
				if !archSeen[arch] {
					archSeen[arch] = true
					w.WriteString("\n\n// APIs for ")
					w.WriteString(r.Namespace.String())
				}
				w.WriteString("\n")

				options := gowinmd.MethodOptions{GoName: override, Projection: projection}
				if err := context.WriteMethodWithOptions(w, j, md, arch, options); err != nil {
					// Include context in the error for diag purposes.
					// writeSys may have partially written into b. This is actually convenient for diag.
					lines := strings.Split(w.String(), "\n")
					if len(lines) > 5 {
						lines = lines[len(lines)-5:]
					}

					return fmt.Errorf(
						"error context: \n---\n%v\n---\nfailed to write sys line for %v.Apis method %v: %v",
						strings.Join(lines, "\n"), r.Namespace, md.Name, err)
				}
			}
		}
	}
	if err := context.WriteUsedTypeDefs(b); err != nil {
		return err
	}

	// Log type refs that were used during generation but couldn't be resolved
	// to a TypeDef in the winmd file (e.g. System::Guid).
	for _, ref := range context.UnresolvableTypeRefs() {
		log.Printf("unable to resolve type: %v", ref)
	}

	return nil
}

// methodFilterKey returns the lowercase module name and "module.method" key for a MethodDef,
// matching the format used in //winmd directives.
// If the method has no ImplMap entry (no module), both return values are empty.
func methodFilterKey(ctx *gowinmd.Context, methodIndex winmd.Index, method winmd.MethodDef) (string, string) {
	moduleName := ctx.MethodModuleName(methodIndex)
	if moduleName == "" {
		return "", ""
	}
	return moduleName, strings.ToLower(moduleName + "." + method.Name.String())
}

// generateFileContent wraps the raw syscall content in a valid Go file,
// auto-detecting which imports and type aliases are needed.
func generateFileContent(content string, pkg string) string {
	var b strings.Builder

	b.WriteString("// Code generated by gowinmd; DO NOT EDIT.\n\n")
	b.WriteString("package ")
	b.WriteString(pkg)
	b.WriteString("\n")

	needsUnsafe := strings.Contains(content, "unsafe.Pointer")

	if needsUnsafe {
		b.WriteString("\nimport \"unsafe\"\n")
		b.WriteString("\nvar _ unsafe.Pointer\n")
	}

	b.WriteString(content)
	return b.String()
}
