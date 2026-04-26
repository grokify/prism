package analysis

import "fmt"

// GapType represents the type of gap identified.
type GapType string

const (
	GapTypeMaturity   GapType = "maturity"
	GapTypeSLO        GapType = "slo"
	GapTypeInitiative GapType = "initiative"
)

// Severity represents the severity of a gap.
type Severity string

const (
	SeverityHigh   Severity = "high"
	SeverityMedium Severity = "medium"
	SeverityLow    Severity = "low"
)

// Gap identifies a gap that needs to be addressed.
type Gap struct {
	Type        GapType  `json:"type"`
	GoalID      string   `json:"goalId,omitempty"`
	Description string   `json:"description"`
	Severity    Severity `json:"severity"`
	PhaseID     string   `json:"phaseId,omitempty"`
}

// IdentifyGaps analyzes the result and identifies gaps.
func IdentifyGaps(analysis *Result) []Gap {
	var gaps []Gap

	// Maturity gaps
	for _, ga := range analysis.Goals {
		if ga.Gap > 0 {
			severity := SeverityLow
			if ga.Gap >= 3 {
				severity = SeverityHigh
			} else if ga.Gap >= 2 {
				severity = SeverityMedium
			}

			gaps = append(gaps, Gap{
				Type:        GapTypeMaturity,
				GoalID:      ga.GoalID,
				Description: fmt.Sprintf("%s needs to progress from M%d to M%d (%d levels)", ga.GoalName, ga.CurrentLevel, ga.TargetLevel, ga.Gap),
				Severity:    severity,
			})
		}

		// SLO gaps for this goal
		for _, slo := range ga.SLOsRequired {
			if !slo.IsMet {
				gaps = append(gaps, Gap{
					Type:        GapTypeSLO,
					GoalID:      ga.GoalID,
					Description: fmt.Sprintf("SLO not met: %s (current: %.2f, target: %s) required for M%d", slo.MetricName, slo.Current, slo.Target, slo.Level),
					Severity:    SeverityMedium,
				})
			}
		}
	}

	// Initiative gaps (phases without enough initiatives)
	for _, pa := range analysis.Phases {
		totalSLOsNeeded := 0
		for _, gt := range pa.GoalTargets {
			totalSLOsNeeded += gt.SLOsNeeded
		}

		if totalSLOsNeeded > 0 && pa.Initiatives < totalSLOsNeeded {
			gaps = append(gaps, Gap{
				Type:        GapTypeInitiative,
				PhaseID:     pa.PhaseID,
				Description: fmt.Sprintf("%s has %d initiatives but needs to achieve %d SLOs across %d goal level progressions", pa.PhaseName, pa.Initiatives, totalSLOsNeeded, len(pa.GoalTargets)),
				Severity:    SeverityMedium,
			})
		}
	}

	return gaps
}

// FilterGapsByType returns gaps of a specific type.
func FilterGapsByType(gaps []Gap, gapType GapType) []Gap {
	var result []Gap
	for _, g := range gaps {
		if g.Type == gapType {
			result = append(result, g)
		}
	}
	return result
}

// FilterGapsBySeverity returns gaps of a specific severity.
func FilterGapsBySeverity(gaps []Gap, severity Severity) []Gap {
	var result []Gap
	for _, g := range gaps {
		if g.Severity == severity {
			result = append(result, g)
		}
	}
	return result
}

// CountGapsBySeverity returns counts of gaps by severity.
func CountGapsBySeverity(gaps []Gap) map[Severity]int {
	counts := make(map[Severity]int)
	for _, g := range gaps {
		counts[g.Severity]++
	}
	return counts
}
