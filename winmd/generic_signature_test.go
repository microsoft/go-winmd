// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func genericSignatureType(class bool, tag winmd.TypeDefOrRefOrSpec, index winmd.Index, arguments ...winmd.SigType) winmd.SigType {
	return winmd.SigType{
		Kind: winmd.ElementType_GENERICINST,
		Value: winmd.SigGenericInst{
			Class: class,
			Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: tag, Index: index},
			Type:  arguments,
		},
	}
}

func TestGenericInstantiationSignatures(t *testing.T) {
	t.Parallel()
	i4 := winmd.SigType{Kind: winmd.ElementType_I4}
	class := genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeRef, 0, i4)
	for _, test := range []struct {
		name string
		data []byte
		want winmd.SigType
	}{
		{"class", []byte{0x15, 0x12, 5, 1, 8}, class},
		{"value-type", []byte{0x15, 0x11, 4, 2, 0x0a, 0x0e}, genericSignatureType(false, winmd.TypeDefOrRefOrSpec_TypeDef, 0,
			winmd.SigType{Kind: winmd.ElementType_I8}, winmd.SigType{Kind: winmd.ElementType_STRING})},
		{"type-spec-handle", []byte{0x15, 0x12, 6, 1, 8}, genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeSpec, 0, i4)},
		{"nested", []byte{0x15, 0x12, 5, 2, 0x0e, 0x15, 0x11, 9, 1, 9}, genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeRef, 0,
			winmd.SigType{Kind: winmd.ElementType_STRING}, genericSignatureType(false, winmd.TypeDefOrRefOrSpec_TypeRef, 1, winmd.SigType{Kind: winmd.ElementType_U4}))},
		{"open", []byte{0x15, 0x12, 5, 2, 0x13, 0, 0x1e, 1}, genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeRef, 0,
			winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(0)}, winmd.SigType{Kind: winmd.ElementType_MVAR, Value: uint32(1)})},
		{"array-argument", []byte{0x15, 0x12, 5, 1, 0x14, 8, 1, 1, 3, 1, 0}, genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeRef, 0,
			winmd.SigType{Kind: winmd.ElementType_ARRAY, Value: winmd.SigArray{Type: i4, Rank: 1, Sizes: []uint32{3}, LowerBounds: []int32{0}}})},
		{"array-of-generic", []byte{0x14, 0x15, 0x12, 5, 1, 8, 1, 1, 2, 1, 0},
			winmd.SigType{Kind: winmd.ElementType_ARRAY, Value: winmd.SigArray{Type: class, Rank: 1, Sizes: []uint32{2}, LowerBounds: []int32{0}}}},
		{"szarray-argument", []byte{0x15, 0x12, 5, 1, 0x1d, 8}, genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeRef, 0,
			winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: i4})},
		{"open-szarray-argument", []byte{0x15, 0x12, 5, 1, 0x1d, 0x13, 0}, genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeRef, 0,
			winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(0)}})},
		{"szarray-of-generic", []byte{0x1d, 0x15, 0x12, 5, 1, 8}, winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: class}},
		{"modified-szarray-argument", []byte{0x15, 0x12, 5, 1, 0x1d, 0x20, 9, 8}, genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeRef, 0,
			winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: winmd.SigType{Kind: winmd.ElementType_I4, Mod: []winmd.SigCustomMod{
				{Kind: winmd.SigCustomModKind_Opt, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 1}},
			}}})},
		{"pointer-to-generic", []byte{0x0f, 0x15, 0x12, 5, 1, 8}, winmd.SigType{Kind: winmd.ElementType_PTR, Value: class}},
	} {
		t.Run(test.name, func(t *testing.T) {
			for _, context := range []string{"field", "return", "parameter"} {
				t.Run(context, func(t *testing.T) {
					var m winmd.Metadata
					var got winmd.SigType
					var err error
					switch context {
					case "field":
						var sig winmd.SigField
						sig, err = m.FieldSignature(append([]byte{6}, test.data...))
						got = sig.Type
					case "return":
						var sig winmd.SigMethodDef
						sig, err = m.MethodDefSignature(append([]byte{0, 0}, test.data...))
						got = sig.RetType.Type
					case "parameter":
						var sig winmd.SigMethodDef
						sig, err = m.MethodDefSignature(append([]byte{0, 1, 1}, test.data...))
						if err == nil {
							got = sig.Param[0].Type
						}
					}
					if err != nil || !reflect.DeepEqual(got, test.want) {
						t.Fatalf("decoded type = %#v, %v; want %#v", got, err, test.want)
					}
				})
			}
		})
	}
}

