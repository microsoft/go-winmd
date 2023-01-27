// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"debug/pe"
	"go/format"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd"
)

func TestWriteMethod(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	var b strings.Builder
	r := regexp.MustCompile(`^(Windows\.Win32\.Storage\.FileSystem|Windows\.Win32\.Security\.Cryptography)::`)
	if err := writePrototypes(&b, f, r); err != nil {
		t.Fatal(err)
	}
	formattedContent, err := format.Source([]byte(b.String()))
	if err != nil {
		t.Fatal(err)
	}
	Check(t, "go test ./cmd/genwinmdsigs", filepath.Join("testdata", "prototypes.golden.go"), string(formattedContent))
}

func TestFullFile(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	var b strings.Builder
	if err := writePrototypes(&b, f, nil); err != nil {
		t.Fatal(err)
	}
	_, err = format.Source([]byte(b.String()))
	if err != nil {
		t.Fatal(err)
	}
	// The generated source code is ~4 MB, so don't write it to source control as a golden file.
	// This test only checks that the generation process doesn't fail and doesn't take an
	// exceptionally long time.

	// To see the output, use:
	//   go run ./cmd/genwinmdsigs -o all.go.temp -source .\testdata\Windows.Win32.winmd
}

func openTestWinmd() (*winmd.Metadata, error) {
	pefile, err := pe.Open("../../testdata/Windows.Win32.winmd")
	if err != nil {
		return nil, err
	}
	defer pefile.Close()
	return winmd.New(pefile)
}
