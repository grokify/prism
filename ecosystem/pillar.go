// Package ecosystem provides unified loading and querying across PRISM modules.
package ecosystem

import (
	"sort"
)

// Pillar represents a strategic investment pillar.
type Pillar string

const (
	// PillarCSAT represents customer satisfaction initiatives.
	// Focus: Existing customer retention, NPS improvement, support quality.
	PillarCSAT Pillar = "CSAT"

	// PillarSAMSOM represents Serviceable Addressable Market / Serviceable Obtainable Market initiatives.
	// Focus: Market expansion within current capabilities, competitive positioning.
	PillarSAMSOM Pillar = "SAM-SOM"

	// PillarTAM represents Total Addressable Market initiatives.
	// Focus: New market entry, platform expansion, strategic bets.
	PillarTAM Pillar = "TAM"
)

// AllPillars returns all pillar types in display order.
func AllPillars() []Pillar {
	return []Pillar{PillarCSAT, PillarSAMSOM, PillarTAM}
}

// ContributionLevel represents the degree of contribution to a pillar.
type ContributionLevel string

const (
	// ContributionHigh indicates primary focus on this pillar.
	ContributionHigh ContributionLevel = "high"

	// ContributionMedium indicates secondary contribution.
	ContributionMedium ContributionLevel = "medium"

	// ContributionLow indicates minor or indirect contribution.
	ContributionLow ContributionLevel = "low"
)

// Weight returns the numeric weight for scoring (High=3, Medium=2, Low=1).
func (c ContributionLevel) Weight() int {
	switch c {
	case ContributionHigh:
		return 3
	case ContributionMedium:
		return 2
	case ContributionLow:
		return 1
	default:
		return 0
	}
}

// PillarContribution describes how an initiative contributes to a pillar.
type PillarContribution struct {
	Pillar    Pillar            `json:"pillar"`
	Level     ContributionLevel `json:"level"`
	Rationale string            `json:"rationale,omitempty"`
}

// InitiativePillars maps initiative IDs to their pillar contributions.
// An initiative can contribute to multiple pillars at different levels.
type InitiativePillars map[string][]PillarContribution

// PillarRollup summarizes investment in a single pillar.
type PillarRollup struct {
	Pillar        Pillar  `json:"pillar"`
	Count         int     `json:"count"`         // Number of initiatives contributing
	WeightedScore int     `json:"weightedScore"` // Sum of contribution weights
	Percentage    float64 `json:"percentage"`    // Percentage of total weighted score
}

// PortfolioRollup provides portfolio-level investment distribution.
type PortfolioRollup struct {
	ByPillar         []PillarRollup `json:"byPillar"`
	TotalInitiatives int            `json:"totalInitiatives"`
	TotalWeight      int            `json:"totalWeight"`
}

// TopInitiative represents an initiative ranked by pillar contribution.
type TopInitiative struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Level     ContributionLevel `json:"level"`
	Rationale string            `json:"rationale,omitempty"`
}

// PillarDetail provides detailed information for a pillar view.
type PillarDetail struct {
	Pillar         Pillar          `json:"pillar"`
	Description    string          `json:"description"`
	Count          int             `json:"count"`
	Percentage     float64         `json:"percentage"`
	TopInitiatives []TopInitiative `json:"topInitiatives"`
}

// PillarDescriptions provides human-readable descriptions for each pillar.
var PillarDescriptions = map[Pillar]string{
	PillarCSAT:   "Customer Satisfaction: Retention, NPS, support quality",
	PillarSAMSOM: "Market Expansion: Competitive positioning, market share growth",
	PillarTAM:    "Total Market: New markets, platform expansion, strategic bets",
}

// ComputePortfolioRollup calculates investment distribution across pillars.
// The contributions parameter maps initiative IDs to their pillar contributions.
func (e *Ecosystem) ComputePortfolioRollup(contributions InitiativePillars) PortfolioRollup {
	// Track unique initiatives and weighted scores per pillar
	pillarScores := make(map[Pillar]int)
	pillarCounts := make(map[Pillar]int)
	seenInitiatives := make(map[string]bool)

	for initID, contribs := range contributions {
		seenInitiatives[initID] = true
		for _, c := range contribs {
			pillarScores[c.Pillar] += c.Level.Weight()
			pillarCounts[c.Pillar]++
		}
	}

	// Calculate total weight for percentage
	totalWeight := 0
	for _, score := range pillarScores {
		totalWeight += score
	}

	// Build rollup for all pillars (include zeros)
	rollups := make([]PillarRollup, 0, len(AllPillars()))
	for _, pillar := range AllPillars() {
		score := pillarScores[pillar]
		pct := 0.0
		if totalWeight > 0 {
			pct = float64(score) / float64(totalWeight) * 100
		}
		rollups = append(rollups, PillarRollup{
			Pillar:        pillar,
			Count:         pillarCounts[pillar],
			WeightedScore: score,
			Percentage:    pct,
		})
	}

	return PortfolioRollup{
		ByPillar:         rollups,
		TotalInitiatives: len(seenInitiatives),
		TotalWeight:      totalWeight,
	}
}

// ComputePillarDetails returns detailed pillar information including top initiatives.
// The initiativeNames parameter maps initiative IDs to display names.
func (e *Ecosystem) ComputePillarDetails(
	contributions InitiativePillars,
	initiativeNames map[string]string,
	topN int,
) []PillarDetail {
	rollup := e.ComputePortfolioRollup(contributions)

	// Group initiatives by pillar
	pillarInitiatives := make(map[Pillar][]TopInitiative)
	for initID, contribs := range contributions {
		name := initiativeNames[initID]
		if name == "" {
			name = initID
		}
		for _, c := range contribs {
			pillarInitiatives[c.Pillar] = append(pillarInitiatives[c.Pillar], TopInitiative{
				ID:        initID,
				Name:      name,
				Level:     c.Level,
				Rationale: c.Rationale,
			})
		}
	}

	// Sort initiatives by level (high first) and take top N
	for pillar := range pillarInitiatives {
		inits := pillarInitiatives[pillar]
		sort.Slice(inits, func(i, j int) bool {
			return inits[i].Level.Weight() > inits[j].Level.Weight()
		})
		if len(inits) > topN {
			pillarInitiatives[pillar] = inits[:topN]
		}
	}

	// Build details
	details := make([]PillarDetail, 0, len(rollup.ByPillar))
	for _, pr := range rollup.ByPillar {
		details = append(details, PillarDetail{
			Pillar:         pr.Pillar,
			Description:    PillarDescriptions[pr.Pillar],
			Count:          pr.Count,
			Percentage:     pr.Percentage,
			TopInitiatives: pillarInitiatives[pr.Pillar],
		})
	}

	return details
}
