import { describe, it, expect } from 'vitest';
import {
  PeriodSchema,
  TimeModelSchema,
  CapabilityJourneySchema,
  InitiativeSchema,
  DependencySchema,
  TeamSchema,
  JourneyRoadmapSchema,
  validateJourneyRoadmap,
  safeValidateJourneyRoadmap,
  MaturityLevels,
} from './index.js';

describe('Roadmap Schema', () => {
  describe('Period Schema', () => {
    it('validates a minimal period', () => {
      const period = {
        id: 'q1-2026',
        label: 'Q1 2026',
      };
      expect(PeriodSchema.parse(period)).toEqual(period);
    });

    it('validates a full period', () => {
      const period = {
        id: 'q1-2026',
        label: 'Q1 2026',
        startDate: '2026-01-01',
        endDate: '2026-03-31',
        isCurrent: true,
        description: 'First quarter',
      };
      const result = PeriodSchema.parse(period);
      expect(result.isCurrent).toBe(true);
    });
  });

  describe('TimeModel Schema', () => {
    it('validates a quarterly time model', () => {
      const timeModel = {
        type: 'quarterly' as const,
        periods: [
          { id: 'q1', label: 'Q1' },
          { id: 'q2', label: 'Q2' },
        ],
      };
      const result = TimeModelSchema.parse(timeModel);
      expect(result.type).toBe('quarterly');
      expect(result.periods).toHaveLength(2);
    });

    it('validates all time model types', () => {
      const types = ['quarterly', 'monthly', 'sprint', 'milestone', 'half', 'custom'] as const;
      for (const t of types) {
        const model = {
          type: t,
          periods: [{ id: 'p1', label: 'Period 1' }],
        };
        expect(() => TimeModelSchema.parse(model)).not.toThrow();
      }
    });
  });

  describe('CapabilityJourney Schema', () => {
    it('validates a minimal capability journey', () => {
      const journey = {
        id: 'cj-1',
        capabilityId: 'cap-1',
        name: 'CI/CD Pipeline',
        targetStates: [],
      };
      expect(CapabilityJourneySchema.parse(journey)).toEqual(journey);
    });

    it('validates a journey with current and target states', () => {
      const journey = {
        id: 'cj-1',
        capabilityId: 'cap-1',
        name: 'CI/CD Pipeline',
        description: 'Continuous integration and deployment',
        owner: 'platform-team',
        currentState: {
          periodId: 'q1-2026',
          maturityLevel: 'M2',
          summary: 'Basic CI in place',
        },
        targetStates: [
          {
            periodId: 'q2-2026',
            maturityLevel: 'M3',
            confidence: 0.8,
            commitment: 'committed' as const,
          },
          {
            periodId: 'q4-2026',
            maturityLevel: 'M4',
            confidence: 0.5,
            commitment: 'planned' as const,
          },
        ],
        tags: ['devops', 'automation'],
      };
      const result = CapabilityJourneySchema.parse(journey);
      expect(result.targetStates).toHaveLength(2);
      expect(result.currentState?.maturityLevel).toBe('M2');
    });

    it('validates commitment levels', () => {
      const commitments = ['committed', 'planned', 'targeted', 'aspirant'] as const;
      for (const c of commitments) {
        const journey = {
          id: 'cj-1',
          capabilityId: 'cap-1',
          name: 'Test',
          targetStates: [
            {
              periodId: 'q1',
              maturityLevel: 'M3',
              commitment: c,
            },
          ],
        };
        expect(() => CapabilityJourneySchema.parse(journey)).not.toThrow();
      }
    });
  });

  describe('Initiative Schema', () => {
    it('validates a minimal initiative', () => {
      const initiative = {
        id: 'init-1',
        name: 'Deploy Automation',
      };
      expect(InitiativeSchema.parse(initiative)).toEqual(initiative);
    });

    it('validates a full initiative', () => {
      const initiative = {
        id: 'init-1',
        name: 'Deploy Automation',
        description: 'Automate deployment pipelines',
        status: 'in_progress' as const,
        ownerTeam: 'platform',
        contributingTeams: ['devops', 'sre'],
        periods: ['q1-2026', 'q2-2026'],
        advances: [
          {
            capabilityId: 'cap-1',
            from: 'M2',
            to: 'M3',
          },
        ],
        expectedOutcomes: ['Faster deployments', 'Fewer failures'],
        requiredCapacity: {
          storyPoints: 40,
          ftes: 2,
        },
        dependencies: ['init-0'],
        risks: ['Resource constraints'],
        links: [
          {
            url: 'https://example.com/project',
            title: 'Project page',
            type: 'documentation',
          },
        ],
        tags: ['automation', 'priority'],
      };
      const result = InitiativeSchema.parse(initiative);
      expect(result.status).toBe('in_progress');
      expect(result.advances).toHaveLength(1);
    });

    it('validates all status values', () => {
      const statuses = [
        'proposed',
        'planned',
        'in_progress',
        'completed',
        'on_hold',
        'cancelled',
      ] as const;
      for (const s of statuses) {
        const initiative = {
          id: 'init-1',
          name: 'Test',
          status: s,
        };
        expect(() => InitiativeSchema.parse(initiative)).not.toThrow();
      }
    });
  });

  describe('Dependency Schema', () => {
    it('validates a dependency between entities', () => {
      const dependency = {
        id: 'dep-1',
        from: { type: 'initiative' as const, id: 'init-1' },
        to: { type: 'capability' as const, id: 'cap-1' },
        type: 'requires' as const,
      };
      const result = DependencySchema.parse(dependency);
      expect(result.from.type).toBe('initiative');
      expect(result.to.type).toBe('capability');
    });

    it('validates all entity types', () => {
      const types = [
        'capability',
        'initiative',
        'team',
        'milestone',
        'outcome',
        'external',
        'decision',
      ] as const;
      for (const t of types) {
        const dependency = {
          from: { type: t, id: 'from-1' },
          to: { type: t, id: 'to-1' },
          type: 'requires' as const,
        };
        expect(() => DependencySchema.parse(dependency)).not.toThrow();
      }
    });

    it('validates all dependency types', () => {
      const depTypes = [
        'requires',
        'blocked_by',
        'resource',
        'external',
        'informs',
        'contributes',
      ] as const;
      for (const dt of depTypes) {
        const dependency = {
          from: { type: 'initiative' as const, id: 'from-1' },
          to: { type: 'initiative' as const, id: 'to-1' },
          type: dt,
        };
        expect(() => DependencySchema.parse(dependency)).not.toThrow();
      }
    });
  });

  describe('Team Schema', () => {
    it('validates a minimal team', () => {
      const team = {
        id: 'team-1',
        name: 'Platform Team',
      };
      expect(TeamSchema.parse(team)).toEqual(team);
    });

    it('validates a full team', () => {
      const team = {
        id: 'team-1',
        name: 'Platform Team',
        description: 'Core platform engineering',
        type: 'platform' as const,
        level: 'team' as const,
        leaderId: 'user-1',
        leaderName: 'Jane Doe',
        capacity: {
          ftes: 5,
          storyPointsPerSprint: 40,
          allocatedPercent: 80,
        },
        skills: ['kubernetes', 'terraform', 'go'],
        tags: ['core', 'infrastructure'],
      };
      const result = TeamSchema.parse(team);
      expect(result.type).toBe('platform');
      expect(result.capacity?.ftes).toBe(5);
    });
  });

  describe('JourneyRoadmap Schema', () => {
    it('validates a minimal roadmap', () => {
      const roadmap = {
        id: 'roadmap-1',
        name: 'Platform Roadmap',
      };
      expect(JourneyRoadmapSchema.parse(roadmap)).toEqual(roadmap);
    });

    it('validates a full roadmap', () => {
      const roadmap = {
        id: 'roadmap-1',
        type: 'capability',
        name: 'Platform Maturity Roadmap',
        vision: 'World-class platform engineering',
        description: 'Multi-quarter plan to improve platform capabilities',
        timeModel: {
          type: 'quarterly' as const,
          periods: [
            { id: 'q1-2026', label: 'Q1 2026', isCurrent: true },
            { id: 'q2-2026', label: 'Q2 2026' },
            { id: 'q3-2026', label: 'Q3 2026' },
          ],
        },
        capabilityJourneys: [
          {
            id: 'cj-1',
            capabilityId: 'cap-1',
            name: 'CI/CD',
            targetStates: [
              { periodId: 'q2-2026', maturityLevel: 'M3' },
            ],
          },
        ],
        initiatives: [
          {
            id: 'init-1',
            name: 'Pipeline Automation',
            status: 'planned' as const,
          },
        ],
        teams: [
          { id: 'team-1', name: 'Platform' },
        ],
        risks: [
          {
            id: 'risk-1',
            description: 'Resource constraints',
            probability: 'medium',
            impact: 'high',
          },
        ],
        narrative: {
          title: 'Journey to Excellence',
          destination: 'Leading platform capabilities',
        },
      };
      const result = JourneyRoadmapSchema.parse(roadmap);
      expect(result.timeModel?.periods).toHaveLength(3);
      expect(result.capabilityJourneys).toHaveLength(1);
      expect(result.initiatives).toHaveLength(1);
    });
  });

  describe('validateJourneyRoadmap helper', () => {
    it('returns parsed roadmap on valid data', () => {
      const roadmap = { id: 'r1', name: 'Test' };
      const result = validateJourneyRoadmap(roadmap);
      expect(result.name).toBe('Test');
    });

    it('throws on invalid data', () => {
      expect(() => validateJourneyRoadmap({})).toThrow();
    });
  });

  describe('safeValidateJourneyRoadmap helper', () => {
    it('returns success result on valid data', () => {
      const roadmap = { id: 'r1', name: 'Test' };
      const result = safeValidateJourneyRoadmap(roadmap);
      expect(result.success).toBe(true);
    });

    it('returns error result on invalid data', () => {
      const result = safeValidateJourneyRoadmap({});
      expect(result.success).toBe(false);
    });
  });

  describe('MaturityLevels constants', () => {
    it('has all maturity levels defined', () => {
      expect(MaturityLevels.M0).toBe('M0');
      expect(MaturityLevels.M1).toBe('M1');
      expect(MaturityLevels.M2).toBe('M2');
      expect(MaturityLevels.M3).toBe('M3');
      expect(MaturityLevels.M4).toBe('M4');
      expect(MaturityLevels.M5).toBe('M5');
    });
  });
});
