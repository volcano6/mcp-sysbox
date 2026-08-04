package sysinfo

import (
	"testing"
)

func TestGetMemoryStatus(t *testing.T) {
	status, err := GetMemoryStatus()
	if err != nil {
		t.Fatalf("GetMemoryStatus returned error: %v", err)
	}

	if status.TotalBytes == 0 {
		t.Error("TotalBytes should not be zero")
	}

	if status.UsedBytes == 0 {
		t.Error("UsedBytes should not be zero")
	}

	if status.AvailableBytes == 0 {
		t.Error("AvailableBytes should not be zero")
	}

	if status.UsedPercent <= 0 || status.UsedPercent > 100 {
		t.Errorf("UsedPercent should be between 0 and 100, got: %.1f", status.UsedPercent)
	}

	t.Logf("Memory: Total=%s Used=%s (%.1f%%) Available=%s",
		FormatBytes(status.TotalBytes),
		FormatBytes(status.UsedBytes),
		status.UsedPercent,
		FormatBytes(status.AvailableBytes),
	)
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "0 B"},
		{512, "512 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{1048576, "1.0 MB"},
		{1073741824, "1.0 GB"},
		{1099511627776, "1.0 TB"},
		{17179869184, "16.0 GB"},
	}

	for _, tt := range tests {
		result := FormatBytes(tt.input)
		if result != tt.expected {
			t.Errorf("FormatBytes(%d) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}
