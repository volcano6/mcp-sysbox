package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Version information (set via ldflags at build time)
var (
	version = "dev"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"mcp-sysbox",
		version,
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	// Register tools
	registerTools(s)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

// registerTools registers all MCP tools on the server.
func registerTools(s *server.MCPServer) {
	// --- ping: health-check / connectivity test ---
	pingTool := mcp.NewTool("ping",
		mcp.WithDescription(
			"Health check tool. Returns 'pong' along with server metadata "+
				"to verify the MCP server is running and reachable.",
		),
	)
	s.AddTool(pingTool, pingHandler)
}

// pingHandler handles the "ping" tool invocation.
func pingHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	response := fmt.Sprintf(
		"pong 🏓\n\nServer: mcp-sysbox %s\nGo: %s\nOS/Arch: %s/%s\nTime: %s",
		version,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		time.Now().Format(time.RFC3339),
	)

	return mcp.NewToolResultText(response), nil
}
