package prism

import (
	"strings"
	"testing"
)

func TestGenerateSLOReport(t *testing.T) {
	doc := &PRISMDocument{
		Metadata: &Metadata{
			Name: "Test Document",
		},
		Metrics: []Metric{
			{
				ID:       "ops-availability",
				Name:     "Service Availability",
				Domain:   DomainOperations,
				Stage:    StageRuntime,
				Category: CategoryReliability,
				Unit:     "%",
			},
			{
				ID:       "ops-mttr",
				Name:     "Mean Time to Recovery",
				Domain:   DomainOperations,
				Stage:    StageResponse,
				Category: CategoryResponse,
				Unit:     "hours",
			},
			{
				ID:       "ops-deploy-freq",
				Name:     "Deployment Frequency",
				Domain:   DomainOperations,
				Stage:    StageBuild,
				Category: CategoryEfficiency,
				Unit:     "/day",
			},
		},
		Goals: []Goal{
			{
				ID:   "goal-reliability",
				Name: "High Reliability",
				MaturityModel: &GoalMaturityModel{
					Levels: []GoalMaturityLevel{
						{
							Level: 2,
							Name:  "Basic",
							MetricCriteria: []MetricCriterion{
								{MetricID: "ops-availability", Operator: SLOOperatorGTE, Value: 99.0},
							},
						},
						{
							Level: 3,
							Name:  "Defined",
							MetricCriteria: []MetricCriterion{
								{MetricID: "ops-availability", Operator: SLOOperatorGTE, Value: 99.5},
								{MetricID: "ops-mttr", Operator: SLOOperatorLTE, Value: 4},
							},
						},
						{
							Level: 4,
							Name:  "Managed",
							MetricCriteria: []MetricCriterion{
								{MetricID: "ops-availability", Operator: SLOOperatorGTE, Value: 99.9},
								{MetricID: "ops-mttr", Operator: SLOOperatorLTE, Value: 1},
							},
						},
					},
				},
			},
			{
				ID:   "goal-velocity",
				Name: "Delivery Velocity",
				MaturityModel: &GoalMaturityModel{
					Levels: []GoalMaturityLevel{
						{
							Level: 2,
							Name:  "Basic",
							MetricCriteria: []MetricCriterion{
								{MetricID: "ops-deploy-freq", Operator: SLOOperatorGTE, Value: 1},
							},
						},
						{
							Level: 3,
							Name:  "Defined",
							MetricCriteria: []MetricCriterion{
								{MetricID: "ops-deploy-freq", Operator: SLOOperatorGTE, Value: 2},
							},
						},
						{
							Level: 4,
							Name:  "Managed",
							MetricCriteria: []MetricCriterion{
								{MetricID: "ops-deploy-freq", Operator: SLOOperatorGTE, Value: 5},
							},
						},
					},
				},
			},
		},
	}

	report := doc.GenerateSLOReport()

	// Test report title
	if !strings.Contains(report.Title, "Test Document") {
		t.Errorf("Expected title to contain 'Test Document', got %s", report.Title)
	}

	// Test categories are sorted
	if len(report.Categories) != 3 {
		t.Errorf("Expected 3 categories, got %d", len(report.Categories))
	}

	expectedCategories := []string{CategoryEfficiency, CategoryReliability, CategoryResponse}
	for i, cat := range report.Categories {
		if cat.Category != expectedCategories[i] {
			t.Errorf("Category %d: expected %s, got %s", i, expectedCategories[i], cat.Category)
		}
	}

	// Test flattened entries
	// goal-reliability: L2(1) + L3(2) + L4(2) = 5
	// goal-velocity: L2(1) + L3(1) + L4(1) = 3
	// Total: 8
	if len(report.Entries) != 8 {
		t.Errorf("Expected 8 entries, got %d", len(report.Entries))
	}

	// Verify availability ladder shows increasing stringency
	var availEntries []SLOReportEntry
	for _, e := range report.Entries {
		if e.MetricID == "ops-availability" {
			availEntries = append(availEntries, e)
		}
	}

	if len(availEntries) != 3 {
		t.Fatalf("Expected 3 availability entries, got %d", len(availEntries))
	}

	// Level 2: >= 99.0, Level 3: >= 99.5, Level 4: >= 99.9
	expectedValues := []float64{99.0, 99.5, 99.9}
	for i, e := range availEntries {
		if e.Value != expectedValues[i] {
			t.Errorf("Availability level %d: expected value %.1f, got %.1f", e.Level, expectedValues[i], e.Value)
		}
	}
}

func TestSLOReportTableOutput(t *testing.T) {
	report := &SLOReport{
		Title: "Test Report",
		Entries: []SLOReportEntry{
			{
				Category:    CategoryReliability,
				MetricID:    "ops-availability",
				MetricName:  "Service Availability",
				Domain:      DomainOperations,
				Stage:       StageRuntime,
				Level:       2,
				LevelName:   "Basic",
				Operator:    SLOOperatorGTE,
				Value:       99.0,
				Requirement: ">=99.00%",
				GoalName:    "Reliability",
			},
		},
	}

	columns := report.TableColumns()
	if len(columns) != 9 {
		t.Errorf("Expected 9 columns, got %d", len(columns))
	}

	rows := report.TableRows()
	if len(rows) != 1 {
		t.Errorf("Expected 1 row, got %d", len(rows))
	}

	if rows[0][0] != CategoryReliability {
		t.Errorf("Expected category '%s', got '%s'", CategoryReliability, rows[0][0])
	}
}

