// Package ecosystem provides unified loading and querying across PRISM modules.
package ecosystem

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

// =============================================================================
// Scoring Dimensions (RMI-014)
// =============================================================================

// ScoringDimension represents a dimension used in initiative scoring.
type ScoringDimension string

const (
	// DimCustomerValue measures direct customer value and satisfaction impact.
	DimCustomerValue ScoringDimension = "customer_value"

	// DimMarketOpportunity measures market size and growth potential.
	DimMarketOpportunity ScoringDimension = "market_opportunity"

	// DimStrategicAlignment measures alignment with company strategy and pillars.
	DimStrategicAlignment ScoringDimension = "strategic_alignment"

	// DimEngineeringCost measures engineering effort (inverse: higher cost = lower score).
	DimEngineeringCost ScoringDimension = "engineering_cost"

	// DimRisk measures risk level (inverse: higher risk = lower score).
	DimRisk ScoringDimension = "risk"

	// DimAdoptionEvidence measures evidence of customer demand/adoption.
	DimAdoptionEvidence ScoringDimension = "adoption_evidence"

	// DimConfidence measures confidence in estimates.
	DimConfidence ScoringDimension = "confidence"
)

// AllScoringDimensions returns all scoring dimensions in display order.
func AllScoringDimensions() []ScoringDimension {
	return []ScoringDimension{
		DimCustomerValue,
		DimMarketOpportunity,
		DimStrategicAlignment,
		DimEngineeringCost,
		DimRisk,
		DimAdoptionEvidence,
		DimConfidence,
	}
}

// DimensionDescriptions provides human-readable descriptions for dimensions.
var DimensionDescriptions = map[ScoringDimension]string{
	DimCustomerValue:      "Customer Value: Direct customer satisfaction and value impact",
	DimMarketOpportunity:  "Market Opportunity: Market size, TAM/SAM/SOM potential",
	DimStrategicAlignment: "Strategic Alignment: Fit with company strategy and pillars",
	DimEngineeringCost:    "Engineering Cost: Development effort (inverted: lower cost = higher score)",
	DimRisk:               "Risk: Technical and execution risk (inverted: lower risk = higher score)",
	DimAdoptionEvidence:   "Adoption Evidence: Customer demand signals, votes, requests",
	DimConfidence:         "Confidence: Certainty in estimates and assumptions",
}

// IsInverseDimension returns true if higher raw values should result in lower scores.
func (d ScoringDimension) IsInverse() bool {
	return d == DimEngineeringCost || d == DimRisk
}

// =============================================================================
// Scoring Configuration (RMI-014)
// =============================================================================

// ScoringConfig defines the weights for each scoring dimension.
// Weights are configuration, not code.
type ScoringConfig struct {
	// Weights maps dimensions to their weights (0.0 to 1.0+).
	// Missing dimensions use DefaultWeight.
	Weights map[ScoringDimension]float64 `json:"weights"`

	// DefaultWeight is used for dimensions not in Weights.
	DefaultWeight float64 `json:"defaultWeight"`

	// Name is an optional name for this configuration.
	Name string `json:"name,omitempty"`

	// Description explains the scoring philosophy.
	Description string `json:"description,omitempty"`
}

// GetWeight returns the weight for a dimension, using DefaultWeight if not set.
func (c *ScoringConfig) GetWeight(dim ScoringDimension) float64 {
	if w, ok := c.Weights[dim]; ok {
		return w
	}
	return c.DefaultWeight
}

// DefaultScoringConfig returns a balanced default configuration.
func DefaultScoringConfig() *ScoringConfig {
	return &ScoringConfig{
		Name:        "balanced",
		Description: "Balanced weights across all dimensions",
		Weights: map[ScoringDimension]float64{
			DimCustomerValue:      1.0,
			DimMarketOpportunity:  0.8,
			DimStrategicAlignment: 1.0,
			DimEngineeringCost:    0.7,
			DimRisk:               0.6,
			DimAdoptionEvidence:   0.9,
			DimConfidence:         0.5,
		},
		DefaultWeight: 1.0,
	}
}

// CustomerFocusedConfig returns a config weighted toward customer value.
func CustomerFocusedConfig() *ScoringConfig {
	return &ScoringConfig{
		Name:        "customer-focused",
		Description: "Prioritizes customer value and adoption evidence",
		Weights: map[ScoringDimension]float64{
			DimCustomerValue:      1.5,
			DimMarketOpportunity:  0.6,
			DimStrategicAlignment: 0.8,
			DimEngineeringCost:    0.5,
			DimRisk:               0.5,
			DimAdoptionEvidence:   1.2,
			DimConfidence:         0.4,
		},
		DefaultWeight: 1.0,
	}
}

// GrowthFocusedConfig returns a config weighted toward market opportunity.
func GrowthFocusedConfig() *ScoringConfig {
	return &ScoringConfig{
		Name:        "growth-focused",
		Description: "Prioritizes market opportunity and strategic alignment",
		Weights: map[ScoringDimension]float64{
			DimCustomerValue:      0.7,
			DimMarketOpportunity:  1.5,
			DimStrategicAlignment: 1.2,
			DimEngineeringCost:    0.6,
			DimRisk:               0.8,
			DimAdoptionEvidence:   0.7,
			DimConfidence:         0.5,
		},
		DefaultWeight: 1.0,
	}
}

