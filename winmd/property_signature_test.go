// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestPropertySignature(t *testing.T) {
	t.Parallel()
	i4 := winmd.SigType{Kind: winmd.ElementType_I4}
	str := winmd.SigType{Kind: winmd.ElementType_STRING}
	handle := winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 0}
	vector := winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(0)}}
	for _, test := range []struct {
		name  string
		data  []byte // ParamCount, property type, then index parameters.
		typ   winmd.SigType
		param []winmd.SigParam
	}{
		{"scalar", []byte{0, 8}, i4, nil},
		{"indexer", []byte{2, 0x0e, 8, 0x0e}, str, []winmd.SigParam{{Type: i4}, {Type: str}}},
		{"class", []byte{0, 0x12, 5}, winmd.SigType{Kind: winmd.ElementType_CLASS, Value: handle}, nil},
		{"byref-index", []byte{1, 8, 0x10, 8}, i4, []winmd.SigParam{
			{Kind: winmd.SigParamKind_ByRef, Type: winmd.SigType{Kind: winmd.ElementType_BYREF, Value: i4}},
		}},
		{"typedref-index", []byte{1, 8, 0x16}, i4, []winmd.SigParam{
			{Kind: winmd.SigParamKind_TypedByRef, Type: winmd.SigType{Kind: winmd.ElementType_TYPEDBYREF}},
		}},
		{"void-pointer", []byte{0, 0x0f, 1}, winmd.SigType{Kind: winmd.ElementType_PTR, Value: winmd.SigType{Kind: winmd.ElementType_VOID}}, nil},
		{"array", []byte{0, 0x14, 8, 2, 1, 3, 1, 0x7f}, winmd.SigType{Kind: winmd.ElementType_ARRAY,
			Value: winmd.SigArray{Type: i4, Rank: 2, Sizes: []uint32{3}, LowerBounds: []int32{-1}}}, nil},
		{"vector", []byte{0, 0x1d, 0x13, 0}, vector, nil},
		{"generic-indexer", []byte{1, 0x15, 0x12, 5, 1, 0x1d, 0x13, 0, 0x1d, 0x13, 0},
			winmd.SigType{Kind: winmd.ElementType_GENERICINST, Value: winmd.SigGenericInst{Class: true, Index: handle, Type: []winmd.SigType{vector}}},
			[]winmd.SigParam{{Type: vector}}},
	} {
		for _, header := range []byte{8, 0x28} {
			t.Run(fmt.Sprintf("%s/header-%x", test.name, header), func(t *testing.T) {
				var m winmd.Metadata
				data := append([]byte{header}, test.data...)
				before := bytes.Clone(data)
				got, err := m.PropertySignature(data)
				want := winmd.SigProperty{HasThis: header == 0x28, SigField: winmd.SigField{Type: test.typ}, Param: test.param}
				if err != nil || !reflect.DeepEqual(got, want) {
					t.Fatalf("property signature = %#v, %v; want %#v", got, err, want)
				}
				if !bytes.Equal(data, before) {
					t.Fatal("property signature decoder changed its input")
				}
			})
		}
	}
}

func TestPropertySignatureHeader(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for header := range 256 {
		t.Run(fmt.Sprintf("%02x", header), func(t *testing.T) {
			sig, err := m.PropertySignature([]byte{byte(header), 0, 8})
			if header == 8 || header == 0x28 {
				if err != nil || sig.HasThis != (header == 0x28) || sig.Type.Kind != winmd.ElementType_I4 {
					t.Fatalf("valid property header: %+v, %v", sig, err)
				}
			} else if err == nil {
				t.Fatalf("invalid property header accepted: %+v", sig)
			}
		})
	}
}

