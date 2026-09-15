// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

// Synthetic tables exercise constructor/type resolution without building a PE file.
func customAttributeTestTable[T any](rows ...T) Table[T] {
	data := make([]byte, len(rows))
	for i := range data {
		data[i] = byte(i)
	}
	return Table[T]{
		len: uint32(len(rows)), width: 1, data: data,
		decode: func(r recordReader) (T, error) { return rows[r.data[0]], nil },
	}
}

func customAttributeTestString(s string) String { return String{data: []byte(s)} }

func customAttributeTestMetadata(signature []byte) *Metadata {
	str := customAttributeTestString
	return &Metadata{Tables: &Tables{
		Module:      customAttributeTestTable(Module{Name: str("Test.winmd")}),
		Assembly:    customAttributeTestTable(Assembly{Name: str("TestAssembly")}),
		AssemblyRef: customAttributeTestTable(AssemblyRef{Name: str("System.Runtime"), MajorVersion: 4}, AssemblyRef{Name: str("Other"), MajorVersion: 1}),
		MemberRef:   customAttributeTestTable(MemberRef{Name: str(".ctor"), Signature: signature}),
		MethodDef:   customAttributeTestTable(MethodDef{Name: str(".ctor"), Signature: signature}),
		TypeRef: customAttributeTestTable(
			TypeRef{Name: str("Enum"), Namespace: str("System"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_AssemblyRef}},
			TypeRef{Name: str("Mode"), Namespace: str("Test"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_Module}},
			TypeRef{Name: str("Type"), Namespace: str("System"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_AssemblyRef}},
			TypeRef{Name: str("Mode"), Namespace: str("External"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_AssemblyRef, Index: 1}},
			TypeRef{Name: str("Outer"), Namespace: str("Test"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_Module}},
			TypeRef{Name: str("Inner"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_TypeRef, Index: 4}},
			TypeRef{Name: str("NotEnum"), Namespace: str("Test"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_Module}},
		),
		TypeDef: customAttributeTestTable(
			TypeDef{Name: str("Attribute"), Namespace: str("Test"), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_Null}},
			TypeDef{Name: str("Mode"), Namespace: str("Test"), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeRef}, FieldList: Slice{Start: 0, End: 1}},
			TypeDef{Name: str("Outer"), Namespace: str("Test"), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_Null}},
			TypeDef{Name: str("Inner"), Flags: TypeAttributes_NestedPublic, Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeRef}, FieldList: Slice{Start: 1, End: 2}},
			TypeDef{Name: str("NotEnum"), Namespace: str("Test"), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_Null}},
		),
		Field: customAttributeTestTable(
			Field{Name: str("value__"), Signature: []byte{0x06, 0x06}},
			Field{Name: str("value__"), Signature: []byte{0x06, 0x05}},
		),
		NestedClass: customAttributeTestTable(NestedClass{NestedClass: 3, EnclosingClass: 2}),
		TypeSpec:    customAttributeTestTable(TypeSpec{}),
	}}
}

func customAttributeTestValue(payload []byte) []byte {
	data := append([]byte{1, 0}, payload...)
	return append(data, 0, 0)
}

