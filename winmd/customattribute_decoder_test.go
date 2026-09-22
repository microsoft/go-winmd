// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"bytes"
	"errors"
	"fmt"
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
		decode: func(r recordReader) (T, string, error) { return rows[r.data[0]], "", nil },
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
			TypeDef{Name: str("Inner"), Flags: TypeAttributes(TypeVisibility_NestedPublic), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeRef}, FieldList: Slice{Start: 1, End: 2}},
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

func customAttributeTestEnum(tag TypeDefOrRefOrSpec, index Index, namespace, name string, underlying ElementType, declaring ...string) CustomAttributeArgumentType {
	return CustomAttributeArgumentType{
		Kind: ElementType_ENUM,
		Enum: EnumReference{Metadata: &EnumReferenceMetadata{
			Handle:    CodedIndex[TypeDefOrRefOrSpec]{Tag: tag, Index: index},
			Namespace: namespace, Name: name, DeclaringTypes: declaring,
		}},
		EnumUnderlyingType: underlying,
	}
}

func TestCustomAttributeDecoderConstructors(t *testing.T) {
	modeDef := customAttributeTestEnum(TypeDefOrRefOrSpec_TypeDef, 1, "Test", "Mode", ElementType_I2)
	modeRef := customAttributeTestEnum(TypeDefOrRefOrSpec_TypeRef, 1, "Test", "Mode", ElementType_I2)
	innerDef := customAttributeTestEnum(TypeDefOrRefOrSpec_TypeDef, 3, "Test", "Inner", ElementType_U1, "Outer")
	innerRef := customAttributeTestEnum(TypeDefOrRefOrSpec_TypeRef, 5, "Test", "Inner", ElementType_U1, "Outer")
	i2 := CustomAttributeArgumentType{Kind: ElementType_I2}
	for _, test := range []struct {
		name      string
		signature []byte
		payload   []byte
		want      []CustomAttributeArgument
	}{
		{"empty", []byte{0x20, 0, 1}, nil, nil},
		{"primitive", []byte{0x20, 1, 1, 6}, []byte{3, 0}, []CustomAttributeArgument{{Type: i2, Value: int16(3)}}},
		{"type-def-enum", []byte{0x20, 1, 1, 0x11, 8}, []byte{0xFF, 0xFF}, []CustomAttributeArgument{{Type: modeDef, Value: int16(-1)}}},
		{"type-ref-enum", []byte{0x20, 1, 1, 0x11, 9}, []byte{0xFF, 0xFF}, []CustomAttributeArgument{{Type: modeRef, Value: int16(-1)}}},
		{"nested-type-def", []byte{0x20, 1, 1, 0x11, 16}, []byte{7}, []CustomAttributeArgument{{Type: innerDef, Value: uint8(7)}}},
		{"nested-type-ref", []byte{0x20, 1, 1, 0x11, 25}, []byte{7}, []CustomAttributeArgument{{Type: innerRef, Value: uint8(7)}}},
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
		{Tag: CustomAttributeType_Reserved0},
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
	assembly, _ := m.Tables.AssemblyRef.At(1)
	want := EnumReference{Metadata: &EnumReferenceMetadata{
		Handle:    CodedIndex[TypeDefOrRefOrSpec]{Tag: TypeDefOrRefOrSpec_TypeRef, Index: 3},
		Namespace: "External", Name: "Mode", Assembly: &assembly,
	}}
	var unresolved *UnresolvedEnumError
	if value, err := NewCustomAttributeDecoder(m).Decode(a); !errors.As(err, &unresolved) || !reflect.DeepEqual(unresolved.Reference, want) || !reflect.DeepEqual(value, CustomAttributeValue{}) {
		t.Fatalf("unresolved enum = (%#v, %v); want (zero, UnresolvedEnumError for %+v)", value, err, want)
	}
	d := NewCustomAttributeDecoder(m)
	var calls int
	d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
		calls++
		if !reflect.DeepEqual(ref, want) {
			t.Fatalf("external enum reference = %+v; want %+v", ref, want)
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
	d.ResolveEnum = func(EnumReference) (ElementType, error) { return 0, wantErr }
	if _, err := d.Decode(a); !errors.Is(err, wantErr) || errors.As(err, &unresolved) {
		t.Fatalf("resolver error = %v; want unclassified %v", err, wantErr)
	}
	wantUnresolved := &UnresolvedEnumError{Reference: want}
	d = NewCustomAttributeDecoder(m)
	d.ResolveEnum = func(EnumReference) (ElementType, error) { return 0, wantUnresolved }
	if _, err := d.Decode(a); !errors.As(err, &unresolved) || unresolved != wantUnresolved {
		t.Fatalf("resolver error = %v; want wrapped %v", err, wantUnresolved)
	}
}

func TestCustomAttributeDecoderRawAssemblyReference(t *testing.T) {
	ecmaKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0}
	for _, test := range []struct {
		name    string
		key     []byte
		flags   AssemblyFlags
		culture string
	}{
		{"unsigned", nil, 0, ""},
		{"token", []byte{0xb7, 0x7a, 0x5c, 0x56, 0x19, 0x34, 0xe0, 0x89}, 0, "fr-FR"},
		{"public-key", ecmaKey, AssemblyFlags_PublicKey, ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			key := bytes.Clone(test.key)
			assembly := AssemblyRef{
				Name: customAttributeTestString("Shared"), MajorVersion: 1, MinorVersion: 2, BuildNumber: 3, RevisionNumber: 4,
				Culture: customAttributeTestString(test.culture), Flags: test.flags, PublicKeyOrToken: key, HashValue: []byte{1, 2, 3},
			}
			m := customAttributeTestMetadata([]byte{0x20, 1, 1, 0x11, 17})
			m.Tables.AssemblyRef = customAttributeTestTable(AssemblyRef{}, assembly)
			d := NewCustomAttributeDecoder(m)
			calls := 0
			d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
				calls++
				if ref.Metadata == nil || ref.SerializedName != "" || !reflect.DeepEqual(ref.Metadata.Assembly, &assembly) {
					t.Fatalf("original AssemblyRef was not preserved: %+v", ref)
				}
				return ElementType_I2, nil
			}
			value, err := d.Decode(CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: customAttributeTestValue([]byte{1, 0})})
			if err != nil || calls != 1 || len(value.FixedArguments) != 1 || value.FixedArguments[0].Value != int16(1) {
				t.Fatalf("raw reference decode = %+v, %v; calls %d", value, err, calls)
			}
			if !bytes.Equal(key, test.key) {
				t.Fatal("assembly public key/token was modified")
			}
		})
	}
}

