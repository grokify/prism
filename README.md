# PRISM

[![Go Reference](https://pkg.go.dev/badge/github.com/grokify/prism.svg)](https://pkg.go.dev/github.com/grokify/prism)

**P**latform for **R**eliability, **I**ntelligence, **S**trategy & **M**aturity

A unified framework for capability-driven organizational intelligence, connecting **what you need** (capabilities), **how you measure** (maturity), and **how you act** (execution).

## Ecosystem

```
┌─────────────────────────────────────────────────────────────────────┐
│                              prism                                   │
│                    Unified Orchestration Layer                       │
└─────────────────────────────────────────────────────────────────────┘
        │                    │                    │
        ▼                    ▼                    ▼
┌───────────────┐    ┌───────────────┐    ┌───────────────┐
│prism-capability│   │prism-intelligence│  │prism-execution│
│               │    │                 │    │               │
│ "What we need"│───▶│ "How we measure"│───▶│ "How we act" │
└───────────────┘    └───────────────┘    └───────────────┘
```

| Module | Purpose | Key Artifacts |
|--------|---------|---------------|
| [prism-capability](https://github.com/grokify/prism-capability) | What capabilities exist | Capability stacks, layers, dependencies |
| [prism-intelligence](https://github.com/grokify/prism-intelligence) | How we measure maturity | SLI/SLO definitions, maturity state |
| [prism-execution](https://github.com/grokify/prism-execution) | How we improve | OKRs, roadmaps, initiatives |
| **prism** (this repo) | Orchestration | Cross-module queries, dashboards |

## Installation

```bash
go install github.com/grokify/prism/cmd/prism@latest
```

## Quick Start

```bash
# Load an ecosystem configuration
prism ecosystem load --config prism.yaml

# Analyze maturity gaps
prism gaps analyze --sort-by=impact

# Generate dashboard
prism dashboard generate --format=html -o dashboard.html

# Validate all documents
prism validate --all
```

## Ecosystem Configuration

```yaml
# prism.yaml
ecosystem:
  name: my-organization

  capability:
    files:
      - capabilities/security.json
      - capabilities/platform.json

  intelligence:
    model: maturity/model.json
    state: maturity/state.json

  execution:
    okrs: plans/okrs/
    roadmaps: plans/roadmaps/
    initiatives: plans/initiatives/
```

## Core Concepts

### Capability Stack → Maturity Model → Execution Plan

1. **Capability Stack** defines *what* the organization needs
2. **Maturity Model** defines *how well* each capability should perform
3. **Maturity State** tracks *where we are today*
4. **Execution Plan** defines *how we improve*

### Gap-Driven Planning

```
Current State (M2) ──────────────────▶ Target State (M4)
                         │
                         │ Gap Analysis
                         ▼
              ┌─────────────────────┐
              │    Initiatives      │
              │  • Enable CI SAST   │
              │  • Custom rules     │
              │  • SLA automation   │
              └─────────────────────┘
```

### Maturity Levels (M1-M5)

| Level | Name | Characteristics |
|-------|------|-----------------|
| M1 | Ad hoc | Missing or informal, inconsistent outcomes |
| M2 | Basic | Some tooling, individual-dependent |
| M3 | Defined | Standardized, repeatable across teams |
| M4 | Managed | Measured, governed, integrated |
| M5 | Optimized | Continuously improved, adaptive |

## Library Usage

```go
package main

import (
    "github.com/grokify/prism"
)

func main() {
    // Load ecosystem
    eco, err := prism.LoadEcosystem("prism.yaml")
    if err != nil {
        log.Fatal(err)
    }

    // Find capabilities with maturity gaps
    gaps := eco.CapabilitiesWithMaturityGap()
    for _, gap := range gaps {
        fmt.Printf("%s: %s → %s (gap: %d levels)\n",
            gap.Capability.Name,
            gap.CurrentLevel,
            gap.TargetLevel,
            gap.GapSize)
    }

    // Get full context for a capability
    ctx := eco.CapabilityContext("sast")
    fmt.Printf("Capability: %s\n", ctx.Capability.Name)
    fmt.Printf("Current Maturity: %s\n", ctx.CurrentMaturity.Level)
    fmt.Printf("Active Initiatives: %d\n", len(ctx.Initiatives))
}
```

## Documentation

- [Ecosystem Design](docs/ECOSYSTEM.md) - Full architecture and design
- [prism-capability](https://github.com/grokify/prism-capability) - Capability stack documentation
- [prism-intelligence](https://github.com/grokify/prism-intelligence) - Maturity model documentation
- [prism-execution](https://github.com/grokify/prism-execution) - Execution planning documentation

## Status

This project is in active development. Current phase: **Foundation**

| Module | Latest Version | Status |
|--------|----------------|--------|
| [prism-capability](https://github.com/grokify/prism-capability) | v0.2.0 | Released |
| [prism-intelligence](https://github.com/grokify/prism-intelligence) | v0.8.0 | Released |
| [prism-execution](https://github.com/grokify/prism-execution) | v0.11.0 | Released |
| prism (this repo) | v0.1.0 | In development |

### Roadmap

- [x] prism-capability v0.2.0 released
- [x] prism-intelligence v0.8.0 released
- [x] prism-execution v0.11.0 released
- [x] Repository restructuring complete
- [ ] Orchestrator types and loading
- [ ] Cross-module validation
- [ ] Unified CLI
- [ ] Dashboard generation

## License

MIT
