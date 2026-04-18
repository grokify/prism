package prism

import "testing"

func TestMetricCriterionIsMet(t *testing.T) {
	tests := []struct {
		name     string
		operator string
		value    float64
		current  float64
		want     bool
	}{
		{"gte met exactly", SLOOperatorGTE, 95, 95, true},
		{"gte met above", SLOOperatorGTE, 95, 100, true},
		{"gte not met", SLOOperatorGTE, 95, 90, false},
		{"lte met exactly", SLOOperatorLTE, 7, 7, true},
		{"lte met below", SLOOperatorLTE, 7, 5, true},
		{"lte not met", SLOOperatorLTE, 7, 10, false},
		{"gt met", SLOOperatorGT, 95, 96, true},
		{"gt not met equal", SLOOperatorGT, 95, 95, false},
		{"gt not met below", SLOOperatorGT, 95, 90, false},
		{"lt met", SLOOperatorLT, 7, 5, true},
		{"lt not met equal", SLOOperatorLT, 7, 7, false},
		{"lt not met above", SLOOperatorLT, 7, 10, false},
		{"eq met", SLOOperatorEQ, 100, 100, true},
		{"eq not met", SLOOperatorEQ, 100, 99, false},
		{"invalid operator", "invalid", 100, 100, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &MetricCriterion{
				MetricID: "test-metric",
				Operator: tt.operator,
				Value:    tt.value,
			}
			got := mc.IsMet(tt.current)
			if got != tt.want {
				t.Errorf("IsMet(%v) = %v, want %v", tt.current, got, tt.want)
			}
		})
	}
}

func TestGoalMaturityModelGetLevel(t *testing.T) {
	gmm := &GoalMaturityModel{
		Levels: []GoalMaturityLevel{
			{Level: 1, Name: "Reactive"},
			{Level: 2, Name: "Basic"},
			{Level: 3, Name: "Defined"},
			{Level: 4, Name: "Managed"},
			{Level: 5, Name: "Optimizing"},
		},
	}

	tests := []struct {
		level    int
		wantName string
		wantNil  bool
	}{
		{1, "Reactive", false},
		{3, "Defined", false},
		{5, "Optimizing", false},
		{0, "", true},
		{6, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.wantName, func(t *testing.T) {
			got := gmm.GetLevel(tt.level)
			if tt.wantNil {
				if got != nil {
					t.Errorf("GetLevel(%d) = %v, want nil", tt.level, got)
				}
			} else {
				if got == nil {
					t.Errorf("GetLevel(%d) = nil, want %v", tt.level, tt.wantName)
				} else if got.Name != tt.wantName {
					t.Errorf("GetLevel(%d).Name = %v, want %v", tt.level, got.Name, tt.wantName)
				}
			}
		})
	}
}

