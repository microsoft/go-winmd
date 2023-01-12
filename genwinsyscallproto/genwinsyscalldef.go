// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Package genwinsyscallproto generates Windows syscall function prototypes ("//sys ..." comments)
// using given win32metadata information parsed by go-winmd as specified by ECMA-335.
package genwinsyscallproto

import (
	"encoding/binary"
	"errors"
	"fmt"
	"go/token"
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

	// neededTypeDefs is the set of all TypeDefs used by syscall function parameters. Go definitions
	// of these refs also need to be generated to make the syscalls usable. During initial
	// generation of "//sys" lines, this is filled with parameter and return value types. When
	// WriteUsedTypeDefs runs, it reads the keys of this map, clears the map, then runs the type
	// generation, so new transitive dependencies are found and put into this map for the next pass.
	// It runs many passes are needed to discover and define all needed types.
	neededTypeDefs map[winmd.Index]struct{}
	// DefinedTypeDefs is the set of all TypeDefs that have been defined so far by
	// WriteUsedTypeDefs. It ensures WriteUsedTypeDefs doesn't write the same definition twice.
	definedTypeDefs map[winmd.Index]struct{}

	currentMethodIndex winmd.Index
	// methodNeedingTypeRef is a map of each needed TypeRef to the MethodDef(s) that need it. This
	// is maintained for debugging purposes, e.g. in case it is difficult to determine why certain
	// type definitions were generated. This also lets you start from a poorly generated type
	// definition from this library, find the syscall(s) that need it, then find an implementation
	// in an existing syscall library (Go src/syscall/syscall_windows.go, x/sys) to compare against.
	methodNeedingTypeRef map[winmd.Index][]winmd.Index

	// typeDefGoName is the Go identifier generated upon first seeing this TypeDef. It is stored
	// here rather than calculating it upon demand: there may be information available upon first
	// encounter that's hard to recalculate later, such as the parent of a nested def.
	typeDefGoName map[winmd.Index]string

	// AllTypeDefs is the set of all TypeDefs in the winmd file indexed by namespace + "::" + name.
	AllTypeDefs map[string]winmd.Index
	// MethodDefImplMap maps MethodDef -> ImplMap with matching MemberForwarded index.
	MethodDefImplMap map[winmd.Index]*winmd.ImplMap
	// FieldConstant maps Field -> Constant with the field as its parent.
	FieldConstant map[winmd.Index]*winmd.Constant
	// TypeDefNativeTypedefAttribute maps TypeDef -> CustomAttribute (if NativeTypedefAttribute).
	TypeDefNativeTypedefAttribute map[winmd.Index]*winmd.CustomAttribute
	// FieldOffset maps Field -> the Value of the FieldIndexAttribute on that field.
	FieldOffset map[winmd.Index]uint32
	// NestedTypeDefParent	maps each nested TypeDef -> its parent TypeDef.
	NestedTypeDefParent   map[winmd.Index]winmd.Index
	NestedTypeDefChildren map[winmd.Index][]winmd.Index
}

