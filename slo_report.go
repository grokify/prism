package prism

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// SLOReport provides a structured view of all SLO requirements across maturity levels,
// organized by category. This enables visualization of how SLOs become more stringent
// as maturity increases.
type SLOReport struct {
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	GeneratedAt string             `json:"generatedAt,omitempty"`
	Categories  []SLOCategoryGroup `json:"categories"`
	// Flattened entries for tabular output
	Entries []SLOReportEntry `json:"entries,omitempty"`
}

// SLOCategoryGroup groups SLO requirements by metric category.
type SLOCategoryGroup struct {
	Category    string            `json:"category"`
	Description string            `json:"description,omitempty"`
	Metrics     []SLOMetricLadder `json:"metrics"`
}

// SLOMetricLadder shows how a single metric's SLO requirements escalate across maturity levels.
type SLOMetricLadder struct {
	MetricID   string             `json:"metricId"`
	MetricName string             `json:"metricName"`
	Unit       string             `json:"unit,omitempty"`
	Domain     string             `json:"domain,omitempty"`
	Stage      string             `json:"stage,omitempty"`
	Levels     []SLOMaturityEntry `json:"levels"`
}

// SLOMaturityEntry represents an SLO requirement at a specific maturity level.
type SLOMaturityEntry struct {
	Level       int     `json:"level"`
	LevelName   string  `json:"levelName"`
	GoalID      string  `json:"goalId,omitempty"`
	GoalName    string  `json:"goalName,omitempty"`
	Operator    string  `json:"operator"`
	Value       float64 `json:"value"`
	Description string  `json:"description,omitempty"`
}

// SLOReportEntry is a flattened row for tabular output (XLSX, CSV).
type SLOReportEntry struct {
	Category    string  `json:"category"`
	MetricID    string  `json:"metricId"`
	MetricName  string  `json:"metricName"`
	Domain      string  `json:"domain"`
	Stage       string  `json:"stage"`
	Unit        string  `json:"unit"`
	Level       int     `json:"level"`
	LevelName   string  `json:"levelName"`
	GoalID      string  `json:"goalId"`
	GoalName    string  `json:"goalName"`
	Operator    string  `json:"operator"`
	Value       float64 `json:"value"`
	Requirement string  `json:"requirement"` // Human-readable: ">=99.9%"
	Description string  `json:"description"`
}

// GenerateSLOReport builds an SLO report from a PRISM document by extracting
// metric criteria from goal maturity models, sorted by category then maturity level.
func (doc *PRISMDocument) GenerateSLOReport() *SLOReport {
	report := &SLOReport{
		Title:       "SLO Requirements by Maturity Level",
		Description: "Shows how SLO targets become more stringent as maturity increases",
		Categories:  []SLOCategoryGroup{},
		Entries:     []SLOReportEntry{},
	}

	if doc.Metadata != nil && doc.Metadata.Name != "" {
		report.Title = fmt.Sprintf("SLO Requirements: %s", doc.Metadata.Name)
	}

	// Build a map of metricID -> []SLOMaturityEntry (from all goals)
	metricEntries := make(map[string][]SLOMaturityEntry)
	metricInfo := make(map[string]*Metric)

	// Index metrics by ID
	for i := range doc.Metrics {
		m := &doc.Metrics[i]
		if m.ID != "" {
			metricInfo[m.ID] = m
		}
	}

	// Extract metric criteria from all goals' maturity models
	for _, goal := range doc.Goals {
		if goal.MaturityModel == nil {
			continue
		}
		for _, level := range goal.MaturityModel.Levels {
			for _, criterion := range level.MetricCriteria {
				entry := SLOMaturityEntry{
					Level:     level.Level,
					LevelName: MaturityLevelName(level.Level),
					GoalID:    goal.ID,
					GoalName:  goal.Name,
					Operator:  criterion.Operator,
					Value:     criterion.Value,
				}
				metricEntries[criterion.MetricID] = append(metricEntries[criterion.MetricID], entry)
			}
		}
	}

	// Group by category
	categoryMetrics := make(map[string][]SLOMetricLadder)

	for metricID, entries := range metricEntries {
		metric := metricInfo[metricID]
		if metric == nil {
			continue
		}

		// Sort entries by level
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Level < entries[j].Level
		})

		ladder := SLOMetricLadder{
			MetricID:   metricID,
			MetricName: metric.Name,
			Unit:       metric.Unit,
			Domain:     metric.Domain,
			Stage:      metric.Stage,
			Levels:     entries,
		}

		categoryMetrics[metric.Category] = append(categoryMetrics[metric.Category], ladder)
	}

	// Build category groups, sorted by category name
	categories := make([]string, 0, len(categoryMetrics))
	for cat := range categoryMetrics {
		categories = append(categories, cat)
	}
	sort.Strings(categories)

	for _, cat := range categories {
		ladders := categoryMetrics[cat]
		// Sort ladders by metric name within category
		sort.Slice(ladders, func(i, j int) bool {
			return ladders[i].MetricName < ladders[j].MetricName
		})

		group := SLOCategoryGroup{
			Category:    cat,
			Description: categoryDescription(cat),
			Metrics:     ladders,
		}
		report.Categories = append(report.Categories, group)
	}

	// Build flattened entries for tabular output
	report.Entries = report.flattenEntries()

	return report
}

