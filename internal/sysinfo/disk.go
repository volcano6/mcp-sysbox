package sysinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/disk"
)

// DiskPartitionStatus holds information about a single disk partition.
type DiskPartitionStatus struct {
	Device      string
	MountPoint  string
	FSType      string
	TotalBytes  uint64
	UsedBytes   uint64
	FreeBytes   uint64
	UsedPercent float64
}

// GetDiskStatus retrieves the status of all disk partitions.
func GetDiskStatus() ([]DiskPartitionStatus, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, fmt.Errorf("failed to get disk partitions: %w", err)
	}

	var results []DiskPartitionStatus
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			// Skip partitions we can't read (e.g. CD-ROM drives)
			continue
		}

		if usage.Total == 0 {
			// Skip virtual filesystems with no real storage (e.g. macOS autofs)
			continue
		}

		results = append(results, DiskPartitionStatus{
			Device:      p.Device,
			MountPoint:  p.Mountpoint,
			FSType:      p.Fstype,
			TotalBytes:  usage.Total,
			UsedBytes:   usage.Used,
			FreeBytes:   usage.Free,
			UsedPercent: usage.UsedPercent,
		})
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no readable disk partitions found")
	}

	return results, nil
}
