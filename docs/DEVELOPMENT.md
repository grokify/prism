# Development Guide

This guide covers the architecture, code patterns, and development workflows for the PRISM orchestrator.

## Architecture Overview

```
prism/
├── cmd/prism/           # CLI entry point
│   └── cmd/             # Cobra command implementations
├── ecosystem/           # Cross-module loading and queries
├── sitegen/             # Static site generator
├── viewer/              # @grokify/prism npm package
│   ├── src/components/  # Lit web components
│   ├── src/schema/      # Zod schemas
│   ├── src/html/        # HTML string renderers
│   └── src/styles/      # CSS stylesheets
├── examples/            # Example capability stacks
└── docs/                # MkDocs documentation
```

### Module Dependencies

```
prism (orchestrator)
    ├── prism-capability  # Capability stack types
    ├── prism-maturity    # SLI/SLO, maturity model types
    ├── prism-roadmap     # OKR, V2MOM, journey roadmap types
    └── prism-core        # Shared primitives
```

Schemas are owned by module repos. This repo imports and orchestrates them.

## Go Development

### ecosystem/ Package

The `ecosystem` package loads documents from all module types and provides cross-module queries.

**Key Types:**

```go
// Config defines what to load
type Config struct {
    CapabilityStacks []CapabilityStackConfig
    Maturity         MaturityConfig
    Roadmap          RoadmapConfig
}

// Ecosystem holds loaded documents
type Ecosystem struct {
    Stacks   []*capstack.CapabilityStack
    Maturity *prism.PRISMDocument
    OKRSets  []*roadmap.OKRSet
    Roadmaps []*journey.JourneyRoadmap
}
```

**Cross-Module Queries:**

```go
eco.AllCapabilities()           // All capabilities from all stacks
eco.GetCapabilityContext(id)    // Capability + linked SLIs + initiatives
eco.AllMetrics()                // All metrics from maturity documents
eco.AllInitiatives()            // All initiatives from roadmaps
eco.Validate()                  // Cross-reference validation
eco.Stats()                     // Aggregate statistics
```

### sitegen/ Package

Generates static HTML sites from capability stacks with maturity overlays.

**Data Flow:**

1. `loadStacks()` — Discovers and loads documents from standard directories
2. `StackData` — Holds loaded documents per capability stack
3. `Generate()` — Orchestrates page generation
4. `generate*Pages()` — Render specific page types

**Standard Directory Structure:**

```
{stack-name}/
├── stack.json      # prism-capability CapabilityStack
├── model.json      # prism-maturity Spec (SLIs, criteria)
├── state.json      # prism-maturity PRISMDocument (current values)
└── roadmap.json    # prism-roadmap JourneyRoadmap
```

**Adding a New Page Type:**

1. Add field to `StackData` if loading a new document type:
   ```go
   type StackData struct {
       // ...existing fields
       NewDoc  *newpkg.NewType
   }
   ```

2. Load in `loadStacks()`:
   ```go
   if src.newDocPath != "" {
       data, _ := os.ReadFile(src.newDocPath)
       var doc newpkg.NewType
       json.Unmarshal(data, &doc)
       sd.NewDoc = &doc
   }
   ```

3. Create page generator:
   ```go
   func (g *Generator) generateNewPages() error {
       for _, sd := range g.stacks {
           if sd.NewDoc == nil {
               continue
           }
           // Create directory, render template
       }
       return nil
   }
   ```

4. Add template:
   ```go
   const newPageTemplate = `<!DOCTYPE html>...`
   ```

5. Call from `Generate()`:
   ```go
   if err := g.generateNewPages(); err != nil {
       return fmt.Errorf("generating new pages: %w", err)
   }
   ```

**Template Data:**

Templates receive typed data structs. Use `template.JS()` for JSON embedding (with nolint for gosec G203 when data is trusted).

### CLI (cmd/prism/)

Uses Cobra for command structure. Commands are organized by module:

```
prism
├── capability   # From prism-capability
├── maturity     # From prism-maturity  
├── roadmap      # From prism-roadmap
├── ecosystem    # Cross-module operations
├── site         # Static site generation
└── validate     # Document validation
```

**Adding a Command:**

1. Create file in `cmd/prism/cmd/`:
   ```go
   var newCmd = &cobra.Command{
       Use:   "new",
       Short: "Description",
       RunE:  runNew,
   }
   
   func init() {
       rootCmd.AddCommand(newCmd)
   }
   ```

## TypeScript Development

### viewer/ Package Structure

```
viewer/
├── src/
│   ├── components/       # Lit web components
│   │   ├── maturity-grid.ts
│   │   ├── roadmap-timeline.ts
│   │   ├── initiative-tracker.ts
│   │   └── index.ts      # Exports
│   ├── schema/           # Zod schemas + types
│   │   ├── maturity/
│   │   └── roadmap/
│   ├── html/             # HTML string renderers
│   └── styles/           # CSS
├── demo/                 # Development demos
├── vitest.config.ts
└── package.json
```

