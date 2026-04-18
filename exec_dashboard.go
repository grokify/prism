package prism

import (
	"sort"
)

// ExecutiveDashboard provides a high-level view of maturity progress for executives.
type ExecutiveDashboard struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle,omitempty"`
	GeneratedAt string `json:"generatedAt,omitempty"`

	// Overall summary metrics
	Summary DashboardSummary `json:"summary"`

	// Goal-level maturity scorecard
	MaturityScorecard []GoalMaturityStatus `json:"maturityScorecard"`

	// SLO compliance by category
	SLOCompliance SLOComplianceSummary `json:"sloCompliance"`

	// Phase progress timeline
	PhaseProgress []PhaseProgressSummary `json:"phaseProgress"`

	// Gap analysis - sorted by priority
	Gaps []GapAnalysisEntry `json:"gaps,omitempty"`

	// Benchmark comparison (if available)
	Benchmarks []BenchmarkComparison `json:"benchmarks,omitempty"`
}

// DashboardSummary provides top-level metrics for the executive summary.
type DashboardSummary struct {
	OverallMaturity       float64 `json:"overallMaturity"`       // Average current level
	TargetMaturity        float64 `json:"targetMaturity"`        // Average target level
	MaturityGap           float64 `json:"maturityGap"`           // Target - Current
	SLOCompliancePercent  float64 `json:"sloCompliancePercent"`  // % of SLOs met
	GoalsOnTrack          int     `json:"goalsOnTrack"`          // Goals meeting targets
	GoalsAtRisk           int     `json:"goalsAtRisk"`           // Goals behind
	GoalsTotal            int     `json:"goalsTotal"`            // Total goals
	CurrentPhase          string  `json:"currentPhase"`          // Current phase name
	PhaseCompletionPct    float64 `json:"phaseCompletionPct"`    // Current phase progress
	InitiativesCompleted  int     `json:"initiativesCompleted"`  // Completed initiatives
	InitiativesInProgress int     `json:"initiativesInProgress"` // In-progress initiatives
	InitiativesTotal      int     `json:"initiativesTotal"`      // Total initiatives
}

// GoalMaturityStatus shows a goal's current vs target maturity.
type GoalMaturityStatus struct {
	GoalID         string  `json:"goalId"`
	GoalName       string  `json:"goalName"`
	CurrentLevel   int     `json:"currentLevel"`
	TargetLevel    int     `json:"targetLevel"`
	Gap            int     `json:"gap"`
	Trend          string  `json:"trend"` // "up", "down", "stable"
	SLOsMet        int     `json:"slosMet"`
	SLOsTotal      int     `json:"slosTotal"`
	SLOsMetPercent float64 `json:"slosMetPercent"`
	Status         string  `json:"status"` // "on_track", "at_risk", "behind"
}

// SLOComplianceSummary shows SLO compliance by category.
type SLOComplianceSummary struct {
	Categories        []CategoryCompliance `json:"categories"`
	OverallMet        int                  `json:"overallMet"`
	OverallAtRisk     int                  `json:"overallAtRisk"`
	OverallMissed     int                  `json:"overallMissed"`
	OverallTotal      int                  `json:"overallTotal"`
	OverallCompliance float64              `json:"overallCompliance"`
}

// CategoryCompliance shows SLO status for a specific category.
type CategoryCompliance struct {
	Category    string  `json:"category"`
	Met         int     `json:"met"`
	AtRisk      int     `json:"atRisk"`
	Missed      int     `json:"missed"`
	NotTargeted int     `json:"notTargeted"`
	Total       int     `json:"total"`
	Compliance  float64 `json:"compliance"`
}

// PhaseProgressSummary shows progress for a roadmap phase.
type PhaseProgressSummary struct {
	PhaseID        string  `json:"phaseId"`
	PhaseName      string  `json:"phaseName"`
	Quarter        string  `json:"quarter,omitempty"`
	Year           int     `json:"year,omitempty"`
	StartDate      string  `json:"startDate"`
	EndDate        string  `json:"endDate"`
	Status         string  `json:"status"` // "completed", "in_progress", "planned"
	IsCurrent      bool    `json:"isCurrent"`
	CompletionPct  float64 `json:"completionPct"`
	GoalsTargeted  int     `json:"goalsTargeted"`
	GoalsAchieved  int     `json:"goalsAchieved"`
	InitCompleted  int     `json:"initCompleted"`
	InitInProgress int     `json:"initInProgress"`
	InitTotal      int     `json:"initTotal"`
}

// GapAnalysisEntry identifies a gap between current and target state.
type GapAnalysisEntry struct {
	Category    string  `json:"category"`
	MetricID    string  `json:"metricId"`
	MetricName  string  `json:"metricName"`
	CurrentVal  float64 `json:"currentVal"`
	TargetVal   float64 `json:"targetVal"`
	TargetLevel int     `json:"targetLevel"`
	Gap         float64 `json:"gap"`
	GapPercent  float64 `json:"gapPercent"`
	Priority    string  `json:"priority"` // "critical", "high", "medium", "low"
	GoalName    string  `json:"goalName,omitempty"`
}