func TestCustomAttributeDecoderExternalEnumIdentities(t *testing.T) {
	for _, discriminator := range []string{"token", "culture"} {
		t.Run(discriminator, func(t *testing.T) {
			str := customAttributeTestString
			m := customAttributeTestMetadata(nil)
			first := AssemblyRef{Name: str("Shared"), MajorVersion: 1, PublicKeyOrToken: bytes.Repeat([]byte{0x11}, 8)}
			second := first
			if discriminator == "token" {
				second.PublicKeyOrToken = bytes.Repeat([]byte{0x22}, 8)
			} else {
				second.Culture = str("fr-FR")
			}
			assemblies := []AssemblyRef{first, second}
			m.Tables.AssemblyRef = customAttributeTestTable(first, second)
			m.Tables.TypeRef = customAttributeTestTable(
				TypeRef{Name: str("Mode"), Namespace: str("External"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_AssemblyRef, Index: 0}},
				TypeRef{Name: str("Mode"), Namespace: str("External"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_AssemblyRef, Index: 1}},
			)
			m.Tables.MemberRef = customAttributeTestTable(
				MemberRef{Name: str(".ctor"), Signature: []byte{0x20, 1, 1, 0x11, 5}},
				MemberRef{Name: str(".ctor"), Signature: []byte{0x20, 1, 1, 0x11, 9}},
			)
			d := NewCustomAttributeDecoder(m)
			var calls []Index
			d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
				if ref.Metadata == nil || ref.SerializedName != "" || ref.Metadata.Namespace != "External" || ref.Metadata.Name != "Mode" || ref.Metadata.Handle.Tag != TypeDefOrRefOrSpec_TypeRef {
					t.Fatalf("incomplete external enum reference: %+v", ref)
				}
				index := ref.Metadata.Handle.Index
				if index > 1 || !reflect.DeepEqual(ref.Metadata.Assembly, &assemblies[index]) {
					t.Fatalf("incorrect raw assembly reference: %+v", ref.Metadata)
				}
				calls = append(calls, index)
				switch index {
				case 0:
					return ElementType_I2, nil
				case 1:
					return ElementType_I4, nil
				default:
					t.Fatalf("unexpected enum handle: %v", ref.Metadata.Handle)
					return 0, nil
				}
			}
			for range 2 { // Each identity must retain its own cached underlying type.
				for i, test := range []struct {
					payload []byte
					typ     ElementType
					value   any
				}{
					{[]byte{1, 0}, ElementType_I2, int16(1)},
					{[]byte{2, 0, 0, 0}, ElementType_I4, int32(2)},
				} {
					a := CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef, Index: Index(i)}, Value: customAttributeTestValue(test.payload)}
					value, err := d.Decode(a)
					if err != nil {
						t.Fatal(err)
					}
					if len(value.FixedArguments) != 1 {
						t.Fatalf("decoded arguments = %+v; want one enum", value.FixedArguments)
					}
					arg := value.FixedArguments[0]
					if arg.Type.Enum.Metadata == nil || arg.Type.Enum.SerializedName != "" || arg.Type.Enum.Metadata.Handle.Index != Index(i) || !reflect.DeepEqual(arg.Type.Enum.Metadata.Assembly, &assemblies[i]) || arg.Type.EnumUnderlyingType != test.typ || arg.Value != test.value {
						t.Fatalf("enum identity %d decoded with the wrong type: %+v", i, arg)
					}
				}
			}
			if !reflect.DeepEqual(calls, []Index{0, 1}) {
				t.Fatalf("resolver calls = %v; want one call per metadata handle", calls)
			}
		})
	}
}

