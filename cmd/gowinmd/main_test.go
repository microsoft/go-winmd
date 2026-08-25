// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	gotypes "go/types"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/cmd/gowinmd/internal/gowinmd"
	"github.com/microsoft/go-winmd/winmd"
)

func newArchBuilders() map[gowinmd.Arch]*strings.Builder {
	return map[gowinmd.Arch]*strings.Builder{
		gowinmd.Arch386:   {},
		gowinmd.ArchAMD64: {},
		gowinmd.ArchARM64: {},
		gowinmd.ArchAll:   {},
		gowinmd.ArchNone:  {},
	}
}

func TestParseInputFilesTypeDirectives(t *testing.T) {
	path := filepath.Join(t.TempDir(), "types.go")
	source := `package cryptography

//winmd:type Windows.Win32.Security.Cryptography.BCRYPT_RSAKEY_BLOB
//winmd:type Windows.Win32.Security.Cryptography.BCRYPT_RSAKEY_BLOB
//winmd:type Windows.Win32.Security.Cryptography.BCRYPT_RSAKEY_BLOB -name RSAKEY_BLOB
//winmd:func bcrypt.dll.BCryptEncrypt
`
	if err := os.WriteFile(path, []byte(source), 0o666); err != nil {
		t.Fatal(err)
	}

	methods, types, pkg, err := parseInputFiles([]string{path})
	if err != nil {
		t.Fatal(err)
	}
	if pkg != "cryptography" {
		t.Fatalf("package = %q; want cryptography", pkg)
	}
	if len(methods) != 1 || methods["bcrypt.dll.bcryptencrypt"] != "" {
		t.Fatalf("methods = %#v; want BCryptEncrypt", methods)
	}
	if len(types) != 1 {
		t.Fatalf("types = %#v; want one deduplicated type", types)
	}
	want := typeSelection{
		Namespace: "Windows.Win32.Security.Cryptography",
		Name:      "BCRYPT_RSAKEY_BLOB",
		GoName:    "RSAKEY_BLOB",
	}
	if got := types["Windows.Win32.Security.Cryptography.BCRYPT_RSAKEY_BLOB"]; got != want {
		t.Fatalf("type selection = %#v; want %#v", got, want)
	}
}

func TestParseInputFilesOnlyTypeDirective(t *testing.T) {
	path := filepath.Join(t.TempDir(), "types.go")
	if err := os.WriteFile(path, []byte("package p\n//winmd:type Windows.Win32.Foundation.POINT\n"), 0o666); err != nil {
		t.Fatal(err)
	}
	methods, types, _, err := parseInputFiles([]string{path})
	if err != nil {
		t.Fatal(err)
	}
	if len(methods) != 0 || len(types) != 1 {
		t.Fatalf("got %d methods and %d types; want 0 and 1", len(methods), len(types))
	}
}

func TestParseInputFilesRejectsInvalidTypeDirectives(t *testing.T) {
	tests := []struct {
		name      string
		directive string
		wantError string
	}{
		{"unqualified", "//winmd:type BCRYPT_RSAKEY_BLOB", "not fully qualified"},
		{"missing type", "//winmd:type Windows.Win32.", "not fully qualified"},
		{"missing name value", "//winmd:type Windows.Win32.POINT -name", "expected <namespace>.<type>"},
		{"unexpected option", "//winmd:type Windows.Win32.POINT -rename POINT", "unexpected option"},
		{"invalid name", "//winmd:type Windows.Win32.POINT -name 3D_POINT", "invalid Go type name"},
		{"keyword name", "//winmd:type Windows.Win32.POINT -name type", "invalid Go type name"},
		{"blank name", "//winmd:type Windows.Win32.POINT -name _", "invalid Go type name"},
		{"missing whitespace", "//winmd:typeWindows.Win32.POINT", "malformed //winmd:type directive"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "types.go")
			source := "package p\n" + test.directive + "\n"
			if err := os.WriteFile(path, []byte(source), 0o666); err != nil {
				t.Fatal(err)
			}
			_, _, _, err := parseInputFiles([]string{path})
			if err == nil || !strings.Contains(err.Error(), test.wantError) {
				t.Fatalf("parseInputFiles() error = %v; want error containing %q", err, test.wantError)
			}
		})
	}
}

