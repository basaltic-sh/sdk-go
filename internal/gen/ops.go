package main

import (
	"fmt"
	"sort"
	"strings"
)

type resultKind int

const (
	resultNone   resultKind = iota // no response body worth returning
	resultValue                    // a single value, unwrapped from its envelope
	resultPage                     // a page of a list operation
	resultStream                   // raw bytes: an object body, a PDF, a screenshot
)

type bodyKind int

const (
	bodyNone bodyKind = iota
	bodyJSON
	bodyStream // io.Reader sent verbatim
	bodyText   // a string sent verbatim under an explicit content type
)

type param struct {
	Name     string // Go identifier
	WireName string // query or path parameter name
	Type     string // Go type
	Doc      []string
	Required bool
}

type result struct {
	Kind resultKind
	// Type is the Go return type: "*Instance", "map[string]any",
	// "*basaltic.Page[Instance]", "io.ReadCloser", or "" for none.
	Type string
	// EnvelopeJSON names the single property the value is wrapped in, when
	// the operation returns something like {"instance": {...}}. Empty means
	// the body decodes straight into Type.
	EnvelopeJSON  string
	EnvelopeField string
	// DecodeType is what the body is decoded into before unwrapping.
	DecodeType string

	// Page fields.
	ItemType  string
	ItemsJSON string
	HasMeta   bool
}

type operation struct {
	ID              string
	GoName          string
	Doc             []string
	Method          string
	Path            string
	PathParams      []*param
	QueryParams     []*param
	ParamsType      string
	Body            bodyKind
	BodyType        string
	BodyContentType string
	Unauthenticated bool
	Result          result
	// Paginated marks a list operation that also gets an "All" method
	// walking every page.
	Paginated bool
	// Idempotent notes that the operation accepts an Idempotency-Key, which
	// the doc comment mentions because it is also what makes it retryable.
	Idempotent bool
}

var httpMethods = []string{"get", "put", "post", "delete", "patch", "head", "options"}

// buildOperations walks the service's paths and turns each operation into the
// intermediate form the emitter renders.
func (b *builder) buildOperations(spec map[string]any, specFile string) error {
	paths := obj(spec, "paths")
	for _, p := range orderedKeys(paths) {
		item, itemRef, err := b.ld.deref(paths[p], specFile)
		if err != nil {
			return err
		}
		itemBase := specFile
		if itemRef.file != "" {
			itemBase = itemRef.file
		}
		im, ok := mapOf(item)
		if !ok {
			continue
		}
		for _, method := range httpMethods {
			raw, ok := im[method]
			if !ok {
				continue
			}
			om, ok := mapOf(raw)
			if !ok {
				continue
			}
			op, err := b.buildOperation(p, method, om, itemBase)
			if err != nil {
				return fmt.Errorf("%s %s: %w", strings.ToUpper(method), p, err)
			}
			b.ops = append(b.ops, op)
		}
	}
	sort.Slice(b.ops, func(i, j int) bool { return b.ops[i].GoName < b.ops[j].GoName })
	return nil
}

func (b *builder) buildOperation(path, method string, om map[string]any, base string) (*operation, error) {
	id := str(om, "operationId")
	if id == "" {
		return nil, fmt.Errorf("operation has no operationId")
	}
	op := &operation{
		ID:     id,
		GoName: exportedName(id),
		Method: strings.ToUpper(method),
		Path:   path,
	}

	// An explicitly empty security list marks the few operations whose
	// credentials are the request itself — the token endpoint.
	if sec, present := om["security"]; present {
		if l, ok := sec.([]any); ok && len(l) == 0 {
			op.Unauthenticated = true
		}
	}

	if err := b.buildParams(op, om, base); err != nil {
		return nil, err
	}
	if err := b.buildBody(op, om, base); err != nil {
		return nil, err
	}
	if err := b.buildResult(op, om, base); err != nil {
		return nil, err
	}
	op.Doc = b.operationDoc(op, om)
	return op, nil
}

