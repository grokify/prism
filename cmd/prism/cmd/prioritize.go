package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/grokify/prism/ecosystem"
	"github.com/spf13/cobra"
)

var (
	prioritizeConfigFile string
	prioritizePillar     string
	prioritizeFormat     string
	prioritizeTop        int
	prioritizeDir        string
)

var prioritizeCmd = &cobra.Command{
	Use:   "prioritize",
	Short: "Rank initiatives by configurable multi-dimensional scoring",
	Long: `Prioritize initiatives using a configurable scoring engine.

Scoring dimensions:
  - customer_value:      Direct customer satisfaction and value impact
  - market_opportunity:  Market size, TAM/SAM/SOM potential
  - strategic_alignment: Fit with company strategy and pillars
  - engineering_cost:    Development effort (inverted: lower cost = higher score)
  - risk:                Technical and execution risk (inverted)
  - adoption_evidence:   Customer demand signals, votes, requests
  - confidence:          Certainty in estimates and assumptions

The scoring weights can be configured via a JSON file (--config).
Default weights provide a balanced scoring approach.

Examples:
  # Rank all initiatives with default weights
  prism prioritize --dir ./ecosystem

  # Rank by specific pillar
  prism prioritize --dir ./ecosystem --pillar CSAT

  # Use custom weights and show top 10
  prism prioritize --dir ./ecosystem --config weights.json --top 10

  # Output as JSON
  prism prioritize --dir ./ecosystem --format json`,
	RunE: runPrioritize,
}

func init() {
	prioritizeCmd.Flags().StringVar(&prioritizeConfigFile, "config", "", "Path to scoring weights JSON file")
	prioritizeCmd.Flags().StringVar(&prioritizePillar, "pillar", "all", "Filter by pillar: CSAT, SAM-SOM, TAM, or all")
	prioritizeCmd.Flags().StringVar(&prioritizeFormat, "format", "table", "Output format: table or json")
	prioritizeCmd.Flags().IntVar(&prioritizeTop, "top", 0, "Limit to top N results (0 = all)")
	prioritizeCmd.Flags().StringVar(&prioritizeDir, "dir", "", "Ecosystem directory to load")
}

func runPrioritize(cmd *cobra.Command, args []string) error {
	// Load scoring config
	var config *ecosystem.ScoringConfig
	if prioritizeConfigFile != "" {
		var err error
		config, err = ecosystem.LoadScoringConfig(prioritizeConfigFile)
		if err != nil {
			return fmt.Errorf("loading scoring config: %w", err)
		}
	} else {
		config = ecosystem.DefaultScoringConfig()
	}

	// Load ecosystem if directory provided
	var eco *ecosystem.Ecosystem
	if prioritizeDir != "" {
		var err error
		eco, err = ecosystem.LoadFromDirectory(prioritizeDir)
		if err != nil {
			return fmt.Errorf("loading ecosystem: %w", err)
		}
	}

	// Build initiative inputs
	inputs := buildInitiativeInputs(eco)

	if len(inputs) == 0 {
		fmt.Println("No initiatives found to prioritize.")
		fmt.Println("\nTo provide initiative data, either:")
		fmt.Println("  1. Load an ecosystem with --dir containing PRISM documents")
		fmt.Println("  2. Provide initiative data via stdin (future feature)")
		return nil
	}

	// Rank initiatives
	var result *ecosystem.RankingResult
	pillarUpper := strings.ToUpper(prioritizePillar)

	if pillarUpper == "ALL" || prioritizePillar == "" {
		result = ecosystem.RankInitiatives(inputs, config)
	} else {
		pillar := ecosystem.Pillar(pillarUpper)
		result = ecosystem.RankInitiativesByPillar(inputs, config, pillar)
	}

	// Apply top N filter
	if prioritizeTop > 0 {
		result.Initiatives = result.TopN(prioritizeTop)
	}

	// Output results
	switch prioritizeFormat {
	case "json":
		return outputPrioritizeJSON(result)
	default:
		return outputPrioritizeTable(result)
	}
}

func buildInitiativeInputs(eco *ecosystem.Ecosystem) []*ecosystem.InitiativeInput {
	var inputs []*ecosystem.InitiativeInput

	if eco == nil {
		return inputs
	}

	// Convert maturity initiatives to scoring inputs
	for _, init := range eco.AllInitiatives() {
		input := &ecosystem.InitiativeInput{
			ID:   init.ID,
			Name: init.Name,
			Sources: map[ecosystem.ScoringDimension]string{
				ecosystem.DimCustomerValue:      "maturity-initiative",
				ecosystem.DimStrategicAlignment: "maturity-initiative",
			},
		}

		// Map priority to strategic alignment (P0=100, P1=75, P2=50, P3=25)
		switch init.Priority {
		case 0:
			input.StrategicAlignment = 100
		case 1:
			input.StrategicAlignment = 75
		case 2:
			input.StrategicAlignment = 50
		case 3:
			input.StrategicAlignment = 25
		default:
			input.StrategicAlignment = 50
		}

		// Use metric count as proxy for customer value
		input.CustomerValue = float64(len(init.MetricIDs)) * 20
		if input.CustomerValue > 100 {
			input.CustomerValue = 100
		}

		// Check for evidence links and enrich if resolver available
		if eco.CapabilityEvidence != nil {
			if links, ok := eco.CapabilityEvidence[init.ID]; ok && len(links) > 0 {
				if eco.EvidenceResolver != nil {
					resolved, err := eco.EvidenceResolver.ResolveAll(links)
					if err == nil {
						summary := ecosystem.ComputeEvidenceSummary(links, resolved)
						ecosystem.EnrichInputFromEvidence(input, summary)
					}
				}
			}
		}

		inputs = append(inputs, input)
	}

	return inputs
}

func outputPrioritizeTable(result *ecosystem.RankingResult) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Header
	fmt.Fprintf(w, "RANK\tID\tNAME\tSCORE\tNORM\tPILLAR\tEXPLANATION\n")
	fmt.Fprintf(w, "----\t--\t----\t-----\t----\t------\t-----------\n")

	for _, ri := range result.Initiatives {
		s := ri.Score
		name := s.InitiativeName
		if len(name) > 30 {
			name = name[:27] + "..."
		}

		explanation := s.Explanation
		if len(explanation) > 50 {
			explanation = explanation[:47] + "..."
		}

		pillar := string(s.Pillar)
		if pillar == "" {
			pillar = "-"
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%.1f\t%.0f\t%s\t%s\n",
			ri.Rank,
			s.InitiativeID,
			name,
			s.TotalScore,
			s.NormalizedScore,
			pillar,
			explanation,
		)
	}

	w.Flush()

	// Summary
	fmt.Printf("\nTotal: %d initiatives", result.TotalCount)
	if result.Pillar != "" {
		fmt.Printf(" (pillar: %s)", result.Pillar)
	}
	fmt.Printf(" | Config: %s\n", result.Config.Name)

	return nil
}

func outputPrioritizeJSON(result *ecosystem.RankingResult) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(result)
}