func TestCustomAttributeDecoderSeparateEnumCaches(t *testing.T) {
	for _, order := range []string{"metadata-first", "serialized-first"} {
		t.Run(order, func(t *testing.T) {
			m := customAttributeTestMetadata(nil)
			m.Tables.MemberRef = customAttributeTestTable(
				MemberRef{Name: customAttributeTestString(".ctor"), Signature: []byte{0x20, 2, 1, 0x11, 17, 0x1c}},
				MemberRef{Name: customAttributeTestString(".ctor"), Signature: []byte{0x20, 1, 1, 0x1c}},
			)
			d := NewCustomAttributeDecoder(m)
			calls := make(map[string]int)
			d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
				if ref.Metadata != nil {
					if ref.SerializedName != "" || ref.Metadata.Namespace != "External" || ref.Metadata.Name != "Mode" {
						t.Fatalf("unexpected metadata enum: %+v", ref)
					}
					calls["metadata"]++
					return ElementType_I2, nil
				}
				if ref.SerializedName != "External.Mode" {
					t.Fatalf("unexpected serialized enum: %+v", ref)
				}
				calls["serialized"]++
				return ElementType_I4, nil
			}
			enumType := append([]byte{byte(ElementType_ENUM)}, attributeString("External.Mode")...)
			boxed := append(bytes.Clone(enumType), 2, 0, 0, 0)
			if order == "serialized-first" {
				_, err := d.Decode(CustomAttribute{
					Type:  CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef, Index: 1},
					Value: attributeBlob(boxed),
				})
				if err != nil {
					t.Fatal(err)
				}
			}
			a := CustomAttribute{
				Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef},
				Value: attributeBlob(append([]byte{1, 0}, boxed...),
					namedAttribute(ElementType_FIELD, enumType, "Mode", []byte{3, 0, 0, 0})),
			}
			for range 2 {
				value, err := d.Decode(a)
				if err != nil {
					t.Fatal(err)
				}
				if len(value.FixedArguments) != 2 || len(value.NamedArguments) != 1 {
					t.Fatalf("unexpected arguments: %+v", value)
				}
				fixed := value.FixedArguments[0]
				boxed := value.FixedArguments[1].Value.(CustomAttributeArgument)
				named := value.NamedArguments[0]
				if fixed.Type.Enum.Metadata == nil || fixed.Value != int16(1) || boxed.Value != int32(2) || named.Value != int32(3) {
					t.Fatalf("metadata and serialized enum widths collided: %+v", value)
				}
				if boxed.Type.Enum.Metadata != nil || boxed.Type.Enum.SerializedName != "External.Mode" || named.Type.Enum != boxed.Type.Enum {
					t.Fatal("serialized enum references were replaced with metadata references")
				}
			}
			if calls["metadata"] != 1 || calls["serialized"] != 1 {
				t.Fatalf("resolver calls = %v; want one per reference form", calls)
			}
		})
	}
}

