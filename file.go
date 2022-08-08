// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"bytes"
	"debug/pe"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// A File represents an open Windows Metadata file.
type File struct {
	CLIHeader      CLIHeader
	MetadataHeader MetadataHeader
	Streams        []*Stream
}

// NewFile creates a new File from an underlying PE file.
func NewFile(pefile *pe.File) (*File, error) {
	if pefile.OptionalHeader == nil {
		return nil, errors.New("pe optional header is required to parse as winmd, but it is missing")
	}

	_, pe64 := pefile.OptionalHeader.(*pe.OptionalHeader64)

	// grab the number of data directory entries.
	var ddLength uint32
	if pe64 {
		ddLength = pefile.OptionalHeader.(*pe.OptionalHeader64).NumberOfRvaAndSizes
	} else {
		ddLength = pefile.OptionalHeader.(*pe.OptionalHeader32).NumberOfRvaAndSizes
	}

	// check that the length of data directory entries is large
	// enough to include the COM descriptor directory.
	if ddLength < pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR+1 {
		return nil, fmt.Errorf("data directory entries length (%d) is less than minimum length (%d) to include the COM descriptor directory", ddLength, pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR)
	}

	clihdr, err := readCLIHeader(pefile, pe64)
	if err != nil {
		return nil, err
	}
	mdhdr, streams, err := readMetadata(pefile, clihdr.Metadata.VirtualAddress)
	if err != nil {
		return nil, err
	}
	return &File{
		CLIHeader:      *clihdr,
		MetadataHeader: *mdhdr,
		Streams:        streams,
	}, nil
}

// Stream returns the stream with the given name, or nil if no such
// stream exists.
func (m *File) Stream(name string) *Stream {
	for _, s := range m.Streams {
		if s.Name == name {
			return s
		}
	}
	return nil
}

// readCLIHeader reads the CLI header from pefile.
func readCLIHeader(pefile *pe.File, pe64 bool) (*CLIHeader, error) {
	// grab the com descriptor data directory entry.
	var comdd pe.DataDirectory
	if pe64 {
		comdd = pefile.OptionalHeader.(*pe.OptionalHeader64).DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR]
	} else {
		comdd = pefile.OptionalHeader.(*pe.OptionalHeader32).DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR]
	}

	// figure out which section contains the COM descriptor directory table
	var ds *pe.Section
	for _, s := range pefile.Sections {
		if s.VirtualAddress <= comdd.VirtualAddress && comdd.VirtualAddress < s.VirtualAddress+s.VirtualSize {
			ds = s
			break
		}
	}

	// didn't find a section.
	if ds == nil {
		return nil, errors.New("COM descriptor directory table is missing")
	}

	// read COM descriptor directory table.
	d, err := ds.Data()
	if err != nil {
		return nil, err
	}

	// seek to the virtual address specified in the COM descriptor data directory.
	d = d[comdd.VirtualAddress-ds.VirtualAddress:]

	var hdr CLIHeader
	// verify there is enough space for the CLI header.
	cliHeaderSize := binary.Size(hdr)
	if len(d) < cliHeaderSize {
		return nil, fmt.Errorf("COM descriptor data directory size (%d) is less than minimum size (%d) for CLI header", len(d), cliHeaderSize)
	}

	// parse CLI header.
	hdr.Size = binary.LittleEndian.Uint32(d[0:4])
	if hdr.Size < uint32(cliHeaderSize) {
		return nil, fmt.Errorf("size field in CLI header (%d) is too smaller than the minimum size (%d)", hdr.Size, cliHeaderSize)
	}
	hdr.MajorRuntimeVersion = binary.LittleEndian.Uint16(d[4:6])
	hdr.MinorRuntimeVersion = binary.LittleEndian.Uint16(d[6:8])
	hdr.Metadata = readDataDirectory(d[8:16])
	hdr.Flags = COMIMAGE_FLAGS(binary.LittleEndian.Uint32(d[16:20]))
	hdr.EntryPointToken = binary.LittleEndian.Uint32(d[20:24])
	hdr.Resources = readDataDirectory(d[24:32])
	hdr.StrongNameSignature = readDataDirectory(d[32:40])
	hdr.CodeManagerTable = readDataDirectory(d[40:48])
	hdr.VTableFixups = readDataDirectory(d[48:56])
	hdr.ExportAddressTableJumps = readDataDirectory(d[56:64])
	hdr.ManagedNativeHeader = readDataDirectory(d[64:72])
	return &hdr, nil
}

