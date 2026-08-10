package tools

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcano6/mcp-sysbox/internal/sysinfo"
)

// CPUTool returns the MCP tool definition for the system CPU probe.
func CPUTool() mcp.Tool {
	return mcp.NewTool("system_cpu",
		mcp.WithDescription(
			"System CPU probe. Returns CPU model, core counts (physical/logical), "+
				"overall usage percentage, and per-core usage breakdown.",
		),
	)
}

// CPUHandler handles the "system_cpu" tool invocation.
func CPUHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	status, err := sysinfo.GetCPUStatus(500 * time.Millisecond)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get cpu status: %v", err)), nil
	}

	// Format per-core usage
	coreUsages := make([]string, len(status.UsagePerCPU))
	for i, usage := range status.UsagePerCPU {
		coreUsages[i] = fmt.Sprintf("%.0f%%", usage)
	}

	response := fmt.Sprintf(
		"🖥️ CPU Status\n\n"+
			"Model:   %s\n"+
			"Cores:   %d Physical / %d Logical\n"+
			"Usage:   %.1f%%\n\n"+
			"Per-Core: [%s]",
		status.ModelName,
		status.PhysicalCPU,
		status.LogicalCPU,
		status.UsageTotal,
		strings.Join(coreUsages, " "),
	)

	return mcp.NewToolResultText(response), nil
}
