// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"github.com/microsoft/go-winmd/flags"
)

type SigType uint8

const (
	SigType_MethodRef SigType = iota
	SigType_MethodDef
	SigType_Field
	SigType_Property
	SigType_LocalVar
	SigType_Type
	SigType_Method
)

// MethodDefSig is defined in §II.23.2.1.
type MethodDefSig struct {
	HasThis      bool
	ExplicitThis bool
	VarArgs      bool
	Generic      uint32
	RetType      RetType
	Param        []SigParam
}

// MethodRefSig is defined in §II.23.2.2.
type MethodRefSig struct {
	MethodDefSig
	VariableParam []Param
}

// StandAloneMethodSig is defined in §II.23.2.3 but is not supported.

// FieldSig is defined in §II.23.2.4.
type FieldSig struct {
	Mod  []CustomMod
	Type Type
}

// PropertySig is defined in §II.23.2.5.
type PropertySig struct {
	HasThis bool
	FieldSig
	Param []SigParam
}

// LocalVarSig is defined in §II.23.2.6.
type LocalVarSig []LocalVar

// CodedIndex indexes a record on any table.
//type CodedIndex struct {
//	Index uint32
//	Tag   int8
//}

type Constraint struct {
	Pinned bool
}

type LocalVarMod struct {
	Mod        *CustomMod
	Constraint Constraint
}

type LocalVarKind uint8

const (
	LocalVarKind_ByValue LocalVarKind = iota
	LocalVarKind_ByRef
	LocalVarKind_TypedByRef
)

type LocalVar struct {
	Kind LocalVarKind
	Mod  []LocalVarMod // empty if Kind is TypedByRef
	Type Type          // empty if Kind is TypedByRef
}

type CustomModKind uint8

const (
	CustomModKind_Opt CustomModKind = iota
	CustomModKind_Reqd
)

// CustomMod is defined in §II.23.2.7.
type CustomMod struct {
	Kind  CustomModKind
	Index CodedIndex
}

type ParamKind uint8

const (
	ParamKind_ByValue ParamKind = iota
	ParamKind_ByRef
	ParamKind_TypedByRef
)

// SigParam is defined in §II.23.2.10.
type SigParam struct {
	Kind ParamKind
	Type Type // empty if Kind is TypedByRef
}

type RetTypeKind uint8

const (
	RetTypeKind_ByValue RetTypeKind = iota
	RetTypeKind_ByRef
	RetTypeKind_TypedByRef
	RetTypeKind_Void
)

// RetType is defined in §II.23.2.11.
type RetType struct {
	Kind RetTypeKind
	Type Type // empty if Kind is TypedByRef or Void
}

// Type is defined in §II.23.2.12.
type Type struct {
	Kind  flags.ElementType
	Mod   []CustomMod
	Value any // optional
}

// Array is a Type with an ArrayShape, where ArrayShape is defined in §II.23.2.13.
type Array struct {
	Type        Type
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
type SigMethodSpec []Type

type CustomModType struct {
	Mod  *CustomMod
	Type *Type // can be nil if type is void
}

type GenericInst struct {
	Class bool
	Index CodedIndex
	Type  []Type
}
