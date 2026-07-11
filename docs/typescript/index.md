# TypeScript / JavaScript

The `@grokify/prism` npm package provides TypeScript-first visualization tools for PRISM capability and maturity models.

## Installation

```bash
npm install @grokify/prism
# or
pnpm add @grokify/prism
# or
yarn add @grokify/prism
```

## Package Overview

The package is organized into subpath exports for tree-shaking:

| Import Path | Description | Dependencies |
|-------------|-------------|--------------|
| `@grokify/prism/schema` | Zod schemas + TypeScript types | `zod` |
| `@grokify/prism/html` | HTML string renderers | `zod` |
| `@grokify/prism/components` | Lit web components | `zod`, `lit` |
| `@grokify/prism/styles` | CSS stylesheets | none |

## Schemas

### Maturity Grid Schema

```typescript
import {
  MaturityGridDataSchema,
  LayerSchema,
  CategorySchema,
  CapabilitySchema,
} from '@grokify/prism/schema/maturity';

// Validate JSON data
const result = MaturityGridDataSchema.safeParse(jsonData);

if (result.success) {
  const grid = result.data;
  console.log(`Loaded ${grid.capabilities.length} capabilities`);
} else {
  console.error('Validation failed:', result.error.issues);
}
```

### Journey Roadmap Schema

```typescript
import {
  JourneyRoadmapSchema,
  CapabilityJourneySchema,
  DependencySchema,
  TeamSchema,
  InitiativeSchema,
} from '@grokify/prism/schema/roadmap';

const roadmap = JourneyRoadmapSchema.parse(roadmapJson);
```

### Type Inference

Types are inferred from Zod schemas:

```typescript
import { MaturityGridDataSchema } from '@grokify/prism/schema/maturity';
import { z } from 'zod';

type MaturityGridData = z.infer<typeof MaturityGridDataSchema>;
```

Or import types directly:

```typescript
import type {
  MaturityGridData,
  Capability,
  CapabilityStatus,
  Layer,
  Category,
} from '@grokify/prism/schema/maturity';

import type {
  JourneyRoadmap,
  CapabilityJourney,
  Dependency,
  Team,
} from '@grokify/prism/schema/roadmap';
```

## HTML Renderers

Framework-agnostic functions that return HTML strings. Use with any frontend framework or vanilla JavaScript.

### Maturity Views

```typescript
import { renderDomainView, renderFrameworkView } from '@grokify/prism/html/maturity';
import type { MaturityGridData } from '@grokify/prism/schema/maturity';

const data: MaturityGridData = { /* ... */ };

// Render domain/layer view
const domainHtml = renderDomainView(data, {
  showLegend: true,
  theme: 'light',
});

// Render framework overview
const frameworkHtml = renderFrameworkView(data);

document.getElementById('container').innerHTML = domainHtml;
```

### Roadmap Views

```typescript
import {
  renderTimelineView,
  renderStoryboardView,
  renderDependencyView,
  renderTeamView,
} from '@grokify/prism/html/roadmap';
import type { JourneyRoadmap } from '@grokify/prism/schema/roadmap';

const roadmap: JourneyRoadmap = { /* ... */ };

// Timeline/heatmap of capability maturity over time
const timelineHtml = renderTimelineView(roadmap, {
  showLegend: true,
  colorScheme: 'heatmap',
});

// Period-based narrative cards
const storyboardHtml = renderStoryboardView(roadmap, {
  orientation: 'horizontal',
});

// Dependency table with critical path
const dependencyHtml = renderDependencyView(roadmap);

// Team hierarchy with capacity
const teamHtml = renderTeamView(roadmap);
```

### Including Styles

The HTML renderers require CSS styles. Include the stylesheet:

```html
<link rel="stylesheet" href="node_modules/@grokify/prism/dist/styles/prism-roadmap.css">
```

Or import in your bundler:

```typescript
import '@grokify/prism/styles/prism-roadmap.css';
```

For SSR or minimal styling:

```typescript
import { INLINE_STYLES } from '@grokify/prism/styles';

// Inject minimal inline styles
document.head.insertAdjacentHTML('beforeend', `<style>${INLINE_STYLES}</style>`);
```

## Lit Web Components

