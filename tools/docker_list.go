package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcano6/mcp-sysbox/internal/docker"
)

// DockerListTool returns the MCP tool definition for listing Docker containers.
func DockerListTool() mcp.Tool {
	return mcp.NewTool("docker_list",
		mcp.WithDescription(
			"List all Docker containers (running and stopped). "+
				"Returns container name, state, image, ports and status.",
		),
	)
}

// DockerListHandler handles the "docker_list" tool invocation.
func DockerListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	containers, err := docker.ListContainers(ctx)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list containers: %v", err)), nil
	}

	if len(containers) == 0 {
		return mcp.NewToolResultText("🐳 No containers found."), nil
	}

	// Count by state
	running := 0
	for _, c := range containers {
		if c.State == "running" {
			running++
		}
	}
	stopped := len(containers) - running

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("🐳 Containers (%d running, %d stopped)\n", running, stopped))

	for _, c := range containers {
		stateIcon := "⏹"
		if c.State == "running" {
			stateIcon = "▶"
		} else if c.State == "paused" {
			stateIcon = "⏸"
		}

		sb.WriteString(fmt.Sprintf(
			"\n%s %-20s %-10s %-25s %s\n"+
				"  Status: %s  |  Ports: %s",
			stateIcon,
			c.Name,
			c.State,
			c.Image,
			c.ID,
			c.Status,
			c.Ports,
		))
	}

	return mcp.NewToolResultText(sb.String()), nil
}
