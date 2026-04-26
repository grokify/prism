# Maturity Model Design: SLO-Backed Levels

## Problem

The current presentations conflate two different concepts:

| Concept | Example | Purpose |
|---------|---------|---------|
| **SLO/Metric** | "MTTR < 7 days" | Measurable outcome that defines level |
| **Requirement/Task** | "Implement SIEM/SOAR" | Work to achieve the SLO |

This creates confusion in presentations and makes it unclear what "achieving a level" actually means.

## Solution

Separate maturity definitions into:

1. **Level Criteria** - OpenSLO definitions that must be met
2. **Enablers** - Tasks/projects that help achieve the criteria

## Proposed Structure

### Naming Convention

Change from `L1-L5` to `M1-M5` where M = Maturity:

| Current | Proposed | Name |
|---------|----------|------|
| L1 | M1 | Reactive |
| L2 | M2 | Basic |
| L3 | M3 | Defined |
| L4 | M4 | Managed |
| M5 | M5 | Optimizing |

### Domain Maturity Model JSON

Each domain defines its own M1-M5 with OpenSLO-backed criteria:

```json
{
  "domainMaturityModels": [
    {
      "domain": "security",
      "description": "Security domain maturity model",
      "levels": [
        {
          "level": 1,
          "name": "Reactive",
          "description": "Ad-hoc security, firefighting mode",
          "criteria": [],
          "enablers": []
        },
        {
          "level": 2,
          "name": "Basic",
          "description": "Basic security controls in place",
          "criteria": [
            {
              "id": "sec-m2-sast",
              "name": "SAST Coverage",
              "slo": {
                "apiVersion": "openslo/v1",
                "kind": "SLO",
                "metadata": {
                  "name": "sast-coverage",
                  "displayName": "SAST Scanning Coverage"
                },
                "spec": {
                  "description": "Percentage of repositories with SAST scanning",
                  "service": "security-tooling",
                  "indicator": {
                    "spec": {
                      "ratioMetric": {
                        "good": { "metricSource": { "type": "Custom", "spec": { "query": "repos_with_sast" }}},
                        "total": { "metricSource": { "type": "Custom", "spec": { "query": "total_repos" }}}
                      }
                    }
                  },
                  "objectives": [
                    { "displayName": "Basic coverage", "target": 0.50 }
                  ],
                  "timeWindow": [{ "duration": "4w", "isRolling": true }]
                }
              }
            }
          ],
          "enablers": [
            {
              "id": "enable-sast-deploy",
              "name": "Deploy SAST tooling",
              "description": "Install and configure SAST scanner in CI pipeline",
              "type": "implementation",
              "effort": "2 weeks"
            }
          ]
        },
        {
          "level": 3,
          "name": "Defined",
          "description": "Integrated security with enforcement",
          "criteria": [
            {
              "id": "sec-m3-sast",
              "name": "SAST Coverage",
              "slo": {
                "spec": {
                  "objectives": [
                    { "displayName": "Full coverage", "target": 1.0 }
                  ]
                }
              }
            },
            {
              "id": "sec-m3-gates",
              "name": "Security Gates Active",
              "slo": {
                "spec": {
                  "description": "Critical findings block deployment",
                  "objectives": [
                    { "displayName": "Gates enforced", "target": 1.0 }
                  ]
                }
              }
            },
            {
              "id": "sec-m3-secrets",
              "name": "No Secrets in Code",
              "slo": {
                "spec": {
                  "description": "Zero secrets detected in repositories",
                  "objectives": [
                    { "displayName": "Zero secrets", "target": 1.0, "value": 0, "op": "eq" }
                  ]
                }
              }
            }
          ],
          "enablers": [
            {
              "id": "enable-security-gates",
              "name": "Implement security gates",
              "description": "Configure CI to block merges with critical/high findings"
            },
            {
              "id": "enable-secrets-scanning",
              "name": "Deploy secrets scanning",
              "description": "Install pre-commit hooks for secrets detection"
            }
          ]
        },
        {
          "level": 4,
          "name": "Managed",
          "description": "Real-time security with measurement",
          "criteria": [
            {
              "id": "sec-m4-mttr",
              "name": "Security MTTR",
              "slo": {
                "spec": {
                  "description": "Mean time to remediate critical security findings",
                  "indicator": {
                    "spec": {
                      "thresholdMetric": {
                        "metricSource": { "type": "Custom", "spec": { "query": "security_mttr_days" }}
                      }
                    }
                  },
                  "objectives": [
                    { "displayName": "Fast remediation", "target": 0.95, "value": 7, "op": "lte" }
                  ]
                }
              }
            },
            {
              "id": "sec-m4-detection",
              "name": "Threat Detection Coverage",
              "slo": {
                "spec": {
                  "description": "MITRE ATT&CK technique coverage",
                  "objectives": [
                    { "displayName": "Good coverage", "target": 0.70 }
                  ]
                }
              }
            },
            {
              "id": "sec-m4-realtime",
              "name": "Real-time Visibility",
              "slo": {
                "spec": {
                  "description": "Time from event to alert",
                  "objectives": [
                    { "displayName": "Near real-time", "target": 0.95, "value": 60, "op": "lte" }
                  ]
                }
              }
            }
          ],
          "enablers": [
            {
              "id": "enable-siem",
              "name": "Deploy SIEM/SOAR",
              "description": "Implement security information and event management"
            },
            {
              "id": "enable-mitre-mapping",
              "name": "Map detections to MITRE",
              "description": "Categorize detection rules by ATT&CK technique"
            }
          ]
        },
        {
          "level": 5,
          "name": "Optimizing",
          "description": "Proactive, automated security",
          "criteria": [
            {
              "id": "sec-m5-mttr",
              "name": "Security MTTR",
              "slo": {
                "spec": {
                  "objectives": [
                    { "displayName": "Rapid remediation", "target": 0.95, "value": 1, "op": "lte" }
                  ]
                }
              }
            },
            {
              "id": "sec-m5-auto-remediation",
              "name": "Auto-Remediation Rate",
              "slo": {
                "spec": {
                  "description": "Percentage of known issues auto-remediated",
                  "objectives": [
                    { "displayName": "High automation", "target": 0.80 }
                  ]
                }
              }
            },
            {
              "id": "sec-m5-detection",
              "name": "Threat Detection Coverage",
              "slo": {
                "spec": {
                  "objectives": [
                    { "displayName": "Excellent coverage", "target": 0.90 }
                  ]
                }
              }
            }
          ],
          "enablers": [
            {
              "id": "enable-auto-remediation",
              "name": "Implement auto-remediation",
              "description": "Automated patching and configuration fixes"
            },
            {
              "id": "enable-chaos-security",
              "name": "Security chaos engineering",
              "description": "Proactive testing of security controls"
            }
          ]
        }
      ]
    }
  ]
}
```

