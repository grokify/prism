# PRISM Architecture

## Purpose

PRISM (Platform for Reliability, Intelligence, Strategy & Maturity) is a unified framework for capability-driven organizational intelligence. It connects **what you need** (capabilities), **how you measure** (maturity), and **how you act** (execution).

This repository (`grokify/prism`) is the **orchestration layer**: it loads documents from the PRISM module repositories, cross-references them, validates referential integrity, and generates dashboards and static sites.

## Module Structure

PRISM is a family of repositories with this one at the top:

```text
┌─────────────────────────────────────────────────────────────┐
│                             prism                           │
│                  Unified Orchestration Layer                │
│       Cross-module queries, validation, dashboards          │
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
│ "What we need" │───>│ "How we measure" │───>│ "How we act"  │
└────────────────┘    └──────────────────┘    └───────────────┘
```

| Module | Purpose | Key Artifacts |
|---|---|---|
| prism-core | Shared primitives | Domain, Layer, Stage, MaturityLevel, TeamType |
| prism-capability | What capabilities exist | Capability stacks, layers, dependencies |
| prism-maturity | How maturity is measured | SLI/SLO definitions, maturity state, services, metrics |
| prism-roadmap | How to improve | OKRs, V2MOM, roadmaps, initiatives, opportunity specs |
| prism (this repo) | Orchestration | Ecosystem loader, cross-module queries, site generation |

## Document Flow

Each artifact informs the next in a top-down planning flow:

```text
Capability Stack          What capabilities do we need?
       │
       ▼
Maturity Model + State    Where are we for each capability?
       │
       ▼
OKR / V2MOM + Roadmap     What goals and sequence to improve?
       │
       ▼
MRD / PRD / TRD           How do we execute each roadmap item?
```

## Position in the ProductContext Ecosystem

PRISM-Roadmap is the **decision layer** of the four-domain ProductContext architecture:

| Repository | Owns | Primary Question |
|---|---|---|
| OrganizationSpec | Customers, prospects, partners, accounts | Who are we serving? |
| MarketSpec | Markets, competitors, analyst intelligence, trends, TAM/SAM/SOM | What market are we competing in? |
| OmniSignal | Normalized evidence and observations | What signals are we observing? |
| PRISM-Roadmap | Strategy, prioritization, roadmaps, portfolios | Given the evidence, what should we build next? |

The traceable chain runs:

```text
Signals (OmniSignal)
        ↓
Evidence (Customer, Market, Analyst, Competitor)
        ↓
Strategic Pillars (CSAT, SAM/SOM, TAM)
        ↓
Roadmap Themes
        ↓
Capabilities (PRISM Capability)
        ↓
Epics / Features / Work Items
```

PRISM references canonical entities defined elsewhere by typed ID — for example `market:identity-governance` (MarketSpec) or `customer:acme` (OrganizationSpec) — and consumes evidence aggregated by OmniSignal. It does not redefine those entities.

## Strategic Pillars

Roadmap decisions are organized around three pillars rather than a single priority score:

| Pillar | Question | Typical Evidence |
|---|---|---|
| CSAT | How do we better serve existing customers? | Support volume, frustration scores, ARR affected, churn risk |
| SAM/SOM Expansion | How do we win more share in markets we serve? | Competitive gaps, win/loss, analyst recommendations |
| TAM Expansion | What new markets should we enter? | Market sizing, trends, strategic fit |

Initiatives can contribute to multiple pillars (e.g. SCIM Improvements: CSAT High, SAM/SOM Medium, TAM Low). Portfolio-level rollups show investment distribution — for example 60% CSAT, 25% SAM/SOM, 15% TAM.

## Multi-Dimensional Prioritization

Rather than an opaque score, each signal contributes evidence along explainable dimensions:

| Dimension | Example Inputs |
|---|---|
| Customer Value | Named customer requests, ARR affected, vote counts |
| Market Opportunity | TAM/SAM/SOM estimates, analyst research, competitor gaps |
| Strategic Alignment | Product vision, OKRs, capability targets, long-term themes |
| Engineering Cost | Estimated effort, dependencies, technical debt |
| Risk | Security, compliance, reliability, architectural risk |
| Adoption Evidence | Usage telemetry, feature adoption, support frequency |
| Confidence | Number and quality of corroborating signals |

## Orchestration Layer Components

This repository contains:

- `ecosystem/` — Loads capability stacks, maturity documents, OKR sets, and roadmaps from configured file sources into a single `Ecosystem`; provides cross-module queries (`AllCapabilities`, `AllMetrics`, `AllInitiatives`, `GetCapabilityContext`), referential validation, and stats.
- `cmd/prism/` — CLI with `ecosystem`, `site`, `stats`, and `validate` subcommands.
- `sitegen/` — Static site generator for capability stacks with maturity badges and theme support.
- `ui/` — Lit web components (`@prism/ui`) for interactive visualization.
- `viewer/` — TypeScript viewer package.

The orchestration layer never owns document schemas; those live in the module repositories. Its job is loading, joining, validating, and rendering.
