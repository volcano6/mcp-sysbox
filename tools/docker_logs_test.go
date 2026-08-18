package tools

import (
	"context"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcano6/mcp-sysbox/internal/docker"
)

func TestDockerLogsHandler(t *testing.T) {
	ctx := context.Background()

	if err := docker.Ping(ctx); err != nil {
		t.Skipf("Docker not available, skipping: %v", err)
	}

	containers, err := docker.ListContainers(ctx)
	if err != nil || len(containers) == 0 {
		t.Skip("No containers available")
	}

	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"container": containers[0].Name,
		"tail":      "5",
	}

	result, err := DockerLogsHandler(ctx, request)
	if err != nil {
		t.Fatalf("DockerLogsHandler returned error: %v", err)
	}

	if result == nil || len(result.Content) == 0 {
		t.Fatal("DockerLogsHandler returned empty result")
	}

	textContent, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", result.Content[0])
	}

	if !strings.Contains(textContent.Text, "Logs:") {
		t.Errorf("expected 'Logs:' header in response, got: %s", textContent.Text)
	}

	t.Logf("docker_logs response:\n%s", textContent.Text)
}

func TestDockerLogsHandler_MissingParam(t *testing.T) {
	request := mcp.CallToolRequest{}

	result, err := DockerLogsHandler(context.Background(), request)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result.IsError {
		t.Error("expected error result when container param is missing")
	}
}
