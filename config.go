package basaltic

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// Environment variables consulted by [NewConfig].
const (
	// EnvAccessKeyID and EnvSecretAccessKey are a service account's key pair.
	// Present together, they configure client-credentials authentication.
	EnvAccessKeyID     = "BASALTIC_ACCESS_KEY_ID"
	EnvSecretAccessKey = "BASALTIC_SECRET_ACCESS_KEY"
	// EnvAccessToken supplies a bearer token directly, skipping the exchange.
	// It takes precedence over the key pair. Useful where something upstream
	// already holds a token and the SDK should not mint its own.
	EnvAccessToken = "BASALTIC_ACCESS_TOKEN"
	// EnvRegion is the region regional services are addressed in.
	EnvRegion = "BASALTIC_REGION"
	// EnvAccountID selects the account requests act on.
	EnvAccountID = "BASALTIC_ACCOUNT_ID"
	// EnvDomain swaps the domain every service endpoint is built from —
	// "cloud.example.dev" instead of the platform default.
	EnvDomain = "BASALTIC_DOMAIN"
	// EnvEndpointPrefix prefixes a per-service endpoint override:
	// BASALTIC_ENDPOINT_URL_COMPUTE points just the compute client somewhere
	// else and leaves the rest alone.
	EnvEndpointPrefix = "BASALTIC_ENDPOINT_URL_"
)

// defaultHTTPTimeout bounds an ordinary API call. Streaming operations —
// object bodies, image uploads — are exempt; see [Client.DoStream].
const defaultHTTPTimeout = 30 * time.Second

// Config is the shared setup every service client is built from: how to
// authenticate, which region and account to act in, and how to reach the
// platform.
//
// Build one with [NewConfig] and share it. Service clients hold a reference
// rather than a copy, so a single Config means a single token cache serving
// every client, and one exchange rather than one per service.
//
// A Config is safe for concurrent use once built. Do not mutate its fields
// after handing it to a client; use the options instead.
type Config struct {
	// TokenSource supplies the bearer token for each request.
	TokenSource TokenSource
	// Region is the region regional services are addressed in. Global
	// services ignore it.
	Region string
	// AccountID is sent as X-Account-Id, selecting the account a request acts
	// on. Empty omits the header, which makes the platform use the
	// credential's own account.
	AccountID string
	// HTTPClient performs every request. Replace it to install a custom
	// transport, a proxy, or instrumentation.
	HTTPClient *http.Client
	// EndpointResolver decides which host each service is reached at.
	EndpointResolver EndpointResolver
	// Retry governs how failed requests are retried.
	Retry RetryConfig
	// UserAgentExtra is appended to the SDK's User-Agent. Set it to identify
	// the calling application in the platform's logs.
	UserAgentExtra string

	// baseHeaders are stamped on every request, set by WithHeader.
	baseHeaders http.Header
	// requestOptions apply to every request made through this config.
	requestOptions []RequestOption
}

// Option configures a [Config].
type Option func(*configBuilder) error

type configBuilder struct {
	cfg       *Config
	domain    string
	overrides map[string]string
	// credentials held until the token endpoint is known, since the source
	// needs an endpoint and the endpoint needs the resolver to be settled.
	accessKeyID     string
	secretAccessKey string
	tokenDuration   time.Duration
	staticToken     string
	// explicitTokenURL overrides the derived IAM token endpoint.
	tokenURL string
}

// NewConfig builds a [Config] from the environment, then applies opts.
//
// Options win over the environment, so a program can take its region from
// BASALTIC_REGION and still override it for one call site. With neither an
// option nor an environment variable supplying credentials, NewConfig fails —
// an unauthenticated client would only fail later, one request at a time.
//
// The environment variables are [EnvAccessKeyID], [EnvSecretAccessKey],
// [EnvAccessToken], [EnvRegion], [EnvAccountID], [EnvDomain] and any
// [EnvEndpointPrefix] override.
func NewConfig(ctx context.Context, opts ...Option) (*Config, error) {
	b := &configBuilder{
		cfg: &Config{
			Region:      os.Getenv(EnvRegion),
			AccountID:   os.Getenv(EnvAccountID),
			Retry:       DefaultRetryConfig(),
			baseHeaders: http.Header{},
		},
		domain:          cmpOr(os.Getenv(EnvDomain), defaultDomain),
		overrides:       endpointOverridesFromEnv(),
		accessKeyID:     os.Getenv(EnvAccessKeyID),
		secretAccessKey: os.Getenv(EnvSecretAccessKey),
		staticToken:     os.Getenv(EnvAccessToken),
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(b); err != nil {
			return nil, err
		}
	}
	return b.build()
}

