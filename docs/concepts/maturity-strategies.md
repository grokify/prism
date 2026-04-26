# Maturity Progression Strategies

Organizations can choose different strategies for advancing maturity across domains and layers. Each strategy has trade-offs between speed, risk, resources, and organizational alignment.

## Overview

| Strategy | Approach | Speed | Risk | Resources |
|----------|----------|-------|------|-----------|
| Spike | One area to M5, others at M2 | Fast (focused) | Higher | Concentrated |
| Wave | All areas together (2→3→4→5) | Slower | Lower | Distributed |
| Targeted | Priority to M4, others to M3 | Moderate | Moderate | Balanced |
| Foundation First | All to M3, then selective | Slower initially | Lowest | Front-loaded |

## Strategy 1: Spike

**Description:** Advance one critical capability to Level 5 (Optimizing) while maintaining other areas at Level 2 (Basic).

```
                M5  ████████████████████  ← Critical Area
                M4
                M3
                M2  ████  ████  ████  ████  ← Other Areas
                M1
                    Area1 Area2 Area3 Area4
```

### When to Use

- Competitive differentiation requires excellence in a specific area
- One capability is on the critical path for business success
- Need to demonstrate "what great looks like" to the organization
- Executive mandate for a specific initiative (e.g., "security first")

### Pros

- **Rapid excellence** in the most important area
- **Creates champions** who can help other teams later
- **Demonstrates value** of maturity investment quickly
- **Focused resources** maximize impact

### Cons

- **Organizational imbalance** may cause friction
- **Dependencies unmet** when advanced area needs immature inputs
- **Isolated excellence** doesn't lift overall capability
- **Burnout risk** on the spiked team

### Example

A fintech company spikes **Security** to Level 5:

| Domain | Layer | Current | Target | Rationale |
|--------|-------|---------|--------|-----------|
| Security | Code | M2 | M5 | Regulatory requirement |
| Security | Runtime | M2 | M5 | Customer trust |
| Operations | Runtime | M2 | M2 | Maintain baseline |
| Quality | Code | M2 | M2 | Maintain baseline |

### Risk Mitigation

- Ensure M2 baseline is truly stable before spiking
- Plan knowledge transfer from spike team to others
- Set timeline for bringing other areas to M3

---

## Strategy 2: Wave

**Description:** Advance all areas together through each maturity level: M2 → M3 → M4 → M5.

```
    Wave 1: All to M2    Wave 2: All to M3    Wave 3: All to M4

    M5                   M5                   M5
    M4                   M4                   M4  ████████████████
    M3                   M3  ████████████████ M3
    M2  ████████████████ M2                   M2
    M1                   M1                   M1
```

### When to Use

- Organization values consistency and fairness across teams
- Regulatory or compliance requirements apply uniformly
- Interconnected systems require balanced capabilities
- Risk tolerance is low

### Pros

- **Balanced growth** across the organization
- **Shared learning** as teams progress together
- **No capability gaps** between areas
- **Organizational alignment** and culture building

### Cons

- **Slower to show excellence** in any single area
- **Coordination overhead** across many teams
- **Lowest common denominator** may slow fast teams
- **Resource contention** when all teams need similar skills

### Example

A healthcare company uses Wave strategy for compliance:

| Phase | Duration | All Domains/Layers Target |
|-------|----------|---------------------------|
| Q1-Q2 | 6 months | All to M2 (Basic) |
| Q3-Q4 | 6 months | All to M3 (Defined) |
| Year 2 H1 | 6 months | All to M4 (Managed) |
| Year 2 H2 | 6 months | All to M5 (Optimizing) |

### Risk Mitigation

- Allow some flexibility for teams that are ready to advance
- Define clear "wave completion" criteria before advancing
- Create peer support networks across teams at same level

---

## Strategy 3: Targeted

**Description:** Advance priority areas to Level 4 (Managed) while bringing others to Level 3 (Defined).

```
                M5
                M4  ████████████  ← Priority Areas
                M3  ████  ████  ████  ← Supporting Areas
                M2
                M1
                    Pri1  Pri2  Sup1  Sup2
```

### When to Use

- Clear strategic priorities exist
- Resources are constrained
- Some areas have higher business impact than others
- Need balance between excellence and foundation

### Pros

- **Strategic focus** on what matters most
- **Pragmatic allocation** of limited resources
- **Foundation built** across all areas (M3)
- **Flexibility** to adjust priorities

### Cons

- **Requires clear prioritization** (may be politically difficult)
- **Some teams may feel deprioritized**
- **M3 areas may stagnate** without advancement path

### Example

An e-commerce company targets customer-facing capabilities:

| Area | Priority | Target | Rationale |
|------|----------|--------|-----------|
| Operations/Runtime | High | M4 | Customer experience |
| Operations/Response | High | M4 | Incident impact |
| Security/Runtime | High | M4 | Trust and compliance |
| Quality/Test | Medium | M3 | Foundation |
| Security/Code | Medium | M3 | Foundation |
| Operations/Build | Medium | M3 | Foundation |

