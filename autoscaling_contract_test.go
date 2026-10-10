package basaltic_test

import (
	"encoding/json"
	"testing"

	"github.com/basaltic-sh/sdk-go/compute"
	"github.com/basaltic-sh/sdk-go/loadbalancer"
)

func TestAutoscalingOptionalWireValues(t *testing.T) {
	zero, two := 0, 2
	for _, tc := range []struct {
		name    string
		request any
		want    string
	}{
		{"bounds only", loadbalancer.UpdateLoadBalancerRequest{MaxCount: &two}, `{"max_count":2}`},
		{"explicit desired zero", loadbalancer.UpdateLoadBalancerRequest{DesiredCount: &zero}, `{"desired_count":0}`},
		{"replica alias", loadbalancer.UpdateLoadBalancerRequest{ReplicaCount: &two}, `{"replica_count":2}`},
		{"disabled policy keeps zero", compute.InstancePoolUpdateRequest{Autoscaling: &compute.AutoscalingPolicy{Enabled: false, DrainSeconds: &zero, Metrics: []*compute.ScalingMetric{}}}, `{"autoscaling":{"drain_seconds":0,"enabled":false,"metrics":[]}}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.request)
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != tc.want {
				t.Fatalf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestFlavorLimitsSurviveClientDecoding(t *testing.T) {
	var flavor compute.Flavor
	if err := json.Unmarshal([]byte(`{"name":"d1.xlarge","net_mbps":10000,"cpu_baseline_pct":100,"cpu_burst_pct":100}`), &flavor); err != nil {
		t.Fatal(err)
	}
	if flavor.NetMbps != 10000 || flavor.CPUBaselinePct != 100 || flavor.CPUBurstPct != 100 {
		t.Fatalf("limits lost: %#v", flavor)
	}
}
