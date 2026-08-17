package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcano6/mcp-sysbox/internal/docker"
)

// DockerInspectTool returns the MCP tool definition for inspecting a Docker container.
func DockerInspectTool() mcp.Tool {
	return mcp.NewTool("docker_inspect",
		mcp.WithDescription(
			"Inspect a Docker container by name or ID. "+
				"Returns detailed information including state, ports, mounts, "+
				"networks, environment variables and command.",
		),
		mcp.WithString("container",
			mcp.Required(),
			mcp.Description("Container name or ID to inspect"),
		),
	)
}

// DockerInspectHandler handles the "docker_inspect" tool invocation.
func DockerInspectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	nameOrID := request.GetString("container", "")
	if nameOrID == "" {
		return mcp.NewToolResultError("parameter 'container' is required (container name or ID)"), nil
	}

	detail, err := docker.InspectContainer(ctx, nameOrID)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to inspect container: %v", err)), nil
	}

	return mcp.NewToolResultText(docker.FormatDetail(detail)), nil
}
