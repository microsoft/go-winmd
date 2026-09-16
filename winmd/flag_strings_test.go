// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"fmt"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestFlagStrings(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name  string
		value fmt.Stringer
		want  string
	}{
		{"assembly-zero", winmd.AssemblyFlags(0), "0"},
		{"assembly-flags", winmd.AssemblyFlags_PublicKey | winmd.AssemblyFlags_Retargetable |
			winmd.AssemblyFlags_DisableJITcompileOptimizer | winmd.AssemblyFlags_EnableJITcompileTracking,
			"PublicKey|Retargetable|DisableJITcompileOptimizer|EnableJITcompileTracking"},
		{"assembly-unknown", winmd.AssemblyFlags_PublicKey | 0x80000002, "PublicKey|0x80000002"},
		{"event-zero", winmd.EventAttributes(0), "0"},
		{"event-flags", winmd.EventAttributes_SpecialName | winmd.EventAttributes_RTSpecialName, "SpecialName|RTSpecialName"},
		{"event-unknown", winmd.EventAttributes_SpecialName | 0x8001, "SpecialName|0x8001"},
		{"field-zero", winmd.FieldAttributes(0), "CompilerControlled"},
		{"field-flags", winmd.FieldAttributes(winmd.MemberAccess_Public) | winmd.FieldAttributes(winmd.FieldFlags_Static|winmd.FieldFlags_InitOnly), "Public|Static|InitOnly"},
		{"field-invalid-access", winmd.FieldAttributes(0x7) | winmd.FieldAttributes(winmd.FieldFlags_Static), "Static|0x7"},
		{"field-rva", winmd.FieldAttributes(winmd.MemberAccess_Public) | winmd.FieldAttributes(winmd.FieldFlags_HasFieldRVA|winmd.FieldFlags_HasDefault), "Public|HasFieldRVA|HasDefault"},
		{"field-unknown", winmd.FieldAttributes(winmd.MemberAccess_Private) | 0x4800, "Private|0x4800"},
		{"file-zero", winmd.FileAttributes(0), "ContainsMetaData"},
		{"file-no-metadata", winmd.FileAttributes(winmd.FileContent_ContainsNoMetaData), "ContainsNoMetaData"},
		{"file-unknown-default", winmd.FileAttributes(2), "ContainsMetaData|0x2"},
		{"file-unknown-high", winmd.FileAttributes(winmd.FileContent_ContainsNoMetaData) | 0x80000000, "ContainsNoMetaData|0x80000000"},
		{"generic-zero", winmd.GenericParamAttributes(0), "None"},
		{"generic-constraints", winmd.GenericParamAttributes(winmd.GenericVariance_Covariant) |
			winmd.GenericParamAttributes(winmd.GenericConstraints_ReferenceTypeConstraint|winmd.GenericConstraints_DefaultConstructorConstraint), "Covariant|ReferenceTypeConstraint|DefaultConstructorConstraint"},
		{"generic-invalid-variance", winmd.GenericParamAttributes(0x3), "0x3"},
		{"generic-constraint-bits", winmd.GenericParamAttributes(0x1c),
			"None|ReferenceTypeConstraint|NotNullableValueTypeConstraint|DefaultConstructorConstraint"},
		{"generic-unknown", winmd.GenericParamAttributes(winmd.GenericVariance_Contravariant) | 0x8000, "Contravariant|0x8000"},
		{"pinvoke-zero", winmd.PInvokeAttributes(0), "NotSpecified"},
		{"pinvoke-flags", winmd.PInvokeAttributes(winmd.PInvokeCharSet_Unicode) | winmd.PInvokeAttributes(winmd.PInvokeCallingConvention_Stdcall) |
			winmd.PInvokeAttributes(winmd.PInvokeFlags_NoMangle|winmd.PInvokeFlags_SupportsLastError),
			"Unicode|Stdcall|NoMangle|SupportsLastError"},
		{"pinvoke-charset-auto", winmd.PInvokeAttributes(0x6), "Auto"},
		{"pinvoke-invalid-callconv", winmd.PInvokeAttributes(0x600), "NotSpecified|0x600"},
		{"pinvoke-callconv-bits", winmd.PInvokeAttributes(0x700), "NotSpecified|0x700"},
		{"pinvoke-unknown", winmd.PInvokeAttributes(winmd.PInvokeCharSet_Ansi) | winmd.PInvokeAttributes(winmd.PInvokeCallingConvention_Cdecl) | 0x8000,
			"Ansi|Cdecl|0x8000"},
		{"resource-zero", winmd.ManifestResourceAttributes(0), "0"},
		{"resource-public", winmd.ManifestResourceAttributes(winmd.ResourceVisibility_Public), "Public"},
		{"resource-private", winmd.ManifestResourceAttributes(winmd.ResourceVisibility_Private), "Private"},
		{"resource-invalid-visibility", winmd.ManifestResourceAttributes(0x3), "0x3"},
		{"resource-visibility-bits", winmd.ManifestResourceAttributes(0x7), "0x7"},
		{"resource-unknown", winmd.ManifestResourceAttributes(winmd.ResourceVisibility_Private) | 0x80000000, "Private|0x80000000"},
		{"method-zero", winmd.MethodAttributes(0), "CompilerControlled|ReuseSlot"},
		{"method-flags", winmd.MethodAttributes(winmd.MemberAccess_Public) | winmd.MethodAttributes(winmd.MethodVtableLayout_NewSlot) |
			winmd.MethodAttributes(winmd.MethodFlags_Virtual|winmd.MethodFlags_Abstract), "Public|NewSlot|Virtual|Abstract"},
		{"method-static", winmd.MethodAttributes(winmd.MemberAccess_Public) | winmd.MethodAttributes(winmd.MethodFlags_Static), "Public|ReuseSlot|Static"},
		{"method-invalid-access", winmd.MethodAttributes(0x7), "ReuseSlot|0x7"},
		{"method-slot-bit", winmd.MethodAttributes(0x100), "CompilerControlled|NewSlot"},
		{"method-all-bits", winmd.MethodAttributes(0xffff),
			"NewSlot|UnmanagedExport|Static|Final|Virtual|HideBySig|Strict|Abstract|SpecialName|RTSpecialName|PInvokeImpl|HasSecurity|RequireSecObject|0x7"},
		{"methodimpl-zero", winmd.MethodImplAttributes(0), "IL|Managed"},
		{"methodimpl-flags", winmd.MethodImplAttributes(winmd.MethodCodeType_Native) | winmd.MethodImplAttributes(winmd.MethodManagedness_Unmanaged) |
			winmd.MethodImplAttributes(winmd.MethodImplFlags_PreserveSig), "Native|Unmanaged|PreserveSig"},
		{"methodimpl-code-bits", winmd.MethodImplAttributes(0x3), "Runtime|Managed"},
		{"methodimpl-managed-bit", winmd.MethodImplAttributes(0x4), "IL|Unmanaged"},
		{"methodimpl-all-bits", winmd.MethodImplAttributes(0xffff),
			"Runtime|Unmanaged|NoInlining|ForwardRef|Synchronized|NoOptimization|PreserveSig|InternalCall|0xef00"},
		{"methodimpl-unknown", winmd.MethodImplAttributes(winmd.MethodCodeType_OPTIL) | 0x8000, "OPTIL|Managed|0x8000"},
		{"semantics-zero", winmd.MethodSemanticsAttributes(0), "0"},
		{"semantics-accessors", winmd.MethodSemanticsAttributes_Getter | winmd.MethodSemanticsAttributes_Setter, "Setter|Getter"},
		{"semantics-all-flags", winmd.MethodSemanticsAttributes(0x3f), "Setter|Getter|Other|AddOn|RemoveOn|Fire"},
		{"semantics-unknown", winmd.MethodSemanticsAttributes_Getter | 0x8040, "Getter|0x8040"},
		{"param-zero", winmd.ParamAttributes(0), "0"},
		{"param-flags", winmd.ParamAttributes_In | winmd.ParamAttributes_Out | winmd.ParamAttributes_Optional |
			winmd.ParamAttributes_HasDefault | winmd.ParamAttributes_HasFieldMarshal, "In|Out|Optional|HasDefault|HasFieldMarshal"},
		{"param-unused-bits", winmd.ParamAttributes(0xcfe0), "0xcfe0"},
		{"param-unknown", winmd.ParamAttributes_In | 0x8004, "In|0x8004"},
		{"property-zero", winmd.PropertyAttributes(0), "0"},
		{"property-flags", winmd.PropertyAttributes_SpecialName | winmd.PropertyAttributes_RTSpecialName |
			winmd.PropertyAttributes_HasDefault, "SpecialName|RTSpecialName|HasDefault"},
		{"property-unused-bits", winmd.PropertyAttributes(0xe9ff), "0xe9ff"},
		{"property-all-bits", winmd.PropertyAttributes(0xffff), "SpecialName|RTSpecialName|HasDefault|0xe9ff"},
		{"type-zero", winmd.TypeAttributes(0), "NotPublic|AutoLayout|Class|AnsiClass"},
		{"type-flags", winmd.TypeAttributes(winmd.TypeVisibility_Public) | winmd.TypeAttributes(winmd.TypeLayout_SequentialLayout) |
			winmd.TypeAttributes(winmd.TypeSemantics_Interface) | winmd.TypeAttributes(winmd.TypeStringFormat_UnicodeClass) |
			winmd.TypeAttributes(winmd.TypeFlags_Abstract), "Public|SequentialLayout|Interface|UnicodeClass|Abstract"},
		{"type-nested", winmd.TypeAttributes(winmd.TypeVisibility_NestedFamANDAssem) | winmd.TypeAttributes(winmd.TypeLayout_ExplicitLayout) | winmd.TypeAttributes(winmd.TypeFlags_Sealed),
			"NestedFamANDAssem|ExplicitLayout|Class|AnsiClass|Sealed"},
		{"type-invalid-layout", winmd.TypeAttributes(0x18), "NotPublic|Class|AnsiClass|0x18"},
		{"type-custom-format", winmd.TypeAttributes(winmd.TypeStringFormat_CustomFormatClass) | 0xc00000,
			"NotPublic|AutoLayout|Class|CustomFormatClass|0xc00000"},
		{"type-high-bit", winmd.TypeAttributes(winmd.TypeVisibility_Public) | 0x80000000, "Public|AutoLayout|Class|AnsiClass|0x80000000"},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := test.value.String(); got != test.want {
				t.Fatalf("String() = %q; want %q", got, test.want)
			}
			if got := fmt.Sprint(test.value); got != test.want {
				t.Fatalf("fmt.Sprint = %q; want %q", got, test.want)
			}
		})
	}
}

