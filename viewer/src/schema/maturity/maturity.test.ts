import { describe, it, expect } from 'vitest';
import {
  SpecSchema,
  SLISchema,
  CriterionSchema,
  LevelSchema,
  DomainModelSchema,
  validateSpec,
  safeValidateSpec,
} from './index.js';

describe('Maturity Schema', () => {
  describe('SLI Schema', () => {
    it('validates a minimal SLI', () => {
      const sli = {
        id: 'sli-1',
        name: 'Deployment Frequency',
        metricName: 'deployments_per_day',
      };
      expect(SLISchema.parse(sli)).toEqual(sli);
    });

    it('validates a full SLI with framework mappings', () => {
      const sli = {
        id: 'sli-1',
        name: 'Deployment Frequency',
        description: 'How often deployments occur',
        metricName: 'deployments_per_day',
        unit: 'per day',
        type: 'quantitative' as const,
        layer: 'operations',
        category: 'devops',
        frameworkMappings: [
          {
            framework: 'DORA',
            reference: 'DF-1',
            name: 'Deployment Frequency',
            description: 'DORA deployment frequency metric',
          },
        ],
      };
      const result = SLISchema.parse(sli);
      expect(result.id).toBe('sli-1');
      expect(result.frameworkMappings).toHaveLength(1);
    });

    it('rejects SLI without required fields', () => {
      const invalid = { id: 'sli-1' };
      expect(() => SLISchema.parse(invalid)).toThrow();
    });
  });

  describe('Criterion Schema', () => {
    it('validates a criterion with SLO target', () => {
      const criterion = {
        id: 'crit-1',
        name: 'Deploy daily',
        operator: 'gte' as const,
        target: 1,
      };
      expect(CriterionSchema.parse(criterion)).toEqual(criterion);
    });

    it('validates all operator types', () => {
      const operators = ['gte', 'lte', 'gt', 'lt', 'eq', 'exists'] as const;
      for (const op of operators) {
        const criterion = {
          id: 'crit-1',
          name: 'Test',
          operator: op,
          target: 1,
        };
        expect(() => CriterionSchema.parse(criterion)).not.toThrow();
      }
    });
  });

  describe('Level Schema', () => {
    it('validates a maturity level with criteria', () => {
      const level = {
        level: 1,
        name: 'Developing',
        description: 'Initial processes defined',
        criteria: [
          {
            id: 'crit-1',
            name: 'Basic deployment',
            operator: 'gte' as const,
            target: 1,
          },
        ],
      };
      const result = LevelSchema.parse(level);
      expect(result.level).toBe(1);
      expect(result.criteria).toHaveLength(1);
    });

    it('validates a level with enablers', () => {
      const level = {
        level: 2,
        name: 'Defined',
        description: 'Standardized processes',
        enablers: [
          {
            id: 'en-1',
            name: 'CI Pipeline',
            description: 'Set up continuous integration',
            type: 'tooling',
            effort: 'medium',
          },
        ],
      };
      const result = LevelSchema.parse(level);
      expect(result.enablers).toHaveLength(1);
    });
  });

  describe('DomainModel Schema', () => {
    it('validates a domain model with levels', () => {
      const domain = {
        name: 'Operations',
        description: 'Operational maturity',
        owner: 'ops-team',
        levels: [
          {
            level: 1,
            name: 'Basic',
            description: 'Basic operations',
          },
          {
            level: 2,
            name: 'Defined',
            description: 'Standardized operations',
          },
        ],
      };
      const result = DomainModelSchema.parse(domain);
      expect(result.name).toBe('Operations');
      expect(result.levels).toHaveLength(2);
    });
  });

  describe('Spec Schema', () => {
    it('validates a minimal spec', () => {
      const spec = {
        domains: {
          operations: {
            name: 'Operations',
            levels: [
              {
                level: 1,
                name: 'Basic',
                description: 'Basic operations',
              },
            ],
          },
        },
      };
      const result = SpecSchema.parse(spec);
      expect(result.domains.operations).toBeDefined();
    });

    it('validates a full spec with metadata and SLIs', () => {
      const spec = {
        $schema: 'https://example.com/prism/maturity.schema.json',
        metadata: {
          name: 'Platform Maturity Model',
          version: '1.0.0',
          organization: 'ACME Corp',
        },
        slis: {
          'deploy-freq': {
            id: 'deploy-freq',
            name: 'Deployment Frequency',
            metricName: 'deployments_per_day',
            type: 'quantitative' as const,
          },
        },
        domains: {
          devops: {
            name: 'DevOps',
            levels: [
              {
                level: 1,
                name: 'Initial',
                description: 'Ad-hoc processes',
              },
            ],
          },
        },
        assessments: {
          devops: {
            domain: 'devops',
            currentLevel: 2,
            targetLevel: 4,
          },
        },
      };
      const result = SpecSchema.parse(spec);
      expect(result.metadata?.name).toBe('Platform Maturity Model');
      expect(result.slis?.['deploy-freq']).toBeDefined();
      expect(result.assessments?.devops.currentLevel).toBe(2);
    });
  });

  describe('validateSpec helper', () => {
    it('returns parsed spec on valid data', () => {
      const spec = {
        domains: {
          test: {
            name: 'Test',
            levels: [{ level: 1, name: 'L1', description: 'Level 1' }],
          },
        },
      };
      const result = validateSpec(spec);
      expect(result.domains.test.name).toBe('Test');
    });

    it('throws on invalid data', () => {
      expect(() => validateSpec({})).toThrow();
    });
  });

  describe('safeValidateSpec helper', () => {
    it('returns success result on valid data', () => {
      const spec = {
        domains: {
          test: {
            name: 'Test',
            levels: [{ level: 1, name: 'L1', description: 'Level 1' }],
          },
        },
      };
      const result = safeValidateSpec(spec);
      expect(result.success).toBe(true);
      if (result.success) {
        expect(result.data.domains.test.name).toBe('Test');
      }
    });

    it('returns error result on invalid data', () => {
      const result = safeValidateSpec({});
      expect(result.success).toBe(false);
      if (!result.success) {
        expect(result.error.issues.length).toBeGreaterThan(0);
      }
    });
  });
});
