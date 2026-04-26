# Maturity Model Code Generation Architecture

## Overview

Define maturity models in Go structs → Generate JSON IR → Output to XLSX and Marp presentations.

```
┌─────────────────────────────────────────────────────────────────────────┐
│                         GO STRUCTS (Source of Truth)                    │
│                                                                         │
│  DomainMaturityModel → MaturityLevel → LevelCriterion (SLO)            │
│                                      → LevelEnabler (Requirement)       │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                            JSON IR                                      │
│                                                                         │
│  maturity-models/security.json                                          │
│  maturity-models/operations.json                                        │
│  maturity-models/quality.json                                           │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                    ┌───────────────┼───────────────┐
                    ▼               ▼               ▼
            ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
            │    XLSX     │  │    Marp     │  │   Reports   │
            │             │  │             │  │             │
            │ • Requirements│ │ • Tables    │  │ • Markdown  │
            │ • SLOs      │  │ • Charts    │  │ • HTML      │
            │ • Progress  │  │ • Narrative │  │ • PDF       │
            └─────────────┘  └─────────────┘  └─────────────┘
                                    │
                                    ▼
                           ┌─────────────────┐
                           │  LLM Enhancement│
                           │                 │
                           │ • Storytelling  │
                           │ • Context       │
                           │ • Audience fit  │
                           └─────────────────┘
```

## Go Types

### Core Types

```go
package prism

// DomainMaturityModel defines maturity levels for a domain.
type DomainMaturityModel struct {
    Domain      string          `json:"domain"`                // security, operations, quality
    Name        string          `json:"name"`                  // Display name
    Description string          `json:"description,omitempty"`
    Owner       string          `json:"owner,omitempty"`       // Accountable team/person
    Levels      []MaturityLevel `json:"levels"`
}

// MaturityLevel defines M1-M5 for a domain.
type MaturityLevel struct {
    Level       int              `json:"level"`       // 1-5
    Name        string           `json:"name"`        // Reactive, Basic, Defined, Managed, Optimizing
    Description string           `json:"description"`
    Criteria    []LevelCriterion `json:"criteria"`    // SLOs that define the level
    Enablers    []LevelEnabler   `json:"enablers"`    // Tasks to achieve the level
}

// LevelCriterion is a measurable SLO that defines level achievement.
type LevelCriterion struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description,omitempty"`

    // SLO Definition
    MetricName  string  `json:"metricName"`           // Human-readable metric
    Operator    string  `json:"operator"`             // gte, lte, gt, lt, eq
    Target      float64 `json:"target"`               // Target value
    Unit        string  `json:"unit,omitempty"`       // %, days, count, etc.

    // OpenSLO (optional, for full SLO definition)
    OpenSLO     *OpenSLOSpec `json:"openslo,omitempty"`

    // Classification
    Layer       string  `json:"layer,omitempty"`      // code, infra, runtime, etc.
    Category    string  `json:"category,omitempty"`   // prevention, detection, response

    // Assessment
    Current     float64 `json:"current,omitempty"`    // Current value
    IsMet       bool    `json:"isMet,omitempty"`      // Calculated: meets target?
    Weight      float64 `json:"weight,omitempty"`     // Relative importance (default 1.0)
    Required    bool    `json:"required"`             // Must pass for level (default true)
}

// LevelEnabler is implementation work to achieve criteria.
type LevelEnabler struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description,omitempty"`

    // Classification
    Type        string   `json:"type,omitempty"`      // implementation, process, training, tooling
    Layer       string   `json:"layer,omitempty"`     // code, infra, runtime, etc.

    // Effort
    Effort      string   `json:"effort,omitempty"`    // T-shirt size or duration
    Team        string   `json:"team,omitempty"`      // Responsible team

    // Tracking
    Status      string   `json:"status,omitempty"`    // not_started, in_progress, completed
    CriteriaIDs []string `json:"criteriaIds,omitempty"` // Which SLOs this enables

    // Dependencies
    DependsOn   []string `json:"dependsOn,omitempty"` // Other enabler IDs
}

