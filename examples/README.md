# PRISM Examples

This directory contains example PRISM configurations demonstrating the full capability maturity workflow.

## Platform Team Example

The `platform-team/` directory contains a complete PRISM stack for a platform engineering team:

```
platform-team/
├── stack.json    # Capability stack definition
├── model.json    # Maturity model with SLIs
├── state.json    # Current capability levels and SLI values
└── roadmap.json  # Journey roadmap with initiatives
```

### Files

- **stack.json**: Defines 14 platform engineering capabilities across 4 layers (Developer Experience, Delivery, Observability, Infrastructure)
- **model.json**: Defines DORA-style SLIs (deployment frequency, lead time, change failure rate, MTTR) and 5 maturity levels
- **state.json**: Current capability maturity levels (M1-M5) and SLI measurements
- **roadmap.json**: Multi-quarter journey roadmap with 5 capability journeys, 5 initiatives, teams, and narrative

### Generate Static Site

To generate a static dashboard from this example:

```bash
# Build the PRISM CLI
go build -o prism ./cmd/prism

# Generate static site
./prism site generate \
  --stack=./examples/platform-team/ \
  --output=./dist \
  --title="Platform Engineering Dashboard" \
  --theme=dark

# View the generated site
open ./dist/index.html
```

### With Lit Components

To use the interactive Lit web components:

1. Build the @grokify/prism npm package:
   ```bash
   cd viewer
   pnpm install
   pnpm build
   ```

2. Generate with Lit component support:
   ```bash
   ./prism site generate \
     --stack=./examples/platform-team/ \
     --output=./dist \
     --prism-ui-js=./viewer/dist/components/index.js \
     --theme=dark
   ```

### Data Flow

```
Go Types (prism-*)    →  JSON Schema  →  JSON Files  →  Generated Site
                                                              ↓
                      TypeScript Types ← Zod Schemas ← @grokify/prism
                                                              ↓
                                              Lit Web Components
                                              (maturity-grid, roadmap-timeline,
                                               initiative-tracker)
```

### Validate Configuration

```bash
# Validate capability stack
./prism capability validate ./examples/platform-team/stack.json

# Validate maturity model
./prism maturity validate ./examples/platform-team/model.json

# Validate journey roadmap
./prism roadmap journey validate ./examples/platform-team/roadmap.json
```

## Creating Your Own

1. Copy the `platform-team/` directory as a template
2. Edit `stack.json` to define your capabilities
3. Edit `model.json` to define your maturity levels and SLIs
4. Edit `state.json` with your current measurements
5. Edit `roadmap.json` with your journey plans
6. Run `prism site generate` to create your dashboard
