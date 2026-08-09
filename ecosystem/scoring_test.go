package ecosystem

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultScoringConfig(t *testing.T) {
	config := DefaultScoringConfig()

	if config.Name != "balanced" {
		t.Errorf("expected name 'balanced', got %q", config.Name)
	}

	// Check all dimensions have weights
	for _, dim := range AllScoringDimensions() {
		weight := config.GetWeight(dim)
		if weight <= 0 {
			t.Errorf("dimension %s has non-positive weight: %f", dim, weight)
		}
	}
}

func TestGetWeight(t *testing.T) {
	config := &ScoringConfig{
		Weights: map[ScoringDimension]float64{
			DimCustomerValue: 1.5,
		},
		DefaultWeight: 0.5,
	}

	// Explicit weight
	if w := config.GetWeight(DimCustomerValue); w != 1.5 {
		t.Errorf("expected 1.5 for CustomerValue, got %f", w)
	}

	// Default weight
	if w := config.GetWeight(DimRisk); w != 0.5 {
		t.Errorf("expected 0.5 (default) for Risk, got %f", w)
	}
}

func TestLoadSaveScoringConfig(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "scoring.json")

	// Save config
	original := CustomerFocusedConfig()
	if err := SaveScoringConfig(original, configPath); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatal("config file was not created")
	}

	// Load config
	loaded, err := LoadScoringConfig(configPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if loaded.Name != original.Name {
		t.Errorf("expected name %q, got %q", original.Name, loaded.Name)
	}

	if loaded.GetWeight(DimCustomerValue) != original.GetWeight(DimCustomerValue) {
		t.Errorf("weight mismatch for CustomerValue")
	}
}

func TestComputeInitiativeScore(t *testing.T) {
	config := DefaultScoringConfig()

	input := &InitiativeInput{
		ID:                 "init-1",
		Name:               "Test Initiative",
		CustomerValue:      80,
		MarketOpportunity:  60,
		StrategicAlignment: 90,
		EngineeringCost:    30, // Low cost = good
		Risk:               20, // Low risk = good
		AdoptionEvidence:   70,
		Confidence:         85,
		PrimaryPillar:      PillarCSAT,
	}

	score := ComputeInitiativeScore(input, config)

	if score.InitiativeID != "init-1" {
		t.Errorf("expected ID 'init-1', got %q", score.InitiativeID)
	}

	if score.TotalScore <= 0 {
		t.Error("expected positive total score")
	}

	if score.NormalizedScore <= 0 || score.NormalizedScore > 100 {
		t.Errorf("normalized score out of range: %f", score.NormalizedScore)
	}

	if len(score.DimensionScores) != len(AllScoringDimensions()) {
		t.Errorf("expected %d dimension scores, got %d",
			len(AllScoringDimensions()), len(score.DimensionScores))
	}

	if score.Pillar != PillarCSAT {
		t.Errorf("expected pillar CSAT, got %s", score.Pillar)
	}

	if score.Explanation == "" {
		t.Error("expected non-empty explanation")
	}
}

func TestInverseDimensions(t *testing.T) {
	config := &ScoringConfig{
		Weights: map[ScoringDimension]float64{
			DimEngineeringCost: 1.0,
			DimRisk:            1.0,
		},
		DefaultWeight: 0,
	}

	// High cost and high risk should result in low scores
	input := &InitiativeInput{
		ID:              "high-cost",
		EngineeringCost: 90, // Very costly
		Risk:            80, // Very risky
	}

	score := ComputeInitiativeScore(input, config)

	// Find the dimension scores
	var costScore, riskScore float64
	for _, ds := range score.DimensionScores {
		if ds.Dimension == DimEngineeringCost {
			costScore = ds.WeightedScore
		}
		if ds.Dimension == DimRisk {
			riskScore = ds.WeightedScore
		}
	}

	// High raw values should result in low weighted scores due to inversion
	if costScore > 20 {
		t.Errorf("expected low cost score for high cost, got %f", costScore)
	}
	if riskScore > 30 {
		t.Errorf("expected low risk score for high risk, got %f", riskScore)
	}
}

