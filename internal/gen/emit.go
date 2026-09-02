package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const genHeader = `// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi
`

// emitService writes one service package.
func emitService(dir string, b *builder, info map[string]any) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	files := map[string]func() ([]byte, error){
		"client_gen.go": func() ([]byte, error) { return b.emitClient(info) },
		"types_gen.go":  b.emitTypes,
		"api_gen.go":    b.emitAPI,
	}
	for name, render := range files {
		src, err := render()
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		formatted, err := format.Source(src)
		if err != nil {
			// Write the unformatted source so the syntax error can be read
			// where it happened rather than guessed at from a line number.
			_ = os.WriteFile(filepath.Join(dir, name+".broken"), src, 0o644)
			return fmt.Errorf("%s: %w (unformatted source written to %s.broken)", name, err, name)
		}
		if err := os.WriteFile(filepath.Join(dir, name), formatted, 0o644); err != nil {
			return err
		}
	}
	return nil
}

// fileHeader renders the generated-code banner, package clause and imports.
func fileHeader(pkg string, doc []string, body string, extraImports ...string) []byte {
	var b bytes.Buffer
	b.WriteString(genHeader)
	b.WriteString("\n")
	b.WriteString(commentLines("", doc))
	fmt.Fprintf(&b, "package %s\n\n", pkg)

	std := map[string]bool{}
	for _, imp := range extraImports {
		std[imp] = true
	}
	// The import set is small and fixed, so deducing it from the rendered
	// body is reliable and saves threading a set through every emitter.
	// Comments are stripped first: a doc comment mentioning "context" is
	// prose, not a dependency, and would otherwise import a package the file
	// never uses.
	code := stripComments(body)
	for marker, imp := range map[string]string{
		"context.":        "context",
		"time.Time":       "time",
		"time.Duration":   "time",
		"url.Values":      "net/url",
		"strconv.":        "strconv",
		"io.Reader":       "io",
		"io.ReadCloser":   "io",
		"iter.Seq2":       "iter",
		"json.RawMessage": "encoding/json",
		"basaltic.":       "github.com/basaltic-sh/sdk-go",
		"fmt.":            "fmt",
		"strings.":        "strings",
	} {
		if strings.Contains(code, marker) {
			std[imp] = true
		}
	}

	var stdlib, external []string
	for imp := range std {
		if strings.Contains(imp, ".") || strings.Contains(imp, "/") && strings.Contains(strings.SplitN(imp, "/", 2)[0], ".") {
			external = append(external, imp)
			continue
		}
		stdlib = append(stdlib, imp)
	}
	sort.Strings(stdlib)
	sort.Strings(external)

	if len(stdlib)+len(external) > 0 {
		b.WriteString("import (\n")
		for _, imp := range stdlib {
			fmt.Fprintf(&b, "\t%q\n", imp)
		}
		if len(external) > 0 && len(stdlib) > 0 {
			b.WriteString("\n")
		}
		for _, imp := range external {
			if imp == "github.com/basaltic-sh/sdk-go" {
				fmt.Fprintf(&b, "\tbasaltic %q\n", imp)
				continue
			}
			fmt.Fprintf(&b, "\t%q\n", imp)
		}
		b.WriteString(")\n\n")
	}
	b.WriteString(body)
	return b.Bytes()
}

