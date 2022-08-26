// Copyright (c) Microsoft Corporation
// Licensed under the MIT License.

// Code generated by "genlayout"; DO NOT EDIT.

package winmd

import (
	"fmt"
	"github.com/microsoft/go-winmd/flags"
)

// Define CodedTable function

// CodedTable returns the table associated to c.
func (t *Tables) CodedTable(c CodedIndex) *Table[Record] {
	switch c.table {
	case tableAssembly:
		return (*Table[Record])(&t.Assembly)
	case tableAssemblyRef:
		return (*Table[Record])(&t.AssemblyRef)
	case tableClassLayout:
		return (*Table[Record])(&t.ClassLayout)
	case tableConstant:
		return (*Table[Record])(&t.Constant)
	case tableCustomAttribute:
		return (*Table[Record])(&t.CustomAttribute)
	case tableDeclSecurity:
		return (*Table[Record])(&t.DeclSecurity)
	case tableEventMap:
		return (*Table[Record])(&t.EventMap)
	case tableEvent:
		return (*Table[Record])(&t.Event)
	case tableExportedType:
		return (*Table[Record])(&t.ExportedType)
	case tableField:
		return (*Table[Record])(&t.Field)
	case tableFieldLayout:
		return (*Table[Record])(&t.FieldLayout)
	case tableFieldMarshal:
		return (*Table[Record])(&t.FieldMarshal)
	case tableFieldRVA:
		return (*Table[Record])(&t.FieldRVA)
	case tableFile:
		return (*Table[Record])(&t.File)
	case tableGenericParam:
		return (*Table[Record])(&t.GenericParam)
	case tableGenericParamConstraint:
		return (*Table[Record])(&t.GenericParamConstraint)
	case tableImplMap:
		return (*Table[Record])(&t.ImplMap)
	case tableInterfaceImpl:
		return (*Table[Record])(&t.InterfaceImpl)
	case tableManifestResource:
		return (*Table[Record])(&t.ManifestResource)
	case tableMemberRef:
		return (*Table[Record])(&t.MemberRef)
	case tableMethodDef:
		return (*Table[Record])(&t.MethodDef)
	case tableMethodImpl:
		return (*Table[Record])(&t.MethodImpl)
	case tableMethodSemantics:
		return (*Table[Record])(&t.MethodSemantics)
	case tableMethodSpec:
		return (*Table[Record])(&t.MethodSpec)
	case tableModule:
		return (*Table[Record])(&t.Module)
	case tableModuleRef:
		return (*Table[Record])(&t.ModuleRef)
	case tableNestedClass:
		return (*Table[Record])(&t.NestedClass)
	case tableParam:
		return (*Table[Record])(&t.Param)
	case tableProperty:
		return (*Table[Record])(&t.Property)
	case tablePropertyMap:
		return (*Table[Record])(&t.PropertyMap)
	case tableStandAloneSig:
		return (*Table[Record])(&t.StandAloneSig)
	case tableTypeDef:
		return (*Table[Record])(&t.TypeDef)
	case tableTypeRef:
		return (*Table[Record])(&t.TypeRef)
	case tableTypeSpec:
		return (*Table[Record])(&t.TypeSpec)
	default:
		return nil
	}
}

// Define table values

const (
	tableAssembly               table = 32
	tableAssemblyOS             table = 34
	tableAssemblyProcessor      table = 33
	tableAssemblyRef            table = 35
	tableAssemblyRefOS          table = 37
	tableAssemblyRefProcessor   table = 36
	tableClassLayout            table = 15
	tableConstant               table = 11
	tableCustomAttribute        table = 12
	tableDeclSecurity           table = 14
	tableEventMap               table = 18
	tableEvent                  table = 20
	tableExportedType           table = 39
	tableField                  table = 4
	tableFieldLayout            table = 16
	tableFieldMarshal           table = 13
	tableFieldRVA               table = 29
	tableFile                   table = 38
	tableGenericParam           table = 42
	tableGenericParamConstraint table = 44
	tableImplMap                table = 28
	tableInterfaceImpl          table = 9
	tableManifestResource       table = 40
	tableMemberRef              table = 10
	tableMethodDef              table = 6
	tableMethodImpl             table = 25
	tableMethodSemantics        table = 24
	tableMethodSpec             table = 43
	tableModule                 table = 0
	tableModuleRef              table = 26
	tableNestedClass            table = 41
	tableParam                  table = 8
	tableProperty               table = 23
	tablePropertyMap            table = 21
	tableStandAloneSig          table = 17
	tableTypeDef                table = 2
	tableTypeRef                table = 1
	tableTypeSpec               table = 27
)

