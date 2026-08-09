package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/grokify/prism-maturity/analysis"
	"github.com/grokify/prism/ecosystem"
	"github.com/spf13/cobra"
)

var (
	gapsConfigFile string
	gapsDirectory  string
	gapsJSON       bool
	gapsSortBy     string
)

var gapsCmd = &cobra.Command{
	Use:   "gaps",
	Short: "Gap analysis commands",
	Long:  `Commands for identifying and analyzing gaps across the ecosystem.`,
}

var gapsAnalyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze gaps across ecosystem documents",
	Long: `Analyze maturity gaps, SLO gaps, and initiative gaps across all loaded documents.

Gaps are sorted by impact (default), which combines severity and maturity gap size.

Output Format:
  TYPE       Gap type: maturity, slo, or initiative
  SEVERITY   High, medium, or low based on gap magnitude
  IMPACT     Numeric score combining severity and context
  GOAL/PHASE Related goal or phase ID
  DESCRIPTION Details about the gap

Examples:
  prism gaps analyze --config prism.yaml
  prism gaps analyze --dir ./ecosystem
  prism gaps analyze --sort-by=severity
  prism gaps analyze --json`,
	RunE: runGapsAnalyze,
}

func init() {
	gapsCmd.AddCommand(gapsAnalyzeCmd)

	gapsAnalyzeCmd.Flags().StringVarP(&gapsConfigFile, "config", "c", "", "Configuration file (JSON)")
	gapsAnalyzeCmd.Flags().StringVarP(&gapsDirectory, "dir", "d", "", "Directory to load from")
	gapsAnalyzeCmd.Flags().BoolVar(&gapsJSON, "json", false, "Output as JSON")
	gapsAnalyzeCmd.Flags().StringVar(&gapsSortBy, "sort-by", "impact", "Sort gaps by: impact, severity, type")
}

// GapResult extends analysis.Gap with computed fields for sorting and display.
type GapResult struct {
	Type        string `json:"type"`
	Severity    string `json:"severity"`
	Impact      int    `json:"impact"`
	GoalID      string `json:"goalId,omitempty"`
	PhaseID     string `json:"phaseId,omitempty"`
	Description string `json:"description"`
	DocumentID  string `json:"documentId,omitempty"`
}

// GapsOutput is the structured output for gap analysis.
type GapsOutput struct {
	Summary GapsSummary `json:"summary"`
	Gaps    []GapResult `json:"gaps"`
}

// GapsSummary provides aggregate statistics.
type GapsSummary struct {
	TotalGaps      int            `json:"totalGaps"`
	BySeverity     map[string]int `json:"bySeverity"`
	ByType         map[string]int `json:"byType"`
	DocumentsCount int            `json:"documentsAnalyzed"`
}

func runGapsAnalyze(cmd *cobra.Command, args []string) error {
	var eco *ecosystem.Ecosystem
	var err error

	switch {
	case gapsConfigFile != "":
		eco, err = ecosystem.LoadFromFile(gapsConfigFile)
	case gapsDirectory != "":
		eco, err = ecosystem.LoadFromDirectory(gapsDirectory)
	default:
		return fmt.Errorf("either --config or --dir is required")
	}

	if err != nil {
		return fmt.Errorf("loading ecosystem: %w", err)
	}

	output := analyzeEcosystemGaps(eco)

	sortGaps(output.Gaps, gapsSortBy)

	if gapsJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(output)
	}

	return printGapsTable(output)
}

