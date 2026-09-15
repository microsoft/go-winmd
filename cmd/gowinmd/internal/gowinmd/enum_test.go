// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package gowinmd

import (
	"bytes"
	"debug/pe"
	"encoding/binary"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

// enumMetadata builds a small in-memory PE containing System.Enum and Test.Mode.
// Mode's base uses either a TypeDef or TypeRef to the same System.Enum type.
func enumMetadata(t *testing.T, typeDefBase bool, kind winmd.ElementType, value []byte) *winmd.Metadata {
	t.Helper()
	write := func(b *bytes.Buffer, values ...any) {
		t.Helper()
		for _, value := range values {
			if err := binary.Write(b, binary.LittleEndian, value); err != nil {
				t.Fatal(err)
			}
		}
	}
	strs := []byte{0}
	addString := func(s string) uint16 {
		offset := uint16(len(strs))
		strs = append(strs, s...)
		strs = append(strs, 0)
		return offset
	}
	module, system, enum := addString("Test.winmd"), addString("System"), addString("Enum")
	namespace, name := addString("Test"), addString("Mode")
	backingName, memberName := addString("MyValue"), addString("Member")
	blobs := []byte{0}
	addBlob := func(data []byte) uint16 {
		offset := uint16(len(blobs))
		blobs = append(blobs, byte(len(data)))
		blobs = append(blobs, data...)
		return offset
	}
	backingSig := addBlob([]byte{0x06, byte(kind)})
	memberSig := addBlob([]byte{0x06, byte(winmd.ElementType_VALUETYPE), 3 << 2})
	constantValue := addBlob(value)
	blobs = append(blobs, 0)

	var tables bytes.Buffer
	// #~ header and row counts: Module, TypeRef, TypeDef, Field, Constant.
	write(&tables, uint32(0), byte(2), byte(0), byte(0), byte(1),
		uint64(1<<0|1<<1|1<<2|1<<4|1<<11), uint64(0), []uint32{1, 1, 3, 2, 1})
	write(&tables, uint16(0), module, uint16(0), uint16(0), uint16(0))
	write(&tables, uint16(1<<2), enum, system) // TypeRef scoped to Module 1.
	write(&tables, uint32(0), addString("<Module>"), uint16(0), uint16(0), uint16(1), uint16(1))
	write(&tables, uint32(winmd.TypeAttributes_Public), enum, system, uint16(0), uint16(1), uint16(1))
	extends := uint16(1<<2 | 1) // TypeRef 1.
	if typeDefBase {
		extends = 2 << 2 // TypeDef 2.
	}
	write(&tables, uint32(winmd.TypeAttributes_Public|winmd.TypeAttributes_Sealed), name, namespace, extends, uint16(1), uint16(1))
	write(&tables, uint16(winmd.FieldAttributes_Public), backingName, backingSig)
	write(&tables, uint16(winmd.FieldAttributes_Public|winmd.FieldAttributes_Static|winmd.FieldAttributes_Literal|winmd.FieldAttributes_HasDefault), memberName, memberSig)
	write(&tables, byte(kind), byte(0), uint16(2<<2), constantValue)

	streams := []struct {
		name string
		data []byte
	}{{"#~", tables.Bytes()}, {"#Strings", strs}, {"#Blob", blobs}}
	var root bytes.Buffer
	write(&root, uint32(0x424A5342), uint16(1), uint16(1), uint32(0), uint32(4), []byte{'v', '1', 0, 0}, uint16(0), uint16(len(streams)))
	offset := root.Len()
	for _, stream := range streams {
		offset += 8 + (len(stream.name)+4)&^3
	}
	for _, stream := range streams {
		write(&root, uint32(offset), uint32(len(stream.data)))
		root.WriteString(stream.name)
		root.Write(make([]byte, (len(stream.name)+4)&^3-len(stream.name)))
		offset += (len(stream.data) + 3) &^ 3
	}
	for _, stream := range streams {
		root.Write(stream.data)
		root.Write(make([]byte, (4-root.Len()%4)%4))
	}

	const sectionRVA = 0x2000
	var section bytes.Buffer
	write(&section, uint32(72), uint16(2), uint16(5), uint32(sectionRVA+72), uint32(root.Len()))
	section.Write(make([]byte, 72-section.Len()))
	section.Write(root.Bytes())
	optional := pe.OptionalHeader32{Magic: 0x10B, NumberOfRvaAndSizes: 16}
	optional.DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR] = pe.DataDirectory{VirtualAddress: sectionRVA, Size: 72}
	header := pe.FileHeader{Machine: pe.IMAGE_FILE_MACHINE_I386, NumberOfSections: 1, SizeOfOptionalHeader: uint16(binary.Size(optional))}
	sectionHeader := pe.SectionHeader32{
		Name: [8]byte{'.', 'm', 'e', 't', 'a'}, VirtualAddress: sectionRVA,
		VirtualSize: uint32(section.Len()), SizeOfRawData: uint32(section.Len()),
		PointerToRawData: uint32(binary.Size(header) + binary.Size(optional) + binary.Size(pe.SectionHeader32{})),
	}
	var image bytes.Buffer
	write(&image, header, optional, sectionHeader)
	image.Write(section.Bytes())
	file, err := pe.NewFile(bytes.NewReader(image.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	metadata, err := winmd.New(file)
	if err != nil {
		t.Fatal(err)
	}
	return metadata
}

func TestEnumBaseDispatch(t *testing.T) {
	for _, test := range []struct {
		name       string
		kind       winmd.ElementType
		value      []byte
		underlying string
		literal    string
		size       uint32
	}{
		{"int64", winmd.ElementType_I8, []byte{3, 0, 0, 0, 0, 0, 0, 0}, "int64", "0x3", 8},
		{"boolean", winmd.ElementType_BOOLEAN, []byte{1}, "bool", "true", 1},
		{"character", winmd.ElementType_CHAR, []byte{0xAC, 0x20}, "uint16", "0x20ac", 2},
	} {
		t.Run(test.name, func(t *testing.T) {
			for _, base := range []string{"TypeRef", "TypeDef"} {
				t.Run(base, func(t *testing.T) {
					metadata := enumMetadata(t, base == "TypeDef", test.kind, test.value)
					context, err := NewContext(metadata)
					if err != nil {
						t.Fatal(err)
					}
					if err := context.SelectTypeDef("Test", "Mode", ""); err != nil {
						t.Fatal(err)
					}
					def := context.resolvedDefsByIndex[2]
					for _, arch := range []Arch{Arch386, ArchAMD64, ArchARM64} {
						var output strings.Builder
						if err := context.writeTypeDef(&output, def, arch); err != nil {
							t.Fatal(err)
						}
						want := "type Mode " + test.underlying + "\n\nconst (\n\tMember Mode = " + test.literal + "\n)\n"
						if output.String() != want {
							t.Fatalf("writeTypeDef() = %q; want %q", output.String(), want)
						}
						fset := token.NewFileSet()
						file, err := parser.ParseFile(fset, "", "package test\n"+output.String(), 0)
						if err != nil {
							t.Fatal(err)
						}
						config := types.Config{Sizes: types.SizesFor("gc", arch.String())}
						if _, err := config.Check("test", fset, []*ast.File{file}, nil); err != nil {
							t.Fatal(err)
						}
						layout, err := context.resolvedDefABITypeLayout(def, arch, nil)
						if want := scalarABITypeLayout(test.size, test.size, arch); err != nil || layout != want {
							t.Fatalf("layout = (%#v, %v); want %#v", layout, err, want)
						}
						fingerprint, err := context.typeDefABILayoutFingerprint(def, arch)
						if err != nil || fingerprint.size != test.size || fingerprint.align != test.size || len(fingerprint.fields) != 0 {
							t.Fatalf("fingerprint = (%#v, %v); want a scalar of size/alignment %d", fingerprint, err, test.size)
						}
					}
				})
			}
		})
	}
}
