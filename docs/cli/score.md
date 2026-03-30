# prism score

Calculate the composite PRISM score for a document.

## Synopsis

```bash
prism score <file> [flags]
```

## Description

Calculates the PRISM score by:

1. Computing maturity scores for each domain/stage cell
2. Computing performance scores from metric progress
3. Weighting cells by domain and stage importance
4. Applying awareness multiplier (if customer awareness data present)
5. Producing a final score between 0.0 and 1.0

## Arguments

| Argument | Description |
|----------|-------------|
| `file` | Path to the PRISM JSON document |

## Flags

| Flag | Description |
|------|-------------|
| `--detailed` | Show detailed breakdown by domain and stage |
| `--json` | Output results as JSON |

## Examples

### Basic Score

```bash
prism score prism.json
```

Output:

```
PRISM Score: 72.5% (Strong)

Security:   68.0%
Operations: 77.0%
```

### Detailed Breakdown

```bash
prism score prism.json --detailed
```

Output:

```
PRISM Score: 72.5% (Strong)

By Domain:
  Security:   68.0%
  Operations: 77.0%

By Stage:
  Design:   65.0%
  Build:    80.0%
  Test:     70.0%
  Runtime:  75.0%
  Response: 72.5%

Component Averages:
  Maturity:    3.2 / 5.0 (64.0%)
  Performance: 78.3%

Health Status: Yellow
  System health needs attention in some areas
```

### JSON Output

```bash
prism score prism.json --json
```

Output:

```json
{
  "overall": 0.725,
  "baseScore": 0.725,
  "awarenessScore": 1.0,
  "securityScore": 0.68,
  "operationsScore": 0.77,
  "interpretation": "Strong",
  "maturityAverage": 0.64,
  "performanceAverage": 0.783,
  "cellScores": [
    {
      "domain": "security",
      "stage": "design",
      "maturityScore": 0.6,
      "performanceScore": 0.75,
      "cellScore": 0.69,
      "weight": 0.075
    }
  ]
}
```

## Score Interpretation

| Score | Level | Description |
|-------|-------|-------------|
| ≥0.90 | Elite | Industry-leading practices |
| ≥0.75 | Strong | Well-managed, proactive |
| ≥0.50 | Medium | Adequate, room for improvement |
| ≥0.25 | Weak | Significant gaps |
| <0.25 | Critical | Immediate attention required |

## Default Weights

### Component Weights

- Maturity: 40%
- Performance: 60%

### Stage Weights

| Stage | Weight |
|-------|--------|
| Design | 15% |
| Build | 20% |
| Test | 15% |
| Runtime | 30% |
| Response | 20% |

### Domain Weights

| Domain | Weight |
|--------|--------|
| Security | 50% |
| Operations | 50% |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Score calculated successfully |
| 1 | Validation errors in document |
| 2 | File not found or unreadable |
