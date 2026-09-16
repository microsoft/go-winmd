// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package gowinmd

import (
	"bytes"
	"debug/pe"
	"encoding/binary"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/microsoft/go-winmd/winmd"
)

// generationMetadataBuilder holds rows in table order, allowing focused tests
// to construct real metadata without modifying the shared Win32 fixture.
type generationMetadataBuilder struct {
	t     *testing.T
	rows  [64][][]byte
	strs  []byte
	blobs []byte
}

func newGenerationMetadataBuilder(t *testing.T) *generationMetadataBuilder {
	b := &generationMetadataBuilder{t: t, strs: []byte{0}, blobs: []byte{0}}
	b.add(0, uint16(0), b.str("Test.winmd"), uint16(0), uint16(0), uint16(0))
	return b
}

func (b *generationMetadataBuilder) str(value string) uint16 {
	b.t.Helper()
	if len(b.strs)+len(value)+1 > 1<<16 {
		b.t.Fatal("test string heap exceeds two-byte indexes")
	}
	offset := uint16(len(b.strs))
	b.strs = append(b.strs, value...)
	b.strs = append(b.strs, 0)
	return offset
}

func (b *generationMetadataBuilder) blob(data []byte) uint16 {
	b.t.Helper()
	if len(data) >= 128 || len(b.blobs)+len(data)+1 > 1<<16 {
		b.t.Fatal("test blob is too large")
	}
	offset := uint16(len(b.blobs))
	b.blobs = append(b.blobs, byte(len(data)))
	b.blobs = append(b.blobs, data...)
	return offset
}

func (b *generationMetadataBuilder) add(table int, fields ...any) {
	b.t.Helper()
	var row bytes.Buffer
	for _, field := range fields {
		if err := binary.Write(&row, binary.LittleEndian, field); err != nil {
			b.t.Fatal(err)
		}
	}
	b.rows[table] = append(b.rows[table], row.Bytes())
}

func (b *generationMetadataBuilder) metadata() *winmd.Metadata {
	b.t.Helper()
	var valid uint64
	for table, rows := range b.rows {
		if len(rows) != 0 {
			valid |= 1 << table
		}
	}
	tables := make([]byte, 24)
	tables[4], tables[7] = 2, 1
	binary.LittleEndian.PutUint64(tables[8:], valid)
	for _, rows := range b.rows {
		if len(rows) != 0 {
			tables = binary.LittleEndian.AppendUint32(tables, uint32(len(rows)))
		}
	}
	for _, rows := range b.rows {
		for _, row := range rows {
			tables = append(tables, row...)
		}
	}
	streams := []struct {
		name string
		data []byte
	}{{"#~", tables}, {"#Strings", b.strs}, {"#Blob", b.blobs}}
	root := make([]byte, 16)
	binary.LittleEndian.PutUint32(root, 0x424a5342)
	binary.LittleEndian.PutUint16(root[4:], 1)
	binary.LittleEndian.PutUint16(root[6:], 1)
	binary.LittleEndian.PutUint32(root[12:], 4)
	root = append(root, 'v', '1', 0, 0, 0, 0)
	root = binary.LittleEndian.AppendUint16(root, uint16(len(streams)))
	offset := len(root)
	for _, stream := range streams {
		offset += 8 + (len(stream.name)+4)&^3
	}
	for _, stream := range streams {
		root = binary.LittleEndian.AppendUint32(root, uint32(offset))
		root = binary.LittleEndian.AppendUint32(root, uint32(len(stream.data)))
		root = append(root, stream.name...)
		root = append(root, make([]byte, 4-len(stream.name)%4)...)
		offset += (len(stream.data) + 3) &^ 3
	}
	for _, stream := range streams {
		root = append(root, stream.data...)
		root = append(root, make([]byte, (4-len(stream.data)%4)%4)...)
	}
	section := make([]byte, 72)
	binary.LittleEndian.PutUint32(section, 72)
	binary.LittleEndian.PutUint16(section[4:], 2)
	binary.LittleEndian.PutUint32(section[8:], 0x2000+72)
	binary.LittleEndian.PutUint32(section[12:], uint32(len(root)))
	section = append(section, root...)
	optional := pe.OptionalHeader32{Magic: 0x10b, NumberOfRvaAndSizes: 16}
	optional.DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR] = pe.DataDirectory{VirtualAddress: 0x2000, Size: 72}
	var image bytes.Buffer
	for _, header := range []any{
		pe.FileHeader{Machine: pe.IMAGE_FILE_MACHINE_I386, NumberOfSections: 1, SizeOfOptionalHeader: uint16(binary.Size(optional))},
		optional,
		pe.SectionHeader32{Name: [8]byte{'.', 't', 'e', 'x', 't'}, VirtualAddress: 0x2000, VirtualSize: uint32(len(section)), SizeOfRawData: uint32(len(section)), PointerToRawData: 0x200},
	} {
		if err := binary.Write(&image, binary.LittleEndian, header); err != nil {
			b.t.Fatal(err)
		}
	}
	image.Write(make([]byte, 0x200-image.Len()))
	image.Write(section)
	f, err := pe.NewFile(bytes.NewReader(image.Bytes()))
	if err != nil {
		b.t.Fatal(err)
	}
	defer f.Close()
	m, err := winmd.New(f)
	if err != nil {
		b.t.Fatal(err)
	}
	return m
}

