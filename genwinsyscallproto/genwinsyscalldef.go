// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Package genwinsyscallproto generates Windows syscall function prototypes ("//sys ..." comments)
// using given win32metadata information parsed by go-winmd as specified by ECMA-335.
package genwinsyscallproto

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/microsoft/go-winmd"
	"github.com/microsoft/go-winmd/coded"
	"github.com/microsoft/go-winmd/flags"
)

// Context stores data about syscall generation so far, and to improve generation performance. It
// keeps track of the list of used typedefs that may need to also be defined in generated Go code.
// For performance, indexing the data helps with (e.g.) traversing one-way pointers backwards rather
// than scanning the entire table each time.
type Context struct {
	Metadata *winmd.Metadata

	// UsedTypeRefs	is the set of all TypeRefs used by syscall function parameters. Go
	// definitions of these refs also need to be generated to make the syscalls usable.
	UsedTypeRefs map[winmd.Index]struct{}

	// AllTypeDefs is the set of all TypeDefs in the winmd file indexed by namespace + "::" + name.
	AllTypeDefs map[string]winmd.Index
	// MethodDefImplMap maps MethodDef -> ImplMap with matching MemberForwarded index.
	MethodDefImplMap map[winmd.Index]*winmd.ImplMap
	// FieldConstant maps Field -> Constant with the field as its parent.
	FieldConstant map[winmd.Index]*winmd.Constant
}

// NewContext creates a Context. Performs some pre-processing to improve generation performance.
func NewContext(f *winmd.Metadata) (*Context, error) {
	l := &Context{
		Metadata: f,

		UsedTypeRefs: make(map[winmd.Index]struct{}),

		AllTypeDefs:      make(map[string]winmd.Index, f.Tables.TypeDef.Len),
		MethodDefImplMap: make(map[winmd.Index]*winmd.ImplMap),
		FieldConstant:    make(map[winmd.Index]*winmd.Constant),
	}
	// Scan some tables to create lookups for later.
	for i := uint32(0); i < f.Tables.TypeDef.Len; i++ {
		r, err := f.Tables.TypeDef.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		// TODO: Determine if disambiguating types in winmd files with multiple modules is necessary.
		l.AllTypeDefs[r.Namespace.String()+"::"+r.Name.String()] = winmd.Index(i)
	}
	for i := uint32(0); i < f.Tables.ImplMap.Len; i++ {
		im, err := f.Tables.ImplMap.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		if im.MemberForwarded.Tag != coded.MemberForwarded_MethodDef {
			continue
		}
		if existing, ok := l.MethodDefImplMap[im.MemberForwarded.Index]; ok {
			return nil, fmt.Errorf(
				"multiple ImplMap rows found pointing at MethodDef %v: found %v; already found %v",
				im.MemberForwarded.Index,
				i,
				existing)
		}
		l.MethodDefImplMap[im.MemberForwarded.Index] = im
	}
	for i := uint32(0); i < f.Tables.Constant.Len; i++ {
		c, err := f.Tables.Constant.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		if c.Parent.Tag != coded.HasConstant_Field {
			continue
		}
		if existing, ok := l.FieldConstant[c.Parent.Index]; ok {
			return nil, fmt.Errorf(
				"multiple Constant rows found pointing at Field %v: found %v; already found %v",
				c.Parent.Index,
				i,
				existing)
		}
		l.FieldConstant[c.Parent.Index] = c
	}
	return l, nil
}

