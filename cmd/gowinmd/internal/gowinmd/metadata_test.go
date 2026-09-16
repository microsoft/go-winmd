// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package gowinmd

import (
	"bytes"
	"debug/pe"
	"encoding/binary"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

type generatorTestOptions struct {
	tailCount        byte
	duplicateNames   bool
	invalidUnusedRef bool
	nestedVisibility winmd.TypeVisibility
	extraTypeDefs    []qualifiedTypeName
	extraTypeRefs    []qualifiedTypeName
}

// generatorTestMetadata constructs a small module containing Sample with a
// uint32 field and a byte-array field, plus an optional nested type. Repeated
// names can deliberately use different heap entries to test textual matching.
func generatorTestMetadata(t *testing.T, opts generatorTestOptions) *winmd.Metadata {
	t.Helper()
	stringsHeap := []byte{0}
	addString := func(s string) uint16 {
		offset := uint16(len(stringsHeap))
		stringsHeap = append(stringsHeap, s...)
		stringsHeap = append(stringsHeap, 0)
		return offset
	}
	moduleName := addString("Module")
	moduleTypeName := addString("<Module>")
	namespace, name := addString("Test"), addString("Sample")
	hiddenName := addString("Hidden")
	refNamespace, refName, hiddenRefName := namespace, name, hiddenName
	if opts.duplicateNames {
		// Reference valid suffixes after unreachable, invalid UTF-8 bytes.
		// Canonicalization must inspect references, not scan physical entries.
		refNamespace = addString("\xffTest") + 1
		refName = addString("\xffSample") + 1
		hiddenRefName = addString("\xffHidden") + 1
	}
	if opts.invalidUnusedRef {
		refName = 0xffff
	}
	system, valueType := addString("System"), addString("ValueType")
	assemblyName := addString("System.Runtime")
	lengthName, tailName := addString("Length"), addString("Tail")
	var payload bytes.Buffer
	write := func(values ...any) {
		for _, value := range values {
			if err := binary.Write(&payload, binary.LittleEndian, value); err != nil {
				t.Fatal(err)
			}
		}
	}
	write(uint16(0), moduleName, uint16(1), uint16(0), uint16(0)) // Module.
	write(uint16(6), valueType, system)                           // System.ValueType TypeRef.
	write(uint16(4), refName, refNamespace)                       // Local Sample TypeRef.
	if opts.nestedVisibility != 0 {
		write(uint16(11), hiddenRefName, uint16(0)) // Nested TypeRef in Sample.
	}
	for _, extra := range opts.extraTypeRefs {
		write(uint16(4), addString(extra.Name), addString(extra.Namespace))
	}
	write(uint32(0), moduleTypeName, uint16(0), uint16(0), uint16(1), uint16(1))
	publicType := uint32(winmd.TypeVisibility_Public) | uint32(winmd.TypeLayout_SequentialLayout) | uint32(winmd.TypeFlags_Sealed)
	write(publicType, name, namespace, uint16(5), uint16(1), uint16(1))
	for _, extra := range opts.extraTypeDefs {
		write(publicType, addString(extra.Name), addString(extra.Namespace), uint16(5), uint16(3), uint16(1))
	}
	if opts.nestedVisibility != 0 {
		write(uint32(opts.nestedVisibility)|uint32(winmd.TypeFlags_Sealed), hiddenName, uint16(0), uint16(5), uint16(3), uint16(1))
	}
	write(uint16(6), lengthName, uint16(1)) // Length: U4.
	write(uint16(6), tailName, uint16(4))   // Tail: U1[tailCount].
	write(uint16(4), uint16(0), uint16(0), uint16(0), uint32(0), uint16(0), assemblyName, uint16(0), uint16(0))
	valid := uint64(1<<0 | 1<<1 | 1<<2 | 1<<4 | 1<<35)
	rows := []uint32{1, 2 + uint32(len(opts.extraTypeRefs)), 2 + uint32(len(opts.extraTypeDefs)), 2, 1}
	if opts.nestedVisibility != 0 {
		valid |= 1 << 41
		rows[1]++
		rows[2]++
		rows = append(rows, 1)
		write(uint16(rows[2]), uint16(2)) // Hidden nested in Sample.
	}
	tables := make([]byte, 24)
	tables[4], tables[7] = 2, 1
	binary.LittleEndian.PutUint64(tables[8:], valid)
	for _, count := range rows {
		tables = binary.LittleEndian.AppendUint32(tables, count)
	}
	tables = append(tables, payload.Bytes()...)
	streams := []struct {
		name string
		data []byte
	}{
		{"#~", tables},
		{"#Strings", stringsHeap},
		{"#Blob", []byte{0, 2, 6, 9, 7, 6, 0x14, 5, 1, 1, opts.tailCount, 0}},
		{"#GUID", bytes.Repeat([]byte{1}, 16)},
	}
	metadata := make([]byte, 16)
	binary.LittleEndian.PutUint32(metadata, 0x424a5342)
	binary.LittleEndian.PutUint16(metadata[4:], 1)
	binary.LittleEndian.PutUint16(metadata[6:], 1)
	const version = "v4.0.30319\x00\x00"
	binary.LittleEndian.PutUint32(metadata[12:], uint32(len(version)))
	metadata = append(metadata, version...)
	metadata = binary.LittleEndian.AppendUint16(metadata, 0)
	metadata = binary.LittleEndian.AppendUint16(metadata, uint16(len(streams)))
	offset := len(metadata)
	for _, stream := range streams {
		offset += 8 + (len(stream.name)+4)&^3
	}
	for _, stream := range streams {
		size := (len(stream.data) + 3) &^ 3
		metadata = binary.LittleEndian.AppendUint32(metadata, uint32(offset))
		metadata = binary.LittleEndian.AppendUint32(metadata, uint32(size))
		metadata = append(metadata, stream.name...)
		metadata = append(metadata, make([]byte, 4-len(stream.name)%4)...)
		offset += size
	}
	for _, stream := range streams {
		metadata = append(metadata, stream.data...)
		metadata = append(metadata, make([]byte, (4-len(stream.data)%4)%4)...)
	}
	section := make([]byte, 72)
	binary.LittleEndian.PutUint32(section, 72)
	binary.LittleEndian.PutUint16(section[4:], 2)
	binary.LittleEndian.PutUint32(section[8:], 0x2000+72)
	binary.LittleEndian.PutUint32(section[12:], uint32(len(metadata)))
	section = append(section, metadata...)
	optional := pe.OptionalHeader32{Magic: 0x10b, NumberOfRvaAndSizes: 16}
	optional.DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR] = pe.DataDirectory{VirtualAddress: 0x2000, Size: 72}
	var image bytes.Buffer
	for _, value := range []any{
		pe.FileHeader{Machine: pe.IMAGE_FILE_MACHINE_I386, NumberOfSections: 1, SizeOfOptionalHeader: uint16(binary.Size(optional))},
		optional,
		pe.SectionHeader32{Name: [8]byte{'.', 't', 'e', 'x', 't'}, VirtualSize: uint32(len(section)), VirtualAddress: 0x2000, SizeOfRawData: uint32(len(section)), PointerToRawData: 0x200},
	} {
		if err := binary.Write(&image, binary.LittleEndian, value); err != nil {
			t.Fatal(err)
		}
	}
	image.Write(make([]byte, 0x200-image.Len()))
	image.Write(section)
	f, err := pe.NewFile(bytes.NewReader(image.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	m, err := winmd.New(f)
	if err != nil {
		t.Fatal(err)
	}
	return m
}

func TestResolveTypeRefDuplicateHeapStrings(t *testing.T) {
	t.Parallel()
	for _, duplicate := range []bool{false, true} {
		t.Run(fmt.Sprintf("duplicate-%v", duplicate), func(t *testing.T) {
			c, err := NewContext(generatorTestMetadata(t, generatorTestOptions{
				tailCount: 1, duplicateNames: duplicate, nestedVisibility: winmd.TypeVisibility_NestedPublic,
			}))
			if err != nil {
				t.Fatal(err)
			}
			if len(c.typeDefCache.aliases) != 0 {
				t.Fatal("unused TypeRefs were canonicalized eagerly")
			}
			for range 2 { // Test both unresolved and resolved lookup paths.
				for _, index := range []winmd.Index{1, 2} {
					def, err := c.resolveTypeRef(index, ArchARM64)
					if err != nil || def == nil || def.Index != index {
						t.Fatalf("TypeRef %d = %+v, %v; want TypeDef %d", index, def, err, index)
					}
				}
				// A second lookup must use the cached integer key, not text.
				c.typeDefsByName = nil
			}
			if (len(c.typeDefCache.aliases) != 0) != duplicate {
				t.Fatalf("alias count = %d; duplicate names = %v", len(c.typeDefCache.aliases), duplicate)
			}
		})
	}
}

func TestContextDoesNotDecodeUnusedTypeRefs(t *testing.T) {
	t.Parallel()
	c, err := NewContext(generatorTestMetadata(t, generatorTestOptions{invalidUnusedRef: true}))
	if err != nil {
		t.Fatalf("unused invalid TypeRef was decoded during context creation: %v", err)
	}
	if _, err := c.resolveTypeRef(1, ArchARM64); err == nil {
		t.Fatal("invalid TypeRef was not rejected when used")
	}
}

func TestUnresolvableTypeRefDuplicateNames(t *testing.T) {
	t.Parallel()
	name := qualifiedTypeName{Namespace: "Test", Name: "Missing"}
	c, err := NewContext(generatorTestMetadata(t, generatorTestOptions{extraTypeRefs: []qualifiedTypeName{name, name}}))
	if err != nil {
		t.Fatal(err)
	}
	for _, index := range []winmd.Index{2, 3} {
		sig := winmd.SigType{
			Kind:  winmd.ElementType_CLASS,
			Value: winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]{Tag: winmd.TypeDefOrRefOrSpec_TypeRef, Index: index},
		}
		var output strings.Builder
		if err := c.writeType(&output, &sig, ArchARM64); err != nil {
			t.Fatal(err)
		}
	}
	if names := c.UnresolvableTypeRefs(); len(names) != 1 {
		t.Fatalf("duplicate unresolved references produced %v; want one diagnostic name", names)
	}
}

func TestNestedTypeVisibilityLookup(t *testing.T) {
	t.Parallel()
	for _, visibility := range []winmd.TypeVisibility{
		winmd.TypeVisibility_NestedPublic, winmd.TypeVisibility_NestedPrivate,
		winmd.TypeVisibility_NestedFamily, winmd.TypeVisibility_NestedAssembly,
		winmd.TypeVisibility_NestedFamANDAssem, winmd.TypeVisibility_NestedFamORAssem,
	} {
		t.Run(fmt.Sprintf("%#x", uint32(visibility)), func(t *testing.T) {
			c, err := NewContext(generatorTestMetadata(t, generatorTestOptions{
				tailCount: 1, duplicateNames: true, nestedVisibility: visibility,
			}))
			if err != nil {
				t.Fatal(err)
			}
			if len(c.typeDefsByName[qualifiedTypeName{Name: "Hidden"}]) != 0 {
				t.Fatal("nested type indexed as a top-level definition")
			}
			if err := c.SelectTypeDef("", "Hidden", ""); err == nil {
				t.Fatal("nested type selectable at module scope")
			}
			child, err := c.resolveTypeRef(2, ArchARM64)
			if err != nil || child == nil || child.Index != 2 || child.Parent == nil || child.Parent.Index != 1 {
				t.Fatalf("nested reference = %+v, %v; want child of Sample", child, err)
			}
			if got := c.typeDefCache.get(child.Namespace, child.Name, ArchARM64); got != nil {
				t.Fatal("resolving nested type populated the module-level cache")
			}
			key := c.typeDefCache.canonicalKey(typeDefKey(child.def))
			if _, ok := c.typeDefCache.unresolved[key]; ok {
				t.Fatal("nested type left in unresolved module-level lookup")
			}
		})
	}
}

func TestStructABITrailingZeroSize(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name        string
		count       byte
		classSize   uint32
		onlyZero    bool
		wantSize    uint32
		wantTailPad uint32
		wantErr     bool
	}{
		{name: "nonzero-control", count: 1, wantSize: 8},
		{name: "unrepresentable-zero-tail", count: 0, wantErr: true},
		{name: "zero-tail-with-room", count: 0, classSize: 8, wantSize: 8},
		{name: "zero-tail-with-explicit-padding", count: 0, classSize: 12, wantSize: 12, wantTailPad: 8},
		{name: "all-zero-size", count: 0, onlyZero: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			for _, arch := range []Arch{Arch386, ArchAMD64, ArchARM64} {
				t.Run(arch.String(), func(t *testing.T) {
					c, err := NewContext(generatorTestMetadata(t, generatorTestOptions{tailCount: test.count}))
					if err != nil {
						t.Fatal(err)
					}
					def, err := c.resolveTypeDef(1)
					if err != nil {
						t.Fatal(err)
					}
					if test.onlyZero {
						def.def.FieldList.Start++
					}
					c.classLayout[1] = winmd.ClassLayout{ClassSize: test.classSize}
					c.abiLayoutTypeDefs[1] = true
					plan, err := c.planStructABI(def, arch, nil)
					if test.wantErr {
						if err == nil || !strings.Contains(err.Error(), "cannot be represented") {
							t.Fatalf("unrepresentable layout error = %v", err)
						}
						var output strings.Builder
						if err := c.writeTypeDef(&output, def, arch); err == nil {
							t.Fatal("generated an incompatible Windows ABI type")
						}
						return
					}
					if err != nil {
						t.Fatal(err)
					}
					if plan.typeLayout.abiSize != test.wantSize || plan.typeLayout.goSize != test.wantSize || plan.tailPadding != test.wantTailPad {
						t.Fatalf("layout = %+v, tail padding %d; want size %d, tail padding %d", plan.typeLayout, plan.tailPadding, test.wantSize, test.wantTailPad)
					}
					var output strings.Builder
					if err := c.writeTypeDef(&output, def, arch); err != nil {
						t.Fatal(err)
					}
					fset := token.NewFileSet()
					file, err := parser.ParseFile(fset, "generated.go", "package test\n"+output.String(), 0)
					if err != nil {
						t.Fatal(err)
					}
					sizes := types.SizesFor("gc", arch.String())
					pkg, err := (&types.Config{Sizes: sizes}).Check("test", fset, []*ast.File{file}, nil)
					if err != nil {
						t.Fatal(err)
					}
					actual := sizes.Sizeof(pkg.Scope().Lookup(def.GoName).Type())
					if actual != int64(test.wantSize) {
						t.Fatalf("generated Go size %d differs from Windows size %d:\n%s", actual, test.wantSize, output.String())
					}
				})
			}
		})
	}
}