func TestSLOReportMarkdownOutput(t *testing.T) {
	doc := &PRISMDocument{
		Metrics: []Metric{
			{
				ID:       "test-metric",
				Name:     "Test Metric",
				Domain:   DomainSecurity,
				Stage:    StageBuild,
				Category: CategoryPrevention,
				Unit:     "%",
			},
		},
		Goals: []Goal{
			{
				ID:   "test-goal",
				Name: "Test Goal",
				MaturityModel: &GoalMaturityModel{
					Levels: []GoalMaturityLevel{
						{
							Level: 2,
							Name:  "Basic",
							MetricCriteria: []MetricCriterion{
								{MetricID: "test-metric", Operator: SLOOperatorGTE, Value: 80},
							},
						},
					},
				},
			},
		},
	}

	report := doc.GenerateSLOReport()
	md := report.ToMarkdown()

	// Check for expected content
	if !strings.Contains(md, "# SLO Requirements") {
		t.Error("Markdown should contain title")
	}
	if !strings.Contains(md, "## Prevention") {
		t.Error("Markdown should contain category header")
	}
	if !strings.Contains(md, "Test Metric") {
		t.Error("Markdown should contain metric name")
	}
	if !strings.Contains(md, ">=80.00%") {
		t.Error("Markdown should contain requirement value")
	}
}

func TestSLOReportMarpOutput(t *testing.T) {
	doc := &PRISMDocument{
		Metrics: []Metric{
			{
				ID:       "test-metric",
				Name:     "Test Metric",
				Domain:   DomainSecurity,
				Stage:    StageBuild,
				Category: CategoryPrevention,
				Unit:     "%",
			},
		},
		Goals: []Goal{
			{
				ID:   "test-goal",
				Name: "Test Goal",
				MaturityModel: &GoalMaturityModel{
					Levels: []GoalMaturityLevel{
						{
							Level: 2,
							Name:  "Basic",
							MetricCriteria: []MetricCriterion{
								{MetricID: "test-metric", Operator: SLOOperatorGTE, Value: 80},
							},
						},
					},
				},
			},
		},
	}

	report := doc.GenerateSLOReport()
	marp := report.ToMarp()

	// Check for Marp front matter
	if !strings.Contains(marp, "marp: true") {
		t.Error("Marp output should contain front matter")
	}
	if !strings.Contains(marp, "---") {
		t.Error("Marp output should contain slide separators")
	}
	if !strings.Contains(marp, "Categories Overview") {
		t.Error("Marp output should contain overview slide")
	}
}

func TestSLOReportMatrixMarkdown(t *testing.T) {
	doc := &PRISMDocument{
		Metrics: []Metric{
			{
				ID:       "test-metric",
				Name:     "Test Metric",
				Domain:   DomainSecurity,
				Stage:    StageBuild,
				Category: CategoryPrevention,
				Unit:     "%",
			},
		},
		Goals: []Goal{
			{
				ID:   "test-goal",
				Name: "Test Goal",
				MaturityModel: &GoalMaturityModel{
					Levels: []GoalMaturityLevel{
						{Level: 2, Name: "Basic", MetricCriteria: []MetricCriterion{
							{MetricID: "test-metric", Operator: SLOOperatorGTE, Value: 50},
						}},
						{Level: 3, Name: "Defined", MetricCriteria: []MetricCriterion{
							{MetricID: "test-metric", Operator: SLOOperatorGTE, Value: 80},
						}},
						{Level: 4, Name: "Managed", MetricCriteria: []MetricCriterion{
							{MetricID: "test-metric", Operator: SLOOperatorGTE, Value: 95},
						}},
					},
				},
			},
		},
	}

	report := doc.GenerateSLOReport()
	matrix := report.ToMatrixMarkdown()

	// Check for matrix headers
	if !strings.Contains(matrix, "L1 Reactive") {
		t.Error("Matrix should contain L1 header")
	}
	if !strings.Contains(matrix, "L5 Optimizing") {
		t.Error("Matrix should contain L5 header")
	}

	// Check for "-" placeholder for missing levels
	if !strings.Contains(matrix, "| - |") {
		t.Error("Matrix should contain '-' for levels without requirements")
	}

	// Check for actual values
	if !strings.Contains(matrix, ">=50.00%") {
		t.Error("Matrix should contain L2 requirement")
	}
	if !strings.Contains(matrix, ">=95.00%") {
		t.Error("Matrix should contain L4 requirement")
	}
}

func TestFormatRequirement(t *testing.T) {
	tests := []struct {
		operator string
		value    float64
		unit     string
		expected string
	}{
		{SLOOperatorGTE, 99.9, "%", ">=99.90%"},
		{SLOOperatorLTE, 7, "days", "<=7.00 days"},
		{SLOOperatorEQ, 100, "%", "=100.00%"},
		{SLOOperatorGT, 5, "", ">5.00"},
		{SLOOperatorLT, 10, "ms", "<10.00 ms"},
	}

	for _, tc := range tests {
		result := formatRequirement(tc.operator, tc.value, tc.unit)
		if result != tc.expected {
			t.Errorf("formatRequirement(%s, %.2f, %s) = %s, expected %s",
				tc.operator, tc.value, tc.unit, result, tc.expected)
		}
	}
}

func TestOperatorSymbol(t *testing.T) {
	tests := []struct {
		op       string
		expected string
	}{
		{SLOOperatorGTE, ">="},
		{SLOOperatorLTE, "<="},
		{SLOOperatorGT, ">"},
		{SLOOperatorLT, "<"},
		{SLOOperatorEQ, "="},
		{"unknown", "unknown"},
	}

	for _, tc := range tests {
		result := operatorSymbol(tc.op)
		if result != tc.expected {
			t.Errorf("operatorSymbol(%s) = %s, expected %s", tc.op, result, tc.expected)
		}
	}
}
