package ecosystem

import (
	"testing"
)

func TestContributionLevelWeight(t *testing.T) {
	tests := []struct {
		level    ContributionLevel
		expected int
	}{
		{ContributionHigh, 3},
		{ContributionMedium, 2},
		{ContributionLow, 1},
		{ContributionLevel("unknown"), 0},
	}

	for _, tt := range tests {
		t.Run(string(tt.level), func(t *testing.T) {
			if got := tt.level.Weight(); got != tt.expected {
				t.Errorf("ContributionLevel(%q).Weight() = %d, want %d", tt.level, got, tt.expected)
			}
		})
	}
}

func TestComputePortfolioRollup(t *testing.T) {
	eco := &Ecosystem{}

	contributions := InitiativePillars{
		"init-1": {
			{Pillar: PillarCSAT, Level: ContributionHigh, Rationale: "NPS improvement"},
			{Pillar: PillarSAMSOM, Level: ContributionLow, Rationale: "Minor market impact"},
		},
		"init-2": {
			{Pillar: PillarCSAT, Level: ContributionMedium},
		},
		"init-3": {
			{Pillar: PillarTAM, Level: ContributionHigh, Rationale: "New market entry"},
		},
	}

	rollup := eco.ComputePortfolioRollup(contributions)

	// Check total initiatives
	if rollup.TotalInitiatives != 3 {
		t.Errorf("TotalInitiatives = %d, want 3", rollup.TotalInitiatives)
	}

	// Check total weight: CSAT(3+2) + SAM-SOM(1) + TAM(3) = 9
	if rollup.TotalWeight != 9 {
		t.Errorf("TotalWeight = %d, want 9", rollup.TotalWeight)
	}

	// Check pillar rollups
	if len(rollup.ByPillar) != 3 {
		t.Fatalf("ByPillar length = %d, want 3", len(rollup.ByPillar))
	}

	// CSAT should be first (pillar order)
	csat := rollup.ByPillar[0]
	if csat.Pillar != PillarCSAT {
		t.Errorf("First pillar = %s, want CSAT", csat.Pillar)
	}
	if csat.Count != 2 {
		t.Errorf("CSAT Count = %d, want 2", csat.Count)
	}
	if csat.WeightedScore != 5 {
		t.Errorf("CSAT WeightedScore = %d, want 5", csat.WeightedScore)
	}

	// SAM-SOM
	samsom := rollup.ByPillar[1]
	if samsom.Pillar != PillarSAMSOM {
		t.Errorf("Second pillar = %s, want SAM-SOM", samsom.Pillar)
	}
	if samsom.WeightedScore != 1 {
		t.Errorf("SAM-SOM WeightedScore = %d, want 1", samsom.WeightedScore)
	}

	// TAM
	tam := rollup.ByPillar[2]
	if tam.Pillar != PillarTAM {
		t.Errorf("Third pillar = %s, want TAM", tam.Pillar)
	}
	if tam.WeightedScore != 3 {
		t.Errorf("TAM WeightedScore = %d, want 3", tam.WeightedScore)
	}

	// Check percentages sum to ~100
	totalPct := 0.0
	for _, pr := range rollup.ByPillar {
		totalPct += pr.Percentage
	}
	if totalPct < 99.9 || totalPct > 100.1 {
		t.Errorf("Total percentage = %.2f, want ~100", totalPct)
	}
}

func TestComputePortfolioRollupEmpty(t *testing.T) {
	eco := &Ecosystem{}
	contributions := InitiativePillars{}

	rollup := eco.ComputePortfolioRollup(contributions)

	if rollup.TotalInitiatives != 0 {
		t.Errorf("TotalInitiatives = %d, want 0", rollup.TotalInitiatives)
	}
	if rollup.TotalWeight != 0 {
		t.Errorf("TotalWeight = %d, want 0", rollup.TotalWeight)
	}

	// All pillars should still be present with zero values
	if len(rollup.ByPillar) != 3 {
		t.Errorf("ByPillar length = %d, want 3", len(rollup.ByPillar))
	}
	for _, pr := range rollup.ByPillar {
		if pr.Percentage != 0 {
			t.Errorf("Empty rollup should have 0%% for %s, got %.2f", pr.Pillar, pr.Percentage)
		}
	}
}

func TestComputePillarDetails(t *testing.T) {
	eco := &Ecosystem{}

	contributions := InitiativePillars{
		"init-1": {{Pillar: PillarCSAT, Level: ContributionHigh}},
		"init-2": {{Pillar: PillarCSAT, Level: ContributionMedium}},
		"init-3": {{Pillar: PillarCSAT, Level: ContributionLow}},
		"init-4": {{Pillar: PillarTAM, Level: ContributionHigh}},
	}

	names := map[string]string{
		"init-1": "Initiative One",
		"init-2": "Initiative Two",
		"init-3": "Initiative Three",
		"init-4": "Initiative Four",
	}

	details := eco.ComputePillarDetails(contributions, names, 2)

	if len(details) != 3 {
		t.Fatalf("Expected 3 pillar details, got %d", len(details))
	}

	// CSAT should have top 2 (high and medium, sorted by weight)
	csat := details[0]
	if csat.Pillar != PillarCSAT {
		t.Errorf("First detail pillar = %s, want CSAT", csat.Pillar)
	}
	if len(csat.TopInitiatives) != 2 {
		t.Errorf("CSAT TopInitiatives = %d, want 2 (topN limit)", len(csat.TopInitiatives))
	}
	if csat.TopInitiatives[0].Level != ContributionHigh {
		t.Errorf("First CSAT initiative should be high level, got %s", csat.TopInitiatives[0].Level)
	}
}

func TestAllPillars(t *testing.T) {
	pillars := AllPillars()
	if len(pillars) != 3 {
		t.Errorf("AllPillars() length = %d, want 3", len(pillars))
	}
	if pillars[0] != PillarCSAT {
		t.Errorf("First pillar = %s, want CSAT", pillars[0])
	}
	if pillars[1] != PillarSAMSOM {
		t.Errorf("Second pillar = %s, want SAM-SOM", pillars[1])
	}
	if pillars[2] != PillarTAM {
		t.Errorf("Third pillar = %s, want TAM", pillars[2])
	}
}