func (b *builder) emitClient(info map[string]any) ([]byte, error) {
	title := str(info, "title")
	description := strings.TrimSpace(str(info, "description"))

	pkgDoc := []string{
		fmt.Sprintf("Package %s is the %s.", b.service, strings.TrimPrefix(title, "Basaltic ")),
	}
	if description != "" {
		pkgDoc = append(pkgDoc, "")
		pkgDoc = append(pkgDoc, strings.Split(wrapText(description, 70), "\n")...)
	}
	pkgDoc = append(pkgDoc,
		"",
		"Build a client from a shared [basaltic.Config]:",
		"",
		fmt.Sprintf("\tc := %s.New(cfg)", b.service),
		"",
		"Clients are safe for concurrent use.",
	)

	var body strings.Builder
	fmt.Fprintf(&body, "// ServiceID is the short name the SDK addresses this service by. Use it\n"+
		"// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.\n"+
		"const ServiceID = %q\n\n", b.service)

	fmt.Fprintf(&body, "// endpointTemplate is the server URL this service's specification\n"+
		"// declares. Any {region} in it is substituted per request.\n"+
		"const endpointTemplate = %q\n\n", b.serverTemplate)

	body.WriteString("func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }\n\n")

	fmt.Fprintf(&body, "// Client calls the %s.\n//\n"+
		"// Build one with [New]. It is safe for concurrent use.\ntype Client struct {\n\trt *basaltic.Client\n}\n\n",
		strings.TrimPrefix(title, "Basaltic "))

	fmt.Fprintf(&body, "// New builds a %s client from a shared configuration.\n//\n"+
		"// Share one [basaltic.Config] across every service client: they then\n"+
		"// share a token, so authenticating costs one exchange rather than one\n"+
		"// per service.\nfunc New(cfg *basaltic.Config) *Client {\n"+
		"\treturn &Client{rt: basaltic.NewClient(cfg, ServiceID)}\n}\n\n", b.service)

	body.WriteString("// Transport returns the underlying transport, for reaching an endpoint\n" +
		"// this package does not generate. See [basaltic.Client.Do].\nfunc (c *Client) Transport() *basaltic.Client { return c.rt }\n")

	return fileHeader(b.service, pkgDoc, body.String()), nil
}

func (b *builder) emitTypes() ([]byte, error) {
	var body strings.Builder
	for _, name := range orderedKeys(b.types) {
		nt := b.types[name]
		switch nt.Kind {
		case kindEnum:
			body.WriteString(commentLines("", nt.Doc))
			fmt.Fprintf(&body, "type %s %s\n\n", nt.Name, nt.Underlying)
			if len(nt.Enum) > 0 {
				fmt.Fprintf(&body, "// Values %s accepts.\nconst (\n", nt.Name)
				for _, c := range nt.Enum {
					fmt.Fprintf(&body, "\t%s %s = %q\n", c.Name, nt.Name, c.Value)
				}
				body.WriteString(")\n\n")
			}
		case kindAlias:
			body.WriteString(commentLines("", nt.Doc))
			fmt.Fprintf(&body, "type %s = %s\n\n", nt.Name, nt.Underlying)
		default:
			body.WriteString(commentLines("", nt.Doc))
			fmt.Fprintf(&body, "type %s struct {\n", nt.Name)
			for i, f := range nt.Fields {
				if i > 0 && len(f.Doc) > 0 {
					body.WriteString("\n")
				}
				body.WriteString(commentLines("\t", f.Doc))
				tag := f.JSONName
				if f.OmitEmpty {
					tag += ",omitempty"
				}
				fmt.Fprintf(&body, "\t%s %s `json:\"%s\"`\n", f.Name, f.Type, tag)
			}
			body.WriteString("}\n\n")
		}
	}
	return fileHeader(b.service, nil, body.String()), nil
}

func (b *builder) emitAPI() ([]byte, error) {
	var body strings.Builder
	for _, op := range b.ops {
		b.emitParamsType(&body, op)
	}
	for _, op := range b.ops {
		if err := b.emitOperation(&body, op); err != nil {
			return nil, err
		}
	}
	return fileHeader(b.service, nil, body.String()), nil
}

