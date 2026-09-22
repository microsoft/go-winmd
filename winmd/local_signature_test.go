// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestLocalVarsSignature(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	i4 := winmd.SigType{Kind: winmd.ElementType_I4}
	byref := winmd.SigType{Kind: winmd.ElementType_BYREF, Value: i4}
	pinned := []winmd.SigLocalVarMod{{Constraint: winmd.SigConstraint{Pinned: true}}}
	typeParam := winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(128)}
	methodParam := winmd.SigType{Kind: winmd.ElementType_MVAR, Value: uint32(0x1fffffff)}
	fnptr := winmd.SigStandAloneMethod{CallingConvention: winmd.SigCallingConvention_Cdecl,
		SigMethodRef: winmd.SigMethodRef{SigMethodDef: winmd.SigMethodDef{
			RetType: winmd.SigRetType{Type: typeParam}, Param: []winmd.SigParam{{Type: methodParam}},
		}},
	}
	for _, test := range []struct {
		name string
		data []byte
		want winmd.SigLocalVar
	}{
		{"primitive", []byte{8}, winmd.SigLocalVar{Type: i4}},
		{"byref", []byte{0x10, 8}, winmd.SigLocalVar{Kind: winmd.SigLocalVarKind_ByRef, Type: byref}},
		{"typedref", []byte{0x16}, winmd.SigLocalVar{Kind: winmd.SigLocalVarKind_TypedByRef, Type: winmd.SigType{Kind: winmd.ElementType_TYPEDBYREF}}},
		{"pinned-string", []byte{0x45, 0x0e}, winmd.SigLocalVar{Mod: pinned, Type: winmd.SigType{Kind: winmd.ElementType_STRING}}},
		{"pinned-object", []byte{0x45, 0x1c}, winmd.SigLocalVar{Mod: pinned, Type: winmd.SigType{Kind: winmd.ElementType_OBJECT}}},
		{"pinned-byref", []byte{0x45, 0x10, 8}, winmd.SigLocalVar{Kind: winmd.SigLocalVarKind_ByRef, Mod: pinned, Type: byref}},
		{"pinned-value", []byte{0x45, 8}, winmd.SigLocalVar{Mod: pinned, Type: i4}}, // No pinning eligibility checks.
		{"var", []byte{0x13, 0x80, 0x80}, winmd.SigLocalVar{Type: typeParam}},
		{"mvar", []byte{0x1e, 0xdf, 0xff, 0xff, 0xff}, winmd.SigLocalVar{Type: methodParam}},
		{"pinned-var", []byte{0x45, 0x13, 0x80, 0x80}, winmd.SigLocalVar{Mod: pinned, Type: typeParam}},
		{"generic", []byte{0x15, 0x12, 5, 1, 0x13, 0x80, 0x80}, winmd.SigLocalVar{Type: genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeRef, 0, typeParam)}},
		{"fnptr", []byte{0x1b, 1, 1, 0x13, 0x80, 0x80, 0x1e, 0xdf, 0xff, 0xff, 0xff}, winmd.SigLocalVar{Type: winmd.SigType{Kind: winmd.ElementType_FNPTR, Value: fnptr}}},
	} {
		t.Run(test.name, func(t *testing.T) {
			data := append([]byte{7, 1}, test.data...)
			before := bytes.Clone(data)
			got, err := m.LocalVarsSignature(data)
			if err != nil || !reflect.DeepEqual(got, winmd.SigLocalVars{test.want}) {
				t.Fatalf("signature = %#v, %v; want %#v", got, err, test.want)
			}
			if !bytes.Equal(data, before) {
				t.Fatal("decoder changed its input")
			}
			for _, suffix := range []byte{0, 8, 0xff} {
				if _, err := m.LocalVarsSignature(append(bytes.Clone(data), suffix)); err == nil {
					t.Fatalf("trailing byte %x accepted", suffix)
				}
			}
		})
	}
}

func TestLocalVarsSignatureModifiers(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	mods := []winmd.SigCustomMod{
		{Kind: winmd.SigCustomModKind_Opt, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeDef, Index: 0}},
		{Kind: winmd.SigCustomModKind_Reqd, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 1}},
		{Kind: winmd.SigCustomModKind_Opt, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeSpec, Index: 2}},
	}
	pinned := winmd.SigLocalVarMod{Constraint: winmd.SigConstraint{Pinned: true}}
	prefix := []byte{0x20, 4, 0x45, 0x1f, 9, 0x45, 0x20, 14}
	wantMod := []winmd.SigLocalVarMod{{Mod: &mods[0]}, pinned, {Mod: &mods[1]}, pinned, {Mod: &mods[2]}}
	data := []byte{7, 3}
	var want winmd.SigLocalVars
	for _, kind := range []winmd.ElementType{winmd.ElementType_PTR, winmd.ElementType_SZARRAY, winmd.ElementType_BYREF} {
		data = append(data, prefix...)
		data = append(data, byte(kind))
		inner := winmd.SigType{Kind: winmd.ElementType_I4}
		local := winmd.SigLocalVar{Mod: wantMod}
		if kind == winmd.ElementType_BYREF {
			local.Kind = winmd.SigLocalVarKind_ByRef
		} else {
			data = append(data, 0x1f, 9)
			inner.Mod = []winmd.SigCustomMod{mods[1]}
		}
		data = append(data, 8)
		local.Type = winmd.SigType{Kind: kind, Value: inner}
		want = append(want, local)
	}
	got, err := m.LocalVarsSignature(data)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("modified locals = %#v, %v; want %#v", got, err, want)
	}
	// Mutating any modifier must not change another local's modifier.
	// Distinct values above also detect aliasing within each local's prefix.
	for i, local := range got {
		for _, prefix := range local.Mod {
			if prefix.Mod != nil {
				prefix.Mod.Index.Index++
				if !reflect.DeepEqual(local.Type, want[i].Type) || !reflect.DeepEqual(got[i+1:], want[i+1:]) {
					t.Fatal("modifier mutation changed a nested type or another local")
				}
				prefix.Mod.Index.Index--
			}
		}
	}
}

