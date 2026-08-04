package sysinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

// MemoryStatus holds system memory information.
type MemoryStatus struct {
	TotalBytes     uint64
	UsedBytes      uint64
	AvailableBytes uint64
	UsedPercent    float64
}

// GetMemoryStatus retrieves the current system memory status.
func GetMemoryStatus() (*MemoryStatus, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory info: %w", err)
	}

	return &MemoryStatus{
		TotalBytes:     v.Total,
		UsedBytes:      v.Used,
		AvailableBytes: v.Available,
		UsedPercent:    v.UsedPercent,
	}, nil
}

// FormatBytes converts bytes to a human-readable string (e.g. "8.0 GB").
func FormatBytes(bytes uint64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.1f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.1f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.1f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.1f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}
