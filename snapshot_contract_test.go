package basaltic_test

import (
	"encoding/json"
	"github.com/basaltic-sh/sdk-go/storage"
	"testing"
)

func TestSnapshotSizeAndReferences(t *testing.T) {
	raw := `{"id":"snapshot-id","size_bytes":10737418240,"volume":{"id":"volume-id","crn":"crn:storage:r:acct:volume/data"},"snapshot_policy":{"id":"policy-id","crn":"crn:storage:r:acct:volume/data/snapshot-policy/daily"}}`
	var snapshot storage.Snapshot
	if err := json.Unmarshal([]byte(raw), &snapshot); err != nil {
		t.Fatal(err)
	}
	if snapshot.SizeBytes != 10737418240 || snapshot.Volume.ID != "volume-id" || snapshot.SnapshotPolicy.ID != "policy-id" {
		t.Fatalf("lost snapshot size or identity: %+v", snapshot)
	}
	encoded, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &fields); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"snapshot_usage", "size_gb", "logical_size_bytes", "volume_id", "snapshot_policy_id"} {
		if _, exists := fields[key]; exists {
			t.Fatalf("removed field %s present", key)
		}
	}
}
