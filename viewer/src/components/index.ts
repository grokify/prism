/**
 * Lit Web Components
 *
 * Custom elements for PRISM visualizations.
 * Requires Lit as a peer dependency.
 *
 * @example
 * ```typescript
 * import '@grokify/prism/components';
 *
 * // Then use in HTML:
 * // <maturity-grid view="by-layer">...</maturity-grid>
 * // <roadmap-timeline theme="dark">...</roadmap-timeline>
 * // <initiative-tracker group-by="team">...</initiative-tracker>
 * ```
 */

export { MaturityGrid } from './maturity-grid.js';
export { RoadmapTimeline } from './roadmap-timeline.js';
export { InitiativeTracker } from './initiative-tracker.js';

export type {
  Capability,
  CapabilityStatus,
  Layer,
  Category,
  MaturityData,
  MaturityGridData,
  SLIThreshold,
  SLIBulletData,
  SLIGroup,
  MetricsTableRow,
  Tool,
} from './types.js';
