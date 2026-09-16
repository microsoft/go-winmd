// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"bytes"
	"debug/pe"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

type metadataStream struct {
	name string
	data []byte
}

func metadataRoot(streams ...metadataStream) []byte {
	return metadataRootVersion("v4.0.30319", streams...)
}

func metadataRootVersion(version string, streams ...metadataStream) []byte {
	versionBytes := append([]byte(version), 0)
	versionBytes = append(versionBytes, make([]byte, (4-len(versionBytes)%4)%4)...)
	data := make([]byte, 16)
	binary.LittleEndian.PutUint32(data, 0x424a5342)
	binary.LittleEndian.PutUint16(data[4:], 1)
	binary.LittleEndian.PutUint16(data[6:], 1)
	binary.LittleEndian.PutUint32(data[12:], uint32(len(versionBytes)))
	data = append(data, versionBytes...)
	data = binary.LittleEndian.AppendUint16(data, 0)
	data = binary.LittleEndian.AppendUint16(data, uint16(len(streams)))
	offset := len(data)
	for _, stream := range streams {
		offset += 8 + (len(stream.name)+4)&^3
	}
	for _, stream := range streams {
		data = binary.LittleEndian.AppendUint32(data, uint32(offset))
		data = binary.LittleEndian.AppendUint32(data, uint32(len(stream.data)))
		data = append(data, stream.name...)
		data = append(data, make([]byte, 4-len(stream.name)%4)...)
		offset += (len(stream.data) + 3) &^ 3
	}
	for _, stream := range streams {
		data = append(data, stream.data...)
		data = append(data, make([]byte, (4-len(stream.data)%4)%4)...)
	}
	return data
}

// metadataPE constructs an in-memory PE image containing a CLI header and metadata root.
func metadataPE(t *testing.T, metadata []byte) *pe.File {
	t.Helper()
	return metadataPEWithDirectorySize(t, metadata, uint32(len(metadata)))
}

func metadataPEWithDirectorySize(t *testing.T, metadata []byte, directorySize uint32) *pe.File {
	t.Helper()
	return metadataPEForArch(t, metadata, directorySize, false)
}

