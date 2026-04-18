package prism

import "fmt"

// Service represents a deployable service or application owned by a team.
type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	// Ownership
	OwnerTeamID string `json:"ownerTeamId,omitempty"` // Team responsible for this service
	LayerID     string `json:"layerId,omitempty"`     // Primary layer (code, infra, runtime)

	// Metrics associated with this service
	MetricIDs []string `json:"metricIds,omitempty"`

	// Additional metadata
	Repository string `json:"repository,omitempty"` // Git repository URL
	Tier       string `json:"tier,omitempty"`       // Service tier (tier1, tier2, tier3)
}

// Validate validates a Service and returns validation errors.
func (s *Service) Validate(doc *PRISMDocument) ValidationErrors {
	var errs ValidationErrors

	if s.ID == "" {
		errs = append(errs, ValidationError{Field: "id", Message: "is required"})
	}

	if s.Name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "is required"})
	}

	// Validate layer reference
	if s.LayerID != "" {
		if err := ValidateLayer(s.LayerID); err != nil {
			errs = append(errs, ValidationError{Field: "layerId", Value: s.LayerID, Message: err.Error()})
		}
	}

	// Validate team reference
	if doc != nil && s.OwnerTeamID != "" {
		if doc.GetTeamByID(s.OwnerTeamID) == nil {
			errs = append(errs, ValidationError{
				Field:   "ownerTeamId",
				Value:   s.OwnerTeamID,
				Message: "references non-existent team ID",
			})
		}
	}

	// Validate metric references
	if doc != nil {
		for i, metricID := range s.MetricIDs {
			if doc.GetMetricByID(metricID) == nil {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("metricIds[%d]", i),
					Value:   metricID,
					Message: "references non-existent metric ID",
				})
			}
		}
	}

	return errs
}
