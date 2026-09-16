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

func TestFieldSignatureHeader(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for high := range byte(16) {
		t.Run(fmt.Sprintf("%#x", high<<4|6), func(t *testing.T) {
			sig, err := m.FieldSignature([]byte{high<<4 | 6, 8})
			if high != 0 {
				if err == nil {
					t.Fatal("nonzero high bits accepted in FIELD signature header")
				}
			} else if err != nil || sig.Type.Kind != winmd.ElementType_I4 {
				t.Fatalf("valid field signature = %+v, %v; want I4, nil", sig, err)
			}
		})
	}
}

func TestMethodDefSignatureFlags(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, test := range []struct {
		name         string
		data         []byte
		varArgs      bool
		hasThis      bool
		explicitThis bool
		generic      uint32
	}{
		{"default", []byte{0, 0, 1}, false, false, false, 0},
		{"vararg", []byte{5, 0, 1}, true, false, false, 0},
		{"instance", []byte{0x20, 0, 1}, false, true, false, 0},
		{"instance-vararg", []byte{0x25, 0, 1}, true, true, false, 0},
		{"explicit-instance", []byte{0x60, 0, 1}, false, true, true, 0},
		{"generic", []byte{0x10, 1, 0, 1}, false, false, false, 1},
		{"generic-instance", []byte{0x30, 2, 0, 1}, false, true, false, 2},
	} {
		t.Run(test.name, func(t *testing.T) {
			sig, err := m.MethodDefSignature(test.data)
			if err != nil {
				t.Fatal(err)
			}
			if sig.VarArgs != test.varArgs || sig.HasThis != test.hasThis || sig.ExplicitThis != test.explicitThis || sig.Generic != test.generic {
				t.Fatalf("signature flags = %+v; want varargs %v, has-this %v, explicit-this %v, generic %d", sig, test.varArgs, test.hasThis, test.explicitThis, test.generic)
			}
		})
	}
}

func TestMethodDefSignatureInvalidHeaders(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, test := range []struct {
		name string
		data []byte
	}{
		{"reserved-bit", []byte{0x80, 0, 1}},
		{"reserved-instance", []byte{0xa0, 0, 1}},
		{"explicit-without-instance", []byte{0x40, 0, 1}},
		{"explicit-vararg-without-instance", []byte{0x45, 0, 1}},
		{"zero-generic-parameters", []byte{0x10, 0, 0, 1}},
		{"instance-zero-generic-parameters", []byte{0x30, 0, 0, 1}},
		{"generic-vararg", []byte{0x15, 1, 0, 1}},
		{"instance-generic-vararg", []byte{0x35, 1, 0, 1}},
		{"generic-vararg-zero-count", []byte{0x15, 0, 0, 1}},
		{"generic-vararg-missing-count", []byte{0x15}},
		{"unmanaged-cdecl", []byte{1, 0, 1}},
		{"unmanaged-stdcall", []byte{2, 0, 1}},
		{"unmanaged-thiscall", []byte{3, 0, 1}},
		{"unmanaged-fastcall", []byte{4, 0, 1}},
	} {
		t.Run(test.name, func(t *testing.T) {
			if sig, err := m.MethodDefSignature(test.data); err == nil {
				t.Fatalf("invalid method header accepted: %+v", sig)
			}
		})
	}
}