func metadataPEForArch(t *testing.T, metadata []byte, directorySize uint32, pe64 bool) *pe.File {
	t.Helper()
	const sectionRVA = 0x2000
	const sectionOffset = 0x200
	const cliSize = 72
	section := make([]byte, cliSize)
	binary.LittleEndian.PutUint32(section, cliSize)
	binary.LittleEndian.PutUint16(section[4:], 2)
	binary.LittleEndian.PutUint16(section[6:], 5)
	binary.LittleEndian.PutUint32(section[8:], sectionRVA+cliSize)
	binary.LittleEndian.PutUint32(section[12:], directorySize)
	section = append(section, metadata...)

	dos := make([]byte, 64)
	copy(dos, "MZ")
	binary.LittleEndian.PutUint32(dos[60:], uint32(len(dos)))
	b := bytes.NewBuffer(dos)
	b.WriteString("PE\x00\x00")
	var directories [16]pe.DataDirectory
	directories[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR] = pe.DataDirectory{
		VirtualAddress: sectionRVA,
		Size:           cliSize,
	}
	machine := uint16(pe.IMAGE_FILE_MACHINE_I386)
	var optional any = pe.OptionalHeader32{Magic: 0x10b, NumberOfRvaAndSizes: 16, DataDirectory: directories}
	if pe64 {
		machine = pe.IMAGE_FILE_MACHINE_AMD64
		optional = pe.OptionalHeader64{Magic: 0x20b, NumberOfRvaAndSizes: 16, DataDirectory: directories}
	}
	for _, header := range []any{
		pe.FileHeader{
			Machine:              machine,
			NumberOfSections:     1,
			SizeOfOptionalHeader: uint16(binary.Size(optional)),
		},
		optional,
		pe.SectionHeader32{
			Name:             [8]byte{'.', 't', 'e', 'x', 't'},
			VirtualSize:      uint32(len(section)),
			VirtualAddress:   sectionRVA,
			SizeOfRawData:    uint32(len(section)),
			PointerToRawData: sectionOffset,
		},
	} {
		if err := binary.Write(b, binary.LittleEndian, header); err != nil {
			t.Fatal(err)
		}
	}
	b.Write(make([]byte, sectionOffset-b.Len()))
	b.Write(section)
	f, err := pe.NewFile(bytes.NewReader(b.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	return f
}

func metadataTables(heapSizes byte, valid uint64, rows []uint32, payloadSize int) []byte {
	data := make([]byte, 24)
	data[4] = 2
	data[6] = heapSizes
	data[7] = 1
	binary.LittleEndian.PutUint64(data[8:], valid)
	for _, count := range rows {
		data = binary.LittleEndian.AppendUint32(data, count)
	}
	return append(data, make([]byte, payloadSize)...)
}

func TestNewDeclSecurityCustomAttributes(t *testing.T) {
	t.Parallel()
	for _, count := range []uint32{1, 2047, 2048} {
		t.Run(fmt.Sprint(count), func(t *testing.T) {
			tables := metadataTables(0, 1<<2|1<<10|1<<12|1<<14, []uint32{1, 1, 2, count}, 0)
			// One TypeDef, followed by a MemberRef constructor on that type.
			tables = append(tables, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0)
			tables = append(tables, 8, 0, 6, 0, 1, 0)
			for _, parent := range []uint32{1<<5 | 6, count<<5 | 8} {
				// DeclSecurity contributes to the width even when the attribute
				// is attached to a different table, such as MemberRef.
				if count >= 2048 {
					tables = binary.LittleEndian.AppendUint32(tables, parent)
				} else {
					tables = binary.LittleEndian.AppendUint16(tables, uint16(parent))
				}
				tables = binary.LittleEndian.AppendUint16(tables, 11) // MemberRef row 1.
				tables = binary.LittleEndian.AppendUint16(tables, 5)
			}
			for range count {
				tables = append(tables, 2, 0, 4, 0, 0, 0) // DeclSecurity on TypeDef row 1.
			}
			m, err := winmd.New(metadataPE(t, metadataRoot(
				metadataStream{"#~", tables},
				metadataStream{"#Strings", []byte("\x00Type\x00.ctor\x00")},
				metadataStream{"#Blob", []byte{0, 3, 0x20, 0, 1, 4, 1, 0, 0, 0}},
			)))
			if err != nil {
				t.Fatal(err)
			}
			for i, want := range []winmd.CodedIndex[winmd.HasCustomAttribute]{
				{Tag: winmd.HasCustomAttribute_MemberRef, Index: 0},
				{Tag: winmd.HasCustomAttribute_DeclSecurity, Index: winmd.Index(count - 1)},
			} {
				attribute, err := m.Tables.CustomAttribute.At(winmd.Index(i))
				if err != nil {
					t.Fatal(err)
				}
				if attribute.Parent != want || attribute.Type.Tag != winmd.CustomAttributeType_MemberRef || attribute.Type.Index != 0 || !bytes.Equal(attribute.Value, []byte{1, 0, 0, 0}) {
					t.Fatalf("CustomAttribute row %d = %+v; want parent %+v, MemberRef constructor and empty arguments", i, attribute, want)
				}
			}
			last, err := m.Tables.DeclSecurity.At(winmd.Index(count - 1))
			if err != nil || last.Action != 2 || last.Parent.Tag != winmd.HasDeclSecurity_TypeDef || last.Parent.Index != 0 {
				t.Fatalf("following DeclSecurity row is misaligned: %+v, %v", last, err)
			}
		})
	}
}

func TestNewCodedReferenceBounds(t *testing.T) {
	t.Parallel()
	for _, row := range []uint16{1, 2} {
		t.Run(fmt.Sprint(row), func(t *testing.T) {
			tables := metadataTables(0, 1<<2|1<<9, []uint32{1, 1}, 0)
			tables = append(tables, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0)
			tables = binary.LittleEndian.AppendUint16(tables, 1)
			tables = binary.LittleEndian.AppendUint16(tables, row<<2)
			m, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables})))
			if err != nil {
				t.Fatal(err)
			}
			impl, err := m.Tables.InterfaceImpl.At(0)
			if row == 1 {
				if err != nil || impl.Interface.Tag != winmd.TypeDefOrRef_TypeDef || impl.Interface.Index != 0 {
					t.Fatalf("valid interface reference = %+v, %v", impl.Interface, err)
				}
			} else if err == nil {
				t.Fatalf("out-of-range interface reference accepted: %+v", impl.Interface)
			}
		})
	}
}