func methodGenerationMetadata(t *testing.T, signature []byte, module, entry string, params ...string) *generationMetadataBuilder {
	t.Helper()
	b := newGenerationMetadataBuilder(t)
	b.add(2, uint32(0), b.str("<Module>"), uint16(0), uint16(0), uint16(1), uint16(1))
	b.add(6, uint32(0), uint16(0), uint16(winmd.MethodAttributes_Public|winmd.MethodAttributes_Static|winmd.MethodAttributes_PInvokeImpl), b.str("Invoke"), b.blob(signature), uint16(1))
	for i, name := range params {
		b.add(8, uint16(0), uint16(i+1), b.str(name))
	}
	b.add(26, b.str(module))
	b.add(28, uint16(winmd.PInvokeAttributes_CallConvPlatformapi), uint16(3), b.str(entry), uint16(1))
	return b
}

func (b *generationMetadataBuilder) sizeAttribute(paramRow uint16, inBytes bool, countIndex int16) {
	name, field := "NativeArrayInfoAttribute", "CountParamIndex"
	if inBytes {
		name, field = "MemorySizeAttribute", "BytesParamIndex"
	}
	// A module-scoped attribute TypeRef with a parameterless constructor.
	refRow := uint16(len(b.rows[1]) + 1)
	b.add(1, uint16(4), b.str(name), b.str("Windows.Win32.Foundation.Metadata"))
	memberRow := uint16(len(b.rows[10]) + 1)
	b.add(10, refRow<<3|1, b.str(".ctor"), b.blob([]byte{0x20, 0, 1}))
	value := append([]byte{1, 0, 1, 0, 0x53, 6, byte(len(field))}, field...)
	value = binary.LittleEndian.AppendUint16(value, uint16(countIndex))
	b.add(12, paramRow<<5|4, memberRow<<3|3, b.blob(value))
}

func writeTestMethod(t *testing.T, b *generationMetadataBuilder, options MethodOptions) (string, error) {
	t.Helper()
	m := b.metadata()
	c, err := NewContext(m)
	if err != nil {
		t.Fatal(err)
	}
	method, err := m.Tables.MethodDef.At(0)
	if err != nil {
		t.Fatal(err)
	}
	var output strings.Builder
	err = c.WriteMethodWithOptions(&output, 0, method, ArchARM64, options)
	return output.String(), err
}

func TestWriteMethodNativeEntryPoint(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name, module, entry, goName, want, wantError string
	}{
		{"same-name", "native.dll", "Invoke", "", "//sys\tInvoke() = native.Invoke", ""},
		{"alias", "native.dll", "ActualEntry", "", "//sys\tInvoke() = native.ActualEntry", ""},
		{"custom-go-name", "native.dll", "ActualEntry", "Call", "//sys\tCall() = native.ActualEntry", ""},
		{"default-module", "kernel32.dll", "Invoke", "", "//sys\tInvoke()", ""},
		{"default-module-alias", "kernel32.dll", "ActualEntry", "", "//sys\tInvoke() = ActualEntry", ""},
		{"go-name-matches-entry", "kernel32.dll", "ActualEntry", "ActualEntry", "//sys\tActualEntry()", ""},
		{"ordinal", "native.dll", "#660", "", "", "ordinal import"},
		{"missing-entry", "native.dll", "", "", "", "missing native entry point"},
	} {
		t.Run(test.name, func(t *testing.T) {
			b := methodGenerationMetadata(t, []byte{0, 0, 1}, test.module, test.entry)
			got, err := writeTestMethod(t, b, MethodOptions{GoName: test.goName})
			if test.wantError != "" {
				if err == nil || !strings.Contains(err.Error(), test.wantError) {
					t.Fatalf("error = %v; want %q", err, test.wantError)
				}
				if got != "" {
					t.Fatalf("invalid entry point produced partial output: %q", got)
				}
				if test.name == "ordinal" && !errors.Is(err, ErrOrdinalImport) {
					t.Fatalf("ordinal error = %v; want ErrOrdinalImport", err)
				}
			} else if err != nil || got != test.want {
				t.Fatalf("method = %q, %v; want %q", got, err, test.want)
			}
		})
	}
}

