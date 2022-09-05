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

// Tables provides access to the tables and records stored in the #~ stream
// as defined in §II.24.2.6
type Tables struct {
	tables
	data    []byte
	strings StringHeap
	layout  layout
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
	table() table
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
	Len    uint32
	tables *Tables
}

func newTable[T Record](t *Tables, table table) Table[T] {
	return Table[T]{t.layout.tables[table].rowCount, t}
}

// Record returns the record at row.
func (t Table[T]) Record(row Index) (*T, error) {
	if uint32(row) >= t.Len {
		return nil, fmt.Errorf("row %d is beyond the end of the table", row)
	}
	var rec T
	info := t.tables.layout.tables[rec.table()]
	offset := info.rowOffset(uint32(row))
	if offset+int(info.width) >= len(t.tables.data) {
		return nil, io.ErrUnexpectedEOF
	}
	// instantiate and decode the record
	r := recordReader{
		data:    t.tables.data[offset : offset+(int(info.width)*int(info.rowCount-uint32(row)))],
		strings: t.tables.strings,
		layout:  &t.tables.layout,
	}
	err := any(&rec).(interface{ decode(r recordReader) error }).decode(r)
	if err != nil {
		return nil, err
	}
	return &rec, err
}
