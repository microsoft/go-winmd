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

func methodSignatureDecoders() map[string]func([]byte) (winmd.SigStandAloneMethod, error) {
	var m winmd.Metadata
	return map[string]func([]byte) (winmd.SigStandAloneMethod, error){
		"reference": func(data []byte) (winmd.SigStandAloneMethod, error) {
			sig, err := m.MethodRefSignature(data)
			return winmd.SigStandAloneMethod{SigMethodRef: sig}, err
		},
		"standalone": func(data []byte) (winmd.SigStandAloneMethod, error) {
			return m.StandAloneMethodSignature(data)
		},
		"function-pointer": func(data []byte) (winmd.SigStandAloneMethod, error) {
			sig, err := m.FieldSignature(append([]byte{6, 0x1b}, data...))
			if err != nil {
				return winmd.SigStandAloneMethod{}, err
			}
			method, ok := sig.Type.Value.(winmd.SigStandAloneMethod)
			if !ok || sig.Type.Kind != winmd.ElementType_FNPTR {
				return method, fmt.Errorf("FNPTR representation = %#v", sig.Type)
			}
			return method, nil
		},
	}
}

func TestMethodCallSignatureHeaders(t *testing.T) {
	t.Parallel()
	conventions := [...]winmd.SigCallingConvention{
		winmd.SigCallingConvention_Default, winmd.SigCallingConvention_Cdecl,
		winmd.SigCallingConvention_Stdcall, winmd.SigCallingConvention_Thiscall,
		winmd.SigCallingConvention_Fastcall, winmd.SigCallingConvention_Vararg,
	}
	for name, decode := range methodSignatureDecoders() {
		for header := range 256 {
			t.Run(fmt.Sprintf("%s/%02x", name, header), func(t *testing.T) {
				data, generic := []byte{byte(header)}, uint32(0)
				if header&0x10 != 0 {
					data, generic = append(data, 1), 1
				}
				kind := header & 15
				valid := header&0x80 == 0 && (header&0x40 == 0 || header&0x20 != 0) && kind <= 5
				valid = valid && (name != "reference" || kind == 0 || kind == 5)
				valid = valid && (generic == 0 || name != "standalone" && kind == 0)
				sig, err := decode(append(data, 0, 1))
				if (err == nil) != valid {
					t.Fatalf("header %#x: error = %v; want valid %v", header, err, valid)
				}
				if err != nil {
					return
				}
				if sig.HasThis != (header&0x20 != 0) || sig.ExplicitThis != (header&0x40 != 0) || sig.VarArgs != (kind == 5) || sig.Generic != generic || sig.RetType.Kind != winmd.SigRetTypeKind_Void || len(sig.Param)+len(sig.VariableParam) != 0 {
					t.Fatalf("header %#x decoded incorrectly: %+v", header, sig)
				}
				if name != "reference" && sig.CallingConvention != conventions[kind] {
					t.Fatalf("calling convention = %v; want %v", sig.CallingConvention, conventions[kind])
				}
			})
		}
	}
}