// Implement table interface

func (Assembly) table() table { return tableAssembly }

func (assemblyOS) table() table { return tableAssemblyOS }

func (assemblyProcessor) table() table { return tableAssemblyProcessor }

func (AssemblyRef) table() table { return tableAssemblyRef }

func (assemblyRefOS) table() table { return tableAssemblyRefOS }

func (assemblyRefProcessor) table() table { return tableAssemblyRefProcessor }

func (ClassLayout) table() table { return tableClassLayout }

func (Constant) table() table { return tableConstant }

func (CustomAttribute) table() table { return tableCustomAttribute }

func (DeclSecurity) table() table { return tableDeclSecurity }

func (EventMap) table() table { return tableEventMap }

func (Event) table() table { return tableEvent }

func (ExportedType) table() table { return tableExportedType }

func (Field) table() table { return tableField }

func (FieldLayout) table() table { return tableFieldLayout }

func (FieldMarshal) table() table { return tableFieldMarshal }

func (FieldRVA) table() table { return tableFieldRVA }

func (File) table() table { return tableFile }

func (GenericParam) table() table { return tableGenericParam }

func (GenericParamConstraint) table() table { return tableGenericParamConstraint }

func (ImplMap) table() table { return tableImplMap }

func (InterfaceImpl) table() table { return tableInterfaceImpl }

func (ManifestResource) table() table { return tableManifestResource }

func (MemberRef) table() table { return tableMemberRef }

func (MethodDef) table() table { return tableMethodDef }

func (MethodImpl) table() table { return tableMethodImpl }

func (MethodSemantics) table() table { return tableMethodSemantics }

func (MethodSpec) table() table { return tableMethodSpec }

func (Module) table() table { return tableModule }

func (ModuleRef) table() table { return tableModuleRef }

func (NestedClass) table() table { return tableNestedClass }

func (Param) table() table { return tableParam }

func (Property) table() table { return tableProperty }

func (PropertyMap) table() table { return tablePropertyMap }

func (StandAloneSig) table() table { return tableStandAloneSig }

func (TypeDef) table() table { return tableTypeDef }

func (TypeRef) table() table { return tableTypeRef }

func (TypeSpec) table() table { return tableTypeSpec }

// Define static info

