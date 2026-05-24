// Package cmd provides the prism CLI commands.
package cmd

import (
	"github.com/spf13/cobra"

	capcli "github.com/grokify/prism-capability/cli"
	maturitycli "github.com/grokify/prism-maturity/cli"
	roadmapcli "github.com/grokify/prism-roadmap/cli"
)

var rootCmd = &cobra.Command{
	Use:   "prism",
	Short: "PRISM - Platform for Reliability, Intelligence, Strategy & Maturity",
	Long: `PRISM is a unified framework for capability-driven organizational intelligence,
connecting what you need (capabilities), how you measure (maturity), and how you act (roadmap).

Subcommands:
  - capability: Capability stacks and dependencies (prism-capability)
  - maturity:   SLIs, SLOs, and maturity models (prism-maturity)
  - roadmap:    OKRs, V2MOMs, and planning documents (prism-roadmap)

Additional commands:
  - ecosystem: View PRISM ecosystem overview
  - validate:  Validate PRISM documents
  - stats:     Show repository statistics
  - site:      Generate static websites from PRISM data`,
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add sub-repo command trees
	rootCmd.AddCommand(capcli.RootCmd)      // prism capability ...
	rootCmd.AddCommand(maturitycli.RootCmd) // prism maturity ...
	rootCmd.AddCommand(roadmapcli.RootCmd)  // prism roadmap ...

	// Add umbrella-specific commands
	rootCmd.AddCommand(ecosystemCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(statsCmd)
	rootCmd.AddCommand(siteCmd)
}
