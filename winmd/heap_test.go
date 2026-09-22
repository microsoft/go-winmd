// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestHeapResultOwnership(t *testing.T) {
	t.Parallel()
	// These heaps are caller-owned, not read-only heaps from Metadata.
	stringsHeap := winmd.StringHeap("\x00ab\x00")
	view, err := stringsHeap.String(1)
	if err != nil || view.String() != "ab" {
		t.Fatalf("String() = %q, %v", view.String(), err)
	}
	text := view.String()
	stringsHeap[1] = 'x'
	if view.String() != "xb" || text != "ab" {
		t.Fatalf("heap view = %q, string copy = %q; want xb, ab", view.String(), text)
	}

	blobHeap := winmd.BlobHeap{0, 3, 1, 2, 3}
	blob, err := blobHeap.Bytes(1)
	if err != nil || !bytes.Equal(blob, []byte{1, 2, 3}) {
		t.Fatalf("Bytes() = %v, %v", blob, err)
	}
	copyOfBlob := bytes.Clone(blob)
	blob[0] = 9
	if blobHeap[2] != 9 || !bytes.Equal(copyOfBlob, []byte{1, 2, 3}) {
		t.Fatalf("blob view or independent copy changed unexpectedly: heap %v, copy %v", blobHeap, copyOfBlob)
	}

	guidHeap := winmd.GUIDHeap(bytes.Repeat([]byte{0xab}, 16))
	guid, err := guidHeap.GUID(0)
	if err != nil || !bytes.Equal(guid[:], guidHeap) {
		t.Fatalf("GUID() = %x, %v", guid, err)
	}
	guid[0] = 0
	guidHeap[1] = 0
	if guidHeap[0] != 0xab || guid[1] != 0xab {
		t.Fatal("GUID result aliases its heap")
	}
}

func TestStringHeapBounds(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name    string
		heap    winmd.StringHeap
		start   uint32
		want    string
		wantErr bool
	}{
		{"empty-heap", nil, 0, "", true},
		{"empty-string", winmd.StringHeap{0}, 0, "", false},
		{"string", winmd.StringHeap("\x00ab\x00"), 1, "ab", false},
		{"suffix", winmd.StringHeap("\x00ab\x00"), 2, "b", false},
		{"last-byte", winmd.StringHeap("\x00ab\x00"), 3, "", false},
		{"past-end", winmd.StringHeap("\x00ab\x00"), 4, "", true},
		{"unterminated", winmd.StringHeap("ab"), 0, "", true},
		{"int32-overflow", winmd.StringHeap{0}, 0x80000000, "", true},
		{"max-offset", winmd.StringHeap{0}, 0xffffffff, "", true},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.heap.String(test.start)
			if (err != nil) != test.wantErr {
				t.Fatalf("String(%d) error = %v; want error %v", test.start, err, test.wantErr)
			}
			if !test.wantErr && (got.String() != test.want || got.Start != test.start) {
				t.Fatalf("String(%d) = %q at %d; want %q at %d", test.start, got.String(), got.Start, test.want, test.start)
			}
		})
	}
}

func TestStringHeapUTF8(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name    string
		heap    winmd.StringHeap
		start   uint32
		want    string
		wantErr bool
	}{
		{"two-byte", winmd.StringHeap{0, 0xc3, 0xa9, 0}, 1, "é", false},
		{"three-byte", winmd.StringHeap{0, 0xe2, 0x82, 0xac, 0}, 1, "€", false},
		{"four-byte", winmd.StringHeap{0, 0xf0, 0x9f, 0x98, 0x80, 0}, 1, "😀", false},
		{"replacement-character", winmd.StringHeap{0, 0xef, 0xbf, 0xbd, 0}, 1, "\ufffd", false},
		{"trailing-garbage", winmd.StringHeap{0, 'x', 0, 0xff}, 1, "x", false},
		{"leading-garbage", winmd.StringHeap{0, 0xff, 0, 'x', 0}, 3, "x", false},
		{"invalid-lead", winmd.StringHeap{0, 0xff, 0}, 1, "", true},
		{"lone-continuation", winmd.StringHeap{0, 0x80, 0}, 1, "", true},
		{"overlong", winmd.StringHeap{0, 0xc0, 0x80, 0}, 1, "", true},
		{"truncated-sequence", winmd.StringHeap{0, 0xc3, 0}, 1, "", true},
		{"surrogate", winmd.StringHeap{0, 0xed, 0xa0, 0x80, 0}, 1, "", true},
		{"out-of-range", winmd.StringHeap{0, 0xf4, 0x90, 0x80, 0x80, 0}, 1, "", true},
		{"middle-of-sequence", winmd.StringHeap{0, 0xc3, 0xa9, 0}, 2, "", true},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.heap.String(test.start)
			if test.wantErr {
				if err == nil || !strings.Contains(err.Error(), "invalid UTF-8") || got.String() != "" {
					t.Fatalf("String(%d) = %q, %v; want zero, invalid UTF-8 error", test.start, got.String(), err)
				}
			} else if err != nil || got.String() != test.want || got.Start != test.start {
				t.Fatalf("String(%d) = %q at %d, %v; want %q", test.start, got.String(), got.Start, err, test.want)
			}
		})
	}
}