func staticTableInfo(tbl table) []columnInfo {
	switch tbl {
	case tableAssembly:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeBlob},
			{columnType: columnTypeString},
			{columnType: columnTypeString},
		}
	case tableAssemblyOS:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeUint, size: 4},
		}
	case tableAssemblyProcessor:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
		}
	case tableAssemblyRef:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeBlob},
			{columnType: columnTypeString},
			{columnType: columnTypeString},
			{columnType: columnTypeBlob},
		}
	case tableAssemblyRefOS:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeIndex, table: tableAssemblyRef},
		}
	case tableAssemblyRefProcessor:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeIndex, table: tableAssemblyRef},
		}
	case tableClassLayout:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeIndex, table: tableTypeDef},
		}
	case tableConstant:
		return []columnInfo{
			{columnType: columnTypeUint, size: 1},
			{columnType: columnTypeUint, size: 1},
			{columnType: columnTypeCodedIndex, coded: codedHasConstant},
			{columnType: columnTypeBlob},
		}
	case tableCustomAttribute:
		return []columnInfo{
			{columnType: columnTypeCodedIndex, coded: codedHasCustomAttribute},
			{columnType: columnTypeCodedIndex, coded: codedCustomAttributeType},
			{columnType: columnTypeBlob},
		}
	case tableDeclSecurity:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeCodedIndex, coded: codedHasDeclSecurity},
			{columnType: columnTypeBlob},
		}
	case tableEventMap:
		return []columnInfo{
			{columnType: columnTypeIndex, table: tableTypeDef},
			{columnType: columnTypeSlice, table: tableEvent},
		}
	case tableEvent:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeString},
			{columnType: columnTypeCodedIndex, coded: codedTypeDefOrRef},
		}
	case tableExportedType:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeString},
			{columnType: columnTypeString},
			{columnType: columnTypeCodedIndex, coded: codedImplementation},
		}
	case tableField:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeString},
			{columnType: columnTypeBlob},
		}
	case tableFieldLayout:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeIndex, table: tableField},
		}
	case tableFieldMarshal:
		return []columnInfo{
			{columnType: columnTypeCodedIndex, coded: codedHasFieldMarshal},
			{columnType: columnTypeBlob},
		}
	case tableFieldRVA:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeIndex, table: tableField},
		}
	case tableFile:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeString},
			{columnType: columnTypeBlob},
		}
	case tableGenericParam:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeCodedIndex, coded: codedTypeOrMethodDef},
			{columnType: columnTypeString},
		}
	case tableGenericParamConstraint:
		return []columnInfo{
			{columnType: columnTypeIndex, table: tableGenericParam},
			{columnType: columnTypeCodedIndex, coded: codedTypeDefOrRef},
		}
	case tableImplMap:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeCodedIndex, coded: codedMemberForwarded},
			{columnType: columnTypeString},
			{columnType: columnTypeIndex, table: tableModuleRef},
		}
	case tableInterfaceImpl:
		return []columnInfo{
			{columnType: columnTypeIndex, table: tableTypeDef},
			{columnType: columnTypeCodedIndex, coded: codedTypeDefOrRef},
		}
	case tableManifestResource:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeString},
			{columnType: columnTypeCodedIndex, coded: codedImplementation},
		}
	case tableMemberRef:
		return []columnInfo{
			{columnType: columnTypeCodedIndex, coded: codedMemberRefParent},
			{columnType: columnTypeString},
			{columnType: columnTypeBlob},
		}
	case tableMethodDef:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeString},
			{columnType: columnTypeBlob},
			{columnType: columnTypeSlice, table: tableParam},
		}
	case tableMethodImpl:
		return []columnInfo{
			{columnType: columnTypeIndex, table: tableTypeDef},
			{columnType: columnTypeCodedIndex, coded: codedMethodDefOrRef},
			{columnType: columnTypeCodedIndex, coded: codedMethodDefOrRef},
		}
	case tableMethodSemantics:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeIndex, table: tableMethodDef},
			{columnType: columnTypeCodedIndex, coded: codedHasSemantics},
		}
	case tableMethodSpec:
		return []columnInfo{
			{columnType: columnTypeCodedIndex, coded: codedMethodDefOrRef},
			{columnType: columnTypeBlob},
		}
	case tableModule:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeString},
			{columnType: columnTypeGUID},
			{columnType: columnTypeGUID},
			{columnType: columnTypeGUID},
		}
	case tableModuleRef:
		return []columnInfo{
			{columnType: columnTypeString},
		}
	case tableNestedClass:
		return []columnInfo{
			{columnType: columnTypeIndex, table: tableTypeDef},
			{columnType: columnTypeIndex, table: tableTypeDef},
		}
	case tableParam:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeString},
		}
	case tableProperty:
		return []columnInfo{
			{columnType: columnTypeUint, size: 2},
			{columnType: columnTypeString},
			{columnType: columnTypeBlob},
		}
	case tablePropertyMap:
		return []columnInfo{
			{columnType: columnTypeIndex, table: tableTypeDef},
			{columnType: columnTypeSlice, table: tableProperty},
		}
	case tableStandAloneSig:
		return []columnInfo{
			{columnType: columnTypeBlob},
		}
	case tableTypeDef:
		return []columnInfo{
			{columnType: columnTypeUint, size: 4},
			{columnType: columnTypeString},
			{columnType: columnTypeString},
			{columnType: columnTypeCodedIndex, coded: codedTypeDefOrRef},
			{columnType: columnTypeIndex, table: tableField},
			{columnType: columnTypeIndex, table: tableMethodDef},
		}
	case tableTypeRef:
		return []columnInfo{
			{columnType: columnTypeCodedIndex, coded: codedResolutionScope},
			{columnType: columnTypeString},
			{columnType: columnTypeString},
		}
	case tableTypeSpec:
		return []columnInfo{
			{columnType: columnTypeBlob},
		}
	default:
		panic(fmt.Sprintf("table %v not supported", tbl))
	}
}

// Define tables struct