func TestWriteMethodByRef(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name   string
		sig    []byte
		params []string
		want   string
	}{
		{"pointer-control", []byte{0, 1, 1, 0x0f, 8}, []string{"value"}, "value *int32"},
		{"parameter", []byte{0, 1, 1, 0x10, 8}, []string{"value"}, "value *int32"},
		{"return", []byte{0, 0, 0x10, 8}, nil, "r *int32"},
		{"pointer-parameter", []byte{0, 1, 1, 0x10, 0x0f, 8}, []string{"value"}, "value **int32"},
		{"void-pointer-parameter", []byte{0, 1, 1, 0x10, 0x0f, 1}, []string{"value"}, "value *unsafe.Pointer"},
	} {
		t.Run(test.name, func(t *testing.T) {
			for _, projection := range []Projection{ProjectionRaw, ProjectionIdiomatic} {
				b := methodGenerationMetadata(t, test.sig, "native.dll", "Invoke", test.params...)
				got, err := writeTestMethod(t, b, MethodOptions{Projection: projection})
				if err != nil || !strings.Contains(got, test.want) {
					t.Fatalf("projection %d: method = %q, %v; want %q", projection, got, err, test.want)
				}
			}
		})
	}
}

func TestWriteTypeRejectsManagedSignatures(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name string
		typ  []byte
	}{
		{"szarray", []byte{0x1d, 8}},
		{"genericinst", []byte{0x15, 0x12, 5, 1, 8}},
		{"var", []byte{0x13, 0}},
		{"mvar", []byte{0x1e, 0}},
	} {
		for _, wrapper := range []struct {
			name   string
			prefix []byte
			suffix []byte
		}{
			{"direct", nil, nil},
			{"pointer", []byte{0x0f}, nil},
			{"byref", []byte{0x10}, nil},
			{"array", []byte{0x14}, []byte{1, 1, 2, 0}},
		} {
			for _, arch := range []Arch{Arch386, ArchAMD64, ArchARM64} {
				t.Run(test.name+"/"+wrapper.name+"/"+arch.String(), func(t *testing.T) {
					var m winmd.Metadata
					// A generic method with one parameter; no metadata is needed
					// to decode the type or reject an unsupported projection.
					data := append([]byte{0x10, 1, 1, 1}, wrapper.prefix...)
					data = append(data, test.typ...)
					data = append(data, wrapper.suffix...)
					sig, err := m.MethodDefSignature(data)
					if err != nil {
						t.Fatal(err)
					}
					var c Context
					var output strings.Builder
					err = c.writeType(&output, &sig.Param[0].Type, arch)
					want := "unsupported type for Go generation: " + winmd.ElementType(test.typ[0]).String()
					if err == nil || !strings.Contains(err.Error(), want) {
						t.Fatalf("rendered managed signature as %q, %v; want %q", output.String(), err, want)
					}
					if wrapper.name == "direct" {
						if output.Len() != 0 {
							t.Fatalf("unsupported type emitted output: %q", output.String())
						}
						if _, err := c.sigTypeABITypeLayout(&sig.Param[0].Type, arch, nil); err == nil {
							t.Fatal("managed type accepted as a native ABI field")
						}
					}
				})
			}
		}
	}
}

