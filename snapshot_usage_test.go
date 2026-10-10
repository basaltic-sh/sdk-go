package basaltic_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/basaltic-sh/sdk-go/storage"
)

func TestSnapshotUsageJSONRoundTrip(t *testing.T) {
	for _, tc := range []struct{ name, raw string }{
		{"unknown", `{"state":"unknown","scope":"volume_lineage","billable":false,"measured_at":null,"lineage_retained_bytes":null}`},
		{"stale", `{"state":"stale","scope":"volume_lineage","billable":false,"measured_at":"2026-10-10T12:30:00Z","lineage_retained_bytes":null}`},
		{"measured_zero", `{"state":"measured","scope":"volume_lineage","billable":false,"measured_at":"2026-10-10T12:30:00Z","lineage_retained_bytes":0}`},
		{"measured_nonzero", `{"state":"measured","scope":"volume_lineage","billable":false,"measured_at":"2026-10-10T12:30:00Z","lineage_retained_bytes":4096}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var usage storage.SnapshotUsage
			if err := json.Unmarshal([]byte(tc.raw), &usage); err != nil {
				t.Fatal(err)
			}
			encoded, err := json.Marshal(usage)
			if err != nil {
				t.Fatal(err)
			}
			var want, got any
			if err := json.Unmarshal([]byte(tc.raw), &want); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("round trip changed measurement: want %s got %s", tc.raw, encoded)
			}
		})
	}
	t.Run("omitted_optional_usage", func(t *testing.T) {
		var snapshot storage.Snapshot
		if err := json.Unmarshal([]byte(`{"id":"snapshot-id"}`), &snapshot); err != nil {
			t.Fatal(err)
		}
		if snapshot.SnapshotUsage != nil {
			t.Fatal("missing usage became a measurement")
		}
		encoded, err := json.Marshal(snapshot)
		if err != nil {
			t.Fatal(err)
		}
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(encoded, &fields); err != nil {
			t.Fatal(err)
		}
		if _, exists := fields["snapshot_usage"]; exists {
			t.Fatal("omitted usage appeared on the wire")
		}
	})
}
