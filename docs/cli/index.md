# CLI Reference

The PRISM CLI provides commands for creating, validating, and scoring PRISM documents.

## Installation

```bash
go install github.com/grokify/prism/cmd/prism@latest
```

## Commands

| Command | Description |
|---------|-------------|
| [`prism init`](init.md) | Initialize a new PRISM document |
| [`prism validate`](validate.md) | Validate a PRISM document |
| [`prism score`](score.md) | Calculate the PRISM score |
| [`prism catalog`](catalog.md) | List available constants |

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