// LoadScoringConfig loads a scoring configuration from a JSON file.
func LoadScoringConfig(path string) (*ScoringConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading scoring config: %w", err)
	}

	var config ScoringConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parsing scoring config: %w", err)
	}

	// Initialize weights map if nil
	if config.Weights == nil {
		config.Weights = make(map[ScoringDimension]float64)
	}

	// Set default weight if not specified
	if config.DefaultWeight == 0 {
		config.DefaultWeight = 1.0
	}

	return &config, nil
}

// SaveScoringConfig saves a scoring configuration to a JSON file.
func SaveScoringConfig(config *ScoringConfig, path string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling scoring config: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("writing scoring config: %w", err)
	}

	return nil
}

// =============================================================================
// Explainable Scoring (RMI-015)
// =============================================================================

// DimensionScore represents a single dimension's contribution to the total score.
type DimensionScore struct {
	Dimension     ScoringDimension `json:"dimension"`
	RawValue      float64          `json:"rawValue"`      // Original value (0-100 scale)
	Weight        float64          `json:"weight"`        // Applied weight
	WeightedScore float64          `json:"weightedScore"` // RawValue * Weight (adjusted for inverse)
	Source        string           `json:"source"`        // Where the value came from
}

// InitiativeScore represents the complete score for an initiative.
type InitiativeScore struct {
	InitiativeID    string           `json:"initiativeId"`
	InitiativeName  string           `json:"initiativeName,omitempty"`
	TotalScore      float64          `json:"totalScore"`
	NormalizedScore float64          `json:"normalizedScore"` // 0-100 scale
	DimensionScores []DimensionScore `json:"dimensionScores"`
	Explanation     string           `json:"explanation"`
	Pillar          Pillar           `json:"pillar,omitempty"` // Primary pillar if assigned
}

// InitiativeInput provides the raw values for scoring an initiative.
type InitiativeInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	// Raw dimension values (0-100 scale)
	CustomerValue      float64 `json:"customerValue"`
	MarketOpportunity  float64 `json:"marketOpportunity"`
	StrategicAlignment float64 `json:"strategicAlignment"`
	EngineeringCost    float64 `json:"engineeringCost"` // Higher = more costly
	Risk               float64 `json:"risk"`            // Higher = more risky
	AdoptionEvidence   float64 `json:"adoptionEvidence"`
	Confidence         float64 `json:"confidence"`

	// Optional pillar assignment
	PrimaryPillar Pillar `json:"primaryPillar,omitempty"`

	// Sources for traceability
	Sources map[ScoringDimension]string `json:"sources,omitempty"`
}

// GetDimensionValue returns the raw value for a dimension.
func (i *InitiativeInput) GetDimensionValue(dim ScoringDimension) float64 {
	switch dim {
	case DimCustomerValue:
		return i.CustomerValue
	case DimMarketOpportunity:
		return i.MarketOpportunity
	case DimStrategicAlignment:
		return i.StrategicAlignment
	case DimEngineeringCost:
		return i.EngineeringCost
	case DimRisk:
		return i.Risk
	case DimAdoptionEvidence:
		return i.AdoptionEvidence
	case DimConfidence:
		return i.Confidence
	default:
		return 0
	}
}

// GetSource returns the source for a dimension's value.
func (i *InitiativeInput) GetSource(dim ScoringDimension) string {
	if i.Sources == nil {
		return "input"
	}
	if src, ok := i.Sources[dim]; ok {
		return src
	}
	return "input"
}

// ComputeInitiativeScore calculates the score for an initiative.
// Every score includes its contributing inputs for explainability.
func ComputeInitiativeScore(input *InitiativeInput, config *ScoringConfig) *InitiativeScore {
	if config == nil {
		config = DefaultScoringConfig()
	}

	score := &InitiativeScore{
		InitiativeID:    input.ID,
		InitiativeName:  input.Name,
		DimensionScores: make([]DimensionScore, 0, len(AllScoringDimensions())),
		Pillar:          input.PrimaryPillar,
	}

	var totalWeightedScore float64
	var totalWeight float64
	var explanationParts []string

	for _, dim := range AllScoringDimensions() {
		rawValue := input.GetDimensionValue(dim)
		weight := config.GetWeight(dim)

		// For inverse dimensions, higher raw value = lower score
		adjustedValue := rawValue
		if dim.IsInverse() {
			adjustedValue = 100 - rawValue // Invert: 100 becomes 0, 0 becomes 100
		}

		weightedScore := adjustedValue * weight

		dimScore := DimensionScore{
			Dimension:     dim,
			RawValue:      rawValue,
			Weight:        weight,
			WeightedScore: weightedScore,
			Source:        input.GetSource(dim),
		}
		score.DimensionScores = append(score.DimensionScores, dimScore)

		totalWeightedScore += weightedScore
		totalWeight += weight

		// Build explanation for significant contributors
		if weightedScore > 50 {
			explanationParts = append(explanationParts,
				fmt.Sprintf("%s (%.0f×%.1f=%.0f)", dim, rawValue, weight, weightedScore))
		}
	}

	score.TotalScore = totalWeightedScore

	// Normalize to 0-100 scale
	if totalWeight > 0 {
		score.NormalizedScore = totalWeightedScore / totalWeight
	}

	// Build explanation
	if len(explanationParts) > 0 {
		score.Explanation = "Key contributors: " + strings.Join(explanationParts, ", ")
	} else {
		score.Explanation = "No standout contributors"
	}

	return score
}

