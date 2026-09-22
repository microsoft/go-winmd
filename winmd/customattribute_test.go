// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func attributeString(s string) []byte {
	n := len(s)
	var data []byte
	switch {
	case n < 0x80:
		data = []byte{byte(n)}
	case n < 0x4000:
		data = []byte{0x80 | byte(n>>8), byte(n)}
	default:
		data = []byte{0xC0 | byte(n>>24), byte(n >> 16), byte(n >> 8), byte(n)}
	}
	return append(data, s...)
}

func attributeBlob(fixed []byte, named ...[]byte) []byte {
	data := append([]byte{1, 0}, fixed...)
	data = binary.LittleEndian.AppendUint16(data, uint16(len(named)))
	for _, argument := range named {
		data = append(data, argument...)
	}
	return data
}

func namedAttribute(kind ElementType, typ []byte, name string, value []byte) []byte {
	data := append([]byte{byte(kind)}, typ...)
	data = append(data, attributeString(name)...)
	return append(data, value...)
}

func attributeType(kind ElementType) CustomAttributeArgumentType {
	return CustomAttributeArgumentType{Kind: kind}
}

func attributeArray(kind ElementType) CustomAttributeArgumentType {
	element := attributeType(kind)
	return CustomAttributeArgumentType{Kind: ElementType_SZARRAY, Element: &element}
}

func TestDecodeCustomAttributeFixedArguments(t *testing.T) {
	for _, test := range []struct {
		name  string
		typ   CustomAttributeArgumentType
		bytes []byte
		want  any
	}{
		{"false", attributeType(ElementType_BOOLEAN), []byte{0}, false},
		{"true", attributeType(ElementType_BOOLEAN), []byte{1}, true},
		{"char", attributeType(ElementType_CHAR), []byte{0xAC, 0x20}, uint16(0x20AC)},
		{"i1", attributeType(ElementType_I1), []byte{0x80}, int8(-128)},
		{"u1", attributeType(ElementType_U1), []byte{0xFF}, uint8(255)},
		{"i2", attributeType(ElementType_I2), []byte{0xFE, 0xFF}, int16(-2)},
		{"u2", attributeType(ElementType_U2), []byte{0x34, 0x12}, uint16(0x1234)},
		{"i4", attributeType(ElementType_I4), []byte{0, 0, 0, 0x80}, int32(-0x80000000)},
		{"u4", attributeType(ElementType_U4), []byte{0xFF, 0xFF, 0xFF, 0xFF}, uint32(0xFFFFFFFF)},
		{"i8", attributeType(ElementType_I8), []byte{0, 0, 0, 0, 0, 0, 0, 0x80}, int64(-0x8000000000000000)},
		{"u8", attributeType(ElementType_U8), bytes.Repeat([]byte{0xFF}, 8), uint64(0xFFFFFFFFFFFFFFFF)},
		{"r4", attributeType(ElementType_R4), []byte{0, 0, 0xC0, 0x3F}, float32(1.5)},
		{"r8", attributeType(ElementType_R8), []byte{0, 0, 0, 0, 0, 0, 0xF8, 0xBF}, float64(-1.5)},
		{"string", attributeType(ElementType_STRING), attributeString("a\x00\u20ac"), "a\x00\u20ac"},
		{"empty-string", attributeType(ElementType_STRING), []byte{0}, ""},
		{"null-string", attributeType(ElementType_STRING), []byte{0xFF}, nil},
		{"long-string", attributeType(ElementType_STRING), attributeString(strings.Repeat("a", 0x4000)), strings.Repeat("a", 0x4000)},
		{"system-type", attributeType(ElementType_TYPE), attributeString("Example.Type, Assembly"), "Example.Type, Assembly"},
		{"null-type", attributeType(ElementType_TYPE), []byte{0xFF}, nil},
		{"enum", CustomAttributeArgumentType{Kind: ElementType_ENUM, Enum: EnumReference{SerializedName: "Example.Mode"}, EnumUnderlyingType: ElementType_I2}, []byte{0xFE, 0xFF}, int16(-2)},
		{"array", attributeArray(ElementType_I2), []byte{2, 0, 0, 0, 1, 0, 0xFE, 0xFF}, []CustomAttributeArgument{
			{Type: attributeType(ElementType_I2), Value: int16(1)},
			{Type: attributeType(ElementType_I2), Value: int16(-2)},
		}},
		{"empty-array", attributeArray(ElementType_U1), []byte{0, 0, 0, 0}, []CustomAttributeArgument{}},
		{"null-array", attributeArray(ElementType_U1), []byte{0xFF, 0xFF, 0xFF, 0xFF}, nil},
		{"boxed-int", attributeType(ElementType_BOXED_OBJECT), []byte{byte(ElementType_I4), 0x34, 0x12, 0, 0}, CustomAttributeArgument{
			Type: attributeType(ElementType_I4), Value: int32(0x1234),
		}},
		{"boxed-null-string", attributeType(ElementType_BOXED_OBJECT), []byte{byte(ElementType_STRING), 0xFF}, CustomAttributeArgument{
			Type: attributeType(ElementType_STRING),
		}},
		{"boxed-array", attributeType(ElementType_BOXED_OBJECT), []byte{byte(ElementType_SZARRAY), byte(ElementType_U1), 1, 0, 0, 0, 7}, CustomAttributeArgument{
			Type: attributeArray(ElementType_U1), Value: []CustomAttributeArgument{{Type: attributeType(ElementType_U1), Value: uint8(7)}},
		}},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := decodeCustomAttributeValue(attributeBlob(test.bytes), []CustomAttributeArgumentType{test.typ}, nil)
			if err != nil {
				t.Fatal(err)
			}
			want := CustomAttributeValue{FixedArguments: []CustomAttributeArgument{{Type: test.typ, Value: test.want}}}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("decodeCustomAttributeValue() = %#v; want %#v", got, want)
			}
		})
	}
}

