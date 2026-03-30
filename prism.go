// Package prism provides the PRISM (Proactive Reliability & Security Maturity Model)
// framework for B2B SaaS health metrics combining SLOs, DMAIC, OKRs, and maturity modeling.
package prism

import "time"

// PRISMDocument represents the top-level PRISM document.
type PRISMDocument struct {
	Schema      string         `json:"$schema,omitempty"`
	Metadata    *Metadata      `json:"metadata,omitempty"`
	Domains     []DomainDef    `json:"domains,omitempty"`
	Metrics     []Metric       `json:"metrics"`
	Maturity    *MaturityModel `json:"maturity,omitempty"`
	OKRs        []OKRMapping   `json:"okrs,omitempty"`
	Initiatives []Initiative   `json:"initiatives,omitempty"`
}

// Metadata contains document-level metadata.
type Metadata struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Author      string `json:"author,omitempty"`
	Created     string `json:"created,omitempty"`
	Updated     string `json:"updated,omitempty"`
}

// DomainDef defines a PRISM domain (security or operations).
type DomainDef struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
}

// Metric represents a PRISM metric with SLO, maturity, and framework mappings.
type Metric struct {
	// Core identity
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	// PRISM classification
	Domain   string `json:"domain"`
	Stage    string `json:"stage"`
	Category string `json:"category"`

	// Measurement
	MetricType     string  `json:"metricType"`
	TrendDirection string  `json:"trendDirection,omitempty"`
	Unit           string  `json:"unit,omitempty"`
	Baseline       float64 `json:"baseline"`
	Current        float64 `json:"current"`
	Target         float64 `json:"target"`

	// SLI/SLO
	SLI *SLI `json:"sli,omitempty"`
	SLO *SLO `json:"slo,omitempty"`

	// Thresholds & Status
	Thresholds *Thresholds `json:"thresholds,omitempty"`
	Status     string      `json:"status,omitempty"`

	// Maturity mapping
	MaturityMapping *MaturityMapping `json:"maturityMapping,omitempty"`

	// DMAIC mapping
	DMAIC *DMAICMapping `json:"dmaic,omitempty"`

	// Customer awareness
	CustomerAwareness *CustomerAwarenessConfig `json:"customerAwareness,omitempty"`

	// Framework mappings
	FrameworkMappings []FrameworkMapping `json:"frameworkMappings,omitempty"`

	// Ownership
	Owner      string `json:"owner,omitempty"`
	DataSource string `json:"dataSource,omitempty"`

	// History
	DataPoints []DataPoint `json:"dataPoints,omitempty"`
}

// SLI represents a Service Level Indicator.
type SLI struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Formula     string `json:"formula,omitempty"`
}

// SLO represents a Service Level Objective.
type SLO struct {
	Target     string      `json:"target"`               // Display string (e.g., ">=99.9%")
	Operator   string      `json:"operator,omitempty"`   // Machine-readable: "gte", "lte", "eq", "gt", "lt"
	Value      float64     `json:"value,omitempty"`      // Numeric target value
	Window     string      `json:"window,omitempty"`     // "7d", "30d", "90d"
	Thresholds *Thresholds `json:"thresholds,omitempty"`
}

// SLO operator constants.
const (
	SLOOperatorGTE = "gte" // Greater than or equal
	SLOOperatorLTE = "lte" // Less than or equal
	SLOOperatorEQ  = "eq"  // Equal
	SLOOperatorGT  = "gt"  // Greater than
	SLOOperatorLT  = "lt"  // Less than
)

// Thresholds defines threshold values for status calculation.
type Thresholds struct {
	Green  float64 `json:"green"`
	Yellow float64 `json:"yellow"`
	Red    float64 `json:"red"`
}

// MaturityMapping maps metric values to maturity levels.
type MaturityMapping struct {
	Level1 string `json:"level1,omitempty"`
	Level2 string `json:"level2,omitempty"`
	Level3 string `json:"level3,omitempty"`
	Level4 string `json:"level4,omitempty"`
	Level5 string `json:"level5,omitempty"`
}

// DMAICMapping maps the metric to DMAIC phases.
type DMAICMapping struct {
	Define  string `json:"define,omitempty"`
	Measure string `json:"measure,omitempty"`
	Analyze string `json:"analyze,omitempty"`
	Improve string `json:"improve,omitempty"`
	Control string `json:"control,omitempty"`
}

// FrameworkMapping maps a metric to an external framework reference.
type FrameworkMapping struct {
	Framework string `json:"framework"`
	Reference string `json:"reference"`
}

