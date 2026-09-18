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
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

type specSignatureCase struct {
	name string
	data []byte
	want winmd.SigType
}

func specSignatureCases() []specSignatureCase {
	i4 := winmd.SigType{Kind: winmd.ElementType_I4}
	var0 := winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(0)}
	mvar := winmd.SigType{Kind: winmd.ElementType_MVAR, Value: uint32(0x1fffffff)}
	def := winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeDef, Index: 0}
	ref := winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 0}
	mods := []winmd.SigCustomMod{{Kind: winmd.SigCustomModKind_Opt, Index: ref}, {Kind: winmd.SigCustomModKind_Reqd, Index: def}}
	vector := winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: var0}
	array := winmd.SigType{Kind: winmd.ElementType_ARRAY, Value: winmd.SigArray{Type: i4, Rank: 2, Sizes: []uint32{3, 128}, LowerBounds: []int32{-1, 2}}}
	inst := winmd.SigType{Kind: winmd.ElementType_GENERICINST, Value: winmd.SigGenericInst{Class: true, Index: ref, Type: []winmd.SigType{vector, mvar}}}
	instData := []byte{0x15, 0x12, 5, 2, 0x1d, 0x13, 0, 0x1e, 0xdf, 0xff, 0xff, 0xff}
	fn := winmd.SigType{Kind: winmd.ElementType_FNPTR, Value: winmd.SigStandAloneMethod{
		CallingConvention: winmd.SigCallingConvention_Cdecl,
		SigMethodRef: winmd.SigMethodRef{
			SigMethodDef:  winmd.SigMethodDef{RetType: winmd.SigRetType{Kind: winmd.SigRetTypeKind_Void, Type: winmd.SigType{Kind: winmd.ElementType_VOID}}, Param: []winmd.SigParam{{Type: i4}}},
			VariableParam: []winmd.SigParam{{Type: var0}},
		},
	}}
	fnData := []byte{0x1b, 1, 2, 1, 8, 0x41, 0x13, 0}
	cases := []specSignatureCase{
		{"class", []byte{0x12, 5}, winmd.SigType{Kind: winmd.ElementType_CLASS, Value: ref}},
		{"value-type", []byte{0x11, 4}, winmd.SigType{Kind: winmd.ElementType_VALUETYPE, Value: def}},
		{"type-spec-handle", []byte{0x12, 6}, winmd.SigType{Kind: winmd.ElementType_CLASS, Value: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeSpec, Index: 0}}},
		{"compressed-handle", []byte{0x12, 0x80, 0x81}, winmd.SigType{Kind: winmd.ElementType_CLASS, Value: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 31}}},
		{"var", []byte{0x13, 0}, var0},
		{"compressed-var", []byte{0x13, 0x80, 0x80}, winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(128)}},
		{"max-mvar", []byte{0x1e, 0xdf, 0xff, 0xff, 0xff}, mvar},
		{"pointer", []byte{0x0f, 8}, winmd.SigType{Kind: winmd.ElementType_PTR, Value: i4}},
		{"void-pointer", []byte{0x0f, 1}, winmd.SigType{Kind: winmd.ElementType_PTR, Value: winmd.SigType{Kind: winmd.ElementType_VOID}}},
		{"modified-pointer", []byte{0x0f, 0x20, 5, 0x1f, 4, 8}, winmd.SigType{Kind: winmd.ElementType_PTR, Value: winmd.SigType{Kind: winmd.ElementType_I4, Mod: mods}}},
		{"vector", []byte{0x1d, 0x13, 0}, vector},
		{"nested-modifiers", []byte{0x1d, 0x20, 5, 0x1f, 4, 0x0f, 0x20, 5, 8}, winmd.SigType{Kind: winmd.ElementType_SZARRAY,
			Value: winmd.SigType{Kind: winmd.ElementType_PTR, Mod: mods, Value: winmd.SigType{Kind: winmd.ElementType_I4, Mod: mods[:1]}}}},
		{"array", []byte{0x14, 8, 2, 2, 3, 0x80, 0x80, 2, 0x7f, 4}, array},
		{"open-generic", instData, inst},
		{"nested-generic", append([]byte{0x15, 0x11, 4, 1}, instData...), winmd.SigType{Kind: winmd.ElementType_GENERICINST,
			Value: winmd.SigGenericInst{Index: def, Type: []winmd.SigType{inst}}}},
		{"generic-array-argument", []byte{0x15, 0x12, 5, 1, 0x14, 8, 2, 2, 3, 0x80, 0x80, 2, 0x7f, 4}, winmd.SigType{Kind: winmd.ElementType_GENERICINST,
			Value: winmd.SigGenericInst{Class: true, Index: ref, Type: []winmd.SigType{array}}}},
		{"array-of-generic", append(append([]byte{0x14}, instData...), 1, 1, 3, 1, 0), winmd.SigType{Kind: winmd.ElementType_ARRAY,
			Value: winmd.SigArray{Type: inst, Rank: 1, Sizes: []uint32{3}, LowerBounds: []int32{0}}}},
		{"function-pointer", fnData, fn},
		{"vector-of-function-pointers", append([]byte{0x1d}, fnData...), winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: fn}},
	}
	for _, kind := range []byte{2, 3, 4, 5, 6, 7, 8, 9, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x18, 0x19, 0x1c} {
		cases = append(cases, specSignatureCase{fmt.Sprintf("primitive-%02x", kind), []byte{kind}, winmd.SigType{Kind: winmd.ElementType(kind)}})
	}
	return cases
}