func TestParseInputFilesRejectsConflictingTypeNames(t *testing.T) {
	path := filepath.Join(t.TempDir(), "types.go")
	source := `package p
//winmd:type Windows.Win32.Foundation.POINT -name FirstPoint
//winmd:type Windows.Win32.Foundation.POINT -name SecondPoint
`
	if err := os.WriteFile(path, []byte(source), 0o666); err != nil {
		t.Fatal(err)
	}
	_, _, _, err := parseInputFiles([]string{path})
	if err == nil || !strings.Contains(err.Error(), "conflicting Go names") {
		t.Fatalf("parseInputFiles() error = %v; want conflicting Go names error", err)
	}
}

func TestWriteMethod(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
	filter, err := buildNamespaceFilter(f,
		"Windows.Win32.Storage.FileSystem",
		"Windows.Win32.Security.Cryptography",
		"Windows.Win32.System.Diagnostics.Debug",
	)
	if err != nil {
		t.Fatal(err)
	}
	if err := writePrototypes(b, f, filter); err != nil {
		t.Fatal(err)
	}
	for arch, w := range b {
		if w.Len() == 0 {
			continue
		}
		formattedContent, err := format.Source([]byte(w.String()))
		if err != nil {
			t.Fatal(err)
		}
		target := "prototypes.golden"
		if arch != gowinmd.ArchAll {
			target += "_" + arch.String()
		}
		target += ".go"
		Check(t, "go test ./cmd/gowinmd", filepath.Join("testdata", target), string(formattedContent))
	}
}

func TestWriteProjectedBCryptMethods(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
	filter := methodFilter{}
	for _, name := range []string{
		"BCryptGetFipsAlgorithmMode", "BCryptGetProperty", "BCryptSetProperty",
		"BCryptOpenAlgorithmProvider", "BCryptCreateHash", "BCryptHashData",
		"BCryptEncrypt", "BCryptDecrypt", "BCryptGenerateSymmetricKey",
		"BCryptSignHash", "BCryptDeriveKey",
	} {
		filter[strings.ToLower("bcrypt.dll."+name)] = ""
	}
	filter["bcrypt.dll.bcrypthashdata"] = "rawEncrypt"

	if err := writePrototypesWithProjection(b, f, filter, gowinmd.ProjectionIdiomatic); err != nil {
		t.Fatal(err)
	}
	got := b[gowinmd.ArchAll].String()
	wants := []string{
		"//sys\tBCryptGetProperty(hObject BCRYPT_HANDLE, pszProperty *uint16, pbOutput []byte, pcbResult *uint32, dwFlags uint32) (ntstatus error) = bcrypt.BCryptGetProperty",
		"//sys\tBCryptSetProperty(hObject BCRYPT_HANDLE, pszProperty *uint16, pbInput []byte, dwFlags uint32) (ntstatus error) = bcrypt.BCryptSetProperty",
		"//sys\trawEncrypt(hHash BCRYPT_HASH_HANDLE, pbInput []byte, dwFlags uint32) (ntstatus error) = bcrypt.BCryptHashData",
	}
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("projected output does not contain %q\noutput:\n%s", want, got)
		}
	}
	if strings.Contains(got, ".dll.") || strings.Contains(got, "PWSTRElement") {
		t.Errorf("projected output contains an unnormalized DLL or synthetic string element type:\n%s", got)
	}
	if !strings.Contains(got, "BCRYPT_BLOCK_PADDING BCRYPT_FLAGS = 0x1") ||
		!strings.Contains(got, "BCRYPT_PAD_NONE BCRYPT_FLAGS = 0x1") {
		t.Errorf("projected output did not preserve duplicate constant values:\n%s", got)
	}
}

func TestWriteStandaloneType(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
	types := typeFilter{
		"Windows.Win32.Security.Cryptography.BCRYPT_RSAKEY_BLOB": {
			Namespace: "Windows.Win32.Security.Cryptography",
			Name:      "BCRYPT_RSAKEY_BLOB",
			GoName:    "RSAKEY_BLOB",
		},
	}
	if err := writeSelectionsWithProjection(b, f, methodFilter{}, types, gowinmd.ProjectionRaw); err != nil {
		t.Fatal(err)
	}
	got := b[gowinmd.ArchAll].String()
	if !strings.Contains(got, "type RSAKEY_BLOB struct {") {
		t.Fatalf("standalone type was not emitted with its custom name:\n%s", got)
	}
	for arch, output := range b {
		if arch != gowinmd.ArchAll && strings.Contains(output.String(), "type RSAKEY_BLOB struct {") {
			t.Fatalf("standalone type was also emitted for %s:\n%s", arch, output.String())
		}
	}
}