func TestDecodeCustomAttributeBoxedNull(t *testing.T) {
	objectType := attributeType(ElementType_BOXED_OBJECT)
	boxedNull := CustomAttributeArgument{Type: attributeType(ElementType_STRING)}
	for _, test := range []struct {
		name  string
		data  []byte
		types []CustomAttributeArgumentType
		want  CustomAttributeValue
	}{
		{
			name:  "fixed",
			data:  attributeBlob([]byte{byte(ElementType_STRING), 0xFF, 3, 0}),
			types: []CustomAttributeArgumentType{objectType, attributeType(ElementType_I2)},
			want: CustomAttributeValue{FixedArguments: []CustomAttributeArgument{
				{Type: objectType, Value: boxedNull},
				{Type: attributeType(ElementType_I2), Value: int16(3)},
			}},
		},
		{
			name: "named",
			data: attributeBlob(nil, namedAttribute(ElementType_PROPERTY, []byte{byte(ElementType_BOXED_OBJECT)}, "Object", []byte{byte(ElementType_STRING), 0xFF})),
			want: CustomAttributeValue{NamedArguments: []CustomAttributeNamedArgument{
				{Kind: ElementType_PROPERTY, Name: "Object", CustomAttributeArgument: CustomAttributeArgument{Type: objectType, Value: boxedNull}},
			}},
		},
		{
			name:  "array",
			data:  attributeBlob([]byte{2, 0, 0, 0, byte(ElementType_STRING), 0xFF, byte(ElementType_I1), 7}),
			types: []CustomAttributeArgumentType{attributeArray(ElementType_BOXED_OBJECT)},
			want: CustomAttributeValue{FixedArguments: []CustomAttributeArgument{{
				Type: attributeArray(ElementType_BOXED_OBJECT),
				Value: []CustomAttributeArgument{
					{Type: objectType, Value: boxedNull},
					{Type: objectType, Value: CustomAttributeArgument{Type: attributeType(ElementType_I1), Value: int8(7)}},
				},
			}}},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := decodeCustomAttributeValue(test.data, test.types, nil)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("boxed null = %#v; want %#v", got, test.want)
			}
		})
	}
}

