# PRISM Implementation Plan

Phased plan for evolving the orchestration layer and the PRISM-Roadmap decision capabilities. Schema changes land in the module repos (prism-core, prism-capability, prism-maturity, prism-roadmap); this repo integrates them.

## Phase 1: Consolidate the Capability Model (Current)

Status: largely complete.

- Ecosystem loader over capability stacks, maturity documents, OKR sets, and roadmaps.
- Cross-module queries and `GetCapabilityContext` join.
- Referential validation and stats.
- Static site generation with maturity badges; `@prism/ui` Lit components.

Remaining:

- Harden `prism gaps analyze` (impact-sorted gap analysis) with tests and documented output format.
- Remove the local `replace` directive on prism-roadmap once its required version is tagged.

## Phase 2: Roadmap Artifact Coverage

Extend prism-roadmap (module repo) and integrate here:

- V2MOM documents alongside OKR sets in `ecosystem.Config`.
- Opportunity Spec loading and cross-referencing to capabilities.
- Business Model Canvas documents.
- Validation rules for each new artifact type (dangling capability/metric references).

Exit criteria: `prism validate --all` covers OKR, V2MOM, roadmap, opportunity spec, and BMC documents.

## Phase 3: Pillar Classification

- Add CSAT / SAM-SOM / TAM contribution levels to initiatives and roadmap items (schema change in prism-roadmap).
- Portfolio rollup query: investment distribution across pillars.
- Pillar view in generated sites and dashboards.

Exit criteria: an executive dashboard shows the pillar mix (e.g. 60% / 25% / 15%) computed from tagged initiatives.

## Phase 4: OmniSignal Evidence Consumption

- Define the evidence-link format: roadmap items reference canonical signals by typed ID (`signal:oauth-idp-compatibility`).
- Resolver interface to validate evidence links against OmniSignal exports without a hard dependency.
- Per-initiative evidence summary in `GetCapabilityContext` and site output: supporting signal count, frustration score, ARR affected, named customers.

Exit criteria: a roadmap item's page lists its supporting evidence with source counts.

## Phase 5: Multi-Pillar Scoring Engine

- Configurable scoring across dimensions: customer value, market opportunity, strategic alignment, engineering cost, risk, adoption evidence, confidence.
- Weights as configuration, not code.
- Explainable output: every score ships with its contributing inputs.
- Ranked initiative lists per pillar and overall.

Exit criteria: `prism prioritize` produces a ranked, explainable initiative list from evidence links and configured weights.

## Phase 6: Portfolio Dashboard

- Unified dashboard: capability maturity heatmap, pillar investment mix, top initiatives with evidence, gap analysis.
- Interactive filtering via `@prism/ui` components.
- Export to HTML for distribution.

## Sequencing Notes

- Phases 2 and 3 are schema-first: land types in prism-roadmap, tag a release, then integrate here (no local replace directives at push time).
- Phase 4 depends on OmniSignal defining its canonical-signal export format; coordinate the typed-ID contract before implementation.
- Each phase follows the repo release flow: commits pushed, CI green, then tag.
