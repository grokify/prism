# Integration Guide

This guide covers how to integrate PRISM modules for end-to-end planning workflows.

## Document Flow

```
Capability Stack          What capabilities do we need?
       │
       ▼
Maturity Model + State    Where are we for each capability?
       │
       ▼
OKR/V2MOM + Roadmap       What goals and sequence to improve?
       │
       ▼
MRD / PRD / TRD           How do we execute each roadmap item?
```

## Step-by-Step Integration

### 1. Start with Capabilities

Define your capability stack first. This is the stable foundation:

```bash
# Create capability stack
prism capability init --domain operations --output ops-stack.json

# Edit to add layers and capabilities
# ...

# Validate
prism capability validate ops-stack.json
```

Key decisions:

- What layers organize your capabilities? (e.g., Observe, Respond, Recover)
- What capabilities exist at each layer?
- What are the dependencies between capabilities?

### 2. Define Maturity Model

Create a maturity model with SLIs that measure each capability:

```json
{
  "domains": {
    "operations": {
      "levels": [
        {
          "level": 1,
          "criteria": [
            {
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
      "unit": "%"
    }
  }
}
```

### 3. Link Capabilities to SLIs

Add `PRISMRef` to capabilities:

```json
{
  "id": "observability-platform",
  "name": "Observability Platform",
  "prismRef": {
    "sliIds": [
      "sli-monitoring-coverage",
      "sli-alerting-quality"
    ]
  }
}
```

### 4. Track Current State

Create a state document with current measurements:

```json
{
  "sliState": {
    "sli-availability": {
      "windows": {
        "30d": { "value": 99.5 }
      }
    }
  }
}
```

### 5. Generate Maturity Dashboard

View maturity with capability layer grouping:

```bash
prism maturity model dashboard model.json \
  --state state.json \
  --capstack ops-stack.json \
  --aggregation min \
  -o dashboard.html
```

### 6. Identify Gaps

The dashboard shows:

- Current maturity level per SLI
- Target level (M5)
- Capabilities with lowest maturity (improvement opportunities)

### 7. Create Goals

Based on gaps, create OKRs:

```json
{
  "objectives": [
    {
      "id": "obj-reliability",
      "title": "Improve platform reliability to M4",
      "keyResults": [
        {
          "description": "Achieve 99.9% availability",
          "target": 99.9,
          "current": 99.5
        }
      ]
    }
  ]
}
```

### 8. Build Roadmap

Sequence work items to achieve goals:

```json
{
  "phases": [
    {
      "id": "q1-2026",
      "name": "Q1 2026",
      "items": [
        {
          "id": "rmi-monitoring-upgrade",
          "title": "Upgrade monitoring coverage",
          "capabilityId": "observability-platform"
        }
      ]
    }
  ]
}
```

### 9. Create Requirements

For each roadmap item, create detailed specs:

```bash
# Market requirements
prism roadmap requirements mrd render rmi-monitoring-mrd.json -o mrd.md

# Product requirements
prism roadmap requirements prd render rmi-monitoring-prd.json -o prd.md

# Technical requirements
prism roadmap requirements trd render rmi-monitoring-trd.json -o trd.md
```

## Cross-Module Validation

Load the full ecosystem to validate cross-references:

```bash
prism ecosystem load --config prism.yaml
```

This validates:

- Capability → SLI references exist
- Roadmap item → Capability references exist
- All documents are structurally valid

## Automation

### CI/CD Integration

Add validation to your CI pipeline:

```yaml
# .github/workflows/prism-validate.yaml
name: PRISM Validation
on: [push, pull_request]
jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      - run: go install github.com/grokify/prism/cmd/prism@latest
      - run: prism capability validate capabilities/*.json
      - run: prism maturity model validate maturity/model.json
      - run: prism ecosystem load --config prism.yaml
```

### Dashboard Generation

Generate dashboards on state updates:

```bash
#!/bin/bash
# generate-dashboard.sh
prism maturity model dashboard \
  maturity/model.json \
  --state maturity/state.json \
  --capstack capabilities/ops-stack.json \
  --aggregation min \
  -o docs/dashboard.html
```

## Best Practices

1. **Version your documents** - Include version in metadata
2. **Validate on commit** - Use pre-commit hooks or CI
3. **Keep capabilities stable** - They change less often than state
4. **Snapshot state regularly** - Track maturity over time (e.g., `state-q1-2026.json`)
5. **Link everything** - Use `PRISMRef` and document references
