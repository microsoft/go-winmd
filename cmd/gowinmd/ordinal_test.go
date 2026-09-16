// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"errors"
	"log"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/cmd/gowinmd/internal/gowinmd"
)

func TestOrdinalImportSelections(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		name    string
		filter  methodFilter
		wantErr bool
	}{
		{"explicit", methodFilter{"shell32.dll.fileiconinit": ""}, true},
		{"wildcard", methodFilter{"shell32.dll.*": ""}, false},
		{"explicit-with-wildcard", methodFilter{"shell32.dll.*": "", "shell32.dll.fileiconinit": ""}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			var diagnostics strings.Builder
			previousOutput := log.Writer()
			log.SetOutput(&diagnostics)
			t.Cleanup(func() { log.SetOutput(previousOutput) })
			builders := newArchBuilders()
			err := writePrototypes(builders, f, test.filter)
			if test.wantErr {
				if !errors.Is(err, gowinmd.ErrOrdinalImport) {
					t.Fatalf("explicit ordinal request error = %v; want ErrOrdinalImport", err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !strings.Contains(diagnostics.String(), "skipping") || !strings.Contains(diagnostics.String(), "FileIconInit") {
					t.Fatalf("missing diagnostic for skipped ordinal: %q", diagnostics.String())
				}
				if !strings.Contains(builders[gowinmd.ArchAll].String(), "//sys\tSHGetKnownFolderPath(") {
					t.Fatal("wildcard selection did not generate supported shell functions")
				}
			}
			for arch, builder := range builders {
				if strings.Contains(builder.String(), "FileIconInit(") {
					t.Fatalf("%s output contains a partial or unsupported ordinal declaration", arch)
				}
			}
		})
	}
}
