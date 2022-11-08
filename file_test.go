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

	testLen(t, f.Tables.Module, 1)
	testLen(t, f.Tables.TypeRef, 14595)
	testLen(t, f.Tables.TypeDef, 33749)
	testLen(t, f.Tables.Field, 225577)
	testLen(t, f.Tables.MethodDef, 66944)
	testLen(t, f.Tables.Param, 206634)
	testLen(t, f.Tables.ClassLayout, 1046)
	testLen(t, f.Tables.Assembly, 1)
	testLen(t, f.Tables.AssemblyRef, 4)
	testLen(t, f.Tables.InterfaceImpl, 7550)
	testLen(t, f.Tables.MemberRef, 24)
	testLen(t, f.Tables.Constant, 144249)
	testLen(t, f.Tables.CustomAttribute, 74268)
	testLen(t, f.Tables.FieldLayout, 3860)
	testLen(t, f.Tables.ModuleRef, 344)
	testLen(t, f.Tables.ImplMap, 17148)
	testLen(t, f.Tables.NestedClass, 1744)
}

func testLen[T any, TP winmd.Record[T]](t *testing.T, table winmd.Table[T, TP], size uint32) {
	t.Helper()
	if table.Len != size {
		t.Errorf("len = %v, want %v", table.Len, size)
	}
}