func TestPropertySignatureModifiers(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	data := []byte{
		0x28, 1,
		0x20, 5, 0x1f, 9, // Property modifiers.
		0x1d, 0x20, 13, 8, // Vector with a separately modified I4 element.
		0x1f, 17, 0x10, 8, // Modified BYREF index parameter.
	}
	mod := func(kind winmd.SigCustomModKind, index winmd.Index) winmd.SigCustomMod {
		return winmd.SigCustomMod{Kind: kind, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: index}}
	}
	want := winmd.SigProperty{
		HasThis: true,
		SigField: winmd.SigField{Type: winmd.SigType{
			Kind:  winmd.ElementType_SZARRAY,
			Mod:   []winmd.SigCustomMod{mod(winmd.SigCustomModKind_Opt, 0), mod(winmd.SigCustomModKind_Reqd, 1)},
			Value: winmd.SigType{Kind: winmd.ElementType_I4, Mod: []winmd.SigCustomMod{mod(winmd.SigCustomModKind_Opt, 2)}},
		}},
		Param: []winmd.SigParam{{Kind: winmd.SigParamKind_ByRef, Type: winmd.SigType{
			Kind:  winmd.ElementType_BYREF,
			Mod:   []winmd.SigCustomMod{mod(winmd.SigCustomModKind_Reqd, 3)},
			Value: winmd.SigType{Kind: winmd.ElementType_I4},
		}}},
	}
	got, err := m.PropertySignature(data)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("modified property = %#v, %v; want %#v", got, err, want)
	}
}

func TestPropertySignatureParameterCounts(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, test := range []struct {
		count int
		data  []byte
	}{
		{0, []byte{0}},
		{1, []byte{1}},
		{127, []byte{0x7f}},
		{128, []byte{0x80, 0x80}},
		{16384, []byte{0xc0, 0, 0x40, 0}},
	} {
		t.Run(fmt.Sprint(test.count), func(t *testing.T) {
			data := append([]byte{8}, test.data...)
			data = append(data, 8) // The property type is not an index parameter.
			data = append(data, bytes.Repeat([]byte{0x0e}, test.count)...)
			sig, err := m.PropertySignature(data)
			if err != nil || len(sig.Param) != test.count || sig.Type.Kind != winmd.ElementType_I4 {
				t.Fatalf("count = %d, type = %v, error = %v; want count %d and I4", len(sig.Param), sig.Type.Kind, err, test.count)
			}
			for _, param := range sig.Param {
				if param.Kind != winmd.SigParamKind_ByValue || param.Type.Kind != winmd.ElementType_STRING {
					t.Fatalf("index parameter = %+v; want by-value STRING", param)
				}
			}
		})
	}
	for _, count := range [][]byte{{2}, {0x80, 0x80}, {0xc0, 1, 0, 0}, {0xdf, 0xff, 0xff, 0xff}} {
		t.Run(fmt.Sprintf("truncated-%x", count), func(t *testing.T) {
			data := append([]byte{8}, count...)
			sig, err := m.PropertySignature(append(data, 8, 8))
			if !errors.Is(err, io.ErrUnexpectedEOF) || cap(sig.Param) != 0 {
				t.Fatalf("impossible count: error = %v, capacity = %d; want unexpected EOF without allocation", err, cap(sig.Param))
			}
		})
	}
	t.Run("invalid-first-parameter", func(t *testing.T) {
		data := append([]byte{8, 0x80, 0x80, 8}, make([]byte, 128)...)
		sig, err := m.PropertySignature(data)
		if err == nil || cap(sig.Param) != 0 {
			t.Fatalf("invalid first parameter: error = %v, capacity = %d; want error without allocation", err, cap(sig.Param))
		}
	})
}

func TestPropertySignatureTruncated(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, data := range [][]byte{
		{8, 0, 8},
		{0x28, 1, 0x0e, 8},
		{8, 0, 0x20, 5, 0x12, 0x80, 0x81},
		{8, 0, 0x14, 8, 1, 1, 3, 1, 0},
		{0x28, 1, 8, 0x20, 5, 0x10, 0x15, 0x12, 9, 1, 0x1d, 0x13, 0x80, 0x80},
		append([]byte{8, 0x80, 0x80, 8}, bytes.Repeat([]byte{8}, 128)...),
	} {
		t.Run(fmt.Sprintf("%x", data), func(t *testing.T) {
			if _, err := m.PropertySignature(data); err != nil {
				t.Fatalf("complete property signature: %v", err)
			}
			for length := range len(data) {
				if _, err := m.PropertySignature(data[:length]); !errors.Is(err, io.ErrUnexpectedEOF) {
					t.Fatalf("prefix length %d: error = %v; want unexpected EOF", length, err)
				}
			}
		})
	}
}

