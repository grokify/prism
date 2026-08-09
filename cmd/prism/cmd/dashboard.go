package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/prism/ecosystem"
	"github.com/grokify/prism/sitegen"
	"github.com/spf13/cobra"
)

var (
	dashboardDir            string
	dashboardConfig         string
	dashboardOutput         string
	dashboardExport         string
	dashboardFormat         string
	dashboardTheme          string
	dashboardIncludeGaps    bool
	dashboardIncludeHeatmap bool
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Generate portfolio dashboard",
	Long: `Generate a unified portfolio dashboard showing:
  - Pillar investment mix (CSAT / SAM-SOM / TAM distribution)
  - Top initiatives with prioritization scores
  - Capability maturity heatmap by domain
  - Gap analysis summary

The dashboard can be generated as an HTML page for the site or exported
as a self-contained HTML file for distribution.`,
}

var dashboardGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate dashboard HTML for the site",
	Long: `Generate the portfolio dashboard as part of the static site.

Example:
  prism dashboard generate --dir ./ecosystem --output ./dist`,
	RunE: runDashboardGenerate,
}

var dashboardExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export self-contained dashboard HTML",
	Long: `Export the portfolio dashboard as a single self-contained HTML file
with all CSS and JavaScript inlined. The exported file works offline
and can be shared or embedded.

Example:
  prism dashboard export --dir ./ecosystem --export ./dashboard.html`,
	RunE: runDashboardExport,
}

var dashboardDataCmd = &cobra.Command{
	Use:   "data",
	Short: "Output dashboard data as JSON",
	Long: `Generate dashboard data and output as JSON for custom visualization
or integration with other tools.

Example:
  prism dashboard data --dir ./ecosystem > dashboard.json`,
	RunE: runDashboardData,
}

func init() {
	// Common flags
	dashboardCmd.PersistentFlags().StringVar(&dashboardDir, "dir", "", "Ecosystem directory to load")
	dashboardCmd.PersistentFlags().StringVar(&dashboardConfig, "config", "", "Ecosystem config file (JSON)")
	dashboardCmd.PersistentFlags().StringVar(&dashboardTheme, "theme", "light", "Theme: light or dark")
	dashboardCmd.PersistentFlags().BoolVar(&dashboardIncludeGaps, "gaps", true, "Include gap analysis")
	dashboardCmd.PersistentFlags().BoolVar(&dashboardIncludeHeatmap, "heatmap", true, "Include maturity heatmap")

	// Generate-specific flags
	dashboardGenerateCmd.Flags().StringVar(&dashboardOutput, "output", "./dist", "Output directory for site")

	// Export-specific flags
	dashboardExportCmd.Flags().StringVar(&dashboardExport, "export", "./dashboard.html", "Output file path")

	// Data-specific flags
	dashboardDataCmd.Flags().StringVar(&dashboardFormat, "format", "json", "Output format: json")

	dashboardCmd.AddCommand(dashboardGenerateCmd)
	dashboardCmd.AddCommand(dashboardExportCmd)
	dashboardCmd.AddCommand(dashboardDataCmd)
}

func loadEcosystemForDashboard() (*ecosystem.Ecosystem, error) {
	if dashboardConfig != "" {
		return ecosystem.LoadFromFile(dashboardConfig)
	}
	if dashboardDir != "" {
		return ecosystem.LoadFromDirectory(dashboardDir)
	}
	return nil, fmt.Errorf("either --dir or --config is required")
}

func buildDashboardConfig() sitegen.DashboardConfig {
	return sitegen.DashboardConfig{
		ScoringConfig:          *ecosystem.DefaultScoringConfig(),
		Contributions:          make(ecosystem.InitiativePillars),
		InitiativeNames:        make(map[string]string),
		TopN:                   10,
		IncludeMaturityHeatmap: dashboardIncludeHeatmap,
		IncludeGapAnalysis:     dashboardIncludeGaps,
	}
}

func runDashboardGenerate(cmd *cobra.Command, args []string) error {
	eco, err := loadEcosystemForDashboard()
	if err != nil {
		return fmt.Errorf("loading ecosystem: %w", err)
	}

	// Build initiative names from ecosystem
	config := buildDashboardConfig()
	for _, init := range eco.AllInitiatives() {
		config.InitiativeNames[init.ID] = init.Name
	}

	// Create generator
	gen := sitegen.NewGenerator(sitegen.Config{
		Title:     "PRISM Portfolio",
		OutputDir: dashboardOutput,
		Theme:     dashboardTheme,
	})

	if err := gen.GenerateDashboardPage(eco, config); err != nil {
		return fmt.Errorf("generating dashboard: %w", err)
	}

	fmt.Printf("Dashboard generated: %s/dashboard.html\n", dashboardOutput)
	return nil
}

func runDashboardExport(cmd *cobra.Command, args []string) error {
	eco, err := loadEcosystemForDashboard()
	if err != nil {
		return fmt.Errorf("loading ecosystem: %w", err)
	}

	config := buildDashboardConfig()
	for _, init := range eco.AllInitiatives() {
		config.InitiativeNames[init.ID] = init.Name
	}

	if err := sitegen.ExportDashboardHTML(eco, config, dashboardExport); err != nil {
		return fmt.Errorf("exporting dashboard: %w", err)
	}

	fmt.Printf("Dashboard exported: %s\n", dashboardExport)
	return nil
}

func runDashboardData(cmd *cobra.Command, args []string) error {
	eco, err := loadEcosystemForDashboard()
	if err != nil {
		return fmt.Errorf("loading ecosystem: %w", err)
	}

	config := buildDashboardConfig()
	for _, init := range eco.AllInitiatives() {
		config.InitiativeNames[init.ID] = init.Name
	}

	data := sitegen.GenerateDashboard(eco, config)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling JSON: %w", err)
	}

	_, err = os.Stdout.Write(jsonData)
	return err
}
