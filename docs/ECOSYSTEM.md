# PRISM Ecosystem Design

## Overview

PRISM is a unified framework for capability-driven organizational intelligence, connecting **what you need** (capabilities), **how you measure** (maturity), and **how you act** (execution).

```
┌────────────────────────────────────────────────────────────┐
│                           prism                            │
│                Unified Orchestration Layer                 │
│      Queries, cross-references, dashboards, workflows      │
└────────────────────────────────────────────────────────────┘
         │                    │                    │
         ▼                    ▼                    ▼
┌────────────────┐    ┌─────────────────┐    ┌───────────────┐
│prism-capability│    │  prism-maturity │    │ prism-roadmap │
│                │    │                 │    │               │
│ "What we need" │───>│ "How we measure"│───>│ "How we act"  │
│                │    │                 │    │               │
│ Capabilities   │    │ Maturity Models │    │ OKRs/V2MOM    │
│ Layers         │    │ SLI/SLO Defs    │    │ Roadmaps      │
│ Categories     │    │ Current State   │    │ Initiatives   │
│ Dependencies   │    │ Gap Analysis    │    │ Requirements  │
└────────────────┘    └─────────────────┘    └───────────────┘
```

## Document Flow

The PRISM ecosystem follows a top-down planning flow where each artifact informs the next:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           CAPABILITY STACK                                  │
│                    "What capabilities do we need?"                          │
│         Layers, capabilities, dependencies, maturity targets                │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                      MATURITY MODEL + STATE                                 │
│                    "Where are we for each capability?"                      │
│              SLIs, criteria (M1-M5), current measurements                   │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                        OKR / V2MOM + ROADMAP                                │
│              "What do we want to achieve, and in what order?"               │
│         Goals (definition of success), phased roadmap items (RMIs)          │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                           MRD / PRD / TRD                                   │
│                    "How do we execute each roadmap item?"                   │
│     Market requirements, product requirements, technical requirements       │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Flow Summary

| Stage | Artifact | Question Answered | Module |
|-------|----------|-------------------|--------|
| 1 | Capability Stack | What capabilities do we need? | prism-capability |
| 2 | Maturity Model + State | Where are we for each capability? | prism-maturity |
| 3 | OKR/V2MOM + Roadmap | What goals and sequence to improve? | prism-roadmap |
| 4 | MRD/PRD/TRD | How do we execute each roadmap item? | prism-roadmap |

### Relationships

- **Capability → Maturity**: Each capability links to SLIs via `PRISMRef.SLIIDs`
- **Maturity → Roadmap**: Gaps between current state and target drive roadmap priorities
- **Roadmap → Requirements**: Each Roadmap Item (RMI) is executed via MRD/PRD/TRD documents

## Module Structure

| Module | Repository | Purpose |
|--------|------------|---------|
| **prism** | `github.com/grokify/prism` | Orchestrator connecting all modules |
| **prism-capability** | `github.com/grokify/prism-capability` | Capability stack definitions |
| **prism-maturity** | `github.com/grokify/prism-maturity` | Maturity models, SLIs, state tracking |
| **prism-roadmap** | `github.com/grokify/prism-roadmap` | Goals, roadmaps, requirements |

## Core Concepts

### Core Principles

The PRISM ecosystem is grounded in these principles:

1. **Capability Stack = Stable Structure**
   - What the organization/system must be able to do
   - Domain/function-first decomposition
   - Relatively stable over time

2. **Maturity = Outcome Realization**
   - Not just "having something in place"
   - How well it works in practice, at scale
   - Composite of multiple dimensions

3. **Roadmap = Time-Phased Delivery**
   - Capability-to-initiative mapping
   - Dependency-driven sequencing
   - Two-way traceability (strategy ↔ execution)

### Maturity Dimensions

From the ideation analysis, maturity is multi-dimensional:

| Dimension | What It Measures | Metric Type |
|-----------|------------------|-------------|
| **Operational Fitness** | Reliability, latency, uptime | SLOs |
| **Adoption** | Usage, penetration, coverage | Usage metrics |
| **Outcome Impact** | Business value, decision quality | KPIs |
| **Integration** | Embeddedness in flows | Dependency metrics |
| **Agility** | Time-to-change, evolution speed | Lead time metrics |

### Importance and Dynamic Priority

PRISM uses a two-tier priority system:

1. **Static Importance** - Inherent weight of a capability (defined in prism-capability)
2. **Dynamic Priority (P0-P3)** - Calculated from importance and maturity gap

| Importance | Weight | Description |
|------------|--------|-------------|
| `critical` | 4 | Critical "-ilities" (security, availability, resiliency) |
| `high` | 3 | High importance capabilities |
| `medium` | 2 | Standard importance (default) |
| `low` | 1 | Nice-to-have capabilities |

