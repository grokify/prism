/**
 * Schema exports
 *
 * Zod schemas and TypeScript types for PRISM models.
 * These have no UI dependencies - just Zod for validation.
 *
 * @example
 * ```typescript
 * import { JourneyRoadmapSchema, SpecSchema } from '@grokify/prism/schema';
 * ```
 */

// Maturity model schemas
export * from './maturity/index.js';

// Journey roadmap schemas
export * from './roadmap/index.js';
