package prism

import "testing"

func TestCalculateGoalProgress(t *testing.T) {
	doc := &PRISMDocument{
		Metrics: []Metric{
			{
				ID:         "metric-a",
				Name:       "Metric A",
				Domain:     DomainSecurity,
				Stage:      StageBuild,
				Category:   CategoryPrevention,
				MetricType: MetricTypeCoverage,
				Current:    95,
				SLO:        &SLO{Operator: SLOOperatorGTE, Value: 95},
			},
			{
				ID:         "metric-b",
				Name:       "Metric B",
				Domain:     DomainSecurity,
				Stage:      StageResponse,
				Category:   CategoryResponse,
				MetricType: MetricTypeLatency,
				Current:    7,
				SLO:        &SLO{Operator: SLOOperatorLTE, Value: 7},
			},
		},
		Goals: []Goal{
			{
				ID:   "goal-sec",
				Name: "Security Goal",
				MaturityModel: &GoalMaturityModel{
					Levels: []GoalMaturityLevel{
						{Level: 1, Name: "Reactive"},
						{Level: 2, Name: "Basic", RequiredSLOs: []SLORequirement{{MetricID: "metric-a"}}},
						{Level: 3, Name: "Defined", RequiredSLOs: []SLORequirement{{MetricID: "metric-a"}, {MetricID: "metric-b"}}},
					},
				},
			},
		},
		Phases: []Phase{
			{
				ID:        "phase-q1",
				Name:      "Q1 2026",
				StartDate: "2026-01-01",
				EndDate:   "2026-03-31",
				GoalTargets: []PhaseGoalTarget{
					{GoalID: "goal-sec", EnterLevel: 2, ExitLevel: 3},
				},
			},
		},
		Initiatives: []Initiative{
			{ID: "init-1", Name: "Init 1", Status: InitiativeStatusCompleted, GoalIDs: []string{"goal-sec"}, PhaseID: "phase-q1"},
			{ID: "init-2", Name: "Init 2", Status: InitiativeStatusInProgress, GoalIDs: []string{"goal-sec"}, PhaseID: "phase-q1"},
			{ID: "init-3", Name: "Init 3", Status: InitiativeStatusCompleted, GoalIDs: []string{"goal-sec"}, PhaseID: "phase-q1"},
		},
	}

	goal := doc.GetGoalByID("goal-sec")
	phase := doc.GetPhaseByID("phase-q1")

	progress := CalculateGoalProgress(goal, phase, doc)

	if progress == nil {
		t.Fatal("CalculateGoalProgress() returned nil")
	}

	if progress.GoalID != "goal-sec" {
		t.Errorf("GoalID = %q, want %q", progress.GoalID, "goal-sec")
	}

	if progress.EnterLevel != 2 {
		t.Errorf("EnterLevel = %d, want 2", progress.EnterLevel)
	}

	if progress.TargetLevel != 3 {
		t.Errorf("TargetLevel = %d, want 3", progress.TargetLevel)
	}

	if progress.InitiativesTotal != 3 {
		t.Errorf("InitiativesTotal = %d, want 3", progress.InitiativesTotal)
	}

	if progress.InitiativesCompleted != 2 {
		t.Errorf("InitiativesCompleted = %d, want 2", progress.InitiativesCompleted)
	}

	// 2/3 = 66.67%
	if progress.CompletionPercent < 66.66 || progress.CompletionPercent > 66.67 {
		t.Errorf("CompletionPercent = %v, want ~66.67", progress.CompletionPercent)
	}

	if progress.SLOsRequired != 2 {
		t.Errorf("SLOsRequired = %d, want 2", progress.SLOsRequired)
	}

	if progress.SLOsMet != 2 {
		t.Errorf("SLOsMet = %d, want 2", progress.SLOsMet)
	}
}