func (b *builder) emitParamsType(body *strings.Builder, op *operation) {
	if op.ParamsType == "" {
		return
	}
	fmt.Fprintf(body, "// %s are the optional filters and pagination controls for\n// [Client.%s]. A nil *%s sends none of them.\n",
		op.ParamsType, op.GoName, op.ParamsType)
	fmt.Fprintf(body, "type %s struct {\n", op.ParamsType)
	for i, p := range op.QueryParams {
		if i > 0 && len(p.Doc) > 0 {
			body.WriteString("\n")
		}
		body.WriteString(commentLines("\t", p.Doc))
		fmt.Fprintf(body, "\t%s %s\n", p.Name, p.Type)
	}
	body.WriteString("}\n\n")

	fmt.Fprintf(body, "// query renders the parameters that are set. A zero value means \"no\n// filter\", which is what leaving one out asks for.\n")
	fmt.Fprintf(body, "func (p *%s) query() url.Values {\n\tq := url.Values{}\n\tif p == nil {\n\t\treturn q\n\t}\n", op.ParamsType)
	for _, p := range op.QueryParams {
		body.WriteString(b.queryEncode(p))
	}
	body.WriteString("\treturn q\n}\n\n")

	if op.Paginated {
		fmt.Fprintf(body, "// withMarker copies p with the pagination cursor replaced, leaving the\n// caller's value untouched across pages.\n")
		fmt.Fprintf(body, "func (p *%s) withMarker(marker string) *%s {\n\tvar out %s\n\tif p != nil {\n\t\tout = *p\n\t}\n\tout.Marker = marker\n\treturn &out\n}\n\n",
			op.ParamsType, op.ParamsType, op.ParamsType)
	}
}

// queryEncode renders the code that adds one parameter to the query, when it
// has been set.
func (b *builder) queryEncode(p *param) string {
	field := "p." + p.Name
	wire := fmt.Sprintf("%q", p.WireName)
	set := func(expr string) string {
		return fmt.Sprintf("\t\tq.Set(%s, %s)\n", wire, expr)
	}
	guard := func(cond, expr string) string {
		return fmt.Sprintf("\tif %s {\n%s\t}\n", cond, set(expr))
	}

	typ := p.Type
	switch {
	case typ == "*bool":
		return guard(field+" != nil", "strconv.FormatBool(*"+field+")")
	case typ == "bool":
		return guard(field, `"true"`)
	case typ == "string":
		return guard(field+` != ""`, field)
	case typ == "int" || typ == "int32":
		return guard(field+" != 0", "strconv.Itoa(int("+field+"))")
	case typ == "int64":
		return guard(field+" != 0", "strconv.FormatInt("+field+", 10)")
	case typ == "float64":
		return guard(field+" != 0", `strconv.FormatFloat(`+field+`, 'f', -1, 64)`)
	case typ == "time.Time":
		return guard("!"+field+".IsZero()", field+".UTC().Format(time.RFC3339)")
	case typ == "[]string":
		return fmt.Sprintf("\tfor _, v := range %s {\n\t\tq.Add(%s, v)\n\t}\n", field, wire)
	case strings.HasPrefix(typ, "[]"):
		return fmt.Sprintf("\tfor _, v := range %s {\n\t\tq.Add(%s, fmt.Sprint(v))\n\t}\n", field, wire)
	}
	// A named type over string — an enum from the specification.
	if nt, ok := b.types[typ]; ok && nt.Kind == kindEnum {
		return guard(field+` != ""`, "string("+field+")")
	}
	return guard(fmt.Sprintf("fmt.Sprint(%s) != \"\"", field), "fmt.Sprint("+field+")")
}

