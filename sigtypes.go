// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"github.com/microsoft/go-winmd/flags"
)

type (
	SigMethodDefBlob []byte
	SigFieldBlob     []byte
	SigPropertyBlob  []byte
)

// SigMethodDef is defined in §II.23.2.1.
type SigMethodDef struct {
	HasThis      bool
	ExplicitThis bool
	VarArgs      bool
	Generic      uint32
	RetType      SigRetType
	Param        []SigParam
}

// SigMethodRef is defined in §II.23.2.2.
type SigMethodRef struct {
	SigMethodDef
	VariableParam []Param
}

// StandAloneMethodSig is defined in §II.23.2.3 but is not supported.

// SigField is defined in §II.23.2.4.
type SigField struct {
	Type SigType
}

// SigProperty is defined in §II.23.2.5.
type SigProperty struct {
	HasThis bool
	SigField
	Param []SigParam
}

// SigLocalVars is defined as "LocalVarSig" in §II.23.2.6.
// This type represents the type of all local vars in a method, and the name has been changed for
// clarity and to make it easier to name "SigLocalVar".
type SigLocalVars []SigLocalVar

type SigConstraint struct {
	Pinned bool
}

type SigLocalVarMod struct {
	Mod        *SigCustomMod
	Constraint SigConstraint
}

type SigLocalVarKind uint8

const (
	SigLocalVarKind_ByValue SigLocalVarKind = iota
	SigLocalVarKind_ByRef
	SigLocalVarKind_TypedByRef
)

type SigLocalVar struct {
	Kind SigLocalVarKind
	Mod  []SigLocalVarMod // empty if Kind is TypedByRef
	Type SigType          // empty if Kind is TypedByRef
}

type SigCustomModKind uint8

const (
	SigCustomModKind_Opt SigCustomModKind = iota
	SigCustomModKind_Reqd
)

// SigCustomMod is defined in §II.23.2.7.
type SigCustomMod struct {
	Kind  SigCustomModKind
	Index CodedIndex
}

type SigParamKind uint8

const (
	SigParamKind_ByValue SigParamKind = iota
	SigParamKind_ByRef
	SigParamKind_TypedByRef
)

// SigParam is defined in §II.23.2.10.
type SigParam struct {
	Kind SigParamKind
	Type SigType // empty if Kind is TypedByRef
}

type SigRetTypeKind uint8

const (
	SigRetTypeKind_ByValue SigRetTypeKind = iota
	SigRetTypeKind_ByRef
	SigRetTypeKind_TypedByRef
	SigRetTypeKind_Void
)

// SigRetType is defined in §II.23.2.11.
type SigRetType struct {
	Kind SigRetTypeKind
	Type SigType // empty if Kind is TypedByRef or Void
}

// SigType is defined in §II.23.2.12.
type SigType struct {
	Kind  flags.ElementType
	Mod   []SigCustomMod
	Value any // optional
}

// SigArray is a SigType with an ArrayShape, where ArrayShape is defined in §II.23.2.13.
type SigArray struct {
	Type        SigType
	Rank        uint32
	Sizes       []uint32
	LowerBounds []int32
}

// SigTypeSpec is defined in §II.23.2.14.
type SigTypeSpec struct {
	Kind  flags.ElementType
	Value any
}

// SigMethodSpec is defined in §II.23.2.15
type SigMethodSpec []SigType

type SigGenericInst struct {
	Class bool
	Index CodedIndex
	Type  []SigType
}
