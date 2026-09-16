// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestDecodeCompressedUint32(t *testing.T) {
	type args struct {
		bh []byte
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint32
		wantN      int
		wantErr    bool
	}{
		// Test cases in §II.23.2.
		{"1", args{[]byte{0x03}}, 0x03, 1, false},
		{"2", args{[]byte{0x7F}}, 0x7F, 1, false},
		{"3", args{[]byte{0x80, 0x80}}, 0x80, 2, false},
		{"4", args{[]byte{0xAE, 0x57}}, 0x2E57, 2, false},
		{"5", args{[]byte{0xBF, 0xFF}}, 0x3FFF, 2, false},
		{"5", args{[]byte{0xC0, 0x00, 0x40, 0x00}}, 0x4000, 4, false},
		{"6", args{[]byte{0xDF, 0xFF, 0xFF, 0xFF}}, 0x1FFF_FFFF, 4, false},
		// Zero and trailing bytes outside the compressed integer.
		{"zero", args{[]byte{0x00}}, 0, 1, false},
		{"one-byte-trailing", args{[]byte{0x03, 0xFF}}, 3, 1, false},
		{"two-byte-trailing", args{[]byte{0x80, 0x80, 0xFF}}, 0x80, 2, false},
		{"four-byte-trailing", args{[]byte{0xC0, 0x00, 0x40, 0x00, 0xFF}}, 0x4000, 4, false},
		// Example invalid data.
		{"invalid", args{[]byte{0xE0}}, 0, 0, true},
		{"invalid-high", args{[]byte{0xFE}}, 0, 0, true},
		{"null-string-marker", args{[]byte{0xFF}}, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotN, err := winmd.DecodeCompressedUint32(tt.args.bh)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeCompressedUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("DecodeCompressedUint32() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotN != tt.wantN {
				t.Errorf("DecodeCompressedUint32() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestDecodeCompressedIntegersTruncated(t *testing.T) {
	for _, test := range []struct {
		name string
		data []byte
	}{
		{"nil", nil},
		{"empty", []byte{}},
		{"two-byte-min", []byte{0x80}},
		{"two-byte-max", []byte{0xBF}},
		{"four-byte-min-1", []byte{0xC0}},
		{"four-byte-min-2", []byte{0xC0, 0x00}},
		{"four-byte-min-3", []byte{0xC0, 0x00, 0x40}},
		{"four-byte-max-1", []byte{0xDF}},
		{"four-byte-max-2", []byte{0xDF, 0xFF}},
		{"four-byte-max-3", []byte{0xDF, 0xFF, 0xFF}},
	} {
		t.Run(test.name, func(t *testing.T) {
			if value, n, err := winmd.DecodeCompressedUint32(test.data); value != 0 || n != 0 || !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Errorf("DecodeCompressedUint32(%x) = (%d, %d, %v); want (0, 0, %v)", test.data, value, n, err, io.ErrUnexpectedEOF)
			}
			if value, n, err := winmd.DecodeCompressedInt32(test.data); value != 0 || n != 0 || !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Errorf("DecodeCompressedInt32(%x) = (%d, %d, %v); want (0, 0, %v)", test.data, value, n, err, io.ErrUnexpectedEOF)
			}
		})
	}
}

func TestDecodeCompressedInt32(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name       string
		args       args
		wantResult int32
		wantN      int
		wantErr    bool
	}{
		// Test cases in §II.23.2.
		{"1", args{[]byte{0x06}}, 3, 1, false},
		{"2", args{[]byte{0x7B}}, -3, 1, false},
		{"3", args{[]byte{0x80, 0x80}}, 64, 2, false},
		{"4", args{[]byte{0x01}}, -64, 1, false},
		{"5", args{[]byte{0xC0, 0x00, 0x40, 0x00}}, 8192, 4, false},
		{"6", args{[]byte{0x80, 0x01}}, -8192, 2, false},
		{"7", args{[]byte{0xDF, 0xFF, 0xFF, 0xFE}}, 268435455, 4, false},
		{"8", args{[]byte{0xC0, 0x00, 0x00, 0x01}}, -268435456, 4, false},
		// Zero and signed-width boundaries.
		{"zero", args{[]byte{0x00}}, 0, 1, false},
		{"negative-one", args{[]byte{0x7F}}, -1, 1, false},
		{"one-byte-max", args{[]byte{0x7E}}, 63, 1, false},
		{"two-byte-negative-boundary", args{[]byte{0xBF, 0x7F}}, -65, 2, false},
		{"two-byte-max", args{[]byte{0xBF, 0xFE}}, 8191, 2, false},
		{"four-byte-negative-boundary", args{[]byte{0xDF, 0xFF, 0xBF, 0xFF}}, -8193, 4, false},
		// Trailing bytes must not affect sign extension or bytes consumed.
		{"one-byte-trailing", args{[]byte{0x7B, 0xFF}}, -3, 1, false},
		{"two-byte-trailing", args{[]byte{0x80, 0x01, 0xFF}}, -8192, 2, false},
		{"four-byte-trailing", args{[]byte{0xC0, 0x00, 0x00, 0x01, 0xFF}}, -268435456, 4, false},
		// Invalid unsigned prefixes are also invalid signed prefixes.
		{"invalid", args{[]byte{0xE0}}, 0, 0, true},
		{"invalid-high", args{[]byte{0xFE}}, 0, 0, true},
		{"null-string-marker", args{[]byte{0xFF}}, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotN, err := winmd.DecodeCompressedInt32(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeCompressedInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("DecodeCompressedInt32() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotN != tt.wantN {
				t.Errorf("DecodeCompressedInt32() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestDecodeSerString(t *testing.T) {
	oneByteString := strings.Repeat("a", 0x7F)
	twoByteString := strings.Repeat("b", 0x80)
	twoByteMaxString := strings.Repeat("c", 0x3FFF)
	fourByteString := strings.Repeat("d", 0x4000)
	tests := []struct {
		name  string
		data  []byte
		want  []byte
		wantN int
	}{
		{"null", []byte{0xFF}, nil, 1},
		{"empty", []byte{0}, []byte{}, 1},
		{"ascii", []byte{3, 'a', 'b', 'c'}, []byte("abc"), 4},
		{"embedded-null", []byte{3, 'a', 0, 'b'}, []byte("a\x00b"), 4},
		{"utf8-two-byte", []byte{2, 0xC3, 0xA9}, []byte("\u00e9"), 3},
		{"utf8-three-byte", []byte{3, 0xE2, 0x82, 0xAC}, []byte("\u20ac"), 4},
		{"utf8-four-byte", []byte{4, 0xF0, 0x9F, 0x98, 0x80}, []byte("\U0001f600"), 5},
		{"replacement-character", []byte{3, 0xEF, 0xBF, 0xBD}, []byte("\ufffd"), 4},
		{"one-byte-max", append([]byte{0x7F}, oneByteString...), []byte(oneByteString), 1 + len(oneByteString)},
		{"two-byte-min", append([]byte{0x80, 0x80}, twoByteString...), []byte(twoByteString), 2 + len(twoByteString)},
		{"two-byte-max", append([]byte{0xBF, 0xFF}, twoByteMaxString...), []byte(twoByteMaxString), 2 + len(twoByteMaxString)},
		{"four-byte-min", append([]byte{0xC0, 0x00, 0x40, 0x00}, fourByteString...), []byte(fourByteString), 4 + len(fourByteString)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, tail := range [][]byte{nil, {0xFF, 0xFE}} {
				data := append(bytes.Clone(test.data), tail...)
				got, n, err := winmd.DecodeSerString(data)
				if err != nil {
					t.Fatal(err)
				}
				if (got == nil) != (test.want == nil) || !bytes.Equal(got, test.want) || n != test.wantN {
					t.Fatalf("DecodeSerString() = (%q, %d), null = %v; want (%q, %d), null = %v", got, n, got == nil, test.want, test.wantN, test.want == nil)
				}
				if cap(got) != len(got) {
					t.Errorf("capacity = %d; want %d", cap(got), len(got))
				}
				if len(got) > 0 && &got[0] != &data[n-len(got)] {
					t.Error("decoded bytes do not alias the input")
				}
				if !bytes.Equal(data[n:], tail) {
					t.Errorf("remaining bytes = %x; want %x", data[n:], tail)
				}
			}
		})
	}
}

func TestDecodeSerStringErrors(t *testing.T) {
	for _, test := range []struct {
		name    string
		data    []byte
		wantEOF bool
	}{
		{"nil", nil, true},
		{"empty", []byte{}, true},
		{"two-byte-prefix", []byte{0x80}, true},
		{"four-byte-prefix-1", []byte{0xC0}, true},
		{"four-byte-prefix-2", []byte{0xC0, 0x00}, true},
		{"four-byte-prefix-3", []byte{0xC0, 0x00, 0x40}, true},
		{"one-byte-payload", []byte{3, 'a', 'b'}, true},
		{"two-byte-payload", []byte{0x80, 0x80, 'a'}, true},
		{"four-byte-payload", []byte{0xC0, 0x00, 0x40, 0x00, 'b'}, true},
		{"maximum-length", []byte{0xDF, 0xFF, 0xFF, 0xFF}, true},
		{"invalid-prefix", []byte{0xE0}, false},
		{"invalid-prefix-high", []byte{0xFE}, false},
		{"null-marker-in-payload", []byte{1, 0xFF}, false},
		{"utf8-continuation", []byte{1, 0x80}, false},
		{"utf8-incomplete", []byte{1, 0xC3}, false},
		{"utf8-exceeds-payload", []byte{1, 0xC3, 0xA9}, false},
		{"utf8-overlong", []byte{2, 0xC0, 0x80}, false},
		{"utf8-surrogate", []byte{3, 0xED, 0xA0, 0x80}, false},
		{"utf8-out-of-range", []byte{4, 0xF4, 0x90, 0x80, 0x80}, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, n, err := winmd.DecodeSerString(test.data)
			if err == nil || got != nil || n != 0 {
				t.Fatalf("DecodeSerString(%x) = (%q, %d, %v); want (nil, 0, error)", test.data, got, n, err)
			}
			if errors.Is(err, io.ErrUnexpectedEOF) != test.wantEOF {
				t.Errorf("error = %v; want unexpected EOF %v", err, test.wantEOF)
			}
		})
	}
}

func TestDecodeConstant(t *testing.T) {
	for _, test := range []struct {
		name string
		typ  winmd.ElementType
		data []byte
		want any
	}{
		{"false", winmd.ElementType_BOOLEAN, []byte{0}, false},
		{"true", winmd.ElementType_BOOLEAN, []byte{1}, true},
		{"nonzero-boolean", winmd.ElementType_BOOLEAN, []byte{0xFF}, true},
		{"char", winmd.ElementType_CHAR, []byte{0xAC, 0x20}, uint16(0x20AC)},
		{"char-surrogate", winmd.ElementType_CHAR, []byte{0x00, 0xD8}, uint16(0xD800)},
		{"i1-min", winmd.ElementType_I1, []byte{0x80}, int8(-0x80)},
		{"i1-max", winmd.ElementType_I1, []byte{0x7F}, int8(0x7F)},
		{"u1-max", winmd.ElementType_U1, []byte{0xFF}, uint8(0xFF)},
		{"i2-min", winmd.ElementType_I2, []byte{0, 0x80}, int16(-0x8000)},
		{"i2-max", winmd.ElementType_I2, []byte{0xFF, 0x7F}, int16(0x7FFF)},
		{"u2", winmd.ElementType_U2, []byte{0x34, 0x12}, uint16(0x1234)},
		{"u2-max", winmd.ElementType_U2, []byte{0xFF, 0xFF}, uint16(0xFFFF)},
		{"i4-min", winmd.ElementType_I4, []byte{0, 0, 0, 0x80}, int32(-0x8000_0000)},
		{"i4-max", winmd.ElementType_I4, []byte{0xFF, 0xFF, 0xFF, 0x7F}, int32(0x7FFF_FFFF)},
		{"u4", winmd.ElementType_U4, []byte{0x78, 0x56, 0x34, 0x12}, uint32(0x1234_5678)},
		{"u4-max", winmd.ElementType_U4, []byte{0xFF, 0xFF, 0xFF, 0xFF}, uint32(0xFFFF_FFFF)},
		{"i8-min", winmd.ElementType_I8, []byte{0, 0, 0, 0, 0, 0, 0, 0x80}, int64(-0x8000_0000_0000_0000)},
		{"i8-max", winmd.ElementType_I8, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, int64(0x7FFF_FFFF_FFFF_FFFF)},
		{"u8", winmd.ElementType_U8, []byte{0xEF, 0xCD, 0xAB, 0x89, 0x67, 0x45, 0x23, 0x01}, uint64(0x0123_4567_89AB_CDEF)},
		{"u8-max", winmd.ElementType_U8, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, uint64(0xFFFF_FFFF_FFFF_FFFF)},
		{"r4", winmd.ElementType_R4, []byte{0, 0, 0xC0, 0x3F}, float32(1.5)},
		{"r8", winmd.ElementType_R8, []byte{0, 0, 0, 0, 0, 0, 0xF8, 0xBF}, float64(-1.5)},
		{"empty-string", winmd.ElementType_STRING, nil, ""},
		{"ascii-string", winmd.ElementType_STRING, []byte{'A', 0, 'B', 0}, "AB"},
		{"unicode-string", winmd.ElementType_STRING, []byte{0xAC, 0x20}, "\u20ac"},
		{"surrogate-pair", winmd.ElementType_STRING, []byte{0x3D, 0xD8, 0x00, 0xDE}, "\U0001f600"},
		{"embedded-null", winmd.ElementType_STRING, []byte{'A', 0, 0, 0, 'B', 0}, "A\x00B"},
		{"string-not-null", winmd.ElementType_STRING, []byte{0, 0, 0, 0}, "\x00\x00"},
		{"byte-order-marker", winmd.ElementType_STRING, []byte{0xFF, 0xFE}, "\ufeff"},
		{"unpaired-high-surrogate", winmd.ElementType_STRING, []byte{0x00, 0xD8, 'A', 0}, "\ufffdA"},
		{"unpaired-low-surrogate", winmd.ElementType_STRING, []byte{0x00, 0xDC}, "\ufffd"},
		{"null-reference", winmd.ElementType_CLASS, []byte{0, 0, 0, 0}, nil},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := winmd.DecodeConstant(test.typ, test.data)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Errorf("DecodeConstant() = %#v (%T); want %#v (%T)", got, got, test.want, test.want)
			}
		})
	}
}