func TestDecodeCustomAttributeNamedArguments(t *testing.T) {
	enumName := "Example.Mode, Assembly"
	enumType := append([]byte{byte(ElementType_ENUM)}, attributeString(enumName)...)
	data := attributeBlob([]byte{7, 0},
		namedAttribute(ElementType_FIELD, []byte{byte(ElementType_I2)}, "Count", []byte{0xFF, 0xFF}),
		namedAttribute(ElementType_PROPERTY, []byte{byte(ElementType_I2)}, "Count", []byte{3, 0}),
		namedAttribute(ElementType_FIELD, enumType, "Mode", []byte{2, 0, 0, 0}),
		namedAttribute(ElementType_FIELD, []byte{byte(ElementType_BOXED_OBJECT)}, "Box", append(bytes.Clone(enumType), 3, 0, 0, 0)),
		namedAttribute(ElementType_PROPERTY, []byte{byte(ElementType_STRING)}, "Text", []byte{0xFF}),
		namedAttribute(ElementType_PROPERTY, []byte{byte(ElementType_TYPE)}, "Type", attributeString("Example.Type")),
		namedAttribute(ElementType_FIELD, append([]byte{byte(ElementType_SZARRAY)}, enumType...), "Modes", []byte{2, 0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0}),
	)
	var calls int
	resolveEnum := func(ref EnumReference) (ElementType, error) {
		calls++
		if ref.Metadata != nil || ref.SerializedName != enumName {
			t.Fatalf("enum reference = %+v; want serialized name %q", ref, enumName)
		}
		return ElementType_U4, nil
	}
	got, err := decodeCustomAttributeValue(data, []CustomAttributeArgumentType{attributeType(ElementType_I2)}, resolveEnum)
	if err != nil {
		t.Fatal(err)
	}
	if calls != 1 {
		t.Fatalf("enum resolved %d times; want 1", calls)
	}
	resolvedEnum := CustomAttributeArgumentType{Kind: ElementType_ENUM, Enum: EnumReference{SerializedName: enumName}, EnumUnderlyingType: ElementType_U4}
	want := CustomAttributeValue{
		FixedArguments: []CustomAttributeArgument{{Type: attributeType(ElementType_I2), Value: int16(7)}},
		NamedArguments: []CustomAttributeNamedArgument{
			{Kind: ElementType_FIELD, Name: "Count", CustomAttributeArgument: CustomAttributeArgument{Type: attributeType(ElementType_I2), Value: int16(-1)}},
			{Kind: ElementType_PROPERTY, Name: "Count", CustomAttributeArgument: CustomAttributeArgument{Type: attributeType(ElementType_I2), Value: int16(3)}},
			{Kind: ElementType_FIELD, Name: "Mode", CustomAttributeArgument: CustomAttributeArgument{Type: resolvedEnum, Value: uint32(2)}},
			{Kind: ElementType_FIELD, Name: "Box", CustomAttributeArgument: CustomAttributeArgument{Type: attributeType(ElementType_BOXED_OBJECT), Value: CustomAttributeArgument{Type: resolvedEnum, Value: uint32(3)}}},
			{Kind: ElementType_PROPERTY, Name: "Text", CustomAttributeArgument: CustomAttributeArgument{Type: attributeType(ElementType_STRING)}},
			{Kind: ElementType_PROPERTY, Name: "Type", CustomAttributeArgument: CustomAttributeArgument{Type: attributeType(ElementType_TYPE), Value: "Example.Type"}},
			{Kind: ElementType_FIELD, Name: "Modes", CustomAttributeArgument: CustomAttributeArgument{
				Type: CustomAttributeArgumentType{Kind: ElementType_SZARRAY, Element: &resolvedEnum},
				Value: []CustomAttributeArgument{
					{Type: resolvedEnum, Value: uint32(1)},
					{Type: resolvedEnum, Value: uint32(2)},
				},
			}},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("decodeCustomAttributeValue() = %#v; want %#v", got, want)
	}
}

func TestDecodeCustomAttributeArrayCount(t *testing.T) {
	// 128 uses a two-byte compressed encoding, but array counts are always uint32.
	payload := append([]byte{128, 0, 0, 0}, bytes.Repeat([]byte{7}, 128)...)
	got, err := decodeCustomAttributeValue(attributeBlob(payload), []CustomAttributeArgumentType{attributeArray(ElementType_U1)}, nil)
	if err != nil {
		t.Fatal(err)
	}
	values := got.FixedArguments[0].Value.([]CustomAttributeArgument)
	if len(values) != 128 || values[127].Value != uint8(7) {
		t.Fatalf("array = %#v; want 128 uint8(7) elements", values)
	}
}

func TestDecodeCustomAttributeTruncated(t *testing.T) {
	data := attributeBlob([]byte{7, 0}, namedAttribute(ElementType_FIELD,
		[]byte{byte(ElementType_STRING)}, "Text", attributeString("hello")))
	for n := range len(data) {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			got, err := decodeCustomAttributeValue(data[:n], []CustomAttributeArgumentType{attributeType(ElementType_I2)}, nil)
			if !errors.Is(err, io.ErrUnexpectedEOF) || !reflect.DeepEqual(got, CustomAttributeValue{}) {
				t.Fatalf("truncated input = (%#v, %v); want (zero, unexpected EOF)", got, err)
			}
		})
	}
}

