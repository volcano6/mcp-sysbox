# mcp-sysbox 🛠️

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![CI](https://github.com/volcano6/mcp-sysbox/actions/workflows/ci.yml/badge.svg)](https://github.com/volcano6/mcp-sysbox/actions/workflows/ci.yml)

**mcp-sysbox** (Go MCP Ops Server) is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server built in Go. It provides LLMs (like Claude) with secure, controlled access to operating system probes and operational control capabilities.

## 🎯 Vision

Turn your LLM into a full-powered ops assistant — from system monitoring to container troubleshooting to controlled operational actions.

## ✨ Features

> 🚧 Under active development

| Category | Capabilities | Status |
|----------|-------------|--------|
| **Connectivity** | `ping` health check tool | ✅ Done |
| **System Probes** | `system_memory` memory usage monitoring | ✅ Done |
| **System Probes** | `system_cpu` CPU usage monitoring | ✅ Done |
| **System Probes** | `system_disk` disk space monitoring | ✅ Done |
| **Docker Observability** | Container listing, log retrieval | Planned |
| **Ops Control** | Container restart (whitelist-gated), port checking | Planned |
| **Security** | YAML-based whitelist, controlled execution | Planned |

## 🏗️ Tech Stack

- **Language:** Go 1.21+
- **Protocol:** [Model Context Protocol (MCP)](https://modelcontextprotocol.io/)
- **MCP SDK:** [mcp-go](https://github.com/mark3labs/mcp-go) v0.56.0
- **System Info:** [gopsutil](https://github.com/shirou/gopsutil) v4 (cross-platform)
- **Transport:** Stdio

## 📁 Project Structure

```
mcp-sysbox/
├── main.go                        # Entry point, server bootstrap
├── tools/                         # MCP tool definitions & handlers
│   ├── ping.go                    # ping health-check tool
│   ├── ping_test.go
│   ├── memory.go                  # system_memory probe tool
│   ├── memory_test.go
│   ├── cpu.go                     # system_cpu probe tool
│   ├── cpu_test.go
│   ├── disk.go                    # system_disk probe tool
│   └── disk_test.go
├── internal/
│   └── sysinfo/                   # Cross-platform system info layer
│       ├── memory.go              # Memory status retrieval
│       ├── memory_test.go
│       ├── cpu.go                 # CPU status retrieval
│       ├── cpu_test.go
│       ├── disk.go                # Disk partition retrieval
│       └── disk_test.go
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
- [ ] **Phase 3: Docker Container Observability** — Container listing, inspect, logs, stats
- [ ] **Phase 4: MCP Guard Security & Audit** — Config center, audit logger, guard middleware, controlled ops
- [ ] **Phase 5: Advanced MCP** — SSE transport, MCP Resources & Prompts, GitHub Actions CI

## 📄 License

This project is licensed under the [MIT License](LICENSE).
