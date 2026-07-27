# PRISM Technical Requirements Document

## Repository Layout

```text
prism/
├── ecosystem/        Ecosystem loader, cross-module queries, validation, stats
├── cmd/prism/        CLI (Cobra): ecosystem, site, stats, validate
├── sitegen/          Static site generator (Go templates)
├── ui/               Lit web components (@prism/ui, Vite + TypeScript)
├── viewer/           TypeScript viewer package
├── examples/         Example ecosystem (platform-team)
└── docs/             MkDocs documentation
```

## Module Dependencies

This repo orchestrates; the schemas live in the module repos:

| Dependency | Provides |
|---|---|
| `github.com/grokify/prism-core` | Domain, Layer, Stage, MaturityLevel, TeamType primitives |
| `github.com/grokify/prism-capability` | `CapabilityStack`, capabilities, layers, dependencies |
| `github.com/grokify/prism-maturity` | `PRISMDocument`, metrics (SLI/SLO), services, initiatives, maturity state |
| `github.com/grokify/prism-roadmap` | `okr.OKRSet`, `roadmap.Roadmap`, V2MOM, opportunity specs |

Dependency versions must be verified against the latest tags before upgrading (see global dependency-verification workflow). Local `replace` directives (currently one pointing at `../prism-roadmap`) must be removed before pushing.

## Ecosystem Package

`ecosystem.Config` declares file sources per module:

```go
type Config struct {
    Name       string
    Capability CapabilityConfig // capability stack files
    Maturity   MaturityConfig   // maturity document files
    Roadmap    RoadmapConfig    // OKR and roadmap files
}
```

`ecosystem.Load` reads all files into an `Ecosystem` holding `CapabilityStacks`, `PRISMDocuments`, `OKRSets`, and `Roadmaps`.

### Query Surface

| Method group | Examples |
|---|---|
| Capabilities | `AllCapabilities`, `GetCapabilityByID`, `CapabilitiesByStatus`, `CapabilitiesByDomain` |
| Maturity | `AllMetrics`, `GetMetricByID`, `AllServices`, `AllInitiatives` |
| Roadmap | `AllObjectives`, `AllPhases`, `GetPhaseByID` |
| Joins | `GetCapabilityContext` — one capability with its metrics, objectives, and roadmap items |
| Integrity | `Validate() ValidationErrors`, `Stats()` |

### Requirements

- Loading must fail with a wrapped error naming the offending file; never silently skip.
- Validation must detect dangling cross-document references (capability IDs in OKRs/roadmaps that do not exist, metric IDs not defined, etc.).
- Queries are read-only over loaded documents; no mutation of source files.

## Cross-Repo Reference Format

References to entities outside PRISM use typed string IDs:

```text
capability:scim-group-push
market:identity-governance
customer:acme
signal:oauth-idp-compatibility
```

Planned: a resolver interface so evidence links can be validated against OmniSignal exports and MarketSpec/OrganizationSpec entity lists without hard dependencies on those repos.

## Scoring Engine (Planned)

Multi-dimensional prioritization computes per-initiative scores along fixed dimensions (customer value, market opportunity, strategic alignment, engineering cost, risk, adoption evidence, confidence) with configurable weights:

- Formulas are data (configuration), not code — organizations can tune weights without forking.
- Output includes the contributing inputs so every score is explainable.
- Pillar classification (CSAT / SAM-SOM / TAM) is a tagged contribution level per pillar, not a single category.

## Schema Approach

Schemas follow the Go-first workflow: Go structs (in module repos) are the source of truth; JSON Schemas are generated via `github.com/invopop/jsonschema`, linted with `schemago lint`, and embedded via `//go:embed`. This repo does not hand-write schemas.

## Visualization Stack

| Component | Technology | Output |
|---|---|---|
| `sitegen/` | Go templates | Static capability site with maturity badges, metrics-count indicators, light/dark themes |
| `ui/` | Lit + Vite + TypeScript | `@prism/ui` web components embedded in generated sites |
| `viewer/` | TypeScript + Vitest | Standalone document viewer |

Generated sites must remain self-contained static output deployable to GitHub Pages.

## CLI

Cobra-based `prism` binary:

| Command | Function |
|---|---|
| `prism ecosystem load --config prism.yaml` | Load and report an ecosystem |
| `prism validate --all` | Cross-document validation |
| `prism stats` | Ecosystem statistics |
| `prism site generate --stack=./stacks/ --output=./dist` | Static site generation |
| `prism gaps analyze --sort-by=impact` | Maturity gap analysis |

## Error Handling and Testing

- Follow the standard error-handling priority: return wrapped errors; panic only on invariant violations; `t.Fatal` in tests.
- `ecosystem` package has unit tests (`ecosystem_test.go`); new query and validation functions require table-driven tests with fixture documents under `examples/`.
- CI: build, `golangci-lint`, `go test ./...` must pass before tagging releases.