func TestCustomAttributeDecoderSerializedReferenceUnchanged(t *testing.T) {
	const name = `External.Outer\+Part+Mode, Shared, PublicKey=000102, Culture=fr-FR, Version=1.2.3.4`
	m := customAttributeTestMetadata([]byte{0x20, 1, 1, 0x1c})
	d := NewCustomAttributeDecoder(m)
	var calls int
	d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
		calls++
		if ref.Metadata != nil || ref.SerializedName != name {
			t.Fatalf("serialized identity was normalized or parsed: %+v", ref)
		}
		return ElementType_I2, nil
	}
	enumType := append([]byte{byte(ElementType_ENUM)}, attributeString(name)...)
	a := CustomAttribute{
		Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef},
		Value: attributeBlob(append(bytes.Clone(enumType), 1, 0),
			namedAttribute(ElementType_PROPERTY, enumType, "Mode", []byte{2, 0})),
	}
	for range 2 {
		value, err := d.Decode(a)
		if err != nil {
			t.Fatal(err)
		}
		if value.NamedArguments[0].Type.Enum.SerializedName != name || value.FixedArguments[0].Value.(CustomAttributeArgument).Type.Enum.SerializedName != name {
			t.Fatal("serialized name not retained in decoded values")
		}
	}
	if calls != 1 || len(d.enumsByName) != 1 || len(d.enumsByType) != 0 {
		t.Fatalf("serialized cache: calls=%d, names=%d, handles=%d", calls, len(d.enumsByName), len(d.enumsByType))
	}
}

func TestCustomAttributeDecoderTypeDefHandleIdentity(t *testing.T) {
	m := customAttributeTestMetadata(nil)
	m.Tables.TypeDef = customAttributeTestTable(
		TypeDef{Name: customAttributeTestString("Mode"), Namespace: customAttributeTestString("Test"), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeRef}, FieldList: Slice{Start: 0, End: 1}},
		TypeDef{Name: customAttributeTestString("Mode"), Namespace: customAttributeTestString("Test"), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeRef}, FieldList: Slice{Start: 1, End: 2}},
	)
	m.Tables.NestedClass = Table[NestedClass]{}
	m.Tables.Field = customAttributeTestTable(
		Field{Signature: []byte{6, byte(ElementType_I2)}},
		Field{Signature: []byte{6, byte(ElementType_I4)}},
	)
	m.Tables.MemberRef = customAttributeTestTable(
		MemberRef{Name: customAttributeTestString(".ctor"), Signature: []byte{0x20, 1, 1, 0x11, 4}},
		MemberRef{Name: customAttributeTestString(".ctor"), Signature: []byte{0x20, 1, 1, 0x11, 8}},
	)
	d := NewCustomAttributeDecoder(m)
	d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
		t.Fatalf("local TypeDef unexpectedly delegated: %+v", ref)
		return 0, nil
	}
	for range 2 {
		for i, payload := range [][]byte{{1, 0}, {2, 0, 0, 0}} {
			value, err := d.Decode(CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef, Index: Index(i)}, Value: customAttributeTestValue(payload)})
			if err != nil {
				t.Fatal(err)
			}
			arg := value.FixedArguments[0]
			if arg.Type.Enum.Metadata.Handle.Index != Index(i) || i == 0 && arg.Value != int16(1) || i == 1 && arg.Value != int32(2) {
				t.Fatalf("same-named TypeDef handles were conflated: %+v", arg)
			}
		}
	}
	if _, err := d.resolveEnum(EnumReference{SerializedName: "Test.Mode"}); err == nil || !strings.Contains(err.Error(), "ambiguous") {
		t.Fatalf("ambiguous serialized name error = %v", err)
	}
}

