// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Package gowinmd generates Windows syscall function prototypes ("//sys ..." comments)
// using given win32metadata information parsed by go-winmd as specified by ECMA-335.
package gowinmd

import (
	"cmp"
	"errors"
	"fmt"
	"go/token"
	"io"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/microsoft/go-winmd/winmd"
)

// Arch is a bitmask of architectures.
type Arch uint32

const (
	ArchNone  Arch = 0
	Arch386   Arch = 1
	ArchAMD64 Arch = 2
	ArchARM64 Arch = 4
	ArchAll        = Arch386 | ArchAMD64 | ArchARM64
)

// Unique returns a list of unique architectures in the bitmask.
func (a Arch) Unique() []Arch {
	if a&ArchAll == ArchAll {
		return []Arch{ArchAll}
	}
	var ret []Arch
	if a&Arch386 == Arch386 {
		ret = append(ret, Arch386)
	}
	if a&ArchAMD64 == ArchAMD64 {
		ret = append(ret, ArchAMD64)
	}
	if a&ArchARM64 == ArchARM64 {
		ret = append(ret, ArchARM64)
	}
	if len(ret) == 0 {
		return []Arch{ArchNone}
	}
	return ret
}

func (a Arch) String() string {
	switch a {
	case Arch386:
		return "386"
	case ArchAMD64:
		return "amd64"
	case ArchARM64:
		return "arm64"
	case ArchNone:
		return "none"
	case ArchAll:
		return "all"
	default:
		return "unknown"
	}
}

// Context stores data about syscall generation so far, and to improve generation performance. It
// keeps track of the list of used typedefs that may need to also be defined in generated Go code.
type Context struct {
	Metadata *winmd.Metadata

	typeDefCache typeDefCache
	// typeDefsByName maps case-sensitive qualified names to module-level TypeDef indices.
	typeDefsByName map[qualifiedTypeName][]winmd.Index

	// resolvedDefsByIndex maps TypeDef Index -> resolved TypeDef information.
	resolvedDefsByIndex map[winmd.Index]*resolvedDef
	// unresolvableTypeRefs is a set of TypeRefs that were discovered but unable to be resolved
	// inside the current module. Diagnostic names are compared by text only on this error path.
	unresolvableTypeRefs map[qualifiedTypeName]winmd.TypeRef
	// projectedTypeDefs contains native typedefs replaced with their underlying Go pointer type.
	projectedTypeDefs map[winmd.Index]bool
	// requiredTypeDefs contains projected typedefs that are also referenced by generated declarations.
	requiredTypeDefs map[winmd.Index]bool
	projectingType   bool
	// abiLayoutTypeDefs contains explicitly selected types and their field dependencies.
	abiLayoutTypeDefs map[winmd.Index]bool
	writingABIType    bool

	// The maps below index commonly used winmd table relationships to allow fast access when
	// interpreting the metadata and writing the Go source code. This helps with (e.g.) traversing
	// one-way pointers backwards rather than scanning the entire table each time.

	// methodDefImplMap maps MethodDef index -> ImplMap with matching MemberForwarded index.
	methodDefImplMap map[winmd.Index]winmd.ImplMap
	// fieldConstant maps Field index -> the Constant with the field as its parent.
	fieldConstant map[winmd.Index]winmd.Constant
	// typeDefNativeTypedefAttribute maps TypeDef -> CustomAttribute (if NativeTypedefAttribute).
	typeDefNativeTypedefAttribute map[winmd.Index]winmd.CustomAttribute
	// typeDefSupportedArch maps TypeDef -> the Value of the SupportedArchitectureAttribute on that type.
	typeDefSupportedArch map[winmd.Index]Arch
	// methodDefSupportedArch maps MethodDef -> the Value of the SupportedArchitectureAttribute on that method.
	methodDefSupportedArch map[winmd.Index]Arch
	// paramSizeIndex maps Param -> the zero-based parameter index that supplies its element or byte count.
	paramSizeIndex map[winmd.Index]paramSizeInfo
	// fieldOffset maps Field -> the Value of the FieldIndexAttribute on that field.
	fieldOffset map[winmd.Index]uint32
	// classLayout maps TypeDef -> its explicit packing and class size metadata.
	classLayout map[winmd.Index]winmd.ClassLayout
	// nestedTypeDefChildren maps TypeDef -> all of its child (nested) TypeDefs.
	nestedTypeDefChildren map[winmd.Index][]winmd.Index
}

type paramSizeInfo struct {
	index   uint16
	inBytes bool
}

// Projection controls whether signatures preserve the literal WinMD ABI types or use
// mkwinsyscall conventions that provide more idiomatic Go wrappers.
type Projection uint8

const (
	ProjectionRaw Projection = iota
	ProjectionIdiomatic
)

// MethodOptions controls how WriteMethodWithOptions formats a method.
type MethodOptions struct {
	GoName     string
	Projection Projection
}

// ErrOrdinalImport reports an entry point that mkwinsyscall cannot bind by name.
var ErrOrdinalImport = errors.New("mkwinsyscall does not support ordinal import")