// NewContext creates a Context. Performs some pre-processing to improve generation performance.
func NewContext(f *winmd.Metadata) (*Context, error) {
	l := &Context{
		Metadata: f,

		neededTypeDefs:       make(map[winmd.Index]struct{}),
		definedTypeDefs:      make(map[winmd.Index]struct{}),
		methodNeedingTypeRef: make(map[winmd.Index][]winmd.Index),

		typeDefGoName: make(map[winmd.Index]string),

		AllTypeDefs:                   make(map[string]winmd.Index, f.Tables.TypeDef.Len),
		MethodDefImplMap:              make(map[winmd.Index]*winmd.ImplMap),
		FieldConstant:                 make(map[winmd.Index]*winmd.Constant),
		TypeDefNativeTypedefAttribute: make(map[winmd.Index]*winmd.CustomAttribute),
		FieldOffset:                   make(map[winmd.Index]uint32),
		NestedTypeDefParent:           make(map[winmd.Index]winmd.Index),
		NestedTypeDefChildren:         make(map[winmd.Index][]winmd.Index),
	}
	// Scan some tables to create lookups for later.
	for i := uint32(0); i < f.Tables.TypeDef.Len; i++ {
		r, err := f.Tables.TypeDef.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		// Nested types can't be resolved at module scope. Skip.
		if r.Flags&flags.TypeAttributes_NestedPublic != 0 {
			continue
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
	for i := uint32(0); i < f.Tables.CustomAttribute.Len; i++ {
		a, err := f.Tables.CustomAttribute.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		if a.Type.Tag != coded.CustomAttributeType_MemberRef {
			continue
		}
		m, err := f.Tables.MemberRef.Record(a.Type.Index)
		if err != nil {
			return nil, err
		}
		if m.Class.Tag != coded.MemberRefParent_TypeRef {
			continue
		}
		c, err := f.Tables.TypeRef.Record(m.Class.Index)
		if err != nil {
			return nil, err
		}
		if c.Namespace.String() != "Windows.Win32.Interop" || c.Name.String() != "NativeTypedefAttribute" {
			continue
		}
		if a.Parent.Tag != coded.HasCustomAttribute_TypeDef {
			continue
		}
		if existing, ok := l.TypeDefNativeTypedefAttribute[a.Parent.Index]; ok {
			return nil, fmt.Errorf(
				"multiple NativeTypedefAttribute rows found pointing at TypeDef %v: found %v; already found %v",
				a.Parent.Index,
				i,
				existing)
		}
		l.TypeDefNativeTypedefAttribute[a.Parent.Index] = a
	}
	for i := uint32(0); i < f.Tables.FieldLayout.Len; i++ {
		layout, err := f.Tables.FieldLayout.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		l.FieldOffset[layout.Field] = layout.Offset
	}
	for i := uint32(0); i < f.Tables.NestedClass.Len; i++ {
		nest, err := f.Tables.NestedClass.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		if existing, ok := l.NestedTypeDefParent[nest.NestedClass]; ok {
			return nil, fmt.Errorf(
				"multiple TypeDef nest parents for TypeDef %v: found %v; already found %v",
				nest.NestedClass,
				nest.EnclosingClass,
				existing)
		}
		l.NestedTypeDefParent[nest.NestedClass] = nest.EnclosingClass
		l.NestedTypeDefChildren[nest.EnclosingClass] = append(l.NestedTypeDefChildren[nest.EnclosingClass], nest.NestedClass)
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
func (c *Context) WriteMethod(w io.StringWriter, methodIndex winmd.Index, method *winmd.MethodDef) error {
	goName := method.Name.String()
	c.currentMethodIndex = methodIndex
	defer func() { c.currentMethodIndex = 0 }()

	w.WriteString("//sys\t")
	writeEscapedUpper(w, goName)
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

	// Find the DllImport pseudo-custom attribute (§II.21.2.1) for module name and return semantics.
	var moduleName string
	var lastErr bool
	if implMap, ok := c.MethodDefImplMap[methodIndex]; ok {
		// TODO: Map of parsed module refs?
		if implMap.MappingFlags&flags.PInvokeAttributes_SupportsLastError != 0 {
			lastErr = true
		}
		mr, err := c.Metadata.Tables.ModuleRef.Record(implMap.ImportScope)
		if err != nil {
			return err
		}
		moduleName = strings.ToLower(mr.Name.String())
		if moduleName == "kernel32" {
			moduleName = ""
		}
	}

	// Find and write return value(s), if they exist.
	if value := sig.RetType.Kind != winmd.SigRetTypeKind_Void; value || lastErr {
		w.WriteString(" (")
		if value {
			// General return value name, because mkwinsyscall needs one.
			// TODO: Make more useful names, like x/sys has. Might need to be guided by human input, because better names don't exist in the winmd file.
			w.WriteString("r ")
			if err := c.writeType(w, &sig.RetType.Type); err != nil {
				return err
			}
			if lastErr {
				w.WriteString(", ")
			}
		}
		if lastErr {
			w.WriteString("err error")
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
		writeEscapedUpper(w, method.Name.String())
	}
	return nil
}

func (c *Context) writeParam(w io.StringWriter, p *winmd.Param) {
	// TODO: Use p.Flags to determine how to translate each param to a Go type?
	writeEscapedParam(w, p.Name.String())
}

func (c *Context) writeType(b io.StringWriter, p *winmd.SigType) error {
	// Keep track of visited types to detect a cycle.
	var visited map[*winmd.SigType]struct{}
	markVisited := func(p *winmd.SigType) {
		// Allocate the visited map at the last possible moment. In simple cases, it isn't needed.
		if visited == nil {
			visited = make(map[*winmd.SigType]struct{})
		}
		visited[p] = struct{}{}
	}

	// Declare the func before defining it to let it capture the variable and recurse.
	var visitType func(p *winmd.SigType) error
	visitType = func(p *winmd.SigType) error {
		if _, ok := visited[p]; ok {
			return fmt.Errorf("cycle detected in type definition: already visited %v", p)
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
			switch v := p.Value.(type) {
			case winmd.CodedIndex:
				switch v.Tag {
				case coded.TypeDefOrRefOrSpec_TypeDef:
					record, err := c.Metadata.Tables.TypeDef.Record(v.Index)
					if err != nil {
						return err
					}
					if _, ok := c.definedTypeDefs[v.Index]; !ok {
						if v.Index == 0 {
							panic("")
						}
						c.neededTypeDefs[v.Index] = struct{}{}
					}
					b.WriteString(escapedUpper(record.Name.String()))

				case coded.TypeDefOrRefOrSpec_TypeRef:
					record, err := c.Metadata.Tables.TypeRef.Record(v.Index)
					if err != nil {
						return err
					}
					defIndex, err := c.findTypeDefIndex(record)
					if err != nil {
						if errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
							log.Printf("skipping type: %v", err)
						} else {
							return err
						}
					} else {
						if _, ok := c.definedTypeDefs[defIndex]; !ok {
							if defIndex == 0 {
								panic("")
							}
							c.neededTypeDefs[defIndex] = struct{}{}
						}
						c.methodNeedingTypeRef[v.Index] = append(c.methodNeedingTypeRef[v.Index], c.currentMethodIndex)
					}
					b.WriteString(escapedUpper(record.Name.String()))

				default:
					return fmt.Errorf("unexpected coded index tag for type Value: %#v", v)
				}

				// Types can nest. A pointer to another type is a very common case.
			case winmd.SigType:
				markVisited(p)
				return visitType(&v)
			case winmd.SigArray:
				b.WriteString("[]")
				markVisited(p)
				return visitType(&v.Type)

			default:
				return fmt.Errorf("unexpected type for type Value: %#v", p.Value)
			}
			return nil
		}
		return nil
	}
	return visitType(p)
}

func (c *Context) WriteTypeDef(b io.StringWriter, i winmd.Index) error {
	def, err := c.Metadata.Tables.TypeDef.Record(i)
	if err != nil {
		return err
	}

	name := escapedUpper(def.Name.String())
	if parentIndex, ok := c.NestedTypeDefParent[i]; ok {
		// Since this is a nested type, we know the parent has already been named.
		name = c.typeDefGoName[parentIndex] + "_nest_" + name
	}
	if len(name) == 0 {
		return fmt.Errorf("unable to find name for %q, %v", name, i)
	}
	c.typeDefGoName[i] = name

	if def.Flags&flags.TypeAttributes_ClassSemanticsMask == flags.TypeAttributes_Interface {
		// TODO: Handle interfaces. Currently writes a struct with no members.
		b.WriteString("// Interface type is likely missing members. Not yet implemented in go-winmd.\n")
		return c.writeTypeDefStruct(b, i, def)
	}
	switch def.Extends.Tag {
	case coded.TypeDefOrRef_TypeRef:
		// TODO: Keep track of this index rather than looking it up for each enum type.
		record, err := c.Metadata.Tables.TypeRef.Record(def.Extends.Index)
		if err != nil {
			return err
		}
		if record.Namespace.String() == "System" {
			if record.Name.String() == "Enum" {
				return c.writeTypeDefEnum(b, def)
			}
			if record.Name.String() == "MulticastDelegate" {
				b.WriteString("type ")
				b.WriteString(escapedUpper(def.Name.String()))
				b.WriteString(" uintptr\n")
				return nil
			}
		}
		if _, ok := c.TypeDefNativeTypedefAttribute[i]; ok {
			return c.writeTypeDefNative(b, def)
		}
		return c.writeTypeDefStruct(b, i, def)
	default:
		return fmt.Errorf("unexpected type extends coded index %#v in def %v :: %v", def.Extends, def.Namespace, def.Name)
	}
}

func (c *Context) writeTypeDefEnum(b io.StringWriter, def *winmd.TypeDef) error {
	b.WriteString("type ")
	writeEscapedUpper(b, def.Name.String())

	// Per §I.8.5.2 CLS Rule 7, the underlying type is the type of the field "__value". Find it.
	var underlyingType *winmd.SigType

	type member struct {
		Name     winmd.String
		HexValue string
	}
	// The number of enum members is the total number of fields minus the special "value__".
	members := make([]member, 0, def.FieldList.End-def.FieldList.Start-1)

	for i := def.FieldList.Start; i < def.FieldList.End; i++ {
		fd, err := c.Metadata.Tables.Field.Record(i)
		if err != nil {
			return err
		}
		if fd.Name.String() == "value__" {
			signature, err := c.Metadata.FieldSignature(fd.Signature)
			if err != nil {
				return err
			}
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
			if hex[0] == '-' {
				hex = "-0x" + hex[1:]
			} else {
				hex = "0x" + hex
			}
		} else {
			return fmt.Errorf("unable to find default value for field %v", fd.Name)
		}

		// Don't write the members yet. We haven't written the enum type definition yet, and the
		// order is important for readability in docs.
		p := member{fd.Name, hex}
		members = append(members, p)
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
		name := pair.Name.String()
		b.WriteString("\t")
		// Add enum name prefix to generated name if the original member name doesn't already have
		// the prefix. This may be necessary to avoid collisions, and also makes the API easier to
		// find in Go via autocomplete.
		// TODO: Be a bit smarter? E.g. accept BCRYPT_CIPHER_OPERATION for BCRYPT_OPERATION enum without changing it.
		if !strings.HasPrefix(name, def.Name.String()) {
			writeEscapedUpper(b, def.Name.String())
			b.WriteString("_")
		}
		writeEscapedUpper(b, name)
		b.WriteString(" ")
		writeEscapedUpper(b, def.Name.String())
		b.WriteString(" = ")
		b.WriteString(pair.HexValue)
		b.WriteString("\n")
	}
	b.WriteString(")\n")
	return nil
}

func (c *Context) writeTypeDefNative(b io.StringWriter, def *winmd.TypeDef) error {
	b.WriteString("type ")
	writeEscapedUpper(b, def.Name.String())
	b.WriteString(" ")
	if def.FieldList.Start+1 != def.FieldList.End {
		return fmt.Errorf("expected exactly one field for native typedef %v", def.Name)
	}
	fd, err := c.Metadata.Tables.Field.Record(def.FieldList.Start)
	if err != nil {
		return err
	}
	signature, err := c.Metadata.FieldSignature(fd.Signature)
	if err != nil {
		return err
	}
	if err := c.writeType(b, &signature.Type); err != nil {
		return err
	}
	b.WriteString("\n")
	return nil
}

func (c *Context) writeTypeDefStruct(b io.StringWriter, i winmd.Index, def *winmd.TypeDef) error {
	b.WriteString("type ")
	writeEscapedUpper(b, c.typeDefGoName[i])
	b.WriteString(" struct {\n")
	// TODO: Look at field index values to find union types. From each union, only define a struct with the first field.
	//
	// Example where the union is not critical and ignored in x/sys:
	//
	//sys	GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) = iphlpapi.GetAdaptersAddresses
	//
	// type IpAdapterAddresses struct {
	//	Length                 uint32
	//	IfIndex                uint32
	//	Next                   *IpAdapterAddresses
	//	AdapterName            *byte
	//
	// public struct IP_ADAPTER_ADDRESSES_LH
	// {
	// 	[StructLayout(LayoutKind.Explicit)]
	//	public struct _Anonymous1_e__Union
	//	{
	//		public struct _Anonymous_e__Struct
	//		{
	//			public uint Length;
	//			public uint IfIndex;
	//		}
	//
	//		[FieldOffset(0)]
	//		public ulong Alignment;
	//
	//		[FieldOffset(0)]
	//		public _Anonymous_e__Struct Anonymous;
	//	}
	//	public _Anonymous1_e__Union Anonymous1;
	//	public unsafe IP_ADAPTER_ADDRESSES_LH* Next;
	//	...
	if err := c.writeStructFields(b, def); err != nil {
		return err
	}
	b.WriteString("}\n")
	return nil
}

func (c *Context) writeStructFields(b io.StringWriter, def *winmd.TypeDef) error {
	// Write a definition for each field. Take FieldOffset into account, but minimally: once a
	// FieldOffset is "used", don't write another field that uses the same explicit field offset.
	// This isn't accurate, but it may be good enough. In many cases, only "0" is used.
	usedFieldOffset := make(map[uint32]struct{})
	for i := def.FieldList.Start; i < def.FieldList.End; i++ {
		if o, ok := c.FieldOffset[i]; ok {
			if _, used := usedFieldOffset[o]; used {
				continue
			}
			usedFieldOffset[o] = struct{}{}
		}
		if err := c.writeStructField(b, i); err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) writeStructField(b io.StringWriter, fieldIndex winmd.Index) error {
	fd, err := c.Metadata.Tables.Field.Record(fieldIndex)
	if err != nil {
		return err
	}
	signature, err := c.Metadata.FieldSignature(fd.Signature)
	if err != nil {
		return err
	}
	if signature.Type.Kind == flags.ElementType_VALUETYPE {
		if v, ok := signature.Type.Value.(winmd.CodedIndex); ok && v.Tag == coded.TypeDefOrRefOrSpec_TypeRef {
			ref, err := c.Metadata.Tables.TypeRef.Record(v.Index)
			if err != nil {
				return err
			}
			if ref.ResolutionScope.Tag == coded.ResolutionScope_TypeRef {
				// This is a reference to a nested type. Embed each of its fields. Don't create a
				// new named type, because in the winmd files we work with, the nested structs don't
				// have meaningful names. It makes the API clunky if we generate unique names.
				index, err := c.findTypeDefIndex(ref)
				if err != nil {
					return err
				}
				def, err := c.Metadata.Tables.TypeDef.Record(index)
				if err != nil {
					return err
				}
				if err := c.writeStructFields(b, def); err != nil {
					return err
				}
				return nil
			}
		}
	}
	b.WriteString("\t")
	writeEscapedUpper(b, fd.Name.String())
	b.WriteString(" ")
	if err := c.writeType(b, &signature.Type); err != nil {
		return err
	}
	b.WriteString("\n")
	return nil
}

func (c *Context) findModuleTypeDefByName(namespace, name string) (i winmd.Index, ok bool) {
	// TODO: Use better keys for this lookup. ("struct{Namespace, Name Index}" where Index is a blob heap index?)
	i, ok = c.AllTypeDefs[namespace+"::"+name]
	return
}

var errTypeDefNotDefinedInCurrentModule = errors.New("TypeRef points to TypeDef not defined in this module")

func (c *Context) findTypeDefIndex(ref *winmd.TypeRef) (winmd.Index, error) {
	// Nested TypeDefs don't have a unique namespace+name, so we need to traverse upward to find
	// the ancestor that's a module-level TypeRef, find its TypeDef, then traverse back down,
	// converting TypeRef -> TypeDef for each level to finally find the nested struct's TypeDef.

	// Limit how much nesting we handle. This is much more permissive than any nesting we've
	// actually seen in winmd files, and will prevent infinite recursion and make degenerate cases
	// (large table consisting of 100% nested types) fail more quickly.
	visitsRemaining := 64

	// TODO: Cache visit results in Context, so finding the chain for sibling and other related nested types doesn't iterate over the same data yet again?
	// Nesting chains in the winmd files tend to be short, so this may not be a significant improvement.
	var visit func(r *winmd.TypeRef) (winmd.Index, error)
	visit = func(r *winmd.TypeRef) (winmd.Index, error) {
		visitsRemaining--
		if visitsRemaining == 0 {
			return 0, errors.New("exceeded recursion limit while finding TypeDef for nested type")
		}
		switch r.ResolutionScope.Tag {
		// We assume module-level TypeRefs refer to the current module.
		case coded.ResolutionScope_Module:
			if def, ok := c.findModuleTypeDefByName(r.Namespace.String(), r.Name.String()); ok {
				return def, nil
			}
		// TypeRef scope indicates that this is a nested type. The Index is the immediate parent.
		case coded.ResolutionScope_TypeRef:
			parentRef, err := c.Metadata.Tables.TypeRef.Record(r.ResolutionScope.Index)
			if err != nil {
				return 0, err
			}
			// Recurse to find the parent def.
			parentDefIndex, err := visit(parentRef)
			if err != nil {
				return 0, err
			}
			// Look in the parent def for the def matching the ref we're looking for.
			// TODO: String lookup with a pre-created map?
			for _, child := range c.NestedTypeDefChildren[parentDefIndex] {
				possibleDef, err := c.Metadata.Tables.TypeDef.Record(child)
				if err != nil {
					return 0, err
				}
				if possibleDef.Name.String() == r.Name.String() {
					return child, nil
				}
			}
		}
		return 0, fmt.Errorf("could not find %v :: %v, %w", r.Namespace, r.Name, errTypeDefNotDefinedInCurrentModule)
	}
	return visit(ref)
}

func (c *Context) WriteUsedTypeDefs(b io.StringWriter) error {
	b.WriteString("\n\n// Types used in generated APIs\n\n")
	var usedTypeDefIndices []winmd.Index
	// Keep going until we stop finding new type refs that need definitions.
	for len(c.neededTypeDefs) > 0 {
		for i := range c.neededTypeDefs {
			usedTypeDefIndices = append(usedTypeDefIndices, i)
			c.definedTypeDefs[i] = struct{}{}
		}
		c.neededTypeDefs = make(map[winmd.Index]struct{})
		sort.Slice(usedTypeDefIndices, func(i, j int) bool {
			return usedTypeDefIndices[i] < usedTypeDefIndices[j]
		})
		if len(usedTypeDefIndices) == 0 {
			return nil
		}
		for _, index := range usedTypeDefIndices {
			// Writing the type def adds new entries to neededTypeRefs if we haven't seen them yet.
			if err := c.WriteTypeDef(b, index); err != nil {
				return err
			}
			b.WriteString("\n")
		}
		usedTypeDefIndices = usedTypeDefIndices[:0]
	}
	return nil
}

// writeEscapedParam writes the given string, adding a suffix if it is a reserved Go keyword. Leave
// the case as-is (unlike writeEscapedUpper) because lowercase is desirable for params.
func writeEscapedParam(b io.StringWriter, s string) {
	if token.IsKeyword(s) {
		s += "Param"
	}
	b.WriteString(s)
}

// writeEscapedUpper writes the given string with the first character in uppercase. All Go keywords
// are lowercase, so uppercasing the first letter does two things: escapes names like "type" and
// exports the generated types/fields.
func writeEscapedUpper(b io.StringWriter, s string) {
	if len(s) > 1 {
		s = strings.ToUpper(string(s[0])) + s[1:]
	}
	b.WriteString(s)
}

func escapedUpper(s string) string {
	if len(s) > 1 {
		s = strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}
