// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const tablesFile = "tables.go"

// These are the schema inputs, not the files needed to compile the package.
// Imports, function bodies, and generated declarations are not resolved.
var schemaFiles = [...]string{tablesFile, "sigtypes.go", "winmd.go", "tags.go"}

type columnType string

const (
	columnTypeIndex      columnType = "columnTypeIndex"
	columnTypeUint       columnType = "columnTypeUint"
	columnTypeString     columnType = "columnTypeString"
	columnTypeGUID       columnType = "columnTypeGUID"
	columnTypeBlob       columnType = "columnTypeBlob"
	columnTypeCodedIndex columnType = "columnTypeCodedIndex"
	columnTypeSlice      columnType = "columnTypeSlice"
)

type tableInfo struct {
	name      string
	tableName string
	exported  bool
	code      uint8
	fields    []columnInfo
}

type columnInfo struct {
	name       string
	typeName   string
	size       int
	needsCast  bool
	columnType columnType
	tableName  string
	coded      string
	nullable   bool
}

type schema struct {
	fset      *token.FileSet
	types     map[string]*ast.TypeSpec
	tables    *ast.File
	codedTags map[string]bool
}

func readTables(dir string) ([]tableInfo, error) {
	s := &schema{fset: token.NewFileSet(), types: make(map[string]*ast.TypeSpec)}
	for _, name := range schemaFiles {
		file, err := parser.ParseFile(s.fset, filepath.Join(dir, name), nil, parser.ParseComments|parser.SkipObjectResolution)
		if err != nil {
			return nil, err
		}
		if err := s.addFile(file); err != nil {
			return nil, err
		}
		if name == tablesFile {
			s.tables = file
		}
	}
	tables, err := s.parseTables()
	if err != nil {
		return nil, err
	}
	const want = 38
	if len(tables) != want {
		return nil, fmt.Errorf("got %d tables, want %d", len(tables), want)
	}
	return tables, nil
}

func (s *schema) errorf(pos token.Pos, format string, args ...any) error {
	return fmt.Errorf("%s: %s", s.fset.Position(pos), fmt.Sprintf(format, args...))
}

func (s *schema) addFile(file *ast.File) error {
	if file.Name.Name != "winmd" {
		return s.errorf(file.Name.Pos(), "schema file must belong to package winmd")
	}
	for _, decl := range file.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			continue
		}
		for _, spec := range decl.Specs {
			spec := spec.(*ast.TypeSpec)
			name := spec.Name.Name
			if _, ok := s.types[name]; ok {
				return s.errorf(spec.Pos(), "duplicate type %s", name)
			}
			if uintSize(name) != 0 {
				return s.errorf(spec.Pos(), "schema type shadows built-in type %s", name)
			}
			s.types[name] = spec
		}
	}
	return nil
}

func (s *schema) parseTables() ([]tableInfo, error) {
	if err := s.readCodedTags(); err != nil {
		return nil, err
	}
	var tables []tableInfo
	codes := make(map[uint8]bool)
	names := make(map[string]bool)
	tableTypes := make(map[string]bool)
	for _, decl := range s.tables.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		_, annotated, err := s.annotation("@table", decl.Doc)
		if err != nil {
			return nil, err
		}
		if annotated && (decl.Tok != token.TYPE || len(decl.Specs) != 1) {
			return nil, s.errorf(decl.Pos(), "@table must annotate an individual type declaration")
		}
		if decl.Tok != token.TYPE {
			continue
		}
		for _, spec := range decl.Specs {
			spec := spec.(*ast.TypeSpec)
			code, ok, err := s.annotation("@table", decl.Doc, spec.Doc)
			if err != nil {
				return nil, err
			}
			if !ok {
				continue
			}
			value, err := strconv.ParseUint(code, 0, 8)
			if err != nil {
				return nil, s.errorf(spec.Pos(), "invalid @table code %q: %v", code, err)
			}
			info, err := s.parseTable(spec)
			if err != nil {
				return nil, err
			}
			info.code = uint8(value)
			if codes[info.code] {
				return nil, s.errorf(spec.Pos(), "table code %#x is duplicated", info.code)
			}
			if names[info.tableName] {
				return nil, s.errorf(spec.Pos(), "generated table name %s is duplicated", info.tableName)
			}
			codes[info.code], names[info.tableName] = true, true
			tableTypes[info.name] = true
			tables = append(tables, info)
		}
	}
	for i := range tables {
		for j := range tables[i].fields {
			field := &tables[i].fields[j]
			if field.tableName != "" {
				if !tableTypes[field.tableName] {
					return nil, s.errorf(s.types[tables[i].name].Pos(), "%s.%s refers to unknown table %s", tables[i].name, field.name, field.tableName)
				}
				field.tableName = tableName(field.tableName)
			}
		}
	}
	return tables, nil
}

