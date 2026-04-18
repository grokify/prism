# prism dashboard

Generate executive dashboards from a PRISM document.

## Synopsis

```bash
prism dashboard <file> [options]
```

## Description

Generate an executive dashboard showing high-level status views including goal progress, phase metrics, and SLO compliance summaries.

## Options

| Option | Description |
|--------|-------------|
| `-o`, `--output <file>` | Output file (default: stdout) |
| `--format <format>` | Output format: `json`, `markdown` (default: json) |
| `--title <title>` | Dashboard title |

## Examples

Generate dashboard JSON:

```bash
prism dashboard prism.json -o dashboard.json
```

Generate dashboard to stdout:

```bash
prism dashboard prism.json
```

Generate markdown dashboard:

```bash
prism dashboard prism.json --format markdown -o dashboard.md
```

## Output Format

The dashboard JSON includes:

```json
{
  "title": "Executive Dashboard",
  "generatedAt": "2026-04-18T10:00:00Z",
  "summary": {
    "prismScore": 0.78,
    "interpretation": "Strong",
    "goalsOnTrack": 2,
    "goalsTotal": 2,
    "currentPhase": "Q1 2026",
    "phaseProgress": 0.75
  },
  "goals": [
    {
      "id": "goal-reliability",
      "name": "Achieve High Reliability",
      "currentLevel": 3,
      "targetLevel": 4,
      "progress": 0.75,
      "sloCompliance": 0.67
    }
  ],
  "metrics": {
    "sloCompliance": 0.75,
    "initiativeCompletion": 0.50,
    "deploymentProgress": 0.80
  }
}
```

## Related Commands

### prism dashforge

Convert a PRISM document to dashforge format for dashboard generation.

```bash
prism dashforge <file> [options]
```

**Options:**

| Option | Description |
|--------|-------------|
| `-o`, `--output <file>` | Output file (default: stdout) |
| `--widgets` | Widget types to include: `all`, `metrics`, `goals`, `phases` |

**Examples:**

```bash
# Convert to dashforge format
prism dashforge prism.json -o dashforge.json

# Generate metrics widgets only
prism dashforge prism.json --widgets metrics
```

## See Also

- [prism report](report.md) - Roadmap reports
- [prism slo-report](slo-report.md) - SLO compliance reports
- [Dashforge Integration](../integrations/dashforge.md) - Dashforge integration guide
