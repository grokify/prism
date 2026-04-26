package prism

import (
	"testing"
)

func TestCalculatePRISMScore_SkipEmptyCells(t *testing.T) {
	// Create a document with metrics only in operations/runtime
	doc := &PRISMDocument{
		Metrics: []Metric{
			{
				ID:       "test-metric",
				Name:     "Test Metric",
				Domain:   DomainOperations,
				Stage:    StageRuntime,
				Baseline: 0,
				Current:  90,
				Target:   100,
			},
		},
	}

	// Test with default config (skip empty cells)
	defaultConfig := DefaultScoreConfig()
	scoreSkip := doc.CalculatePRISMScore(defaultConfig, nil)

	// Test with legacy config (include empty cells)
	legacyConfig := LegacyScoreConfig()
	scoreLegacy := doc.CalculatePRISMScore(legacyConfig, nil)

	// Skip empty should give higher score since empty cells don't penalize
	if scoreSkip.Overall <= scoreLegacy.Overall {
		t.Errorf("Expected skip empty (%f) > legacy (%f)", scoreSkip.Overall, scoreLegacy.Overall)
	}

	// Skip empty should score based on actual data (90% progress * 60% perf weight = 54%)
	expectedMin := 0.50
	if scoreSkip.Overall < expectedMin {
		t.Errorf("Expected skip empty score >= %f, got %f", expectedMin, scoreSkip.Overall)
	}

	// Legacy should be much lower due to empty cells
	expectedMax := 0.20
	if scoreLegacy.Overall > expectedMax {
		t.Errorf("Expected legacy score <= %f, got %f", expectedMax, scoreLegacy.Overall)
	}
}

func TestCalculatePRISMScore_EmptyDocument(t *testing.T) {
	doc := &PRISMDocument{}

	// Should not panic on empty document
	score := doc.CalculatePRISMScore(nil, nil)

	if score.Overall != 0.0 {
		t.Errorf("Expected 0.0 for empty document, got %f", score.Overall)
	}
}

func TestCalculatePRISMScore_GoalMaturity(t *testing.T) {
	// Create a document with goals but no global maturity
	doc := &PRISMDocument{
		Metrics: []Metric{
			{
				ID:       "availability",
				Name:     "Availability",
				Domain:   DomainOperations,
				Stage:    StageRuntime,
				Baseline: 99.0,
				Current:  99.9,
				Target:   99.99,
				SLO: &SLO{
					Target:   ">=99.9%",
					Operator: SLOOperatorGTE,
					Value:    99.9,
				},
			},
		},
		Goals: []Goal{
			{
				ID:           "goal-reliability",
				Name:         "Reliability Goal",
				Owner:        "SRE Team",
				CurrentLevel: 3,
				TargetLevel:  4,
			},
		},
	}

	// Test without goal maturity
	configNoGoals := DefaultScoreConfig()
	configNoGoals.UseGoalMaturity = false
	scoreNoGoals := doc.CalculatePRISMScore(configNoGoals, nil)

	// Test with goal maturity
	configWithGoals := DefaultScoreConfig()
	configWithGoals.UseGoalMaturity = true
	scoreWithGoals := doc.CalculatePRISMScore(configWithGoals, nil)

	// With goals, maturity should be higher (goal at level 3 = 60%)
	if scoreWithGoals.MaturityAverage <= scoreNoGoals.MaturityAverage {
		t.Errorf("Expected goal maturity (%f) > no goal maturity (%f)",
			scoreWithGoals.MaturityAverage, scoreNoGoals.MaturityAverage)
	}
}

func TestLegacyScoreConfig(t *testing.T) {
	config := LegacyScoreConfig()

	if config.SkipEmptyCells {
		t.Error("LegacyScoreConfig should have SkipEmptyCells = false")
	}
}

func TestDefaultScoreConfig(t *testing.T) {
	config := DefaultScoreConfig()

	if !config.SkipEmptyCells {
		t.Error("DefaultScoreConfig should have SkipEmptyCells = true")
	}

	if config.UseGoalMaturity {
		t.Error("DefaultScoreConfig should have UseGoalMaturity = false")
	}
}

func TestGetScopedDomains(t *testing.T) {
	config := DefaultScoreConfig()

	// Default should return all domains
	domains := config.GetScopedDomains()
	if len(domains) == 0 {
		t.Error("GetScopedDomains should return all domains by default")
	}

	// Custom scope should return specified domains
	config.ScopedDomains = []string{DomainSecurity}
	domains = config.GetScopedDomains()
	if len(domains) != 1 || domains[0] != DomainSecurity {
		t.Errorf("Expected [%s], got %v", DomainSecurity, domains)
	}
}
