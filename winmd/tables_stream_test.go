// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestNewTablesStreamRequired(t *testing.T) {
	t.Parallel()
	tables := metadataTables(0, 0, nil, 0)
	for _, test := range []struct {
		name      string
		streams   []metadataStream
		wantError string
	}{
		{"missing", nil, "missing metadata tables stream"},
		{"heaps-only", []metadataStream{{"#Strings", []byte{0}}, {"#Blob", []byte{0}}}, "missing metadata tables stream"},
		{"unknown-only", []metadataStream{{"#Other", tables}}, "missing metadata tables stream"},
		{"uncompressed", []metadataStream{{"#-", tables}}, "uncompressed metadata tables stream"},
		{"uncompressed-after-compressed", []metadataStream{{"#~", tables}, {"#-", tables}}, "uncompressed metadata tables stream"},
		{"uncompressed-before-compressed", []metadataStream{{"#-", tables}, {"#~", tables}}, "uncompressed metadata tables stream"},
		{"empty-stream", []metadataStream{{"#~", nil}}, "tables stream header"},
		{"empty-tables", []metadataStream{{"#~", tables}}, ""},
		{"unknown-with-tables", []metadataStream{{"#Other", []byte{1, 2, 3, 4}}, {"#~", tables}}, ""},
		{"duplicate-tables", []metadataStream{{"#~", tables}, {"#~", tables}}, "duplicated #~ stream"},
	} {
		t.Run(test.name, func(t *testing.T) {
			for _, pe64 := range []bool{false, true} {
				t.Run(fmt.Sprintf("pe64-%v", pe64), func(t *testing.T) {
					root := metadataRoot(test.streams...)
					m, err := winmd.New(metadataPEForArch(t, root, uint32(len(root)), pe64))
					if test.wantError != "" {
						if err == nil || m != nil || !strings.Contains(err.Error(), test.wantError) {
							t.Fatalf("New() = %+v, %v; want nil metadata and %q", m, err, test.wantError)
						}
						return
					}
					if err != nil || m == nil || m.Tables == nil || m.Tables.Module.Len() != 0 {
						t.Fatalf("empty tables = %+v, %v; want non-nil empty tables", m, err)
					}
					if _, err := m.EnumUnderlyingType(0); err == nil {
						t.Fatal("enum lookup in empty tables succeeded")
					}
				})
			}
		})
	}
}

func TestNewTablesExtraData(t *testing.T) {
	t.Parallel()
	for _, heapSizes := range []byte{0, 7} {
		t.Run(fmt.Sprintf("heap-sizes-%d", heapSizes), func(t *testing.T) {
			appendIndex := func(data []byte, index uint32) []byte {
				if heapSizes == 0 {
					return binary.LittleEndian.AppendUint16(data, uint16(index))
				}
				return binary.LittleEndian.AppendUint32(data, index)
			}
			for _, test := range []struct {
				name    string
				present bool
				value   uint32
			}{
				{"absent", false, 0},
				{"zero", true, 0},
				{"nonzero", true, 0x00020000},
				{"maximum", true, 0xffffffff},
			} {
				t.Run(test.name, func(t *testing.T) {
					flags := heapSizes
					if test.present {
						flags |= 0x40
					}
					tables := metadataTables(flags, 1<<0|1<<26, []uint32{1, 2}, 0)
					if test.present {
						tables = binary.LittleEndian.AppendUint32(tables, test.value)
					}
					// Module.Generation, then Name, Mvid, EncId, and EncBaseId.
					tables = binary.LittleEndian.AppendUint16(tables, 0)
					for _, index := range []uint32{1, 1, 0, 0} {
						tables = appendIndex(tables, index)
					}
					tables = appendIndex(tables, 8)  // ModuleRef row 1.
					tables = appendIndex(tables, 14) // ModuleRef row 2.
					m, err := winmd.New(metadataPE(t, metadataRoot(
						metadataStream{"#~", tables},
						metadataStream{"#Strings", []byte("\x00Module\x00First\x00Second\x00")},
						metadataStream{"#GUID", bytes.Repeat([]byte{1}, 16)},
					)))
					if err != nil {
						t.Fatal(err)
					}
					module, err := m.Tables.Module.At(0)
					if err != nil || module.Generation != 0 || module.Name.String() != "Module" || !bytes.Equal(module.Mvid[:], bytes.Repeat([]byte{1}, 16)) {
						t.Fatalf("first table is misaligned: %+v, %v", module, err)
					}
					for i, name := range []string{"First", "Second"} {
						ref, err := m.Tables.ModuleRef.At(winmd.Index(i))
						if err != nil || ref.Name.String() != name {
							t.Fatalf("following table row %d = %+v, %v; want %q", i, ref, err, name)
						}
					}
				})
			}
		})
	}
}

