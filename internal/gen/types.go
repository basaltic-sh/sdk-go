package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

type kind int

const (
	kindStruct kind = iota
	kindEnum
	kindAlias
)

// namedType is a Go type the generator will emit for this service.
type namedType struct {
	Name       string
	Doc        []string
	Kind       kind
	Fields     []*field
	Enum       []enumConst
	Underlying string // for kindEnum and kindAlias
	Origin     ref
	// RequestSide marks a type reachable from a request body. Optional
	// scalars on those become pointers, so that sending a zero value stays
	// distinguishable from leaving the field alone.
	RequestSide bool
	// resolved guards against resolving a recursive schema twice.
	resolved bool
}

type field struct {
	Name     string
	JSONName string
	Type     string
	Doc      []string
	Required bool
	// Nilable types carry no omitempty-relevant zero value distinction.
	OmitEmpty bool
}

type enumConst struct {
	Name  string
	Value string
	Doc   []string
}

// builder turns one service's specification into the Go types and operations
// the emitter writes out.
type builder struct {
	ld             *loader
	service        string
	specDir        string
	serverTemplate string

	types   map[string]*namedType
	byRef   map[string]string // ref identity -> Go type name
	origins map[string]ref    // Go type name -> where it came from
	ops     []*operation

	// overrides renames a schema whose Go name would collide with another
	// file's. Keyed by "<file>.yaml#/<Name>".
	overrides map[string]string
}

// nameOverrides resolves the collisions the specifications currently carry.
//
// Two files define a schema called Image and the compute service reaches
// both: components/schemas/images.yaml#/Image is the /v1/images resource,
// while components/schemas/compute.yaml#/Image is the narrower summary
// embedded in an instance. They are genuinely different types, so the
// generator cannot pick one — it renames the embedded one and says so here.
//
// An entry is a bug in the specifications, not a feature of the generator: a
// new collision fails the build rather than landing silently.
var nameOverrides = map[string]string{
	"compute.yaml#/Image":             "InstanceImage",
	"compute.yaml#/ImageListResponse": "InstanceImageListResponse",
}

func newBuilder(ld *loader, service, specDir string) *builder {
	return &builder{
		ld:        ld,
		service:   service,
		specDir:   specDir,
		types:     map[string]*namedType{},
		byRef:     map[string]string{},
		origins:   map[string]ref{},
		overrides: nameOverrides,
	}
}

// refKey identifies a schema for override lookup and collision reporting.
func refKey(r ref) string { return filepath.Base(r.file) + "#/" + r.pointer }

// namedTypeFor registers (or returns) the Go type for a schema that has a
// name of its own in the specification.
func (b *builder) namedTypeFor(schema node, r ref, requestSide bool) (string, error) {
	if existing, ok := b.byRef[refKey(r)]; ok {
		if requestSide {
			b.markRequestSide(existing)
		}
		return existing, nil
	}

	name := exportedName(r.name())
	if override, ok := b.overrides[refKey(r)]; ok {
		name = override
	}
	if prev, taken := b.origins[name]; taken && refKey(prev) != refKey(r) {
		return "", fmt.Errorf(
			"service %s: two schemas both want the Go name %s — %s and %s.\n"+
				"Give one of them an entry in nameOverrides, or de-duplicate them in the specifications",
			b.service, name, refKey(prev), refKey(r))
	}

	nt := &namedType{Name: name, Origin: r, RequestSide: requestSide}
	b.types[name] = nt
	b.byRef[refKey(r)] = name
	b.origins[name] = r

	// Register before resolving, so a schema that refers to itself finds the
	// name already present instead of recursing forever.
	if err := b.fillNamedType(nt, schema, r.file); err != nil {
		return "", err
	}
	return name, nil
}

// markRequestSide propagates request-side-ness to a type and everything it
// contains, since a nested schema first seen on the response side may later
// turn up inside a request body.
func (b *builder) markRequestSide(name string) {
	nt, ok := b.types[name]
	if !ok || nt.RequestSide {
		return
	}
	nt.RequestSide = true
	for _, f := range nt.Fields {
		b.markRequestSide(strings.TrimLeft(f.Type, "*[]"))
	}
}

func (b *builder) fillNamedType(nt *namedType, schema node, base string) error {
	m, ok := mapOf(schema)
	if !ok {
		nt.Kind = kindAlias
		nt.Underlying = "any"
		return nil
	}
	nt.Doc = goDoc(nt.Name, wrapText(str(m, "description"), 70))

	if enums := list(m, "enum"); len(enums) > 0 && str(m, "type") == "string" {
		nt.Kind = kindEnum
		nt.Underlying = "string"
		for _, e := range enums {
			v := fmt.Sprint(e)
			nt.Enum = append(nt.Enum, enumConst{
				Name:  nt.Name + exportedName(v),
				Value: v,
			})
		}
		return nil
	}

	merged, err := b.mergeSchema(m, base)
	if err != nil {
		return err
	}
	props := obj(merged, "properties")
	if len(props) == 0 {
		// Not a struct: a map, an array or a scalar that happens to be named.
		t, err := b.typeForSchema(merged, base, nt.Name, nt.RequestSide)
		if err != nil {
			return err
		}
		nt.Kind = kindAlias
		nt.Underlying = t
		return nil
	}

	nt.Kind = kindStruct
	required := map[string]bool{}
	for _, r := range list(merged, "required") {
		required[fmt.Sprint(r)] = true
	}
	for _, propName := range orderedKeys(props) {
		propSchema := props[propName]
		f, err := b.buildField(nt, propName, propSchema, base, required[propName])
		if err != nil {
			return err
		}
		nt.Fields = append(nt.Fields, f)
	}
	return nil
}

