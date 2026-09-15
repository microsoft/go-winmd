// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"errors"
	"io"
	"testing"
)

func enumTestMetadata(fields ...Field) *Metadata {
	str := customAttributeTestString
	return &Metadata{Tables: &Tables{
		TypeDef: customAttributeTestTable(
			TypeDef{Name: str("Mode"), Namespace: str("Test"), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeRef}, FieldList: Slice{End: Index(len(fields))}},
			TypeDef{Name: str("Enum"), Namespace: str("System"), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_Null}},
		),
		TypeRef:  customAttributeTestTable(TypeRef{Name: str("Enum"), Namespace: str("System"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_AssemblyRef}}),
		TypeSpec: customAttributeTestTable(TypeSpec{}),
		Field:    customAttributeTestTable(fields...),
	}}
}

func TestEnumUnderlyingType(t *testing.T) {
	for _, typ := range []ElementType{
		ElementType_BOOLEAN, ElementType_CHAR,
		ElementType_I1, ElementType_U1, ElementType_I2, ElementType_U2,
		ElementType_I4, ElementType_U4, ElementType_I8, ElementType_U8,
		ElementType_I, ElementType_U,
	} {
		t.Run(typ.String(), func(t *testing.T) {
			for _, name := range []string{"value__", "MyValue"} {
				m := enumTestMetadata(
					Field{Name: customAttributeTestString("Member"), Flags: FieldAttributes_Static | FieldAttributes_Literal},
					Field{Name: customAttributeTestString(name), Signature: []byte{sigKind_FIELD, byte(typ)}},
				)
				if got, err := m.EnumUnderlyingType(0); got != typ || err != nil {
					t.Fatalf("TypeRef base, field %s: got (%v, %v); want (%v, nil)", name, got, err, typ)
				}
				def, _ := m.Tables.TypeDef.At(0)
				base, _ := m.Tables.TypeDef.At(1)
				def.Extends = CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeDef, Index: 1}
				m.Tables.TypeDef = customAttributeTestTable(def, base)
				if got, err := m.EnumUnderlyingType(0); got != typ || err != nil {
					t.Fatalf("TypeDef base, field %s: got (%v, %v); want (%v, nil)", name, got, err, typ)
				}
			}
		})
	}
}

func TestEnumUnderlyingTypeModifiers(t *testing.T) {
	for _, signature := range []SigFieldBlob{
		{6, byte(ElementType_CMOD_OPT), 5, byte(ElementType_I2)},
		{6, byte(ElementType_CMOD_REQD), 8, byte(ElementType_CMOD_OPT), 5, byte(ElementType_I2)},
		{6, byte(ElementType_CMOD_OPT), 6, byte(ElementType_I2)},
	} {
		m := enumTestMetadata(Field{Name: customAttributeTestString("MyValue"), Signature: signature})
		if got, err := m.EnumUnderlyingType(0); got != ElementType_I2 || err != nil {
			t.Fatalf("signature %x: got (%v, %v); want (I2, nil)", signature, got, err)
		}
	}
}

func TestEnumUnderlyingTypeInvalidSignatures(t *testing.T) {
	for _, test := range []struct {
		name      string
		signature SigFieldBlob
		wantEOF   bool
	}{
		{"empty", nil, true},
		{"missing-type", []byte{6}, true},
		{"wrong-kind", []byte{0, 6}, false},
		{"float", []byte{6, byte(ElementType_R4)}, false},
		{"string", []byte{6, byte(ElementType_STRING)}, false},
		{"array", []byte{6, byte(ElementType_ARRAY), byte(ElementType_U1)}, false},
		{"trailing-data", []byte{6, byte(ElementType_I2), 0}, false},
		{"missing-modifier-handle", []byte{6, byte(ElementType_CMOD_OPT)}, true},
		{"truncated-modifier-handle", []byte{6, byte(ElementType_CMOD_OPT), 0x80}, true},
		{"null-modifier-handle", []byte{6, byte(ElementType_CMOD_OPT), 0, 6}, false},
		{"invalid-modifier-tag", []byte{6, byte(ElementType_CMOD_OPT), 7, 6}, false},
		{"zero-modifier-row", []byte{6, byte(ElementType_CMOD_OPT), 1, 6}, false},
		{"modifier-out-of-range", []byte{6, byte(ElementType_CMOD_OPT), 0x7D, 6}, false},
		{"modifier-without-type", []byte{6, byte(ElementType_CMOD_OPT), 5}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			m := enumTestMetadata(Field{Name: customAttributeTestString("MyValue"), Signature: test.signature})
			got, err := m.EnumUnderlyingType(0)
			if got != 0 || err == nil || errors.Is(err, io.ErrUnexpectedEOF) != test.wantEOF {
				t.Fatalf("EnumUnderlyingType() = (%v, %v); want (zero, error), unexpected EOF %v", got, err, test.wantEOF)
			}
		})
	}
}