func TestCustomAttributeDecoderLocalTypeRefIdentity(t *testing.T) {
	for _, test := range []struct {
		name              string
		namespace1, name1 string
		namespace2, name2 string
	}{
		{"namespace-boundary", "Test", "Sub.Mode", "Test.Sub", "Mode"},
		{"empty-namespace", "", "Test.Mode", "Test", "Mode"},
	} {
		t.Run(test.name, func(t *testing.T) {
			str := customAttributeTestString
			m := customAttributeTestMetadata(nil)
			base, err := m.Tables.TypeRef.At(0)
			if err != nil {
				t.Fatal(err)
			}
			m.Tables.TypeRef = customAttributeTestTable(base,
				TypeRef{Namespace: str(test.namespace1), Name: str(test.name1), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_Module}},
				TypeRef{Namespace: str(test.namespace2), Name: str(test.name2), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_Module}},
			)
			m.Tables.TypeDef = customAttributeTestTable(
				TypeDef{Namespace: str(test.namespace1), Name: str(test.name1), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeRef}, FieldList: Slice{0, 1}},
				TypeDef{Namespace: str(test.namespace2), Name: str(test.name2), Extends: CodedIndex[TypeDefOrRef]{Tag: TypeDefOrRef_TypeRef}, FieldList: Slice{1, 2}},
			)
			m.Tables.NestedClass = Table[NestedClass]{}
			m.Tables.Field = customAttributeTestTable(
				Field{Signature: []byte{6, byte(ElementType_I2)}},
				Field{Signature: []byte{6, byte(ElementType_I4)}},
			)
			m.Tables.MemberRef = customAttributeTestTable(
				MemberRef{Name: str(".ctor"), Signature: []byte{0x20, 1, 1, 0x11, 9}},
				MemberRef{Name: str(".ctor"), Signature: []byte{0x20, 1, 1, 0x11, 13}},
			)
			d := NewCustomAttributeDecoder(m)
			d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
				t.Fatalf("local metadata identity unexpectedly delegated: %+v", ref)
				return 0, nil
			}
			for range 2 {
				for i, payload := range [][]byte{{1, 0}, {2, 0, 0, 0}} {
					value, err := d.Decode(CustomAttribute{
						Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef, Index: Index(i)}, Value: customAttributeTestValue(payload),
					})
					if err != nil {
						t.Fatal(err)
					}
					want := []any{int16(1), int32(2)}[i]
					if len(value.FixedArguments) != 1 || value.FixedArguments[0].Value != want {
						t.Fatalf("namespace/name identities collided: %+v; want %v", value, want)
					}
				}
			}
		})
	}
}

func TestMetadataEnumIdentityBoundaries(t *testing.T) {
	refs := []EnumReferenceMetadata{
		{Namespace: "N", Name: "Outer.Mode"},
		{Namespace: "N.Outer", Name: "Mode"},
		{Namespace: "N", Name: "Mode", DeclaringTypes: []string{"Outer+Inner"}},
		{Namespace: "N", Name: "Mode", DeclaringTypes: []string{"Outer", "Inner"}},
		{Namespace: "N", Name: "Mode", DeclaringTypes: []string{"Outer", "Inner.Part"}},
		{Namespace: "N", Name: "Mode", DeclaringTypes: []string{"Outer.Inner", "Part"}},
	}
	seen := make(map[metadataTypeIdentity]bool)
	for _, ref := range refs {
		key := ref.identity()
		if seen[key] {
			t.Fatalf("raw name or nesting boundaries lost for %+v", ref)
		}
		seen[key] = true
	}
	if refs[0].fullName() != refs[1].fullName() {
		t.Fatal("test must include distinct identities with identical serialized names")
	}
}

