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

	layout *layout
}

// New creates a new File from an underlying PE file.
func New(pefile *pe.File) (*Metadata, error) {
	return newMetadata(pefile)
}

func (m *Metadata) FieldSignature(bytes SigFieldBlob) (SigField, error) {
	r := m.sigReader(bytes)
	return r.fieldSig(), r.err
}

func (m *Metadata) MethodDefSignature(data SigMethodDefBlob) (SigMethodDef, error) {
	r := m.sigReader(data)
	return r.methodDefSig(), r.err
}

func (m *Metadata) sigReader(data []byte) sigReader {
	return sigReader{
		ecma335Reader{
			data:   data,
			layout: m.layout,
		},
	}
}

// Record is an item in a metadata table.
type Record[T any] interface {
	decode(r recordReader) error
	*T
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
type String struct {
	// Start is the offset in the #Strings heap where the string starts. This is the parameter that
	// was passed to StringHeap.String to create this String. The strings heap doesn't contain
	// duplicate strings, so this value can be used to uniquely identify strings that come from the
	// same heap.
	Start uint32
	data  []byte
}

func (s String) String() string {
	return string(s.data)
}

// Slice indexes the range of records [Start,End) on the table T.
type Slice struct {
	Start Index
	End   Index
}

// Table is a record container as defined in §II.22.
type Table[T any, TP Record[T]] struct {
	Len uint32

	width  uint8
	data   []byte
	heaps  heaps
	layout *layout
}

func newTable[T any, TP Record[T]](data []byte, hps heaps, layout *layout, table table) Table[T, TP] {
	info := layout.tables[table]
	return Table[T, TP]{
		Len:    info.rowCount,
		width:  uint8(info.width),
		data:   data[info.offset : info.offset+int(info.width)*int(info.rowCount)],
		heaps:  hps,
		layout: layout,
	}
}

// Record returns the record at row.
func (t Table[T, TP]) Record(row Index) (TP, error) {
	if uint32(row) >= t.Len {
		return nil, fmt.Errorf("row %d is beyond the end of the table", row)
	}
	offset := int(t.width) * int(row)
	if offset+int(t.width) > len(t.data) {
		return nil, io.ErrUnexpectedEOF
	}
	r := recordReader{
		ecma335Reader: ecma335Reader{
			data:   t.data[offset:],
			layout: t.layout,
		},
		heaps: t.heaps,
	}
	rec := TP(new(T))
	err := rec.decode(r)
	if err != nil {
		return nil, err
	}
	return rec, err
}
