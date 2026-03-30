# SLIs & SLOs

PRISM supports Service Level Indicators (SLIs) and Service Level Objectives (SLOs) following SRE best practices.

## SLI (Service Level Indicator)

An SLI defines what is being measured.

### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | No | SLI name |
| `description` | string | No | SLI description |
| `formula` | string | No | Calculation formula |

### Example

```json
{
  "sli": {
    "name": "Availability",
    "description": "Percentage of successful requests",
    "formula": "successful_requests / total_requests * 100"
  }
}
```

## SLO (Service Level Objective)

An SLO defines the target for an SLI.

### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `target` | string | Yes | Human-readable target |
| `operator` | string | No | Comparison operator for machine evaluation |
| `value` | number | No | Numeric target value |
| `window` | string | No | Measurement window |
| `thresholds` | object | No | Additional thresholds |

### SLO Operators

PRISM supports machine-evaluable SLOs with these operators:

| Operator | Constant | Description | Example |
|----------|----------|-------------|---------|
| `gte` | `>=` | Greater than or equal | Availability ≥99.99% |
| `lte` | `<=` | Less than or equal | Latency ≤200ms |
| `gt` | `>` | Greater than | Score >80 |
| `lt` | `<` | Less than | Error rate <0.1% |
| `eq` | `=` | Equal to | Target exactly 100 |

### Example with Machine-Evaluable SLO

```json
{
  "slo": {
    "target": ">=99.99%",
    "operator": "gte",
    "value": 99.99,
    "window": "30d"
  }
}
```

### Measurement Windows

Common window values:

| Window | Description |
|--------|-------------|
| `7d` | 7-day rolling window |
| `30d` | 30-day rolling window |
| `90d` | 90-day rolling window |
| `monthly` | Calendar month |
| `quarterly` | Calendar quarter |

## Programmatic SLO Checking

PRISM provides a `MeetsSLO()` method for programmatic checking:

```go
metric := prism.Metric{
    Current: 99.95,
    SLO: &prism.SLO{
        Target:   ">=99.99%",
        Operator: prism.SLOOperatorGTE,
        Value:    99.99,
    },
}

if metric.MeetsSLO() {
    fmt.Println("SLO met!")
} else {
    fmt.Println("SLO not met")
}
```

### Operator Behavior

| Operator | Current | Value | MeetsSLO() |
|----------|---------|-------|------------|
| `gte` | 99.99 | 99.99 | true |
| `gte` | 99.95 | 99.99 | false |
| `lte` | 200 | 250 | true |
| `lte` | 300 | 250 | false |
| `eq` | 100 | 100 | true |

## Complete Metric Example

```json
{
  "id": "ops-availability",
  "name": "Service Availability",
  "description": "Percentage of time the service is available",
  "domain": "operations",
  "stage": "runtime",
  "category": "reliability",
  "metricType": "rate",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 99.0,
  "current": 99.95,
  "target": 99.99,
  "sli": {
    "name": "Availability",
    "description": "Successful requests / total requests",
    "formula": "1 - (error_count / total_requests)"
  },
  "slo": {
    "target": ">=99.99%",
    "operator": "gte",
    "value": 99.99,
    "window": "30d"
  },
  "thresholds": {
    "green": 99.95,
    "yellow": 99.9,
    "red": 99.0
  },
  "frameworkMappings": [
    {"framework": "SRE", "reference": "availability-slo"},
    {"framework": "DORA", "reference": "availability"}
  ]
}
```

## Best Practices

1. **Set Realistic Targets** - SLOs should be achievable but challenging
2. **Include Error Budgets** - Use thresholds to define acceptable ranges
3. **Document Formulas** - Include SLI formulas for clarity
4. **Use Machine-Evaluable Operators** - Enable automated SLO checking
5. **Define Measurement Windows** - Clarify the evaluation period
6. **Map to Frameworks** - Reference industry standards (SRE, DORA)
