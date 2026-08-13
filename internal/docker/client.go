package docker

import (
	"context"
	"fmt"
	"sync"

	"github.com/moby/moby/client"
)

var (
	once       sync.Once
	dockerCli  *client.Client
	initErr    error
)

// GetClient returns a shared Docker client instance.
// It lazily initializes the client on first call.
func GetClient() (*client.Client, error) {
	once.Do(func() {
		dockerCli, initErr = client.NewClientWithOpts(
			client.FromEnv,
			client.WithAPIVersionNegotiation(),
		)
	})

	if initErr != nil {
		return nil, fmt.Errorf("failed to create docker client: %w", initErr)
	}

	return dockerCli, nil
}

// Ping checks if the Docker daemon is reachable.
func Ping(ctx context.Context) error {
	cli, err := GetClient()
	if err != nil {
		return err
	}

	_, err = cli.Ping(ctx, client.PingOptions{})
	if err != nil {
		return fmt.Errorf("docker daemon not reachable: %w", err)
	}

	return nil
}