func TestPropertySignatureMalformed(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, test := range []struct {
		name string
		data []byte
	}{
		{"void-property", []byte{8, 0, 1}},
		{"byref-property", []byte{8, 0, 0x10, 8}},
		{"typedref-property", []byte{8, 0, 0x16}},
		{"void-index", []byte{8, 1, 8, 1}},
		{"vararg-sentinel", []byte{8, 1, 8, 0x41, 8}},
		{"pinned-property", []byte{8, 0, 0x45, 8}},
		{"pinned-index", []byte{8, 1, 8, 0x45, 8}},
		{"modifier-after-byref", []byte{8, 1, 8, 0x10, 0x20, 5, 8}},
		{"invalid-count", []byte{8, 0xff, 8}},
		{"wide-element-code", []byte{8, 0, 0x81, 0x08}},
		{"null-property-handle", []byte{8, 0, 0x12, 0}},
		{"zero-row-index-handle", []byte{8, 1, 8, 0x11, 1}},
		{"reserved-property-handle", []byte{8, 0, 0x12, 7}},
		{"null-modifier-handle", []byte{8, 0, 0x20, 0, 8}},
		{"invalid-generic-kind", []byte{8, 0, 0x15, 8, 5, 1, 8}},
		{"vector-of-void", []byte{8, 0, 0x1d, 1}},
	} {
		t.Run(test.name, func(t *testing.T) {
			if sig, err := m.PropertySignature(test.data); err == nil {
				t.Fatalf("malformed property signature accepted: %+v", sig)
			}
		})
	}
	for _, data := range [][]byte{{8, 0, 8}, {0x28, 1, 8, 8}} {
		for _, suffix := range []byte{0, 8, 0xff} {
			_, err := m.PropertySignature(append(bytes.Clone(data), suffix))
			if err == nil || !strings.Contains(err.Error(), "trailing") {
				t.Fatalf("signature %x with suffix %x: %v; want trailing-data error", data, suffix, err)
			}
		}
	}
}

func TestPropertySignatureNesting(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, kind := range []struct {
		name   string
		prefix []byte
		suffix []byte
	}{
		{"pointer", []byte{0x0f}, nil},
		{"vector", []byte{0x1d}, nil},
		{"generic", []byte{0x15, 0x12, 5, 1}, nil},
		{"array", []byte{0x14}, []byte{1, 0, 0}},
	} {
		for _, layers := range []int{0, 63, 64, 65, 4096} {
			t.Run(fmt.Sprintf("%s/%d", kind.name, layers), func(t *testing.T) {
				typ := append(bytes.Repeat(kind.prefix, layers), 8)
				typ = append(typ, bytes.Repeat(kind.suffix, layers)...)
				for _, header := range [][]byte{{8, 0}, {8, 1, 8}} {
					_, err := m.PropertySignature(append(bytes.Clone(header), typ...))
					if layers >= 64 {
						if err == nil || !strings.Contains(err.Error(), "nesting limit") {
							t.Fatalf("header %x: error = %v; want nesting limit error", header, err)
						}
					} else if err != nil {
						t.Fatalf("header %x: valid depth %d rejected: %v", header, layers+1, err)
					}
				}
			})
		}
	}
	t.Run("independent-type-and-index-parameters", func(t *testing.T) {
		typ := append(bytes.Repeat([]byte{0x1d}, 63), 8)
		data := []byte{0x28, 2}
		for range 3 { // The property type and two index parameters each get 64 levels.
			data = append(data, typ...)
		}
		if _, err := m.PropertySignature(data); err != nil {
			t.Fatalf("independent property types share a depth budget: %v", err)
		}
	})
}