// =============================================================================
// Initiative Ranking (RMI-016)
// =============================================================================

// RankedInitiative combines an initiative with its score for ranking.
type RankedInitiative struct {
	Rank  int              `json:"rank"`
	Score *InitiativeScore `json:"score"`
}

// RankingResult contains ranked initiatives with metadata.
type RankingResult struct {
	Config      *ScoringConfig     `json:"config"`
	Pillar      Pillar             `json:"pillar,omitempty"` // Empty means all pillars
	Initiatives []RankedInitiative `json:"initiatives"`
	TotalCount  int                `json:"totalCount"`
}

// RankInitiatives scores and ranks initiatives by their total score.
func RankInitiatives(inputs []*InitiativeInput, config *ScoringConfig) *RankingResult {
	if config == nil {
		config = DefaultScoringConfig()
	}

	// Score all initiatives
	scores := make([]*InitiativeScore, 0, len(inputs))
	for _, input := range inputs {
		score := ComputeInitiativeScore(input, config)
		scores = append(scores, score)
	}

	// Sort by total score descending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].TotalScore > scores[j].TotalScore
	})

	// Build ranked result
	ranked := make([]RankedInitiative, len(scores))
	for i, score := range scores {
		ranked[i] = RankedInitiative{
			Rank:  i + 1,
			Score: score,
		}
	}

	return &RankingResult{
		Config:      config,
		Initiatives: ranked,
		TotalCount:  len(ranked),
	}
}

// RankInitiativesByPillar scores and ranks initiatives filtered by pillar.
func RankInitiativesByPillar(inputs []*InitiativeInput, config *ScoringConfig, pillar Pillar) *RankingResult {
	// Filter by pillar
	filtered := make([]*InitiativeInput, 0)
	for _, input := range inputs {
		if input.PrimaryPillar == pillar {
			filtered = append(filtered, input)
		}
	}

	result := RankInitiatives(filtered, config)
	result.Pillar = pillar
	return result
}

// TopN returns the top N initiatives from a ranking result.
func (r *RankingResult) TopN(n int) []RankedInitiative {
	if n <= 0 || n >= len(r.Initiatives) {
		return r.Initiatives
	}
	return r.Initiatives[:n]
}

// =============================================================================
// Evidence-Based Scoring
// =============================================================================

// ComputeAdoptionScore derives an adoption evidence score from evidence summary.
// Returns a 0-100 score based on signal count, votes, and customer impact.
func ComputeAdoptionScore(summary *EvidenceSummary) float64 {
	if summary == nil || !summary.HasEvidence() {
		return 0
	}

	// Weighted components (adjust these based on your needs)
	signalScore := float64(summary.SignalCount) * 10      // Each signal worth 10 points
	voteScore := float64(summary.TotalVotes) * 0.5        // Each vote worth 0.5 points
	customerScore := float64(summary.UniqueCustomers) * 5 // Each customer worth 5 points
	arrScore := summary.ARRInDollars() / 10000            // $10k ARR = 1 point

	// Combine and cap at 100
	total := signalScore + voteScore + customerScore + arrScore
	if total > 100 {
		total = 100
	}

	return total
}

// ComputeConfidenceScore derives a confidence score from evidence resolution.
// Returns a 0-100 score based on how much evidence was successfully resolved.
func ComputeConfidenceScore(summary *EvidenceSummary) float64 {
	if summary == nil || summary.TotalLinks == 0 {
		return 50 // Default confidence when no evidence
	}

	// Base confidence on resolution rate
	resolutionRate := float64(summary.ResolvedCount) / float64(summary.TotalLinks)

	// Scale to 0-100, with 50 as baseline
	return 50 + (resolutionRate * 50)
}

// EnrichInputFromEvidence adds evidence-derived scores to an initiative input.
func EnrichInputFromEvidence(input *InitiativeInput, summary *EvidenceSummary) {
	if summary == nil {
		return
	}

	// Only override if not already set
	if input.AdoptionEvidence == 0 {
		input.AdoptionEvidence = ComputeAdoptionScore(summary)
		if input.Sources == nil {
			input.Sources = make(map[ScoringDimension]string)
		}
		input.Sources[DimAdoptionEvidence] = "evidence-summary"
	}

	if input.Confidence == 0 {
		input.Confidence = ComputeConfidenceScore(summary)
		if input.Sources == nil {
			input.Sources = make(map[ScoringDimension]string)
		}
		input.Sources[DimConfidence] = "evidence-resolution"
	}
}
