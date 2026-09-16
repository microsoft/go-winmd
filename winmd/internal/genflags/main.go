// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// genflags generates checked-in flag and enum lookup tables from tables.go and
// sigtypes.go. It uses explicit mask rules and Go syntax only, without loading
// or type-checking the winmd package. The schema's constants must be explicitly
// typed integer literals; other forms are rejected rather than evaluated or
// guessed. It also generates read-only attribute accessors and bitset helpers.
package main

import (
	"bytes"
	"cmp"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"math/bits"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

const outputFile = "zflag.go"

func main() {
	log.SetFlags(0)
	check := flag.Bool("check", false, "check the generated file without writing it")
	flag.Parse()
	if flag.NArg() != 0 {
		log.Fatal("genflags takes no positional arguments")
	}
	src, err := generate()
	if err != nil {
		log.Fatal(err)
	}
	if *check {
		current, err := os.ReadFile(outputFile)
		if err != nil {
			log.Fatal(err)
		}
		if !bytes.Equal(bytes.ReplaceAll(current, []byte("\r\n"), []byte("\n")), src) {
			log.Fatalf("%s is out of date; run go generate", outputFile)
		}
		return
	}
	if err := os.WriteFile(outputFile, src, 0644); err != nil {
		log.Fatal(err)
	}
}

type namedConstant struct {
	name  string
	value uint64
}

type flagType struct {
	name      string
	width     int
	enum      bool
	constants []namedConstant
}

func readConstants() ([]flagType, error) {
	fset := token.NewFileSet()
	errorf := func(pos token.Pos, format string, args ...any) error {
		return fmt.Errorf("%s: %s", fset.Position(pos), fmt.Sprintf(format, args...))
	}
	var decls []ast.Decl
	for _, name := range []string{"tables.go", "sigtypes.go"} {
		file, err := parser.ParseFile(fset, name, nil, parser.SkipObjectResolution)
		if err != nil {
			return nil, err
		}
		if file.Name.Name != "winmd" {
			return nil, errorf(file.Name.Pos(), "expected package winmd")
		}
		decls = append(decls, file.Decls...)
	}
	// Composition rules identify the choice domains; pure flag sets are still
	// discovered by name. No inventory of individual constants is needed.
	choices := make(map[string]bool)
	required := make(map[string]bool)
	for _, name := range standaloneEnums {
		required[name], choices[name] = true, true
	}
	maskOwners := make(map[string]string)
	for name, rule := range flagRules {
		required[name] = true
		if rule.flags != "" {
			required[rule.flags] = true
		}
		for _, field := range rule.fields {
			required[field.typ], choices[field.typ] = true, true
			if _, ok := maskOwners[field.mask]; ok {
				return nil, fmt.Errorf("mask %s is used more than once", field.mask)
			}
			maskOwners[field.mask] = name
		}
	}
	// Preserve type declaration order for deterministic output.
	var types []flagType
	byName := make(map[string]int)
	for _, decl := range decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			continue
		}
		for _, spec := range decl.Specs {
			spec := spec.(*ast.TypeSpec)
			name := spec.Name.Name
			if !required[name] && !strings.HasSuffix(name, "Flags") && !strings.HasSuffix(name, "Attributes") {
				continue
			}
			if _, ok := byName[name]; ok {
				return nil, errorf(spec.Pos(), "duplicate flag type %s", name)
			}
			var width int
			if base, ok := ast.Unparen(spec.Type).(*ast.Ident); ok && !spec.Assign.IsValid() && spec.TypeParams == nil {
				switch base.Name {
				case "uint8":
					width = 8
				case "uint16":
					width = 16
				case "uint32":
					width = 32
				}
			}
			if width == 0 {
				return nil, errorf(spec.Pos(), "%s must be a uint8, uint16, or uint32 type definition", name)
			}
			byName[name] = len(types)
			types = append(types, flagType{name: name, width: width, enum: choices[name]})
		}
	}
	if len(types) == 0 {
		return nil, fmt.Errorf("no flag types found")
	}
	for name := range required {
		if _, ok := byName[name]; !ok {
			return nil, fmt.Errorf("rule refers to missing flag type %s", name)
		}
	}
	seen := make(map[string]bool)
	for _, decl := range decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok || decl.Tok != token.CONST {
			continue
		}
		for _, spec := range decl.Specs {
			spec := spec.(*ast.ValueSpec)
			typeName := ""
			if typ, ok := spec.Type.(*ast.Ident); ok {
				typeName = typ.Name
			}
			for _, name := range spec.Names {
				if owner, ok := maskOwners[name.Name]; ok && typeName != owner {
					return nil, errorf(name.Pos(), "%s must have explicit type %s", name.Name, owner)
				}
				prefix, _, _ := strings.Cut(name.Name, "_")
				if _, ok := byName[prefix]; ok && typeName != prefix {
					return nil, errorf(name.Pos(), "%s must have explicit type %s", name.Name, prefix)
				}
			}
			index, ok := byName[typeName]
			if !ok {
				continue
			}
			// Only public enum/flag values and masks named by composition rules
			// participate. Other private helper constants are not display values.
			if len(spec.Names) == 1 && !spec.Names[0].IsExported() && maskOwners[spec.Names[0].Name] == "" {
				continue
			}
			if len(spec.Names) != 1 || len(spec.Values) != 1 {
				return nil, errorf(spec.Pos(), "flag constants must have one name and one integer literal")
			}
			name := spec.Names[0].Name
			_, raw := flagRules[typeName]
			if raw && maskOwners[name] != typeName {
				return nil, errorf(spec.Pos(), "declare %s in an enum or independent flag type, not %s", name, typeName)
			}
			if suffix, ok := strings.CutPrefix(name, typeName+"_"); (!raw && (!ok || suffix == "")) || seen[name] {
				return nil, errorf(spec.Pos(), "invalid or duplicate flag constant %s", name)
			}
			literal, ok := ast.Unparen(spec.Values[0]).(*ast.BasicLit)
			if !ok || literal.Kind != token.INT {
				return nil, errorf(spec.Pos(), "%s must use an integer literal", name)
			}
			value, err := strconv.ParseUint(literal.Value, 0, types[index].width)
			if err != nil {
				return nil, errorf(spec.Pos(), "%s: %v", name, err)
			}
			seen[name] = true
			types[index].constants = append(types[index].constants, namedConstant{name, value})
		}
	}
	return types, nil
}

