// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type recordReader struct {
	data []byte
	i    int
	err  error

	strings StringHeap
	layout  *layout
}

func (r *recordReader) coded(coded coded) (_ CodedIndex) {
	if r.err != nil {
		return
	}
	tagbits := codedTagBits(coded)
	bitmask := (1 << tagbits) - 1
	code := r.uint(r.layout.codedSizes[coded])
	if code < 1 {
		return CodedIndex{table: tableNone}
	}
	row, tag := code>>tagbits-1, code&uint32(bitmask)
	table, ok := codedTable(coded, uint8(tag))
	if !ok {
		r.err = fmt.Errorf("unknown coded %d tag %d", coded, tag)
		return
	}
	return CodedIndex{
		Index: Index(row),
		table: table,
	}
}

// slice reads a Slice from r.
// ownTable is the table code of the table being read.
// targetTable is the table code of the table being referenced.
func (r *recordReader) slice(ownTable, targetTable table) Slice {
	if r.err != nil {
		return Slice{}
	}
	ownWidth := int(r.layout.tables[ownTable].width)
	var sl Slice
	// read first end of the slice so r.i+ownWidth
	// points at the same column of the next row.
	sl.End = r.peekIndex(ownWidth, targetTable)
	sl.Start = r.index(targetTable)
	if r.err == nil && sl.Start > sl.End {
		r.err = fmt.Errorf("invalid slice end: value=%d, max=%d", sl.End, sl.Start)
		return Slice{}
	}
	if sl.Start == sl.End {
		// this is a valid situation which means range is null.
		return Slice{}
	}
	return sl
}

// peekIndex reads an index at r.i+offset without advancing r.i.
// It returns the number of rows in target if r.i+offset is greater than len(r.data).
func (r *recordReader) peekIndex(offset int, target table) Index {
	if r.err != nil {
		return 0
	}
	if r.i+offset > len(r.data) {
		// we are reading the last row
		return Index(r.layout.tables[target].rowCount)
	}
	baseOffset := r.i
	r.i += offset
	v := r.index(target)
	r.i = baseOffset
	return v
}

func (r *recordReader) uint8() uint8 {
	if r.err != nil {
		return 0
	}
	v := r.data[r.i]
	r.i += 1
	return v
}

func (r *recordReader) uint16() uint16 {
	if r.err != nil {
		return 0
	}
	v := binary.LittleEndian.Uint16(r.data[r.i:])
	r.i += 2
	return v
}

func (r *recordReader) uint32() uint32 {
	if r.err != nil {
		return 0
	}
	v := binary.LittleEndian.Uint32(r.data[r.i:])
	r.i += 4
	return v
}

func (r *recordReader) string() (v string) {
	if r.err != nil {
		return ""
	}
	idx := r.uint(r.layout.stringSize)
	v, r.err = r.strings.String(idx)
	return
}

func (r *recordReader) blob() (v BlobIndex) {
	if r.err != nil {
		return 0
	}
	return BlobIndex(r.uint(r.layout.blobSize))
}

func (r *recordReader) guid() (v GUIDIndex) {
	if r.err != nil {
		return 0
	}
	return GUIDIndex(r.uint(r.layout.guidSize))
}

func (r *recordReader) index(tbl table) Index {
	if r.err != nil {
		return 0
	}
	v := r.uint(r.layout.simpleSizes[tbl])
	if v == 0 {
		r.err = errors.New("record index must be greater than 0")
		return 0
	}
	// ECMA-335 table indices are 1-based, but we follow Go notation instead.
	v -= 1
	if max := r.layout.tables[tbl].rowCount; v > max {
		r.err = fmt.Errorf("record index %d must be smaller than %d", v, max)
		return 0
	}
	return Index(v)
}

func (r *recordReader) uint(size uint8) uint32 {
	switch size {
	case 1:
		return uint32(r.uint8())
	case 2:
		return uint32(r.uint16())
	case 4:
		return r.uint32()
	default:
		panic(fmt.Errorf("columns size %d is not supported", size))
	}
}
