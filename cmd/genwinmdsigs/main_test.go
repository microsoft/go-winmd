// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"debug/pe"
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
	Check(t, "go test ./cmd/genwinmdsigs", filepath.Join("testdata", "prototypes.golden.go"), b.String())
}

func openTestWinmd() (*winmd.Metadata, error) {
	pefile, err := pe.Open("../../testdata/Windows.Win32.winmd")
	if err != nil {
		return nil, err
	}
	defer pefile.Close()
	return winmd.New(pefile)
}
