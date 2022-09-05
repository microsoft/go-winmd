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

func (r *recordReader) slice(tbl, target table) (sl Slice) {
	if r.err != nil {
		return
	}
	width := int(r.layout.tables[tbl].width)
	nrowsref := Index(r.layout.tables[target].rowCount)
	validIdx := func(i Index) bool {
		return i <= nrowsref
	}
	base := r.i
	sl.Start = r.index(target)
	if !validIdx(sl.Start) {
		r.err = fmt.Errorf("invalid slice start: value=%d, min=1, max=%d", sl.Start, nrowsref)
		return
	}
	if r.i+width > len(r.data) {
		// we are reading the last row
		sl.End = nrowsref
	} else {
		// read the same column in the next row
		r.i = base + width
		sl.End = r.index(target)
		r.i -= width
		if !validIdx(sl.End) || sl.Start > sl.End {
			r.err = fmt.Errorf("invalid slice end: value=%d, min=%d, max=%d", sl.End, sl.Start, nrowsref)
			return
		}
		if sl.Start == sl.End {
			// range is not defined
			return
		}
	}
	return
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

func (r *recordReader) index(tbl table) (v Index) {
	if r.err != nil {
		return 0
	}
	v = Index(r.uint(r.layout.simpleSizes[tbl]))
	if v == 0 {
		r.err = errors.New("record index must be greater than 0")
	}
	return v - 1
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