// OpenSLOSpec is an embedded OpenSLO definition (simplified).
type OpenSLOSpec struct {
    Service     string            `json:"service,omitempty"`
    Indicator   *SLOIndicator     `json:"indicator,omitempty"`
    Objectives  []SLOObjective    `json:"objectives,omitempty"`
    TimeWindow  string            `json:"timeWindow,omitempty"`  // e.g., "30d"
    Budgeting   string            `json:"budgeting,omitempty"`   // Timeslices, Occurrences
}
```

### Assessment Types

```go
// MaturityAssessment captures current state against a maturity model.
type MaturityAssessment struct {
    Domain        string                   `json:"domain"`
    AssessedAt    time.Time                `json:"assessedAt"`
    AssessedBy    string                   `json:"assessedBy,omitempty"`
    CurrentLevel  int                      `json:"currentLevel"`  // Achieved level (1-5)
    TargetLevel   int                      `json:"targetLevel"`   // Goal level
    LevelProgress map[int]LevelProgress    `json:"levelProgress"` // Progress per level
}

// LevelProgress tracks progress toward a maturity level.
type LevelProgress struct {
    Level           int     `json:"level"`
    CriteriaMet     int     `json:"criteriaMet"`
    CriteriaTotal   int     `json:"criteriaTotal"`
    ProgressPercent float64 `json:"progressPercent"`
    EnablersDone    int     `json:"enablersDone"`
    EnablersTotal   int     `json:"enablersTotal"`
}
```

## XLSX Output Structure

### Sheet 1: Requirements (Enablers)

| Column | Description |
|--------|-------------|
| ID | Enabler ID |
| Domain | security, operations, quality |
| Level | M1, M2, M3, M4, M5 |
| Name | Enabler name |
| Description | What needs to be done |
| Type | implementation, process, training, tooling |
| Layer | code, infra, runtime, etc. |
| Team | Responsible team |
| Effort | T-shirt size or duration |
| Status | not_started, in_progress, completed |
| Enables | Comma-separated criteria IDs |
| Depends On | Comma-separated enabler IDs |

### Sheet 2: SLOs (Criteria)

| Column | Description |
|--------|-------------|
| ID | Criterion ID |
| Domain | security, operations, quality |
| Level | M1, M2, M3, M4, M5 |
| Name | SLO name |
| Metric | What is measured |
| Operator | gte, lte, eq, etc. |
| Target | Target value |
| Unit | %, days, count, etc. |
| Current | Current value |
| Met | Yes/No |
| Layer | code, infra, runtime, etc. |
| Category | prevention, detection, response |
| Required | Yes/No |
| Weight | Relative importance |

### Sheet 3: Progress Summary

| Column | Description |
|--------|-------------|
| Domain | security, operations, quality |
| Current Level | M1-M5 |
| Target Level | M1-M5 |
| M2 Progress | % criteria met |
| M3 Progress | % criteria met |
| M4 Progress | % criteria met |
| M5 Progress | % criteria met |
| Next Actions | Key enablers needed |

### Sheet 4: Level Definitions

| Column | Description |
|--------|-------------|
| Level | M1, M2, M3, M4, M5 |
| Name | Reactive, Basic, etc. |
| Security | Description for security |
| Operations | Description for operations |
| Quality | Description for quality |

## Marp Generation

### Deterministic Content (Go Code)

Generated programmatically from JSON IR:

```go
// GenerateLevelSlide creates a Marp slide for a level transition.
func GenerateLevelSlide(from, to int, level MaturityLevel) string {
    var sb strings.Builder

    sb.WriteString(fmt.Sprintf("# M%d → M%d: %s\n\n", from, to, level.Name))

    // Criteria table (SLOs)
    sb.WriteString("## Level Criteria (SLOs)\n\n")
    sb.WriteString("| SLO | Target | Current | Status |\n")
    sb.WriteString("|-----|--------|---------|--------|\n")
    for _, c := range level.Criteria {
        status := "Not Met"
        if c.IsMet {
            status = "Met"
        }
        sb.WriteString(fmt.Sprintf("| %s | %s %v%s | %v%s | %s |\n",
            c.Name, operatorSymbol(c.Operator), c.Target, c.Unit,
            c.Current, c.Unit, status))
    }

    // Enablers table (Tasks)
    sb.WriteString("\n## Enablers (Tasks)\n\n")
    sb.WriteString("| Task | Type | Status | Enables |\n")
    sb.WriteString("|------|------|--------|--------|\n")
    for _, e := range level.Enablers {
        sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
            e.Name, e.Type, e.Status, strings.Join(e.CriteriaIDs, ", ")))
    }

    sb.WriteString("\n---\n\n")
    return sb.String()
}
```

### LLM-Enhanced Content

Templates with placeholders for LLM enhancement:

```go
// SlideTemplate for LLM enhancement.
type SlideTemplate struct {
    Type        string            // title, overview, level, project, roadmap
    Data        map[string]any    // Structured data from JSON IR
    Placeholders []string         // Fields for LLM to fill
}

