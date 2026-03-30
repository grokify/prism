package prism

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// ValidationError represents a validation error with context.
type ValidationError struct {
	Field   string
	Value   string
	Message string
}

func (e *ValidationError) Error() string {
	if e.Value != "" {
		return fmt.Sprintf("%s: %s (value: %q)", e.Field, e.Message, e.Value)
	}
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors is a collection of validation errors.
type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return ""
	}
	var msgs []string
	for _, e := range ve {
		msgs = append(msgs, e.Error())
	}
	return strings.Join(msgs, "; ")
}

// HasErrors returns true if there are any validation errors.
func (ve ValidationErrors) HasErrors() bool {
	return len(ve) > 0
}

// ValidateDomain validates a domain value.
func ValidateDomain(domain string) error {
	if domain == "" {
		return errors.New("domain is required")
	}
	if !slices.Contains(AllDomains(), domain) {
		return fmt.Errorf("invalid domain %q, must be one of: %s", domain, strings.Join(AllDomains(), ", "))
	}
	return nil
}

// ValidateStage validates a stage value.
func ValidateStage(stage string) error {
	if stage == "" {
		return errors.New("stage is required")
	}
	if !slices.Contains(AllStages(), stage) {
		return fmt.Errorf("invalid stage %q, must be one of: %s", stage, strings.Join(AllStages(), ", "))
	}
	return nil
}

// ValidateCategory validates a category value.
func ValidateCategory(category string) error {
	if category == "" {
		return errors.New("category is required")
	}
	if !slices.Contains(AllCategories(), category) {
		return fmt.Errorf("invalid category %q, must be one of: %s", category, strings.Join(AllCategories(), ", "))
	}
	return nil
}

// ValidateMaturityLevel validates a maturity level value.
func ValidateMaturityLevel(level int) error {
	if level < MaturityLevel1 || level > MaturityLevel5 {
		return fmt.Errorf("invalid maturity level %d, must be between %d and %d", level, MaturityLevel1, MaturityLevel5)
	}
	return nil
}

// ValidateAwarenessState validates an awareness state value.
func ValidateAwarenessState(state string) error {
	if state == "" {
		return errors.New("awareness state is required")
	}
	if !slices.Contains(AllAwarenessStates(), state) {
		return fmt.Errorf("invalid awareness state %q, must be one of: %s", state, strings.Join(AllAwarenessStates(), ", "))
	}
	return nil
}

// ValidateFramework validates a framework value.
func ValidateFramework(framework string) error {
	if framework == "" {
		return errors.New("framework is required")
	}
	if !slices.Contains(AllFrameworks(), framework) {
		return fmt.Errorf("invalid framework %q, must be one of: %s", framework, strings.Join(AllFrameworks(), ", "))
	}
	return nil
}

// ValidateMetricType validates a metric type value.
func ValidateMetricType(metricType string) error {
	if metricType == "" {
		return errors.New("metric type is required")
	}
	if !slices.Contains(AllMetricTypes(), metricType) {
		return fmt.Errorf("invalid metric type %q, must be one of: %s", metricType, strings.Join(AllMetricTypes(), ", "))
	}
	return nil
}

// ValidateTrendDirection validates a trend direction value.
func ValidateTrendDirection(trend string) error {
	if trend == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllTrendDirections(), trend) {
		return fmt.Errorf("invalid trend direction %q, must be one of: %s", trend, strings.Join(AllTrendDirections(), ", "))
	}
	return nil
}

// ValidateStatus validates a status value.
func ValidateStatus(status string) error {
	if status == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllStatuses(), status) {
		return fmt.Errorf("invalid status %q, must be one of: %s", status, strings.Join(AllStatuses(), ", "))
	}
	return nil
}

// ValidateWindow validates an SLO window value.
func ValidateWindow(window string) error {
	if window == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllWindows(), window) {
		return fmt.Errorf("invalid window %q, must be one of: %s", window, strings.Join(AllWindows(), ", "))
	}
	return nil
}

