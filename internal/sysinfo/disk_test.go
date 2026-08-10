package sysinfo

import (
	"testing"
)

func TestGetDiskStatus(t *testing.T) {
	partitions, err := GetDiskStatus()
	if err != nil {
		t.Fatalf("GetDiskStatus returned error: %v", err)
	}

	if len(partitions) == 0 {
		t.Fatal("GetDiskStatus returned no partitions")
	}

	for _, p := range partitions {
		if p.MountPoint == "" {
			t.Error("MountPoint should not be empty")
		}

		if p.TotalBytes == 0 {
			t.Errorf("TotalBytes should not be zero for %s", p.MountPoint)
		}

		if p.UsedPercent < 0 || p.UsedPercent > 100 {
			t.Errorf("UsedPercent should be between 0 and 100 for %s, got: %.1f", p.MountPoint, p.UsedPercent)
		}

		t.Logf("Disk: %s (%s)  Total=%s  Used=%s (%.1f%%)  Free=%s",
			p.MountPoint,
			p.FSType,
			FormatBytes(p.TotalBytes),
			FormatBytes(p.UsedBytes),
			p.UsedPercent,
			FormatBytes(p.FreeBytes),
		)
	}
}
