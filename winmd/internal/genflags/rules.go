// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

// Standalone enums are not fields within an attribute word. Their names omit
// the usual Type_ prefix, and unnamed values use Type(decimal).
var standaloneEnums = []string{
	"AssemblyHashAlgorithm",
	"ElementType",
}

type flagRule struct {
	fields []maskedField
	flags  string // Optional override for the inferred <owner>Flags type.
}

type maskedField struct {
	mask, typ string
}

// Rules describe how a raw metadata word is composed from typed choices and
// independent flags. Choices (including zero defaults) now come from their own
// enum types, so there are no per-value rules or exclusions. Mask names refer to
// private constants in tables.go; output order is inferred from their bits.
// Accessors follow <lowercase owner><Accessor>Mask and the Flags/Constraints
// suffix of a bitset type. Ordinary <owner>Flags types are inferred when present.
var flagRules = map[string]flagRule{
	"FieldAttributes": {
		fields: []maskedField{{"fieldAccessMask", "MemberAccess"}},
	},
	"FileAttributes": {
		fields: []maskedField{{"fileContentMask", "FileContent"}},
	},
	"GenericParamAttributes": {
		fields: []maskedField{{"genericVarianceMask", "GenericVariance"}},
		flags:  "GenericConstraints",
	},
	"PInvokeAttributes": {
		fields: []maskedField{
			{"pinvokeCharSetMask", "PInvokeCharSet"},
			{"pinvokeCallingConventionMask", "PInvokeCallingConvention"},
		},
	},
	"ManifestResourceAttributes": {
		fields: []maskedField{{"resourceVisibilityMask", "ResourceVisibility"}},
	},
	"MethodAttributes": {
		fields: []maskedField{
			{"methodAccessMask", "MemberAccess"},
			{"methodVtableLayoutMask", "MethodVtableLayout"},
		},
	},
	"MethodImplAttributes": {
		fields: []maskedField{
			{"methodCodeTypeMask", "MethodCodeType"},
			{"methodManagednessMask", "MethodManagedness"},
		},
	},
	"TypeAttributes": {
		fields: []maskedField{
			{"typeVisibilityMask", "TypeVisibility"},
			{"typeLayoutMask", "TypeLayout"},
			{"typeSemanticsMask", "TypeSemantics"},
			{"typeStringFormatMask", "TypeStringFormat"},
		},
	},
}
