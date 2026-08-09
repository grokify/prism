# prism prioritize

Rank initiatives by a configurable, multi-dimensional scoring model.

```bash
prism prioritize [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--dir` | Ecosystem directory to load |
| `--config` | Path to scoring weights JSON file |
| `--pillar` | Filter by pillar: `CSAT`, `SAM-SOM`, `TAM`, or `all` (default) |
| `--format` | Output format: `table` (default) or `json` |
| `--top` | Limit to top N results (`0` = all) |

**Examples:**

```bash
# Rank all initiatives as a table
prism prioritize --dir ./ecosystem

# Top 10 initiatives for the CSAT pillar, as JSON
prism prioritize --dir ./ecosystem --pillar CSAT --top 10 --format json

# Use custom scoring weights
prism prioritize --dir ./ecosystem --config weights.json
```

## Scoring Configuration

Scoring weights can be supplied as a JSON file via `--config`. Built-in presets
are available in the ecosystem library (`DefaultScoringConfig`,
`CustomerFocusedConfig`, `GrowthFocusedConfig`) and can be saved to disk as a
starting point. See the [Ecosystem API guide](../guide/ecosystem-api.md) for
details on the scoring dimensions and how evidence enriches scores.

## See Also

- [`prism gaps`](gaps.md) - Identify the gaps initiatives address
- [`prism dashboard`](dashboard.md) - Portfolio dashboard with ranked initiatives
