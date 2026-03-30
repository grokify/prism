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
# Create default document with both domains
prism init

# Create security-focused document
prism init -d security -o security.json

# Create operations-focused document
prism init -d operations -o ops.json
```

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
| `design` | Architecture, threat modeling, requirements |
| `build` | CI/CD, SAST, dependency scanning |
| `test` | Testing coverage, penetration testing |
| `runtime` | Production monitoring, availability, detection |
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
| `coverage` | Percentage of coverage | SAST coverage |
| `rate` | Frequency or percentage | Error rate |
| `latency` | Time duration | P99 latency, MTTR |
| `ratio` | Proportion | Success ratio |
| `count` | Absolute count | Incident count |
| `distribution` | Statistical distribution | Latency percentiles |
| `score` | Composite score | Security score |

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
| `NIST_CSF` | NIST Cybersecurity Framework |
| `MITRE_ATTACK` | MITRE ATT&CK Framework |
| `DORA` | DevOps Research and Assessment |
| `SRE` | Site Reliability Engineering |

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

## Examples

See the `examples/` directory:

- `security-metrics.json` - Security-focused metrics (vulnerability management, SAST, threat modeling)
- `operations-metrics.json` - Operations-focused metrics (DORA metrics, SLOs, reliability)

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