func TestWriteStandaloneTypeArchitectureLayout(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
	types := typeFilter{
		"Windows.Win32.Security.Cryptography.BCRYPT_AUTHENTICATED_CIPHER_MODE_INFO": {
			Namespace: "Windows.Win32.Security.Cryptography",
			Name:      "BCRYPT_AUTHENTICATED_CIPHER_MODE_INFO",
			GoName:    "AUTHENTICATED_CIPHER_MODE_INFO",
		},
	}
	if err := writeSelectionsWithProjection(b, f, methodFilter{}, types, gowinmd.ProjectionRaw); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(b[gowinmd.ArchAll].String(), "type AUTHENTICATED_CIPHER_MODE_INFO struct {") {
		t.Fatalf("architecture-dependent type was also emitted in common output:\n%s", b[gowinmd.ArchAll].String())
	}
	want386 := "CbAAD uint32\n\t_ [4]byte\n\tCbData uint64\n\tDwFlags uint32\n\t_ [4]byte"
	if got := b[gowinmd.Arch386].String(); !strings.Contains(got, want386) {
		t.Fatalf("386 layout does not contain required uint64 and tail padding:\n%s", got)
	}
	for _, arch := range []gowinmd.Arch{gowinmd.ArchAMD64, gowinmd.ArchARM64} {
		got := b[arch].String()
		if !strings.Contains(got, "type AUTHENTICATED_CIPHER_MODE_INFO struct {") {
			t.Fatalf("%s layout was not emitted:\n%s", arch, got)
		}
		if strings.Contains(got, "_ [4]byte") {
			t.Fatalf("%s layout contains explicit 386 padding:\n%s", arch, got)
		}
	}
	layoutTests := []struct {
		arch              gowinmd.Arch
		wantSize          int64
		wantCbDataOffset  int64
		wantDwFlagsOffset int64
	}{
		{gowinmd.Arch386, 64, 48, 56},
		{gowinmd.ArchAMD64, 88, 72, 80},
		{gowinmd.ArchARM64, 88, 72, 80},
	}
	for _, test := range layoutTests {
		t.Run("GoLayout/"+test.arch.String(), func(t *testing.T) {
			source := "package p\n" + b[gowinmd.ArchAll].String() + b[test.arch].String()
			fileSet := token.NewFileSet()
			file, err := parser.ParseFile(fileSet, "generated.go", source, 0)
			if err != nil {
				t.Fatal(err)
			}
			sizes := gotypes.SizesFor("gc", test.arch.String())
			if sizes == nil {
				t.Fatalf("no Go sizes available for %s", test.arch)
			}
			checked, err := (&gotypes.Config{Sizes: sizes}).Check("p", fileSet, []*ast.File{file}, nil)
			if err != nil {
				t.Fatal(err)
			}
			object := checked.Scope().Lookup("AUTHENTICATED_CIPHER_MODE_INFO")
			if object == nil {
				t.Fatal("generated type not found")
			}
			structure := object.Type().Underlying().(*gotypes.Struct)
			fields := make([]*gotypes.Var, structure.NumFields())
			fieldIndices := make(map[string]int)
			for i := range fields {
				fields[i] = structure.Field(i)
				fieldIndices[fields[i].Name()] = i
			}
			offsets := sizes.Offsetsof(fields)
			if got := sizes.Sizeof(object.Type()); got != test.wantSize {
				t.Errorf("Go size = %d; want %d", got, test.wantSize)
			}
			if got := offsets[fieldIndices["CbData"]]; got != test.wantCbDataOffset {
				t.Errorf("CbData offset = %d; want %d", got, test.wantCbDataOffset)
			}
			if got := offsets[fieldIndices["DwFlags"]]; got != test.wantDwFlagsOffset {
				t.Errorf("DwFlags offset = %d; want %d", got, test.wantDwFlagsOffset)
			}
		})
	}
}

