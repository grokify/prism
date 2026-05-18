// Package cmd provides the prism CLI commands.
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prism",
	Short: "PRISM - Platform for Reliability, Intelligence, Strategy & Maturity",
	Long: `PRISM is a unified framework for capability-driven organizational intelligence,
connecting what you need (capabilities), how you measure (maturity), and how you act (execution).

Ecosystem orchestrator for:
  - prism-capability: Capability stacks and dependencies
  - prism-intelligence: SLIs, SLOs, and maturity models
  - prism-execution: OKRs, roadmaps, and initiatives`,
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(ecosystemCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(statsCmd)
}