func TestGoalCurrentMaturityLevel(t *testing.T) {
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
				SLO: &SLO{
					Operator: SLOOperatorGTE,
					Value:    95,
				},
			},
			{
				ID:         "metric-b",
				Name:       "Metric B",
				Domain:     DomainSecurity,
				Stage:      StageResponse,
				Category:   CategoryResponse,
				MetricType: MetricTypeLatency,
				Current:    7,
				SLO: &SLO{
					Operator: SLOOperatorLTE,
					Value:    7,
				},
			},
			{
				ID:         "metric-c",
				Name:       "Metric C",
				Domain:     DomainSecurity,
				Stage:      StageDesign,
				Category:   CategoryPrevention,
				MetricType: MetricTypeCoverage,
				Current:    50, // Not meeting level 3 requirement
				SLO: &SLO{
					Operator: SLOOperatorGTE,
					Value:    90,
				},
			},
		},
	}

	goal := &Goal{
		ID:   "goal-test",
		Name: "Test Goal",
		MaturityModel: &GoalMaturityModel{
			Levels: []GoalMaturityLevel{
				{
					Level:        1,
					Name:         "Reactive",
					RequiredSLOs: []SLORequirement{},
				},
				{
					Level: 2,
					Name:  "Basic",
					RequiredSLOs: []SLORequirement{
						{MetricID: "metric-a"},
					},
					MetricCriteria: []MetricCriterion{
						{MetricID: "metric-a", Operator: SLOOperatorGTE, Value: 50},
					},
				},
				{
					Level: 3,
					Name:  "Defined",
					RequiredSLOs: []SLORequirement{
						{MetricID: "metric-a"},
						{MetricID: "metric-b"},
					},
					MetricCriteria: []MetricCriterion{
						{MetricID: "metric-a", Operator: SLOOperatorGTE, Value: 95},
						{MetricID: "metric-b", Operator: SLOOperatorLTE, Value: 14},
					},
				},
				{
					Level: 4,
					Name:  "Managed",
					RequiredSLOs: []SLORequirement{
						{MetricID: "metric-a"},
						{MetricID: "metric-b"},
						{MetricID: "metric-c"},
					},
					MetricCriteria: []MetricCriterion{
						{MetricID: "metric-a", Operator: SLOOperatorEQ, Value: 100},
						{MetricID: "metric-b", Operator: SLOOperatorLTE, Value: 7},
						{MetricID: "metric-c", Operator: SLOOperatorGTE, Value: 90},
					},
				},
				{
					Level: 5,
					Name:  "Optimizing",
					RequiredSLOs: []SLORequirement{
						{MetricID: "metric-a"},
						{MetricID: "metric-b"},
						{MetricID: "metric-c"},
					},
					MetricCriteria: []MetricCriterion{
						{MetricID: "metric-a", Operator: SLOOperatorEQ, Value: 100},
						{MetricID: "metric-b", Operator: SLOOperatorLTE, Value: 3},
						{MetricID: "metric-c", Operator: SLOOperatorEQ, Value: 100},
					},
				},
			},
		},
	}

	// Current values: metric-a=95, metric-b=7, metric-c=50
	// Should be at level 3 (metric-c SLO not met for level 4)
	level := goal.CurrentMaturityLevel(doc)
	if level != 3 {
		t.Errorf("CurrentMaturityLevel() = %d, want 3", level)
	}
}

func TestGoalSLOsMetForLevel(t *testing.T) {
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
				SLO: &SLO{
					Operator: SLOOperatorGTE,
					Value:    95,
				},
			},
			{
				ID:         "metric-b",
				Name:       "Metric B",
				Domain:     DomainSecurity,
				Stage:      StageResponse,
				Category:   CategoryResponse,
				MetricType: MetricTypeLatency,
				Current:    7,
				SLO: &SLO{
					Operator: SLOOperatorLTE,
					Value:    7,
				},
			},
			{
				ID:         "metric-c",
				Name:       "Metric C",
				Domain:     DomainSecurity,
				Stage:      StageDesign,
				Category:   CategoryPrevention,
				MetricType: MetricTypeCoverage,
				Current:    50,
				SLO: &SLO{
					Operator: SLOOperatorGTE,
					Value:    90, // Not met
				},
			},
		},
	}

	goal := &Goal{
		ID:   "goal-test",
		Name: "Test Goal",
		MaturityModel: &GoalMaturityModel{
			Levels: []GoalMaturityLevel{
				{
					Level: 3,
					Name:  "Defined",
					RequiredSLOs: []SLORequirement{
						{MetricID: "metric-a"},
						{MetricID: "metric-b"},
						{MetricID: "metric-c"},
					},
				},
			},
		},
	}

	met, total := goal.SLOsMetForLevel(3, doc)
	if met != 2 {
		t.Errorf("SLOsMetForLevel() met = %d, want 2", met)
	}
	if total != 3 {
		t.Errorf("SLOsMetForLevel() total = %d, want 3", total)
	}
}

func TestGoalWithoutMaturityModel(t *testing.T) {
	doc := &PRISMDocument{}
	goal := &Goal{
		ID:   "goal-no-model",
		Name: "Goal without model",
	}

	level := goal.CurrentMaturityLevel(doc)
	if level != 1 {
		t.Errorf("CurrentMaturityLevel() without model = %d, want 1", level)
	}

	meets := goal.MeetsLevelRequirements(1, doc)
	if !meets {
		t.Error("MeetsLevelRequirements(1) without model should return true")
	}
}
