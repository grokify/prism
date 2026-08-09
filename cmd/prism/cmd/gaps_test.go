package cmd

import (
	"testing"

	"github.com/grokify/prism-maturity/analysis"
)

func TestCalculateImpact(t *testing.T) {
	tests := []struct {
		name     string
		gap      analysis.Gap
		expected int
	}{
		{
			name: "high severity maturity gap",
			gap: analysis.Gap{
				Type:     analysis.GapTypeMaturity,
				Severity: analysis.SeverityHigh,
			},
			expected: 9, // 3 * 3
		},
		{
			name: "medium severity SLO gap",
			gap: analysis.Gap{
				Type:     analysis.GapTypeSLO,
				Severity: analysis.SeverityMedium,
			},
			expected: 4, // 2 * 2
		},
		{
			name: "low severity initiative gap",
			gap: analysis.Gap{
				Type:     analysis.GapTypeInitiative,
				Severity: analysis.SeverityLow,
			},
			expected: 1, // 1 * 1
		},
		{
			name: "high severity SLO gap",
			gap: analysis.Gap{
				Type:     analysis.GapTypeSLO,
				Severity: analysis.SeverityHigh,
			},
			expected: 6, // 3 * 2
		},
		{
			name: "low severity maturity gap",
			gap: analysis.Gap{
				Type:     analysis.GapTypeMaturity,
				Severity: analysis.SeverityLow,
			},
			expected: 3, // 1 * 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateImpact(tt.gap)
			if got != tt.expected {
				t.Errorf("calculateImpact() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestSortGaps(t *testing.T) {
	gaps := []GapResult{
		{Type: "maturity", Severity: "low", Impact: 3},
		{Type: "slo", Severity: "high", Impact: 6},
		{Type: "initiative", Severity: "medium", Impact: 2},
		{Type: "maturity", Severity: "high", Impact: 9},
	}

	t.Run("sort by impact", func(t *testing.T) {
		g := make([]GapResult, len(gaps))
		copy(g, gaps)
		sortGaps(g, "impact")

		if g[0].Impact != 9 {
			t.Errorf("first gap impact = %d, want 9", g[0].Impact)
		}
		if g[len(g)-1].Impact != 2 {
			t.Errorf("last gap impact = %d, want 2", g[len(g)-1].Impact)
		}
	})

	t.Run("sort by severity", func(t *testing.T) {
		g := make([]GapResult, len(gaps))
		copy(g, gaps)
		sortGaps(g, "severity")

		if g[0].Severity != "high" {
			t.Errorf("first gap severity = %s, want high", g[0].Severity)
		}
		if g[len(g)-1].Severity != "low" {
			t.Errorf("last gap severity = %s, want low", g[len(g)-1].Severity)
		}
	})

	t.Run("sort by type", func(t *testing.T) {
		g := make([]GapResult, len(gaps))
		copy(g, gaps)
		sortGaps(g, "type")

		if g[0].Type != "initiative" {
			t.Errorf("first gap type = %s, want initiative", g[0].Type)
		}
	})
}

func TestSeverityOrder(t *testing.T) {
	tests := []struct {
		severity string
		expected int
	}{
		{"high", 0},
		{"medium", 1},
		{"low", 2},
		{"unknown", 3},
	}

	for _, tt := range tests {
		t.Run(tt.severity, func(t *testing.T) {
			got := severityOrder(tt.severity)
			if got != tt.expected {
				t.Errorf("severityOrder(%s) = %d, want %d", tt.severity, got, tt.expected)
			}
		})
	}
}

func TestGapsSummaryInitialization(t *testing.T) {
	output := &GapsOutput{
		Summary: GapsSummary{
			BySeverity:     make(map[string]int),
			ByType:         make(map[string]int),
			DocumentsCount: 0,
		},
		Gaps: make([]GapResult, 0),
	}

	if output.Summary.BySeverity == nil {
		t.Error("BySeverity map should be initialized")
	}
	if output.Summary.ByType == nil {
		t.Error("ByType map should be initialized")
	}
	if output.Gaps == nil {
		t.Error("Gaps slice should be initialized")
	}
}
