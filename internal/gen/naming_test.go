package main

import "testing"

func TestExportedName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"instance_id", "InstanceID"},
		{"id", "ID"},
		{"vpc_id", "VPCID"},
		{"primary_ipv6", "PrimaryIPv6"},
		{"crn", "CRN"},
		{"iam_role_id", "IAMRoleID"},
		{"has_more", "HasMore"},
		{"listInstances", "ListInstances"},
		{"getOAuthToken", "GetOAuthToken"},
		{"min_disk_gb", "MinDiskGB"},
		// Repeated-array query parameters must not carry brackets into an
		// identifier.
		{"by[]", "By"},
		{"match[]", "Match"},
		{"l7_policy_id", "L7PolicyID"},
		{"nat_gateway_id", "NATGatewayID"},
		{"dnssec", "DNSSEC"},
		{"", "Field"},
	}
	for _, tc := range tests {
		if got := exportedName(tc.in); got != tc.want {
			t.Errorf("exportedName(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestUnexportedName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"instance_id", "instanceID"},
		{"id", "id"},
		{"bucket", "bucket"},
		{"key", "key"},
		{"vpc_id", "vpcID"},
		{"l7_rule_id", "l7RuleID"},
		// A parameter that would shadow a builtin or keyword.
		{"type", "type_"},
		{"range", "range_"},
	}
	for _, tc := range tests {
		if got := unexportedName(tc.in); got != tc.want {
			t.Errorf("unexportedName(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestConjugateSummary(t *testing.T) {
	tests := []struct {
		name, goName, summary, want string
	}{
		{"verb matches the method", "ListInstances", "List instances.", "lists instances."},
		{"third person spelling", "AttachVolume", "Attach a volume.", "attaches a volume."},
		{"y becomes ies", "QueryLogs", "Query logs.", "queries logs."},
		{"known verb that differs from the method", "GetObject", "Download object.", "downloads object."},
		{"trailing comma is kept with the sentence", "UpdateLoadBalancer", "Rename, scale, or resize.", "renames, scale, or resize."},
		// "Instant" is a noun here; bending it into a verb would be wrong,
		// so the summary is kept verbatim behind a dash.
		{"noun summary is left alone", "QueryMetricsInstant", "Instant structured metric query.", "— Instant structured metric query."},
		{"empty", "GetThing", "", ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := conjugateSummary(tc.goName, tc.summary); got != tc.want {
				t.Errorf("conjugateSummary(%q, %q) = %q, want %q", tc.goName, tc.summary, got, tc.want)
			}
		})
	}
}

func TestThirdPerson(t *testing.T) {
	tests := []struct{ in, want string }{
		{"get", "gets"}, {"list", "lists"}, {"attach", "attaches"},
		{"detach", "detaches"}, {"query", "queries"}, {"push", "pushes"},
		{"fix", "fixes"}, {"do", "does"}, {"destroy", "destroys"},
	}
	for _, tc := range tests {
		if got := thirdPerson(tc.in); got != tc.want {
			t.Errorf("thirdPerson(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestPathPlaceholdersAreInOrder(t *testing.T) {
	got := pathPlaceholders("/v1/buckets/{bucket}/objects/{key}")
	if len(got) != 2 || got[0] != "bucket" || got[1] != "key" {
		t.Errorf("pathPlaceholders() = %v, want [bucket key]", got)
	}
	if got := pathPlaceholders("/v1/instances"); len(got) != 0 {
		t.Errorf("pathPlaceholders() = %v, want none", got)
	}
}

func TestWrapTextReflowsProseAndLeavesCodeAlone(t *testing.T) {
	in := "a description that\narrives already broken\nacross lines\n\n```\ncurl -s https://example.test \\\n  --data x\n```\n\n- a list item that is quite long and should not be reflowed either"
	got := wrapText(in, 40)

	// The prose paragraph is rejoined and rewrapped, so no line exceeds the
	// width and the original break points are gone.
	lines := splitLines(got)
	for _, l := range lines {
		if len(l) > 40 && !hasPrefix(l, "  ") && !hasPrefix(l, "-") && !hasPrefix(l, "curl") {
			t.Errorf("line exceeds the width: %q", l)
		}
	}
	if !containsLine(lines, "```") {
		t.Error("the code fence was lost")
	}
	if !containsLine(lines, "  --data x") {
		t.Errorf("an indented code line was reflowed: %q", got)
	}
}

func splitLines(s string) []string {
	var out, cur = []string{}, ""
	for _, r := range s {
		if r == '\n' {
			out = append(out, cur)
			cur = ""
			continue
		}
		cur += string(r)
	}
	return append(out, cur)
}

func hasPrefix(s, p string) bool { return len(s) >= len(p) && s[:len(p)] == p }

func containsLine(lines []string, want string) bool {
	for _, l := range lines {
		if l == want {
			return true
		}
	}
	return false
}