func TestWriteStandaloneTypeRejectsUnknownType(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	types := typeFilter{
		"Windows.Win32.Security.Cryptography.NOT_A_REAL_TYPE": {
			Namespace: "Windows.Win32.Security.Cryptography",
			Name:      "NOT_A_REAL_TYPE",
		},
	}
	err = writeSelectionsWithProjection(newArchBuilders(), f, methodFilter{}, types, gowinmd.ProjectionRaw)
	if err == nil || !strings.Contains(err.Error(), "unknown WinMD type") {
		t.Fatalf("writeSelectionsWithProjection() error = %v; want unknown WinMD type", err)
	}
}

func TestWriteStandaloneTypeMatchesCaseSensitively(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	types := typeFilter{
		"Windows.Win32.Security.Cryptography.bcrypt_rsakey_blob": {
			Namespace: "Windows.Win32.Security.Cryptography",
			Name:      "bcrypt_rsakey_blob",
		},
	}
	err = writeSelectionsWithProjection(newArchBuilders(), f, methodFilter{}, types, gowinmd.ProjectionRaw)
	if err == nil || !strings.Contains(err.Error(), "unknown WinMD type") {
		t.Fatalf("writeSelectionsWithProjection() error = %v; want case-sensitive unknown WinMD type", err)
	}
}

func TestWriteSelectedTypeUsesCustomNameInMethodsAndDependencies(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name       string
		methods    methodFilter
		types      typeFilter
		want       []string
		projection gowinmd.Projection
	}{
		{
			name:    "raw method",
			methods: methodFilter{"kernel32.dll.comparefiletime": ""},
			types: typeFilter{
				"Windows.Win32.Foundation.FILETIME": {
					Namespace: "Windows.Win32.Foundation",
					Name:      "FILETIME",
					GoName:    "WinFileTime",
				},
			},
			want: []string{
				"//sys\tCompareFileTime(lpFileTime1 *WinFileTime, lpFileTime2 *WinFileTime)",
				"type WinFileTime struct {",
			},
			projection: gowinmd.ProjectionRaw,
		},
		{
			name:    "idiomatic method",
			methods: methodFilter{"kernel32.dll.comparefiletime": ""},
			types: typeFilter{
				"Windows.Win32.Foundation.FILETIME": {
					Namespace: "Windows.Win32.Foundation",
					Name:      "FILETIME",
					GoName:    "WinFileTime",
				},
			},
			want: []string{
				"//sys\tCompareFileTime(lpFileTime1 *WinFileTime, lpFileTime2 *WinFileTime)",
				"type WinFileTime struct {",
			},
			projection: gowinmd.ProjectionIdiomatic,
		},
		{
			name: "struct dependency",
			types: typeFilter{
				"Windows.Win32.Security.Cryptography.BCRYPT_DSA_PARAMETER_HEADER_V2": {
					Namespace: "Windows.Win32.Security.Cryptography",
					Name:      "BCRYPT_DSA_PARAMETER_HEADER_V2",
				},
				"Windows.Win32.Security.Cryptography.HASHALGORITHM_ENUM": {
					Namespace: "Windows.Win32.Security.Cryptography",
					Name:      "HASHALGORITHM_ENUM",
					GoName:    "HashAlgorithm",
				},
			},
			want: []string{
				"HashAlgorithm HashAlgorithm",
				"type HashAlgorithm int32",
			},
			projection: gowinmd.ProjectionRaw,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := newArchBuilders()
			if err := writeSelectionsWithProjection(b, f, test.methods, test.types, test.projection); err != nil {
				t.Fatal(err)
			}
			got := b[gowinmd.ArchAll].String()
			for _, want := range test.want {
				if !strings.Contains(got, want) {
					t.Fatalf("output does not contain %q:\n%s", want, got)
				}
			}
		})
	}
}