func TestDecodeConstantFloatBits(t *testing.T) {
	for _, bits := range []uint32{0, 0x8000_0000, 1, 0x7F7F_FFFF, 0x7F80_0000, 0xFF80_0000, 0x7FC1_2345} {
		got, err := winmd.DecodeConstant(winmd.ElementType_R4, binary.LittleEndian.AppendUint32(nil, bits))
		value, ok := got.(float32)
		if err != nil || !ok || math.Float32bits(value) != bits {
			t.Errorf("R4 bits %#x: got %v (%T), error %v", bits, got, got, err)
		}
	}
	for _, bits := range []uint64{0, 0x8000_0000_0000_0000, 1, 0x7FEF_FFFF_FFFF_FFFF, 0x7FF0_0000_0000_0000, 0xFFF0_0000_0000_0000, 0x7FF8_0000_0000_1234} {
		got, err := winmd.DecodeConstant(winmd.ElementType_R8, binary.LittleEndian.AppendUint64(nil, bits))
		value, ok := got.(float64)
		if err != nil || !ok || math.Float64bits(value) != bits {
			t.Errorf("R8 bits %#x: got %v (%T), error %v", bits, got, got, err)
		}
	}
}

func TestDecodeConstantLengths(t *testing.T) {
	for _, test := range []struct {
		typ  winmd.ElementType
		size int
	}{
		{winmd.ElementType_BOOLEAN, 1},
		{winmd.ElementType_CHAR, 2},
		{winmd.ElementType_I1, 1},
		{winmd.ElementType_U1, 1},
		{winmd.ElementType_I2, 2},
		{winmd.ElementType_U2, 2},
		{winmd.ElementType_I4, 4},
		{winmd.ElementType_U4, 4},
		{winmd.ElementType_I8, 8},
		{winmd.ElementType_U8, 8},
		{winmd.ElementType_R4, 4},
		{winmd.ElementType_R8, 8},
		{winmd.ElementType_CLASS, 4},
	} {
		t.Run(test.typ.String(), func(t *testing.T) {
			for size := 0; size < test.size; size++ {
				got, err := winmd.DecodeConstant(test.typ, make([]byte, size))
				if got != nil || !errors.Is(err, io.ErrUnexpectedEOF) {
					t.Errorf("length %d: got (%v, %v); want (nil, %v)", size, got, err, io.ErrUnexpectedEOF)
				}
			}
			got, err := winmd.DecodeConstant(test.typ, make([]byte, test.size+1))
			if got != nil || err == nil || errors.Is(err, io.ErrUnexpectedEOF) {
				t.Errorf("extra byte: got (%v, %v); want (nil, length error)", got, err)
			}
		})
	}
}

