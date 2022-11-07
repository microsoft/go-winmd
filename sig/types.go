// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package sig

import "github.com/microsoft/go-winmd/flags"

// CodedIndex indexes a record on any table.
type CodedIndex struct {
	Index uint32
	Tag   int8
}

type FieldSig struct {
	Mod  []CustomMod
	Type Type
}

type MethodDefSig struct {
	HasThis      bool
	ExplicitThis bool
	VarArgs      bool
	Generic      uint32
	RetType      RetType
	Param        []Param
}

type LocalVarSig []LocalVar

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

type Type struct {
	Kind  flags.ElementType
	Value any // optional
}

type Array struct {
	Type        Type
	Rank        uint32
	Sizes       []uint32
	LowerBounds []int32
}

type CustomModKind uint8

const (
	CustomModKind_Opt CustomModKind = iota
	CustomModKind_Reqd
)

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

type Param struct {
	Kind ParamKind
	Mod  []CustomMod
	Type Type // empty if Kind is TypedByRef
}

type RetTypeKind uint8

const (
	RetTypeKind_ByValue RetTypeKind = iota
	RetTypeKind_ByRef
	RetTypeKind_TypedByRef
	RetTypeKind_Void
)

type RetType struct {
	Kind RetTypeKind
	Mod  []CustomMod
	Type Type // empty if Kind is TypedByRef or Void
}

type CustomModType struct {
	Mod  *CustomMod
	Type *Type // can be nil if type is void
}

type GenericInst struct {
	Class bool
	Index CodedIndex
	Type  []Type
}

type TypeSpec struct {
	Kind  flags.ElementType
	Value any
}

type MethodSpec []Type