type entry struct{ mask, value, name string }

func resolveRule(name string, types map[string]flagType) (flagRule, bool) {
	rule, raw := flagRules[name]
	if raw && rule.flags == "" {
		if owner, ok := strings.CutSuffix(name, "Attributes"); ok {
			if _, exists := types[owner+"Flags"]; exists {
				rule.flags = owner + "Flags"
			}
		}
	}
	return rule, raw
}

func orderedValues(typ flagType) ([]namedConstant, error) {
	values := slices.Clone(typ.constants)
	if len(values) == 0 {
		return nil, fmt.Errorf("no named values for %s", typ.name)
	}
	slices.SortFunc(values, func(a, b namedConstant) int { return cmp.Compare(a.value, b.value) })
	for i, c := range values {
		if i != 0 && values[i-1].value == c.value {
			return nil, fmt.Errorf("ambiguous values %s and %s", values[i-1].name, c.name)
		}
		if !typ.enum && (c.value == 0 || c.value&(c.value-1) != 0) {
			return nil, fmt.Errorf("%s is not an independent single-bit flag", c.name)
		}
	}
	return values, nil
}

func tableEntries(typ flagType, types map[string]flagType) ([]entry, error) {
	rule, raw := resolveRule(typ.name, types)
	if !raw {
		values, err := orderedValues(typ)
		if err != nil {
			return nil, err
		}
		var entries []entry
		for _, c := range values {
			mask := c.name
			if typ.enum {
				// A standalone choice must match exactly, even if an unknown
				// encoding has bits outside its field in the raw metadata word.
				mask = "^" + typ.name + "(0)"
			}
			entries = append(entries, entry{mask, c.name, strings.TrimPrefix(c.name, typ.name+"_")})
		}
		return entries, nil
	}
	type resolvedField struct {
		mask namedConstant
		typ  flagType
	}
	var fields []resolvedField
	var fieldBits uint64
	for _, field := range rule.fields {
		var mask namedConstant
		for _, c := range typ.constants {
			if c.name == field.mask {
				mask = c
			}
		}
		choice := types[field.typ]
		if mask.value == 0 || mask.value&fieldBits != 0 || choice.width != typ.width || !choice.enum {
			return nil, fmt.Errorf("invalid or overlapping field %s for %s", field.mask, typ.name)
		}
		fieldBits |= mask.value
		fields = append(fields, resolvedField{mask, choice})
	}
	slices.SortFunc(fields, func(a, b resolvedField) int {
		return cmp.Compare(bits.TrailingZeros64(a.mask.value), bits.TrailingZeros64(b.mask.value))
	})
	var entries []entry
	for _, field := range fields {
		values, err := orderedValues(field.typ)
		if err != nil {
			return nil, err
		}
		for _, c := range values {
			if c.value&^field.mask.value != 0 {
				return nil, fmt.Errorf("%s does not fit %s", c.name, field.mask.name)
			}
			entries = append(entries, entry{field.mask.name, typ.name + "(" + c.name + ")", strings.TrimPrefix(c.name, field.typ.name+"_")})
		}
	}
	if rule.flags != "" {
		flags := types[rule.flags]
		if flags.width != typ.width || flags.enum {
			return nil, fmt.Errorf("invalid independent flag type %s for %s", rule.flags, typ.name)
		}
		values, err := orderedValues(flags)
		if err != nil {
			return nil, err
		}
		for _, c := range values {
			if c.value&fieldBits != 0 {
				return nil, fmt.Errorf("%s overlaps an exclusive field in %s", c.name, typ.name)
			}
			value := typ.name + "(" + c.name + ")"
			entries = append(entries, entry{value, value, strings.TrimPrefix(c.name, flags.name+"_")})
		}
	}
	return entries, nil
}