func TestMethodCallSignatureValues(t *testing.T) {
	t.Parallel()
	i4 := winmd.SigType{Kind: winmd.ElementType_I4}
	mods := []winmd.SigCustomMod{
		{Kind: winmd.SigCustomModKind_Opt, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 0}},
		{Kind: winmd.SigCustomModKind_Reqd, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 1}},
	}
	for _, ret := range []struct {
		data []byte
		kind winmd.SigRetTypeKind
		typ  winmd.SigType
	}{
		{[]byte{1}, winmd.SigRetTypeKind_Void, winmd.SigType{Kind: winmd.ElementType_VOID}},
		{[]byte{8}, winmd.SigRetTypeKind_ByValue, i4},
		{[]byte{0x10, 8}, winmd.SigRetTypeKind_ByRef, winmd.SigType{Kind: winmd.ElementType_BYREF, Value: i4}},
		{[]byte{0x16}, winmd.SigRetTypeKind_TypedByRef, winmd.SigType{Kind: winmd.ElementType_TYPEDBYREF}},
	} {
		ret.typ.Mod = mods
		data := append([]byte{0x65, 3, 0x20, 5, 0x1f, 9}, ret.data...)
		data = append(data, 0x13, 0x80, 0x80, 0x41, 0x20, 5, 0x1f, 9, 0x10, 8, 0x20, 5, 0x1f, 9, 0x16)
		want := winmd.SigMethodRef{
			SigMethodDef: winmd.SigMethodDef{
				HasThis: true, ExplicitThis: true, VarArgs: true,
				RetType: winmd.SigRetType{Kind: ret.kind, Type: ret.typ},
				Param:   []winmd.SigParam{{Type: winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(128)}}},
			},
			VariableParam: []winmd.SigParam{
				{Kind: winmd.SigParamKind_ByRef, Type: winmd.SigType{Kind: winmd.ElementType_BYREF, Mod: mods, Value: i4}},
				{Kind: winmd.SigParamKind_TypedByRef, Type: winmd.SigType{Kind: winmd.ElementType_TYPEDBYREF, Mod: mods}},
			},
		}
		for name, decode := range methodSignatureDecoders() {
			t.Run(fmt.Sprintf("%s/return-%x", name, ret.data), func(t *testing.T) {
				if got, err := decode(data); err != nil || !reflect.DeepEqual(got.SigMethodRef, want) {
					t.Fatalf("signature = %#v, %v; want %#v", got, err, want)
				}
			})
		}
	}
}

func TestMethodCallSignatureGenerics(t *testing.T) {
	t.Parallel()
	for name, decode := range methodSignatureDecoders() {
		for _, generic := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/generic-%v", name, generic), func(t *testing.T) {
				// Open parameter references need not be bounded by a declaration's arity.
				data := []byte{0, 1, 0x1e, 0xdf, 0xff, 0xff, 0xff, 0x13, 0}
				var arity uint32
				if generic {
					data, arity = append([]byte{0x10, 0x80, 0x80}, data[1:]...), 128
				}
				got, err := decode(data)
				if generic && name == "standalone" {
					if err == nil {
						t.Fatal("standalone generic declaration accepted")
					}
					return
				}
				want := winmd.SigType{Kind: winmd.ElementType_MVAR, Value: uint32(0x1fffffff)}
				if err != nil || got.Generic != arity || !reflect.DeepEqual(got.RetType.Type, want) || len(got.Param) != 1 || got.Param[0].Type.Kind != winmd.ElementType_VAR || got.Param[0].Type.Value != uint32(0) {
					t.Fatalf("open generic signature = %+v, %v", got, err)
				}
			})
		}
	}
}

func TestMethodCallSignatureSentinel(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for name, decode := range methodSignatureDecoders() {
		for convention := range byte(6) {
			for _, counts := range [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {127, 1}} {
				t.Run(fmt.Sprintf("%s/%d/%d+%d", name, convention, counts[0], counts[1]), func(t *testing.T) {
					data := []byte{convention}
					count := counts[0] + counts[1]
					if count >= 128 {
						data = append(data, 0x80)
					}
					data = append(data, byte(count), 1)
					data = append(data, bytes.Repeat([]byte{8}, counts[0])...)
					if counts[1] != 0 {
						data = append(data, 0x41)
						data = append(data, bytes.Repeat([]byte{0x0e}, counts[1])...)
					}
					valid := name != "reference" || convention == 0 || convention == 5
					valid = valid && (counts[1] == 0 || convention == 5 || convention == 1)
					got, err := decode(data)
					if (err == nil) != valid || err == nil && (len(got.Param) != counts[0] || len(got.VariableParam) != counts[1] || got.VarArgs != (convention == 5)) {
						t.Fatalf("signature %x = %+v, %v; want valid %v, split %v", data, got, err, valid, counts)
					}
					if convention == 5 && counts[1] != 0 {
						if _, err := m.MethodDefSignature(data); err == nil {
							t.Fatal("MethodDef accepted SENTINEL")
						}
					}
				})
			}
		}
	}
}

