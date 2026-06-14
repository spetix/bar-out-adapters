---
description: "Instructions for AI agents working on the bar-out-adapters Go library project"
---

# Bar-out-adapters Project Instructions

This project is a Go library that adapts blocklet/widget output data to different bar formats (waybar, i3bar, JSON, raw text).

## Build and Test

- Use `make build` to compile for linux-amd64 and linux-arm64.
- Use `make test` to run tests with coverage.
- Use `make coverage` to generate coverage.xml.

## Architecture

- Public API in `pkg/barout/`: Factory function `New(protocol string)` returns `BlockletOutput` interface.
- Internal protocols in `internal/protocols/`: Implementations for json, raw, waybar.
- Data interface in `pkg/barout/data/`: Methods for short, long, label, colors.

## Conventions

- Package structure: Public in `pkg/`, internal in `internal/`.
- Interface-driven design.
- Factory pattern with fallback to "raw".
- Testing: Example-based tests with `//Output:` comments.

## Potential Pitfalls

- Test files in `internal/protocols/` are mostly commented-out templates; implement them.
- No trailing newline in JSON outputs; verify if intended.
- Error handling: Marshal errors are logged but not returned.

For more details, see [README.md](../README.md).