// NewContext creates a Context. Reads some tables in their entirety to fill in internal index maps
// that improve generation performance at the cost of startup time.
func NewContext(f *winmd.Metadata) (*Context, error) {
	l := &Context{
		Metadata:             f,
		typeDefCache:         *newTypeDefCache(),
		typeDefsByName:       make(map[qualifiedTypeName][]winmd.Index),
		resolvedDefsByIndex:  make(map[winmd.Index]*resolvedDef),
		unresolvableTypeRefs: make(map[qualifiedTypeName]winmd.TypeRef),
		projectedTypeDefs:    make(map[winmd.Index]bool),
		requiredTypeDefs:     make(map[winmd.Index]bool),
		abiLayoutTypeDefs:    make(map[winmd.Index]bool),

		methodDefImplMap:              make(map[winmd.Index]winmd.ImplMap),
		fieldConstant:                 make(map[winmd.Index]winmd.Constant),
		typeDefNativeTypedefAttribute: make(map[winmd.Index]winmd.CustomAttribute),
		typeDefSupportedArch:          make(map[winmd.Index]Arch),
		methodDefSupportedArch:        make(map[winmd.Index]Arch),
		paramSizeIndex:                make(map[winmd.Index]paramSizeInfo),
		fieldOffset:                   make(map[winmd.Index]uint32),
		classLayout:                   make(map[winmd.Index]winmd.ClassLayout),
		nestedTypeDefChildren:         make(map[winmd.Index][]winmd.Index),
	}
	// We index tables and resolved defs with the assumption that there is exactly one module. For
	// winmd files, we expect this to always be the case.
	if f.Tables.Module.Len() != 1 {
		return nil, fmt.Errorf("expected exactly one module in the file, but found %v", f.Tables.Module.Len())
	}
	if err := l.indexTypeNames(); err != nil {
		return nil, err
	}
	for idx := range f.Tables.ImplMap.Indices() {
		im, err := f.Tables.ImplMap.At(idx)
		if err != nil {
			return nil, err
		}
		if im.MemberForwarded.Tag != winmd.MemberForwarded_MethodDef {
			continue
		}
		if existing, ok := l.methodDefImplMap[im.MemberForwarded.Index]; ok {
			return nil, fmt.Errorf(
				"multiple ImplMap rows found pointing at MethodDef %v: found %v; already found %v",
				im.MemberForwarded.Index,
				idx,
				existing)
		}
		l.methodDefImplMap[im.MemberForwarded.Index] = im
	}
	for idx := range f.Tables.Constant.Indices() {
		c, err := f.Tables.Constant.At(idx)
		if err != nil {
			return nil, err
		}
		if c.Parent.Tag != winmd.HasConstant_Field {
			continue
		}
		if existing, ok := l.fieldConstant[c.Parent.Index]; ok {
			return nil, fmt.Errorf(
				"multiple Constant rows found pointing at Field %v: found %v; already found %v",
				c.Parent.Index,
				idx,
				existing)
		}
		l.fieldConstant[c.Parent.Index] = c
	}
	attributeDecoder := winmd.NewCustomAttributeDecoder(f)
	for idx := range f.Tables.CustomAttribute.Indices() {
		a, err := f.Tables.CustomAttribute.At(idx)
		if err != nil {
			return nil, err
		}
		if a.Type.Tag != winmd.CustomAttributeType_MemberRef {
			continue
		}
		m, err := f.Tables.MemberRef.At(a.Type.Index)
		if err != nil {
			return nil, err
		}
		if m.Class.Tag != winmd.MemberRefParent_TypeRef {
			continue
		}
		c, err := f.Tables.TypeRef.At(m.Class.Index)
		if err != nil {
			return nil, err
		}
		switch c.Namespace.String() {
		case "Windows.Win32.Foundation.Metadata", "Windows.Win32.Interop": // The former is legacy
			switch c.Name.String() {
			case "NativeArrayInfoAttribute", "MemorySizeAttribute":
				if a.Parent.Tag != winmd.HasCustomAttribute_Param {
					break
				}
				fieldName := "CountParamIndex"
				if c.Name.String() == "MemorySizeAttribute" {
					fieldName = "BytesParamIndex"
				}
				value, err := attributeDecoder.Decode(a)
				if err != nil {
					return nil, fmt.Errorf("decode %s on parameter %v: %w", c.Name, a.Parent.Index, err)
				}
				index, ok, err := int16AttributeField(value, fieldName)
				if err != nil {
					return nil, fmt.Errorf("decode %s on parameter %v: %w", c.Name, a.Parent.Index, err)
				}
				if ok && index >= 0 {
					l.paramSizeIndex[a.Parent.Index] = paramSizeInfo{
						index:   uint16(index),
						inBytes: fieldName == "BytesParamIndex",
					}
				}
			case "NativeTypedefAttribute":
				if a.Parent.Tag != winmd.HasCustomAttribute_TypeDef {
					break
				}
				if existing, ok := l.typeDefNativeTypedefAttribute[a.Parent.Index]; ok {
					return nil, fmt.Errorf(
						"multiple NativeTypedefAttribute rows found pointing at TypeDef %v: found %v; already found %v",
						a.Parent.Index,
						idx,
						existing)
				}
				l.typeDefNativeTypedefAttribute[a.Parent.Index] = a
			case "SupportedArchitectureAttribute":
				if a.Parent.Tag != winmd.HasCustomAttribute_MethodDef && a.Parent.Tag != winmd.HasCustomAttribute_TypeDef {
					break
				}
				value, err := attributeDecoder.Decode(a)
				if err != nil {
					return nil, fmt.Errorf("decode %s on %v: %w", c.Name, a.Parent, err)
				}
				arch, err := supportedArchitecture(value)
				if err != nil {
					return nil, fmt.Errorf("decode %s on %v: %w", c.Name, a.Parent, err)
				}
				switch a.Parent.Tag {
				case winmd.HasCustomAttribute_MethodDef:
					if existing, ok := l.methodDefSupportedArch[a.Parent.Index]; ok {
						return nil, fmt.Errorf(
							"multiple SupportedArchitectureAttribute rows found pointing at MethodDef %v: found %v; already found %v",
							a.Parent.Index,
							idx,
							existing)
					}
					l.methodDefSupportedArch[a.Parent.Index] = arch
				case winmd.HasCustomAttribute_TypeDef:
					if existing, ok := l.typeDefSupportedArch[a.Parent.Index]; ok {
						return nil, fmt.Errorf(
							"multiple SupportedArchitectureAttribute rows found pointing at TypeDef %v: found %v; already found %v",
							a.Parent.Index,
							idx,
							existing)
					}
					l.typeDefSupportedArch[a.Parent.Index] = arch
				}
			}
		}
	}
	for idx := range f.Tables.FieldLayout.Indices() {
		layout, err := f.Tables.FieldLayout.At(idx)
		if err != nil {
			return nil, err
		}
		l.fieldOffset[layout.Field] = layout.Offset
	}
	for idx := range f.Tables.ClassLayout.Indices() {
		layout, err := f.Tables.ClassLayout.At(idx)
		if err != nil {
			return nil, err
		}
		if existing, ok := l.classLayout[layout.Parent]; ok {
			return nil, fmt.Errorf("multiple ClassLayout rows found for TypeDef %v: found %v; already found %#v", layout.Parent, idx, existing)
		}
		l.classLayout[layout.Parent] = layout
	}
	for idx := range f.Tables.NestedClass.Indices() {
		nest, err := f.Tables.NestedClass.At(idx)
		if err != nil {
			return nil, err
		}
		l.nestedTypeDefChildren[nest.EnclosingClass] = append(l.nestedTypeDefChildren[nest.EnclosingClass], nest.NestedClass)
	}
	return l, nil
}