Interactive web components built with [Lit](https://lit.dev). Requires Lit as a peer dependency.

### maturity-grid

```bash
# Install Lit peer dependency
npm install lit
```

```typescript
// Register the component
import '@grokify/prism/components';
```

#### Inline Data

```html
<maturity-grid theme="dark" view="by-layer" show-legend show-view-toggle>
  <script type="application/json">
    {
      "title": "Security Capabilities",
      "layers": [
        { "id": "prevention", "name": "Prevention", "order": 1 }
      ],
      "categories": [
        { "id": "appsec", "name": "Application Security" }
      ],
      "capabilities": [
        {
          "id": "sast",
          "name": "SAST",
          "layerId": "prevention",
          "categoryId": "appsec",
          "status": "operational"
        }
      ]
    }
  </script>
</maturity-grid>
```

#### Remote Data

```html
<maturity-grid
  src="/api/capabilities.json"
  theme="light"
  view="by-category"
></maturity-grid>
```

#### Attributes

| Attribute | Type | Default | Description |
|-----------|------|---------|-------------|
| `theme` | `'light'` \| `'dark'` | `'light'` | Color theme |
| `view` | `'by-layer'` \| `'by-category'` | `'by-layer'` | Grouping mode |
| `src` | `string` | - | URL to fetch JSON data |
| `show-legend` | `boolean` | `true` | Show filter legend |
| `show-view-toggle` | `boolean` | `true` | Show view toggle buttons |

#### Events

| Event | Detail | Description |
|-------|--------|-------------|
| `view-change` | `{ view: ViewMode }` | Fired when view mode changes |

#### Programmatic Usage

```typescript
import { MaturityGrid } from '@grokify/prism/components';

const grid = document.querySelector('maturity-grid') as MaturityGrid;

// Change view programmatically
grid.view = 'by-category';

// Listen for events
grid.addEventListener('view-change', (e) => {
  console.log('View changed to:', e.detail.view);
});
```

## Integration Examples

### React

```tsx
import { useEffect, useRef } from 'react';
import { renderTimelineView } from '@grokify/prism/html/roadmap';
import type { JourneyRoadmap } from '@grokify/prism/schema/roadmap';

function TimelineView({ roadmap }: { roadmap: JourneyRoadmap }) {
  const containerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (containerRef.current) {
      containerRef.current.innerHTML = renderTimelineView(roadmap);
    }
  }, [roadmap]);

  return <div ref={containerRef} />;
}
```

### Vue

```vue
<template>
  <div v-html="timelineHtml"></div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { renderTimelineView } from '@grokify/prism/html/roadmap';
import type { JourneyRoadmap } from '@grokify/prism/schema/roadmap';

const props = defineProps<{ roadmap: JourneyRoadmap }>();

const timelineHtml = computed(() => renderTimelineView(props.roadmap));
</script>
```

### Vanilla JavaScript

```html
<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" href="node_modules/@grokify/prism/dist/styles/prism-roadmap.css">
</head>
<body>
  <div id="timeline"></div>

  <script type="module">
    import { renderTimelineView } from '@grokify/prism/html/roadmap';
    import { JourneyRoadmapSchema } from '@grokify/prism/schema/roadmap';

    const response = await fetch('/api/roadmap.json');
    const data = await response.json();

    const roadmap = JourneyRoadmapSchema.parse(data);
    document.getElementById('timeline').innerHTML = renderTimelineView(roadmap);
  </script>
</body>
</html>
```

## Data Format

### MaturityGridData

```json
{
  "title": "Platform Capabilities",
  "layers": [
    { "id": "infra", "name": "Infrastructure", "order": 1 },
    { "id": "platform", "name": "Platform", "order": 2 }
  ],
  "categories": [
    { "id": "compute", "name": "Compute" },
    { "id": "storage", "name": "Storage" }
  ],
  "capabilities": [
    {
      "id": "k8s",
      "name": "Kubernetes",
      "fullName": "Kubernetes Container Orchestration",
      "description": "Container orchestration platform",
      "layerId": "infra",
      "categoryId": "compute",
      "status": "operational",
      "owner": "Platform Team"
    }
  ],
  "maturity": {
    "k8s": { "level": 4 }
  }
}
```

### JourneyRoadmap

See the [prism-roadmap documentation](https://github.com/grokify/prism-roadmap) for the full schema.

## Related

- [PRISM CLI](../cli/index.md) - Go command-line tool
- [prism-roadmap](https://github.com/grokify/prism-roadmap) - Go roadmap types
- [prism-maturity](https://github.com/grokify/prism-maturity) - Go maturity models
