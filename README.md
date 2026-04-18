# PRISM

**Proactive Reliability & Security Maturity Model**

PRISM is a unified framework for B2B SaaS health metrics that combines SLOs, DMAIC, OKRs, and maturity modeling into a single coherent system. It provides structured schemas for defining metrics, calculating composite health scores, and tracking organizational maturity across security and operations domains.

## Installation

```bash
go install github.com/grokify/prism/cmd/prism@latest
```

Or add as a library dependency:

```bash
go get github.com/grokify/prism
```

## CLI Usage

### Initialize a new PRISM document

```bash
# Create default document with operations metrics
prism init

# Create operations-focused document
prism init -d operations -o ops.json
```

For security metrics examples, see [prism-security](https://github.com/grokify/prism-security).

### Validate a document

```bash
prism validate prism.json
```

### Calculate PRISM score

```bash
# Basic score
prism score prism.json

# Detailed breakdown
prism score prism.json --detailed

# JSON output
prism score prism.json --json
```

### List available constants

```bash
prism catalog
```

## Schema Overview

### Domains

PRISM organizes metrics into two primary domains:

| Domain | Description |
|--------|-------------|
| `security` | Application and infrastructure security metrics |
| `operations` | Reliability, performance, and efficiency metrics |

### Lifecycle Stages

Metrics are mapped to software delivery lifecycle stages:

| Stage | Description |
|-------|-------------|
| `design` | Architecture, requirements, planning |
| `build` | CI/CD, code quality, dependency management |
| `test` | Testing coverage, quality assurance |
| `runtime` | Production monitoring, availability, performance |
| `response` | Incident response, remediation, recovery |

### Categories

| Category | Description |
|----------|-------------|
| `prevention` | Proactive controls that prevent issues |
| `detection` | Monitoring and alerting capabilities |
| `response` | Incident handling and remediation |
| `reliability` | Availability and durability |
| `efficiency` | Performance and resource utilization |
| `quality` | Code and process quality |

### Metric Types

| Type | Description | Example |
|------|-------------|---------|
| `coverage` | Percentage of coverage | Test coverage |
| `rate` | Frequency or percentage | Error rate |
| `latency` | Time duration | P99 latency, MTTR |
| `ratio` | Proportion | Success ratio |
| `count` | Absolute count | Deployment count |
| `distribution` | Statistical distribution | Latency percentiles |
| `score` | Composite score | Health score |

## Example Metric

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

## PRISM Score Calculation

The PRISM score combines maturity levels, metric performance, and customer awareness into a composite health score (0.0-1.0).

### Formula

```
CellScore = (MaturityWeight × MaturityScore) + (PerformanceWeight × PerformanceScore)
BaseScore = Σ(CellScore × Weight) / Σ(Weight)
Overall = BaseScore × AwarenessScore
```

### Default Weights

**Component weights:**

- Maturity: 40%
- Performance: 60%

**Stage weights:**

- Design: 15%
- Build: 20%
- Test: 15%
- Runtime: 30%
- Response: 20%

**Domain weights:**

- Security: 50%
- Operations: 50%

### Score Interpretation

| Score | Level | Description |
|-------|-------|-------------|
| ≥0.90 | Elite | Industry-leading practices |
| ≥0.75 | Strong | Well-managed, proactive |
| ≥0.50 | Medium | Adequate, room for improvement |
| ≥0.25 | Weak | Significant gaps |
| <0.25 | Critical | Immediate attention required |

## Maturity Levels

PRISM uses a 5-level maturity model:

| Level | Name | Description |
|-------|------|-------------|
| 1 | Reactive | Ad-hoc processes, firefighting mode |
| 2 | Basic | Basic controls, some documentation |
| 3 | Defined | Standardized processes, consistent execution |
| 4 | Managed | Data-driven, measured and controlled |
| 5 | Optimizing | Continuous improvement, automated optimization |

## Framework Mappings

PRISM metrics can be mapped to external frameworks:

| Framework | Description |
|-----------|-------------|
| `DORA` | DevOps Research and Assessment |
| `SRE` | Site Reliability Engineering |
| `NIST_CSF` | NIST Cybersecurity Framework (see [prism-security](https://github.com/grokify/prism-security)) |
| `MITRE_ATTACK` | MITRE ATT&CK Framework (see [prism-security](https://github.com/grokify/prism-security)) |

## JSON Schema

The JSON Schema is auto-generated from Go types:

```bash
cd schema && go run generate.go
```

Schema location: `schema/prism.schema.json`

Use in your editor for validation:

```json
{
  "$schema": "https://github.com/grokify/prism/schema/prism.schema.json",
  "metrics": [...]
}
```

## Goal-Driven Maturity Roadmap

PRISM supports goal-driven maturity tracking with multi-phase roadmaps.

### Goals

Goals represent strategic objectives with their own 5-level maturity models:

```json
{
  "id": "goal-reliability",
  "name": "Achieve High Reliability",
  "owner": "VP Engineering",
  "currentLevel": 3,
  "targetLevel": 4,
  "maturityModel": {
    "levels": [
      {
        "level": 3,
        "name": "Defined",
        "requiredSLOs": [
          { "metricId": "ops-availability" },
          { "metricId": "ops-mttr" }
        ],
        "metricCriteria": [
          { "metricId": "ops-availability", "operator": "gte", "value": 99.5 }
        ]
      }
    ]
  }
}
```

Each maturity level specifies which SLOs must be met to achieve that level.

### Phases

Phases organize work into time-bounded periods (quarters) with goal targets:

```json
{
  "id": "phase-q1-2026",
  "name": "Q1 2026",
  "quarter": "Q1",
  "year": 2026,
  "startDate": "2026-01-01",
  "endDate": "2026-03-31",
  "goalTargets": [
    { "goalId": "goal-reliability", "enterLevel": 2, "exitLevel": 3 }
  ],
  "swimlanes": [
    {
      "name": "Platform Initiatives",
      "domain": "operations",
      "initiativeIds": ["init-monitoring", "init-ci-cd"]
    }
  ]
}
```

### Initiatives

Initiatives link to goals and phases with deployment tracking:

```json
{
  "id": "init-monitoring",
  "name": "Observability Platform",
  "goalIds": ["goal-reliability"],
  "phaseId": "phase-q1-2026",
  "status": "completed",
  "deploymentStatus": {
    "status": "completed",
    "totalCustomers": 50,
    "deployedCustomers": 50,
    "adoptionPercent": 100
  }
}
```

## Examples

See the `examples/` directory:

- `operations-metrics.json` - Operations-focused metrics (DORA metrics, SLOs, reliability)
- `goal-roadmap.json` - Goal-driven maturity roadmap with phases and initiatives

For security metrics examples, see [prism-security](https://github.com/grokify/prism-security).

## Library Usage

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/grokify/prism"
)

func main() {
    // Load document
    data, _ := os.ReadFile("prism.json")
    var doc prism.PRISMDocument
    json.Unmarshal(data, &doc)

    // Validate
    if errs := doc.Validate(); errs.HasErrors() {
        fmt.Println("Validation errors:", errs)
        return
    }

    // Calculate score
    score := doc.CalculatePRISMScore(nil, nil)
    fmt.Printf("PRISM Score: %.1f%% (%s)\n", score.Overall*100, score.Interpretation)

    // Check individual metrics
    for _, m := range doc.Metrics {
        status := m.CalculateStatus()
        meetsSLO := m.MeetsSLO()
        fmt.Printf("  %s: %s (SLO met: %v)\n", m.Name, status, meetsSLO)
    }
}
```

## License

MIT