## Type Definitions

### DomainMaturityModel

```go
// DomainMaturityModel defines maturity levels for a specific domain.
type DomainMaturityModel struct {
    Domain      string           `json:"domain"`
    Description string           `json:"description,omitempty"`
    Levels      []MaturityLevel  `json:"levels"`
}

// MaturityLevel defines what M1-M5 means for a domain.
type MaturityLevel struct {
    Level       int                `json:"level"`       // 1-5
    Name        string             `json:"name"`        // Reactive, Basic, Defined, Managed, Optimizing
    Description string             `json:"description"`
    Criteria    []LevelCriterion   `json:"criteria"`    // SLOs that must be met
    Enablers    []LevelEnabler     `json:"enablers"`    // Tasks to achieve criteria
}

// LevelCriterion is an SLO that must be met for the level.
type LevelCriterion struct {
    ID          string      `json:"id"`
    Name        string      `json:"name"`
    Description string      `json:"description,omitempty"`
    SLO         *OpenSLO    `json:"slo"`           // OpenSLO definition
    Weight      float64     `json:"weight,omitempty"` // Relative importance (default 1.0)
    Required    bool        `json:"required,omitempty"` // Must pass (default true)
}

// LevelEnabler is a task/project that helps achieve criteria.
type LevelEnabler struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description,omitempty"`
    Type        string   `json:"type,omitempty"`      // implementation, process, training
    Effort      string   `json:"effort,omitempty"`    // Estimated effort
    CriteriaIDs []string `json:"criteriaIds,omitempty"` // Which criteria this enables
}
```

## Presentation Slide Template

With this separation, slides become clearer:

### Before (Confusing)

```markdown
# M3 → M4: Managed

