package prism

// GoalRoadmapView shows a goal's progress across all phases (Goal → Initiative → Phase).
type GoalRoadmapView struct {
	GoalID        string              `json:"goalId"`
	GoalName      string              `json:"goalName"`
	Description   string              `json:"description,omitempty"`
	CurrentLevel  int                 `json:"currentLevel"`
	TargetLevel   int                 `json:"targetLevel"`
	PhaseProgress []GoalPhaseProgress `json:"phaseProgress"`
}

// GoalPhaseProgress tracks a goal's progress within a specific phase.
type GoalPhaseProgress struct {
	PhaseID              string              `json:"phaseId"`
	PhaseName            string              `json:"phaseName"`
	Quarter              string              `json:"quarter,omitempty"`
	Year                 int                 `json:"year,omitempty"`
	EnterLevel           int                 `json:"enterLevel"`
	ExitLevel            int                 `json:"exitLevel"`
	InitiativesTotal     int                 `json:"initiativesTotal"`
	InitiativesCompleted int                 `json:"initiativesCompleted"`
	CompletionPercent    float64             `json:"completionPercent"`
	Initiatives          []InitiativeSummary `json:"initiatives,omitempty"`
}

// InitiativeSummary provides a brief summary of an initiative.
type InitiativeSummary struct {
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	Status               string  `json:"status,omitempty"`
	Team                 string  `json:"team,omitempty"`
	DevCompletionPercent float64 `json:"devCompletionPercent"`
}

// PhaseRoadmapView shows a phase's goals and initiatives (Phase → Goal → Initiative).
type PhaseRoadmapView struct {
	PhaseID           string          `json:"phaseId"`
	PhaseName         string          `json:"phaseName"`
	Quarter           string          `json:"quarter,omitempty"`
	Year              int             `json:"year,omitempty"`
	StartDate         string          `json:"startDate"`
	EndDate           string          `json:"endDate"`
	Status            string          `json:"status,omitempty"`
	GoalViews         []PhaseGoalView `json:"goalViews"`
	OverallCompletion float64         `json:"overallCompletion"`
}

// PhaseGoalView shows a goal's status within a phase.
type PhaseGoalView struct {
	GoalID               string              `json:"goalId"`
	GoalName             string              `json:"goalName"`
	EnterLevel           int                 `json:"enterLevel"`
	ExitLevel            int                 `json:"exitLevel"`
	CurrentLevel         int                 `json:"currentLevel"`
	InitiativesTotal     int                 `json:"initiativesTotal"`
	InitiativesCompleted int                 `json:"initiativesCompleted"`
	CompletionPercent    float64             `json:"completionPercent"`
	Initiatives          []InitiativeSummary `json:"initiatives,omitempty"`
}

// RoadmapReport contains both views of the roadmap.
type RoadmapReport struct {
	Metadata    *Metadata          `json:"metadata,omitempty"`
	ByGoal      []GoalRoadmapView  `json:"byGoal"`
	ByPhase     []PhaseRoadmapView `json:"byPhase"`
	GeneratedAt string             `json:"generatedAt,omitempty"`
}

// GenerateGoalRoadmapView creates a Goal-centric view for a specific goal.
func (doc *PRISMDocument) GenerateGoalRoadmapView(goalID string) *GoalRoadmapView {
	goal := doc.GetGoalByID(goalID)
	if goal == nil {
		return nil
	}

	view := &GoalRoadmapView{
		GoalID:       goal.ID,
		GoalName:     goal.Name,
		Description:  goal.Description,
		CurrentLevel: goal.CurrentLevel,
		TargetLevel:  goal.TargetLevel,
	}

	// Sort phases by year and quarter
	phases := doc.GetPhasesSorted()

	for _, phase := range phases {
		// Check if this goal is targeted in this phase
		target := phase.GetGoalTarget(goalID)
		if target == nil {
			continue
		}

		progress := GoalPhaseProgress{
			PhaseID:    phase.ID,
			PhaseName:  phase.Name,
			Quarter:    phase.Quarter,
			Year:       phase.Year,
			EnterLevel: target.EnterLevel,
			ExitLevel:  target.ExitLevel,
		}

		// Find initiatives for this goal in this phase
		for _, init := range doc.Initiatives {
			if init.PhaseID != phase.ID {
				continue
			}
			for _, gid := range init.GoalIDs {
				if gid == goalID {
					progress.InitiativesTotal++
					if init.IsDevComplete() {
						progress.InitiativesCompleted++
					}
					progress.Initiatives = append(progress.Initiatives, InitiativeSummary{
						ID:                   init.ID,
						Name:                 init.Name,
						Status:               init.Status,
						Team:                 init.Team,
						DevCompletionPercent: init.DevCompletionPercent,
					})
					break
				}
			}
		}

		if progress.InitiativesTotal > 0 {
			progress.CompletionPercent = float64(progress.InitiativesCompleted) / float64(progress.InitiativesTotal) * 100
		}

		view.PhaseProgress = append(view.PhaseProgress, progress)
	}

	return view
}