func decodeSpecSignature(m *winmd.Metadata, method bool, data []byte) (any, error) {
	if method {
		return m.MethodSpecSignature(data)
	}
	return m.TypeSpecSignature(data)
}

func TestSpecSignatureShapes(t *testing.T) {
	var m winmd.Metadata
	for _, test := range specSignatureCases() {
		for _, method := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/method-%v", test.name, method), func(t *testing.T) {
				data := bytes.Clone(test.data)
				var want any = winmd.SigTypeSpec{Kind: test.want.Kind, Value: test.want.Value}
				if method {
					data = append([]byte{0x0a, 1}, data...)
					want = winmd.SigMethodSpec{test.want}
				}
				before := bytes.Clone(data)
				got, err := decodeSpecSignature(&m, method, data)
				if method && (test.want.Kind == winmd.ElementType_PTR || test.want.Kind == winmd.ElementType_FNPTR) {
					if err == nil {
						t.Fatal("direct pointer accepted as a method type argument")
					}
					return
				}
				if err != nil || !reflect.DeepEqual(got, want) {
					t.Fatalf("signature = %#v, %v; want %#v", got, err, want)
				}
				if !bytes.Equal(data, before) {
					t.Fatal("decoder changed its input")
				}
				// Validate the complete shape before checking every truncated prefix.
				for length := range len(data) {
					if _, err := decodeSpecSignature(&m, method, data[:length]); !errors.Is(err, io.ErrUnexpectedEOF) {
						t.Fatalf("prefix %d: %v; want unexpected EOF", length, err)
					}
				}
				for _, suffix := range []byte{0, 8, 0xff} {
					if _, err := decodeSpecSignature(&m, method, append(bytes.Clone(data), suffix)); err == nil {
						t.Fatalf("trailing byte %02x accepted", suffix)
					}
				}
			})
		}
	}
}

func TestSpecSignatureMalformed(t *testing.T) {
	var m winmd.Metadata
	for _, data := range [][]byte{
		{0x1f, 4, 8}, {0x20, 5, 8}, // Root modifiers are not part of either grammar.
		{1}, {0x10, 8}, {0x16}, {0x45, 8}, {0x41, 8}, {0, 0, 8},
		{0x1d, 1}, {0x14, 0x20, 5, 8, 1, 0, 0}, {0x0f, 0x10, 8},
		{0x12, 0}, {0x12, 1}, {0x12, 2}, {0x12, 7}, {0x0f, 0x20, 0, 8},
		{0x13, 0xff}, {0x1e, 0xff}, {0x81, 8}, {0x21},
		{0x15, 0x12, 5, 0}, {0x15, 8, 5, 1, 8},
		{0x15, 0x12, 5, 1, 0x20, 5, 8},
		{0x15, 0x12, 5, 1, 0x0f, 8}, {0x15, 0x12, 5, 1, 0x1b, 0, 0, 1},
	} {
		for _, method := range []bool{false, true} {
			t.Run(fmt.Sprintf("%x/method-%v", data, method), func(t *testing.T) {
				blob := bytes.Clone(data)
				if method {
					blob = append([]byte{0x0a, 1}, blob...)
				}
				if got, err := decodeSpecSignature(&m, method, blob); err == nil {
					t.Fatalf("malformed signature accepted: %#v", got)
				}
			})
		}
	}
	// In particular, GENERICINST (0x15) is not the MethodSpec marker (0x0a).
	for header := range 256 {
		if _, err := m.MethodSpecSignature([]byte{byte(header), 1, 8}); (err == nil) != (header == 0x0a) {
			t.Fatalf("MethodSpec header %02x: %v", header, err)
		}
	}
}

