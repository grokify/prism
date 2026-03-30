# Thresholds

Thresholds define the boundaries for metric status calculation, enabling traffic-light style health indicators.

## Threshold Structure

```json
{
  "thresholds": {
    "green": 99.95,
    "yellow": 99.9,
    "red": 99.0
  }
}
```

### Fields

| Field | Type | Description |
|-------|------|-------------|
| `green` | number | Threshold for "healthy" status |
| `yellow` | number | Threshold for "warning" status |
| `red` | number | Threshold for "critical" status |

## Status Calculation

Status is calculated based on the metric's trend direction.

### Higher Better (e.g., Availability)

For metrics where higher values are better:

| Current Value | Status |
|---------------|--------|
| `>= green` | Green |
| `>= yellow` | Yellow |
| `< yellow` | Red |

Example:

```json
{
  "current": 99.92,
  "trendDirection": "higher_better",
  "thresholds": {
    "green": 99.95,
    "yellow": 99.9,
    "red": 99.0
  }
}
// Status: Yellow (99.92 >= 99.9 but < 99.95)
```

### Lower Better (e.g., Latency)

For metrics where lower values are better:

| Current Value | Status |
|---------------|--------|
| `<= green` | Green |
| `<= yellow` | Yellow |
| `> yellow` | Red |

Example:

```json
{
  "current": 180,
  "trendDirection": "lower_better",
  "thresholds": {
    "green": 100,
    "yellow": 200,
    "red": 500
  }
}
// Status: Yellow (180 <= 200 but > 100)
```

## Programmatic Status Calculation

Use `CalculateStatus()` to compute status in code:

```go
metric := prism.Metric{
    Current:        99.92,
    TrendDirection: prism.TrendHigherBetter,
    Thresholds: &prism.Thresholds{
        Green:  99.95,
        Yellow: 99.9,
        Red:    99.0,
    },
}

status := metric.CalculateStatus()
fmt.Println(status) // "Yellow"
```

## Status Values

| Status | Constant | Color | Meaning |
|--------|----------|-------|---------|
| Green | `green` | 🟢 | Healthy, meeting targets |
| Yellow | `yellow` | 🟡 | Warning, needs attention |
| Red | `red` | 🔴 | Critical, requires action |

## Examples by Metric Type

### Availability (Higher Better)

```json
{
  "id": "ops-availability",
  "name": "Service Availability",
  "metricType": "rate",
  "trendDirection": "higher_better",
  "unit": "%",
  "current": 99.95,
  "thresholds": {
    "green": 99.95,
    "yellow": 99.9,
    "red": 99.0
  }
}
```

### Latency (Lower Better)

```json
{
  "id": "ops-p99-latency",
  "name": "P99 Latency",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "ms",
  "current": 250,
  "thresholds": {
    "green": 200,
    "yellow": 500,
    "red": 1000
  }
}
```

### Error Rate (Lower Better)

```json
{
  "id": "ops-error-rate",
  "name": "Error Rate",
  "metricType": "rate",
  "trendDirection": "lower_better",
  "unit": "%",
  "current": 0.15,
  "thresholds": {
    "green": 0.1,
    "yellow": 0.5,
    "red": 1.0
  }
}
```

### MTTR (Lower Better)

```json
{
  "id": "sec-vuln-mttr",
  "name": "Vulnerability MTTR",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "days",
  "current": 7,
  "thresholds": {
    "green": 3,
    "yellow": 7,
    "red": 14
  }
}
```

## Threshold vs. Target vs. SLO

| Concept | Purpose | Use Case |
|---------|---------|----------|
| **Target** | Aspirational goal | Planning, roadmaps |
| **SLO** | Contractual commitment | SLA compliance |
| **Thresholds** | Status visualization | Dashboards, alerts |

These can be different:

```json
{
  "current": 99.95,
  "target": 99.99,
  "slo": {
    "target": ">=99.9%",
    "operator": "gte",
    "value": 99.9
  },
  "thresholds": {
    "green": 99.95,
    "yellow": 99.9,
    "red": 99.0
  }
}
```

In this example:

- **Target**: 99.99% (aspirational)
- **SLO**: 99.9% (minimum commitment)
- **Green threshold**: 99.95% (above SLO, approaching target)
