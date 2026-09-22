// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd_test

import (
	"bytes"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"testing/fstest"
)

func TestGeneratePristine(t *testing.T) {
	t.Parallel()
	// Copy Go sources, including generators and checked-in output, but not the
	// metadata fixtures. Run generation outside the checkout so failures cannot
	// modify source files or interfere with other tests.
	sources := fstest.MapFS{}
	err := fs.WalkDir(os.DirFS("."), ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if entry.Name() == "testdata" {
				return fs.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		data, err := os.ReadFile(filepath.FromSlash(path))
		if err != nil {
			return err
		}
		sources["winmd/"+path] = &fstest.MapFile{Data: data}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	module, err := os.ReadFile("../go.mod")
	if err != nil {
		t.Fatal(err)
	}
	sources["go.mod"] = &fstest.MapFile{Data: module}
	dir := t.TempDir()
	if err := os.CopyFS(dir, sources); err != nil {
		t.Fatal(err)
	}
	cmd := exec.CommandContext(t.Context(), filepath.Join(runtime.GOROOT(), "bin", "go"), "generate", "./...")
	cmd.Dir = dir
	cmd.Env = append(cmd.Environ(), "GOWORK=off")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go generate failed: %v\n%s", err, output)
	}
	// Ignore Git's checkout line-ending conversion, but no other differences.
	normalize := func(data []byte) []byte {
		return bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
	}
	generated := os.DirFS(dir)
	err = fs.WalkDir(generated, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		before, ok := sources[path]
		if !ok {
			t.Errorf("go generate added %s; run go generate ./... and check in the output", path)
			return nil
		}
		after, err := fs.ReadFile(generated, path)
		if err != nil {
			return err
		}
		if !bytes.Equal(normalize(before.Data), normalize(after)) {
			t.Errorf("go generate changed %s; run go generate ./...", path)
		}
		delete(sources, path)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for path := range sources {
		t.Errorf("go generate removed %s", path)
	}
}