func TestWriteMethodCountUnits(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name    string
		typ     []byte
		inBytes bool
		want    string
	}{
		{"byte-size", []byte{0x0f, 5}, true, "buffer []byte"},
		{"signed-byte-size", []byte{0x0f, 4}, true, "buffer []int8"},
		{"boolean-size", []byte{0x0f, 2}, true, "buffer []bool"},
		{"uint16-byte-size", []byte{0x0f, 7}, true, "buffer *uint16, count uint32"},
		{"uint32-byte-size", []byte{0x0f, 9}, true, "buffer *uint32, count uint32"},
		{"pointer-byte-size", []byte{0x0f, 0x0f, 5}, true, "buffer **uint8, count uint32"},
		{"void-byte-size", []byte{0x0f, 1}, true, "buffer unsafe.Pointer, count uint32"},
		{"byref-byte-size", []byte{0x10, 9}, true, "buffer *uint32, count uint32"},
		{"byte-count", []byte{0x0f, 5}, false, "buffer []byte"},
		{"uint16-count", []byte{0x0f, 7}, false, "buffer []uint16"},
		{"pointer-count", []byte{0x0f, 0x0f, 5}, false, "buffer []*uint8"},
	} {
		t.Run(test.name, func(t *testing.T) {
			sig := append([]byte{0, 2, 1}, test.typ...)
			sig = append(sig, 9)
			b := methodGenerationMetadata(t, sig, "native.dll", "Invoke", "buffer", "count")
			b.sizeAttribute(1, test.inBytes, 1)
			got, err := writeTestMethod(t, b, MethodOptions{Projection: ProjectionIdiomatic})
			want := "//sys\tInvoke(" + test.want + ") = native.Invoke"
			if err != nil || got != want {
				t.Fatalf("method = %q, %v; want %q", got, err, want)
			}
			got, err = writeTestMethod(t, b, MethodOptions{Projection: ProjectionRaw})
			if err != nil || !strings.Contains(got, "count uint32") || strings.Contains(got, "[]") {
				t.Fatalf("raw projection changed parameter shapes: %q, %v", got, err)
			}
		})
	}
	t.Run("mixed-buffers", func(t *testing.T) {
		b := methodGenerationMetadata(t, []byte{0, 4, 1, 0x0f, 5, 9, 0x0f, 7, 9}, "native.dll", "Invoke", "bytes", "byteCount", "words", "wordBytes")
		b.sizeAttribute(1, true, 1)
		b.sizeAttribute(3, true, 3)
		got, err := writeTestMethod(t, b, MethodOptions{Projection: ProjectionIdiomatic})
		want := "//sys\tInvoke(bytes []byte, words *uint16, wordBytes uint32) = native.Invoke"
		if err != nil || got != want {
			t.Fatalf("method = %q, %v; want %q", got, err, want)
		}
	})
}

func TestWriteMethodRealMetadataCorrections(t *testing.T) {
	m, err := winmd.Open("../../../../winmd/testdata/Windows.Win32.winmd")
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewContext(m)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		want      string
		wantError string
	}{
		"RtlGenRandom":         {want: "= advapi32.SystemFunction036"},
		"RtlEncryptMemory":     {want: "= advapi32.SystemFunction040"},
		"RtlDecryptMemory":     {want: "= advapi32.SystemFunction041"},
		"waveOutPrepareHeader": {want: "pwh *WAVEHDR, cbwh uint32"},
		"BCryptGenRandom":      {want: "pbBuffer []byte"},
		"FileIconInit":         {wantError: "ordinal import"},
	}
	found := make(map[string]bool)
	for index := range m.Tables.MethodDef.Indices() {
		method, err := m.Tables.MethodDef.At(index)
		if err != nil {
			t.Fatal(err)
		}
		name := method.Name.String()
		test, ok := tests[name]
		if !ok {
			continue
		}
		found[name] = true
		t.Run(name, func(t *testing.T) {
			var output strings.Builder
			err := c.WriteMethodWithOptions(&output, index, method, ArchARM64, MethodOptions{Projection: ProjectionIdiomatic})
			if test.wantError != "" {
				if err == nil || !strings.Contains(err.Error(), test.wantError) {
					t.Fatalf("error = %v; want %q", err, test.wantError)
				}
			} else if err != nil || !strings.Contains(output.String(), test.want) {
				t.Fatalf("method = %q, %v; want %q", output.String(), err, test.want)
			}
		})
	}
	if len(found) != len(tests) {
		t.Fatalf("fixture contains %v; want all %d methods", found, len(tests))
	}
}