func (s *schema) parseTable(spec *ast.TypeSpec) (tableInfo, error) {
	st, ok := ast.Unparen(spec.Type).(*ast.StructType)
	if !ok || spec.Assign.IsValid() || spec.TypeParams != nil {
		return tableInfo{}, s.errorf(spec.Pos(), "@table type %s must be a non-generic struct definition", spec.Name.Name)
	}
	if spec.Name.Name == "_" || len(st.Fields.List) == 0 {
		return tableInfo{}, s.errorf(spec.Pos(), "table must have a name and at least one field")
	}
	info := tableInfo{name: spec.Name.Name, tableName: tableName(spec.Name.Name), exported: spec.Name.IsExported()}
	fields := make(map[string]bool)
	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			return tableInfo{}, s.errorf(field.Pos(), "%s: embedded table fields are not supported", info.name)
		}
		col, err := s.column(field.Type, make(map[string]bool))
		if err != nil {
			return tableInfo{}, s.errorf(field.Pos(), "%s.%s: %v", info.name, field.Names[0].Name, err)
		}
		ref, hasRef, err := s.annotation("@ref", field.Doc, field.Comment)
		if err != nil {
			return tableInfo{}, err
		}
		if col.columnType == columnTypeIndex || col.columnType == columnTypeSlice {
			if !hasRef || !token.IsIdentifier(ref) || ref == "_" {
				return tableInfo{}, s.errorf(field.Pos(), "%s.%s requires @ref=<table name>", info.name, field.Names[0].Name)
			}
			// Keep the exact source name until all annotated tables are known.
			col.tableName = ref
		} else if hasRef {
			return tableInfo{}, s.errorf(field.Pos(), "@ref is only supported on Index and Slice fields")
		}
		nullable, hasNullable, err := s.annotation("@nullable", field.Doc, field.Comment)
		if err != nil {
			return tableInfo{}, err
		}
		if hasNullable {
			if col.columnType != columnTypeCodedIndex {
				return tableInfo{}, s.errorf(field.Pos(), "@nullable is only supported on CodedIndex fields")
			}
			if nullable != "true" && nullable != "false" {
				return tableInfo{}, s.errorf(field.Pos(), "invalid @nullable value %q; want true or false", nullable)
			}
			col.nullable = nullable == "true"
		}
		for _, name := range field.Names {
			if name.Name == "_" || fields[name.Name] {
				return tableInfo{}, s.errorf(name.Pos(), "blank or duplicate table field %s", name.Name)
			}
			fields[name.Name] = true
			col.name = name.Name
			info.fields = append(info.fields, col)
		}
	}
	return info, nil
}

// column recognizes the schema's supported Go type shapes. Named integer
// widths and aliases are resolved from declarations, not from a name/size list.
func (s *schema) column(expr ast.Expr, visiting map[string]bool) (columnInfo, error) {
	switch expr := ast.Unparen(expr).(type) {
	case *ast.Ident:
		name := expr.Name
		if size := uintSize(name); size != 0 {
			return columnInfo{columnType: columnTypeUint, typeName: name, size: size}, nil
		}
		spec := s.types[name]
		if spec == nil {
			return columnInfo{}, fmt.Errorf("unknown or unsupported type %s", name)
		}
		switch name {
		case "Index":
			return columnInfo{columnType: columnTypeIndex}, nil
		case "Slice":
			return columnInfo{columnType: columnTypeSlice}, nil
		case "String":
			return columnInfo{columnType: columnTypeString}, nil
		}
		if spec.TypeParams != nil {
			return columnInfo{}, fmt.Errorf("generic type %s requires a supported instantiation", name)
		}
		if visiting[name] {
			return columnInfo{}, fmt.Errorf("cyclic type definition involving %s", name)
		}
		visiting[name] = true
		defer delete(visiting, name)
		col, err := s.column(spec.Type, visiting)
		if err != nil {
			return columnInfo{}, err
		}
		if col.columnType == columnTypeUint {
			col.typeName = name
			col.needsCast = col.needsCast || !spec.Assign.IsValid()
		} else if !spec.Assign.IsValid() && col.columnType != columnTypeBlob && col.columnType != columnTypeGUID {
			return columnInfo{}, fmt.Errorf("defined wrapper %s of %s is not supported; use a type alias", name, col.columnType)
		}
		return col, nil
	case *ast.ArrayType:
		element, err := s.column(expr.Elt, visiting)
		if err != nil || element.columnType != columnTypeUint || element.size != 1 || element.needsCast {
			return columnInfo{}, fmt.Errorf("only byte slices and [16]byte arrays are supported")
		}
		if expr.Len == nil {
			return columnInfo{columnType: columnTypeBlob}, nil
		}
		if length, ok := ast.Unparen(expr.Len).(*ast.BasicLit); ok && length.Kind == token.INT {
			if value, err := strconv.ParseUint(length.Value, 0, 64); err == nil && value == 16 {
				return columnInfo{columnType: columnTypeGUID}, nil
			}
		}
		return columnInfo{}, fmt.Errorf("GUID arrays must have literal length 16")
	case *ast.IndexExpr:
		return s.codedColumn(expr.X, []ast.Expr{expr.Index})
	case *ast.IndexListExpr:
		return s.codedColumn(expr.X, expr.Indices)
	default:
		return columnInfo{}, fmt.Errorf("unsupported field type syntax %T", expr)
	}
}