**Priority Calculation:**

```
Priority Score = Importance Weight × (Target Level - Current Level)
```

| Score | Priority | Action Required |
|-------|----------|-----------------|
| ≥8 | P0 | Immediate action |
| ≥4 | P1 | High priority |
| ≥2 | P2 | Scheduled improvement |
| <2 | P3 | Low priority |

**Example:**
- Critical capability at M2, target M4 → Score = 4 × 2 = 8 → **P0**
- Medium capability at M3, target M4 → Score = 2 × 1 = 2 → **P2**

### SLI Hierarchy

```
┌─────────────────────────────────────────────────────────────────┐
│                    Capability Level Objectives                  │
│         Unified threshold-based measurement framework           │
└─────────────────────────────────────────────────────────────────┘
        │                    │                    │
        ▼                    ▼                    ▼
┌───────────────┐    ┌───────────────┐    ┌───────────────┐
│     SLOs      │    │   Adoption    │    │   Outcome     │
│               │    │   Objectives  │    │   Objectives  │
│ Availability  │    │ Team coverage │    │ Business KPIs │
│ Latency       │    │ Transaction % │    │ Decision lift │
│ Error rate    │    │ API usage     │    │ Conversion    │
└───────────────┘    └───────────────┘    └───────────────┘
```

## Module Details

### prism-capability

**Purpose:** Define what capabilities exist in an organization's technology landscape.

**Key Types:**
- `CapabilityStack` - Root document
- `Capability` - Individual capability with status, dependencies, tooling
- `Layer` - Horizontal grouping (phases, stages)
- `Category` - Visual grouping with colors
- `PRISMRef` - Link to maturity model

**Document Example:**
```json
{
  "metadata": { "name": "security-stack", "domain": "security" },
  "layers": [
    { "id": "shift-left", "name": "Shift-Left Security", "phase": "build" }
  ],
  "capabilities": [
    {
      "id": "sast",
      "name": "SAST",
      "layerId": "shift-left",
      "status": "operational",
      "prismRef": {
        "domainId": "security",
        "sliIds": ["sli-sast-coverage", "sli-sast-findings"]
      }
    }
  ]
}
```

### prism-maturity

**Purpose:** Define how capabilities are measured and track current maturity state.

**Key Types:**
- `MaturityModel` - SLI definitions, level thresholds
- `MaturityState` - Current measurements per capability
- `SLI` - Service Level Indicator definition
- `Domain` - Grouping of related SLIs
- `LevelThreshold` - M1-M5 criteria

**Document Relationships:**
```
MaturityModel (what good looks like)
    │
    ▼
MaturityState (where we are today)
    │
    ▼
Gap Analysis (what needs improvement)
```

**SLI Categories:**
| Category | Focus | Examples |
|----------|-------|----------|
| Operational | System health | Availability, latency, error rate |
| Adoption | Usage uptake | Team coverage, transaction flow |
| Outcome | Business impact | Revenue lift, risk reduction |
| Quality | Decision effectiveness | Precision, recall, override rate |

### prism-roadmap

**Purpose:** Plan and track improvement initiatives aligned to capability maturity.

**Key Types:**
- `OKR` - Objectives and Key Results
- `V2MOM` - Vision, Values, Methods, Obstacles, Measures
- `Roadmap` - Time-phased delivery plan
- `Initiative` - Specific improvement project
- `PRD/MRD/TRD` - Requirements documents
- `DMAIC` - Six Sigma process improvement (migrated from structured-goals)

**Capability-to-Initiative Mapping:**
```json
{
  "initiative": "Implement SAST pipeline",
  "capabilityId": "sast",
  "maturityTarget": {
    "from": "M2",
    "to": "M4"
  },
  "phase": "Q2-2026",
  "keyResults": [
    "100% of repos have SAST enabled",
    "<24h remediation SLA for critical findings"
  ]
}
```

### prism (Orchestrator)

**Purpose:** Connect all modules with unified queries, cross-references, and workflows.

**Key Features:**

1. **Unified Document Loading**
   ```go
   ecosystem, err := prism.LoadEcosystem(prism.EcosystemConfig{
       CapabilityStack: "capabilities/security.json",
       MaturityModel:   "maturity/security-model.json",
       MaturityState:   "maturity/security-state.json",
       Roadmap:         "plans/security-roadmap.json",
   })
   ```

2. **Cross-Module Queries**
   ```go
   // Find capabilities below target maturity
   gaps := ecosystem.CapabilitiesWithMaturityGap()

   // Find initiatives for a capability
   initiatives := ecosystem.InitiativesForCapability("sast")

   // Get full context for a capability
   context := ecosystem.CapabilityContext("sast")
   // Returns: capability + current maturity + target + initiatives + dependencies
   ```

