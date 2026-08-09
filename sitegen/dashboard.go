package sitegen

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"

	"github.com/grokify/prism/ecosystem"
)

// DashboardConfig holds configuration for dashboard generation.
type DashboardConfig struct {
	// ScoringConfig for initiative prioritization.
	ScoringConfig ecosystem.ScoringConfig

	// Contributions maps initiative IDs to pillar contributions.
	Contributions ecosystem.InitiativePillars

	// InitiativeNames maps initiative IDs to display names.
	InitiativeNames map[string]string

	// TopN is the number of top initiatives to show.
	TopN int

	// IncludeMaturityHeatmap enables the maturity heatmap section.
	IncludeMaturityHeatmap bool

	// IncludeGapAnalysis enables the gap analysis section.
	IncludeGapAnalysis bool
}

// DashboardData holds all data for the unified dashboard.
type DashboardData struct {
	// Pillar investment mix
	PillarRollup  ecosystem.PortfolioRollup `json:"pillarRollup"`
	PillarDetails []ecosystem.PillarDetail  `json:"pillarDetails"`

	// Top initiatives with scores
	TopInitiatives []RankedInitiative `json:"topInitiatives"`

	// Maturity heatmap data (domain -> status -> count)
	MaturityHeatmap []MaturityHeatmapRow `json:"maturityHeatmap,omitempty"`

	// Gap analysis summary
	GapSummary *GapSummary `json:"gapSummary,omitempty"`

	// Stats
	TotalCapabilities int `json:"totalCapabilities"`
	TotalInitiatives  int `json:"totalInitiatives"`
	TotalMetrics      int `json:"totalMetrics"`
}

// RankedInitiative represents a scored and ranked initiative.
type RankedInitiative struct {
	Rank            int                            `json:"rank"`
	ID              string                         `json:"id"`
	Name            string                         `json:"name"`
	Score           float64                        `json:"score"`
	Pillars         []ecosystem.PillarContribution `json:"pillars,omitempty"`
	EvidenceCount   int                            `json:"evidenceCount"`
	DimensionScores []ecosystem.DimensionScore     `json:"dimensionScores,omitempty"`
}

// MaturityHeatmapRow represents one row in the maturity heatmap.
type MaturityHeatmapRow struct {
	Domain      string         `json:"domain"`
	StatusCount map[string]int `json:"statusCount"`
	Total       int            `json:"total"`
}

// GapSummary provides a summary of capability gaps.
type GapSummary struct {
	TotalGaps  int            `json:"totalGaps"`
	BySeverity map[string]int `json:"bySeverity"`
	TopGaps    []GapItem      `json:"topGaps"`
}

// GapItem represents a single gap.
type GapItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
	Type        string `json:"type"`
}

// GenerateDashboard creates dashboard data from an ecosystem.
func GenerateDashboard(eco *ecosystem.Ecosystem, config DashboardConfig) DashboardData {
	if config.TopN == 0 {
		config.TopN = 10
	}

	data := DashboardData{
		TotalCapabilities: len(eco.AllCapabilities()),
		TotalInitiatives:  len(eco.AllInitiatives()),
		TotalMetrics:      len(eco.AllMetrics()),
	}

	// Compute pillar rollup
	data.PillarRollup = eco.ComputePortfolioRollup(config.Contributions)
	data.PillarDetails = eco.ComputePillarDetails(config.Contributions, config.InitiativeNames, 5)

	// Compute top initiatives with scores
	data.TopInitiatives = computeTopInitiatives(eco, config)

	// Compute maturity heatmap if enabled
	if config.IncludeMaturityHeatmap {
		data.MaturityHeatmap = computeMaturityHeatmap(eco)
	}

	// Compute gap summary if enabled
	if config.IncludeGapAnalysis {
		data.GapSummary = computeGapSummary(eco)
	}

	return data
}

