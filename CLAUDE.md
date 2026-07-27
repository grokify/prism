# CLAUDE.md — prism

PRISM (Platform for Reliability, Intelligence, Strategy & Maturity) is a capability-driven organizational intelligence framework. This repo is the **orchestration layer** over the PRISM module repos.

## Architecture

```
prism (this repo)
├── cmd/prism/         # CLI entry point
├── ecosystem/         # Cross-module loading, queries, validation
├── sitegen/           # Static site generator (HTML dashboards)
└── viewer/            # @grokify/prism npm package (TypeScript/Lit)

Module repos (schemas owned there, not here):
├── prism-core         # Shared primitives (Domain, Layer, Stage, MaturityLevel)
├── prism-capability   # Capability stacks, layers, categories
├── prism-maturity     # SLIs, SLOs, maturity models, state tracking
└── prism-roadmap      # OKRs, V2MOMs, journey roadmaps, initiatives
```

## Specs

Specs are the source of truth for design decisions — read `docs/specs/` before implementing:

- `ARCHITECTURE.md` — System design, component interactions
- `PRD.md` — Product requirements, user stories
- `TRD.md` — Technical requirements, implementation details
- `PLAN.md` — Implementation phases, milestones
- `ROADMAP.md` — Feature roadmap with RMI tracking

## PRISM Control

Roadmap items are tracked in [prism-control](https://github.com/ProductBuildersHQ/prism-control). Use `prismctl work ready --repo github.com/grokify/prism` to find claimable work. Commits implementing an RMI carry the trailer `Refs: RMI-PRISM-<NNN>`.

## Common Tasks

### Update module dependencies

```bash
go get github.com/grokify/prism-maturity@latest
go get github.com/grokify/prism-capability@latest
go get github.com/grokify/prism-roadmap@latest
go mod tidy
```

### Run tests

```bash
# Go tests
go test ./...

# TypeScript tests (viewer/)
cd viewer && pnpm test
```

### Lint

```bash
golangci-lint run
```

### Build CLI

```bash
go build -o prism ./cmd/prism
```

### Generate static site

```bash
prism site generate --stack=./examples/platform-team/ --output=./dist
```

## Code Patterns

### sitegen/ — HTML generation

Templates are Go `html/template` strings embedded in `site.go`. Data flows:

1. `loadStacks()` — reads stack.json, model.json, state.json, roadmap.json
2. `StackData` struct holds loaded documents + computed `MaturityAggregator`
3. `generate*Pages()` methods render HTML via templates
4. Lit components via `--prism-ui-js` flag inject JSON into `<script type="application/json">`

When adding a new page type:
1. Add field to `StackData` if loading new document type
2. Create `generate*Pages()` method
3. Add template const (e.g., `roadmapPageTemplate`)
4. Call from `Generate()` method

### viewer/ — TypeScript/Lit components

Structure:
```
viewer/
├── src/
│   ├── components/    # Lit web components
│   ├── schema/        # Zod schemas + TypeScript types
│   ├── html/          # HTML string renderers
│   └── styles/        # CSS stylesheets
├── vitest.config.ts   # Test configuration
└── package.json       # @grokify/prism package
```

Component pattern:
```typescript
@customElement('component-name')
export class ComponentName extends LitElement {
  static override styles = css`...`;
  
  @property({ type: String }) theme: 'light' | 'dark' = 'light';
  @state() private data: DataType | null = null;
  
  // Load JSON from <script type="application/json"> child
  override firstUpdated() {
    const script = this.querySelector('script[type="application/json"]');
    if (script) this.data = JSON.parse(script.textContent || '{}');
  }
  
  override render() {
    return html`...`;
  }
}
```

### ecosystem/ — Cross-module queries

The `Ecosystem` type loads all document types and provides unified queries:

```go
eco := ecosystem.Load(config)
eco.AllCapabilities()           // From all stacks
eco.GetCapabilityContext(id)    // Capability + linked SLIs + initiatives
eco.Validate()                  // Cross-module reference checks
eco.Stats()                     // Aggregate statistics
```

## Release Process

1. Update dependencies if needed
2. Run tests: `go test ./...` and `cd viewer && pnpm test`
3. Run lint: `golangci-lint run`
4. Update `CHANGELOG.json` with new version entry
5. Generate `CHANGELOG.md`: `schangelog generate CHANGELOG.json -o CHANGELOG.md`
6. Create release notes: `docs/releases/vX.Y.Z.md`
7. Update `mkdocs.yml` navigation
8. Push and verify CI passes
9. Tag: `git tag vX.Y.Z && git push origin vX.Y.Z`

## Module Boundaries

**This repo does:**
- Orchestrate loading documents from all module types
- Cross-module queries and validation
- Static site generation
- CLI integrating all module commands
- TypeScript visualization components

**This repo does NOT:**
- Own schemas (they live in module repos)
- Define SLI/SLO types (prism-maturity)
- Define capability stack structure (prism-capability)
- Define OKR/roadmap types (prism-roadmap)

When a schema change is needed, it lands in the module repo first, gets tagged, then this repo updates the dependency.

## Testing Guidance

- Go: `go test ./...` — ecosystem package has integration tests
- TypeScript: `cd viewer && pnpm test` — vitest with jsdom
- Manual: `prism site generate --stack=./examples/platform-team/ --output=./dist` then open `dist/index.html`

## Troubleshooting

**"undefined: ..." errors in sitegen**: Templates reference variables from `StackData`. Check that the field exists and is populated in `loadStacks()`.

**Lit component not rendering**: Check that JSON is valid in `<script type="application/json">`. The component's `firstUpdated()` parses it.

**Cross-module validation failures**: Usually a missing or renamed ID. Check that capability IDs match SLI references in maturity documents.
