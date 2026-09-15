// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// UnresolvedEnumError reports a custom-attribute enum whose underlying integer
// type could not be resolved. Use [errors.As] to distinguish this condition from
// malformed metadata and report or skip the attribute without stopping other work.
type UnresolvedEnumError struct {
	Name string // Enum type name, possibly assembly-qualified.
}

func (e *UnresolvedEnumError) Error() string {
	return fmt.Sprintf("cannot resolve custom attribute enum %q", e.Name)
}

// CustomAttributeDecoder resolves constructor signatures and enum types before
// decoding custom-attribute values. It caches metadata lookups for repeated use.
// A decoder must not be used concurrently, and its metadata and ResolveEnum
// callback must not be changed after decoding starts.
type CustomAttributeDecoder struct {
	// ResolveEnum supplies the underlying integer type of enums not resolved
	// in the current metadata. Names can be assembly-qualified and use backslash
	// escapes for special characters in type identifiers. Return an
	// *UnresolvedEnumError when an enum's underlying type is unavailable; other
	// errors are propagated without being classified as unresolved enums.
	// A nil callback reports unresolved enums with *UnresolvedEnumError.
	// The decoder never assumes an enum's size.
	// Enum arguments support ElementType_BOOLEAN, ElementType_CHAR, and the
	// fixed-width integer types ElementType_I1 through ElementType_U8. Native-sized
	// enum arguments are unsupported because their serialized width is unknown.
	ResolveEnum func(name string) (ElementType, error)

	metadata     *Metadata
	constructors map[CodedIndex[CustomAttributeType]][]CustomAttributeArgumentType
	enums        map[string]ElementType
	typeNames    map[Index]string
	typesByName  map[string][]Index
	typesIndexed bool
	typesError   error
}

// NewCustomAttributeDecoder creates a reusable decoder for attributes in m.
// It resolves enums declared in m, including nested enums. Unqualified names
// and names qualified with m's simple assembly name can be resolved locally;
// other assembly-qualified names are passed to ResolveEnum.
func NewCustomAttributeDecoder(m *Metadata) *CustomAttributeDecoder {
	return &CustomAttributeDecoder{
		metadata:     m,
		constructors: make(map[CodedIndex[CustomAttributeType]][]CustomAttributeArgumentType),
		enums:        make(map[string]ElementType),
	}
}

// Decode reads a's MethodDef or MemberRef constructor signature and decodes its
// value. Constructor signatures must be non-generic instance constructors with
// a void return type and valid custom-attribute parameter types. External enums
// require ResolveEnum; unsupported TypeSpec parameters are reported as errors.
//
// An unresolved enum stops decoding and returns a zero CustomAttributeValue.
// Use [errors.As] to detect *UnresolvedEnumError in the returned error. The
// original a.Value is unchanged and can be retained for later decoding.
func (d *CustomAttributeDecoder) Decode(a CustomAttribute) (CustomAttributeValue, error) {
	if d == nil || d.metadata == nil || d.metadata.Tables == nil {
		return CustomAttributeValue{}, errors.New("missing custom attribute metadata")
	}
	types, err := d.constructorTypes(a.Type)
	if err != nil {
		return CustomAttributeValue{}, fmt.Errorf("custom attribute constructor: %w", err)
	}
	return decodeCustomAttributeValue(a.Value, types, d.resolveEnum)
}

func (d *CustomAttributeDecoder) constructorTypes(index CodedIndex[CustomAttributeType]) ([]CustomAttributeArgumentType, error) {
	if types, ok := d.constructors[index]; ok {
		return types, nil
	}
	var name String
	var signature []byte
	switch index.Tag {
	case CustomAttributeType_MethodDef:
		ctor, err := d.metadata.Tables.MethodDef.At(index.Index)
		if err != nil {
			return nil, err
		}
		name, signature = ctor.Name, ctor.Signature
	case CustomAttributeType_MemberRef:
		ctor, err := d.metadata.Tables.MemberRef.At(index.Index)
		if err != nil {
			return nil, err
		}
		name, signature = ctor.Name, ctor.Signature
	default:
		return nil, fmt.Errorf("invalid constructor index %v", index)
	}
	if name.String() != ".ctor" {
		return nil, fmt.Errorf("%q is not an instance constructor", name)
	}
	// Read the subset of signatures allowed for attribute constructors. Using
	// a bounded reader also rejects unsupported type shapes before allocating
	// arrays or recursively decoding arbitrary signature types.
	r := customAttributeReader{data: signature}
	if header := r.byte(); r.err != nil {
		return nil, r.err
	} else if header != sigAbbrev_HASTHIS|sigKind_DEFAULT {
		return nil, errors.New("invalid custom attribute constructor calling convention")
	}
	count := r.compressedUint32()
	if ret := d.signatureKind(&r); r.err != nil {
		return nil, r.err
	} else if ret != ElementType_VOID {
		return nil, errors.New("custom attribute constructor must return void")
	}
	if uint64(count) > uint64(len(r.data)) {
		return nil, io.ErrUnexpectedEOF
	}
	var types []CustomAttributeArgumentType
	for i := range count {
		typ := d.signatureType(&r, 0)
		if r.err != nil {
			return nil, fmt.Errorf("parameter %d: %w", i, r.err)
		}
		types = append(types, typ)
	}
	if len(r.data) != 0 {
		return nil, errors.New("trailing custom attribute constructor signature data")
	}
	d.constructors[index] = types
	return types, nil
}