func TestCustomAttributeDecoderTypeRefDepthCache(t *testing.T) {
	for _, levels := range []int{maxCustomAttributeDepth, maxCustomAttributeDepth + 1, maxCustomAttributeDepth + 2} {
		t.Run(fmt.Sprint(levels), func(t *testing.T) {
			str := customAttributeTestString
			m := customAttributeTestMetadata(nil)
			refs := make([]TypeRef, levels)
			for i := range refs {
				refs[i].Name = str(fmt.Sprintf("T%d", i))
				if i == 0 {
					refs[i].Namespace = str("External")
					refs[i].ResolutionScope = CodedIndex[ResolutionScope]{Tag: ResolutionScope_AssemblyRef, Index: 1}
				} else {
					refs[i].ResolutionScope = CodedIndex[ResolutionScope]{Tag: ResolutionScope_TypeRef, Index: Index(i - 1)}
				}
			}
			m.Tables.TypeRef = customAttributeTestTable(refs...)
			code := uint32(levels)<<2 | 1
			m.Tables.MemberRef = customAttributeTestTable(MemberRef{Name: str(".ctor"), Signature: []byte{0x20, 1, 1, 0x11, 0x80 | byte(code>>8), byte(code)}})
			for _, warm := range []bool{false, true} {
				t.Run(fmt.Sprintf("warm-%v", warm), func(t *testing.T) {
					d := NewCustomAttributeDecoder(m)
					if warm {
						_, err := d.typeReference(CodedIndex[TypeDefOrRefOrSpec]{Tag: TypeDefOrRefOrSpec_TypeRef, Index: maxCustomAttributeDepth - 2}, 0)
						if err != nil {
							t.Fatalf("valid ancestor rejected: %v", err)
						}
					}
					calls := 0
					d.ResolveEnum = func(EnumReference) (ElementType, error) { calls++; return ElementType_I2, nil }
					for range 2 {
						value, err := d.Decode(CustomAttribute{
							Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: customAttributeTestValue([]byte{1, 0}),
						})
						if levels > maxCustomAttributeDepth {
							if err == nil || !strings.Contains(err.Error(), "excessively nested") || calls != 0 {
								t.Fatalf("excessive cached depth accepted: error=%v, calls=%d", err, calls)
							}
						} else if err != nil || len(value.FixedArguments) != 1 || len(value.FixedArguments[0].Type.Enum.Metadata.DeclaringTypes)+1 != levels {
							t.Fatalf("valid depth %d rejected: value=%+v, error=%v", levels, value, err)
						}
					}
				})
			}
		})
	}
}