func TestDecodeConstantErrors(t *testing.T) {
	for _, test := range []struct {
		name    string
		typ     winmd.ElementType
		data    []byte
		wantEOF bool
	}{
		{"odd-string", winmd.ElementType_STRING, []byte{'A', 0, 'B'}, true},
		{"ser-string-null-marker", winmd.ElementType_STRING, []byte{0xFF}, true},
		{"nonzero-null-low", winmd.ElementType_CLASS, []byte{1, 0, 0, 0}, false},
		{"nonzero-null-high", winmd.ElementType_CLASS, []byte{0, 0, 0, 1}, false},
		{"void", winmd.ElementType_VOID, nil, false},
		{"object", winmd.ElementType_OBJECT, []byte{0, 0, 0, 0}, false},
		{"native-int", winmd.ElementType_I, []byte{0, 0, 0, 0}, false},
		{"native-uint", winmd.ElementType_U, []byte{0, 0, 0, 0}, false},
		{"unknown", winmd.ElementType(0xFF), nil, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := winmd.DecodeConstant(test.typ, test.data)
			if got != nil || err == nil || errors.Is(err, io.ErrUnexpectedEOF) != test.wantEOF {
				t.Errorf("DecodeConstant() = (%v, %v); want (nil, error), unexpected EOF %v", got, err, test.wantEOF)
			}
		})
	}
}

