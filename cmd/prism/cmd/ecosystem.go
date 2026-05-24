package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/prism/ecosystem"
	"github.com/spf13/cobra"
)

var (
	configFile string
	directory  string
	outputJSON bool
)

var ecosystemCmd = &cobra.Command{
	Use:   "ecosystem",
	Short: "Ecosystem operations",
	Long:  `Load and inspect PRISM ecosystem configurations.`,
}

var ecosystemLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load an ecosystem configuration",
	Long: `Load an ecosystem from a configuration file or directory structure.

Examples:
  prism ecosystem load --config prism.yaml
  prism ecosystem load --dir ./ecosystem`,
	RunE: runEcosystemLoad,
}

func init() {
	ecosystemCmd.AddCommand(ecosystemLoadCmd)

	ecosystemLoadCmd.Flags().StringVarP(&configFile, "config", "c", "", "Configuration file (JSON)")
	ecosystemLoadCmd.Flags().StringVarP(&directory, "dir", "d", "", "Directory to load from")
	ecosystemLoadCmd.Flags().BoolVar(&outputJSON, "json", false, "Output as JSON")
}

func runEcosystemLoad(cmd *cobra.Command, args []string) error {
	var eco *ecosystem.Ecosystem
	var err error

	switch {
	case configFile != "":
		eco, err = ecosystem.LoadFromFile(configFile)
	case directory != "":
		eco, err = ecosystem.LoadFromDirectory(directory)
	default:
		return fmt.Errorf("either --config or --dir is required")
	}

	if err != nil {
		return fmt.Errorf("loading ecosystem: %w", err)
	}

	if outputJSON {
		stats := eco.Stats()
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(stats)
	}

	// Text output
	stats := eco.Stats()
	fmt.Printf("Ecosystem: %s\n\n", eco.Config.Name)
	fmt.Printf("Capability:\n")
	fmt.Printf("  Stacks:       %d\n", stats.CapabilityStacks)
	fmt.Printf("  Capabilities: %d\n", stats.TotalCapabilities)
	fmt.Printf("\nMaturity:\n")
	fmt.Printf("  Documents:    %d\n", stats.PRISMDocuments)
	fmt.Printf("  Metrics:      %d\n", stats.TotalMetrics)
	fmt.Printf("  Services:     %d\n", stats.TotalServices)
	fmt.Printf("  Initiatives:  %d\n", stats.TotalInitiatives)
	fmt.Printf("\nRoadmap:\n")
	fmt.Printf("  OKR Sets:     %d\n", stats.TotalOKRSets)
	fmt.Printf("  Objectives:   %d\n", stats.TotalObjectives)
	fmt.Printf("  Roadmaps:     %d\n", stats.TotalRoadmaps)
	fmt.Printf("  Phases:       %d\n", stats.TotalPhases)

	if len(stats.ByDomain) > 0 {
		fmt.Printf("\nBy Domain:\n")
		for domain, count := range stats.ByDomain {
			fmt.Printf("  %s: %d\n", domain, count)
		}
	}

	if len(stats.ByStatus) > 0 {
		fmt.Printf("\nBy Status:\n")
		for status, count := range stats.ByStatus {
			fmt.Printf("  %s: %d\n", status, count)
		}
	}

	return nil
}
