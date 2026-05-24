# PRISM

**P**latform for **R**eliability, **I**ntelligence, **S**trategy & **M**aturity

A unified framework for capability-driven organizational intelligence, connecting **what you need** (capabilities), **how you measure** (maturity), and **how you act** (execution).

## Ecosystem Overview

```
┌─────────────────────────────────────────────────────────────┐
│                             prism                           │
│                  Unified Orchestration Layer                │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                         prism-core                          │
│       Shared primitives: Domain, Layer, Stage, Maturity     │
└─────────────────────────────────────────────────────────────┘
        │                       │                     │
        ▼                       ▼                     ▼
┌────────────────┐    ┌──────────────────┐    ┌───────────────┐
│prism-capability│    │  prism-maturity  │    │ prism-roadmap │
│                │    │                  │    │               │
│ "What we need" │───>│ "How we measure" │───>│ "How we act"  │
└────────────────┘    └──────────────────┘    └───────────────┘
```

## Modules

| Module | Purpose | Documentation |
|--------|---------|---------------|
| [prism-core](https://github.com/grokify/prism-core) | Shared primitives | Domain, Layer, Stage, MaturityLevel |
| [prism-capability](https://github.com/grokify/prism-capability) | What capabilities exist | Capability stacks, layers |
| [prism-maturity](https://grokify.github.io/prism-maturity/) | How we measure maturity | SLIs, SLOs, maturity models |
| [prism-roadmap](https://grokify.github.io/prism-roadmap/) | How we improve | OKRs, roadmaps, PRDs |
| **prism** (this repo) | Orchestration | Cross-module queries |

## Installation

```bash
go get github.com/grokify/prism@latest
```

## Quick Start

```go
package main

import (
    "github.com/grokify/prism/ecosystem"
)

func main() {
    // Create an ecosystem
    eco := ecosystem.New(ecosystem.Config{
        Name: "my-organization",
    })

    // Load capability stacks
    eco.LoadCapabilityStack("capabilities/security.json")

    // Load maturity documents
    eco.LoadPRISMDocument("maturity/model.json")

    // Query across modules
    stats := eco.Statistics()
    fmt.Printf("Capabilities: %d\n", stats.TotalCapabilities)
    fmt.Printf("Metrics: %d\n", stats.TotalMetrics)
}
```

## Core Concepts

### Gap-Driven Planning

1. **Capability Stack** defines *what* the organization needs
2. **Maturity Model** defines *how well* each capability should perform
3. **Maturity State** tracks *where we are today*
4. **Execution Plan** defines *how we improve*

### Maturity Levels (M1-M5)

| Level | Name | Characteristics |
|-------|------|-----------------|
| M1 | Reactive | Missing or informal, inconsistent outcomes |
| M2 | Basic | Some tooling, individual-dependent |
| M3 | Defined | Standardized, repeatable across teams |
| M4 | Managed | Measured, governed, integrated |
| M5 | Optimizing | Continuously improved, adaptive |

## Site Generator

Generate static sites from capability stacks with maturity visualization:

```bash
prism site generate --stack=./stacks/ --output=./dist
```

### Features

- **Metrics Count Indicator** - Badge showing SLI count per capability
- **Light/Dark Theme** - Automatic theme synchronization
- **Navigation** - Browse by category or layer
- **Custom Footer** - Use `--hide-generated-by` to remove attribution

### CLI Options

| Flag | Description |
|------|-------------|
| `--stack` | Path to capability stack directory |
| `--output` | Output directory for generated site |
| `--hide-generated-by` | Hide footer attribution |

## Next Steps

- [Ecosystem Design](ECOSYSTEM.md) - Full architecture documentation
- [v0.7.0 Release Notes](releases/v0.7.0.md) - Latest release