func TestGenericMethodParameterReferences(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	data := []byte{
		0x10, 2, 2, // Generic method with two generic parameters and two ordinary parameters.
		0x1e, 0, // Return !!0.
		0x15, 0x12, 5, 2, 0x13, 0, 0x1e, 1, // Class instantiated with !0 and !!1.
		0x10, 0x15, 0x11, 9, 1, 0x1e, 0, // Byref to a value type instantiated with !!0.
	}
	got, err := m.MethodDefSignature(data)
	mvar0 := winmd.SigType{Kind: winmd.ElementType_MVAR, Value: uint32(0)}
	want := winmd.SigMethodDef{
		Generic: 2,
		RetType: winmd.SigRetType{Kind: winmd.SigRetTypeKind_ByValue, Type: mvar0},
		Param: []winmd.SigParam{
			{Kind: winmd.SigParamKind_ByValue, Type: genericSignatureType(true, winmd.TypeDefOrRefOrSpec_TypeRef, 0,
				winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(0)}, winmd.SigType{Kind: winmd.ElementType_MVAR, Value: uint32(1)})},
			{Kind: winmd.SigParamKind_ByRef, Type: winmd.SigType{Kind: winmd.ElementType_BYREF,
				Value: genericSignatureType(false, winmd.TypeDefOrRefOrSpec_TypeRef, 1, mvar0)}},
		},
	}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("generic method = %#v, %v; want %#v", got, err, want)
	}
}

func TestGenericParameterNumberEncoding(t *testing.T) {
	t.Parallel()
	for _, kind := range []winmd.ElementType{winmd.ElementType_VAR, winmd.ElementType_MVAR} {
		for _, test := range []struct {
			number uint32
			data   []byte
		}{
			{0, []byte{0}},
			{127, []byte{0x7f}},
			{128, []byte{0x80, 0x80}},
			{16383, []byte{0xbf, 0xff}},
			{16384, []byte{0xc0, 0, 0x40, 0}},
			{0x1fffffff, []byte{0xdf, 0xff, 0xff, 0xff}},
		} {
			t.Run(fmt.Sprintf("%s/%d", kind, test.number), func(t *testing.T) {
				var m winmd.Metadata
				sig, err := m.FieldSignature(append([]byte{6, byte(kind)}, test.data...))
				if err != nil || sig.Type.Kind != kind || sig.Type.Value != test.number {
					t.Fatalf("parameter type = %+v, %v; want %s ordinal %d", sig.Type, err, kind, test.number)
				}
			})
		}
	}
}

func TestGenericInstantiationModifiersAndCount(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	data := []byte{6, 0x20, 5, 0x1f, 9, 0x15, 0x11, 4, 0x80, 0x80}
	data = append(data, bytes.Repeat([]byte{8}, 128)...)
	before := bytes.Clone(data)
	sig, err := m.FieldSignature(data)
	if err != nil {
		t.Fatal(err)
	}
	inst, ok := sig.Type.Value.(winmd.SigGenericInst)
	if !ok || inst.Class || inst.Index.Tag != winmd.TypeDefOrRefOrSpec_TypeDef || inst.Index.Index != 0 || len(inst.Type) != 128 {
		t.Fatalf("generic instantiation = %+v; want value type with 128 arguments", sig.Type.Value)
	}
	for _, arg := range inst.Type {
		if arg.Kind != winmd.ElementType_I4 {
			t.Fatalf("argument = %+v; want I4", arg)
		}
	}
	if len(sig.Type.Mod) != 2 || sig.Type.Mod[0].Kind != winmd.SigCustomModKind_Opt || sig.Type.Mod[1].Kind != winmd.SigCustomModKind_Reqd {
		t.Fatalf("generic field modifiers reordered: %+v", sig.Type.Mod)
	}
	if !bytes.Equal(data, before) {
		t.Fatal("signature decoder changed its input")
	}
}