// BenchmarkComparison compares metrics against industry benchmarks.
type BenchmarkComparison struct {
	MetricID      string  `json:"metricId"`
	MetricName    string  `json:"metricName"`
	CurrentVal    float64 `json:"currentVal"`
	IndustryAvg   float64 `json:"industryAvg"`
	Top10Percent  float64 `json:"top10Percent"`
	Position      string  `json:"position"` // "below_avg", "avg", "above_avg", "top_10"
	PercentileEst int     `json:"percentileEst"`
}

// GenerateExecutiveDashboard creates an executive dashboard from a PRISM document.
func (doc *PRISMDocument) GenerateExecutiveDashboard() *ExecutiveDashboard {
	dashboard := &ExecutiveDashboard{
		Title:    "Executive Security Dashboard",
		Subtitle: "Maturity & SLO Progress Overview",
	}

	if doc.Metadata != nil && doc.Metadata.Name != "" {
		dashboard.Title = doc.Metadata.Name
	}

	// Generate maturity scorecard
	dashboard.MaturityScorecard = doc.generateMaturityScorecard()

	// Calculate summary metrics
	dashboard.Summary = doc.generateDashboardSummary(dashboard.MaturityScorecard)

	// Generate SLO compliance summary
	dashboard.SLOCompliance = doc.generateSLOCompliance()

	// Generate phase progress
	dashboard.PhaseProgress = doc.generatePhaseProgress()

	// Generate gap analysis
	dashboard.Gaps = doc.generateGapAnalysis()

	return dashboard
}

func (doc *PRISMDocument) generateMaturityScorecard() []GoalMaturityStatus {
	var scorecard []GoalMaturityStatus

	for _, goal := range doc.Goals {
		status := GoalMaturityStatus{
			GoalID:       goal.ID,
			GoalName:     goal.Name,
			CurrentLevel: goal.CurrentLevel,
			TargetLevel:  goal.TargetLevel,
			Gap:          goal.TargetLevel - goal.CurrentLevel,
			Trend:        "stable", // Could be calculated from historical data
		}

		// Calculate SLO compliance for this goal
		if goal.MaturityModel != nil {
			targetLevel := goal.MaturityModel.GetLevel(goal.TargetLevel)
			if targetLevel != nil {
				status.SLOsTotal = len(targetLevel.MetricCriteria)
				for _, criterion := range targetLevel.MetricCriteria {
					// Check if metric meets criterion
					metric := doc.GetMetricByID(criterion.MetricID)
					if metric != nil {
						if criterion.IsMet(metric.Current) {
							status.SLOsMet++
						}
					}
				}
				if status.SLOsTotal > 0 {
					status.SLOsMetPercent = float64(status.SLOsMet) / float64(status.SLOsTotal) * 100
				}
			}
		}

		// Determine status
		if status.Gap <= 0 {
			status.Status = "on_track"
		} else if status.Gap == 1 {
			status.Status = "at_risk"
		} else {
			status.Status = "behind"
		}

		scorecard = append(scorecard, status)
	}

	return scorecard
}

func (doc *PRISMDocument) generateDashboardSummary(scorecard []GoalMaturityStatus) DashboardSummary {
	summary := DashboardSummary{
		GoalsTotal: len(scorecard),
	}

	var totalCurrent, totalTarget float64
	for _, g := range scorecard {
		totalCurrent += float64(g.CurrentLevel)
		totalTarget += float64(g.TargetLevel)

		switch g.Status {
		case "on_track":
			summary.GoalsOnTrack++
		case "at_risk", "behind":
			summary.GoalsAtRisk++
		}
	}

	if len(scorecard) > 0 {
		summary.OverallMaturity = totalCurrent / float64(len(scorecard))
		summary.TargetMaturity = totalTarget / float64(len(scorecard))
		summary.MaturityGap = summary.TargetMaturity - summary.OverallMaturity
	}

	// Count initiatives
	for _, init := range doc.Initiatives {
		summary.InitiativesTotal++
		switch init.Status {
		case "completed":
			summary.InitiativesCompleted++
		case "in_progress":
			summary.InitiativesInProgress++
		}
	}

	// Find current phase
	phases := doc.GetPhasesSorted()
	for _, phase := range phases {
		if phase.Status == "in_progress" {
			summary.CurrentPhase = phase.Name
			if phase.Quarter != "" && phase.Year > 0 {
				summary.CurrentPhase = phase.Quarter + " " + string(rune('0'+phase.Year%10))
			}
			// Calculate phase completion
			view := doc.GeneratePhaseRoadmapView(phase.ID)
			if view != nil {
				summary.PhaseCompletionPct = view.OverallCompletion
			}
			break
		}
	}

	return summary
}