// Example template
template := SlideTemplate{
    Type: "domain_overview",
    Data: map[string]any{
        "domain": "security",
        "currentLevel": 2,
        "targetLevel": 4,
        "criteriaGaps": []string{"MTTR", "Detection Coverage"},
    },
    Placeholders: []string{
        "executive_summary",      // 2-3 sentence overview
        "key_challenge",          // Main obstacle
        "business_impact",        // Why this matters
    },
}
```

### Generation Pipeline

```go
// GeneratePresentation creates a Marp deck from a maturity model.
func GeneratePresentation(model DomainMaturityModel, assessment MaturityAssessment, opts GenerateOpts) (string, error) {
    var slides []string

    // 1. Title slide (deterministic)
    slides = append(slides, generateTitleSlide(model))

    // 2. Overview slide (LLM-enhanced)
    if opts.UseLLM {
        overview, err := enhanceWithLLM(generateOverviewTemplate(model, assessment))
        if err != nil {
            return "", err
        }
        slides = append(slides, overview)
    } else {
        slides = append(slides, generateOverviewSlide(model, assessment))
    }

    // 3. Current state (deterministic)
    slides = append(slides, generateCurrentStateSlide(assessment))

    // 4. Level transition slides (deterministic tables + LLM narrative)
    for level := assessment.CurrentLevel + 1; level <= assessment.TargetLevel; level++ {
        levelDef := model.Levels[level-1]
        slide := generateLevelSlide(level-1, level, levelDef)
        if opts.UseLLM {
            slide, _ = addLLMNarrative(slide, levelDef)
        }
        slides = append(slides, slide)
    }

    // 5. Project prioritization (deterministic)
    slides = append(slides, generateProjectSlide(model, assessment))

    // 6. Roadmap (deterministic)
    slides = append(slides, generateRoadmapSlide(model, assessment))

    // 7. Summary (LLM-enhanced)
    if opts.UseLLM {
        summary, _ := enhanceWithLLM(generateSummaryTemplate(model, assessment))
        slides = append(slides, summary)
    }

    return assembleMarpDeck(slides, opts.Theme), nil
}
```

## CLI Commands

```bash
# Generate JSON IR from Go definitions
prism maturity export --domain security -o maturity-models/security.json