func analyzeEcosystemGaps(eco *ecosystem.Ecosystem) *GapsOutput {
	output := &GapsOutput{
		Summary: GapsSummary{
			BySeverity:     make(map[string]int),
			ByType:         make(map[string]int),
			DocumentsCount: len(eco.PRISMDocuments),
		},
		Gaps: make([]GapResult, 0),
	}

	for i, doc := range eco.PRISMDocuments {
		result := analysis.Analyze(doc)
		docID := fmt.Sprintf("doc[%d]", i)
		if doc.Metadata != nil && doc.Metadata.Name != "" {
			docID = doc.Metadata.Name
		}

		for _, gap := range result.Gaps {
			gr := GapResult{
				Type:        string(gap.Type),
				Severity:    string(gap.Severity),
				Impact:      calculateImpact(gap),
				GoalID:      gap.GoalID,
				PhaseID:     gap.PhaseID,
				Description: gap.Description,
				DocumentID:  docID,
			}
			output.Gaps = append(output.Gaps, gr)

			output.Summary.BySeverity[gr.Severity]++
			output.Summary.ByType[gr.Type]++
		}
	}

	output.Summary.TotalGaps = len(output.Gaps)
	return output
}

// calculateImpact computes a numeric impact score for sorting.
// Impact = severity weight * type weight, allowing consistent prioritization.
func calculateImpact(gap analysis.Gap) int {
	severityWeight := map[analysis.Severity]int{
		analysis.SeverityHigh:   3,
		analysis.SeverityMedium: 2,
		analysis.SeverityLow:    1,
	}

	typeWeight := map[analysis.GapType]int{
		analysis.GapTypeMaturity:   3,
		analysis.GapTypeSLO:        2,
		analysis.GapTypeInitiative: 1,
	}

	sw := severityWeight[gap.Severity]
	tw := typeWeight[gap.Type]
	if sw == 0 {
		sw = 1
	}
	if tw == 0 {
		tw = 1
	}

	return sw * tw
}

func sortGaps(gaps []GapResult, sortBy string) {
	switch sortBy {
	case "severity":
		sort.Slice(gaps, func(i, j int) bool {
			si := severityOrder(gaps[i].Severity)
			sj := severityOrder(gaps[j].Severity)
			if si != sj {
				return si < sj
			}
			return gaps[i].Impact > gaps[j].Impact
		})
	case "type":
		sort.Slice(gaps, func(i, j int) bool {
			if gaps[i].Type != gaps[j].Type {
				return gaps[i].Type < gaps[j].Type
			}
			return gaps[i].Impact > gaps[j].Impact
		})
	default: // "impact"
		sort.Slice(gaps, func(i, j int) bool {
			return gaps[i].Impact > gaps[j].Impact
		})
	}
}

func severityOrder(s string) int {
	switch s {
	case "high":
		return 0
	case "medium":
		return 1
	case "low":
		return 2
	default:
		return 3
	}
}

func printGapsTable(output *GapsOutput) error {
	fmt.Printf("Gap Analysis Results\n")
	fmt.Printf("====================\n\n")

	fmt.Printf("Summary:\n")
	fmt.Printf("  Documents analyzed: %d\n", output.Summary.DocumentsCount)
	fmt.Printf("  Total gaps:         %d\n", output.Summary.TotalGaps)
	if output.Summary.BySeverity["high"] > 0 {
		fmt.Printf("  High severity:      %d\n", output.Summary.BySeverity["high"])
	}
	if output.Summary.BySeverity["medium"] > 0 {
		fmt.Printf("  Medium severity:    %d\n", output.Summary.BySeverity["medium"])
	}
	if output.Summary.BySeverity["low"] > 0 {
		fmt.Printf("  Low severity:       %d\n", output.Summary.BySeverity["low"])
	}
	fmt.Println()

	if len(output.Gaps) == 0 {
		fmt.Println("No gaps identified.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "TYPE\tSEVERITY\tIMPACT\tGOAL/PHASE\tDESCRIPTION")
	fmt.Fprintln(w, "----\t--------\t------\t----------\t-----------")

	for _, g := range output.Gaps {
		ref := g.GoalID
		if ref == "" {
			ref = g.PhaseID
		}
		desc := g.Description
		if len(desc) > 60 {
			desc = desc[:57] + "..."
		}
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\n",
			g.Type, strings.ToUpper(g.Severity), g.Impact, ref, desc)
	}

	return w.Flush()
}