type tables struct {
	Assembly               Table[Assembly]
	AssemblyRef            Table[AssemblyRef]
	ClassLayout            Table[ClassLayout]
	Constant               Table[Constant]
	CustomAttribute        Table[CustomAttribute]
	DeclSecurity           Table[DeclSecurity]
	EventMap               Table[EventMap]
	Event                  Table[Event]
	ExportedType           Table[ExportedType]
	Field                  Table[Field]
	FieldLayout            Table[FieldLayout]
	FieldMarshal           Table[FieldMarshal]
	FieldRVA               Table[FieldRVA]
	File                   Table[File]
	GenericParam           Table[GenericParam]
	GenericParamConstraint Table[GenericParamConstraint]
	ImplMap                Table[ImplMap]
	InterfaceImpl          Table[InterfaceImpl]
	ManifestResource       Table[ManifestResource]
	MemberRef              Table[MemberRef]
	MethodDef              Table[MethodDef]
	MethodImpl             Table[MethodImpl]
	MethodSemantics        Table[MethodSemantics]
	MethodSpec             Table[MethodSpec]
	Module                 Table[Module]
	ModuleRef              Table[ModuleRef]
	NestedClass            Table[NestedClass]
	Param                  Table[Param]
	Property               Table[Property]
	PropertyMap            Table[PropertyMap]
	StandAloneSig          Table[StandAloneSig]
	TypeDef                Table[TypeDef]
	TypeRef                Table[TypeRef]
	TypeSpec               Table[TypeSpec]
}

func initTables(t *Tables) {
	t.Assembly = newTable[Assembly](t, tableAssembly)
	t.AssemblyRef = newTable[AssemblyRef](t, tableAssemblyRef)
	t.ClassLayout = newTable[ClassLayout](t, tableClassLayout)
	t.Constant = newTable[Constant](t, tableConstant)
	t.CustomAttribute = newTable[CustomAttribute](t, tableCustomAttribute)
	t.DeclSecurity = newTable[DeclSecurity](t, tableDeclSecurity)
	t.EventMap = newTable[EventMap](t, tableEventMap)
	t.Event = newTable[Event](t, tableEvent)
	t.ExportedType = newTable[ExportedType](t, tableExportedType)
	t.Field = newTable[Field](t, tableField)
	t.FieldLayout = newTable[FieldLayout](t, tableFieldLayout)
	t.FieldMarshal = newTable[FieldMarshal](t, tableFieldMarshal)
	t.FieldRVA = newTable[FieldRVA](t, tableFieldRVA)
	t.File = newTable[File](t, tableFile)
	t.GenericParam = newTable[GenericParam](t, tableGenericParam)
	t.GenericParamConstraint = newTable[GenericParamConstraint](t, tableGenericParamConstraint)
	t.ImplMap = newTable[ImplMap](t, tableImplMap)
	t.InterfaceImpl = newTable[InterfaceImpl](t, tableInterfaceImpl)
	t.ManifestResource = newTable[ManifestResource](t, tableManifestResource)
	t.MemberRef = newTable[MemberRef](t, tableMemberRef)
	t.MethodDef = newTable[MethodDef](t, tableMethodDef)
	t.MethodImpl = newTable[MethodImpl](t, tableMethodImpl)
	t.MethodSemantics = newTable[MethodSemantics](t, tableMethodSemantics)
	t.MethodSpec = newTable[MethodSpec](t, tableMethodSpec)
	t.Module = newTable[Module](t, tableModule)
	t.ModuleRef = newTable[ModuleRef](t, tableModuleRef)
	t.NestedClass = newTable[NestedClass](t, tableNestedClass)
	t.Param = newTable[Param](t, tableParam)
	t.Property = newTable[Property](t, tableProperty)
	t.PropertyMap = newTable[PropertyMap](t, tablePropertyMap)
	t.StandAloneSig = newTable[StandAloneSig](t, tableStandAloneSig)
	t.TypeDef = newTable[TypeDef](t, tableTypeDef)
	t.TypeRef = newTable[TypeRef](t, tableTypeRef)
	t.TypeSpec = newTable[TypeSpec](t, tableTypeSpec)
}

// Define table decoding functions

func (rec *Assembly) decode(r recordReader) error {
	rec.HashAlgID = flags.AssemblyHashAlgorithm(r.uint32())
	rec.MajorVersion = r.uint16()
	rec.MinorVersion = r.uint16()
	rec.BuildNumber = r.uint16()
	rec.RevisionNumber = r.uint16()
	rec.Flags = flags.AssemblyFlags(r.uint32())
	rec.PublicKey = r.blob()
	rec.Name = r.string()
	rec.Culture = r.string()
	return r.err
}

