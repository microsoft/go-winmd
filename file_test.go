// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"debug/pe"
	"testing"

	"github.com/microsoft/go-winmd"
)

func TestNew(t *testing.T) {
	pefile, err := pe.Open("./testdata/Windows.Win32.winmd")
	if err != nil {
		t.Fatal(err)
	}
	defer pefile.Close()
	f, err := winmd.New(pefile)
	if err != nil {
		t.Fatal(err)
	}
	wantVersion := "v4.0.30319"
	if f.Version != wantVersion {
		t.Errorf("Version = %v, want %v", f.Version, wantVersion)

	}
	testHeap := func(data []byte, size int) {
		t.Helper()
		if len(data) == 0 {
			t.Error("heap missing")
			return
		}
		if len(data) != size {
			t.Errorf("Size = %v, want %v", len(data), size)
		}
	}

	testHeap(f.Strings, 6081516)
	testHeap(f.US, 4)
	testHeap(f.GUID, 16)
	testHeap(f.Blob, 1169700)
}