func TestNewNullTypeDefLists(t *testing.T) {
	t.Parallel()
	for _, fields := range []uint16{0, 1} {
		for _, methods := range []uint16{0, 1} {
			t.Run(fmt.Sprintf("fields-%d/methods-%d", fields, methods), func(t *testing.T) {
				tables := metadataTables(0, 1<<2, []uint32{1}, 14)
				binary.LittleEndian.PutUint16(tables[32:], 1) // TypeDef.Name.
				binary.LittleEndian.PutUint16(tables[38:], fields)
				binary.LittleEndian.PutUint16(tables[40:], methods)
				m, err := winmd.New(metadataPE(t, metadataRoot(
					metadataStream{"#~", tables},
					metadataStream{"#Strings", []byte("\x00<Module>\x00")},
				)))
				if err != nil {
					t.Fatal(err)
				}
				typ, err := m.Tables.TypeDef.At(0)
				if err != nil || typ.Name.String() != "<Module>" || typ.FieldList != (winmd.Slice{}) || typ.MethodList != (winmd.Slice{}) {
					t.Fatalf("TypeDef with null/empty lists = %+v, %v; want empty lists", typ, err)
				}
			})
		}
	}
}

func TestNewOmittedEmptyBlobHeap(t *testing.T) {
	t.Parallel()
	for _, heap := range []struct {
		name    string
		present bool
		data    []byte
	}{
		{"omitted", false, nil},
		{"empty", true, nil},
		{"empty-entry", true, []byte{0}},
	} {
		for _, index := range []uint16{0, 1} {
			t.Run(fmt.Sprintf("%s/index-%d", heap.name, index), func(t *testing.T) {
				tables := metadataTables(0, 1<<32, []uint32{1}, 22)
				binary.LittleEndian.PutUint16(tables[44:], index) // Assembly.PublicKey.
				binary.LittleEndian.PutUint16(tables[46:], 1)     // Assembly.Name.
				streams := []metadataStream{
					{"#~", tables}, {"#Strings", []byte("\x00Example\x00")},
				}
				if heap.present {
					streams = append(streams, metadataStream{"#Blob", heap.data})
				}
				m, err := winmd.New(metadataPE(t, metadataRoot(streams...)))
				if err != nil {
					t.Fatal(err)
				}
				assembly, err := m.Tables.Assembly.At(0)
				if index != 0 {
					if err == nil {
						t.Fatal("non-null reference into an empty heap was accepted")
					}
				} else if err != nil || assembly.Name.String() != "Example" || assembly.PublicKey != nil {
					t.Fatalf("null Assembly.PublicKey = %+v, %v", assembly, err)
				}
			})
		}
	}
}

func TestNewStringHeap(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"empty", nil, true},
		{"missing-empty-entry", []byte{'x'}, true},
		{"nonempty-first-string", []byte{'x', 0}, true},
		{"empty-string", []byte{0}, false},
		{"strings", []byte{0, 'x', 0}, false},
		{"trailing-garbage", []byte{0, 'x', 0, 0xfe}, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			m, err := winmd.New(metadataPE(t, metadataRoot(
				metadataStream{"#Strings", test.data},
				metadataStream{"#~", metadataTables(0, 0, nil, 0)},
			)))
			if test.wantErr {
				if err == nil || !strings.Contains(err.Error(), "string heap must start with the empty string") {
					t.Fatalf("New() error = %v; want a string heap first-entry error", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if m.Version != "v4.0.30319" || !bytes.Equal(m.Strings, test.data) {
				t.Fatalf("New() = version %q, strings %v; want version v4.0.30319, strings %v", m.Version, m.Strings, test.data)
			}
		})
	}
}