# Generate XLSX from JSON IR
prism maturity xlsx maturity-models/*.json -o maturity-report.xlsx

# Generate Marp slides (deterministic only)
prism maturity slides maturity-models/security.json -o security-slides.md

# Generate Marp slides with LLM enhancement
prism maturity slides maturity-models/security.json \
  --llm \
  --audience "security leadership" \
  -o security-slides.md

# Generate all outputs
prism maturity generate maturity-models/*.json \
  --xlsx maturity-report.xlsx \
  --slides presentations/ \
  --llm
```

## File Structure

```
prism/
├── maturity/
│   ├── model.go           # Core types
│   ├── security.go        # Security domain model (Go source of truth)
│   ├── operations.go      # Operations domain model
│   ├── quality.go         # Quality domain model
│   ├── product.go         # Product domain model (requirements + adoption)
│   ├── assessment.go      # Assessment types and calculations
│   ├── xlsx.go            # XLSX generation
│   ├── marp.go            # Marp generation (deterministic)
│   └── llm.go             # LLM enhancement
├── maturity-models/       # Generated JSON IR
│   ├── security.json
│   ├── operations.json
│   ├── quality.json
│   └── product.json
└── cmd/prism/
    └── maturity.go        # CLI commands
```

## Example: Security Domain in Go

```go
package maturity

// SecurityMaturityModel returns the security domain maturity model.
func SecurityMaturityModel() DomainMaturityModel {
    return DomainMaturityModel{
        Domain:      "security",
        Name:        "Security",
        Description: "Application and infrastructure security maturity",
        Owner:       "Security Team",
        Levels: []MaturityLevel{
            {
                Level:       1,
                Name:        "Reactive",
                Description: "Ad-hoc security, firefighting mode",
                Criteria:    []LevelCriterion{}, // No criteria for M1
                Enablers:    []LevelEnabler{},
            },
            {
                Level:       2,
                Name:        "Basic",
                Description: "Basic security controls in place",
                Criteria: []LevelCriterion{
                    {
                        ID:         "sec-m2-sast-coverage",
                        Name:       "SAST Coverage",
                        MetricName: "Repositories with SAST scanning",
                        Operator:   "gte",
                        Target:     50,
                        Unit:       "%",
                        Layer:      "code",
                        Category:   "prevention",
                        Required:   true,
                    },
                    {
                        ID:         "sec-m2-vulns-tracked",
                        Name:       "Vulnerabilities Tracked",
                        MetricName: "Vulnerabilities logged in tracking system",
                        Operator:   "eq",
                        Target:     100,
                        Unit:       "%",
                        Layer:      "code",
                        Category:   "detection",
                        Required:   true,
                    },
                },
                Enablers: []LevelEnabler{
                    {
                        ID:          "sec-m2-deploy-sast",
                        Name:        "Deploy SAST tooling",
                        Description: "Install and configure SAST scanner",
                        Type:        "tooling",
                        Layer:       "code",
                        Effort:      "2 weeks",
                        Team:        "Security",
                        CriteriaIDs: []string{"sec-m2-sast-coverage"},
                    },
                    {
                        ID:          "sec-m2-vuln-tracking",
                        Name:        "Implement vulnerability tracking",
                        Description: "Set up vulnerability tracking workflow",
                        Type:        "process",
                        Layer:       "code",
                        Effort:      "1 week",
                        Team:        "Security",
                        CriteriaIDs: []string{"sec-m2-vulns-tracked"},
                    },
                },
            },
            {
                Level:       3,
                Name:        "Defined",
                Description: "Integrated security with enforcement",
                Criteria: []LevelCriterion{
                    {
                        ID:         "sec-m3-sast-coverage",
                        Name:       "SAST Coverage",
                        MetricName: "Repositories with SAST in CI",
                        Operator:   "eq",
                        Target:     100,
                        Unit:       "%",
                        Layer:      "code",
                        Category:   "prevention",
                        Required:   true,
                    },
                    {
                        ID:         "sec-m3-gates-active",
                        Name:       "Security Gates Active",
                        MetricName: "Pipelines with security gates",
                        Operator:   "eq",
                        Target:     100,
                        Unit:       "%",
                        Layer:      "code",
                        Category:   "prevention",
                        Required:   true,
                    },
                    {
                        ID:         "sec-m3-secrets-zero",
                        Name:       "No Secrets in Code",
                        MetricName: "Secrets detected in repositories",
                        Operator:   "eq",
                        Target:     0,
                        Unit:       "count",
                        Layer:      "code",
                        Category:   "prevention",
                        Required:   true,
                    },
                    {
                        ID:         "sec-m3-critical-findings",
                        Name:       "Zero Critical Findings",
                        MetricName: "Critical/high findings in production",
                        Operator:   "eq",
                        Target:     0,
                        Unit:       "count",
                        Layer:      "runtime",
                        Category:   "prevention",
                        Required:   true,
                    },
                },
                Enablers: []LevelEnabler{
                    {
                        ID:          "sec-m3-ci-integration",
                        Name:        "Integrate SAST in CI/CD",
                        Description: "Add SAST scanning to all pipelines",
                        Type:        "implementation",
                        Layer:       "code",
                        Effort:      "3 weeks",
                        Team:        "Platform",
                        CriteriaIDs: []string{"sec-m3-sast-coverage"},
                        DependsOn:   []string{"sec-m2-deploy-sast"},
                    },
                    {
                        ID:          "sec-m3-security-gates",
                        Name:        "Implement security gates",
                        Description: "Block merges with critical findings",
                        Type:        "implementation",
                        Layer:       "code",
                        Effort:      "2 weeks",
                        Team:        "Platform",
                        CriteriaIDs: []string{"sec-m3-gates-active", "sec-m3-critical-findings"},
                        DependsOn:   []string{"sec-m3-ci-integration"},
                    },
                    {
                        ID:          "sec-m3-secrets-scanning",
                        Name:        "Deploy secrets scanning",
                        Description: "Pre-commit hooks for secrets detection",
                        Type:        "tooling",
                        Layer:       "code",
                        Effort:      "1 week",
                        Team:        "Security",
                        CriteriaIDs: []string{"sec-m3-secrets-zero"},
                    },
                },
            },
            {
                Level:       4,
                Name:        "Managed",
                Description: "Real-time security with measurement",
                Criteria: []LevelCriterion{
                    {
                        ID:         "sec-m4-mttr",
                        Name:       "Security MTTR",
                        MetricName: "Mean time to remediate critical findings",
                        Operator:   "lte",
                        Target:     7,
                        Unit:       "days",
                        Layer:      "runtime",
                        Category:   "response",
                        Required:   true,
                    },
                    {
                        ID:         "sec-m4-detection-coverage",
                        Name:       "Threat Detection Coverage",
                        MetricName: "MITRE ATT&CK technique coverage",
                        Operator:   "gte",
                        Target:     70,
                        Unit:       "%",
                        Layer:      "runtime",
                        Category:   "detection",
                        Required:   true,
                    },
                    {
                        ID:         "sec-m4-detection-latency",
                        Name:       "Detection Latency",
                        MetricName: "Time from event to alert",
                        Operator:   "lte",
                        Target:     60,
                        Unit:       "seconds",
                        Layer:      "runtime",
                        Category:   "detection",
                        Required:   true,
                    },
                },
                Enablers: []LevelEnabler{
                    {
                        ID:          "sec-m4-siem",
                        Name:        "Deploy SIEM/SOAR",
                        Description: "Security information and event management",
                        Type:        "tooling",
                        Layer:       "runtime",
                        Effort:      "3 months",
                        Team:        "Security",
                        CriteriaIDs: []string{"sec-m4-detection-coverage", "sec-m4-detection-latency"},
                    },
                    {
                        ID:          "sec-m4-mitre-mapping",
                        Name:        "Map detections to MITRE",
                        Description: "Categorize detection rules by ATT&CK technique",
                        Type:        "process",
                        Layer:       "runtime",
                        Effort:      "4 weeks",
                        Team:        "Security",
                        CriteriaIDs: []string{"sec-m4-detection-coverage"},
                        DependsOn:   []string{"sec-m4-siem"},
                    },
                    {
                        ID:          "sec-m4-remediation-sla",
                        Name:        "Implement remediation SLAs",
                        Description: "Define and enforce fix timelines",
                        Type:        "process",
                        Layer:       "runtime",
                        Effort:      "2 weeks",
                        Team:        "Security",
                        CriteriaIDs: []string{"sec-m4-mttr"},
                    },
                },
            },
            {
                Level:       5,
                Name:        "Optimizing",
                Description: "Proactive, automated security",
                Criteria: []LevelCriterion{
                    {
                        ID:         "sec-m5-mttr",
                        Name:       "Security MTTR",
                        MetricName: "Mean time to remediate critical findings",
                        Operator:   "lte",
                        Target:     1,
                        Unit:       "days",
                        Layer:      "runtime",
                        Category:   "response",
                        Required:   true,
                    },
                    {
                        ID:         "sec-m5-auto-remediation",
                        Name:       "Auto-Remediation Rate",
                        MetricName: "Known issues auto-remediated",
                        Operator:   "gte",
                        Target:     80,
                        Unit:       "%",
                        Layer:      "runtime",
                        Category:   "response",
                        Required:   true,
                    },
                    {
                        ID:         "sec-m5-detection-coverage",
                        Name:       "Threat Detection Coverage",
                        MetricName: "MITRE ATT&CK technique coverage",
                        Operator:   "gte",
                        Target:     90,
                        Unit:       "%",
                        Layer:      "runtime",
                        Category:   "detection",
                        Required:   true,
                    },
                },
                Enablers: []LevelEnabler{
                    {
                        ID:          "sec-m5-auto-remediation",
                        Name:        "Implement auto-remediation",
                        Description: "Automated patching and configuration fixes",
                        Type:        "implementation",
                        Layer:       "runtime",
                        Effort:      "3 months",
                        Team:        "Security + Platform",
                        CriteriaIDs: []string{"sec-m5-auto-remediation", "sec-m5-mttr"},
                    },
                    {
                        ID:          "sec-m5-chaos-security",
                        Name:        "Security chaos engineering",
                        Description: "Proactive testing of security controls",
                        Type:        "process",
                        Layer:       "runtime",
                        Effort:      "Ongoing",
                        Team:        "Security",
                        CriteriaIDs: []string{"sec-m5-detection-coverage"},
                    },
                },
            },
        },
    }
}
```

## Benefits

1. **Single Source of Truth**: Go code defines everything
2. **Type Safety**: Compile-time validation of model structure
3. **Multiple Outputs**: XLSX for collaboration, Marp for presentations
4. **Deterministic + LLM**: Tables/metrics generated exactly, narrative enhanced
5. **Auditable**: JSON IR can be version controlled
6. **Extensible**: Add new output formats easily

## Next Steps

1. [ ] Implement core types in `maturity/model.go`
2. [ ] Implement Security domain model in `maturity/security.go`
3. [ ] Implement XLSX generator in `maturity/xlsx.go`
4. [ ] Implement Marp generator in `maturity/marp.go`
5. [ ] Add CLI commands in `cmd/prism/maturity.go`
6. [ ] Add LLM enhancement in `maturity/llm.go`
