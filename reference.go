package basaltic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// ReferenceKind is how a reference string is read: as a CRN, as a resource id
// or as a name. It is decided by syntax alone, before any request is made.
type ReferenceKind string

const (
	// ReferenceCRN is a "crn:" string with exactly five colon-separated
	// segments, such as crn:compute:sa-saopaulo-1:acme:instance/web-01.
	ReferenceCRN ReferenceKind = "crn"
	// ReferenceID is the canonical 36-character UUID form, 8-4-4-4-12 hex
	// groups. No other spelling of a UUID counts: braces, a urn:uuid:
	// prefix or 32 bare hex digits are names.
	ReferenceID ReferenceKind = "id"
	// ReferenceName is anything else.
	ReferenceName ReferenceKind = "name"
)

// Reference is a reference string and its kind.
//
// The platform accepts a UUID, a CRN or a name wherever a request names
// another resource, and classifies the string by its syntax alone; it never
// tries one kind and then another. [ParseReference] applies the same rule
// here, so the SDK and the platform always agree on what a string is.
type Reference struct {
	Kind  ReferenceKind
	Value string
}

// ErrInvalidReference is returned for a string that starts with "crn:" but
// is not a well-formed CRN. Such a string is never read as a name.
var ErrInvalidReference = errors.New("invalid reference: a crn: value must have five colon-separated segments")

// ParseReference classifies s the way the platform does.
//
// The rule, in order: a "crn:" prefix commits to CRN — with anything other
// than five colon-separated segments it is an error, not a name; the
// canonical 36-character UUID form is an id; everything else, including an
// empty string and every alternate spelling of a UUID, is a name. Nothing is
// trimmed.
func ParseReference(s string) (Reference, error) {
	if strings.HasPrefix(s, "crn:") {
		if strings.Count(s, ":") != 4 {
			return Reference{}, ErrInvalidReference
		}
		return Reference{Kind: ReferenceCRN, Value: s}, nil
	}
	if isCanonicalUUID(s) {
		return Reference{Kind: ReferenceID, Value: s}, nil
	}
	return Reference{Kind: ReferenceName, Value: s}, nil
}

// isCanonicalUUID matches exactly the platform's rule: 36 bytes, hyphens at
// positions 8, 13, 18 and 23, hex digits of either case everywhere else.
func isCanonicalUUID(s string) bool {
	if len(s) != 36 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if s[i] != '-' {
				return false
			}
			continue
		}
		c := s[i]
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F') {
			return false
		}
	}
	return true
}

// AmbiguousReferenceError reports that a name or CRN matched more than one
// resource, so no single one can be returned.
//
// List filters are exact, so this only happens for a name on a resource that
// is unique within a parent — a subnet within a VPC, a snapshot within a
// volume — when the request did not fix the parent. Narrow the scope and try
// again.
type AmbiguousReferenceError struct {
	// Resource is the kind that was looked up, such as "subnet".
	Resource string
	// Reference is the string that was resolved.
	Reference string
	// Count is how many resources matched, when the page said; at least 2.
	Count int
}

func (e *AmbiguousReferenceError) Error() string {
	return fmt.Sprintf("reference %q matches more than one %s; narrow the scope", e.Reference, e.Resource)
}

// IsAmbiguousReference reports whether err says a reference matched more
// than one resource.
func IsAmbiguousReference(err error) bool {
	var e *AmbiguousReferenceError
	return errors.As(err, &e)
}

// referenceNotFound builds the not-found error a miss on a name or CRN
// returns, shaped like the platform's own so that [IsNotFound] recognises it.
func referenceNotFound(operationID, resource string, ref Reference) *Error {
	return &Error{
		StatusCode:  http.StatusNotFound,
		Code:        strings.ToUpper(strings.ReplaceAll(resource, "-", "_")) + "_NOT_FOUND",
		Message:     fmt.Sprintf("no %s matches %s %q", resource, ref.Kind, ref.Value),
		OperationID: operationID,
	}
}

// ResolveByReference fetches one resource by id, CRN or name.
//
// It is the engine behind every generated Get<Resource>ByReference method:
// an id goes to getByID; a CRN or a name goes to list with that single exact
// filter. A miss is a not-found error for the kind the string was read as —
// there is no second attempt under another kind — and more than one match is
// an [AmbiguousReferenceError]. hasName is false for resources that carry no
// name, where a name reference is refused before any request.
//
// listByFilter receives exactly one of name and crn set. It should ask for a
// page of at least two items, so that ambiguity is visible without walking
// the whole collection.
func ResolveByReference[T any](
	ctx context.Context,
	ref string,
	resource, listOperationID string,
	hasName bool,
	getByID func(ctx context.Context, id string) (*T, error),
	listByFilter func(ctx context.Context, name, crn string) (*Page[T], error),
) (*T, error) {
	r, err := ParseReference(ref)
	if err != nil {
		return nil, err
	}
	var page *Page[T]
	switch r.Kind {
	case ReferenceID:
		return getByID(ctx, r.Value)
	case ReferenceCRN:
		page, err = listByFilter(ctx, "", r.Value)
	default:
		if !hasName {
			return nil, fmt.Errorf("%s has no name: %q must be an id or a CRN", resource, ref)
		}
		page, err = listByFilter(ctx, r.Value, "")
	}
	if err != nil {
		return nil, err
	}
	switch {
	case page == nil || len(page.Items) == 0:
		return nil, referenceNotFound(listOperationID, resource, r)
	case len(page.Items) > 1 || page.HasMore:
		count := len(page.Items)
		if page.Total > count {
			count = page.Total
		}
		return nil, &AmbiguousReferenceError{Resource: resource, Reference: ref, Count: count}
	}
	return &page.Items[0], nil
}
