package prism

// Goal represents a strategic objective with its own maturity model.
type Goal struct {
	ID            string             `json:"id,omitempty"`
	Name          string             `json:"name"`
	Description   string             `json:"description,omitempty"`
	Owner         string             `json:"owner,omitempty"`
	Priority      int                `json:"priority,omitempty"`
	Status        string             `json:"status,omitempty"` // active, on_hold, completed, cancelled
	StartDate     string             `json:"startDate,omitempty"`
	TargetDate    string             `json:"targetDate,omitempty"`
	MaturityModel *GoalMaturityModel `json:"maturityModel,omitempty"`
	CurrentLevel  int                `json:"currentLevel,omitempty"`
	TargetLevel   int                `json:"targetLevel,omitempty"`
}

// Goal status constants.
const (
	GoalStatusActive    = "active"
	GoalStatusOnHold    = "on_hold"
	GoalStatusCompleted = "completed"
	GoalStatusCancelled = "cancelled"
)

// AllGoalStatuses returns all valid goal status values.
func AllGoalStatuses() []string {
	return []string{GoalStatusActive, GoalStatusOnHold, GoalStatusCompleted, GoalStatusCancelled}
}

// GoalMaturityModel defines the 5-level maturity model for a specific goal.
type GoalMaturityModel struct {
	Levels []GoalMaturityLevel `json:"levels"`
}

// GoalMaturityLevel defines what a maturity level means for a goal,
// including the SLOs that must be met to achieve that level.
type GoalMaturityLevel struct {
	Level          int               `json:"level"`
	Name           string            `json:"name"`
	Description    string            `json:"description,omitempty"`
	RequiredSLOs   []SLORequirement  `json:"requiredSLOs,omitempty"`
	MetricCriteria []MetricCriterion `json:"metricCriteria,omitempty"`
}

// SLORequirement specifies an SLO that must be met for a maturity level.
type SLORequirement struct {
	MetricID    string `json:"metricId"`
	Description string `json:"description,omitempty"`
}

// MetricCriterion specifies a metric value requirement for a maturity level.
type MetricCriterion struct {
	MetricID string  `json:"metricId"`
	Operator string  `json:"operator"` // gte, lte, gt, lt, eq
	Value    float64 `json:"value"`
}

// IsMet returns whether the criterion is met given the current metric value.
func (mc *MetricCriterion) IsMet(current float64) bool {
	switch mc.Operator {
	case SLOOperatorGTE:
		return current >= mc.Value
	case SLOOperatorLTE:
		return current <= mc.Value
	case SLOOperatorGT:
		return current > mc.Value
	case SLOOperatorLT:
		return current < mc.Value
	case SLOOperatorEQ:
		return current == mc.Value
	default:
		return false
	}
}

// GetLevel returns the maturity level definition for the specified level number.
func (gmm *GoalMaturityModel) GetLevel(level int) *GoalMaturityLevel {
	if gmm == nil {
		return nil
	}
	for i := range gmm.Levels {
		if gmm.Levels[i].Level == level {
			return &gmm.Levels[i]
		}
	}
	return nil
}

// CurrentMaturityLevel calculates the current maturity level for a goal
// based on which SLOs and metric criteria are met.
// It checks from level 5 down to 1 and returns the highest level where
// all requirements are satisfied.
func (g *Goal) CurrentMaturityLevel(doc *PRISMDocument) int {
	if g.MaturityModel == nil {
		return 1
	}

	for level := MaturityLevel5; level >= MaturityLevel1; level-- {
		if g.MeetsLevelRequirements(level, doc) {
			return level
		}
	}
	return MaturityLevel1
}

// MeetsLevelRequirements returns whether all requirements for the specified
// maturity level are met.
func (g *Goal) MeetsLevelRequirements(level int, doc *PRISMDocument) bool {
	if g.MaturityModel == nil {
		return level == MaturityLevel1
	}

	levelDef := g.MaturityModel.GetLevel(level)
	if levelDef == nil {
		return false
	}

	// Check all required SLOs
	for _, sloReq := range levelDef.RequiredSLOs {
		metric := doc.GetMetricByID(sloReq.MetricID)
		if metric == nil {
			return false // Metric not found
		}
		if !metric.MeetsSLO() {
			return false
		}
	}

	// Check all metric criteria
	for _, criterion := range levelDef.MetricCriteria {
		metric := doc.GetMetricByID(criterion.MetricID)
		if metric == nil {
			return false // Metric not found
		}
		if !criterion.IsMet(metric.Current) {
			return false
		}
	}

	return true
}

// SLOsMetForLevel returns the count of SLOs met and total SLOs required
// for the specified maturity level.
func (g *Goal) SLOsMetForLevel(level int, doc *PRISMDocument) (met, total int) {
	if g.MaturityModel == nil {
		return 0, 0
	}

	levelDef := g.MaturityModel.GetLevel(level)
	if levelDef == nil {
		return 0, 0
	}

	total = len(levelDef.RequiredSLOs)
	for _, sloReq := range levelDef.RequiredSLOs {
		metric := doc.GetMetricByID(sloReq.MetricID)
		if metric != nil && metric.MeetsSLO() {
			met++
		}
	}
	return met, total
}

// CriteriaMetForLevel returns the count of metric criteria met and total criteria
// for the specified maturity level.
func (g *Goal) CriteriaMetForLevel(level int, doc *PRISMDocument) (met, total int) {
	if g.MaturityModel == nil {
		return 0, 0
	}

	levelDef := g.MaturityModel.GetLevel(level)
	if levelDef == nil {
		return 0, 0
	}

	total = len(levelDef.MetricCriteria)
	for _, criterion := range levelDef.MetricCriteria {
		metric := doc.GetMetricByID(criterion.MetricID)
		if metric != nil && criterion.IsMet(metric.Current) {
			met++
		}
	}
	return met, total
}
