package sitegen

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/grokify/prism/ecosystem"
)

// PillarConfig holds configuration for pillar dashboard generation.
type PillarConfig struct {
	// Contributions maps initiative IDs to their pillar contributions.
	Contributions ecosystem.InitiativePillars

	// InitiativeNames maps initiative IDs to display names.
	InitiativeNames map[string]string

	// TopN is the number of top initiatives to show per pillar.
	TopN int
}

// GeneratePillarDashboard generates the pillar investment dashboard page.
func (g *Generator) GeneratePillarDashboard(config PillarConfig) error {
	if config.TopN == 0 {
		config.TopN = 5
	}

	// Compute rollup and details
	eco := &ecosystem.Ecosystem{}
	rollup := eco.ComputePortfolioRollup(config.Contributions)
	details := eco.ComputePillarDetails(config.Contributions, config.InitiativeNames, config.TopN)

	// Create template with custom functions
	funcMap := template.FuncMap{
		"pillarClass": func(p ecosystem.Pillar) string {
			switch p {
			case ecosystem.PillarCSAT:
				return "csat"
			case ecosystem.PillarSAMSOM:
				return "sam-som"
			case ecosystem.PillarTAM:
				return "tam"
			default:
				return strings.ToLower(string(p))
			}
		},
	}

	tmpl, err := template.New("pillar-dashboard").Funcs(funcMap).Parse(pillarDashboardTemplate)
	if err != nil {
		return err
	}

	// Prepare template data
	data := struct {
		Title           string
		Theme           string
		BaseURL         string
		HasSiteNavJS    bool
		HideGeneratedBy bool
		Rollup          ecosystem.PortfolioRollup
		Details         []ecosystem.PillarDetail
	}{
		Title:           g.config.Title,
		Theme:           g.config.Theme,
		BaseURL:         g.config.BaseURL,
		HasSiteNavJS:    g.config.SiteNavJS != "",
		HideGeneratedBy: g.config.HideGeneratedBy,
		Rollup:          rollup,
		Details:         details,
	}

	// Ensure output directory exists
	if err := os.MkdirAll(g.config.OutputDir, 0755); err != nil {
		return err
	}

	// Write pillar dashboard
	outPath := filepath.Join(g.config.OutputDir, "pillars.html")
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}

// GeneratePillarDashboardFromEcosystem generates the pillar dashboard from ecosystem initiatives.
// This is a convenience method that extracts initiative names from the ecosystem.
func (g *Generator) GeneratePillarDashboardFromEcosystem(eco *ecosystem.Ecosystem, contributions ecosystem.InitiativePillars) error {
	// Build initiative names from ecosystem
	names := make(map[string]string)
	for _, init := range eco.AllInitiatives() {
		names[init.ID] = init.Name
	}

	return g.GeneratePillarDashboard(PillarConfig{
		Contributions:   contributions,
		InitiativeNames: names,
		TopN:            5,
	})
}
