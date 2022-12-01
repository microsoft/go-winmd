// Copyright (c) Microsoft Corporation
// Licensed under the MIT License.

package main

import (
	"go/ast"
	"go/types"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
)

const tablesFile = "tables.go"

type columnType string

const (
	columnTypeIndex        columnType = "columnTypeIndex"
	columnTypeUint         columnType = "columnTypeUint"
	columnTypeString       columnType = "columnTypeString"
	columnTypeGUID         columnType = "columnTypeGUID"
	columnTypeBlob         columnType = "columnTypeBlob"
	columnTypeCodedIndex   columnType = "columnTypeCodedIndex"
	columnTypeSlice        columnType = "columnTypeSlice"
	columnTypeMethodDefSig columnType = "columnTypeMethodDefSig"
)

type tableInfo struct {
	name string
	// same as name but prefixed with "table" and with the first name letter upper-cased
	tableName string
	exported  bool
	code      uint8
	fields    []columnInfo
}

type columnInfo struct {
	name       string
	typeName   string
	size       int
	columnType columnType
	tableName  string
	coded      string
}

// parsePackage analyzes the single package constructed from the current directory.
func parsePackage() *packages.Package {
	cfg := &packages.Config{
		Mode:  packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		Tests: false,
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	return pkgs[0]
}

func findTableFile(pkg *packages.Package) *ast.File {
	for _, s := range pkg.Syntax {
		f := pkg.Fset.File(s.Pos())
		if f != nil && filepath.Base(f.Name()) == tablesFile {
			return s
		}
	}
	log.Fatalf("%q not found", tablesFile)
	return nil
}

func parseTables(pkg *packages.Package, file *ast.File) []tableInfo {
	var tables []tableInfo
	// traverse the file looking for ast.TypeSpec
	ast.Inspect(file, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok || len(decl.Specs) != 1 {
			return true
		}
		spec, ok := decl.Specs[0].(*ast.TypeSpec)
		if !ok {
			// some specs do not define a type
			return false
		}
		info := parseTable(pkg, spec)
		info.code = tableCode(decl)
		info.tableName = tableName(info.name)
		tables = append(tables, info)
		return false
	})

	// verify that there are exactly 38 tables
	const want = 38
	if got := len(tables); got != want {
		log.Fatalf("got %d tables, want %d", got, want)
	}

	// verify there are no duplicated codes
	codes := make(map[uint8]struct{}, len(tables))
	for _, t := range tables {
		if _, ok := codes[t.code]; ok {
			log.Fatalf("code %#x is duplicated", t.code)
		}
		codes[t.code] = struct{}{}
	}
	return tables
}

// parseTable parses a table from the given spec.
func parseTable(pkg *packages.Package, spec *ast.TypeSpec) (info tableInfo) {
	// this code does not do any effort to do check if the type-casted interfaces
	// will panic at runtime or not. If something panic it means that tables.go does
	// not met genlayout rules.
	t := pkg.TypesInfo.TypeOf(spec.Type).(*types.Struct)
	n := t.NumFields()
	info.name = spec.Name.Name
	info.exported = spec.Name.IsExported()
	info.fields = make([]columnInfo, 0, n)
	for i := 0; i < n; i++ {
		f := t.Field(i)
		var col columnInfo
		tp := f.Type()
		col.name = f.Name()
		switch tp := tp.(type) {
		case *types.Named:
			obj := tp.Obj()
			col.typeName = obj.Name()
			switch objName := obj.Name(); objName {
			case "Index":
				col.columnType = columnTypeIndex
				col.tableName = tableName(fieldComment(spec, i, objName, tp.String(), "@ref"))
			case "CodedIndex":
				col.columnType = columnTypeCodedIndex
				col.coded = fieldComment(spec, i, objName, tp.String(), "@code")
			case "Slice":
				col.columnType = columnTypeSlice
				col.tableName = tableName(fieldComment(spec, i, objName, tp.String(), "@ref"))
			case "String":
				col.columnType = columnTypeString
			case "MethodDefSig":
				col.columnType = columnTypeMethodDefSig
			default:
				if obj.Pkg().Name() == "flags" {
					col.columnType = columnTypeUint
					col.typeName = "flags." + col.typeName
					switch tp.Underlying().(*types.Basic).Kind() {
					case types.Uint8:
						col.size = 1
					case types.Uint16:
						col.size = 2
					case types.Uint32:
						col.size = 4
					}
				}
			}
		case *types.Basic:
			col.typeName = tp.Name()
			switch tp.Kind() {
			case types.Uint8:
				col.columnType = columnTypeUint
				col.size = 1
			case types.Uint16:
				col.columnType = columnTypeUint
				col.size = 2
			case types.Uint32:
				col.columnType = columnTypeUint
				col.size = 4
			}
		case *types.Array:
			if tp.Len() == 16 {
				col.columnType = columnTypeGUID
			}
		case *types.Slice:
			col.columnType = columnTypeBlob
		}
		if col.columnType == "" {
			log.Panicf("unsupported type %s", tp.String())
		}
		info.fields = append(info.fields, col)
	}
	return
}

func fieldComment(spec *ast.TypeSpec, i int, t, tp, marker string) (v string) {
	comment := spec.Type.(*ast.StructType).Fields.List[i].Comment
	if comment != nil && len(comment.List) == 1 {
		v = strings.TrimPrefix(comment.List[0].Text, "// "+marker+"=")
	}
	if v == "" {
		log.Panicf("%s %s requires %s comment", t, tp, marker)
	}
	return v
}

func tableCode(decl *ast.GenDecl) uint8 {
	for _, c := range decl.Doc.List {
		const tablePrefix = "// @table="
		if strings.HasPrefix(c.Text, tablePrefix) {
			code, err := strconv.ParseUint(c.Text[len(tablePrefix):], 0, 8)
			if err != nil {
				log.Panic(err)
			}
			return uint8(code)
		}
	}
	log.Panic("table code not found")
	return 0
}

func tableName(s string) string {
	return "table" + strings.ToUpper(s[:1]) + s[1:]
}
