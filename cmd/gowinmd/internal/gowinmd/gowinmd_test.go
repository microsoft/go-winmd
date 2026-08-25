// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package gowinmd

import (
	"encoding/binary"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func TestDecodeInt16AttributeField(t *testing.T) {
	value := []byte{1, 0, 1, 0, 0x53, byte(winmd.ElementType_I2), 15}
	value = append(value, "CountParamIndex"...)
	value = binary.LittleEndian.AppendUint16(value, 3)

	got, ok, err := decodeInt16AttributeField(value, "CountParamIndex")
	if err != nil {
		t.Fatal(err)
	}
	if !ok || got != 3 {
		t.Fatalf("decodeInt16AttributeField() = %v, %v; want 3, true", got, ok)
	}
}

func TestMkwinsyscallModuleName(t *testing.T) {
	for _, test := range []struct {
		name string
		want string
	}{
		{"bcrypt.dll", "bcrypt"},
		{"BCRYPT.DLL", "bcrypt"},
		{"bcrypt", "bcrypt"},
	} {
		if got := mkwinsyscallModuleName(test.name); got != test.want {
			t.Errorf("mkwinsyscallModuleName(%q) = %q; want %q", test.name, got, test.want)
		}
	}
}

func TestContext_writeType_cycle(t *testing.T) {
	t.Skip("cycles can't be built with SigType rather than *SigType, and this code only supports SigType")

	p1 := winmd.SigType{Kind: winmd.ElementType_PTR}
	p2 := winmd.SigType{Kind: winmd.ElementType_PTR}
	// These are copies, not pointers...
	var p1a any = p1
	var p2a any = p2
	// ...So this doesn't cause a cycle.
	p2.Value = p1a
	p1.Value = p2a

	// If we can create a cycle (e.g. if we change to *SigType e.g. for performance reasons) then
	// this code would test it. The strings.Contains check should likely be changed to errors.Is if
	// we do that.
	var b strings.Builder
	var c Context
	if err := c.writeType(&b, &p1, ArchAll); err == nil {
		t.Fatalf("expected error due to detected cycle, but no error was returned")
	} else if !strings.Contains(err.Error(), "cycle") {
		t.Fatalf("got an error, but not a cycle detection error: %v", err)
	}
}
