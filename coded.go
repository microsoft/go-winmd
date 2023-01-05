// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"fmt"
	"math/bits"
)

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
	// codedTypeDefOrRefOrSpec is for signature decoding, defined in §II.23.2.8.
	codedTypeDefOrRefOrSpec
	codedMax
)

// codedMap maps each coded type to the list of table types that it may encode, in order.
var codedMap = [codedMax][]table{
	// The following entries are taken from §II.24.2.6.
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
	// codedTypeDefOrRefOrSpec is for signature decoding, defined in §II.23.2.8. It isn't
	// technically called a coded index by the spec, but it's encoded like one.
	codedTypeDefOrRefOrSpec: {tableTypeDef, tableTypeRef, tableTypeSpec},
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

// parseCoded parses an encoded CodedIndex.
func parseCoded(coded coded, code uint32) (CodedIndex, error) {
	tagbits := codedTagBits(coded)
	bitmask := (1 << tagbits) - 1
	if code < 1 {
		return CodedIndex{Tag: -1}, nil
	}
	row, tag := code>>tagbits-1, code&uint32(bitmask)
	_, ok := codedTable(coded, uint8(tag))
	if !ok {
		return CodedIndex{}, fmt.Errorf("unknown coded %d tag %d", coded, tag)
	}
	return CodedIndex{
		Index: Index(row),
		Tag:   int8(tag),
	}, nil
}