// readMetadata reads the Metadata from pefile.
// f.CLIHeader must be already filled.
func readMetadata(pefile *pe.File, rva uint32) (*MetadataHeader, []*Stream, error) {
	// figure out which section contains the metadata.
	var ds *pe.Section
	for _, s := range pefile.Sections {
		if s.VirtualAddress <= rva && rva < s.VirtualAddress+s.VirtualSize {
			ds = s
			break
		}
	}

	// didn't find a section.
	if ds == nil {
		return nil, nil, errors.New("metadata section is missing")
	}

	// The metadata section can be huge, we better don't call ds.Data()
	r := ds.Open()

	// seek to the virtual address specified in the COM descriptor data directory.
	rootOffset := int64(rva - ds.VirtualAddress)
	_, err := r.Seek(rootOffset, io.SeekStart)
	if err != nil {
		return nil, nil, fmt.Errorf("failure to seek to the metadata root: %v", err)
	}

	read := func(data interface{}) bool {
		err = binary.Read(r, binary.LittleEndian, data)
		return err == nil
	}
	readStr := func(n int, data *string) bool {
		buf := make([]byte, n)
		err = binary.Read(r, binary.LittleEndian, buf)
		if err == nil {
			*data = cstring(buf)
		}
		return err == nil
	}

	// parse hdr header.
	var hdr MetadataHeader
	if !read(&hdr.Signature) {
		return nil, nil, fmt.Errorf("failure to read the metadata header signature: %v", err)
	}
	// magic signature from II.24.2.1
	const signature = 0x424A5342
	if hdr.Signature != signature {
		return nil, nil, fmt.Errorf("metadata header signature (%#X) must be (%#X)", hdr.Signature, signature)
	}
	var streamsCount uint16
	var cstringLength uint32
	if !read(&hdr.MajorVersion) ||
		!read(&hdr.MinorVersion) ||
		!read(&hdr.Reserved) ||
		!read(&cstringLength) ||
		!readStr(int(cstringLength), &hdr.Version) ||
		!read(&hdr.Flags) ||
		!read(&streamsCount) {
		return nil, nil, fmt.Errorf("failure to read the metadata header: %v", err)
	}
	streams := make([]*Stream, streamsCount)
	for i := 0; i < int(streamsCount); i++ {
		var s Stream
		if !read(&s.Offset) ||
			!read(&s.Size) ||
			!readStr(32, &s.Name) { // the stream header name is limited to 32 characters.
			return nil, nil, fmt.Errorf("failure to read the stream header (%d): %v", i, err)
		}
		s.sr = io.NewSectionReader(ds, rootOffset+int64(s.Offset), int64(s.Size))
		s.ReaderAt = s.sr
		streams[i] = &s
		// name is null-terminated and padded to the next 4-byte boundary when layed out in disk.
		l := len(s.Name) + 1
		padding := (4 - l) % 4
		// seek backwards, we probably read more than necessary.
		_, err = r.Seek(-int64(32-l-padding), io.SeekCurrent)
		if err != nil {
			return nil, nil, fmt.Errorf("failure to seek to the stream header (%d): %v", i, err)
		}
	}
	return &hdr, streams, nil
}

func readDataDirectory(d []byte) pe.DataDirectory {
	return pe.DataDirectory{
		VirtualAddress: binary.LittleEndian.Uint32(d[0:4]),
		Size:           binary.LittleEndian.Uint32(d[4:8]),
	}
}

// cstring converts UTF8 byte sequence b to string.
// It stops once it finds 0 or reaches end of b.
func cstring(b []byte) string {
	i := bytes.IndexByte(b, 0)
	if i == -1 {
		i = len(b)
	}
	return string(b[:i])
}
