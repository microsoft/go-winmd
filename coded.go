// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import "math/bits"

type coded uint8

const (
	codedTypeDefOrRef coded = iota
	codedHasConstant
	codedHasFieldMarshal
	codedHasDeclSecurity
	codedMemberRefParent
	codedHasSemantics
	codedMethodDefOrRef
	codedMemberForwarded
	codedImplementation
	codedCustomAttributeType
	codedResolutionScope
	codedTypeOrMethodDef
	codedHasCustomAttribute
	codedMax
)

// codedMap is taken from §II.24.2.6.
var codedMap = [codedMax][]table{
	codedTypeDefOrRef:        {tableTypeDef, tableTypeRef, tableTypeSpec},
	codedHasConstant:         {tableField, tableParam, tableProperty},
	codedHasFieldMarshal:     {tableField, tableParam},
	codedHasDeclSecurity:     {tableTypeDef, tableMethodDef, tableAssembly},
	codedMemberRefParent:     {tableTypeDef, tableTypeRef, tableModuleRef, tableMethodDef, tableTypeSpec},
	codedHasSemantics:        {tableEvent, tableProperty},
	codedMethodDefOrRef:      {tableMethodDef, tableMemberRef},
	codedMemberForwarded:     {tableField, tableMethodDef},
	codedImplementation:      {tableFile, tableAssemblyRef, tableExportedType},
	codedCustomAttributeType: {tableNone, tableNone, tableMethodDef, tableMemberRef, tableNone},
	codedResolutionScope:     {tableModule, tableModuleRef, tableAssemblyRef, tableTypeRef},
	codedTypeOrMethodDef:     {tableTypeDef, tableMethodDef},
	codedHasCustomAttribute: {
		tableMethodDef,
		tableField,
		tableTypeRef,
		tableTypeDef,
		tableParam,
		tableInterfaceImpl,
		tableMemberRef,
		tableModule,
		tableNone,
		tableProperty,
		tableEvent,
		tableStandAloneSig,
		tableModuleRef,
		tableTypeSpec,
		tableAssembly,
		tableAssemblyRef,
		tableFile,
		tableExportedType,
		tableManifestResource,
		tableGenericParam,
		tableGenericParamConstraint,
		tableMethodSpec,
	},
}

// codedTagBits returns the minimum number of bits
// to identify a table from the possible options.
func codedTagBits(c coded) int {
	return bits.Len8(uint8(len(codedMap[c]) - 1))
}

// codedTable returns the table associated to the tag
// of the coded c as defined in §II.24.2.6.
func codedTable(c coded, tag uint8) (table, bool) {
	tbls := codedMap[c]
	if int(tag) < len(tbls) {
		return tbls[tag], true
	}
	return tableNone, false
}
