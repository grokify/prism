package prism

import "fmt"

// CustomerAwarenessConfig defines whether customer awareness tracking is enabled for a metric.
type CustomerAwarenessConfig struct {
	Enabled bool     `json:"enabled"`
	States  []string `json:"states,omitempty"`
}

// CustomerAwarenessData represents customer awareness distribution for a period.
type CustomerAwarenessData struct {
	Period       string                  `json:"period"`
	Distribution []AwarenessDistribution `json:"distribution"`
}

// AwarenessDistribution represents the count and percentage for a single awareness state.
type AwarenessDistribution struct {
	State   string  `json:"state"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

// DefaultAwarenessStates returns the default four awareness states.
func DefaultAwarenessStates() []string {
	return AllAwarenessStates()
}

// NewCustomerAwarenessConfig creates a new config with defaults.
func NewCustomerAwarenessConfig(enabled bool) *CustomerAwarenessConfig {
	return &CustomerAwarenessConfig{
		Enabled: enabled,
		States:  DefaultAwarenessStates(),
	}
}

// NewCustomerAwarenessData creates awareness data with zero counts.
func NewCustomerAwarenessData(period string) *CustomerAwarenessData {
	distribution := make([]AwarenessDistribution, len(AllAwarenessStates()))
	for i, state := range AllAwarenessStates() {
		distribution[i] = AwarenessDistribution{
			State:   state,
			Count:   0,
			Percent: 0,
		}
	}
	return &CustomerAwarenessData{
		Period:       period,
		Distribution: distribution,
	}
}

// TotalCount returns the total count across all awareness states.
func (d *CustomerAwarenessData) TotalCount() int {
	total := 0
	for _, dist := range d.Distribution {
		total += dist.Count
	}
	return total
}

// GetStateCount returns the count for a specific awareness state.
func (d *CustomerAwarenessData) GetStateCount(state string) int {
	for _, dist := range d.Distribution {
		if dist.State == state {
			return dist.Count
		}
	}
	return 0
}

// GetStatePercent returns the percentage for a specific awareness state.
func (d *CustomerAwarenessData) GetStatePercent(state string) float64 {
	for _, dist := range d.Distribution {
		if dist.State == state {
			return dist.Percent
		}
	}
	return 0
}

// SetCount sets the count for a specific awareness state and recalculates percentages.
func (d *CustomerAwarenessData) SetCount(state string, count int) error {
	if err := ValidateAwarenessState(state); err != nil {
		return err
	}

	found := false
	for i := range d.Distribution {
		if d.Distribution[i].State == state {
			d.Distribution[i].Count = count
			found = true
			break
		}
	}

	if !found {
		d.Distribution = append(d.Distribution, AwarenessDistribution{
			State: state,
			Count: count,
		})
	}

	d.RecalculatePercentages()
	return nil
}

// RecalculatePercentages recalculates all percentages based on counts.
func (d *CustomerAwarenessData) RecalculatePercentages() {
	total := d.TotalCount()
	if total == 0 {
		for i := range d.Distribution {
			d.Distribution[i].Percent = 0
		}
		return
	}

	for i := range d.Distribution {
		d.Distribution[i].Percent = float64(d.Distribution[i].Count) / float64(total) * 100
	}
}

// UnawareRate returns the rate (0.0-1.0) of customers who are unaware.
func (d *CustomerAwarenessData) UnawareRate() float64 {
	total := d.TotalCount()
	if total == 0 {
		return 0
	}
	return float64(d.GetStateCount(AwarenessUnaware)) / float64(total)
}

// ProactiveDetectionRate returns 1 - unaware rate (rate of customers who are aware).
func (d *CustomerAwarenessData) ProactiveDetectionRate() float64 {
	return 1 - d.UnawareRate()
}

// ProactiveResolutionRate returns the rate of customers who have remediated.
func (d *CustomerAwarenessData) ProactiveResolutionRate() float64 {
	total := d.TotalCount()
	if total == 0 {
		return 0
	}
	return float64(d.GetStateCount(AwarenessAwareRemediated)) / float64(total)
}

// RemediationInProgressRate returns the rate of customers actively remediating.
func (d *CustomerAwarenessData) RemediationInProgressRate() float64 {
	total := d.TotalCount()
	if total == 0 {
		return 0
	}
	return float64(d.GetStateCount(AwarenessAwareRemediating)) / float64(total)
}

// AwareNotActingRate returns the rate of customers who are aware but not remediating.
func (d *CustomerAwarenessData) AwareNotActingRate() float64 {
	total := d.TotalCount()
	if total == 0 {
		return 0
	}
	return float64(d.GetStateCount(AwarenessAwareNotActing)) / float64(total)
}

// AwarenessScore returns a composite awareness score (0.0-1.0).
// Higher scores indicate better awareness/remediation state.
// Uses mutually exclusive states with weighted values:
//   - unaware: 0.0 (worst - customer doesn't know about the issue)
//   - aware_not_acting: 0.25 (customer knows but hasn't started remediation)
//   - remediating: 0.5 (customer is actively working on it)
//   - remediated: 1.0 (best - customer has resolved the issue)
func (d *CustomerAwarenessData) AwarenessScore() float64 {
	unaware := d.UnawareRate()
	notActing := d.AwareNotActingRate()
	remediating := d.RemediationInProgressRate()
	remediated := d.ProactiveResolutionRate()

	return (unaware * 0.0) + (notActing * 0.25) + (remediating * 0.5) + (remediated * 1.0)
}

// Validate validates the awareness data.
func (d *CustomerAwarenessData) Validate() ValidationErrors {
	var errs ValidationErrors

	if d.Period == "" {
		errs = append(errs, ValidationError{Field: "period", Message: "is required"})
	}

	for i, dist := range d.Distribution {
		if err := ValidateAwarenessState(dist.State); err != nil {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("distribution[%d].state", i),
				Value:   dist.State,
				Message: err.Error(),
			})
		}
		if dist.Count < 0 {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("distribution[%d].count", i),
				Value:   fmt.Sprintf("%d", dist.Count),
				Message: "must be non-negative",
			})
		}
	}

	return errs
}

// AwarenessSummary provides a summary view of awareness metrics.
type AwarenessSummary struct {
	TotalCustomers          int     `json:"totalCustomers"`
	UnawareCount            int     `json:"unawareCount"`
	AwareCount              int     `json:"awareCount"`
	RemediatingCount        int     `json:"remediatingCount"`
	RemediatedCount         int     `json:"remediatedCount"`
	ProactiveDetectionRate  float64 `json:"proactiveDetectionRate"`
	ProactiveResolutionRate float64 `json:"proactiveResolutionRate"`
	AwarenessScore          float64 `json:"awarenessScore"`
}

// Summary returns a summary of the awareness data.
func (d *CustomerAwarenessData) Summary() *AwarenessSummary {
	return &AwarenessSummary{
		TotalCustomers:          d.TotalCount(),
		UnawareCount:            d.GetStateCount(AwarenessUnaware),
		AwareCount:              d.GetStateCount(AwarenessAwareNotActing),
		RemediatingCount:        d.GetStateCount(AwarenessAwareRemediating),
		RemediatedCount:         d.GetStateCount(AwarenessAwareRemediated),
		ProactiveDetectionRate:  d.ProactiveDetectionRate(),
		ProactiveResolutionRate: d.ProactiveResolutionRate(),
		AwarenessScore:          d.AwarenessScore(),
	}
}
