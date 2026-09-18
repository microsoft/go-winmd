// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package gowinmd

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd/winmd"
)

func (b *generationMetadataBuilder) typeArchitecture(row uint16, arch Arch) {
	refRow := uint16(len(b.rows[1]) + 1)
	b.add(1, uint16(4), b.str("SupportedArchitectureAttribute"), b.str("Windows.Win32.Foundation.Metadata"))
	memberRow := uint16(len(b.rows[10]) + 1)
	b.add(10, refRow<<3|1, b.str(".ctor"), b.blob([]byte{0x20, 1, 1, 8}))
	value := binary.LittleEndian.AppendUint32([]byte{1, 0}, uint32(arch))
	value = append(value, 0, 0) // No named arguments.
	b.add(12, row<<5|3, memberRow<<3|3, b.blob(value))
}

func TestContextAttributeClassification(t *testing.T) {
	t.Parallel()
	for _, namespace := range []string{"Windows.Win32.Foundation.Metadata", "Windows.Win32.Interop"} {
		t.Run(namespace, func(t *testing.T) {
			b := newGenerationMetadataBuilder(t)
			for _, name := range []string{"First", "Second"} {
				b.add(2, uint32(0), b.str(name), b.str("Test"), uint16(0), uint16(1), uint16(1))
			}
			for _, ns := range []string{namespace, "Other"} {
				b.add(1, uint16(4), b.str("SupportedArchitectureAttribute"), b.str(ns))
				b.add(10, uint16(len(b.rows[1]))<<3|1, b.str(".ctor"), b.blob([]byte{0x20, 1, 1, 8}))
			}
			// An unused invalid attribute type must not be read eagerly.
			b.add(1, uint16(4), uint16(0xffff), b.str(namespace))
			b.add(10, uint16(3<<3|1), b.str(".ctor"), b.blob([]byte{0x20, 1, 1, 8}))
			invalid := b.blob([]byte{0xff})
			for i, arch := range []Arch{Arch386, ArchARM64} {
				parent := uint16(i+1)<<5 | 3
				// Same attribute name in another namespace: ignore even a bad value.
				b.add(12, parent, uint16(2<<3|3), invalid)
				value := binary.LittleEndian.AppendUint32([]byte{1, 0}, uint32(arch))
				b.add(12, parent, uint16(1<<3|3), b.blob(append(value, 0, 0)))
			}
			c, err := NewContext(b.metadata())
			if err != nil {
				t.Fatal(err)
			}
			if c.TypeDefSupportedArch(0) != Arch386 || c.TypeDefSupportedArch(1) != ArchARM64 {
				t.Fatal("shared constructor lost per-attribute values or namespace identity")
			}
			for _, test := range []struct {
				name, want string
				ctor       uint16
				value      []byte
			}{
				{"duplicate", "multiple SupportedArchitectureAttribute", 1, []byte{1, 0, 1, 0, 0, 0, 0, 0}},
				{"invalid-value", "decode SupportedArchitectureAttribute", 1, []byte{0xff}},
				{"invalid-reference", "TypeRef[2].Name", 3, []byte{0xff}},
			} {
				t.Run(test.name, func(t *testing.T) {
					b.add(12, uint16(1<<5|3), test.ctor<<3|3, b.blob(test.value))
					_, err := NewContext(b.metadata())
					b.rows[12] = b.rows[12][:len(b.rows[12])-1]
					if err == nil || !strings.Contains(err.Error(), test.want) {
						t.Fatalf("NewContext() error = %v; want %q", err, test.want)
					}
				})
			}
		})
	}
}