func (b *configBuilder) build() (*Config, error) {
	cfg := b.cfg
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: defaultHTTPTimeout}
	}
	if cfg.EndpointResolver == nil {
		cfg.EndpointResolver = &defaultResolver{
			domain:    b.domain,
			overrides: b.overrides,
			templates: registeredTemplates,
		}
	}
	if cfg.TokenSource == nil {
		switch {
		case b.staticToken != "":
			cfg.TokenSource = StaticTokenSource(b.staticToken)
		case b.accessKeyID != "" && b.secretAccessKey != "":
			tokenURL := b.tokenURL
			if tokenURL == "" {
				iam, err := cfg.EndpointResolver.ResolveEndpoint("iam", cfg.Region)
				if err != nil {
					return nil, fmt.Errorf("basaltic: resolving the IAM token endpoint: %w", err)
				}
				tokenURL = iam + "/v1/oauth/token"
			}
			cfg.TokenSource = &ClientCredentialsSource{
				AccessKeyID:     b.accessKeyID,
				SecretAccessKey: b.secretAccessKey,
				TokenURL:        tokenURL,
				Duration:        b.tokenDuration,
				HTTPClient:      cfg.HTTPClient,
			}
		default:
			return nil, fmt.Errorf(
				"basaltic: no credentials — pass WithClientCredentials or WithAccessToken, or set %s and %s",
				EnvAccessKeyID, EnvSecretAccessKey)
		}
	}
	return cfg, nil
}

// WithClientCredentials authenticates as a service account, exchanging its
// access key pair for bearer tokens and refreshing them before expiry.
//
// This is the ordinary choice. The same key pair is also the AWS SigV4
// credential for the S3-compatible object endpoint, which this SDK does not
// wrap — use an AWS SDK there with these same values.
func WithClientCredentials(accessKeyID, secretAccessKey string) Option {
	return func(b *configBuilder) error {
		if accessKeyID == "" || secretAccessKey == "" {
			return fmt.Errorf("basaltic: WithClientCredentials needs both an access key id and a secret")
		}
		b.accessKeyID, b.secretAccessKey = accessKeyID, secretAccessKey
		b.staticToken = ""
		return nil
	}
}

// WithTokenDuration requests a token lifetime from the token endpoint.
//
// The platform serves 15 minutes to 12 hours and clamps values outside that
// range rather than refusing them, so asking for a day yields the longest
// token allowed. Zero takes the platform default of one hour. Only meaningful
// alongside [WithClientCredentials].
func WithTokenDuration(d time.Duration) Option {
	return func(b *configBuilder) error {
		b.tokenDuration = d
		return nil
	}
}

// WithAccessToken presents a bearer token the caller already holds.
//
// The SDK cannot refresh it, so a process that outlives the token starts
// failing authentication. Prefer [WithClientCredentials] wherever the key
// pair is available.
func WithAccessToken(token string) Option {
	return func(b *configBuilder) error {
		if token == "" {
			return fmt.Errorf("basaltic: WithAccessToken needs a token")
		}
		b.staticToken = token
		b.accessKeyID, b.secretAccessKey = "", ""
		return nil
	}
}

// WithTokenSource authenticates from an arbitrary source, for tokens the SDK
// cannot mint itself — an assumed-role session, or a sidecar.
func WithTokenSource(ts TokenSource) Option {
	return func(b *configBuilder) error {
		if ts == nil {
			return fmt.Errorf("basaltic: WithTokenSource needs a source")
		}
		b.cfg.TokenSource = ts
		return nil
	}
}

// WithRegion sets the region regional services are addressed in, such as
// "sa-saopaulo-1". Global services — iam, dns, billing, audit, quota —
// ignore it.
func WithRegion(region string) Option {
	return func(b *configBuilder) error {
		b.cfg.Region = region
		return nil
	}
}

// WithAccountID selects the account requests act on, sent as X-Account-Id.
//
// Omit it to act on the credential's own account. Set it to reach another
// account the credential is authorized for — without it those resources
// answer 404 rather than 403, since an account you are not acting in is one
// whose resources you cannot see.
func WithAccountID(accountID string) Option {
	return func(b *configBuilder) error {
		b.cfg.AccountID = accountID
		return nil
	}
}