func TestSignatureTypeNesting(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	// The documented limit counts the leaf type as one of the 64 levels.
	const maxDepth = 64
	for _, kind := range []string{"pointer", "array", "szarray", "mixed"} {
		t.Run(kind, func(t *testing.T) {
			for _, layers := range []int{0, maxDepth - 1, maxDepth, maxDepth + 1, 4096} {
				t.Run(fmt.Sprint(layers), func(t *testing.T) {
					data := []byte{6}
					arrays := 0
					for i := range layers {
						if kind == "array" || kind == "mixed" && i%3 == 0 {
							data = append(data, 0x14)
							arrays++
						} else if kind == "szarray" || kind == "mixed" && i%3 == 1 {
							data = append(data, 0x1d)
						} else {
							data = append(data, 0x0f)
						}
					}
					data = append(data, 8)
					for range arrays {
						data = append(data, 1, 0, 0) // Rank 1, unspecified bounds.
					}
					_, err := m.FieldSignature(data)
					if layers >= maxDepth {
						if err == nil || !strings.Contains(err.Error(), "nesting limit") {
							t.Fatalf("depth %d error = %v; want nesting limit error", layers+1, err)
						}
					} else if err != nil {
						t.Fatalf("valid depth %d rejected: %v", layers+1, err)
					}
				})
			}
		})
	}
	t.Run("separate-return-and-parameters", func(t *testing.T) {
		typ := append(bytes.Repeat([]byte{0x0f}, maxDepth-1), 8)
		data := []byte{0, 2}
		for range 3 { // Return type and two parameters each get their own limit.
			data = append(data, typ...)
		}
		if _, err := m.MethodDefSignature(data); err != nil {
			t.Fatalf("independent type nesting rejected: %v", err)
		}
	})
	t.Run("byref-counts-as-a-level", func(t *testing.T) {
		for _, layers := range []int{maxDepth - 2, maxDepth - 1} {
			data := append([]byte{0, 0, 0x10}, bytes.Repeat([]byte{0x0f}, layers)...)
			_, err := m.MethodDefSignature(append(data, 8))
			if (err != nil) != (layers == maxDepth-1) {
				t.Fatalf("BYREF with %d pointer levels: %v", layers, err)
			}
		}
	})
	t.Run("flat-modifiers-do-not-add-depth", func(t *testing.T) {
		data := append([]byte{6}, bytes.Repeat([]byte{0x20, 5}, 128)...)
		sig, err := m.FieldSignature(append(data, 8))
		if err != nil || len(sig.Type.Mod) != 128 {
			t.Fatalf("flat modifier sequence: modifier count %d, error %v", len(sig.Type.Mod), err)
		}
	})
}

func TestMethodDefSignatureReturnKind(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, test := range []struct {
		name string
		data []byte
		kind winmd.SigRetTypeKind
	}{
		{"value", []byte{8}, winmd.SigRetTypeKind_ByValue},
		{"byref", []byte{0x10, 8}, winmd.SigRetTypeKind_ByRef},
		{"typedbyref", []byte{0x16}, winmd.SigRetTypeKind_TypedByRef},
		{"void", []byte{1}, winmd.SigRetTypeKind_Void},
	} {
		t.Run(test.name, func(t *testing.T) {
			sig, err := m.MethodDefSignature(append([]byte{0, 0}, test.data...))
			if err != nil || sig.RetType.Kind != test.kind {
				t.Fatalf("return kind = %v, %v; want %v, nil", sig.RetType.Kind, err, test.kind)
			}
		})
	}
}

func TestSignaturesTruncated(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, test := range []struct {
		name  string
		data  []byte
		field bool
	}{
		{"field", []byte{6, 8}, true},
		{"pointer-field", []byte{6, 0x0f, 8}, true},
		{"type-handle", []byte{6, 0x11, 0x80, 0x81}, true},
		{"modifier", []byte{6, 0x20, 5, 8}, true},
		{"array", []byte{6, 0x14, 8, 1, 1, 2, 1, 0}, true},
		{"szarray", []byte{6, 0x1d, 8}, true},
		{"szarray-modifier", []byte{6, 0x1d, 0x20, 5, 8}, true},
		{"szarray-type-handle", []byte{6, 0x1d, 0x12, 0x80, 0x81}, true},
		{"method", []byte{0, 1, 1, 8}, false},
		{"szarray-method", []byte{0, 1, 1, 0x1d, 8}, false},
		{"generic-method", []byte{0x10, 1, 0, 1}, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			for length := 0; length <= len(test.data); length++ {
				t.Run(fmt.Sprint(length), func(t *testing.T) {
					var err error
					if test.field {
						_, err = m.FieldSignature(test.data[:length])
					} else {
						_, err = m.MethodDefSignature(test.data[:length])
					}
					if length == len(test.data) {
						if err != nil {
							t.Fatalf("complete signature failed: %v", err)
						}
					} else if !errors.Is(err, io.ErrUnexpectedEOF) {
						t.Fatalf("truncated signature error = %v; want %v", err, io.ErrUnexpectedEOF)
					}
				})
			}
		})
	}
}

