/**
 * HTML Renderers
 *
 * Framework-agnostic HTML string rendering for PRISM visualizations.
 * These components return HTML strings that can be used in any context
 * (React dangerouslySetInnerHTML, server-side rendering, vanilla JS, etc.)
 *
 * @example
 * ```typescript
 * import { renderTimelineView, renderDomainView } from '@grokify/prism/html';
 *
 * const roadmapHtml = renderTimelineView(roadmap);
 * const maturityHtml = renderDomainView(spec);
 * ```
 */

// Maturity model renderers
export * from './maturity/index.js';

// Journey roadmap renderers
export * from './roadmap/index.js';
