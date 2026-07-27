# PRISM — Roadmap

**Initiative:** `INIT-PRISM-001`
**Repository:** `github.com/grokify/prism`
**Status:** Proposed — 6 phases, 19 RMIs

> RMI IDs are stable and permanent. Commits implementing an item carry the trailer `Refs: RMI-PRISM-NNN`. Phase status is derived from member RMIs — a phase is complete only when all its required RMIs are complete. Schema changes land in the module repos (prism-core, prism-capability, prism-maturity, prism-roadmap); this repo orchestrates and integrates them. See [PLAN.md](PLAN.md) for phase rationale and exit criteria.

## Phase 1 — Capability Baseline Hardening

**Theme:** Finish the consolidation phase; clean module dependencies.
**Status:** Not started — 0 of 3 items completed

- [ ] `RMI-PRISM-001` Harden `prism gaps analyze`: impact-sorted gap analysis with unit tests and documented output format
- [ ] `RMI-PRISM-002` Remove local `replace github.com/grokify/prism-roadmap => ../prism-roadmap` from go.mod
  - Acceptance: required prism-roadmap version is tagged and released; `go.mod` has no local replace directives at push time
- [ ] `RMI-PRISM-003` Document the ecosystem loader configuration and cross-module query surface (`AllCapabilities`, `GetCapabilityContext`, `Validate`, `Stats`)

## Phase 2 — Roadmap Artifact Coverage

**Theme:** Load and validate every strategic planning artifact type.
**Status:** Not started — 0 of 4 items completed

- [ ] `RMI-PRISM-004` V2MOM document loading alongside OKR sets in `ecosystem.Config`
  - Acceptance: V2MOM schema lands in prism-roadmap (module repo) and is tagged before integration here
- [ ] `RMI-PRISM-005` Opportunity Spec loading with cross-references to capabilities
  - Acceptance: Opportunity Spec schema lands in prism-roadmap (module repo) and is tagged before integration here
- [ ] `RMI-PRISM-006` Business Model Canvas document loading
  - Acceptance: BMC schema lands in prism-roadmap (module repo) and is tagged before integration here
- [ ] `RMI-PRISM-007` Validation rules for new artifact types: dangling capability and metric references
  - Depends on: `RMI-PRISM-004`, `RMI-PRISM-005`, `RMI-PRISM-006`
  - Acceptance: `prism validate --all` covers OKR, V2MOM, roadmap, opportunity spec, and BMC documents

## Phase 3 — Pillar Classification

**Theme:** Classify investment by CSAT / SAM-SOM / TAM strategic pillars.
**Status:** Not started — 0 of 3 items completed

- [ ] `RMI-PRISM-008` CSAT / SAM-SOM / TAM contribution levels on initiatives and roadmap items
  - Acceptance: schema change lands in prism-roadmap (module repo); an initiative can contribute to multiple pillars at different levels
- [ ] `RMI-PRISM-009` Portfolio rollup query: investment distribution across pillars
  - Depends on: `RMI-PRISM-008`
- [ ] `RMI-PRISM-010` Pillar view in generated sites and dashboards
  - Depends on: `RMI-PRISM-009`
  - Acceptance: an executive dashboard shows the pillar mix (e.g. 60% / 25% / 15%) computed from tagged initiatives

## Phase 4 — OmniSignal Evidence Consumption

**Theme:** Link roadmap items to canonical signals as supporting evidence.
**Status:** Not started — 0 of 3 items completed

- [ ] `RMI-PRISM-011` Evidence-link format: roadmap items reference canonical signals by typed ID (`signal:oauth-idp-compatibility`)
  - Acceptance: typed-ID contract coordinated with OmniSignal's canonical-signal export format before implementation
- [ ] `RMI-PRISM-012` Resolver interface validating evidence links against OmniSignal exports without a hard dependency
  - Depends on: `RMI-PRISM-011`
- [ ] `RMI-PRISM-013` Per-initiative evidence summary in `GetCapabilityContext` and site output: supporting signal count, frustration score, ARR affected, named customers
  - Depends on: `RMI-PRISM-012`
  - Acceptance: a roadmap item's page lists its supporting evidence with source counts

## Phase 5 — Multi-Pillar Scoring Engine

**Theme:** Explainable, configurable prioritization from evidence.
**Status:** Not started — 0 of 3 items completed

- [ ] `RMI-PRISM-014` Configurable scoring across dimensions: customer value, market opportunity, strategic alignment, engineering cost, risk, adoption evidence, confidence
  - Depends on: `RMI-PRISM-008`, `RMI-PRISM-011`
  - Acceptance: weights are configuration, not code
- [ ] `RMI-PRISM-015` Explainable score output: every score ships with its contributing inputs
  - Depends on: `RMI-PRISM-014`
- [ ] `RMI-PRISM-016` `prism prioritize`: ranked initiative lists per pillar and overall
  - Depends on: `RMI-PRISM-015`
  - Acceptance: produces a ranked, explainable initiative list from evidence links and configured weights

## Phase 6 — Portfolio Dashboard

**Theme:** One view of maturity, investment mix, and evidence.
**Status:** Not started — 0 of 3 items completed

- [ ] `RMI-PRISM-017` Unified dashboard: capability maturity heatmap, pillar investment mix, top initiatives with evidence, gap analysis
  - Depends on: `RMI-PRISM-010`, `RMI-PRISM-013`, `RMI-PRISM-016`
- [ ] `RMI-PRISM-018` Interactive filtering via `@prism/ui` components
  - Depends on: `RMI-PRISM-017`
- [ ] `RMI-PRISM-019` Self-contained HTML export for distribution
  - Depends on: `RMI-PRISM-017`

## Long-Term Direction (unscheduled)

Beyond the phases above, held as direction rather than RMIs:

- AI-driven investment recommendations: given capability gaps, evidence, and pillar targets, propose portfolio adjustments with cited reasoning.
- Continuous re-prioritization as new OmniSignal evidence arrives.
- What-if analysis across investment mixes and scenario comparison with traceable evidence chains.

## Guiding Constraints

- Every ranking stays explainable — evidence in, reasoning out; no opaque scores.
- Schemas remain owned by module repos; this repo orchestrates.
- Documents remain plain JSON/YAML in git; static-site outputs remain self-contained.
