// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Package genwinsyscallproto generates Windows syscall function prototypes ("//sys ..." comments)
// using given win32metadata information parsed by go-winmd as specified by ECMA-335.
package genwinsyscallproto

import (
	"errors"
	"fmt"
	"io"

	"github.com/microsoft/go-winmd"
	"github.com/microsoft/go-winmd/coded"
	"github.com/microsoft/go-winmd/flags"
)

// WriteMethod writes to w the signature for "method" in x/sys/windows/mkwinsyscall format.
// Uses the parsed metadata to determine the meaning of data inside the given method.
//
// goName defines what the name of the method will be in Go after mkwinsyscall is applied. It may
// be different from the method's actual name.
//
// moduleName should be the name of the module that contains the syscall, or empty string if none is
// required. (mkwinsyscall has defaults that may be acceptable.) It is recommended to read the
// DllImport pseudo-attribute (§II.21.2.1) to determine this value.
func WriteMethod(w io.StringWriter, metadata *winmd.Metadata, method *winmd.MethodDef, sig *winmd.SigMethodDef, moduleName, goName string) error {
	w.WriteString("//sys\t")
	w.WriteString(goName)
	w.WriteString("(")

	for paramRowIndex := method.ParamList.Start; paramRowIndex < method.ParamList.End; paramRowIndex++ {
		param, err := metadata.Tables.Param.Record(paramRowIndex)
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
		writeParam(w, param)
		w.WriteString(" ")

		if int(i) >= len(sig.Param) {
			return fmt.Errorf("param record Sequence value %v is out of range of parsed signature params, length %v", i, len(sig.Param))
		}
		if err := writeType(w, metadata, &sig.Param[i].Type); err != nil {
			return fmt.Errorf("failed to interpret type of param %v of method %v: %w", i, method.Name, err)
		}
	}
	w.WriteString(")")

	// Write return value, if one exists.
	if sig.RetType.Kind != winmd.SigRetTypeKind_Void {
		w.WriteString(" (")
		if err := writeType(w, metadata, &sig.RetType.Type); err != nil {
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

func writeParam(w io.StringWriter, p *winmd.Param) {
	// TODO: Use p.Flags to determine how to translate each param to a Go type?
	w.WriteString(p.Name.String())
}

func writeType(b io.StringWriter, f *winmd.Metadata, p *winmd.SigType) error {
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
		// TODO: Keep track of visited types to avoid infinite recursion.
		return writeTypeValue(b, f, p.Value)
	}
	return nil
}

func writeTypeValue(b io.StringWriter, f *winmd.Metadata, value any) error {
	switch v := value.(type) {
	case winmd.CodedIndex:
		switch v.Tag {
		case coded.TypeDefOrRefOrSpec_TypeDef:
			record, err := f.Tables.TypeDef.Record(v.Index)
			if err != nil {
				return err
			}
			b.WriteString(record.Name.String())

		case coded.TypeDefOrRefOrSpec_TypeRef:
			record, err := f.Tables.TypeRef.Record(v.Index)
			if err != nil {
				return err
			}
			b.WriteString(record.Name.String())

		default:
			return fmt.Errorf("unexpected coded index value: %#v", v)
		}

	// Types can nest. A pointer to another type is a very common case.
	case winmd.SigType:
		return writeType(b, f, &v)
	case winmd.SigArray:
		b.WriteString("[]")
		return writeType(b, f, &v.Type)

	default:
		return fmt.Errorf("unexpected type value: %#v", value)
	}
	return nil
}

func WriteTypeDef(b io.StringWriter, f *winmd.Metadata, def *winmd.TypeDef) error {
	switch def.Extends.Tag {
	case coded.TypeDefOrRef_TypeRef:
		// TODO: Keep track of this index rather than looking it up for each enum type.
		record, err := f.Tables.TypeRef.Record(def.Extends.Index)
		if err != nil {
			return err
		}
		if record.Namespace.String() == "System" && record.Name.String() == "Enum" {
			return writeTypeDefEnum(b, f, def)
		}
		// TODO: Detect NativeTypedefAttribute and generate a simple type.
		return writeTypeDefStruct(b, f, def)
	default:
		return fmt.Errorf("unexpected type extends coded index: %#v", def.Extends)
	}
}

func writeTypeDefEnum(b io.StringWriter, f *winmd.Metadata, def *winmd.TypeDef) error {
	b.WriteString("type ")
	b.WriteString(def.Name.String())

	// Per §I.8.5.2 CLS Rule 7, the underlying type is the type of the field "__value". Find it.
	var underlyingType *winmd.SigType

	type nameValuePair struct {
		Name  string
		Value any
	}
	// The number of enum members is the total number of fields minus the special "value__".
	members := make([]nameValuePair, 0, def.FieldList.End-def.FieldList.Start-1)

	maxNameLen := 0

	for i := def.FieldList.Start; i < def.FieldList.End; i++ {
		fd, err := f.Tables.Field.Record(i)
		if err != nil {
			return err
		}
		signature, err := f.FieldSignature(fd.Signature)
		if err != nil {
			return err
		}
		if fd.Name.String() == "value__" {
			underlyingType = &signature.Type
			continue
		}

		p := nameValuePair{fd.Name.String(), signature.Type}
		members = append(members, p)
		if len(p.Name) > maxNameLen {
			maxNameLen = len(p.Name)
		}
	}
	if underlyingType == nil {
		return errors.New("failed to find underlying type for enum")
	}

	b.WriteString(" ")
	if err := writeType(b, f, underlyingType); err != nil {
		return err
	}
	b.WriteString("\n\nconst (\n")
	for _, pair := range members {
		b.WriteString("\t")
		b.WriteString(pair.Name)
		for i := 0; i < maxNameLen-len(pair.Name)+1; i++ {
			b.WriteString(" ")
		}
		// TODO: Add enum name prefix to enum entries if the names don't already have the prefix? This may be necessary to avoid collisions. Also might be useful to make the API clear in Go.
		b.WriteString(def.Name.String())
		b.WriteString(" = ")
		// TODO: Look up enum member values in Constant table. Index values to avoid O(len(enums)*len(Constants))
		b.WriteString("42")
		b.WriteString("\n")
	}
	b.WriteString(")\n")
	return nil
}

func writeTypeDefStruct(b io.StringWriter, f *winmd.Metadata, def *winmd.TypeDef) error {
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
		fd, err := f.Tables.Field.Record(i)
		if err != nil {
			return err
		}
		signature, err := f.FieldSignature(fd.Signature)
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
		if err := writeType(b, f, pair.Type); err != nil {
			return err
		}
		b.WriteString("\n")
	}
	b.WriteString("}\n")
	return nil
}
