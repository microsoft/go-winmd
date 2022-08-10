// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Stringer is only implemented by types which do not contain bit flags nor bitmasks,
// as those can't be autogenerated by stringer.
// TODO(#4): implement a custom stringer tool that understands bit flags and bitmasks.
//go:generate stringer -type=ElementType -trimprefix ElementType_
//go:generate stringer -type=AssemblyHashAlgorithm -trimprefix AssemblyHash_Algorithm

package flags

// AssemblyHashAlgorithm is defined in §II.23.1.1.
type AssemblyHashAlgorithm uint32

const (
	AssemblyHash_AlgorithmNone AssemblyHashAlgorithm = 0x0000
	AssemblyHash_AlgorithmMD5  AssemblyHashAlgorithm = 0x8003
	AssemblyHash_AlgorithmSHA1 AssemblyHashAlgorithm = 0x8004
)

// AssemblyFlags is defined in §II.23.1.2.
type AssemblyFlags uint32

const (
	AssemblyFlags_PublicKey                  AssemblyFlags = 0x0001
	AssemblyFlags_Retargetable               AssemblyFlags = 0x0100
	AssemblyFlags_DisableJITcompileOptimizer AssemblyFlags = 0x4000
	AssemblyFlags_EnableJITcompileTracking   AssemblyFlags = 0x8000
)

// EventAttributes is defined in §II.23.1.4.
type EventAttributes uint16

const (
	EventAttributes_SpecialName   EventAttributes = 0x0200
	EventAttributes_RTSpecialName EventAttributes = 0x0400
)

// FieldAttributes is defined in §II.23.1.5.
type FieldAttributes uint16

const (
	FieldAttributes_FieldAccessMask    FieldAttributes = 0x0007
	FieldAttributes_CompilerControlled FieldAttributes = 0x0000
	FieldAttributes_Private            FieldAttributes = 0x0001
	FieldAttributes_FamANDAssem        FieldAttributes = 0x0002
	FieldAttributes_Assembly           FieldAttributes = 0x0003
	FieldAttributes_Family             FieldAttributes = 0x0004
	FieldAttributes_FamORAssem         FieldAttributes = 0x0005
	FieldAttributes_Public             FieldAttributes = 0x0006
	FieldAttributes_Static             FieldAttributes = 0x0010
	FieldAttributes_InitOnly           FieldAttributes = 0x0020
	FieldAttributes_Literal            FieldAttributes = 0x0040
	FieldAttributes_NotSerialized      FieldAttributes = 0x0080
	FieldAttributes_SpecialName        FieldAttributes = 0x0200
	FieldAttributes_PInvokeImpl        FieldAttributes = 0x2000
	FieldAttributes_RTSpecialName      FieldAttributes = 0x0400
	FieldAttributes_HasFieldMarshal    FieldAttributes = 0x1000
	FieldAttributes_HasDefault         FieldAttributes = 0x8000
	FieldAttributes_HasFieldRVA        FieldAttributes = 0x0100
)

// FileAttributes is defined in §II.23.1.6.
type FileAttributes uint16

const (
	FileAttributes_ContainsMetaData   FileAttributes = 0x0000
	FileAttributes_ContainsNoMetaData FileAttributes = 0x0001
)

// GenericParamAttributes is defined in §II.23.1.7.
type GenericParamAttributes uint16

const (
	GenericParamAttributes_VarianceMask                   GenericParamAttributes = 0x0003
	GenericParamAttributes_None                           GenericParamAttributes = 0x0000
	GenericParamAttributes_Covariant                      GenericParamAttributes = 0x0001
	GenericParamAttributes_Contravariant                  GenericParamAttributes = 0x0002
	GenericParamAttributes_SpecialConstraintMask          GenericParamAttributes = 0x001C
	GenericParamAttributes_ReferenceTypeConstraint        GenericParamAttributes = 0x0004
	GenericParamAttributes_NotNullableValueTypeConstraint GenericParamAttributes = 0x0008
	GenericParamAttributes_DefaultConstructorConstraint   GenericParamAttributes = 0x0010
)

