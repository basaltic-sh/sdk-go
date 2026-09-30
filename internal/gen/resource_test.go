package main

import (
	"strings"
	"testing"
)

func TestGatewayRoutesHaveDistinctCommands(t *testing.T) {
	ops := []*operation{
		{ID: "listEgressOnlyGatewayRoutes", Method: "GET", Path: "/v1/egress-only-gateways/{id}/routes", xResource: "Route"},
		{ID: "listInternetGatewayRoutes", Method: "GET", Path: "/v1/internet-gateways/{id}/routes", xResource: "Route"},
		{ID: "listNATGatewayRoutes", Method: "GET", Path: "/v1/nat-gateways/{id}/routes", xResource: "Route"},
		{ID: "listRoutes", Method: "GET", Path: "/v1/route-tables/{id}/routes", xResource: "Route"},
	}
	want := []string{"egress-only-gateway list-routes", "internet-gateway list-routes", "nat-gateway list-routes", "route list"}
	r := newResolver("network", ops)
	for i, op := range ops {
		op.Resource, op.Verb = r.resolve(op, op.xResource)
		if got := op.Resource + " " + op.Verb; got != want[i] {
			t.Errorf("%s: got %q, want %q", op.ID, got, want[i])
		}
	}
	if err := validateCommandNames("network", ops); err != nil {
		t.Fatal(err)
	}
	ops[0].Resource, ops[0].Verb = "route", "list"
	if err := validateCommandNames("network", ops); err == nil || !strings.Contains(err.Error(), "listEgressOnlyGatewayRoutes and listRoutes") {
		t.Fatalf("expected actionable collision error, got %v", err)
	}
}

func TestSingular(t *testing.T) {
	tests := []struct{ in, want string }{
		{"instances", "instance"},
		{"databases", "database"},
		{"policies", "policy"},
		{"addresses", "address"},
		{"boxes", "box"},
		{"nics", "nic"},
		{"volumes", "volume"},
		// An acronym is one word, not a plural.
		{"cors", "cors"},
		{"dns", "dns"},
		// Already singular.
		{"console", "console"},
		{"start", "start"},
		{"refresh", "refresh"},
		{"status", "status"},
	}
	for _, tc := range tests {
		if got := singular(tc.in); got != tc.want {
			t.Errorf("singular(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestJoinVerbDoesNotStutter(t *testing.T) {
	tests := []struct{ verb, noun, want string }{
		{"cancel", "cancel-deletion", "cancel-deletion"},
		{"convert", "convert-to-ha", "convert-to-ha"},
		{"rotate", "rotate-password", "rotate-password"},
		{"set", "encryption", "set-encryption"},
		{"list", "volumes", "list-volumes"},
		{"attach", "volume", "attach-volume"},
		{"refresh", "refresh", "refresh"},
	}
	for _, tc := range tests {
		if got := joinVerb(tc.verb, tc.noun); got != tc.want {
			t.Errorf("joinVerb(%q, %q) = %q, want %q", tc.verb, tc.noun, got, tc.want)
		}
	}
}

// locateXResource is what stops a service-wide tag collapsing several
// resources into one: every billing operation is tagged "Billing", which
// matches no segment of /v1/invoices and is therefore ignored.
func TestLocateXResource(t *testing.T) {
	tests := []struct {
		name      string
		segs      []string
		xResource string
		want      int
	}{
		{"exact match", []string{"instances", "{id}"}, "Instance", 0},
		{"last word matches", []string{"security-groups", "{id}", "rules"}, "SecurityGroupRule", 2},
		{"nested collection", []string{"buckets", "{b}", "objects", "{k}"}, "Object", 2},
		{"parent when tagged as parent", []string{"instances", "{id}", "volumes"}, "Instance", 0},
		{"service tag matches nothing", []string{"invoices", "{id}"}, "Billing", -1},
		{"absent", []string{"invoices"}, "", -1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := locateXResource(tc.segs, tc.xResource); got != tc.want {
				t.Errorf("locateXResource(%v, %q) = %d, want %d", tc.segs, tc.xResource, got, tc.want)
			}
		})
	}
}
