# prism export

Export a PRISM document to OKR, V2MOM, or Roadmap format for use with structured-plan.

## Synopsis

```bash
prism export <subcommand> <prism-file> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `okr` | Export as OKR (Objectives and Key Results) document |
| `v2mom` | Export as V2MOM (Vision, Values, Methods, Obstacles, Measures) document |
| `roadmap` | Export as Roadmap with deployment/adoption tracking |

## Description

The `export` command converts PRISM goals, SLOs, phases, and initiatives into planning document formats compatible with the [structured-plan](https://github.com/grokify/structured-plan) repository.

### Mapping

| PRISM Concept | Roadmap Concept | OKR Concept | V2MOM Concept |
|---------------|-----------------|-------------|---------------|
| Goal | — | Objective | Method |
| Goal (per maturity level) | — | Objective (one per level) | — |
| SLO/MetricCriterion | — | Key Result | Measure |
| Phase | Phase | — | — |
| Phase.GoalTargets | Phase.Goals | PhaseTargets | Method timeline |
| Initiative | Deliverable | Referenced in KR | Project |
| DeploymentStatus | RolloutStatus | — | — |

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--output` | `-o` | Output file or directory (default: stdout) |

## prism export okr

Export PRISM as an OKR document.

### Usage

```bash
prism export okr <prism-file> [flags]
```

### Examples

```bash
# Output to stdout
prism export okr prism.json

# Output to file
prism export okr prism.json -o roadmap.okr.json

# Output to directory (uses default filename)
prism export okr prism.json -o ./exports/
```

### Output Format

```json
{
  "$schema": "https://github.com/grokify/structured-plan/schema/okr.schema.json",
  "metadata": {
    "title": "PRISM Roadmap",
    "description": "OKR document generated from PRISM",
    "version": "1.0.0",
    "generatedFrom": "prism"
  },
  "objectives": [
    {
      "id": "obj-goal-reliability",
      "title": "Achieve High Reliability",
      "description": "Build world-class reliability with elite SLO compliance",
      "progress": 60.0,
      "tags": ["operations", "reliability"],
      "keyResults": [
        {
          "id": "kr-ops-availability",
          "title": "Service Availability >= 99.9%",
          "description": "Achieve 99.9% availability SLO",
          "target": ">=99.9%",
          "score": 0.85,
          "phaseTargets": [
            {
              "phaseId": "phase-q1-2026",
              "target": ">=99.5%"
            },
            {
              "phaseId": "phase-q2-2026",
              "target": ">=99.9%"
            }
          ]
        }
      ]
    }
  ],
  "alignment": {
    "goalToObjective": {
      "goal-reliability": "obj-goal-reliability"
    },
    "sloToKeyResult": {
      "ops-availability": "kr-ops-availability"
    }
  }
}
```

## prism export v2mom

Export PRISM as a V2MOM document.

### Usage

```bash
prism export v2mom <prism-file> [flags]
```

### Examples

```bash
# Output to stdout
prism export v2mom prism.json

# Output to file
prism export v2mom prism.json -o roadmap.v2mom.json
```

### Output Format

```json
{
  "$schema": "https://github.com/grokify/structured-plan/schema/v2mom.schema.json",
  "metadata": {
    "title": "PRISM Roadmap",
    "description": "V2MOM document generated from PRISM",
    "version": "1.0.0",
    "generatedFrom": "prism"
  },
  "vision": "Achieve operational excellence through systematic maturity improvement",
  "values": [
    "Reliability",
    "Data-Driven Decisions",
    "Continuous Improvement"
  ],
  "methods": [
    {
      "id": "method-goal-reliability",
      "name": "Achieve High Reliability",
      "description": "Build world-class reliability with elite SLO compliance",
      "priority": 1,
      "measures": [
        {
          "id": "measure-ops-availability",
          "name": "Service Availability",
          "target": ">=99.9%",
          "current": 99.5,
          "progress": 0.85
        }
      ],
      "projects": [
        {
          "id": "proj-init-observability",
          "name": "Observability Platform",
          "status": "in_progress"
        }
      ]
    }
  ]
}
```

## prism export roadmap

Export PRISM as a roadmap document with deployment and adoption tracking.

### Usage