// WriteMethod writes to w the signature for "method" in x/sys/windows/mkwinsyscall format.
// Uses the parsed metadata to determine the meaning of data inside the given method.
//
// goName defines what the name of the method will be in Go after mkwinsyscall is applied. It may
// be different from the method's actual name.
//
// moduleName should be the name of the module that contains the syscall, or empty string if none is
// required. (mkwinsyscall has defaults that may be acceptable.) It is recommended to read the
// DllImport pseudo-attribute (§II.21.2.1) to determine this value.
func (c *Context) WriteMethod(w io.StringWriter, method *winmd.MethodDef, moduleName, goName string) error {
	w.WriteString("//sys\t")
	w.WriteString(goName)
	w.WriteString("(")

	sig, err := c.Metadata.MethodDefSignature(method.Signature)
	if err != nil {
		return err
	}

	for paramRowIndex := method.ParamList.Start; paramRowIndex < method.ParamList.End; paramRowIndex++ {
		param, err := c.Metadata.Tables.Param.Record(paramRowIndex)
		if err != nil {
			return fmt.Errorf("failed to read param row %v defined by method %v: %w", paramRowIndex, method.Name, err)
		}
		if param.Sequence == 0 {
			// Sequence 0 is ".param". A row with Sequence 0 is often included before other rows,
			// but Sequence 0 doesn't contain any data in cases seen so far.
			if param.Flags == 0 && param.Name.String() == "" {
				continue
			}
			// TODO: Add support for assemblies that do have data in Sequence 0.
			return fmt.Errorf("unsupported method: expected param row with sequence 0 to be empty, but: %#v", param)
		}

		// Param data is in two places: a param table row and the signature. Index into the
		// signature param slice by converting the 1-based sequence value. Note that we assume
		// ascending Sequence values for proper formatting: this is stated to be true in the
		// informational section of §II.22.33, point 4.
		// TODO: Handle gaps in Sequence values? §II.22.33 information point 5.
		// Gaps have not been seen in Windows.Win32.winmd. The meaning of a gap is not yet known.
		i := param.Sequence - 1
		if i > 0 {
			w.WriteString(", ")
		}
		c.writeParam(w, param)
		w.WriteString(" ")

		if int(i) >= len(sig.Param) {
			return fmt.Errorf("param record Sequence value %v is out of range of parsed signature params, length %v", i, len(sig.Param))
		}
		if err := c.writeType(w, &sig.Param[i].Type); err != nil {
			return fmt.Errorf("failed to interpret type of param %v of method %v: %w", i, method.Name, err)
		}
	}
	w.WriteString(")")

	// Write return value, if one exists.
	if sig.RetType.Kind != winmd.SigRetTypeKind_Void {
		w.WriteString(" (")
		if err := c.writeType(w, &sig.RetType.Type); err != nil {
			return err
		}
		w.WriteString(")")
	}

	// Write syscall name and module if non-defaults are needed.
	if goName != method.Name.String() || moduleName != "" {
		w.WriteString(" = ")
		if moduleName != "" {
			w.WriteString(moduleName)
			w.WriteString(".")
		}
		w.WriteString(method.Name.String())
	}
	return nil
}

func (c *Context) writeParam(w io.StringWriter, p *winmd.Param) {
	// TODO: Use p.Flags to determine how to translate each param to a Go type?
	w.WriteString(p.Name.String())
}

func (c *Context) writeType(b io.StringWriter, p *winmd.SigType) error {
	return c.writeTypeCore(b, p, nil)
}

func (c *Context) writeTypeCore(b io.StringWriter, p *winmd.SigType, visited map[*winmd.SigType]struct{}) error {
	if visited != nil {
		if _, ok := visited[p]; ok {
			return fmt.Errorf("cycle detected in type definition: already visited %v", p)
		}
	}

	// Special case: *void is unsafe.Pointer
	if p.Kind == flags.ElementType_PTR {
		if t, ok := p.Value.(winmd.SigType); ok {
			if t.Kind == flags.ElementType_VOID {
				b.WriteString("unsafe.Pointer")
				return nil
			}
		}
	}

	switch p.Kind {
	// Translate ECMA-335 primitive types to Go types.
	case flags.ElementType_BOOLEAN:
		b.WriteString("bool")
	case flags.ElementType_CHAR:
		// TODO: Is there a better representation of CHAR?
		b.WriteString("uint16")
	case flags.ElementType_I1:
		b.WriteString("int8")
	case flags.ElementType_U1:
		b.WriteString("uint8")
	case flags.ElementType_I2:
		b.WriteString("int16")
	case flags.ElementType_U2:
		b.WriteString("uint16")
	case flags.ElementType_I4:
		b.WriteString("int32")
	case flags.ElementType_U4:
		b.WriteString("uint32")
	case flags.ElementType_I8:
		b.WriteString("int64")
	case flags.ElementType_U8:
		b.WriteString("uint64")
	case flags.ElementType_R4:
		b.WriteString("float32")
	case flags.ElementType_R8:
		b.WriteString("float64")

	// ECMA-335 distinguishes uintptr and intptr, Go only has uintptr used in both cases.
	case flags.ElementType_I, flags.ElementType_U:
		b.WriteString("uintptr")

	case flags.ElementType_VOID:
		// We catch "*void" with a special case above. We should never see simply VOID.
		return errors.New("unexpected primitive type: VOID")

	case flags.ElementType_OBJECT:
		b.WriteString("any")

	// If this is not a simple value type, there will be p.Value. Handle all those cases here.
	default:
		if p.Kind == flags.ElementType_PTR {
			b.WriteString("*")
		}
		if visited == nil {
			visited = make(map[*winmd.SigType]struct{})
		}
		visited[p] = struct{}{}
		return c.writeTypeValue(b, p.Value, visited)
	}
	return nil
}

