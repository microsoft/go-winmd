// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"debug/pe"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestNew(t *testing.T) {
	t.Parallel()
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

	testHeap(f.Strings, 6660408)
	testHeap(f.US, 4)
	testHeap(f.GUID, 16)
	testHeap(f.Blob, 7219532)

	testLen(t, f.Tables.Module, 1)
	testLen(t, f.Tables.TypeRef, 16505)
	testLen(t, f.Tables.TypeDef, 37284)
	testLen(t, f.Tables.Field, 247559)
	testLen(t, f.Tables.MethodDef, 70599)
	testLen(t, f.Tables.Param, 218990)
	testLen(t, f.Tables.ClassLayout, 1250)
	testLen(t, f.Tables.Assembly, 1)
	testLen(t, f.Tables.AssemblyRef, 5)
	testLen(t, f.Tables.InterfaceImpl, 7945)
	testLen(t, f.Tables.MemberRef, 41)
	testLen(t, f.Tables.Constant, 156721)
	testLen(t, f.Tables.CustomAttribute, 151929)
	testLen(t, f.Tables.FieldLayout, 4521)
	testLen(t, f.Tables.ModuleRef, 377)
	testLen(t, f.Tables.ImplMap, 18315)
	testLen(t, f.Tables.NestedClass, 2165)
}

func testLen[T any](t *testing.T, table winmd.Table[T], size uint32) {
	t.Helper()
	if table.Len() != size {
		t.Errorf("table = %v, len = %v, want %v", table.Name(), table.Len(), size)
	}
}

func TestTable(t *testing.T) {
	t.Parallel()
	f, err := winmd.Open("./testdata/Windows.Win32.winmd")
	if err != nil {
		t.Fatal(err)
	}

	testTable(t, f.Tables.Assembly)
	testTable(t, f.Tables.AssemblyRef)
	testTable(t, f.Tables.ClassLayout)
	testTable(t, f.Tables.Constant)
	testTable(t, f.Tables.CustomAttribute)
	testTable(t, f.Tables.DeclSecurity)
	testTable(t, f.Tables.EventMap)
	testTable(t, f.Tables.Event)
	testTable(t, f.Tables.ExportedType)
	testTable(t, f.Tables.Field)
	testTable(t, f.Tables.FieldLayout)
	testTable(t, f.Tables.FieldMarshal)
	testTable(t, f.Tables.FieldRVA)
	testTable(t, f.Tables.File)
	testTable(t, f.Tables.GenericParam)
	testTable(t, f.Tables.GenericParamConstraint)
	testTable(t, f.Tables.ImplMap)
	testTable(t, f.Tables.InterfaceImpl)
	testTable(t, f.Tables.ManifestResource)
	testTable(t, f.Tables.MemberRef)
	testTable(t, f.Tables.MethodDef)
	testTable(t, f.Tables.MethodImpl)
	testTable(t, f.Tables.MethodSemantics)
	testTable(t, f.Tables.MethodSpec)
	testTable(t, f.Tables.Module)
	testTable(t, f.Tables.ModuleRef)
	testTable(t, f.Tables.NestedClass)
	testTable(t, f.Tables.Param)
	testTable(t, f.Tables.Property)
	testTable(t, f.Tables.PropertyMap)
	testTable(t, f.Tables.StandAloneSig)
	testTable(t, f.Tables.TypeDef)
	testTable(t, f.Tables.TypeRef)
	testTable(t, f.Tables.TypeSpec)
}

func testTable[T any](t *testing.T, table winmd.Table[T]) {
	t.Run(table.Name(), func(t *testing.T) {
		t.Parallel()
		for idx := range table.Indices() {
			if _, err := table.At(idx); err != nil {
				t.Errorf("%d: %v", idx, err)
			}
		}
	})
}

func BenchmarkReadAllTableEntries(b *testing.B) {
	b.ReportAllocs()

	metadata, err := winmd.Open("./testdata/Windows.Win32.winmd")
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		benchmarkReadTable(b, metadata.Tables.Assembly)
		benchmarkReadTable(b, metadata.Tables.AssemblyRef)
		benchmarkReadTable(b, metadata.Tables.ClassLayout)
		benchmarkReadTable(b, metadata.Tables.Constant)
		benchmarkReadTable(b, metadata.Tables.CustomAttribute)
		benchmarkReadTable(b, metadata.Tables.DeclSecurity)
		benchmarkReadTable(b, metadata.Tables.EventMap)
		benchmarkReadTable(b, metadata.Tables.Event)
		benchmarkReadTable(b, metadata.Tables.ExportedType)
		benchmarkReadTable(b, metadata.Tables.Field)
		benchmarkReadTable(b, metadata.Tables.FieldLayout)
		benchmarkReadTable(b, metadata.Tables.FieldMarshal)
		benchmarkReadTable(b, metadata.Tables.FieldRVA)
		benchmarkReadTable(b, metadata.Tables.File)
		benchmarkReadTable(b, metadata.Tables.GenericParam)
		benchmarkReadTable(b, metadata.Tables.GenericParamConstraint)
		benchmarkReadTable(b, metadata.Tables.ImplMap)
		benchmarkReadTable(b, metadata.Tables.InterfaceImpl)
		benchmarkReadTable(b, metadata.Tables.ManifestResource)
		benchmarkReadTable(b, metadata.Tables.MemberRef)
		benchmarkReadTable(b, metadata.Tables.MethodDef)
		benchmarkReadTable(b, metadata.Tables.MethodImpl)
		benchmarkReadTable(b, metadata.Tables.MethodSemantics)
		benchmarkReadTable(b, metadata.Tables.MethodSpec)
		benchmarkReadTable(b, metadata.Tables.Module)
		benchmarkReadTable(b, metadata.Tables.ModuleRef)
		benchmarkReadTable(b, metadata.Tables.NestedClass)
		benchmarkReadTable(b, metadata.Tables.Param)
		benchmarkReadTable(b, metadata.Tables.Property)
		benchmarkReadTable(b, metadata.Tables.PropertyMap)
		benchmarkReadTable(b, metadata.Tables.StandAloneSig)
		benchmarkReadTable(b, metadata.Tables.TypeDef)
		benchmarkReadTable(b, metadata.Tables.TypeRef)
		benchmarkReadTable(b, metadata.Tables.TypeSpec)
	}
}

func benchmarkReadTable[T any](b *testing.B, table winmd.Table[T]) {
	b.Helper()
	for idx := range table.Indices() {
		if _, err := table.At(idx); err != nil {
			b.Fatal(err)
		}
	}
}
