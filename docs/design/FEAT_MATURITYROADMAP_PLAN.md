# FEAT_MATURITYROADMAP: Implementation Plan

## Overview

This plan implements the Goal-driven Maturity Roadmap feature as specified in FEAT_MATURITYROADMAP_PRD.md.

## Implementation Phases

### Phase 1: Core Goal Types

**Files to create/modify:**

- `goal.go` - Goal, GoalMaturityModel, GoalMaturityLevel, SLORequirement, MetricCriterion types

**Tasks:**

1. Create Goal struct with all fields from PRD R1
2. Create GoalMaturityModel struct (R2)
3. Create GoalMaturityLevel struct with RequiredSLOs and MetricCriteria (R3)
4. Create SLORequirement struct
5. Create MetricCriterion struct with IsMet() method
6. Implement Goal.CurrentMaturityLevel() calculation
7. Implement Goal.MeetsLevelRequirements() method
8. Add validation for Goal types

### Phase 2: Phase and Swimlane Types

**Files to create/modify:**

- `phase.go` - Phase, PhaseGoalTarget, Swimlane types

**Tasks:**

1. Create Phase struct with all fields from PRD R4
2. Create PhaseGoalTarget struct
3. Create Swimlane struct (R5)
4. Add validation for Phase types

### Phase 3: Enhanced Initiatives

**Files to modify:**

- `prism.go` - Add GoalIds, PhaseId, DeploymentStatus to Initiative

**Tasks:**

1. Add GoalIds []string field to Initiative
2. Add PhaseId string field to Initiative
3. Add DevCompletionPercent float64 field
4. Create DeploymentStatus struct (R6)
5. Update Initiative validation

### Phase 4: Phase Metrics

**Files to create:**

- `phase_metrics.go` - PhaseMetrics, GoalProgress, InitiativeMetrics, SLOCompliance types

**Tasks:**

1. Create PhaseMetrics struct (R7)
2. Create GoalProgress struct with all calculation fields
3. Create InitiativeMetrics struct
4. Create SLOCompliance struct
5. Implement Phase.GoalProgress() calculation method
6. Implement Phase.InitiativeMetrics() calculation

### Phase 5: PRISMDocument Integration

**Files to modify:**

- `prism.go` - Add Goals, Phases, Roadmap fields to PRISMDocument

**Tasks:**

1. Add Goals []Goal field
2. Add Phases []Phase field
3. Add Roadmap *RoadmapConfig field (optional config)
4. Update PRISMDocument validation to include new types
5. Ensure cross-references are validated (metric IDs in SLORequirements exist)

### Phase 6: CLI Commands

**Files to create:**

- `cmd/prism/goal.go` - goal list, show, progress commands
- `cmd/prism/phase.go` - phase list, show, metrics commands
- `cmd/prism/roadmap.go` - roadmap show, progress commands

**Tasks:**

1. Implement `prism goal list` command
2. Implement `prism goal show <goal-id>` command
3. Implement `prism goal progress <goal-id>` command
4. Implement `prism phase list` command
5. Implement `prism phase show <phase-id>` command
6. Implement `prism phase metrics <phase-id>` command
7. Implement `prism roadmap show` command
8. Implement `prism roadmap progress` command
9. Enhance `prism score` with --goals and --phase flags

### Phase 7: Examples and Tests

**Files to create:**

- `examples/goal-roadmap.json` - Complete example with goals, phases, initiatives
- `goal_test.go` - Unit tests for goal types and calculations
- `phase_test.go` - Unit tests for phase types
- `phase_metrics_test.go` - Unit tests for metrics calculations

**Tasks:**

1. Create comprehensive example JSON
2. Write unit tests for Goal.CurrentMaturityLevel()
3. Write unit tests for Goal.MeetsLevelRequirements()
4. Write unit tests for MetricCriterion.IsMet()
5. Write unit tests for Phase.GoalProgress()
6. Write integration tests with example file

### Phase 8: Schema and Documentation

**Files to modify:**

- `schema/generate.go` - Include new types in schema generation
- `docs/concepts/goals.md` - New documentation page
- `docs/concepts/phases.md` - New documentation page
- `docs/cli/goal.md` - New CLI reference
- `docs/cli/phase.md` - New CLI reference
- `docs/cli/roadmap.md` - New CLI reference

**Tasks:**

1. Regenerate JSON schema with new types
2. Write goals concept documentation
3. Write phases concept documentation
4. Write CLI command documentation
5. Update mkdocs.yml navigation

## File Summary

### New Files

| File | Description |
|------|-------------|
| `goal.go` | Goal, GoalMaturityModel, GoalMaturityLevel types |
| `phase.go` | Phase, PhaseGoalTarget, Swimlane types |
| `phase_metrics.go` | PhaseMetrics, GoalProgress, InitiativeMetrics types |
| `cmd/prism/goal.go` | Goal CLI commands |
| `cmd/prism/phase.go` | Phase CLI commands |
| `cmd/prism/roadmap.go` | Roadmap CLI commands |
| `goal_test.go` | Goal unit tests |
| `phase_test.go` | Phase unit tests |
| `phase_metrics_test.go` | Phase metrics tests |
| `examples/goal-roadmap.json` | Complete example |
| `docs/concepts/goals.md` | Goals documentation |
| `docs/concepts/phases.md` | Phases documentation |
| `docs/cli/goal.md` | Goal CLI docs |
| `docs/cli/phase.md` | Phase CLI docs |
| `docs/cli/roadmap.md` | Roadmap CLI docs |

### Modified Files

| File | Changes |
|------|---------|
| `prism.go` | Add Goals, Phases, Roadmap to PRISMDocument; enhance Initiative |
| `validation.go` | Add validation for new types |
| `schema/generate.go` | Include new types |
| `mkdocs.yml` | Add new navigation entries |

## Dependency Order

```
Phase 1 (goal.go)
    ↓
Phase 2 (phase.go)
    ↓
Phase 3 (initiative enhancements) ← depends on Phase 1 for GoalIds
    ↓
Phase 4 (phase_metrics.go) ← depends on all above
    ↓
Phase 5 (PRISMDocument integration)
    ↓
Phase 6 (CLI) ← can start after Phase 5
Phase 7 (Tests) ← can start after Phase 4
Phase 8 (Docs) ← can start after Phase 6
```

## Verification Checklist

- [ ] All new types have JSON tags with omitempty where appropriate
- [ ] All validation functions return meaningful error messages
- [ ] MeetsSLO() integration works with existing Metric type
- [ ] CurrentMaturityLevel() correctly iterates from 5 to 1
- [ ] Phase metrics calculations match PRD formulas
- [ ] CLI commands have consistent help text
- [ ] Example JSON validates against schema
- [ ] All tests pass
- [ ] golangci-lint passes
