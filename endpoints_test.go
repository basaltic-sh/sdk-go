package basaltic

import (
	"context"
	"strings"
	"testing"
)

// The endpoint templates come from the service specifications at init time,
// so a test in this package has to register the ones it exercises.
func registerTestTemplates(t *testing.T) {
	t.Helper()
	RegisterServiceEndpoint("compute", "https://compute.{region}.basaltic.sh")
	RegisterServiceEndpoint("iam", "https://iam.basaltic.sh")
}

func TestResolveEndpoint(t *testing.T) {
	registerTestTemplates(t)

	tests := []struct {
		name    string
		opts    []Option
		service string
		want    string
		wantErr string
	}{
		{
			name:    "regional service substitutes the region",
			opts:    []Option{WithRegion("sa-saopaulo-1")},
			service: "compute",
			want:    "https://compute.sa-saopaulo-1.basaltic.sh",
		},
		{
			name:    "global service ignores the region",
			opts:    []Option{WithRegion("sa-saopaulo-1")},
			service: "iam",
			want:    "https://iam.basaltic.sh",
		},
		{
			// A regional call with no region would otherwise be sent to a
			// host with a literal {region} in it.
			name:    "regional service without a region is refused",
			service: "compute",
			wantErr: "regional service and no region is set",
		},
		{
			name:    "global service works without a region",
			service: "iam",
			want:    "https://iam.basaltic.sh",
		},
		{
			name:    "domain swap keeps the service and region labels",
			opts:    []Option{WithRegion("sa-saopaulo-1"), WithDomain("cloud.example.dev")},
			service: "compute",
			want:    "https://compute.sa-saopaulo-1.cloud.example.dev",
		},
		{
			name:    "domain swap applies to global services too",
			opts:    []Option{WithDomain("cloud.example.dev")},
			service: "iam",
			want:    "https://iam.cloud.example.dev",
		},
		{
			name: "per-service override wins over the domain",
			opts: []Option{
				WithRegion("sa-saopaulo-1"),
				WithDomain("cloud.example.dev"),
				WithServiceEndpoint("compute", "http://127.0.0.1:8080/"),
			},
			service: "compute",
			want:    "http://127.0.0.1:8080",
		},
		{
			name: "an override on one service leaves the others alone",
			opts: []Option{
				WithServiceEndpoint("compute", "http://127.0.0.1:8080"),
			},
			service: "iam",
			want:    "https://iam.basaltic.sh",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := append([]Option{WithAccessToken("t")}, tc.opts...)
			cfg, err := NewConfig(context.Background(), opts...)
			if err != nil {
				t.Fatalf("NewConfig: %v", err)
			}
			got, err := cfg.EndpointResolver.ResolveEndpoint(tc.service, cfg.Region)
			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("ResolveEndpoint() error = %v, want it to contain %q", err, tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("ResolveEndpoint: %v", err)
			}
			if got != tc.want {
				t.Errorf("ResolveEndpoint() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestEndpointOverrideFromEnvironment(t *testing.T) {
	registerTestTemplates(t)
	t.Setenv(EnvAccessToken, "t")
	t.Setenv(EnvEndpointPrefix+"COMPUTE", "http://localhost:9999")

	cfg, err := NewConfig(context.Background(), WithRegion("sa-saopaulo-1"))
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	got, err := cfg.EndpointResolver.ResolveEndpoint("compute", cfg.Region)
	if err != nil {
		t.Fatalf("ResolveEndpoint: %v", err)
	}
	if want := "http://localhost:9999"; got != want {
		t.Errorf("ResolveEndpoint() = %q, want %q", got, want)
	}
}

func TestCustomEndpointResolver(t *testing.T) {
	cfg, err := NewConfig(context.Background(),
		WithAccessToken("t"),
		WithEndpointResolver(EndpointResolverFunc(func(service, region string) (string, error) {
			return "https://" + region + "-" + service + ".internal", nil
		})),
		WithRegion("sa-saopaulo-1"),
	)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	got, _ := cfg.EndpointResolver.ResolveEndpoint("compute", cfg.Region)
	if want := "https://sa-saopaulo-1-compute.internal"; got != want {
		t.Errorf("ResolveEndpoint() = %q, want %q", got, want)
	}
}

func TestNewIdempotencyKeyLooksLikeAUUID(t *testing.T) {
	seen := map[string]bool{}
	for i := 0; i < 100; i++ {
		k := NewIdempotencyKey()
		if len(k) != 36 {
			t.Fatalf("key %q has length %d, want 36", k, len(k))
		}
		if k[8] != '-' || k[13] != '-' || k[18] != '-' || k[23] != '-' {
			t.Fatalf("key %q is not in UUID layout", k)
		}
		if seen[k] {
			t.Fatalf("key %q was generated twice", k)
		}
		seen[k] = true
	}
}
