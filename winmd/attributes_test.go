// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"fmt"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func checkAttributeValue[T comparable](t *testing.T, name string, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("%s = %v; want %v", name, got, want)
	}
}

func TestAttributeAccessors(t *testing.T) {
	t.Parallel()
	// Decode raw words, including unknown bits, rather than constructing them
	// through the API under test. Typed getters must not change the original word.
	t.Run("field", func(t *testing.T) {
		a := winmd.FieldAttributes(0x4816)
		checkAttributeValue(t, "Access", a.Access(), winmd.MemberAccess_Public)
		checkAttributeValue(t, "Flags", a.Flags(), winmd.FieldFlags_Static)
		checkAttributeValue(t, "Bits", a.Bits(), uint16(0x4816))
		if !a.HasAll(0) || !a.HasAll(winmd.FieldFlags_Static) || a.HasAll(winmd.FieldFlags_Static|winmd.FieldFlags_InitOnly) || a.HasAll(winmd.FieldFlags(0x4800)) {
			t.Fatal("HasAll must test only known independent flags")
		}
	})
	t.Run("file", func(t *testing.T) {
		a := winmd.FileAttributes(0x80000001)
		checkAttributeValue(t, "Content", a.Content(), winmd.FileContent_ContainsNoMetaData)
		checkAttributeValue(t, "Bits", a.Bits(), uint32(0x80000001))
		checkAttributeValue(t, "zero Content", winmd.FileAttributes(0).Content(), winmd.FileContent_ContainsMetaData)
	})
	t.Run("generic-parameter", func(t *testing.T) {
		a := winmd.GenericParamAttributes(0x8015)
		want := winmd.GenericConstraints_ReferenceTypeConstraint | winmd.GenericConstraints_DefaultConstructorConstraint
		checkAttributeValue(t, "Variance", a.Variance(), winmd.GenericVariance_Covariant)
		checkAttributeValue(t, "Constraints", a.Constraints(), want)
		checkAttributeValue(t, "Bits", a.Bits(), uint16(0x8015))
		if !a.HasAll(0) || !a.HasAll(want) || a.HasAll(winmd.GenericConstraints_NotNullableValueTypeConstraint) || a.HasAll(winmd.GenericConstraints(0x8000)) {
			t.Fatal("HasAll must test only known constraints")
		}
	})
	t.Run("pinvoke", func(t *testing.T) {
		a := winmd.PInvokeAttributes(0x8345)
		want := winmd.PInvokeFlags_NoMangle | winmd.PInvokeFlags_SupportsLastError
		checkAttributeValue(t, "CharSet", a.CharSet(), winmd.PInvokeCharSet_Unicode)
		checkAttributeValue(t, "CallingConvention", a.CallingConvention(), winmd.PInvokeCallingConvention_Stdcall)
		checkAttributeValue(t, "Flags", a.Flags(), want)
		checkAttributeValue(t, "Bits", a.Bits(), uint16(0x8345))
		if !a.HasAll(0) || !a.HasAll(want) || a.HasAll(winmd.PInvokeFlags(0x8000)) {
			t.Fatal("HasAll must test only known independent flags")
		}
	})
	t.Run("resource", func(t *testing.T) {
		a := winmd.ManifestResourceAttributes(0x80000002)
		checkAttributeValue(t, "Visibility", a.Visibility(), winmd.ResourceVisibility_Private)
		checkAttributeValue(t, "Bits", a.Bits(), uint32(0x80000002))
	})
	t.Run("method", func(t *testing.T) {
		a := winmd.MethodAttributes(0x0546)
		want := winmd.MethodFlags_Virtual | winmd.MethodFlags_Abstract
		checkAttributeValue(t, "Access", a.Access(), winmd.MemberAccess_Public)
		checkAttributeValue(t, "VtableLayout", a.VtableLayout(), winmd.MethodVtableLayout_NewSlot)
		checkAttributeValue(t, "Flags", a.Flags(), want)
		checkAttributeValue(t, "Bits", a.Bits(), uint16(0x0546))
		if !a.HasAll(0) || !a.HasAll(want) || a.HasAll(winmd.MethodFlags_Final) || a.HasAll(winmd.MethodFlags(0x100)) {
			t.Fatal("HasAll must not treat the slot choice as an independent flag")
		}
	})
	t.Run("method-implementation", func(t *testing.T) {
		a := winmd.MethodImplAttributes(0x8085)
		checkAttributeValue(t, "CodeType", a.CodeType(), winmd.MethodCodeType_Native)
		checkAttributeValue(t, "Managedness", a.Managedness(), winmd.MethodManagedness_Unmanaged)
		checkAttributeValue(t, "Flags", a.Flags(), winmd.MethodImplFlags_PreserveSig)
		checkAttributeValue(t, "Bits", a.Bits(), uint16(0x8085))
		if !a.HasAll(0) || !a.HasAll(winmd.MethodImplFlags_PreserveSig) || a.HasAll(winmd.MethodImplFlags_NoInlining) || a.HasAll(winmd.MethodImplFlags(0x8000)) {
			t.Fatal("HasAll must test only known independent flags")
		}
	})
	t.Run("type", func(t *testing.T) {
		a := winmd.TypeAttributes(0x80c10131)
		checkAttributeValue(t, "Visibility", a.Visibility(), winmd.TypeVisibility_Public)
		checkAttributeValue(t, "Layout", a.Layout(), winmd.TypeLayout_ExplicitLayout)
		checkAttributeValue(t, "Semantics", a.Semantics(), winmd.TypeSemantics_Interface)
		checkAttributeValue(t, "StringFormat", a.StringFormat(), winmd.TypeStringFormat_UnicodeClass)
		checkAttributeValue(t, "Flags", a.Flags(), winmd.TypeFlags_Sealed)
		checkAttributeValue(t, "Bits", a.Bits(), uint32(0x80c10131))
		if !a.HasAll(0) || !a.HasAll(winmd.TypeFlags_Sealed) || a.HasAll(winmd.TypeFlags_Abstract) || a.HasAll(winmd.TypeFlags(0x80c00000)) {
			t.Fatal("HasAll must not expose unknown or custom string-format bits as flags")
		}
	})
}

