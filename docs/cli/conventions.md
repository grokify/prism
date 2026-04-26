# CLI Conventions

This document defines the conventions for the PRISM CLI.

## Output Format Flag

### Display Commands (stdout)

Commands that display data to stdout use a single format flag:

```
-f, --format <type>    Output format (default: text)
```

**Supported formats:**

| Format | Description |
|--------|-------------|
| `text` | Human-readable tabular output (default) |
| `json` | JSON for scripting and piping |
| `markdown` | Markdown tables |
| `toon` | Token-optimized notation for LLM consumption |

**Examples:**

```bash
prism goal list prism.json              # text output (default)
prism goal list prism.json -f json      # JSON output
prism goal list prism.json -f markdown  # Markdown output
prism goal show prism.json goal-1 -f json
prism phase list prism.json -f toon
```

### Export Commands (files)

Commands that generate file artifacts support multiple formats:

```
-f, --format <type>    Output format(s), can be repeated or comma-separated
-o, --output-dir <dir> Output directory (required for export)
```

**Examples:**

```bash
# Single format export
prism export goals prism.json -o ./out/ -f json

# Multiple formats (repeated flag)
prism export goals prism.json -o ./out/ -f json -f markdown-pandoc -f markdown-marp

# Multiple formats (comma-separated)
prism export goals prism.json -o ./out/ -f json,markdown-pandoc,markdown-marp

# Export phases
prism export phases prism.json -o ./out/ -f json,markdown

# Export roadmap
prism export roadmap prism.json -o ./out/ -f json,markdown-marp
```

**Export-specific formats:**

| Format | Description |
|--------|-------------|
| `json` | JSON data file |
| `markdown-pandoc` | Markdown optimized for Pandoc processing |
| `markdown-marp` | Markdown optimized for Marp presentations |
| `yaml` | YAML data file |
| `csv` | CSV for spreadsheet import |

## Command Structure

### Subcommand Pattern

Commands follow a noun-verb pattern:

```
prism <resource> <action> [args] [flags]
```

**Resources:** `goal`, `phase`, `roadmap`, `metric`, `team`, `service`, `layer`

**Actions:**

| Action | Purpose | Output |
|--------|---------|--------|
| `list` | List all resources | stdout |
| `show` | Show single resource details | stdout |
| `progress` | Show progress/metrics | stdout |
| `metrics` | Show detailed metrics | stdout |
| `export` | Generate file artifacts | files |

### Common Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--format` | `-f` | Output format |
| `--output-dir` | `-o` | Output directory (export only) |
| `--help` | `-h` | Show help |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |
| 3 | File not found |
| 4 | Validation error |
