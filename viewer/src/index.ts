/**
 * @grokify/prism
 *
 * PRISM visualization library for capability and maturity models.
 *
 * This package provides:
 * - Zod schemas and TypeScript types (source of truth)
 * - HTML string renderers (framework-agnostic)
 * - Lit web components
 *
 * @example
 * ```typescript
 * // Import schemas/types only (no UI deps)
 * import { JourneyRoadmapSchema, type JourneyRoadmap } from '@grokify/prism/schema';
 *
 * // Import HTML renderers (minimal deps)
 * import { renderTimelineView } from '@grokify/prism/html';
 *
 * // Import Lit components (requires Lit)
 * import { MaturityGrid } from '@grokify/prism/components';
 * ```
 */

// Re-export everything for convenience
export * from './schema/index.js';
export * from './html/index.js';
export * from './components/index.js';
export * from './styles/index.js';
