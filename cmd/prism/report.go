package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/prism"
	"github.com/grokify/prism/report"
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report <file>",
	Short: "Generate roadmap report from a PRISM document",
	Long: `Generate a roadmap report in Markdown or JSON format.

The report can be generated in two views:
  - by-phase: Phase → Goal → Initiative (timeline view)
  - by-goal:  Goal → Phase → Initiative (strategic view)
  - both:     Both views in a single document (default)

Examples:
  prism report roadmap.json                    # Markdown to stdout
  prism report roadmap.json -o report.md       # Markdown to file
  prism report roadmap.json --format json      # JSON output
  prism report roadmap.json --view by-phase    # Phase-centric only
  prism report roadmap.json --view by-goal     # Goal-centric only`,
	Args: cobra.ExactArgs(1),
	RunE: runReport,
}

var (
	reportOutput   string
	reportFormat   string
	reportView     string
	reportTitle    string
	reportAuthor   string
	reportNoMeta   bool
	reportNoDetail bool
)

func init() {
	reportCmd.Flags().StringVarP(&reportOutput, "output", "o", "", "Output file (default: stdout)")
	reportCmd.Flags().StringVarP(&reportFormat, "format", "f", "markdown", "Output format: markdown, json")
	reportCmd.Flags().StringVarP(&reportView, "view", "v", "both", "View type: both, by-phase, by-goal")
	reportCmd.Flags().StringVar(&reportTitle, "title", "", "Report title (default: from metadata or 'PRISM Roadmap Report')")
	reportCmd.Flags().StringVar(&reportAuthor, "author", "", "Report author (default: from metadata)")
	reportCmd.Flags().BoolVar(&reportNoMeta, "no-meta", false, "Omit YAML front matter (Markdown only)")
	reportCmd.Flags().BoolVar(&reportNoDetail, "no-detail", false, "Omit initiative details")
	rootCmd.AddCommand(reportCmd)
}

func runReport(cmd *cobra.Command, args []string) error {
	filename := args[0]

	// Read and parse document
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var doc prism.PRISMDocument
	if err := json.Unmarshal(data, &doc); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Validate view type
	switch reportView {
	case "both", "by-phase", "by-goal":
		// Valid
	default:
		return fmt.Errorf("invalid view type: %s (must be: both, by-phase, by-goal)", reportView)
	}

	var output string

	switch reportFormat {
	case "markdown", "md":
		opts := report.DefaultMarkdownOptions()
		opts.ViewType = reportView
		opts.IncludeYAMLMeta = !reportNoMeta
		opts.IncludeDetails = !reportNoDetail

		if reportTitle != "" {
			opts.Title = reportTitle
		} else if doc.Metadata != nil && doc.Metadata.Name != "" {
			opts.Title = doc.Metadata.Name + " Roadmap"
		}

		if reportAuthor != "" {
			opts.Author = reportAuthor
		}

		output = report.GenerateMarkdown(&doc, opts)

	case "json":
		roadmapReport := doc.GenerateRoadmapReport()
		jsonData, err := json.MarshalIndent(roadmapReport, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		output = string(jsonData)

	default:
		return fmt.Errorf("invalid format: %s (must be: markdown, json)", reportFormat)
	}

	// Write output
	if reportOutput != "" {
		if err := os.WriteFile(reportOutput, []byte(output), 0600); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		fmt.Printf("Report written to %s\n", reportOutput)
	} else {
		fmt.Print(output)
	}

	return nil
}
