// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/bits"
)

type layout struct {
	tables      [tableMax]tableInfo
	stringSize  uint8
	guidSize    uint8
	blobSize    uint8
	simpleSizes [tableMax]uint8
	codedSizes  [codedMax]uint8
}

type tableInfo struct {
	rowCount uint32
	width    uint8
	offset   int
}

// generateLayout generates the bit-accurate layout for the given heapSizes and tableRowCounts.
func generateLayout(heapSizes uint8, tableRowCounts [tableMax]uint32) *layout {
	var la layout
	// String, GUID, and blob index column sizes only depend on the heapSize.
	la.stringSize, la.guidSize, la.blobSize = heapIndexSize(heapSizes)

	// Simple index column sizes only depend on the number of rows of the referenced table.
	for e := table(0); e < tableMax; e++ {
		la.simpleSizes[e] = simpleIndexSize(e, tableRowCounts)
	}

	// Coded index column sizes depend on the maximum number of rows in the set of allowed tables to reference.
	for e := coded(0); e < codedMax; e++ {
		la.codedSizes[e] = codedIndexSize(e, tableRowCounts)
	}

	// We now have all the static and dynamic information to calculate the size of each table column.
	var offset int
	for t := table(0); t < tableMax; t++ {
		rowCount := tableRowCounts[t]
		if rowCount == 0 {
			continue
		}
		info := tableInfo{
			rowCount: rowCount,
			offset:   offset,
		}
		info.width += t.width(&la)
		la.tables[t] = info
		offset += int(info.width) * int(rowCount)
	}
	return &la
}

// simpleIndexSize calculates the size of the simple index e.
// Algorithm defined in §II.24.2.6.
func simpleIndexSize(e table, tableRowCounts [tableMax]uint32) uint8 {
	// e is a simple index into a table with index i, it is stored using 2 bytes if table i has
	// less than 2^16 rows, otherwise it is stored using 4 bytes.
	if tableRowCounts[e] < 1<<16 {
		return 2
	}
	return 4
}

// codedIndexSize calculates the size of the coded index e.
// Algorithm defined in §II.24.2.6.
func codedIndexSize(e coded, tableRowCounts [tableMax]uint32) uint8 {
	// e is a coded index that points into table t[i] out of n possible tables {t[0], t[n-1]}.
	tables := codedMap[e]
	// The index is stored using 2 bytes if the maximum number of rows of tables is less than 2^(16 – (log2(n))),
	// and using 4 bytes otherwise.
	var maxRowCount uint32
	for _, r := range tables {
		if r != tableNone && tableRowCounts[r] > maxRowCount {
			maxRowCount = tableRowCounts[r]
		}
	}

	var logn byte
	if len(tables) > 0 {
		// bits.Len is effectively calculating log2(n)+1.
		// We need log2(n), so subtract 1.
		n := uint(len(tables))
		logn = byte(bits.Len(n)) - 1
	}
	if maxRowCount < 1<<(16-logn) {
		return 2
	}
	return 4
}

// heapIndexSize calculates the size of indexes into the various heaps.
// The heapSizes field is a bitvector that encodes the width of indexes
// into the various heaps as retrieved from the #~ stream header.
// Algorithm defined in §II.24.2.6.
func heapIndexSize(heapSizes uint8) (strs uint8, guid uint8, blob uint8) {
	// If bit 0 is set, indexes into the “#String” heap are 4 bytes wide; if bit 1 is set, indexes into the “#GUID” heap are
	// 4 bytes wide; if bit 2 is set, indexes into the “#Blob” heap are 4 bytes wide. Conversely, if the
	// HeapSize bit for a particular heap is not set, indexes into that heap are 2 bytes wide.
	const (
		heapSizesStringBit = 1 << iota
		heapSizesGUIDBit
		heapSizesBlobBit
	)
	sizefn := func(bit uint8) uint8 {
		if heapSizes&bit != 0 {
			return 4
		}
		return 2
	}
	return sizefn(heapSizesStringBit), sizefn(heapSizesGUIDBit), sizefn(heapSizesBlobBit)
}

type recordReader struct {
	data []byte
	i    int
	err  error

	heaps  heaps
	layout *layout
}

func (r *recordReader) coded(coded coded) (_ CodedIndex) {
	if r.err != nil {
		return
	}
	tagbits := codedTagBits(coded)
	bitmask := (1 << tagbits) - 1
	code := r.uint(r.layout.codedSizes[coded])
	if code < 1 {
		return CodedIndex{Tag: -1}
	}
	row, tag := code>>tagbits-1, code&uint32(bitmask)
	_, ok := codedTable(coded, uint8(tag))
	if !ok {
		r.err = fmt.Errorf("unknown coded %d tag %d", coded, tag)
		return
	}
	return CodedIndex{
		Index: Index(row),
		Tag:   int8(tag),
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

func (r *recordReader) string() (v String) {
	if r.err != nil {
		return
	}
	idx := r.uint(r.layout.stringSize)
	v, r.err = r.heaps.strs.String(idx)
	return
}

func (r *recordReader) blob() (v []byte) {
	if r.err != nil {
		return
	}
	idx := r.uint(r.layout.blobSize)
	v, r.err = r.heaps.blobs.Bytes(idx)
	return
}

func (r *recordReader) guid() (v [16]byte) {
	if r.err != nil {
		return
	}
	idx := r.uint(r.layout.guidSize)
	if idx == 0 {
		return
	}
	// ECMA-335 GUID indices are 1-based, but we follow Go notation instead.
	v, r.err = r.heaps.guids.GUID(idx - 1)
	return
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
