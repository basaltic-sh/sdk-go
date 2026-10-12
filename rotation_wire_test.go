package basaltic_test

import (
	"context"
	"encoding/json"
	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/kms"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestKMSRotationWireCompatibility(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/keys/key/rotation":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			if r.Method != "PUT" || body["enabled"] != false || body["rotation_period_days"] != float64(90) {
				t.Error(body, r.Method)
			}
			w.Write([]byte(`{"rotation":{"enabled":false,"rotation_period_days":90,"next_rotation_at":null}}`))
		case "/v1/keys/key/rotate":
			if r.Header.Get("Idempotency-Key") != "durable-request" {
				t.Error("request token not sent")
			}
			w.WriteHeader(202)
			w.Write([]byte(`{"rotation":{"id":"rotation","key_id":"key","trigger":"on_demand","status":"pending","previous_version":null,"key_version":null,"requested_at":"2026-10-12T12:00:00Z","started_at":null,"completed_at":null,"error_code":null}}`))
		case "/v1/keys/key/verify":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			if body["key_version"] != float64(1) || body["signature"] != "c2ln" {
				t.Error(body)
			}
			w.Write([]byte(`{"signature_valid":true}`))
		default:
			t.Error(r.URL.Path)
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()
	cfg, err := basaltic.NewConfig(context.Background(), basaltic.WithServiceEndpoint("kms", srv.URL), basaltic.WithAccessToken("token"), basaltic.WithRegion("test-1"), basaltic.WithoutRetry())
	if err != nil {
		t.Fatal(err)
	}
	c := kms.New(cfg)
	days, version := 90, 1
	p, err := c.SetKeyRotation(context.Background(), "key", &kms.SetKeyRotationRequest{Enabled: false, RotationPeriodDays: &days})
	if err != nil || p.Enabled || p.NextRotationAt != nil {
		t.Fatal(p, err)
	}
	r, err := c.RotateKeyOnDemand(context.Background(), "key", basaltic.WithIdempotencyKey("durable-request"))
	if err != nil || r.Status != "pending" || r.KeyVersion != nil {
		t.Fatal(r, err)
	}
	v, err := c.Verify(context.Background(), "key", &kms.VerifyRequest{Message: []byte("msg"), Signature: []byte("sig"), KeyVersion: &version})
	if err != nil || !v || calls != 3 {
		t.Fatal(v, err, calls)
	}
}
