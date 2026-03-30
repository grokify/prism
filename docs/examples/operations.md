# Operations Metrics Example

This example demonstrates operations-focused metrics aligned with DORA and SRE practices.

## Overview

The operations metrics example includes 8 metrics:

| Metric | Stage | Category | Framework |
|--------|-------|----------|-----------|
| Service Availability | Runtime | Reliability | SRE |
| P99 Latency | Runtime | Reliability | SRE |
| Error Rate | Runtime | Reliability | SRE |
| Deployment Frequency | Build | Efficiency | DORA |
| Lead Time for Changes | Build | Efficiency | DORA |
| Mean Time to Recovery | Response | Reliability | DORA |
| Change Failure Rate | Build | Quality | DORA |
| Infrastructure as Code Coverage | Build | Quality | - |

## File Location

```
examples/operations-metrics.json
```

## DORA Metrics

### Deployment Frequency

How often deployments occur.

```json
{
  "id": "ops-deploy-frequency",
  "name": "Deployment Frequency",
  "description": "Number of deployments per day",
  "domain": "operations",
  "stage": "build",
  "category": "efficiency",
  "metricType": "rate",
  "trendDirection": "higher_better",
  "unit": "deploys/day",
  "baseline": 1,
  "current": 5,
  "target": 10,
  "frameworkMappings": [
    {"framework": "DORA", "reference": "deployment-frequency"}
  ]
}
```

### Lead Time for Changes

Time from commit to production.

```json
{
  "id": "ops-lead-time",
  "name": "Lead Time for Changes",
  "description": "Time from code commit to production deployment",
  "domain": "operations",
  "stage": "build",
  "category": "efficiency",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "hours",
  "baseline": 168,
  "current": 24,
  "target": 1,
  "thresholds": {
    "green": 24,
    "yellow": 168,
    "red": 720
  },
  "frameworkMappings": [
    {"framework": "DORA", "reference": "lead-time"}
  ]
}
```

### Mean Time to Recovery

Time to recover from incidents.

```json
{
  "id": "ops-mttr",
  "name": "Mean Time to Recovery",
  "description": "Average time to recover from production incidents",
  "domain": "operations",
  "stage": "response",
  "category": "reliability",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "hours",
  "baseline": 24,
  "current": 2,
  "target": 1,
  "thresholds": {
    "green": 1,
    "yellow": 4,
    "red": 24
  },
  "slo": {
    "target": "<=4 hours",
    "operator": "lte",
    "value": 4,
    "window": "30d"
  },
  "frameworkMappings": [
    {"framework": "DORA", "reference": "mttr"},
    {"framework": "SRE", "reference": "recovery-slo"}
  ]
}
```

### Change Failure Rate

Percentage of deployments causing failures.

```json
{
  "id": "ops-change-failure-rate",
  "name": "Change Failure Rate",
  "description": "Percentage of deployments causing production failures",
  "domain": "operations",
  "stage": "build",
  "category": "quality",
  "metricType": "rate",
  "trendDirection": "lower_better",
  "unit": "%",
  "baseline": 25,
  "current": 8,
  "target": 5,
  "thresholds": {
    "green": 5,
    "yellow": 15,
    "red": 30
  },
  "frameworkMappings": [
    {"framework": "DORA", "reference": "change-failure-rate"}
  ]
}
```

## SRE Metrics

### Service Availability

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
  "thresholds": {
    "green": 99.95,
    "yellow": 99.9,
    "red": 99.0
  },
  "slo": {
    "target": ">=99.99%",
    "operator": "gte",
    "value": 99.99,
    "window": "30d"
  },
  "frameworkMappings": [
    {"framework": "SRE", "reference": "availability-slo"},
    {"framework": "DORA", "reference": "availability"}
  ]
}
```

### P99 Latency

```json
{
  "id": "ops-p99-latency",
  "name": "P99 Latency",
  "description": "99th percentile response latency",
  "domain": "operations",
  "stage": "runtime",
  "category": "reliability",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "ms",
  "baseline": 500,
  "current": 180,
  "target": 100,
  "thresholds": {
    "green": 100,
    "yellow": 250,
    "red": 500
  },
  "slo": {
    "target": "<=200ms",
    "operator": "lte",
    "value": 200,
    "window": "30d"
  },
  "frameworkMappings": [
    {"framework": "SRE", "reference": "latency-slo"}
  ]
}
```

### Error Rate

```json
{
  "id": "ops-error-rate",
  "name": "Error Rate",
  "description": "Percentage of requests resulting in errors",
  "domain": "operations",
  "stage": "runtime",
  "category": "reliability",
  "metricType": "rate",
  "trendDirection": "lower_better",
  "unit": "%",
  "baseline": 1.0,
  "current": 0.15,
  "target": 0.1,
  "thresholds": {
    "green": 0.1,
    "yellow": 0.5,
    "red": 1.0
  },
  "slo": {
    "target": "<=0.1%",
    "operator": "lte",
    "value": 0.1,
    "window": "30d"
  },
  "frameworkMappings": [
    {"framework": "SRE", "reference": "error-rate-slo"}
  ]
}
```

## Infrastructure Metric

### Infrastructure as Code Coverage

```json
{
  "id": "ops-iac-coverage",
  "name": "Infrastructure as Code Coverage",
  "description": "Percentage of infrastructure managed via IaC",
  "domain": "operations",
  "stage": "build",
  "category": "quality",
  "metricType": "coverage",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 40,
  "current": 85,
  "target": 100,
  "thresholds": {
    "green": 90,
    "yellow": 70,
    "red": 50
  }
}
```

## Usage

### Validate

```bash
prism validate examples/operations-metrics.json
```

### Score

```bash
prism score examples/operations-metrics.json --detailed
```

### Expected Output

```
PRISM Score: 78.2% (Strong)

Operations: 78.2%

By Stage:
  Build:    76.5%
  Runtime:  82.0%
  Response: 75.0%
```

## DORA Performance Levels

Based on the metrics in this example:

| Metric | Current | DORA Level |
|--------|---------|------------|
| Deployment Frequency | 5/day | High |
| Lead Time | 24 hours | High |
| MTTR | 2 hours | Elite |
| Change Failure Rate | 8% | Elite |