func (b *builder) emitOperation(body *strings.Builder, op *operation) error {
	args := []string{"ctx context.Context"}
	for _, p := range op.PathParams {
		args = append(args, p.Name+" "+p.Type)
	}
	switch {
	case op.ParamsType != "":
		args = append(args, "params *"+op.ParamsType)
	case op.Body == bodyJSON:
		args = append(args, "body "+op.BodyType)
	case op.Body == bodyStream:
		args = append(args, "body io.Reader")
	case op.Body == bodyText:
		args = append(args, "body string")
	}
	args = append(args, "opts ...basaltic.RequestOption")

	ret := "error"
	if op.Result.Kind != resultNone {
		ret = "(" + op.Result.Type + ", error)"
	}

	body.WriteString(commentLines("", op.Doc))
	fmt.Fprintf(body, "func (c *Client) %s(%s) %s {\n", op.GoName, strings.Join(args, ", "), ret)

	// The operation descriptor.
	fmt.Fprintf(body, "\top := &basaltic.Operation{\n\t\tID:     %q,\n\t\tMethod: %q,\n\t\tPath:   %q,\n", op.ID, op.Method, op.Path)
	if len(op.PathParams) > 0 {
		names := make([]string, 0, len(op.PathParams))
		for _, p := range op.PathParams {
			names = append(names, p.Name)
		}
		fmt.Fprintf(body, "\t\tPathArgs: []string{%s},\n", strings.Join(names, ", "))
	}
	switch op.Body {
	case bodyJSON:
		body.WriteString("\t\tBody: body,\n")
	case bodyStream:
		fmt.Fprintf(body, "\t\tStream:      body,\n\t\tContentType: %q,\n", op.BodyContentType)
	case bodyText:
		fmt.Fprintf(body, "\t\tRawBody:     []byte(body),\n\t\tContentType: %q,\n", op.BodyContentType)
	}
	if op.Unauthenticated {
		body.WriteString("\t\tUnauthenticated: true,\n")
	}
	body.WriteString("\t}\n")
	if op.ParamsType != "" {
		body.WriteString("\top.Query = params.query()\n")
	}

	fail := "return err"
	if op.Result.Kind != resultNone {
		fail = "return " + b.zeroValue(op.Result.Type) + ", err"
	}

	switch op.Result.Kind {
	case resultNone:
		fmt.Fprintf(body, "\tif err := c.rt.Do(ctx, op, nil, opts...); err != nil {\n\t\t%s\n\t}\n\treturn nil\n", fail)

	case resultStream:
		fmt.Fprintf(body, "\tstream, _, err := c.rt.DoStream(ctx, op, opts...)\n\tif err != nil {\n\t\t%s\n\t}\n\treturn stream, nil\n", fail)

	case resultPage:
		fmt.Fprintf(body, "\tvar out struct {\n\t\tItems []%s `json:%q`\n", op.Result.ItemType, op.Result.ItemsJSON)
		if op.Result.HasMeta {
			body.WriteString("\t\tMeta  *struct {\n" +
				"\t\t\tTotal   int    `json:\"total\"`\n" +
				"\t\t\tLimit   int    `json:\"limit\"`\n" +
				"\t\t\tMarker  string `json:\"marker\"`\n" +
				"\t\t\tHasMore bool   `json:\"has_more\"`\n" +
				"\t\t} `json:\"meta\"`\n")
		}
		body.WriteString("\t}\n")
		fmt.Fprintf(body, "\tif err := c.rt.Do(ctx, op, &out, opts...); err != nil {\n\t\t%s\n\t}\n", fail)
		fmt.Fprintf(body, "\tpage := &basaltic.Page[%s]{Items: out.Items}\n", op.Result.ItemType)
		if op.Result.HasMeta {
			body.WriteString("\tif out.Meta != nil {\n" +
				"\t\tpage.Total = out.Meta.Total\n" +
				"\t\tpage.Limit = out.Meta.Limit\n" +
				"\t\tpage.Marker = out.Meta.Marker\n" +
				"\t\tpage.HasMore = out.Meta.HasMore\n\t}\n")
		}
		body.WriteString("\treturn page, nil\n")

	case resultValue:
		if op.Result.EnvelopeJSON != "" {
			fmt.Fprintf(body, "\tvar out struct {\n\t\t%s %s `json:%q`\n\t}\n",
				op.Result.EnvelopeField, op.Result.Type, op.Result.EnvelopeJSON)
			fmt.Fprintf(body, "\tif err := c.rt.Do(ctx, op, &out, opts...); err != nil {\n\t\t%s\n\t}\n", fail)
			fmt.Fprintf(body, "\treturn out.%s, nil\n", op.Result.EnvelopeField)
		} else if strings.HasPrefix(op.Result.Type, "*") {
			base := strings.TrimPrefix(op.Result.Type, "*")
			fmt.Fprintf(body, "\tvar out %s\n", base)
			fmt.Fprintf(body, "\tif err := c.rt.Do(ctx, op, &out, opts...); err != nil {\n\t\t%s\n\t}\n", fail)
			body.WriteString("\treturn &out, nil\n")
		} else {
			fmt.Fprintf(body, "\tvar out %s\n", op.Result.Type)
			fmt.Fprintf(body, "\tif err := c.rt.Do(ctx, op, &out, opts...); err != nil {\n\t\t%s\n\t}\n", fail)
			body.WriteString("\treturn out, nil\n")
		}
	}
	body.WriteString("}\n\n")

	if op.Paginated {
		b.emitPaginator(body, op)
	}
	return nil
}

