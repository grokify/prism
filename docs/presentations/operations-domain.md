---
marp: true
theme: default
paginate: true
header: 'PRISM: Operations Domain'
footer: 'Engineering & SRE Leadership'
---

<!-- _class: lead -->

# Operations Domain

## PRISM Maturity Framework

**Reliability, Efficiency, and Performance**

---

# Operations Overview

## What We Measure

The Operations domain focuses on **system reliability and delivery efficiency**:

- **Reliability** - Availability, latency, error rates
- **Efficiency** - Deployment frequency, lead time, resource utilization
- **Response** - Incident detection, recovery, post-mortems

**Goal:** Deliver reliable services efficiently with rapid recovery

---

# Operations Across the Value Stream

## Layer Responsibilities

| Layer | Operations Focus | Key Metrics |
|-------|------------------|-------------|
| **Code** | Build quality, CI/CD | Build success rate, lead time |
| **Infra** | Platform reliability | Resource utilization, drift |
| **Runtime** | Service health | Availability, latency, errors |
| **Support** | Incident management | MTTR, escalation rate |

---

# DORA Metrics

## Industry-Standard Delivery Metrics

| Metric | Elite | High | Medium | Low |
|--------|-------|------|--------|-----|
| **Deployment Frequency** | On-demand | Daily-Weekly | Weekly-Monthly | Monthly+ |
| **Lead Time for Changes** | <1 hour | 1 day-1 week | 1 week-1 month | 1 month+ |
| **Mean Time to Recovery** | <1 hour | <1 day | 1 day-1 week | 1 week+ |
| **Change Failure Rate** | 0-15% | 16-30% | 31-45% | 46%+ |

**Where are we today?**

---

# SRE Golden Signals

## Per-Layer Observability

| Signal | Description | Example KPIs |
|--------|-------------|--------------|
| **Latency** | Time to serve requests | P50, P95, P99 response time |
| **Traffic** | Demand on the system | Requests/sec, concurrent users |
| **Errors** | Failed requests | Error rate %, 5xx count |
| **Saturation** | Resource utilization | CPU, memory, queue depth |

Each layer defines its golden signals in PRISM.

---

# SLOs and Error Budgets

## Target-Based Reliability

| Service | SLO | Current | Error Budget |
|---------|-----|---------|--------------|
| API Gateway | 99.95% availability | 99.92% | -3 hours |
| Payments | 99.99% availability | 99.995% | +4 hours |
| User Service | P99 < 200ms | 185ms | Within budget |

**Error Budget = Allowed failures before SLO breach**

---

# Operations Maturity Levels

## What Each Level Means

| Level | Description | Indicators |
|-------|-------------|------------|
| **M1** | Reactive | No SLOs, manual deploys, firefighting |
| **M2** | Basic | Basic monitoring, some CI/CD, incident response |
| **M3** | Defined | SLOs defined, CI/CD standard, runbooks |
| **M4** | Managed | Error budgets, DORA tracked, auto-scaling |
| **M5** | Optimizing | Self-healing, <1h lead time, proactive |

---

# M1 → M2: Basic

## Establishing Foundations

**Requirements:**

- [ ] Basic monitoring in place (uptime checks)
- [ ] CI/CD pipeline exists (even if inconsistent)
- [ ] Incident response process documented
- [ ] On-call rotation established

**Key Metrics:**

- Uptime measured
- Deployments tracked
- Incidents logged

---

# M2 → M3: Defined

## Standardization

**Requirements:**

- [ ] SLOs defined for all tier-1 services
- [ ] CI/CD pipelines standardized across teams
- [ ] Runbooks exist for common incidents
- [ ] DORA metrics collected

**Key Metrics:**

- SLO compliance > 95%
- Deployment frequency ≥ weekly
- Runbook coverage > 80%

---

# M3 → M4: Managed

## Data-Driven Operations

**Requirements:**

- [ ] Error budgets implemented and enforced
- [ ] DORA metrics meet "High Performer" thresholds
- [ ] Auto-scaling based on traffic patterns
- [ ] Incident post-mortems with action items

**Key Metrics:**

- MTTR < 1 day
- Change failure rate < 30%
- Lead time < 1 week
- Error budget tracking automated

---

# M4 → M5: Optimizing

## Continuous Improvement

**Requirements:**

- [ ] Self-healing for common failure modes
- [ ] Lead time < 1 hour (on-demand deploys)
- [ ] Proactive capacity management
- [ ] Chaos engineering practices

**Key Metrics:**

- MTTR < 1 hour
- Change failure rate < 15%
- Deployment frequency: on-demand
- Zero manual incident remediation for known issues

---

# Current State Assessment

## Operations Maturity by Layer

| Layer | Current | Target | Gap |
|-------|---------|--------|-----|
| Code | M3 | M4 | CI/CD metrics, quality gates |
| Infra | M3 | M4 | IaC compliance, drift detection |
| Runtime | M3 | M4 | Error budgets, auto-remediation |
| Support | M2 | M3 | Runbooks, incident process |

