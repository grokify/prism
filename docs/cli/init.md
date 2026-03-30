# prism init

Initialize a new PRISM document with scaffold structure.

## Synopsis

```bash
prism init [flags]
```

## Description

Creates a new PRISM document with default structure including:

- Metadata section
- Sample metrics for selected domains
- Maturity model with level definitions
- Maturity cells for each domain/stage combination

## Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--output` | `-o` | `prism.json` | Output file path |
| `--domain` | `-d` | (all) | Domain filter: `security`, `operations`, or both |

## Examples

### Create Default Document

Create a document with both security and operations domains:

```bash
prism init
```

Output: `prism.json`

### Specify Output Path

```bash
prism init -o my-metrics.json
```

### Security-Only Document

Create a document with only security domain metrics and maturity cells:

```bash
prism init -d security -o security.json
```

### Operations-Only Document

Create a document with only operations domain metrics and maturity cells:

```bash
prism init -d operations -o ops.json
```

## Output Structure

The generated document includes:

```json
{
  "$schema": "https://github.com/grokify/prism/schema/prism.schema.json",
  "metadata": {
    "name": "PRISM Document",
    "version": "1.0.0"
  },
  "metrics": [
    // Sample metrics for selected domains
  ],
  "maturity": {
    "levels": [
      // 5-level maturity definitions
    ],
    "cells": [
      // Maturity cells for selected domains × all stages
    ]
  }
}
```

## Notes

- If the output file already exists, it will be overwritten
- Use `prism validate` after editing to check for errors
- Maturity cells are filtered to only include selected domains
