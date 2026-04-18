package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/prism"
	"github.com/grokify/prism/report"
	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard <file>",
	Short: "Generate executive dashboard from a PRISM document",
	Long: `Generate an executive-level dashboard showing maturity progress,
SLO compliance, and roadmap status.

Output formats:
  --format json      JSON data (default)
  --format markdown  Pandoc-compatible markdown
  --format marp      Marp presentation slides
  --format html      Standalone HTML dashboard

Examples:
  prism dashboard roadmap.json                     # JSON to stdout
  prism dashboard roadmap.json -f markdown         # Pandoc markdown
  prism dashboard roadmap.json -f html -o dash.html # HTML file
  prism dashboard roadmap.json -f marp -o slides.md # Marp slides`,
	Args: cobra.ExactArgs(1),
	RunE: runDashboard,
}

var (
	dashboardFormat  string
	dashboardOutput  string
	dashboardTitle   string
	dashboardAuthor  string
	dashboardMaxGaps int
)

func init() {
	dashboardCmd.Flags().StringVarP(&dashboardFormat, "format", "f", "json", "Output format: json, markdown, marp, html")
	dashboardCmd.Flags().StringVarP(&dashboardOutput, "output", "o", "", "Output file (default: stdout)")
	dashboardCmd.Flags().StringVar(&dashboardTitle, "title", "", "Dashboard title (default: from metadata)")
	dashboardCmd.Flags().StringVar(&dashboardAuthor, "author", "", "Dashboard author")
	dashboardCmd.Flags().IntVar(&dashboardMaxGaps, "max-gaps", 10, "Maximum gaps to show (0 = all)")
	rootCmd.AddCommand(dashboardCmd)
}

func runDashboard(cmd *cobra.Command, args []string) error {
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

	// Generate dashboard data
	dashboard := doc.GenerateExecutiveDashboard()

	// Apply title override
	if dashboardTitle != "" {
		dashboard.Title = dashboardTitle
	}

	var output string

	switch dashboardFormat {
	case "json":
		jsonData, err := json.MarshalIndent(dashboard, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		output = string(jsonData)

	case "markdown", "md":
		opts := report.DefaultDashboardOptions()
		if dashboardTitle != "" {
			opts.Title = dashboardTitle
		} else if dashboard.Title != "" {
			opts.Title = dashboard.Title
		}
		opts.Author = dashboardAuthor
		opts.MaxGaps = dashboardMaxGaps
		output = report.GenerateDashboardMarkdown(dashboard, opts)

	case "marp":
		opts := report.DefaultDashboardOptions()
		if dashboardTitle != "" {
			opts.Title = dashboardTitle
		} else if dashboard.Title != "" {
			opts.Title = dashboard.Title
		}
		opts.Author = dashboardAuthor
		opts.MaxGaps = dashboardMaxGaps
		output = report.GenerateDashboardMarp(dashboard, opts)

	case "html":
		opts := report.DefaultDashboardOptions()
		if dashboardTitle != "" {
			opts.Title = dashboardTitle
		} else if dashboard.Title != "" {
			opts.Title = dashboard.Title
		}
		opts.Author = dashboardAuthor
		opts.MaxGaps = dashboardMaxGaps
		output = report.GenerateDashboardHTML(dashboard, opts)

	default:
		return fmt.Errorf("unknown format: %s (must be: json, markdown, marp, html)", dashboardFormat)
	}

	// Write output
	if dashboardOutput != "" {
		if err := os.WriteFile(dashboardOutput, []byte(output), 0600); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		fmt.Printf("Dashboard written to %s\n", dashboardOutput)
	} else {
		fmt.Print(output)
	}

	return nil
}
