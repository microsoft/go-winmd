// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

/*
Package coded defines coded indices tags as appear in §II.24.2.6.
*/
package coded

// Null is the sentinel value for null coded indices.
const Null = -1

const (
	TypeDefOrRef_TypeDef = iota
	TypeDefOrRef_TypeRef
	TypeDefOrRef_TypeSpec
)

const (
	HasConstant_Field = iota
	HasConstant_Param
	HasConstant_Property
)

const (
	HasFieldMarshal_Field = iota
	HasFieldMarshal_Param
)

const (
	HasDeclSecurity_TypeDef = iota
	HasDeclSecurity_MethodDef
	HasDeclSecurity_tableAssembly
)

const (
	MemberRefParent_TypeDef = iota
	MemberRefParent_TypeRef
	MemberRefParent_ModuleRef
	MemberRefParent_MethodDef
	MemberRefParent_TypeSpec
)

const (
	HasSemantics_Event = iota
	HasSemantics_Property
)

const (
	MethodDefOrRef_MethodDef = iota
	MethodDefOrRef_MemberRef
)
const (
	MemberForwarded_Field = iota
	MemberForwarded_MethodDef
)

const (
	Implementation_File = iota
	Implementation_AssemblyRef
	Implementation_ExportedType
)

const (
	_ = iota
	_
	CustomAttributeType_MethodDef
	CustomAttributeType_MemberRef
	_
)

const (
	ResolutionScope_Module = iota
	ResolutionScope_ModuleRef
	ResolutionScope_AssemblyRef
	ResolutionScope_TypeRef
)

const (
	TypeOrMethodDef_TypeDef = iota
	TypeOrMethodDef_MethodDef
)

const (
	HasCustomAttribute_MethodDef = iota
	HasCustomAttribute_Field
	HasCustomAttribute_TypeRef
	HasCustomAttribute_TypeDef
	HasCustomAttribute_Param
	HasCustomAttribute_InterfaceImpl
	HasCustomAttribute_MemberRef
	HasCustomAttribute_Module
	HasCustomAttribute_None
	HasCustomAttribute_Property
	HasCustomAttribute_Event
	HasCustomAttribute_StandAloneSig
	HasCustomAttribute_ModuleRef
	HasCustomAttribute_TypeSpec
	HasCustomAttribute_Assembly
	HasCustomAttribute_AssemblyRef
	HasCustomAttribute_File
	HasCustomAttribute_ExportedType
	HasCustomAttribute_ManifestResource
	HasCustomAttribute_GenericParam
	HasCustomAttribute_GenericParamConstraint
	HasCustomAttribute_MethodSpec
)
