# Installation

## Prerequisites

- Go 1.21 or later
- Git (for cloning repositories)

## Install from Source

```bash
go install github.com/grokify/prism/cmd/prism@latest
```

## Verify Installation

```bash
prism --version
```

## Install Individual Modules

If you only need specific functionality, you can install standalone CLIs:

```bash
# Capability stack CLI
go install github.com/grokify/prism-capability/cmd/capstack@latest

# Maturity model CLI
go install github.com/grokify/prism-maturity/cmd/prism@latest

# Roadmap/planning CLI
go install github.com/grokify/prism-roadmap/cmd/splan@latest
```

## Development Setup

For contributing or local development:

```bash
# Clone the ecosystem
mkdir -p ~/go/src/github.com/grokify
cd ~/go/src/github.com/grokify

git clone https://github.com/grokify/prism-core.git
git clone https://github.com/grokify/prism-capability.git
git clone https://github.com/grokify/prism-maturity.git
git clone https://github.com/grokify/prism-roadmap.git
git clone https://github.com/grokify/prism.git

# Build umbrella CLI
cd prism
go build ./cmd/prism
```

The umbrella repo uses `replace` directives in `go.mod` to reference local sibling directories during development.
