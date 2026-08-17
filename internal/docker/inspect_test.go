package docker

import (
	"context"
	"testing"
)

func TestInspectContainer(t *testing.T) {
	ctx := context.Background()

	// Skip if Docker is not available
	if err := Ping(ctx); err != nil {
		t.Skipf("Docker not available, skipping: %v", err)
	}

	// List containers to find one to inspect
	containers, err := ListContainers(ctx)
	if err != nil {
		t.Fatalf("ListContainers returned error: %v", err)
	}
	if len(containers) == 0 {
		t.Skip("No containers available to inspect")
	}

	// Inspect the first container
	detail, err := InspectContainer(ctx, containers[0].ID)
	if err != nil {
		t.Fatalf("InspectContainer returned error: %v", err)
	}

	if detail.Name == "" {
		t.Error("Name should not be empty")
	}
	if detail.Image == "" {
		t.Error("Image should not be empty")
	}

	t.Logf("Inspect result:\n%s", FormatDetail(detail))
}
