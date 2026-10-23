// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package gowinmd

import (
	"encoding/binary"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func bufferMethodMetadata(t *testing.T, countType byte) *generationMetadataBuilder {
	t.Helper()
	b := methodGenerationMetadata(t, []byte{0, 2, 0x11, 5, 0x0f, 5, countType}, "native.dll", "Invoke", "buffer", "count")
	b.add(1, uint16(4), b.str("NTSTATUS"), b.str("Windows.Win32.Foundation"))
	b.sizeAttribute(1, true, 1)
	return b
}

func TestWriteMethodBufferChecks(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name      string
		countType byte
		limit     string
	}{
		{"int8", 4, "1<<7-1"},
		{"uint8", 5, "1<<8-1"},
		{"int16", 6, "1<<15-1"},
		{"uint16", 7, "1<<16-1"},
		{"int32", 8, "1<<31-1"},
		{"uint32", 9, "1<<32-1"},
	} {
		t.Run(test.name, func(t *testing.T) {
			b := bufferMethodMetadata(t, test.countType)
			got, err := writeTestMethod(t, b, MethodOptions{GoName: "Read", Projection: ProjectionIdiomatic})
			if err != nil {
				t.Fatal(err)
			}
			for _, want := range []string{
				"//sys\t_winmd_Read(buffer []byte) (ntstatus error) = native.Invoke",
				"func Read(buffer []byte) (ntstatus error)",
				"if uint64(len(buffer)) > " + test.limit,
				"ntstatus = syscall.EINVAL",
				"return _winmd_Read(buffer)",
			} {
				if !strings.Contains(got, want) {
					t.Errorf("missing %q:\n%s", want, got)
				}
			}
			raw, err := writeTestMethod(t, b, MethodOptions{GoName: "Read"})
			if err != nil || strings.Contains(raw, "func ") || strings.Contains(raw, "[]") {
				t.Fatalf("raw projection changed: %q, %v", raw, err)
			}
		})
	}
}

func TestWriteMethodWideBufferCounts(t *testing.T) {
	t.Parallel()
	for _, countType := range []byte{0x0a, 0x0b, 0x18, 0x19} {
		for _, arch := range []Arch{ArchAll, Arch386, ArchAMD64, ArchARM64} {
			b := bufferMethodMetadata(t, countType)
			m := b.metadata()
			c, err := NewContext(m)
			if err != nil {
				t.Fatal(err)
			}
			method, err := m.Tables.MethodDef.At(0)
			if err != nil {
				t.Fatal(err)
			}
			var output strings.Builder
			err = c.WriteMethodWithOptions(&output, 0, method, arch, MethodOptions{Projection: ProjectionIdiomatic})
			got := output.String()
			wantSlice := countType >= 0x18 || arch == ArchAMD64 || arch == ArchARM64
			if err != nil || strings.Contains(got, "buffer []byte") != wantSlice || strings.Contains(got, "func ") {
				t.Fatalf("count %#x, arch %s: %q, %v; want slice: %v, no guard", countType, arch, got, err, wantSlice)
			}
		}
	}
}

func TestWriteMethodExplicitBufferCounts(t *testing.T) {
	t.Parallel()
	for _, countType := range [][]byte{{0x0f, 9}, {0x10, 9}, {2}, {0x0c}} {
		sig := append([]byte{0, 2, 1, 0x0f, 5}, countType...)
		b := methodGenerationMetadata(t, sig, "native.dll", "Invoke", "buffer", "count")
		b.sizeAttribute(1, true, 1)
		got, err := writeTestMethod(t, b, MethodOptions{Projection: ProjectionIdiomatic})
		if err != nil || !strings.HasPrefix(got, "//sys\tInvoke(buffer *uint8, count ") || strings.Contains(got, "func ") {
			t.Fatalf("count %x: %q, %v; want explicit ABI parameters", countType, got, err)
		}
	}
}

