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
	testParseCodedRows[TypeDefOrRef](t, true)
	testParseCodedRows[HasConstant](t, false)
	testParseCodedRows[HasFieldMarshal](t, false)
	testParseCodedRows[HasDeclSecurity](t, false)
	testParseCodedRows[MemberRefParent](t, false)
	testParseCodedRows[HasSemantics](t, false)
	testParseCodedRows[MethodDefOrRef](t, false)
	testParseCodedRows[MemberForwarded](t, false)
	testParseCodedRows[Implementation](t, true)
	testParseCodedRows[CustomAttributeType](t, false)
	testParseCodedRows[ResolutionScope](t, true)
	testParseCodedRows[TypeOrMethodDef](t, false)
	testParseCodedRows[HasCustomAttribute](t, false)
	testParseCodedRows[TypeDefOrRefOrSpec](t, false)
}

func testParseCodedRows[T CodedTag](t *testing.T, nullable bool) {
	t.Helper()
	var zero T
	t.Run(fmt.Sprintf("%T", zero), func(t *testing.T) {
		kind := zero.kind()
		tagBits := codedTagBits(kind)
		got, err := parseCoded[T](0)
		if nullable {
			if err != nil || got.Tag != codedFromInt8[T](-1) || got.Index != 0 {
				t.Fatalf("null code = %+v, %v; want null, nil", got, err)
			}
		} else if err == nil {
			t.Fatalf("null code accepted: %+v", got)
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
		testReadCodedBounds[T](t, nullable)
	})
}

func testReadCodedBounds[T CodedTag](t *testing.T, nullable bool) {
	t.Helper()
	var zero T
	kind := zero.kind()
	tagBits := codedTagBits(kind)
	for _, width := range []uint8{2, 4} {
		var la layout
		la.codedSizes[kind] = width
		for _, allowNull := range []bool{false, true} {
			r := ecma335Reader{data: make([]byte, width), layout: &la}
			got := readCoded[T](&r, allowNull)
			if nullable && allowNull {
				if r.err != nil || got.Tag != codedFromInt8[T](-1) {
					t.Fatalf("null reference to absent tables = %+v, %v", got, r.err)
				}
			} else if r.err == nil {
				t.Fatalf("null reference accepted with nullable %v: %+v", allowNull, got)
			}
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
				got := readCoded[T](&r, nullable)
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

func TestTableAtCodedIndexNullability(t *testing.T) {
	t.Parallel()
	// One row per table, with valid non-null references. Values are written as
	// uint16 words; four-byte columns occupy two words. Signatures are not read.
	rows := [tableMax][]uint16{
		0x00: {0, 0, 0, 0, 0},                // Module.
		0x01: {4, 0, 0},                      // TypeRef: ResolutionScope -> Module.
		0x02: {0, 0, 0, 0, 5, 1, 1},          // TypeDef: Extends -> TypeRef.
		0x04: {0, 0, 0},                      // Field.
		0x06: {0, 0, 0, 0, 0, 0, 1},          // MethodDef, with an empty ParamList.
		0x09: {1, 4},                         // InterfaceImpl: Class and Interface -> TypeDef.
		0x0a: {8, 0, 0},                      // MemberRef: Class -> TypeDef.
		0x0b: {uint16(ElementType_I4), 4, 0}, // Constant: Parent -> Field.
		0x0c: {35, 10, 0},                    // CustomAttribute: Parent -> TypeDef, Type -> MethodDef.
		0x0d: {2, 0},                         // FieldMarshal: Parent -> Field.
		0x0e: {0, 4, 0},                      // DeclSecurity: Parent -> TypeDef.
		0x14: {0, 0, 4},                      // Event: EventType -> TypeDef.
		0x18: {8, 1, 2},                      // MethodSemantics: Method -> MethodDef, Association -> Event.
		0x19: {1, 2, 2},                      // MethodImpl: Class -> TypeDef, methods -> MethodDef.
		0x1a: {0},                            // ModuleRef.
		0x1c: {0, 3, 0, 1},                   // ImplMap: MemberForwarded -> MethodDef, ImportScope -> ModuleRef.
		0x26: {0, 0, 0, 0},                   // File.
		0x27: {0, 0, 0, 0, 0, 0, 4},          // ExportedType: Implementation -> File.
		0x28: {0, 0, 0, 0, 0, 4},             // ManifestResource: Implementation -> File.
		0x2a: {0, 0, 2, 0},                   // GenericParam: Owner -> TypeDef.
		0x2b: {2, 0},                         // MethodSpec: Method -> MethodDef.
		0x2c: {1, 4},                         // GenericParamConstraint: Owner -> GenericParam, Constraint -> TypeDef.
	}
	for _, test := range []struct {
		name     string
		table    int
		column   int
		nullable bool
		read     func(*Tables) error
	}{
		{"Constant.Parent", 0x0b, 1, false, func(t *Tables) error { return readFirstRow(t.Constant) }},
		{"CustomAttribute.Parent", 0x0c, 0, false, func(t *Tables) error { return readFirstRow(t.CustomAttribute) }},
		{"CustomAttribute.Type", 0x0c, 1, false, func(t *Tables) error { return readFirstRow(t.CustomAttribute) }},
		{"DeclSecurity.Parent", 0x0e, 1, false, func(t *Tables) error { return readFirstRow(t.DeclSecurity) }},
		{"Event.EventType", 0x14, 2, true, func(t *Tables) error { return readFirstRow(t.Event) }},
		{"ExportedType.Implementation", 0x27, 6, false, func(t *Tables) error { return readFirstRow(t.ExportedType) }},
		{"FieldMarshal.Parent", 0x0d, 0, false, func(t *Tables) error { return readFirstRow(t.FieldMarshal) }},
		{"GenericParam.Owner", 0x2a, 2, false, func(t *Tables) error { return readFirstRow(t.GenericParam) }},
		{"GenericParamConstraint.Constraint", 0x2c, 1, false, func(t *Tables) error { return readFirstRow(t.GenericParamConstraint) }},
		{"ImplMap.MemberForwarded", 0x1c, 1, false, func(t *Tables) error { return readFirstRow(t.ImplMap) }},
		{"InterfaceImpl.Interface", 0x09, 1, false, func(t *Tables) error { return readFirstRow(t.InterfaceImpl) }},
		{"ManifestResource.Implementation", 0x28, 5, true, func(t *Tables) error { return readFirstRow(t.ManifestResource) }},
		{"MemberRef.Class", 0x0a, 0, false, func(t *Tables) error { return readFirstRow(t.MemberRef) }},
		{"MethodImpl.MethodBody", 0x19, 1, false, func(t *Tables) error { return readFirstRow(t.MethodImpl) }},
		{"MethodImpl.MethodDeclaration", 0x19, 2, false, func(t *Tables) error { return readFirstRow(t.MethodImpl) }},
		{"MethodSemantics.Association", 0x18, 2, false, func(t *Tables) error { return readFirstRow(t.MethodSemantics) }},
		{"MethodSpec.Method", 0x2b, 0, false, func(t *Tables) error { return readFirstRow(t.MethodSpec) }},
		{"TypeDef.Extends", 0x02, 4, true, func(t *Tables) error { return readFirstRow(t.TypeDef) }},
		{"TypeRef.ResolutionScope", 0x01, 0, true, func(t *Tables) error { return readFirstRow(t.TypeRef) }},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Check the positive control before replacing just this column with zero.
			for _, null := range []bool{false, true} {
				t.Run(fmt.Sprintf("null-%v", null), func(t *testing.T) {
					var counts [tableMax]uint32
					var data []byte
					for table, row := range rows {
						if row == nil {
							continue
						}
						counts[table] = 1
						for column, value := range row {
							if null && table == test.table && column == test.column {
								value = 0
							}
							data = binary.LittleEndian.AppendUint16(data, value)
						}
					}
					layout, err := generateLayout(0, counts, len(data))
					if err != nil {
						t.Fatal(err)
					}
					tables := newTables(data, &heaps{strs: StringHeap{0}}, layout)
					wantErr := null && !test.nullable
					if err := test.read(tables); (err != nil) != wantErr {
						t.Fatalf("At() error = %v; want error %v", err, wantErr)
					}
				})
			}
		})
	}
}

func readFirstRow[T any](table Table[T]) error {
	_, err := table.At(0)
	return err
}