func TestCustomAttributeDecoderConstructors(t *testing.T) {
	mode := CustomAttributeArgumentType{Kind: ElementType_ENUM, EnumName: "Test.Mode", EnumUnderlyingType: ElementType_I2}
	inner := CustomAttributeArgumentType{Kind: ElementType_ENUM, EnumName: "Test.Outer+Inner", EnumUnderlyingType: ElementType_U1}
	i2 := CustomAttributeArgumentType{Kind: ElementType_I2}
	for _, test := range []struct {
		name      string
		signature []byte
		payload   []byte
		want      []CustomAttributeArgument
	}{
		{"empty", []byte{0x20, 0, 1}, nil, nil},
		{"primitive", []byte{0x20, 1, 1, 6}, []byte{3, 0}, []CustomAttributeArgument{{Type: i2, Value: int16(3)}}},
		{"type-def-enum", []byte{0x20, 1, 1, 0x11, 8}, []byte{0xFF, 0xFF}, []CustomAttributeArgument{{Type: mode, Value: int16(-1)}}},
		{"type-ref-enum", []byte{0x20, 1, 1, 0x11, 9}, []byte{0xFF, 0xFF}, []CustomAttributeArgument{{Type: mode, Value: int16(-1)}}},
		{"nested-type-def", []byte{0x20, 1, 1, 0x11, 16}, []byte{7}, []CustomAttributeArgument{{Type: inner, Value: uint8(7)}}},
		{"nested-type-ref", []byte{0x20, 1, 1, 0x11, 25}, []byte{7}, []CustomAttributeArgument{{Type: inner, Value: uint8(7)}}},
		{"system-type", []byte{0x20, 1, 1, 0x12, 13}, append([]byte{12}, "Example.Type"...), []CustomAttributeArgument{{Type: CustomAttributeArgumentType{Kind: ElementType_TYPE}, Value: "Example.Type"}}},
		{"boxed", []byte{0x20, 1, 1, 0x1C}, []byte{4, 0xFE}, []CustomAttributeArgument{{Type: CustomAttributeArgumentType{Kind: ElementType_BOXED_OBJECT}, Value: CustomAttributeArgument{Type: CustomAttributeArgumentType{Kind: ElementType_I1}, Value: int8(-2)}}}},
		{"array", []byte{0x20, 1, 1, 0x1D, 6}, []byte{1, 0, 0, 0, 3, 0}, []CustomAttributeArgument{{Type: CustomAttributeArgumentType{Kind: ElementType_SZARRAY, Element: &i2}, Value: []CustomAttributeArgument{{Type: i2, Value: int16(3)}}}}},
		{"modifier", []byte{0x20, 1, 1, 0x1F, 13, 6}, []byte{3, 0}, []CustomAttributeArgument{{Type: i2, Value: int16(3)}}},
		{"fixed-and-named-boundary", []byte{0x20, 2, 1, 6, 0x0E}, []byte{3, 0, 1, 'A'}, []CustomAttributeArgument{{Type: i2, Value: int16(3)}, {Type: CustomAttributeArgumentType{Kind: ElementType_STRING}, Value: "A"}}},
	} {
		t.Run(test.name, func(t *testing.T) {
			for _, tag := range []CustomAttributeType{CustomAttributeType_MethodDef, CustomAttributeType_MemberRef} {
				d := NewCustomAttributeDecoder(customAttributeTestMetadata(test.signature))
				a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: tag}, Value: customAttributeTestValue(test.payload)}
				for range 2 { // Both uncached and cached constructor paths.
					got, err := d.Decode(a)
					if err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(got.FixedArguments, test.want) || len(got.NamedArguments) != 0 {
						t.Fatalf("decoded attribute = %#v; want fixed %#v", got, test.want)
					}
				}
			}
		})
	}
}

