package gowinmd

import (
	"github.com/microsoft/go-winmd/winmd"
)

type typeNameKey struct {
	NamespaceStart uint32
	NameStart      uint32
}

func typeRefKey(ref winmd.TypeRef) typeNameKey {
	return typeNameKey{ref.Namespace.Start, ref.Name.Start}
}

func typeDefKey(def winmd.TypeDef) typeNameKey {
	return typeNameKey{def.Namespace.Start, def.Name.Start}
}

type typeDefCache struct {
	// aliases maps duplicate heap-offset pairs to the first pair with the same
	// namespace and name. Identity mappings are omitted, so deduplicated heaps
	// need no extra map lookup. Definition aliases are indexed up front;
	// reference aliases are added only when an offset lookup misses.
	aliases map[typeNameKey]typeNameKey
	// resolved maps type name keys -> resolved TypeDef information.
	resolved map[typeNameKey]*resolvedDef
	// unresolved maps type name keys -> winmd.TypeDef index.
	unresolved map[typeNameKey]winmd.Index

	// Duplicated TypeDefs are an uncommon special case, treat them separately to avoid
	// allocating a slice for every TypeDef, only for those that are duplicated.
	// WinMD TypeDefs use duplicated TypeDef names to represent functions with the same name but
	// different signatures due to architecture-specific overloads.

	// resolvedDuplicated maps type name keys -> list of TypeDef indices with the same name.
	resolvedDuplicated map[typeNameKey][]*resolvedDef
	// unresolvedDuplicated maps type name keys -> list of winmd.TypeDef indices.
	unresolvedDuplicated map[typeNameKey][]winmd.Index
}

func newTypeDefCache(capacity uint32) *typeDefCache {
	return &typeDefCache{
		resolved:             make(map[typeNameKey]*resolvedDef),
		unresolved:           make(map[typeNameKey]winmd.Index, capacity),
		unresolvedDuplicated: make(map[typeNameKey][]winmd.Index),
		resolvedDuplicated:   make(map[typeNameKey][]*resolvedDef),
	}
}

// indexTypeNames canonicalizes definitions during the existing top-level index
// pass. References are canonicalized on demand; neither TypeRef nor the raw
// #Strings heap needs an eager scan.
func (c *Context) indexTypeNames() error {
	// Own one string per namespace instead of copying it for every type.
	namespaces := make(map[string]string)
	for idx := range c.Metadata.Tables.TypeDef.Indices() {
		def, err := c.Metadata.Tables.TypeDef.At(idx)
		if err != nil {
			return err
		}
		// Nested types cannot be resolved at module scope.
		if def.Flags.Visibility().IsNested() {
			continue
		}
		namespace, ok := namespaces[def.Namespace.String()]
		if !ok {
			namespace = def.Namespace.String()
			namespaces[namespace] = namespace
		}
		name := qualifiedTypeName{Namespace: namespace, Name: def.Name.String()}
		if firstIndex, ok := c.typeDefsByName[name]; ok {
			first, err := c.Metadata.Tables.TypeDef.At(firstIndex)
			if err != nil {
				return err
			}
			c.typeDefCache.addAlias(typeDefKey(def), typeDefKey(first))
			if c.typeDefNameDuplicates == nil {
				c.typeDefNameDuplicates = make(map[winmd.Index][]winmd.Index)
			}
			if indices := c.typeDefNameDuplicates[firstIndex]; len(indices) != 0 {
				c.typeDefNameDuplicates[firstIndex] = append(indices, idx)
			} else {
				c.typeDefNameDuplicates[firstIndex] = []winmd.Index{firstIndex, idx}
			}
		} else {
			c.typeDefsByName[name] = idx
		}
		c.typeDefCache.add(idx, def)
	}
	return nil
}

func (tc *typeDefCache) addAlias(key, canonical typeNameKey) {
	if key == canonical {
		return
	}
	if tc.aliases == nil {
		tc.aliases = make(map[typeNameKey]typeNameKey)
	}
	tc.aliases[key] = canonical
}

func (tc *typeDefCache) canonicalKey(key typeNameKey) typeNameKey {
	if len(tc.aliases) != 0 {
		if canonical, ok := tc.aliases[key]; ok {
			return canonical
		}
	}
	return key
}

func (tc *typeDefCache) add(i winmd.Index, typ winmd.TypeDef) {
	key := tc.canonicalKey(typeDefKey(typ))
	if _, ok := tc.unresolvedDuplicated[key]; ok {
		tc.unresolvedDuplicated[key] = append(tc.unresolvedDuplicated[key], i)
	} else if _, ok := tc.unresolved[key]; ok {
		tc.unresolvedDuplicated[key] = append(tc.unresolvedDuplicated[key], tc.unresolved[key], i)
		delete(tc.unresolved, key)
	} else {
		tc.unresolved[key] = i
	}
}

func (tc *typeDefCache) resolve(r *resolvedDef) {
	key := tc.canonicalKey(typeNameKey{r.Namespace.Start, r.Name.Start})
	if tc.unresolvedDuplicated[key] == nil {
		tc.resolved[key] = r
		delete(tc.unresolved, key)
	} else {
		tc.resolvedDuplicated[key] = append(tc.resolvedDuplicated[key], r)
		delete(tc.unresolved, key)
	}
}

func (tc *typeDefCache) get(namespace, name winmd.String, arch Arch) *resolvedDef {
	key := tc.canonicalKey(typeNameKey{namespace.Start, name.Start})
	// ArchAll must visit the complete duplicate list in metadata order,
	// not return whichever variant happened to be resolved first.
	if arch != ArchAll {
		for _, def := range tc.resolvedDuplicated[key] {
			if def.Arch&arch == arch {
				return def
			}
		}
	}
	if def, ok := tc.resolved[key]; ok && (arch == ArchAll || def.Arch&arch == arch) {
		return def
	}
	return nil
}

func (tc *typeDefCache) collect(filter func(r *resolvedDef) bool) []*resolvedDef {
	ret := make([]*resolvedDef, 0, len(tc.resolved)+len(tc.resolvedDuplicated))
	for _, r := range tc.resolved {
		if filter(r) {
			ret = append(ret, r)
		}
	}
	for _, def := range tc.resolvedDuplicated {
		for _, r := range def {
			if filter(r) {
				ret = append(ret, r)
			}
		}
	}
	return ret
}
