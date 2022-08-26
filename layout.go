// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"math"
)

type table uint8

const (
	// the following values does not appear in the spec
	tableMax  table = 0x2C + 1
	tableNone table = 255
)

type layout struct {
	tables      [tableMax]tableInfo
	stringSize  uint32
	guidSize    uint32
	blobSize    uint32
	simpleSizes [tableMax]uint32
	codedSizes  [codedMax]uint32
}

type tableInfo struct {
	rows   uint32
	width  uint32
	offset int
}

func (t tableInfo) rowOffset(row uint32) int {
	return t.offset + int(t.width)*int(row)
}

type columnType uint8

const (
	columnTypeIndex columnType = iota
	columnTypeUint
	columnTypeString
	columnTypeGUID
	columnTypeBlob
	columnTypeCodedIndex
	columnTypeSlice
)

type columnInfo struct {
	size       uint32
	columnType columnType
	table      table
	coded      coded
}

func generateLayout(heapSizes uint8, rows [tableMax]uint32) layout {
	var la layout
	la.stringSize = stringHeapIndexSize(heapSizes)
	la.guidSize = guidHeapIndexSize(heapSizes)
	la.blobSize = blobHeapIndexSize(heapSizes)
	for i := table(0); i < tableMax; i++ {
		la.simpleSizes[i] = refSimpleWidth(rows[i])
	}
	for i := coded(0); i < codedMax; i++ {
		la.codedSizes[i] = refCodedWidth(rows, codedMap[i]...)
	}
	var offset int
	for i := 0; i < len(rows); i++ {
		if rows[i] == 0 {
			continue
		}
		// deep copy table columns
		staticColInfo := staticTableInfo(table(i))
		info := tableInfo{
			rows:   rows[i],
			offset: offset,
		}
		for j := 0; j < len(staticColInfo); j++ {
			c := staticColInfo[j]
			var size uint32
			switch c.columnType {
			case columnTypeCodedIndex:
				size = la.codedSizes[c.coded]
			case columnTypeIndex, columnTypeSlice:
				size = la.simpleSizes[c.table]
			case columnTypeString:
				size = la.stringSize
			case columnTypeGUID:
				size = la.guidSize
			case columnTypeBlob:
				size = la.blobSize
			default:
				size = c.size
			}
			info.width += size
		}
		la.tables[i] = info
		offset += int(info.width * rows[i])
	}
	return la
}

func refSimpleWidth(rows uint32) uint32 {
	if rows < 1<<16 {
		return 2
	}
	return 4
}

func refCodedWidth(rows [tableMax]uint32, refs ...table) uint32 {
	// algorithm defined in §II.24.2.6.
	logn := byte(math.Ceil(math.Log2(float64(len(refs)))))
	var maxRows uint32
	for _, r := range refs {
		if r != tableNone && rows[r] > maxRows {
			maxRows = rows[r]
		}
	}

	if maxRows < 1<<(16-logn) {
		return 2
	}
	return 4
}

func stringHeapIndexSize(heapSizes uint8) uint32 {
	if (heapSizes & 1) == 1 {
		return 4
	}
	return 2
}

func guidHeapIndexSize(heapSizes uint8) uint32 {
	if (heapSizes >> 1 & 1) == 1 {
		return 4
	}
	return 2
}

func blobHeapIndexSize(heapSizes uint8) uint32 {
	if (heapSizes >> 2 & 1) == 1 {
		return 4
	}
	return 2
}
