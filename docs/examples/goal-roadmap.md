# Goal Roadmap Example

This example demonstrates the Goal-driven Maturity Roadmap feature, showing how to:

- Define strategic goals with their own maturity models
- Organize work into phases (quarters)
- Link initiatives to goals and phases
- Track progress with SLO-backed maturity levels

## Full Example

The complete example is in `examples/goal-roadmap.json`.

## Goals

The example defines two strategic operations goals:

### 1. Achieve High Reliability

**Owner:** VP Engineering
**Priority:** 1 (highest)
**Current Level:** 3 (Defined)
**Target Level:** 5 (Optimizing)

This goal's maturity model defines what each level means for reliability:

| Level | Name | Key Requirements |
|-------|------|------------------|
| 1 | Reactive | Frequent outages, no SLOs |
| 2 | Basic | Basic monitoring, 99% availability |
| 3 | Defined | SLOs tracked, incident response documented |
| 4 | Managed | Proactive monitoring, error budgets, 99.9% availability |
| 5 | Optimizing | Self-healing systems, chaos engineering, 99.99% availability |

### 2. Accelerate Delivery Velocity

**Owner:** Platform Team Lead
**Priority:** 2
**Current Level:** 3 (Defined)
**Target Level:** 4 (Managed)

This goal's maturity model focuses on DORA metrics:

| Level | Name | Key Requirements |
|-------|------|------------------|
| 1 | Reactive | Manual deployments, weekly releases |
| 2 | Basic | Some automation, bi-weekly releases |
| 3 | Defined | CI/CD pipeline, daily deployments possible |
| 4 | Managed | Multiple daily deployments, low change failure rate |
| 5 | Optimizing | On-demand deployments, elite DORA metrics |

## Phases

The roadmap spans two quarters:

### Q1 2026 (In Progress)

**Goal Targets:**

- Reliability: Level 2 → Level 3
- Delivery Velocity: Level 2 → Level 3

**Swimlanes:**

- **Platform Initiatives:**
  - CI/CD Pipeline Enhancement (completed, 100% deployed)
  - Observability Platform (90% dev, 90% deployed)

### Q2 2026 (Planning)

**Goal Targets:**

- Reliability: Level 3 → Level 4
- Delivery Velocity: Level 3 → Level 4

**Planned Initiatives:**

- SLO Dashboard
- Deployment Automation

## Metrics

The example includes 6 operations metrics that support the goals:

| Metric | Category | Current | SLO Target |
|--------|----------|---------|------------|
| Service Availability | Reliability | 99.95% | ≥99.9% |
| Deployment Frequency | Efficiency | 5/day | ≥5/day |
| Lead Time for Changes | Efficiency | 24h | ≤24h |
| Mean Time to Recovery | Response | 2h | ≤1h |
| Change Failure Rate | Quality | 5% | ≤5% |
| P99 Latency | Efficiency | 150ms | ≤200ms |

## Key Concepts Demonstrated

### 1. SLO-Backed Maturity Levels

Each maturity level specifies required SLOs:

```json
{
  "level": 3,
  "name": "Defined",
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

### 2. Initiative-Goal Linkage

Initiatives reference goals they contribute to:

```json
{
  "id": "init-monitoring",
  "name": "Observability Platform",
  "goalIds": ["goal-reliability"],
  "phaseId": "phase-q1-2026"
}
```

### 3. Deployment Tracking

Track customer adoption for completed initiatives:

```json
{
  "deploymentStatus": {
    "status": "completed",
    "totalCustomers": 50,
    "deployedCustomers": 50,
    "adoptionPercent": 100
  }
}
```

### 4. Phase Goal Targets

Define enter/exit maturity levels per phase:

```json
{
  "goalTargets": [
    {
      "goalId": "goal-reliability",
      "enterLevel": 2,
      "exitLevel": 3
    }
  ]
}
```

## Using the Example

```bash
# Validate the document
prism validate examples/goal-roadmap.json

# Show all goals
prism goal list examples/goal-roadmap.json

# Check reliability goal progress
prism goal progress goal-reliability examples/goal-roadmap.json

# Show Q1 phase metrics
prism phase metrics phase-q1-2026 examples/goal-roadmap.json

# View roadmap progress
prism roadmap progress examples/goal-roadmap.json
```

## See Also

- [Goals Concept](../concepts/goals.md)
- [Phases Concept](../concepts/phases.md)
- [Operations Metrics Example](operations.md)

For security roadmap examples, see [prism-security](https://github.com/grokify/prism-security).
