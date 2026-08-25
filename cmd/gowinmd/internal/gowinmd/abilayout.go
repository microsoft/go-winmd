package gowinmd

import (
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/microsoft/go-winmd/winmd"
)

type abiTypeLayout struct {
	abiSize  uint32
	abiAlign uint32
	goSize   uint32
	goAlign  uint32
}

type abiStructField struct {
	index    winmd.Index
	offset   uint32
	padding  uint32
	abiSize  uint32
	abiAlign uint32
}

type abiStructLayout struct {
	fields      []abiStructField
	tailPadding uint32
	typeLayout  abiTypeLayout
}

type abiLayoutFingerprint struct {
	size   uint32
	align  uint32
	fields []abiFieldFingerprint
}

type abiFieldFingerprint struct {
	offset uint32
	size   uint32
	align  uint32
}

func (c *Context) structABITypeLayout(def *resolvedDef, arch Arch, visiting map[winmd.Index]bool) (abiTypeLayout, error) {
	layout, err := c.planStructABI(def, arch, visiting)
	if err != nil {
		return abiTypeLayout{}, err
	}
	return layout.typeLayout, nil
}

func (c *Context) planStructABI(def *resolvedDef, arch Arch, visiting map[winmd.Index]bool) (abiStructLayout, error) {
	if visiting[def.Index] {
		return abiStructLayout{}, fmt.Errorf("value-type layout cycle involving %s.%s", def.Namespace, def.Name)
	}
	if visiting == nil {
		visiting = make(map[winmd.Index]bool)
	}
	visiting[def.Index] = true
	defer delete(visiting, def.Index)

	packing := defaultPacking(arch)
	classSize := uint32(0)
	if classLayout, ok := c.classLayout[def.Index]; ok {
		if classLayout.PackingSize != 0 {
			packing = uint32(classLayout.PackingSize)
		}
		classSize = classLayout.ClassSize
	}

	explicit := def.def.Flags&winmd.TypeAttributes_LayoutMask == winmd.TypeAttributes_ExplicitLayout
	usedExplicitOffsets := make(map[uint32]bool)
	abiOffset := uint32(0)
	abiEnd := uint32(0)
	abiAlign := uint32(1)
	goOffset := uint32(0)
	goAlign := uint32(1)
	var fields []abiStructField

	for fieldIndex := range def.def.FieldList.All() {
		field, err := c.Metadata.Tables.Field.At(fieldIndex)
		if err != nil {
			return abiStructLayout{}, err
		}
		if field.Flags&winmd.FieldAttributes_Static != 0 {
			continue
		}
		signature, err := c.Metadata.FieldSignature(field.Signature)
		if err != nil {
			return abiStructLayout{}, fmt.Errorf("read layout of field %s.%s: %w", def.Name, field.Name, err)
		}
		fieldLayout, err := c.sigTypeABITypeLayout(&signature.Type, arch, visiting)
		if err != nil {
			return abiStructLayout{}, fmt.Errorf("calculate layout of field %s.%s: %w", def.Name, field.Name, err)
		}
		if fieldLayout.abiSize != fieldLayout.goSize {
			return abiStructLayout{}, fmt.Errorf("field %s.%s has Windows size %d but generated Go size %d", def.Name, field.Name, fieldLayout.abiSize, fieldLayout.goSize)
		}

		fieldABIAlign := min(fieldLayout.abiAlign, packing)
		abiAlign = max(abiAlign, fieldABIAlign)
		var fieldOffset uint32
		if explicit {
			var ok bool
			fieldOffset, ok = c.fieldOffset[fieldIndex]
			if !ok {
				return abiStructLayout{}, fmt.Errorf("explicit-layout field %s.%s has no FieldLayout offset", def.Name, field.Name)
			}
		} else {
			fieldOffset = alignUp(abiOffset, fieldABIAlign)
			abiOffset = fieldOffset + fieldLayout.abiSize
		}
		abiEnd = max(abiEnd, fieldOffset+fieldLayout.abiSize)

		if explicit && usedExplicitOffsets[fieldOffset] {
			continue
		}
		usedExplicitOffsets[fieldOffset] = true

		if fieldOffset < goOffset {
			return abiStructLayout{}, fmt.Errorf("overlapping field %s.%s at offset %d cannot be represented as a Go struct", def.Name, field.Name, fieldOffset)
		}
		naturalGoOffset := alignUp(goOffset, fieldLayout.goAlign)
		if naturalGoOffset > fieldOffset {
			return abiStructLayout{}, fmt.Errorf("packed field %s.%s at offset %d cannot be represented with Go alignment %d", def.Name, field.Name, fieldOffset, fieldLayout.goAlign)
		}
		padding := uint32(0)
		if naturalGoOffset < fieldOffset {
			padding = fieldOffset - goOffset
			if alignUp(goOffset+padding, fieldLayout.goAlign) != fieldOffset {
				return abiStructLayout{}, fmt.Errorf("field %s.%s at offset %d cannot be represented with Go padding", def.Name, field.Name, fieldOffset)
			}
		}
		fields = append(fields, abiStructField{
			index:    fieldIndex,
			offset:   fieldOffset,
			padding:  padding,
			abiSize:  fieldLayout.abiSize,
			abiAlign: fieldABIAlign,
		})
		goOffset = fieldOffset + fieldLayout.goSize
		goAlign = max(goAlign, fieldLayout.goAlign)
	}

	abiSize := alignUp(max(abiEnd, classSize), abiAlign)
	naturalGoSize := alignUp(goOffset, goAlign)
	if naturalGoSize > abiSize {
		return abiStructLayout{}, fmt.Errorf("Windows size %d for %s cannot be represented by Go size %d", abiSize, def.Name, naturalGoSize)
	}
	tailPadding := uint32(0)
	if naturalGoSize < abiSize {
		tailPadding = abiSize - goOffset
		if alignUp(goOffset+tailPadding, goAlign) != abiSize {
			return abiStructLayout{}, fmt.Errorf("Windows size %d for %s cannot be represented with Go padding", abiSize, def.Name)
		}
	}

	return abiStructLayout{
		fields:      fields,
		tailPadding: tailPadding,
		typeLayout: abiTypeLayout{
			abiSize:  abiSize,
			abiAlign: abiAlign,
			goSize:   abiSize,
			goAlign:  goAlign,
		},
	}, nil
}