func TestCustomAttributeDecoderInvalidSignatures(t *testing.T) {
	for _, test := range []struct {
		name string
		sig  []byte
	}{
		{"nil", nil},
		{"header-only", []byte{0x20}},
		{"missing-return", []byte{0x20, 0}},
		{"missing-param", []byte{0x20, 1, 1}},
		{"truncated-count", []byte{0x20, 0x80}},
		{"huge-count", []byte{0x20, 0xDF, 0xFF, 0xFF, 0xFF, 1}},
		{"static", []byte{0, 0, 1}},
		{"varargs", []byte{0x25, 0, 1}},
		{"generic", []byte{0x30, 1, 0, 1}},
		{"return-type", []byte{0x20, 0, 2}},
		{"trailing", []byte{0x20, 0, 1, 0}},
		{"null-handle", []byte{0x20, 1, 1, 0x11, 0}},
		{"truncated-handle", []byte{0x20, 1, 1, 0x11, 0x80}},
		{"invalid-handle-tag", []byte{0x20, 1, 1, 0x11, 7}},
		{"out-of-range-handle", []byte{0x20, 1, 1, 0x11, 0x7D}},
		{"type-spec", []byte{0x20, 1, 1, 0x11, 6}},
		{"non-enum", []byte{0x20, 1, 1, 0x11, 29}},
		{"non-system-type-class", []byte{0x20, 1, 1, 0x12, 9}},
		{"byref", []byte{0x20, 1, 1, 0x10, 6}},
		{"multidimensional-array", []byte{0x20, 1, 1, 0x14}},
		{"jagged-array", []byte{0x20, 1, 1, 0x1D, 0x1D, 6}},
		{"missing-modifier-handle", []byte{0x20, 1, 1, 0x1F}},
		{"serialization-type-marker", []byte{0x20, 1, 1, 0x50}},
		{"serialization-object-marker", []byte{0x20, 1, 1, 0x51}},
	} {
		t.Run(test.name, func(t *testing.T) {
			d := NewCustomAttributeDecoder(customAttributeTestMetadata(test.sig))
			a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: []byte{1, 0, 0, 0}}
			got, err := d.Decode(a)
			if err == nil || !reflect.DeepEqual(got, CustomAttributeValue{}) {
				t.Fatalf("invalid signature = (%#v, %v); want (zero, error)", got, err)
			}
		})
	}
}

func TestCustomAttributeDecoderInvalidConstructors(t *testing.T) {
	m := customAttributeTestMetadata([]byte{0x20, 0, 1})
	for _, index := range []CodedIndex[CustomAttributeType]{
		{Tag: CustomAttributeType_Null}, {Tag: CustomAttributeType_Reserved0},
		{Tag: CustomAttributeType_MethodDef, Index: 1}, {Tag: CustomAttributeType_MemberRef, Index: 1},
	} {
		if _, err := NewCustomAttributeDecoder(m).Decode(CustomAttribute{Type: index}); err == nil {
			t.Errorf("invalid constructor %v was accepted", index)
		}
	}
	m.Tables.MemberRef = customAttributeTestTable(MemberRef{Name: customAttributeTestString("Other"), Signature: []byte{0x20, 0, 1}})
	a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}}
	if _, err := NewCustomAttributeDecoder(m).Decode(a); err == nil {
		t.Fatal("non-constructor member was accepted")
	}
	for _, d := range []*CustomAttributeDecoder{nil, {}, NewCustomAttributeDecoder(nil)} {
		if _, err := d.Decode(a); err == nil {
			t.Fatal("missing metadata was accepted")
		}
	}
}

func TestCustomAttributeDecoderExternalEnum(t *testing.T) {
	m := customAttributeTestMetadata([]byte{0x20, 1, 1, 0x11, 17})
	a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: customAttributeTestValue([]byte{0xFF, 0xFF, 0xFF, 0xFF})}
	const enumName = "External.Mode, Other, Version=1.0.0.0"
	var unresolved *UnresolvedEnumError
	if value, err := NewCustomAttributeDecoder(m).Decode(a); !errors.As(err, &unresolved) || unresolved.Name != enumName || !reflect.DeepEqual(value, CustomAttributeValue{}) {
		t.Fatalf("unresolved enum = (%#v, %v); want (zero, UnresolvedEnumError for %q)", value, err, enumName)
	}
	d := NewCustomAttributeDecoder(m)
	var calls int
	d.ResolveEnum = func(name string) (ElementType, error) {
		calls++
		if name != enumName {
			t.Fatalf("external enum name = %q", name)
		}
		return ElementType_I4, nil
	}
	for range 2 {
		value, err := d.Decode(a)
		if err != nil || value.FixedArguments[0].Value != int32(-1) {
			t.Fatalf("external enum = (%#v, %v)", value, err)
		}
	}
	if calls != 1 {
		t.Fatalf("resolver called %d times; want 1", calls)
	}
	wantErr := errors.New("external assembly unavailable")
	d = NewCustomAttributeDecoder(m)
	d.ResolveEnum = func(string) (ElementType, error) { return 0, wantErr }
	if _, err := d.Decode(a); !errors.Is(err, wantErr) || errors.As(err, &unresolved) {
		t.Fatalf("resolver error = %v; want unclassified %v", err, wantErr)
	}
	wantUnresolved := &UnresolvedEnumError{Name: enumName}
	d = NewCustomAttributeDecoder(m)
	d.ResolveEnum = func(string) (ElementType, error) { return 0, wantUnresolved }
	if _, err := d.Decode(a); !errors.As(err, &unresolved) || unresolved != wantUnresolved {
		t.Fatalf("resolver error = %v; want wrapped %v", err, wantUnresolved)
	}
}

