package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcano6/mcp-sysbox/internal/sysinfo"
)

// MemoryTool returns the MCP tool definition for the system memory probe.
func MemoryTool() mcp.Tool {
	return mcp.NewTool("system_memory",
		mcp.WithDescription(
			"System memory probe. Returns current memory usage including "+
				"total, used, available memory and usage percentage.",
		),
	)
}

// MemoryHandler handles the "system_memory" tool invocation.
func MemoryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	status, err := sysinfo.GetMemoryStatus()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get memory status: %v", err)), nil
	}

	response := fmt.Sprintf(
		"💾 Memory Status\n\n"+
			"Total:      %s\n"+
			"Used:       %s (%.1f%%)\n"+
			"Available:  %s",
		sysinfo.FormatBytes(status.TotalBytes),
		sysinfo.FormatBytes(status.UsedBytes),
		status.UsedPercent,
		sysinfo.FormatBytes(status.AvailableBytes),
	)

	return mcp.NewToolResultText(response), nil
}