func (c *Context) discoverABILayoutDependencies() error {
	discovered := make(map[winmd.Index]bool)
	for {
		indices := make([]winmd.Index, 0, len(c.abiLayoutTypeDefs))
		for index := range c.abiLayoutTypeDefs {
			if !discovered[index] {
				indices = append(indices, index)
			}
		}
		if len(indices) == 0 {
			return nil
		}
		sort.Slice(indices, func(i, j int) bool { return indices[i] < indices[j] })
		for _, index := range indices {
			discovered[index] = true
			def, err := c.resolveTypeDef(index)
			if err != nil {
				return err
			}
			arches := c.TypeDefSupportedArch(index)
			if arches == ArchAll {
				arches = Arch386 | ArchAMD64 | ArchARM64
			}
			for _, arch := range arches.Unique() {
				for fieldIndex := range def.def.FieldList.All() {
					field, err := c.Metadata.Tables.Field.At(fieldIndex)
					if err != nil {
						return err
					}
					if field.Flags&winmd.FieldAttributes_Static != 0 {
						continue
					}
					signature, err := c.Metadata.FieldSignature(field.Signature)
					if err != nil {
						return err
					}
					if err := c.markSigTypeABILayoutDependencies(&signature.Type, arch); err != nil {
						return fmt.Errorf("resolve layout dependencies of %s.%s: %w", def.Name, field.Name, err)
					}
				}
			}
		}
	}
}

