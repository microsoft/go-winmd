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

	flag.Parse()

	if *source == "" {
		return errors.New("source is required: pass the path to a win32metadata file using the -source flag")
	}
	if *formatFlag != "mkwinsyscall" {
		return fmt.Errorf("format is required: pass -format mkwinsyscall (got %q)", *formatFlag)
	}

	inputFiles := flag.Args()
	if len(inputFiles) == 0 {
		return errors.New("no input files provided: pass Go files containing //winmd directives as non-flag arguments")
	}

	// Parse //winmd directives and infer package name from the input files.
	filter, pkg, err := parseInputFiles(inputFiles)
	if err != nil {
		return err
	}

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

	if err := writePrototypes(b, f, filter); err != nil {
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
			os.WriteFile(target, formattedContent, 0666)
		} else {
			log.Printf("Printing signature results for %s because no output path was specified:\n", arch)
			log.Println("---")
			log.Println(finalContent)
		}
	}
	return nil
}

// parseInputFiles reads Go source files to extract //winmd directives and infer the package name.
// Each //winmd directive should contain a module.method reference (e.g. "kernel32.CreateFileW"),
// optionally followed by -name GoFuncName to rename the generated function.
// Returns a map of "module.method" -> Go name (empty string if no rename) and the inferred package name.
func parseInputFiles(files []string) (methodFilter, string, error) {
	filter := make(methodFilter)
	var pkg string
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return nil, "", err
		}

		s := bufio.NewScanner(file)
		for s.Scan() {
			t := strings.TrimSpace(s.Text())
			if !strings.HasPrefix(t, "//winmd") {
				continue
			}
			t = t[len("//winmd"):]
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
				ref = "kernel32." + ref
			}
			filter[strings.ToLower(ref)] = goName
		}
		if err := s.Err(); err != nil {
			file.Close()
			return nil, "", err
		}

		// Infer package name from the Go source file.
		if pkg == "" {
			if _, err := file.Seek(0, 0); err != nil {
				file.Close()
				return nil, "", err
			}
			fset := token.NewFileSet()
			parsed, err := parser.ParseFile(fset, path, file, parser.PackageClauseOnly)
			if err != nil {
				file.Close()
				return nil, "", fmt.Errorf("failed to parse package name from %s: %w", path, err)
			}
			pkg = parsed.Name.Name
		}

		file.Close()
	}
	if pkg == "" {
		return nil, "", errors.New("could not determine package name from input files")
	}
	if len(filter) == 0 {
		return nil, "", errors.New("no //winmd directives found in input files")
	}
	return filter, pkg, nil
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
	context, err := gowinmd.NewContext(f)
	if err != nil {
		return err
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

				if err := context.WriteMethod(w, j, md, arch, override); err != nil {
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
