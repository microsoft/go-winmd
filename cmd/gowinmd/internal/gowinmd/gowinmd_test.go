// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package gowinmd

import (
	"encoding/binary"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestDecodeInt16AttributeField(t *testing.T) {
	value := []byte{1, 0, 1, 0, 0x53, byte(winmd.ElementType_I2), 15}
	value = append(value, "CountParamIndex"...)
	value = binary.LittleEndian.AppendUint16(value, 3)

	got, ok, err := decodeInt16AttributeField(value, "CountParamIndex")
	if err != nil {
		t.Fatal(err)
	}
	if !ok || got != 3 {
		t.Fatalf("decodeInt16AttributeField() = %v, %v; want 3, true", got, ok)
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
		typeDefsByName: map[qualifiedTypeName][]winmd.Index{
			key: {1, 2},
		},
		typeDefSupportedArch: make(map[winmd.Index]Arch),
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
	indices := context.typeDefsByName[qualifiedTypeName{Namespace: namespace, Name: name}]
	if len(indices) != 1 {
		t.Fatalf("found %d TypeDefs; want 1", len(indices))
	}
	def := context.resolvedDefsByIndex[indices[0]]

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
	for key, indices := range context.typeDefsByName {
		if key.Name != "PWSTR" {
			continue
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
