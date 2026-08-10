package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcano6/mcp-sysbox/internal/sysinfo"
)

// DiskTool returns the MCP tool definition for the system disk probe.
func DiskTool() mcp.Tool {
	return mcp.NewTool("system_disk",
		mcp.WithDescription(
			"System disk probe. Returns all mounted disk partitions with "+
				"total capacity, used space, free space and usage percentage.",
		),
	)
}

// DiskHandler handles the "system_disk" tool invocation.
func DiskHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	partitions, err := sysinfo.GetDiskStatus()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get disk status: %v", err)), nil
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("💿 Disk Status (%d partitions)\n", len(partitions)))

	for _, p := range partitions {
		sb.WriteString(fmt.Sprintf(
			"\n%s  (%s)   Total: %s  Used: %s (%.1f%%)  Free: %s",
			p.MountPoint,
			p.FSType,
			sysinfo.FormatBytes(p.TotalBytes),
			sysinfo.FormatBytes(p.UsedBytes),
			p.UsedPercent,
			sysinfo.FormatBytes(p.FreeBytes),
		))
	}

	return mcp.NewToolResultText(sb.String()), nil
}