func (rec *assemblyOS) decode(r recordReader) error {
	rec.OSPlatformID = r.uint32()
	rec.OSMajorVersion = r.uint32()
	rec.OSMinorVersion = r.uint32()
	return r.err
}

func (rec *assemblyProcessor) decode(r recordReader) error {
	rec.Processor = r.uint32()
	return r.err
}

func (rec *AssemblyRef) decode(r recordReader) error {
	rec.MajorVersion = r.uint16()
	rec.MinorVersion = r.uint16()
	rec.BuildNumber = r.uint16()
	rec.RevisionNumber = r.uint16()
	rec.Flags = flags.AssemblyFlags(r.uint32())
	rec.PublicKeyOrToken = r.blob()
	rec.Name = r.string()
	rec.Culture = r.string()
	rec.HashValue = r.blob()
	return r.err
}

func (rec *assemblyRefOS) decode(r recordReader) error {
	rec.OSPlatformID = r.uint32()
	rec.OSMajorVersion = r.uint32()
	rec.OSMinorVersion = r.uint32()
	rec.AssemblyRef = r.index(tableAssemblyRef)
	return r.err
}

func (rec *assemblyRefProcessor) decode(r recordReader) error {
	rec.Processor = r.uint32()
	rec.AssemblyRef = r.index(tableAssemblyRef)
	return r.err
}

func (rec *ClassLayout) decode(r recordReader) error {
	rec.PackingSize = r.uint16()
	rec.ClassSize = r.uint32()
	rec.Parent = r.index(tableTypeDef)
	return r.err
}

func (rec *Constant) decode(r recordReader) error {
	rec.Type = flags.ElementType(r.uint8())
	rec.Padding = r.uint8()
	rec.Parent = r.coded(codedHasConstant)
	rec.Value = r.blob()
	return r.err
}

func (rec *CustomAttribute) decode(r recordReader) error {
	rec.Parent = r.coded(codedHasCustomAttribute)
	rec.Type = r.coded(codedCustomAttributeType)
	rec.Value = r.blob()
	return r.err
}

func (rec *DeclSecurity) decode(r recordReader) error {
	rec.Action = r.uint16()
	rec.Parent = r.coded(codedHasDeclSecurity)
	rec.PermissionSet = r.blob()
	return r.err
}

func (rec *EventMap) decode(r recordReader) error {
	rec.Parent = r.index(tableTypeDef)
	rec.EventList = r.slice(tableEventMap, tableEvent)
	return r.err
}

func (rec *Event) decode(r recordReader) error {
	rec.EventFlags = flags.EventAttributes(r.uint16())
	rec.Name = r.string()
	rec.EventType = r.coded(codedTypeDefOrRef)
	return r.err
}

func (rec *ExportedType) decode(r recordReader) error {
	rec.Flags = flags.TypeAttributes(r.uint32())
	rec.TypeDefID = r.uint32()
	rec.Name = r.string()
	rec.Namespace = r.string()
	rec.Implementation = r.coded(codedImplementation)
	return r.err
}

func (rec *Field) decode(r recordReader) error {
	rec.Flags = flags.FieldAttributes(r.uint16())
	rec.Name = r.string()
	rec.Signature = r.blob()
	return r.err
}

func (rec *FieldLayout) decode(r recordReader) error {
	rec.Offset = r.uint32()
	rec.Field = r.index(tableField)
	return r.err
}

func (rec *FieldMarshal) decode(r recordReader) error {
	rec.Parent = r.coded(codedHasFieldMarshal)
	rec.NativeType = r.blob()
	return r.err
}

func (rec *FieldRVA) decode(r recordReader) error {
	rec.RVA = r.uint32()
	rec.Field = r.index(tableField)
	return r.err
}

func (rec *File) decode(r recordReader) error {
	rec.Flags = flags.FileAttributes(r.uint16())
	rec.Name = r.string()
	rec.HashValue = r.blob()
	return r.err
}

func (rec *GenericParam) decode(r recordReader) error {
	rec.Number = r.uint16()
	rec.Flags = flags.GenericParamAttributes(r.uint16())
	rec.Owner = r.coded(codedTypeOrMethodDef)
	rec.Name = r.string()
	return r.err
}

