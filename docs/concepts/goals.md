# Goals

Goals represent strategic objectives that an organization wants to achieve. Each goal has its own maturity model, allowing you to track progress from reactive practices to optimized, industry-leading capabilities.

## Goal Structure

A goal contains:

- **Identity**: ID, name, description
- **Ownership**: Owner, priority
- **Status**: active, on_hold, completed, cancelled
- **Timeline**: Start date, target date
- **Maturity**: Current level, target level, and a complete maturity model

## Goal Maturity Model

Each goal has its own 5-level maturity model that defines what it means to progress from Level 1 (Reactive) to Level 5 (Optimizing) for that specific goal.

### Maturity Levels

| Level | Name | Description |
|-------|------|-------------|
| 1 | Reactive | Ad-hoc practices, no systematic approach |
| 2 | Basic | Basic practices in place, manual processes |
| 3 | Defined | Standardized processes, documented practices |
| 4 | Managed | Metrics-driven, SLOs actively monitored |
| 5 | Optimizing | Continuous improvement, industry-leading |

### SLO Requirements

Each maturity level specifies which SLOs must be met to achieve that level:

```json
{
  "level": 3,
  "name": "Defined",
  "description": "SLOs tracked, incident response documented",
  "requiredSLOs": [
    { "metricId": "ops-availability", "description": "Production SLO" },
    { "metricId": "ops-mttr", "description": "Recovery time tracked" }
  ],
  "metricCriteria": [
    { "metricId": "ops-availability", "operator": "gte", "value": 99.5 },
    { "metricId": "ops-mttr", "operator": "lte", "value": 4 }
  ]
}
```

### Level Achievement

A goal achieves maturity level N when:

1. **All required SLOs are met** - Each metric referenced in `requiredSLOs` must have `MeetsSLO() == true`
2. **All metric criteria are satisfied** - Each criterion's value requirement must be met

The system checks levels from 5 down to 1 and returns the highest level where all requirements are satisfied.

## Example Goal

```json
{
  "id": "goal-reliability",
  "name": "Achieve High Reliability",
  "description": "Build world-class reliability with elite SLO compliance",
  "owner": "VP Engineering",
  "priority": 1,
  "status": "active",
  "targetDate": "2026-12-31",
  "currentLevel": 3,
  "targetLevel": 5,
  "maturityModel": {
    "levels": [
      {
        "level": 1,
        "name": "Reactive",
        "description": "Frequent outages, no SLOs",
        "requiredSLOs": []
      },
      {
        "level": 2,
        "name": "Basic",
        "description": "Basic monitoring in place",
        "requiredSLOs": [
          { "metricId": "ops-availability" }
        ],
        "metricCriteria": [
          { "metricId": "ops-availability", "operator": "gte", "value": 99.0 }
        ]
      },
      {
        "level": 3,
        "name": "Defined",
        "description": "SLOs tracked, incident response documented",
        "requiredSLOs": [
          { "metricId": "ops-availability" },
          { "metricId": "ops-mttr" }
        ],
        "metricCriteria": [
          { "metricId": "ops-availability", "operator": "gte", "value": 99.5 },
          { "metricId": "ops-mttr", "operator": "lte", "value": 4 }
        ]
      },
      {
        "level": 4,
        "name": "Managed",
        "description": "Proactive monitoring, error budgets",
        "requiredSLOs": [
          { "metricId": "ops-availability" },
          { "metricId": "ops-mttr" },
          { "metricId": "ops-p99-latency" }
        ],
        "metricCriteria": [
          { "metricId": "ops-availability", "operator": "gte", "value": 99.9 },
          { "metricId": "ops-mttr", "operator": "lte", "value": 1 },
          { "metricId": "ops-p99-latency", "operator": "lte", "value": 200 }
        ]
      },
      {
        "level": 5,
        "name": "Optimizing",
        "description": "Self-healing systems, chaos engineering",
        "requiredSLOs": [
          { "metricId": "ops-availability" },
          { "metricId": "ops-mttr" },
          { "metricId": "ops-p99-latency" }
        ],
        "metricCriteria": [
          { "metricId": "ops-availability", "operator": "gte", "value": 99.99 },
          { "metricId": "ops-mttr", "operator": "lte", "value": 0.25 },
          { "metricId": "ops-p99-latency", "operator": "lte", "value": 100 }
        ]
      }
    ]
  }
}
```

## Linking Initiatives to Goals

Initiatives can be linked to one or more goals via the `goalIds` field:

```json
{
  "id": "init-monitoring",
  "name": "Observability Platform",
  "goalIds": ["goal-reliability"],
  "phaseId": "phase-q1-2026"
}
```

## Goal Progress Tracking

### Calculate Achieved Level

```go
level := goal.CalculateAchievedLevel(doc.Metrics)
fmt.Printf("Achieved: Level %d (%s)\n", level.Level, level.Name)
```

### Check Requirements for Next Level

```go
requirements := goal.RequirementsForLevel(nextLevel, doc.Metrics)
for _, req := range requirements {
    status := "Met"
    if !req.Met {
        status = "Not Met"
    }
    fmt.Printf("%s: %s\n", req.MetricName, status)
}
```

## CLI Commands

```bash
# List all goals
prism goal list document.json

# Show goal progress
prism goal progress goal-reliability document.json

# Show all goals status
prism goal status document.json
```

## Domain Extensions

For security-specific goal examples (security posture, compliance readiness, etc.), see [prism-security](https://github.com/grokify/prism-security).
