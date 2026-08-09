// Package ecosystem provides unified loading and querying across PRISM modules.
package ecosystem

import (
	"errors"
	"fmt"
	"strings"
)

// =============================================================================
// Evidence Types (RMI-011)
// =============================================================================

// EvidenceType represents the type of evidence link.
type EvidenceType string

const (
	// SignalEvidence links to a canonical signal from OmniSignal.
	SignalEvidence EvidenceType = "signal"
	// CustomerEvidence links to a specific customer request or feedback.
	CustomerEvidence EvidenceType = "customer"
	// MarketEvidence links to market research or competitive analysis.
	MarketEvidence EvidenceType = "market"
	// IdeaEvidence links to an idea from a product management tool.
	IdeaEvidence EvidenceType = "idea"
)

// AllEvidenceTypes returns all valid evidence types.
func AllEvidenceTypes() []EvidenceType {
	return []EvidenceType{SignalEvidence, CustomerEvidence, MarketEvidence, IdeaEvidence}
}

// IsValidEvidenceType returns true if the type is valid.
func IsValidEvidenceType(t EvidenceType) bool {
	switch t {
	case SignalEvidence, CustomerEvidence, MarketEvidence, IdeaEvidence:
		return true
	default:
		return false
	}
}

// EvidenceLink represents a typed reference to external evidence.
// Format: "type:id" (e.g., "signal:oauth-idp-compatibility")
type EvidenceLink struct {
	Type EvidenceType `json:"type"`
	ID   string       `json:"id"`
	raw  string
}

// String returns the canonical string representation.
func (e EvidenceLink) String() string {
	if e.raw != "" {
		return e.raw
	}
	return fmt.Sprintf("%s:%s", e.Type, e.ID)
}

// ParseEvidenceLink parses a typed evidence link string.
// Format: "type:id" where type is one of: signal, customer, market, idea
func ParseEvidenceLink(s string) (EvidenceLink, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return EvidenceLink{}, errors.New("empty evidence link")
	}

	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return EvidenceLink{}, fmt.Errorf("invalid evidence link format: %q (expected type:id)", s)
	}

	evidenceType := EvidenceType(parts[0])
	if !IsValidEvidenceType(evidenceType) {
		return EvidenceLink{}, fmt.Errorf("invalid evidence type: %q (valid types: signal, customer, market, idea)", parts[0])
	}

	id := strings.TrimSpace(parts[1])
	if id == "" {
		return EvidenceLink{}, fmt.Errorf("empty evidence ID in: %q", s)
	}

	return EvidenceLink{
		Type: evidenceType,
		ID:   id,
		raw:  s,
	}, nil
}

// ParseEvidenceLinks parses multiple evidence link strings.
// Returns successfully parsed links and any parse errors.
func ParseEvidenceLinks(links []string) ([]EvidenceLink, []error) {
	var parsed []EvidenceLink
	var errs []error

	for _, s := range links {
		link, err := ParseEvidenceLink(s)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		parsed = append(parsed, link)
	}

	return parsed, errs
}

// =============================================================================
// Evidence Resolver (RMI-012)
// =============================================================================

// ResolvedEvidence contains the resolved data for an evidence link.
type ResolvedEvidence struct {
	Link EvidenceLink `json:"link"`

	// Signal data (when Type == SignalEvidence)
	SignalName       string  `json:"signalName,omitempty"`
	FrustrationScore float64 `json:"frustrationScore,omitempty"` // 0-100 scale

	// Customer impact
	CustomerCount int      `json:"customerCount,omitempty"`
	CustomerNames []string `json:"customerNames,omitempty"`
	TotalARR      int64    `json:"totalArr,omitempty"` // in cents

	// Idea aggregation
	IdeaCount  int `json:"ideaCount,omitempty"`
	TotalVotes int `json:"totalVotes,omitempty"`

	// Source tracking
	Sources []string `json:"sources,omitempty"`

	// Raw data from the external system (for custom processing)
	RawData map[string]any `json:"rawData,omitempty"`
}

