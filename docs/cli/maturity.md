# prism maturity

Maturity model and state operations for measuring organizational maturity.

## Commands

### model dashboard

Generate an HTML dashboard from a maturity model:

```bash
prism maturity model dashboard <model-file> [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--state` | State document with current measurements |
| `--capstack` | Capability stack for layer-based views |
| `--aggregation` | Aggregation method: `min` (default), `avg` |
| `--format`, `-f` | Output format: `html` (default), `json` |
| `--output`, `-o` | Output file path |

**Examples:**

```bash
# Basic dashboard
prism maturity model dashboard model.json -o dashboard.html

# With state overlay
prism maturity model dashboard model.json \
  --state state.json \
  -o dashboard.html

# With capability layer views
prism maturity model dashboard model.json \
  --state state.json \
  --capstack capabilities.json \
  --aggregation min \
  -o dashboard.html
```

### model validate

Validate a maturity model specification:

```bash
prism maturity model validate <model-file>
```

Checks:

- Valid JSON structure
- Required fields present (domains, levels, criteria)
- SLI references resolve
- Level numbers sequential (1-5)
- Valid operators

### model lint

Lint a maturity model for common issues:

```bash
prism maturity model lint <model-file> [--strict]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--strict` | Treat warnings as errors |

Checks:

- Criteria without `sliId` (won't show in dashboard)
- Missing operator or target
- Unused SLI definitions
- Missing units
- Threshold coverage gaps

**Exit codes:**

- `0` - No issues
- `1` - Warnings found
- `2` - Errors found

### Other Commands

```bash
# Analyze maturity gaps
prism maturity analyze <document>

# Export to different formats
prism maturity export okr <document> --output goals.json
prism maturity export v2mom <document> --output v2mom.json

# Generate reports
prism maturity report <document> --format markdown
```

## Dashboard Views

When a capability stack is provided, the dashboard includes:

1. **Layer Maturity Overview** - Horizontal bar chart showing aggregate maturity per layer
2. **Layer Summary Cards** - Metric cards with capability counts and maturity levels
3. **Capability Bullets** - Bullet charts grouped by layer showing individual capability maturity

### Aggregation Methods

| Method | Description | Use Case |
|--------|-------------|----------|
| `min` | Minimum value across SLIs | Conservative; "you're only as strong as your weakest link" |
| `avg` | Average value across SLIs | Balanced view of overall maturity |

## Model Format

```json
{
  "$schema": "https://github.com/grokify/prism-maturity/schema/maturity-model.schema.json",
  "domains": {
    "operations": {
      "name": "Operations",
      "levels": [
        {
          "level": 1,
          "name": "Initial",
          "criteria": [
            {
              "id": "crit-availability-m1",
              "sliId": "sli-availability",
              "operator": "gte",
              "target": 95
            }
          ]
        }
      ]
    }
  },
  "slis": {
    "sli-availability": {
      "name": "Service Availability",
      "unit": "%",
      "sliType": "availability"
    }
  }
}
```