func TestEnumUnderlyingTypeInstanceFields(t *testing.T) {
	str := customAttributeTestString
	backing := Field{Name: str("value__"), Signature: []byte{6, byte(ElementType_I4)}}
	for _, fields := range [][]Field{
		nil,
		{{Name: str("Member"), Flags: FieldAttributes_Static | FieldAttributes_Literal}},
		{backing, {Name: str("Other"), Signature: backing.Signature}},
	} {
		if got, err := enumTestMetadata(fields...).EnumUnderlyingType(0); got != 0 || err == nil {
			t.Fatalf("fields %#v: got (%v, %v); want (zero, error)", fields, got, err)
		}
	}
	m := enumTestMetadata(backing)
	def, _ := m.Tables.TypeDef.At(0)
	def.FieldList.End++
	m.Tables.TypeDef = customAttributeTestTable(def)
	if got, err := m.EnumUnderlyingType(0); got != 0 || err == nil {
		t.Fatalf("invalid field index: got (%v, %v); want (zero, error)", got, err)
	}
}

func TestEnumUnderlyingTypeInvalidBase(t *testing.T) {
	str := customAttributeTestString
	backing := Field{Name: str("MyValue"), Signature: []byte{6, byte(ElementType_I4)}}
	for _, base := range []TypeRef{
		{Name: str("ValueType"), Namespace: str("System")},
		{Name: str("Enum"), Namespace: str("Other")},
		{Name: str("Enum"), Namespace: str("System"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_TypeRef}},
	} {
		m := enumTestMetadata(backing)
		m.Tables.TypeRef = customAttributeTestTable(base)
		if got, err := m.EnumUnderlyingType(0); got != 0 || err == nil {
			t.Fatalf("invalid base %#v: got (%v, %v); want (zero, error)", base, got, err)
		}
	}
	for _, extends := range []CodedIndex[TypeDefOrRef]{
		{Tag: TypeDefOrRef_Null}, {Tag: TypeDefOrRef_TypeSpec},
		{Tag: TypeDefOrRef_TypeRef, Index: 1}, {Tag: TypeDefOrRef_TypeDef, Index: 2},
	} {
		m := enumTestMetadata(backing)
		def, _ := m.Tables.TypeDef.At(0)
		base, _ := m.Tables.TypeDef.At(1)
		def.Extends = extends
		m.Tables.TypeDef = customAttributeTestTable(def, base)
		if got, err := m.EnumUnderlyingType(0); got != 0 || err == nil {
			t.Fatalf("invalid base %v: got (%v, %v); want (zero, error)", extends, got, err)
		}
	}
	m := enumTestMetadata(backing)
	def, _ := m.Tables.TypeDef.At(0)
	base, _ := m.Tables.TypeDef.At(1)
	def.Extends = CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeDef, Index: 1}
	base.Flags = TypeAttributes_NestedPublic
	m.Tables.TypeDef = customAttributeTestTable(def, base)
	if got, err := m.EnumUnderlyingType(0); got != 0 || err == nil {
		t.Fatalf("nested base: got (%v, %v); want (zero, error)", got, err)
	}
	if got, err := m.EnumUnderlyingType(2); got != 0 || err == nil {
		t.Fatalf("invalid TypeDef index: got (%v, %v); want (zero, error)", got, err)
	}
}