// GeneratePhaseRoadmapView creates a Phase-centric view for a specific phase.
func (doc *PRISMDocument) GeneratePhaseRoadmapView(phaseID string) *PhaseRoadmapView {
	phase := doc.GetPhaseByID(phaseID)
	if phase == nil {
		return nil
	}

	view := &PhaseRoadmapView{
		PhaseID:   phase.ID,
		PhaseName: phase.Name,
		Quarter:   phase.Quarter,
		Year:      phase.Year,
		StartDate: phase.StartDate,
		EndDate:   phase.EndDate,
		Status:    phase.Status,
	}

	totalInitiatives := 0
	completedInitiatives := 0

	for _, target := range phase.GoalTargets {
		goal := doc.GetGoalByID(target.GoalID)
		if goal == nil {
			continue
		}

		goalView := PhaseGoalView{
			GoalID:       goal.ID,
			GoalName:     goal.Name,
			EnterLevel:   target.EnterLevel,
			ExitLevel:    target.ExitLevel,
			CurrentLevel: goal.CurrentMaturityLevel(doc),
		}

		// Find initiatives for this goal in this phase
		for _, init := range doc.Initiatives {
			if init.PhaseID != phaseID {
				continue
			}
			for _, gid := range init.GoalIDs {
				if gid == goal.ID {
					goalView.InitiativesTotal++
					totalInitiatives++
					if init.IsDevComplete() {
						goalView.InitiativesCompleted++
						completedInitiatives++
					}
					goalView.Initiatives = append(goalView.Initiatives, InitiativeSummary{
						ID:                   init.ID,
						Name:                 init.Name,
						Status:               init.Status,
						Team:                 init.Team,
						DevCompletionPercent: init.DevCompletionPercent,
					})
					break
				}
			}
		}

		if goalView.InitiativesTotal > 0 {
			goalView.CompletionPercent = float64(goalView.InitiativesCompleted) / float64(goalView.InitiativesTotal) * 100
		}

		view.GoalViews = append(view.GoalViews, goalView)
	}

	if totalInitiatives > 0 {
		view.OverallCompletion = float64(completedInitiatives) / float64(totalInitiatives) * 100
	}

	return view
}

// GenerateRoadmapReport creates a complete roadmap report with both views.
func (doc *PRISMDocument) GenerateRoadmapReport() *RoadmapReport {
	report := &RoadmapReport{
		Metadata: doc.Metadata,
	}

	// Generate Goal-centric views
	for _, goal := range doc.Goals {
		view := doc.GenerateGoalRoadmapView(goal.ID)
		if view != nil {
			report.ByGoal = append(report.ByGoal, *view)
		}
	}

	// Generate Phase-centric views
	phases := doc.GetPhasesSorted()
	for _, phase := range phases {
		view := doc.GeneratePhaseRoadmapView(phase.ID)
		if view != nil {
			report.ByPhase = append(report.ByPhase, *view)
		}
	}

	return report
}

// GetPhasesSorted returns phases sorted by year and quarter.
func (doc *PRISMDocument) GetPhasesSorted() []Phase {
	if len(doc.Phases) == 0 {
		return nil
	}

	// Create a copy to avoid modifying original
	phases := make([]Phase, len(doc.Phases))
	copy(phases, doc.Phases)

	// Sort by year, then quarter
	for i := 0; i < len(phases)-1; i++ {
		for j := i + 1; j < len(phases); j++ {
			if phases[j].Year < phases[i].Year ||
				(phases[j].Year == phases[i].Year && quarterOrder(phases[j].Quarter) < quarterOrder(phases[i].Quarter)) {
				phases[i], phases[j] = phases[j], phases[i]
			}
		}
	}

	return phases
}

func quarterOrder(q string) int {
	switch q {
	case QuarterQ1:
		return 1
	case QuarterQ2:
		return 2
	case QuarterQ3:
		return 3
	case QuarterQ4:
		return 4
	default:
		return 0
	}
}