func TestNewStringHeapReferencedEntries(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name      string
		data      []byte
		wantError string
	}{
		{"terminated-with-garbage", []byte{0, 'x', 0, 0xfe}, ""},
		{"unterminated-reference", []byte{0, 'x', 'y', 0xfe}, "not null-terminated"},
		{"invalid-utf8-reference", []byte{0, 0xff, 0, 0}, "invalid UTF-8"},
	} {
		t.Run(test.name, func(t *testing.T) {
			tables := metadataTables(0, 1, []uint32{1}, 10)
			binary.LittleEndian.PutUint16(tables[30:], 1) // Module.Name.
			binary.LittleEndian.PutUint16(tables[32:], 1) // Module.Mvid.
			m, err := winmd.New(metadataPE(t, metadataRoot(
				metadataStream{"#~", tables},
				metadataStream{"#Strings", test.data},
				metadataStream{"#GUID", bytes.Repeat([]byte{1}, 16)},
			)))
			if err != nil {
				t.Fatal(err)
			}
			module, err := m.Tables.Module.At(0)
			if test.wantError != "" {
				if err == nil || !strings.Contains(err.Error(), test.wantError) {
					t.Fatalf("Module.At() error = %v; want %q", err, test.wantError)
				}
			} else if err != nil || module.Name.String() != "x" {
				t.Fatalf("Module.Name = %q, %v; want x, nil", module.Name.String(), err)
			}
		})
	}
}

func TestNewCLIDirectoryBounds(t *testing.T) {
	t.Parallel()
	metadata := metadataRoot(metadataStream{"#~", metadataTables(0, 0, nil, 0)})
	for _, pe64 := range []bool{false, true} {
		t.Run(fmt.Sprintf("pe64-%v", pe64), func(t *testing.T) {
			for _, size := range []uint32{0, 4, 8, 12, 16, 71, 72, uint32(72 + len(metadata) + 1), 0xffffffff} {
				t.Run(fmt.Sprint(size), func(t *testing.T) {
					f := metadataPEForArch(t, metadata, uint32(len(metadata)), pe64)
					if pe64 {
						f.OptionalHeader.(*pe.OptionalHeader64).DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR].Size = size
					} else {
						f.OptionalHeader.(*pe.OptionalHeader32).DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR].Size = size
					}
					m, err := winmd.New(f)
					if size == 72 {
						if err != nil || m == nil {
							t.Fatalf("valid CLI directory = %+v, %v", m, err)
						}
					} else if !errors.Is(err, io.ErrUnexpectedEOF) || m != nil {
						t.Fatalf("CLI directory size %d: metadata=%+v, error=%v; want unexpected EOF", size, m, err)
					}
				})
			}
			t.Run("offset-plus-size-exceeds-section", func(t *testing.T) {
				f := metadataPEForArch(t, metadata, uint32(len(metadata)), pe64)
				section := f.Sections[0]
				rva := section.VirtualAddress + section.Size - 71
				if pe64 {
					f.OptionalHeader.(*pe.OptionalHeader64).DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR].VirtualAddress = rva
				} else {
					f.OptionalHeader.(*pe.OptionalHeader32).DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR].VirtualAddress = rva
				}
				if _, err := winmd.New(f); !errors.Is(err, io.ErrUnexpectedEOF) || !strings.Contains(err.Error(), "COM descriptor directory exceeds") {
					t.Fatalf("directory crossing the section end: %v", err)
				}
			})
		})
	}
}