// PInvokeAttributes is defined in §II.23.1.8.
type PInvokeAttributes uint16

const (
	PInvokeAttributes_NoMangle            PInvokeAttributes = 0x0001
	PInvokeAttributes_CharSetMask         PInvokeAttributes = 0x0006
	PInvokeAttributes_CharSetNotSpec      PInvokeAttributes = 0x0000
	PInvokeAttributes_CharSetAnsi         PInvokeAttributes = 0x0002
	PInvokeAttributes_CharSetUnicode      PInvokeAttributes = 0x0004
	PInvokeAttributes_CharSetAuto         PInvokeAttributes = 0x0006
	PInvokeAttributes_SupportsLastError   PInvokeAttributes = 0x0040
	PInvokeAttributes_CallConvMask        PInvokeAttributes = 0x0700
	PInvokeAttributes_CallConvPlatformapi PInvokeAttributes = 0x0100
	PInvokeAttributes_CallConvCdecl       PInvokeAttributes = 0x0200
	PInvokeAttributes_CallConvStdcall     PInvokeAttributes = 0x0300
	PInvokeAttributes_CallConvThiscall    PInvokeAttributes = 0x0400
	PInvokeAttributes_CallConvFastcall    PInvokeAttributes = 0x0500
)

// ManifestResourceAttributes is defined in §II.23.1.9.
type ManifestResourceAttributes uint32

const (
	ManifestResourceAttributes_VisibilityMask ManifestResourceAttributes = 0x0007
	ManifestResourceAttributes_Public         ManifestResourceAttributes = 0x0001
	ManifestResourceAttributes_Private        ManifestResourceAttributes = 0x0002
)

// MethodAttributes is defined in §II.23.1.10.
type MethodAttributes uint16

const (
	MethodImplAttributes_MemberAccessMask   MethodImplAttributes = 0x0007
	MethodImplAttributes_CompilerControlled MethodImplAttributes = 0x0000
	MethodImplAttributes_Private            MethodImplAttributes = 0x0001
	MethodImplAttributes_FamANDAssem        MethodImplAttributes = 0x0002
	MethodImplAttributes_Assem              MethodImplAttributes = 0x0003
	MethodImplAttributes_Family             MethodImplAttributes = 0x0004
	MethodImplAttributes_FamORAssem         MethodImplAttributes = 0x0005
	MethodImplAttributes_Public             MethodImplAttributes = 0x0006
	MethodImplAttributes_Static             MethodImplAttributes = 0x0010
	MethodImplAttributes_Final              MethodImplAttributes = 0x0020
	MethodImplAttributes_Virtual            MethodImplAttributes = 0x0040
	MethodImplAttributes_HideBySig          MethodImplAttributes = 0x0080
	MethodImplAttributes_VtableLayoutMask   MethodImplAttributes = 0x0100
	MethodImplAttributes_ReuseSlot          MethodImplAttributes = 0x0000
	MethodImplAttributes_NewSlot            MethodImplAttributes = 0x0100
	MethodImplAttributes_Strict             MethodImplAttributes = 0x0200
	MethodImplAttributes_Abstract           MethodImplAttributes = 0x0400
	MethodImplAttributes_SpecialName        MethodImplAttributes = 0x0800
	MethodImplAttributes_PInvokeImpl        MethodImplAttributes = 0x2000
	MethodImplAttributes_UnmanagedExport    MethodImplAttributes = 0x0008
	MethodImplAttributes_RTSpecialName      MethodImplAttributes = 0x1000
	MethodImplAttributes_HasSecurity        MethodImplAttributes = 0x4000
	MethodImplAttributes_RequireSecObject   MethodImplAttributes = 0x8000
)