func TestSpecSignatureCounts(t *testing.T) {
	var m winmd.Metadata
	for _, test := range []struct {
		count int
		data  []byte
	}{{1, []byte{1}}, {127, []byte{0x7f}}, {128, []byte{0x80, 0x80}}, {16383, []byte{0xbf, 0xff}}, {16384, []byte{0xc0, 0, 0x40, 0}}} {
		header := append([]byte{0x0a}, test.data...)
		data := append(bytes.Clone(header), bytes.Repeat([]byte{8}, test.count)...)
		got, err := m.MethodSpecSignature(data)
		if err != nil || len(got) != test.count {
			t.Fatalf("count %d: got %d arguments, %v", test.count, len(got), err)
		}
		for _, arg := range got {
			if !reflect.DeepEqual(arg, winmd.SigType{Kind: winmd.ElementType_I4}) {
				t.Fatalf("argument = %#v; want I4", arg)
			}
		}
		for length := range len(header) {
			if _, err := m.MethodSpecSignature(data[:length]); !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatalf("count %d, prefix %d: %v; want unexpected EOF", test.count, length, err)
			}
		}
		if _, err := m.MethodSpecSignature(data[:len(data)-1]); !errors.Is(err, io.ErrUnexpectedEOF) {
			t.Fatalf("count %d, missing argument: %v; want unexpected EOF", test.count, err)
		}
	}
	for _, count := range [][]byte{{0}, {0x80, 0}, {0xc0, 0, 0, 0}, {0xe0}, {0xff}} {
		if _, err := m.MethodSpecSignature(append([]byte{0x0a}, count...)); err == nil {
			t.Fatalf("invalid count %x accepted", count)
		}
	}
	for _, count := range [][]byte{{2}, {0x80, 0x80}, {0xc0, 0, 0x40, 0}, {0xdf, 0xff, 0xff, 0xff}} {
		data := append(append([]byte{0x0a}, count...), 8)
		if _, err := m.MethodSpecSignature(data); !errors.Is(err, io.ErrUnexpectedEOF) {
			t.Fatalf("impossible count %x: %v; want unexpected EOF", count, err)
		}
		if allocs := testing.AllocsPerRun(10, func() { _, _ = m.MethodSpecSignature(data) }); allocs != 0 {
			t.Fatalf("impossible count %x allocated %g times", count, allocs)
		}
	}
}

func TestSpecSignatureDepth(t *testing.T) {
	var m winmd.Metadata
	for _, pattern := range []struct {
		name           string
		prefix, suffix []byte
	}{
		{"pointer", []byte{0x0f}, nil}, {"vector", []byte{0x1d}, nil},
		{"array", []byte{0x14}, []byte{1, 0, 0}},
		{"generic", []byte{0x15, 0x12, 5, 1}, nil},
		{"function-pointer", []byte{0x1b, 0, 0}, nil},
	} {
		for _, layers := range []int{63, 64, 4096} {
			for _, method := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/%d/method-%v", pattern.name, layers, method), func(t *testing.T) {
					n, data := layers, []byte(nil)
					if method {
						// A vector allows nested pointers without making them direct arguments.
						n, data = layers-1, []byte{0x1d}
					}
					data = append(data, bytes.Repeat(pattern.prefix, n)...)
					data = append(data, 8)
					data = append(data, bytes.Repeat(pattern.suffix, n)...)
					if method {
						// Each sibling gets its own complete 64-level budget.
						data = append(append([]byte{0x0a, 2}, data...), data...)
					}
					if _, err := decodeSpecSignature(&m, method, data); (err != nil) != (layers >= 64) {
						t.Fatalf("depth %d: %v; want error %v", layers+1, err, layers >= 64)
					}
				})
			}
		}
	}
}

