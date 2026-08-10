package sysinfo

import (
	"testing"
	"time"
)

func TestGetCPUStatus(t *testing.T) {
	status, err := GetCPUStatus(200 * time.Millisecond)
	if err != nil {
		t.Fatalf("GetCPUStatus returned error: %v", err)
	}

	if status.ModelName == "" {
		t.Error("ModelName should not be empty")
	}

	if status.PhysicalCPU <= 0 {
		t.Errorf("PhysicalCPU should be > 0, got: %d", status.PhysicalCPU)
	}

	if status.LogicalCPU <= 0 {
		t.Errorf("LogicalCPU should be > 0, got: %d", status.LogicalCPU)
	}

	if status.LogicalCPU < status.PhysicalCPU {
		t.Errorf("LogicalCPU (%d) should be >= PhysicalCPU (%d)", status.LogicalCPU, status.PhysicalCPU)
	}

	if status.UsageTotal < 0 || status.UsageTotal > 100 {
		t.Errorf("UsageTotal should be between 0 and 100, got: %.1f", status.UsageTotal)
	}

	if len(status.UsagePerCPU) == 0 {
		t.Error("UsagePerCPU should not be empty")
	}

	for i, usage := range status.UsagePerCPU {
		if usage < 0 || usage > 100 {
			t.Errorf("UsagePerCPU[%d] should be between 0 and 100, got: %.1f", i, usage)
		}
	}

	t.Logf("CPU: %s | %d Physical / %d Logical | Usage: %.1f%%",
		status.ModelName,
		status.PhysicalCPU,
		status.LogicalCPU,
		status.UsageTotal,
	)
	t.Logf("Per-Core: %v", status.UsagePerCPU)
}
