package maturity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateXLSX(t *testing.T) {
	// Read the security maturity model
	specFile := filepath.Join("..", "maturity-models", "security.json")
	spec, err := ReadSpecFile(specFile)
	if err != nil {
		t.Fatalf("Failed to read spec file: %v", err)
	}

	// Verify spec was parsed correctly
	if spec == nil {
		t.Fatal("Spec is nil")
	}

	if len(spec.Domains) == 0 {
		t.Fatal("No domains in spec")
	}

	securityDomain, ok := spec.Domains["security"]
	if !ok {
		t.Fatal("Security domain not found")
	}

	if len(securityDomain.Levels) != 5 {
		t.Errorf("Expected 5 levels, got %d", len(securityDomain.Levels))
	}

	// Test level criteria
	level3, ok := securityDomain.GetLevel(3)
	if !ok {
		t.Fatal("Level 3 not found")
	}

	if len(level3.Criteria) == 0 {
		t.Error("Level 3 has no criteria")
	}

	if len(level3.Enablers) == 0 {
		t.Error("Level 3 has no enablers")
	}

	// Test criterion checking
	for _, c := range level3.Criteria {
		if c.ID == "" {
			t.Error("Criterion has empty ID")
		}
		if c.MetricName == "" {
			t.Error("Criterion has empty MetricName")
		}
	}

	// Generate XLSX
	gen := NewXLSXGenerator(spec)
	if err := gen.Generate(); err != nil {
		t.Fatalf("Failed to generate XLSX: %v", err)
	}

	// Save to temp file
	tmpFile := filepath.Join(os.TempDir(), "maturity-test.xlsx")
	if err := gen.SaveAs(tmpFile); err != nil {
		t.Fatalf("Failed to save XLSX: %v", err)
	}

	// Verify file was created
	info, err := os.Stat(tmpFile)
	if err != nil {
		t.Fatalf("Failed to stat output file: %v", err)
	}

	if info.Size() == 0 {
		t.Error("Output file is empty")
	}

	t.Logf("Generated XLSX file: %s (%d bytes)", tmpFile, info.Size())

	// Cleanup
	os.Remove(tmpFile)
}

func TestCriterionCheckMet(t *testing.T) {
	tests := []struct {
		name     string
		operator string
		target   float64
		current  float64
		expected bool
	}{
		{"gte met", OpGTE, 80, 85, true},
		{"gte not met", OpGTE, 80, 75, false},
		{"gte equal", OpGTE, 80, 80, true},
		{"lte met", OpLTE, 7, 5, true},
		{"lte not met", OpLTE, 7, 10, false},
		{"eq met", OpEQ, 0, 0, true},
		{"eq not met", OpEQ, 0, 5, false},
		{"gt met", OpGT, 50, 51, true},
		{"gt not met", OpGT, 50, 50, false},
		{"lt met", OpLT, 100, 99, true},
		{"lt not met", OpLT, 100, 100, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Criterion{
				Operator: tt.operator,
				Target:   tt.target,
			}
			result := c.CheckMet(tt.current)
			if result != tt.expected {
				t.Errorf("CheckMet(%v) = %v, expected %v", tt.current, result, tt.expected)
			}
		})
	}
}

func TestOperatorSymbol(t *testing.T) {
	tests := []struct {
		op       string
		expected string
	}{
		{OpGTE, ">="},
		{OpLTE, "<="},
		{OpGT, ">"},
		{OpLT, "<"},
		{OpEQ, "="},
		{"unknown", "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.op, func(t *testing.T) {
			result := OperatorSymbol(tt.op)
			if result != tt.expected {
				t.Errorf("OperatorSymbol(%q) = %q, expected %q", tt.op, result, tt.expected)
			}
		})
	}
}

func TestLevelProgress(t *testing.T) {
	level := Level{
		Level: 3,
		Criteria: []Criterion{
			{ID: "c1", Operator: OpGTE, Target: 80, Required: true},
			{ID: "c2", Operator: OpEQ, Target: 0, Required: true},
			{ID: "c3", Operator: OpGTE, Target: 100, Required: true},
		},
		Enablers: []Enabler{
			{ID: "e1"},
			{ID: "e2"},
		},
	}

	values := map[string]float64{
		"c1": 85, // Met
		"c2": 0,  // Met
		"c3": 80, // Not met
	}

	enablerStatus := map[string]string{
		"e1": StatusCompleted,
		"e2": StatusInProgress,
	}

	progress := level.CalculateLevelProgress(values, enablerStatus)

	if progress.CriteriaMet != 2 {
		t.Errorf("Expected 2 criteria met, got %d", progress.CriteriaMet)
	}

	if progress.CriteriaTotal != 3 {
		t.Errorf("Expected 3 total criteria, got %d", progress.CriteriaTotal)
	}

	if progress.EnablersDone != 1 {
		t.Errorf("Expected 1 enabler done, got %d", progress.EnablersDone)
	}

	if progress.ProgressPercent < 66 || progress.ProgressPercent > 67 {
		t.Errorf("Expected ~66.67%% progress, got %.2f%%", progress.ProgressPercent)
	}
}