func (b *builder) buildParams(op *operation, om map[string]any, base string) error {
	// Path placeholders are filled positionally, so their order must be the
	// order they appear in the path, not the order the spec lists them.
	placeholders := pathPlaceholders(op.Path)
	byWireName := map[string]*param{}

	for _, raw := range list(om, "parameters") {
		resolved, r, err := b.ld.deref(raw, base)
		if err != nil {
			return err
		}
		pm, ok := mapOf(resolved)
		if !ok {
			continue
		}
		paramBase := base
		if r.file != "" {
			paramBase = r.file
		}
		wire := str(pm, "name")
		in := str(pm, "in")

		if in == "header" {
			// Idempotency-Key is the only header parameter in the platform,
			// and it is a per-call concern rather than a signature one — the
			// SDK carries it as basaltic.WithIdempotencyKey.
			if wire == "Idempotency-Key" {
				op.Idempotent = true
			}
			continue
		}

		schema := pm["schema"]
		typ, err := b.typeForSchema(schema, paramBase, op.GoName+exportedName(wire), false)
		if err != nil {
			return err
		}
		required := boolOf(pm, "required")

		doc := str(pm, "description")
		if sm, ok := mapOf(schema); ok {
			if enums := list(sm, "enum"); len(enums) > 0 {
				vals := make([]string, 0, len(enums))
				for _, e := range enums {
					vals = append(vals, fmt.Sprintf("%q", fmt.Sprint(e)))
				}
				doc = strings.TrimSpace(doc + "\n\nOne of: " + strings.Join(vals, ", ") + ".")
			}
		}

		p := &param{
			WireName: wire,
			Type:     typ,
			Required: required,
		}
		switch in {
		case "path":
			p.Name = unexportedName(wire)
			p.Doc = goDoc("", wrapText(doc, 68))
			byWireName[wire] = p
		case "query":
			p.Name = exportedName(wire)
			// A boolean filter is the one query parameter where "false" and
			// "not set" mean different things, so it takes a pointer. The
			// rest use plain values and are simply omitted when zero, which
			// is what an empty filter means anyway.
			if typ == "bool" && !required {
				p.Type = "*bool"
			}
			p.Doc = goDoc(p.Name, wrapText(doc, 68))
			op.QueryParams = append(op.QueryParams, p)
		}
	}

	for _, ph := range placeholders {
		p, ok := byWireName[ph]
		if !ok {
			// A placeholder the specification does not document. Take it as a
			// string rather than dropping it, which would emit a call that
			// cannot address the resource.
			p = &param{Name: unexportedName(ph), WireName: ph, Type: "string", Required: true}
		}
		if p.Type != "string" {
			p.Type = "string"
		}
		op.PathParams = append(op.PathParams, p)
	}

	if len(op.QueryParams) > 0 {
		sort.Slice(op.QueryParams, func(i, j int) bool { return op.QueryParams[i].Name < op.QueryParams[j].Name })
		op.ParamsType = op.GoName + "Params"
	}
	return nil
}

func (b *builder) buildBody(op *operation, om map[string]any, base string) error {
	rbRaw, ok := om["requestBody"]
	if !ok {
		return nil
	}
	resolved, r, err := b.ld.deref(rbRaw, base)
	if err != nil {
		return err
	}
	rb, ok := mapOf(resolved)
	if !ok {
		return nil
	}
	bodyBase := base
	if r.file != "" {
		bodyBase = r.file
	}
	content := obj(rb, "content")
	if len(content) == 0 {
		return nil
	}

	// JSON wins wherever the operation accepts it, which keeps the three
	// multi-format bodies — a zone file import, the two OAuth endpoints —
	// on one encoding rather than exposing the alternatives.
	if mt, ok := content["application/json"]; ok {
		mm, _ := mapOf(mt)
		typ, err := b.typeForSchema(mm["schema"], bodyBase, op.GoName+"Request", true)
		if err != nil {
			return err
		}
		op.Body = bodyJSON
		op.BodyType = typ
		op.BodyContentType = "application/json"
		return nil
	}

	ct := orderedKeys(content)[0]
	mm, _ := mapOf(content[ct])
	sm, _ := mapOf(mm["schema"])
	op.BodyContentType = ct
	if sm != nil && str(sm, "type") == "string" && str(sm, "format") != "binary" {
		op.Body = bodyText
		op.BodyType = "string"
		return nil
	}
	op.Body = bodyStream
	op.BodyType = "io.Reader"
	return nil
}

