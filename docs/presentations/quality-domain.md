---
marp: true
theme: default
paginate: true
header: 'PRISM: Quality Domain'
footer: 'QE Leadership'
---

<!-- _class: lead -->

# Quality Domain

## PRISM Maturity Framework

**Testing, Defects, and Quality Assurance**

---

# Quality Overview

## What We Measure

The Quality domain focuses on **software quality and testing effectiveness**:

- **Coverage** - Test coverage across code and features
- **Defects** - Finding and fixing bugs early
- **Prevention** - Quality gates that prevent escapes

**Goal:** High-quality releases with minimal production defects

---

# Quality Across the Value Stream

## Layer Responsibilities

| Layer | Quality Focus | Key Metrics |
|-------|---------------|-------------|
| **Requirements** | Acceptance criteria, testability | Spec coverage |
| **Code** | Unit tests, static analysis | Code coverage |
| **Infra** | Infrastructure testing | IaC test coverage |
| **Runtime** | Production quality monitoring | Defect escape rate |

---

# ISO 25010 Quality Characteristics

## Software Quality Model

| Characteristic | Description | Example Metrics |
|----------------|-------------|-----------------|
| **Functional** | Does it work correctly? | Test pass rate |
| **Reliability** | Does it work consistently? | Defect density |
| **Performance** | Is it fast enough? | Response time tests |
| **Security** | Is it secure? | Security test coverage |
| **Maintainability** | Can we change it? | Code complexity |
| **Usability** | Is it easy to use? | Accessibility tests |

---

# Quality Categories

## Test Types and Coverage

| Category | Stage | Description |
|----------|-------|-------------|
| **Unit Testing** | Build | Code-level verification |
| **Integration Testing** | Build/Test | Component interaction |
| **E2E Testing** | Test | Full workflow validation |
| **Performance Testing** | Test | Load and stress testing |
| **Production Monitoring** | Runtime | Quality in production |

---

# Quality KPIs

## Core Metrics

| Metric | Description | Target |
|--------|-------------|--------|
| **Code Coverage** | % of code tested | >80% |
| **Test Pass Rate** | % of tests passing | >99% |
| **Defect Density** | Defects per KLOC | <0.5 |
| **Escape Rate** | Defects found post-release | <5% |
| **MTTR (Defects)** | Time to fix defects | <3 days |
| **Flaky Test Rate** | Unreliable tests | <1% |

---

# Quality Maturity Levels

## What Each Level Means

| Level | Description | Indicators |
|-------|-------------|------------|
| **M1** | Reactive | Manual testing, no automation |
| **M2** | Basic | Some unit tests, basic CI |
| **M3** | Defined | Test automation, quality gates |
| **M4** | Managed | Coverage targets, defect tracking |
| **M5** | Optimizing | Continuous testing, predictive quality |

---

# M1 → M2: Basic

## Establishing Quality Foundations

**Requirements:**

- [ ] Unit testing framework in place
- [ ] CI pipeline runs tests on commit
- [ ] Basic defect tracking process
- [ ] Manual test cases documented

**Key Metrics:**

- Some test coverage (any %)
- Tests run in CI
- Defects logged

---

# M2 → M3: Defined

## Standardized Quality

**Requirements:**

- [ ] Code coverage > 60% across all repos
- [ ] Quality gates block merges below threshold
- [ ] E2E tests for critical paths
- [ ] Defect triage process defined

**Key Metrics:**

- Coverage tracked and enforced
- Quality gates active
- E2E coverage for happy paths
- Defect SLAs defined

---

# M3 → M4: Managed

## Measured Quality

**Requirements:**

- [ ] Code coverage > 80%
- [ ] Escape rate tracked and < 10%
- [ ] Performance testing automated
- [ ] Defect density tracked per team

**Key Metrics:**

- Coverage targets met
- Escape rate declining
- Performance baselines established
- Quality dashboards operational

---

# M4 → M5: Optimizing

## Continuous Quality

**Requirements:**

- [ ] Production quality monitoring
- [ ] Predictive defect analysis
- [ ] Continuous testing in production
- [ ] Quality metrics drive prioritization

**Key Metrics:**

- Escape rate < 2%
- Defect prediction accuracy > 70%
- Canary testing standard
- Quality trends improving

---

# Current State Assessment

## Quality Maturity by Layer

| Layer | Current | Target | Gap |
|-------|---------|--------|-----|
| Requirements | M2 | M3 | Acceptance criteria coverage |
| Code | M3 | M4 | Coverage targets, tracking |
| Infra | M2 | M3 | IaC testing |
| Runtime | M2 | M4 | Production monitoring |

*Replace with actual assessment*

---

# Quality KPI Dashboard

## Current vs Target

| KPI | Current | Target | Status |
|-----|---------|--------|--------|
| Code Coverage | 72% | 80% | In Progress |
| Test Pass Rate | 97% | 99% | At Risk |
| Defect Density | 0.8/KLOC | 0.5/KLOC | At Risk |
| Escape Rate | 12% | 5% | At Risk |
| Flaky Test Rate | 3% | 1% | In Progress |
| E2E Coverage | 60% | 90% | In Progress |

---

# Improvement Projects

## Prioritized Initiatives

