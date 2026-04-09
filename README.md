# config-loader

[![Build Status](https://img.shields.io/github/actions/workflow/status/user/config-loader/ci.yml?branch=main)]()
[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)]()
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

**config-loader** universal config loader supporting YAML, TOML, JSON, and env vars. Built with simplicity and performance in mind.

## Features

- **Graceful Shutdown** — Clean shutdown handling with configurable drain timeout
- **Structured Logging** — Built-in structured logging with slog compatibility
- **Minimal Allocations** — Designed to minimize GC pressure in hot paths

## Installation

```bash
go get github.com/user/config-loader@latest
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/user/config-loader"
)

func main() {
	client := configloader.New(
		configloader.WithTimeout(30 * time.Second),
	)

	result, err := client.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
```

## Configuration

Configuration can be provided via environment variables, a config file, or programmatically.

| Variable | Description | Default |
|----------|-------------|--------|
| `CONFIG_LOADER_TIMEOUT` | Request timeout in seconds | `30` |
| `CONFIG_LOADER_RETRIES` | Number of retry attempts | `3` |
| `CONFIG_LOADER_LOG_LEVEL` | Log verbosity (debug, info, warn, error) | `info` |

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
