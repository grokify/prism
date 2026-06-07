/**
 * PRISM UI - Lit web components for capability and maturity visualization.
 *
 * @example
 * ```html
 * <script type="module" src="prism-ui.js"></script>
 *
 * <maturity-grid view="by-layer" theme="dark">
 *   <script type="application/json">
 *     {"layers": [...], "capabilities": [...]}
 *   </script>
 * </maturity-grid>
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

export {
  STATUS_COLORS,
  STATUS_TEXT_COLORS,
  MATURITY_COLORS,
} from './types.js';
