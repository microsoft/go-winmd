// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"bytes"
	"fmt"
	"io"
)

// StringHeap provides access to #Strings heap as defined in §II.24.2.3.
type StringHeap []byte

// String extracts string from the string heap st at offset start.
func (sh StringHeap) String(start uint32) (String, error) {
	if int(start) >= len(sh) {
		return nil, fmt.Errorf("offset %d is beyond the end of string heap", start)
	}
	length := bytes.IndexByte(sh[start:], '\x00')
	if length == -1 {
		return nil, fmt.Errorf("offset %d is not null-terminated", start)
	}
	end := int(start) + length
	return String(sh[start:end:end]), nil
}

// GUIDHeap provides access to the #GUID heap as defined in §II.24.2.5.
type GUIDHeap []byte

// GUID extracts the GUID from the guid heap gh at idx.
func (gh GUIDHeap) GUID(idx uint32) ([16]byte, error) {
	offset := int(idx * 16)
	if offset+16 > len(gh) {
		return [16]byte{}, fmt.Errorf("offset %d is beyond the end of the heap", offset)
	}
	var v [16]byte
	copy(v[:], gh[offset:])
	return v, nil
}

// USHeap provides access to the #US heap as defined in §II.24.2.4.
type USHeap []byte

// BlobHeap provides access to the #Blob heap as defined in §II.24.2.4.
type BlobHeap []byte

// Bytes extracts data from the blob heap bh at offset start.
func (bh BlobHeap) Bytes(start uint32) ([]byte, error) {
	if int(start) >= len(bh) {
		return nil, fmt.Errorf("offset %d is beyond the end of the heap", start)
	}
	// Blob length is stored in the first few bytes.
	const (
		mask1 byte = 0b_1000_0000
		mask2 byte = 0b_1100_0000
		mask3 byte = 0b_1110_0000
	)
	var size uint32
	switch v := bh[start]; {
	case v&mask1 == 0:
		// If the first one byte is 0bbb_bbbb, then the rest of the blob contains
		// 0bbb_bbbb bytes of actual data.
		size = uint32(v & ^mask1)
		start += 1
	case v&mask2 == mask1:
		// If the first two bytes are 10bb_bbbb and x, then the rest of the blob
		// contains the (00bb_bbbb << 8 + x) bytes of actual data.
		if int(start)+1 >= len(bh) {
			return nil, io.ErrUnexpectedEOF
		}
		size = uint32(v & ^mask2)<<8 + uint32(bh[start+1])
		start += 2
	case v&mask3 == mask2:
		// If the first four bytes are 110b_bbbb, x, y, and z, then the rest of the
		// blob contains the (000b_bbbb << 24 + x << 16 + y << 8 + z) bytes of actual data.
		if int(start)+3 >= len(bh) {
			return nil, io.ErrUnexpectedEOF
		}
		size = uint32(v & ^mask3)<<24 + uint32(bh[start+1])<<16 + uint32(bh[start+2])<<8 + uint32(bh[start+3])
		start += 4
	default:
		return nil, fmt.Errorf("invalid length %d", v)
	}
	if int(uint32(start)+size) >= len(bh) {
		return nil, io.ErrUnexpectedEOF
	}
	return bh[start : uint32(start)+size : uint32(start)+size], nil
}

type heaps struct {
	strs  StringHeap
	blobs BlobHeap
	guids GUIDHeap
}