func TestDecodeCustomAttributeErrors(t *testing.T) {
	for _, test := range []struct {
		name  string
		data  []byte
		types []CustomAttributeArgumentType
	}{
		{"prolog", []byte{0, 0, 0, 0}, nil},
		{"missing-named-count", []byte{1, 0}, nil},
		{"extra-data", []byte{1, 0, 0, 0, 0}, nil},
		{"invalid-named-kind", []byte{1, 0, 1, 0, 0x52}, nil},
		{"invalid-named-type", []byte{1, 0, 1, 0, 0x53, 0x0F}, nil},
		{"null-name", []byte{1, 0, 1, 0, 0x53, 0x06, 0xFF, 0, 0}, nil},
		{"empty-name", []byte{1, 0, 1, 0, 0x53, 0x06, 0, 0, 0}, nil},
		{"invalid-utf8-name", []byte{1, 0, 1, 0, 0x53, 0x06, 1, 0xFF, 0, 0}, nil},
		{"Boolean", attributeBlob([]byte{2}), []CustomAttributeArgumentType{attributeType(ElementType_BOOLEAN)}},
		{"utf8", attributeBlob([]byte{1, 0xFF}), []CustomAttributeArgumentType{attributeType(ElementType_STRING)}},
		{"unresolved-enum", attributeBlob([]byte{1, 0, 0, 0}), []CustomAttributeArgumentType{{Kind: ElementType_ENUM, Enum: EnumReference{SerializedName: "Example.E"}}}},
		{"enum-name", attributeBlob([]byte{1}), []CustomAttributeArgumentType{{Kind: ElementType_ENUM, EnumUnderlyingType: ElementType_U1}}},
		{"enum-width", attributeBlob([]byte{1}), []CustomAttributeArgumentType{{Kind: ElementType_ENUM, Enum: EnumReference{SerializedName: "Example.E"}, EnumUnderlyingType: ElementType_R4}}},
		{"ambiguous-enum-reference", attributeBlob([]byte{1}), []CustomAttributeArgumentType{{Kind: ElementType_ENUM, Enum: EnumReference{Metadata: &EnumReferenceMetadata{Name: "E"}, SerializedName: "Example.E"}, EnumUnderlyingType: ElementType_U1}}},
		{"unnamed-metadata-enum", attributeBlob([]byte{1}), []CustomAttributeArgumentType{{Kind: ElementType_ENUM, Enum: EnumReference{Metadata: &EnumReferenceMetadata{}}, EnumUnderlyingType: ElementType_U1}}},
		{"invalid-fixed-type", attributeBlob(nil), []CustomAttributeArgumentType{attributeType(ElementType_PTR)}},
		{"array-element-type", attributeBlob(nil), []CustomAttributeArgumentType{attributeType(ElementType_SZARRAY)}},
		{"array-count", attributeBlob([]byte{0, 0, 0, 0x80}), []CustomAttributeArgumentType{attributeArray(ElementType_U1)}},
		{"array-count-too-large", attributeBlob([]byte{0xFF, 0xFF, 0xFF, 0x7F}), []CustomAttributeArgumentType{attributeArray(ElementType_U1)}},
		{"array-element-truncated", attributeBlob([]byte{1, 0, 0, 0, 1}), []CustomAttributeArgumentType{attributeArray(ElementType_I8)}},
		{"boxed-object", attributeBlob([]byte{0x51}), []CustomAttributeArgumentType{attributeType(ElementType_BOXED_OBJECT)}},
		{"boxed-null-marker", attributeBlob([]byte{0xFF}), []CustomAttributeArgumentType{attributeType(ElementType_BOXED_OBJECT)}},
		{"named-boxed-null-marker", attributeBlob(nil, namedAttribute(ElementType_PROPERTY, []byte{byte(ElementType_BOXED_OBJECT)}, "Object", []byte{0xFF})), nil},
		{"array-boxed-null-marker", attributeBlob([]byte{1, 0, 0, 0, 0xFF}), []CustomAttributeArgumentType{attributeArray(ElementType_BOXED_OBJECT)}},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := decodeCustomAttributeValue(test.data, test.types, nil)
			if err == nil || !reflect.DeepEqual(got, CustomAttributeValue{}) {
				t.Fatalf("decodeCustomAttributeValue() = (%#v, %v); want (zero, error)", got, err)
			}
		})
	}
	for _, data := range [][]byte{nil, {}, {1, 0, 0, 0}} {
		if _, err := decodeCustomAttributeValue(data, nil, nil); err != nil {
			t.Errorf("empty attribute %x: %v", data, err)
		}
	}
}

