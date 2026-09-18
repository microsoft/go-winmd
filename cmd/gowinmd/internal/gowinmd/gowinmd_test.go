// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package gowinmd

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestInt16AttributeField(t *testing.T) {
	value := winmd.CustomAttributeValue{NamedArguments: []winmd.CustomAttributeNamedArgument{
		{Name: "Other", CustomAttributeArgument: winmd.CustomAttributeArgument{Value: int16(3)}},
		{Name: "CountParamIndex", CustomAttributeArgument: winmd.CustomAttributeArgument{Value: int16(0)}},
	}}
	if got, ok, err := int16AttributeField(value, "CountParamIndex"); got != 0 || !ok || err != nil {
		t.Fatalf("zero index = (%v, %v, %v); want (0, true, nil)", got, ok, err)
	}
	if got, ok, err := int16AttributeField(value, "BytesParamIndex"); got != 0 || ok || err != nil {
		t.Fatalf("missing field = (%v, %v, %v); want (0, false, nil)", got, ok, err)
	}
}

func TestSupportedArchitecture(t *testing.T) {
	for _, test := range []struct {
		name      string
		arguments []winmd.CustomAttributeArgument
		want      Arch
		wantErr   bool
	}{
		{"int32", []winmd.CustomAttributeArgument{{Value: int32(3)}}, Arch386 | ArchAMD64, false},
		{"uint32", []winmd.CustomAttributeArgument{{Value: uint32(3)}}, ArchNone, true},
		{"missing", nil, ArchNone, true},
		{"multiple", []winmd.CustomAttributeArgument{{Value: int32(1)}, {Value: int32(2)}}, ArchNone, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			value := winmd.CustomAttributeValue{FixedArguments: test.arguments}
			got, err := supportedArchitecture(value)
			if got != test.want || (err != nil) != test.wantErr {
				t.Fatalf("supportedArchitecture() = (%v, %v); want (%v, error %v)", got, err, test.want, test.wantErr)
			}
		})
	}
}

func TestMkwinsyscallModuleName(t *testing.T) {
	for _, test := range []struct {
		name string
		want string
	}{
		{"bcrypt.dll", "bcrypt"},
		{"BCRYPT.DLL", "bcrypt"},
		{"bcrypt", "bcrypt"},
	} {
		if got := mkwinsyscallModuleName(test.name); got != test.want {
			t.Errorf("mkwinsyscallModuleName(%q) = %q; want %q", test.name, got, test.want)
		}
	}
}

func TestFormatEnumConstant(t *testing.T) {
	for _, test := range []struct {
		name string
		typ  winmd.ElementType
		data []byte
		want string
	}{
		{"false", winmd.ElementType_BOOLEAN, []byte{0}, "false"},
		{"true", winmd.ElementType_BOOLEAN, []byte{1}, "true"},
		{"char", winmd.ElementType_CHAR, []byte{0xAC, 0x20}, "0x20ac"},
		{"char-surrogate", winmd.ElementType_CHAR, []byte{0, 0xD8}, "0xd800"},
		{"zero", winmd.ElementType_I1, []byte{0}, "0x0"},
		{"negative-one", winmd.ElementType_I1, []byte{0xFF}, "-0x1"},
		{"i1", winmd.ElementType_I1, []byte{0x80}, "-0x80"},
		{"i2", winmd.ElementType_I2, []byte{0, 0x80}, "-0x8000"},
		{"positive-i2", winmd.ElementType_I2, []byte{0x34, 0x12}, "0x1234"},
		{"i4", winmd.ElementType_I4, []byte{0, 0, 0, 0x80}, "-0x80000000"},
		{"i8", winmd.ElementType_I8, []byte{0, 0, 0, 0, 0, 0, 0, 0x80}, "-0x8000000000000000"},
		{"u1", winmd.ElementType_U1, []byte{0xFF}, "0xff"},
		{"u2", winmd.ElementType_U2, []byte{0xFF, 0xFF}, "0xffff"},
		{"u4", winmd.ElementType_U4, []byte{0xFF, 0xFF, 0xFF, 0xFF}, "0xffffffff"},
		{"u8", winmd.ElementType_U8, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, "0xffffffffffffffff"},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := formatEnumConstant(winmd.Constant{Type: test.typ, Value: test.data})
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Errorf("formatEnumConstant() = %q; want %q", got, test.want)
			}
		})
	}
}