// MethodImplAttributes is defined in §II.23.1.11.
type MethodImplAttributes uint16

const (
	MethodImplAttributes_CodeTypeMask     MethodImplAttributes = 0x0003
	MethodImplAttributes_IL               MethodImplAttributes = 0x0000
	MethodImplAttributes_Native           MethodImplAttributes = 0x0001
	MethodImplAttributes_OPTIL            MethodImplAttributes = 0x0002
	MethodImplAttributes_Runtime          MethodImplAttributes = 0x0003
	MethodImplAttributes_ManagedMask      MethodImplAttributes = 0x0004
	MethodImplAttributes_Unmanaged        MethodImplAttributes = 0x0004
	MethodImplAttributes_Managed          MethodImplAttributes = 0x0000
	MethodImplAttributes_ForwardRef       MethodImplAttributes = 0x0010
	MethodImplAttributes_PreserveSig      MethodImplAttributes = 0x0080
	MethodImplAttributes_InternalCall     MethodImplAttributes = 0x1000
	MethodImplAttributes_Synchronized     MethodImplAttributes = 0x0020
	MethodImplAttributes_NoInlining       MethodImplAttributes = 0x0008
	MethodImplAttributes_MaxMethodImplVal MethodImplAttributes = 0xffff
	MethodImplAttributes_NoOptimization   MethodImplAttributes = 0x0040
)

// MethodSemanticsAttributes is defined in §II.23.1.12.
type MethodSemanticsAttributes uint16

const (
	MethodSemanticsAttributes_Setter   MethodSemanticsAttributes = 0x0001
	MethodSemanticsAttributes_Getter   MethodSemanticsAttributes = 0x0002
	MethodSemanticsAttributes_Other    MethodSemanticsAttributes = 0x0004
	MethodSemanticsAttributes_AddOn    MethodSemanticsAttributes = 0x0008
	MethodSemanticsAttributes_RemoveOn MethodSemanticsAttributes = 0x0010
	MethodSemanticsAttributes_Fire     MethodSemanticsAttributes = 0x0020
)

// ParamAttributes is defined in §II.23.1.13.
type ParamAttributes uint16

const (
	ParamAttributes_In              ParamAttributes = 0x0001
	ParamAttributes_Out             ParamAttributes = 0x0002
	ParamAttributes_Optional        ParamAttributes = 0x0010
	ParamAttributes_HasDefault      ParamAttributes = 0x1000
	ParamAttributes_HasFieldMarshal ParamAttributes = 0x2000
	ParamAttributes_Unused          ParamAttributes = 0xcfe0
)

// PropertyAttributes is defined in §II.23.1.14.
type PropertyAttributes uint16

const (
	PropertyAttributes_SpecialName   PropertyAttributes = 0x0200
	PropertyAttributes_RTSpecialName PropertyAttributes = 0x0400
	PropertyAttributes_HasDefault    PropertyAttributes = 0x1000
	PropertyAttributes_Unused        PropertyAttributes = 0xe9ff
)

// TypeAttributes is defined in §II.23.1.15.
type TypeAttributes uint32

