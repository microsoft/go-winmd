// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"slices"
	"strings"
	"testing"
)

func TestECMA335ReaderIntegers(t *testing.T) {
	t.Parallel()
	for _, width := range []uint8{1, 2, 4} {
		t.Run(fmt.Sprint(width), func(t *testing.T) {
			for length := 0; length <= int(width); length++ {
				data := bytes.Repeat([]byte{1}, length)
				r := ecma335Reader{data: data}
				got := r.uint(width)
				if length < int(width) {
					if got != 0 || !errors.Is(r.err, io.ErrUnexpectedEOF) || len(r.data) != length {
						t.Fatalf("length %d: got %d, %v, remaining %d; want 0, unexpected EOF, remaining %d", length, got, r.err, len(r.data), length)
					}
				} else {
					want := uint32(0x01010101) >> (32 - 8*width)
					if got != want || r.err != nil || len(r.data) != 0 {
						t.Fatalf("complete integer = %#x, %v, remaining %d; want %#x, nil, remaining 0", got, r.err, len(r.data), want)
					}
				}
			}
			wantErr := errors.New("earlier error")
			r := ecma335Reader{data: make([]byte, width), err: wantErr}
			if got := r.uint(width); got != 0 || r.err != wantErr || len(r.data) != int(width) {
				t.Fatalf("read after error = %d, %v, remaining %d", got, r.err, len(r.data))
			}
		})
	}
}

func TestTableIndexBounds(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name  string
		count uint32
		code  uint32
	}{
		{"empty-zero", 0, 0},
		{"empty-sentinel", 0, 1},
		{"empty-out-of-range", 0, 2},
		{"zero", 2, 0},
		{"first", 2, 1},
		{"last", 2, 2},
		{"sentinel", 2, 3},
		{"out-of-range", 2, 4},
		{"wide-last", 65536, 65536},
		{"wide-sentinel", 65536, 65537},
		{"wide-out-of-range", 65536, 65538},
		{"max-row", 0xffffffff, 0xffffffff},
		{"max-sentinel", 0xfffffffe, 0xffffffff},
	} {
		t.Run(test.name, func(t *testing.T) {
			var la layout
			la.tables[tableTypeDef].rowCount = test.count
			width := uint8(2)
			if test.count >= 65536 {
				width = 4
			}
			la.simpleSizes[tableTypeDef] = width
			data := binary.LittleEndian.AppendUint32(nil, test.code)[:width]
			for _, list := range []bool{false, true} {
				r := ecma335Reader{data: data, layout: &la}
				var got Index
				max := uint64(test.count)
				if list {
					max++
					got = r.listIndex(tableTypeDef)
				} else {
					got = r.index(tableTypeDef)
				}
				wantErr := test.code == 0 || uint64(test.code) > max
				if (r.err != nil) != wantErr || !wantErr && uint32(got) != test.code-1 {
					t.Fatalf("list=%v: index = %d, %v; want error %v", list, got, r.err, wantErr)
				}
			}
		})
	}
}

