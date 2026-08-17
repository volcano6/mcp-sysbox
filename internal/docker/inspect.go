package docker

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/moby/moby/client"
)

// ContainerDetail holds detailed information about a single Docker container.
type ContainerDetail struct {
	ID          string
	Name        string
	Image       string
	State       string
	Status      string
	Pid         int
	StartedAt   string
	FinishedAt  string
	RestartCount int
	Platform    string
	Driver      string
	LogPath     string
	Env         []string
	Cmd         []string
	Ports       map[string][]string // container_port -> []host_port
	Mounts      []MountInfo
	Networks    []NetworkInfo
}

// MountInfo holds mount point information.
type MountInfo struct {
	Source      string
	Destination string
	Mode        string
	RW          bool
}

// NetworkInfo holds container network information.
type NetworkInfo struct {
	Name      string
	IPAddress string
	Gateway   string
	MacAddress string
}

// InspectContainer returns detailed info for a container by name or ID.
func InspectContainer(ctx context.Context, nameOrID string) (*ContainerDetail, error) {
	cli, err := GetClient()
	if err != nil {
		return nil, err
	}

	result, err := cli.ContainerInspect(ctx, nameOrID, client.ContainerInspectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to inspect container %q: %w", nameOrID, err)
	}

	c := result.Container

	detail := &ContainerDetail{
		ID:           c.ID[:12],
		Name:         strings.TrimPrefix(c.Name, "/"),
		Image:        c.Config.Image,
		RestartCount: c.RestartCount,
		Platform:     c.Platform,
		Driver:       c.Driver,
		LogPath:      c.LogPath,
	}

	// State
	if c.State != nil {
		detail.State = string(c.State.Status)
		detail.Status = string(c.State.Status)
		detail.Pid = c.State.Pid
		detail.StartedAt = c.State.StartedAt
		detail.FinishedAt = c.State.FinishedAt
	}

	// Config
	if c.Config != nil {
		detail.Env = c.Config.Env
		detail.Cmd = c.Config.Cmd
		if c.Config.Image != "" {
			detail.Image = c.Config.Image
		}
	}

	// Ports
	detail.Ports = make(map[string][]string)
	if c.NetworkSettings != nil {
		for port, bindings := range c.NetworkSettings.Ports {
			var hostPorts []string
			for _, b := range bindings {
				if b.HostPort != "" {
					hostPorts = append(hostPorts, fmt.Sprintf("%s:%s", b.HostIP, b.HostPort))
				}
			}
			detail.Ports[port.String()] = hostPorts
		}
	}

	// Mounts
	for _, m := range c.Mounts {
		detail.Mounts = append(detail.Mounts, MountInfo{
			Source:      m.Source,
			Destination: m.Destination,
			Mode:        m.Mode,
			RW:          m.RW,
		})
	}

	// Networks
	if c.NetworkSettings != nil {
		for name, net := range c.NetworkSettings.Networks {
			detail.Networks = append(detail.Networks, NetworkInfo{
				Name:       name,
				IPAddress:  net.IPAddress.String(),
				Gateway:    net.Gateway.String(),
				MacAddress: net.MacAddress.String(),
			})
		}
	}

	return detail, nil
}

// FormatDetail formats a ContainerDetail into a human-readable string.
func FormatDetail(d *ContainerDetail) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("🔍 Container: %s (%s)\n", d.Name, d.ID))
	sb.WriteString(fmt.Sprintf("Image:    %s\n", d.Image))
	sb.WriteString(fmt.Sprintf("State:    %s\n", d.State))
	sb.WriteString(fmt.Sprintf("PID:      %d\n", d.Pid))
	sb.WriteString(fmt.Sprintf("Restarts: %d\n", d.RestartCount))

	if d.StartedAt != "" {
		if t, err := time.Parse(time.RFC3339Nano, d.StartedAt); err == nil {
			sb.WriteString(fmt.Sprintf("Started:  %s\n", t.Format("2006-01-02 15:04:05")))
		}
	}

	// Command
	if len(d.Cmd) > 0 {
		sb.WriteString(fmt.Sprintf("Cmd:      %s\n", strings.Join(d.Cmd, " ")))
	}

	// Ports
	if len(d.Ports) > 0 {
		sb.WriteString("\nPorts:\n")
		for cPort, hPorts := range d.Ports {
			if len(hPorts) > 0 {
				sb.WriteString(fmt.Sprintf("  %s → %s\n", cPort, strings.Join(hPorts, ", ")))
			} else {
				sb.WriteString(fmt.Sprintf("  %s (not published)\n", cPort))
			}
		}
	}

	// Mounts
	if len(d.Mounts) > 0 {
		sb.WriteString("\nMounts:\n")
		for _, m := range d.Mounts {
			rw := "ro"
			if m.RW {
				rw = "rw"
			}
			sb.WriteString(fmt.Sprintf("  %s → %s (%s)\n", m.Source, m.Destination, rw))
		}
	}

	// Networks
	if len(d.Networks) > 0 {
		sb.WriteString("\nNetworks:\n")
		for _, n := range d.Networks {
			sb.WriteString(fmt.Sprintf("  %s: IP=%s  GW=%s\n", n.Name, n.IPAddress, n.Gateway))
		}
	}

	// Environment (filter sensitive vars)
	if len(d.Env) > 0 {
		sb.WriteString("\nEnvironment:\n")
		for _, e := range d.Env {
			parts := strings.SplitN(e, "=", 2)
			key := strings.ToUpper(parts[0])
			if containsSensitive(key) {
				sb.WriteString(fmt.Sprintf("  %s=***\n", parts[0]))
			} else {
				sb.WriteString(fmt.Sprintf("  %s\n", e))
			}
		}
	}

	return sb.String()
}

// containsSensitive checks if an env var key looks sensitive.
func containsSensitive(key string) bool {
	sensitive := []string{"PASSWORD", "SECRET", "TOKEN", "KEY", "CREDENTIAL", "PRIVATE"}
	for _, s := range sensitive {
		if strings.Contains(key, s) {
			return true
		}
	}
	return false
}
