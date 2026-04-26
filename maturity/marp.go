package maturity

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// MarpGenerator generates Marp presentations from maturity specs.
type MarpGenerator struct {
	spec *Spec
}

// NewMarpGenerator creates a new Marp generator.
func NewMarpGenerator(spec *Spec) *MarpGenerator {
	return &MarpGenerator{spec: spec}
}

// Generate creates the Marp presentation content.
func (g *MarpGenerator) Generate() (string, error) {
	var sb strings.Builder

	// Write frontmatter
	sb.WriteString(marpFrontmatter(g.spec.Metadata.Name))

	// Executive Overview
	sb.WriteString(g.execOverviewSlides())

	// Domain sections
	domainOrder := []string{"security", "operational-excellence", "quality", "product", "ai"}
	for _, domainKey := range domainOrder {
		domain, ok := g.spec.Domains[domainKey]
		if !ok {
			continue
		}
		sb.WriteString(g.domainSlides(domainKey, domain))
	}

	// Summary and appendix
	sb.WriteString(g.summarySlides())
	sb.WriteString(g.appendixSlides())

	return sb.String(), nil
}

// SaveAs saves the presentation to a file.
func (g *MarpGenerator) SaveAs(filename string) error {
	content, err := g.Generate()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, []byte(content), 0600)
}

func marpFrontmatter(title string) string {
	return fmt.Sprintf(`---
marp: true
theme: default
paginate: true
header: '%s'
footer: 'Confidential'
---

<!-- _class: lead -->

# %s

## A Unified Framework for B2B SaaS Health Metrics

**Security | Operational Excellence | Quality | Product | AI**

---

`, title, title)
}

func (g *MarpGenerator) execOverviewSlides() string {
	var sb strings.Builder

	sb.WriteString(`<!-- _class: lead -->

# Executive Overview

---

# The Challenge

## Fragmented Metrics Across the Organization

- **Security** tracks vulnerabilities, compliance, threat detection
- **Engineering** tracks DORA metrics, SLOs, incidents
- **Quality** tracks coverage, defects, test results
- **Product** tracks adoption, activation, churn
- **AI** tracks adoption, productivity, governance

**Result:** No unified view of organizational health

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

# Current State Summary

## Where We Are Today

`)

	// Build current state table from assessments
	sb.WriteString("| Domain | Current | Target | Gap |\n")
	sb.WriteString("|--------|---------|--------|-----|\n")

	domainOrder := []string{"security", "operational-excellence", "quality", "product", "ai"}
	for _, domainKey := range domainOrder {
		assessment, ok := g.spec.Assessments[domainKey]
		if !ok {
			continue
		}
		domain := g.spec.Domains[domainKey]
		gap := assessment.TargetLevel - assessment.CurrentLevel
		sb.WriteString(fmt.Sprintf("| **%s** | M%d | M%d | %d levels |\n",
			domain.Name, assessment.CurrentLevel, assessment.TargetLevel, gap))
	}

	sb.WriteString("\n---\n\n")
	return sb.String()
}

