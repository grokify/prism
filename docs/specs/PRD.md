# PRISM Product Requirements Document

## Product Summary

PRISM is a capability-driven organizational intelligence framework. This repository is its orchestration layer: it unifies capability stacks (prism-capability), maturity models and state (prism-maturity), and strategic planning artifacts (prism-roadmap) into cross-referenced views, validation, and dashboards.

Within the ProductContext ecosystem, PRISM-Roadmap is the decision layer — it answers "given all available evidence, what should we build next?" using evidence from OmniSignal and entities from MarketSpec and OrganizationSpec.

## Personas

| Persona | Needs |
|---|---|
| Executive | Portfolio-level visibility: where is investment going across CSAT, SAM/SOM, and TAM pillars; are capability maturity targets on track |
| Product leader | Evidence-based prioritization: which initiatives are supported by which signals; explainable rankings instead of opaque scores |
| Engineering manager | Capability maturity assessment: current state vs. target per capability; gap analysis; what work closes which gaps |
| Program / strategy team | Consistent planning artifacts: OKRs, V2MOM, roadmaps, opportunity specs that cross-reference cleanly |

## Use Cases

1. **Capability assessment** — Load capability stacks and maturity state; identify gaps between current and target maturity per capability, sorted by impact.

2. **Evidence-based roadmap prioritization** — Link roadmap items to capabilities and to supporting evidence (customer requests, competitive gaps, analyst findings). Every priority is explainable: "ranked high because 20 enterprise customers, 2 competitor gaps, aligned with OKR-3."

3. **Portfolio allocation across pillars** — Classify initiatives by CSAT / SAM-SOM / TAM contribution and report the investment mix (e.g. 60% / 25% / 15%), enabling deliberate rebalancing.

4. **Cross-module validation** — Verify that OKRs reference real capabilities, roadmap items reference real metrics, and all cross-document IDs resolve.

5. **Dashboard and site generation** — Generate static sites and HTML dashboards from capability stacks with maturity badges, metric counts, and theme support.

## Requirements

### Functional

| ID | Requirement | Status |
|---|---|---|
| F1 | Load ecosystem configuration (capability stacks, maturity docs, OKR sets, roadmaps) from YAML/JSON | Implemented |
| F2 | Cross-module queries: capabilities, metrics, services, initiatives, objectives, phases by ID/status/domain | Implemented |
| F3 | Capability context: for one capability, join its maturity state, metrics, objectives, and roadmap items | Implemented |
| F4 | Referential validation across all loaded documents | Implemented |
| F5 | Ecosystem statistics rollup | Implemented |
| F6 | Static site generation with maturity badges and Lit components | Implemented |
| F7 | Gap analysis sorted by impact | Partial (CLI surface exists) |
| F8 | Pillar classification (CSAT / SAM-SOM / TAM) on initiatives and portfolio rollup | Planned |
| F9 | Evidence links from roadmap items to OmniSignal canonical signals | Planned |
| F10 | Multi-dimensional scoring (customer value, market opportunity, strategic alignment, cost, risk, adoption, confidence) | Planned |

### Non-Functional

- Documents are plain JSON/YAML files in git; no database required.
- All cross-repo references use typed IDs (`capability:x`, `market:y`, `customer:z`).
- Scores must be explainable — every ranking traceable to its contributing evidence and weights.
- Schemas are owned by the module repos; this repo consumes them via Go module dependencies.

## Success Criteria

- A roadmap review can answer "why is this ranked here?" from evidence links alone.
- Executives can see pillar-level investment distribution without manual spreadsheet work.
- Validation catches dangling references before documents are published.
- New signal sources (via OmniSignal) require no changes to the prioritization model.
