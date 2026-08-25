// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Command getwinmd downloads Windows.Win32.winmd from NuGet.
package main

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const (
	packageID             = "Microsoft.Windows.SDK.Win32Metadata"
	metadataFileName      = "Windows.Win32.winmd"
	defaultVersionFeedURL = "https://www.nuget.org/packages/" + packageID + "/atom.xml"
	defaultPackageBaseURL = "https://www.nuget.org/api/v2/package/" + packageID
	maxPackageSize        = int64(128 << 20)
	maxMetadataSize       = int64(256 << 20)
)

func main() {
	client := &http.Client{Timeout: 5 * time.Minute}
	if err := run(context.Background(), os.Args[1:], os.Stdout, os.Stderr, client, nugetEndpoints{
		versionFeedURL: defaultVersionFeedURL,
		packageBaseURL: defaultPackageBaseURL,
	}); err != nil {
		log.Fatal(err)
	}
}

type nugetEndpoints struct {
	versionFeedURL string
	packageBaseURL string
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer, client *http.Client, endpoints nugetEndpoints) error {
	flags := flag.NewFlagSet("getwinmd", flag.ContinueOnError)
	flags.SetOutput(stderr)
	version := flags.String("version", "", "NuGet package version (latest if omitted)")
	output := flags.String("output", metadataFileName, "output WinMD file")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("unexpected arguments: %s", strings.Join(flags.Args(), " "))
	}
	if strings.TrimSpace(*output) == "" {
		return errors.New("output path must not be empty")
	}

	resolvedVersion := strings.TrimSpace(*version)
	if resolvedVersion == "" {
		var err error
		resolvedVersion, err = latestVersion(ctx, client, endpoints.versionFeedURL)
		if err != nil {
			return err
		}
	}
	if strings.ContainsAny(resolvedVersion, `/\\`) {
		return fmt.Errorf("invalid package version %q", resolvedVersion)
	}

	packageData, err := downloadPackage(ctx, client, endpoints.packageBaseURL, resolvedVersion)
	if err != nil {
		return err
	}
	if err := extractMetadata(packageData, *output); err != nil {
		return err
	}
	fmt.Fprintf(stdout, "Downloaded %s %s to %s\n", packageID, resolvedVersion, *output)
	return nil
}

type atomFeed struct {
	Entries []atomEntry `xml:"entry"`
}

type atomEntry struct {
	ID string `xml:"id"`
}

func latestVersion(ctx context.Context, client *http.Client, versionFeedURL string) (string, error) {
	response, err := get(ctx, client, versionFeedURL)
	if err != nil {
		return "", fmt.Errorf("resolve latest %s version: %w", packageID, err)
	}
	defer response.Body.Close()

	var feed atomFeed
	if err := xml.NewDecoder(response.Body).Decode(&feed); err != nil {
		return "", fmt.Errorf("decode %s version feed: %w", packageID, err)
	}
	// NuGet's package Atom feed is ordered from newest to oldest.
	for _, entry := range feed.Entries {
		entryURL, err := url.Parse(strings.TrimSpace(entry.ID))
		if err != nil {
			continue
		}
		version, err := url.PathUnescape(path.Base(strings.TrimSuffix(entryURL.Path, "/")))
		if err == nil && version != "" && version != "." {
			return version, nil
		}
	}
	return "", fmt.Errorf("NuGet returned no versions for %s", packageID)
}

func downloadPackage(ctx context.Context, client *http.Client, packageBaseURL, version string) ([]byte, error) {
	normalizedVersion := strings.ToLower(version)
	escapedVersion := url.PathEscape(normalizedVersion)
	packageURL := strings.TrimRight(packageBaseURL, "/") + "/" + escapedVersion
	response, err := get(ctx, client, packageURL)
	if err != nil {
		return nil, fmt.Errorf("download %s %s: %w", packageID, version, err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(io.LimitReader(response.Body, maxPackageSize+1))
	if err != nil {
		return nil, fmt.Errorf("read %s %s: %w", packageID, version, err)
	}
	if int64(len(data)) > maxPackageSize {
		return nil, fmt.Errorf("%s %s exceeds the %d-byte download limit", packageID, version, maxPackageSize)
	}
	return data, nil
}

func get(ctx context.Context, client *http.Client, address string) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, address, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "go-winmd/getwinmd")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		message, _ := io.ReadAll(io.LimitReader(response.Body, 4<<10))
		return nil, fmt.Errorf("GET %s: %s%s", address, response.Status, formatResponseMessage(message))
	}
	return response, nil
}

func formatResponseMessage(message []byte) string {
	trimmed := strings.TrimSpace(string(message))
	if trimmed == "" {
		return ""
	}
	return ": " + trimmed
}

func extractMetadata(packageData []byte, output string) error {
	archive, err := zip.NewReader(bytes.NewReader(packageData), int64(len(packageData)))
	if err != nil {
		return fmt.Errorf("open NuGet package: %w", err)
	}
	metadata, err := findMetadata(archive.File)
	if err != nil {
		return err
	}
	if metadata.UncompressedSize64 > uint64(maxMetadataSize) {
		return fmt.Errorf("%s exceeds the %d-byte extraction limit", metadata.Name, maxMetadataSize)
	}

	reader, err := metadata.Open()
	if err != nil {
		return fmt.Errorf("open %s in NuGet package: %w", metadata.Name, err)
	}
	defer reader.Close()

	directory := filepath.Dir(output)
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	temporary, err := os.CreateTemp(directory, "."+filepath.Base(output)+"-*")
	if err != nil {
		return fmt.Errorf("create temporary output: %w", err)
	}
	temporaryName := temporary.Name()
	defer os.Remove(temporaryName)

	written, copyErr := io.Copy(temporary, io.LimitReader(reader, maxMetadataSize+1))
	if copyErr != nil {
		temporary.Close()
		return fmt.Errorf("extract %s: %w", metadata.Name, copyErr)
	}
	if written > maxMetadataSize {
		temporary.Close()
		return fmt.Errorf("%s exceeds the %d-byte extraction limit", metadata.Name, maxMetadataSize)
	}
	if err := temporary.Chmod(0o644); err != nil {
		temporary.Close()
		return fmt.Errorf("set output permissions: %w", err)
	}
	if err := temporary.Close(); err != nil {
		return fmt.Errorf("close temporary output: %w", err)
	}
	if err := replaceFile(temporaryName, output); err != nil {
		return fmt.Errorf("write %s: %w", output, err)
	}
	return nil
}

func findMetadata(files []*zip.File) (*zip.File, error) {
	var matches []*zip.File
	for _, file := range files {
		if file.FileInfo().IsDir() || !strings.EqualFold(path.Base(file.Name), metadataFileName) {
			continue
		}
		matches = append(matches, file)
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("%s not found in %s package", metadataFileName, packageID)
	}
	if len(matches) == 1 {
		return matches[0], nil
	}
	for _, file := range matches {
		if file.Name == metadataFileName {
			return file, nil
		}
	}
	return nil, fmt.Errorf("multiple %s files found in %s package", metadataFileName, packageID)
}

func replaceFile(source, destination string) error {
	if err := os.Rename(source, destination); err == nil {
		return nil
	}
	if err := os.Remove(destination); err != nil && !os.IsNotExist(err) {
		return err
	}
	return os.Rename(source, destination)
}