func (b *builder) buildField(parent *namedType, propName string, schema node, base string, required bool) (*field, error) {
	goName := exportedName(propName)
	if goName == parent.Name {
		// A field cannot share its struct's name in a way that reads well;
		// Go allows it, but "Instance.Instance" is a trap. Rare enough to
		// just suffix.
		goName += "Value"
	}
	hint := parent.Name + goName

	sm, _ := mapOf(schema)
	doc := ""
	if sm != nil {
		doc = str(sm, "description")
		// A description attached beside a $ref is dropped by conforming
		// resolvers, so the specifications write it under a one-element
		// allOf. Pick it up from either place.
		if doc == "" {
			if inner := list(sm, "allOf"); len(inner) == 1 {
				if im, ok := mapOf(inner[0]); ok {
					doc = str(im, "description")
				}
			}
		}
	}

	typ, err := b.typeForSchema(schema, base, hint, parent.RequestSide)
	if err != nil {
		return nil, fmt.Errorf("%s.%s: %w", parent.Name, propName, err)
	}

	// Optional scalars in a request body become pointers so that an explicit
	// zero — false, 0, "" — is distinguishable from an unset field. Response
	// types keep plain values: nothing marshals them, and a pointer would
	// only make every read a nil check.
	if !required && parent.RequestSide && !b.isNilable(typ) {
		typ = "*" + typ
	}

	f := &field{
		Name:      goName,
		JSONName:  propName,
		Type:      typ,
		Required:  required,
		OmitEmpty: !required,
	}
	extra := ""
	if sm != nil {
		if enums := list(sm, "enum"); len(enums) > 0 && !strings.Contains(typ, ".") {
			vals := make([]string, 0, len(enums))
			for _, e := range enums {
				vals = append(vals, fmt.Sprintf("%q", fmt.Sprint(e)))
			}
			extra = "One of: " + strings.Join(vals, ", ") + "."
		}
	}
	if required && parent.RequestSide {
		if extra != "" {
			extra += "\n\n"
		}
		extra += "Required."
	}
	f.Doc = goDoc(goName, wrapText(doc, 68), extra)
	return f, nil
}

// isNilable reports whether a Go type already has a nil value, so wrapping it
// in a pointer would buy nothing.
//
// It resolves named aliases: Metadata and Tags are declared types over
// map[string]string, and "*Metadata" would make callers take the address of a
// map for no gain.
func (b *builder) isNilable(t string) bool {
	if strings.HasPrefix(t, "*") || strings.HasPrefix(t, "[]") ||
		strings.HasPrefix(t, "map[") || t == "any" || t == "json.RawMessage" {
		return true
	}
	if nt, ok := b.types[t]; ok && nt.Kind == kindAlias {
		return b.isNilable(nt.Underlying)
	}
	return false
}

// mergeSchema flattens allOf into one schema.
//
// The specifications use a one-element allOf as the way to attach a
// description to a $ref, and a multi-element one to compose. Both flatten the
// same way. Branches that carry only validation — the oneOf pairs that say
// "exactly one of these two fields" — contribute nothing to a Go type and are
// skipped.
func (b *builder) mergeSchema(m map[string]any, base string) (map[string]any, error) {
	allOf := list(m, "allOf")
	if len(allOf) == 0 {
		return m, nil
	}
	out := map[string]any{}
	for k, v := range m {
		if k == "allOf" {
			continue
		}
		out[k] = v
	}
	props := map[string]any{}
	for k, v := range obj(m, "properties") {
		props[k] = v
	}
	var required []any
	required = append(required, list(m, "required")...)

	for _, member := range allOf {
		mm, ok := mapOf(member)
		if !ok {
			continue
		}
		resolved, r, err := b.ld.deref(mm, base)
		if err != nil {
			return nil, err
		}
		rm, ok := mapOf(resolved)
		if !ok {
			continue
		}
		memberBase := base
		if r.file != "" {
			memberBase = r.file
		}
		flat, err := b.mergeSchema(rm, memberBase)
		if err != nil {
			return nil, err
		}
		for k, v := range obj(flat, "properties") {
			if _, exists := props[k]; !exists {
				props[k] = v
			}
		}
		required = append(required, list(flat, "required")...)
		for _, k := range []string{"type", "additionalProperties", "items"} {
			if _, exists := out[k]; !exists {
				if v, ok := flat[k]; ok {
					out[k] = v
				}
			}
		}
	}
	if len(props) > 0 {
		out["properties"] = props
		if _, ok := out["type"]; !ok {
			out["type"] = "object"
		}
	}
	if len(required) > 0 {
		out["required"] = dedupeAny(required)
	}
	return out, nil
}