// zeroValue is what an operation returns alongside an error. Most results
// are pointers, but unwrapping a single-field envelope can yield a bare bool
// or string — kms.Verify answers whether a signature checked out, not a
// wrapper around it — and those cannot return nil.
func (b *builder) zeroValue(t string) string {
	switch t {
	case "":
		return "nil"
	case "string":
		return `""`
	case "bool":
		return "false"
	case "int", "int32", "int64", "float64":
		return "0"
	case "time.Time":
		return "time.Time{}"
	}
	if strings.HasPrefix(t, "*") || strings.HasPrefix(t, "[]") ||
		strings.HasPrefix(t, "map[") || t == "any" || t == "io.ReadCloser" {
		return "nil"
	}
	if nt, ok := b.types[t]; ok {
		switch nt.Kind {
		case kindEnum:
			return `""`
		case kindAlias:
			return b.zeroValue(nt.Underlying)
		}
	}
	return t + "{}"
}

func (b *builder) emitPaginator(body *strings.Builder, op *operation) {
	item := op.Result.ItemType

	args := []string{"ctx context.Context"}
	callArgs := []string{"ctx"}
	for _, p := range op.PathParams {
		args = append(args, p.Name+" "+p.Type)
		callArgs = append(callArgs, p.Name)
	}
	args = append(args, "params *"+op.ParamsType, "opts ...basaltic.RequestOption")
	callArgs = append(callArgs, "params.withMarker(marker)", "opts...")

	sample := "c." + op.GoName + "All(ctx, "
	for _, p := range op.PathParams {
		sample += p.Name + ", "
	}
	sample += "nil)"

	doc := goDoc(op.GoName+"All", wrapText(fmt.Sprintf(
		"%sAll walks every page of %s, yielding one item at a time.", op.GoName, op.GoName), 70),
		"The iterator stops at the first error, yielding it alongside a zero\n"+
			"value, so check err on every step:\n\n"+
			fmt.Sprintf("\tfor item, err := range %s {\n", sample)+
			"\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\t...\n\t}\n\n"+
			"Breaking out of the loop stops the walk; no further requests are made.\n"+
			"Any Marker on params is overwritten as the walk advances.")
	body.WriteString(commentLines("", doc))
	fmt.Fprintf(body, "func (c *Client) %sAll(%s) iter.Seq2[%s, error] {\n",
		op.GoName, strings.Join(args, ", "), item)
	fmt.Fprintf(body, "\treturn basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[%s], error) {\n", item)
	fmt.Fprintf(body, "\t\treturn c.%s(%s)\n\t})\n}\n\n", op.GoName, strings.Join(callArgs, ", "))
}

// stripComments removes comment lines so that prose in a doc comment cannot
// be mistaken for a use of a package.
func stripComments(body string) string {
	var b strings.Builder
	for _, line := range strings.Split(body, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "//") {
			continue
		}
		b.WriteString(line)
		b.WriteByte('\n')
	}
	return b.String()
}