func TestRankInitiatives(t *testing.T) {
	inputs := []*InitiativeInput{
		{ID: "low", Name: "Low Priority", CustomerValue: 20, MarketOpportunity: 20},
		{ID: "high", Name: "High Priority", CustomerValue: 90, MarketOpportunity: 85},
		{ID: "med", Name: "Medium Priority", CustomerValue: 50, MarketOpportunity: 50},
	}

	result := RankInitiatives(inputs, DefaultScoringConfig())

	if result.TotalCount != 3 {
		t.Errorf("expected 3 initiatives, got %d", result.TotalCount)
	}

	// High should be rank 1
	if result.Initiatives[0].Score.InitiativeID != "high" {
		t.Errorf("expected 'high' at rank 1, got %q", result.Initiatives[0].Score.InitiativeID)
	}

	// Low should be rank 3
	if result.Initiatives[2].Score.InitiativeID != "low" {
		t.Errorf("expected 'low' at rank 3, got %q", result.Initiatives[2].Score.InitiativeID)
	}

	// Verify ranks are sequential
	for i, ri := range result.Initiatives {
		if ri.Rank != i+1 {
			t.Errorf("expected rank %d, got %d", i+1, ri.Rank)
		}
	}
}

func TestRankInitiativesByPillar(t *testing.T) {
	inputs := []*InitiativeInput{
		{ID: "csat-1", PrimaryPillar: PillarCSAT, CustomerValue: 80},
		{ID: "csat-2", PrimaryPillar: PillarCSAT, CustomerValue: 60},
		{ID: "tam-1", PrimaryPillar: PillarTAM, CustomerValue: 90},
		{ID: "sam-1", PrimaryPillar: PillarSAMSOM, CustomerValue: 70},
	}

	result := RankInitiativesByPillar(inputs, DefaultScoringConfig(), PillarCSAT)

	if result.Pillar != PillarCSAT {
		t.Errorf("expected pillar CSAT, got %s", result.Pillar)
	}

	if result.TotalCount != 2 {
		t.Errorf("expected 2 CSAT initiatives, got %d", result.TotalCount)
	}

	// csat-1 should be ranked higher due to higher CustomerValue
	if result.Initiatives[0].Score.InitiativeID != "csat-1" {
		t.Errorf("expected 'csat-1' at rank 1, got %q", result.Initiatives[0].Score.InitiativeID)
	}
}

func TestTopN(t *testing.T) {
	result := &RankingResult{
		Initiatives: []RankedInitiative{
			{Rank: 1, Score: &InitiativeScore{InitiativeID: "a"}},
			{Rank: 2, Score: &InitiativeScore{InitiativeID: "b"}},
			{Rank: 3, Score: &InitiativeScore{InitiativeID: "c"}},
			{Rank: 4, Score: &InitiativeScore{InitiativeID: "d"}},
		},
	}

	top2 := result.TopN(2)
	if len(top2) != 2 {
		t.Errorf("expected 2, got %d", len(top2))
	}

	// TopN with n > len should return all
	topAll := result.TopN(10)
	if len(topAll) != 4 {
		t.Errorf("expected 4, got %d", len(topAll))
	}

	// TopN with n <= 0 should return all
	topZero := result.TopN(0)
	if len(topZero) != 4 {
		t.Errorf("expected 4 for n=0, got %d", len(topZero))
	}
}

