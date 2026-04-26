package analysis

import (
	"testing"
)

func TestIdentifyGapsMaturity(t *testing.T) {
	result := &Result{
		Goals: []GoalAnalysis{
			{
				GoalID:       "goal-1",
				GoalName:     "Reliability",
				CurrentLevel: 2,
				TargetLevel:  4,
				Gap:          2,
			},
			{
				GoalID:       "goal-2",
				GoalName:     "Security",
				CurrentLevel: 1,
				TargetLevel:  4,
				Gap:          3,
			},
			{
				GoalID:       "goal-3",
				GoalName:     "Quality",
				CurrentLevel: 4,
				TargetLevel:  4,
				Gap:          0,
			},
		},
	}

	gaps := IdentifyGaps(result)

	// Should have 2 maturity gaps (goal-3 has no gap)
	maturityGaps := FilterGapsByType(gaps, GapTypeMaturity)
	if len(maturityGaps) != 2 {
		t.Errorf("len(maturityGaps) = %d, want 2", len(maturityGaps))
	}

	// Check severity
	highGaps := FilterGapsBySeverity(gaps, SeverityHigh)
	if len(highGaps) != 1 {
		t.Errorf("len(highGaps) = %d, want 1 (goal-2 with gap=3)", len(highGaps))
	}

	mediumGaps := FilterGapsBySeverity(gaps, SeverityMedium)
	if len(mediumGaps) != 1 {
		t.Errorf("len(mediumGaps) = %d, want 1 (goal-1 with gap=2)", len(mediumGaps))
	}
}

func TestIdentifyGapsSLO(t *testing.T) {
	result := &Result{
		Goals: []GoalAnalysis{
			{
				GoalID:   "goal-1",
				GoalName: "Reliability",
				SLOsRequired: []SLORequirement{
					{MetricID: "m1", MetricName: "Availability", IsMet: true},
					{MetricID: "m2", MetricName: "Latency", IsMet: false, Current: 250, Target: "<=200ms", Level: 3},
					{MetricID: "m3", MetricName: "MTTR", IsMet: false, Current: 2, Target: "<=1h", Level: 4},
				},
			},
		},
	}

	gaps := IdentifyGaps(result)

	sloGaps := FilterGapsByType(gaps, GapTypeSLO)
	if len(sloGaps) != 2 {
		t.Errorf("len(sloGaps) = %d, want 2", len(sloGaps))
	}

	// All SLO gaps should be medium severity
	for _, g := range sloGaps {
		if g.Severity != SeverityMedium {
			t.Errorf("SLO gap severity = %q, want %q", g.Severity, SeverityMedium)
		}
	}
}

func TestIdentifyGapsInitiative(t *testing.T) {
	result := &Result{
		Phases: []PhaseAnalysis{
			{
				PhaseID:     "phase-1",
				PhaseName:   "Q1 2024",
				Initiatives: 1,
				GoalTargets: []GoalTarget{
					{GoalID: "goal-1", SLOsNeeded: 3},
					{GoalID: "goal-2", SLOsNeeded: 2},
				},
			},
			{
				PhaseID:     "phase-2",
				PhaseName:   "Q2 2024",
				Initiatives: 5,
				GoalTargets: []GoalTarget{
					{GoalID: "goal-1", SLOsNeeded: 2},
				},
			},
		},
	}

	gaps := IdentifyGaps(result)

	initGaps := FilterGapsByType(gaps, GapTypeInitiative)
	// phase-1 has 1 initiative but needs 5 SLOs
	// phase-2 has 5 initiatives and needs 2 SLOs (no gap)
	if len(initGaps) != 1 {
		t.Errorf("len(initGaps) = %d, want 1", len(initGaps))
	}

	if len(initGaps) > 0 && initGaps[0].PhaseID != "phase-1" {
		t.Errorf("initGaps[0].PhaseID = %q, want %q", initGaps[0].PhaseID, "phase-1")
	}
}

func TestFilterGapsByType(t *testing.T) {
	gaps := []Gap{
		{Type: GapTypeMaturity, GoalID: "g1"},
		{Type: GapTypeSLO, GoalID: "g1"},
		{Type: GapTypeMaturity, GoalID: "g2"},
		{Type: GapTypeInitiative, PhaseID: "p1"},
	}

	tests := []struct {
		gapType GapType
		want    int
	}{
		{GapTypeMaturity, 2},
		{GapTypeSLO, 1},
		{GapTypeInitiative, 1},
	}

	for _, tt := range tests {
		t.Run(string(tt.gapType), func(t *testing.T) {
			got := FilterGapsByType(gaps, tt.gapType)
			if len(got) != tt.want {
				t.Errorf("FilterGapsByType(%q) returned %d gaps, want %d", tt.gapType, len(got), tt.want)
			}
		})
	}
}

func TestFilterGapsBySeverity(t *testing.T) {
	gaps := []Gap{
		{Type: GapTypeMaturity, Severity: SeverityHigh},
		{Type: GapTypeSLO, Severity: SeverityMedium},
		{Type: GapTypeMaturity, Severity: SeverityMedium},
		{Type: GapTypeInitiative, Severity: SeverityLow},
	}

	tests := []struct {
		severity Severity
		want     int
	}{
		{SeverityHigh, 1},
		{SeverityMedium, 2},
		{SeverityLow, 1},
	}

	for _, tt := range tests {
		t.Run(string(tt.severity), func(t *testing.T) {
			got := FilterGapsBySeverity(gaps, tt.severity)
			if len(got) != tt.want {
				t.Errorf("FilterGapsBySeverity(%q) returned %d gaps, want %d", tt.severity, len(got), tt.want)
			}
		})
	}
}

func TestCountGapsBySeverity(t *testing.T) {
	gaps := []Gap{
		{Severity: SeverityHigh},
		{Severity: SeverityMedium},
		{Severity: SeverityMedium},
		{Severity: SeverityLow},
		{Severity: SeverityLow},
		{Severity: SeverityLow},
	}

	counts := CountGapsBySeverity(gaps)

	if counts[SeverityHigh] != 1 {
		t.Errorf("counts[High] = %d, want 1", counts[SeverityHigh])
	}
	if counts[SeverityMedium] != 2 {
		t.Errorf("counts[Medium] = %d, want 2", counts[SeverityMedium])
	}
	if counts[SeverityLow] != 3 {
		t.Errorf("counts[Low] = %d, want 3", counts[SeverityLow])
	}
}

func TestCountGapsBySeverityEmpty(t *testing.T) {
	gaps := []Gap{}
	counts := CountGapsBySeverity(gaps)

	if len(counts) != 0 {
		t.Errorf("counts on empty gaps should be empty, got %v", counts)
	}
}