const (
	TypeAttributes_VisibilityMask         TypeAttributes = 0x00000007
	TypeAttributes_NotPublic              TypeAttributes = 0x00000000
	TypeAttributes_Public                 TypeAttributes = 0x00000001
	TypeAttributes_NestedPublic           TypeAttributes = 0x00000002
	TypeAttributes_NestedPrivate          TypeAttributes = 0x00000003
	TypeAttributes_NestedFamily           TypeAttributes = 0x00000004
	TypeAttributes_NestedAssembly         TypeAttributes = 0x00000005
	TypeAttributes_NestedFamANDAssem      TypeAttributes = 0x00000006
	TypeAttributes_NestedFamORAssem       TypeAttributes = 0x00000007
	TypeAttributes_LayoutMask             TypeAttributes = 0x00000018
	TypeAttributes_AutoLayout             TypeAttributes = 0x00000000
	TypeAttributes_SequentialLayout       TypeAttributes = 0x00000008
	TypeAttributes_ExplicitLayout         TypeAttributes = 0x00000010
	TypeAttributes_ClassSemanticsMask     TypeAttributes = 0x00000020
	TypeAttributes_Class                  TypeAttributes = 0x00000000
	TypeAttributes_Interface              TypeAttributes = 0x00000020
	TypeAttributes_Abstract               TypeAttributes = 0x00000080
	TypeAttributes_Sealed                 TypeAttributes = 0x00000100
	TypeAttributes_SpecialName            TypeAttributes = 0x00000400
	TypeAttributes_Import                 TypeAttributes = 0x00001000
	TypeAttributes_Serializable           TypeAttributes = 0x00002000
	TypeAttributes_StringFormatMask       TypeAttributes = 0x00030000
	TypeAttributes_AnsiClass              TypeAttributes = 0x00000000
	TypeAttributes_UnicodeClass           TypeAttributes = 0x00010000
	TypeAttributes_AutoClass              TypeAttributes = 0x00020000
	TypeAttributes_CustomFormatClass      TypeAttributes = 0x00030000
	TypeAttributes_CustomStringFormatMask TypeAttributes = 0x00C00000
	TypeAttributes_BeforeFieldInit        TypeAttributes = 0x00100000
	TypeAttributes_RTSpecialName          TypeAttributes = 0x00000800
	TypeAttributes_HasSecurity            TypeAttributes = 0x00040000
	TypeAttributes_IsTypeForwarder        TypeAttributes = 0x00200000
)

// ElementType is defined in §II.23.1.16.
type ElementType uint8

const (
	ElementType_END          ElementType = 0x00
	ElementType_VOID         ElementType = 0x01
	ElementType_BOOLEAN      ElementType = 0x02
	ElementType_CHAR         ElementType = 0x03
	ElementType_I1           ElementType = 0x04
	ElementType_U1           ElementType = 0x05
	ElementType_I2           ElementType = 0x06
	ElementType_U2           ElementType = 0x07
	ElementType_I4           ElementType = 0x08
	ElementType_U4           ElementType = 0x09
	ElementType_I8           ElementType = 0x0a
	ElementType_U8           ElementType = 0x0b
	ElementType_R4           ElementType = 0x0c
	ElementType_R8           ElementType = 0x0d
	ElementType_STRING       ElementType = 0x0e
	ElementType_PTR          ElementType = 0x0f
	ElementType_BYREF        ElementType = 0x10
	ElementType_VALUETYPE    ElementType = 0x11
	ElementType_CLASS        ElementType = 0x12
	ElementType_VAR          ElementType = 0x13
	ElementType_ARRAY        ElementType = 0x14
	ElementType_GENERICINST  ElementType = 0x15
	ElementType_TYPEDBYREF   ElementType = 0x16
	ElementType_I            ElementType = 0x18
	ElementType_U            ElementType = 0x19
	ElementType_FNPTR        ElementType = 0x1b
	ElementType_OBJECT       ElementType = 0x1c
	ElementType_SZARRAY      ElementType = 0x1d
	ElementType_MVAR         ElementType = 0x1e
	ElementType_CMOD_REQD    ElementType = 0x1f
	ElementType_CMOD_OPT     ElementType = 0x20
	ElementType_INTERNAL     ElementType = 0x21
	ElementType_MODIFIER     ElementType = 0x40
	ElementType_SENTINEL     ElementType = 0x41
	ElementType_PINNED       ElementType = 0x45
	ElementType_TYPE         ElementType = 0x50
	ElementType_BOXED_OBJECT ElementType = 0x51
	ElementType_RESERVED     ElementType = 0x52
	ElementType_FIELD        ElementType = 0x53
	ElementType_PROPERTY     ElementType = 0x54
	ElementType_ENUM         ElementType = 0x55
)