func TestFieldSignatureZeroRowTypeHandle(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, code := range []byte{0, 1, 2, 3, 4, 5, 6} {
		t.Run(fmt.Sprint(code), func(t *testing.T) {
			sig, err := m.FieldSignature([]byte{6, 0x11, code})
			if code < 4 {
				if err == nil {
					t.Fatalf("zero-row type handle accepted: %+v", sig.Type.Value)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if index := sig.Type.Value.(winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]).Index; index != 0 {
				t.Fatalf("decoded type handle row = %d; want 0", index)
			}
		})
	}
}

func TestSignatureNullTypeHandles(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, data := range [][]byte{
		{6, 0x11, 0},
		{6, 0x12, 0},
		{6, 0x20, 0, 8},
		{6, 0x1f, 0, 8},
		{6, 0x0f, 0x12, 0},
		{6, 0x1d, 0x12, 0},
		{6, 0x1d, 0x20, 0, 8},
	} {
		t.Run(fmt.Sprintf("%x", data), func(t *testing.T) {
			_, err := m.FieldSignature(data)
			if err == nil || !strings.Contains(err.Error(), "must not be null") {
				t.Fatalf("null type handle error = %v; want null-handle error", err)
			}
		})
	}
	t.Run("parameter", func(t *testing.T) {
		if _, err := m.MethodDefSignature([]byte{0, 1, 1, 0x12, 0}); err == nil {
			t.Fatal("null parameter type handle accepted")
		}
	})
}

func TestSignatureTypeHandleBounds(t *testing.T) {
	t.Parallel()
	for _, target := range []struct {
		name    string
		tag     byte
		rowSlot int
	}{
		{"TypeDef", 0, 1},
		{"TypeRef", 1, 0},
		{"TypeSpec", 2, 2},
	} {
		for _, count := range []uint32{0, 1} {
			t.Run(fmt.Sprintf("%s/rows-%d", target.name, count), func(t *testing.T) {
				rows := []uint32{1, 1, 1} // TypeRef, TypeDef, TypeSpec.
				rows[target.rowSlot] = count
				tables := metadataTables(0, 1<<1|1<<2|1<<27, rows, int(6*rows[0]+14*rows[1]+2*rows[2]))
				m, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables})))
				if err != nil {
					t.Fatal(err)
				}
				for _, kind := range []winmd.ElementType{
					winmd.ElementType_CLASS, winmd.ElementType_VALUETYPE,
					winmd.ElementType_CMOD_OPT, winmd.ElementType_CMOD_REQD,
				} {
					for _, handle := range []struct {
						name  string
						data  []byte
						valid bool
					}{
						{"first-row", []byte{4 | target.tag}, count != 0},
						{"past-end", []byte{8 | target.tag}, false},
						{"largest-row", []byte{0xdf, 0xff, 0xff, 0xfc | target.tag}, false},
					} {
						t.Run(kind.String()+"/"+handle.name, func(t *testing.T) {
							typ := append([]byte{byte(kind)}, handle.data...)
							if kind == winmd.ElementType_CMOD_OPT || kind == winmd.ElementType_CMOD_REQD {
								typ = append(typ, byte(winmd.ElementType_I4))
							}
							for _, context := range []string{"field", "return", "parameter", "szarray-field", "szarray-return", "szarray-parameter"} {
								t.Run(context, func(t *testing.T) {
									var err error
									switch context {
									case "field":
										_, err = m.FieldSignature(append([]byte{6}, typ...))
									case "return":
										_, err = m.MethodDefSignature(append([]byte{0, 0}, typ...))
									case "parameter":
										_, err = m.MethodDefSignature(append([]byte{0, 1, 1}, typ...))
									case "szarray-field":
										_, err = m.FieldSignature(append([]byte{6, 0x1d}, typ...))
									case "szarray-return":
										_, err = m.MethodDefSignature(append([]byte{0, 0, 0x1d}, typ...))
									case "szarray-parameter":
										_, err = m.MethodDefSignature(append([]byte{0, 1, 1, 0x1d}, typ...))
									}
									if handle.valid {
										if err != nil {
											t.Fatalf("valid target rejected: %v", err)
										}
									} else if err == nil || !strings.Contains(err.Error(), "beyond the end of table") {
										t.Fatalf("out-of-range target error = %v; want table bounds error", err)
									}
								})
							}
						})
					}
				}
				if _, err := m.FieldSignature([]byte{6, 0x12, 0x80}); !errors.Is(err, io.ErrUnexpectedEOF) {
					t.Fatalf("truncated handle error = %v; want unexpected EOF", err)
				}
			})
		}
	}
}