func TestComputeAdoptionScore(t *testing.T) {
	tests := []struct {
		name     string
		summary  *EvidenceSummary
		minScore float64
		maxScore float64
	}{
		{
			name:     "nil summary",
			summary:  nil,
			minScore: 0,
			maxScore: 0,
		},
		{
			name:     "empty summary",
			summary:  &EvidenceSummary{},
			minScore: 0,
			maxScore: 0,
		},
		{
			name: "signals only",
			summary: &EvidenceSummary{
				TotalLinks:  5,
				SignalCount: 5,
			},
			minScore: 40,
			maxScore: 60,
		},
		{
			name: "high evidence",
			summary: &EvidenceSummary{
				TotalLinks:      10,
				SignalCount:     5,
				TotalVotes:      100,
				UniqueCustomers: 10,
				TotalARR:        500000, // $5000
			},
			minScore: 90,
			maxScore: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := ComputeAdoptionScore(tt.summary)
			if score < tt.minScore || score > tt.maxScore {
				t.Errorf("expected score in [%f, %f], got %f", tt.minScore, tt.maxScore, score)
			}
		})
	}
}

func TestComputeConfidenceScore(t *testing.T) {
	tests := []struct {
		name     string
		summary  *EvidenceSummary
		expected float64
	}{
		{
			name:     "nil summary",
			summary:  nil,
			expected: 50,
		},
		{
			name:     "no links",
			summary:  &EvidenceSummary{TotalLinks: 0},
			expected: 50,
		},
		{
			name: "all resolved",
			summary: &EvidenceSummary{
				TotalLinks:    10,
				ResolvedCount: 10,
			},
			expected: 100,
		},
		{
			name: "half resolved",
			summary: &EvidenceSummary{
				TotalLinks:    10,
				ResolvedCount: 5,
			},
			expected: 75,
		},
		{
			name: "none resolved",
			summary: &EvidenceSummary{
				TotalLinks:      10,
				ResolvedCount:   0,
				UnresolvedCount: 10,
			},
			expected: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := ComputeConfidenceScore(tt.summary)
			if score != tt.expected {
				t.Errorf("expected %f, got %f", tt.expected, score)
			}
		})
	}
}

func TestEnrichInputFromEvidence(t *testing.T) {
	input := &InitiativeInput{
		ID:   "test",
		Name: "Test Initiative",
	}

	summary := &EvidenceSummary{
		TotalLinks:      5,
		SignalCount:     3,
		ResolvedCount:   4,
		UnresolvedCount: 1,
		TotalVotes:      50,
		UniqueCustomers: 5,
	}

	EnrichInputFromEvidence(input, summary)

	if input.AdoptionEvidence == 0 {
		t.Error("expected non-zero AdoptionEvidence")
	}

	if input.Confidence == 0 {
		t.Error("expected non-zero Confidence")
	}

	if input.Sources[DimAdoptionEvidence] != "evidence-summary" {
		t.Errorf("expected source 'evidence-summary', got %q", input.Sources[DimAdoptionEvidence])
	}

	if input.Sources[DimConfidence] != "evidence-resolution" {
		t.Errorf("expected source 'evidence-resolution', got %q", input.Sources[DimConfidence])
	}
}

func TestAllScoringDimensions(t *testing.T) {
	dims := AllScoringDimensions()

	if len(dims) != 7 {
		t.Errorf("expected 7 dimensions, got %d", len(dims))
	}

	// Check all dimensions have descriptions
	for _, dim := range dims {
		if _, ok := DimensionDescriptions[dim]; !ok {
			t.Errorf("dimension %s has no description", dim)
		}
	}
}

func TestIsInverseDimension(t *testing.T) {
	inverseDims := []ScoringDimension{DimEngineeringCost, DimRisk}
	normalDims := []ScoringDimension{DimCustomerValue, DimMarketOpportunity, DimStrategicAlignment, DimAdoptionEvidence, DimConfidence}

	for _, dim := range inverseDims {
		if !dim.IsInverse() {
			t.Errorf("expected %s to be inverse", dim)
		}
	}

	for _, dim := range normalDims {
		if dim.IsInverse() {
			t.Errorf("expected %s to NOT be inverse", dim)
		}
	}
}