func TestCustomAttributeDecoderUnresolvedNamedEnum(t *testing.T) {
	const enumName = "Missing.Mode"
	d := NewCustomAttributeDecoder(customAttributeTestMetadata([]byte{0x20, 0, 1}))
	enumType := append([]byte{byte(ElementType_ENUM)}, attributeString(enumName)...)
	a := CustomAttribute{
		Type:  CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef},
		Value: attributeBlob(nil, namedAttribute(ElementType_FIELD, enumType, "Mode", []byte{1, 0, 0, 0})),
	}
	value, err := d.Decode(a)
	var unresolved *UnresolvedEnumError
	if !errors.As(err, &unresolved) || unresolved.Name != enumName || !reflect.DeepEqual(value, CustomAttributeValue{}) {
		t.Fatalf("unresolved named enum = (%#v, %v); want (zero, UnresolvedEnumError for %q)", value, err, enumName)
	}
	// A caller can skip this attribute and continue using the same decoder.
	a.Value = attributeBlob(nil)
	if _, err := d.Decode(a); err != nil {
		t.Fatalf("decode after unresolved attribute: %v", err)
	}
}

func TestCustomAttributeDecoderNamedEnum(t *testing.T) {
	for _, name := range []string{"Test.Outer+Inner", "Test.Outer+Inner, TestAssembly"} {
		t.Run(name, func(t *testing.T) {
			d := NewCustomAttributeDecoder(customAttributeTestMetadata([]byte{0x20, 0, 1}))
			data := append([]byte{1, 0, 1, 0, 0x53, 0x55, byte(len(name))}, name...)
			data = append(data, 4, 'M', 'o', 'd', 'e', 7)
			value, err := d.Decode(CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: data})
			if err != nil {
				t.Fatal(err)
			}
			argument := value.NamedArguments[0]
			if argument.Type.EnumName != name || argument.Type.EnumUnderlyingType != ElementType_U1 || argument.Value != uint8(7) {
				t.Fatalf("named enum = %#v", argument)
			}
		})
	}
}

func TestCustomAttributeDecoderMalformedEnum(t *testing.T) {
	for _, field := range []Field{
		{Name: customAttributeTestString("value__")},
		{Name: customAttributeTestString("value__"), Signature: []byte{6, byte(ElementType_R4)}},
		{Name: customAttributeTestString("value__"), Signature: []byte{6, byte(ElementType_I2), 0}},
		{Name: customAttributeTestString("value__"), Signature: []byte{6, byte(ElementType_I2)}, Flags: FieldAttributes_Static},
	} {
		m := customAttributeTestMetadata([]byte{0x20, 1, 1, 0x11, 9})
		m.Tables.Field = customAttributeTestTable(field)
		a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: customAttributeTestValue([]byte{0, 0})}
		var unresolved *UnresolvedEnumError
		if _, err := NewCustomAttributeDecoder(m).Decode(a); err == nil || errors.As(err, &unresolved) {
			t.Fatalf("malformed enum field %#v: got %v; want a malformed-metadata error", field, err)
		}
	}
}

