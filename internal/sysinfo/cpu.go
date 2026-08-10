package sysinfo

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

// CPUStatus holds system CPU information.
type CPUStatus struct {
	ModelName    string
	PhysicalCPU int
	LogicalCPU  int
	UsageTotal  float64   // overall usage percentage
	UsagePerCPU []float64 // per-logical-core usage percentage
}

// GetCPUStatus retrieves the current system CPU status.
// It samples CPU usage over the given interval duration.
func GetCPUStatus(interval time.Duration) (*CPUStatus, error) {
	// Get CPU model info
	infos, err := cpu.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get cpu info: %w", err)
	}

	modelName := "Unknown"
	if len(infos) > 0 {
		modelName = infos[0].ModelName
	}

	// Get core counts
	physical, err := cpu.Counts(false)
	if err != nil {
		return nil, fmt.Errorf("failed to get physical cpu count: %w", err)
	}

	logical, err := cpu.Counts(true)
	if err != nil {
		return nil, fmt.Errorf("failed to get logical cpu count: %w", err)
	}

	// Get overall CPU usage
	totalPercent, err := cpu.Percent(interval, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get cpu usage: %w", err)
	}

	usageTotal := 0.0
	if len(totalPercent) > 0 {
		usageTotal = totalPercent[0]
	}

	// Get per-core CPU usage
	perCPU, err := cpu.Percent(0, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get per-cpu usage: %w", err)
	}

	return &CPUStatus{
		ModelName:    modelName,
		PhysicalCPU:  physical,
		LogicalCPU:   logical,
		UsageTotal:   usageTotal,
		UsagePerCPU:  perCPU,
	}, nil
}
