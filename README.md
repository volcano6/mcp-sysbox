# mcp-sysbox 🛠️

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)](https://go.dev)

**mcp-sysbox** (Go MCP Ops Server) is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server built in Go. It provides LLMs (like Claude) with secure, controlled access to operating system probes and operational control capabilities.

## 🎯 Vision

Turn your LLM into a full-powered ops assistant — from system monitoring to container troubleshooting to controlled operational actions.

## ✨ Features

> 🚧 Under active development

| Category | Capabilities | Status |
|----------|-------------|--------|
| **System Probes** | CPU usage, memory status, disk space | Planned |
| **Docker Observability** | Container listing, log retrieval | Planned |
| **Ops Control** | Container restart (whitelist-gated), port checking | Planned |
| **Security** | YAML-based whitelist, controlled execution | Planned |

## 🏗️ Tech Stack

- **Language:** Go 1.21+
- **Protocol:** [Model Context Protocol (MCP)](https://modelcontextprotocol.io/)
- **MCP SDK:** [mcp-go](https://github.com/mark3labs/mcp-go)
- **Transport:** Stdio

## 🚀 Roadmap

- [x] **Phase 1: Project Scaffold & Infrastructure** *(current)*
- [ ] **Phase 2: Read-only System Probes** — CPU, Memory, Disk monitoring
- [ ] **Phase 3: Docker Container Observability** — Status queries, log retrieval
- [ ] **Phase 4: Secure Execution & Ops Control** — Whitelist validation, port probing, service restart

## 📦 Quick Start

> Coming soon — stay tuned for Phase 1 completion!

## 📄 License

This project is licensed under the [MIT License](LICENSE).
