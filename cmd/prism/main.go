// Package main provides the prism CLI tool for working with PRISM documents.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "prism",
	Short: "PRISM CLI - Platform for Reliability, Improvement, and Strategic Maturity",
	Long: `PRISM is an Operational Product Management platform for COO-level
organizational health monitoring combining SLOs, maturity modeling, and OKRs.

Use the subcommands to create, validate, and score PRISM documents.`,
	Version: version,
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(scoreCmd)
	rootCmd.AddCommand(catalogCmd)
	rootCmd.AddCommand(layerCmd)
	rootCmd.AddCommand(teamCmd)
	rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(goalCmd)
	rootCmd.AddCommand(phaseCmd)
	rootCmd.AddCommand(roadmapCmd)
	rootCmd.AddCommand(initiativeCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(analyzeCmd)
}
