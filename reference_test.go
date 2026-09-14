package basaltic

import (
	"context"
	"errors"
	"strings"
	"testing"
)

// The classification table is the platform's own (services/pkg/reference):
// the SDK must read a string exactly as the server will, or a client could
// send a name where the server sees an id and the two would disagree on what
// the request even asked for.
func TestParseReferenceMatchesThePlatform(t *testing.T) {
	const id = "abcdef01-2345-6789-abcd-ef0123456789"
	for _, tc := range []struct {
		input string
		kind  ReferenceKind
	}{
		{id, ReferenceID}, {strings.ToUpper(id), ReferenceID}, {"00000000-0000-0000-0000-000000000000", ReferenceID},
		{"{" + id + "}", ReferenceName}, {"urn:uuid:" + id, ReferenceName}, {strings.ReplaceAll(id, "-", ""), ReferenceName},
		{"gbcdef01-2345-6789-abcd-ef0123456789", ReferenceName}, {"abcdef0102345-6789-abcd-ef0123456789", ReferenceName},
		{"ubuntu:24.04", ReferenceName}, {"", ReferenceName}, {" " + id, ReferenceName}, {id + "\n", ReferenceName},
		{"CRN:compute:dev:acct:image/id", ReferenceName}, {"crn", ReferenceName},
		{"crn:compute:dev:acct:image/ubuntu", ReferenceCRN}, {"crn:iam:::user/id", ReferenceCRN},
		{"crn:storage:dev:acct:bucket/name/path/key", ReferenceCRN},
		{"crn:network:sa-saopaulo-1:acme:vpc/main/subnet/public", ReferenceCRN},
	} {
		t.Run(tc.input, func(t *testing.T) {
			r, err := ParseReference(tc.input)
			if err != nil || r.Kind != tc.kind || r.Value != tc.input {
				t.Fatalf("ParseReference(%q) = %+v, %v; want kind %s", tc.input, r, err, tc.kind)
			}
		})
	}
}

func TestParseReferenceRejectsMalformedCRNs(t *testing.T) {
	// A "crn:" prefix commits to CRN: these are errors, never names.
	for _, input := range []string{
		"crn:", "crn:compute:dev", "crn:compute:dev:acct",
		"crn:compute:dev:acct:image/id:tag", "crn:compute:dev:acct:image/id:",
	} {
		t.Run(input, func(t *testing.T) {
			r, err := ParseReference(input)
			if !errors.Is(err, ErrInvalidReference) {
				t.Fatalf("error = %v; want ErrInvalidReference", err)
			}
			if r != (Reference{}) {
				t.Fatalf("malformed CRN returned a usable reference: %+v", r)
			}
		})
	}
}

type thing struct{ ID, Name string }

// resolver records what ResolveByReference asked of it.
type resolver struct {
	gets  []string
	lists []struct{ name, crn string }
	page  *Page[thing]
	err   error
}

func (r *resolver) get(_ context.Context, id string) (*thing, error) {
	r.gets = append(r.gets, id)
	return &thing{ID: id}, r.err
}

func (r *resolver) list(_ context.Context, name, crn string) (*Page[thing], error) {
	r.lists = append(r.lists, struct{ name, crn string }{name, crn})
	return r.page, r.err
}

func (r *resolver) resolve(ref string, hasName bool) (*thing, error) {
	return ResolveByReference(context.Background(), ref, "thing", "listThings", hasName, r.get, r.list)
}

