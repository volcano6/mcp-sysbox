package main

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestPingHandler(t *testing.T) {
	request := mcp.CallToolRequest{}

	result, err := pingHandler(context.Background(), request)
	if err != nil {
		t.Fatalf("pingHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("pingHandler returned nil result")
	}

	// Check that the result contains text content
	if len(result.Content) == 0 {
		t.Fatal("pingHandler returned empty content")
	}

	// Verify the response starts with "pong"
	textContent, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", result.Content[0])
	}

	if len(textContent.Text) == 0 {
		t.Fatal("pingHandler returned empty text")
	}

	// Check that the response contains expected substrings
	expectedSubstrings := []string{"pong", "mcp-sysbox", "OS/Arch"}
	for _, sub := range expectedSubstrings {
		if !contains(textContent.Text, sub) {
			t.Errorf("expected response to contain %q, got: %s", sub, textContent.Text)
		}
	}

	t.Logf("ping response:\n%s", textContent.Text)
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
