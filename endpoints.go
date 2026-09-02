package basaltic

import (
	"fmt"
	"strings"
)

// defaultDomain is the platform's public domain. Every service endpoint is a
// subdomain of it, so pointing the SDK at another deployment is a matter of
// swapping this one label — see [WithDomain].
const defaultDomain = "basaltic.sh"

// EndpointResolver decides which host an operation is sent to.
//
// service is the service's short name ("compute", "iam"); region is the
// configured region, empty when none was set. A resolver returning an error
// fails the request before it is sent.
//
// Implement one only to reach a deployment whose layout is not
// "<service>.<domain>" — [WithDomain] and [WithServiceEndpoint] cover the
// ordinary cases.
type EndpointResolver interface {
	ResolveEndpoint(service, region string) (string, error)
}

// EndpointResolverFunc adapts a function to [EndpointResolver].
type EndpointResolverFunc func(service, region string) (string, error)

func (f EndpointResolverFunc) ResolveEndpoint(service, region string) (string, error) {
	return f(service, region)
}

// defaultResolver builds endpoints from a service's OpenAPI server template,
// with per-service overrides taking precedence over a whole-domain swap.
type defaultResolver struct {
	domain    string
	overrides map[string]string
	// templates maps a service to the server template its spec declares, so
	// the SDK does not have to carry a second, hand-kept list of which
	// services are regional. Registered by each service package.
	templates map[string]string
}

func (r *defaultResolver) ResolveEndpoint(service, region string) (string, error) {
	if u, ok := r.overrides[service]; ok {
		return strings.TrimRight(u, "/"), nil
	}
	tmpl, ok := r.templates[service]
	if !ok {
		// A service the SDK knows nothing about. Falling back to the
		// conventional layout keeps a newly generated package working before
		// this table catches up, and a wrong guess fails loudly at DNS.
		tmpl = "https://" + service + ".{region}." + defaultDomain
	}
	if r.domain != "" && r.domain != defaultDomain {
		tmpl = strings.Replace(tmpl, defaultDomain, r.domain, 1)
	}
	if strings.Contains(tmpl, "{region}") {
		if region == "" {
			return "", fmt.Errorf("basaltic: %s is a regional service and no region is set — pass WithRegion or set BASALTIC_REGION", service)
		}
		tmpl = strings.ReplaceAll(tmpl, "{region}", region)
	}
	return strings.TrimRight(tmpl, "/"), nil
}

// registeredTemplates holds the server template each generated service
// package declares for itself, keyed by service name. Populated by package
// init in the service packages, so the table only ever describes services the
// program actually imports.
var registeredTemplates = map[string]string{}

// RegisterServiceEndpoint records a service's OpenAPI server template.
//
// Called from generated code at init time. It is exported because the
// generated packages are separate packages, not because callers should use
// it; registering a template by hand will be overwritten by the generated
// registration for the same service.
func RegisterServiceEndpoint(service, template string) {
	registeredTemplates[service] = template
}
