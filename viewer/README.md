# @grokify/prism

[![npm version](https://img.shields.io/npm/v/@grokify/prism.svg)](https://www.npmjs.com/package/@grokify/prism)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/grokify/prism/blob/main/LICENSE)

TypeScript library for PRISM capability and maturity visualization.

Part of the [PRISM ecosystem](https://github.com/grokify/prism) - **P**latform for **R**eliability, **I**ntelligence, **S**trategy & **M**aturity.

## Features

- **Zod Schemas** - Runtime validation with TypeScript type inference
- **HTML Renderers** - Framework-agnostic string rendering functions
- **Lit Components** - Interactive web components (optional)
- **Tree-shakeable** - Import only what you need via subpath exports

## Installation

```bash
npm install @grokify/prism
# or
pnpm add @grokify/prism
```

## Subpath Exports

Import specific modules for optimal bundle size:

```typescript
// Schemas only (Zod + types)
import { MaturityGridDataSchema, type MaturityGridData } from '@grokify/prism/schema';
import { JourneyRoadmapSchema, type JourneyRoadmap } from '@grokify/prism/schema/roadmap';

// HTML renderers (framework-agnostic)
import { renderDomainView, renderFrameworkView } from '@grokify/prism/html/maturity';
import { renderTimelineView, renderStoryboardView } from '@grokify/prism/html/roadmap';

// Lit components (requires Lit peer dependency)
import { MaturityGrid } from '@grokify/prism/components';
```

## Usage

### Validating Data with Zod

```typescript
import { MaturityGridDataSchema } from '@grokify/prism/schema/maturity';

const result = MaturityGridDataSchema.safeParse(jsonData);
if (result.success) {
  const data = result.data; // Fully typed MaturityGridData
} else {
  console.error(result.error.issues);
}
```

### Rendering HTML

```typescript
import { renderTimelineView } from '@grokify/prism/html/roadmap';
import type { JourneyRoadmap } from '@grokify/prism/schema/roadmap';

const roadmap: JourneyRoadmap = { /* ... */ };

// Returns HTML string
const html = renderTimelineView(roadmap, {
  showLegend: true,
  colorScheme: 'heatmap',
});

document.getElementById('container').innerHTML = html;
```

### Using Lit Components

```typescript
import '@grokify/prism/components';

// In your HTML or template:
// <maturity-grid theme="dark" view="by-layer"></maturity-grid>
```

```html
<maturity-grid theme="dark" view="by-layer">
  <script type="application/json">
    {
      "title": "Security Capabilities",
      "layers": [...],
      "capabilities": [...]
    }
  </script>
</maturity-grid>
```

Or load data from a URL:

```html
<maturity-grid
  src="/api/capabilities.json"
  theme="light"
  view="by-category"
  show-legend
  show-view-toggle
></maturity-grid>
```

## API Reference

### Schemas

| Export Path | Schemas |
|-------------|---------|
| `@grokify/prism/schema/maturity` | `LayerSchema`, `CategorySchema`, `CapabilitySchema`, `MaturityGridDataSchema` |
| `@grokify/prism/schema/roadmap` | `JourneyRoadmapSchema`, `CapabilityJourneySchema`, `DependencySchema`, `TeamSchema`, `InitiativeSchema`, and 25+ more |

### HTML Renderers

| Export Path | Functions |
|-------------|-----------|
| `@grokify/prism/html/maturity` | `renderDomainView()`, `renderFrameworkView()` |
| `@grokify/prism/html/roadmap` | `renderTimelineView()`, `renderStoryboardView()`, `renderDependencyView()`, `renderTeamView()` |

### Components

| Export Path | Components |
|-------------|------------|
| `@grokify/prism/components` | `<maturity-grid>` |

### Styles

```typescript
import { ROADMAP_STYLES_PATH, INLINE_STYLES } from '@grokify/prism/styles';
```

Include the CSS file for HTML renderers:

```html
<link rel="stylesheet" href="node_modules/@grokify/prism/dist/styles/prism-roadmap.css">
```

## TypeScript Support

All exports include full TypeScript types. Types are inferred from Zod schemas:

```typescript
import { MaturityGridDataSchema } from '@grokify/prism/schema/maturity';
import { z } from 'zod';

// Infer type from schema
type MaturityGridData = z.infer<typeof MaturityGridDataSchema>;
```

Or import types directly:

```typescript
import type {
  MaturityGridData,
  Capability,
  Layer,
  Category,
} from '@grokify/prism/schema/maturity';
```

## Related Packages

This is the TypeScript/JavaScript package. For Go:

- [github.com/grokify/prism](https://github.com/grokify/prism) - Go CLI and library
- [github.com/grokify/prism-core](https://github.com/grokify/prism-core) - Shared Go primitives
- [github.com/grokify/prism-maturity](https://github.com/grokify/prism-maturity) - Go maturity models
- [github.com/grokify/prism-roadmap](https://github.com/grokify/prism-roadmap) - Go roadmap types

## License

MIT
