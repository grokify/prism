# Concepts Overview

PRISM combines several established methodologies into a unified metrics framework:

## Core Concepts

### [Goals](goals.md)

Strategic objectives with their own maturity models. Each goal defines:

- 5-level maturity progression specific to that goal
- SLO requirements for each maturity level
- Initiative linkage for tracking progress

### [Phases](phases.md)

Time-bounded planning periods (quarters) that organize:

- Goal maturity targets (enter/exit levels)
- Swimlanes for initiative grouping
- Phase metrics (completion %, SLO compliance)

### [Maturity Model](maturity.md)

A 5-level capability maturity model adapted for operations and extensible domains:

| Level | Name | Description |
|-------|------|-------------|
| 1 | Reactive | Ad-hoc processes, firefighting |
| 2 | Basic | Basic controls, some documentation |
| 3 | Defined | Standardized processes |
| 4 | Managed | Measured and controlled |
| 5 | Optimizing | Continuous improvement |

### [Customer Awareness](awareness.md)

Track customer awareness of issues through four mutually exclusive states:

| State | Weight | Description |
|-------|--------|-------------|
| Unaware | 0.0 | Customer not aware |
| Aware (not acting) | 0.25 | Aware but not remediating |
| Remediating | 0.5 | Actively working on fix |
| Remediated | 1.0 | Issue resolved |

### [PRISM Score](scoring.md)

A composite health score (0.0-1.0) combining:

- **Maturity scores** (40% weight) - organizational capability
- **Performance scores** (60% weight) - metric achievement
- **Awareness multiplier** - customer communication effectiveness

### [Framework Mappings](frameworks.md)

Map metrics to industry standards:

- DORA metrics
- SRE practices
- NIST Cybersecurity Framework (see [prism-security](https://github.com/grokify/prism-security))
- MITRE ATT&CK (see [prism-security](https://github.com/grokify/prism-security))

## Methodology Foundations

PRISM draws from several established methodologies:

### DMAIC (Six Sigma)

- **Define** - Define what you're measuring
- **Measure** - Collect baseline data
- **Analyze** - Understand current performance
- **Improve** - Set targets and improve
- **Control** - Monitor and maintain

PRISM metrics support DMAIC with:

- Baseline values
- Current measurements
- Thresholds for control
- Targets for improvement

### OKRs (Objectives & Key Results)

PRISM supports OKR alignment through:

- Metric targets as key results
- Progress tracking (`ProgressToTarget()`)
- Domain/stage organization for objectives

### SRE (Site Reliability Engineering)

PRISM adopts SRE concepts:

- SLIs (Service Level Indicators)
- SLOs (Service Level Objectives)
- Error budgets (via thresholds)

### Capability Maturity

PRISM's maturity model is adapted from:

- CMMI (Capability Maturity Model Integration)
- BSIMM (Building Security In Maturity Model)
- OpenSAMM (Software Assurance Maturity Model)

## Score Hierarchy

```
PRISM Score (Overall)
├── Base Score
│   ├── Domain Score (e.g., Operations)
│   │   ├── Design Cell Score
│   │   ├── Build Cell Score
│   │   ├── Test Cell Score
│   │   ├── Runtime Cell Score
│   │   └── Response Cell Score
│   └── Additional Domains (extensible)
│       └── ...
└── Awareness Multiplier
```

For security domain scoring, see [prism-security](https://github.com/grokify/prism-security).

## Getting Started

1. **Define Metrics** - Identify key metrics for each domain/stage
2. **Set Baselines** - Establish current state
3. **Set Targets** - Define improvement goals
4. **Assess Maturity** - Rate organizational capability
5. **Calculate Score** - Get composite health score
6. **Track Progress** - Monitor over time