func generatedStruct(t *testing.T, text, name string, arch Arch) (*types.Struct, types.Sizes) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "generated.go", "package test\n"+text, 0)
	if err != nil {
		t.Fatal(err)
	}
	sizes := types.SizesFor("gc", arch.String())
	pkg, err := (&types.Config{Sizes: sizes}).Check("test", fset, []*ast.File{file}, nil)
	if err != nil {
		t.Fatal(err)
	}
	return pkg.Scope().Lookup(name).Type().Underlying().(*types.Struct), sizes
}

func TestNestedStructABIBoundaries(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name      string
		last      []byte
		classSize uint32
		explicit  bool
		childSize int64
	}{
		{"full-child", []byte{9}, 0, false, 8},
		{"implicit-tail-padding", []byte{5}, 0, false, 8},
		{"explicit-tail-padding", []byte{5}, 16, false, 16},
		{"zero-sized-last-field", []byte{0x14, 5, 1, 1, 0, 0}, 8, false, 8},
		{"explicit-field-offset", []byte{5}, 0, true, 12},
	} {
		t.Run(test.name, func(t *testing.T) {
			b := newGenerationMetadataBuilder(t)
			b.add(35, uint16(4), uint16(0), uint16(0), uint16(0), uint32(0), uint16(0), b.str("System.Runtime"), uint16(0), uint16(0))
			b.add(1, uint16(6), b.str("ValueType"), b.str("System"))
			b.add(1, uint16(4), b.str("Parent"), b.str("Test"))
			b.add(1, uint16(11), b.str("Child"), uint16(0))
			b.add(2, uint32(0), b.str("<Module>"), uint16(0), uint16(0), uint16(1), uint16(1))
			b.add(2, uint32(winmd.TypeAttributes_Public|winmd.TypeAttributes_SequentialLayout|winmd.TypeAttributes_Sealed), b.str("Parent"), b.str("Test"), uint16(5), uint16(1), uint16(1))
			layoutFlag := winmd.TypeAttributes_SequentialLayout
			if test.explicit {
				layoutFlag = winmd.TypeAttributes_ExplicitLayout
				b.add(16, uint32(0), uint16(4))
				b.add(16, uint32(8), uint16(5))
			}
			b.add(2, uint32(winmd.TypeAttributes_NestedPublic|layoutFlag|winmd.TypeAttributes_Sealed), b.str("Child"), uint16(0), uint16(5), uint16(4), uint16(1))
			if test.classSize != 0 {
				b.add(15, uint16(0), test.classSize, uint16(3))
			}
			b.add(4, uint16(winmd.FieldAttributes_Public), b.str("Before"), b.blob([]byte{6, 5}))
			b.add(4, uint16(winmd.FieldAttributes_Public), b.str("Data"), b.blob([]byte{6, 0x11, 13}))
			// Reusing a member name across struct boundaries must remain valid Go.
			b.add(4, uint16(winmd.FieldAttributes_Public), b.str("A"), b.blob([]byte{6, 5}))
			b.add(4, uint16(winmd.FieldAttributes_Public), b.str("A"), b.blob([]byte{6, 9}))
			b.add(4, uint16(winmd.FieldAttributes_Public), b.str("B"), b.blob(append([]byte{6}, test.last...)))
			b.add(41, uint16(3), uint16(2))
			c, err := NewContext(b.metadata())
			if err != nil {
				t.Fatal(err)
			}
			if err := c.SelectTypeDef("Test", "Parent", ""); err != nil {
				t.Fatal(err)
			}
			def := c.resolvedDefsByIndex[1]
			for _, arch := range []Arch{Arch386, ArchAMD64, ArchARM64} {
				t.Run(arch.String(), func(t *testing.T) {
					plan, err := c.planStructABI(def, arch, nil)
					if err != nil {
						t.Fatal(err)
					}
					var output strings.Builder
					if err := c.writeTypeDef(&output, def, arch); err != nil {
						t.Fatal(err)
					}
					actual, sizes := generatedStruct(t, output.String(), "Parent", arch)
					wantSize := test.childSize + 8 // Prefix alignment and final byte, rounded to 4.
					if got := sizes.Sizeof(actual); got != wantSize || got != int64(plan.typeLayout.goSize) {
						t.Fatalf("actual size %d, plan %d; want %d:\n%s", got, plan.typeLayout.goSize, wantSize, output.String())
					}
					if actual.NumFields() != 3 || actual.Field(1).Name() != "Data" {
						t.Fatalf("nested field boundary not preserved:\n%s", output.String())
					}
					fields := []*types.Var{actual.Field(0), actual.Field(1), actual.Field(2)}
					offsets := sizes.Offsetsof(fields)
					if offsets[1] != 4 || offsets[2] != 4+test.childSize {
						t.Fatalf("nested field offsets = %v; want [0 4 %d]", offsets, 4+test.childSize)
					}
					if got := sizes.Sizeof(actual.Field(1).Type()); got != test.childSize {
						t.Fatalf("child size %d; want %d", got, test.childSize)
					}
				})
			}
		})
	}
}

