# Metric Types

PRISM supports various metric types to represent different kinds of measurements.

## Available Metric Types

| Type | Constant | Description | Example |
|------|----------|-------------|---------|
| Coverage | `coverage` | Percentage of coverage | SAST coverage |
| Rate | `rate` | Frequency or percentage | Error rate, availability |
| Latency | `latency` | Time duration | P99 latency, MTTR |
| Ratio | `ratio` | Proportion | Success ratio |
| Count | `count` | Absolute count | Incident count |
| Distribution | `distribution` | Statistical distribution | Latency percentiles |
| Score | `score` | Composite score | Security score |

## Coverage

Coverage metrics measure the extent of coverage for a control or practice.

### Characteristics

- Range: 0-100%
- Trend direction: `higher_better`
- Common unit: `%`

### Examples

```json
{
  "id": "sec-sast-coverage",
  "name": "SAST Coverage",
  "metricType": "coverage",
  "trendDirection": "higher_better",
  "unit": "%",
  "current": 95,
  "target": 100
}
```

## Rate

Rate metrics measure frequency or percentage over time.

### Characteristics

- Can be percentages (0-100%) or counts per time period
- Trend direction varies by metric
- Common units: `%`, `per hour`, `per day`

### Examples

```json
{
  "id": "ops-availability",
  "name": "Service Availability",
  "metricType": "rate",
  "trendDirection": "higher_better",
  "unit": "%",
  "current": 99.95,
  "target": 99.99
}
```

```json
{
  "id": "ops-error-rate",
  "name": "Error Rate",
  "metricType": "rate",
  "trendDirection": "lower_better",
  "unit": "%",
  "current": 0.1,
  "target": 0.01
}
```

## Latency

Latency metrics measure time duration, typically for response times or remediation.

### Characteristics

- Trend direction: Usually `lower_better`
- Common units: `ms`, `seconds`, `minutes`, `hours`, `days`

### Examples

```json
{
  "id": "ops-p99-latency",
  "name": "P99 Latency",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "ms",
  "current": 250,
  "target": 200
}
```

```json
{
  "id": "sec-vuln-mttr",
  "name": "Vulnerability MTTR",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "days",
  "current": 7,
  "target": 3
}
```

## Ratio

Ratio metrics measure proportions, often as decimals or percentages.

### Characteristics

- Range: 0-1 (decimal) or 0-100 (percentage)
- Trend direction varies
- Common units: none, `%`

### Examples

```json
{
  "id": "ops-success-ratio",
  "name": "Request Success Ratio",
  "metricType": "ratio",
  "trendDirection": "higher_better",
  "current": 0.998,
  "target": 0.999
}
```

## Count

Count metrics measure absolute numbers.

### Characteristics

- Integer values
- Trend direction varies by metric
- No unit or custom unit

### Examples

```json
{
  "id": "sec-open-vulns",
  "name": "Open Critical Vulnerabilities",
  "metricType": "count",
  "trendDirection": "lower_better",
  "current": 3,
  "target": 0
}
```

```json
{
  "id": "ops-deploy-count",
  "name": "Monthly Deployments",
  "metricType": "count",
  "trendDirection": "higher_better",
  "current": 45,
  "target": 60
}
```

## Distribution

Distribution metrics capture statistical distributions, often with percentiles.

### Characteristics

- Multiple values (p50, p90, p99, etc.)
- Typically used with `dataPoints` array
- Trend direction usually `lower_better` for latency distributions

### Examples

```json
{
  "id": "ops-latency-distribution",
  "name": "Latency Distribution",
  "metricType": "distribution",
  "trendDirection": "lower_better",
  "unit": "ms",
  "dataPoints": [
    {"timestamp": "2024-01-01T00:00:00Z", "value": 50, "label": "p50"},
    {"timestamp": "2024-01-01T00:00:00Z", "value": 150, "label": "p90"},
    {"timestamp": "2024-01-01T00:00:00Z", "value": 250, "label": "p99"}
  ]
}
```

## Score

Score metrics represent composite or aggregated scores.

### Characteristics

- Range: Typically 0-100 or 0-1
- May be computed from other metrics
- Trend direction: `higher_better`

### Examples

```json
{
  "id": "sec-overall-score",
  "name": "Security Score",
  "metricType": "score",
  "trendDirection": "higher_better",
  "current": 78,
  "target": 90
}
```

## Trend Directions

| Direction | Constant | Description |
|-----------|----------|-------------|
| Higher Better | `higher_better` | Higher values indicate improvement |
| Lower Better | `lower_better` | Lower values indicate improvement |
| Target Value | `target_value` | Target a specific value |

### Choosing Trend Direction

| Metric Type | Typical Direction |
|-------------|-------------------|
| Coverage | `higher_better` |
| Availability | `higher_better` |
| Error Rate | `lower_better` |
| Latency | `lower_better` |
| MTTR | `lower_better` |
| Deployment Frequency | `higher_better` |
| Security Score | `higher_better` |
