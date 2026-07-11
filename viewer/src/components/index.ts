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
 * ```
 */

export { MaturityGrid } from './maturity-grid.js';

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