func (d *CustomAttributeDecoder) signatureKind(r *customAttributeReader) ElementType {
	for {
		kind := ElementType(r.byte())
		if r.err != nil || (kind != ElementType_CMOD_OPT && kind != ElementType_CMOD_REQD) {
			return kind
		}
		d.signatureHandle(r)
	}
}

func (d *CustomAttributeDecoder) signatureHandle(r *customAttributeReader) CodedIndex[TypeDefOrRefOrSpec] {
	code := r.compressedUint32()
	if r.err != nil {
		return CodedIndex[TypeDefOrRefOrSpec]{}
	}
	index, err := parseCoded[TypeDefOrRefOrSpec](code)
	if err != nil {
		r.err = err
		return index
	}
	var count uint32
	switch index.Tag {
	case TypeDefOrRefOrSpec_TypeDef:
		count = d.metadata.Tables.TypeDef.Len()
	case TypeDefOrRefOrSpec_TypeRef:
		count = d.metadata.Tables.TypeRef.Len()
	case TypeDefOrRefOrSpec_TypeSpec:
		count = d.metadata.Tables.TypeSpec.Len()
	default:
		r.err = errors.New("invalid custom attribute signature type handle")
	}
	if r.err == nil && uint32(index.Index) >= count {
		r.err = errors.New("custom attribute signature type handle is out of range")
	}
	return index
}

func (d *CustomAttributeDecoder) signatureType(r *customAttributeReader, depth int) CustomAttributeArgumentType {
	if !r.checkDepth(depth) {
		return CustomAttributeArgumentType{}
	}
	typ := CustomAttributeArgumentType{Kind: d.signatureKind(r)}
	switch typ.Kind {
	case ElementType_BOOLEAN, ElementType_CHAR,
		ElementType_I1, ElementType_U1, ElementType_I2, ElementType_U2,
		ElementType_I4, ElementType_U4, ElementType_I8, ElementType_U8,
		ElementType_R4, ElementType_R8, ElementType_STRING:
	case ElementType_OBJECT:
		typ.Kind = ElementType_BOXED_OBJECT
	case ElementType_SZARRAY:
		element := d.signatureType(r, depth+1)
		typ.Element = &element
	case ElementType_CLASS, ElementType_VALUETYPE:
		handle := d.signatureHandle(r)
		if r.err != nil {
			break
		}
		name, err := d.typeName(handle, 0)
		if err != nil {
			r.err = err
			break
		}
		if typ.Kind == ElementType_CLASS {
			unqualified, _, _ := splitTypeName(name)
			if unqualified != "System.Type" {
				r.err = fmt.Errorf("unsupported custom attribute class %q", name)
				break
			}
			typ.Kind = ElementType_TYPE
		} else {
			typ.Kind = ElementType_ENUM
			typ.EnumName = name
			typ.EnumUnderlyingType, r.err = d.resolveEnum(name)
		}
	default:
		if r.err == nil {
			r.err = fmt.Errorf("unsupported custom attribute signature type %v", typ.Kind)
		}
	}
	return r.normalizeType(typ, depth)
}

func (d *CustomAttributeDecoder) indexTypes() error {
	if d.typesIndexed {
		return d.typesError
	}
	d.typesIndexed = true
	d.typesError = d.buildTypeIndex()
	return d.typesError
}

func (d *CustomAttributeDecoder) buildTypeIndex() error {
	parents := make(map[Index]Index)
	for index := range d.metadata.Tables.NestedClass.Indices() {
		nested, err := d.metadata.Tables.NestedClass.At(index)
		if err != nil {
			return err
		}
		if _, ok := parents[nested.NestedClass]; ok {
			return fmt.Errorf("duplicate enclosing type for TypeDef %d", nested.NestedClass)
		}
		parents[nested.NestedClass] = nested.EnclosingClass
	}
	d.typeNames = make(map[Index]string)
	d.typesByName = make(map[string][]Index)
	var nameOf func(Index, int) (string, error)
	nameOf = func(index Index, depth int) (string, error) {
		if name, ok := d.typeNames[index]; ok {
			return name, nil
		}
		if depth >= maxCustomAttributeDepth {
			return "", errors.New("cyclic or excessively nested TypeDef")
		}
		typ, err := d.metadata.Tables.TypeDef.At(index)
		if err != nil {
			return "", err
		}
		name := typeNameEscaper.Replace(typ.Name.String())
		if parent, ok := parents[index]; ok {
			parentName, err := nameOf(parent, depth+1)
			if err != nil {
				return "", err
			}
			name = parentName + "+" + name
		} else if namespace := typ.Namespace.String(); namespace != "" {
			name = typeNameEscaper.Replace(namespace) + "." + name
		}
		d.typeNames[index] = name
		return name, nil
	}
	for index := range d.metadata.Tables.TypeDef.Indices() {
		name, err := nameOf(index, 0)
		if err != nil {
			return err
		}
		d.typesByName[name] = append(d.typesByName[name], index)
	}
	return nil
}

