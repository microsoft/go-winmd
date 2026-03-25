// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"go/format"
	"path/filepath"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/cmd/gowinmd/internal/gowinmd"
	"github.com/microsoft/go-winmd/winmd"
)

func TestWriteMethod(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := map[gowinmd.Arch]*strings.Builder{
		gowinmd.Arch386:   {},
		gowinmd.ArchAMD64: {},
		gowinmd.ArchARM64: {},
		gowinmd.ArchAll:   {},
		gowinmd.ArchNone:  {},
	}
	filter, err := buildNamespaceFilter(f,
		"Windows.Win32.Storage.FileSystem",
		"Windows.Win32.Security.Cryptography",
		"Windows.Win32.System.Diagnostics.Debug",
	)
	if err != nil {
		t.Fatal(err)
	}
	if err := writePrototypes(b, f, filter); err != nil {
		t.Fatal(err)
	}
	for arch, w := range b {
		formattedContent, err := format.Source([]byte(w.String()))
		if err != nil {
			t.Fatal(err)
		}
		target := "prototypes.golden"
		if arch != gowinmd.ArchAll {
			target += "_" + arch.String()
		}
		target += ".go"
		Check(t, "go test ./cmd/gowinmd", filepath.Join("testdata", target), string(formattedContent))
	}
}

func TestFullFile(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := map[gowinmd.Arch]*strings.Builder{
		gowinmd.Arch386:   {},
		gowinmd.ArchAMD64: {},
		gowinmd.ArchARM64: {},
		gowinmd.ArchAll:   {},
		gowinmd.ArchNone:  {},
	}
	if err := writePrototypes(b, f, nil); err != nil {
		t.Fatal(err)
	}
	for _, w := range b {
		_, err = format.Source([]byte(w.String()))
		if err != nil {
			t.Fatal(err)
		}
	}

	// The generated source code is ~4 MB, so don't write it to source control as a golden file.
	// This test only checks that the generation process doesn't fail and doesn't take an
	// exceptionally long time.

	// To see the output, use:
	//   go run ./cmd/gowinmd -o all.go.temp -source .\winmd\testdata\Windows.Win32.winmd
}

func openTestWinmd() (*winmd.Metadata, error) {
	return winmd.Open("../../winmd/testdata/Windows.Win32.winmd")
}

// buildNamespaceFilter pre-scans the winmd file to build a methodFilter containing
// all module.method entries for the given namespaces.
func buildNamespaceFilter(f *winmd.Metadata, namespaces ...string) (methodFilter, error) {
	ctx, err := gowinmd.NewContext(f)
	if err != nil {
		return nil, err
	}
	nsSet := make(map[string]bool, len(namespaces))
	for _, ns := range namespaces {
		nsSet[ns] = true
	}
	filter := make(methodFilter)
	for idx := range f.Tables.TypeDef.Indices() {
		r, err := f.Tables.TypeDef.At(idx)
		if err != nil {
			return nil, err
		}
		if !nsSet[r.Namespace.String()] || !strings.Contains(r.Name.String(), "Apis") {
			continue
		}
		for j := range r.MethodList.All() {
			md, err := f.Tables.MethodDef.At(j)
			if err != nil {
				return nil, err
			}
			moduleName := ctx.MethodModuleName(j)
			if moduleName == "" {
				continue
			}
			filter[strings.ToLower(moduleName+"."+md.Name.String())] = ""
		}
	}
	return filter, nil
}