func TestTypeDefCacheCanonicalKeys(t *testing.T) {
	t.Parallel()
	c, err := NewContext(generatorTestMetadata(t, generatorTestOptions{
		duplicateNames: true,
		extraTypeDefs:  []qualifiedTypeName{{Namespace: "Test", Name: "Sample"}},
	}))
	if err != nil {
		t.Fatal(err)
	}
	first, err := c.Metadata.Tables.TypeDef.At(1)
	if err != nil {
		t.Fatal(err)
	}
	second, err := c.Metadata.Tables.TypeDef.At(2)
	if err != nil {
		t.Fatal(err)
	}
	key := typeDefKey(first)
	if typeDefKey(second) == key || c.typeDefCache.canonicalKey(typeDefKey(second)) != key {
		t.Fatal("equal names at different heap offsets did not canonicalize")
	}
	if len(c.typeDefCache.unresolvedDuplicated[key]) != 2 {
		t.Fatal("equal names at different offsets were not grouped as duplicate definitions")
	}
	for i, arch := range []Arch{Arch386, ArchAMD64} {
		c.typeDefSupportedArch[winmd.Index(i+1)] = arch
	}
	if _, err := c.resolveTypeRef(1, ArchARM64); !errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
		t.Fatalf("cold unavailable architecture error = %v", err)
	}
	for range 2 {
		for i, arch := range []Arch{Arch386, ArchAMD64} {
			got, err := c.resolveTypeRef(1, arch)
			if err != nil || got == nil || got.Index != winmd.Index(i+1) {
				t.Fatalf("canonical architecture lookup returned %+v, %v", got, err)
			}
		}
	}
	if _, err := c.resolveTypeRef(1, ArchARM64); !errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
		t.Fatalf("unavailable architecture error = %v", err)
	}
	if len(c.typeDefCache.resolvedDuplicated[key]) != 2 {
		t.Fatal("resolved architecture variants were not cached under the canonical key")
	}
	ref, err := c.Metadata.Tables.TypeRef.At(1)
	if err != nil {
		t.Fatal(err)
	}
	if ref.Name.Start == first.Name.Start || ref.Namespace.Start == first.Namespace.Start {
		t.Fatal("canonicalization changed the original metadata heap offsets")
	}
	if got := c.typeDefCache.get(second.Namespace, second.Name, ArchAMD64); got == nil || got.Index != 2 {
		t.Fatalf("lookup using duplicate definition offsets = %+v; want TypeDef 2", got)
	}
}

func TestTypeNameCanonicalizationBoundaries(t *testing.T) {
	t.Parallel()
	names := []qualifiedTypeName{
		{Namespace: "Test", Name: "Sub.Mode"},
		{Namespace: "Test.Sub", Name: "Mode"},
		{Namespace: "test", Name: "Sub.Mode"},
		{Namespace: "Test", Name: "Sub.mode"},
	}
	c, err := NewContext(generatorTestMetadata(t, generatorTestOptions{extraTypeDefs: names}))
	if err != nil {
		t.Fatal(err)
	}
	if len(c.typeDefCache.aliases) != 0 {
		t.Fatal("distinct case-sensitive namespace/name pairs were canonicalized together")
	}
	for i, name := range names {
		index := winmd.Index(i + 2)
		if indices := c.typeDefsByName[name]; len(indices) != 1 || indices[0] != index {
			t.Fatalf("top-level lookup for %+v = %v; want [%d]", name, indices, index)
		}
		def, err := c.resolveTypeDef(index)
		if err != nil {
			t.Fatal(err)
		}
		if got := c.typeDefCache.get(def.Namespace, def.Name, ArchAll); got != def {
			t.Fatalf("canonical lookup for %+v returned %+v; want %+v", name, got, def)
		}
	}
}