func TestGUIDHeapBounds(t *testing.T) {
	t.Parallel()
	first := bytes.Repeat([]byte{0xaa}, 16)
	second := bytes.Repeat([]byte{0xbb}, 16)
	heap := winmd.GUIDHeap(append(bytes.Clone(first), second...))
	for _, test := range []struct {
		name string
		heap winmd.GUIDHeap
		idx  uint32
		want []byte
	}{
		{"first", heap, 0, first},
		{"last", heap, 1, second},
		{"past-end", heap, 2, nil},
		{"empty", nil, 0, nil},
		{"short-first", heap[:15], 0, nil},
		{"short-last", heap[:31], 1, nil},
		{"int32-end-overflow", heap, 0x07ffffff, nil},
		{"int32-offset-overflow", heap, 0x08000000, nil},
		{"uint32-product-overflow", heap, 0x10000000, nil},
		{"max-index", heap, 0xffffffff, nil},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.heap.GUID(test.idx)
			if test.want == nil {
				if err == nil || got != [16]byte{} {
					t.Fatalf("GUID(%#x) = %x, %v; want zero, error", test.idx, got, err)
				}
			} else if err != nil || !bytes.Equal(got[:], test.want) {
				t.Fatalf("GUID(%d) = %x, %v; want %x, nil", test.idx, got, err, test.want)
			}
		})
	}
}

func TestBlobHeapBounds(t *testing.T) {
	t.Parallel()
	medium := bytes.Repeat([]byte{0xaa}, 128)
	large := bytes.Repeat([]byte{0xbb}, 16384)
	for _, test := range []struct {
		name    string
		heap    winmd.BlobHeap
		start   uint32
		want    []byte
		wantErr bool
		wantEOF bool
	}{
		{"empty-heap", nil, 0, nil, true, false},
		{"only-empty-blob", winmd.BlobHeap{0}, 0, []byte{}, false, false},
		{"padded", winmd.BlobHeap{0, 1, 0xaa, 0}, 1, []byte{0xaa}, false, false},
		{"exact-end", winmd.BlobHeap{0, 2, 0xaa, 0xbb}, 1, []byte{0xaa, 0xbb}, false, false},
		{"two-byte-length", append(winmd.BlobHeap{0, 0x80, 0x80}, medium...), 1, medium, false, false},
		{"four-byte-length", append(winmd.BlobHeap{0, 0xc0, 0, 0x40, 0}, large...), 1, large, false, false},
		{"past-end", winmd.BlobHeap{0}, 1, nil, true, false},
		{"short-length", winmd.BlobHeap{0, 0x80}, 1, nil, true, true},
		{"short-long-length", winmd.BlobHeap{0, 0xc0, 0, 0}, 1, nil, true, true},
		{"invalid-length", winmd.BlobHeap{0, 0xff}, 1, nil, true, false},
		{"short-payload", winmd.BlobHeap{0, 2, 0xaa}, 1, nil, true, true},
		{"huge-payload", winmd.BlobHeap{0, 0xdf, 0xff, 0xff, 0xff}, 1, nil, true, true},
		{"int32-overflow", winmd.BlobHeap{0}, 0x80000000, nil, true, false},
		{"max-offset", winmd.BlobHeap{0}, 0xffffffff, nil, true, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.heap.Bytes(test.start)
			if (err != nil) != test.wantErr || errors.Is(err, io.ErrUnexpectedEOF) != test.wantEOF {
				t.Fatalf("Bytes(%d) error = %v; want error %v, unexpected EOF %v", test.start, err, test.wantErr, test.wantEOF)
			}
			if test.wantErr {
				if got != nil {
					t.Fatalf("Bytes() = %x on error; want nil", got)
				}
				return
			}
			if got == nil || !bytes.Equal(got, test.want) || cap(got) != len(got) {
				t.Fatalf("Bytes() = %x (capacity %d); want %x with capacity %d", got, cap(got), test.want, len(test.want))
			}
		})
	}
}