// DataPoint represents a historical measurement.
type DataPoint struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
	Note      string    `json:"note,omitempty"`
}

// OKRMapping represents alignment between metrics and OKRs.
type OKRMapping struct {
	ObjectiveID   string   `json:"objectiveId,omitempty"`
	ObjectiveName string   `json:"objectiveName"`
	KeyResultID   string   `json:"keyResultId,omitempty"`
	KeyResultName string   `json:"keyResultName,omitempty"`
	MetricIDs     []string `json:"metricIds,omitempty"`
}

// Initiative represents an improvement initiative.
type Initiative struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Status      string   `json:"status,omitempty"`
	Priority    int      `json:"priority,omitempty"`
	MetricIDs   []string `json:"metricIds,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	StartDate   string   `json:"startDate,omitempty"`
	EndDate     string   `json:"endDate,omitempty"`
}

// CalculateStatus computes the status based on current value and thresholds.
// For higher_better trends: value >= green threshold = Green, etc.
// For lower_better trends: value <= green threshold = Green, etc.
func (m *Metric) CalculateStatus() string {
	if m.Thresholds == nil {
		return ""
	}

	current := m.Current
	th := m.Thresholds

	switch m.TrendDirection {
	case TrendHigherBetter:
		if current >= th.Green {
			return StatusGreen
		} else if current >= th.Yellow {
			return StatusYellow
		}
		return StatusRed
	case TrendLowerBetter:
		if current <= th.Green {
			return StatusGreen
		} else if current <= th.Yellow {
			return StatusYellow
		}
		return StatusRed
	case TrendTargetValue:
		// For target value, check distance from target
		diff := abs(current - m.Target)
		greenRange := abs(th.Green - m.Target)
		yellowRange := abs(th.Yellow - m.Target)
		if diff <= greenRange {
			return StatusGreen
		} else if diff <= yellowRange {
			return StatusYellow
		}
		return StatusRed
	default:
		// Default to higher_better behavior
		if current >= th.Green {
			return StatusGreen
		} else if current >= th.Yellow {
			return StatusYellow
		}
		return StatusRed
	}
}

// ProgressToTarget returns the progress as a ratio (0.0-1.0) toward the target.
func (m *Metric) ProgressToTarget() float64 {
	if m.Target == m.Baseline {
		if m.Current >= m.Target {
			return 1.0
		}
		return 0.0
	}

	progress := (m.Current - m.Baseline) / (m.Target - m.Baseline)
	if progress < 0 {
		return 0.0
	}
	if progress > 1 {
		return 1.0
	}
	return progress
}

// MeetsSLO returns whether the metric's current value meets its SLO.
// Returns true if no SLO is defined or if Operator/Value are not set.
// Uses the structured Operator and Value fields for evaluation.
func (m *Metric) MeetsSLO() bool {
	if m.SLO == nil {
		return true // No SLO defined
	}
	if m.SLO.Operator == "" {
		return true // No machine-readable operator defined
	}

	current := m.Current
	target := m.SLO.Value

	switch m.SLO.Operator {
	case SLOOperatorGTE:
		return current >= target
	case SLOOperatorLTE:
		return current <= target
	case SLOOperatorGT:
		return current > target
	case SLOOperatorLT:
		return current < target
	case SLOOperatorEQ:
		return current == target
	default:
		return true // Unknown operator, assume met
	}
}

// GetMetricsByDomain returns all metrics for the specified domain.
func (doc *PRISMDocument) GetMetricsByDomain(domain string) []Metric {
	var result []Metric
	for _, m := range doc.Metrics {
		if m.Domain == domain {
			result = append(result, m)
		}
	}
	return result
}

// GetMetricsByStage returns all metrics for the specified stage.
func (doc *PRISMDocument) GetMetricsByStage(stage string) []Metric {
	var result []Metric
	for _, m := range doc.Metrics {
		if m.Stage == stage {
			result = append(result, m)
		}
	}
	return result
}

// GetMetricsByCategory returns all metrics for the specified category.
func (doc *PRISMDocument) GetMetricsByCategory(category string) []Metric {
	var result []Metric
	for _, m := range doc.Metrics {
		if m.Category == category {
			result = append(result, m)
		}
	}
	return result
}

// GetMetricByID returns a metric by its ID.
func (doc *PRISMDocument) GetMetricByID(id string) *Metric {
	for i := range doc.Metrics {
		if doc.Metrics[i].ID == id {
			return &doc.Metrics[i]
		}
	}
	return nil
}

// abs returns the absolute value of x.
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
