package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/prism"
	"github.com/spf13/cobra"
)

var (
	initDomain string
	initOutput string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new PRISM document",
	Long: `Create a new PRISM document scaffold with default structure.

Examples:
  prism init                          # Create default prism.json with operations metrics
  prism init -d operations -o ops.json # Create ops-focused document

For security metrics examples, see: https://github.com/grokify/prism-security`,
	RunE: runInit,
}

func init() {
	initCmd.Flags().StringVarP(&initDomain, "domain", "d", "", "Focus domain (operations). For security examples, see prism-security.")
	initCmd.Flags().StringVarP(&initOutput, "output", "o", "prism.json", "Output file path")
}

func runInit(cmd *cobra.Command, args []string) error {
	// Determine which domains to include
	var domains []prism.DomainDef
	var domainNames []string

	if initDomain == "" {
		// Include both domains
		domains = []prism.DomainDef{
			{Name: prism.DomainSecurity, Description: "Security metrics and controls", Weight: 0.5},
			{Name: prism.DomainOperations, Description: "Operational metrics and SLOs", Weight: 0.5},
		}
		domainNames = prism.AllDomains()
	} else {
		// Include only the specified domain
		switch initDomain {
		case prism.DomainSecurity:
			domains = []prism.DomainDef{
				{Name: prism.DomainSecurity, Description: "Security metrics and controls", Weight: 1.0},
			}
		case prism.DomainOperations:
			domains = []prism.DomainDef{
				{Name: prism.DomainOperations, Description: "Operational metrics and SLOs", Weight: 1.0},
			}
		default:
			return fmt.Errorf("invalid domain %q, must be 'security' or 'operations'", initDomain)
		}
		domainNames = []string{initDomain}
	}

	// Create document with metadata
	doc := &prism.PRISMDocument{
		Schema: "https://github.com/grokify/prism/schema/prism.schema.json",
		Metadata: &prism.Metadata{
			Name:        "My PRISM Document",
			Description: "PRISM metrics for SaaS health monitoring",
			Version:     "1.0.0",
		},
		Domains:  domains,
		Maturity: prism.NewMaturityModelForDomains(domainNames),
		Metrics:  make([]prism.Metric, 0),
	}

	// Add example metrics (operations only - for security examples see prism-security)
	if initDomain == "" || initDomain == prism.DomainOperations {
		doc.Metrics = append(doc.Metrics, createOperationsMetrics()...)
	}
	if initDomain == prism.DomainSecurity {
		fmt.Println("Note: For security metric examples, see https://github.com/grokify/prism-security")
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal document: %w", err)
	}

	// Write to file (0644 is appropriate for shareable config files)
	if err := os.WriteFile(initOutput, data, 0644); err != nil { //nolint:gosec // G306: PRISM docs are shareable configs
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("Created %s\n", initOutput)
	return nil
}

func createOperationsMetrics() []prism.Metric {
	return []prism.Metric{
		{
			ID:             "ops-availability-01",
			Name:           "Service Availability",
			Description:    "Percentage of time the service is available",
			Domain:         prism.DomainOperations,
			Stage:          prism.StageRuntime,
			Category:       prism.CategoryReliability,
			MetricType:     prism.MetricTypeRate,
			TrendDirection: prism.TrendHigherBetter,
			Unit:           "%",
			Baseline:       99.0,
			Current:        99.9,
			Target:         99.95,
			Thresholds:     &prism.Thresholds{Green: 99.9, Yellow: 99.5, Red: 99.0},
			SLI:            &prism.SLI{Name: "Availability", Formula: "successful_requests / total_requests"},
			SLO:            &prism.SLO{Target: ">=99.95%", Window: prism.Window30Days},
			FrameworkMappings: []prism.FrameworkMapping{
				{Framework: prism.FrameworkSRE, Reference: "availability-slo"},
				{Framework: prism.FrameworkDORA, Reference: "availability"},
			},
		},
		{
			ID:             "ops-latency-01",
			Name:           "P99 Latency",
			Description:    "99th percentile response latency",
			Domain:         prism.DomainOperations,
			Stage:          prism.StageRuntime,
			Category:       prism.CategoryEfficiency,
			MetricType:     prism.MetricTypeLatency,
			TrendDirection: prism.TrendLowerBetter,
			Unit:           "ms",
			Baseline:       500,
			Current:        200,
			Target:         100,
			Thresholds:     &prism.Thresholds{Green: 150, Yellow: 300, Red: 500},
			SLI:            &prism.SLI{Name: "Latency", Formula: "percentile(response_time, 99)"},
			SLO:            &prism.SLO{Target: "<=100ms", Window: prism.Window7Days},
		},
	}
}
