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
func WriteMethod(w io.StringWriter, metadata *winmd.Metadata, method *winmd.MethodDef, moduleName, goName string) error {
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

		if int(i) >= len(method.Signature.Param) {
			return fmt.Errorf("param record Sequence value %v is out of range of parsed signature params, length %v", i, len(method.Signature.Param))
		}
		if err := writeType(w, metadata, method.Signature.Param[i].Type); err != nil {
			return fmt.Errorf("failed to interpret type of param %v of method %v: %w", i, method.Name, err)
		}
	}
	w.WriteString(")")

	// Write return value, if one exists.
	if method.Signature.RetType.Kind != winmd.RetTypeKind_Void {
		w.WriteString(" (")
		if err := writeType(w, metadata, method.Signature.RetType.Type); err != nil {
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
	// TODO: Remove these? Or use them to help determine how to translate return types for syscalls.
	// Include comments for param attributes to make them visible for diagnosis purposes.
	if p.Flags&flags.ParamAttributes_In != 0 {
		w.WriteString("/*in*/")
	}
	if p.Flags&flags.ParamAttributes_Out != 0 {
		w.WriteString("/*out*/")
	}
	if p.Flags&flags.ParamAttributes_Optional != 0 {
		w.WriteString("/*optional*/")
	}
	if p.Flags&flags.ParamAttributes_HasDefault != 0 {
		w.WriteString("/*hasdefault*/")
	}
	if p.Flags&flags.ParamAttributes_HasFieldMarshal != 0 {
		w.WriteString("/*hasfieldmarshal*/")
	}
	if p.Flags&flags.ParamAttributes_Unused != 0 {
		w.WriteString("/*unused*/")
	}
	w.WriteString(p.Name.String())
}

func writeType(b io.StringWriter, f *winmd.Metadata, p winmd.Type) error {
	// Special case: *void is unsafe.Pointer
	if p.Kind == flags.ElementType_PTR {
		if t, ok := p.Value.(winmd.Type); ok {
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

	// If this is not a simple value type, there will be p.Value. Handle all those cases here.
	default:
		if p.Kind == flags.ElementType_PTR {
			b.WriteString("*")
		}
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
	case winmd.Type:
		return writeType(b, f, v)

	default:
		return fmt.Errorf("unexpected type value: %#v", value)
	}
	return nil
}