func TestNewCLIHeaderSize(t *testing.T) {
	t.Parallel()
	metadata := metadataRoot(metadataStream{"#~", metadataTables(0, 0, nil, 0)})
	for _, pe64 := range []bool{false, true} {
		t.Run(fmt.Sprintf("pe64-%v", pe64), func(t *testing.T) {
			for _, size := range []uint32{0, 4, 16, 71, 72, 73, 0xffffffff} {
				t.Run(fmt.Sprint(size), func(t *testing.T) {
					f := metadataPEForArch(t, metadata, uint32(len(metadata)), pe64)
					data, err := f.Sections[0].Data()
					if err != nil {
						t.Fatal(err)
					}
					binary.LittleEndian.PutUint32(data, size)
					f.Sections[0].ReaderAt = bytes.NewReader(data)
					m, err := winmd.New(f)
					if size == 72 {
						if err != nil || m == nil {
							t.Fatalf("valid CLI header = %+v, %v", m, err)
						}
					} else if err == nil || m != nil || !strings.Contains(err.Error(), "invalid CLI header size") {
						t.Fatalf("CLI header size %d: metadata=%+v, error=%v", size, m, err)
					}
				})
			}
		})
	}
}

func TestNewCLIHeaderTruncated(t *testing.T) {
	t.Parallel()
	metadata := metadataRoot(metadataStream{"#~", metadataTables(0, 0, nil, 0)})
	for _, size := range []int{0, 4, 8, 12, 16, 40, 71} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			f := metadataPE(t, metadata)
			data, err := f.Sections[0].Data()
			if err != nil {
				t.Fatal(err)
			}
			f.Sections[0].ReaderAt = bytes.NewReader(data[:size])
			want := io.ErrUnexpectedEOF
			if size == 0 {
				want = io.EOF
			}
			if _, err := winmd.New(f); !errors.Is(err, want) || !strings.Contains(err.Error(), "failure to read the CLI header") {
				t.Fatalf("truncated CLI header error = %v; want %v", err, want)
			}
		})
	}
}

func TestNewMetadataDirectoryBounds(t *testing.T) {
	t.Parallel()
	metadata := metadataRoot(
		metadataStream{"#~", metadataTables(0, 0, nil, 0)},
		metadataStream{"#Blob", []byte{0, 1, 2, 3}},
	)
	end := uint32(len(metadata))
	for _, size := range []uint32{0, 3, 16, 31, 41, end - 4, end - 1, end, end + 1, 0xffffffff} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			m, err := winmd.New(metadataPEWithDirectorySize(t, metadata, size))
			if size != end {
				if err == nil || m != nil {
					t.Fatalf("directory size %d accepted: metadata=%+v, error=%v", size, m, err)
				}
			} else if err != nil || m.Tables == nil || !bytes.Equal(m.Blob, []byte{0, 1, 2, 3}) {
				t.Fatalf("exact directory boundary failed: metadata=%+v, error=%v", m, err)
			}
		})
	}
	t.Run("section-tail-outside-metadata", func(t *testing.T) {
		section := append(bytes.Clone(metadata), bytes.Repeat([]byte{0xfe}, 16)...)
		m, err := winmd.New(metadataPEWithDirectorySize(t, section, end))
		if err != nil || !bytes.Equal(m.Blob, []byte{0, 1, 2, 3}) {
			t.Fatalf("metadata within larger section: metadata=%+v, error=%v", m, err)
		}
	})
}

func TestNewStreamDirectoryBounds(t *testing.T) {
	t.Parallel()
	for _, name := range []string{"#Blob", "#Unknown"} {
		t.Run(name, func(t *testing.T) {
			metadata := metadataRoot(
				metadataStream{"#~", metadataTables(0, 0, nil, 0)},
				metadataStream{name, []byte{0, 1, 2, 3}},
			)
			end := uint32(len(metadata))
			headerOffset := len(metadataRoot()) + 12 // Skip the #~ stream header.
			for _, test := range []struct {
				name    string
				offset  uint32
				size    uint32
				wantErr bool
			}{
				{"exact-end", end - 4, 4, false},
				{"empty-at-end", end, 0, false},
				{"empty-past-end", end + 1, 0, true},
				{"crosses-end", end - 3, 4, true},
				{"offset-wrap", 0xfffffff0, 0x20, true},
				{"size-wrap", end - 4, 0xffffffff, true},
			} {
				t.Run(test.name, func(t *testing.T) {
					data := bytes.Clone(metadata)
					binary.LittleEndian.PutUint32(data[headerOffset:], test.offset)
					binary.LittleEndian.PutUint32(data[headerOffset+4:], test.size)
					_, err := winmd.New(metadataPE(t, data))
					if test.wantErr {
						if !errors.Is(err, io.ErrUnexpectedEOF) {
							t.Fatalf("out-of-range stream error = %v; want unexpected EOF", err)
						}
					} else if err != nil {
						t.Fatal(err)
					}
				})
			}
		})
	}
}