func TestSignatureCustomModifierOrder(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	prefix := []byte{0x20, 5, 0x1f, 9, 0x20, 13}
	want := []winmd.SigCustomMod{
		{Kind: winmd.SigCustomModKind_Opt, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 0}},
		{Kind: winmd.SigCustomModKind_Reqd, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 1}},
		{Kind: winmd.SigCustomModKind_Opt, Index: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 2}},
	}
	check := func(t *testing.T, got []winmd.SigCustomMod) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("modifiers = %+v; want encoded order %+v", got, want)
		}
	}
	t.Run("field", func(t *testing.T) {
		data := append([]byte{6}, prefix...)
		sig, err := m.FieldSignature(append(data, 8))
		if err != nil {
			t.Fatal(err)
		}
		check(t, sig.Type.Mod)
	})
	t.Run("pointer-target", func(t *testing.T) {
		data := append([]byte{6}, prefix...)
		data = append(data, 0x0f)
		data = append(data, prefix...)
		sig, err := m.FieldSignature(append(data, 1))
		if err != nil {
			t.Fatal(err)
		}
		check(t, sig.Type.Mod)
		check(t, sig.Type.Value.(winmd.SigType).Mod)
	})
	t.Run("szarray-element", func(t *testing.T) {
		data := append([]byte{6}, prefix...)
		data = append(data, 0x1d)
		data = append(data, prefix...)
		sig, err := m.FieldSignature(append(data, 8))
		if err != nil {
			t.Fatal(err)
		}
		check(t, sig.Type.Mod)
		check(t, sig.Type.Value.(winmd.SigType).Mod)
	})
	t.Run("return-and-parameter", func(t *testing.T) {
		data := append([]byte{0, 1}, prefix...)
		data = append(data, 1)
		data = append(data, prefix...)
		sig, err := m.MethodDefSignature(append(data, 0x10, 8))
		if err != nil {
			t.Fatal(err)
		}
		check(t, sig.RetType.Type.Mod)
		check(t, sig.Param[0].Type.Mod)
		if sig.Param[0].Kind != winmd.SigParamKind_ByRef || sig.Param[0].Type.Value.(winmd.SigType).Kind != winmd.ElementType_I4 {
			t.Fatalf("by-reference parameter representation changed: %+v", sig.Param[0])
		}
	})
}