func TestCustomAttributeDecoderBooleanAndCharacterEnums(t *testing.T) {
	for _, test := range []struct {
		name    string
		typ     ElementType
		payload []byte
		want    any
	}{
		{"false", ElementType_BOOLEAN, []byte{0}, false},
		{"true", ElementType_BOOLEAN, []byte{1}, true},
		{"char", ElementType_CHAR, []byte{0xAC, 0x20}, uint16(0x20AC)},
		{"surrogate", ElementType_CHAR, []byte{0, 0xD8}, uint16(0xD800)},
	} {
		t.Run(test.name, func(t *testing.T) {
			m := customAttributeTestMetadata([]byte{0x20, 1, 1, 0x11, 9})
			m.Tables.Field = customAttributeTestTable(Field{Name: customAttributeTestString("MyValue"), Signature: []byte{6, byte(test.typ)}})
			d := NewCustomAttributeDecoder(m)
			a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: customAttributeTestValue(test.payload)}
			got, err := d.Decode(a)
			if err != nil {
				t.Fatal(err)
			}
			if arg := got.FixedArguments[0]; arg.Type.EnumUnderlyingType != test.typ || arg.Value != test.want {
				t.Fatalf("fixed enum = %#v; want %v (%T)", arg, test.want, test.want)
			}
			// Also exercise serialized enum names resolved by the external callback.
			enumType := append([]byte{byte(ElementType_ENUM)}, attributeString("External.Mode")...)
			m = customAttributeTestMetadata([]byte{0x20, 0, 1})
			d = NewCustomAttributeDecoder(m)
			d.ResolveEnum = func(string) (ElementType, error) { return test.typ, nil }
			a.Value = attributeBlob(nil, namedAttribute(ElementType_FIELD, enumType, "Mode", test.payload))
			got, err = d.Decode(a)
			if err != nil {
				t.Fatal(err)
			}
			if arg := got.NamedArguments[0]; arg.Type.EnumUnderlyingType != test.typ || arg.Value != test.want {
				t.Fatalf("named enum = %#v; want %v (%T)", arg, test.want, test.want)
			}
		})
	}
}

func TestCustomAttributeDecoderNativeEnum(t *testing.T) {
	for _, typ := range []ElementType{ElementType_I, ElementType_U} {
		m := customAttributeTestMetadata([]byte{0x20, 1, 1, 0x11, 9})
		m.Tables.Field = customAttributeTestTable(Field{Name: customAttributeTestString("MyValue"), Signature: []byte{6, byte(typ)}})
		a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: customAttributeTestValue(make([]byte, 8))}
		if _, err := NewCustomAttributeDecoder(m).Decode(a); err == nil || !strings.Contains(err.Error(), "unsupported underlying type") {
			t.Fatalf("native enum %v: got %v; want unsupported-type error", typ, err)
		}
	}
}