func TestCalculateInitiativeMetrics(t *testing.T) {
	doc := &PRISMDocument{
		Metrics: []Metric{
			{ID: "m1", Name: "M1", Domain: DomainSecurity, Stage: StageBuild, Category: CategoryPrevention, MetricType: MetricTypeCoverage},
		},
		Phases: []Phase{
			{ID: "phase-q1", Name: "Q1", StartDate: "2026-01-01", EndDate: "2026-03-31"},
		},
		Initiatives: []Initiative{
			{
				ID:                   "init-1",
				Name:                 "Init 1",
				Status:               InitiativeStatusCompleted,
				PhaseID:              "phase-q1",
				DevCompletionPercent: 100,
				DeploymentStatus:     &DeploymentStatus{Status: InitiativeStatusCompleted, TotalCustomers: 50, DeployedCustomers: 50},
			},
			{
				ID:                   "init-2",
				Name:                 "Init 2",
				Status:               InitiativeStatusCompleted,
				PhaseID:              "phase-q1",
				DevCompletionPercent: 100,
				DeploymentStatus:     &DeploymentStatus{Status: InitiativeStatusInProgress, TotalCustomers: 50, DeployedCustomers: 25},
			},
			{
				ID:                   "init-3",
				Name:                 "Init 3",
				Status:               InitiativeStatusInProgress,
				PhaseID:              "phase-q1",
				DevCompletionPercent: 50,
			},
			{
				ID:      "init-other",
				Name:    "Other Phase Init",
				Status:  InitiativeStatusCompleted,
				PhaseID: "phase-q2", // Different phase
			},
		},
	}

	phase := doc.GetPhaseByID("phase-q1")
	metrics := CalculateInitiativeMetrics(phase, doc)

	if metrics == nil {
		t.Fatal("CalculateInitiativeMetrics() returned nil")
	}

	if metrics.Total != 3 {
		t.Errorf("Total = %d, want 3", metrics.Total)
	}

	if metrics.Completed != 2 {
		t.Errorf("Completed = %d, want 2", metrics.Completed)
	}

	if metrics.Deployed != 1 {
		t.Errorf("Deployed = %d, want 1", metrics.Deployed)
	}

	// (100 + 50) / 2 = 75%
	expectedAdoption := 75.0
	if metrics.AvgAdoptionPercent != expectedAdoption {
		t.Errorf("AvgAdoptionPercent = %v, want %v", metrics.AvgAdoptionPercent, expectedAdoption)
	}
}

func TestCalculateSLOCompliance(t *testing.T) {
	doc := &PRISMDocument{
		Metrics: []Metric{
			{
				ID:         "metric-a",
				Name:       "Metric A",
				Domain:     DomainSecurity,
				Stage:      StageBuild,
				Category:   CategoryPrevention,
				MetricType: MetricTypeCoverage,
				Current:    95,
				SLO:        &SLO{Target: ">=95%", Operator: SLOOperatorGTE, Value: 95},
			},
			{
				ID:         "metric-b",
				Name:       "Metric B",
				Domain:     DomainSecurity,
				Stage:      StageResponse,
				Category:   CategoryResponse,
				MetricType: MetricTypeLatency,
				Current:    10,
				SLO:        &SLO{Target: "<=7 days", Operator: SLOOperatorLTE, Value: 7}, // Not met
			},
		},
		Goals: []Goal{
			{
				ID:   "goal-a",
				Name: "Goal A",
				MaturityModel: &GoalMaturityModel{
					Levels: []GoalMaturityLevel{
						{Level: 2, Name: "Basic", RequiredSLOs: []SLORequirement{{MetricID: "metric-a"}}},
						{Level: 3, Name: "Defined", RequiredSLOs: []SLORequirement{{MetricID: "metric-a"}, {MetricID: "metric-b"}}},
					},
				},
			},
		},
		Phases: []Phase{
			{ID: "phase-q1", Name: "Q1", StartDate: "2026-01-01", EndDate: "2026-03-31"},
		},
	}

	phase := doc.GetPhaseByID("phase-q1")
	compliance := CalculateSLOCompliance(phase, doc)

	if len(compliance) != 2 {
		t.Errorf("CalculateSLOCompliance() returned %d items, want 2", len(compliance))
	}

	// Find metric-a compliance
	var metricACompliance *SLOCompliance
	var metricBCompliance *SLOCompliance
	for i := range compliance {
		if compliance[i].MetricID == "metric-a" {
			metricACompliance = &compliance[i]
		}
		if compliance[i].MetricID == "metric-b" {
			metricBCompliance = &compliance[i]
		}
	}

	if metricACompliance == nil {
		t.Error("metric-a not found in compliance")
	} else {
		if !metricACompliance.IsMet {
			t.Error("metric-a should be met")
		}
		if metricACompliance.Current != 95 {
			t.Errorf("metric-a Current = %v, want 95", metricACompliance.Current)
		}
	}

	if metricBCompliance == nil {
		t.Error("metric-b not found in compliance")
	} else {
		if metricBCompliance.IsMet {
			t.Error("metric-b should not be met")
		}
		if metricBCompliance.Current != 10 {
			t.Errorf("metric-b Current = %v, want 10", metricBCompliance.Current)
		}
	}
}