func TestRecordReaderSliceSentinels(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name    string
		count   uint32
		starts  []uint16
		row     int
		want    Slice
		wantErr bool
	}{
		{"next-sentinel", 2, []uint16{1, 3, 3}, 0, Slice{0, 2}, false},
		{"empty-middle", 2, []uint16{1, 3, 3}, 1, Slice{}, false},
		{"empty-last", 2, []uint16{1, 3, 3}, 2, Slice{}, false},
		{"last-nonempty", 2, []uint16{1}, 0, Slice{0, 2}, false},
		{"empty-target", 0, []uint16{1, 1}, 0, Slice{}, false},
		{"empty-target-last", 0, []uint16{1, 1}, 1, Slice{}, false},
		{"past-end", 2, []uint16{4}, 0, Slice{}, true},
		{"zero", 2, []uint16{0}, 0, Slice{}, true},
		{"decreasing", 2, []uint16{3, 1}, 0, Slice{}, true},
		{"next-past-end", 2, []uint16{1, 4}, 0, Slice{}, true},
		{"next-zero", 2, []uint16{1, 0}, 0, Slice{}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			var la layout
			la.tables[tableEventMap].width = 4
			la.tables[tableEvent].rowCount = test.count
			la.simpleSizes[tableEvent] = 2
			var data []byte
			for _, start := range test.starts {
				data = binary.LittleEndian.AppendUint16(data, 1) // Parent.
				data = binary.LittleEndian.AppendUint16(data, start)
			}
			r := recordReader{ecma335Reader: ecma335Reader{data: data[test.row*4+2:], layout: &la}}
			got := r.slice(tableEventMap, tableEvent)
			if (r.err != nil) != test.wantErr || !test.wantErr && got != test.want {
				t.Fatalf("slice = %v, %v; want %v, error %v", got, r.err, test.want, test.wantErr)
			}
		})
	}
}

func TestRecordReaderNullLists(t *testing.T) {
	t.Parallel()
	for _, width := range []uint8{2, 4} {
		for _, tables := range [][2]table{
			{tableTypeDef, tableField},
			{tableTypeDef, tableMethodDef},
			{tableMethodDef, tableParam},
		} {
			owner, target := tables[0], tables[1]
			t.Run(fmt.Sprintf("width-%d/owner-%d/target-%d", width, owner, target), func(t *testing.T) {
				for _, test := range []struct {
					name    string
					count   uint32
					starts  []uint32
					row     int
					want    Slice
					wantErr bool
				}{
					{"null-empty", 0, []uint32{0}, 0, Slice{}, false},
					{"null-with-target-rows", 3, []uint32{0}, 0, Slice{}, false},
					{"null-first", 3, []uint32{0, 1}, 0, Slice{}, false},
					{"non-null-after-null", 3, []uint32{0, 1}, 1, Slice{0, 3}, false},
					{"skip-null", 3, []uint32{1, 0, 3}, 0, Slice{0, 2}, false},
					{"empty-middle", 3, []uint32{1, 0, 3}, 1, Slice{}, false},
					{"last-non-null", 3, []uint32{1, 0, 3}, 2, Slice{2, 3}, false},
					{"skip-many-nulls", 3, []uint32{1, 0, 0, 4}, 0, Slice{0, 3}, false},
					{"only-nulls-follow", 3, []uint32{1, 0, 0}, 0, Slice{0, 3}, false},
					{"empty-at-sentinel", 3, []uint32{4, 0}, 0, Slice{}, false},
					{"out-of-range-after-null", 3, []uint32{1, 0, 5}, 0, Slice{}, true},
					{"decreasing-after-null", 3, []uint32{3, 0, 2}, 0, Slice{}, true},
				} {
					t.Run(test.name, func(t *testing.T) {
						var la layout
						la.tables[owner].width = 2 + width
						la.tables[target].rowCount = test.count
						la.simpleSizes[target] = width
						var data []byte
						for _, start := range test.starts {
							data = append(data, 0, 0) // Columns preceding the list.
							encoded := binary.LittleEndian.AppendUint32(nil, start)
							data = append(data, encoded[:width]...)
						}
						data = data[test.row*int(2+width)+2:]
						r := recordReader{ecma335Reader: ecma335Reader{data: data, layout: &la}}
						got := r.slice(owner, target)
						if (r.err != nil) != test.wantErr || !test.wantErr && got != test.want {
							t.Fatalf("slice = %v, %v; want %v, error %v", got, r.err, test.want, test.wantErr)
						}
						if len(r.data) != len(data)-int(width) {
							t.Fatal("lookahead changed the current row position")
						}
					})
				}
			})
		}
	}
}