// flattenEntries converts the hierarchical structure to flat rows for tabular output.
func (r *SLOReport) flattenEntries() []SLOReportEntry {
	var entries []SLOReportEntry

	for _, cat := range r.Categories {
		for _, ladder := range cat.Metrics {
			for _, level := range ladder.Levels {
				entry := SLOReportEntry{
					Category:    cat.Category,
					MetricID:    ladder.MetricID,
					MetricName:  ladder.MetricName,
					Domain:      ladder.Domain,
					Stage:       ladder.Stage,
					Unit:        ladder.Unit,
					Level:       level.Level,
					LevelName:   level.LevelName,
					GoalID:      level.GoalID,
					GoalName:    level.GoalName,
					Operator:    level.Operator,
					Value:       level.Value,
					Requirement: formatRequirement(level.Operator, level.Value, ladder.Unit),
					Description: level.Description,
				}
				entries = append(entries, entry)
			}
		}
	}

	return entries
}

// formatRequirement creates a human-readable requirement string.
func formatRequirement(operator string, value float64, unit string) string {
	opSymbol := operatorSymbol(operator)
	if unit == "%" {
		return fmt.Sprintf("%s%.2f%%", opSymbol, value)
	}
	if unit != "" {
		return fmt.Sprintf("%s%.2f %s", opSymbol, value, unit)
	}
	return fmt.Sprintf("%s%.2f", opSymbol, value)
}

// operatorSymbol converts operator constants to display symbols.
func operatorSymbol(op string) string {
	switch op {
	case SLOOperatorGTE:
		return ">="
	case SLOOperatorLTE:
		return "<="
	case SLOOperatorGT:
		return ">"
	case SLOOperatorLT:
		return "<"
	case SLOOperatorEQ:
		return "="
	default:
		return op
	}
}

// categoryDescription returns a description for a category.
func categoryDescription(cat string) string {
	switch cat {
	case CategoryPrevention:
		return "Proactive measures to prevent issues"
	case CategoryDetection:
		return "Capabilities to identify issues"
	case CategoryResponse:
		return "Ability to respond to incidents"
	case CategoryReliability:
		return "System availability and stability"
	case CategoryEfficiency:
		return "Operational efficiency metrics"
	case CategoryQuality:
		return "Code and service quality"
	default:
		return ""
	}
}

// TableColumns returns the column headers for tabular output.
func (r *SLOReport) TableColumns() []string {
	return []string{
		"Category",
		"Metric ID",
		"Metric Name",
		"Domain",
		"Stage",
		"Level",
		"Level Name",
		"Requirement",
		"Goal",
	}
}

// TableRows returns the data rows for tabular output.
func (r *SLOReport) TableRows() [][]string {
	rows := make([][]string, len(r.Entries))
	for i, e := range r.Entries {
		rows[i] = []string{
			e.Category,
			e.MetricID,
			e.MetricName,
			e.Domain,
			e.Stage,
			fmt.Sprintf("%d", e.Level),
			e.LevelName,
			e.Requirement,
			e.GoalName,
		}
	}
	return rows
}

// titleCase converts a string to title case.
func titleCase(s string) string {
	return cases.Title(language.English).String(s)
}