// WithHTTPClient installs the client used for every request, including token
// exchanges. Use it for a custom transport, a proxy, or instrumentation.
//
// The client's Timeout bounds ordinary API calls. Streaming operations get
// their own client without a total-request deadline, because a multi-gigabyte
// body cannot fit a fixed one.
func WithHTTPClient(hc *http.Client) Option {
	return func(b *configBuilder) error {
		if hc == nil {
			return fmt.Errorf("basaltic: WithHTTPClient needs a client")
		}
		b.cfg.HTTPClient = hc
		return nil
	}
}

// WithDomain builds every service endpoint from another domain, keeping the
// per-service and per-region labels: compute in region sa-saopaulo-1 becomes
// "https://compute.sa-saopaulo-1.<domain>".
//
// This is how to point the whole SDK at a non-production deployment.
func WithDomain(domain string) Option {
	return func(b *configBuilder) error {
		if domain == "" {
			return fmt.Errorf("basaltic: WithDomain needs a domain")
		}
		b.domain = strings.TrimSuffix(strings.TrimPrefix(domain, "."), "/")
		return nil
	}
}

// WithServiceEndpoint points one service at an explicit URL and leaves the
// rest alone — useful for testing one client against a local stub.
//
// service is the short name the SDK uses: "compute", "iam", "network".
func WithServiceEndpoint(service, endpoint string) Option {
	return func(b *configBuilder) error {
		if service == "" || endpoint == "" {
			return fmt.Errorf("basaltic: WithServiceEndpoint needs a service and an endpoint")
		}
		if b.overrides == nil {
			b.overrides = map[string]string{}
		}
		b.overrides[strings.ToLower(service)] = endpoint
		return nil
	}
}

// WithEndpointResolver takes over endpoint resolution entirely, for a
// deployment whose layout the built-in options cannot describe.
func WithEndpointResolver(r EndpointResolver) Option {
	return func(b *configBuilder) error {
		if r == nil {
			return fmt.Errorf("basaltic: WithEndpointResolver needs a resolver")
		}
		b.cfg.EndpointResolver = r
		return nil
	}
}

// WithTokenEndpoint overrides where client credentials are exchanged. By
// default this is the IAM service's /v1/oauth/token, resolved like any other
// endpoint.
func WithTokenEndpoint(url string) Option {
	return func(b *configBuilder) error {
		b.tokenURL = url
		return nil
	}
}

// WithRetry replaces the retry policy. See [RetryConfig] and
// [DefaultRetryConfig].
func WithRetry(rc RetryConfig) Option {
	return func(b *configBuilder) error {
		b.cfg.Retry = rc
		return nil
	}
}

// WithoutRetry disables retries. Every request is sent once and its outcome
// returned as-is.
func WithoutRetry() Option {
	return func(b *configBuilder) error {
		b.cfg.Retry.MaxAttempts = 1
		return nil
	}
}

// WithUserAgent appends a token to the SDK's User-Agent, identifying the
// calling application in the platform's logs. Conventionally "name/version".
func WithUserAgent(extra string) Option {
	return func(b *configBuilder) error {
		b.cfg.UserAgentExtra = extra
		return nil
	}
}

// WithHeader stamps a header on every request this config makes. Setting a
// header the SDK manages — Authorization, X-Account-Id — replaces the SDK's
// value, which is rarely what you want.
func WithHeader(name, value string) Option {
	return func(b *configBuilder) error {
		if b.cfg.baseHeaders == nil {
			b.cfg.baseHeaders = http.Header{}
		}
		b.cfg.baseHeaders.Set(name, value)
		return nil
	}
}

// WithRequestOptions applies the given per-request options to every request
// made through this config. Options passed at the call site are applied after
// these and win.
func WithRequestOptions(opts ...RequestOption) Option {
	return func(b *configBuilder) error {
		b.cfg.requestOptions = append(b.cfg.requestOptions, opts...)
		return nil
	}
}

// endpointOverridesFromEnv collects BASALTIC_ENDPOINT_URL_<SERVICE> values.
func endpointOverridesFromEnv() map[string]string {
	out := map[string]string{}
	for _, kv := range os.Environ() {
		name, value, ok := strings.Cut(kv, "=")
		if !ok || value == "" || !strings.HasPrefix(name, EnvEndpointPrefix) {
			continue
		}
		service := strings.ToLower(strings.TrimPrefix(name, EnvEndpointPrefix))
		if service == "" {
			continue
		}
		out[service] = value
	}
	return out
}

// cmpOr returns the first non-empty argument. (cmp.Or, without requiring the
// Go version that introduced it.)
func cmpOr(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