func TestLocalVarsSignatureMalformed(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for header := range 256 {
		if _, err := m.LocalVarsSignature([]byte{byte(header), 1, 8}); (err == nil) != (header == 7) {
			t.Fatalf("header %02x: %v; only exact LOCAL_SIG is valid", header, err)
		}
	}
	for _, data := range [][]byte{
		{7, 0}, {7, 0xff}, {7, 0xdf, 0xff, 0xff, 0xff, 8},
		append([]byte{7, 0xc0, 0, 0xff, 0xff}, bytes.Repeat([]byte{8}, 65535)...),
		{7, 1, 1}, {7, 1, 0x10, 1}, {7, 1, 0x10, 0x10, 8}, {7, 1, 0x10, 0x16},
		{7, 1, 0x20, 5, 0x16}, {7, 1, 0x1f, 5, 0x16}, {7, 1, 0x45, 0x16},
		{7, 1, 0x45, 0x20, 5, 0x16}, {7, 1, 0x10, 0x45, 8}, {7, 1, 0x10, 0x20, 5, 8},
		{7, 1, 0x0f, 0x45, 8}, {7, 1, 0x1d, 0x45, 8}, {7, 1, 0x41, 8},
		{7, 1, 0x12, 0}, {7, 1, 0x12, 1}, {7, 1, 0x12, 7}, {7, 1, 0x20, 0, 8},
		{7, 1, 0x1f, 2, 8}, {7, 1, 0x20, 7, 8}, {7, 1, 0x81, 8},
	} {
		if _, err := m.LocalVarsSignature(data); err == nil {
			t.Fatalf("malformed signature accepted: %x", data[:min(len(data), 24)])
		}
	}
}

func TestLocalVarsSignatureCountsAndTruncation(t *testing.T) {
	var m winmd.Metadata
	for _, test := range []struct {
		count int
		data  []byte
	}{
		{1, []byte{7, 1, 8}},
		{128, append([]byte{7, 0x80, 0x80}, bytes.Repeat([]byte{8}, 128)...)},
		{16384, append([]byte{7, 0xc0, 0, 0x40, 0}, bytes.Repeat([]byte{8}, 16384)...)},
		{65534, append([]byte{7, 0xc0, 0, 0xff, 0xfe}, bytes.Repeat([]byte{8}, 65534)...)},
		{1, []byte{7, 1, 0x20, 0xc0, 0, 0x40, 1, 0x45, 0x10, 0x1d, 0x1f, 0x80, 0x81, 0x12, 5}},
		{1, []byte{7, 1, 0x13, 0xc0, 0, 0x40, 0}},
		{1, []byte{7, 1, 0x1b, 0, 1, 1, 0x1e, 0x80, 0x80}},
	} {
		got, err := m.LocalVarsSignature(test.data)
		if err != nil || len(got) != test.count {
			t.Fatalf("complete signature: %d locals, %v; want %d", len(got), err, test.count)
		}
		for length := range len(test.data) {
			if _, err := m.LocalVarsSignature(test.data[:length]); !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatalf("count %d, prefix length %d: %v; want unexpected EOF", test.count, length, err)
			}
		}
	}
	// An otherwise legal count cannot allocate locals before checking available bytes.
	short := winmd.SigLocalVarsBlob{7, 0xc0, 0, 0xff, 0xfe, 8}
	if allocs := testing.AllocsPerRun(1, func() { _, _ = m.LocalVarsSignature(short) }); allocs != 0 {
		t.Fatalf("impossible local count allocated %g times", allocs)
	}
}

