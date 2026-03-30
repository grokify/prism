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
    "description": "Security and operations health metrics",
    "owner": "Platform Team",
    "lastUpdated": "2024-01-15"
  }
}
```

## Step 2: Define Security Metrics

Add security metrics covering the software delivery lifecycle:

### Design Stage - Threat Modeling

```json
{
  "id": "sec-threat-modeling",
  "name": "Threat Modeling Coverage",
  "description": "Percentage of new features with completed threat models",
  "domain": "security",
  "stage": "design",
  "category": "prevention",
  "metricType": "coverage",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 25,
  "current": 75,
  "target": 100,
  "thresholds": {
    "green": 90,
    "yellow": 70,
    "red": 50
  },
  "slo": {
    "target": ">=90%",
    "operator": "gte",
    "value": 90,
    "window": "quarterly"
  }
}
```

### Build Stage - SAST Coverage

```json
{
  "id": "sec-sast-coverage",
  "name": "SAST Coverage",
  "description": "Percentage of repositories with SAST scanning enabled",
  "domain": "security",
  "stage": "build",
  "category": "prevention",
  "metricType": "coverage",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 60,
  "current": 95,
  "target": 100,
  "slo": {
    "target": ">=95%",
    "operator": "gte",
    "value": 95
  }
}
```

### Response Stage - Vulnerability MTTR

```json
{
  "id": "sec-vuln-mttr",
  "name": "Vulnerability MTTR",
  "description": "Mean time to remediate critical vulnerabilities",
  "domain": "security",
  "stage": "response",
  "category": "response",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "days",
  "baseline": 30,
  "current": 7,
  "target": 3,
  "thresholds": {
    "green": 7,
    "yellow": 14,
    "red": 30
  },
  "slo": {
    "target": "<=7 days",
    "operator": "lte",
    "value": 7,
    "window": "30d"
  }
}
```

## Step 3: Define Operations Metrics

Add operations metrics aligned with DORA:

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

## Step 4: Define Maturity Levels

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
        "domain": "security",
        "stage": "design",
        "currentLevel": 3,
        "targetLevel": 4,
        "primaryKPI": "sec-threat-modeling"
      },
      {
        "domain": "security",
        "stage": "build",
        "currentLevel": 4,
        "targetLevel": 5,
        "primaryKPI": "sec-sast-coverage"
      }
    ]
  }
}
```

## Step 5: Validate and Score

Save your document and run:

```bash
# Validate
prism validate prism.json

# Score
prism score prism.json --detailed
```

## Complete Example

See the [examples directory](https://github.com/grokify/prism/tree/main/examples) for complete working documents:

- `security-metrics.json` - Security-focused metrics
- `operations-metrics.json` - DORA-aligned operations metrics
