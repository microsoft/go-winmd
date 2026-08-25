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

func newArchBuilders() map[gowinmd.Arch]*strings.Builder {
	return map[gowinmd.Arch]*strings.Builder{
		gowinmd.Arch386:   {},
		gowinmd.ArchAMD64: {},
		gowinmd.ArchARM64: {},
		gowinmd.ArchAll:   {},
		gowinmd.ArchNone:  {},
	}
}

func TestWriteMethod(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
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
		if w.Len() == 0 {
			continue
		}
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

func TestWriteProjectedBCryptMethods(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
	filter := methodFilter{}
	for _, name := range []string{
		"BCryptGetFipsAlgorithmMode", "BCryptGetProperty", "BCryptSetProperty",
		"BCryptOpenAlgorithmProvider", "BCryptCreateHash", "BCryptHashData",
		"BCryptEncrypt", "BCryptDecrypt", "BCryptGenerateSymmetricKey",
		"BCryptSignHash", "BCryptDeriveKey",
	} {
		filter[strings.ToLower("bcrypt.dll."+name)] = ""
	}
	filter["bcrypt.dll.bcrypthashdata"] = "rawEncrypt"

	if err := writePrototypesWithProjection(b, f, filter, gowinmd.ProjectionIdiomatic); err != nil {
		t.Fatal(err)
	}
	got := b[gowinmd.ArchAll].String()
	wants := []string{
		"//sys\tBCryptGetProperty(hObject BCRYPT_HANDLE, pszProperty *uint16, pbOutput []byte, pcbResult *uint32, dwFlags uint32) (ntstatus error) = bcrypt.BCryptGetProperty",
		"//sys\tBCryptSetProperty(hObject BCRYPT_HANDLE, pszProperty *uint16, pbInput []byte, dwFlags uint32) (ntstatus error) = bcrypt.BCryptSetProperty",
		"//sys\trawEncrypt(hHash BCRYPT_HASH_HANDLE, pbInput []byte, dwFlags uint32) (ntstatus error) = bcrypt.BCryptHashData",
	}
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("projected output does not contain %q\noutput:\n%s", want, got)
		}
	}
	if strings.Contains(got, ".dll.") || strings.Contains(got, "PWSTRElement") {
		t.Errorf("projected output contains an unnormalized DLL or synthetic string element type:\n%s", got)
	}
	if !strings.Contains(got, "BCRYPT_BLOCK_PADDING BCRYPT_FLAGS = 0x1") ||
		!strings.Contains(got, "BCRYPT_PAD_NONE BCRYPT_FLAGS = 0x1") {
		t.Errorf("projected output did not preserve duplicate constant values:\n%s", got)
	}
}

func TestIdiomaticProjectionDoesNotProjectHRESULT(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
	filter := methodFilter{"ncrypt.dll.ncryptgetproperty": ""}
	if err := writePrototypesWithProjection(b, f, filter, gowinmd.ProjectionIdiomatic); err != nil {
		t.Fatal(err)
	}
	got := b[gowinmd.ArchAll].String()
	if !strings.Contains(got, "(r HRESULT) = ncrypt.NCryptGetProperty") {
		t.Fatalf("HRESULT return was unexpectedly projected:\n%s", got)
	}
}

func TestFullFile(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
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
		if !nsSet[r.Namespace.String()] {
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