func TestWriteBCryptTypesGolden(t *testing.T) {
	methods, types, pkg, err := parseInputFiles([]string{filepath.Join("testdata", "bcrypt_types.go")})
	if err != nil {
		t.Fatal(err)
	}
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}

	generate := func() map[gowinmd.Arch]string {
		b := newArchBuilders()
		if err := writeSelectionsWithProjection(b, f, methods, types, gowinmd.ProjectionRaw); err != nil {
			t.Fatal(err)
		}
		result := make(map[gowinmd.Arch]string)
		for arch, builder := range b {
			if builder.Len() == 0 {
				continue
			}
			formatted, err := format.Source([]byte(generateFileContent(builder.String(), pkg)))
			if err != nil {
				t.Fatal(err)
			}
			result[arch] = string(formatted)
		}
		return result
	}

	first := generate()
	second := generate()
	for arch, actual := range first {
		if second[arch] != actual {
			t.Fatalf("generation for %s is not deterministic", arch)
		}
		target := "bcrypt_types.golden"
		if arch != gowinmd.ArchAll {
			target += "_" + arch.String()
		}
		target += ".go"
		Check(t, "go test ./cmd/gowinmd", filepath.Join("testdata", target), actual)
	}
	if len(first) != 4 {
		t.Fatalf("generated %d files; want common, 386, amd64, and arm64 files", len(first))
	}
	commonTypeNames := make(map[string]bool)
	for _, line := range strings.Split(first[gowinmd.ArchAll], "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[0] == "type" {
			commonTypeNames[fields[1]] = true
		}
	}
	for _, arch := range []gowinmd.Arch{gowinmd.Arch386, gowinmd.ArchAMD64, gowinmd.ArchARM64} {
		for _, line := range strings.Split(first[arch], "\n") {
			fields := strings.Fields(line)
			if len(fields) >= 2 && fields[0] == "type" && commonTypeNames[fields[1]] {
				t.Fatalf("type %s was emitted in both common and %s output", fields[1], arch)
			}
		}
	}
	for _, typeName := range []string{"PKCS1_PADDING_INFO", "PSS_PADDING_INFO", "OAEP_PADDING_INFO", "PQDSA_PADDING_INFO", "AUTHENTICATED_CIPHER_MODE_INFO"} {
		if commonTypeNames[typeName] {
			t.Errorf("architecture-dependent type %s was emitted in common output", typeName)
		}
		for _, arch := range []gowinmd.Arch{gowinmd.Arch386, gowinmd.ArchAMD64, gowinmd.ArchARM64} {
			if !strings.Contains(first[arch], "type "+typeName+" ") {
				t.Errorf("architecture-dependent type %s was not emitted for %s", typeName, arch)
			}
		}
	}
}

func TestIdiomaticProjectionDoesNotProjectHRESULT(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
	filter := methodFilter{"ncrypt.dll.ncryptgetproperty": ""}
	if err := writePrototypesWithProjection(b, f, filter, gowinmd.ProjectionIdiomatic); err != nil {
		t.Fatal(err)
	}
	got := b[gowinmd.ArchAll].String()
	if !strings.Contains(got, "(r HRESULT) = ncrypt.NCryptGetProperty") {
		t.Fatalf("HRESULT return was unexpectedly projected:\n%s", got)
	}
}

func TestFullFile(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := newArchBuilders()
	if err := writePrototypes(b, f, nil); err != nil {
		t.Fatal(err)
	}
	for _, w := range b {
		_, err = format.Source([]byte(w.String()))
		if err != nil {
			t.Fatal(err)
		}
	}

	// The generated source code is ~4 MB, so don't write it to source control as a golden file.
	// This test only checks that the generation process doesn't fail and doesn't take an
	// exceptionally long time.

	// To see the output, use:
	//   go run ./cmd/gowinmd -o all.go.temp -source .\winmd\testdata\Windows.Win32.winmd
}

func openTestWinmd() (*winmd.Metadata, error) {
	return winmd.Open("../../winmd/testdata/Windows.Win32.winmd")
}

// buildNamespaceFilter pre-scans the winmd file to build a methodFilter containing
// all module.method entries for the given namespaces.
func buildNamespaceFilter(f *winmd.Metadata, namespaces ...string) (methodFilter, error) {
	ctx, err := gowinmd.NewContext(f)
	if err != nil {
		return nil, err
	}
	nsSet := make(map[string]bool, len(namespaces))
	for _, ns := range namespaces {
		nsSet[ns] = true
	}
	filter := make(methodFilter)
	for idx := range f.Tables.TypeDef.Indices() {
		r, err := f.Tables.TypeDef.At(idx)
		if err != nil {
			return nil, err
		}
		if !nsSet[r.Namespace.String()] {
			continue
		}
		for j := range r.MethodList.All() {
			md, err := f.Tables.MethodDef.At(j)
			if err != nil {
				return nil, err
			}
			moduleName := ctx.MethodModuleName(j)
			if moduleName == "" {
				continue
			}
			filter[strings.ToLower(moduleName+"."+md.Name.String())] = ""
		}
	}
	return filter, nil
}
