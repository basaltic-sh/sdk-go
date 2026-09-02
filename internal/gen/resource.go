package main

import (
	"strings"
)

// Deriving the CLI's noun and verb for each operation.
//
// The command tree is `basaltic <service> <resource> <verb>`, so every
// operation needs a resource and a verb. Neither can come from the OpenAPI
// tags: those are a documentation grouping, which is why the previous CLI
// listed `buckets` beside `storage` and `images` beside `compute`.
//
// The hard case is a nested collection. Some are genuinely a resource — a load
// balancer's listeners have a lifecycle of their own — and some are an
// attachment or a sub-document belonging to the parent. Guessing from the
// shape of the path gets this wrong in both directions.
//
// The specifications already draw the line, in the operation id they give the
// POST: `createListener` makes listeners a resource, while `attachTarget` and
// `putUserInlinePolicy` do not. So the derivation runs in two passes — first
// find every collection the API creates members of, then place each operation
// against that.
//
// x-resource is honoured where it can be located in the path. It is set on
// only two thirds of operations and it is sometimes a service name rather than
// a resource (every billing operation is tagged `Billing`, across six distinct
// collections), so "can I find this noun in the path?" is what separates the
// useful settings from the ones that would collapse six resources into one.

// resourceOverrides settles the cases no rule gets right. Each is a judgment
// made once, keyed by operation id.
var resourceOverrides = map[string]struct{ Resource, Verb string }{
	// The console is a facet of an instance. Left to the rules, `console`
	// would become a noun of its own.
	"getConsoleOutput":     {"instance", "console-output"},
	"getConsoleScreenshot": {"instance", "console-screenshot"},
	"startSerialConsole":   {"instance", "serial-console"},
	// Mints a short-lived ticket for a browser, whose WebSocket constructor
	// takes a URL and cannot send an Authorization header. Named for what it
	// produces, not for the verb the path uses.
	"createSerialConsoleTicket": {"instance", "console-ticket"},

	// Two paths delete the same rule. The one nested under its listener is
	// the documented route and keeps the plain verb.
	"deleteRuleInListener": {"rule", "delete"},
	"deleteRule":           {"rule", "delete-orphaned"},

	// Reads as an action on the pool rather than a resource called "instance".
	"listPoolInstances": {"instance-pool", "list-instances"},

	// Metric queries are facets of metrics; the POST forms differ only in how
	// the query is carried.
	"queryMetricsInstant":     {"metric", "query"},
	"queryMetricsInstantPost": {"metric", "query-body"},
	"queryMetricsRange":       {"metric", "query-range"},
	"queryMetricsRangePost":   {"metric", "query-range-body"},
	"listMetricSeries":        {"metric", "list-series"},
	"listMetricSeriesPost":    {"metric", "list-series-body"},
	"listMetricNames":         {"metric", "list-names"},
	"listMetricNamesPost":     {"metric", "list-names-body"},
	"writeMetrics":            {"metric", "write"},

	// One settings document per account, not a collection.
	"getTraceSettings": {"trace-settings", "get"},
	"putTraceSettings": {"trace-settings", "set"},

	// Multipart upload is a resource with its own lifecycle, but the API
	// never "creates" one — it initiates one — so the create-POST rule cannot
	// see it. Without this, listParts and listMultipartUploads collapse onto
	// the same command.
	"initiateMultipartUpload": {"multipart-upload", "initiate"},
	"completeMultipartUpload": {"multipart-upload", "complete"},
	"abortMultipartUpload":    {"multipart-upload", "abort"},
	"listMultipartUploads":    {"multipart-upload", "list"},
	"listParts":               {"multipart-upload", "list-parts"},
	"uploadPart":              {"multipart-upload", "upload-part"},

	// PUT on an object is an upload, not an update, and it pairs with get.
	"putObject": {"object", "put"},

	// Session actions on the caller rather than CRUD.
	"getOAuthToken":             {"token", "create"},
	"revokeOAuthToken":          {"token", "revoke"},
	"assumeRole":                {"role", "assume"},
	"assumeRoleWithWebIdentity": {"role", "assume-with-web-identity"},
}

