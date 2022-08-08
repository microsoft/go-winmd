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
	ds := sectionByRVA(pefile, comdd.VirtualAddress)
	if ds == nil {
		return nil, errors.New("COM descriptor directory table is missing")
	}

	// read COM descriptor directory table might be in a large section,
	// we better don't call ds.Data()
	r := ds.Open()

	// seek to the COM descriptor data directory virtual address.
	_, err := r.Seek(int64(comdd.VirtualAddress-ds.VirtualAddress), io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("failure to seek to the COM descriptor data directory root: %v", err)
	}

	read := func(data interface{}) bool {
		err = binary.Read(r, binary.LittleEndian, data)
		return err == nil
	}
	readDataDirectory := func(data *pe.DataDirectory) bool {
		return read(&data.VirtualAddress) && read(&data.Size)
	}

	var hdr CLIHeader
	if !read(&hdr.Size) ||
		!read(&hdr.MajorRuntimeVersion) ||
		!read(&hdr.MinorRuntimeVersion) ||
		!readDataDirectory(&hdr.Metadata) ||
		!read(&hdr.Flags) ||
		!read(&hdr.EntryPointToken) ||
		!readDataDirectory(&hdr.Resources) ||
		!readDataDirectory(&hdr.StrongNameSignature) ||
		!readDataDirectory(&hdr.CodeManagerTable) ||
		!readDataDirectory(&hdr.VTableFixups) ||
		!readDataDirectory(&hdr.ExportAddressTableJumps) ||
		!readDataDirectory(&hdr.ManagedNativeHeader) {
		return nil, fmt.Errorf("failure to read the CLI header: %v", err)
	}
	return &hdr, nil
}

// readMetadata reads the Metadata from pefile.
func readMetadata(pefile *pe.File, rva uint32) (*MetadataHeader, []*Stream, error) {
	// figure out which section contains the metadata.
	ds := sectionByRVA(pefile, rva)
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

// sectionByRVA returns the section which contains rva.
func sectionByRVA(pefile *pe.File, rva uint32) *pe.Section {
	var ds *pe.Section
	for _, s := range pefile.Sections {
		if s.VirtualAddress <= rva && rva < s.VirtualAddress+s.VirtualSize {
			ds = s
			break
		}
	}
	return ds
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
