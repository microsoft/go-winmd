// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package ecma335

import "testing"

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
		// Example invalid data.
		{"invalid", args{[]byte{0xE0}}, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotN, err := DecodeCompressedUint32(tt.args.bh)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotN, err := DecodeCompressedInt32(tt.args.data)
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