func TestResolveByReferenceRoutesEachKindOnce(t *testing.T) {
	const id = "0f9c1c8a-8c3e-4c7b-9c2e-1a2b3c4d5e6f"
	const crn = "crn:compute:sa-saopaulo-1:acme:thing/web-01"

	t.Run("id goes to get and nowhere else", func(t *testing.T) {
		r := &resolver{}
		got, err := r.resolve(id, true)
		if err != nil || got.ID != id {
			t.Fatalf("got %+v, %v", got, err)
		}
		if len(r.gets) != 1 || len(r.lists) != 0 {
			t.Fatalf("gets=%v lists=%v; want one get, no list", r.gets, r.lists)
		}
	})

	t.Run("name goes to list with the name filter", func(t *testing.T) {
		r := &resolver{page: &Page[thing]{Items: []thing{{ID: id, Name: "web-01"}}}}
		got, err := r.resolve("web-01", true)
		if err != nil || got.ID != id {
			t.Fatalf("got %+v, %v", got, err)
		}
		if len(r.gets) != 0 || len(r.lists) != 1 || r.lists[0].name != "web-01" || r.lists[0].crn != "" {
			t.Fatalf("gets=%v lists=%v; want one list by name", r.gets, r.lists)
		}
	})

	t.Run("crn goes to list with the crn filter", func(t *testing.T) {
		r := &resolver{page: &Page[thing]{Items: []thing{{ID: id}}}}
		if _, err := r.resolve(crn, true); err != nil {
			t.Fatal(err)
		}
		if len(r.gets) != 0 || len(r.lists) != 1 || r.lists[0].crn != crn || r.lists[0].name != "" {
			t.Fatalf("gets=%v lists=%v; want one list by crn", r.gets, r.lists)
		}
	})

	t.Run("a miss is not found for that kind, with no second attempt", func(t *testing.T) {
		r := &resolver{page: &Page[thing]{}}
		_, err := r.resolve("web-01", true)
		if !IsNotFound(err) {
			t.Fatalf("error = %v; want not found", err)
		}
		e, _ := AsError(err)
		if e.Code != "THING_NOT_FOUND" || e.OperationID != "listThings" || !strings.Contains(e.Message, `name "web-01"`) {
			t.Fatalf("error = %+v", e)
		}
		if len(r.gets) != 0 || len(r.lists) != 1 {
			t.Fatalf("gets=%v lists=%v; a name miss must not retry as an id", r.gets, r.lists)
		}
	})

	t.Run("more than one match is ambiguous", func(t *testing.T) {
		r := &resolver{page: &Page[thing]{Items: []thing{{ID: "a"}, {ID: "b"}}, Total: 2}}
		_, err := r.resolve("public", true)
		if !IsAmbiguousReference(err) {
			t.Fatalf("error = %v; want ambiguous", err)
		}
		var amb *AmbiguousReferenceError
		errors.As(err, &amb)
		if amb.Resource != "thing" || amb.Reference != "public" || amb.Count != 2 {
			t.Fatalf("ambiguous = %+v", amb)
		}
	})

	t.Run("a page that says more exist is ambiguous too", func(t *testing.T) {
		r := &resolver{page: &Page[thing]{Items: []thing{{ID: "a"}}, HasMore: true}}
		if _, err := r.resolve("public", true); !IsAmbiguousReference(err) {
			t.Fatalf("error = %v; want ambiguous", err)
		}
	})

	t.Run("a nameless resource refuses a name before any request", func(t *testing.T) {
		r := &resolver{}
		_, err := r.resolve("web-01", false)
		if err == nil || !strings.Contains(err.Error(), "has no name") {
			t.Fatalf("error = %v", err)
		}
		if len(r.gets)+len(r.lists) != 0 {
			t.Fatalf("gets=%v lists=%v; want no request", r.gets, r.lists)
		}
	})

	t.Run("a malformed crn is refused before any request", func(t *testing.T) {
		r := &resolver{}
		_, err := r.resolve("crn:compute:thing", true)
		if !errors.Is(err, ErrInvalidReference) {
			t.Fatalf("error = %v", err)
		}
		if len(r.gets)+len(r.lists) != 0 {
			t.Fatalf("gets=%v lists=%v; want no request", r.gets, r.lists)
		}
	})

	t.Run("a list failure is returned as is", func(t *testing.T) {
		boom := errors.New("boom")
		r := &resolver{err: boom}
		if _, err := r.resolve("web-01", true); !errors.Is(err, boom) {
			t.Fatalf("error = %v; want boom", err)
		}
	})
}