func TestUnicodeTypeAndFieldNames(t *testing.T) {
	t.Parallel()
	for _, test := range []struct{ name, want string }{
		{"sample", "Sample"}, {"éclair", "Éclair"}, {"Δelta", "Δelta"}, {"𐐨name", "𐐀name"},
	} {
		t.Run(test.name, func(t *testing.T) {
			b := newGenerationMetadataBuilder(t)
			b.add(2, uint32(0), b.str("<Module>"), uint16(0), uint16(0), uint16(1), uint16(1))
			b.add(2, uint32(winmd.TypeAttributes_Public|winmd.TypeAttributes_SequentialLayout), b.str(test.name), b.str("Test"), uint16(0), uint16(1), uint16(1))
			b.add(4, uint16(winmd.FieldAttributes_Public), b.str(test.name), b.blob([]byte{6, 9}))
			c, err := NewContext(b.metadata())
			if err != nil {
				t.Fatal(err)
			}
			def, err := c.resolveTypeDef(1)
			if err != nil {
				t.Fatal(err)
			}
			var output strings.Builder
			if err := c.writeTypeDef(&output, def, ArchARM64); err != nil {
				t.Fatal(err)
			}
			if def.GoName != test.want || !utf8.ValidString(output.String()) {
				t.Fatalf("name %q generated %q; want valid UTF-8 name %q", test.name, output.String(), test.want)
			}
			actual, _ := generatedStruct(t, output.String(), test.want, ArchARM64)
			if actual.Field(0).Name() != test.want {
				t.Fatalf("field name %q; want %q", actual.Field(0).Name(), test.want)
			}
		})
	}
}

func TestArrayABIOverflow(t *testing.T) {
	t.Parallel()
	compressed := func(value uint32) []byte {
		if value < 128 {
			return []byte{byte(value)}
		}
		if value < 16384 {
			return []byte{byte(value>>8) | 0x80, byte(value)}
		}
		return []byte{byte(value>>24) | 0xc0, byte(value >> 16), byte(value >> 8), byte(value)}
	}
	for _, test := range []struct {
		name    string
		kind    byte
		dims    []uint32
		want    uint32
		wantErr bool
	}{
		{"small", 5, []uint32{2, 3, 4}, 24, false},
		{"large-valid", 5, []uint32{0x1fffffff, 8}, 0xfffffff8, false},
		{"uint32-overflow", 5, []uint32{1 << 22, 1 << 22}, 0, true},
		{"uint64-overflow", 5, []uint32{1 << 22, 1 << 22, 1 << 20}, 0, true},
		{"element-size-overflow", 0x0b, []uint32{1 << 28, 4}, 0, true},
		{"empty-inner", 5, []uint32{1 << 22, 1 << 22, 1 << 20, 0}, 0, false},
		{"empty-outer", 5, []uint32{0, 3, 4}, 0, false},
		{"oversized-inner-of-empty-outer", 5, []uint32{0, 1 << 22, 1 << 22}, 0, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			data := []byte{6, 0x14, test.kind, byte(len(test.dims)), byte(len(test.dims))}
			for _, dimension := range test.dims {
				data = append(data, compressed(dimension)...)
			}
			data = append(data, 0)
			var m winmd.Metadata
			sig, err := m.FieldSignature(data)
			if err != nil {
				t.Fatal(err)
			}
			for _, arch := range []Arch{Arch386, ArchAMD64, ArchARM64} {
				var c Context
				layout, err := c.sigTypeABITypeLayout(&sig.Type, arch, nil)
				if test.wantErr {
					if err == nil || !strings.Contains(err.Error(), "overflows") {
						t.Fatalf("%s: overflowing dimensions %v returned %+v, %v", arch, test.dims, layout, err)
					}
				} else if err != nil || layout.abiSize != test.want || layout.goSize != test.want {
					t.Fatalf("%s: dimensions %v returned %+v, %v; want size %d", arch, test.dims, layout, err, test.want)
				}
			}
		})
	}
}
