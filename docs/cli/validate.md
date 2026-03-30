# prism validate

Validate a PRISM document for structural and semantic correctness.

## Synopsis

```bash
prism validate <file>
```

## Description

Validates a PRISM document by checking:

- JSON syntax
- Required fields presence
- Valid constant values (domains, stages, categories, metric types)
- SLO operator validity
- Threshold consistency
- Maturity level ranges (1-5)

## Arguments

| Argument | Description |
|----------|-------------|
| `file` | Path to the PRISM JSON document |

## Examples

### Validate a Document

```bash
prism validate prism.json
```

Success output:

```
✓ prism.json is valid
```

### Validation Errors

If the document has errors:

```
✗ prism.json has 2 validation errors:
  - metrics[0].domain: invalid domain "sec" (valid: security, operations)
  - metrics[2].stage: invalid stage "deployment" (valid: design, build, test, runtime, response)
```

## Validation Rules

### Required Fields

Each metric must have:

- `id` or `name`
- `domain`
- `stage`
- `category`
- `metricType`

### Valid Constants

| Field | Valid Values |
|-------|--------------|
| `domain` | `security`, `operations` |
| `stage` | `design`, `build`, `test`, `runtime`, `response` |
| `category` | `prevention`, `detection`, `response`, `reliability`, `efficiency`, `quality` |
| `metricType` | `coverage`, `rate`, `latency`, `ratio`, `count`, `distribution`, `score` |
| `trendDirection` | `higher_better`, `lower_better`, `target_value` |
| `slo.operator` | `gte`, `lte`, `gt`, `lt`, `eq` |

### Maturity Levels

- `currentLevel` must be between 1 and 5
- `targetLevel` must be between 1 and 5
- `targetLevel` should be >= `currentLevel` (warning)

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Document is valid |
| 1 | Validation errors found |
| 2 | File not found or unreadable |