func TestSpecSignatureMetadata(t *testing.T) {
	wantOpen := winmd.SigMethodSpec{{Kind: winmd.ElementType_VAR, Value: uint32(128)}, {Kind: winmd.ElementType_MVAR, Value: uint32(0x1fffffff)}}
	for _, wrapper := range []struct{ prefix, suffix []byte }{
		{[]byte{0x12}, nil}, {[]byte{0x11}, nil}, {[]byte{0x15, 0x12}, []byte{1, 8}},
		{[]byte{0x1d, 0x1f}, []byte{8}}, {[]byte{0x1d, 0x0f, 0x20}, []byte{8}},
		{[]byte{0x1d, 0x1b, 0, 0, 0x12}, nil},
	} {
		for _, handle := range []byte{4, 5, 6, 8, 9, 10} {
			t.Run(fmt.Sprintf("%x/handle-%d", wrapper.prefix, handle), func(t *testing.T) {
				data := append(append(bytes.Clone(wrapper.prefix), handle), wrapper.suffix...)
				methodData := append([]byte{0x0a, 3}, data...)
				methodData = append(methodData, 0x13, 0x80, 0x80, 0x1e, 0xdf, 0xff, 0xff, 0xff)
				blobs := []byte{0, 4, 0x10, 1, 0, 1} // MethodDef: one generic parameter, no ordinary parameters, VOID result.
				addBlob := func(data []byte) uint16 {
					index := uint16(len(blobs))
					blobs = append(blobs, byte(len(data)))
					blobs = append(blobs, data...)
					return index
				}
				typeIndex, methodIndex := addBlob(data), addBlob(methodData)
				tables := metadataTables(0, 1<<1|1<<2|1<<6|1<<27|1<<42|1<<43, []uint32{1, 1, 1, 1, 1, 1}, 6+14)
				binary.LittleEndian.PutUint16(tables[len(tables)-2:], 1) // TypeDef.MethodList owns MethodDef row 1.
				tables = append(tables, make([]byte, 10)...)             // MethodDef through Name.
				// MethodDef tail, TypeSpec, GenericParam (value-type constraint), then MethodSpec.
				// MethodSpec.Method = 2 is a non-null reference to MethodDef row 1.
				for _, column := range []uint16{1, 0, typeIndex, 0, 8, 3, 0, 2, methodIndex} {
					tables = binary.LittleEndian.AppendUint16(tables, column)
				}
				m, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables}, metadataStream{"#Blob", blobs})))
				if err != nil {
					t.Fatal(err)
				}
				typ, err := m.Tables.TypeSpec.At(0)
				if err != nil || !bytes.Equal(typ.Signature, data) {
					t.Fatalf("TypeSpec row = %+v, %v", typ, err)
				}
				method, err := m.Tables.MethodSpec.At(0)
				if err != nil || !bytes.Equal(method.Instantiation, methodData) || method.Method != (winmd.CodedIndex[winmd.MethodDefOrRef]{Tag: winmd.MethodDefOrRef_MethodDef, Index: 0}) {
					t.Fatalf("MethodSpec row = %+v, %v", method, err)
				}
				_, typeErr := m.TypeSpecSignature(typ.Signature)
				args, methodErr := m.MethodSpecSignature(method.Instantiation)
				if (typeErr != nil) != (handle >= 8) || (methodErr != nil) != (handle >= 8) {
					t.Fatalf("handle %d: TypeSpec error %v, MethodSpec error %v", handle, typeErr, methodErr)
				}
				// Decoding does not enforce the target's arity or value-type constraint.
				if handle < 8 && (len(args) != 3 || !reflect.DeepEqual(args[1:], wantOpen)) {
					t.Fatalf("open arguments = %#v; want three arguments ending in %#v", args, wantOpen)
				}
			})
		}
	}
}

func ExampleMetadata_TypeSpecSignature() {
	var m winmd.Metadata
	// SZARRAY !0 has no calling-convention header.
	sig, err := m.TypeSpecSignature(winmd.SigTypeSpecBlob{0x1d, 0x13, 0})
	if err != nil {
		panic(err)
	}
	element := sig.Value.(winmd.SigType)
	fmt.Println(sig.Kind, element.Kind, element.Value)
	// Output: SZARRAY VAR 0
}

func ExampleMetadata_MethodSpecSignature() {
	var m winmd.Metadata
	sig, err := m.MethodSpecSignature(winmd.SigMethodSpecBlob{0x0a, 2, 8, 0x1e, 1})
	if err != nil {
		panic(err)
	}
	fmt.Println("arguments:", len(sig))
	fmt.Println(sig[0].Kind, sig[1].Kind, sig[1].Value)
	// Output:
	// arguments: 2
	// I4 MVAR 1
}

func FuzzSpecSignatures(f *testing.F) {
	for _, test := range specSignatureCases() {
		f.Add(false, test.data)
		f.Add(true, append([]byte{0x0a, 1}, test.data...))
	}
	for _, data := range [][]byte{{}, {0x0a}, {0x15, 1, 8}, {0x0a, 0}, {0x0a, 0xdf, 0xff, 0xff, 0xff, 8}, append(bytes.Repeat([]byte{0x1d}, 64), 8)} {
		f.Add(false, data)
		f.Add(true, data)
	}
	f.Fuzz(func(t *testing.T, method bool, data []byte) {
		var m winmd.Metadata
		before := bytes.Clone(data)
		_, err := decodeSpecSignature(&m, method, data)
		if !bytes.Equal(data, before) {
			t.Fatal("decoder changed its input")
		}
		if err == nil {
			if _, err := decodeSpecSignature(&m, method, append(bytes.Clone(data), 0)); err == nil {
				t.Fatal("decoder accepted trailing data")
			}
		}
	})
}