func TestNewVersionLength(t *testing.T) {
	t.Parallel()
	// The upper half of uint32 must also be rejected on 32-bit platforms.
	for _, length := range []uint32{257, 260, 0x7fffffff, 0x80000000, 0xffffffff} {
		t.Run(fmt.Sprintf("%#x", length), func(t *testing.T) {
			data := metadataRoot()
			binary.LittleEndian.PutUint32(data[12:], length)
			_, err := winmd.New(metadataPE(t, data))
			want := fmt.Sprintf("padded string length (%d) is higher than the maximum length (256)", length)
			if err == nil || !strings.Contains(err.Error(), want) {
				t.Fatalf("New() error = %v; want %q", err, want)
			}
		})
	}
}

func TestNewVersionPadding(t *testing.T) {
	t.Parallel()
	for _, length := range []int{0, 1, 3, 4, 251, 252, 253, 254, 255} {
		t.Run(fmt.Sprint(length), func(t *testing.T) {
			version := strings.Repeat("v", length)
			data := metadataRootVersion(version,
				metadataStream{"#Strings", []byte{0}},
				metadataStream{"#~", metadataTables(0, 0, nil, 0)},
			)
			m, err := winmd.New(metadataPE(t, data))
			if length == 255 { // The terminating NUL also counts towards the limit.
				if err == nil || !strings.Contains(err.Error(), "maximum length (255)") {
					t.Fatalf("New() error = %v; want an overlong version error", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if m.Version != version || !bytes.Equal(m.Strings, []byte{0}) {
				t.Fatalf("New() = version %q, strings %v; want version %q, strings [0]", m.Version, m.Strings, version)
			}
		})
	}
}

func TestNewVersionUnterminated(t *testing.T) {
	t.Parallel()
	for _, length := range []int{0, 4, 256} {
		t.Run(fmt.Sprint(length), func(t *testing.T) {
			data := metadataRoot()[:16]
			binary.LittleEndian.PutUint32(data[12:], uint32(length))
			data = append(data, bytes.Repeat([]byte{'v'}, length)...)
			data = append(data, 0, 0, 0, 0) // Flags and stream count are not part of the string.
			_, err := winmd.New(metadataPE(t, data))
			if err == nil || !strings.Contains(err.Error(), "version string must be null-terminated") {
				t.Fatalf("New() error = %v; want an unterminated version error", err)
			}
		})
	}
}

func TestNewStreamNamePadding(t *testing.T) {
	t.Parallel()
	for length := 1; length < 32; length++ {
		t.Run(fmt.Sprint(length), func(t *testing.T) {
			data := metadataRoot(
				metadataStream{strings.Repeat("a", length), []byte{1, 2, 3, 4}},
				metadataStream{"#Strings", []byte{0}},
				metadataStream{"#~", metadataTables(0, 0, nil, 0)},
			)
			m, err := winmd.New(metadataPE(t, data))
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(m.Strings, []byte{0}) {
				t.Fatalf("second stream = %v; want [0]", m.Strings)
			}
		})
	}
}

func TestNewStreamNameErrors(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name string
		data []byte
		want error
	}{
		{"EOF", nil, io.EOF},
		{"partial-chunk", []byte("#~"), io.ErrUnexpectedEOF},
		{"later-EOF", []byte("abcd"), io.EOF},
		{"later-partial-chunk", []byte("abcdef"), io.ErrUnexpectedEOF},
		{"missing-padding", []byte{'a', 0}, io.ErrUnexpectedEOF},
		{"unterminated", bytes.Repeat([]byte{'a'}, 32), nil},
	} {
		t.Run(test.name, func(t *testing.T) {
			data := metadataRoot()
			binary.LittleEndian.PutUint16(data[len(data)-2:], 1)
			data = append(data, make([]byte, 8)...) // Stream offset and size.
			data = append(data, test.data...)
			_, err := winmd.New(metadataPE(t, data))
			if err == nil {
				t.Fatal("New() accepted a malformed stream name")
			}
			if test.want != nil {
				if !errors.Is(err, test.want) {
					t.Fatalf("New() error = %v; want %v", err, test.want)
				}
			} else if !strings.Contains(err.Error(), "name not found") {
				t.Fatalf("New() error = %v; want an unterminated-name error", err)
			}
		})
	}
}

func TestNewTableValidBits(t *testing.T) {
	t.Parallel()
	// Table numbers supported by the #~ stream, including unexported tables.
	supported := []int{
		0, 1, 2, 4, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21,
		23, 24, 25, 26, 27, 28, 29, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44,
	}
	for bit := range 64 {
		for _, count := range []uint32{0, 1} {
			t.Run(fmt.Sprintf("bit-%d/rows-%d", bit, count), func(t *testing.T) {
				valid := uint64(1) << bit
				tables := metadataTables(0, valid, []uint32{count}, 64)
				_, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables})))
				if slices.Contains(supported, bit) {
					if err != nil {
						t.Fatal(err)
					}
					return
				}
				want := fmt.Sprintf("invalid bit vector of present tables: 0b%b", valid)
				if err == nil || !strings.Contains(err.Error(), want) {
					t.Fatalf("New() error = %v; want %q", err, want)
				}
			})
		}
	}
}

