# Examples Overview

PRISM includes example documents to help you get started quickly.

## Available Examples

| Example | Description | Metrics |
|---------|-------------|---------|
| [Operations Metrics](operations.md) | DORA-aligned operations metrics | 8 |
| [Goal Roadmap](goal-roadmap.md) | Goal-driven maturity roadmap with phases | 6 |

## Example Files

Example files are located in the `examples/` directory:

```
examples/
├── operations-metrics.json  # Operations-focused metrics
└── goal-roadmap.json        # Goal-driven maturity roadmap
```

## Using Examples

### View an Example

```bash
cat examples/operations-metrics.json
```

### Validate an Example

```bash
prism validate examples/operations-metrics.json
```

### Score an Example

```bash
prism score examples/operations-metrics.json --detailed
```

### Copy as Starting Point

```bash
cp examples/operations-metrics.json my-ops.json
# Edit my-ops.json with your values
prism validate my-ops.json
```

## Example Patterns

### Coverage Metric Pattern

```json
{
  "id": "coverage-metric",
  "name": "Coverage Metric",
  "metricType": "coverage",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 50,
  "current": 85,
  "target": 100,
  "thresholds": {
    "green": 90,
    "yellow": 70,
    "red": 50
  }
}
```

### Latency Metric Pattern

```json
{
  "id": "latency-metric",
  "name": "Response Time",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "ms",
  "baseline": 500,
  "current": 200,
  "target": 100,
  "thresholds": {
    "green": 100,
    "yellow": 250,
    "red": 500
  }
}
```

### Rate Metric Pattern

```json
{
  "id": "rate-metric",
  "name": "Availability",
  "metricType": "rate",
  "trendDirection": "higher_better",
  "unit": "%",
  "current": 99.95,
  "target": 99.99,
  "slo": {
    "target": ">=99.99%",
    "operator": "gte",
    "value": 99.99,
    "window": "30d"
  }
}
```