3. **Gap Analysis**
   ```go
   analysis := ecosystem.AnalyzeGaps()
   // Returns prioritized list of capabilities needing investment
   // Sorted by: strategic importance × maturity gap × dependency impact
   ```

4. **Dashboard Generation**
   ```go
   dashboard := ecosystem.GenerateDashboard(prism.DashboardOptions{
       Format: "dashforge",
       Widgets: []string{"maturity-heatmap", "initiative-timeline", "gap-chart"},
   })
   ```

5. **Validation Across Documents**
   ```go
   errs := ecosystem.Validate()
   // Checks:
   // - All prismRef.sliIds exist in MaturityModel
   // - All initiative.capabilityId exist in CapabilityStack
   // - No orphaned measurements in MaturityState
   ```

## Data Flow

### Strategy to Execution

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  Strategy   │────▶│ Capability  │────▶│  Maturity   │────▶│  Roadmap    │
│  Priorities │     │  Importance │     │    Gaps     │     │ Initiatives │
└─────────────┘     └─────────────┘     └─────────────┘     └─────────────┘
```

### Execution to Outcomes

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  Roadmap    │────▶│ Initiative  │────▶│  Maturity   │────▶│  Strategic  │
│  Delivery   │     │  Completion │     │  Increase   │     │  Outcomes   │
└─────────────┘     └─────────────┘     └─────────────┘     └─────────────┘
```

## CLI Integration

Each module has its own CLI, plus unified commands in `prism`:

```bash
# Module CLIs
prism-capability validate capabilities.json
prism-maturity assess maturity-state.json
prism-roadmap show roadmap.json

# Unified prism CLI
prism ecosystem load --config ecosystem.yaml
prism gaps analyze --sort-by=impact
prism dashboard generate --format=html
prism validate --all
```

## File Organization

Recommended project structure:

```
my-org-prism/
├── capabilities/
│   ├── security.json
│   ├── platform.json
│   └── data.json
├── maturity/
│   ├── model.json          # SLI definitions, level thresholds
│   └── state.json          # Current measurements
├── plans/
│   ├── okrs/
│   │   └── fy2026-q2.json
│   ├── roadmaps/
│   │   └── security-roadmap.json
│   └── initiatives/
│       └── sast-pipeline.json
├── requirements/
│   ├── prd/
│   ├── mrd/
│   └── trd/
└── prism.yaml              # Ecosystem configuration
```

## Migration Path

### From Existing Repos

| Old Location | New Location | Action |
|--------------|--------------|--------|
| `plexusone/capability-stack-spec` | `grokify/prism-capability` | ✅ Done (v0.1.0 tagged) |
| `grokify/prism` | `grokify/prism-maturity` | ✅ Done |
| `grokify/structured-plan` | `grokify/prism-roadmap` | ✅ Done |
| `grokify/structured-goals` | Merge DMAIC → prism-roadmap | Pending |
| New | `grokify/prism` | This repo |

### Module Path Updates

Each module needs `go.mod` path updates:

```go
// prism-capability/go.mod
module github.com/grokify/prism-capability

// prism-maturity/go.mod
module github.com/grokify/prism-maturity

// prism-roadmap/go.mod
module github.com/grokify/prism-roadmap

// prism/go.mod
module github.com/grokify/prism

require (
    github.com/grokify/prism-capability v0.1.0
    github.com/grokify/prism-maturity v0.x.0
    github.com/grokify/prism-roadmap v0.x.0
)
```

## Implementation Phases

### Phase 1: Foundation (Current)
- [x] Tag prism-capability v0.1.0
- [x] Rename repos to prism-* naming
- [ ] Create prism orchestrator repo structure
- [ ] Update go.mod paths in all modules
- [ ] Create this design document

### Phase 2: Integration Types
- [ ] Define shared types in prism (or prism-core)
- [ ] Add cross-reference validation
- [ ] Implement EcosystemLoader

### Phase 3: Cross-Module Queries
- [ ] CapabilityContext aggregation
- [ ] Gap analysis engine
- [ ] Initiative-to-capability tracing

### Phase 4: Unified CLI
- [ ] `prism ecosystem` commands
- [ ] `prism gaps` analysis
- [ ] `prism dashboard` generation

### Phase 5: Visualization
- [ ] Unified rendering (D2, HTML, dashforge)
- [ ] Maturity heatmaps
- [ ] Initiative timelines
- [ ] Gap charts

## Related Documents

- [prism-capability README](https://github.com/grokify/prism-capability)
- [prism-maturity README](https://github.com/grokify/prism-maturity)
- [prism-roadmap README](https://github.com/grokify/prism-roadmap)
