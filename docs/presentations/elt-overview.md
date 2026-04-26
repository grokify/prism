---
marp: true
theme: default
paginate: true
header: 'PRISM Framework'
footer: 'Confidential'
---

<!-- _class: lead -->

# PRISM

## Platform for Reliability, Improvement, and Strategic Maturity

**Operational Product Management for COO-Level Health Monitoring**

---

# The Challenge

## Fragmented Metrics Across the Organization

- **Security** tracks vulnerabilities, compliance, threat detection
- **Engineering** tracks DORA metrics, SLOs, incidents
- **Quality** tracks coverage, defects, test results
- **Product** tracks adoption, activation, churn
- **Support** tracks tickets, resolution time, CSAT

**Result:** No unified view of organizational health

---

# The Impact

## What Fragmentation Costs Us

| Problem | Impact |
|---------|--------|
| Siloed dashboards | Leaders see different pictures |
| Inconsistent maturity | Some areas advanced, others neglected |
| Unclear accountability | Who owns cross-cutting metrics? |
| Competing priorities | No framework for trade-offs |
| Reactive posture | Firefighting instead of improving |

---

# PRISM: The Solution

## One Framework, Unified View

**PRISM** provides:

- **Unified metrics** across all domains
- **Maturity model** to measure capability
- **Clear accountability** through team ownership
- **Progression strategies** to guide improvement
- **Roadmap** to track advancement

---

# The PRISM Model

## Three Dimensions

```
┌─────────────────────────────────────────────────────────────────┐
│                      VALUE STREAM LAYERS                        │
│   Requirements → Code → Infra → Runtime → Adoption → Support    │
└─────────────────────────────────────────────────────────────────┘
                              │
        ┌─────────────────────┼─────────────────────┐
        │                     │                     │
        ▼                     ▼                     ▼
  ┌───────────┐         ┌───────────┐         ┌───────────┐
  │ OPERATIONS│         │ SECURITY  │         │  QUALITY  │
  └───────────┘         └───────────┘         └───────────┘
                              │
┌─────────────────────────────────────────────────────────────────┐
│                      LIFECYCLE STAGES                           │
│          Design → Build → Test → Runtime → Response             │
└─────────────────────────────────────────────────────────────────┘
```

---

# Value Stream Layers

## From Ideation to Support

| Layer | Description | Owner |
|-------|-------------|-------|
| **Requirements** | Product ideation, specs, design | Product |
| **Code** | Application code, libraries, dependencies | Engineering |
| **Infra** | Cloud resources, networking, platform | Platform |
| **Runtime** | Running services, production workloads | SRE |
| **Adoption** | Product analytics, user engagement | Product/Growth |
| **Support** | Customer support, incident management | Support |

---

# Three Domains

## Functional Areas with Standards

| Domain | Focus | Key Metrics |
|--------|-------|-------------|
| **Operations** | Reliability, efficiency | DORA, SLOs, golden signals |
| **Security** | Protection, compliance | Vulnerabilities, MTTR, compliance |
| **Quality** | Testing, defects | Coverage, defect density, escape rate |

Each domain has an **overlay team** that sets standards and measures across all layers.

---

# Team Accountability

## Team Topologies Integration

| Team Type | Role | Accountability |
|-----------|------|----------------|
| **Stream-Aligned** | Build and run services | Layer metrics (code, runtime) |
| **Platform** | Provide infrastructure | Infra layer metrics |
| **Enabling** | Help teams adopt practices | Cross-team improvement |
| **Overlay** | Define domain standards | Domain-wide metrics |

**Every metric has an owner. Every team knows their KPIs.**

---

# Maturity Model

## 5 Levels of Capability

| Level | Name | Description |
|-------|------|-------------|
| **M1** | Reactive | Ad-hoc, firefighting, heroics |
| **M2** | Basic | Documented, some repeatability |
| **M3** | Defined | Standardized, consistent execution |
| **M4** | Managed | Data-driven, measured, controlled |
| **M5** | Optimizing | Continuous improvement, automated |

---

# Maturity Examples

## What Each Level Looks Like

| Level | Operations | Security | Quality |
|-------|------------|----------|---------|
| M1 | No SLOs, manual deploys | Ad-hoc scanning | No automation |
| M2 | Basic monitoring | Scheduled scans | Some unit tests |
| M3 | SLOs defined, CI/CD | Integrated scanning | Test automation |
| M4 | Error budgets, DORA | Real-time detection | Quality gates |
| M5 | Self-healing, <1h lead time | Auto-remediation | Continuous testing |

---

# Current State

## Where We Are Today (Example)

|  | Requirements | Code | Infra | Runtime | Adoption | Support |
|--|--------------|------|-------|---------|----------|---------|
| **Operations** | M2 | M3 | M3 | M3 | M2 | M2 |
| **Security** | M2 | M3 | M2 | M3 | M1 | M2 |
| **Quality** | M2 | M3 | M2 | M2 | M2 | M2 |

**Average Maturity: M2.4 (Basic+)**

*Note: Replace with actual assessment data*

---

# Progression Strategies

## How Should We Advance?