func TestRecordReaderBlobReferences(t *testing.T) {
	t.Parallel()
	for _, width := range []uint8{2, 4} {
		for _, test := range []struct {
			name    string
			heap    BlobHeap
			index   uint32
			want    []byte
			wantErr bool
		}{
			{"absent-null", nil, 0, nil, false},
			{"present-null", BlobHeap{0}, 0, nil, false},
			{"absent-non-null", nil, 1, nil, true},
			{"explicit-empty", BlobHeap{0, 0}, 1, []byte{}, false},
			{"nonempty", BlobHeap{0, 1, 0xaa}, 1, []byte{0xaa}, false},
			{"past-end", BlobHeap{0}, 1, nil, true},
		} {
			t.Run(fmt.Sprintf("%d/%s", width, test.name), func(t *testing.T) {
				data := binary.LittleEndian.AppendUint32(nil, test.index)[:width]
				r := recordReader{ecma335Reader: ecma335Reader{data: data, layout: &layout{blobSize: width}}, heaps: &heaps{blobs: test.heap}}
				got := r.blob()
				if (r.err != nil) != test.wantErr || !bytes.Equal(got, test.want) {
					t.Fatalf("blob = %x, %v; want %x, error %v", got, r.err, test.want, test.wantErr)
				}
				if test.index == 0 && got != nil || !test.wantErr && test.index != 0 && got == nil {
					t.Fatal("null and non-null empty blob references were conflated")
				}
			})
		}
	}
}

func TestRecordReaderTruncatedIndices(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name string
		read func(*recordReader)
	}{
		{"string", func(r *recordReader) { r.string() }},
		{"blob", func(r *recordReader) { r.blob() }},
		{"guid", func(r *recordReader) { r.guid() }},
		{"index", func(r *recordReader) { r.index(tableTypeDef) }},
		{"list-index", func(r *recordReader) { r.listIndex(tableTypeDef) }},
		{"coded", func(r *recordReader) { readCoded[TypeDefOrRef](&r.ecma335Reader, false) }},
	} {
		t.Run(test.name, func(t *testing.T) {
			la := layout{stringSize: 2, blobSize: 2, guidSize: 2}
			la.simpleSizes[tableTypeDef] = 2
			la.tables[tableTypeDef].rowCount = 1
			la.codedSizes[codedTypeDefOrRef] = 2
			for length := range 2 {
				r := recordReader{
					ecma335Reader: ecma335Reader{data: make([]byte, length), layout: &la},
					heaps:         &heaps{strs: StringHeap{0}, blobs: BlobHeap{0}},
				}
				test.read(&r)
				if !errors.Is(r.err, io.ErrUnexpectedEOF) {
					t.Fatalf("length %d: read error = %v; want %v", length, r.err, io.ErrUnexpectedEOF)
				}
			}
		})
	}
}

func decodeErrorTestTables(t *testing.T) *Tables {
	t.Helper()
	counts := [tableMax]uint32{tableModule: 1, tableTypeDef: 2, tableField: 2, tableClassLayout: 1}
	words := []uint16{
		0, 1, 1, 0, 0, // Module: Name and Mvid at heap index 1.
		0, 0, 1, 0, 0, 1, 1, // First TypeDef: empty field and method lists.
		0, 0, 1, 0, 0, 1, 1, // Second TypeDef: owns both fields.
		6, 1, 1, // First Field: public, Name and Signature at heap offset 1.
		6, 1, 1, // Second Field.
		0, 0, 0, 1, // ClassLayout: default packing/size, Parent is TypeDef row 1.
	}
	var data []byte
	for _, word := range words {
		data = binary.LittleEndian.AppendUint16(data, word)
	}
	la, err := generateLayout(0, counts, len(data))
	if err != nil {
		t.Fatal(err)
	}
	return newTables(data, &heaps{
		strs:  StringHeap("\x00Name\x00"),
		blobs: BlobHeap{0, 2, 6, 8}, // An int32 field signature.
		guids: GUIDHeap(bytes.Repeat([]byte{1}, 16)),
	}, la)
}

