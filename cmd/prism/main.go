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
	Short: "PRISM CLI - Proactive Reliability & Security Maturity Model",
	Long: `PRISM is a unified framework for B2B SaaS health metrics combining
SLOs, DMAIC, OKRs, and maturity modeling.

Use the subcommands to create, validate, and score PRISM documents.`,
	Version: version,
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(scoreCmd)
	rootCmd.AddCommand(catalogCmd)
}