func TestCustomAttributeDecoderTypeDefDepthIndex(t *testing.T) {
	for _, levels := range []int{maxCustomAttributeDepth, maxCustomAttributeDepth + 1} {
		t.Run(fmt.Sprint(levels), func(t *testing.T) {
			m := customAttributeTestMetadata(nil)
			defs := make([]TypeDef, levels)
			var parents []NestedClass
			for i := range defs {
				defs[i].Name = customAttributeTestString(fmt.Sprintf("T%d", i))
				if i > 0 {
					defs[i].Flags = TypeAttributes(TypeVisibility_NestedPublic)
					parents = append(parents, NestedClass{NestedClass: Index(i), EnclosingClass: Index(i - 1)})
				}
			}
			m.Tables.TypeDef = customAttributeTestTable(defs...)
			m.Tables.NestedClass = customAttributeTestTable(parents...)
			d := NewCustomAttributeDecoder(m)
			for range 2 {
				// The normal topological row order caches every ancestor before
				// visiting a child. That must not bypass the complete depth limit.
				err := d.indexTypes()
				if levels > maxCustomAttributeDepth {
					if err == nil || !strings.Contains(err.Error(), "excessively nested") {
						t.Fatalf("TypeDef depth %d error = %v; want depth error", levels, err)
					}
				} else if err != nil {
					t.Fatalf("valid TypeDef depth %d rejected: %v", levels, err)
				}
			}
		})
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
	if !errors.As(err, &unresolved) || unresolved.Reference.Metadata != nil || unresolved.Reference.SerializedName != enumName || !reflect.DeepEqual(value, CustomAttributeValue{}) {
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
			if argument.Type.Enum.Metadata != nil || argument.Type.Enum.SerializedName != name || argument.Type.EnumUnderlyingType != ElementType_U1 || argument.Value != uint8(7) {
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
		{Name: customAttributeTestString("value__"), Signature: []byte{6, byte(ElementType_I2)}, Flags: FieldAttributes(FieldFlags_Static)},
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
			d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
				if ref.Metadata != nil || ref.SerializedName != "External.Mode" {
					t.Fatalf("serialized reference changed: %+v", ref)
				}
				return test.typ, nil
			}
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
				if got, err := d.typeReference(handle, 0); err != nil || got.Namespace != test.namespace || got.Name != test.name || got.fullName() != test.serialized || got.Assembly != nil {
					t.Fatalf("typeReference() = (%+v, %v); want raw names and local lookup name %q", got, err, test.serialized)
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
				if arg := value.NamedArguments[0]; arg.Type.Enum.Metadata != nil || arg.Type.Enum.SerializedName != name || arg.Value != int16(1) {
					t.Fatalf("decoded named enum = %#v; want name %q, value int16(1)", arg, name)
				}
			}
		})
	}
}

func TestCustomAttributeDecoderEscapedNestedTypeRef(t *testing.T) {
	m := customAttributeTestMetadata(nil)
	m.Tables.TypeRef = customAttributeTestTable(
		TypeRef{Name: customAttributeTestString("Outer,+Part"), Namespace: customAttributeTestString("Test+Space"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_AssemblyRef, Index: 1}},
		TypeRef{Name: customAttributeTestString("Inner[]"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_TypeRef}},
		TypeRef{Name: customAttributeTestString("Mode&"), ResolutionScope: CodedIndex[ResolutionScope]{Tag: ResolutionScope_TypeRef, Index: 1}},
	)
	m.Tables.MemberRef = customAttributeTestTable(MemberRef{Name: customAttributeTestString(".ctor"), Signature: []byte{0x20, 1, 1, 0x11, 13}})
	d := NewCustomAttributeDecoder(m)
	assembly, _ := m.Tables.AssemblyRef.At(1)
	want := &EnumReferenceMetadata{
		Handle:    CodedIndex[TypeDefOrRefOrSpec]{Tag: TypeDefOrRefOrSpec_TypeRef, Index: 2},
		Namespace: "Test+Space", Name: "Mode&", DeclaringTypes: []string{"Outer,+Part", "Inner[]"}, Assembly: &assembly,
	}
	if got, err := d.typeReference(want.Handle, 0); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("nested type reference = (%+v, %v); want (%+v, nil)", got, err, want)
	}
	d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
		if ref.SerializedName != "" || !reflect.DeepEqual(ref.Metadata, want) {
			t.Fatalf("nested reference was escaped or flattened: %+v", ref)
		}
		return ElementType_I2, nil
	}
	if _, err := d.Decode(CustomAttribute{Type: CodedIndex[CustomAttributeType]{Tag: CustomAttributeType_MemberRef}, Value: customAttributeTestValue([]byte{1, 0})}); err != nil {
		t.Fatal(err)
	}
	outer, err := d.typeReference(CodedIndex[TypeDefOrRefOrSpec]{Tag: TypeDefOrRefOrSpec_TypeRef}, 0)
	if err != nil || len(outer.DeclaringTypes) != 0 {
		t.Fatalf("building nested references changed the enclosing type: %+v, %v", outer, err)
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