func dedupeAny(in []any) []any {
	seen := map[string]bool{}
	var out []any
	for _, v := range in {
		s := fmt.Sprint(v)
		if seen[s] {
			continue
		}
		seen[s] = true
		out = append(out, v)
	}
	sort.Slice(out, func(i, j int) bool { return fmt.Sprint(out[i]) < fmt.Sprint(out[j]) })
	return out
}

// typeForSchema returns the Go type expression for a schema, registering any
// named types it needs along the way.
//
// hint names an inline object that has to become a type of its own; it is
// ignored when the schema resolves to something already named.
func (b *builder) typeForSchema(schema node, base, hint string, requestSide bool) (string, error) {
	m, ok := mapOf(schema)
	if !ok {
		return "any", nil
	}

	if refStr, isRef := m["$ref"].(string); isRef {
		resolved, r, err := b.ld.resolve(refStr, base)
		if err != nil {
			return "", err
		}
		name, err := b.namedTypeFor(resolved, r, requestSide)
		if err != nil {
			return "", err
		}
		if b.types[name].Kind == kindStruct {
			// Struct-valued fields are pointers throughout, so that "absent"
			// and "present but empty" stay distinguishable on a response.
			return "*" + name, nil
		}
		return name, nil
	}

	// A one-element allOf wrapping a $ref is the specifications' way of
	// attaching prose to a reference; unwrap to the reference itself.
	if allOf := list(m, "allOf"); len(allOf) == 1 && len(obj(m, "properties")) == 0 {
		if inner, ok := mapOf(allOf[0]); ok {
			if _, isRef := inner["$ref"]; isRef {
				return b.typeForSchema(inner, base, hint, requestSide)
			}
		}
	}
	merged, err := b.mergeSchema(m, base)
	if err != nil {
		return "", err
	}
	m = merged

	switch str(m, "type") {
	case "array":
		items, ok := m["items"]
		if !ok {
			return "[]any", nil
		}
		elem, err := b.typeForSchema(items, base, singularHint(hint), requestSide)
		if err != nil {
			return "", err
		}
		return "[]" + elem, nil

	case "object", "":
		if props := obj(m, "properties"); len(props) > 0 {
			// An inline object with fields becomes a type of its own, named
			// after where it appears.
			return b.inlineStruct(m, base, hint, requestSide)
		}
		if ap, ok := m["additionalProperties"]; ok {
			if apBool, isBool := ap.(bool); isBool {
				if apBool {
					return "map[string]any", nil
				}
				return "map[string]any", nil
			}
			elem, err := b.typeForSchema(ap, base, hint+"Value", requestSide)
			if err != nil {
				return "", err
			}
			return "map[string]" + elem, nil
		}
		if str(m, "type") == "object" {
			// Deliberately free-form — the metrics query responses, which
			// mirror a Prometheus body the platform does not model.
			return "map[string]any", nil
		}
		return "any", nil

	case "string":
		switch str(m, "format") {
		case "date-time":
			return "time.Time", nil
		case "byte":
			return "[]byte", nil
		}
		return "string", nil

	case "integer":
		switch str(m, "format") {
		case "int64":
			return "int64", nil
		case "int32":
			return "int32", nil
		}
		return "int", nil

	case "number":
		return "float64", nil

	case "boolean":
		return "bool", nil
	}
	return "any", nil
}

// inlineStruct registers a type for an object written inline in the spec.
func (b *builder) inlineStruct(m map[string]any, base, hint string, requestSide bool) (string, error) {
	name := hint
	for i := 2; ; i++ {
		if _, taken := b.types[name]; !taken {
			break
		}
		name = fmt.Sprintf("%s%d", hint, i)
	}
	nt := &namedType{Name: name, RequestSide: requestSide, Kind: kindStruct}
	b.types[name] = nt
	b.origins[name] = ref{file: base, pointer: "inline:" + name}
	if err := b.fillNamedType(nt, m, base); err != nil {
		return "", err
	}
	return "*" + name, nil
}

// singularHint trims a trailing plural from a type-name hint so a slice
// element type reads as one thing rather than many. It appends "Item" when it
// cannot, so the hint is always distinct from its container's name.
func singularHint(s string) string {
	switch {
	case strings.HasSuffix(s, "ies"):
		return strings.TrimSuffix(s, "ies") + "y"
	case strings.HasSuffix(s, "ses"), strings.HasSuffix(s, "xes"):
		return strings.TrimSuffix(s, "es")
	case strings.HasSuffix(s, "s") && !strings.HasSuffix(s, "ss"):
		return strings.TrimSuffix(s, "s")
	}
	return s + "Item"
}
