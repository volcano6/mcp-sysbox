package docker

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

// ContainerInfo holds summary information about a Docker container.
type ContainerInfo struct {
	ID      string
	Name    string
	Image   string
	State   string // running, exited, paused, etc.
	Status  string // e.g. "Up 3 days", "Exited (0) 2 days ago"
	Ports   string
	Created time.Time
}

// ListContainers returns a list of all containers (including stopped ones).
func ListContainers(ctx context.Context) ([]ContainerInfo, error) {
	cli, err := GetClient()
	if err != nil {
		return nil, err
	}

	result, err := cli.ContainerList(ctx, client.ContainerListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	containers := result.Items
	infos := make([]ContainerInfo, 0, len(containers))
	for _, c := range containers {
		name := ""
		if len(c.Names) > 0 {
			// Docker prepends "/" to container names
			name = strings.TrimPrefix(c.Names[0], "/")
		}

		ports := formatPorts(c.Ports)

		infos = append(infos, ContainerInfo{
			ID:      c.ID[:12],
			Name:    name,
			Image:   c.Image,
			State:   string(c.State),
			Status:  c.Status,
			Ports:   ports,
			Created: time.Unix(c.Created, 0),
		})
	}

	return infos, nil
}

// formatPorts formats port bindings into a readable string.
func formatPorts(ports []container.PortSummary) string {
	if len(ports) == 0 {
		return "—"
	}

	seen := make(map[string]bool)
	var parts []string

	for _, p := range ports {
		var s string
		if p.PublicPort != 0 {
			s = fmt.Sprintf("%d:%d/%s", p.PublicPort, p.PrivatePort, p.Type)
		} else {
			s = fmt.Sprintf("%d/%s", p.PrivatePort, p.Type)
		}
		if !seen[s] {
			seen[s] = true
			parts = append(parts, s)
		}
	}

	return strings.Join(parts, ", ")
}