func TestMethodCallSignatureMalformed(t *testing.T) {
	t.Parallel()
	for name, decode := range methodSignatureDecoders() {
		for _, data := range [][]byte{
			{0x10, 0, 0, 1},                   // Zero generic arity.
			{0, 0xe0, 1},                      // Invalid compressed count.
			{0x10, 0xe0, 0, 1},                // Invalid compressed arity.
			{5, 0, 0x41, 1},                   // SENTINEL before the return type.
			{5, 0, 1, 0x41},                   // Empty optional list.
			{5, 1, 1, 8, 0x41},                // Marker after the declared count.
			{5, 2, 1, 8, 0x41},                // Missing optional parameter.
			{5, 1, 1, 0x41, 0x41, 8},          // Adjacent markers.
			{5, 2, 1, 0x41, 8, 0x41, 8},       // Duplicate marker between parameters.
			{1, 0, 1, 0x41},                   // Cdecl also forbids an empty optional list.
			{1, 2, 1, 8, 0x41},                // Cdecl missing optional parameter.
			{1, 2, 1, 0x41, 8, 0x41, 8},       // Cdecl duplicate marker.
			{5, 1, 1, 0x20, 5, 0x41, 8},       // Marker within a parameter's prefixes.
			{5, 1, 1, 0x80, 0x41, 8},          // Noncanonical SENTINEL.
			{5, 1, 1, 0x81, 0x41, 8},          // Element code must not narrow to SENTINEL.
			{0, 1, 1, 0x81, 0x1b, 0, 0, 1},    // Nor to FNPTR.
			{0, 1, 1, 0x80, 0x45, 8},          // PINNED is not a parameter prefix.
			{0, 1, 1, 0x10, 0x80, 0x20, 5, 8}, // Modifier after BYREF.
			{0, 0, 0x80, 0x10, 1},             // BYREF VOID, even with a wide prefix.
			{0, 0, 0x10, 0x80, 0x16},          // BYREF TYPEDBYREF.
			{0, 1, 1, 1},                      // VOID parameter.
		} {
			t.Run(fmt.Sprintf("%s/%x", name, data), func(t *testing.T) {
				if _, err := decode(data); err == nil {
					t.Fatalf("malformed signature %x accepted", data)
				}
			})
		}
	}
}

func TestMethodCallSignatureFraming(t *testing.T) {
	t.Parallel()
	for name, decode := range methodSignatureDecoders() {
		t.Run(name, func(t *testing.T) {
			for _, data := range [][]byte{
				{0, 0, 1},
				{5, 2, 0x20, 0x80, 0x81, 0x16, 0x10, 0x1e, 0xc0, 0, 0x40, 0, 0x41, 0x1b, 1, 1, 1, 0x41, 0x0e},
				{0x10, 0x80, 0x80, 1, 0x1e, 0x80, 0x80, 0x12, 0xc0, 0, 0x40, 1},
			} {
				if data[0] == 0x10 && name == "standalone" {
					continue
				}
				if _, err := decode(data); err != nil {
					t.Fatalf("positive control %x: %v", data, err)
				}
				for length := range len(data) {
					if _, err := decode(data[:length]); !errors.Is(err, io.ErrUnexpectedEOF) {
						t.Fatalf("prefix %x: %v; want unexpected EOF", data[:length], err)
					}
				}
				for _, suffix := range []byte{0, 8, 0x41, 0xff} {
					if _, err := decode(append(bytes.Clone(data), suffix)); err == nil {
						t.Fatalf("trailing byte %#x accepted after %x", suffix, data)
					}
				}
			}
			for _, count := range [][]byte{{0x80, 0x80}, {0xc0, 1, 0, 0}, {0xdf, 0xff, 0xff, 0xff}} {
				data := append([]byte{0}, count...)
				if _, err := decode(append(data, 1, 8)); !errors.Is(err, io.ErrUnexpectedEOF) {
					t.Fatalf("oversized count %x: %v; want unexpected EOF", count, err)
				}
			}
		})
	}
}