func (doc *PRISMDocument) generateSLOCompliance() SLOComplianceSummary {
	compliance := SLOComplianceSummary{}

	// Group metrics by category
	categoryMetrics := make(map[string][]Metric)
	for _, m := range doc.Metrics {
		categoryMetrics[m.Category] = append(categoryMetrics[m.Category], m)
	}

	// Get all categories sorted
	var categories []string
	for cat := range categoryMetrics {
		categories = append(categories, cat)
	}
	sort.Strings(categories)

	for _, cat := range categories {
		metrics := categoryMetrics[cat]
		catCompliance := CategoryCompliance{
			Category: cat,
			Total:    len(metrics),
		}

		for _, m := range metrics {
			// Check against SLOs
			hasSLO := false
			sloMet := false

			for _, goal := range doc.Goals {
				if goal.MaturityModel == nil {
					continue
				}
				for _, level := range goal.MaturityModel.Levels {
					for _, criterion := range level.MetricCriteria {
						if criterion.MetricID == m.ID {
							hasSLO = true
							if criterion.IsMet(m.Current) {
								sloMet = true
							}
						}
					}
				}
			}

			if !hasSLO {
				catCompliance.NotTargeted++
			} else if sloMet {
				catCompliance.Met++
			} else {
				catCompliance.Missed++
			}
		}

		if catCompliance.Total-catCompliance.NotTargeted > 0 {
			catCompliance.Compliance = float64(catCompliance.Met) / float64(catCompliance.Total-catCompliance.NotTargeted) * 100
		}

		compliance.Categories = append(compliance.Categories, catCompliance)
		compliance.OverallMet += catCompliance.Met
		compliance.OverallAtRisk += catCompliance.AtRisk
		compliance.OverallMissed += catCompliance.Missed
		compliance.OverallTotal += catCompliance.Total - catCompliance.NotTargeted
	}

	if compliance.OverallTotal > 0 {
		compliance.OverallCompliance = float64(compliance.OverallMet) / float64(compliance.OverallTotal) * 100
	}

	return compliance
}

func (doc *PRISMDocument) generatePhaseProgress() []PhaseProgressSummary {
	var progress []PhaseProgressSummary

	phases := doc.GetPhasesSorted()
	for _, phase := range phases {
		summary := PhaseProgressSummary{
			PhaseID:       phase.ID,
			PhaseName:     phase.Name,
			Quarter:       phase.Quarter,
			Year:          phase.Year,
			StartDate:     phase.StartDate,
			EndDate:       phase.EndDate,
			Status:        phase.Status,
			IsCurrent:     phase.Status == "in_progress",
			GoalsTargeted: len(phase.GoalTargets),
		}

		// Count initiatives
		for _, init := range doc.Initiatives {
			if init.PhaseID == phase.ID {
				summary.InitTotal++
				switch init.Status {
				case "completed":
					summary.InitCompleted++
				case "in_progress":
					summary.InitInProgress++
				}
			}
		}

		if summary.InitTotal > 0 {
			summary.CompletionPct = float64(summary.InitCompleted) / float64(summary.InitTotal) * 100
		}

		// Count goals achieved (meeting exit level)
		for _, target := range phase.GoalTargets {
			goal := doc.GetGoalByID(target.GoalID)
			if goal != nil && goal.CurrentLevel >= target.ExitLevel {
				summary.GoalsAchieved++
			}
		}

		progress = append(progress, summary)
	}

	return progress
}

func (doc *PRISMDocument) generateGapAnalysis() []GapAnalysisEntry {
	var gaps []GapAnalysisEntry

	for _, goal := range doc.Goals {
		if goal.MaturityModel == nil {
			continue
		}

		// Get target level criteria
		targetLevel := goal.MaturityModel.GetLevel(goal.TargetLevel)
		if targetLevel == nil {
			continue
		}

		for _, criterion := range targetLevel.MetricCriteria {
			metric := doc.GetMetricByID(criterion.MetricID)
			if metric == nil {
				continue
			}

			if !criterion.IsMet(metric.Current) {
				gap := GapAnalysisEntry{
					Category:    metric.Category,
					MetricID:    metric.ID,
					MetricName:  metric.Name,
					CurrentVal:  metric.Current,
					TargetVal:   criterion.Value,
					TargetLevel: goal.TargetLevel,
					GoalName:    goal.Name,
				}

				// Calculate gap
				gap.Gap = criterion.Value - metric.Current
				if criterion.Operator == SLOOperatorLTE || criterion.Operator == SLOOperatorLT {
					gap.Gap = metric.Current - criterion.Value
				}

				if criterion.Value != 0 {
					gap.GapPercent = (gap.Gap / criterion.Value) * 100
					if gap.GapPercent < 0 {
						gap.GapPercent = -gap.GapPercent
					}
				}

				// Determine priority based on gap size and category
				switch {
				case gap.GapPercent > 50:
					gap.Priority = "critical"
				case gap.GapPercent > 25:
					gap.Priority = "high"
				case gap.GapPercent > 10:
					gap.Priority = "medium"
				default:
					gap.Priority = "low"
				}

				gaps = append(gaps, gap)
			}
		}
	}

	// Sort by priority
	priorityOrder := map[string]int{"critical": 0, "high": 1, "medium": 2, "low": 3}
	sort.Slice(gaps, func(i, j int) bool {
		return priorityOrder[gaps[i].Priority] < priorityOrder[gaps[j].Priority]
	})

	return gaps
}