func TestWriteEnumWithNonstandardBackingField(t *testing.T) {
	for _, test := range []struct {
		name       string
		kind       winmd.ElementType
		value      []byte
		underlying string
		literal    string
	}{
		{"integer", winmd.ElementType_I4, []byte{3, 0, 0, 0}, "int32", "0x3"},
		{"boolean", winmd.ElementType_BOOLEAN, []byte{1}, "bool", "true"},
		{"character", winmd.ElementType_CHAR, []byte{0xAC, 0x20}, "uint16", "0x20ac"},
	} {
		t.Run(test.name, func(t *testing.T) {
			metadata, err := winmd.Open("../../../../winmd/testdata/Windows.Win32.winmd")
			if err != nil {
				t.Fatal(err)
			}
			context, err := NewContext(metadata)
			if err != nil {
				t.Fatal(err)
			}
			index, ok := context.typeDefsByName[qualifiedTypeName{Namespace: "Windows.Win32.Foundation.Metadata", Name: "Architecture"}]
			if !ok || len(context.typeDefNameDuplicates[index]) != 0 {
				t.Fatal("expected one Architecture definition")
			}
			def, err := context.resolveTypeDef(index)
			if err != nil {
				t.Fatal(err)
			}
			def.GoName = "Mode"
			for index := range def.def.FieldList.All() {
				field, err := metadata.Tables.Field.At(index)
				if err != nil {
					t.Fatal(err)
				}
				if !field.Flags.HasAll(winmd.FieldFlags_Static) {
					if field.Name.String() != "value__" || len(field.Signature) != 2 {
						t.Fatal("unexpected fixture backing field")
					}
					// Mutate only this test's in-memory metadata, retaining heap lengths.
					copy(metadata.Strings[field.Name.Start:], "MyValue")
					field.Signature[1] = byte(test.kind)
					continue
				}
				context.fieldConstant[index] = winmd.Constant{Type: test.kind, Value: test.value}
			}
			var output strings.Builder
			if err := context.writeTypeDef(&output, def, ArchAll); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(output.String(), "type Mode "+test.underlying) || !strings.Contains(output.String(), " = "+test.literal) {
				t.Fatalf("unexpected enum declaration:\n%s", output.String())
			}
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, "", "package test\n"+output.String(), 0)
			if err != nil {
				t.Fatal(err)
			}
			var config types.Config
			if _, err := config.Check("test", fset, []*ast.File{file}, nil); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestFormatEnumConstantErrors(t *testing.T) {
	for _, test := range []struct {
		name    string
		typ     winmd.ElementType
		data    []byte
		wantEOF bool
	}{
		{"truncated-boolean", winmd.ElementType_BOOLEAN, nil, true},
		{"truncated-char", winmd.ElementType_CHAR, []byte{'A'}, true},
		{"float", winmd.ElementType_R4, []byte{0, 0, 0, 0}, false},
		{"string", winmd.ElementType_STRING, []byte{'A', 0}, false},
		{"null", winmd.ElementType_CLASS, []byte{0, 0, 0, 0}, false},
		{"truncated", winmd.ElementType_I4, []byte{0}, true},
		{"extra-byte", winmd.ElementType_U2, []byte{0, 0, 0}, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := formatEnumConstant(winmd.Constant{Type: test.typ, Value: test.data})
			if got != "" || err == nil || errors.Is(err, io.ErrUnexpectedEOF) != test.wantEOF {
				t.Errorf("formatEnumConstant() = (%q, %v); want (empty, error), unexpected EOF %v", got, err, test.wantEOF)
			}
		})
	}
}

func TestEscapedUpper(t *testing.T) {
	for _, test := range []struct {
		name string
		want string
	}{
		{"", ""},
		{"q", "Q"},
		{"type", "Type"},
	} {
		if got := escapedUpper(test.name); got != test.want {
			t.Errorf("escapedUpper(%q) = %q; want %q", test.name, got, test.want)
		}
	}
}

func TestContext_writeType_cycle(t *testing.T) {
	t.Skip("cycles can't be built with SigType rather than *SigType, and this code only supports SigType")

	p1 := winmd.SigType{Kind: winmd.ElementType_PTR}
	p2 := winmd.SigType{Kind: winmd.ElementType_PTR}
	// These are copies, not pointers...
	var p1a any = p1
	var p2a any = p2
	// ...So this doesn't cause a cycle.
	p2.Value = p1a
	p1.Value = p2a

	// If we can create a cycle (e.g. if we change to *SigType e.g. for performance reasons) then
	// this code would test it. The strings.Contains check should likely be changed to errors.Is if
	// we do that.
	var b strings.Builder
	var c Context
	if err := c.writeType(&b, &p1, ArchAll); err == nil {
		t.Fatalf("expected error due to detected cycle, but no error was returned")
	} else if !strings.Contains(err.Error(), "cycle") {
		t.Fatalf("got an error, but not a cycle detection error: %v", err)
	}
}

func TestContextSelectTypeDefRejectsAmbiguousMatch(t *testing.T) {
	key := qualifiedTypeName{Namespace: "Windows.Win32.Test", Name: "AMBIGUOUS"}
	context := Context{
		typeDefsByName:        map[qualifiedTypeName]winmd.Index{key: 1},
		typeDefNameDuplicates: map[winmd.Index][]winmd.Index{1: {1, 2}},
		typeDefSupportedArch:  make(map[winmd.Index]Arch),
	}
	err := context.SelectTypeDef(key.Namespace, key.Name, "")
	if err == nil || !strings.Contains(err.Error(), "ambiguous WinMD type") {
		t.Fatalf("SelectTypeDef() error = %v; want ambiguous WinMD type", err)
	}
}

func TestAuthenticatedCipherModeInfoABILayout(t *testing.T) {
	metadata, err := winmd.Open("../../../../winmd/testdata/Windows.Win32.winmd")
	if err != nil {
		t.Fatal(err)
	}
	context, err := NewContext(metadata)
	if err != nil {
		t.Fatal(err)
	}
	const namespace = "Windows.Win32.Security.Cryptography"
	const name = "BCRYPT_AUTHENTICATED_CIPHER_MODE_INFO"
	if err := context.SelectTypeDef(namespace, name, "AUTHENTICATED_CIPHER_MODE_INFO"); err != nil {
		t.Fatal(err)
	}
	index, ok := context.typeDefsByName[qualifiedTypeName{Namespace: namespace, Name: name}]
	if !ok || len(context.typeDefNameDuplicates[index]) != 0 {
		t.Fatal("expected one TypeDef")
	}
	def := context.resolvedDefsByIndex[index]

	tests := []struct {
		arch           Arch
		wantSize       uint32
		wantABIAlign   uint32
		wantGoAlign    uint32
		wantDataPad    uint32
		wantTailPad    uint32
		wantDataOffset uint32
	}{
		{Arch386, 64, 8, 4, 4, 4, 48},
		{ArchAMD64, 88, 8, 8, 0, 0, 72},
		{ArchARM64, 88, 8, 8, 0, 0, 72},
	}
	for _, test := range tests {
		t.Run(test.arch.String(), func(t *testing.T) {
			layout, err := context.planStructABI(def, test.arch, nil)
			if err != nil {
				t.Fatal(err)
			}
			if layout.typeLayout.abiSize != test.wantSize || layout.typeLayout.goSize != test.wantSize {
				t.Fatalf("size = ABI %d, Go %d; want %d", layout.typeLayout.abiSize, layout.typeLayout.goSize, test.wantSize)
			}
			if layout.typeLayout.abiAlign != test.wantABIAlign || layout.typeLayout.goAlign != test.wantGoAlign {
				t.Fatalf("alignment = ABI %d, Go %d; want ABI %d, Go %d", layout.typeLayout.abiAlign, layout.typeLayout.goAlign, test.wantABIAlign, test.wantGoAlign)
			}
			if layout.tailPadding != test.wantTailPad {
				t.Fatalf("tail padding = %d; want %d", layout.tailPadding, test.wantTailPad)
			}
			var foundData bool
			for _, fieldLayout := range layout.fields {
				field, err := metadata.Tables.Field.At(fieldLayout.index)
				if err != nil {
					t.Fatal(err)
				}
				if field.Name.String() == "cbData" {
					foundData = true
					if fieldLayout.offset != test.wantDataOffset || fieldLayout.padding != test.wantDataPad {
						t.Fatalf("cbData = offset %d, padding %d; want offset %d, padding %d", fieldLayout.offset, fieldLayout.padding, test.wantDataOffset, test.wantDataPad)
					}
				}
			}
			if !foundData {
				t.Fatal("cbData field not found")
			}
		})
	}
}

func TestDiscoverABILayoutDependenciesThroughPointerTypedef(t *testing.T) {
	metadata, err := winmd.Open("../../../../winmd/testdata/Windows.Win32.winmd")
	if err != nil {
		t.Fatal(err)
	}
	context, err := NewContext(metadata)
	if err != nil {
		t.Fatal(err)
	}
	if err := context.SelectTypeDef(
		"Windows.Win32.Security.Cryptography",
		"BCRYPT_OAEP_PADDING_INFO",
		"OAEP_PADDING_INFO",
	); err != nil {
		t.Fatal(err)
	}
	if err := context.discoverABILayoutDependencies(); err != nil {
		t.Fatal(err)
	}

	var foundPWSTR bool
	for key, first := range context.typeDefsByName {
		if key.Name != "PWSTR" {
			continue
		}
		indices := context.typeDefNameDuplicates[first]
		if len(indices) == 0 {
			indices = []winmd.Index{first}
		}
		for _, index := range indices {
			if context.abiLayoutTypeDefs[index] {
				foundPWSTR = true
			}
		}
	}
	if !foundPWSTR {
		t.Fatal("layout dependency PWSTR was not discovered through BCRYPT_OAEP_PADDING_INFO")
	}
}
