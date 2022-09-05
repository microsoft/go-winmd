// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import "math/bits"

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
	width    uint32
	offset   int
}

func (t tableInfo) rowOffset(row uint32) int {
	return t.offset + int(t.width)*int(row)
}

// generateLayout generates the bit-accurate layout for the given heapSizes and tableRowCounts.
func generateLayout(heapSizes uint8, tableRowCounts [tableMax]uint32) (la layout) {
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
		info.width += uint32(t.width(&la))
		la.tables[t] = info
		offset += int(info.width * rowCount)
	}
	return la
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
