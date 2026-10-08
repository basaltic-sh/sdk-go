package basaltic_test

import (
	"context"
	"encoding/json"
	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/compute"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInstanceIAMRoleUpdateWireStates(t *testing.T) {
	role, empty := "worker", ""
	for _, tc := range []struct {
		name    string
		role    *string
		present bool
	}{{"omitted", nil, false}, {"attach", &role, true}, {"detach", &empty, true}} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "PATCH" || r.URL.Path != "/v1/instances/vm" {
					t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
				}
				var body map[string]any
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					t.Error(err)
				}
				got, ok := body["iam_role"]
				if ok != tc.present {
					t.Errorf("iam_role present=%v", ok)
				}
				if tc.present && got != *tc.role {
					t.Errorf("iam_role=%v", got)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"instance":{"id":"vm"}}`))
			}))
			defer server.Close()
			cfg, err := basaltic.NewConfig(context.Background(), basaltic.WithAccessToken("test-token"), basaltic.WithServiceEndpoint("compute", server.URL), basaltic.WithRegion("sa-saopaulo-1"), basaltic.WithoutRetry())
			if err != nil {
				t.Fatal(err)
			}
			if _, err = compute.New(cfg).UpdateInstance(context.Background(), "vm", &compute.InstanceUpdateRequest{IAMRole: tc.role}); err != nil {
				t.Fatal(err)
			}
		})
	}
}
