package prism

// LayerDef defines an ownership layer in the stack (code, infra, runtime).
// Layers represent where metrics are measured and who is accountable.
type LayerDef struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	Weight      float64      `json:"weight,omitempty"`
	Signals     GoldenSignal `json:"signals,omitempty"`
}

// GoldenSignal defines the golden signals for a layer.
// Based on Google SRE's four golden signals: latency, traffic, errors, saturation.
type GoldenSignal struct {
	Latency    string `json:"latency,omitempty"`    // Response time metric ID
	Traffic    string `json:"traffic,omitempty"`    // Throughput metric ID
	Errors     string `json:"errors,omitempty"`     // Error rate metric ID
	Saturation string `json:"saturation,omitempty"` // Resource utilization metric ID
}

// DefaultLayers returns the default layer definitions.
func DefaultLayers() []LayerDef {
	return []LayerDef{
		{
			ID:          LayerCode,
			Name:        "Code",
			Description: "Application code, libraries, and dependencies",
		},
		{
			ID:          LayerInfra,
			Name:        "Infrastructure",
			Description: "Cloud resources, networking, and platform services",
		},
		{
			ID:          LayerRuntime,
			Name:        "Runtime",
			Description: "Running services, containers, and workloads",
		},
	}
}

// Validate validates a LayerDef and returns validation errors.
func (l *LayerDef) Validate() ValidationErrors {
	var errs ValidationErrors

	if l.ID == "" {
		errs = append(errs, ValidationError{Field: "id", Message: "is required"})
	}

	if l.Name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "is required"})
	}

	return errs
}
