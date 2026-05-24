# CLI Reference

The `prism` CLI provides a unified interface to the PRISM ecosystem.

## Command Structure

```
prism <module> <command> [subcommand] [flags]
```

## Modules

| Module | Description |
|--------|-------------|
| [`capability`](capability.md) | Capability stack operations |
| [`maturity`](maturity.md) | Maturity model and state operations |
| [`roadmap`](roadmap.md) | Goals and requirements operations |
| [`ecosystem`](ecosystem.md) | Cross-module ecosystem operations |
| [`site`](site.md) | Static site generation |

## Global Commands

### validate

Validate PRISM documents:

```bash
prism validate --all
prism validate document.json
```

### stats

Show repository statistics:

```bash
prism stats
```

## Common Flags

| Flag | Description |
|------|-------------|
| `--help`, `-h` | Show help for command |
| `--version` | Show version information |
| `--output`, `-o` | Output file path |
| `--format`, `-f` | Output format (json, html, md) |

## Examples

```bash
# Capability operations
prism capability list stack.json
prism capability validate stack.json

# Maturity operations
prism maturity model dashboard model.json --state state.json -o dash.html

# Roadmap operations
prism roadmap goals okr render goals.json -o goals.md

# Ecosystem operations
prism ecosystem load --config prism.yaml --json

# Site generation
prism site generate --stack=./stacks/ --output=./dist
prism site generate --stack=./security/ --stack=./reliability/ --title="Dashboard"
```