// canonicalVerbs maps an operation id's leading word onto the CLI's verb.
// Words absent from it are kept as written, which is what preserves the
// domain's own vocabulary: failover, reinstall, rotate.
var canonicalVerbs = map[string]string{
	"list": "list", "get": "get", "describe": "get",
	"create": "create", "update": "update", "patch": "update",
	"put": "set", "delete": "delete", "search": "search",
}

// resolver holds the per-service pre-pass: which collections the API creates
// members of.
type resolver struct {
	service string
	// creatable is keyed by a normalised collection path, e.g.
	// "/service-accounts/{}/credentials".
	creatable map[string]bool
}

func newResolver(service string, ops []*operation) *resolver {
	r := &resolver{service: service, creatable: map[string]bool{}}
	for _, op := range ops {
		if op.Method != "POST" {
			continue
		}
		lead, _ := splitOperationID(op.ID)
		if lead != "create" {
			continue
		}
		segs := pathSegments(op.Path)
		if _, idx := lastCollection(segs); idx >= 0 {
			r.creatable[normalisePath(segs[:idx+1])] = true
		}
	}
	return r
}

// resolve assigns an operation's CLI noun and verb.
func (r *resolver) resolve(op *operation, xResource string) (resource, verb string) {
	if ov, ok := resourceOverrides[op.ID]; ok {
		return ov.Resource, ov.Verb
	}
	segs := pathSegments(op.Path)
	lead, _ := splitOperationID(op.ID)

	idx := r.resourceCollectionIndex(segs, xResource)
	if idx < 0 {
		return singular(kebab(r.service)), canonicalVerb(lead)
	}

	name := singular(kebab(segs[idx]))
	if x := locateXResource(segs, xResource); x >= 0 && x == idx {
		// The extension names this collection; prefer its wording, which
		// carries context the bare path segment loses — SecurityGroupRule
		// rather than Rule.
		name = kebab(camelWords(xResource))
	}
	name = stripServicePrefix(name, r.service)

	return name, r.verbFor(op, segs, idx, lead)
}

// resourceCollectionIndex picks the collection an operation acts on: the
// deepest one that is a resource in its own right.
func (r *resolver) resourceCollectionIndex(segs []string, xResource string) int {
	if x := locateXResource(segs, xResource); x >= 0 {
		return x
	}
	first := -1
	for i := len(segs) - 1; i >= 0; i-- {
		if isPlaceholder(segs[i]) {
			continue
		}
		first = i
		if r.creatable[normalisePath(segs[:i+1])] {
			return i
		}
		// Keep walking outwards; a collection the API never creates members
		// of belongs to its parent.
	}
	// Nothing creatable: the outermost collection owns the operation.
	for i := 0; i < len(segs); i++ {
		if !isPlaceholder(segs[i]) {
			return i
		}
	}
	return first
}

// verbFor names the action, given which collection owns the operation.
func (r *resolver) verbFor(op *operation, segs []string, idx int, lead string) string {
	v := canonicalVerb(lead)
	tail := segs[idx+1:]
	// Drop the resource's own id placeholder.
	if len(tail) > 0 && isPlaceholder(tail[0]) {
		tail = tail[1:]
	}
	if len(tail) == 0 {
		// A leading word that is not one of the CRUD verbs is the domain's own
		// word for what the operation does, and it beats anything the method
		// implies: ingestLogs is an ingest, not a create, and revokeSTSSession
		// is a revoke rather than a delete.
		if _, isCRUD := canonicalVerbs[lead]; !isCRUD {
			return v
		}
		// Plain CRUD, decided by the method and whether an id was addressed.
		addressed := idx+1 < len(segs) && isPlaceholder(segs[idx+1])
		switch {
		case !addressed && op.Method == "GET":
			return "list"
		case !addressed && op.Method == "POST":
			return "create"
		case addressed && op.Method == "GET":
			return "get"
		case addressed && (op.Method == "PATCH" || op.Method == "PUT"):
			return "update"
		case addressed && op.Method == "DELETE":
			return "delete"
		}
		return v
	}

	// A trailing action segment — /instances/{id}/start — is the verb itself.
	noun := kebab(tail[0])
	if len(tail) == 1 && !r.isCollection(segs, idx, tail[0]) {
		return joinVerb(v, noun)
	}
	if v != "list" {
		noun = singular(noun)
	}
	return joinVerb(v, noun)
}

