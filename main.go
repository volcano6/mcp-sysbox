package main

import (
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/volcano6/mcp-sysbox/tools"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"mcp-sysbox",
		tools.Version,
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	// Register tools
	s.AddTool(tools.PingTool(), tools.PingHandler)
	s.AddTool(tools.MemoryTool(), tools.MemoryHandler)
	s.AddTool(tools.CPUTool(), tools.CPUHandler)
	s.AddTool(tools.DiskTool(), tools.DiskHandler)
	s.AddTool(tools.DockerListTool(), tools.DockerListHandler)
	s.AddTool(tools.DockerInspectTool(), tools.DockerInspectHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