// Validate validates a Metric and returns validation errors.
func (m *Metric) Validate() ValidationErrors {
	var errs ValidationErrors

	if m.Name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "is required"})
	}

	if err := ValidateDomain(m.Domain); err != nil {
		errs = append(errs, ValidationError{Field: "domain", Value: m.Domain, Message: err.Error()})
	}

	if err := ValidateStage(m.Stage); err != nil {
		errs = append(errs, ValidationError{Field: "stage", Value: m.Stage, Message: err.Error()})
	}

	if err := ValidateCategory(m.Category); err != nil {
		errs = append(errs, ValidationError{Field: "category", Value: m.Category, Message: err.Error()})
	}

	if err := ValidateMetricType(m.MetricType); err != nil {
		errs = append(errs, ValidationError{Field: "metricType", Value: m.MetricType, Message: err.Error()})
	}

	if err := ValidateTrendDirection(m.TrendDirection); err != nil {
		errs = append(errs, ValidationError{Field: "trendDirection", Value: m.TrendDirection, Message: err.Error()})
	}

	if err := ValidateStatus(m.Status); err != nil {
		errs = append(errs, ValidationError{Field: "status", Value: m.Status, Message: err.Error()})
	}

	// Validate SLO window if present
	if m.SLO != nil {
		if err := ValidateWindow(m.SLO.Window); err != nil {
			errs = append(errs, ValidationError{Field: "slo.window", Value: m.SLO.Window, Message: err.Error()})
		}
	}

	// Validate framework mappings
	for i, fm := range m.FrameworkMappings {
		if err := ValidateFramework(fm.Framework); err != nil {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("frameworkMappings[%d].framework", i),
				Value:   fm.Framework,
				Message: err.Error(),
			})
		}
		if fm.Reference == "" {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("frameworkMappings[%d].reference", i),
				Message: "is required",
			})
		}
	}

	return errs
}

// Validate validates the entire PRISMDocument.
func (doc *PRISMDocument) Validate() ValidationErrors {
	var errs ValidationErrors

	if len(doc.Metrics) == 0 {
		errs = append(errs, ValidationError{Field: "metrics", Message: "at least one metric is required"})
	}

	// Validate each metric
	for i, m := range doc.Metrics {
		metricErrs := m.Validate()
		for _, e := range metricErrs {
			e.Field = fmt.Sprintf("metrics[%d].%s", i, e.Field)
			errs = append(errs, e)
		}
	}

	// Validate maturity model if present
	if doc.Maturity != nil {
		maturityErrs := doc.Maturity.Validate()
		for _, e := range maturityErrs {
			e.Field = "maturity." + e.Field
			errs = append(errs, e)
		}
	}

	// Check for duplicate metric IDs
	seenIDs := make(map[string]int)
	for i, m := range doc.Metrics {
		if m.ID != "" {
			if prevIdx, exists := seenIDs[m.ID]; exists {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("metrics[%d].id", i),
					Value:   m.ID,
					Message: fmt.Sprintf("duplicate ID, also used at metrics[%d]", prevIdx),
				})
			}
			seenIDs[m.ID] = i
		}
	}

	// Validate OKR metric references
	for i, okr := range doc.OKRs {
		for j, metricID := range okr.MetricIDs {
			if doc.GetMetricByID(metricID) == nil {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("okrs[%d].metricIds[%d]", i, j),
					Value:   metricID,
					Message: "references non-existent metric ID",
				})
			}
		}
	}

	// Validate initiative metric references
	for i, init := range doc.Initiatives {
		for j, metricID := range init.MetricIDs {
			if doc.GetMetricByID(metricID) == nil {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("initiatives[%d].metricIds[%d]", i, j),
					Value:   metricID,
					Message: "references non-existent metric ID",
				})
			}
		}
	}

	return errs
}