func (c *Context) writeTypeValue(b io.StringWriter, value any, visited map[*winmd.SigType]struct{}) error {
	switch v := value.(type) {
	case winmd.CodedIndex:
		switch v.Tag {
		case coded.TypeDefOrRefOrSpec_TypeDef:
			record, err := c.Metadata.Tables.TypeDef.Record(v.Index)
			if err != nil {
				return err
			}
			b.WriteString(record.Name.String())

		case coded.TypeDefOrRefOrSpec_TypeRef:
			record, err := c.Metadata.Tables.TypeRef.Record(v.Index)
			if err != nil {
				return err
			}
			c.UsedTypeRefs[v.Index] = struct{}{}
			b.WriteString(record.Name.String())

		default:
			return fmt.Errorf("unexpected coded index value: %#v", v)
		}

	// Types can nest. A pointer to another type is a very common case.
	case winmd.SigType:
		return c.writeTypeCore(b, &v, visited)
	case winmd.SigArray:
		b.WriteString("[]")
		return c.writeTypeCore(b, &v.Type, visited)

	default:
		return fmt.Errorf("unexpected type value: %#v", value)
	}
	return nil
}

func (c *Context) WriteTypeDef(b io.StringWriter, def *winmd.TypeDef) error {
	switch def.Extends.Tag {
	case coded.TypeDefOrRef_TypeRef:
		// TODO: Keep track of this index rather than looking it up for each enum type.
		record, err := c.Metadata.Tables.TypeRef.Record(def.Extends.Index)
		if err != nil {
			return err
		}
		if record.Namespace.String() == "System" && record.Name.String() == "Enum" {
			return c.writeTypeDefEnum(b, def)
		}
		// TODO: Detect NativeTypedefAttribute and generate a simple type.
		return c.writeTypeDefStruct(b, def)
	default:
		return fmt.Errorf("unexpected type extends coded index: %#v", def.Extends)
	}
}

func (c *Context) writeTypeDefEnum(b io.StringWriter, def *winmd.TypeDef) error {
	b.WriteString("type ")
	b.WriteString(def.Name.String())

	// Per §I.8.5.2 CLS Rule 7, the underlying type is the type of the field "__value". Find it.
	var underlyingType *winmd.SigType

	type nameValuePair struct {
		Name  string
		Value string
	}
	// The number of enum members is the total number of fields minus the special "value__".
	members := make([]nameValuePair, 0, def.FieldList.End-def.FieldList.Start-1)

	maxNameLen := 0

	for i := def.FieldList.Start; i < def.FieldList.End; i++ {
		fd, err := c.Metadata.Tables.Field.Record(i)
		if err != nil {
			return err
		}
		signature, err := c.Metadata.FieldSignature(fd.Signature)
		if err != nil {
			return err
		}
		if fd.Name.String() == "value__" {
			underlyingType = &signature.Type
			continue
		}

		var hex string
		if constant, ok := c.FieldConstant[i]; ok {
			// Read the value as a hex string. §II.22.9 informative text 1 and 2 restrict the
			// possible types further than what we can handle here, but we handle all simple integer
			// types for simplicity and in case this code needs to be moved and reused elsewhere.
			switch constant.Type {
			case flags.ElementType_I1:
				hex = strconv.FormatInt(int64(int8(constant.Value[0])), 16)
			case flags.ElementType_I2:
				hex = strconv.FormatInt(int64(int16(binary.LittleEndian.Uint16(constant.Value))), 16)
			case flags.ElementType_I4:
				hex = strconv.FormatInt(int64(int32(binary.LittleEndian.Uint32(constant.Value))), 16)
			case flags.ElementType_I8:
				hex = strconv.FormatInt(int64(binary.LittleEndian.Uint64(constant.Value)), 16)
			case flags.ElementType_U1:
				hex = strconv.FormatUint(uint64(constant.Value[0]), 16)
			case flags.ElementType_U2:
				hex = strconv.FormatUint(uint64(binary.LittleEndian.Uint16(constant.Value)), 16)
			case flags.ElementType_U4:
				hex = strconv.FormatUint(uint64(binary.LittleEndian.Uint32(constant.Value)), 16)
			case flags.ElementType_U8:
				hex = strconv.FormatUint(binary.LittleEndian.Uint64(constant.Value), 16)
			default:
				return fmt.Errorf("enum member has unexpected type: %v, field %v", constant.Type, fd.Name)
			}
		} else {
			return fmt.Errorf("unable to find default value for field %v", fd.Name)
		}

		p := nameValuePair{fd.Name.String(), hex}
		members = append(members, p)
		if len(p.Name) > maxNameLen {
			maxNameLen = len(p.Name)
		}
	}
	if underlyingType == nil {
		return errors.New("failed to find underlying type for enum")
	}

	b.WriteString(" ")
	if err := c.writeType(b, underlyingType); err != nil {
		return err
	}
	b.WriteString("\n\nconst (\n")
	for _, pair := range members {
		b.WriteString("\t")
		// Add enum name prefix to generated name if the original member name doesn't already have
		// the prefix. This may be necessary to avoid collisions, and also makes the API easier to
		// find in Go via autocomplete.
		// TODO: Be a bit smarter? E.g. accept BCRYPT_CIPHER_OPERATION for BCRYPT_OPERATION enum without changing it.
		if !strings.HasPrefix(pair.Name, def.Name.String()) {
			b.WriteString(def.Name.String())
			b.WriteString("_")
		}
		b.WriteString(pair.Name)
		for i := 0; i < maxNameLen-len(pair.Name)+1; i++ {
			b.WriteString(" ")
		}
		b.WriteString(def.Name.String())
		b.WriteString(" = 0x")
		b.WriteString(pair.Value)
		b.WriteString("\n")
	}
	b.WriteString(")\n")
	return nil
}