### Risk Mitigation

- Communicate prioritization criteria transparently
- Create advancement path for M3 areas
- Review priorities quarterly

---

## Strategy 4: Foundation First

**Description:** Bring all areas to Level 3 (Defined) first, then selectively advance based on business value.

```
    Phase 1: Foundation          Phase 2: Selective Advancement

    M5                           M5  ████
    M4                           M4  ████  ████
    M3  ████████████████████     M3  ████  ████  ████  ████
    M2                           M2
    M1                           M1
```

### When to Use

- Large organization with inconsistent practices
- M&A integration requiring standardization
- Technical debt creates unpredictable outcomes
- Need to establish baseline before optimizing

### Pros

- **Consistent baseline** across organization
- **Reduces variability** and risk
- **Builds organizational capability** broadly
- **Informed decisions** about where to invest further

### Cons

- **Delays competitive advantages** from advanced maturity
- **May frustrate high-performers** who want to advance
- **Front-loaded investment** before seeing differentiation

### Example

A company post-acquisition standardizes first:

| Phase | Focus | Target | Duration |
|-------|-------|--------|----------|
| 1 | Foundation | All areas to M3 | 12 months |
| 2a | Differentiate | Security to M4 | 6 months |
| 2b | Differentiate | Operations/Runtime to M4 | 6 months |
| 3 | Optimize | Selected areas to M5 | 12 months |

### Risk Mitigation

- Set clear timeline for Phase 2 to maintain momentum
- Identify and protect high-performing pockets
- Celebrate M3 achievements to maintain morale

---

## Decision Matrix

Use this matrix to select the appropriate strategy:

| Factor | Spike | Wave | Targeted | Foundation |
|--------|-------|------|----------|------------|
| **Time pressure** | High | Low | Medium | Low |
| **Resource constraints** | High | Low | Medium | Medium |
| **Risk tolerance** | High | Low | Medium | Low |
| **Organizational consistency** | Low priority | High priority | Medium | High priority |
| **Clear priorities** | Yes | Not needed | Yes | Not needed |
| **Regulatory pressure** | Specific area | Uniform | Mixed | Uniform |
| **Current state** | Some areas strong | Uniform low | Mixed | Uniform low |

### Decision Flow

```
START
  │
  ├─► Is there a critical area that MUST be excellent?
  │     │
  │     ├─► YES: Consider SPIKE
  │     │
  │     └─► NO: Continue
  │
  ├─► Is the current state highly inconsistent?
  │     │
  │     ├─► YES: Consider FOUNDATION FIRST
  │     │
  │     └─► NO: Continue
  │
  ├─► Are resources constrained with clear priorities?
  │     │
  │     ├─► YES: Consider TARGETED
  │     │
  │     └─► NO: Consider WAVE
  │
  END
```

---

## Hybrid Approaches

Organizations often combine strategies:

### Spike + Wave

1. Spike one critical area to M4
2. Wave all areas to M3
3. Continue advancing the spiked area to M5
4. Wave remaining areas to M4

### Foundation + Targeted

1. Foundation: All areas to M3
2. Targeted: Priority areas to M4
3. Selective: Highest-value area to M5

### Progressive Targeting

1. Start with 2 priority areas to M4
2. Add 2 more areas to M4 each quarter
3. Eventually all areas at M4, select few to M5

---

## Measuring Progress

Regardless of strategy, track:

| Metric | Description | Frequency |
|--------|-------------|-----------|
| Maturity Score | Current level per area | Monthly |
| Level Transitions | Areas that moved up | Quarterly |
| SLO Compliance | % of SLOs met | Weekly |
| Initiative Completion | Projects done vs planned | Monthly |
| Time in Level | Months at current level | Quarterly |

---

## Common Pitfalls

### 1. Strategy Drift

Starting with one strategy but drifting to another without intentional decision.

**Mitigation:** Quarterly strategy review, document strategy choice

### 2. Premature Advancement

Advancing to next level before current level is stable.

**Mitigation:** Define clear level completion criteria, require sign-off

### 3. Ignoring Dependencies

Advancing one area that depends on immature inputs.

**Mitigation:** Map dependencies, coordinate advancement

### 4. Resource Starvation

Spreading resources too thin across all areas.

**Mitigation:** Match strategy to available resources, prioritize

### 5. Measurement Theater

Claiming level advancement without real capability improvement.

**Mitigation:** Require SLO compliance for level claims, audit

---

## Next Steps

1. **Assess current state** across all domains and layers
2. **Identify constraints** (resources, time, priorities)
3. **Select strategy** using decision matrix
4. **Define phase targets** with timelines
5. **Communicate strategy** to all stakeholders
6. **Track progress** with consistent metrics
