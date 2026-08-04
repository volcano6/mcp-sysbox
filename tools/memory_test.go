package tools

import (
	"context"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestMemoryHandler(t *testing.T) {
	request := mcp.CallToolRequest{}

	result, err := MemoryHandler(context.Background(), request)
	if err != nil {
		t.Fatalf("MemoryHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("MemoryHandler returned nil result")
	}

	if len(result.Content) == 0 {
		t.Fatal("MemoryHandler returned empty content")
	}

	textContent, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", result.Content[0])
	}

	expectedSubstrings := []string{"Memory Status", "Total:", "Used:", "Available:"}
	for _, sub := range expectedSubstrings {
		if !strings.Contains(textContent.Text, sub) {
			t.Errorf("expected response to contain %q, got: %s", sub, textContent.Text)
		}
	}

	t.Logf("memory response:\n%s", textContent.Text)
}
