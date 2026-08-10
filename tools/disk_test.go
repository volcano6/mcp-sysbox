package tools

import (
	"context"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestDiskHandler(t *testing.T) {
	request := mcp.CallToolRequest{}

	result, err := DiskHandler(context.Background(), request)
	if err != nil {
		t.Fatalf("DiskHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("DiskHandler returned nil result")
	}

	if len(result.Content) == 0 {
		t.Fatal("DiskHandler returned empty content")
	}

	textContent, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", result.Content[0])
	}

	expectedSubstrings := []string{"Disk Status", "Total:", "Used:", "Free:"}
	for _, sub := range expectedSubstrings {
		if !strings.Contains(textContent.Text, sub) {
			t.Errorf("expected response to contain %q, got: %s", sub, textContent.Text)
		}
	}

	t.Logf("disk response:\n%s", textContent.Text)
}