func TestSignatureSZArray(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	i4 := winmd.SigType{Kind: winmd.ElementType_I4}
	for _, test := range []struct {
		name    string
		element []byte
		want    winmd.SigType
	}{
		{"primitive", []byte{8}, i4},
		{"class", []byte{0x12, 5}, winmd.SigType{Kind: winmd.ElementType_CLASS,
			Value: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: 0}}},
		{"value-type", []byte{0x11, 4}, winmd.SigType{Kind: winmd.ElementType_VALUETYPE,
			Value: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeDef, Index: 0}}},
		{"jagged", []byte{0x1d, 8}, winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: i4}},
		{"pointer", []byte{0x0f, 1}, winmd.SigType{Kind: winmd.ElementType_PTR, Value: winmd.SigType{Kind: winmd.ElementType_VOID}}},
		{"multidimensional", []byte{0x14, 8, 2, 0, 0}, winmd.SigType{Kind: winmd.ElementType_ARRAY,
			Value: winmd.SigArray{Type: i4, Rank: 2, Sizes: []uint32{}, LowerBounds: []int32{}}}},
		{"type-parameter", []byte{0x13, 0}, winmd.SigType{Kind: winmd.ElementType_VAR, Value: uint32(0)}},
		{"method-parameter", []byte{0x1e, 1}, winmd.SigType{Kind: winmd.ElementType_MVAR, Value: uint32(1)}},
	} {
		t.Run(test.name, func(t *testing.T) {
			want := winmd.SigType{Kind: winmd.ElementType_SZARRAY, Value: test.want}
			field, err := m.FieldSignature(append([]byte{6, 0x1d}, test.element...))
			if err != nil || !reflect.DeepEqual(field.Type, want) {
				t.Fatalf("field type = %+v, %v; want %+v", field.Type, err, want)
			}
			// Return and parameter types must remain distinct vectors too.
			data := append([]byte{0, 1, 0x1d}, test.element...)
			data = append(data, 0x1d)
			data = append(data, test.element...)
			method, err := m.MethodDefSignature(data)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(method.RetType.Type, want) || len(method.Param) != 1 || !reflect.DeepEqual(method.Param[0].Type, want) {
				t.Fatalf("method types = %+v; want SZARRAY return and parameter %+v", method, want)
			}
		})
	}
}

func TestSignatureArrayShapes(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, test := range []struct {
		name      string
		shape     []byte
		rank      uint32
		sizes     []uint32
		bounds    []int32
		wantError bool
	}{
		{"unspecified", []byte{1, 0, 0}, 1, []uint32{}, []int32{}, false},
		{"partial-dimensions", []byte{3, 1, 4, 2, 0x7f, 0}, 3, []uint32{4}, []int32{-1, 0}, false},
		{"zero-sized-dimension", []byte{1, 1, 0, 1, 0}, 1, []uint32{0}, []int32{0}, false},
		{"zero-rank", []byte{0, 0, 0}, 0, nil, nil, true},
		{"too-many-sizes", []byte{1, 2, 3, 4, 0}, 0, nil, nil, true},
		{"too-many-bounds", []byte{1, 0, 2, 0, 0}, 0, nil, nil, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			sig, err := m.FieldSignature(append([]byte{6, 0x14, 8}, test.shape...))
			if test.wantError {
				if err == nil {
					t.Fatalf("invalid shape accepted: %+v", sig.Type.Value)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			array := sig.Type.Value.(winmd.SigArray)
			if array.Rank != test.rank || !reflect.DeepEqual(array.Sizes, test.sizes) || !reflect.DeepEqual(array.LowerBounds, test.bounds) {
				t.Fatalf("decoded array = %+v; want rank %d, sizes %v, bounds %v", array, test.rank, test.sizes, test.bounds)
			}
		})
	}
}

func TestSignatureArrayCountsBounded(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	// Keep the regression bounded even against an old decoder: the count is
	// 65536, not the maximum compressed integer. Rank equals the count so
	// rejecting counts above rank alone is insufficient.
	for _, test := range []struct {
		name string
		data []byte
	}{
		{"sizes", []byte{6, 0x14, 8, 0xc0, 1, 0, 0, 0xc0, 1, 0, 0}},
		{"bounds", []byte{6, 0x14, 8, 0xc0, 1, 0, 0, 0, 0xc0, 1, 0, 0}},
	} {
		t.Run(test.name, func(t *testing.T) {
			sig, err := m.FieldSignature(test.data)
			if !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatalf("truncated shape error = %v; want %v", err, io.ErrUnexpectedEOF)
			}
			if array, ok := sig.Type.Value.(winmd.SigArray); ok && (cap(array.Sizes) != 0 || cap(array.LowerBounds) != 0) {
				t.Fatalf("truncated shape allocated slices with capacities %d, %d", cap(array.Sizes), cap(array.LowerBounds))
			}
		})
	}
}

func TestSignatureElementTypeCodes(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, data := range [][]byte{
		{6, 0x81, 0x08},
		{6, 0x81, 0x12, 5},
		{6, 0xc0, 0, 1, 8},
		{6, 0x81, 0x20, 5, 8},
	} {
		t.Run(fmt.Sprintf("%x", data), func(t *testing.T) {
			if sig, err := m.FieldSignature(data); err == nil {
				t.Fatalf("invalid element code narrowed to a valid type: %+v", sig)
			}
		})
	}
}

