// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

// StringHeap provides access to #Strings heap as defined in §II.24.2.3.
type StringHeap []byte

// String extracts string from the string heap st at offset start.
func (sh StringHeap) String(start uint32) (String, error) {
	if uint64(start) >= uint64(len(sh)) {
		return String{}, fmt.Errorf("offset %d is beyond the end of string heap", start)
	}
	length := bytes.IndexByte(sh[start:], '\x00')
	if length == -1 {
		return String{}, fmt.Errorf("offset %d is not null-terminated", start)
	}
	end := int(start) + length
	data := sh[start:end:end]
	if !utf8.Valid(data) {
		return String{}, fmt.Errorf("offset %d contains invalid UTF-8 in string heap", start)
	}
	return String{start, data}, nil
}

// GUIDHeap provides access to the #GUID heap as defined in §II.24.2.5.
type GUIDHeap []byte

// GUID extracts the GUID from the guid heap gh at idx.
func (gh GUIDHeap) GUID(idx uint32) ([16]byte, error) {
	offset := uint64(idx) * 16
	if offset+16 > uint64(len(gh)) {
		return [16]byte{}, fmt.Errorf("offset %d is beyond the end of the heap", offset)
	}
	var v [16]byte
	copy(v[:], gh[offset:offset+16])
	return v, nil
}

// USHeap provides access to the #US heap as defined in §II.24.2.4.
type USHeap []byte

// BlobHeap provides access to the #Blob heap as defined in §II.24.2.4.
type BlobHeap []byte

// Bytes extracts data from the blob heap bh at offset start.
func (bh BlobHeap) Bytes(start uint32) ([]byte, error) {
	if uint64(start) >= uint64(len(bh)) {
		return nil, fmt.Errorf("offset %d is beyond the end of the heap", start)
	}
	data := bh[start:]
	size, n, err := DecodeCompressedUint32(data)
	if err != nil {
		return nil, err
	}
	data = data[n:]
	if uint64(size) > uint64(len(data)) {
		return nil, io.ErrUnexpectedEOF
	}
	return data[:size:size], nil
}

type heaps struct {
	strs  StringHeap
	blobs BlobHeap
	guids GUIDHeap
}
