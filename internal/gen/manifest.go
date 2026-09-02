package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// The manifest describes the SDK's own Go surface as data.
//
// It exists so that tooling which generates code CALLING this SDK — the CLI —
// does not have to re-derive the Go names from the OpenAPI specifications. Two
// independent implementations of the naming rules would drift, and the drift
// would land as a wall of compile errors in the downstream tool rather than
// anywhere near its cause.
//
// It also means the CLI's generator needs only this file, not a checkout of
// the specification repository, which is private.
//
// Committed alongside the generated packages and regenerated with them.

type manifest struct {
	Module   string            `json:"module"`
	Version  string            `json:"version"`
	Services []manifestService `json:"services"`
}

type manifestService struct {
	Name             string       `json:"name"`
	Package          string       `json:"package"`
	Title            string       `json:"title"`
	Description      string       `json:"description,omitempty"`
	EndpointTemplate string       `json:"endpoint_template"`
	Regional         bool         `json:"regional"`
	Operations       []manifestOp `json:"operations"`
}

type manifestOp struct {
	ID      string `json:"id"`
	GoName  string `json:"go_name"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	Summary string `json:"summary,omitempty"`
	// Resource is the noun this operation acts on, taken from x-resource
	// where the specification sets it and derived from the path where it does
	// not. Only two thirds of operations carry the extension, so the path is
	// the authority and the extension is a cross-check.
	Resource string `json:"resource"`
	// Verb is the action, reduced to a small canonical set where the
	// operation is ordinary CRUD and left as the domain's own word where it
	// is not: start, failover, rotate.
	Verb string `json:"verb"`

	PathParams []manifestParam `json:"path_params,omitempty"`

	ParamsType string          `json:"params_type,omitempty"`
	Params     []manifestField `json:"params,omitempty"`

	// BodyType is the Go type exactly as the method takes it, pointer and
	// all: whether a body is passed by value or by pointer depends on the
	// schema and cannot be guessed from the name.
	BodyType   string          `json:"body_type,omitempty"`
	BodyKind   string          `json:"body_kind,omitempty"` // json | stream | text
	BodyFields []manifestField `json:"body_fields,omitempty"`

	ResultKind string `json:"result_kind"` // none | value | page | stream
	ResultType string `json:"result_type,omitempty"`
	ItemType   string `json:"item_type,omitempty"`

	Paginated  bool `json:"paginated,omitempty"`
	Idempotent bool `json:"idempotent,omitempty"`
	Deprecated bool `json:"deprecated,omitempty"`
}

type manifestParam struct {
	Wire   string `json:"wire"`
	GoName string `json:"go_name"`
	Doc    string `json:"doc,omitempty"`
}

// manifestField is one settable input: a query parameter or a request-body
// field. FlagKind says how a command-line tool should surface it.
type manifestField struct {
	Wire     string   `json:"wire"`
	GoName   string   `json:"go_name"`
	GoType   string   `json:"go_type"`
	FlagKind string   `json:"flag_kind"` // string|bool|int|int64|float|duration|stringSlice|json
	Pointer  bool     `json:"pointer,omitempty"`
	Required bool     `json:"required,omitempty"`
	Enum     []string `json:"enum,omitempty"`
	Doc      string   `json:"doc,omitempty"`
}

// flagKind classifies a Go type for a command-line flag.
//
// Anything that is not a scalar or a list of scalars becomes a JSON flag: a
// nested object has no honest flat representation, and inventing dotted flag
// names for it produces a worse interface than asking for the object.
func flagKind(goType string) string {
	t := strings.TrimPrefix(goType, "*")
	switch t {
	case "string":
		return "string"
	case "bool":
		return "bool"
	case "int", "int32":
		return "int"
	case "int64":
		return "int64"
	case "float64":
		return "float"
	case "time.Time":
		// Not a string: a caller has to parse it, so say so rather than
		// letting a downstream tool bind a *time.Time to a string flag.
		return "time"
	case "[]string":
		return "stringSlice"
	case "[]byte":
		return "bytes"
	}
	if strings.HasPrefix(t, "[]") || strings.HasPrefix(t, "map[") {
		return "json"
	}
	return ""
}

// buildManifest projects one service's builder state into manifest form.
func (b *builder) buildManifest(info map[string]any) manifestService {
	svc := manifestService{
		Name:             b.service,
		Package:          b.service,
		Title:            str(info, "title"),
		Description:      strings.TrimSpace(str(info, "description")),
		EndpointTemplate: b.serverTemplate,
		Regional:         strings.Contains(b.serverTemplate, "{region}"),
	}
	for _, op := range b.ops {
		svc.Operations = append(svc.Operations, b.manifestOp(op))
	}
	sort.Slice(svc.Operations, func(i, j int) bool {
		return svc.Operations[i].GoName < svc.Operations[j].GoName
	})
	return svc
}

func (b *builder) manifestOp(op *operation) manifestOp {
	m := manifestOp{
		ID:         op.ID,
		GoName:     op.GoName,
		Method:     op.Method,
		Path:       op.Path,
		Summary:    op.Summary,
		Resource:   op.Resource,
		Verb:       op.Verb,
		ParamsType: op.ParamsType,
		BodyType:   op.BodyType,
		ResultType: op.Result.Type,
		ItemType:   op.Result.ItemType,
		Paginated:  op.Paginated,
		Idempotent: op.Idempotent,
	}
	switch op.Result.Kind {
	case resultNone:
		m.ResultKind = "none"
	case resultValue:
		m.ResultKind = "value"
	case resultPage:
		m.ResultKind = "page"
	case resultStream:
		m.ResultKind = "stream"
	}
	switch op.Body {
	case bodyJSON:
		m.BodyKind = "json"
	case bodyStream:
		m.BodyKind = "stream"
	case bodyText:
		m.BodyKind = "text"
	}

	for _, p := range op.PathParams {
		m.PathParams = append(m.PathParams, manifestParam{
			Wire:   p.WireName,
			GoName: p.Name,
			Doc:    usageText(p.RawDoc),
		})
	}
	for _, p := range op.QueryParams {
		m.Params = append(m.Params, manifestField{
			Wire:     p.WireName,
			GoName:   p.Name,
			GoType:   p.Type,
			FlagKind: fallbackKind(flagKind(p.Type), b, p.Type),
			Pointer:  strings.HasPrefix(p.Type, "*"),
			Required: p.Required,
			Enum:     b.enumValues(p.Type),
			Doc:      usageText(p.RawDoc),
		})
	}
	if op.Body == bodyJSON {
		if nt, ok := b.types[strings.TrimPrefix(op.BodyType, "*")]; ok {
			for _, f := range nt.Fields {
				m.BodyFields = append(m.BodyFields, manifestField{
					Wire:     f.JSONName,
					GoName:   f.Name,
					GoType:   f.Type,
					FlagKind: fallbackKind(flagKind(f.Type), b, f.Type),
					Pointer:  strings.HasPrefix(f.Type, "*"),
					Required: f.Required,
					Enum:     mergeEnum(b.enumValues(f.Type), f.Enum),
					Doc:      usageText(f.RawDoc),
				})
			}
		}
	}
	return m
}

// fallbackKind resolves a named type — an enum over string, or an alias over a
// map — to the flag kind of what it actually is.
func fallbackKind(kind string, b *builder, goType string) string {
	if kind != "" {
		return kind
	}
	base := strings.TrimPrefix(goType, "*")
	if nt, ok := b.types[base]; ok {
		switch nt.Kind {
		case kindEnum:
			return "string"
		case kindAlias:
			return fallbackKind(flagKind(nt.Underlying), b, nt.Underlying)
		}
	}
	return "json"
}

// mergeEnum prefers a named type's values and falls back to an inline list.
func mergeEnum(named, inline []string) []string {
	if len(named) > 0 {
		return named
	}
	return inline
}

func (b *builder) enumValues(goType string) []string {
	nt, ok := b.types[strings.TrimPrefix(goType, "*")]
	if !ok || nt.Kind != kindEnum {
		return nil
	}
	out := make([]string, 0, len(nt.Enum))
	for _, e := range nt.Enum {
		out = append(out, e.Value)
	}
	return out
}

// usageText reduces a description to one line fit for a flag's usage string.
//
// Three things have to go. The rendered Go doc comment is prefixed with the
// field's own name, which reads as "AssignPublicIP the older spelling of…" on
// a command line. Markdown backticks are load-bearing to pflag, which reads
// the first backticked word as the flag's value placeholder and would print
// "--assign-public-ip networks[0].assign_public_ip". And a paragraph is too
// long, so it is cut at the first sentence rather than the first line, which
// would end mid-clause.
func usageText(doc string) string {
	s := strings.TrimSpace(strings.Join(strings.Fields(strings.ReplaceAll(doc, "\n", " ")), " "))
	s = strings.ReplaceAll(s, "`", "")
	s = strings.ReplaceAll(s, "**", "")
	if i := strings.Index(s, ". "); i > 0 {
		s = s[:i]
	}
	s = strings.TrimSuffix(s, ".")
	if s == "" {
		return ""
	}
	// A usage string sits after the flag name, so it reads better lowercase
	// unless it opens with a name or an acronym.
	return s
}

func writeManifest(outDir string, m manifest) error {
	sort.Slice(m.Services, func(i, j int) bool { return m.Services[i].Name < m.Services[j].Name })
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(filepath.Join(outDir, "api.json"), data, 0o644)
}