func TestSplitTypeName(t *testing.T) {
	for _, test := range []struct {
		name, typ, assembly string
		qualified           bool
	}{
		{"Example.Mode", "Example.Mode", "", false},
		{"Example.Mode, Assembly", "Example.Mode", " Assembly", true},
		{`Example.Mode\,Part`, `Example.Mode\,Part`, "", false},
		{`Example.Mode\,Part, Assembly`, `Example.Mode\,Part`, " Assembly", true},
		{`Example.Mode\\, Assembly`, `Example.Mode\\`, " Assembly", true},
		{`Example.Mode\\\,Part`, `Example.Mode\\\,Part`, "", false},
		{`Example.Mode\[A\,B\], Assembly`, `Example.Mode\[A\,B\]`, " Assembly", true},
		{" Example.Mode , Assembly ", " Example.Mode ", " Assembly ", true},
		{`Example.Mode\`, `Example.Mode\`, "", false},
	} {
		t.Run(test.name, func(t *testing.T) {
			typ, assembly, qualified := splitTypeName(test.name)
			if typ != test.typ || assembly != test.assembly || qualified != test.qualified {
				t.Fatalf("splitTypeName() = (%q, %q, %v); want (%q, %q, %v)", typ, assembly, qualified, test.typ, test.assembly, test.qualified)
			}
		})
	}
}

func TestCustomAttributeDecoderEscapedTypeNames(t *testing.T) {
	for _, test := range []struct {
		namespace, name, serialized string
	}{
		{"Test", "Mode+Part", `Test.Mode\+Part`},
		{"Test+Space", "Mode,Part", `Test\+Space.Mode\,Part`},
		{"Test", `Mode\&*[]`, `Test.Mode\\\&\*\[\]`},
		{"", " Mode ", " Mode "},
		{"Test", "Mode ", "Test.Mode "},
	} {
		t.Run(test.serialized, func(t *testing.T) {
			m := customAttributeTestMetadata([]byte{0x20, 0, 1})
			def, _ := m.Tables.TypeDef.At(1)
			base, _ := m.Tables.TypeRef.At(0)
			ref, _ := m.Tables.TypeRef.At(1)
			def.Namespace, def.Name = customAttributeTestString(test.namespace), customAttributeTestString(test.name)
			ref.Namespace, ref.Name = def.Namespace, def.Name
			m.Tables.TypeDef = customAttributeTestTable(def)
			m.Tables.TypeRef = customAttributeTestTable(base, ref)
			m.Tables.NestedClass = Table[NestedClass]{}
			d := NewCustomAttributeDecoder(m)
			for _, handle := range []CodedIndex[TypeDefOrRefOrSpec]{{Tag: TypeDefOrRefOrSpec_TypeDef}, {Tag: TypeDefOrRefOrSpec_TypeRef, Index: 1}} {
				if got, err := d.typeName(handle, 0); got != test.serialized || err != nil {
					t.Fatalf("typeName() = (%q, %v); want (%q, nil)", got, err, test.serialized)
				}
			}
			for _, suffix := range []string{"", ", TestAssembly"} {
				name := test.serialized + suffix
				enumType := append([]byte{byte(ElementType_ENUM)}, attributeString(name)...)
				a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: attributeBlob(nil, namedAttribute(ElementType_FIELD, enumType, "Mode", []byte{1, 0}))}
				value, err := d.Decode(a)
				if err != nil {
					t.Fatal(err)
				}
				if arg := value.NamedArguments[0]; arg.Type.EnumName != name || arg.Value != int16(1) {
					t.Fatalf("decoded named enum = %#v; want name %q, value int16(1)", arg, name)
				}
			}
		})
	}
}

func TestCustomAttributeDecoderEscapedNestedTypeRef(t *testing.T) {
	m := customAttributeTestMetadata(nil)
	m.Tables.TypeRef = customAttributeTestTable(
		TypeRef{Name: customAttributeTestString("Outer,+Part"), Namespace: customAttributeTestString("Test"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_AssemblyRef, Index: 1}},
		TypeRef{Name: customAttributeTestString("Inner[]"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_TypeRef}},
	)
	d := NewCustomAttributeDecoder(m)
	want := `Test.Outer\,\+Part+Inner\[\], Other, Version=1.0.0.0`
	if got, err := d.typeName(CodedIndex[TypeDefOrRefOrSpec]{Tag: TypeDefOrRefOrSpec_TypeRef, Index: 1}, 0); got != want || err != nil {
		t.Fatalf("nested type name = (%q, %v); want (%q, nil)", got, err, want)
	}
}

func TestCustomAttributeDecoderTypeCycles(t *testing.T) {
	m := customAttributeTestMetadata([]byte{0x20, 1, 1, 0x11, 9})
	m.Tables.NestedClass = customAttributeTestTable(NestedClass{NestedClass: 1, EnclosingClass: 1})
	a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: customAttributeTestValue([]byte{0, 0})}
	if _, err := NewCustomAttributeDecoder(m).Decode(a); err == nil || !strings.Contains(err.Error(), "cyclic") {
		t.Fatalf("cyclic TypeDef error = %v", err)
	}
	m = customAttributeTestMetadata([]byte{0x20, 1, 1, 0x11, 5})
	m.Tables.TypeRef = customAttributeTestTable(TypeRef{Name: customAttributeTestString("Self"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_TypeRef}})
	if _, err := NewCustomAttributeDecoder(m).Decode(a); err == nil || !strings.Contains(err.Error(), "cyclic") {
		t.Fatalf("cyclic TypeRef error = %v", err)
	}
}