func TestDecodeCustomAttributeUnresolvedEnum(t *testing.T) {
	const enumName = "Example.Mode, Assembly"
	enumType := append([]byte{byte(ElementType_ENUM)}, attributeString(enumName)...)
	for _, test := range []struct {
		name  string
		data  []byte
		types []CustomAttributeArgumentType
	}{
		{"fixed", attributeBlob([]byte{1, 0}), []CustomAttributeArgumentType{{Kind: ElementType_ENUM, Enum: EnumReference{SerializedName: enumName}}}},
		{"named", attributeBlob(nil, namedAttribute(ElementType_FIELD, enumType, "Mode", []byte{1, 0})), nil},
		{"boxed", attributeBlob(append(bytes.Clone(enumType), 1, 0)), []CustomAttributeArgumentType{attributeType(ElementType_BOXED_OBJECT)}},
	} {
		t.Run(test.name, func(t *testing.T) {
			value, err := decodeCustomAttributeValue(test.data, test.types, nil)
			var unresolved *UnresolvedEnumError
			if !errors.As(err, &unresolved) || unresolved.Reference.Metadata != nil || unresolved.Reference.SerializedName != enumName || !reflect.DeepEqual(value, CustomAttributeValue{}) {
				t.Fatalf("unresolved enum = (%#v, %v); want (zero, UnresolvedEnumError for %q)", value, err, enumName)
			}
		})
	}
}

func TestDecodeCustomAttributeEnumResolver(t *testing.T) {
	typ := CustomAttributeArgumentType{Kind: ElementType_ENUM, Enum: EnumReference{SerializedName: "Example.Mode"}}
	data := attributeBlob([]byte{0xFF, 0xFF})
	wantErr := errors.New("unresolved external enum")
	if _, err := decodeCustomAttributeValue(data, []CustomAttributeArgumentType{typ}, func(EnumReference) (ElementType, error) {
		return 0, wantErr
	}); !errors.Is(err, wantErr) {
		t.Errorf("resolver error = %v; want %v", err, wantErr)
	}
	var unresolved *UnresolvedEnumError
	if _, err := decodeCustomAttributeValue(data, []CustomAttributeArgumentType{typ}, func(EnumReference) (ElementType, error) {
		return ElementType_STRING, nil
	}); err == nil || errors.As(err, &unresolved) {
		t.Fatalf("invalid resolved enum underlying type: got %v; want an invalid-type error", err)
	}
	value, err := decodeCustomAttributeValue(data, []CustomAttributeArgumentType{typ}, func(EnumReference) (ElementType, error) {
		return ElementType_I2, nil
	})
	if err != nil || value.FixedArguments[0].Value != int16(-1) || typ.EnumUnderlyingType != 0 {
		t.Fatalf("resolved enum = (%#v, %v); input type = %#v", value, err, typ)
	}
}

