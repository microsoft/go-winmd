// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"archive/zip"
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunDownloadsExplicitVersion(t *testing.T) {
	packageData := makePackage(t, map[string][]byte{
		"metadata/Windows.Win32.winmd": []byte("explicit metadata"),
	})
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/package/1.2.3-preview" {
			t.Errorf("request path = %q; want explicit package path", request.URL.Path)
			http.NotFound(response, request)
			return
		}
		response.Write(packageData)
	}))
	defer server.Close()

	output := filepath.Join(t.TempDir(), "metadata", metadataFileName)
	if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(output, []byte("old metadata"), 0o644); err != nil {
		t.Fatal(err)
	}
	var stdout, stderr strings.Builder
	err := run(context.Background(), []string{"-version", "1.2.3-Preview", "-output", output}, &stdout, &stderr, server.Client(), nugetEndpoints{
		versionFeedURL: server.URL + "/atom.xml",
		packageBaseURL: server.URL + "/package",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q; want empty", stderr.String())
	}
	if !strings.Contains(stdout.String(), "1.2.3-Preview") {
		t.Fatalf("stdout = %q; want resolved version", stdout.String())
	}
	got, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "explicit metadata" {
		t.Fatalf("output = %q; want explicit metadata", got)
	}
}

func TestRunDownloadsLatestVersion(t *testing.T) {
	packageData := makePackage(t, map[string][]byte{
		metadataFileName: []byte("latest metadata"),
	})
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/atom.xml":
			io.WriteString(response, `<feed xmlns="http://www.w3.org/2005/Atom"><entry><id>https://www.nuget.org/packages/`+packageID+`/2.0.0-preview</id></entry><entry><id>https://www.nuget.org/packages/`+packageID+`/1.0.0</id></entry></feed>`)
		case "/package/2.0.0-preview":
			response.Write(packageData)
		default:
			http.NotFound(response, request)
		}
	}))
	defer server.Close()

	output := filepath.Join(t.TempDir(), metadataFileName)
	var stdout, stderr strings.Builder
	if err := run(context.Background(), []string{"-output", output}, &stdout, &stderr, server.Client(), nugetEndpoints{
		versionFeedURL: server.URL + "/atom.xml",
		packageBaseURL: server.URL + "/package",
	}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stdout.String(), "2.0.0-preview") {
		t.Fatalf("stdout = %q; want latest version", stdout.String())
	}
	got, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "latest metadata" {
		t.Fatalf("output = %q; want latest metadata", got)
	}
}

func TestLatestVersionRejectsEmptyIndex(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		io.WriteString(response, `<feed xmlns="http://www.w3.org/2005/Atom"></feed>`)
	}))
	defer server.Close()

	_, err := latestVersion(context.Background(), server.Client(), server.URL)
	if err == nil || !strings.Contains(err.Error(), "no versions") {
		t.Fatalf("latestVersion() error = %v; want no versions error", err)
	}
}

func TestRunReportsHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		http.Error(response, "not available", http.StatusNotFound)
	}))
	defer server.Close()

	err := run(context.Background(), []string{"-version", "1.0.0"}, io.Discard, io.Discard, server.Client(), nugetEndpoints{
		versionFeedURL: server.URL + "/atom.xml",
		packageBaseURL: server.URL + "/package",
	})
	if err == nil || !strings.Contains(err.Error(), "404 Not Found") || !strings.Contains(err.Error(), "not available") {
		t.Fatalf("run() error = %v; want HTTP status and response", err)
	}
}

func TestExtractMetadataRejectsMissingFile(t *testing.T) {
	packageData := makePackage(t, map[string][]byte{"README.md": []byte("no metadata")})
	err := extractMetadata(packageData, filepath.Join(t.TempDir(), metadataFileName))
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Fatalf("extractMetadata() error = %v; want not found error", err)
	}
}

func TestFindMetadataPrefersArchiveRoot(t *testing.T) {
	packageData := makePackage(t, map[string][]byte{
		metadataFileName:                []byte("root"),
		"duplicate/Windows.Win32.winmd": []byte("duplicate"),
	})
	archive, err := zip.NewReader(bytes.NewReader(packageData), int64(len(packageData)))
	if err != nil {
		t.Fatal(err)
	}
	file, err := findMetadata(archive.File)
	if err != nil {
		t.Fatal(err)
	}
	if file.Name != metadataFileName {
		t.Fatalf("findMetadata() = %q; want root file", file.Name)
	}
}

func TestRunRejectsUnexpectedArguments(t *testing.T) {
	err := run(context.Background(), []string{"unexpected"}, io.Discard, io.Discard, http.DefaultClient, nugetEndpoints{
		versionFeedURL: "https://example.invalid/atom.xml",
		packageBaseURL: "https://example.invalid/package",
	})
	if err == nil || !strings.Contains(err.Error(), "unexpected arguments") {
		t.Fatalf("run() error = %v; want unexpected arguments error", err)
	}
}

func makePackage(t *testing.T, files map[string][]byte) []byte {
	t.Helper()
	var buffer bytes.Buffer
	archive := zip.NewWriter(&buffer)
	for name, contents := range files {
		file, err := archive.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := file.Write(contents); err != nil {
			t.Fatal(err)
		}
	}
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	return buffer.Bytes()
}