func TestTableAtDecodeError(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name    string
		column  string
		wantEOF bool
		mutate  func(*Table[Field])
	}{
		{"truncated-flags", "Flags", true, func(t *Table[Field]) { t.data = t.data[:7] }},
		{"truncated-name", "Name", true, func(t *Table[Field]) { t.data = t.data[:9] }},
		{"truncated-signature", "Signature", true, func(t *Table[Field]) { t.data = t.data[:11] }},
		{"invalid-utf8", "Name", false, func(t *Table[Field]) { t.heaps.strs[1] = 0xff }},
		{"truncated-blob", "Signature", true, func(t *Table[Field]) { t.heaps.blobs = BlobHeap{0, 2, 6} }},
		{"first-error", "Name", false, func(t *Table[Field]) {
			binary.LittleEndian.PutUint16(t.data[8:], 99) // Second Field.Name.
			t.heaps.blobs = BlobHeap{0, 2, 6}             // A later error must not replace it.
		}},
	} {
		t.Run(test.name, func(t *testing.T) {
			table := decodeErrorTestTables(t).Field
			if field, err := table.At(1); err != nil || field.Name.String() != "Name" || !bytes.Equal(field.Signature, []byte{6, 8}) {
				t.Fatalf("valid field = %+v, %v", field, err)
			}
			test.mutate(&table)
			_, err := table.At(1)
			var decodeErr *DecodeError
			if !errors.As(err, &decodeErr) || decodeErr.Table != "Field" || decodeErr.Row != 1 || decodeErr.Column != test.column {
				t.Fatalf("At() error = %v; want DecodeError for Field[1].%s", err, test.column)
			}
			if errors.Is(err, io.ErrUnexpectedEOF) != test.wantEOF {
				t.Fatalf("At() error = %v; want unexpected EOF %v", err, test.wantEOF)
			}
			if decodeErr.Err == nil || !strings.Contains(err.Error(), "Field[1]."+test.column) {
				t.Fatalf("At() error lost its cause or location: %v", err)
			}
		})
	}
}

func TestTableAtDecodeErrorReferences(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name   string
		table  string
		column string
		read   func(*Tables) error
		mutate func(*Tables)
	}{
		{"guid", "Module", "Mvid", func(t *Tables) error { return readFirstRow(t.Module) }, func(t *Tables) {
			binary.LittleEndian.PutUint16(t.Module.data[4:], 2)
		}},
		{"simple-index", "ClassLayout", "Parent", func(t *Tables) error { return readFirstRow(t.ClassLayout) }, func(t *Tables) {
			binary.LittleEndian.PutUint16(t.ClassLayout.data[6:], 3)
		}},
		{"coded-index", "TypeDef", "Extends", func(t *Tables) error { return readFirstRow(t.TypeDef) }, func(t *Tables) {
			binary.LittleEndian.PutUint16(t.TypeDef.data[8:], 3<<2) // TypeDef row 3 is absent.
		}},
		{"list-lookahead", "TypeDef", "FieldList", func(t *Tables) error { return readFirstRow(t.TypeDef) }, func(t *Tables) {
			binary.LittleEndian.PutUint16(t.TypeDef.data[24:], 4) // Invalid next row's FieldList.
		}},
	} {
		t.Run(test.name, func(t *testing.T) {
			tables := decodeErrorTestTables(t)
			if err := test.read(tables); err != nil {
				t.Fatalf("valid row rejected: %v", err)
			}
			test.mutate(tables)
			err := test.read(tables)
			var decodeErr *DecodeError
			if !errors.As(err, &decodeErr) || decodeErr.Table != test.table || decodeErr.Row != 0 || decodeErr.Column != test.column || decodeErr.Err == nil {
				t.Fatalf("At() error = %v; want DecodeError for %s[0].%s", err, test.table, test.column)
			}
		})
	}
}