// isCollection reports whether a trailing segment names a sub-collection
// rather than an action. A collection is addressed by an id somewhere, or is
// the plural the operation lists.
func (r *resolver) isCollection(segs []string, idx int, seg string) bool {
	full := normalisePath(append(append([]string{}, segs[:idx+1]...), seg))
	if r.creatable[full] {
		return true
	}
	return singular(kebab(seg)) != kebab(seg)
}

// locateXResource finds the path segment the extension names, or -1.
//
// Matching on the extension's last word is what lets SecurityGroupRule find
// the segment "rules" while Billing finds nothing in "/v1/invoices" and is
// correctly ignored.
func locateXResource(segs []string, xResource string) int {
	if xResource == "" {
		return -1
	}
	words := splitWords(xResource)
	if len(words) == 0 {
		return -1
	}
	want := words[len(words)-1]
	full := strings.Join(words, "-")
	for i := len(segs) - 1; i >= 0; i-- {
		if isPlaceholder(segs[i]) {
			continue
		}
		s := singular(kebab(segs[i]))
		if s == full || s == want {
			return i
		}
	}
	return -1
}

// stripServicePrefix removes a redundant service name from a resource, so kms
// exposes `key` rather than `kms-key`.
func stripServicePrefix(resource, service string) string {
	p := kebab(service) + "-"
	if strings.HasPrefix(resource, p) && len(resource) > len(p) {
		return strings.TrimPrefix(resource, p)
	}
	return resource
}

func canonicalVerb(lead string) string {
	if v, ok := canonicalVerbs[lead]; ok {
		return v
	}
	return kebab(lead)
}

// joinVerb builds a compound verb without stuttering: the segment
// "cancel-deletion" under the verb "cancel" is already the whole verb.
func joinVerb(verb, noun string) string {
	switch {
	case noun == "" || noun == verb:
		return verb
	case strings.HasPrefix(noun, verb+"-"):
		return noun
	}
	return verb + "-" + noun
}

func pathSegments(p string) []string {
	var out []string
	for _, s := range strings.Split(strings.Trim(p, "/"), "/") {
		if s == "" || s == "v1" {
			continue
		}
		out = append(out, s)
	}
	return out
}

// normalisePath keys a collection independently of what its ids are called.
func normalisePath(segs []string) string {
	parts := make([]string, len(segs))
	for i, s := range segs {
		if isPlaceholder(s) {
			parts[i] = "{}"
			continue
		}
		parts[i] = s
	}
	return "/" + strings.Join(parts, "/")
}

func isPlaceholder(s string) bool { return strings.HasPrefix(s, "{") }

func lastCollection(segs []string) (string, int) {
	for i := len(segs) - 1; i >= 0; i-- {
		if !isPlaceholder(segs[i]) {
			return segs[i], i
		}
	}
	return "", -1
}

func splitOperationID(id string) (lead, rest string) {
	words := splitWords(id)
	if len(words) == 0 {
		return id, ""
	}
	return words[0], strings.Join(words[1:], "-")
}

func kebab(s string) string      { return strings.Join(splitWords(s), "-") }
func camelWords(s string) string { return strings.Join(splitWords(s), "-") }

// singular trims a plural noun, returning it unchanged when it cannot. Unlike
// singularHint in types.go it never invents a suffix: a noun that is already
// singular has to stay as written.
func singular(s string) string {
	// An acronym ending in "s" may be singular (CORS, DNS) or a plural of a
	// shorter one (NICs, IPs, IDs). What separates them is whether trimming
	// the "s" leaves another acronym.
	if _, isAcronym := initialisms[s]; isAcronym {
		if _, trimmedIsAcronym := initialisms[strings.TrimSuffix(s, "s")]; !trimmedIsAcronym {
			return s
		}
	}
	switch {
	case strings.HasSuffix(s, "ies") && len(s) > 3:
		return strings.TrimSuffix(s, "ies") + "y"
	case strings.HasSuffix(s, "sses"), strings.HasSuffix(s, "ches"),
		strings.HasSuffix(s, "shes"), strings.HasSuffix(s, "xes"),
		strings.HasSuffix(s, "zes"):
		return strings.TrimSuffix(s, "es")
	case strings.HasSuffix(s, "ss"), strings.HasSuffix(s, "us"):
		return s
	case strings.HasSuffix(s, "s") && len(s) > 1:
		return strings.TrimSuffix(s, "s")
	}
	return s
}