func uintSize(name string) int {
	switch name {
	case "byte", "uint8":
		return 1
	case "uint16":
		return 2
	case "uint32":
		return 4
	}
	return 0
}

func (s *schema) codedColumn(base ast.Expr, args []ast.Expr) (columnInfo, error) {
	name, ok := ast.Unparen(base).(*ast.Ident)
	if !ok || name.Name != "CodedIndex" || len(args) != 1 {
		return columnInfo{}, fmt.Errorf("only CodedIndex with exactly one type argument is supported")
	}
	tag, err := s.namedType(args[0], make(map[string]bool))
	if err != nil {
		return columnInfo{}, err
	}
	if !s.codedTags[tag.Name.Name] {
		return columnInfo{}, fmt.Errorf("%s is not a CodedTag", tag.Name.Name)
	}
	return columnInfo{columnType: columnTypeCodedIndex, coded: tag.Name.Name}, nil
}

// namedType follows aliases to the declared name used by a coded tag family.
func (s *schema) namedType(expr ast.Expr, visiting map[string]bool) (*ast.TypeSpec, error) {
	name, ok := ast.Unparen(expr).(*ast.Ident)
	if !ok || s.types[name.Name] == nil {
		return nil, fmt.Errorf("coded tag must name a declared type")
	}
	if visiting[name.Name] {
		return nil, fmt.Errorf("cyclic type alias involving %s", name.Name)
	}
	spec := s.types[name.Name]
	if !spec.Assign.IsValid() {
		return spec, nil
	}
	visiting[name.Name] = true
	defer delete(visiting, name.Name)
	return s.namedType(spec.Type, visiting)
}

func (s *schema) readCodedTags() error {
	spec := s.types["CodedTag"]
	if spec == nil {
		return fmt.Errorf("missing CodedTag declaration")
	}
	constraint, ok := ast.Unparen(spec.Type).(*ast.InterfaceType)
	if !ok {
		return s.errorf(spec.Pos(), "CodedTag must be an interface with named type terms")
	}
	s.codedTags = make(map[string]bool)
	var addTerm func(ast.Expr) error
	addTerm = func(expr ast.Expr) error {
		if union, ok := ast.Unparen(expr).(*ast.BinaryExpr); ok && union.Op == token.OR {
			if err := addTerm(union.X); err != nil {
				return err
			}
			return addTerm(union.Y)
		}
		tag, err := s.namedType(expr, make(map[string]bool))
		if err != nil {
			return s.errorf(expr.Pos(), "invalid CodedTag term: %v", err)
		}
		s.codedTags[tag.Name.Name] = true
		return nil
	}
	for _, field := range constraint.Methods.List {
		if len(field.Names) == 0 {
			if err := addTerm(field.Type); err != nil {
				return err
			}
		}
	}
	return nil
}

// Annotations are standalone // @key=value comments. Both leading and trailing
// field comments are accepted, but duplicate or empty annotations are errors.
func (s *schema) annotation(key string, groups ...*ast.CommentGroup) (string, bool, error) {
	var value string
	found := false
	for _, group := range groups {
		if group == nil {
			continue
		}
		for _, comment := range group.List {
			text, ok := strings.CutPrefix(comment.Text, "//")
			if !ok {
				continue
			}
			text = strings.TrimSpace(text)
			if text == key || strings.HasPrefix(text, key+" ") {
				return "", false, s.errorf(comment.Pos(), "malformed %s annotation; want %s=value", key, key)
			}
			v, ok := strings.CutPrefix(text, key+"=")
			if !ok {
				continue
			}
			if found {
				return "", false, s.errorf(comment.Pos(), "duplicate %s annotation", key)
			}
			value = strings.TrimSpace(v)
			if value == "" {
				return "", false, s.errorf(comment.Pos(), "empty %s annotation", key)
			}
			found = true
		}
	}
	return value, found, nil
}

func tableName(s string) string {
	first, size := utf8.DecodeRuneInString(s)
	return "table" + string(unicode.ToUpper(first)) + s[size:]
}