func writeAttributeHelpers(out *bytes.Buffer, typ flagType, rule flagRule, types map[string]flagType) error {
	names := map[string]bool{"String": true, "Bits": true, "HasAll": true}
	checkName := func(name string) error {
		if !token.IsIdentifier(name) || !ast.IsExported(name) || names[name] {
			return fmt.Errorf("invalid or duplicate accessor %s.%s", typ.name, name)
		}
		names[name] = true
		return nil
	}
	fmt.Fprintln(out, "\n// Bits returns the complete metadata word, including unknown bits.")
	fmt.Fprintf(out, "func (a %s) Bits() uint%d { return uint%d(a) }\n", typ.name, typ.width, typ.width)
	for _, field := range rule.fields {
		stem, ok := strings.CutSuffix(field.mask, "Mask")
		start := strings.IndexFunc(stem, unicode.IsUpper)
		if !ok || start <= 0 {
			return fmt.Errorf("cannot derive accessor from mask %s", field.mask)
		}
		accessor := stem[start:]
		if err := checkName(accessor); err != nil {
			return err
		}
		fmt.Fprintf(out, "\n// %s returns the encoded %s choice, including unnamed values.\n", accessor, field.typ)
		fmt.Fprintf(out, "func (a %s) %s() %s { return %s(a & %s) }\n", typ.name, accessor, field.typ, field.typ, field.mask)
	}
	if rule.flags != "" {
		var accessor string
		switch {
		case strings.HasSuffix(rule.flags, "Flags"):
			accessor = "Flags"
		case strings.HasSuffix(rule.flags, "Constraints"):
			accessor = "Constraints"
		default:
			return fmt.Errorf("cannot derive flag accessor from type %s", rule.flags)
		}
		if err := checkName(accessor); err != nil {
			return err
		}
		values, err := orderedValues(types[rule.flags])
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "\nconst flagMask%s = ", typ.name)
		for i, c := range values {
			if i != 0 {
				fmt.Fprint(out, " | ")
			}
			fmt.Fprint(out, c.name)
		}
		fmt.Fprintln(out)
		fmt.Fprintf(out, "\n// %s returns only the known independent %s bits.\n", accessor, rule.flags)
		fmt.Fprintf(out, "func (a %s) %s() %s { return %s(a) & flagMask%s }\n", typ.name, accessor, rule.flags, rule.flags, typ.name)
		fmt.Fprintln(out, "\n// HasAll reports whether all requested known independent bits are set. HasAll(0) is true.")
		fmt.Fprintf(out, "func (a %s) HasAll(v %s) bool { return a.%s()&v == v }\n", typ.name, rule.flags, accessor)
	}
	return nil
}

func generate() ([]byte, error) {
	types, err := readConstants()
	if err != nil {
		return nil, err
	}
	byName := make(map[string]flagType)
	for _, typ := range types {
		byName[typ.name] = typ
	}
	var out bytes.Buffer
	fmt.Fprintln(&out, "// Copyright (c) Microsoft Corporation.")
	fmt.Fprintln(&out, "// Licensed under the MIT License.")
	fmt.Fprintln(&out, "\n// Code generated by \"genflags\"; DO NOT EDIT.")
	fmt.Fprintln(&out, "\npackage winmd")
	for _, typ := range types {
		entries, err := tableEntries(typ, byName)
		if err != nil {
			return nil, err
		}
		if slices.Contains(standaloneEnums, typ.name) {
			fmt.Fprintln(&out, "\n// String returns the enum name, or Type(decimal) for an unnamed value.")
			fmt.Fprintf(&out, "func (v %s) String() string {\n\treturn formatEnum(v, flagNames%s[:], %q)\n}\n", typ.name, typ.name, typ.name)
		} else {
			fmt.Fprintln(&out, "\n// String returns flag and masked-field names, with unknown bits in hexadecimal.")
			fmt.Fprintf(&out, "func (f %s) String() string {\n\treturn formatFlags(f, flagNames%s[:])\n}\n", typ.name, typ.name)
		}
		if rule, raw := resolveRule(typ.name, byName); raw {
			if err := writeAttributeHelpers(&out, typ, rule, byName); err != nil {
				return nil, err
			}
		} else if !typ.enum {
			fmt.Fprintln(&out, "\n// HasAll reports whether all requested bits are set. HasAll(0) is true.")
			fmt.Fprintf(&out, "func (f %s) HasAll(v %s) bool { return f&v == v }\n", typ.name, typ.name)
			fmt.Fprintln(&out, "\n// Bits returns the raw bitset, including unknown bits.")
			fmt.Fprintf(&out, "func (f %s) Bits() uint%d { return uint%d(f) }\n", typ.name, typ.width, typ.width)
		}
		fmt.Fprintf(&out, "\nvar flagNames%s = [...]flagName[%s]{\n", typ.name, typ.name)
		for _, e := range entries {
			fmt.Fprintf(&out, "\t{%s, %s, %q},\n", e.mask, e.value, e.name)
		}
		fmt.Fprintln(&out, "}")
	}
	return format.Source(out.Bytes())
}
