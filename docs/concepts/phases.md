# Phases

Phases organize work into time-bounded periods, typically quarters, allowing teams to plan incremental progress toward maturity goals.

## Phase Structure

A phase contains:

- **Identity**: ID, name, description
- **Time bounds**: Quarter, year, start date, end date
- **Status**: planning, in_progress, completed
- **Goal targets**: Enter and exit maturity levels for each goal
- **Swimlanes**: Grouped initiatives by domain

## Phase Statuses

| Status | Description |
|--------|-------------|
| `planning` | Phase is being planned, not yet started |
| `in_progress` | Phase is currently active |
| `completed` | Phase has ended |

## Goal Targets

Each phase specifies target maturity levels for goals:

```json
{
  "goalTargets": [
    {
      "goalId": "goal-reliability",
      "enterLevel": 2,
      "exitLevel": 3
    },
    {
      "goalId": "goal-velocity",
      "enterLevel": 2,
      "exitLevel": 3
    }
  ]
}
```

- **enterLevel**: Expected maturity at phase start
- **exitLevel**: Target maturity at phase end

## Swimlanes

Swimlanes group initiatives within a phase by domain:

```json
{
  "swimlanes": [
    {
      "id": "sw-platform",
      "name": "Platform Initiatives",
      "domain": "operations",
      "initiativeIds": ["init-monitoring", "init-ci-cd"]
    }
  ]
}
```

## Example Phase

```json
{
  "id": "phase-q1-2026",
  "name": "Q1 2026",
  "quarter": "Q1",
  "year": 2026,
  "startDate": "2026-01-01",
  "endDate": "2026-03-31",
  "status": "in_progress",
  "goalTargets": [
    {
      "goalId": "goal-reliability",
      "enterLevel": 2,
      "exitLevel": 3
    },
    {
      "goalId": "goal-velocity",
      "enterLevel": 2,
      "exitLevel": 3
    }
  ],
  "swimlanes": [
    {
      "id": "sw-platform",
      "name": "Platform Initiatives",
      "domain": "operations",
      "initiativeIds": ["init-monitoring", "init-ci-cd"]
    }
  ]
}
```

## Phase Metrics

Phase metrics show the performance of metrics tied to phase initiatives:

```json
{
  "metricId": "ops-availability",
  "metricName": "Service Availability",
  "current": 99.95,
  "target": 99.99,
  "unit": "%",
  "meetsSLO": true
}
```

## Initiatives

Initiatives are projects linked to phases and goals:

```json
{
  "id": "init-monitoring",
  "name": "Observability Platform",
  "description": "Deploy comprehensive monitoring and alerting",
  "status": "in_progress",
  "priority": 1,
  "owner": "SRE Team",
  "goalIds": ["goal-reliability"],
  "phaseId": "phase-q1-2026",
  "metricIds": ["ops-availability", "ops-mttr"],
  "devCompletionPercent": 90,
  "deploymentStatus": {
    "status": "in_progress",
    "totalCustomers": 50,
    "deployedCustomers": 45,
    "adoptionPercent": 90
  }
}
```

### Initiative Statuses

| Status | Description |
|--------|-------------|
| `planned` | Not yet started |
| `in_progress` | Currently being worked on |
| `completed` | Fully deployed |
| `blocked` | Blocked by dependencies |
| `cancelled` | No longer planned |

## CLI Commands

```bash
# List all phases
prism phase list document.json

# Show phase details
prism phase show phase-q1-2026 document.json

# Show phase metrics
prism phase metrics phase-q1-2026 document.json

# Show phase progress
prism phase progress phase-q1-2026 document.json
```

### Example Output

```
Phase: Q1 2026 (in_progress)
Period: 2026-01-01 to 2026-03-31

Goal Targets:
  goal-reliability: Level 2 → Level 3
  goal-velocity: Level 2 → Level 3

Platform Initiatives:
  - init-monitoring: Observability Platform [in_progress] (90% dev, 90% deployed)
  - init-ci-cd: CI/CD Pipeline Enhancement [completed] (100% deployed)

Metrics:
  ✓ ops-availability: 99.95% (target: >=99.9%)
  ✓ ops-deploy-frequency: 5/day (target: >=5/day)
```

## Roadmap View

The roadmap combines all phases to show the journey:

```bash
prism roadmap progress document.json
```

## Domain Extensions

For security-specific phase examples, see [prism-security](https://github.com/grokify/prism-security).
