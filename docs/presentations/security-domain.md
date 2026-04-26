---
marp: true
theme: default
paginate: true
header: 'PRISM: Security Domain'
footer: 'Security Leadership'
---

<!-- _class: lead -->

# Security Domain

## PRISM Maturity Framework

**Prevention, Detection, and Response**

---

# Security Overview

## What We Measure

The Security domain focuses on **protecting systems and data**:

- **Prevention** - Secure coding, vulnerability management
- **Detection** - Threat detection, anomaly identification
- **Response** - Incident response, remediation time

**Goal:** Shift security left while maintaining runtime protection

---

# Security Across the Value Stream

## Layer Responsibilities

| Layer | Security Focus | Key Metrics |
|-------|----------------|-------------|
| **Requirements** | Threat modeling, security requirements | Coverage % |
| **Code** | SAST, SCA, secrets scanning | Findings count, MTTR |
| **Infra** | Cloud security, IaC scanning | Misconfigurations |
| **Runtime** | DAST, WAF, runtime protection | Threats blocked |

---

# Security Categories

## Prevention, Detection, Response

| Category | Stage | Description |
|----------|-------|-------------|
| **Prevention** | Design, Build | Stop vulnerabilities before production |
| **Detection** | Test, Runtime | Find issues early and in production |
| **Response** | Response | Remediate quickly when issues arise |

**Balanced investment across all three is essential.**

---

# Security KPIs

## Core Metrics

| Metric | Description | Target |
|--------|-------------|--------|
| **SAST Findings** | Static analysis vulnerabilities | 0 critical/high |
| **SCA Vulnerabilities** | Dependency vulnerabilities | 0 critical |
| **Secrets in Code** | Exposed credentials | 0 |
| **MTTR (Security)** | Time to remediate | <7 days critical |
| **Cloud Misconfigs** | Infrastructure security gaps | 0 critical |
| **Threat Detection Rate** | % of threats caught | >95% |

---

# Framework Alignment

## Industry Standards

| Framework | PRISM Mapping |
|-----------|---------------|
| **NIST CSF** | Identify, Protect, Detect, Respond, Recover |
| **MITRE ATT&CK** | Detection coverage across techniques |
| **OWASP Top 10** | Code layer vulnerability categories |
| **CIS Benchmarks** | Infrastructure compliance |
| **SOC 2** | Control coverage metrics |

---

# Security Maturity Levels

## What Each Level Means

| Level | Description | Indicators |
|-------|-------------|------------|
| **M1** | Reactive | Ad-hoc scanning, no process |
| **M2** | Basic | Scheduled scans, manual review |
| **M3** | Defined | Integrated scanning, policies enforced |
| **M4** | Managed | Real-time detection, automated response |
| **M5** | Optimizing | Predictive security, auto-remediation |

---

# M1 → M2: Basic

## Establishing Security Foundations

**Requirements:**

- [ ] SAST tool deployed (even if not integrated)
- [ ] Dependency scanning scheduled
- [ ] Security incident process documented
- [ ] Basic security training completed

**Key Metrics:**

- Scans running (any frequency)
- Vulnerabilities tracked
- Incidents logged

---

# M2 → M3: Defined

## Integrated Security

**Requirements:**

- [ ] SAST/SCA integrated in CI/CD pipeline
- [ ] Security gates block critical findings
- [ ] Secrets scanning prevents commits
- [ ] Cloud security posture managed

**Key Metrics:**

- 100% of repos scanned in CI
- <24h to block critical findings
- 0 secrets in code
- Weekly cloud posture reports

---

# M3 → M4: Managed

## Real-Time Security

**Requirements:**

- [ ] Runtime protection (WAF, RASP) deployed
- [ ] Threat detection with SIEM/SOAR
- [ ] MTTR < 7 days for critical issues
- [ ] MITRE ATT&CK coverage > 70%

**Key Metrics:**

- Real-time threat visibility
- Automated incident triage
- Mean time to detect < 1 hour
- Security SLOs defined

---

# M4 → M5: Optimizing

## Proactive Security

**Requirements:**

- [ ] Auto-remediation for common vulnerabilities
- [ ] Predictive threat intelligence
- [ ] Red team/purple team exercises
- [ ] Security chaos engineering

**Key Metrics:**

- 0 manual remediation for known patterns
- Threat prediction accuracy > 80%
- Quarterly adversarial testing
- Security built into deployment pipelines

---

# Current State Assessment

## Security Maturity by Layer

| Layer | Current | Target | Gap |
|-------|---------|--------|-----|
| Requirements | M2 | M3 | Threat modeling coverage |
| Code | M3 | M4 | Real-time, faster MTTR |
| Infra | M2 | M4 | Posture management, compliance |
| Runtime | M3 | M4 | Detection coverage, automation |

*Replace with actual assessment*

---

# Security KPI Dashboard

## Current vs Target

| KPI | Current | Target | Status |
|-----|---------|--------|--------|
| Critical SAST Findings | 5 | 0 | At Risk |
| Critical SCA Vulns | 12 | 0 | At Risk |
| Secrets Detected | 0 | 0 | On Track |
| MTTR (Critical) | 14 days | 7 days | At Risk |
| Cloud Misconfigs | 25 | 0 | At Risk |
| MITRE Coverage | 55% | 80% | In Progress |

---

# Improvement Projects

## Prioritized Initiatives