func (b *builder) buildResult(op *operation, om map[string]any, base string) error {
	responses := obj(om, "responses")
	codes := orderedKeys(responses)
	sort.Strings(codes)

	var chosen map[string]any
	chosenBase := base
	for _, code := range codes {
		if !strings.HasPrefix(code, "2") {
			continue
		}
		resolved, r, err := b.ld.deref(responses[code], base)
		if err != nil {
			return err
		}
		rm, ok := mapOf(resolved)
		if !ok {
			continue
		}
		if len(obj(rm, "content")) == 0 {
			continue
		}
		chosen = rm
		if r.file != "" {
			chosenBase = r.file
		}
		break
	}
	if chosen == nil {
		op.Result = result{Kind: resultNone}
		return nil
	}

	content := obj(chosen, "content")
	// Anything the platform answers with that is not JSON is bytes: an
	// object's contents, an invoice PDF, a console screenshot, a zone file.
	// Those stream rather than decode.
	jsonMT, hasJSON := content["application/json"]
	if !hasJSON || len(content) > 1 {
		nonJSON := false
		for ct := range content {
			if ct != "application/json" {
				nonJSON = true
			}
		}
		if nonJSON {
			op.Result = result{Kind: resultStream, Type: "io.ReadCloser"}
			return nil
		}
	}

	mm, _ := mapOf(jsonMT)
	schema := mm["schema"]
	sm, ok := mapOf(schema)
	if !ok {
		op.Result = result{Kind: resultNone}
		return nil
	}

	// A $ref to a named schema: either a list envelope, which becomes a Page,
	// or a resource, which is returned as-is.
	if refStr, isRef := sm["$ref"].(string); isRef {
		resolved, r, err := b.ld.resolve(refStr, chosenBase)
		if err != nil {
			return err
		}
		rm, _ := mapOf(resolved)
		if itemsJSON, itemsSchema, hasMeta, isList := listEnvelope(rm); isList {
			itemType, err := b.typeForSchema(itemsSchema, r.file, exportedName(singular(itemsJSON)), false)
			if err != nil {
				return err
			}
			itemType = strings.TrimPrefix(itemType, "[]")
			op.Result = result{
				Kind:      resultPage,
				Type:      "*basaltic.Page[" + strings.TrimPrefix(itemType, "*") + "]",
				ItemType:  strings.TrimPrefix(itemType, "*"),
				ItemsJSON: itemsJSON,
				HasMeta:   hasMeta,
			}
			op.Paginated = hasMeta && hasQueryParam(op, "marker")
			return nil
		}
		// A named "…Response" schema with a single property is an envelope,
		// the same one the inline responses write out longhand. Unwrap it so
		// that getInstancePool returns *InstancePool like getInstance returns
		// *Instance, instead of a wrapper type that exists only to be
		// dereferenced. Registering the wrapper is skipped entirely, so the
		// package does not carry ~45 types nothing names.
		if key, inner, ok := singleFieldEnvelope(rm); ok && strings.HasSuffix(r.name(), "Response") {
			typ, err := b.typeForSchema(inner, r.file, op.GoName+exportedName(key), false)
			if err != nil {
				return err
			}
			op.Result = result{
				Kind:          resultValue,
				Type:          typ,
				EnvelopeJSON:  key,
				EnvelopeField: exportedName(key),
			}
			return nil
		}

		typ, err := b.typeForSchema(sm, chosenBase, op.GoName+"Response", false)
		if err != nil {
			return err
		}
		op.Result = result{Kind: resultValue, Type: typ, DecodeType: typ}
		return nil
	}

	// An inline object. One property is an envelope around the real result;
	// anything else becomes a small result type of its own.
	props := obj(sm, "properties")
	switch len(props) {
	case 0:
		if str(sm, "type") == "object" {
			op.Result = result{Kind: resultValue, Type: "map[string]any", DecodeType: "map[string]any"}
			return nil
		}
		op.Result = result{Kind: resultNone}
		return nil
	case 1:
		key := orderedKeys(props)[0]
		typ, err := b.typeForSchema(props[key], chosenBase, op.GoName+exportedName(key), false)
		if err != nil {
			return err
		}
		op.Result = result{
			Kind:          resultValue,
			Type:          typ,
			EnvelopeJSON:  key,
			EnvelopeField: exportedName(key),
		}
		return nil
	default:
		typ, err := b.typeForSchema(sm, chosenBase, op.GoName+"Result", false)
		if err != nil {
			return err
		}
		op.Result = result{Kind: resultValue, Type: typ, DecodeType: typ}
		return nil
	}
}