func TestSignaturesTrailingData(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, suffix := range []byte{0, 8, 0xff} {
		t.Run(fmt.Sprintf("%#x", suffix), func(t *testing.T) {
			_, fieldErr := m.FieldSignature([]byte{6, 8, suffix})
			_, methodErr := m.MethodDefSignature([]byte{0, 0, 1, suffix})
			for _, err := range []error{fieldErr, methodErr} {
				if err == nil || !strings.Contains(err.Error(), "trailing") {
					t.Fatalf("signature error = %v; want trailing-data error", err)
				}
			}
		})
	}
}

func TestSignatureTypeContexts(t *testing.T) {
	t.Parallel()
	var m winmd.Metadata
	for _, test := range []struct {
		name    string
		data    []byte
		method  bool
		wantErr bool
	}{
		{"void-pointer-field", []byte{6, 0x0f, 1}, false, false},
		{"nested-void-pointer", []byte{6, 0x0f, 0x0f, 1}, false, false},
		{"nested-array", []byte{6, 0x14, 0x14, 8, 1, 0, 0, 1, 0, 0}, false, false},
		{"void-pointer-array", []byte{6, 0x14, 0x0f, 1, 1, 0, 0}, false, false},
		{"szarray-field", []byte{6, 0x1d, 8}, false, false},
		{"szarray-pointer", []byte{6, 0x0f, 0x1d, 8}, false, false},
		{"typed-reference-parameter", []byte{0, 1, 1, 0x16}, true, false},
		{"byref-parameter", []byte{0, 1, 1, 0x10, 8}, true, false},
		{"byref-pointer-return", []byte{0, 0, 0x10, 0x0f, 1}, true, false},
		{"byref-szarray-parameter", []byte{0, 1, 1, 0x10, 0x1d, 8}, true, false},
		{"byref-szarray-return", []byte{0, 0, 0x10, 0x1d, 8}, true, false},
		{"void-field", []byte{6, 1}, false, true},
		{"byref-field", []byte{6, 0x10, 8}, false, true},
		{"typed-reference-field", []byte{6, 0x16}, false, true},
		{"void-parameter", []byte{0, 1, 1, 1}, true, true},
		{"byref-void-return", []byte{0, 0, 0x10, 1}, true, true},
		{"byref-typed-reference", []byte{0, 0, 0x10, 0x16}, true, true},
		{"nested-byref", []byte{0, 0, 0x10, 0x10, 8}, true, true},
		{"pointer-to-byref", []byte{6, 0x0f, 0x10, 8}, false, true},
		{"pointer-to-typed-reference", []byte{6, 0x0f, 0x16}, false, true},
		{"array-of-void", []byte{6, 0x14, 1, 1, 0, 0}, false, true},
		{"array-of-byref", []byte{6, 0x14, 0x10, 8, 1, 0, 0}, false, true},
		{"array-of-typed-reference", []byte{6, 0x14, 0x16, 1, 0, 0}, false, true},
		{"szarray-of-void", []byte{6, 0x1d, 1}, false, true},
		{"szarray-of-byref", []byte{6, 0x1d, 0x10, 8}, false, true},
		{"szarray-of-typed-reference", []byte{6, 0x1d, 0x16}, false, true},
		{"modified-szarray-of-void", []byte{6, 0x1d, 0x20, 5, 1}, false, true},
		{"szarray-with-shape", []byte{6, 0x1d, 8, 1, 0, 0}, false, true},
		{"modifier-after-byref", []byte{0, 1, 1, 0x10, 0x20, 5, 8}, true, true},
		{"modifier-on-array-element", []byte{6, 0x14, 0x20, 5, 8, 1, 0, 0}, false, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			var err error
			if test.method {
				_, err = m.MethodDefSignature(test.data)
			} else {
				_, err = m.FieldSignature(test.data)
			}
			if (err != nil) != test.wantErr {
				t.Fatalf("signature %x error = %v; want error %v", test.data, err, test.wantErr)
			}
		})
	}
}
