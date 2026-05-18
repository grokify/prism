package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/prism/ecosystem"
	"github.com/spf13/cobra"
)

var (
	statsConfigFile string
	statsDirectory  string
	statsJSON       bool
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show ecosystem statistics",
	Long: `Display summary statistics for the ecosystem.

Examples:
  prism stats --config prism.yaml
  prism stats --dir ./ecosystem
  prism stats --json`,
	RunE: runStats,
}

func init() {
	statsCmd.Flags().StringVarP(&statsConfigFile, "config", "c", "", "Configuration file (JSON)")
	statsCmd.Flags().StringVarP(&statsDirectory, "dir", "d", "", "Directory to load from")
	statsCmd.Flags().BoolVar(&statsJSON, "json", false, "Output as JSON")
}

func runStats(cmd *cobra.Command, args []string) error {
	var eco *ecosystem.Ecosystem
	var err error

	switch {
	case statsConfigFile != "":
		eco, err = ecosystem.LoadFromFile(statsConfigFile)
	case statsDirectory != "":
		eco, err = ecosystem.LoadFromDirectory(statsDirectory)
	default:
		return fmt.Errorf("either --config or --dir is required")
	}

	if err != nil {
		return fmt.Errorf("loading ecosystem: %w", err)
	}

	stats := eco.Stats()

	if statsJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(stats)
	}

	// Text output
	fmt.Printf("PRISM Ecosystem Statistics\n")
	fmt.Printf("==========================\n\n")

	fmt.Printf("Capability Module\n")
	fmt.Printf("  Stacks:        %d\n", stats.CapabilityStacks)
	fmt.Printf("  Capabilities:  %d\n", stats.TotalCapabilities)

	fmt.Printf("\nIntelligence Module\n")
	fmt.Printf("  Documents:     %d\n", stats.PRISMDocuments)
	fmt.Printf("  Metrics:       %d\n", stats.TotalMetrics)
	fmt.Printf("  Services:      %d\n", stats.TotalServices)
	fmt.Printf("  Initiatives:   %d\n", stats.TotalInitiatives)

	fmt.Printf("\nExecution Module\n")
	fmt.Printf("  OKR Sets:      %d\n", stats.TotalOKRSets)
	fmt.Printf("  Objectives:    %d\n", stats.TotalObjectives)
	fmt.Printf("  Roadmaps:      %d\n", stats.TotalRoadmaps)
	fmt.Printf("  Phases:        %d\n", stats.TotalPhases)

	if len(stats.ByDomain) > 0 {
		fmt.Printf("\nCapabilities by Domain\n")
		for domain, count := range stats.ByDomain {
			fmt.Printf("  %-20s %d\n", domain+":", count)
		}
	}

	if len(stats.ByStatus) > 0 {
		fmt.Printf("\nCapabilities by Status\n")
		for status, count := range stats.ByStatus {
			fmt.Printf("  %-20s %d\n", status+":", count)
		}
	}

	return nil
}