func (d *CustomAttributeDecoder) typeName(index CodedIndex[TypeDefOrRefOrSpec], depth int) (string, error) {
	if depth >= maxCustomAttributeDepth {
		return "", errors.New("cyclic or excessively nested TypeRef")
	}
	switch index.Tag {
	case TypeDefOrRefOrSpec_TypeDef:
		if err := d.indexTypes(); err != nil {
			return "", err
		}
		name, ok := d.typeNames[index.Index]
		if !ok {
			return "", errors.New("custom attribute TypeDef is out of range")
		}
		return name, nil
	case TypeDefOrRefOrSpec_TypeRef:
		ref, err := d.metadata.Tables.TypeRef.At(index.Index)
		if err != nil {
			return "", err
		}
		name := typeNameEscaper.Replace(ref.Name.String())
		if ref.ResolutionScope.Tag == ResolutionScope_TypeRef {
			parent, err := d.typeName(CodedIndex[TypeDefOrRefOrSpec]{Tag: TypeDefOrRefOrSpec_TypeRef, Index: ref.ResolutionScope.Index}, depth+1)
			if err != nil {
				return "", err
			}
			parentName, assembly, qualified := splitTypeName(parent)
			name = parentName + "+" + name
			if qualified {
				name += "," + assembly
			}
			return name, nil
		}
		if namespace := ref.Namespace.String(); namespace != "" {
			name = typeNameEscaper.Replace(namespace) + "." + name
		}
		switch ref.ResolutionScope.Tag {
		case ResolutionScope_Module:
			if _, err := d.metadata.Tables.Module.At(ref.ResolutionScope.Index); err != nil {
				return "", err
			}
			return name, nil
		case ResolutionScope_AssemblyRef:
			assembly, err := d.metadata.Tables.AssemblyRef.At(ref.ResolutionScope.Index)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("%s, %s, Version=%d.%d.%d.%d", name, assembly.Name, assembly.MajorVersion,
				assembly.MinorVersion, assembly.BuildNumber, assembly.RevisionNumber), nil
		default:
			return "", fmt.Errorf("unsupported resolution scope for custom attribute type %q", name)
		}
	default:
		return "", errors.New("unsupported custom attribute TypeSpec or null type")
	}
}

func (d *CustomAttributeDecoder) resolveEnum(name string) (ElementType, error) {
	if typ, ok := d.enums[name]; ok {
		return typ, nil
	}
	localName, assemblyName, qualified := splitTypeName(name)
	local := !qualified
	if qualified && !strings.Contains(assemblyName, ",") && d.metadata.Tables.Assembly.Len() == 1 {
		assembly, err := d.metadata.Tables.Assembly.At(0)
		if err != nil {
			return 0, err
		}
		local = strings.EqualFold(strings.TrimSpace(assemblyName), assembly.Name.String())
	}
	if local {
		if err := d.indexTypes(); err != nil {
			return 0, err
		}
		var underlying ElementType
		for _, index := range d.typesByName[localName] {
			typ, err := d.metadata.EnumUnderlyingType(index)
			if err != nil {
				return 0, err
			}
			if underlying != 0 && underlying != typ {
				return 0, fmt.Errorf("ambiguous underlying type for enum %q", name)
			}
			underlying = typ
		}
		if underlying != 0 {
			d.enums[name] = underlying
			return underlying, nil
		}
	}
	if d.ResolveEnum == nil {
		return 0, &UnresolvedEnumError{Name: name}
	}
	typ, err := d.ResolveEnum(name)
	if err != nil {
		return 0, err
	}
	if !isCustomAttributeEnumType(typ) {
		return 0, fmt.Errorf("unsupported underlying type %v for enum %q", typ, name)
	}
	d.enums[name] = typ
	return typ, nil
}

// typeNameEscaper converts literal metadata identifiers into serialized type-name
// components for lookup. Escaping distinguishes literal '+' and ',' characters
// from nesting and assembly separators. The other escapes preserve literal
// identifiers too; they do not imply support for pointer, array, or generic types.
// Names read from attribute blobs are left unchanged.
// See https://learn.microsoft.com/dotnet/fundamentals/reflection/specifying-fully-qualified-type-names#specify-special-characters.
var typeNameEscaper = strings.NewReplacer(
	`\`, `\\`,
	`+`, `\+`,
	`,`, `\,`,
	`&`, `\&`,
	`*`, `\*`,
	`[`, `\[`,
	`]`, `\]`,
)

// splitTypeName separates an optional assembly qualifier at the first unescaped
// comma. Type-name whitespace is significant and is preserved.
func splitTypeName(name string) (typeName, assembly string, qualified bool) {
	for i := 0; i < len(name); i++ {
		switch name[i] {
		case '\\':
			i++
		case ',':
			return name[:i], name[i+1:], true
		}
	}
	return name, "", false
}
