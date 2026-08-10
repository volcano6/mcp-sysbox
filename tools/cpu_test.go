package tools

import (
	"context"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestCPUHandler(t *testing.T) {
	request := mcp.CallToolRequest{}

	result, err := CPUHandler(context.Background(), request)
	if err != nil {
		t.Fatalf("CPUHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("CPUHandler returned nil result")
	}

	if len(result.Content) == 0 {
		t.Fatal("CPUHandler returned empty content")
	}

	textContent, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", result.Content[0])
	}

	expectedSubstrings := []string{"CPU Status", "Model:", "Cores:", "Usage:", "Per-Core:"}
	for _, sub := range expectedSubstrings {
		if !strings.Contains(textContent.Text, sub) {
			t.Errorf("expected response to contain %q, got: %s", sub, textContent.Text)
		}
	}

	t.Logf("cpu response:\n%s", textContent.Text)
}
