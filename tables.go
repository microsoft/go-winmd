// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
//go:generate genlayout
package winmd

import "github.com/microsoft/go-winmd/flags"

// Assembly is defined in §II.22.2.
// @table=0x20
type Assembly struct {
	HashAlgID      flags.AssemblyHashAlgorithm
	MajorVersion   uint16
	MinorVersion   uint16
	BuildNumber    uint16
	RevisionNumber uint16
	Flags          flags.AssemblyFlags
	PublicKey      BlobIndex
	Name           string
	Culture        string
}

// assemblyOS is defined in §II.22.3.
// This record should not be emitted into any PE file,
// but needed temporarily to calculate sizes and offsets for subsequent tables.
// @table=0x22
type assemblyOS struct {
	OSPlatformID   uint32
	OSMajorVersion uint32
	OSMinorVersion uint32
}

// assemblyProcessor is defined in §II.22.4.
// This record should not be emitted into any PE file,
// but needed temporarily to calculate sizes and offsets for subsequent tables.
// @table=0x21
type assemblyProcessor struct {
	Processor uint32
}

// AssemblyRef is defined in §II.22.5.
// @table=0x23
type AssemblyRef struct {
	MajorVersion     uint16
	MinorVersion     uint16
	BuildNumber      uint16
	RevisionNumber   uint16
	Flags            flags.AssemblyFlags
	PublicKeyOrToken BlobIndex
	Name             string
	Culture          string
	HashValue        BlobIndex
}

// assemblyRefOS is defined in §II.22.6.
// This record should not be emitted into any PE file,
// but needed temporarily to calculate sizes and offsets for subsequent tables.
// @table=0x25
type assemblyRefOS struct {
	OSPlatformID   uint32
	OSMajorVersion uint32
	OSMinorVersion uint32
	AssemblyRef    Index // @ref=AssemblyRef
}

// assemblyRefProcessor is defined in §II.22.7.
// This record should not be emitted into any PE file,
// but needed temporarily to calculate sizes and offsets for subsequent tables.
// @table=0x24
type assemblyRefProcessor struct {
	Processor   uint32
	AssemblyRef Index // @ref=AssemblyRef
}

// ClassLayout is defined in §II.22.8.
// @table=0x0F
type ClassLayout struct {
	PackingSize uint16
	ClassSize   uint32
	Parent      Index // @ref=TypeDef
}

// Constant is defined in §II.22.9.
// @table=0x0B
type Constant struct {
	Type    flags.ElementType
	Padding byte       // 1-byte padding zero
	Parent  CodedIndex // @code=HasConstant
	Value   BlobIndex
}

// CustomAttribute is defined in §II.22.10.
// @table=0x0C
type CustomAttribute struct {
	Parent CodedIndex // @code=HasCustomAttribute
	Type   CodedIndex // @code=CustomAttributeType
	Value  BlobIndex
}

// DeclSecurity is defined in §II.22.11.
// @table=0x0E
type DeclSecurity struct {
	Action        uint16
	Parent        CodedIndex // @code=HasDeclSecurity
	PermissionSet BlobIndex
}

// EventMap is defined in §II.22.12.
// @table=0x12
type EventMap struct {
	Parent    Index // @ref=TypeDef
	EventList Slice // @ref=Event
}

// Event is defined in §II.22.13.
// @table=0x14
type Event struct {
	EventFlags flags.EventAttributes
	Name       string
	EventType  CodedIndex // @code=TypeDefOrRef
}

// ExportedType is defined in §II.22.14.
// @table=0x27
type ExportedType struct {
	Flags          flags.TypeAttributes
	TypeDefID      uint32 // index into a TypeDef table, used as hint only
	Name           string
	Namespace      string
	Implementation CodedIndex // @code=Implementation
}

// Field is defined in §II.22.15.
// @table=0x04
type Field struct {
	Flags     flags.FieldAttributes
	Name      string
	Signature BlobIndex
}

// FieldLayout is defined in §II.22.16.
// @table=0x10
type FieldLayout struct {
	Offset uint32
	Field  Index // @ref=Field
}

// FieldMarshal is defined in §II.22.17.
// @table=0x0D
type FieldMarshal struct {
	Parent     CodedIndex // @code=HasFieldMarshal
	NativeType BlobIndex
}

// FieldRVA is defined in §II.22.18.
// @table=0x1D
type FieldRVA struct {
	RVA   uint32
	Field Index // @ref=Field
}

