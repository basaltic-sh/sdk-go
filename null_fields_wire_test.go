package basaltic_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/telemetry"
)

func TestNullablePatchFieldsOnWire(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value *int
		nulls []string
		want  map[string]any
	}{
		{name: "omitted", want: map[string]any{}},
		{name: "value", value: basaltic.Int(7), want: map[string]any{"retention_days": float64(7)}},
		{name: "clear", nulls: []string{"retention_days", "kms_key"}, want: map[string]any{"retention_days": nil, "kms_key": nil}},
		{name: "null overrides value", value: basaltic.Int(7), nulls: []string{"retention_days"}, want: map[string]any{"retention_days": nil}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var got map[string]any
				if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
					t.Error(err)
				}
				if r.Method != "PATCH" || !reflect.DeepEqual(got, tc.want) {
					t.Errorf("%s body = %#v; want %#v", r.Method, got, tc.want)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"log_group":{"id":"group"}}`))
			}))
			defer server.Close()
			cfg, err := basaltic.NewConfig(context.Background(), basaltic.WithAccessToken("test-token"), basaltic.WithServiceEndpoint("telemetry", server.URL), basaltic.WithoutRetry())
			if err != nil {
				t.Fatal(err)
			}
			body := &telemetry.UpdateLogGroupRequest{RetentionDays: tc.value}
			if _, err := telemetry.New(cfg).UpdateLogGroup(context.Background(), "group", body, basaltic.WithNullFields(tc.nulls...)); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(body.RetentionDays, tc.value) {
				t.Fatal("request mutated")
			}
		})
	}
}

func TestNullFieldsRequireJSONObject(t *testing.T) {
	cfg, err := basaltic.NewConfig(context.Background(), basaltic.WithAccessToken("test-token"), basaltic.WithServiceEndpoint("telemetry", "http://invalid.test"))
	if err != nil {
		t.Fatal(err)
	}
	for _, body := range []any{nil, []int{1}, "raw"} {
		op := &basaltic.Operation{ID: "test", Method: "PATCH", Path: "/v1/test", Body: body}
		if err := basaltic.NewClient(cfg, "telemetry").Do(context.Background(), op, nil, basaltic.WithNullFields("field")); err == nil {
			t.Fatal("accepted non-object body")
		}
	}
}
