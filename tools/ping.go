package tools

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

// Version is set at build time via ldflags.
var Version = "dev"

// PingTool returns the MCP tool definition for the ping health-check.
func PingTool() mcp.Tool {
	return mcp.NewTool("ping",
		mcp.WithDescription(
			"Health check tool. Returns 'pong' along with server metadata "+
				"to verify the MCP server is running and reachable.",
		),
	)
}

// PingHandler handles the "ping" tool invocation.
func PingHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	response := fmt.Sprintf(
		"pong 🏓\n\nServer: mcp-sysbox %s\nGo: %s\nOS/Arch: %s/%s\nTime: %s",
		Version,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		time.Now().Format(time.RFC3339),
	)

	return mcp.NewToolResultText(response), nil
}
