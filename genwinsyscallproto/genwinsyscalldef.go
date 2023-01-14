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

	// resolvedDefs maps type name keys -> resolved TypeDef information.
	resolvedDefs map[typeNameKey]*resolvedDef
	// resolvedDefsByIndex maps TypeDef Index -> resolved TypeDef information.
	resolvedDefsByIndex map[winmd.Index]*resolvedDef
	// unresolvedDefs maps type name keys -> winmd.TypeDef index.
	unresolvedDefs map[typeNameKey]winmd.Index
	// unresolvableTypeRefs is a set of TypeRefs that were discovered but unable to be resolved
	// inside the current module.
	unresolvableTypeRefs map[typeNameKey]*winmd.TypeRef

	// The maps below index commonly used winmd table relationships to allow fast access when
	// interpreting the metadata and writing the Go source code.

	// methodDefImplMap maps MethodDef index -> ImplMap with matching MemberForwarded index.
	methodDefImplMap map[winmd.Index]*winmd.ImplMap
	// fieldConstant maps Field index -> the Constant with the field as its parent.
	fieldConstant map[winmd.Index]*winmd.Constant
	// typeDefNativeTypedefAttribute maps TypeDef -> CustomAttribute (if NativeTypedefAttribute).
	typeDefNativeTypedefAttribute map[winmd.Index]*winmd.CustomAttribute
	// fieldOffset maps Field -> the Value of the FieldIndexAttribute on that field.
	fieldOffset map[winmd.Index]uint32
	// nestedTypeDefChildren maps TypeDef -> all of its child (nested) TypeDefs.
	nestedTypeDefChildren map[winmd.Index][]winmd.Index
}