// computeTopInitiatives ranks initiatives by score.
func computeTopInitiatives(eco *ecosystem.Ecosystem, config DashboardConfig) []RankedInitiative {
	initiatives := eco.AllInitiatives()
	ranked := make([]RankedInitiative, 0, len(initiatives))

	for _, init := range initiatives {
		// Get pillar contributions
		pillars := config.Contributions[init.ID]

		// Compute evidence count
		evidenceCount := 0
		if eco.CapabilityEvidence != nil {
			for _, links := range eco.CapabilityEvidence {
				evidenceCount += len(links)
			}
		}

		// Compute score using scoring config
		score := computeInitiativeScore(init, pillars, evidenceCount, config.ScoringConfig)

		name := config.InitiativeNames[init.ID]
		if name == "" {
			name = init.Name
		}

		ranked = append(ranked, RankedInitiative{
			ID:              init.ID,
			Name:            name,
			Score:           score.TotalScore,
			Pillars:         pillars,
			EvidenceCount:   evidenceCount,
			DimensionScores: score.DimensionScores,
		})
	}

	// Sort by score descending
	sort.Slice(ranked, func(i, j int) bool {
		return ranked[i].Score > ranked[j].Score
	})

	// Assign ranks and limit
	for i := range ranked {
		ranked[i].Rank = i + 1
	}

	if len(ranked) > config.TopN {
		ranked = ranked[:config.TopN]
	}

	return ranked
}

// computeInitiativeScore computes the score for an initiative.
func computeInitiativeScore(
	_ interface{},
	pillars []ecosystem.PillarContribution,
	evidenceCount int,
	config ecosystem.ScoringConfig,
) ecosystem.InitiativeScore {
	score := ecosystem.InitiativeScore{
		DimensionScores: make([]ecosystem.DimensionScore, 0),
	}

	// Strategic alignment from pillars
	alignmentScore := 0.0
	for _, p := range pillars {
		alignmentScore += float64(p.Level.Weight()) / 3.0 // Normalize to 0-1
	}
	if len(pillars) > 0 {
		alignmentScore /= float64(len(pillars))
	}

	// Evidence score
	evidenceScore := float64(evidenceCount) / 10.0 // Normalize
	if evidenceScore > 1.0 {
		evidenceScore = 1.0
	}

	// Add dimension scores
	dims := []struct {
		dim   ecosystem.ScoringDimension
		value float64
	}{
		{ecosystem.DimStrategicAlignment, alignmentScore},
		{ecosystem.DimAdoptionEvidence, evidenceScore},
		{ecosystem.DimCustomerValue, 0.5},     // Default mid-value
		{ecosystem.DimMarketOpportunity, 0.5}, // Default mid-value
		{ecosystem.DimConfidence, 0.7},        // Default reasonable confidence
	}

	totalWeightedScore := 0.0
	totalWeight := 0.0

	for _, d := range dims {
		weight := config.GetWeight(d.dim)
		weighted := d.value * weight
		totalWeightedScore += weighted
		totalWeight += weight

		score.DimensionScores = append(score.DimensionScores, ecosystem.DimensionScore{
			Dimension:     d.dim,
			RawValue:      d.value,
			Weight:        weight,
			WeightedScore: weighted,
		})
	}

	if totalWeight > 0 {
		score.TotalScore = totalWeightedScore / totalWeight * 100 // Scale to 0-100
	}

	return score
}

// computeMaturityHeatmap computes capability status distribution by domain.
func computeMaturityHeatmap(eco *ecosystem.Ecosystem) []MaturityHeatmapRow {
	domainStatus := make(map[string]map[string]int)

	// Get domain from capability stacks
	for _, stack := range eco.CapabilityStacks {
		domain := stack.Metadata.Domain
		if domain == "" {
			domain = "unspecified"
		}

		for _, cap := range stack.AllCapabilities() {
			status := cap.Status
			if status == "" {
				status = "unknown"
			}

			if domainStatus[domain] == nil {
				domainStatus[domain] = make(map[string]int)
			}
			domainStatus[domain][status]++
		}
	}

	rows := make([]MaturityHeatmapRow, 0, len(domainStatus))
	for domain, statuses := range domainStatus {
		total := 0
		for _, count := range statuses {
			total += count
		}
		rows = append(rows, MaturityHeatmapRow{
			Domain:      domain,
			StatusCount: statuses,
			Total:       total,
		})
	}

	// Sort by domain name
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Domain < rows[j].Domain
	})

	return rows
}

