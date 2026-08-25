# getwinmd

`getwinmd` downloads `Windows.Win32.winmd` from the
[`Microsoft.Windows.SDK.Win32Metadata`](https://www.nuget.org/packages/Microsoft.Windows.SDK.Win32Metadata)
NuGet package.

## Usage

```text
getwinmd [-version <NuGet version>] [-output <path>]
```

- `-version` selects an exact NuGet package version. When omitted, the latest version in NuGet's
  package feed is used, including prerelease versions.
- `-output` selects the destination file. The default is `Windows.Win32.winmd` in the current
  directory.

Download the latest metadata:

```sh
go run github.com/microsoft/go-winmd/cmd/getwinmd@latest
```

Download a specific version:

```sh
go run github.com/microsoft/go-winmd/cmd/getwinmd@latest \
  -version 71.0.20-preview \
  -output Windows.Win32.winmd
```

The command downloads the NuGet archive, locates `Windows.Win32.winmd` inside it, and extracts only
that file. It does not validate the package against a separately maintained checksum; selecting an
exact package version provides reproducible metadata content.