**Requirements:**
- [ ] Threat detection with SIEM/SOAR        ← Implementation
- [ ] MTTR < 7 days for critical issues      ← Metric
- [ ] Real-time threat visibility            ← Ambiguous
- [ ] MITRE ATT&CK coverage > 70%            ← Metric

**Key Metrics:**
- Real-time threat visibility                ← Ambiguous
- Automated incident triage                  ← Implementation
- Mean time to detect < 1 hour               ← Metric
```

### After (Clear)

```markdown
# M3 → M4: Managed

**Level Criteria (SLOs):**

| SLO | Target | Current |
|-----|--------|---------|
| Security MTTR (critical) | ≤ 7 days | 14 days |
| MITRE ATT&CK Coverage | ≥ 70% | 55% |
| Time to Detection | ≤ 1 hour | 4 hours |
| Real-time Alert Latency | ≤ 60 sec | N/A |

**Enablers (Tasks):**

| Task | Status | Enables |
|------|--------|---------|
| Deploy SIEM/SOAR | Planned | Detection, MTTR |
| Map detections to MITRE | Not Started | Coverage |
| Configure real-time alerting | Not Started | Alert Latency |
```

## Level Achievement Calculation

```go
// IsLevelAchieved checks if all criteria for a level are met.
func (l *MaturityLevel) IsLevelAchieved(doc *PRISMDocument) bool {
    for _, criterion := range l.Criteria {
        if criterion.Required && !criterion.IsMet(doc) {
            return false
        }
    }
    return true
}

// LevelProgress returns the percentage of criteria met.
func (l *MaturityLevel) LevelProgress(doc *PRISMDocument) float64 {
    if len(l.Criteria) == 0 {
        return 1.0
    }

    var met, total float64
    for _, criterion := range l.Criteria {
        weight := criterion.Weight
        if weight == 0 {
            weight = 1.0
        }
        total += weight
        if criterion.IsMet(doc) {
            met += weight
        }
    }
    return met / total
}
```

## Benefits

1. **Clear Definition**: Maturity levels defined by measurable SLOs, not vague requirements
2. **Objective Assessment**: Can automatically calculate current level from metrics
3. **Actionable Roadmap**: Enablers become the project backlog
4. **OpenSLO Integration**: Standard format for SLO definitions via slogo
5. **Progress Tracking**: Clear percentage toward next level

## Migration Path

1. **Phase 1**: Add `DomainMaturityModel` type alongside existing `MaturityModel`
2. **Phase 2**: Create domain-specific maturity models for Security, Operations, Quality
3. **Phase 3**: Update presentations to use new structure
4. **Phase 4**: Add CLI commands: `prism maturity show`, `prism maturity progress`
5. **Phase 5**: Deprecate old `MaturityModel` structure

## Example: Complete Security Domain Model

See `/docs/schema/security-maturity.json` for a complete example.

## CLI Commands

```bash
# Show maturity model for a domain
prism maturity show prism.json --domain security

# Show progress toward next level
prism maturity progress prism.json --domain security

# List criteria not yet met
prism maturity gaps prism.json --domain security --level 4

# Show enablers for next level
prism maturity enablers prism.json --domain security
```

## Open Questions

1. Should enablers link to PRISM `initiatives` for tracking?
2. Should criteria support "stretch" vs "required" distinction?
3. How to handle criteria that span multiple layers?
4. Should we support "partial" level achievement (e.g., M3.5)?
