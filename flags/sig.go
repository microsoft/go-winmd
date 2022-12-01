// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package flags

// SigKind is defined in §II.23.2.
type SigKind uint8

const (
	// SigKind_DEFAULT (0) through SigKind_VARARG (5) are method signature types. Defined in §II.23.2.3.
	SigKind_DEFAULT  SigKind = 0x0
	SigKind_C        SigKind = 0x1
	SigKind_STDCALL  SigKind = 0x2
	SigKind_THISCALL SigKind = 0x3
	SigKind_FASTCALL SigKind = 0x4
	SigKind_VARARG   SigKind = 0x5

	// SigKind_FIELD is a FIELD signature. Defined in §II.23.2.4.
	SigKind_FIELD SigKind = 0x6

	// SigKind_LOCAL is a local variable signature. Defined in §II.23.2.6.
	SigKind_LOCAL SigKind = 0x7

	// SigKind_PROPERTY is a property signature. Defined in §II.23.2.5.
	SigKind_PROPERTY SigKind = 0x8
)

type SigAttributes uint8

const (
	SigAttributes_NONE         SigAttributes = 0x00
	SigAttributes_GENERIC      SigAttributes = 0x10
	SigAttributes_HASTHIS      SigAttributes = 0x20
	SigAttributes_EXPLICITTHIS SigAttributes = 0x40
	SigAttributes_SENTINEL     SigAttributes = 0x41
)