func TestNewTablesExtraDataBounds(t *testing.T) {
	t.Parallel()
	for _, rows := range []uint32{0, 1} {
		for length := 0; length <= 4; length++ {
			t.Run(fmt.Sprintf("rows-%d/extra-bytes-%d", rows, length), func(t *testing.T) {
				tables := metadataTables(0x40, 1, []uint32{rows}, 0)
				tables = append(tables, make([]byte, length)...)
				m, err := winmd.New(metadataPE(t, metadataRoot(
					metadataStream{"#~", tables},
					// Data after the tables stream must not satisfy a truncated DWORD.
					metadataStream{"#Other", make([]byte, 16)},
				)))
				if length == 4 && rows == 0 {
					if err != nil || m == nil || m.Tables == nil {
						t.Fatalf("empty tables with complete extra data = %+v, %v", m, err)
					}
					return
				}
				want := io.ErrUnexpectedEOF
				if length == 0 {
					want = io.EOF
				}
				if !errors.Is(err, want) || m != nil {
					t.Fatalf("truncated tables = %+v, %v; want nil metadata and %v", m, err, want)
				}
				if length < 4 && !strings.Contains(err.Error(), "tables stream extra data") {
					t.Fatalf("truncated extra data reported as a row error: %v", err)
				}
			})
		}
	}
}

func TestNewNullMethodDefParamLists(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name   string
		starts []uint16
		params uint32
		want   []winmd.Slice
	}{
		{"empty-null", []uint16{0}, 0, []winmd.Slice{{}}},
		{"empty-sentinel", []uint16{1}, 0, []winmd.Slice{{}}},
		{"null-first", []uint16{0, 1, 2}, 2, []winmd.Slice{{}, {Start: 0, End: 1}, {Start: 1, End: 2}}},
		{"null-middle", []uint16{1, 0, 2}, 2, []winmd.Slice{{Start: 0, End: 1}, {}, {Start: 1, End: 2}}},
		{"null-last", []uint16{1, 2, 0}, 2, []winmd.Slice{{Start: 0, End: 1}, {Start: 1, End: 2}, {}}},
		{"all-null", []uint16{0, 0}, 0, []winmd.Slice{{}, {}}},
	} {
		t.Run(test.name, func(t *testing.T) {
			tables := metadataTables(0, 1<<6|1<<8, []uint32{uint32(len(test.starts)), test.params}, 0)
			for _, start := range test.starts {
				tables = binary.LittleEndian.AppendUint32(tables, 0) // RVA.
				for _, value := range []uint16{
					uint16(winmd.MethodImplAttributes_Runtime),
					uint16(winmd.MethodAttributes_Public | winmd.MethodAttributes_Static),
					1, 1, start, // Name, signature, ParamList.
				} {
					tables = binary.LittleEndian.AppendUint16(tables, value)
				}
			}
			for range test.params {
				tables = append(tables, 0, 0, 1, 0, 6, 0) // Param: flags, sequence, name.
			}
			signature := []byte{0, 1, 1, 8} // A static method with one int32 parameter.
			if test.params == 0 {
				signature = []byte{0, 0, 1}
			}
			m, err := winmd.New(metadataPE(t, metadataRoot(
				metadataStream{"#~", tables},
				metadataStream{"#Strings", []byte("\x00Call\x00arg\x00")},
				metadataStream{"#Blob", append([]byte{0, byte(len(signature))}, signature...)},
			)))
			if err != nil {
				t.Fatal(err)
			}
			for i, want := range test.want {
				method, err := m.Tables.MethodDef.At(winmd.Index(i))
				if err != nil || method.ParamList != want || method.Name.String() != "Call" {
					t.Fatalf("method %d = %+v, %v; want ParamList %+v", i, method, err, want)
				}
				if _, err := m.MethodDefSignature(method.Signature); err != nil {
					t.Fatal(err)
				}
				for row := range method.ParamList.All() {
					param, err := m.Tables.Param.At(row)
					if err != nil || param.Name.String() != "arg" || param.Sequence != 1 {
						t.Fatalf("parameter %d = %+v, %v", row, param, err)
					}
				}
			}
		})
	}
}