// computeGapSummary computes a summary of capability gaps.
func computeGapSummary(eco *ecosystem.Ecosystem) *GapSummary {
	summary := &GapSummary{
		BySeverity: make(map[string]int),
		TopGaps:    make([]GapItem, 0),
	}

	// Count capabilities by status to identify gaps
	for _, cap := range eco.AllCapabilities() {
		// Consider "planned" or "proposed" as gaps
		if cap.Status == "planned" || cap.Status == "proposed" || cap.Status == "missing" {
			summary.TotalGaps++
			severity := "medium"
			if cap.Status == "missing" {
				severity = "high"
			} else if cap.Status == "proposed" {
				severity = "low"
			}
			summary.BySeverity[severity]++

			if len(summary.TopGaps) < 5 {
				summary.TopGaps = append(summary.TopGaps, GapItem{
					ID:          cap.ID,
					Description: cap.Name + ": " + cap.Description,
					Severity:    severity,
					Type:        "capability",
				})
			}
		}
	}

	return summary
}

// GenerateDashboardPage generates the unified dashboard HTML page.
func (g *Generator) GenerateDashboardPage(eco *ecosystem.Ecosystem, config DashboardConfig) error {
	data := GenerateDashboard(eco, config)

	tmpl, err := template.New("dashboard").Parse(dashboardTemplate)
	if err != nil {
		return fmt.Errorf("parsing dashboard template: %w", err)
	}

	// Prepare template data
	tplData := struct {
		Title           string
		Theme           string
		BaseURL         string
		HasSiteNavJS    bool
		HideGeneratedBy bool
		Dashboard       DashboardData
	}{
		Title:           g.config.Title,
		Theme:           g.config.Theme,
		BaseURL:         g.config.BaseURL,
		HasSiteNavJS:    g.config.SiteNavJS != "",
		HideGeneratedBy: g.config.HideGeneratedBy,
		Dashboard:       data,
	}

	// Ensure output directory exists
	if err := os.MkdirAll(g.config.OutputDir, 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	// Write dashboard page
	outPath := filepath.Join(g.config.OutputDir, "dashboard.html")
	f, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("creating dashboard file: %w", err)
	}
	defer f.Close()

	return tmpl.Execute(f, tplData)
}

// ExportDashboardHTML exports a self-contained HTML file with all assets inlined.
func ExportDashboardHTML(eco *ecosystem.Ecosystem, config DashboardConfig, outputPath string) error {
	data := GenerateDashboard(eco, config)

	// Read UI bundle if it exists
	var uiBundle string
	uiBundlePath := filepath.Join(filepath.Dir(outputPath), "..", "ui", "dist", "prism-ui.js")
	if bundleData, err := os.ReadFile(uiBundlePath); err == nil {
		uiBundle = string(bundleData)
	}

	tmpl, err := template.New("export").Parse(selfContainedDashboardTemplate)
	if err != nil {
		return fmt.Errorf("parsing export template: %w", err)
	}

	// Prepare template data
	tplData := struct {
		Title     string
		Dashboard DashboardData
		UIBundle  string
	}{
		Title:     "PRISM Portfolio Dashboard",
		Dashboard: data,
		UIBundle:  uiBundle,
	}

	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating export file: %w", err)
	}
	defer f.Close()

	return tmpl.Execute(f, tplData)
}

// DashboardDataToJSON exports dashboard data as JSON.
func DashboardDataToJSON(data DashboardData) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}
