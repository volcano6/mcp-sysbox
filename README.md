# mcp-sysbox 🛠️

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![CI](https://github.com/volcano6/mcp-sysbox/actions/workflows/ci.yml/badge.svg)](https://github.com/volcano6/mcp-sysbox/actions/workflows/ci.yml)

**mcp-sysbox** (Go MCP Ops Server) is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server built in Go. It provides LLMs (like Claude) with secure, controlled access to operating system probes and operational control capabilities.

## 🎯 Vision

Turn your LLM into a full-powered ops assistant — from system monitoring to container troubleshooting to controlled operational actions.

## 💡 Why mcp-sysbox?

Giving an AI agent unrestricted shell access is dangerous. **mcp-sysbox** takes a different approach: it exposes only explicit, purpose-built tools through the [Model Context Protocol](https://modelcontextprotocol.io/), so your AI assistant can observe system state and troubleshoot containers **without the risk of arbitrary command execution**.

**Who is it for?**

- **AI/LLM developers** — Connect Claude or any MCP client to real infrastructure data for smarter, context-aware responses.
- **DevOps & SREs** — Let AI assist with diagnostics: CPU spikes, memory pressure, disk usage, container status, and log analysis.
- **Self-hosted / Homelab users** — Monitor your Docker host, NAS, or VPS through natural language conversations.
- **MCP ecosystem builders** — Use it as a reference implementation or extend it with your own system operations tools.

**Design principles:**

- 🔒 **Secure by default** — Read-only probes first; dangerous operations require guard middleware + audit logging (Phase 4 roadmap).
- 🔌 **Pluggable** — Each tool is an independent module. Add, remove, or customize tools without touching the core.
- 🌍 **Cross-platform** — Works on Linux, macOS, and Windows thanks to [gopsutil](https://github.com/shirou/gopsutil) and the official Docker Engine SDK.

## ✨ Features

> 🚧 Under active development

| Category | Capabilities | Status |
|----------|-------------|--------|
| **Connectivity** | `ping` health check tool | ✅ Done |
| **System Probes** | `system_memory` memory usage monitoring | ✅ Done |
| **System Probes** | `system_cpu` CPU usage monitoring | ✅ Done |
| **System Probes** | `system_disk` disk space monitoring | ✅ Done |
| **Docker Observability** | `docker_list` container listing | ✅ Done |
| **Docker Observability** | `docker_inspect` container details | ✅ Done |
| **Docker Observability** | `docker_logs` container log reading | ✅ Done |
| **Docker Observability** | `docker_stats` container resource usage | 🔜 Next |
| **Security & Audit** | YAML config center, audit logging, guard middleware | Planned |
| **Ops Control** | Container restart (guard-gated) | Planned |
| **Advanced MCP** | SSE transport, Resources & Prompts | Planned |

## 🏗️ Tech Stack

- **Language:** Go 1.25+
- **Protocol:** [Model Context Protocol (MCP)](https://modelcontextprotocol.io/)
- **MCP SDK:** [mcp-go](https://github.com/mark3labs/mcp-go)
- **System Info:** [gopsutil](https://github.com/shirou/gopsutil) v4 (cross-platform)
- **Docker:** [moby/moby](https://github.com/moby/moby) SDK (official Docker Engine API)
- **Transport:** Stdio
- **CI/CD:** GitHub Actions (multi-platform test + release)

## 📁 Project Structure

```
mcp-sysbox/
├── main.go                        # Entry point, server bootstrap
├── tools/                         # MCP tool definitions & handlers
│   ├── ping.go                    # ping health-check tool
│   ├── memory.go                  # system_memory probe tool
│   ├── cpu.go                     # system_cpu probe tool
│   ├── disk.go                    # system_disk probe tool
│   ├── docker_list.go             # docker_list container listing
│   ├── docker_inspect.go          # docker_inspect container details
│   ├── docker_logs.go             # docker_logs log reading
│   └── *_test.go                  # Unit tests for each tool
├── internal/
│   ├── sysinfo/                   # Cross-platform system info layer
│   │   ├── memory.go              # Memory status (gopsutil)
│   │   ├── cpu.go                 # CPU status (gopsutil)
│   │   └── disk.go                # Disk partition info (gopsutil)
│   └── docker/                    # Docker Engine API layer
│       ├── client.go              # Docker client singleton
│       ├── container.go           # Container listing
│       ├── inspect.go             # Container inspection
│       └── logs.go                # Container log retrieval
├── .github/
│   ├── workflows/
│   │   ├── ci.yml                 # CI: test + vet (3 platforms)
│   │   └── release.yml            # Release: cross-compile + publish
│   └── dependabot.yml             # Automated dependency updates
├── go.mod
└── go.sum
```

## 🚀 Quick Start

### Build from source

```bash
git clone https://github.com/volcano6/mcp-sysbox.git
cd mcp-sysbox
go build -o mcp-sysbox .
```

### Configure with Claude Desktop

Add the following to your Claude Desktop config file (`claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "mcp-sysbox": {
      "command": "/absolute/path/to/mcp-sysbox"
    }
  }
}
```

### Available Tools

| Tool | Description |
|------|-------------|
| `ping` | Health check — returns `pong` with server metadata (version, Go runtime, OS/Arch, timestamp) |
| `system_memory` | Memory probe — returns total, used, available memory and usage percentage |
| `system_cpu` | CPU probe — returns model name, core counts (physical/logical), overall and per-core usage |
| `system_disk` | Disk probe — returns all mounted partitions with total, used, free space and usage percentage |
| `docker_list` | List all Docker containers (running + stopped) with state, image, ports |
| `docker_inspect` | Inspect a container by name/ID — state, ports, mounts, networks, env (sensitive values masked) |
| `docker_logs` | Read container logs with `tail`, `since`, `timestamps` parameters |

## 🚀 Roadmap

- [x] **Phase 1: Project Scaffold & Infrastructure** ✅
  - [x] Open source essentials (README, LICENSE, .gitignore)
  - [x] Go module init + mcp-go SDK integration
  - [x] MCP Server skeleton with Stdio transport
  - [x] `ping` → `pong` connectivity test tool
- [x] **Phase 2: Read-only System Probes** ✅
  - [x] `system_memory` — memory usage monitoring
  - [x] `system_cpu` — CPU usage monitoring
  - [x] `system_disk` — disk space monitoring
- [ ] **Phase 3: Docker Container Observability** 🚧
  - [x] `docker_list` — container listing (running + stopped)
  - [x] `docker_inspect` — container details (state, ports, mounts, networks, env)
  - [x] `docker_logs` — container log reading (tail, since, timestamps)
  - [ ] `docker_stats` — container resource usage (CPU, memory, network I/O)
- [ ] **Phase 4: MCP Guard Security & Audit** — Config center, audit logger, guard middleware, controlled ops
- [ ] **Phase 5: Advanced MCP** — SSE transport, MCP Resources & Prompts, GitHub Actions CI

## 📄 License

This project is licensed under the [MIT License](LICENSE).
