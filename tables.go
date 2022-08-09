// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"github.com/microsoft/go-winmd/flags"
)

// Record is an item in a metadata table.
type Record = any

// Index indexes a record of type T.
type Index[T Record] int

// CodedIndex indexes a record on any table.
type CodedIndex struct {
	Index[any]
	Table int
}

// Slice indexes the range of records [Start,End) on the table T.
type Slice[T Record] struct {
	Start Index[T]
	End   Index[T]
}

// BlobIndex is an index to a #Blob heap.
type BlobIndex int

// GuidIndex is an index to a #GUID heap.
type GUIDIndex int

// Assembly is defined in §II.22.2.
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
type assemblyOS struct {
	OSPlatformID   uint32
	OSMajorVersion uint32
	OSMinorVersion uint32
}

// assemblyProcessor is defined in §II.22.4.
// This record should not be emitted into any PE file,
// but needed temporarily to calculate sizes and offsets for subsequent tables.
type assemblyProcessor struct {
	Processor uint32
}

// AssemblyRef is defined in §II.22.5.
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
type assemblyRefOS struct {
	OSPlatformID   uint32
	OSMajorVersion uint32
	OSMinorVersion uint32
	AssemblyRef    Index[AssemblyRef]
}

// assemblyRefProcessor is defined in §II.22.7.
// This record should not be emitted into any PE file,
// but needed temporarily to calculate sizes and offsets for subsequent tables.
type assemblyRefProcessor struct {
	Processor   uint32
	AssemblyRef Index[AssemblyRef]
}

// ClassLayout is defined in §II.22.8.
type ClassLayout struct {
	PackingSize uint16
	ClassSize   uint32
	AssemblyRef Index[TypeDef]
}

// Constant is defined in §II.22.9.
type Constant struct {
	Type    flags.ElementType
	Padding byte       // 1-byte padding zero
	Parent  CodedIndex `winmd:"HasConstant"`
	Value   BlobIndex
}

// CustomAttribute is defined in §II.22.10.
type CustomAttribute struct {
	Parent CodedIndex `winmd:"HasCustomAttribute"`
	Type   CodedIndex `winmd:"CustomAttributeType"`
	Value  BlobIndex
}

// DeclSecurity is defined in §II.22.11.
type DeclSecurity struct {
	Action        uint16
	Parent        CodedIndex `winmd:"HasDeclSecurity"`
	PermissionSet BlobIndex
}

// EventMap is defined in §II.22.12.
type EventMap struct {
	Parent    Index[TypeDef]
	EventList Slice[Event]
}

// Event is defined in §II.22.13.
type Event struct {
	EventFlags flags.EventAttributes
	Name       string
	EventType  CodedIndex `winmd:"TypeDefOrRef"`
}

// ExportedType is defined in §II.22.14.
type ExportedType struct {
	Flags          flags.TypeAttributes
	TypeDefID      uint32
	Name           string
	Namespace      string     `winmd:",cache"`
	Implementation CodedIndex `winmd:"Implementation"`
}

// Field is defined in §II.22.15.
type Field struct {
	Flags     flags.FieldAttributes
	Name      string
	Signature BlobIndex
}

// FieldLayout is defined in §II.22.16.
type FieldLayout struct {
	Offset uint32
	Field  Index[Field]
}

// FieldMarshal is defined in §II.22.17.
type FieldMarshal struct {
	Parent     CodedIndex `winmd:"HasFieldMarshal"`
	NativeType BlobIndex
}

// FieldRVA is defined in §II.22.18.
type FieldRVA struct {
	RVA   uint32
	Field Index[Field]
}

// File is defined in §II.22.19.
type File struct {
	Flags     flags.FileAttributes
	Name      string
	HashValue BlobIndex
}

// GenericParam is defined in §II.22.20.
type GenericParam struct {
	Number uint16
	Flags  flags.GenericParamAttributes
	Owner  CodedIndex `winmd:"TypeOrMethodDef"`
	Name   string
}

// GenericParam is defined in §II.22.21.
type GenericParamConstraint struct {
	Owner      Index[GenericParam]
	Constraint CodedIndex `winmd:"TypeDefOrRef"`
}

// ImplMap is defined in §II.22.22.
type ImplMap struct {
	MappingFlags    flags.PInvokeAttributes
	MemberForwarded CodedIndex `winmd:"MemberForwarded"`
	ImportName      string
	ImportScope     Index[ModuleRef]
}

// InterfaceImpl is defined in §II.22.23.
type InterfaceImpl struct {
	Class     Index[TypeDef]
	Interface CodedIndex `winmd:"TypeDefOrRef"`
}

// ManifestResource is defined in §II.22.24.
type ManifestResource struct {
	Offset         uint32
	Flags          flags.ManifestResourceAttributes
	Name           string
	Implementation CodedIndex `winmd:"Implementation"`
}

// MemberRef is defined in §II.22.25.
type MemberRef struct {
	Class     CodedIndex `winmd:"MemberRefParent"`
	Name      string
	Signature BlobIndex
}

// MethodDef is defined in §II.22.26.
type MethodDef struct {
	RVA       uint32
	ImplFlags flags.MethodImplAttributes
	Flags     flags.MethodAttributes
	Name      string
	Signature BlobIndex
	ParamList Slice[Param]
}

// MethodImpl is defined in §II.22.27.
type MethodImpl struct {
	Class             Index[TypeDef]
	MethodBody        CodedIndex `winmd:"MethodDefOrRef"`
	MethodDeclaration CodedIndex `winmd:"MethodDefOrRef"`
}

// MethodImpl is defined in §II.22.28.
type MethodSemantics struct {
	Semantics   flags.MethodSemanticsAttributes
	Method      Index[MethodDef]
	Association CodedIndex `winmd:"HasSemantics"`
}

// MethodSpec is defined in §II.22.29.
type MethodSpec struct {
	Method        CodedIndex `winmd:"MethodDefOrRef"`
	Instantiation BlobIndex
}

// Module is defined in §II.22.30.
type Module struct {
	Generation uint16
	Name       string
	Mvid       GUIDIndex
	EncID      GUIDIndex
	EncBaseID  GUIDIndex
}

// ModuleRef is defined in §II.22.31.
type ModuleRef struct {
	Name string
}

// NestedClass is defined in §II.22.32.
type NestedClass struct {
	NestedClass    Index[TypeDef]
	EnclosingClass Index[TypeDef]
}

// Param is defined in §II.22.33.
type Param struct {
	Flags    flags.ParamAttributes
	Sequence uint16
	Name     string // String heap
}

// Property is defined in §II.22.34.
type Property struct {
	Flags flags.PropertyAttributes
	Name  string
	Type  BlobIndex
}

// PropertyMap is defined in §II.22.35.
type PropertyMap struct {
	Parent       Index[TypeDef]
	PropertyList Slice[Property]
}

// StandAloneSig is defined in §II.22.36.
type StandAloneSig struct {
	Signature BlobIndex
}

// TypeDef is defined in §II.22.37.
type TypeDef struct {
	Flags      flags.TypeAttributes
	Name       string
	Namespace  string     `winmd:",cache"`
	Extends    CodedIndex `winmd:"TypeDefOrRef"`
	FieldList  Slice[Field]
	MethodList Slice[MethodDef]
}

// TypeRef is defined in §II.22.38.
type TypeRef struct {
	ResolutionScope CodedIndex `winmd:"ResolutionScope"`
	Name            string
	Namespace       string `winmd:",cache"`
}

// TypeSpec is defined in §II.22.39.
type TypeSpec struct {
	Signature BlobIndex
}
