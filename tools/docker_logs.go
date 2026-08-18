package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcano6/mcp-sysbox/internal/docker"
)

// DockerLogsTool returns the MCP tool definition for reading container logs.
func DockerLogsTool() mcp.Tool {
	return mcp.NewTool("docker_logs",
		mcp.WithDescription(
			"Read logs from a Docker container. "+
				"Returns stdout and stderr output. "+
				"Supports tail (number of lines) and since (time filter) parameters.",
		),
		mcp.WithString("container",
			mcp.Required(),
			mcp.Description("Container name or ID"),
		),
		mcp.WithString("tail",
			mcp.Description("Number of lines from the end to show (default: 100)"),
		),
		mcp.WithString("since",
			mcp.Description("Show logs since timestamp (e.g. '2024-01-01') or relative duration (e.g. '1h', '30m')"),
		),
		mcp.WithBoolean("timestamps",
			mcp.Description("Prepend RFC3339 timestamps to each log line"),
		),
	)
}

// DockerLogsHandler handles the "docker_logs" tool invocation.
func DockerLogsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	nameOrID := request.GetString("container", "")
	if nameOrID == "" {
		return mcp.NewToolResultError("parameter 'container' is required (container name or ID)"), nil
	}

	opts := docker.LogsOptions{
		Tail:       request.GetString("tail", "100"),
		Since:      request.GetString("since", ""),
		Timestamps: request.GetBool("timestamps", false),
	}

	logs, err := docker.GetContainerLogs(ctx, nameOrID, opts)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get logs: %v", err)), nil
	}

	header := fmt.Sprintf("📋 Logs: %s (tail=%s)\n\n", nameOrID, opts.Tail)
	return mcp.NewToolResultText(header + logs), nil
}