### Zod Schemas

Schemas define runtime validation and TypeScript types:

```typescript
// schema/maturity/index.ts
import { z } from 'zod';

export const CapabilitySchema = z.object({
  id: z.string(),
  name: z.string(),
  layerId: z.string(),
  categoryId: z.string(),
  status: z.enum(['planned', 'building', 'operational', 'deprecated']),
});

export type Capability = z.infer<typeof CapabilitySchema>;
```

### Lit Components

Pattern for data-driven components:

```typescript
import { LitElement, html, css, PropertyValues } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';
import { DataSchema, type DataType } from '../schema/index.js';

@customElement('component-name')
export class ComponentName extends LitElement {
  static override styles = css`
    :host { display: block; }
    /* Theme-aware styles */
    :host([theme="dark"]) { --bg: #0f172a; }
  `;

  @property({ type: String }) theme: 'light' | 'dark' = 'light';
  @property({ type: String }) src?: string;
  
  @state() private data: DataType | null = null;
  @state() private loading = false;
  @state() private error: string | null = null;

  override async firstUpdated() {
    // Load from inline script
    const script = this.querySelector('script[type="application/json"]');
    if (script) {
      try {
        const parsed = JSON.parse(script.textContent || '{}');
        this.data = DataSchema.parse(parsed);
      } catch (e) {
        this.error = `Parse error: ${e}`;
      }
      return;
    }
    
    // Load from src URL
    if (this.src) {
      await this.loadFromUrl();
    }
  }

  private async loadFromUrl() {
    this.loading = true;
    try {
      const res = await fetch(this.src!);
      const json = await res.json();
      this.data = DataSchema.parse(json);
    } catch (e) {
      this.error = `Fetch error: ${e}`;
    } finally {
      this.loading = false;
    }
  }

  override render() {
    if (this.loading) return html`<div>Loading...</div>`;
    if (this.error) return html`<div class="error">${this.error}</div>`;
    if (!this.data) return html`<div>No data</div>`;
    
    return html`
      <div class="container">
        ${this.renderContent()}
      </div>
    `;
  }

  private renderContent() {
    // Render data
  }
}
```

### HTML Renderers

Framework-agnostic functions returning HTML strings:

```typescript
// html/maturity/index.ts
import type { MaturityGridData } from '../../schema/maturity/index.js';

export interface RenderOptions {
  theme?: 'light' | 'dark';
  showLegend?: boolean;
}

export function renderDomainView(
  data: MaturityGridData,
  options: RenderOptions = {}
): string {
  const { theme = 'light', showLegend = true } = options;
  
  return `
    <div class="maturity-grid ${theme}">
      ${showLegend ? renderLegend() : ''}
      ${data.layers.map(layer => renderLayer(layer, data)).join('')}
    </div>
  `;
}
```

### Testing

Vitest with jsdom for browser environment:

```typescript
// components/components.test.ts
import { describe, it, expect, beforeEach } from 'vitest';
import { fixture, html } from '@open-wc/testing-helpers';
import './maturity-grid.js';
import type { MaturityGrid } from './maturity-grid.js';

describe('MaturityGrid', () => {
  let element: MaturityGrid;

  beforeEach(async () => {
    element = await fixture(html`
      <maturity-grid theme="dark">
        <script type="application/json">
          {"layers": [], "capabilities": []}
        </script>
      </maturity-grid>
    `);
  });

  it('renders with dark theme', () => {
    expect(element.theme).toBe('dark');
  });
});
```

## Testing Strategy

### Go Tests

```bash
go test ./...                    # All tests
go test ./ecosystem/...          # Package tests
go test -v ./... -run TestName   # Specific test
```

### TypeScript Tests

```bash
cd viewer
pnpm test              # Run once
pnpm test:watch        # Watch mode
pnpm test:coverage     # Coverage report
```

### Manual Testing

```bash
# Generate site from example
prism site generate --stack=./examples/platform-team/ --output=./dist

# Serve and view
cd dist && python3 -m http.server 8080
```

## Debugging

### Go

Use `log.Printf` or attach debugger. For template issues, check that data structs match template variable references.

### TypeScript

Browser DevTools for Lit components. Check:
- JSON parsing in `firstUpdated()`
- Zod validation errors
- Shadow DOM rendering

### Common Issues

**"undefined" in templates:** Field missing from data struct or not populated in loader.

**Lit component blank:** JSON parse failed, check browser console.

**Cross-module validation errors:** ID mismatch between documents (capability ID vs SLI reference).

## Release Checklist

See [CONTRIBUTING.md](CONTRIBUTING.md) for the full process. Key steps:

1. Update dependencies
2. Run all tests
3. Update CHANGELOG.json
4. Generate CHANGELOG.md
5. Create release notes
6. Push and verify CI
7. Tag release
