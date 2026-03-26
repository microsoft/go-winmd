# gowinmd

`gowinmd` generates a file that can be fed to `x/sys/windows/mkwinsyscall` to generate syscalls for methods in a win32metadata (winmd) file.

The inputs, outputs, and in general the configurability of this tool is a work in progress.
See [go-winmd#8](https://github.com/microsoft/go-winmd/issues/8)

## Usage

```
gowinmd -source <path/to/Windows.Win32.winmd> -format mkwinsyscall [-output <output.go>] <input.go ...>
```
### Flags

- `-source` — Path to the win32metadata (winmd) file to parse (required).
- `-format` — Output format (required). Supported values: `mkwinsyscall`.
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
  
        `//winmd:func kernel32.CreateFileW -name CreateFile`

## Output

`gowinmd` generates a complete Go source file with auto-detected imports and type aliases based on the generated content. The package name is inferred from the input Go files.