func int16AttributeField(value winmd.CustomAttributeValue, wanted string) (int16, bool, error) {
	for _, argument := range value.NamedArguments {
		if argument.Name != wanted {
			continue
		}
		index, ok := argument.Value.(int16)
		if !ok {
			return 0, false, fmt.Errorf("field %s is not an int16", wanted)
		}
		return index, true, nil
	}
	return 0, false, nil
}

func supportedArchitecture(value winmd.CustomAttributeValue) (Arch, error) {
	if len(value.FixedArguments) != 1 {
		return ArchNone, errors.New("expected one architecture argument")
	}
	arch, ok := value.FixedArguments[0].Value.(int32)
	if !ok {
		return ArchNone, errors.New("architecture argument is not an int32")
	}
	return Arch(arch), nil
}

// MethodDefSupportedArch returns the set of architectures that the given method is supported on.
func (c *Context) MethodDefSupportedArch(idx winmd.Index) Arch {
	v, ok := c.methodDefSupportedArch[idx]
	if !ok {
		v = ArchAll
	}
	return v
}

// UnresolvableTypeRefs returns the qualified names (namespace::name) of TypeRefs that were
// referenced during generation but could not be resolved to a TypeDef in the current module.
func (c *Context) UnresolvableTypeRefs() []string {
	result := make([]string, 0, len(c.unresolvableTypeRefs))
	for _, r := range c.unresolvableTypeRefs {
		result = append(result, r.Namespace.String()+"::"+r.Name.String())
	}
	return result
}

// MethodModuleName returns the lowercase DLL module name for a MethodDef via its ImplMap entry.
// Returns an empty string if the method has no ImplMap entry.
func (c *Context) MethodModuleName(methodIndex winmd.Index) string {
	implMap, ok := c.methodDefImplMap[methodIndex]
	if !ok {
		return ""
	}
	mr, err := c.Metadata.Tables.ModuleRef.At(implMap.ImportScope)
	if err != nil {
		return ""
	}
	return strings.ToLower(mr.Name.String())
}

// TypeDefSupportedArch returns the set of architectures that the given type is supported on.
func (c *Context) TypeDefSupportedArch(idx winmd.Index) Arch {
	v, ok := c.typeDefSupportedArch[idx]
	if !ok {
		v = ArchAll
	}
	return v
}

type qualifiedTypeName struct {
	Namespace string
	Name      string
}

// SelectTypeDef marks a module-level TypeDef and its architecture variants for emission.
func (c *Context) SelectTypeDef(namespace, name, goName string) error {
	key := qualifiedTypeName{Namespace: namespace, Name: name}
	indices := c.typeDefsByName[key]
	qualifiedName := namespace + "." + name
	if len(indices) == 0 {
		return fmt.Errorf("unknown WinMD type %q", qualifiedName)
	}
	if goName != "" && (!token.IsIdentifier(goName) || goName == "_") {
		return fmt.Errorf("invalid Go type name %q", goName)
	}

	for i, left := range indices {
		leftArch := c.TypeDefSupportedArch(left)
		for _, right := range indices[i+1:] {
			overlap := leftArch & c.TypeDefSupportedArch(right)
			if overlap != 0 {
				return fmt.Errorf("ambiguous WinMD type %q: multiple definitions apply to architecture %s", qualifiedName, overlap.String())
			}
		}
	}

	for _, index := range indices {
		def, err := c.resolveTypeDef(index)
		if err != nil {
			return fmt.Errorf("resolve WinMD type %q: %w", qualifiedName, err)
		}
		if goName != "" {
			def.GoName = goName
		}
		c.requiredTypeDefs[index] = true
		c.abiLayoutTypeDefs[index] = true
	}
	return nil
}

// WriteMethod writes to w the signature for "method" in x/sys/windows/mkwinsyscall format.
// Uses the parsed metadata to determine the meaning of data inside the given method.
//
// methodIndex and method must match. Only methodIndex is necessary, but method is also accepted to
// avoid unnecessary decoding. The caller has likely already decoded the method to filter by name.
//
// arch is the architecture that the method is being generated for, or ArchAll if the method is
// supported on all architectures.
//
// goNameOverride, if non-empty, replaces the default Go function name.
func (c *Context) WriteMethod(w io.StringWriter, methodIndex winmd.Index, method winmd.MethodDef, arch Arch, goNameOverride string) error {
	return c.WriteMethodWithOptions(w, methodIndex, method, arch, MethodOptions{GoName: goNameOverride})
}

