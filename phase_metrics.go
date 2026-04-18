package prism

// PhaseMetrics tracks progress at phase boundaries.
type PhaseMetrics struct {
	PhaseID           string             `json:"phaseId"`
	GoalProgress      []GoalProgress     `json:"goalProgress,omitempty"`
	InitiativeMetrics *InitiativeMetrics `json:"initiativeMetrics,omitempty"`
	SLOCompliance     []SLOCompliance    `json:"sloCompliance,omitempty"`
}

// GoalProgress tracks progress for a goal within a phase.
type GoalProgress struct {
	GoalID               string  `json:"goalId"`
	EnterLevel           int     `json:"enterLevel"`           // Maturity level at phase entry
	CurrentLevel         int     `json:"currentLevel"`         // Current maturity level
	TargetLevel          int     `json:"targetLevel"`          // Target level for phase exit
	InitiativesTotal     int     `json:"initiativesTotal"`     // Total initiatives for this goal
	InitiativesCompleted int     `json:"initiativesCompleted"` // Completed initiatives
	CompletionPercent    float64 `json:"completionPercent"`    // initiativesCompleted / initiativesTotal
	SLOsRequired         int     `json:"slosRequired"`         // SLOs required for target level
	SLOsMet              int     `json:"slosMet"`              // SLOs currently met
}

// InitiativeMetrics provides aggregate statistics for initiatives in a phase.
type InitiativeMetrics struct {
	Total              int     `json:"total"`                        // Total initiatives in phase
	Completed          int     `json:"completed"`                    // Dev-complete initiatives
	Deployed           int     `json:"deployed"`                     // Fully deployed initiatives
	AvgAdoptionPercent float64 `json:"avgAdoptionPercent,omitempty"` // Average adoption across completed
}

// SLOCompliance tracks SLO compliance status.
type SLOCompliance struct {
	MetricID   string   `json:"metricId"`
	MetricName string   `json:"metricName,omitempty"`
	SLOTarget  string   `json:"sloTarget,omitempty"`
	Current    float64  `json:"current"`
	IsMet      bool     `json:"isMet"`
	GoalIDs    []string `json:"goalIds,omitempty"` // Goals that depend on this SLO
}

// CalculateGoalProgress computes progress for a goal within a phase.
func CalculateGoalProgress(goal *Goal, phase *Phase, doc *PRISMDocument) *GoalProgress {
	if goal == nil || phase == nil || doc == nil {
		return nil
	}

	// Find the goal target for this phase
	target := phase.GetGoalTarget(goal.ID)
	enterLevel := 1
	targetLevel := 1
	if target != nil {
		enterLevel = target.EnterLevel
		targetLevel = target.ExitLevel
	}

	// Count initiatives for this goal in this phase
	initiativesTotal := 0
	initiativesCompleted := 0

	for _, init := range doc.Initiatives {
		// Check if initiative is in this phase and linked to this goal
		if init.PhaseID != phase.ID {
			continue
		}
		for _, gid := range init.GoalIDs {
			if gid == goal.ID {
				initiativesTotal++
				if init.IsDevComplete() {
					initiativesCompleted++
				}
				break
			}
		}
	}

	// Calculate completion percentage
	completionPercent := 0.0
	if initiativesTotal > 0 {
		completionPercent = float64(initiativesCompleted) / float64(initiativesTotal) * 100
	}

	// Get SLO compliance for target level
	slosMet, slosRequired := goal.SLOsMetForLevel(targetLevel, doc)

	return &GoalProgress{
		GoalID:               goal.ID,
		EnterLevel:           enterLevel,
		CurrentLevel:         goal.CurrentMaturityLevel(doc),
		TargetLevel:          targetLevel,
		InitiativesTotal:     initiativesTotal,
		InitiativesCompleted: initiativesCompleted,
		CompletionPercent:    completionPercent,
		SLOsRequired:         slosRequired,
		SLOsMet:              slosMet,
	}
}

// CalculateInitiativeMetrics computes aggregate initiative metrics for a phase.
func CalculateInitiativeMetrics(phase *Phase, doc *PRISMDocument) *InitiativeMetrics {
	if phase == nil || doc == nil {
		return nil
	}

	metrics := &InitiativeMetrics{}
	var totalAdoption float64
	var adoptionCount int

	for _, init := range doc.Initiatives {
		if init.PhaseID != phase.ID {
			continue
		}

		metrics.Total++

		if init.IsDevComplete() {
			metrics.Completed++

			// Track adoption
			if init.DeploymentStatus != nil {
				if init.IsFullyDeployed() {
					metrics.Deployed++
				}
				if init.DeploymentStatus.TotalCustomers > 0 {
					adoption := float64(init.DeploymentStatus.DeployedCustomers) / float64(init.DeploymentStatus.TotalCustomers) * 100
					totalAdoption += adoption
					adoptionCount++
				}
			}
		}
	}

	if adoptionCount > 0 {
		metrics.AvgAdoptionPercent = totalAdoption / float64(adoptionCount)
	}

	return metrics
}

// CalculateSLOCompliance generates SLO compliance records for metrics
// associated with goals in a phase.
func CalculateSLOCompliance(phase *Phase, doc *PRISMDocument) []SLOCompliance {
	if phase == nil || doc == nil {
		return nil
	}

	// Build a map of metric IDs to goals that depend on them
	metricToGoals := make(map[string][]string)
	for _, goal := range doc.Goals {
		if goal.MaturityModel == nil {
			continue
		}
		for _, level := range goal.MaturityModel.Levels {
			for _, sloReq := range level.RequiredSLOs {
				metricToGoals[sloReq.MetricID] = append(metricToGoals[sloReq.MetricID], goal.ID)
			}
		}
	}

	// Generate compliance records
	var compliance []SLOCompliance
	seen := make(map[string]bool)

	for metricID, goalIDs := range metricToGoals {
		if seen[metricID] {
			continue
		}
		seen[metricID] = true

		metric := doc.GetMetricByID(metricID)
		if metric == nil {
			continue
		}

		sloTarget := ""
		if metric.SLO != nil {
			sloTarget = metric.SLO.Target
		}

		compliance = append(compliance, SLOCompliance{
			MetricID:   metricID,
			MetricName: metric.Name,
			SLOTarget:  sloTarget,
			Current:    metric.Current,
			IsMet:      metric.MeetsSLO(),
			GoalIDs:    goalIDs,
		})
	}

	return compliance
}

// CalculatePhaseMetrics generates complete phase metrics.
func CalculatePhaseMetrics(phase *Phase, doc *PRISMDocument) *PhaseMetrics {
	if phase == nil || doc == nil {
		return nil
	}

	metrics := &PhaseMetrics{
		PhaseID:           phase.ID,
		InitiativeMetrics: CalculateInitiativeMetrics(phase, doc),
		SLOCompliance:     CalculateSLOCompliance(phase, doc),
	}

	// Calculate progress for each goal targeted in this phase
	for _, target := range phase.GoalTargets {
		goal := doc.GetGoalByID(target.GoalID)
		if goal == nil {
			continue
		}
		progress := CalculateGoalProgress(goal, phase, doc)
		if progress != nil {
			metrics.GoalProgress = append(metrics.GoalProgress, *progress)
		}
	}

	return metrics
}