func (c *Context) markSigTypeABILayoutDependencies(sig *winmd.SigType, arch Arch) error {
	switch value := sig.Value.(type) {
	case winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]:
		var def *resolvedDef
		var err error
		switch value.Tag {
		case winmd.TypeDefOrRefOrSpec_TypeDef:
			def, err = c.resolveTypeDef(value.Index)
		case winmd.TypeDefOrRefOrSpec_TypeRef:
			def, err = c.resolveTypeRef(value.Index, arch)
			if errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
				return nil
			}
		default:
			return fmt.Errorf("unsupported ABI coded index tag %#v", value)
		}
		if err != nil {
			return err
		}
		c.abiLayoutTypeDefs[def.Index] = true
	case winmd.SigType:
		return c.markSigTypeABILayoutDependencies(&value, arch)
	case winmd.SigArray:
		return c.markSigTypeABILayoutDependencies(&value.Type, arch)
	}
	return nil
}

func (c *Context) typeDefABILayoutFingerprint(def *resolvedDef, arch Arch) (abiLayoutFingerprint, error) {
	if !def.Native && !extendsSystemType(c.Metadata, def, "Enum") && !extendsSystemType(c.Metadata, def, "MulticastDelegate") {
		layout, err := c.planStructABI(def, arch, nil)
		if err != nil {
			return abiLayoutFingerprint{}, err
		}
		fingerprint := abiLayoutFingerprint{
			size:   layout.typeLayout.abiSize,
			align:  layout.typeLayout.abiAlign,
			fields: make([]abiFieldFingerprint, len(layout.fields)),
		}
		for i, field := range layout.fields {
			fingerprint.fields[i] = abiFieldFingerprint{
				offset: field.offset,
				size:   field.abiSize,
				align:  field.abiAlign,
			}
		}
		return fingerprint, nil
	}
	layout, err := c.resolvedDefABITypeLayout(def, arch, nil)
	if err != nil {
		return abiLayoutFingerprint{}, err
	}
	return abiLayoutFingerprint{size: layout.abiSize, align: layout.abiAlign}, nil
}

func equalABILayout(left, right abiLayoutFingerprint) bool {
	if left.size != right.size || left.align != right.align || len(left.fields) != len(right.fields) {
		return false
	}
	for i := range left.fields {
		if left.fields[i] != right.fields[i] {
			return false
		}
	}
	return true
}

func (c *Context) sigTypeABITypeLayout(sig *winmd.SigType, arch Arch, visiting map[winmd.Index]bool) (abiTypeLayout, error) {
	if sig.Kind == winmd.ElementType_PTR || sig.Kind == winmd.ElementType_BYREF {
		return pointerABITypeLayout(arch), nil
	}

	switch sig.Kind {
	case winmd.ElementType_BOOLEAN, winmd.ElementType_I1, winmd.ElementType_U1:
		return scalarABITypeLayout(1, 1, arch), nil
	case winmd.ElementType_CHAR, winmd.ElementType_I2, winmd.ElementType_U2:
		return scalarABITypeLayout(2, 2, arch), nil
	case winmd.ElementType_I4, winmd.ElementType_U4, winmd.ElementType_R4:
		return scalarABITypeLayout(4, 4, arch), nil
	case winmd.ElementType_I8, winmd.ElementType_U8, winmd.ElementType_R8:
		return scalarABITypeLayout(8, 8, arch), nil
	case winmd.ElementType_I, winmd.ElementType_U:
		return pointerABITypeLayout(arch), nil
	case winmd.ElementType_ARRAY:
		array, ok := sig.Value.(winmd.SigArray)
		if !ok {
			return abiTypeLayout{}, fmt.Errorf("unexpected array signature value %#v", sig.Value)
		}
		element, err := c.sigTypeABITypeLayout(&array.Type, arch, visiting)
		if err != nil {
			return abiTypeLayout{}, err
		}
		count := uint64(1)
		for dimension := 0; dimension < int(array.Rank); dimension++ {
			if dimension >= len(array.Sizes) {
				return abiTypeLayout{}, errors.New("variable-length array fields are not representable in Go")
			}
			count *= uint64(array.Sizes[dimension])
		}
		if count > math.MaxUint32 || count*uint64(element.abiSize) > math.MaxUint32 {
			return abiTypeLayout{}, errors.New("array field size overflows uint32")
		}
		return abiTypeLayout{
			abiSize:  uint32(count) * element.abiSize,
			abiAlign: element.abiAlign,
			goSize:   uint32(count) * element.goSize,
			goAlign:  element.goAlign,
		}, nil
	case winmd.ElementType_OBJECT:
		return abiTypeLayout{}, errors.New("object fields are not representable as Windows ABI structs")
	}

	index, ok := sig.Value.(winmd.CodedIndex[winmd.TypeDefOrRefOrSpec])
	if !ok {
		return abiTypeLayout{}, fmt.Errorf("unsupported ABI field type %v", sig.Kind)
	}
	var def *resolvedDef
	var err error
	switch index.Tag {
	case winmd.TypeDefOrRefOrSpec_TypeDef:
		def, err = c.resolveTypeDef(index.Index)
	case winmd.TypeDefOrRefOrSpec_TypeRef:
		def, err = c.resolveTypeRef(index.Index, arch)
		if errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
			return abiTypeLayout{}, fmt.Errorf("field TypeRef is not defined in the current WinMD module: %w", err)
		}
	default:
		return abiTypeLayout{}, fmt.Errorf("unsupported ABI coded index tag %#v", index)
	}
	if err != nil {
		return abiTypeLayout{}, err
	}
	c.abiLayoutTypeDefs[def.Index] = true
	return c.resolvedDefABITypeLayout(def, arch, visiting)
}