func TestPropertySignatureFromMetadata(t *testing.T) {
	t.Parallel()
	i4 := winmd.SigType{Kind: winmd.ElementType_I4}
	tests := []struct {
		name string
		sig  winmd.SigPropertyBlob
		want winmd.SigProperty
	}{
		{"Item", []byte{0x28, 1, 0x1d, 0x12, 4, 8}, winmd.SigProperty{
			HasThis: true,
			SigField: winmd.SigField{Type: winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: winmd.SigType{
				Kind: winmd.ElementType_CLASS, Value: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeDef, Index: 0},
			}}},
			Param: []winmd.SigParam{{Type: i4}},
		}},
		{"Count", []byte{8, 0, 8}, winmd.SigProperty{SigField: winmd.SigField{Type: i4}}},
	}
	// One TypeDef owns two properties through a PropertyMap row.
	tables := metadataTables(0, 1<<2|1<<21|1<<23, []uint32{1, 1, 2}, 0)
	tables = binary.LittleEndian.AppendUint32(tables, uint32(winmd.TypeVisibility_Public))
	for _, value := range []uint16{1, 11, 0, 1, 1, 1, 1} { // TypeDef columns, then PropertyMap.
		tables = binary.LittleEndian.AppendUint16(tables, value)
	}
	strs := []byte("\x00Container\x00Test\x00")
	blobs := []byte{0}
	for _, test := range tests {
		tables = binary.LittleEndian.AppendUint16(tables, 0) // Property.Flags.
		tables = binary.LittleEndian.AppendUint16(tables, uint16(len(strs)))
		tables = binary.LittleEndian.AppendUint16(tables, uint16(len(blobs)))
		strs = append(strs, test.name...)
		strs = append(strs, 0)
		blobs = append(blobs, byte(len(test.sig)))
		blobs = append(blobs, test.sig...)
	}
	m, err := winmd.New(metadataPE(t, metadataRoot(
		metadataStream{"#~", tables}, metadataStream{"#Strings", strs}, metadataStream{"#Blob", blobs},
	)))
	if err != nil {
		t.Fatal(err)
	}
	mapping, err := m.Tables.PropertyMap.At(0)
	if err != nil || mapping.Parent != 0 || mapping.PropertyList.Len() != 2 {
		t.Fatalf("property map = %+v, %v; want two properties on the first type", mapping, err)
	}
	for index := range mapping.PropertyList.All() {
		property, err := m.Tables.Property.At(index)
		if err != nil {
			t.Fatal(err)
		}
		test := tests[index]
		if property.Name.String() != test.name || !bytes.Equal(property.Type, test.sig) {
			t.Fatalf("property row = %+v; want %s with signature %x", property, test.name, test.sig)
		}
		sig, err := m.PropertySignature(property.Type)
		if err != nil || !reflect.DeepEqual(sig, test.want) {
			t.Fatalf("%s signature = %#v, %v; want %#v", test.name, sig, err, test.want)
		}
	}
}

func ExampleMetadata_PropertySignature() {
	var m winmd.Metadata
	// Instance property returning STRING, indexed by one I4 parameter.
	sig, err := m.PropertySignature([]byte{0x28, 1, 0x0e, 8})
	if err != nil {
		panic(err)
	}
	fmt.Println("instance:", sig.HasThis)
	fmt.Println("property type:", sig.Type.Kind)
	fmt.Println("index type:", sig.Param[0].Type.Kind)
	// Output:
	// instance: true
	// property type: STRING
	// index type: I4
}

func FuzzPropertySignature(f *testing.F) {
	for _, data := range [][]byte{
		{8, 0, 8},
		{0x28, 1, 0x0e, 8},
		{0x28, 1, 0x20, 5, 8, 0x1f, 9, 0x10, 8},
		{8, 0, 0x15, 0x12, 5, 1, 0x1d, 0x13, 0},
		{8, 0xdf, 0xff, 0xff, 0xff, 8},
	} {
		f.Add(data)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		var m winmd.Metadata
		_, _ = m.PropertySignature(data)
	})
}