// architectureTestContext creates equally named definitions with real
// SupportedArchitecture attributes, either at module scope or under one parent.
// The TypeRef's heap offsets may match any definition or none of them.
func architectureTestContext(t *testing.T, arches []Arch, offsets string, nested bool) (*Context, winmd.Index, []winmd.Index) {
	t.Helper()
	b := newGenerationMetadataBuilder(t)
	namespace := b.str("Test")
	name := "Sample"
	if nested {
		name = "Child"
	}
	names := make([]uint16, len(arches))
	names[0] = b.str(name)
	for i := 1; i < len(names); i++ {
		names[i] = names[0]
		if offsets != "shared" {
			names[i] = b.str(name)
		}
	}
	refName, refNamespace := names[0], namespace
	switch offsets {
	case "last":
		refName = names[len(names)-1]
	case "new":
		refName, refNamespace = b.str(name), b.str("Test")
	}
	b.add(35, uint16(4), uint16(0), uint16(0), uint16(0), uint32(0), uint16(0), b.str("System.Runtime"), uint16(0), uint16(0))
	b.add(1, uint16(6), b.str("ValueType"), b.str("System"))
	b.add(2, uint32(0), b.str("<Module>"), uint16(0), uint16(0), uint16(1), uint16(1))
	refIndex := winmd.Index(1)
	const commonFlags = uint32(winmd.TypeLayout_SequentialLayout) | uint32(winmd.TypeFlags_Sealed)
	flags := commonFlags | uint32(winmd.TypeVisibility_Public)
	if nested {
		parentName := b.str("Parent")
		b.add(1, uint16(4), parentName, namespace)
		b.add(1, uint16(11), refName, uint16(0))
		b.add(2, flags, parentName, namespace, uint16(5), uint16(1), uint16(1))
		b.add(15, uint16(0), uint32(1), uint16(2)) // The fieldless parent has a nonzero size.
		refIndex = 2
		flags = commonFlags | uint32(winmd.TypeVisibility_NestedPublic)
		namespace = 0
	} else {
		b.add(1, uint16(4), refName, refNamespace)
	}
	indices := make([]winmd.Index, len(arches))
	for i, arch := range arches {
		row := uint16(len(b.rows[2]) + 1)
		indices[i] = winmd.Index(row - 1)
		b.add(2, flags, names[i], namespace, uint16(5), uint16(i+1), uint16(1))
		kind := winmd.ElementType_U4
		if arch&Arch386 == 0 {
			kind = winmd.ElementType_U8
		}
		b.add(4, uint16(winmd.MemberAccess_Public), b.str("Value"), b.blob([]byte{6, byte(kind)}))
		b.typeArchitecture(row, arch)
		if nested {
			b.add(41, row, uint16(2))
		}
	}
	c, err := NewContext(b.metadata())
	if err != nil {
		t.Fatal(err)
	}
	for i, index := range indices {
		if got := c.TypeDefSupportedArch(index); got != arches[i] {
			t.Fatalf("decoded architecture = %d; want %d", got, arches[i])
		}
	}
	return c, refIndex, indices
}

func TestResolveSingleTypeArchitecture(t *testing.T) {
	t.Parallel()
	for _, supported := range []Arch{ArchAMD64, ArchAMD64 | ArchARM64, ArchAll} {
		for _, offsets := range []string{"shared", "new"} {
			for _, state := range []string{"cold", "definition-cached", "reference-cached"} {
				for _, requested := range []Arch{Arch386, ArchAMD64, ArchARM64, ArchAMD64 | ArchARM64, ArchAll} {
					t.Run(fmt.Sprintf("supported-%d/%s/%s/requested-%d", supported, offsets, state, requested), func(t *testing.T) {
						c, refIndex, indices := architectureTestContext(t, []Arch{supported}, offsets, false)
						switch state {
						case "definition-cached":
							if _, err := c.resolveTypeDef(indices[0]); err != nil {
								t.Fatal(err)
							}
						case "reference-cached":
							if _, err := c.resolveTypeRef(refIndex, supported); err != nil {
								t.Fatal(err)
							}
						}
						got, err := c.resolveTypeRef(refIndex, requested)
						if requested == ArchAll || supported&requested == requested {
							if err != nil || got == nil || got.Index != indices[0] || got.Arch != supported {
								t.Fatalf("supported lookup = %+v, %v; want TypeDef %d with architecture %d", got, err, indices[0], supported)
							}
							return
						}
						if got != nil || !errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
							t.Fatalf("unsupported lookup = %+v, %v; want no definition", got, err)
						}
						if state == "cold" && c.resolvedDefsByIndex[indices[0]] != nil {
							t.Fatal("unsupported lookup resolved the wrong-architecture definition")
						}
						if got, err := c.resolveTypeRef(refIndex, supported); err != nil || got == nil || got.Index != indices[0] {
							t.Fatalf("unsupported lookup poisoned a later supported lookup: %+v, %v", got, err)
						}
					})
				}
			}
		}
	}
}

func TestTypeDefCacheSingleArchitecture(t *testing.T) {
	t.Parallel()
	for _, offsets := range []string{"shared", "new"} {
		t.Run(offsets, func(t *testing.T) {
			c, refIndex, _ := architectureTestContext(t, []Arch{ArchAMD64}, offsets, false)
			want, err := c.resolveTypeRef(refIndex, ArchAMD64)
			if err != nil {
				t.Fatal(err)
			}
			ref, err := c.Metadata.Tables.TypeRef.At(refIndex)
			if err != nil {
				t.Fatal(err)
			}
			for _, arch := range []Arch{Arch386, ArchAMD64, ArchARM64, ArchAll} {
				got := c.typeDefCache.get(ref.Namespace, ref.Name, arch)
				if arch == ArchAMD64 || arch == ArchAll {
					if got != want {
						t.Fatalf("cache lookup on %s = %+v; want the cached definition", arch, got)
					}
				} else if got != nil {
					t.Fatalf("cache returned an AMD64-only definition for %s", arch)
				}
			}
		})
	}
}

