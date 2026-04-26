# Project Scoring Framework

A structured approach to ranking and prioritizing improvement projects based on their ability to advance KPIs and maturity levels.

## Overview

Projects compete for limited resources. This framework provides objective criteria for prioritization, ensuring investments align with strategic maturity goals.

## Scoring Factors

| Factor | Weight | Description |
|--------|--------|-------------|
| KPI Impact | 30% | Expected improvement to target KPIs |
| Maturity Advancement | 25% | Enables progression to next maturity level |
| Business Value | 20% | Revenue, cost, risk, or compliance impact |
| Effort | 15% | Time, resources, and complexity (inverse) |
| Dependencies | 10% | External blockers and prerequisites (inverse) |

## Scoring Scale

Each factor is scored 1-5:

| Score | KPI Impact | Maturity | Business Value | Effort | Dependencies |
|-------|------------|----------|----------------|--------|--------------|
| 5 | >50% improvement | Enables M4→M5 | Critical/mandatory | <1 month | None |
| 4 | 25-50% improvement | Enables M3→M4 | High revenue/risk | 1-3 months | Minor |
| 3 | 10-25% improvement | Enables M2→M3 | Moderate impact | 3-6 months | Manageable |
| 2 | 5-10% improvement | Supports current level | Low impact | 6-12 months | Significant |
| 1 | <5% improvement | No maturity impact | Minimal impact | >12 months | Blocking |

## Priority Score Calculation

```
Priority = (KPI × 0.30) + (Maturity × 0.25) + (Business × 0.20) + (Effort × 0.15) + (Dependencies × 0.10)
```

**Score Interpretation:**

| Score Range | Priority | Action |
|-------------|----------|--------|
| 4.0 - 5.0 | Critical | Execute immediately |
| 3.0 - 3.9 | High | Execute this quarter |
| 2.0 - 2.9 | Medium | Plan for next quarter |
| 1.0 - 1.9 | Low | Backlog / reconsider |

---

## Factor Details

### KPI Impact (30%)

Measures the expected improvement to target KPIs.

**Assessment Questions:**

- Which KPIs will this project improve?
- What is the expected percentage improvement?
- Is the improvement measurable and attributable?
- How quickly will the improvement be visible?

**Scoring Guidelines:**

| Score | Criteria | Example |
|-------|----------|---------|
| 5 | >50% improvement to critical KPI | Reduce MTTR from 4h to 1h |
| 4 | 25-50% improvement | Improve deployment frequency 2x |
| 3 | 10-25% improvement | Increase test coverage 70%→85% |
| 2 | 5-10% improvement | Reduce error rate 0.5%→0.45% |
| 1 | <5% improvement | Marginal documentation updates |

### Maturity Advancement (25%)

Measures whether the project enables progression to the next maturity level.

**Assessment Questions:**

- Does this project fulfill a level requirement?
- Is this a prerequisite for the next level?
- Does it close a gap blocking advancement?
- How many areas does it advance?

**Scoring Guidelines:**

| Score | Criteria | Example |
|-------|----------|---------|
| 5 | Required for M4→M5, multiple areas | Self-healing automation |
| 4 | Required for M3→M4 | Implement SLO-based alerting |
| 3 | Required for M2→M3 | Standardize incident runbooks |
| 2 | Supports current level | Improve existing automation |
| 1 | No maturity impact | Cosmetic improvements |

### Business Value (20%)

Measures the direct business impact of the project.

**Assessment Questions:**

- Does this generate revenue or reduce costs?
- Does this mitigate significant risks?
- Is this required for compliance?
- Does this improve customer experience?

**Scoring Guidelines:**

| Score | Criteria | Example |
|-------|----------|---------|
| 5 | Mandatory (compliance/legal) | SOC 2 requirement |
| 4 | High revenue/major risk reduction | Prevent $1M+ incidents |
| 3 | Moderate business impact | Improve customer NPS |
| 2 | Low but positive impact | Developer productivity |
| 1 | Minimal business impact | Internal tooling polish |