// ToMarkdown renders the report as Pandoc-compatible markdown.
func (r *SLOReport) ToMarkdown() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s\n\n", r.Title))
	if r.Description != "" {
		sb.WriteString(fmt.Sprintf("%s\n\n", r.Description))
	}

	for _, cat := range r.Categories {
		sb.WriteString(fmt.Sprintf("## %s\n\n", titleCase(cat.Category)))
		if cat.Description != "" {
			sb.WriteString(fmt.Sprintf("*%s*\n\n", cat.Description))
		}

		for _, ladder := range cat.Metrics {
			sb.WriteString(fmt.Sprintf("### %s\n\n", ladder.MetricName))
			sb.WriteString(fmt.Sprintf("**ID:** `%s` | **Domain:** %s | **Stage:** %s\n\n",
				ladder.MetricID, ladder.Domain, ladder.Stage))

			// Table header
			sb.WriteString("| Level | Name | Requirement | Goal |\n")
			sb.WriteString("|-------|------|-------------|------|\n")

			for _, level := range ladder.Levels {
				req := formatRequirement(level.Operator, level.Value, ladder.Unit)
				sb.WriteString(fmt.Sprintf("| %d | %s | %s | %s |\n",
					level.Level, level.LevelName, req, level.GoalName))
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// ToMarp renders the report as Marp presentation markdown.
func (r *SLOReport) ToMarp() string {
	var sb strings.Builder

	// Marp front matter
	sb.WriteString("---\n")
	sb.WriteString("marp: true\n")
	sb.WriteString("theme: default\n")
	sb.WriteString("paginate: true\n")
	sb.WriteString("---\n\n")

	// Title slide
	sb.WriteString(fmt.Sprintf("# %s\n\n", r.Title))
	if r.Description != "" {
		sb.WriteString(fmt.Sprintf("%s\n\n", r.Description))
	}
	sb.WriteString("---\n\n")

	// Overview slide
	sb.WriteString("# Categories Overview\n\n")
	for _, cat := range r.Categories {
		sb.WriteString(fmt.Sprintf("- **%s**: %d metrics\n", titleCase(cat.Category), len(cat.Metrics)))
	}
	sb.WriteString("\n---\n\n")

	// Category slides
	for _, cat := range r.Categories {
		sb.WriteString(fmt.Sprintf("# %s\n\n", titleCase(cat.Category)))
		if cat.Description != "" {
			sb.WriteString(fmt.Sprintf("*%s*\n\n", cat.Description))
		}

		for _, ladder := range cat.Metrics {
			sb.WriteString(fmt.Sprintf("**%s** (`%s`)\n\n", ladder.MetricName, ladder.MetricID))
		}
		sb.WriteString("\n---\n\n")

		// Detail slides for each metric
		for _, ladder := range cat.Metrics {
			sb.WriteString(fmt.Sprintf("## %s\n\n", ladder.MetricName))
			sb.WriteString(fmt.Sprintf("Domain: %s | Stage: %s\n\n", ladder.Domain, ladder.Stage))

			sb.WriteString("| Level | Requirement |\n")
			sb.WriteString("|-------|-------------|\n")

			for _, level := range ladder.Levels {
				req := formatRequirement(level.Operator, level.Value, ladder.Unit)
				sb.WriteString(fmt.Sprintf("| L%d %s | %s |\n", level.Level, level.LevelName, req))
			}
			sb.WriteString("\n---\n\n")
		}
	}

	return sb.String()
}

// ToMatrixMarkdown renders a matrix view showing all metrics across maturity levels.
func (r *SLOReport) ToMatrixMarkdown() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s - Matrix View\n\n", r.Title))

	for _, cat := range r.Categories {
		sb.WriteString(fmt.Sprintf("## %s\n\n", titleCase(cat.Category)))

		// Header with all 5 levels
		sb.WriteString("| Metric | L1 Reactive | L2 Basic | L3 Defined | L4 Managed | L5 Optimizing |\n")
		sb.WriteString("|--------|-------------|----------|------------|------------|---------------|\n")

		for _, ladder := range cat.Metrics {
			row := make([]string, 6)
			row[0] = ladder.MetricName

			// Initialize with "-" for levels without requirements
			for i := 1; i <= 5; i++ {
				row[i] = "-"
			}

			// Fill in actual requirements
			for _, level := range ladder.Levels {
				if level.Level >= 1 && level.Level <= 5 {
					row[level.Level] = formatRequirement(level.Operator, level.Value, ladder.Unit)
				}
			}

			sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
				row[0], row[1], row[2], row[3], row[4], row[5]))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