func TestFlagStringMaskedChoices(t *testing.T) {
	t.Parallel()
	for access, name := range []string{"CompilerControlled", "Private", "FamANDAssem", "Assembly", "Family", "FamORAssem", "Public"} {
		f := winmd.FieldAttributes(access) | winmd.FieldAttributes(winmd.FieldFlags_Static)
		if got, want := f.String(), name+"|Static"; got != want {
			t.Errorf("field access %d = %q; want %q", access, got, want)
		}
	}
	for access, name := range []string{"CompilerControlled", "Private", "FamANDAssem", "Assembly", "Family", "FamORAssem", "Public"} {
		for slot, slotName := range []string{"ReuseSlot", "NewSlot"} {
			f := winmd.MethodAttributes(access|slot<<8) | winmd.MethodAttributes(winmd.MethodFlags_Virtual)
			if got, want := f.String(), name+"|"+slotName+"|Virtual"; got != want {
				t.Errorf("method access %d, slot %d = %q; want %q", access, slot, got, want)
			}
		}
	}
	for charset, charsetName := range []string{"NotSpecified", "Ansi", "Unicode", "Auto"} {
		for call, callName := range []string{"", "PlatformAPI", "Cdecl", "Stdcall", "Thiscall", "Fastcall", "0x600", "0x700"} {
			f := winmd.PInvokeAttributes(charset<<1 | call<<8)
			want := charsetName
			if callName != "" {
				want += "|" + callName
			}
			if got := f.String(); got != want {
				t.Errorf("charset %d, calling convention %d = %q; want %q", charset, call, got, want)
			}
		}
	}
	for kind, kindName := range []string{"IL", "Native", "OPTIL", "Runtime"} {
		for managed, managedName := range []string{"Managed", "Unmanaged"} {
			f := winmd.MethodImplAttributes(kind | managed<<2)
			if got, want := f.String(), kindName+"|"+managedName; got != want {
				t.Errorf("code kind %d, managedness %d = %q; want %q", kind, managed, got, want)
			}
		}
	}
	for visibility, visibilityName := range []string{"NotPublic", "Public", "NestedPublic", "NestedPrivate", "NestedFamily", "NestedAssembly", "NestedFamANDAssem", "NestedFamORAssem"} {
		for layout, layoutName := range []string{"AutoLayout", "SequentialLayout", "ExplicitLayout"} {
			for kind, kindName := range []string{"Class", "Interface"} {
				for format, formatName := range []string{"AnsiClass", "UnicodeClass", "AutoClass", "CustomFormatClass"} {
					f := winmd.TypeAttributes(visibility | layout<<3 | kind<<5 | format<<16)
					want := visibilityName + "|" + layoutName + "|" + kindName + "|" + formatName
					if got := f.String(); got != want {
						t.Errorf("type flags %#x = %q; want %q", uint32(f), got, want)
					}
				}
			}
		}
	}
}

