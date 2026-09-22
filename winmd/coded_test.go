// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestParseCodedRows(t *testing.T) {
	t.Parallel()
	testParseCodedRows[TypeDefOrRef](t)
	testParseCodedRows[HasConstant](t)
	testParseCodedRows[HasFieldMarshal](t)
	testParseCodedRows[HasDeclSecurity](t)
	testParseCodedRows[MemberRefParent](t)
	testParseCodedRows[HasSemantics](t)
	testParseCodedRows[MethodDefOrRef](t)
	testParseCodedRows[MemberForwarded](t)
	testParseCodedRows[Implementation](t)
	testParseCodedRows[CustomAttributeType](t)
	testParseCodedRows[ResolutionScope](t)
	testParseCodedRows[TypeOrMethodDef](t)
	testParseCodedRows[HasCustomAttribute](t)
	testParseCodedRows[TypeDefOrRefOrSpec](t)
}

func testParseCodedRows[T CodedTag](t *testing.T) {
	t.Helper()
	var zero T
	t.Run(fmt.Sprintf("%T", zero), func(t *testing.T) {
		kind := zero.kind()
		tagBits := codedTagBits(kind)
		if got, err := parseCoded[T](0); err != nil || got.Tag != codedFromInt8[T](-1) || got.Index != 0 {
			t.Fatalf("null code = %+v, %v; want null, nil", got, err)
		}
		for code := uint32(1); code < 1<<tagBits; code++ {
			if got, err := parseCoded[T](code); err == nil {
				t.Errorf("zero-row code %d accepted: %+v", code, got)
			}
		}
		for tag := range 1 << tagBits {
			if tag >= len(codedMap[kind]) || codedMap[kind][tag] == tableNone {
				if got, err := parseCoded[T](1<<tagBits | uint32(tag)); err == nil {
					t.Errorf("reserved tag %d accepted: %+v", tag, got)
				}
				continue
			}
			for _, row := range []uint32{1, 2, 0xffffffff >> tagBits} {
				code := row<<tagBits | uint32(tag)
				got, err := parseCoded[T](code)
				if err != nil || uint32(got.Index) != row-1 || got.Tag != codedFromInt8[T](int8(tag)) {
					t.Errorf("code %#x = %+v, %v; want index %d, tag %d", code, got, err, row-1, tag)
				}
			}
		}
		testReadCodedBounds[T](t)
	})
}

func testReadCodedBounds[T CodedTag](t *testing.T) {
	t.Helper()
	var zero T
	kind := zero.kind()
	tagBits := codedTagBits(kind)
	for _, width := range []uint8{2, 4} {
		var la layout
		la.codedSizes[kind] = width
		r := ecma335Reader{data: make([]byte, width), layout: &la}
		if got := readCoded[T](&r); r.err != nil || got.Tag != codedFromInt8[T](-1) {
			t.Fatalf("null reference to absent tables = %+v, %v", got, r.err)
		}
		maxRow := uint32(0xffffffff) >> (32 - int(width)*8 + tagBits)
		for tag, tbl := range codedMap[kind] {
			if tbl == tableNone {
				continue
			}
			for _, test := range []struct {
				row, count uint32
			}{
				{1, 0}, {1, 1}, {2, 1}, {2, 2}, {maxRow, maxRow}, {maxRow, maxRow - 1},
			} {
				la = layout{}
				la.codedSizes[kind] = width
				la.tables[tbl].rowCount = test.count
				data := binary.LittleEndian.AppendUint32(nil, test.row<<tagBits|uint32(tag))[:width]
				r := ecma335Reader{data: data, layout: &la}
				got := readCoded[T](&r)
				wantErr := test.row > test.count
				if (r.err != nil) != wantErr || !wantErr && (uint32(got.Index) != test.row-1 || got.Tag != codedFromInt8[T](int8(tag))) {
					t.Errorf("width %d, tag %d, row %d, count %d: got %+v, %v; want error %v", width, tag, test.row, test.count, got, r.err, wantErr)
				}
			}
		}
	}
}

func TestHasCustomAttributeDeclSecurity(t *testing.T) {
	t.Parallel()
	if tbl, ok := codedTable(codedHasCustomAttribute, 8); !ok || tbl != tableDeclSecurity {
		t.Fatalf("HasCustomAttribute tag 8 = %d, %v; want DeclSecurity", tbl, ok)
	}
	index, err := parseCoded[HasCustomAttribute](1<<5 | 8)
	if err != nil || index.Tag != HasCustomAttribute_DeclSecurity || index.Index != 0 {
		t.Fatalf("DeclSecurity reference = %+v, %v", index, err)
	}
	for _, count := range []uint32{0, 2047, 2048, 65536} {
		var rows [tableMax]uint32
		rows[tableDeclSecurity] = count
		want := uint8(2)
		if count >= 2048 {
			want = 4
		}
		if got := codedIndexSize(codedHasCustomAttribute, rows); got != want {
			t.Errorf("DeclSecurity rows %d: HasCustomAttribute width = %d; want %d", count, got, want)
		}
	}
}