func TestTableAtDecodeErrorBounds(t *testing.T) {
	t.Parallel()
	for _, table := range []Table[Field]{decodeErrorTestTables(t).Field, {}} {
		for _, row := range []Index{Index(table.Len()), ^Index(0)} {
			_, err := table.At(row)
			var decodeErr *DecodeError
			if !errors.As(err, &decodeErr) || decodeErr.Table != table.Name() || decodeErr.Row != row || decodeErr.Column != "" || decodeErr.Err == nil {
				t.Fatalf("At(%d) error = %v; want row-level DecodeError", row, err)
			}
		}
	}
}

func TestDecodeError(t *testing.T) {
	t.Parallel()
	cause := &fs.PathError{Op: "read", Path: "metadata", Err: io.ErrUnexpectedEOF}
	for _, test := range []struct {
		table  string
		column string
		prefix string
	}{
		{"Field", "Signature", "Field[12].Signature: "},
		{"Field", "", "Field[12]: "},
		{"", "", "table[12]: "},
	} {
		decodeErr := &DecodeError{Table: test.table, Row: 12, Column: test.column, Err: cause}
		if got := decodeErr.Error(); got != test.prefix+cause.Error() {
			t.Fatalf("Error() = %q; want location followed by %q", got, cause.Error())
		}
		err := fmt.Errorf("load metadata: %w", decodeErr)
		var got *DecodeError
		var gotCause *fs.PathError
		if !errors.As(err, &got) || got != decodeErr || !errors.As(err, &gotCause) || gotCause != cause || !errors.Is(err, io.ErrUnexpectedEOF) {
			t.Fatalf("error chain lost context or cause: %v", err)
		}
	}
}

func TestTableAtSuccessfulReadAllocs(t *testing.T) {
	table := decodeErrorTestTables(t).Field
	var field Field
	var err error
	allocs := testing.AllocsPerRun(100, func() {
		field, err = table.At(1)
	})
	if err != nil || field.Name.String() != "Name" || !bytes.Equal(field.Signature, []byte{6, 8}) {
		t.Fatalf("valid field = %+v, %v", field, err)
	}
	if allocs != 0 {
		t.Fatalf("successful Table.At allocated %g times; want zero", allocs)
	}
}

func TestTableAllSuccessfulReadAllocs(t *testing.T) {
	table := decodeErrorTestTables(t).Field
	var count int
	var failure error
	allocs := testing.AllocsPerRun(100, func() {
		count = 0
		for _, err := range table.All() {
			if err != nil {
				failure = err
				return
			}
			count++
		}
	})
	if failure != nil || count != 2 {
		t.Fatalf("iteration returned %d rows, %v; want two rows", count, failure)
	}
	if allocs != 0 {
		t.Fatalf("successful Table.All allocated %g times; want zero", allocs)
	}
}

func TestTableAllLookahead(t *testing.T) {
	t.Parallel()
	table := decodeErrorTestTables(t).TypeDef
	var lists []Slice
	for def, err := range table.All() {
		if err != nil {
			t.Fatal(err)
		}
		lists = append(lists, def.FieldList)
	}
	// The next row bounds the first row's field list; the last owns both fields.
	if want := []Slice{{}, {Start: 0, End: 2}}; !slices.Equal(lists, want) {
		t.Fatalf("field lists = %v; want %v", lists, want)
	}

	// An invalid next row's FieldList must be reported on the current row.
	binary.LittleEndian.PutUint16(table.data[24:], 4)
	var count int
	var failure error
	for _, err := range table.All() {
		count++
		failure = err
	}
	var decodeErr *DecodeError
	if count != 1 || !errors.As(failure, &decodeErr) || decodeErr.Table != "TypeDef" || decodeErr.Row != 0 || decodeErr.Column != "FieldList" || decodeErr.Err == nil {
		t.Fatalf("iteration returned %d rows, %v; want one TypeDef[0].FieldList error", count, failure)
	}
}
