package basaltic_test

import (
	"context"
	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/catalog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCatalogDiscoveryNeedsNoCredentials(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Header.Get("Authorization") != "" {
			t.Error("public discovery sent bearer credentials")
		}
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/regions":
			if r.URL.Query().Get("crn") != "crn:catalog::platform:region/sa-saopaulo-1" {
				t.Error("missing exact catalog filter")
			}
			w.Write([]byte(`{"regions":[{"code":"sa-saopaulo-1","crn":"crn:catalog::platform:region/sa-saopaulo-1","state":"active","accepting_new_resources":true}],"default":"sa-saopaulo-1"}`))
		case "/v1/regions/sa-saopaulo-1":
			w.Write([]byte(`{"code":"sa-saopaulo-1","state":"active","accepting_new_resources":true}`))
		default:
			t.Errorf("unexpected path %s", r.URL.Path)
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()
	cfg, err := basaltic.NewConfig(context.Background(), basaltic.WithServiceEndpoint("catalog", srv.URL), basaltic.WithAnonymousAccess(), basaltic.WithoutRetry())
	if err != nil {
		t.Fatal(err)
	}
	client := catalog.New(cfg)
	if _, err := cfg.TokenSource.Token(context.Background()); err == nil {
		t.Fatal("anonymous configuration permitted authenticated use")
	}
	rows, err := client.ListRegions(context.Background(), &catalog.ListRegionsParams{CRN: "crn:catalog::platform:region/sa-saopaulo-1"})
	if err != nil {
		t.Fatal(err)
	}
	if rows.Default != "sa-saopaulo-1" || len(rows.Regions) != 1 || !rows.Regions[0].AcceptingNewResources {
		t.Fatal("catalog response lost lifecycle or default")
	}
	detail, err := client.GetRegion(context.Background(), "sa-saopaulo-1")
	if err != nil {
		t.Fatal(err)
	}
	if detail.Code != "sa-saopaulo-1" || !detail.AcceptingNewResources || calls != 2 {
		t.Fatal("invalid catalog detail")
	}
}