func TestLocalVarsSignatureNesting(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, shape := range []string{"pointer", "vector", "byref", "fnptr-return", "fnptr-parameter", "mixed"} {
		for _, depth := range []int{2, 64, 65, 4096} { // Positive controls precede excessive depths.
			t.Run(fmt.Sprintf("%s/%d", shape, depth), func(t *testing.T) {
				var typ []byte
				for i := range depth - 1 {
					switch {
					case shape == "byref" && i == 0:
						typ = append(typ, 0x10)
					case shape == "vector":
						typ = append(typ, 0x1d)
					case shape == "fnptr-return":
						typ = append(typ, 0x1b, 0, 0)
					case shape == "fnptr-parameter" || shape == "mixed" && i%2 == 1:
						typ = append(typ, 0x1b, 0, 1, 1)
					default:
						typ = append(typ, 0x0f)
					}
				}
				typ = append(typ, 8) // The leaf is one of the 64 levels.
				data := append([]byte{7, 2}, typ...)
				got, err := m.LocalVarsSignature(append(data, typ...)) // Each local gets its own budget.
				if (err == nil) != (depth <= 64) || err == nil && len(got) != 2 {
					t.Fatalf("depth %d: %d locals, %v", depth, len(got), err)
				}
			})
		}
	}
	data := append([]byte{7, 1}, bytes.Repeat([]byte{0x45, 0x20, 5}, 128)...)
	data = append(data, bytes.Repeat([]byte{0x0f}, 63)...)
	got, err := m.LocalVarsSignature(append(data, 8))
	if err != nil || len(got) != 1 || len(got[0].Mod) != 256 {
		t.Fatalf("flat prefixes consumed type depth: %v", err)
	}
}

func TestLocalVarsSignatureFromMetadata(t *testing.T) {
	t.Parallel()
	var standalone winmd.Metadata
	data := winmd.SigLocalVarsBlob{7, 2, 0x20, 4, 0x45, 0x10, 0x12, 5, 0x11, 6}
	want, err := standalone.LocalVarsSignature(data)
	if err != nil {
		t.Fatal(err)
	}
	for _, rows := range []uint32{1, 0} {
		// TypeRef, TypeDef, StandAloneSig, TypeSpec; only the signature row is read.
		tables := metadataTables(0, 1<<1|1<<2|1<<17|1<<27, []uint32{rows, rows, 1, rows}, int(20*rows))
		tables = append(tables, 1, 0) // StandAloneSig.Signature points at blob offset 1.
		tables = append(tables, make([]byte, 2*rows)...)
		blobs := append([]byte{0, byte(len(data))}, data...)
		m, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables}, metadataStream{"#Blob", blobs})))
		if err != nil {
			t.Fatal(err)
		}
		row, err := m.Tables.StandAloneSig.At(0)
		if err != nil {
			t.Fatal(err)
		}
		got, err := m.LocalVarsSignature(winmd.SigLocalVarsBlob(row.Signature))
		if (err == nil) != (rows == 1) || err == nil && !reflect.DeepEqual(got, want) {
			t.Fatalf("table-backed locals with %d target rows: %#v, %v", rows, got, err)
		}
		for _, kind := range []byte{0x11, 0x12, 0x20, 0x1f} {
			for tag := range byte(3) {
				for _, index := range []byte{1, 2} {
					for _, prefix := range [][]byte{nil, {0x0f}, {0x1d}, {0x1b, 0, 0}} {
						blob := append([]byte{7, 1}, prefix...)
						blob = append(blob, kind, index<<2|tag)
						if kind == 0x20 || kind == 0x1f {
							blob = append(blob, 8)
						}
						if _, err := standalone.LocalVarsSignature(blob); err != nil {
							t.Fatalf("valid signature without layout %x: %v", blob, err)
						}
						if _, err := m.LocalVarsSignature(blob); (err == nil) != (uint32(index) <= rows) {
							t.Fatalf("signature %x with %d target rows: %v", blob, rows, err)
						}
					}
				}
			}
		}
	}
}

func ExampleMetadata_LocalVarsSignature() {
	var m winmd.Metadata
	// Two locals: I4 and a pinned reference to STRING.
	sig, err := m.LocalVarsSignature(winmd.SigLocalVarsBlob{7, 2, 8, 0x45, 0x10, 0x0e})
	if err != nil {
		panic(err)
	}
	fmt.Println("locals:", len(sig))
	fmt.Println("first:", sig[0].Type.Kind)
	fmt.Println("second:", sig[1].Type.Kind, "pinned:", sig[1].Mod[0].Constraint.Pinned)
	// Output:
	// locals: 2
	// first: I4
	// second: BYREF pinned: true
}

func FuzzLocalVarsSignature(f *testing.F) {
	for _, data := range [][]byte{
		{7, 1, 8}, {7, 1, 0x16}, {7, 1, 0x45, 0x10, 0x0e},
		{7, 2, 0x20, 4, 0x45, 0x1f, 9, 0x45, 8, 0x1d, 0x20, 6, 0x13, 0},
		{7, 1, 0x1b, 0, 1, 0x13, 0, 0x1e, 1}, {7, 0}, {7, 0xc0, 0, 0xff, 0xfe, 8},
		append(append([]byte{7, 1}, bytes.Repeat([]byte{0x0f, 0x1b, 0, 0}, 64)...), 8),
	} {
		f.Add(data)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		var m winmd.Metadata
		_, _ = m.LocalVarsSignature(data)
	})
}
