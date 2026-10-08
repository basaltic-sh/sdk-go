package main

import (
	"fmt"
	"regexp"
	"strings"
)

// A by-reference getter pairs a resource's get with its list.
//
// The platform names other resources by id, CRN or name and classifies the
// string by syntax; every list takes exact `name` and `crn` filters so that a
// client can fetch one resource by any of the three without a lookup endpoint
// of its own. The pairing is derived here rather than declared, from the same
// resource/verb placement the command tree uses: a `get` whose result type is
// what the resource's `list` pages over, under the same parent path, gets a
// Get<Resource>ByReference method beside it.
type byReference struct {
	Get  *operation
	List *operation
	// HasName is false for resources that carry no name, where only an id
	// or a CRN can address one.
	HasName bool
	// Name, CRN and Limit are the list's parameter fields, by Go name.
	Name, CRN, Limit string
	// Scope are the list's other reference filters — the parent a name is
	// unique within, such as a subnet's vpc — which the getter passes
	// through and the CLI exposes as flags.
	Scope []*param
}

// GoName is the getter's method name: GetInstance becomes
// GetInstanceByReference.
func (r *byReference) GoName() string { return r.Get.GoName + "ByReference" }

// scopeDoc recognises a list filter that itself takes a reference, from the
// specification's own wording.
var scopeDoc = regexp.MustCompile(`(?i)\breference\b|\bCRN\b`)

// deriveByReference finds every get/list pair once the operations are placed.
func (b *builder) deriveByReference() {
	lists := map[string][]*operation{}
	for _, op := range b.ops {
		if op.Verb == "list" && op.Result.Kind == resultPage {
			lists[op.Resource] = append(lists[op.Resource], op)
		}
	}
	for _, get := range b.ops {
		if get.Verb != "get" || get.Result.Kind != resultValue || len(get.PathParams) == 0 {
			continue
		}
		item := strings.TrimPrefix(get.Result.Type, "*")
		if item == get.Result.Type {
			continue
		}
		for _, list := range lists[get.Resource] {
			if list.Result.ItemType != item || !sameParents(get.PathParams[:len(get.PathParams)-1], list.PathParams) {
				continue
			}
			r := &byReference{Get: get, List: list}
			for _, p := range list.QueryParams {
				switch {
				case p.WireName == "name" && p.Type == "string":
					r.Name = p.Name
					r.HasName = true
				case p.WireName == "crn" && p.Type == "string":
					r.CRN = p.Name
				case p.WireName == "limit" && (p.Type == "int" || p.Type == "int64"):
					r.Limit = p.Name
				case p.Type == "string" && scopeDoc.MatchString(p.RawDoc):
					r.Scope = append(r.Scope, p)
				}
			}
			if r.CRN == "" {
				// A list without a crn filter cannot serve a lookup; the
				// platform's contract gives every list one, so this only
				// happens for a collection that is not a resource.
				continue
			}
			b.byReference = append(b.byReference, r)
			break
		}
	}
}

func sameParents(a, c []*param) bool {
	if len(a) != len(c) {
		return false
	}
	for i := range a {
		if a[i].WireName != c[i].WireName {
			return false
		}
	}
	return true
}

// emitByReference writes one Get<Resource>ByReference method.
func (b *builder) emitByReference(body *strings.Builder, r *byReference) {
	get, list := r.Get, r.List
	item := list.Result.ItemType

	args := []string{"ctx context.Context"}
	parents := []string{}
	for _, p := range list.PathParams {
		args = append(args, p.Name+" string")
		parents = append(parents, p.Name)
	}
	args = append(args, "ref string", "scope *"+list.ParamsType, "opts ...basaltic.RequestOption")

	kinds := "an id, a CRN or a name"
	if !r.HasName {
		kinds = "an id or a CRN"
	}
	extra := []string{wrapText(fmt.Sprintf(
		"The reference is classified by its syntax alone, exactly as the platform "+
			"does (see [basaltic.ParseReference]): an id is fetched with [Client.%s]; "+
			"a CRN or a name goes to [Client.%s] as an exact filter, together with "+
			"any filters already set on scope, which may be nil. A miss is a not-found "+
			"error for the kind the string was read as — no other kind is tried — and "+
			"more than one match is a [basaltic.AmbiguousReferenceError].",
		get.GoName, list.GoName), 70)}
	if !r.HasName {
		extra = append(extra, "This resource has no name; a name reference is refused before any request.")
	}
	if len(r.Scope) > 0 {
		names := make([]string, 0, len(r.Scope))
		for _, p := range r.Scope {
			names = append(names, p.Name)
		}
		extra = append(extra, wrapText(fmt.Sprintf(
			"A name is unique only within its parent; fix it on scope (%s) or the lookup "+
				"can match more than one.", strings.Join(names, ", ")), 70))
	}
	doc := goDoc(r.GoName(), wrapText(fmt.Sprintf(
		"%s fetches one %s by %s.", r.GoName(), strings.ReplaceAll(get.Resource, "-", " "), kinds), 70), extra...)
	body.WriteString(commentLines("", doc))

	fmt.Fprintf(body, "func (c *Client) %s(%s) (*%s, error) {\n", r.GoName(), strings.Join(args, ", "), item)
	fmt.Fprintf(body, "\treturn basaltic.ResolveByReference(ctx, ref, %q, %q, %t,\n", get.Resource, list.ID, r.HasName)

	// The closure parameters carry a ref prefix so that a parent path
	// parameter called id, name or crn is not shadowed.
	getArgs := append([]string{"ctx"}, parents...)
	getArgs = append(getArgs, "refID", "opts...")
	fmt.Fprintf(body, "\t\tfunc(ctx context.Context, refID string) (*%s, error) {\n\t\t\treturn c.%s(%s)\n\t\t},\n",
		item, get.GoName, strings.Join(getArgs, ", "))

	listArgs := append([]string{"ctx"}, parents...)
	listArgs = append(listArgs, "&p", "opts...")
	fmt.Fprintf(body, "\t\tfunc(ctx context.Context, refName, refCRN string) (*basaltic.Page[%s], error) {\n", item)
	fmt.Fprintf(body, "\t\t\tvar p %s\n\t\t\tif scope != nil {\n\t\t\t\tp = *scope\n\t\t\t}\n", list.ParamsType)
	if r.HasName {
		fmt.Fprintf(body, "\t\t\tp.%s = refName\n", r.Name)
	} else {
		body.WriteString("\t\t\t_ = refName\n")
	}
	fmt.Fprintf(body, "\t\t\tp.%s = refCRN\n", r.CRN)
	if r.Limit != "" {
		// Two is enough to tell one match from many without paging.
		fmt.Fprintf(body, "\t\t\tp.%s = 2\n", r.Limit)
	}
	fmt.Fprintf(body, "\t\t\treturn c.%s(%s)\n\t\t})\n}\n\n", list.GoName, strings.Join(listArgs, ", "))
}