func TestAttributeChoiceEncodings(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		value fmt.Stringer
		text  string
	}{
		{winmd.MemberAccess_Public, "Public"},
		{winmd.MemberAccess(0x8001), "0x8001"},
		{winmd.FileContent_ContainsMetaData, "ContainsMetaData"},
		{winmd.FileContent(2), "0x2"},
		{winmd.GenericVariance_Contravariant, "Contravariant"},
		{winmd.GenericVariance(3), "0x3"},
		{winmd.PInvokeCharSet_Unicode, "Unicode"},
		{winmd.PInvokeCharSet(1), "0x1"},
		{winmd.PInvokeCallingConvention_PlatformAPI, "PlatformAPI"},
		{winmd.PInvokeCallingConvention(0), "0"},
		{winmd.ResourceVisibility_Private, "Private"},
		{winmd.ResourceVisibility(0), "0"},
		{winmd.MethodVtableLayout_NewSlot, "NewSlot"},
		{winmd.MethodVtableLayout(0x101), "0x101"},
		// This is a named encoding, even though ECMA reserves its use.
		{winmd.MethodCodeType_OPTIL, "OPTIL"},
		{winmd.MethodCodeType(4), "0x4"},
		{winmd.MethodManagedness_Managed, "Managed"},
		{winmd.MethodManagedness(1), "0x1"},
		{winmd.TypeVisibility_NestedPrivate, "NestedPrivate"},
		{winmd.TypeVisibility(0x80000001), "0x80000001"},
		{winmd.TypeLayout_ExplicitLayout, "ExplicitLayout"},
		{winmd.TypeLayout(0x18), "0x18"},
		{winmd.TypeSemantics_Class, "Class"},
		{winmd.TypeSemantics(0x21), "0x21"},
		{winmd.TypeStringFormat_CustomFormatClass, "CustomFormatClass"},
		{winmd.TypeStringFormat(1), "0x1"},
		{winmd.FieldAttributes(0x8017).Access(), "0x7"},
		{winmd.MethodAttributes(0x17).Access(), "0x7"},
		{winmd.GenericParamAttributes(0x8003).Variance(), "0x3"},
		{winmd.PInvokeAttributes(0x8740).CallingConvention(), "0x700"},
		{winmd.ManifestResourceAttributes(0x80000003).Visibility(), "0x3"},
		{winmd.TypeAttributes(0xc00018).Layout(), "0x18"},
	} {
		t.Run(fmt.Sprintf("%T/%s", test.value, test.text), func(t *testing.T) {
			checkAttributeValue(t, "String", test.value.String(), test.text)
		})
	}
	for _, v := range []winmd.TypeVisibility{0, 1, 2, 3, 4, 5, 6, 7, 8, 0xffffffff} {
		if got, want := v.IsNested(), v >= 2 && v <= 7; got != want {
			t.Errorf("IsNested(%#x) = %v; want %v", uint32(v), got, want)
		}
	}
}

func ExampleTypeAttributes_Visibility() {
	// This is the raw word returned by the metadata table decoder.
	attrs := winmd.TypeAttributes(0x8000010a)
	fmt.Println(attrs.Visibility(), attrs.Visibility().IsNested())
	fmt.Println(attrs.Layout())
	fmt.Println(attrs.HasAll(winmd.TypeFlags_Sealed))
	fmt.Printf("%#x\n", attrs.Bits())
	// Output:
	// NestedPublic true
	// SequentialLayout
	// true
	// 0x8000010a
}