func TestGenericSignaturesTruncated(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, data := range [][]byte{
		{6, 0x15, 0x12, 5, 1, 8},
		{6, 0x15, 0x11, 0x80, 0x80, 1, 0x13, 0x80, 0x80},
		{6, 0x15, 0x12, 5, 1, 0x15, 0x11, 9, 1, 0x1e, 0},
		{6, 0x15, 0x12, 5, 1, 0x1d, 0x13, 0},
		{6, 0x1d, 0x15, 0x12, 5, 1, 0x1e, 0},
		{6, 0x13, 0xc0, 0, 0x40, 0},
		{6, 0x1e, 0xc0, 0, 0x40, 0},
	} {
		t.Run(fmt.Sprintf("%x", data), func(t *testing.T) {
			if _, err := m.FieldSignature(data); err != nil {
				t.Fatalf("complete signature rejected: %v", err)
			}
			for length := range len(data) {
				if _, err := m.FieldSignature(data[:length]); !errors.Is(err, io.ErrUnexpectedEOF) {
					t.Fatalf("prefix length %d: error = %v; want unexpected EOF", length, err)
				}
			}
		})
	}
}

func TestGenericSignaturesMalformed(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, test := range []struct {
		name string
		data []byte
	}{
		{"invalid-kind", []byte{0x15, 8, 5, 1, 8}},
		{"modifier-before-kind", []byte{0x15, 0x20, 5, 0x12, 5, 1, 8}},
		{"wide-invalid-kind", []byte{0x15, 0x81, 0x12, 5, 1, 8}},
		{"null-handle", []byte{0x15, 0x12, 0, 1, 8}},
		{"zero-row-handle", []byte{0x15, 0x12, 1, 1, 8}},
		{"reserved-handle", []byte{0x15, 0x12, 7, 1, 8}},
		{"zero-arguments", []byte{0x15, 0x12, 5, 0}},
		{"void-argument", []byte{0x15, 0x12, 5, 1, 1}},
		{"byref-argument", []byte{0x15, 0x12, 5, 1, 0x10, 8}},
		{"typedref-argument", []byte{0x15, 0x12, 5, 1, 0x16}},
		{"pointer-argument", []byte{0x15, 0x12, 5, 1, 0x0f, 8}},
		{"void-pointer-argument", []byte{0x15, 0x12, 5, 1, 0x0f, 1}},
		{"modifier-on-argument", []byte{0x15, 0x12, 5, 1, 0x20, 9, 8}},
		{"invalid-var-number", []byte{0x13, 0xff}},
		{"invalid-mvar-number", []byte{0x1e, 0xff}},
		{"invalid-argument-count", []byte{0x15, 0x12, 5, 0xff}},
		{"trailing-argument", []byte{0x15, 0x12, 5, 1, 8, 8}},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, err := m.FieldSignature(append([]byte{6}, test.data...)); err == nil {
				t.Fatalf("malformed field type %x accepted", test.data)
			}
			if _, err := m.MethodDefSignature(append([]byte{0, 0}, test.data...)); err == nil {
				t.Fatalf("malformed return type %x accepted", test.data)
			}
		})
	}
}

func TestGenericArgumentCountsBounded(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, count := range [][]byte{{0x80, 0x80}, {0xc0, 1, 0, 0}, {0xdf, 0xff, 0xff, 0xff}} {
		t.Run(fmt.Sprintf("%x", count), func(t *testing.T) {
			data := append([]byte{6, 0x15, 0x12, 5}, count...)
			sig, err := m.FieldSignature(append(data, 8))
			if !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatalf("oversized count: %v; want unexpected EOF", err)
			}
			if inst, ok := sig.Type.Value.(winmd.SigGenericInst); ok && cap(inst.Type) != 0 {
				t.Fatalf("oversized count allocated %d arguments", cap(inst.Type))
			}
		})
	}
	t.Run("invalid-first-argument", func(t *testing.T) {
		data := append([]byte{6, 0x15, 0x12, 5, 0x80, 0x80}, make([]byte, 128)...)
		sig, err := m.FieldSignature(data)
		if err == nil {
			t.Fatal("invalid argument accepted")
		}
		if inst, ok := sig.Type.Value.(winmd.SigGenericInst); ok && cap(inst.Type) != 0 {
			t.Fatal("argument storage allocated before validating the first argument")
		}
	})
}

