package main

import "testing"

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