func TestDecodeCustomAttributeNesting(t *testing.T) {
	var payload []byte
	for range 80 {
		// A boxed array with one boxed element can recurse without jagged array types.
		payload = append(payload, byte(ElementType_SZARRAY), byte(ElementType_BOXED_OBJECT), 1, 0, 0, 0)
	}
	payload = append(payload, byte(ElementType_I1), 0)
	if _, err := decodeCustomAttributeValue(attributeBlob(payload), []CustomAttributeArgumentType{attributeType(ElementType_BOXED_OBJECT)}, nil); err == nil || !strings.Contains(err.Error(), "nesting") {
		t.Fatalf("deeply nested attribute error = %v", err)
	}
	cycle := attributeArray(ElementType_I1)
	cycle.Element = &cycle
	if _, err := decodeCustomAttributeValue(attributeBlob(nil), []CustomAttributeArgumentType{cycle}, nil); err == nil {
		t.Fatal("cyclic array type was accepted")
	}
}

func TestCustomAttributeDecoderMetadata(t *testing.T) {
	m, err := Open("testdata/Windows.Win32.winmd")
	if err != nil {
		t.Fatal(err)
	}
	d := NewCustomAttributeDecoder(m)
	calls := make(map[CodedIndex[TypeDefOrRefOrSpec]]int)
	names := make(map[string]bool)
	d.ResolveEnum = func(ref EnumReference) (ElementType, error) {
		if ref.Metadata == nil || ref.SerializedName != "" || ref.Metadata.Assembly == nil {
			t.Fatalf("external enum lacks raw metadata: %+v", ref)
		}
		calls[ref.Metadata.Handle]++
		name := ref.Metadata.Namespace + "." + ref.Metadata.Name
		names[name] = true
		switch name {
		case "System.AttributeTargets", "System.Runtime.InteropServices.CallingConvention":
			return ElementType_I4, nil
		default:
			return 0, fmt.Errorf("unexpected external enum %q", name)
		}
	}
	var fixed, named int
	for index := range m.Tables.CustomAttribute.Indices() {
		a, err := m.Tables.CustomAttribute.At(index)
		if err != nil {
			t.Fatal(err)
		}
		value, err := d.Decode(a)
		if err != nil {
			t.Fatalf("attribute %d, constructor %v: %v", index, a.Type, err)
		}
		fixed += len(value.FixedArguments)
		named += len(value.NamedArguments)
	}
	if fixed == 0 || named == 0 || len(names) != 2 {
		t.Fatalf("fixed=%d, named=%d, external enums=%v", fixed, named, calls)
	}
	for handle, count := range calls {
		if count != 1 {
			t.Errorf("resolved %v %d times; want 1", handle, count)
		}
	}
}

func FuzzDecodeCustomAttribute(f *testing.F) {
	f.Add([]byte{1, 0, 0, 0})
	f.Add(attributeBlob([]byte{1, 0}))
	f.Add(attributeBlob([]byte{1, 0, 0, 0, byte(ElementType_STRING), 0xFF}))
	f.Fuzz(func(t *testing.T, data []byte) {
		for _, types := range [][]CustomAttributeArgumentType{nil, {attributeType(ElementType_I2)}, {attributeArray(ElementType_BOXED_OBJECT)}} {
			value, err := decodeCustomAttributeValue(data, types, func(EnumReference) (ElementType, error) { return ElementType_I4, nil })
			if err != nil && !reflect.DeepEqual(value, CustomAttributeValue{}) {
				t.Fatal("decoder returned a partial result on error")
			}
		}
	})
}
