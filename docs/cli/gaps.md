# prism gaps

Identify and analyze gaps across ecosystem documents.

## Commands

### analyze

Analyze coverage gaps across the loaded ecosystem:

```bash
prism gaps analyze [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--config`, `-c` | Configuration file (JSON) |
| `--dir`, `-d` | Directory to load from |
| `--json` | Output as JSON |
| `--sort-by` | Sort gaps by: `impact` (default), `severity`, or `type` |

**Examples:**

```bash
# Analyze gaps from a directory, sorted by impact
prism gaps analyze --dir ./ecosystem --sort-by=impact

# Analyze gaps from a config file
prism gaps analyze --config prism.yaml

# Output as JSON for tooling
prism gaps analyze --dir ./ecosystem --json
```

## See Also

- [`prism prioritize`](prioritize.md) - Rank initiatives that close gaps
- [`prism dashboard`](dashboard.md) - Visualize gaps in the portfolio dashboard