### Effort (15%)

Measures the resources and time required (inverse scoring - lower effort = higher score).

**Assessment Questions:**

- How many person-months of effort?
- What skills are required?
- Is external help needed?
- What is the calendar time to completion?

**Scoring Guidelines:**

| Score | Criteria | Example |
|-------|----------|---------|
| 5 | <1 month, existing skills | Configure existing tool |
| 4 | 1-3 months, minor hiring | Small feature development |
| 3 | 3-6 months, team effort | New system implementation |
| 2 | 6-12 months, cross-team | Platform migration |
| 1 | >12 months, major investment | Full architecture redesign |

### Dependencies (10%)

Measures external factors that could block or delay the project (inverse scoring).

**Assessment Questions:**

- What must be completed first?
- Are other teams required?
- Are there vendor dependencies?
- Are there organizational blockers?

**Scoring Guidelines:**

| Score | Criteria | Example |
|-------|----------|---------|
| 5 | No dependencies | Standalone improvement |
| 4 | Minor internal dependencies | Needs config from another team |
| 3 | Manageable dependencies | Requires API from another service |
| 2 | Significant dependencies | Needs platform team work first |
| 1 | Blocking dependencies | Waiting on vendor/legal/exec |

---

## Scoring Process

### Step 1: Project Intake

Collect project proposals with:

- [ ] Project name and description
- [ ] Target KPIs and expected improvement
- [ ] Maturity level advancement claim
- [ ] Business justification
- [ ] Effort estimate
- [ ] Known dependencies

### Step 2: Initial Scoring

Project owner scores each factor with justification:

```
Project: Implement Automated Rollback
==========================================
KPI Impact:      4  (Reduce MTTR by 40%)
Maturity:        4  (Required for M3→M4 in Operations/Response)
Business Value:  4  (Prevent extended outages)
Effort:          3  (3 months, platform team)
Dependencies:    4  (Needs observability in place - done)
------------------------------------------
Weighted Score:  3.85  → HIGH PRIORITY
```

### Step 3: Review and Calibration

Scoring committee reviews and calibrates:

- Compare similar projects for consistency
- Challenge optimistic estimates
- Validate maturity advancement claims
- Confirm dependency status

### Step 4: Ranking

Sort projects by priority score:

| Rank | Project | Score | Priority |
|------|---------|-------|----------|
| 1 | Automated Rollback | 3.85 | High |
| 2 | SLO Dashboard | 3.60 | High |
| 3 | Incident Runbooks | 3.25 | High |
| 4 | Code Coverage Tooling | 2.90 | Medium |
| 5 | Documentation Update | 2.10 | Medium |

### Step 5: Resource Allocation

Match high-priority projects to available resources:

- Critical (4.0+): Assign immediately
- High (3.0-3.9): Assign this quarter
- Medium (2.0-2.9): Queue for next quarter
- Low (<2.0): Backlog or decline

---

## Example: Full Scoring

### Project: Implement SLO-Based Alerting

**Description:** Replace threshold-based alerts with SLO-based alerting using error budgets.

**Factor Scoring:**

| Factor | Score | Weight | Weighted | Justification |
|--------|-------|--------|----------|---------------|
| KPI Impact | 4 | 0.30 | 1.20 | Reduce alert fatigue 50%, improve MTTR 25% |
| Maturity | 5 | 0.25 | 1.25 | Required for Operations M4 (Managed) |
| Business | 4 | 0.20 | 0.80 | Reduce incident costs, improve reliability |
| Effort | 3 | 0.15 | 0.45 | 4 months, SRE team + platform support |
| Dependencies | 4 | 0.10 | 0.40 | Needs SLI instrumentation (80% done) |
| **Total** | | | **4.10** | **CRITICAL** |