func TestFunctionPointerSignatureContexts(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	leaf := winmd.SigType{Kind: winmd.ElementType_FNPTR, Value: winmd.SigStandAloneMethod{
		CallingConvention: winmd.SigCallingConvention_Cdecl,
		SigMethodRef: winmd.SigMethodRef{
			SigMethodDef: winmd.SigMethodDef{
				RetType: winmd.SigRetType{Type: winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(0)}},
				Param:   []winmd.SigParam{{Type: winmd.SigType{Kind: winmd.ElementType_MVAR, Value: uint32(1)}}},
			},
			VariableParam: []winmd.SigParam{{Kind: winmd.SigParamKind_TypedByRef, Type: winmd.SigType{Kind: winmd.ElementType_TYPEDBYREF}}},
		},
	}}
	want := winmd.SigType{Kind: winmd.ElementType_FNPTR, Value: winmd.SigStandAloneMethod{
		SigMethodRef: winmd.SigMethodRef{SigMethodDef: winmd.SigMethodDef{
			RetType: winmd.SigRetType{Type: leaf}, Param: []winmd.SigParam{{Type: leaf}},
		}},
	}}
	inner := []byte{0x1b, 1, 2, 0x13, 0, 0x1e, 1, 0x41, 0x16}
	typ := append(append([]byte{0x1b, 0, 1}, inner...), inner...)
	field, fe := m.FieldSignature(append([]byte{6}, typ...))
	method, me := m.MethodDefSignature(append(append([]byte{0, 1}, typ...), typ...))
	property, pe := m.PropertySignature(append(append([]byte{8, 1}, typ...), typ...))
	if fe != nil || me != nil || pe != nil || len(method.Param) != 1 || len(property.Param) != 1 {
		t.Fatalf("nested FNPTR: field %v, method %v, property %v", fe, me, pe)
	}
	for _, got := range []winmd.SigType{field.Type, method.RetType.Type, method.Param[0].Type, property.Type, property.Param[0].Type} {
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("nested FNPTR = %#v; want %#v", got, want)
		}
	}
	if _, err := m.FieldSignature(append([]byte{6, 0x15, 0x12, 5, 1}, typ...)); err == nil {
		t.Fatal("FNPTR accepted as a generic type argument")
	}
	if _, err := m.MethodSpecSignature(append([]byte{10, 1}, typ...)); err == nil {
		t.Fatal("FNPTR accepted as a generic method argument")
	}
}

func TestFunctionPointerSignatureSentinelBoundary(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, inner := range [][]byte{{0, 0, 1}, {5, 0, 1}, {1, 0, 1}, {5, 1, 1, 8}, {1, 1, 1, 8}, {5, 2, 1, 8, 0x41, 0x0e}, {1, 2, 1, 8, 0x41, 0x0e}} {
		want, err := m.StandAloneMethodSignature(inner)
		if err != nil {
			t.Fatalf("inner positive control %x: %v", inner, err)
		}
		for name, decode := range methodSignatureDecoders() {
			for _, returned := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/%x/return-%v", name, inner, returned), func(t *testing.T) {
					data := []byte{5, 2, 1, 0x1b}
					fixed := 1
					if returned {
						data, fixed = []byte{5, 1, 0x1b}, 0
					}
					// Once the inner count is reached, this SENTINEL belongs to the outer method.
					data = append(append(data, inner...), 0x41, 8)
					got, err := decode(data)
					if err != nil || len(got.Param) != fixed || len(got.VariableParam) != 1 || got.VariableParam[0].Type.Kind != winmd.ElementType_I4 {
						t.Fatalf("outer signature %x = %+v, %v", data, got, err)
					}
					typ := got.RetType.Type
					if !returned {
						typ = got.Param[0].Type
					}
					if typ.Kind != winmd.ElementType_FNPTR || !reflect.DeepEqual(typ.Value, want) {
						t.Fatalf("inner signature = %#v; want %#v", typ, want)
					}
				})
			}
		}
	}
}

