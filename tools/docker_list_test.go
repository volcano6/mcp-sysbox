package tools

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcano6/mcp-sysbox/internal/docker"
)

func TestDockerListHandler(t *testing.T) {
	ctx := context.Background()

	// Skip if Docker is not available
	if err := docker.Ping(ctx); err != nil {
		t.Skipf("Docker not available, skipping: %v", err)
	}

	request := mcp.CallToolRequest{}

	result, err := DockerListHandler(ctx, request)
	if err != nil {
		t.Fatalf("DockerListHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("DockerListHandler returned nil result")
	}

	if len(result.Content) == 0 {
		t.Fatal("DockerListHandler returned empty content")
	}

	textContent, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", result.Content[0])
	}

	t.Logf("docker_list response:\n%s", textContent.Text)
}
