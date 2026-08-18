package docker

import (
	"context"
	"testing"
)

func TestGetContainerLogs(t *testing.T) {
	ctx := context.Background()

	if err := Ping(ctx); err != nil {
		t.Skipf("Docker not available, skipping: %v", err)
	}

	containers, err := ListContainers(ctx)
	if err != nil || len(containers) == 0 {
		t.Skip("No containers available")
	}

	logs, err := GetContainerLogs(ctx, containers[0].Name, LogsOptions{Tail: "10"})
	if err != nil {
		t.Fatalf("GetContainerLogs returned error: %v", err)
	}

	t.Logf("Logs (last 10 lines):\n%s", logs)
}