// NewContext creates a Context. Reads some tables in their entirety to fill in internal index maps
// that improve generation performance at the cost of startup time.
func NewContext(f *winmd.Metadata) (*Context, error) {
	l := &Context{
		Metadata: f,

		resolvedDefs:         make(map[typeNameKey]*resolvedDef),
		resolvedDefsByIndex:  make(map[winmd.Index]*resolvedDef),
		unresolvedDefs:       make(map[typeNameKey]winmd.Index),
		unresolvableTypeRefs: make(map[typeNameKey]*winmd.TypeRef),

		methodDefImplMap:              make(map[winmd.Index]*winmd.ImplMap),
		fieldConstant:                 make(map[winmd.Index]*winmd.Constant),
		typeDefNativeTypedefAttribute: make(map[winmd.Index]*winmd.CustomAttribute),
		fieldOffset:                   make(map[winmd.Index]uint32),
		nestedTypeDefChildren:         make(map[winmd.Index][]winmd.Index),
	}
	// We index tables and resolved defs with the assumption that there is exactly one module. For
	// winmd files, we expect this to always be the case.
	if f.Tables.Module.Len != 1 {
		return nil, fmt.Errorf("expected exactly one module in the file, but found %v", f.Tables.Module.Len)
	}
	for i := uint32(0); i < f.Tables.TypeDef.Len; i++ {
		r, err := f.Tables.TypeDef.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		// Nested types can't be resolved at module scope. Skip.
		if r.Flags&flags.TypeAttributes_NestedPublic != 0 {
			continue
		}
		l.unresolvedDefs[typeDefKey(r)] = winmd.Index(i)
	}
	for i := uint32(0); i < f.Tables.ImplMap.Len; i++ {
		im, err := f.Tables.ImplMap.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		if im.MemberForwarded.Tag != coded.MemberForwarded_MethodDef {
			continue
		}
		if existing, ok := l.methodDefImplMap[im.MemberForwarded.Index]; ok {
			return nil, fmt.Errorf(
				"multiple ImplMap rows found pointing at MethodDef %v: found %v; already found %v",
				im.MemberForwarded.Index,
				i,
				existing)
		}
		l.methodDefImplMap[im.MemberForwarded.Index] = im
	}
	for i := uint32(0); i < f.Tables.Constant.Len; i++ {
		c, err := f.Tables.Constant.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		if c.Parent.Tag != coded.HasConstant_Field {
			continue
		}
		if existing, ok := l.fieldConstant[c.Parent.Index]; ok {
			return nil, fmt.Errorf(
				"multiple Constant rows found pointing at Field %v: found %v; already found %v",
				c.Parent.Index,
				i,
				existing)
		}
		l.fieldConstant[c.Parent.Index] = c
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
		if existing, ok := l.typeDefNativeTypedefAttribute[a.Parent.Index]; ok {
			return nil, fmt.Errorf(
				"multiple NativeTypedefAttribute rows found pointing at TypeDef %v: found %v; already found %v",
				a.Parent.Index,
				i,
				existing)
		}
		l.typeDefNativeTypedefAttribute[a.Parent.Index] = a
	}
	for i := uint32(0); i < f.Tables.FieldLayout.Len; i++ {
		layout, err := f.Tables.FieldLayout.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		l.fieldOffset[layout.Field] = layout.Offset
	}
	for i := uint32(0); i < f.Tables.NestedClass.Len; i++ {
		nest, err := f.Tables.NestedClass.Record(winmd.Index(i))
		if err != nil {
			return nil, err
		}
		l.nestedTypeDefChildren[nest.EnclosingClass] = append(l.nestedTypeDefChildren[nest.EnclosingClass], nest.NestedClass)
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
//
// methodIndex and method must match. Only methodIndex is necessary, but method is also accepted to
// avoid unnecessary decoding. The caller has likely already decoded the method to filter by name.
func (c *Context) WriteMethod(w io.StringWriter, methodIndex winmd.Index, method *winmd.MethodDef) error {
	goName := method.Name.String()

	w.WriteString("//sys\t")
	w.WriteString(escapedUpper(goName))
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
			// A signature with data in ".param" hasn't been encountered. See https://github.com/microsoft/go-winmd/issues/9
			return fmt.Errorf("unsupported method: expected param row with sequence 0 to be empty, but: %#v", param)
		}

		// Param data is in two places: a param table row and the signature. Index into the
		// signature param slice by converting the 1-based sequence value. Note that we assume
		// ascending Sequence values for proper formatting: this is stated to be true in the
		// informational section of §II.22.33, point 4.
		//
		// Note: this assumes there are no gaps in Sequence values, but technically gaps are
		// possible per §II.22.33 information point 5. See https://github.com/microsoft/go-winmd/issues/10
		i := param.Sequence - 1
		if i > 0 {
			w.WriteString(", ")
		}
		writeEscapedParam(w, param.Name.String())
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
	if implMap, ok := c.methodDefImplMap[methodIndex]; ok {
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
			// Generated returns in general could be better. See https://github.com/microsoft/go-winmd/issues/12
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
		w.WriteString(method.Name.String())
	}
	return nil
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
		case flags.ElementType_I1:
			b.WriteString("int8")
		case flags.ElementType_U1:
			b.WriteString("uint8")
		case flags.ElementType_I2:
			b.WriteString("int16")
		case flags.ElementType_U2, flags.ElementType_CHAR:
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
					def, err := c.resolveTypeDef(v.Index)
					if err != nil {
						return err
					}
					if def.NeedsPointerWhenUsed() {
						b.WriteString("*")
					}
					b.WriteString(def.GoName)
				case coded.TypeDefOrRefOrSpec_TypeRef:
					def, err := c.resolveTypeRef(v.Index)
					if err != nil && !errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
						return err
					}
					if def == nil {
						ref, err := c.Metadata.Tables.TypeRef.Record(v.Index)
						if err != nil {
							return err
						}
						b.WriteString(ref.Name.String())
						c.unresolvableTypeRefs[typeRefKey(ref)] = ref
					} else {
						if def.NeedsPointerWhenUsed() {
							b.WriteString("*")
						}
						b.WriteString(def.GoName)
					}
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

type typeNameKey struct {
	NamespaceStart uint32
	NameStart      uint32
}

func typeRefKey(ref *winmd.TypeRef) typeNameKey {
	return typeNameKey{ref.Namespace.Start, ref.Name.Start}
}

func typeDefKey(def *winmd.TypeDef) typeNameKey {
	return typeNameKey{def.Namespace.Start, def.Name.Start}
}

type resolvedDef struct {
	Index     winmd.Index
	Namespace winmd.String
	Name      winmd.String

	GoName string

	Native bool
	// NativePointer enables a workaround for the mkwinsyscall expectations not quite matching a
	// straightforward interpretation of the winmd. If we generate "//sys" with "thing PWSTR" and
	// "typedef PWSTR *uint16", mkwinsyscall misses including the "unsafe.Pointer" in
	// "uintptr(unsafe.Pointer(thing))". This code means we generate "thing *PWSTR" and "typedef
	// PWSTR uint16" instead, which works.
	NativePointer bool

	Parent   *resolvedDef
	Children []*resolvedDef

	def *winmd.TypeDef
}

func (r *resolvedDef) IsInterface() bool {
	return r.def.Flags&flags.TypeAttributes_ClassSemanticsMask == flags.TypeAttributes_Interface
}

func (r *resolvedDef) NeedsPointerWhenUsed() bool {
	return r.NativePointer || r.IsInterface()
}

var errTypeDefNotDefinedInCurrentModule = errors.New("TypeRef points to TypeDef not defined in this module")

func (c *Context) resolveTypeRef(refIndex winmd.Index) (*resolvedDef, error) {
	// Nested TypeDefs don't have a unique namespace+name, so we need to traverse upward to find
	// the ancestor that's a module-level TypeRef, find its TypeDef, then traverse back down,
	// converting TypeRef -> TypeDef for each level to finally find the nested struct's TypeDef.

	// Prevent infinite recursion with a visited set.
	var visited map[winmd.Index]struct{}

	var visit func(refIndex winmd.Index) (*resolvedDef, error)
	visit = func(refIndex winmd.Index) (*resolvedDef, error) {
		if _, ok := visited[refIndex]; ok {
			return nil, fmt.Errorf("cycle: visited %v twice", refIndex)
		}
		if visited == nil {
			visited = make(map[winmd.Index]struct{})
		}
		visited[refIndex] = struct{}{}

		r, err := c.Metadata.Tables.TypeRef.Record(refIndex)
		if err != nil {
			return nil, err
		}
		switch r.ResolutionScope.Tag {
		// We assume module-level TypeRefs refer to the current module.
		case coded.ResolutionScope_Module:
			if def, ok := c.resolvedDefs[typeRefKey(r)]; ok {
				return def, nil
			}
			if defIndex, ok := c.unresolvedDefs[typeRefKey(r)]; ok {
				return c.resolveTypeDef(defIndex)
			}
		// TypeRef scope indicates that this is a nested type. The Index is the immediate parent.
		case coded.ResolutionScope_TypeRef:
			// Recurse to find the parent def.
			parentDefIndex, err := visit(r.ResolutionScope.Index)
			if err != nil {
				return nil, err
			}
			// Look in the parent def for the def matching the ref we're looking for.
			for _, child := range parentDefIndex.Children {
				if child.Name.Start == r.Name.Start {
					return child, nil
				}
			}
		}
		return nil, fmt.Errorf("could not find %v :: %v, %w", r.Namespace, r.Name, errTypeDefNotDefinedInCurrentModule)
	}
	return visit(refIndex)
}

func (c *Context) resolveTypeDef(defIndex winmd.Index) (*resolvedDef, error) {
	// Prevent infinite recursion with a visited set.
	var visited map[winmd.Index]struct{}

	// Resolving a typedef also resolves all its children, transitively.
	var visit func(defIndex winmd.Index) (*resolvedDef, error)
	visit = func(defIndex winmd.Index) (*resolvedDef, error) {
		if _, ok := visited[defIndex]; ok {
			return nil, fmt.Errorf("cycle: visited %v twice", defIndex)
		}
		if visited == nil {
			visited = make(map[winmd.Index]struct{})
		}
		visited[defIndex] = struct{}{}

		// Check if this index has already been resolved to avoid creating duplicate resolvedDefs.
		if r, ok := c.resolvedDefsByIndex[defIndex]; ok {
			return r, nil
		}
		def, err := c.Metadata.Tables.TypeDef.Record(defIndex)
		if err != nil {
			return nil, err
		}
		r := resolvedDef{
			Index:     defIndex,
			Namespace: def.Namespace,
			Name:      def.Name,
			GoName:    escapedUpper(def.Name.String()),
			def:       def,
		}
		if _, ok := c.typeDefNativeTypedefAttribute[defIndex]; ok {
			r.Native = true
			if def.FieldList.Start+1 != def.FieldList.End {
				return nil, fmt.Errorf("expected exactly one field for native typedef %v", r.def.Name)
			}
			fd, err := c.Metadata.Tables.Field.Record(r.def.FieldList.Start)
			if err != nil {
				return nil, err
			}
			signature, err := c.Metadata.FieldSignature(fd.Signature)
			if err != nil {
				return nil, err
			}
			if signature.Type.Kind == flags.ElementType_PTR {
				to := signature.Type.Value.(winmd.SigType)
				if to.Kind != flags.ElementType_VOID {
					r.NativePointer = true
					// Clarify the Go name, to avoid confusion for anyone with a strong expectation
					// about types like PWSTR.
					r.GoName += "Element"
				}
			}
		}
		// Nested types can't be resolved at module scope. Don't add it to the module lookup.
		if def.Flags&flags.TypeAttributes_NestedPublic == 0 {
			c.resolvedDefs[typeDefKey(def)] = &r
		}
		c.resolvedDefsByIndex[defIndex] = &r
		for _, childIndex := range c.nestedTypeDefChildren[defIndex] {
			child, err := visit(childIndex)
			if err != nil {
				return nil, err
			}
			// Establish bidirectional links.
			child.Parent = &r
			r.Children = append(r.Children, child)
		}
		return &r, nil
	}
	return visit(defIndex)
}

func (c *Context) writeTypeDef(b io.StringWriter, r *resolvedDef) error {
	if r.def.Flags&flags.TypeAttributes_ClassSemanticsMask == flags.TypeAttributes_Interface {
		// Issue tracking implementing interface types: https://github.com/microsoft/go-winmd/issues/14
		b.WriteString("// Interface type is likely missing members. Not yet implemented in go-winmd.\n")
	}
	switch r.def.Extends.Tag {
	case coded.TypeDefOrRef_TypeRef:
		extendsRef, err := c.Metadata.Tables.TypeRef.Record(r.def.Extends.Index)
		if err != nil {
			return err
		}
		if extendsRef.Namespace.String() == "System" {
			if extendsRef.Name.String() == "Enum" {
				return c.writeTypeDefEnum(b, r)
			}
			if extendsRef.Name.String() == "MulticastDelegate" {
				b.WriteString("type ")
				b.WriteString(r.GoName)
				b.WriteString(" uintptr\n")
				return nil
			}
		}
		if r.Native {
			return c.writeTypeDefNative(b, r)
		}
		return c.writeTypeDefStruct(b, r)
	case coded.Null:
		return c.writeTypeDefStruct(b, r)
	default:
		return fmt.Errorf("unexpected type extends coded index %#v in def %v :: %v", r.def.Extends, r.def.Namespace, r.def.Name)
	}
}

func (c *Context) writeTypeDefEnum(b io.StringWriter, r *resolvedDef) error {
	b.WriteString("type ")
	b.WriteString(r.GoName)

	// Per §I.8.5.2 CLS Rule 7, the underlying type is the type of the field "__value". Find it.
	var underlyingType *winmd.SigType

	type member struct {
		Name     winmd.String
		HexValue string
	}
	// The number of enum members is the total number of fields minus the special "value__".
	members := make([]member, 0, r.def.FieldList.End-r.def.FieldList.Start-1)

	for i := r.def.FieldList.Start; i < r.def.FieldList.End; i++ {
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
		if constant, ok := c.fieldConstant[i]; ok {
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
		if !strings.HasPrefix(name, r.def.Name.String()) {
			b.WriteString(r.GoName)
			b.WriteString("_")
		}
		b.WriteString(escapedUpper(name))
		b.WriteString(" ")
		b.WriteString(r.GoName)
		b.WriteString(" = ")
		b.WriteString(pair.HexValue)
		b.WriteString("\n")
	}
	b.WriteString(")\n")
	return nil
}

func (c *Context) writeTypeDefNative(b io.StringWriter, r *resolvedDef) error {
	b.WriteString("type ")
	b.WriteString(r.GoName)
	b.WriteString(" ")
	if r.def.FieldList.Start+1 != r.def.FieldList.End {
		return fmt.Errorf("expected exactly one field for native typedef %v", r.def.Name)
	}
	fd, err := c.Metadata.Tables.Field.Record(r.def.FieldList.Start)
	if err != nil {
		return err
	}
	signature, err := c.Metadata.FieldSignature(fd.Signature)
	if err != nil {
		return err
	}
	if r.NativePointer {
		// Modify sig to skip the pointer, for "type PWSTR uint16" rather than "type PWSTR *uint16".
		to := signature.Type.Value.(winmd.SigType)
		signature.Type = to
	}
	if err := c.writeType(b, &signature.Type); err != nil {
		return err
	}
	b.WriteString("\n")
	return nil
}

func (c *Context) writeTypeDefStruct(b io.StringWriter, r *resolvedDef) error {
	b.WriteString("type ")
	b.WriteString(r.GoName)
	b.WriteString(" struct {\n")
	if err := c.writeStructFields(b, r); err != nil {
		return err
	}
	b.WriteString("}\n")
	return nil
}

func (c *Context) writeStructFields(b io.StringWriter, r *resolvedDef) error {
	// Union type support is simple for now. Roughly follow the x/sys approach and pick one union
	// option to implement. See the x/sys/windows "IpAdapterAddresses" struct in syscall_windows.go
	// for an example. Better support is tracked at https://github.com/microsoft/go-winmd/issues/17

	// Write a definition for each field. Take FieldOffset into account, but minimally, for
	// simplicity: once a FieldOffset is "used", don't write another field that uses the same
	// explicit field offset. This isn't accurate, but it may be good enough. In many cases, only
	// "0" is used.
	usedFieldOffset := make(map[uint32]struct{})
	for i := r.def.FieldList.Start; i < r.def.FieldList.End; i++ {
		if o, ok := c.fieldOffset[i]; ok {
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
				def, err := c.resolveTypeRef(v.Index)
				if err != nil {
					return err
				}
				return c.writeStructFields(b, def)
			}
		}
	}
	b.WriteString("\t")
	b.WriteString(escapedUpper(fd.Name.String()))
	b.WriteString(" ")
	if err := c.writeType(b, &signature.Type); err != nil {
		return err
	}
	b.WriteString("\n")
	return nil
}

func (c *Context) WriteUsedTypeDefs(b io.StringWriter) error {
	// Log TypeRefs to the console to let the def know these are expected.
	// This issue tracks better approaches than simply logging: https://github.com/microsoft/go-winmd/issues/18
	for _, r := range c.unresolvableTypeRefs {
		log.Printf("unable to resolve type: %v :: %v", r.Namespace, r.Name)
	}
	b.WriteString("\n\n// Types used in generated APIs\n\n")
	// Keep going until we stop finding new types that need definitions.
	written := make(map[*resolvedDef]struct{})
	for {
		var usedTypeDefs []*resolvedDef
		for _, r := range c.resolvedDefs {
			if _, ok := written[r]; ok {
				continue
			}
			usedTypeDefs = append(usedTypeDefs, r)
			written[r] = struct{}{}
		}
		if len(usedTypeDefs) == 0 {
			break
		}
		// Order is scrambled due to the map. Put it back in some order.
		sort.Slice(usedTypeDefs, func(i, j int) bool {
			return usedTypeDefs[i].Index < usedTypeDefs[j].Index
		})
		for _, r := range usedTypeDefs {
			// Writing the type def (field types in particular) adds new entries to ResolvedDefs if
			// we haven't seen them yet.
			if err := c.writeTypeDef(b, r); err != nil {
				return err
			}
			b.WriteString("\n")
		}
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

// escapedUpper returns the given string with the first character in uppercase. All Go keywords are
// lowercase, so uppercasing the first letter does two things: escapes names like "type" and exports
// the generated types/fields.
func escapedUpper(s string) string {
	if len(s) > 1 {
		s = strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}