// File is defined in §II.22.19.
// @table=0x26
type File struct {
	Flags     flags.FileAttributes
	Name      string
	HashValue BlobIndex
}

// GenericParam is defined in §II.22.20.
// @table=0x2A
type GenericParam struct {
	Number uint16
	Flags  flags.GenericParamAttributes
	Owner  CodedIndex // @code=TypeOrMethodDef
	Name   string
}

// GenericParam is defined in §II.22.21.
// @table=0x2C
type GenericParamConstraint struct {
	Owner      Index      // @ref=GenericParam
	Constraint CodedIndex // @code=TypeDefOrRef
}

// ImplMap is defined in §II.22.22.
// @table=0x1C
type ImplMap struct {
	MappingFlags    flags.PInvokeAttributes
	MemberForwarded CodedIndex // @code=MemberForwarded
	ImportName      string
	ImportScope     Index // @ref=ModuleRef
}

// InterfaceImpl is defined in §II.22.23.
// @table=0x09
type InterfaceImpl struct {
	Class     Index      // @ref=TypeDef
	Interface CodedIndex // @code=TypeDefOrRef
}

// ManifestResource is defined in §II.22.24.
// @table=0x28
type ManifestResource struct {
	Offset         uint32
	Flags          flags.ManifestResourceAttributes
	Name           string
	Implementation CodedIndex // @code=Implementation
}

// MemberRef is defined in §II.22.25.
// @table=0x0A
type MemberRef struct {
	Class     CodedIndex // @code=MemberRefParent
	Name      string
	Signature BlobIndex
}

// MethodDef is defined in §II.22.26.
// @table=0x06
type MethodDef struct {
	RVA       uint32
	ImplFlags flags.MethodImplAttributes
	Flags     flags.MethodAttributes
	Name      string
	Signature BlobIndex
	ParamList Slice // @ref=Param
}

// MethodImpl is defined in §II.22.27.
// @table=0x19
type MethodImpl struct {
	Class             Index      // @ref=TypeDef
	MethodBody        CodedIndex // @code=MethodDefOrRef
	MethodDeclaration CodedIndex // @code=MethodDefOrRef
}

// MethodImpl is defined in §II.22.28.
// @table=0x18
type MethodSemantics struct {
	Semantics   flags.MethodSemanticsAttributes
	Method      Index      // @ref=MethodDef
	Association CodedIndex // @code=HasSemantics
}

// MethodSpec is defined in §II.22.29.
// @table=0x2B
type MethodSpec struct {
	Method        CodedIndex // @code=MethodDefOrRef
	Instantiation BlobIndex
}

// Module is defined in §II.22.30.
// @table=0x00
type Module struct {
	Generation uint16
	Name       string
	Mvid       GUIDIndex
	EncID      GUIDIndex
	EncBaseID  GUIDIndex
}

// ModuleRef is defined in §II.22.31.
// @table=0x1A
type ModuleRef struct {
	Name string
}

// NestedClass is defined in §II.22.32.
// @table=0x29
type NestedClass struct {
	NestedClass    Index // @ref=TypeDef
	EnclosingClass Index // @ref=TypeDef
}

// Param is defined in §II.22.33.
// @table=0x08
type Param struct {
	Flags    flags.ParamAttributes
	Sequence uint16
	Name     string
}

// Property is defined in §II.22.34.
// @table=0x17
type Property struct {
	Flags flags.PropertyAttributes
	Name  string
	Type  BlobIndex
}

// PropertyMap is defined in §II.22.35.
// @table=0x15
type PropertyMap struct {
	Parent       Index // @ref=TypeDef
	PropertyList Slice // @ref=Property
}

// StandAloneSig is defined in §II.22.36.
// @table=0x11
type StandAloneSig struct {
	Signature BlobIndex
}

// TypeDef is defined in §II.22.37.
// @table=0x02
type TypeDef struct {
	Flags      flags.TypeAttributes
	Name       string
	Namespace  string
	Extends    CodedIndex // @code=TypeDefOrRef
	FieldList  Index      // @ref=Field
	MethodList Index      // @ref=MethodDef
}

// TypeRef is defined in §II.22.38.
// @table=0x01
type TypeRef struct {
	ResolutionScope CodedIndex // @code=ResolutionScope
	Name            string
	Namespace       string
}

// TypeSpec is defined in §II.22.39.
// @table=0x1B
type TypeSpec struct {
	Signature BlobIndex
}