func TestCalculatePhaseMetrics(t *testing.T) {
	doc := &PRISMDocument{
		Metrics: []Metric{
			{
				ID:         "metric-a",
				Name:       "Metric A",
				Domain:     DomainSecurity,
				Stage:      StageBuild,
				Category:   CategoryPrevention,
				MetricType: MetricTypeCoverage,
				Current:    95,
				SLO:        &SLO{Operator: SLOOperatorGTE, Value: 95},
			},
		},
		Goals: []Goal{
			{
				ID:   "goal-sec",
				Name: "Security Goal",
				MaturityModel: &GoalMaturityModel{
					Levels: []GoalMaturityLevel{
						{Level: 2, Name: "Basic", RequiredSLOs: []SLORequirement{{MetricID: "metric-a"}}},
					},
				},
			},
		},
		Phases: []Phase{
			{
				ID:          "phase-q1",
				Name:        "Q1",
				StartDate:   "2026-01-01",
				EndDate:     "2026-03-31",
				GoalTargets: []PhaseGoalTarget{{GoalID: "goal-sec", EnterLevel: 1, ExitLevel: 2}},
			},
		},
		Initiatives: []Initiative{
			{ID: "init-1", Name: "Init 1", Status: InitiativeStatusCompleted, GoalIDs: []string{"goal-sec"}, PhaseID: "phase-q1"},
		},
	}

	phase := doc.GetPhaseByID("phase-q1")
	metrics := CalculatePhaseMetrics(phase, doc)

	if metrics == nil {
		t.Fatal("CalculatePhaseMetrics() returned nil")
	}

	if metrics.PhaseID != "phase-q1" {
		t.Errorf("PhaseID = %q, want %q", metrics.PhaseID, "phase-q1")
	}

	if len(metrics.GoalProgress) != 1 {
		t.Errorf("GoalProgress has %d items, want 1", len(metrics.GoalProgress))
	}

	if metrics.InitiativeMetrics == nil {
		t.Error("InitiativeMetrics is nil")
	} else if metrics.InitiativeMetrics.Total != 1 {
		t.Errorf("InitiativeMetrics.Total = %d, want 1", metrics.InitiativeMetrics.Total)
	}
}

func TestCalculateGoalProgressNilInputs(t *testing.T) {
	doc := &PRISMDocument{}
	phase := &Phase{}
	goal := &Goal{}

	if CalculateGoalProgress(nil, phase, doc) != nil {
		t.Error("CalculateGoalProgress(nil goal) should return nil")
	}

	if CalculateGoalProgress(goal, nil, doc) != nil {
		t.Error("CalculateGoalProgress(nil phase) should return nil")
	}

	if CalculateGoalProgress(goal, phase, nil) != nil {
		t.Error("CalculateGoalProgress(nil doc) should return nil")
	}
}
