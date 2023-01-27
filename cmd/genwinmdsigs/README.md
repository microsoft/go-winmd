# genwinmdsigs

`genwinmdsigs` generates a file that can be fed to `x/sys/windows/mkwinsyscall` to generate syscalls for methods in a win32metadata (winmd) file.

The inputs, outputs, and in general the configurability of this tool is a work in progress.
See [go-winmd#8](https://github.com/microsoft/go-winmd/issues/8)

## Template

`genwinmdsigs` accepts a template file (`-template path/to/template.tmpl`) to specify a header and footer for the generated file.
An example is provided at [ExampleTemplate.tmpl](ExampleTemplate.tmpl). 