func (c *Context) resolvedDefABITypeLayout(def *resolvedDef, arch Arch, visiting map[winmd.Index]bool) (abiTypeLayout, error) {
	if def.NeedsPointerWhenUsed() {
		return pointerABITypeLayout(arch), nil
	}
	if def.Native {
		field, err := c.Metadata.Tables.Field.At(def.def.FieldList.Start)
		if err != nil {
			return abiTypeLayout{}, err
		}
		signature, err := c.Metadata.FieldSignature(field.Signature)
		if err != nil {
			return abiTypeLayout{}, err
		}
		return c.sigTypeABITypeLayout(&signature.Type, arch, visiting)
	}
	if extendsSystemType(c.Metadata, def, "Enum") {
		for fieldIndex := range def.def.FieldList.All() {
			field, err := c.Metadata.Tables.Field.At(fieldIndex)
			if err != nil {
				return abiTypeLayout{}, err
			}
			if field.Name.String() != "value__" {
				continue
			}
			signature, err := c.Metadata.FieldSignature(field.Signature)
			if err != nil {
				return abiTypeLayout{}, err
			}
			return c.sigTypeABITypeLayout(&signature.Type, arch, visiting)
		}
		return abiTypeLayout{}, fmt.Errorf("enum %s has no value__ field", def.Name)
	}
	if extendsSystemType(c.Metadata, def, "MulticastDelegate") {
		return pointerABITypeLayout(arch), nil
	}
	return c.structABITypeLayout(def, arch, visiting)
}

func extendsSystemType(metadata *winmd.Metadata, def *resolvedDef, name string) bool {
	if def.def.Extends.Tag != winmd.TypeDefOrRef_TypeRef {
		return false
	}
	extends, err := metadata.Tables.TypeRef.At(def.def.Extends.Index)
	return err == nil && extends.Namespace.String() == "System" && extends.Name.String() == name
}

func scalarABITypeLayout(size, abiAlign uint32, arch Arch) abiTypeLayout {
	goAlign := abiAlign
	if arch == Arch386 && goAlign > 4 {
		goAlign = 4
	}
	return abiTypeLayout{abiSize: size, abiAlign: abiAlign, goSize: size, goAlign: goAlign}
}

func pointerABITypeLayout(arch Arch) abiTypeLayout {
	size := uint32(8)
	if arch == Arch386 {
		size = 4
	}
	return abiTypeLayout{abiSize: size, abiAlign: size, goSize: size, goAlign: size}
}

func defaultPacking(arch Arch) uint32 {
	if arch == ArchAMD64 {
		return 16
	}
	return 8
}

func alignUp(value, alignment uint32) uint32 {
	if alignment <= 1 {
		return value
	}
	return (value + alignment - 1) &^ (alignment - 1)
}