// WriteMethodWithOptions writes a signature with the requested projection options.
func (c *Context) WriteMethodWithOptions(w io.StringWriter, methodIndex winmd.Index, method winmd.MethodDef, arch Arch, options MethodOptions) error {
	goName := method.Name.String()
	if options.GoName != "" {
		goName = options.GoName
		if !token.IsIdentifier(goName) {
			return fmt.Errorf("invalid Go function name %q", goName)
		}
	} else if !token.IsIdentifier(goName) {
		goName = escapedUpper(goName)
	}

	// Validate the native entry point before writing output or resolving types.
	// Bulk callers can skip unsupported ordinal imports without retaining
	// dependencies or partially generated declarations for the skipped method.
	entryPoint := method.Name.String()
	var moduleName string
	var lastErr bool
	if implMap, ok := c.methodDefImplMap[methodIndex]; ok {
		entryPoint = implMap.ImportName.String()
		if entryPoint == "" {
			return fmt.Errorf("missing native entry point for method %s", method.Name)
		}
		if strings.HasPrefix(entryPoint, "#") {
			return fmt.Errorf("%w %q for method %s", ErrOrdinalImport, entryPoint, method.Name)
		}
		if implMap.MappingFlags&winmd.PInvokeAttributes_SupportsLastError != 0 {
			lastErr = true
		}
		mr, err := c.Metadata.Tables.ModuleRef.At(implMap.ImportScope)
		if err != nil {
			return err
		}
		moduleName = mkwinsyscallModuleName(mr.Name.String())
		if moduleName == "kernel32" {
			moduleName = ""
		}
	}

	w.WriteString("//sys\t")
	w.WriteString(goName)
	w.WriteString("(")

	sig, err := c.Metadata.MethodDefSignature(method.Signature)
	if err != nil {
		return err
	}

	params := make([]winmd.Param, len(sig.Param))
	paramRows := make([]winmd.Index, len(sig.Param))
	for paramRowIndex := range method.ParamList.All() {
		param, err := c.Metadata.Tables.Param.At(paramRowIndex)
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
		if int(i) >= len(sig.Param) {
			return fmt.Errorf("param record Sequence value %v is out of range of parsed signature params, length %v", i, len(sig.Param))
		}
		params[i] = param
		paramRows[i] = paramRowIndex
	}

	wroteParam := false
	skipNext := false
	for i, param := range params {
		if skipNext {
			skipNext = false
			continue
		}
		if wroteParam {
			w.WriteString(", ")
		}
		wroteParam = true
		w.WriteString(escapeParam(param.Name.String()))
		w.WriteString(" ")

		projectSlice := false
		var size paramSizeInfo
		if options.Projection == ProjectionIdiomatic && i+1 < len(params) {
			var ok bool
			size, ok = c.paramSizeIndex[paramRows[i]]
			projectSlice = ok && int(size.index) == i+1
		}
		var err error
		if projectSlice {
			skipNext, err = c.writeSliceType(w, &sig.Param[i].Type, arch, size.inBytes)
		} else if options.Projection == ProjectionIdiomatic {
			err = c.writeProjectedType(w, &sig.Param[i].Type, arch)
		} else {
			err = c.writeType(w, &sig.Param[i].Type, arch)
		}
		if err != nil {
			return fmt.Errorf("failed to interpret type of param %v of method %v: %w", i, method.Name, err)
		}
	}
	w.WriteString(")")

	// Find and write return value(s), if they exist.
	if value := sig.RetType.Kind != winmd.SigRetTypeKind_Void; value || lastErr {
		w.WriteString(" (")
		if value {
			if options.Projection == ProjectionIdiomatic && c.isTypeNamed(&sig.RetType.Type, arch, "NTSTATUS") {
				w.WriteString("ntstatus error")
				value = false
			}
			// General return value name, because mkwinsyscall needs one.
			// Generated returns in general could be better. See https://github.com/microsoft/go-winmd/issues/12
			if value {
				w.WriteString("r ")
				if err := c.writeType(w, &sig.RetType.Type, arch); err != nil {
					return err
				}
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
	if goName != entryPoint || moduleName != "" {
		w.WriteString(" = ")
		if moduleName != "" {
			w.WriteString(moduleName)
			w.WriteString(".")
		}
		w.WriteString(entryPoint)
	}
	return nil
}

func mkwinsyscallModuleName(moduleName string) string {
	return strings.TrimSuffix(strings.ToLower(moduleName), ".dll")
}

func (c *Context) writeProjectedType(w io.StringWriter, p *winmd.SigType, arch Arch) error {
	var rendered strings.Builder
	c.projectingType = true
	err := c.writeType(&rendered, p, arch)
	c.projectingType = false
	if err != nil {
		return err
	}
	typeName := rendered.String()
	if strings.HasPrefix(typeName, "*") && strings.HasSuffix(typeName, "Element") {
		for _, def := range c.resolvedDefsByIndex {
			if def.NativePointer && typeName == "*"+def.GoName {
				c.projectedTypeDefs[def.Index] = true
				var native strings.Builder
				if err := c.writeTypeDefNativeUnderlying(&native, def, arch); err != nil {
					return err
				}
				w.WriteString("*")
				w.WriteString(native.String())
				return nil
			}
		}
	}
	w.WriteString(typeName)
	return nil
}

// writeSliceType reports whether the following count parameter was absorbed.
// mkwinsyscall expands slices using len, which is an element count, not a byte size.
func (c *Context) writeSliceType(w io.StringWriter, p *winmd.SigType, arch Arch, sizeInBytes bool) (bool, error) {
	var rendered strings.Builder
	if err := c.writeProjectedType(&rendered, p, arch); err != nil {
		return false, err
	}
	typeName := rendered.String()
	if sizeInBytes && typeName == "unsafe.Pointer" {
		w.WriteString(typeName)
		return false, nil
	}
	if !strings.HasPrefix(typeName, "*") {
		return false, fmt.Errorf("array metadata applied to non-pointer type %s", typeName)
	}
	elementType := strings.TrimPrefix(typeName, "*")
	if sizeInBytes && elementType != "uint8" && elementType != "int8" && elementType != "bool" {
		// Retain the explicit byte count unless the emitted element is known
		// to occupy one byte. This also avoids guessing sizes of named types.
		w.WriteString(typeName)
		return false, nil
	}
	w.WriteString("[]")
	if elementType == "uint8" {
		elementType = "byte"
	}
	w.WriteString(elementType)
	return true, nil
}

func (c *Context) isTypeNamed(p *winmd.SigType, _ Arch, name string) bool {
	v, ok := p.Value.(winmd.CodedIndex[winmd.TypeDefOrRefOrSpec])
	if !ok {
		return false
	}
	switch v.Tag {
	case winmd.TypeDefOrRefOrSpec_TypeDef:
		def, err := c.Metadata.Tables.TypeDef.At(v.Index)
		return err == nil && def.Name.String() == name
	case winmd.TypeDefOrRefOrSpec_TypeRef:
		ref, err := c.Metadata.Tables.TypeRef.At(v.Index)
		return err == nil && ref.Name.String() == name
	}
	return false
}

func (c *Context) writeType(w io.StringWriter, p *winmd.SigType, arch Arch) error {
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
		if p.Kind == winmd.ElementType_PTR {
			if t, ok := p.Value.(winmd.SigType); ok {
				if t.Kind == winmd.ElementType_VOID {
					w.WriteString("unsafe.Pointer")
					return nil
				}
			}
		}

		switch p.Kind {
		// Translate ECMA-335 primitive types to Go types.
		case winmd.ElementType_BOOLEAN:
			w.WriteString("bool")
		case winmd.ElementType_I1:
			w.WriteString("int8")
		case winmd.ElementType_U1:
			w.WriteString("uint8")
		case winmd.ElementType_I2:
			w.WriteString("int16")
		case winmd.ElementType_U2, winmd.ElementType_CHAR:
			w.WriteString("uint16")
		case winmd.ElementType_I4:
			w.WriteString("int32")
		case winmd.ElementType_U4:
			w.WriteString("uint32")
		case winmd.ElementType_I8:
			w.WriteString("int64")
		case winmd.ElementType_U8:
			w.WriteString("uint64")
		case winmd.ElementType_R4:
			w.WriteString("float32")
		case winmd.ElementType_R8:
			w.WriteString("float64")

		// ECMA-335 distinguishes uintptr and intptr, Go only has uintptr used in both cases.
		case winmd.ElementType_I, winmd.ElementType_U:
			w.WriteString("uintptr")

		case winmd.ElementType_VOID:
			// We catch "*void" with a special case above. We should never see simply VOID.
			return errors.New("unexpected primitive type: VOID")

		case winmd.ElementType_OBJECT:
			w.WriteString("any")

		case winmd.ElementType_SZARRAY, winmd.ElementType_GENERICINST,
			winmd.ElementType_VAR, winmd.ElementType_MVAR:
			// These metadata shapes need a managed-to-native projection. Do not
			// recurse through their Value and discard the enclosing type.
			return fmt.Errorf("unsupported type for Go generation: %v", p.Kind)

		// If this is not a simple value type, there will be p.Value. Handle all those cases here.
		default:
			if p.Kind == winmd.ElementType_PTR || p.Kind == winmd.ElementType_BYREF {
				w.WriteString("*")
			}
			switch v := p.Value.(type) {
			case winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]:
				switch v.Tag {
				case winmd.TypeDefOrRefOrSpec_TypeDef:
					def, err := c.resolveTypeDef(v.Index)
					if err != nil {
						return err
					}
					if def.NeedsPointerWhenUsed() {
						w.WriteString("*")
					}
					if !c.projectingType {
						c.requiredTypeDefs[def.Index] = true
					}
					if c.writingABIType {
						c.abiLayoutTypeDefs[def.Index] = true
					}
					w.WriteString(def.GoName)
				case winmd.TypeDefOrRefOrSpec_TypeRef:
					def, err := c.resolveTypeRef(v.Index, arch)
					if err != nil && !errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
						return err
					}
					if def == nil {
						ref, err := c.Metadata.Tables.TypeRef.At(v.Index)
						if err != nil {
							return err
						}
						w.WriteString(ref.Name.String())
						key := qualifiedTypeName{Namespace: ref.Namespace.String(), Name: ref.Name.String()}
						c.unresolvableTypeRefs[key] = ref
					} else {
						if def.NeedsPointerWhenUsed() {
							w.WriteString("*")
						}
						if !c.projectingType {
							c.requiredTypeDefs[def.Index] = true
						}
						if c.writingABIType {
							c.abiLayoutTypeDefs[def.Index] = true
						}
						w.WriteString(def.GoName)
					}
				default:
					return fmt.Errorf("unexpected coded index tag for type Value: %#v", v)
				}

				// Types can nest. A pointer to another type is a very common case.
			case winmd.SigType:
				markVisited(p)
				return visitType(&v)
			case winmd.SigArray:
				for i := range int(v.Rank) {
					if i < len(v.Sizes) {
						w.WriteString("[" + strconv.Itoa(int(v.Sizes[i])) + "]")
					} else {
						w.WriteString("[]")
					}
				}
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

	Arch Arch

	def winmd.TypeDef
}

func (r *resolvedDef) IsInterface() bool {
	return r.def.Flags&winmd.TypeAttributes_ClassSemanticsMask == winmd.TypeAttributes_Interface
}

func (r *resolvedDef) NeedsPointerWhenUsed() bool {
	return r.NativePointer || r.IsInterface()
}

var errTypeDefNotDefinedInCurrentModule = errors.New("TypeRef points to TypeDef not defined in this module")

// resolveTypeRef resolves a TypeRef to a TypeDef, if possible. If the TypeRef is not defined in
// this module, it returns nil, errTypeDefNotDefinedInCurrentModule.
// If refIndex can be resolved to multiple TypeDefs, then the behavior depends on arch:
//   - If arch is ArchAll, then all TypeDefs are resolved and the first one is returned.
//   - If arch is not ArchAll, then only the one matching arch is resolved and returned.
func (c *Context) resolveTypeRef(refIndex winmd.Index, arch Arch) (*resolvedDef, error) {
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

		r, err := c.Metadata.Tables.TypeRef.At(refIndex)
		if err != nil {
			return nil, err
		}
		switch r.ResolutionScope.Tag {
		// We assume module-level TypeRefs refer to the current module.
		case winmd.ResolutionScope_Module:
			if def := c.typeDefCache.get(r.Namespace, r.Name, arch); def != nil {
				return def, nil
			}
			key := c.typeDefCache.canonicalKey(typeRefKey(r))
			if defIndex, ok := c.typeDefCache.unresolved[key]; ok {
				if arch == ArchAll || c.TypeDefSupportedArch(defIndex)&arch == arch {
					return c.resolveTypeDef(defIndex)
				}
				break
			}
			defIndices := c.typeDefCache.unresolvedDuplicated[key]
			if len(defIndices) == 0 {
				// Only an offset miss needs a textual lookup. Memoize an alias
				// so subsequent references use the integer-keyed fast path.
				name := qualifiedTypeName{Namespace: r.Namespace.String(), Name: r.Name.String()}
				defIndices = c.typeDefsByName[name]
				if len(defIndices) != 0 {
					first, err := c.Metadata.Tables.TypeDef.At(defIndices[0])
					if err != nil {
						return nil, err
					}
					c.typeDefCache.addAlias(typeRefKey(r), typeDefKey(first))
				}
			}
			if len(defIndices) != 0 {
				var archDef *resolvedDef
				for _, defIndex := range defIndices {
					if arch == ArchAll || c.TypeDefSupportedArch(defIndex)&arch == arch {
						rdef, err := c.resolveTypeDef(defIndex)
						if err != nil {
							return nil, err
						}
						if archDef == nil {
							archDef = rdef
						}
					}

				}
				if archDef != nil {
					return archDef, nil
				}
			}
		// TypeRef scope indicates that this is a nested type. The Index is the immediate parent.
		case winmd.ResolutionScope_TypeRef:
			// Recurse to find the parent def.
			parentDefIndex, err := visit(r.ResolutionScope.Index)
			if err != nil {
				return nil, err
			}
			// Select nested variants by name and architecture, never by the
			// physical heap offset of an equal string.
			for _, child := range parentDefIndex.Children {
				if child.Name.String() == r.Name.String() && (arch == ArchAll || child.Arch&arch == arch) {
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
		def, err := c.Metadata.Tables.TypeDef.At(defIndex)
		if err != nil {
			return nil, err
		}
		r := resolvedDef{
			Index:     defIndex,
			Namespace: def.Namespace,
			Name:      def.Name,
			GoName:    escapedUpper(def.Name.String()),
			Arch:      c.TypeDefSupportedArch(defIndex),
			def:       def,
		}
		if _, ok := c.typeDefNativeTypedefAttribute[defIndex]; ok {
			r.Native = true
			if def.FieldList.Start+1 != def.FieldList.End {
				return nil, fmt.Errorf("expected exactly one field for native typedef %v", r.def.Name)
			}
			fd, err := c.Metadata.Tables.Field.At(r.def.FieldList.Start)
			if err != nil {
				return nil, err
			}
			signature, err := c.Metadata.FieldSignature(fd.Signature)
			if err != nil {
				return nil, err
			}
			if signature.Type.Kind == winmd.ElementType_PTR {
				to := signature.Type.Value.(winmd.SigType)
				if to.Kind != winmd.ElementType_VOID {
					r.NativePointer = true
					// Clarify the Go name, to avoid confusion for anyone with a strong expectation
					// about types like PWSTR.
					r.GoName += "Element"
				}
			}
		}
		// Nested types can't be resolved at module scope. Don't add it to the module lookup.
		if def.Flags&winmd.TypeAttributes_VisibilityMask <= winmd.TypeAttributes_Public {
			c.typeDefCache.resolve(&r)
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

func (c *Context) writeTypeDef(w io.StringWriter, r *resolvedDef, arch Arch) error {
	previousWritingABIType := c.writingABIType
	c.writingABIType = c.abiLayoutTypeDefs[r.Index]
	defer func() { c.writingABIType = previousWritingABIType }()

	if r.IsInterface() {
		// Issue tracking implementing interface types: https://github.com/microsoft/go-winmd/issues/14
		w.WriteString("// Interface type is likely missing members. Not yet implemented in go-winmd.\n")
	}
	switch r.def.Extends.Tag {
	case winmd.TypeDefOrRef_TypeDef, winmd.TypeDefOrRef_TypeRef:
		base, err := systemBaseType(c.Metadata, r)
		if err != nil {
			return err
		}
		switch base {
		case "Enum":
			return c.writeTypeDefEnum(w, r, arch)
		case "MulticastDelegate":
			w.WriteString("type ")
			w.WriteString(r.GoName)
			w.WriteString(" uintptr\n")
			return nil
		}
		if r.Native {
			return c.writeTypeDefNative(w, r, arch)
		}
		return c.writeTypeDefStruct(w, r, arch)
	case winmd.TypeDefOrRef_Null:
		return c.writeTypeDefStruct(w, r, arch)
	default:
		return fmt.Errorf("unexpected type extends coded index %#v in def %v :: %v", r.def.Extends, r.def.Namespace, r.def.Name)
	}
}

// systemBaseType returns the name of a non-nested base type in System, if any.
func systemBaseType(metadata *winmd.Metadata, def *resolvedDef) (string, error) {
	var namespace, name winmd.String
	switch def.def.Extends.Tag {
	case winmd.TypeDefOrRef_TypeDef:
		base, err := metadata.Tables.TypeDef.At(def.def.Extends.Index)
		if err != nil {
			return "", err
		}
		if base.Flags&winmd.TypeAttributes_VisibilityMask > winmd.TypeAttributes_Public {
			return "", nil
		}
		namespace, name = base.Namespace, base.Name
	case winmd.TypeDefOrRef_TypeRef:
		base, err := metadata.Tables.TypeRef.At(def.def.Extends.Index)
		if err != nil {
			return "", err
		}
		if base.ResolutionScope.Tag == winmd.ResolutionScope_TypeRef {
			return "", nil
		}
		namespace, name = base.Namespace, base.Name
	default:
		return "", nil
	}
	if namespace.String() != "System" {
		return "", nil
	}
	return name.String(), nil
}

func (c *Context) writeTypeDefEnum(w io.StringWriter, r *resolvedDef, arch Arch) error {
	underlyingType, err := c.Metadata.EnumUnderlyingType(r.Index)
	if err != nil {
		return err
	}
	w.WriteString("type ")
	w.WriteString(r.GoName)

	type member struct {
		Name     winmd.String
		HexValue string
	}
	// All fields except the single instance field represent enum members.
	members := make([]member, 0, r.def.FieldList.Len()-1)

	for i := range r.def.FieldList.All() {
		fd, err := c.Metadata.Tables.Field.At(i)
		if err != nil {
			return err
		}
		if fd.Flags&winmd.FieldAttributes_Static == 0 {
			continue
		}

		constant, ok := c.fieldConstant[i]
		if !ok {
			return fmt.Errorf("unable to find default value for field %v", fd.Name)
		}
		hex, err := formatEnumConstant(constant)
		if err != nil {
			return fmt.Errorf("field %v: %w", fd.Name, err)
		}

		// Don't write the members yet. We haven't written the enum type definition yet, and the
		// order is important for readability in docs.
		p := member{fd.Name, hex}
		members = append(members, p)
	}

	w.WriteString(" ")
	if err := c.writeType(w, &winmd.SigType{Kind: underlyingType}, arch); err != nil {
		return err
	}
	w.WriteString("\n\nconst (\n")
	for _, pair := range members {
		name := pair.Name.String()
		w.WriteString("\t")
		w.WriteString(escapedUpper(name))
		w.WriteString(" ")
		w.WriteString(r.GoName)
		w.WriteString(" = ")
		w.WriteString(pair.HexValue)
		w.WriteString("\n")
	}
	w.WriteString(")\n")
	return nil
}

func formatEnumConstant(constant winmd.Constant) (string, error) {
	switch constant.Type {
	case winmd.ElementType_BOOLEAN, winmd.ElementType_CHAR,
		winmd.ElementType_I1, winmd.ElementType_I2, winmd.ElementType_I4, winmd.ElementType_I8,
		winmd.ElementType_U1, winmd.ElementType_U2, winmd.ElementType_U4, winmd.ElementType_U8:
	default:
		return "", fmt.Errorf("enum member has unexpected type: %v", constant.Type)
	}
	value, err := constant.DecodeValue()
	if err != nil {
		return "", err
	}
	if boolean, ok := value.(bool); ok {
		return strconv.FormatBool(boolean), nil
	}
	return fmt.Sprintf("%#x", value), nil
}

func (c *Context) writeTypeDefNative(w io.StringWriter, r *resolvedDef, arch Arch) error {
	w.WriteString("type ")
	w.WriteString(r.GoName)
	w.WriteString(" ")
	if err := c.writeTypeDefNativeUnderlying(w, r, arch); err != nil {
		return err
	}
	w.WriteString("\n")
	return nil
}

func (c *Context) writeTypeDefNativeUnderlying(w io.StringWriter, r *resolvedDef, arch Arch) error {
	if r.def.FieldList.Start+1 != r.def.FieldList.End {
		return fmt.Errorf("expected exactly one field for native typedef %v", r.def.Name)
	}
	fd, err := c.Metadata.Tables.Field.At(r.def.FieldList.Start)
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
	if err := c.writeType(w, &signature.Type, arch); err != nil {
		return err
	}
	return nil
}

func (c *Context) writeTypeDefStruct(w io.StringWriter, r *resolvedDef, arch Arch) error {
	w.WriteString("type ")
	w.WriteString(r.GoName)
	w.WriteString(" struct {\n")
	if err := c.writeStructFields(w, r, arch); err != nil {
		return err
	}
	w.WriteString("}\n")
	return nil
}

func (c *Context) writeStructFields(w io.StringWriter, r *resolvedDef, arch Arch) error {
	if c.abiLayoutTypeDefs[r.Index] && (arch == Arch386 || arch == ArchAMD64 || arch == ArchARM64) {
		layout, err := c.planStructABI(r, arch, nil)
		if err != nil {
			return err
		}
		for _, field := range layout.fields {
			if field.padding != 0 {
				w.WriteString("\t_ [" + strconv.Itoa(int(field.padding)) + "]byte\n")
			}
			if err := c.writeStructField(w, field.index, arch); err != nil {
				return err
			}
		}
		if layout.tailPadding != 0 {
			w.WriteString("\t_ [" + strconv.Itoa(int(layout.tailPadding)) + "]byte\n")
		}
		return nil
	}

	// Union type support is simple for now. Roughly follow the x/sys approach and pick one union
	// option to implement. See the x/sys/windows "IpAdapterAddresses" struct in syscall_windows.go
	// for an example. Better support is tracked at https://github.com/microsoft/go-winmd/issues/17

	// Write a definition for each field. Take FieldOffset into account, but minimally, for
	// simplicity: once a FieldOffset is "used", don't write another field that uses the same
	// explicit field offset. This isn't accurate, but it may be good enough. In many cases, only
	// "0" is used.
	usedFieldOffset := make(map[uint32]struct{})
	for i := range r.def.FieldList.All() {
		if o, ok := c.fieldOffset[i]; ok {
			if _, used := usedFieldOffset[o]; used {
				continue
			}
			usedFieldOffset[o] = struct{}{}
		}
		if err := c.writeStructField(w, i, arch); err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) writeStructField(w io.StringWriter, fieldIndex winmd.Index, arch Arch) error {
	fd, err := c.Metadata.Tables.Field.At(fieldIndex)
	if err != nil {
		return err
	}
	signature, err := c.Metadata.FieldSignature(fd.Signature)
	if err != nil {
		return err
	}
	if signature.Type.Kind == winmd.ElementType_VALUETYPE {
		if v, ok := signature.Type.Value.(winmd.CodedIndex[winmd.TypeDefOrRefOrSpec]); ok && v.Tag == winmd.TypeDefOrRefOrSpec_TypeRef {
			ref, err := c.Metadata.Tables.TypeRef.At(v.Index)
			if err != nil {
				return err
			}
			if ref.ResolutionScope.Tag == winmd.ResolutionScope_TypeRef {
				// Nested types are emitted inline rather than as standalone declarations.
				def, err := c.resolveTypeRef(v.Index, arch)
				if err != nil && !errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
					return err
				}
				if def != nil {
					if c.writingABIType {
						// Keep the field's struct boundary: flattening would discard
						// its alignment and implicit tail padding from the ABI plan.
						w.WriteString("\t" + escapedUpper(fd.Name.String()) + " struct {\n")
						if err := c.writeStructFields(w, def, arch); err != nil {
							return err
						}
						w.WriteString("\t}\n")
						return nil
					}
					// Preserve the legacy flattened representation for types that
					// have not opted into architecture-specific ABI layout.
					return c.writeStructFields(w, def, arch)
				}
				// Fall through to write the field as-is if the nested type can't be resolved.
			}
		}
	}
	w.WriteString("\t")
	w.WriteString(escapedUpper(fd.Name.String()))
	w.WriteString(" ")
	if err := c.writeType(w, &signature.Type, arch); err != nil {
		return err
	}
	w.WriteString("\n")
	return nil
}

// WriteUsedTypeDefs writes Go definitions for TypeDefs that were discovered during WriteMethod
// calls or explicitly selected with SelectTypeDef. For a given Context c, only call this method one
// time, after all WriteMethod and SelectTypeDef calls are complete.
func (c *Context) WriteUsedTypeDefs(b map[Arch]*strings.Builder) error {
	if err := c.discoverABILayoutDependencies(); err != nil {
		return err
	}
	archSeen := make(map[Arch]bool)
	// Keep going until we stop finding new types that need definitions.
	written := make(map[*resolvedDef]struct{})
	for {
		usedTypeDefs := c.typeDefCache.collect(func(r *resolvedDef) bool {
			if c.projectedTypeDefs[r.Index] && !c.requiredTypeDefs[r.Index] {
				return false
			}
			if _, ok := written[r]; !ok {
				written[r] = struct{}{}
				return true
			}
			return false
		})
		if len(usedTypeDefs) == 0 {
			break
		}
		// Order is scrambled due to the map. Put it back in some order.
		slices.SortFunc(usedTypeDefs, func(a, b *resolvedDef) int {
			return cmp.Compare(a.Index, b.Index)
		})
		for _, r := range usedTypeDefs {
			// Writing the type def (field types in particular) adds new entries to ResolvedDefs if
			// we haven't seen them yet.
			supportedArches := c.TypeDefSupportedArch(r.Index)
			type renderedType struct {
				arch   Arch
				text   string
				layout abiLayoutFingerprint
			}
			var rendered []renderedType
			if c.abiLayoutTypeDefs[r.Index] && supportedArches == ArchAll {
				for _, arch := range []Arch{Arch386, ArchAMD64, ArchARM64} {
					var definition strings.Builder
					if err := c.writeTypeDef(&definition, r, arch); err != nil {
						return err
					}
					layout, err := c.typeDefABILayoutFingerprint(r, arch)
					if err != nil {
						return err
					}
					rendered = append(rendered, renderedType{arch: arch, text: definition.String(), layout: layout})
				}
				if rendered[0].text == rendered[1].text && rendered[0].text == rendered[2].text &&
					equalABILayout(rendered[0].layout, rendered[1].layout) && equalABILayout(rendered[0].layout, rendered[2].layout) {
					rendered = []renderedType{{arch: ArchAll, text: rendered[0].text}}
				}
			} else {
				for _, arch := range supportedArches.Unique() {
					var definition strings.Builder
					if err := c.writeTypeDef(&definition, r, arch); err != nil {
						return err
					}
					rendered = append(rendered, renderedType{arch: arch, text: definition.String()})
				}
			}
			for _, definition := range rendered {
				arch := definition.arch
				w := b[arch]
				if !archSeen[arch] {
					w.WriteString("\n\n// Types used in generated APIs for\n\n")
					archSeen[arch] = true
				}
				w.WriteString(definition.text)
				w.WriteString("\n")
			}
		}
	}
	return nil
}

// escapeParam returns the given string, adding a suffix if it is a reserved Go keyword. Leave the
// case as-is (unlike writeEscapedUpper) because lowercase is desirable for params.
func escapeParam(s string) string {
	if token.IsKeyword(s) {
		s += "Param"
	}
	return s
}

// escapedUpper returns the given string with the first character in uppercase. All Go keywords are
// lowercase, so uppercasing the first letter does two things: escapes names like "type" and exports
// the generated types/fields.
func escapedUpper(s string) string {
	if len(s) > 0 {
		_, size := utf8.DecodeRuneInString(s)
		s = strings.ToUpper(s[:size]) + s[size:]
	}
	return s
}