func (c *Context) writeTypeDefStruct(b io.StringWriter, def *winmd.TypeDef) error {
	b.WriteString("type ")
	b.WriteString(def.Name.String())
	b.WriteString(" struct {\n")

	type nameTypePair struct {
		Name string
		Type *winmd.SigType
	}
	fields := make([]nameTypePair, 0, def.FieldList.End-def.FieldList.Start)
	maxNameLen := 0
	for i := def.FieldList.Start; i < def.FieldList.End; i++ {
		fd, err := c.Metadata.Tables.Field.Record(i)
		if err != nil {
			return err
		}
		signature, err := c.Metadata.FieldSignature(fd.Signature)
		if err != nil {
			return err
		}
		// Don't write anything yet, just save the necessary data. We need to wait until we know the
		// longest field name to write the correct number of spaces.
		p := nameTypePair{fd.Name.String(), &signature.Type}
		fields = append(fields, p)

		if len(p.Name) > maxNameLen {
			maxNameLen = len(p.Name)
		}
	}
	for _, pair := range fields {
		b.WriteString("\t")
		b.WriteString(pair.Name)
		for i := 0; i < maxNameLen-len(pair.Name)+1; i++ {
			b.WriteString(" ")
		}
		if err := c.writeType(b, pair.Type); err != nil {
			return err
		}
		b.WriteString("\n")
	}
	b.WriteString("}\n")
	return nil
}

func (c *Context) WriteUsedTypeDefs(b io.StringWriter) error {
	var usedTypeDefIndices []winmd.Index
	for i := range c.UsedTypeRefs {
		ref, err := c.Metadata.Tables.TypeRef.Record(i)
		if err != nil {
			return err
		}
		switch ref.ResolutionScope.Tag {
		case coded.ResolutionScope_Module:
			// This TypeRef refers to the current module. We don't need to look at the Index.
		default:
			log.Printf("Skipping %v::%v: not defined in current module\n", ref.Namespace, ref.Name)
			continue
		}
		if def, ok := c.AllTypeDefs[ref.Namespace.String()+"::"+ref.Name.String()]; ok {
			usedTypeDefIndices = append(usedTypeDefIndices, def)
		} else {
			return fmt.Errorf("TypeRef %v::%v has Module resolution scope, but is not found in the module", ref.Namespace, ref.Name)
		}
	}
	sort.Slice(usedTypeDefIndices, func(i, j int) bool {
		return usedTypeDefIndices[i] < usedTypeDefIndices[j]
	})
	if len(usedTypeDefIndices) == 0 {
		return nil
	}
	b.WriteString("\n\n// Types used in generated APIs\n\n")
	for _, index := range usedTypeDefIndices {
		record, err := c.Metadata.Tables.TypeDef.Record(index)
		if err != nil {
			return err
		}
		if err := c.WriteTypeDef(b, record); err != nil {
			return err
		}
		b.WriteString("\n")
	}
	return nil
}
