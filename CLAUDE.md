# CLAUDE.md — prism

PRISM (Platform for Reliability, Intelligence, Strategy & Maturity) is a capability-driven organizational intelligence framework. This repo is the orchestration layer over the prism module repos (prism-core, prism-capability, prism-maturity, prism-roadmap): it loads their documents, runs cross-module queries and validation, and generates static sites. Schemas are owned by the module repos, not here.

Specs are the source of truth for design decisions — read `docs/specs/` (ARCHITECTURE.md, PRD.md, TRD.md, PLAN.md, ROADMAP.md) before implementing.

## PRISM Control

This repo's roadmap items are tracked in [prism-control](https://github.com/ProductBuildersHQ/prism-control). Use `prismctl work ready --repo github.com/grokify/prism` to find claimable work, and carry the `Refs: RMI-PRISM-<NNN>` trailer on every commit.