| Project | Score | Impact | Timeline |
|---------|-------|--------|----------|
| Coverage enforcement | 4.0 | M3→M4 Code | Q1 |
| E2E test expansion | 3.8 | M2→M3 Runtime | Q1 |
| Flaky test elimination | 3.6 | Quality gate reliability | Q1 |
| Performance test automation | 3.4 | M3→M4 Test | Q2 |
| Production quality monitoring | 3.2 | M2→M4 Runtime | Q2 |

---

# Project: Coverage Enforcement

## Deep Dive

**Current State:** Coverage measured but not enforced, declining trend

**Target State:** Coverage gates block merges below 80%

**Scoring:**

| Factor | Score | Justification |
|--------|-------|---------------|
| KPI Impact | 4 | Increase coverage 72%→80% |
| Maturity | 4 | Required for M4 (Managed) |
| Business | 4 | Reduce production defects |
| Effort | 4 | 2 months, tooling exists |
| Dependencies | 4 | CI/CD infrastructure ready |

**Priority Score: 4.0 (Critical)**

---

# Project: Flaky Test Elimination

## Deep Dive

**Current State:** 3% of tests are flaky, causing CI delays

**Target State:** <1% flaky rate, reliable quality gates

**Scoring:**

| Factor | Score | Justification |
|--------|-------|---------------|
| KPI Impact | 4 | Reliable CI, faster feedback |
| Maturity | 3 | Supports M3 quality |
| Business | 3 | Developer productivity |
| Effort | 4 | 6 weeks, targeted effort |
| Dependencies | 5 | No external dependencies |

**Priority Score: 3.6 (High)**

---

# Roadmap

## Quarterly Plan

```
Q1 2026                    Q2 2026                    Q3 2026
─────────────────────────────────────────────────────────────────

Coverage Gates ██████████

E2E Expansion    ████████████████

Flaky Tests  ██████

                      Performance Tests ████████████

                      Prod Monitoring ████████████████

                                        Predictive Quality ████████
```

---

# Quality Gates

## Blocking Criteria

| Gate | Stage | Criteria | Action |
|------|-------|----------|--------|
| **Unit Test** | Build | Coverage < 80% | Block merge |
| **Integration** | Build | Tests fail | Block merge |
| **E2E** | Test | Critical path fails | Block deploy |
| **Performance** | Test | Regression > 10% | Block deploy |
| **Security** | Build | Critical findings | Block merge |

---

# Dependencies

## Cross-Team Coordination

| Dependency | Team | Status | Blocker? |
|------------|------|--------|----------|
| CI/CD pipeline access | Platform | Available | No |
| Test infrastructure | Platform | Available | No |
| Coverage tooling | Platform | Needs upgrade | Minor |
| Developer training | Engineering | Planned | No |

---

# Collaboration

## Working with Other Domains

| Domain | Collaboration Area |
|--------|-------------------|
| **Operations** | Performance testing, production monitoring |
| **Security** | Security testing integration |
| **Product** | Acceptance criteria, feature coverage |

**Quality enables confidence in releases.**

---

# Success Criteria

## How We Know We've Succeeded

| Milestone | Target Date | Criteria |
|-----------|-------------|----------|
| M3 Code | End Q1 | Coverage > 80%, gates active |
| M3 Runtime | End Q1 | E2E coverage > 80% |
| M4 Code | End Q2 | Escape rate < 8% |
| M4 Runtime | End Q3 | Production monitoring active |

---

# Resource Ask

## What We Need

| Resource | Purpose | Duration |
|----------|---------|----------|
| 1 SDET | E2E test expansion | 6 months |
| Test infrastructure | Performance testing | Q2 |
| Coverage tooling | Enhanced reporting | Q1 |
| Training | Test automation best practices | Ongoing |

---

# Discussion

## Questions?

**Decisions Needed:**

- [ ] Approve Quality maturity roadmap
- [ ] Confirm coverage targets (80%)
- [ ] Allocate SDET resource
- [ ] Assign project owners

---

<!-- _class: lead -->

# Appendix

---

# Test Pyramid

## Balanced Testing Strategy

```
                    ▲
                   /E\         E2E Tests (10%)
                  /2E \        - Slow, expensive
                 /Tests\       - Critical paths only
                /───────\
               /         \     Integration (20%)
              /Integration\    - Component boundaries
             /─────────────\   - API contracts
            /               \
           /    Unit Tests   \ Unit Tests (70%)
          /───────────────────\ - Fast, cheap
                               - High coverage
```

---

# Defect Severity Classification

## SLA by Severity

| Severity | Description | Fix SLA | Current Avg |
|----------|-------------|---------|-------------|
| Critical | System down | 4 hours | 6 hours |
| High | Major feature broken | 24 hours | 36 hours |
| Medium | Feature impaired | 1 week | 5 days |
| Low | Minor issue | 2 weeks | 10 days |

---

# Coverage by Service

## Current State

| Service | Unit | Integration | E2E |
|---------|------|-------------|-----|
| Payments | 85% | 70% | 90% |
| Users | 78% | 65% | 85% |
| Inventory | 68% | 50% | 60% |
| Reporting | 55% | 40% | 50% |
| **Average** | **72%** | **56%** | **71%** |
