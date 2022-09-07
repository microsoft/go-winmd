// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
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

// Record is an item in a metadata table
type Record interface {
	decode(r recordReader) error
}

// Index indexes a record in a table.
type Index uint32

// CodedIndex indexes a record on any table.
type CodedIndex struct {
	Index Index
	Tag   int8
}

// String is complete UTF8 string from the #String heap
// It does not contain the null-terminated character.
//
// It is used as an optimization to avoid allocating
// when reading from the #Strings heap.
type String []byte

func (s String) String() string {
	return string(s)
}

// Slice indexes the range of records [Start,End) on the table T.
type Slice struct {
	Start Index
	End   Index
}

// Table is a record container as defined in §II.22.
type Table[T Record] struct {
	Len uint32

	width  uint8
	data   []byte
	heaps  heaps
	layout *layout
	newFn  func() T
}

func newTable[T Record](data []byte, hps heaps, layout *layout, table table, newFn func() T) Table[T] {
	info := layout.tables[table]
	return Table[T]{
		Len:    info.rowCount,
		width:  uint8(info.width),
		data:   data[info.offset : info.offset+int(info.width)*int(info.rowCount)],
		heaps:  hps,
		layout: layout,
		newFn:  newFn,
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
		data:   t.data,
		i:      offset,
		heaps:  t.heaps,
		layout: t.layout,
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
