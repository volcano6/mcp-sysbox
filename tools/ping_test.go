package tools

import (
	"context"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestPingHandler(t *testing.T) {
	request := mcp.CallToolRequest{}

	result, err := PingHandler(context.Background(), request)
	if err != nil {
		t.Fatalf("PingHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("PingHandler returned nil result")
	}

	if len(result.Content) == 0 {
		t.Fatal("PingHandler returned empty content")
	}

	textContent, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", result.Content[0])
	}

	if len(textContent.Text) == 0 {
		t.Fatal("PingHandler returned empty text")
	}

	expectedSubstrings := []string{"pong", "mcp-sysbox", "OS/Arch"}
	for _, sub := range expectedSubstrings {
		if !strings.Contains(textContent.Text, sub) {
			t.Errorf("expected response to contain %q, got: %s", sub, textContent.Text)
		}
	}

	t.Logf("ping response:\n%s", textContent.Text)
}