func (rec *GenericParamConstraint) decode(r recordReader) error {
	rec.Owner = r.index(tableGenericParam)
	rec.Constraint = r.coded(codedTypeDefOrRef)
	return r.err
}

func (rec *ImplMap) decode(r recordReader) error {
	rec.MappingFlags = flags.PInvokeAttributes(r.uint16())
	rec.MemberForwarded = r.coded(codedMemberForwarded)
	rec.ImportName = r.string()
	rec.ImportScope = r.index(tableModuleRef)
	return r.err
}

func (rec *InterfaceImpl) decode(r recordReader) error {
	rec.Class = r.index(tableTypeDef)
	rec.Interface = r.coded(codedTypeDefOrRef)
	return r.err
}

func (rec *ManifestResource) decode(r recordReader) error {
	rec.Offset = r.uint32()
	rec.Flags = flags.ManifestResourceAttributes(r.uint32())
	rec.Name = r.string()
	rec.Implementation = r.coded(codedImplementation)
	return r.err
}

func (rec *MemberRef) decode(r recordReader) error {
	rec.Class = r.coded(codedMemberRefParent)
	rec.Name = r.string()
	rec.Signature = r.blob()
	return r.err
}

func (rec *MethodDef) decode(r recordReader) error {
	rec.RVA = r.uint32()
	rec.ImplFlags = flags.MethodImplAttributes(r.uint16())
	rec.Flags = flags.MethodAttributes(r.uint16())
	rec.Name = r.string()
	rec.Signature = r.blob()
	rec.ParamList = r.slice(tableMethodDef, tableParam)
	return r.err
}

func (rec *MethodImpl) decode(r recordReader) error {
	rec.Class = r.index(tableTypeDef)
	rec.MethodBody = r.coded(codedMethodDefOrRef)
	rec.MethodDeclaration = r.coded(codedMethodDefOrRef)
	return r.err
}

func (rec *MethodSemantics) decode(r recordReader) error {
	rec.Semantics = flags.MethodSemanticsAttributes(r.uint16())
	rec.Method = r.index(tableMethodDef)
	rec.Association = r.coded(codedHasSemantics)
	return r.err
}

func (rec *MethodSpec) decode(r recordReader) error {
	rec.Method = r.coded(codedMethodDefOrRef)
	rec.Instantiation = r.blob()
	return r.err
}

func (rec *Module) decode(r recordReader) error {
	rec.Generation = r.uint16()
	rec.Name = r.string()
	rec.Mvid = r.guid()
	rec.EncID = r.guid()
	rec.EncBaseID = r.guid()
	return r.err
}

func (rec *ModuleRef) decode(r recordReader) error {
	rec.Name = r.string()
	return r.err
}

func (rec *NestedClass) decode(r recordReader) error {
	rec.NestedClass = r.index(tableTypeDef)
	rec.EnclosingClass = r.index(tableTypeDef)
	return r.err
}

func (rec *Param) decode(r recordReader) error {
	rec.Flags = flags.ParamAttributes(r.uint16())
	rec.Sequence = r.uint16()
	rec.Name = r.string()
	return r.err
}

func (rec *Property) decode(r recordReader) error {
	rec.Flags = flags.PropertyAttributes(r.uint16())
	rec.Name = r.string()
	rec.Type = r.blob()
	return r.err
}

func (rec *PropertyMap) decode(r recordReader) error {
	rec.Parent = r.index(tableTypeDef)
	rec.PropertyList = r.slice(tablePropertyMap, tableProperty)
	return r.err
}

func (rec *StandAloneSig) decode(r recordReader) error {
	rec.Signature = r.blob()
	return r.err
}

func (rec *TypeDef) decode(r recordReader) error {
	rec.Flags = flags.TypeAttributes(r.uint32())
	rec.Name = r.string()
	rec.Namespace = r.string()
	rec.Extends = r.coded(codedTypeDefOrRef)
	rec.FieldList = r.index(tableField)
	rec.MethodList = r.index(tableMethodDef)
	return r.err
}

func (rec *TypeRef) decode(r recordReader) error {
	rec.ResolutionScope = r.coded(codedResolutionScope)
	rec.Name = r.string()
	rec.Namespace = r.string()
	return r.err
}

func (rec *TypeSpec) decode(r recordReader) error {
	rec.Signature = r.blob()
	return r.err
}