func TestGenericSignatureHandleBounds(t *testing.T) {
	t.Parallel()
	tables := metadataTables(0, 1<<1|1<<2|1<<27, []uint32{1, 1, 1}, 6+14+2)
	m, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables})))
	if err != nil {
		t.Fatal(err)
	}
	for _, handle := range []byte{4, 5, 6, 8, 9, 10} {
		for _, nested := range []bool{false, true} {
			t.Run(fmt.Sprintf("handle-%d/nested-%v", handle, nested), func(t *testing.T) {
				data := []byte{6, 0x15, 0x12, handle, 1, 8}
				if nested {
					data = []byte{6, 0x15, 0x12, 5, 1, 0x15, 0x11, handle, 1, 8}
				}
				_, err := m.FieldSignature(data)
				if handle < 8 {
					if err != nil {
						t.Fatalf("valid target row rejected: %v", err)
					}
				} else if err == nil || !strings.Contains(err.Error(), "beyond the end of table") {
					t.Fatalf("out-of-range target: %v; want table bounds error", err)
				}
			})
		}
	}
}

func TestGenericSignatureNesting(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, mixed := range []string{"none", "array", "szarray"} {
		for _, layers := range []int{0, 1, 63, 64, 65, 4096} {
			t.Run(fmt.Sprintf("mixed-%s/layers-%d", mixed, layers), func(t *testing.T) {
				data := []byte{6}
				arrays := 0
				for i := range layers {
					if mixed == "array" && i%2 == 0 {
						data = append(data, 0x14)
						arrays++
					} else if mixed == "szarray" && i%2 == 0 {
						data = append(data, 0x1d)
					} else {
						data = append(data, 0x15, 0x12, 5, 1)
					}
				}
				data = append(data, 8)
				for range arrays {
					data = append(data, 1, 0, 0) // Array shapes follow their element types.
				}
				_, err := m.FieldSignature(data)
				if layers >= 64 {
					if err == nil || !strings.Contains(err.Error(), "nesting limit") {
						t.Fatalf("depth %d: %v; want nesting limit error", layers+1, err)
					}
				} else if err != nil {
					t.Fatalf("valid depth %d rejected: %v", layers+1, err)
				}
			})
		}
	}
}

func TestGenericArgumentsHaveIndependentDepthBudgets(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	// The parent, 62 nested instantiations, and I4 occupy exactly 64 levels.
	child := bytes.Repeat([]byte{0x15, 0x12, 5, 1}, 62)
	child = append(child, 8)
	typ := []byte{0x15, 0x12, 5, 2}
	typ = append(typ, child...)
	typ = append(typ, child...)
	if _, err := m.FieldSignature(append([]byte{6}, typ...)); err != nil {
		t.Fatalf("sibling generic arguments share a depth budget: %v", err)
	}
	data := append([]byte{0, 1}, typ...)
	data = append(data, typ...)
	if _, err := m.MethodDefSignature(data); err != nil {
		t.Fatalf("generic return and parameter share a depth budget: %v", err)
	}
}

func FuzzGenericInstantiationSignature(f *testing.F) {
	for _, data := range [][]byte{
		{0x12, 5, 1, 8},
		{0x11, 4, 2, 0x13, 0, 0x1e, 1},
		{0x12, 5, 1, 0x15, 0x11, 9, 1, 8},
		{0x12, 5, 1, 0x1d, 8},
		{0x12, 5, 1, 0x1d, 0x1d, 0x13, 0},
		{0x12, 5, 0xdf, 0xff, 0xff, 0xff},
	} {
		f.Add(data)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		var m winmd.Metadata
		_, _ = m.FieldSignature(append([]byte{6, 0x15}, data...))
		_, _ = m.MethodDefSignature(append([]byte{0, 1, 1, 0x15}, data...))
	})
}