| Strategy | Approach | Best For |
|----------|----------|----------|
| **Spike** | One area to M5, others M2 | Competitive differentiation |
| **Wave** | All together (M2→M3→M4→M5) | Balanced, low risk |
| **Targeted** | Priority to M4, others M3 | Resource-constrained |
| **Foundation** | All to M3, then selective | Standardization first |

---

# Strategy Comparison

## Trade-offs

```
SPIKE                    WAVE                     TARGETED

M5 ████                  M5                       M5
M4                       M4 ████████████████      M4 ████████████
M3                       M3                       M3 ████████████████
M2 ████████████████      M2                       M2
M1                       M1                       M1

Fast excellence          Balanced growth          Strategic focus
Higher risk              Lower risk               Moderate risk
```

---

# Recommended Approach

## Foundation + Targeted

**Phase 1 (6 months):** Establish foundation

- All areas to **M3** (Defined)
- Standardize processes, metrics, accountability

**Phase 2 (6 months):** Advance priorities

- Operations/Runtime to **M4**
- Security/Code to **M4**

**Phase 3 (12 months):** Optimize

- Selected areas to **M5**
- Based on business value

---

# Roadmap

## Quarterly Progression

| Quarter | Focus | Exit Criteria |
|---------|-------|---------------|
| **Q1** | Foundation | All domains M2→M3 in Code, Infra |
| **Q2** | Runtime | Operations/Runtime to M4 |
| **Q3** | Security | Security/Code, Security/Runtime to M4 |
| **Q4** | Optimization | Top 2 areas to M5 |

---

# Project Prioritization

## Scoring Framework

| Factor | Weight | Description |
|--------|--------|-------------|
| KPI Impact | 30% | Expected metric improvement |
| Maturity | 25% | Enables level progression |
| Business Value | 20% | Revenue, risk, compliance |
| Effort | 15% | Time and resources (inverse) |
| Dependencies | 10% | Blockers (inverse) |

**Projects scored 1-5 on each factor, weighted total determines priority.**

---

# Success Metrics

## How We'll Measure Progress

| Metric | Current | 6 Month | 12 Month |
|--------|---------|---------|----------|
| Average Maturity | M2.4 | M3.0 | M3.5 |
| SLO Compliance | 85% | 95% | 99% |
| MTTR | 4 hours | 2 hours | 1 hour |
| Deployment Frequency | Weekly | Daily | On-demand |
| Security Findings | 50 | 25 | 10 |

---

# Investment Required

## Resources and Commitment

| Area | Investment | Purpose |
|------|------------|---------|
| **Tooling** | Platform improvements | Automation, observability |
| **Training** | Team enablement | SRE, security, quality practices |
| **Headcount** | Enabling team | Cross-functional improvement |
| **Time** | Team allocation | 20% on maturity initiatives |

---

# Risk and Mitigation

## What Could Go Wrong

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Resource contention | High | Clear prioritization framework |
| Strategy drift | Medium | Quarterly strategy reviews |
| Measurement theater | Medium | SLO-backed level requirements |
| Team resistance | Low | Clear communication, quick wins |

---

# Cross-Functional Alignment

## Department Engagement

| Department | Role | Benefit |
|------------|------|---------|
| **Security** | Define security standards | Unified security posture |
| **Engineering** | Implement, measure | Clear expectations |
| **Product** | Requirements, adoption metrics | Product health visibility |
| **Support** | Support layer ownership | Integrated view |
| **QE** | Quality standards | Consistent quality |

---

# Next Steps

## Immediate Actions

1. **Approve** PRISM as the unified metrics framework
2. **Conduct** baseline maturity assessment
3. **Select** progression strategy
4. **Assign** domain overlay owners
5. **Kickoff** Q1 initiatives

---

# Discussion

## Questions?

**Key Decisions Needed:**

- [ ] Approve PRISM framework adoption
- [ ] Select progression strategy (recommend: Foundation + Targeted)
- [ ] Allocate resources for Q1 initiatives
- [ ] Assign executive sponsors per domain

---

<!-- _class: lead -->

# Appendix

---

# PRISM Score Calculation

## Composite Health Score

```
CellScore = (MaturityWeight × MaturityScore) + (PerformanceWeight × PerformanceScore)

Overall = Σ(CellScore × Weight) / Σ(Weight)
```

**Default Weights:**

- Maturity: 40%
- Performance: 60%

**Score Interpretation:**

- >= 0.90: Elite
- >= 0.75: Strong
- >= 0.50: Medium
- < 0.50: Needs Improvement

---

# Framework Mappings

## Industry Alignment

| Framework | PRISM Mapping |
|-----------|---------------|
| **DORA** | Operations metrics (deploy freq, lead time, MTTR, CFR) |
| **SRE** | Golden signals (latency, traffic, errors, saturation) |
| **NIST CSF** | Security domain stages |
| **MITRE ATT&CK** | Security detection metrics |
| **ISO 25010** | Quality characteristics |

---

# Glossary

| Term | Definition |
|------|------------|
| **SLO** | Service Level Objective - target for a metric |
| **SLI** | Service Level Indicator - the measurement |
| **DORA** | DevOps Research and Assessment metrics |
| **MTTR** | Mean Time to Recovery |
| **CFR** | Change Failure Rate |
| **Golden Signals** | Latency, Traffic, Errors, Saturation |
