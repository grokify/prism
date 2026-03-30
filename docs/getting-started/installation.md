# Installation

## CLI Installation

Install the PRISM CLI using Go:

```bash
go install github.com/grokify/prism/cmd/prism@latest
```

Verify the installation:

```bash
prism --version
```

## Library Installation

Add PRISM as a Go module dependency:

```bash
go get github.com/grokify/prism
```

## Requirements

- Go 1.24 or later
- Git (for installation via `go install`)

## Building from Source

Clone the repository and build:

```bash
git clone https://github.com/grokify/prism.git
cd prism
go build -o prism ./cmd/prism
```

Install locally:

```bash
go install ./cmd/prism
```

## Editor Support

PRISM provides a JSON Schema for editor validation:

### VS Code

Add to your JSON file:

```json
{
  "$schema": "https://github.com/grokify/prism/schema/prism.schema.json"
}
```

Or configure in `.vscode/settings.json`:

```json
{
  "json.schemas": [
    {
      "fileMatch": ["*prism*.json", "*metrics*.json"],
      "url": "https://raw.githubusercontent.com/grokify/prism/main/schema/prism.schema.json"
    }
  ]
}
```

### JetBrains IDEs

1. Open Preferences → Languages & Frameworks → Schemas and DTDs → JSON Schema Mappings
2. Add a new mapping with the schema URL
3. Set file pattern to `*prism*.json`