**Recommendation:** Execute immediately, assign dedicated team.

---

### Project: Update Security Documentation

**Description:** Update security policies and procedures documentation.

**Factor Scoring:**

| Factor | Score | Weight | Weighted | Justification |
|--------|-------|--------|----------|---------------|
| KPI Impact | 2 | 0.30 | 0.60 | Marginal improvement to compliance metrics |
| Maturity | 2 | 0.25 | 0.50 | Supports M2, not required for M3 |
| Business | 3 | 0.20 | 0.60 | Needed for audit, moderate compliance value |
| Effort | 5 | 0.15 | 0.75 | 2 weeks, existing staff |
| Dependencies | 5 | 0.10 | 0.50 | None |
| **Total** | | | **2.95** | **MEDIUM** |

**Recommendation:** Plan for next quarter, combine with other doc updates.

---

## Governance

### Scoring Committee

| Role | Responsibility |
|------|----------------|
| Domain Lead | Score projects in their domain |
| Engineering Lead | Validate effort estimates |
| Product Lead | Validate business value |
| Program Manager | Facilitate, track, report |

### Cadence

| Activity | Frequency |
|----------|-----------|
| Project intake | Ongoing |
| Scoring sessions | Bi-weekly |
| Priority review | Monthly |
| Full re-ranking | Quarterly |

### Escalation

Projects can request re-scoring when:

- New information changes estimates
- Dependencies are resolved
- Business priorities shift
- Resources become available

---

## Integration with PRISM

### Link Projects to Goals

```json
{
  "initiatives": [
    {
      "id": "init-slo-alerting",
      "name": "SLO-Based Alerting",
      "goalIds": ["goal-reliability"],
      "phaseId": "phase-q2-2026",
      "priority": 1,
      "scores": {
        "kpiImpact": 4,
        "maturityAdvancement": 5,
        "businessValue": 4,
        "effort": 3,
        "dependencies": 4,
        "priorityScore": 4.10
      }
    }
  ]
}
```

### Track in Phase Metrics

Projects appear in phase progress:

```bash
prism phase metrics prism.json phase-q2-2026
```

Output:

```
Phase: Q2 2026
==============

Initiatives:
  Total: 5
  Completed: 2
  In Progress: 2
  Planned: 1

By Priority:
  Critical (4.0+): 1 in progress
  High (3.0-3.9):  2 (1 complete, 1 in progress)
  Medium (2.0-2.9): 2 planned
```

---

## Templates

### Project Scoring Template

```markdown
## Project: [Name]

**Description:** [One paragraph]

**Target KPIs:**
- [ ] KPI 1: Current → Target
- [ ] KPI 2: Current → Target

**Maturity Impact:**
- Domain: [domain]
- Layer: [layer]
- Current Level: M[n]
- Target Level: L[n+1]
- Requirement Met: [which requirement]

**Scoring:**

| Factor | Score | Justification |
|--------|-------|---------------|
| KPI Impact | | |
| Maturity | | |
| Business Value | | |
| Effort | | |
| Dependencies | | |
| **Weighted Total** | | |

**Recommendation:** [Critical/High/Medium/Low]
```

### Quarterly Review Template

```markdown
## Q[N] Project Review

**Projects Completed:** X
**Projects In Progress:** Y
**Projects Deferred:** Z

### Completed Projects

| Project | Score | Outcome |
|---------|-------|---------|
| | | |

### Scoring Accuracy

| Factor | Avg Estimate | Avg Actual | Variance |
|--------|--------------|------------|----------|
| KPI Impact | | | |
| Effort | | | |

### Lessons Learned

1.
2.
3.

### Next Quarter Priorities

1.
2.
3.
```

---

## Next Steps

1. **Create project intake form** based on template
2. **Establish scoring committee** with domain leads
3. **Score existing backlog** to establish baseline
4. **Communicate framework** to all teams
5. **Track outcomes** to calibrate scoring accuracy
