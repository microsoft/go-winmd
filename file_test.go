// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"debug/pe"
	"reflect"
	"testing"

	"github.com/microsoft/go-winmd"
)

func TestNewFile(t *testing.T) {
	pefile, err := pe.Open("./testdata/Windows.Win32.winmd")
	if err != nil {
		t.Fatal(err)
	}
	defer pefile.Close()
	f, err := winmd.NewFile(pefile)
	if err != nil {
		t.Fatal(err)
	}
	wantCLIHeader := winmd.CLIHeader{
		Size:                72,
		MajorRuntimeVersion: 2,
		MinorRuntimeVersion: 5,
		Metadata:            pe.DataDirectory{VirtualAddress: 8272, Size: 16087908},
		Flags:               1,
		StrongNameSignature: pe.DataDirectory{VirtualAddress: 16096180, Size: 128},
	}
	if !reflect.DeepEqual(f.CLIHeader, wantCLIHeader) {
		t.Errorf("CLIHeader = %v, want %v", f.CLIHeader, wantCLIHeader)
	}
	wantMetadataHeader := winmd.MetadataHeader{
		Signature:    1112167234,
		MajorVersion: 1,
		MinorVersion: 1,
		Version:      "v4.0.30319",
	}
	if !reflect.DeepEqual(f.MetadataHeader, wantMetadataHeader) {
		t.Errorf("Metadata = %v, want %v", f.MetadataHeader, wantMetadataHeader)
	}
	if f.Stream("#~") == nil {
		t.Error("missing stream #~")
	}
	if f.Stream("#Strings") == nil {
		t.Error("missing stream #String")
	}
	if f.Stream("#US") == nil {
		t.Error("missing stream #US")
	}
	if f.Stream("#GUID") == nil {
		t.Error("missing stream #GUID")
	}
	if f.Stream("#Blob") == nil {
		t.Error("missing stream #Blob")
	}
}
