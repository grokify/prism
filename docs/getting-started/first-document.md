# Your First PRISM Document

This guide walks you through creating a comprehensive PRISM document from scratch.

## Document Structure

A PRISM document has the following top-level structure:

```json
{
  "$schema": "https://github.com/grokify/prism/schema/prism.schema.json",
  "metadata": { ... },
  "metrics": [ ... ],
  "maturity": { ... }
}
```

## Step 1: Metadata

Start with basic metadata:

```json
{
  "metadata": {
    "name": "Acme Corp PRISM",
    "version": "1.0.0",
    "description": "Operations health metrics",
    "owner": "Platform Team",
    "lastUpdated": "2024-01-15"
  }
}
```

## Step 2: Define Operations Metrics

Add operations metrics aligned with DORA and SRE practices:

### Runtime Stage - Availability

```json
{
  "id": "ops-availability",
  "name": "Service Availability",
  "description": "Percentage uptime over measurement window",
  "domain": "operations",
  "stage": "runtime",
  "category": "reliability",
  "metricType": "rate",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 99.0,
  "current": 99.95,
  "target": 99.99,
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

### Build Stage - Deployment Frequency

```json
{
  "id": "ops-deploy-frequency",
  "name": "Deployment Frequency",
  "description": "Number of deployments per day",
  "domain": "operations",
  "stage": "build",
  "category": "efficiency",
  "metricType": "count",
  "trendDirection": "higher_better",
  "unit": "/day",
  "baseline": 1,
  "current": 5,
  "target": 10,
  "slo": {
    "target": ">=5/day",
    "operator": "gte",
    "value": 5,
    "window": "7d"
  },
  "frameworkMappings": [
    {"framework": "DORA", "reference": "deployment-frequency"}
  ]
}
```

### Build Stage - Lead Time for Changes

```json
{
  "id": "ops-lead-time",
  "name": "Lead Time for Changes",
  "description": "Time from commit to production deployment",
  "domain": "operations",
  "stage": "build",
  "category": "efficiency",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "hours",
  "baseline": 168,
  "current": 24,
  "target": 1,
  "slo": {
    "target": "<=24h",
    "operator": "lte",
    "value": 24,
    "window": "7d"
  },
  "frameworkMappings": [
    {"framework": "DORA", "reference": "lead-time-for-changes"}
  ]
}
```

### Response Stage - Mean Time to Recovery

```json
{
  "id": "ops-mttr",
  "name": "Mean Time to Recovery",
  "description": "Average time to recover from production incidents",
  "domain": "operations",
  "stage": "response",
  "category": "response",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "hours",
  "baseline": 24,
  "current": 2,
  "target": 1,
  "slo": {
    "target": "<=1h",
    "operator": "lte",
    "value": 1,
    "window": "30d"
  },
  "frameworkMappings": [
    {"framework": "DORA", "reference": "time-to-restore-service"},
    {"framework": "SRE", "reference": "mttr"}
  ]
}
```

## Step 3: Define Maturity Levels

Set your current and target maturity for each domain/stage:

```json
{
  "maturity": {
    "levels": [
      {"level": 1, "name": "Reactive", "description": "Ad-hoc processes"},
      {"level": 2, "name": "Basic", "description": "Basic controls"},
      {"level": 3, "name": "Defined", "description": "Standardized processes"},
      {"level": 4, "name": "Managed", "description": "Measured and controlled"},
      {"level": 5, "name": "Optimizing", "description": "Continuous improvement"}
    ],
    "cells": [
      {
        "domain": "operations",
        "stage": "runtime",
        "currentLevel": 3,
        "targetLevel": 4,
        "primaryKPI": "ops-availability"
      },
      {
        "domain": "operations",
        "stage": "build",
        "currentLevel": 3,
        "targetLevel": 5,
        "primaryKPI": "ops-deploy-frequency"
      },
      {
        "domain": "operations",
        "stage": "response",
        "currentLevel": 4,
        "targetLevel": 5,
        "primaryKPI": "ops-mttr"
      }
    ]
  }
}
```

## Step 4: Validate and Score

Save your document and run:

```bash
# Validate
prism validate prism.json

# Score
prism score prism.json --detailed
```

## Complete Example

See the [examples directory](https://github.com/grokify/prism/tree/main/examples) for complete working documents:

- `operations-metrics.json` - DORA-aligned operations metrics
- `goal-roadmap.json` - Goal-driven maturity roadmap with phases

For security metrics examples, see [prism-security](https://github.com/grokify/prism-security).