func TestWriteMethodBufferCheckExecution(t *testing.T) {
	var generated strings.Builder
	for _, test := range []struct {
		name    string
		retType byte
		lastErr bool
	}{
		{"Checked", 0, false},
		{"LastError", 8, true},
		{"NoError", 8, false},
		{"NoResult", 1, false},
	} {
		b := bufferMethodMetadata(t, 5)
		if test.retType != 0 {
			b = methodGenerationMetadata(t, []byte{0, 2, test.retType, 0x0f, 5, 5}, "native.dll", "Invoke", "buffer", "count")
			b.sizeAttribute(1, true, 1)
		}
		if test.lastErr {
			binary.LittleEndian.PutUint16(b.rows[28][0], uint16(winmd.PInvokeCallingConvention_PlatformAPI)|uint16(winmd.PInvokeFlags_SupportsLastError))
		}
		got, err := writeTestMethod(t, b, MethodOptions{GoName: test.name, Projection: ProjectionIdiomatic})
		if err != nil {
			t.Fatal(err)
		}
		generated.WriteString(got)
		generated.WriteString("\n\n")
	}
	b := methodGenerationMetadata(t, []byte{0, 5, 0x11, 5, 0x0f, 5, 5, 0x0f, 7, 4, 9}, "native.dll", "Invoke", "buffer", "byteCount", "words", "wordCount", "flags")
	b.add(1, uint16(4), b.str("NTSTATUS"), b.str("Windows.Win32.Foundation"))
	b.sizeAttribute(1, true, 1)
	b.sizeAttribute(3, false, 3)
	got, err := writeTestMethod(t, b, MethodOptions{GoName: "Many", Projection: ProjectionIdiomatic})
	if err != nil {
		t.Fatal(err)
	}
	generated.WriteString(got)
	generated.WriteString("\n\n")
	const tests = `
var calls int
var nativeError error
func _winmd_Checked(buffer []byte) error { calls++; return nativeError }
func _winmd_LastError(buffer []byte) (int32, error) { calls++; return 42, nativeError }
func _winmd_NoError(buffer []byte) int32 { calls++; return 42 }
func _winmd_NoResult(buffer []byte) { calls++ }
func _winmd_Many(buffer []byte, words []uint16, flags uint32) error {
	calls++
	if len(buffer) != 255 || len(words) != 127 || flags != 17 { return syscall.EIO }
	return nil
}

func TestWrappers(t *testing.T) {
	for _, size := range []int{0, 1, 254, 255} {
		buffer := make([]byte, size)
		calls = 0
		if err := Checked(buffer); err != nil { t.Fatal(err) }
		if result, err := LastError(buffer); result != 42 || err != nil { t.Fatalf("%v, %v", result, err) }
		if result := NoError(buffer); result != 42 { t.Fatal(result) }
		NoResult(buffer)
		if calls != 4 { t.Fatalf("native calls = %d, want 4", calls) }
	}
	buffer := make([]byte, 256)
	calls = 0
	if err := Checked(buffer); err != syscall.EINVAL { t.Fatalf("got %v, want EINVAL", err) }
	if result, err := LastError(buffer); result != 0 || err != syscall.EINVAL { t.Fatalf("%v, %v", result, err) }
	for _, call := range []func(){func() { NoError(buffer) }, func() { NoResult(buffer) }} {
		func() {
			defer func() { if recover() == nil { t.Error("missing overflow panic") } }()
			call()
		}()
	}
	if calls != 0 { t.Fatalf("overflow reached native code %d times", calls) }
	nativeError = syscall.EIO
	if err := Checked(nil); err != nativeError { t.Fatal(err) }
	if result, err := LastError(nil); result != 42 || err != nativeError { t.Fatalf("%v, %v", result, err) }
	nativeError = nil
	words := make([]uint16, 128)
	if err := Many(buffer[:255], words[:127], 17); err != nil { t.Fatal(err) }
	calls = 0
	if err := Many(buffer, words[:127], 17); err != syscall.EINVAL { t.Fatal(err) }
	if err := Many(buffer[:255], words, 17); err != syscall.EINVAL { t.Fatal(err) }
	if calls != 0 { t.Fatalf("multiple-buffer overflow reached native code %d times", calls) }
	for _, input := range [][]byte{nil, buffer[:255], buffer} {
		if allocs := testing.AllocsPerRun(100, func() { _ = Checked(input) }); allocs != 0 {
			t.Fatalf("len %d: %v allocations, want 0", len(input), allocs)
		}
	}
}
`
	dir := t.TempDir()
	source := "package generated\nimport (\"syscall\"; \"testing\")\n" + generated.String() + tests
	if err := os.WriteFile(filepath.Join(dir, "generated_test.go"), []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module generated\n\ngo 1.26\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", "-count=1", ".")
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOWORK=off")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated wrapper tests: %v\n%s\n%s", err, output, source)
	}
}