// listEnvelope recognises a list response: exactly one array property, with
// nothing beside it but the pagination metadata.
//
// The shapes that fail this test are real and must not become pages — the
// S3-style object listing carries common_prefixes and is_truncated, and the
// price list carries an as_of timestamp. Those are returned whole.
func listEnvelope(m map[string]any) (itemsKey string, itemsSchema node, hasMeta, ok bool) {
	if m == nil {
		return "", nil, false, false
	}
	props := obj(m, "properties")
	if len(props) == 0 || len(props) > 2 {
		return "", nil, false, false
	}
	for name, raw := range props {
		pm, isMap := mapOf(raw)
		if name == "meta" {
			hasMeta = true
			continue
		}
		if !isMap || str(pm, "type") != "array" {
			return "", nil, false, false
		}
		if itemsKey != "" {
			return "", nil, false, false
		}
		itemsKey, itemsSchema = name, pm["items"]
	}
	if itemsKey == "" {
		return "", nil, false, false
	}
	if len(props) == 2 && !hasMeta {
		return "", nil, false, false
	}
	return itemsKey, itemsSchema, hasMeta, true
}

// singleFieldEnvelope recognises a schema whose only job is to wrap one
// value under a key.
func singleFieldEnvelope(m map[string]any) (key string, inner node, ok bool) {
	if m == nil || str(m, "type") == "array" {
		return "", nil, false
	}
	props := obj(m, "properties")
	if len(props) != 1 {
		return "", nil, false
	}
	for k, v := range props {
		return k, v, true
	}
	return "", nil, false
}

func hasQueryParam(op *operation, name string) bool {
	for _, p := range op.QueryParams {
		if p.WireName == name {
			return true
		}
	}
	return false
}

func (b *builder) operationDoc(op *operation, om map[string]any) []string {
	summary := strings.TrimSpace(str(om, "summary"))
	description := strings.TrimSpace(str(om, "description"))

	var body string
	switch {
	case summary != "" && description != "":
		// Avoid saying the same thing twice when the description merely
		// restates the summary.
		if strings.EqualFold(strings.TrimRight(description, "."), strings.TrimRight(summary, ".")) {
			body = summary
		} else {
			body = summary + ".\n\n" + description
		}
	case summary != "":
		body = summary
	default:
		body = description
	}
	body = strings.TrimSuffix(strings.TrimSpace(body), ".")
	if body != "" {
		body += "."
	}
	body = conjugateSummary(op.GoName, body)

	var extra []string
	if op.Paginated {
		extra = append(extra, fmt.Sprintf(
			"Returns one page. Use %sAll to walk every page.", op.GoName))
	}
	if op.Idempotent {
		extra = append(extra, "Accepts basaltic.WithIdempotencyKey, which makes the call\nreplay-safe and therefore retryable.")
	}
	if op.Result.Kind == resultStream {
		extra = append(extra, "The caller must close the returned reader.")
	}
	if op.Unauthenticated {
		extra = append(extra, "Sends no bearer token: the credentials in the request are the\nauthentication.")
	}
	return goDoc(op.GoName, wrapText(body, 70), extra...)
}

// conjugateSummary turns the specification's imperative summary into the
// third-person form a Go doc comment wants: "listInstances / List instances"
// becomes "ListInstances lists instances".
//
// It only fires when the summary opens with the same verb the method name
// does, which is the case for almost every operation and is a precise enough
// test that a summary phrased some other way is left exactly as written.
func conjugateSummary(goName, body string) string {
	if body == "" {
		return body
	}
	first, rest, _ := strings.Cut(body, " ")
	word := strings.ToLower(strings.TrimRight(first, ".,:;"))
	// Whatever trailed the verb — the comma in "Rename, scale, or resize" —
	// belongs to the sentence, not to the verb, and has to come back.
	punct := first[len(word):]

	// The verb the method name itself opens with is always the right one to
	// conjugate; beyond that, only a word known to be a verb. The summaries
	// that open with a noun — "Instant structured metric query" — must not be
	// bent into one.
	if word != leadingWord(goName) && !summaryVerbs[word] {
		return "\u2014 " + body
	}
	conjugated := thirdPerson(word) + punct
	if rest == "" {
		return conjugated
	}
	return conjugated + " " + rest
}