func TestNewTablePayloadBounds(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name        string
		heapSizes   byte
		valid       uint64
		rows        []uint32
		payloadSize int
		wantErr     bool
	}{
		{"empty", 0, 0, nil, 0, false},
		{"exact-row", 0, 1, []uint32{1}, 10, false},
		{"missing-row", 0, 1, []uint32{1}, 0, true},
		{"short-row", 0, 1, []uint32{1}, 9, true},
		{"trailing-padding", 0, 1, []uint32{1}, 12, false},
		{"wide-heap-indices", 7, 1, []uint32{1}, 18, false},
		{"short-wide-heap-indices", 7, 1, []uint32{1}, 17, true},
		{"multiple-tables", 0, 1 | 1<<32, []uint32{1, 1}, 32, false},
		{"short-second-table", 0, 1 | 1<<32, []uint32{1, 1}, 31, true},
		{"unexported-table", 0, 1 << 34, []uint32{1}, 12, false},
		{"short-unexported-table", 0, 1 << 34, []uint32{1}, 11, true},
		{"last-table", 0, 1 << 44, []uint32{1}, 4, false},
		{"short-last-table", 0, 1 << 44, []uint32{1}, 3, true},
		{"large-row-count", 0, 1, []uint32{0xffffffff}, 0, true},
		{"32-bit-size-wraps-to-zero", 0, 1, []uint32{0x80000000}, 0, true},
		{"32-bit-size-wraps-to-payload", 0, 1, []uint32{0x80000001}, 10, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			tables := metadataTables(test.heapSizes, test.valid, test.rows, test.payloadSize)
			m, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables})))
			if test.wantErr {
				if !errors.Is(err, io.ErrUnexpectedEOF) {
					t.Fatalf("New() error = %v; want %v", err, io.ErrUnexpectedEOF)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if m.Tables == nil {
				t.Fatal("New() returned no tables")
			}
			if test.valid&1 != 0 && m.Tables.Module.Len() != test.rows[0] {
				t.Fatalf("Module rows = %d; want %d", m.Tables.Module.Len(), test.rows[0])
			}
		})
	}
}