func TestFunctionPointerSignatureDepth(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, prefix := range [][]byte{{0x1b, 0, 0}, {0x1b, 0, 1, 1}, {0x0f}} {
		for _, layers := range []int{62, 63, 64} {
			// Include both FNPTR-only chains and pointers enclosing a FNPTR.
			typ := []byte{8}
			if prefix[0] == 0x0f {
				typ = []byte{0x1b, 0, 0, 8}
			}
			data := append(append([]byte{6}, bytes.Repeat(prefix, layers)...), typ...)
			valid := layers < 64 && (prefix[0] != 0x0f || layers < 63)
			if _, err := m.FieldSignature(data); (err == nil) != valid {
				t.Fatalf("prefix %x, layers %d: %v; want valid %v", prefix, layers, err, valid)
			}
		}
	}
	for _, overlong := range []int{-1, 0, 1, 2} {
		data := []byte{6, 0x1b, 0, 2}
		for slot := range 3 { // Return and both parameters inherit, but do not share, the budget.
			layers := 62
			if slot == overlong {
				layers++
			}
			data = append(data, bytes.Repeat([]byte{0x0f}, layers)...)
			data = append(data, 8)
		}
		if _, err := m.FieldSignature(data); (err == nil) != (overlong == -1) {
			t.Fatalf("overlong slot %d: %v", overlong, err)
		}
	}
}

func TestFunctionPointerSignatureHandleBounds(t *testing.T) {
	t.Parallel()
	tables := metadataTables(0, 1<<1|1<<2|1<<27, []uint32{1, 1, 1}, 6+14+2)
	m, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables})))
	if err != nil {
		t.Fatal(err)
	}
	for tag := range byte(3) {
		for _, row := range []byte{1, 2} {
			for _, prefix := range [][]byte{{6, 0x1b, 0, 0, 0x12}, {6, 0x1b, 0, 1, 1, 0x12}, {6, 0x1b, 0, 1, 1, 0x20}} {
				data := append(bytes.Clone(prefix), row<<2|tag)
				if prefix[len(prefix)-1] == 0x20 {
					data = append(data, 8)
				}
				if _, err := m.FieldSignature(data); (err == nil) != (row == 1) {
					t.Fatalf("table-backed FNPTR %x: %v; want valid %v", data, err, row == 1)
				}
			}
		}
	}
}

