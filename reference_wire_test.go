package basaltic_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/compute"
)

// What a generated Get<Resource>ByReference puts on the wire, through a
// real generated client: one request, of the right shape, for each kind.
func TestGetByReferenceOnTheWire(t *testing.T) {
	const id = "0f9c1c8a-8c3e-4c7b-9c2e-1a2b3c4d5e6f"
	const crn = "crn:compute:sa-saopaulo-1:acme:instance/web-01"

	var mu sync.Mutex
	var seen []*url.URL
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		seen = append(seen, r.URL)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/v1/instances/"+id:
			json.NewEncoder(w).Encode(map[string]any{"instance": map[string]any{"id": id, "name": "web-01"}})
		case r.URL.Path == "/v1/instances" && r.URL.Query().Get("name") == "web-01":
			json.NewEncoder(w).Encode(map[string]any{
				"instances": []any{map[string]any{"id": id, "name": "web-01"}},
				"meta":      map[string]any{"total": 1, "limit": 2, "has_more": false},
			})
		case r.URL.Path == "/v1/instances" && r.URL.Query().Get("crn") == crn:
			json.NewEncoder(w).Encode(map[string]any{
				"instances": []any{map[string]any{"id": id, "name": "web-01", "crn": crn}},
				"meta":      map[string]any{"total": 1, "limit": 2, "has_more": false},
			})
		case r.URL.Path == "/v1/instances":
			json.NewEncoder(w).Encode(map[string]any{"instances": []any{}, "meta": map[string]any{"total": 0}})
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":{"code":"INSTANCE_NOT_FOUND","message":"no such instance"}}`))
		}
	}))
	defer srv.Close()

	cfg, err := basaltic.NewConfig(context.Background(),
		basaltic.WithAccessToken("test-token"),
		basaltic.WithServiceEndpoint("compute", srv.URL),
		basaltic.WithRegion("sa-saopaulo-1"),
		basaltic.WithoutRetry(),
	)
	if err != nil {
		t.Fatal(err)
	}
	c := compute.New(cfg)
	ctx := context.Background()

	reset := func() []*url.URL {
		mu.Lock()
		defer mu.Unlock()
		out := seen
		seen = nil
		return out
	}

	t.Run("id", func(t *testing.T) {
		reset()
		inst, err := c.GetInstanceByReference(ctx, id, nil)
		if err != nil || inst.ID != id {
			t.Fatalf("got %+v, %v", inst, err)
		}
		reqs := reset()
		if len(reqs) != 1 || reqs[0].Path != "/v1/instances/"+id {
			t.Fatalf("requests = %v; want one GET of the instance by id", reqs)
		}
	})

	t.Run("name", func(t *testing.T) {
		reset()
		inst, err := c.GetInstanceByReference(ctx, "web-01", nil)
		if err != nil || inst.ID != id {
			t.Fatalf("got %+v, %v", inst, err)
		}
		reqs := reset()
		if len(reqs) != 1 || reqs[0].Path != "/v1/instances" {
			t.Fatalf("requests = %v; want one list", reqs)
		}
		q := reqs[0].Query()
		if q.Get("name") != "web-01" || q.Has("crn") || q.Get("limit") != "2" {
			t.Fatalf("query = %v; want name=web-01&limit=2 and no crn", q)
		}
	})

	t.Run("crn", func(t *testing.T) {
		reset()
		inst, err := c.GetInstanceByReference(ctx, crn, nil)
		if err != nil || inst.ID != id {
			t.Fatalf("got %+v, %v", inst, err)
		}
		q := reset()[0].Query()
		if q.Get("crn") != crn || q.Has("name") {
			t.Fatalf("query = %v; want crn only", q)
		}
	})

	t.Run("scope filters ride along", func(t *testing.T) {
		reset()
		_, err := c.GetInstanceByReference(ctx, "web-01", &compute.ListInstancesParams{Flavor: "m1.small", Limit: 50})
		if err != nil {
			t.Fatal(err)
		}
		q := reset()[0].Query()
		if q.Get("flavor") != "m1.small" || q.Get("name") != "web-01" || q.Get("limit") != "2" {
			t.Fatalf("query = %v; want the scope's flavor, the name, and limit forced to 2", q)
		}
	})

	t.Run("a name miss is not found and is not retried as an id", func(t *testing.T) {
		reset()
		_, err := c.GetInstanceByReference(ctx, "nope", nil)
		if !basaltic.IsNotFound(err) {
			t.Fatalf("error = %v; want not found", err)
		}
		if reqs := reset(); len(reqs) != 1 {
			t.Fatalf("requests = %v; want exactly one", reqs)
		}
	})

	t.Run("a malformed crn sends nothing", func(t *testing.T) {
		reset()
		if _, err := c.GetInstanceByReference(ctx, "crn:compute:oops", nil); err == nil {
			t.Fatal("want an error")
		}
		if reqs := reset(); len(reqs) != 0 {
			t.Fatalf("requests = %v; want none", reqs)
		}
	})
}
