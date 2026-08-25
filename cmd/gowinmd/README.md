# gowinmd

`gowinmd` generates a file that can be fed to `x/sys/windows/mkwinsyscall` to generate syscalls for methods in a win32metadata (winmd) file.

The inputs, outputs, and in general the configurability of this tool is a work in progress.
See [go-winmd#8](https://github.com/microsoft/go-winmd/issues/8)

## Usage

```
gowinmd -source <path/to/Windows.Win32.winmd> -format mkwinsyscall [-projection raw|idiomatic] [-output <output.go>] <input.go ...>
```
### Flags

- `-source` — Path to the win32metadata (winmd) file to parse (required).
- `-format` — Output format (required). Supported values: `mkwinsyscall`.
- `-projection` — Signature projection. `raw` (the default) preserves literal ABI parameters. `idiomatic` uses metadata-backed slices, normal Go string element pointers, and the mkwinsyscall NTSTATUS error convention.
- `-output` — Output file name. Prints to stdout if omitted.

Input Go files are scanned for `//winmd:func` and `//winmd:type` directives. The package name for the generated output is inferred from the input files. Function and type directives may be mixed in one or more files, and an input may contain only type directives.

### `//winmd:func` directive

Each directive specifies an API to import using the format `moduleref.methoddef`:

```go
//winmd:func kernel32.CreateFileW
//winmd:func kernel32.ReadFile
//winmd:func advapi32.RegOpenKeyExW
```

If only the module name is specified, all methods from that module are generated:

```go
//winmd:func kernel32.*
```

This directive supports the following optional flags:

- `-name` — Specify a custom Go function name for the imported method. By default, the Go name is the same as the method name in the winmd file. Example:
  
        //winmd:func kernel32.CreateFileW -name CreateFile

The requested `-name` is emitted exactly, including its capitalization. Both exported and
unexported valid Go identifiers are supported.

### `//winmd:type` directive

Each type directive contains a case-sensitive, fully qualified WinMD namespace and type name:

```go
//winmd:type Windows.Win32.Security.Cryptography.BCRYPT_RSAKEY_BLOB
```

The namespace and type must both be present. Matching is case-sensitive because WinMD identifiers
are case-sensitive. An unknown type or a name that matches multiple definitions for the same target
architecture is reported as an error.

Use `-name` to choose the generated Go type name:

```go
//winmd:type Windows.Win32.Security.Cryptography.BCRYPT_RSAKEY_BLOB -name RSAKEY_BLOB
```

Custom names must be non-blank Go identifiers and cannot be Go keywords. The name is applied to the
TypeDef itself and every generated reference to it, including references from selected functions and
other generated types. Repeating an identical directive has no effect. Multiple non-empty custom
names for the same WinMD type are rejected.

Explicitly selected structs, enums, native typedefs, and their constants are emitted even when no
selected function references them. Field types and other TypeDef dependencies are resolved
recursively through the same dependency traversal used for function signatures. Each resolved
TypeDef is emitted once, and enum members are preserved in metadata order, including duplicate
values.

For example, this input generates a standalone struct and its enum dependencies alongside a syscall:

```go
//winmd:type Windows.Win32.Security.Cryptography.BCRYPT_DSA_PARAMETER_HEADER_V2 -name DSA_PARAMETER_HEADER_V2
//winmd:func bcrypt.dll.BCryptGetProperty
```

## Output

`gowinmd` generates a complete Go source file with auto-detected imports and type aliases based on the generated content. The package name is inferred from the input Go files.

Explicitly selected types are checked against the Windows ABI for 386, amd64, and arm64. A definition
is written to the common output only when its ABI size, alignment, and field offsets are identical on
every target and the same Go declaration represents all three layouts. If any layout property differs,
including pointer-sized fields, the definition is omitted from the common file and written to
architecture-suffixed files such as `output_386.go`, `output_amd64.go`, and `output_arm64.go`. This
also permits the additional 386 padding around the `uint64` field in
`BCRYPT_AUTHENTICATED_CIPHER_MODE_INFO`. Definitions are never emitted in both common and
architecture-specific output.

WinMD layouts that require overlapping non-union fields, reduced alignment that Go cannot express
for typed fields, variable-length inline arrays, or unresolved external by-value field types are
rejected with an error instead of being emitted with an incorrect ABI.

Module references in directives use their WinMD names, which commonly include `.dll`:

```go
//winmd:func bcrypt.dll.BCryptGetProperty
```

The mkwinsyscall target omits that suffix because mkwinsyscall adds it while loading the DLL:

```go
//sys BCryptGetProperty(hObject BCRYPT_HANDLE, pszProperty *uint16, pbOutput []byte, pcbResult *uint32, dwFlags uint32) (ntstatus error) = bcrypt.BCryptGetProperty
```

In `idiomatic` projection, a pointer and its immediately following length are coalesced only when
`NativeArrayInfoAttribute` or `MemorySizeAttribute` explicitly associates them. This preserves the
native argument order while letting mkwinsyscall expand a slice back to pointer and length arguments.
Input, output, optional, and nullable buffers all use this convention. An empty or nil slice passes
a nil pointer and zero length. APIs that require a non-nil pointer for zero-length input still need a
handwritten wrapper and can use `-projection raw` as the escape hatch.

Only `NTSTATUS` returns become `(ntstatus error)`. mkwinsyscall preserves a nonzero status as
`windows.NTStatus` and returns nil on success; other numeric return typedefs remain unchanged.