func TestMethodCallSignatureFromMetadata(t *testing.T) {
	t.Parallel()
	for _, handle := range []byte{4, 5, 6, 8, 9, 10} {
		t.Run(fmt.Sprint(handle), func(t *testing.T) {
			refData := []byte{5, 2, 0x12, handle, 0x10, 0x12, handle, 0x41, 0x16}
			callData := []byte{1, 1, 0x20, handle, 0x10, 0x12, handle, 0x41, 0x13, 0}
			var standalone winmd.Metadata
			wantRef, err := standalone.MethodRefSignature(refData)
			if err != nil {
				t.Fatalf("method reference control: %v", err)
			}
			wantCall, err := standalone.StandAloneMethodSignature(callData)
			if err != nil {
				t.Fatalf("standalone method control: %v", err)
			}
			blobs := append([]byte{0, byte(len(refData))}, refData...)
			callOffset := uint16(len(blobs))
			blobs = append(blobs, byte(len(callData)))
			blobs = append(blobs, callData...)
			// TypeRef and TypeDef rows supply handle bounds. MemberRef.Class
			// points to TypeDef row 1; a TypeSpec row supplies the third handle kind.
			tables := metadataTables(0, 1<<1|1<<2|1<<10|1<<17|1<<27, []uint32{1, 1, 1, 1, 1}, 6+14)
			for _, value := range []uint16{8, 0, 1, callOffset, 0} {
				tables = binary.LittleEndian.AppendUint16(tables, value)
			}
			m, err := winmd.New(metadataPE(t, metadataRoot(
				metadataStream{"#~", tables}, metadataStream{"#Strings", []byte{0}}, metadataStream{"#Blob", blobs},
			)))
			if err != nil {
				t.Fatal(err)
			}
			member, err := m.Tables.MemberRef.At(0)
			if err != nil || !bytes.Equal(member.Signature, refData) {
				t.Fatalf("MemberRef = %+v, %v", member, err)
			}
			call, err := m.Tables.StandAloneSig.At(0)
			if err != nil || !bytes.Equal(call.Signature, callData) {
				t.Fatalf("StandAloneSig = %+v, %v", call, err)
			}
			gotRef, refErr := m.MethodRefSignature(member.Signature)
			gotCall, callErr := m.StandAloneMethodSignature(call.Signature)
			if (refErr == nil) != (handle < 8) || (callErr == nil) != (handle < 8) {
				t.Fatalf("handle %d: reference error %v, standalone error %v", handle, refErr, callErr)
			}
			if handle < 8 && (!reflect.DeepEqual(gotRef, wantRef) || !reflect.DeepEqual(gotCall, wantCall)) {
				t.Fatalf("table-backed signatures changed: reference %+v, standalone %+v", gotRef, gotCall)
			}
		})
	}
}

func ExampleMetadata_MethodRefSignature() {
	var m winmd.Metadata
	// The count includes both the fixed I4 and optional STRING arguments.
	sig, err := m.MethodRefSignature([]byte{5, 2, 1, 8, 0x41, 0x0e})
	if err != nil {
		panic(err)
	}
	fmt.Println("fixed:", sig.Param[0].Type.Kind)
	fmt.Println("optional:", sig.VariableParam[0].Type.Kind)
	// Output:
	// fixed: I4
	// optional: STRING
}

func ExampleMetadata_StandAloneMethodSignature() {
	var m winmd.Metadata
	sig, err := m.StandAloneMethodSignature([]byte{1, 1, 1, 0x41, 8})
	if err != nil {
		panic(err)
	}
	fmt.Println("cdecl:", sig.CallingConvention == winmd.SigCallingConvention_Cdecl)
	fmt.Println("managed vararg:", sig.VarArgs)
	fmt.Println("optional arguments:", len(sig.VariableParam))
	// Output:
	// cdecl: true
	// managed vararg: false
	// optional arguments: 1
}

func FuzzMethodCallSignature(f *testing.F) {
	for _, data := range [][]byte{
		{}, {0, 0, 1}, {0x61, 1, 1, 0x41, 8}, {0x10, 1, 1, 0x1e, 0, 0x13, 0},
		{5, 2, 1, 0x1b, 5, 0, 1, 0x41, 8}, {5, 1, 1, 0x41, 0x41, 8},
		{5, 1, 0x20, 5, 0x16, 0x41, 0x10, 8}, {0, 0xdf, 0xff, 0xff, 0xff, 1},
	} {
		f.Add(data)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		before := bytes.Clone(data)
		for name, decode := range methodSignatureDecoders() {
			_, err := decode(data)
			if !bytes.Equal(data, before) {
				t.Fatalf("%s changed its input", name)
			}
			if err == nil {
				if _, err := decode(append(bytes.Clone(data), 0xff)); err == nil {
					t.Fatalf("%s accepted trailing data", name)
				}
			}
		}
	})
}
