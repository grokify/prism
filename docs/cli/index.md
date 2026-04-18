# CLI Reference

The PRISM CLI provides commands for creating, validating, and scoring PRISM documents.

## Installation

```bash
go install github.com/grokify/prism/cmd/prism@latest
```

## Commands

### Core Commands

| Command | Description |
|---------|-------------|
| [`prism init`](init.md) | Initialize a new PRISM document |
| [`prism validate`](validate.md) | Validate a PRISM document |
| [`prism score`](score.md) | Calculate the PRISM score |
| [`prism catalog`](catalog.md) | List available constants |

### Roadmap Commands (v0.2.0)

| Command | Description |
|---------|-------------|
| [`prism goal`](goal.md) | Manage and inspect goals |
| [`prism phase`](phase.md) | Manage and inspect phases |
| [`prism roadmap`](roadmap.md) | View roadmap overview |

### Reporting Commands (v0.2.0)

| Command | Description |
|---------|-------------|
| `prism report` | Generate roadmap reports (Markdown/JSON) |
| `prism slo-report` | Generate SLO compliance reports |
| `prism dashboard` | Generate executive dashboards |
| `prism dashforge` | Convert to dashforge format |

## Global Flags

| Flag | Description |
|------|-------------|
| `--help`, `-h` | Show help for any command |
| `--version` | Show version information |

## Usage

```bash
# Get help
prism --help
prism <command> --help

# Initialize a document
prism init -o prism.json

# Validate a document
prism validate prism.json

# Calculate score
prism score prism.json

# List constants
prism catalog
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Validation errors or general failure |
| 2 | File not found or I/O error |
