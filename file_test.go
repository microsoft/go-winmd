// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"debug/pe"
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
	wantVersion := "v4.0.30319"
	if f.Version != wantVersion {
		t.Errorf("Version = %v, want %v", f.Version, wantVersion)

	}
	testHeap := func(h *winmd.Heap, size uint32) {
		t.Helper()
		if h == nil {
			t.Error("heap missing")
			return
		}
		if h.Size != size {
			t.Errorf("Size = %v, want %v", h.Size, size)
		}
	}
	testHeap(f.Tables, 8836564)
	testHeap(f.Strings, 6081516)
	testHeap(f.US, 4)
	testHeap(f.GUID, 16)
	testHeap(f.Blob, 1169700)
}