// EvidenceResolver resolves evidence links against external systems.
// Implementations should handle their specific data sources (OmniSignal, etc.)
type EvidenceResolver interface {
	// Resolve fetches the evidence data for a single link.
	// Returns nil, nil if the link is valid but no data is found.
	Resolve(link EvidenceLink) (*ResolvedEvidence, error)

	// ValidateLinks checks that all links can be resolved.
	// Returns validation errors for links that cannot be resolved.
	ValidateLinks(links []EvidenceLink) []ValidationError

	// ResolveAll fetches evidence data for multiple links.
	// Returns resolved evidence for each link (nil entries for not found).
	ResolveAll(links []EvidenceLink) ([]*ResolvedEvidence, error)
}

// NoOpResolver is an EvidenceResolver that always returns nil.
// Use when OmniSignal or other evidence sources are not available.
type NoOpResolver struct{}

// Resolve always returns nil, nil.
func (r *NoOpResolver) Resolve(_ EvidenceLink) (*ResolvedEvidence, error) {
	return nil, nil
}

// ValidateLinks always returns no errors (all links considered valid).
func (r *NoOpResolver) ValidateLinks(_ []EvidenceLink) []ValidationError {
	return nil
}

// ResolveAll returns nil for all links.
func (r *NoOpResolver) ResolveAll(links []EvidenceLink) ([]*ResolvedEvidence, error) {
	return make([]*ResolvedEvidence, len(links)), nil
}

// =============================================================================
// Evidence Summary (RMI-013)
// =============================================================================

// EvidenceSummary aggregates evidence data across multiple links.
type EvidenceSummary struct {
	// Link counts by type
	SignalCount   int `json:"signalCount"`
	CustomerCount int `json:"customerCount"`
	MarketCount   int `json:"marketCount"`
	IdeaCount     int `json:"ideaCount"`
	TotalLinks    int `json:"totalLinks"`

	// Aggregated metrics
	TotalFrustrationScore float64  `json:"totalFrustrationScore"` // Sum of frustration scores
	AvgFrustrationScore   float64  `json:"avgFrustrationScore"`   // Average frustration score
	TotalARR              int64    `json:"totalArr"`              // Total ARR in cents
	UniqueCustomers       int      `json:"uniqueCustomers"`       // Deduplicated customer count
	CustomerNames         []string `json:"customerNames"`         // Deduplicated customer names
	TotalVotes            int      `json:"totalVotes"`            // Total idea votes

	// Resolution status
	ResolvedCount   int `json:"resolvedCount"`
	UnresolvedCount int `json:"unresolvedCount"`
}

// ComputeEvidenceSummary aggregates resolved evidence into a summary.
func ComputeEvidenceSummary(links []EvidenceLink, resolved []*ResolvedEvidence) *EvidenceSummary {
	summary := &EvidenceSummary{
		TotalLinks:    len(links),
		CustomerNames: []string{},
	}

	// Count links by type
	for _, link := range links {
		switch link.Type {
		case SignalEvidence:
			summary.SignalCount++
		case CustomerEvidence:
			summary.CustomerCount++
		case MarketEvidence:
			summary.MarketCount++
		case IdeaEvidence:
			summary.IdeaCount++
		}
	}

	// Aggregate resolved data
	customerSet := make(map[string]bool)
	frustrationCount := 0

	for _, r := range resolved {
		if r == nil {
			summary.UnresolvedCount++
			continue
		}
		summary.ResolvedCount++

		// Frustration score
		if r.FrustrationScore > 0 {
			summary.TotalFrustrationScore += r.FrustrationScore
			frustrationCount++
		}

		// ARR
		summary.TotalARR += r.TotalARR

		// Votes
		summary.TotalVotes += r.TotalVotes

		// Deduplicate customers
		for _, name := range r.CustomerNames {
			if !customerSet[name] {
				customerSet[name] = true
				summary.CustomerNames = append(summary.CustomerNames, name)
			}
		}
	}

	summary.UniqueCustomers = len(summary.CustomerNames)

	// Calculate average frustration
	if frustrationCount > 0 {
		summary.AvgFrustrationScore = summary.TotalFrustrationScore / float64(frustrationCount)
	}

	return summary
}

// ARRInDollars returns the total ARR in dollars.
func (s *EvidenceSummary) ARRInDollars() float64 {
	return float64(s.TotalARR) / 100
}

// HasEvidence returns true if there's any evidence data.
func (s *EvidenceSummary) HasEvidence() bool {
	return s.TotalLinks > 0
}

// HasResolvedEvidence returns true if any evidence was resolved.
func (s *EvidenceSummary) HasResolvedEvidence() bool {
	return s.ResolvedCount > 0
}
