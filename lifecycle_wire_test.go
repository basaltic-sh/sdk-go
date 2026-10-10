package basaltic_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/storage"
)

func TestLifecycleOptionalDatesOnWire(t *testing.T) {
	for _, rule := range []string{
		`{"status":"enabled","expiration":{"days":30},"transition":{"days":7,"storage_class":"COLD"}}`,
		`{"status":"enabled","expiration":{"date":"2030-01-01T00:00:00Z"},"transition":{"date":"2029-01-01T00:00:00Z","storage_class":"COLD"}}`,
		`{"status":"enabled","expiration":{"days":30},"transition":{"days":0,"storage_class":"COLD"}}`,
	} {
		t.Run(rule, func(t *testing.T) {
			input := `{"lifecycle":{"rules":[` + rule + `]}}`
			var body storage.PutBucketLifecycleRequest
			if err := json.Unmarshal([]byte(input), &body); err != nil {
				t.Fatal(err)
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var got, want any
				if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
					t.Error(err)
				}
				if err := json.Unmarshal([]byte(input), &want); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("wire body = %#v; want %#v", got, want)
				}
				w.WriteHeader(http.StatusNoContent)
			}))
			defer server.Close()
			cfg, err := basaltic.NewConfig(context.Background(), basaltic.WithAccessToken("test-token"), basaltic.WithServiceEndpoint("storage", server.URL), basaltic.WithRegion("sa-saopaulo-1"), basaltic.WithoutRetry())
			if err != nil {
				t.Fatal(err)
			}
			if err = storage.New(cfg).PutBucketLifecycle(context.Background(), "test", &body); err != nil {
				t.Fatal(err)
			}
		})
	}
}