func (g *MarpGenerator) domainSlides(domainKey string, domain *DomainModel) string {
	var sb strings.Builder

	// Domain title slide
	sb.WriteString(fmt.Sprintf(`<!-- _class: lead -->

# %s Domain

**%s**

---

`, domain.Name, domain.Description))

	// KPI Thresholds table
	thresholds := g.spec.KPIThresholds[domainKey]
	if len(thresholds) > 0 {
		sb.WriteString(fmt.Sprintf("# %s KPIs\n\n## Thresholds by Maturity Level\n\n", domain.Name))
		sb.WriteString("| Metric | M2 | M3 | M4 | M5 |\n")
		sb.WriteString("|--------|-----|-----|-----|-----|\n")

		for _, kpi := range thresholds {
			sb.WriteString(fmt.Sprintf("| **%s** | %s | %s | %s | %s |\n",
				kpi.Name,
				formatThreshold(kpi.Thresholds.M2, kpi.Unit),
				formatThreshold(kpi.Thresholds.M3, kpi.Unit),
				formatThreshold(kpi.Thresholds.M4, kpi.Unit),
				formatThreshold(kpi.Thresholds.M5, kpi.Unit),
			))
		}
		sb.WriteString("\n---\n\n")
	}

	// Current state with level assessment
	if len(thresholds) > 0 {
		sb.WriteString(fmt.Sprintf("# %s Current State\n\n## Assessment Summary\n\n", domain.Name))
		sb.WriteString("| KPI | Current | Level | M3 Target | M4 Target |\n")
		sb.WriteString("|-----|---------|-------|-----------|------------|\n")

		for _, kpi := range thresholds {
			level := determineLevelFromKPI(kpi)
			sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				kpi.Name,
				formatCurrentAny(kpi.Current, kpi.Unit),
				level,
				formatThreshold(kpi.Thresholds.M3, kpi.Unit),
				formatThreshold(kpi.Thresholds.M4, kpi.Unit),
			))
		}
		sb.WriteString("\n---\n\n")
	}

	// Level definitions
	sb.WriteString(fmt.Sprintf("# %s Maturity Levels\n\n## What Each Level Means\n\n", domain.Name))
	sb.WriteString("| Level | Name | Description |\n")
	sb.WriteString("|-------|------|-------------|\n")

	for _, level := range domain.Levels {
		sb.WriteString(fmt.Sprintf("| **M%d** | %s | %s |\n",
			level.Level, level.Name, truncate(level.Description, 50)))
	}
	sb.WriteString("\n---\n\n")

	// Enablers/Roadmap
	sb.WriteString(fmt.Sprintf("# %s Roadmap\n\n## Key Initiatives\n\n", domain.Name))
	sb.WriteString("| Project | Impact | Status |\n")
	sb.WriteString("|---------|--------|--------|\n")

	assessment := g.spec.Assessments[domainKey]
	enablers := g.collectEnablers(domain, assessment)
	for i, e := range enablers {
		if i >= 5 {
			break
		}
		sb.WriteString(fmt.Sprintf("| %s | M%d %s | %s |\n",
			e.Name, e.Level, e.Layer, formatStatus(e.Status)))
	}
	sb.WriteString("\n---\n\n")

	return sb.String()
}

type enablerWithLevel struct {
	Enabler
	Level int
}

func (g *MarpGenerator) collectEnablers(domain *DomainModel, assessment *DomainAssessment) []enablerWithLevel {
	var enablers []enablerWithLevel

	for _, level := range domain.Levels {
		for _, e := range level.Enablers {
			status := e.Status
			if assessment != nil && assessment.EnablerStatus != nil {
				if s, ok := assessment.EnablerStatus[e.ID]; ok {
					status = s
				}
			}
			enablers = append(enablers, enablerWithLevel{
				Enabler: Enabler{
					ID:          e.ID,
					Name:        e.Name,
					Description: e.Description,
					Type:        e.Type,
					Layer:       e.Layer,
					Status:      status,
				},
				Level: level.Level,
			})
		}
	}

	// Sort by status (in_progress first, then not_started, then completed)
	sort.Slice(enablers, func(i, j int) bool {
		statusOrder := map[string]int{
			StatusInProgress: 0,
			StatusNotStarted: 1,
			StatusCompleted:  2,
			StatusBlocked:    3,
		}
		return statusOrder[enablers[i].Status] < statusOrder[enablers[j].Status]
	})

	return enablers
}

func (g *MarpGenerator) summarySlides() string {
	return `<!-- _class: lead -->

# Summary and Next Steps

---

# Cross-Domain Summary

## Maturity Progression Plan

| Domain | Current | Q2 Target | Q4 Target |
|--------|---------|-----------|-----------|
| **Security** | M2 | M3 | M4 |
| **Operational Excellence** | M3 | M4 | M4 |
| **Quality** | M2 | M3 | M4 |
| **Product** | M2 | M3 | M4 |
| **AI** | M2 | M3 | M4 |

---

# Next Steps

## Immediate Actions

1. **Approve** maturity model as the unified metrics framework
2. **Conduct** baseline maturity assessment per domain
3. **Assign** domain overlay owners
4. **Kickoff** Q1 initiatives

---

`
}