func TestSelectTypeDefArchitectureVariants(t *testing.T) {
	t.Parallel()
	arches := []Arch{Arch386, ArchAMD64, ArchARM64}
	for _, offsets := range []string{"shared", "first", "last", "new"} {
		for _, warm := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/warm-%v", offsets, warm), func(t *testing.T) {
				c, refIndex, _ := architectureTestContext(t, arches, offsets, false)
				if warm {
					if _, err := c.resolveTypeRef(refIndex, ArchARM64); err != nil {
						t.Fatal(err)
					}
				}
				if err := c.SelectTypeDef("Test", "Sample", "Selected"); err != nil {
					t.Fatal(err)
				}
				out := map[Arch]*strings.Builder{Arch386: {}, ArchAMD64: {}, ArchARM64: {}, ArchAll: {}}
				if err := c.WriteUsedTypeDefs(out); err != nil {
					t.Fatal(err)
				}
				for _, arch := range arches {
					field := "Value uint64"
					if arch == Arch386 {
						field = "Value uint32"
					}
					text := out[arch].String()
					if strings.Count(text, "type Selected struct {") != 1 || !strings.Contains(text, field) {
						t.Fatalf("selected %s variant missing or duplicated:\n%s", arch, text)
					}
				}
				if out[ArchAll].Len() != 0 {
					t.Fatal("architecture-specific variants emitted as a common definition")
				}
			})
		}
	}
}

func TestResolveAllTypeArchitectures(t *testing.T) {
	t.Parallel()
	arches := []Arch{Arch386, ArchAMD64, ArchARM64}
	for _, offsets := range []string{"shared", "first", "last", "new"} {
		for _, warm := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/warm-%v", offsets, warm), func(t *testing.T) {
				c, refIndex, indices := architectureTestContext(t, arches, offsets, false)
				if warm {
					got, err := c.resolveTypeRef(refIndex, ArchARM64)
					if err != nil || got == nil || got.Index != indices[2] {
						t.Fatalf("warm-up = %+v, %v; want the last variant", got, err)
					}
				}
				for range 2 {
					got, err := c.resolveTypeRef(refIndex, ArchAll)
					if err != nil || got == nil || got.Index != indices[0] {
						t.Fatalf("ArchAll lookup = %+v, %v; want the first metadata variant", got, err)
					}
					for i, index := range indices {
						if def := c.resolvedDefsByIndex[index]; def == nil || def.Arch != arches[i] {
							t.Fatalf("ArchAll did not resolve variant %d with its declared architecture", index)
						}
					}
					// The warmed path must not require another textual lookup.
					c.typeDefsByName = nil
				}
				for i, arch := range arches {
					got, err := c.resolveTypeRef(refIndex, arch)
					if err != nil || got == nil || got.Index != indices[i] {
						t.Fatalf("lookup after ArchAll on %s = %+v, %v", arch, got, err)
					}
				}
			})
		}
	}
}

func TestResolveNestedTypeArchitecture(t *testing.T) {
	t.Parallel()
	for _, offsets := range []string{"shared", "first", "last", "new"} {
		for _, warm := range []bool{false, true} {
			for _, arch := range []Arch{Arch386, ArchAMD64, ArchARM64, ArchAll} {
				t.Run(fmt.Sprintf("%s/warm-%v/%s", offsets, warm, arch), func(t *testing.T) {
					c, refIndex, indices := architectureTestContext(t, []Arch{Arch386, ArchAMD64}, offsets, true)
					if warm {
						if _, err := c.resolveTypeRef(refIndex, ArchAMD64); err != nil {
							t.Fatal(err)
						}
					}
					def, err := c.resolveTypeRef(refIndex, arch)
					if arch == ArchARM64 {
						if def != nil || !errors.Is(err, errTypeDefNotDefinedInCurrentModule) {
							t.Fatalf("unsupported nested lookup returned index instead of an error: %v", err)
						}
						return
					}
					wantIndex, wantSize := indices[0], uint32(4)
					if arch == ArchAMD64 {
						wantIndex, wantSize = indices[1], 8
					}
					if err != nil || def == nil {
						t.Fatalf("nested lookup failed: %v", err)
					}
					if def.Index != wantIndex || def.Parent == nil || def.Parent.Index != 1 {
						t.Fatalf("nested lookup returned TypeDef %d; want child %d of Parent", def.Index, wantIndex)
					}
					layout, err := c.resolvedDefABITypeLayout(def, arch, nil)
					if err != nil || layout.abiSize != wantSize {
						t.Fatalf("nested ABI layout = %+v, %v; want size %d", layout, err, wantSize)
					}
				})
			}
		}
	}
}
