package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// The platform's specifications are modular: each service's openapi.yaml
// re-exports path items from ../paths and schemas from ../components via
// $ref. Nothing here bundles them to disk — the generator resolves refs as it
// walks, so the specification repository stays the single source of truth and
// the SDK carries no copy of it.

// node is a decoded YAML value: map[string]any, []any, or a scalar.
type node = any

// loader reads and caches specification files, and resolves $ref across them.
type loader struct {
	root  string
	files map[string]map[string]any
}

func newLoader(root string) *loader {
	return &loader{root: root, files: map[string]map[string]any{}}
}

func (l *loader) load(path string) (map[string]any, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	if doc, ok := l.files[abs]; ok {
		return doc, nil
	}
	data, err := os.ReadFile(abs)
	if err != nil {
		return nil, err
	}
	var doc map[string]any
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("%s: %w", abs, err)
	}
	l.files[abs] = doc
	return doc, nil
}

// ref is a resolved reference: the file it landed in and the JSON pointer
// within it. Together they identify a schema uniquely, which is what lets the
// generator notice that two files define the same name.
type ref struct {
	file    string // absolute path
	pointer string // "Instance", or "" for a whole document
}

func (r ref) name() string {
	i := strings.LastIndexByte(r.pointer, '/')
	return r.pointer[i+1:]
}

func (r ref) String() string { return filepath.Base(r.file) + "#/" + r.pointer }

// resolve follows a $ref written relative to base, returning the node it
// names along with its identity.
func (l *loader) resolve(refStr, base string) (node, ref, error) {
	file, pointer, _ := strings.Cut(refStr, "#")
	target := base
	if file != "" {
		target = filepath.Join(filepath.Dir(base), file)
	}
	abs, err := filepath.Abs(target)
	if err != nil {
		return nil, ref{}, err
	}
	doc, err := l.load(abs)
	if err != nil {
		return nil, ref{}, fmt.Errorf("resolving %q from %s: %w", refStr, filepath.Base(base), err)
	}
	var cur node = doc
	pointer = strings.TrimPrefix(pointer, "/")
	if pointer != "" {
		for _, tok := range strings.Split(pointer, "/") {
			tok = strings.ReplaceAll(tok, "~1", "/")
			tok = strings.ReplaceAll(tok, "~0", "~")
			m, ok := cur.(map[string]any)
			if !ok {
				return nil, ref{}, fmt.Errorf("resolving %q: %q is not an object", refStr, tok)
			}
			cur, ok = m[tok]
			if !ok {
				return nil, ref{}, fmt.Errorf("resolving %q: no such key %q in %s", refStr, tok, filepath.Base(abs))
			}
		}
	}
	return cur, ref{file: abs, pointer: pointer}, nil
}

// deref follows n if it is a $ref, otherwise returns it unchanged. The
// returned ref identifies where the value came from, so a caller can resolve
// further refs relative to the right file.
func (l *loader) deref(n node, base string) (node, ref, error) {
	m, ok := n.(map[string]any)
	if !ok {
		return n, ref{file: base}, nil
	}
	r, ok := m["$ref"].(string)
	if !ok {
		return n, ref{file: base}, nil
	}
	return l.resolve(r, base)
}

// mapOf narrows a node to an object, or reports that it is not one.
func mapOf(n node) (map[string]any, bool) {
	m, ok := n.(map[string]any)
	return m, ok
}

func str(m map[string]any, key string) string {
	s, _ := m[key].(string)
	return s
}

func boolOf(m map[string]any, key string) bool {
	b, _ := m[key].(bool)
	return b
}

func list(m map[string]any, key string) []any {
	l, _ := m[key].([]any)
	return l
}

func obj(m map[string]any, key string) map[string]any {
	o, _ := m[key].(map[string]any)
	return o
}

// orderedKeys returns a map's keys sorted, so generated output does not
// change from run to run.
func orderedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sortStrings(out)
	return out
}
