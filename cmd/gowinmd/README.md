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

Input Go files are scanned for `//winmd:func` directives that specify which APIs to generate. The package name for the generated output is inferred from the input files.

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

## Output

`gowinmd` generates a complete Go source file with auto-detected imports and type aliases based on the generated content. The package name is inferred from the input Go files.

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