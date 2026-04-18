package prism

// Phase represents a time-bounded planning period (typically a quarter).
type Phase struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name"`
	Quarter     string            `json:"quarter,omitempty"` // Q1, Q2, Q3, Q4
	Year        int               `json:"year,omitempty"`
	StartDate   string            `json:"startDate"`
	EndDate     string            `json:"endDate"`
	Status      string            `json:"status,omitempty"` // planning, in_progress, completed
	GoalTargets []PhaseGoalTarget `json:"goalTargets,omitempty"`
	Swimlanes   []Swimlane        `json:"swimlanes,omitempty"`
}

// Phase status constants.
const (
	PhaseStatusPlanning   = "planning"
	PhaseStatusInProgress = "in_progress"
	PhaseStatusCompleted  = "completed"
)

// AllPhaseStatuses returns all valid phase status values.
func AllPhaseStatuses() []string {
	return []string{PhaseStatusPlanning, PhaseStatusInProgress, PhaseStatusCompleted}
}

// Quarter constants.
const (
	QuarterQ1 = "Q1"
	QuarterQ2 = "Q2"
	QuarterQ3 = "Q3"
	QuarterQ4 = "Q4"
)

// AllQuarters returns all valid quarter values.
func AllQuarters() []string {
	return []string{QuarterQ1, QuarterQ2, QuarterQ3, QuarterQ4}
}

// PhaseGoalTarget specifies the maturity target for a goal within a phase.
type PhaseGoalTarget struct {
	GoalID     string `json:"goalId"`
	EnterLevel int    `json:"enterLevel"` // Expected maturity level at phase start
	ExitLevel  int    `json:"exitLevel"`  // Target maturity level at phase end
}

// Swimlane organizes initiatives within a phase by domain or stage.
type Swimlane struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"name"`
	Domain        string   `json:"domain,omitempty"` // security, operations
	Stage         string   `json:"stage,omitempty"`  // design, build, test, runtime, response
	InitiativeIDs []string `json:"initiativeIds"`
}

// GetGoalTarget returns the goal target for the specified goal ID.
func (p *Phase) GetGoalTarget(goalID string) *PhaseGoalTarget {
	for i := range p.GoalTargets {
		if p.GoalTargets[i].GoalID == goalID {
			return &p.GoalTargets[i]
		}
	}
	return nil
}

// GetSwimlane returns the swimlane with the specified ID.
func (p *Phase) GetSwimlane(swimlaneID string) *Swimlane {
	for i := range p.Swimlanes {
		if p.Swimlanes[i].ID == swimlaneID {
			return &p.Swimlanes[i]
		}
	}
	return nil
}

// AllInitiativeIDs returns all initiative IDs across all swimlanes in this phase.
func (p *Phase) AllInitiativeIDs() []string {
	seen := make(map[string]bool)
	var result []string
	for _, sw := range p.Swimlanes {
		for _, id := range sw.InitiativeIDs {
			if !seen[id] {
				seen[id] = true
				result = append(result, id)
			}
		}
	}
	return result
}

// RoadmapConfig holds configuration options for the roadmap.
type RoadmapConfig struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	FiscalYearStart int    `json:"fiscalYearStart,omitempty"` // Month (1-12), default 1 (January)
}
