// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import "debug/pe"

// CLIHeader contains all of the runtime-specific data entries and other information.
// Defined in §II.25.3.3.
type CLIHeader struct {
	Size                    uint32
	MajorRuntimeVersion     uint16
	MinorRuntimeVersion     uint16
	Metadata                pe.DataDirectory
	Flags                   COMIMAGE_FLAGS
	EntryPointToken         uint32
	Resources               pe.DataDirectory
	StrongNameSignature     pe.DataDirectory
	CodeManagerTable        pe.DataDirectory
	VTableFixups            pe.DataDirectory
	ExportAddressTableJumps pe.DataDirectory
	ManagedNativeHeader     pe.DataDirectory
}

// COMIMAGE_FLAGS describe the runtime image.
// Defined in §II.25.3.3.1.
type COMIMAGE_FLAGS uint32

const (
	COMIMAGE_FLAGS_ILONLY            COMIMAGE_FLAGS = 0x00000001
	COMIMAGE_FLAGS_32BITREQUIRED     COMIMAGE_FLAGS = 0x00000002
	COMIMAGE_FLAGS_STRONGNAMESIGNED  COMIMAGE_FLAGS = 0x00000008
	COMIMAGE_FLAGS_NATIVE_ENTRYPOINT COMIMAGE_FLAGS = 0x00000010
	COMIMAGE_FLAGS_TRACKDEBUGDATA    COMIMAGE_FLAGS = 0x00010000
)

// MetadataHeader is the root of the metadata.
// Defined in §II.24.2.1.
type MetadataHeader struct {
	Signature    uint32
	MajorVersion uint16
	MinorVersion uint16
	Reserved     uint32
	Version      string
	Flags        uint16
}

// StreamHeader gives the names, and the position and
// length of a particular table or heap.
// Defined in §II.24.2.2.
type StreamHeader struct {
	Offset uint32
	Size   uint32
	Name   string
}