```bash
prism export roadmap <prism-file> [flags]
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--output` | `-o` | Output file or directory (default: stdout) |
| `--with-okrs` | | Include OKRs alongside roadmap |

### Examples

```bash
# Output roadmap to stdout
prism export roadmap prism.json

# Output to file
prism export roadmap prism.json -o product-roadmap.json

# Export roadmap with embedded OKRs
prism export roadmap prism.json --with-okrs -o roadmap-full.json
```

### Output Format

The roadmap export uses the [structured-plan/roadmap](https://github.com/grokify/structured-plan) format:

```json
{
  "phases": [
    {
      "id": "phase-q1-2026",
      "name": "Q1 2026",
      "type": "quarter",
      "startDate": "2026-01-01T00:00:00Z",
      "endDate": "2026-03-31T00:00:00Z",
      "status": "in_progress",
      "goals": [
        "Achieve Reliability Level 4"
      ],
      "deliverables": [
        {
          "id": "init-monitoring",
          "title": "Observability Platform",
          "description": "Deploy comprehensive monitoring",
          "type": "feature",
          "status": "in_progress",
          "tags": ["prism", "initiative", "goal:goal-reliability"],
          "rollout": {
            "totalCustomers": 50,
            "deployedCustomers": 45,
            "adoptedCustomers": 40,
            "status": "rolling_out",
            "notes": "Dev completion: 90%"
          }
        }
      ],
      "successCriteria": [
        "availability >= 99.9%"
      ]
    }
  ]
}
```

### RolloutStatus

The roadmap export includes B2B SaaS deployment tracking via `RolloutStatus`:

| Field | Description |
|-------|-------------|
| `totalCustomers` | Total customers in rollout scope |
| `deployedCustomers` | Customers with feature deployed (available) |
| `adoptedCustomers` | Customers actively using the feature |
| `status` | Rollout stage: `not_started`, `rolling_out`, `deployed`, `adopted`, `paused`, `rolled_back` |
| `startDate` | Rollout start date |
| `targetDate` | Target completion date |
| `waves` | Phased rollout waves (beta, GA, etc.) |

**Deployment vs Adoption:**

- **Deployment** = Feature is available to the customer (rolled out)
- **Adoption** = Customer is actively using the feature

```
Deployment %  = deployedCustomers / totalCustomers × 100
Adoption %    = adoptedCustomers / totalCustomers × 100
Adoption Rate = adoptedCustomers / deployedCustomers × 100  (among deployed)
```

### With OKRs

When using `--with-okrs`, the export includes both roadmap and OKRs:

```json
{
  "roadmap": {
    "phases": [...]
  },
  "okrs": {
    "metadata": {...},
    "objectives": [
      {
        "id": "obj-m4-goal-reliability",
        "title": "Achieve Managed Level for Reliability",
        "category": "Operational Maturity",
        "keyResults": [
          {
            "id": "kr-availability-m4",
            "title": "availability meets M4 requirements",
            "target": ">=99.9%"
          }
        ]
      }
    ]
  }
}
```

Each maturity level to achieve becomes a separate OKR objective with key results derived from the SLO requirements at that level.

## Workflow

The export commands are designed for use in a structured planning workflow:

```bash
# 1. Analyze PRISM document
prism analyze prism.json -f prompt > analysis.md

# 2. Generate initiatives with LLM (manual step)
# Use analysis.md as input to LLM

# 3. Export to roadmap for execution tracking
prism export roadmap prism.json --with-okrs -o roadmap.json

# 4. Or export to OKR-only format
prism export okr prism.json -o roadmap.okr.json

# 5. Use with structured-plan tools
splan validate roadmap.json
splan render roadmap.json -o roadmap.html
```

### Choosing an Export Format

| Format | Best For |
|--------|----------|
| **roadmap** | B2B SaaS with customer deployment tracking, phased rollouts |
| **okr** | Goal-focused tracking without deployment details |
| **v2mom** | Salesforce-style planning with vision/values emphasis |

## Related Commands

- [`prism analyze`](analyze.md) - Analyze document and generate recommendations
- [`prism goal`](goal.md) - Manage and inspect goals
- [`prism phase`](phase.md) - Manage and inspect phases
- [`prism roadmap`](roadmap.md) - View roadmap overview
