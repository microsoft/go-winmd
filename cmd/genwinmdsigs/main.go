// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"debug/pe"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/microsoft/go-winmd"
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

	formattedContent, err := format.Source([]byte(finalContent))
	if err != nil {
		log.Printf("Unable to format generated code, writing unformatted code instead. Error: %v", err)
		formattedContent = []byte(finalContent)
	}

	end := time.Now()
	log.Printf("Time elapsed to produce sys signatures: %v\n", end.Sub(start))

	if *output != "" {
		return os.WriteFile(*output, formattedContent, 0666)
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

	for i := uint32(0); i < f.Tables.TypeDef.Len; i++ {
		r, err := f.Tables.TypeDef.Record(winmd.Index(i))
		if err != nil {
			return err
		}

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

			// Write a comment describing this chunk of methods.
			if firstMethod {
				firstMethod = false
				b.WriteString("\n\n// APIs for ")
				b.WriteString(r.Namespace.String())
			}
			b.WriteString("\n")

			if err := context.WriteMethod(b, j, md); err != nil {
				// Include context in the error for diag purposes.
				// writeSys may have partially written into b. This is actually convenient for diag.
				lines := strings.Split(b.String(), "\n")
				if len(lines) > 5 {
					lines = lines[len(lines)-5:]
				}

				return fmt.Errorf(
					"error context: \n---\n%v\n---\nfailed to write sys line for %v.Apis method %v: %v",
					strings.Join(lines, "\n"), r.Namespace, md.Name, err)
			}
		}
	}
	if err := context.WriteUsedTypeDefs(b); err != nil {
		return err
	}
	return nil
}
