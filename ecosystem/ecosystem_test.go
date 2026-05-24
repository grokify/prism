package ecosystem

import (
	"testing"

	capability "github.com/grokify/prism-capability"
	maturity "github.com/grokify/prism-maturity"
)

func TestEcosystemStats(t *testing.T) {
	eco := &Ecosystem{
		CapabilityStacks: []*capability.CapabilityStack{
			{
				Metadata: capability.Metadata{Domain: "security"},
				Capabilities: []capability.Capability{
					{ID: "cap1", Status: "active"},
					{ID: "cap2", Status: "planned"},
				},
			},
		},
		PRISMDocuments: []*maturity.PRISMDocument{
			{
				Metrics: []maturity.Metric{
					{ID: "m1", Name: "Metric 1"},
					{ID: "m2", Name: "Metric 2"},
				},
			},
		},
	}

	stats := eco.Stats()

	if stats.CapabilityStacks != 1 {
		t.Errorf("expected 1 capability stack, got %d", stats.CapabilityStacks)
	}
	if stats.TotalCapabilities != 2 {
		t.Errorf("expected 2 capabilities, got %d", stats.TotalCapabilities)
	}
	if stats.PRISMDocuments != 1 {
		t.Errorf("expected 1 PRISM document, got %d", stats.PRISMDocuments)
	}
	if stats.TotalMetrics != 2 {
		t.Errorf("expected 2 metrics, got %d", stats.TotalMetrics)
	}
	if stats.ByDomain["security"] != 2 {
		t.Errorf("expected 2 capabilities in security domain, got %d", stats.ByDomain["security"])
	}
}

func TestGetCapabilityByID(t *testing.T) {
	eco := &Ecosystem{
		CapabilityStacks: []*capability.CapabilityStack{
			{
				Capabilities: []capability.Capability{
					{ID: "cap1", Name: "Capability 1"},
					{ID: "cap2", Name: "Capability 2"},
				},
			},
		},
	}

	cap := eco.GetCapabilityByID("cap1")
	if cap == nil {
		t.Fatal("expected to find capability cap1")
	}
	if cap.Name != "Capability 1" {
		t.Errorf("expected name 'Capability 1', got %q", cap.Name)
	}

	cap = eco.GetCapabilityByID("nonexistent")
	if cap != nil {
		t.Error("expected nil for nonexistent capability")
	}
}

func TestGetMetricByID(t *testing.T) {
	eco := &Ecosystem{
		PRISMDocuments: []*maturity.PRISMDocument{
			{
				Metrics: []maturity.Metric{
					{ID: "m1", Name: "Metric 1"},
				},
			},
			{
				Metrics: []maturity.Metric{
					{ID: "m2", Name: "Metric 2"},
				},
			},
		},
	}

	m := eco.GetMetricByID("m2")
	if m == nil {
		t.Fatal("expected to find metric m2")
	}
	if m.Name != "Metric 2" {
		t.Errorf("expected name 'Metric 2', got %q", m.Name)
	}
}

func TestCapabilityContext(t *testing.T) {
	eco := &Ecosystem{
		CapabilityStacks: []*capability.CapabilityStack{
			{
				Capabilities: []capability.Capability{
					{
						ID:   "cap1",
						Name: "Capability 1",
						PRISMRef: &capability.PRISMRef{
							SLIIDs: []string{"m1", "m2"},
						},
					},
				},
			},
		},
		PRISMDocuments: []*maturity.PRISMDocument{
			{
				Metrics: []maturity.Metric{
					{ID: "m1", Name: "Metric 1"},
					{ID: "m2", Name: "Metric 2"},
					{ID: "m3", Name: "Metric 3"},
				},
			},
		},
	}

	ctx := eco.GetCapabilityContext("cap1")
	if ctx == nil {
		t.Fatal("expected to find capability context")
	}
	if len(ctx.Metrics) != 2 {
		t.Errorf("expected 2 linked metrics, got %d", len(ctx.Metrics))
	}
}

func TestValidationErrors(t *testing.T) {
	var errs ValidationErrors
	if errs.HasErrors() {
		t.Error("empty ValidationErrors should not have errors")
	}

	errs = append(errs, ValidationError{Message: "test error"})
	if !errs.HasErrors() {
		t.Error("ValidationErrors with one error should have errors")
	}
	if errs.Error() != "1 validation errors" {
		t.Errorf("unexpected error string: %s", errs.Error())
	}
}

func TestLoadConfig(t *testing.T) {
	config := Config{
		Name: "test-ecosystem",
	}

	eco, err := Load(config)
	if err != nil {
		t.Fatalf("failed to load empty config: %v", err)
	}
	if eco.Config.Name != "test-ecosystem" {
		t.Errorf("expected name 'test-ecosystem', got %q", eco.Config.Name)
	}
}
