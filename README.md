# PRISM

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Docs][docs-mkdoc-svg]][docs-mkdoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/prism/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/prism/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/prism/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/prism/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/prism/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/prism/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/prism
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/prism
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/prism
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/prism
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://grokify.github.io/prism
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fprism
 [loc-svg]: https://tokei.rs/b1/github/grokify/prism
 [repo-url]: https://github.com/grokify/prism
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/prism/blob/main/LICENSE

**P**latform for **R**eliability, **I**ntelligence, **S**trategy & **M**aturity

A unified framework for capability-driven organizational intelligence, connecting **what you need** (capabilities), **how you measure** (maturity), and **how you act** (execution).

## Ecosystem

```
┌─────────────────────────────────────────────────────────────┐
│                             prism                           │
│                  Unified Orchestration Layer                │
└─────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────┐
│                          prism-core                         │
│      Shared primitives: Domain, Layer, Stage, Maturity      │
└─────────────────────────────────────────────────────────────┘
        │                       │                     │
        ▼                       ▼                     ▼
┌────────────────┐    ┌──────────────────┐    ┌───────────────┐
│prism-capability│    │  prism-maturity  │    │ prism-roadmap │
│                │    │                  │    │               │
│ "What we need" │───>│ "How we measure" │───>│ "How we act"  │
└────────────────┘    └──────────────────┘    └───────────────┘
```

| Module | Purpose | Key Artifacts |
|--------|---------|---------------|
| [prism-core](https://github.com/grokify/prism-core) | Shared primitives | Domain, Layer, Stage, MaturityLevel, TeamType |
| [prism-capability](https://github.com/grokify/prism-capability) | What capabilities exist | Capability stacks, layers, dependencies |
| [prism-maturity](https://github.com/grokify/prism-maturity) | How we measure maturity | SLI/SLO definitions, maturity state |
| [prism-roadmap](https://github.com/grokify/prism-roadmap) | How we improve | OKRs, roadmaps, initiatives |
| **prism** (this repo) | Orchestration | Cross-module queries, dashboards |

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

# Generate static site from capability stacks
prism site generate --stack=./stacks/ --output=./dist

# Generate site without footer attribution
prism site generate --stack=./stacks/ --output=./dist --hide-generated-by

# Validate all documents
prism validate --all
```

## Site Generator Features

### Metrics Count Indicator

The site generator displays a badge on each capability card showing the number of SLIs/metrics being tracked. This appears as a circular indicator (similar to unread message badges) in the top-right corner of capability cards.

### Theme Support

The generated site supports light and dark themes with automatic synchronization between the navigation bar and page content. The theme toggle in the navigation bar updates all page elements.

### Lit Web Components

For interactive visualization, the site generator supports Lit-based web components via the `@prism/ui` package:

```bash
# Build the UI components
cd ui/ && npm install && npm run build

# Generate site with Lit component integration
prism site generate \
  --stack=./security/ \
  --output=./dist \
  --prism-ui-js=./ui/dist/prism-ui.js
```

The `maturity-grid` component provides:

- Interactive capability grid with maturity overlay
- Toggle between by-layer and by-category views
- Theme-aware rendering synced with site theme
- Click-through to capability detail pages

### CLI Options

| Flag | Description |
|------|-------------|
| `--stack` | Path to capability stack directory or JSON file |
| `--output` | Output directory for generated site |
| `--theme` | Site theme: `light` or `dark` |
| `--aggregation` | Aggregation method: `min` (conservative) or `avg` (average) |
| `--prism-ui-js` | Path to prism-ui.js Lit component bundle (enables interactive rendering) |
| `--hide-generated-by` | Hide the "Generated by PRISM" footer message |

## Standard Directory Structure

PRISM supports a standard directory structure that enables auto-discovery of related files:

```
{stack-name}/
├── stack.json           # Capability stack definition
├── model.json           # Maturity model (SLIs, levels, criteria)
├── state.json           # Current state (SLI values)
└── roadmap.json         # OKRs, initiatives (optional)
```

When using this structure, the `prism site generate` command auto-discovers all related files:

```bash
# Single stack directory
prism site generate --stack=./security/

# Parent directory with multiple stacks
prism site generate --stack=./stacks/ --output=./dist

# Multiple individual stacks
prism site generate --stack=./security/ --stack=./reliability/ --output=./dist
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

  maturity:
    model: maturity/model.json
    state: maturity/state.json

  roadmap:
    okrs: plans/okrs/
    roadmaps: plans/roadmaps/
    requirements: plans/requirements/
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
- [prism-maturity](https://github.com/grokify/prism-maturity) - Maturity model documentation
- [prism-roadmap](https://github.com/grokify/prism-roadmap) - Execution planning documentation

## Status

This project is in active development. Current phase: **Foundation**

| Module | Latest Version | Status |
|--------|----------------|--------|
| [prism-core](https://github.com/grokify/prism-core) | v0.3.0 | Released |
| [prism-capability](https://github.com/grokify/prism-capability) | v0.6.0 | Released |
| [prism-maturity](https://github.com/grokify/prism-maturity) | v0.12.0 | Released |
| [prism-roadmap](https://github.com/grokify/prism-roadmap) | v0.14.1 | Released |
| prism (this repo) | v0.8.0 | Released |

### Roadmap

- [x] prism-core v0.3.0 released (developer productivity frameworks)
- [x] prism-capability v0.6.0 released (operations & dev productivity)
- [x] prism-maturity v0.12.0 released (SPACE, AI-DORA, AI-SPACE)
- [x] prism-roadmap v0.14.1 released
- [x] Repository restructuring complete
- [x] Orchestrator types and loading
- [x] Cross-module integration
- [x] Unified CLI
- [x] Dashboard generation
- [x] Static site generation
- [x] Lit web components for interactive visualization

## License

MIT