// summaryVerbs are the imperative verbs the specifications open summaries
// with. A word absent from this list is left alone rather than guessed at, so
// the cost of an omission is a slightly stilted comment, not a wrong one.
var summaryVerbs = map[string]bool{
	"accept": true, "activate": true, "add": true, "allocate": true,
	"apply": true, "assume": true, "attach": true, "cancel": true,
	"capture": true, "close": true, "configure": true, "count": true,
	"create": true, "deactivate": true, "decrypt": true, "delete": true,
	"deregister": true, "describe": true, "detach": true, "disable": true,
	"disassociate": true, "download": true, "drain": true, "empty": true,
	"enable": true, "encrypt": true, "exchange": true, "export": true,
	"failover": true, "fetch": true, "generate": true, "get": true,
	"give": true, "import": true, "invite": true, "list": true, "move": true,
	"open": true, "patch": true, "promote": true, "purge": true, "put": true,
	"read": true, "reboot": true, "refresh": true, "register": true,
	"reinstall": true, "reject": true, "release": true, "remove": true,
	"rename": true, "replace": true, "resize": true, "restore": true,
	"retry": true, "revoke": true, "roll": true, "rotate": true,
	"scale": true, "schedule": true, "search": true, "send": true,
	"set": true, "sign": true, "start": true, "stop": true, "store": true,
	"suspend": true, "switchover": true, "take": true, "trigger": true,
	"update": true, "upload": true, "validate": true, "verify": true,
}

// leadingWord returns the first PascalCase word of an identifier.
func leadingWord(s string) string {
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}
	return words[0]
}

// thirdPerson applies English third-person-singular spelling to a verb.
func thirdPerson(verb string) string {
	switch {
	case strings.HasSuffix(verb, "s"), strings.HasSuffix(verb, "x"),
		strings.HasSuffix(verb, "z"), strings.HasSuffix(verb, "ch"),
		strings.HasSuffix(verb, "sh"), strings.HasSuffix(verb, "o"):
		return verb + "es"
	case strings.HasSuffix(verb, "y") && len(verb) > 1 && !isVowel(verb[len(verb)-2]):
		return verb[:len(verb)-1] + "ies"
	}
	return verb + "s"
}

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

// unexportedName renders a wire name as an unexported Go identifier suitable
// for a parameter.
//
// It works from the words rather than from the assembled identifier:
// lowercasing the leading run of capitals in "VPCID" would yield "vpcid",
// losing the boundary between the two acronyms.
func unexportedName(s string) string {
	words := splitWords(s)
	if len(words) == 0 {
		return "arg"
	}
	out := words[0]
	for _, w := range words[1:] {
		if up, ok := initialisms[w]; ok {
			out += up
			continue
		}
		out += strings.ToUpper(w[:1]) + w[1:]
	}
	if out == "" {
		return "arg"
	}
	if c := out[0]; c >= '0' && c <= '9' {
		out = "n" + out
	}
	if isGoKeyword(out) {
		return out + "_"
	}
	return out
}

var goKeywords = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true,
	"interface": true, "map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true, "var": true,
	"len": true, "cap": true, "new": true, "make": true, "copy": true, "close": true,
}

func isGoKeyword(s string) bool { return goKeywords[s] }

// pathPlaceholders lists the {placeholder} names in a path, in order.
func pathPlaceholders(path string) []string {
	var out []string
	rest := path
	for {
		open := strings.IndexByte(rest, '{')
		if open < 0 {
			return out
		}
		end := strings.IndexByte(rest[open:], '}')
		if end < 0 {
			return out
		}
		end += open
		out = append(out, rest[open+1:end])
		rest = rest[end+1:]
	}
}
