package docker

import (
	"context"
	"testing"
)

func TestListContainers(t *testing.T) {
	ctx := context.Background()

	// Skip if Docker is not available
	if err := Ping(ctx); err != nil {
		t.Skipf("Docker not available, skipping: %v", err)
	}

	containers, err := ListContainers(ctx)
	if err != nil {
		t.Fatalf("ListContainers returned error: %v", err)
	}

	t.Logf("Found %d containers", len(containers))
	for _, c := range containers {
		t.Logf("  %-15s %-10s %-25s %s", c.Name, c.State, c.Image, c.Ports)
	}
}