func (g *MarpGenerator) appendixSlides() string {
	return `<!-- _class: lead -->

# Appendix

---

# Framework Mappings

## Industry Alignment

| Framework | Mapping |
|-----------|---------|
| **DORA** | Operational Excellence metrics |
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
| **AIOps** | AI for IT Operations |
| **CSPM** | Cloud Security Posture Management |
`
}

func formatThreshold(val any, unit string) string {
	if val == nil {
		return "-"
	}

	switch v := val.(type) {
	case float64:
		if unit == "%" {
			return fmt.Sprintf("≥%.0f%%", v)
		}
		if unit == "days" || unit == "time" {
			return fmt.Sprintf("≤%.0f", v)
		}
		if unit == "count" || unit == "per KLOC" {
			return fmt.Sprintf("≤%.1f", v)
		}
		return fmt.Sprintf("%.0f", v)
	case string:
		return v
	case bool:
		if v {
			return "Yes"
		}
		return "No"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func formatCurrentAny(val any, unit string) string {
	if val == nil {
		return "-"
	}

	switch v := val.(type) {
	case float64:
		if unit == "%" {
			return fmt.Sprintf("%.0f%%", v)
		}
		if unit == "days" {
			return fmt.Sprintf("%.0f days", v)
		}
		return fmt.Sprintf("%.1f", v)
	case bool:
		if v {
			return "Yes"
		}
		return "No"
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func formatStatus(status string) string {
	switch status {
	case StatusCompleted:
		return "Done"
	case StatusInProgress:
		return "In Progress"
	case StatusBlocked:
		return "Blocked"
	default:
		return "Not Started"
	}
}

func determineLevelFromKPI(kpi KPIThreshold) string {
	if kpi.Current == nil {
		return "M1"
	}

	// Handle boolean current values
	if currentBool, ok := kpi.Current.(bool); ok {
		if m5, ok := kpi.Thresholds.M5.(bool); ok && m5 && currentBool {
			return "M5"
		}
		if m4, ok := kpi.Thresholds.M4.(bool); ok && m4 && currentBool {
			return "M4"
		}
		if m3, ok := kpi.Thresholds.M3.(bool); ok && m3 && currentBool {
			return "M3"
		}
		return "M1"
	}

	// Handle numeric current values
	currentVal, ok := toFloat64(kpi.Current)
	if !ok {
		// String current values (like "2x/week") - can't auto-determine level
		return "~M3" // Approximate
	}

	isLowerBetter := kpi.Operator == "lte"

	// Check each level from highest to lowest
	levels := []struct {
		name      string
		threshold any
	}{
		{"M5", kpi.Thresholds.M5},
		{"M4", kpi.Thresholds.M4},
		{"M3", kpi.Thresholds.M3},
		{"M2", kpi.Thresholds.M2},
	}

	for _, level := range levels {
		if level.threshold == nil {
			continue
		}
		if _, isString := level.threshold.(string); isString {
			continue // Skip string thresholds like "tracked"
		}

		threshold, ok := toFloat64(level.threshold)
		if !ok {
			continue
		}

		if isLowerBetter {
			if currentVal <= threshold {
				return level.name
			}
		} else {
			if currentVal >= threshold {
				return level.name
			}
		}
	}

	return "M1"
}

func toFloat64(v any) (float64, bool) {
	switch val := v.(type) {
	case float64:
		return val, true
	case int:
		return float64(val), true
	case int64:
		return float64(val), true
	default:
		return 0, false
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// GenerateMarp is a convenience function to generate Marp from a spec file.
func GenerateMarp(specFile, outputFile string) error {
	spec, err := ReadSpecFile(specFile)
	if err != nil {
		return err
	}

	gen := NewMarpGenerator(spec)
	return gen.SaveAs(outputFile)
}
