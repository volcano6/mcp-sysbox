package docker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/moby/moby/client"
)

// LogsOptions configures the container logs query.
type LogsOptions struct {
	Tail       string // Number of lines from the end, e.g. "100", "all"
	Since      string // Show logs since timestamp (e.g. "2024-01-01") or duration (e.g. "1h")
	Timestamps bool   // Prepend timestamps to each line
}

// GetContainerLogs retrieves logs for a container by name or ID.
func GetContainerLogs(ctx context.Context, nameOrID string, opts LogsOptions) (string, error) {
	cli, err := GetClient()
	if err != nil {
		return "", err
	}

	tail := opts.Tail
	if tail == "" {
		tail = "100" // Default to last 100 lines
	}

	result, err := cli.ContainerLogs(ctx, nameOrID, client.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       tail,
		Since:      opts.Since,
		Timestamps: opts.Timestamps,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get logs for container %q: %w", nameOrID, err)
	}
	defer result.Close()

	// Demux stdout/stderr streams
	var stdout, stderr bytes.Buffer
	_, err = stdcopy.StdCopy(&stdout, &stderr, result)
	if err != nil {
		// Fallback: container might be using TTY (no multiplexing)
		// Re-read from the result stream directly
		raw, readErr := io.ReadAll(result)
		if readErr != nil {
			return "", fmt.Errorf("failed to read logs: %w", err)
		}
		return string(raw), nil
	}

	var sb strings.Builder
	if stdout.Len() > 0 {
		sb.WriteString(stdout.String())
	}
	if stderr.Len() > 0 {
		if sb.Len() > 0 {
			sb.WriteString("\n--- stderr ---\n")
		}
		sb.WriteString(stderr.String())
	}

	if sb.Len() == 0 {
		return "(no logs)", nil
	}

	return sb.String(), nil
}
