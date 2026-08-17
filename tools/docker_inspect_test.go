package tools

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcano6/mcp-sysbox/internal/docker"
)

func TestDockerInspectHandler(t *testing.T) {
	ctx := context.Background()

	// Skip if Docker is not available
	if err := docker.Ping(ctx); err != nil {
		t.Skipf("Docker not available, skipping: %v", err)
	}

	// Need a container to inspect
	containers, err := docker.ListContainers(ctx)
	if err != nil || len(containers) == 0 {
		t.Skip("No containers available to inspect")
	}

	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"container": containers[0].Name,
	}

	result, err := DockerInspectHandler(ctx, request)
	if err != nil {
		t.Fatalf("DockerInspectHandler returned error: %v", err)
	}

	if result == nil || len(result.Content) == 0 {
		t.Fatal("DockerInspectHandler returned empty result")
	}

	textContent, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", result.Content[0])
	}

	t.Logf("docker_inspect response:\n%s", textContent.Text)
}

func TestDockerInspectHandler_MissingParam(t *testing.T) {
	request := mcp.CallToolRequest{}

	result, err := DockerInspectHandler(context.Background(), request)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result.IsError {
		t.Error("expected error result when container param is missing")
	}
}