| Project | Score | Impact | Timeline |
|---------|-------|--------|----------|
| Pipeline security gates | 4.2 | M2→M3 Code | Q1 |
| Cloud posture management | 4.0 | M2→M4 Infra | Q1 |
| SIEM/SOAR integration | 3.8 | M3→M4 Runtime | Q2 |
| Threat modeling program | 3.5 | M2→M3 Requirements | Q2 |
| Auto-remediation | 3.2 | M4→M5 Code | Q3 |

---

# Project: Pipeline Security Gates

## Deep Dive

**Current State:** Scans run but don't block deploys

**Target State:** Critical/high findings block production deployments

**Scoring:**

| Factor | Score | Justification |
|--------|-------|---------------|
| KPI Impact | 5 | Eliminate critical vulns in production |
| Maturity | 4 | Required for M3 (Defined) |
| Business | 5 | Compliance requirement, risk reduction |
| Effort | 4 | 2 months, existing tools |
| Dependencies | 4 | Pipeline infrastructure ready |

**Priority Score: 4.5 (Critical)**

---

# Project: Cloud Posture Management

## Deep Dive

**Current State:** Manual cloud security reviews, inconsistent

**Target State:** Continuous posture assessment, auto-remediation

**Scoring:**

| Factor | Score | Justification |
|--------|-------|---------------|
| KPI Impact | 5 | Eliminate cloud misconfigurations |
| Maturity | 5 | Enables M4 (Managed) |
| Business | 5 | SOC 2, data protection |
| Effort | 3 | 4 months, vendor selection |
| Dependencies | 3 | Cloud team collaboration |

**Priority Score: 4.4 (Critical)**

---

# Roadmap

## Quarterly Plan

```
Q1 2026                    Q2 2026                    Q3 2026
─────────────────────────────────────────────────────────────────

Pipeline Gates ██████████

Cloud Posture    ████████████████

                      SIEM/SOAR ████████████████

                      Threat Modeling ████████

                                        Auto-Remediation ████████

                                        Red Team ██████
```

---

# Compliance Alignment

## Framework Coverage

| Framework | Current | Target | Gap |
|-----------|---------|--------|-----|
| SOC 2 Type II | 85% | 100% | Access controls, logging |
| NIST CSF | 70% | 90% | Detection, response |
| PCI DSS | 90% | 100% | Encryption, access |
| GDPR | 80% | 95% | Data handling |

---

# Dependencies

## Cross-Team Coordination

| Dependency | Team | Status | Blocker? |
|------------|------|--------|----------|
| CI/CD pipeline access | Platform | Available | No |
| Cloud accounts access | Platform | Available | No |
| SIEM infrastructure | IT/Security | Planned Q2 | Yes for M4 |
| Developer training | Engineering | Ongoing | No |

---

# Collaboration

## Working with Other Domains

| Domain | Collaboration Area |
|--------|-------------------|
| **Operations** | Incident response, runtime monitoring |
| **Quality** | Security testing in QA |
| **Product** | Security requirements, privacy |

**Security as an enabler, not a blocker.**

---

# Success Criteria

## How We Know We've Succeeded

| Milestone | Target Date | Criteria |
|-----------|-------------|----------|
| M3 Code | End Q1 | 0 critical findings in production |
| M3 Infra | End Q2 | Cloud posture score > 90% |
| M4 Runtime | End Q2 | SIEM operational, MTTR < 7 days |
| M4 Complete | End Q3 | MITRE coverage > 80% |

---

# Resource Ask

## What We Need

| Resource | Purpose | Duration |
|----------|---------|----------|
| 1 Security Engineer | Pipeline integration | Q1-Q2 |
| Cloud security tool | Posture management | Annual license |
| SIEM/SOAR platform | Detection & response | Annual license |
| Training budget | Developer security | Ongoing |

---

# Discussion

## Questions?

**Decisions Needed:**

- [ ] Approve Security maturity roadmap
- [ ] Budget for cloud security tooling
- [ ] Confirm SIEM/SOAR timeline
- [ ] Assign project owners

---

<!-- _class: lead -->

# Appendix

---

# MITRE ATT&CK Coverage

## Detection Techniques

| Tactic | Techniques | Covered | Gap |
|--------|------------|---------|-----|
| Initial Access | 9 | 6 | Phishing detection |
| Execution | 12 | 8 | Script monitoring |
| Persistence | 19 | 10 | Registry, scheduled tasks |
| Privilege Escalation | 13 | 9 | Token manipulation |
| Defense Evasion | 42 | 20 | File masquerading |

---

# Vulnerability Severity

## SLA by Severity

| Severity | CVSS | Remediation SLA | Current Avg |
|----------|------|-----------------|-------------|
| Critical | 9.0-10.0 | 7 days | 14 days |
| High | 7.0-8.9 | 30 days | 45 days |
| Medium | 4.0-6.9 | 90 days | 60 days |
| Low | 0.1-3.9 | Best effort | N/A |

---

# Security Tooling Stack

## Current and Planned

| Category | Current | Planned |
|----------|---------|---------|
| SAST | Semgrep | - |
| SCA | Snyk | - |
| Secrets | GitLeaks | - |
| Cloud Security | Manual | CSPM (Q1) |
| DAST | Burp (manual) | Automated (Q2) |
| SIEM | ELK (basic) | Full SIEM (Q2) |
| WAF | CloudFlare | - |
