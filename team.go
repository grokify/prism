package prism

import "fmt"

// TeamType constants based on Team Topologies.
const (
	TeamTypeStreamAligned = "stream_aligned"
	TeamTypePlatform      = "platform"
	TeamTypeEnabling      = "enabling"
	TeamTypeOverlay       = "overlay"
)

// AllTeamTypes returns all valid team type values.
func AllTeamTypes() []string {
	return []string{
		TeamTypeStreamAligned,
		TeamTypePlatform,
		TeamTypeEnabling,
		TeamTypeOverlay,
	}
}

// Team represents a team in the organization following Team Topologies patterns.
type Team struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"` // stream_aligned, platform, enabling, overlay

	// Domain accountability (for overlay/enabling teams)
	Domain string `json:"domain,omitempty"` // security, operations, quality

	// Layer accountability (which layers this team is responsible for)
	LayerAccountability []string `json:"layerAccountability,omitempty"` // code, infra, runtime

	// Service ownership (for stream-aligned teams)
	ServiceIDs []string `json:"serviceIds,omitempty"`

	// Contact information
	Owner string `json:"owner,omitempty"`
	Slack string `json:"slack,omitempty"`
	Email string `json:"email,omitempty"`
}

// Validate validates a Team and returns validation errors.
func (t *Team) Validate(doc *PRISMDocument) ValidationErrors {
	var errs ValidationErrors

	if t.ID == "" {
		errs = append(errs, ValidationError{Field: "id", Message: "is required"})
	}

	if t.Name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "is required"})
	}

	if err := ValidateTeamType(t.Type); err != nil {
		errs = append(errs, ValidationError{Field: "type", Value: t.Type, Message: err.Error()})
	}

	// Validate domain if specified
	if t.Domain != "" {
		if err := ValidateDomain(t.Domain); err != nil {
			errs = append(errs, ValidationError{Field: "domain", Value: t.Domain, Message: err.Error()})
		}
	}

	// Validate layer accountability
	for i, layer := range t.LayerAccountability {
		if err := ValidateLayer(layer); err != nil {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("layerAccountability[%d]", i),
				Value:   layer,
				Message: err.Error(),
			})
		}
	}

	// Validate service references if document is provided
	if doc != nil {
		for i, serviceID := range t.ServiceIDs {
			if doc.GetServiceByID(serviceID) == nil {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("serviceIds[%d]", i),
					Value:   serviceID,
					Message: "references non-existent service ID",
				})
			}
		}
	}

	return errs
}