func ExampleTypeAttributes_String() {
	// Raw metadata for a public, sequential-layout, sealed class.
	flags := winmd.TypeAttributes(0x109)
	fmt.Println(flags)
	// Output: Public|SequentialLayout|Class|AnsiClass|Sealed
}

func TestAssemblyHashAlgorithmString(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		value winmd.AssemblyHashAlgorithm
		want  string
	}{
		{0, "None"},
		{1, "AssemblyHashAlgorithm(1)"},
		{0x8002, "AssemblyHashAlgorithm(32770)"},
		{0x8003, "MD5"},
		{0x8004, "SHA1"},
		{0x8005, "AssemblyHashAlgorithm(32773)"},
		{0x8007, "AssemblyHashAlgorithm(32775)"},
		{0x80000000, "AssemblyHashAlgorithm(2147483648)"},
		{0xffffffff, "AssemblyHashAlgorithm(4294967295)"},
	} {
		if got := test.value.String(); got != test.want {
			t.Errorf("String(%#x) = %q; want %q", uint32(test.value), got, test.want)
		}
		if got := fmt.Sprint(test.value); got != test.want {
			t.Errorf("fmt.Sprint(%#x) = %q; want %q", uint32(test.value), got, test.want)
		}
	}
}

func TestElementTypeString(t *testing.T) {
	t.Parallel()
	names := [...]string{
		0x00: "END",
		0x01: "VOID",
		0x02: "BOOLEAN",
		0x03: "CHAR",
		0x04: "I1",
		0x05: "U1",
		0x06: "I2",
		0x07: "U2",
		0x08: "I4",
		0x09: "U4",
		0x0a: "I8",
		0x0b: "U8",
		0x0c: "R4",
		0x0d: "R8",
		0x0e: "STRING",
		0x0f: "PTR",
		0x10: "BYREF",
		0x11: "VALUETYPE",
		0x12: "CLASS",
		0x13: "VAR",
		0x14: "ARRAY",
		0x15: "GENERICINST",
		0x16: "TYPEDBYREF",
		0x18: "I",
		0x19: "U",
		0x1b: "FNPTR",
		0x1c: "OBJECT",
		0x1d: "SZARRAY",
		0x1e: "MVAR",
		0x1f: "CMOD_REQD",
		0x20: "CMOD_OPT",
		0x21: "INTERNAL",
		0x40: "MODIFIER",
		0x41: "SENTINEL",
		0x45: "PINNED",
		0x50: "TYPE",
		0x51: "BOXED_OBJECT",
		0x52: "RESERVED",
		0x53: "FIELD",
		0x54: "PROPERTY",
		0x55: "ENUM",
	}
	// Check every byte, including gaps, high bits, and composite encodings.
	for value := range 256 {
		want := fmt.Sprintf("ElementType(%d)", value)
		if value < len(names) && names[value] != "" {
			want = names[value]
		}
		if got := winmd.ElementType(value).String(); got != want {
			t.Errorf("String(%#x) = %q; want %q", value, got, want)
		}
		if got := fmt.Sprint(winmd.ElementType(value)); got != want {
			t.Errorf("fmt.Sprint(%#x) = %q; want %q", value, got, want)
		}
	}
}
