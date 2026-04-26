// Package maturity provides types and functions for maturity model management.
package maturity

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Spec defines a complete maturity specification for an organization.
// It contains maturity models for multiple domains.
type Spec struct {
	Schema        string                       `json:"$schema,omitempty"`
	Metadata      *SpecMetadata                `json:"metadata,omitempty"`
	KPIThresholds map[string][]KPIThreshold    `json:"kpiThresholds,omitempty"`
	Domains       map[string]*DomainModel      `json:"domains"`
	Assessments   map[string]*DomainAssessment `json:"assessments,omitempty"`
}

// KPIThreshold defines the progression of a KPI across maturity levels.
type KPIThreshold struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Unit        string          `json:"unit,omitempty"`
	Operator    string          `json:"operator,omitempty"` // gte (default), lte for "lower is better"
	Thresholds  LevelThresholds `json:"thresholds"`
	Current     any             `json:"current,omitempty"` // Current value for assessment
}

// LevelThresholds holds threshold values for each maturity level.
type LevelThresholds struct {
	M1 any `json:"m1,omitempty"`
	M2 any `json:"m2,omitempty"`
	M3 any `json:"m3,omitempty"`
	M4 any `json:"m4,omitempty"`
	M5 any `json:"m5,omitempty"`
}

// SpecMetadata contains metadata about the maturity specification.
type SpecMetadata struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Version      string `json:"version,omitempty"`
	Organization string `json:"organization,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
}

// DomainModel defines maturity levels for a specific domain.
type DomainModel struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Owner       string  `json:"owner,omitempty"`
	Levels      []Level `json:"levels"`
}

// Level defines a maturity level (M1-M5) for a domain.
type Level struct {
	Level       int         `json:"level"` // 1-5
	Name        string      `json:"name"`  // Reactive, Basic, Defined, Managed, Optimizing
	Description string      `json:"description"`
	Criteria    []Criterion `json:"criteria,omitempty"` // SLOs that define the level
	Enablers    []Enabler   `json:"enablers,omitempty"` // Tasks to achieve the level
}

// Criterion is a measurable SLO that defines level achievement.
type Criterion struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	// SLO Definition
	MetricName string  `json:"metricName"`     // Human-readable metric description
	Operator   string  `json:"operator"`       // gte, lte, gt, lt, eq
	Target     float64 `json:"target"`         // Target value
	Unit       string  `json:"unit,omitempty"` // %, days, count, seconds, etc.

	// Classification
	Layer    string `json:"layer,omitempty"`    // requirements, code, infra, runtime, adoption, support
	Category string `json:"category,omitempty"` // prevention, detection, response

	// Assessment (populated during evaluation)
	Current float64 `json:"current,omitempty"` // Current value
	IsMet   bool    `json:"isMet,omitempty"`   // Calculated: meets target?

	// Weighting
	Weight   float64 `json:"weight,omitempty"`   // Relative importance (default 1.0)
	Required bool    `json:"required,omitempty"` // Must pass for level (default true if omitted)
}

// Enabler is implementation work to achieve criteria.
type Enabler struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	// Classification
	Type  string `json:"type,omitempty"`  // implementation, process, training, tooling
	Layer string `json:"layer,omitempty"` // requirements, code, infra, runtime, adoption, support

	// Effort
	Effort string `json:"effort,omitempty"` // T-shirt size or duration
	Team   string `json:"team,omitempty"`   // Responsible team

	// Tracking
	Status      string   `json:"status,omitempty"`      // not_started, in_progress, completed
	CriteriaIDs []string `json:"criteriaIds,omitempty"` // Which SLOs this enables

	// Dependencies
	DependsOn []string `json:"dependsOn,omitempty"` // Other enabler IDs
}

// DomainAssessment captures current state against a domain's maturity model.
type DomainAssessment struct {
	Domain         string             `json:"domain"`
	AssessedAt     string             `json:"assessedAt,omitempty"`
	AssessedBy     string             `json:"assessedBy,omitempty"`
	CurrentLevel   int                `json:"currentLevel"`             // Achieved level (1-5)
	TargetLevel    int                `json:"targetLevel"`              // Goal level
	CriteriaValues map[string]float64 `json:"criteriaValues,omitempty"` // Current values by criterion ID
	EnablerStatus  map[string]string  `json:"enablerStatus,omitempty"`  // Status by enabler ID
}

// Enabler status constants.
const (
	StatusNotStarted = "not_started"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
	StatusBlocked    = "blocked"
)

// Enabler type constants.
const (
	TypeImplementation = "implementation"
	TypeProcess        = "process"
	TypeTraining       = "training"
	TypeTooling        = "tooling"
)

// Operator constants.
const (
	OpGTE = "gte" // Greater than or equal
	OpLTE = "lte" // Less than or equal
	OpGT  = "gt"  // Greater than
	OpLT  = "lt"  // Less than
	OpEQ  = "eq"  // Equal
)

// Level name constants.
const (
	LevelNameReactive   = "Reactive"
	LevelNameBasic      = "Basic"
	LevelNameDefined    = "Defined"
	LevelNameManaged    = "Managed"
	LevelNameOptimizing = "Optimizing"
)

// DefaultLevelNames returns the standard M1-M5 level names.
func DefaultLevelNames() map[int]string {
	return map[int]string{
		1: LevelNameReactive,
		2: LevelNameBasic,
		3: LevelNameDefined,
		4: LevelNameManaged,
		5: LevelNameOptimizing,
	}
}

// ReadSpecFile reads a maturity spec from a JSON file.
func ReadSpecFile(filename string) (*Spec, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	var spec Spec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &spec, nil
}

// WriteSpecFile writes a maturity spec to a JSON file.
func (s *Spec) WriteSpecFile(filename string) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal spec: %w", err)
	}

	if err := os.WriteFile(filename, data, 0600); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	return nil
}

// GetDomain returns the domain model by name.
func (s *Spec) GetDomain(name string) (*DomainModel, bool) {
	domain, ok := s.Domains[strings.ToLower(name)]
	return domain, ok
}

// GetLevel returns the level definition for a domain.
func (d *DomainModel) GetLevel(level int) (*Level, bool) {
	for i := range d.Levels {
		if d.Levels[i].Level == level {
			return &d.Levels[i], true
		}
	}
	return nil, false
}

// OperatorSymbol returns the symbol for an operator.
func OperatorSymbol(op string) string {
	switch op {
	case OpGTE:
		return ">="
	case OpLTE:
		return "<="
	case OpGT:
		return ">"
	case OpLT:
		return "<"
	case OpEQ:
		return "="
	default:
		return op
	}
}

// IsMet checks if a criterion is met given a current value.
func (c *Criterion) CheckMet(current float64) bool {
	switch c.Operator {
	case OpGTE:
		return current >= c.Target
	case OpLTE:
		return current <= c.Target
	case OpGT:
		return current > c.Target
	case OpLT:
		return current < c.Target
	case OpEQ:
		return current == c.Target
	default:
		return false
	}
}

// TargetString returns a formatted target string like ">=95%".
func (c *Criterion) TargetString() string {
	symbol := OperatorSymbol(c.Operator)
	if c.Unit != "" {
		return fmt.Sprintf("%s%.0f%s", symbol, c.Target, c.Unit)
	}
	return fmt.Sprintf("%s%.0f", symbol, c.Target)
}

// CurrentString returns a formatted current value string.
func (c *Criterion) CurrentString() string {
	if c.Unit != "" {
		return fmt.Sprintf("%.1f%s", c.Current, c.Unit)
	}
	return fmt.Sprintf("%.1f", c.Current)
}

// LevelProgress tracks progress toward a maturity level.
type LevelProgress struct {
	Level           int     `json:"level"`
	CriteriaMet     int     `json:"criteriaMet"`
	CriteriaTotal   int     `json:"criteriaTotal"`
	ProgressPercent float64 `json:"progressPercent"`
	EnablersDone    int     `json:"enablersDone"`
	EnablersTotal   int     `json:"enablersTotal"`
}

// CalculateLevelProgress calculates progress for a level given current values.
func (l *Level) CalculateLevelProgress(values map[string]float64, enablerStatus map[string]string) LevelProgress {
	progress := LevelProgress{
		Level:         l.Level,
		CriteriaTotal: len(l.Criteria),
		EnablersTotal: len(l.Enablers),
	}

	for _, c := range l.Criteria {
		if current, ok := values[c.ID]; ok {
			if c.CheckMet(current) {
				progress.CriteriaMet++
			}
		}
	}

	for _, e := range l.Enablers {
		if status, ok := enablerStatus[e.ID]; ok && status == StatusCompleted {
			progress.EnablersDone++
		}
	}

	if progress.CriteriaTotal > 0 {
		progress.ProgressPercent = float64(progress.CriteriaMet) / float64(progress.CriteriaTotal) * 100
	} else {
		progress.ProgressPercent = 100 // No criteria means level is achieved
	}

	return progress
}

// IsLevelAchieved checks if all required criteria for a level are met.
func (l *Level) IsLevelAchieved(values map[string]float64) bool {
	for _, c := range l.Criteria {
		// Default to required if not specified
		required := c.Required || c.Weight == 0
		if !required {
			continue
		}

		current, ok := values[c.ID]
		if !ok {
			return false // Missing value for required criterion
		}

		if !c.CheckMet(current) {
			return false
		}
	}
	return true
}

// AllCriteria returns all criteria across all levels for a domain.
func (d *DomainModel) AllCriteria() []Criterion {
	var all []Criterion
	for _, level := range d.Levels {
		for _, c := range level.Criteria {
			all = append(all, c)
		}
	}
	return all
}

// AllEnablers returns all enablers across all levels for a domain.
func (d *DomainModel) AllEnablers() []Enabler {
	var all []Enabler
	for _, level := range d.Levels {
		for _, e := range level.Enablers {
			all = append(all, e)
		}
	}
	return all
}

// CriteriaForLevel returns criteria for a specific level.
func (d *DomainModel) CriteriaForLevel(level int) []Criterion {
	for _, l := range d.Levels {
		if l.Level == level {
			return l.Criteria
		}
	}
	return nil
}

// EnablersForLevel returns enablers for a specific level.
func (d *DomainModel) EnablersForLevel(level int) []Enabler {
	for _, l := range d.Levels {
		if l.Level == level {
			return l.Enablers
		}
	}
	return nil
}