*Replace with actual assessment*

---

# Operations KPIs

## What We Track

| KPI | Current | Target | Owner |
|-----|---------|--------|-------|
| Deployment Frequency | 2/week | Daily | Platform |
| Lead Time | 5 days | 1 day | Platform |
| MTTR | 4 hours | 1 hour | SRE |
| Change Failure Rate | 20% | 10% | Engineering |
| Availability | 99.9% | 99.95% | SRE |
| P99 Latency | 250ms | 150ms | Service Teams |

---

# Improvement Projects

## Prioritized Initiatives

| Project | Score | Impact | Timeline |
|---------|-------|--------|----------|
| SLO-based alerting | 4.1 | M3→M4 Runtime | Q1 |
| Automated rollback | 3.9 | M3→M4 Runtime | Q1 |
| DORA dashboard | 3.6 | M3→M4 Code | Q1 |
| Runbook standardization | 3.3 | M2→M3 Support | Q2 |
| Chaos engineering | 3.0 | M4→M5 Runtime | Q3 |

---

# Project: SLO-Based Alerting

## Deep Dive

**Current State:** Threshold-based alerts, high noise, alert fatigue

**Target State:** SLO-based alerts using error budgets

**Scoring:**

| Factor | Score | Justification |
|--------|-------|---------------|
| KPI Impact | 4 | Reduce MTTR 25%, eliminate alert fatigue |
| Maturity | 5 | Required for M4 (Managed) |
| Business | 4 | Fewer incidents, faster recovery |
| Effort | 3 | 4 months implementation |
| Dependencies | 4 | SLI instrumentation 80% complete |

**Priority Score: 4.1 (Critical)**

---

# Roadmap

## Quarterly Plan

```
Q1 2026                    Q2 2026                    Q3 2026
─────────────────────────────────────────────────────────────────

SLO Alerting ████████████

Auto Rollback   ████████████

DORA Dashboard     ████████

                         Runbooks ████████████

                         IaC Compliance ████████

                                          Chaos Eng ██████████
```

---

# Dependencies

## Cross-Team Coordination

| Dependency | Team | Status | Blocker? |
|------------|------|--------|----------|
| SLI instrumentation | Platform | 80% done | No |
| Observability stack | Platform | Complete | No |
| Runbook templates | SRE | In progress | No |
| Change management | Process | Needed | Minor |

---

# Collaboration

## Working with Other Domains

| Domain | Collaboration Area |
|--------|-------------------|
| **Security** | Runtime security monitoring, incident response |
| **Quality** | Test automation in CI/CD, quality gates |
| **Product** | Adoption metrics, feature flags |

**Operations enables reliability across all domains.**

---

# Success Criteria

## How We Know We've Succeeded

| Milestone | Target Date | Criteria |
|-----------|-------------|----------|
| M3 Complete | End Q1 | SLOs for all tier-1, runbook coverage >80% |
| M4 Runtime | End Q2 | Error budgets active, MTTR <2h |
| M4 Code | End Q2 | DORA "High Performer" |
| M5 Pilot | End Q4 | Self-healing for top 3 failure modes |

---

# Resource Ask

## What We Need

| Resource | Purpose | Duration |
|----------|---------|----------|
| 2 SRE engineers | SLO implementation | 6 months |
| Platform time | Alerting infrastructure | Q1 |
| Training budget | SRE practices | Ongoing |
| Tooling budget | Observability, chaos | Q2 |

---

# Discussion

## Questions?

**Decisions Needed:**

- [ ] Approve Operations maturity roadmap
- [ ] Allocate SRE resources for Q1
- [ ] Confirm DORA targets
- [ ] Assign project owners

---

<!-- _class: lead -->

# Appendix

---

# DORA Research Context

## Why These Metrics Matter

From *Accelerate* (Forsgren, Humble, Kim):

> "High performers deploy 208x more frequently, have 106x faster lead times, recover 2,604x faster, and have 7x lower change failure rate."

**DORA metrics are predictive of:**

- Organizational performance
- Employee satisfaction
- Reduced burnout

---

# SLO Definition Template

## For Each Service

```yaml
service: payments-api
tier: 1

slos:
  - name: Availability
    target: 99.99%
    window: 30d
    sli: successful_requests / total_requests

  - name: Latency
    target: 99th percentile < 200ms
    window: 7d
    sli: request_duration_histogram

error_budget:
  monthly_minutes: 4.32  # 99.99% of 43,200 minutes
```

---

# Incident Severity Levels

## Classification

| Severity | Impact | Response Time | Example |
|----------|--------|---------------|---------|
| SEV1 | Complete outage | 15 min | All services down |
| SEV2 | Major degradation | 30 min | Payments failing |
| SEV3 | Minor degradation | 2 hours | Slow response times |
| SEV4 | Low impact | Next business day | Cosmetic issue |
