// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"bytes"
	"debug/pe"
	"fmt"
	"io"
)

// A Metadata represents an open Windows Metadata file.
type Metadata struct {
	Version string
	Tables  *Tables
	Strings StringHeap
	US      USHeap
	Blob    BlobHeap
	GUID    GUIDHeap
}

// New creates a new File from an underlying PE file.
func New(pefile *pe.File) (*Metadata, error) {
	return newMetadata(pefile)
}

// GUIDHeap provides access to the #GUID heap as defined in §II.24.2.5.
type GUIDHeap []byte

// USHeap provides access to the #US heap as defined in §II.24.2.4.
type USHeap []byte

// BlobHeap provides access to the #Blob heap as defined in §II.24.2.4.
type BlobHeap []byte

// StringHeap provides access to #Strings heap as defined in §II.24.2.3.
type StringHeap []byte

// String extracts string from the string heap st at offset start.
func (sh StringHeap) String(start uint32) (string, error) {
	if int(start) >= len(sh) {
		return "", fmt.Errorf("offset %d is beyond the end of string heap", start)
	}
	end := bytes.IndexByte(sh[start:], '\x00')
	if end == -1 {
		return "", fmt.Errorf("offset %d is not null-terminated", start)
	}
	return string(sh[start : int(start)+end]), nil
}

// Record is an item in a metadata table
type Record interface {
	decode(r recordReader) error
}

// Index indexes a record in a table.
type Index uint32

// CodedIndex indexes a record on any table.
type CodedIndex struct {
	Index Index
	table table
}

// IsNull returns true if the coded index points to a null record.
func (c CodedIndex) IsNull() bool {
	return c.table != tableNone
}

// BlobIndex is an index to a #Blob heap.
type BlobIndex uint32

// GUIDIndex is an index to a #GUID heap.
type GUIDIndex uint32

// Slice indexes the range of records [Start,End) on the table T.
type Slice struct {
	Start Index
	End   Index
}

// Table is a record container as defined in §II.22.
type Table[T Record] struct {
	Len uint32

	width   uint8
	data    []byte
	strings StringHeap
	layout  *layout
	newFn   func() T
}

func newTable[T Record](data []byte, stringHeap StringHeap, layout *layout, table table, newFn func() T) Table[T] {
	info := layout.tables[table]
	return Table[T]{
		Len:     info.rowCount,
		width:   uint8(info.width),
		data:    data[info.offset : info.offset+int(info.width)*int(info.rowCount)],
		strings: stringHeap,
		layout:  layout,
		newFn:   newFn,
	}
}

// Record returns the record at row.
func (t Table[T]) Record(row Index) (T, error) {
	if uint32(row) >= t.Len {
		var rec T
		return rec, fmt.Errorf("row %d is beyond the end of the table", row)
	}
	offset := int(t.width) * int(row)
	if offset+int(t.width) > len(t.data) {
		var rec T
		return rec, io.ErrUnexpectedEOF
	}
	r := recordReader{
		data:    t.data,
		i:       offset,
		strings: t.strings,
		layout:  t.layout,
	}
	// Ideally, we would instantiate rec using `var rec T`,
	// but T is a pointer type, so its default value is nil and it would
	// panic when calling decode().
	// With the newFn approach we avoid importing reflect package,
	// which is another way to instantiate generic pointer types.
	rec := t.newFn()
	err := rec.decode(r)
	if err != nil {
		return rec, err
	}
	return rec, err
}