func TestNewTruncatedTablesHeader(t *testing.T) {
	t.Parallel()
	tables := metadataTables(0, 1, []uint32{1}, 0)
	for length := range len(tables) {
		t.Run(fmt.Sprint(length), func(t *testing.T) {
			_, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables[:length]})))
			if err == nil {
				t.Fatal("New() accepted a truncated tables header")
			}
		})
	}
}

func TestNewFileTableLayout(t *testing.T) {
	t.Parallel()
	for _, heapSizes := range []byte{0, 5} {
		t.Run(fmt.Sprintf("heap-sizes-%d", heapSizes), func(t *testing.T) {
			appendIndex := func(data []byte, index uint32) []byte {
				if heapSizes == 0 {
					return binary.LittleEndian.AppendUint16(data, uint16(index))
				}
				return binary.LittleEndian.AppendUint32(data, index)
			}
			tables := metadataTables(heapSizes, 1<<38|1<<40, []uint32{2, 1}, 0)
			// File.Flags is four bytes, independent of heap index sizes.
			for i := range uint32(2) {
				tables = binary.LittleEndian.AppendUint32(tables, i)
				tables = appendIndex(tables, 1+6*i)
				tables = appendIndex(tables, 1+2*i)
			}
			// A following table verifies the total size of the File table.
			tables = binary.LittleEndian.AppendUint32(tables, 0x12345678)
			tables = binary.LittleEndian.AppendUint32(tables, uint32(winmd.ManifestResourceAttributes_Public))
			tables = appendIndex(tables, 13)
			tables = binary.LittleEndian.AppendUint16(tables, 4) // Implementation: File row 1.
			m, err := winmd.New(metadataPE(t, metadataRoot(
				metadataStream{"#~", tables},
				metadataStream{"#Strings", []byte("\x00a.bin\x00b.bin\x00resource\x00")},
				metadataStream{"#Blob", []byte{0, 1, 0xaa, 1, 0xbb}},
			)))
			if err != nil {
				t.Fatal(err)
			}
			for i, name := range []string{"a.bin", "b.bin"} {
				file, err := m.Tables.File.At(winmd.Index(i))
				if err != nil {
					t.Fatal(err)
				}
				if file.Flags != winmd.FileAttributes(i) || file.Name.String() != name || !bytes.Equal(file.HashValue, []byte{0xaa + byte(i)*0x11}) {
					t.Fatalf("File row %d = %+v; want flags %d, name %q, hash %x", i, file, i, name, 0xaa+byte(i)*0x11)
				}
			}
			resource, err := m.Tables.ManifestResource.At(0)
			if err != nil {
				t.Fatal(err)
			}
			if resource.Offset != 0x12345678 || resource.Name.String() != "resource" || resource.Implementation.Tag != winmd.Implementation_File || resource.Implementation.Index != 0 {
				t.Fatalf("following ManifestResource row is misaligned: %+v", resource)
			}
		})
	}
}

func TestNewRegularTableIndex(t *testing.T) {
	t.Parallel()
	for _, parent := range []uint16{0, 1, 2} {
		t.Run(fmt.Sprint(parent), func(t *testing.T) {
			tables := metadataTables(0, 1<<2|1<<15, []uint32{1, 1}, 0)
			typeDef := make([]byte, 14)
			binary.LittleEndian.PutUint16(typeDef[10:], 1)
			binary.LittleEndian.PutUint16(typeDef[12:], 1)
			tables = append(tables, typeDef...)
			tables = binary.LittleEndian.AppendUint16(tables, 0)
			tables = binary.LittleEndian.AppendUint32(tables, 0)
			tables = binary.LittleEndian.AppendUint16(tables, parent)
			m, err := winmd.New(metadataPE(t, metadataRoot(metadataStream{"#~", tables})))
			if err != nil {
				t.Fatal(err)
			}
			class, err := m.Tables.ClassLayout.At(0)
			if parent == 1 {
				if err != nil || class.Parent != 0 {
					t.Fatalf("ClassLayout.Parent = %d, %v; want 0, nil", class.Parent, err)
				}
			} else if err == nil {
				t.Fatalf("ClassLayout.Parent accepted invalid encoded index %d", parent)
			}
		})
	}
}