func TestConstantDecodeValue(t *testing.T) {
	for _, test := range []struct {
		name     string
		constant winmd.Constant
	}{
		{"integer", winmd.Constant{Type: winmd.ElementType_I2, Value: []byte{0xFF, 0xFF}}},
		{"string", winmd.Constant{Type: winmd.ElementType_STRING, Value: []byte{'A', 0}}},
		{"null", winmd.Constant{Type: winmd.ElementType_CLASS, Value: []byte{0, 0, 0, 0}}},
		{"truncated", winmd.Constant{Type: winmd.ElementType_I4, Value: []byte{0}}},
		{"extra-byte", winmd.Constant{Type: winmd.ElementType_U1, Value: []byte{0, 0}}},
		{"nonzero-null", winmd.Constant{Type: winmd.ElementType_CLASS, Value: []byte{1, 0, 0, 0}}},
		{"unsupported", winmd.Constant{}},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.constant.DecodeValue()
			want, wantErr := winmd.DecodeConstant(test.constant.Type, test.constant.Value)
			if got != want || (err != nil) != (wantErr != nil) {
				t.Fatalf("Constant.DecodeValue() = (%#v, %v); DecodeConstant() = (%#v, %v)", got, err, want, wantErr)
			}
			if err != nil && (err.Error() != wantErr.Error() || errors.Is(err, io.ErrUnexpectedEOF) != errors.Is(wantErr, io.ErrUnexpectedEOF)) {
				t.Errorf("Constant.DecodeValue() error = %v; want %v", err, wantErr)
			}
		})
	}
}

func TestConstantDecodeValueMetadata(t *testing.T) {
	metadata, err := winmd.Open("testdata/Windows.Win32.winmd")
	if err != nil {
		t.Fatal(err)
	}
	for index := range metadata.Tables.Constant.Indices() {
		constant, err := metadata.Tables.Constant.At(index)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := constant.DecodeValue(); err != nil {
			t.Fatalf("constant %d (%v, %d bytes): %v", index, constant.Type, len(constant.Value), err)
		}
	}
}
