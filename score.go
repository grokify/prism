package prism

// PRISMScore represents the composite PRISM score for a document.
type PRISMScore struct {
	Overall            float64     `json:"overall"`
	BaseScore          float64     `json:"baseScore"`
	AwarenessScore     float64     `json:"awarenessScore"`
	SecurityScore      float64     `json:"securityScore"`
	OperationsScore    float64     `json:"operationsScore"`
	CellScores         []CellScore `json:"cellScores,omitempty"`
	Interpretation     string      `json:"interpretation"`
	MaturityAverage    float64     `json:"maturityAverage,omitempty"`
	PerformanceAverage float64     `json:"performanceAverage,omitempty"`
}

// CellScore represents the score for a specific domain/stage cell.
type CellScore struct {
	Domain           string  `json:"domain"`
	Stage            string  `json:"stage"`
	MaturityScore    float64 `json:"maturityScore"`
	PerformanceScore float64 `json:"performanceScore"`
	CellScore        float64 `json:"cellScore"`
	Weight           float64 `json:"weight"`
}

// ScoreConfig configures the PRISM score calculation.
//
// # Weight Normalization Behavior
//
// Cell weights are calculated as: DomainWeight × StageWeight
//
// With default config (domain weights 0.5/0.5, stage weights summing to 1.0),
// each cell weight ranges from 0.075 to 0.15. The total weight across all 10
// cells (2 domains × 5 stages) sums to 0.5 (not 1.0).
//
// The final BaseScore divides weightedSum by totalWeight, which normalizes
// the result. This means:
//
//   - If both domains have equal stage coverage, domain weights are effectively
//     irrelevant (they cancel out in the normalization)
//   - Domain weights only affect the score when domains have different coverage
//   - To make domain weights meaningful, either:
//     (a) Have different numbers of metrics per domain, or
//     (b) Use domain-specific subscores (SecurityScore, OperationsScore)
//
// The awareness multiplier is applied after normalization:
//
//	Overall = BaseScore × AwarenessScore
type ScoreConfig struct {
	MaturityWeight    float64            `json:"maturityWeight"`
	PerformanceWeight float64            `json:"performanceWeight"`
	StageWeights      map[string]float64 `json:"stageWeights"`
	DomainWeights     map[string]float64 `json:"domainWeights"`
}

// DefaultScoreConfig returns the default score configuration.
func DefaultScoreConfig() *ScoreConfig {
	return &ScoreConfig{
		MaturityWeight:    0.4,
		PerformanceWeight: 0.6,
		StageWeights: map[string]float64{
			StageDesign:   0.15,
			StageBuild:    0.20,
			StageTest:     0.15,
			StageRuntime:  0.30,
			StageResponse: 0.20,
		},
		DomainWeights: map[string]float64{
			DomainSecurity:   0.5,
			DomainOperations: 0.5,
		},
	}
}

// GetStageWeight returns the weight for a stage, defaulting to equal weight.
func (c *ScoreConfig) GetStageWeight(stage string) float64 {
	if c.StageWeights == nil {
		return 1.0 / float64(len(AllStages()))
	}
	if w, ok := c.StageWeights[stage]; ok {
		return w
	}
	return 1.0 / float64(len(AllStages()))
}

// GetDomainWeight returns the weight for a domain, defaulting to equal weight.
func (c *ScoreConfig) GetDomainWeight(domain string) float64 {
	if c.DomainWeights == nil {
		return 1.0 / float64(len(AllDomains()))
	}
	if w, ok := c.DomainWeights[domain]; ok {
		return w
	}
	return 1.0 / float64(len(AllDomains()))
}

// CalculatePRISMScore calculates the composite PRISM score for a document.
//
// The score is computed as follows:
//
//  1. For each domain/stage cell, compute a CellScore combining:
//     - MaturityScore: currentLevel / 5 (from maturity model)
//     - PerformanceScore: average ProgressToTarget() of metrics in that cell
//     - CellScore = (MaturityWeight × MaturityScore) + (PerformanceWeight × PerformanceScore)
//
//  2. Compute weighted average of all cell scores:
//     - Each cell has weight = DomainWeight × StageWeight
//     - BaseScore = Σ(CellScore × Weight) / Σ(Weight)
//
//  3. Apply awareness multiplier (if provided):
//     - Overall = BaseScore × AwarenessScore
//     - AwarenessScore ranges from 0.0 (all unaware) to 1.0 (all remediated)
//
//  4. Interpret the score: Elite (≥0.9), Strong (≥0.75), Medium (≥0.5), Weak (≥0.25), Critical (<0.25)
//
// Pass nil for config to use DefaultScoreConfig().
// Pass nil for awareness to skip the awareness multiplier (defaults to 1.0).
func (doc *PRISMDocument) CalculatePRISMScore(config *ScoreConfig, awareness *CustomerAwarenessData) *PRISMScore {
	if config == nil {
		config = DefaultScoreConfig()
	}

	score := &PRISMScore{
		CellScores: make([]CellScore, 0),
	}

	// Calculate awareness score
	if awareness != nil {
		score.AwarenessScore = awareness.AwarenessScore()
	} else {
		score.AwarenessScore = 1.0 // Default to full score if no awareness data
	}

	// Calculate cell scores
	var totalWeight float64
	var weightedSum float64
	var securityWeightedSum, securityWeight float64
	var operationsWeightedSum, operationsWeight float64
	var totalMaturity, totalPerformance float64
	var cellCount int

	for _, domain := range AllDomains() {
		for _, stage := range AllStages() {
			cellScore := doc.calculateCellScore(domain, stage, config)
			score.CellScores = append(score.CellScores, cellScore)

			weight := cellScore.Weight
			totalWeight += weight
			weightedSum += cellScore.CellScore * weight
			totalMaturity += cellScore.MaturityScore
			totalPerformance += cellScore.PerformanceScore
			cellCount++

			if domain == DomainSecurity {
				securityWeight += weight
				securityWeightedSum += cellScore.CellScore * weight
			} else {
				operationsWeight += weight
				operationsWeightedSum += cellScore.CellScore * weight
			}
		}
	}

	// Calculate base score (weighted average of cell scores)
	if totalWeight > 0 {
		score.BaseScore = weightedSum / totalWeight
	}

	// Calculate domain scores
	if securityWeight > 0 {
		score.SecurityScore = securityWeightedSum / securityWeight
	}
	if operationsWeight > 0 {
		score.OperationsScore = operationsWeightedSum / operationsWeight
	}

	// Calculate averages
	if cellCount > 0 {
		score.MaturityAverage = totalMaturity / float64(cellCount)
		score.PerformanceAverage = totalPerformance / float64(cellCount)
	}

	// Apply awareness multiplier to get final score
	score.Overall = score.BaseScore * score.AwarenessScore

	// Interpret the score
	score.Interpretation = InterpretScore(score.Overall)

	return score
}

// calculateCellScore calculates the score for a specific domain/stage cell.
func (doc *PRISMDocument) calculateCellScore(domain, stage string, config *ScoreConfig) CellScore {
	cs := CellScore{
		Domain: domain,
		Stage:  stage,
		Weight: config.GetDomainWeight(domain) * config.GetStageWeight(stage),
	}

	// Get maturity score from maturity model
	if doc.Maturity != nil {
		cell := doc.Maturity.GetCell(domain, stage)
		if cell != nil {
			cs.MaturityScore = cell.CalculateMaturityScore()
		}
	}

	// Calculate performance score from metrics
	metrics := doc.getMetricsForCell(domain, stage)
	if len(metrics) > 0 {
		var progressSum float64
		for _, m := range metrics {
			progressSum += m.ProgressToTarget()
		}
		cs.PerformanceScore = progressSum / float64(len(metrics))
	}

	// Calculate combined cell score
	cs.CellScore = (config.MaturityWeight * cs.MaturityScore) +
		(config.PerformanceWeight * cs.PerformanceScore)

	return cs
}

// getMetricsForCell returns metrics matching the domain and stage.
func (doc *PRISMDocument) getMetricsForCell(domain, stage string) []Metric {
	var result []Metric
	for _, m := range doc.Metrics {
		if m.Domain == domain && m.Stage == stage {
			result = append(result, m)
		}
	}
	return result
}

// InterpretScore returns a human-readable interpretation of the PRISM score.
func InterpretScore(score float64) string {
	switch {
	case score >= 0.9:
		return "Elite"
	case score >= 0.75:
		return "Strong"
	case score >= 0.5:
		return "Medium"
	case score >= 0.25:
		return "Weak"
	default:
		return "Critical"
	}
}

// ScoreBreakdown provides detailed breakdown of score components.
type ScoreBreakdown struct {
	DomainBreakdown map[string]DomainScoreBreakdown `json:"domainBreakdown"`
	StageBreakdown  map[string]StageScoreBreakdown  `json:"stageBreakdown"`
}

// DomainScoreBreakdown breaks down scores by domain.
type DomainScoreBreakdown struct {
	Score         float64 `json:"score"`
	Weight        float64 `json:"weight"`
	MetricCount   int     `json:"metricCount"`
	MaturityLevel float64 `json:"maturityLevel"`
}

// StageScoreBreakdown breaks down scores by stage.
type StageScoreBreakdown struct {
	Score       float64 `json:"score"`
	Weight      float64 `json:"weight"`
	MetricCount int     `json:"metricCount"`
}

// GetScoreBreakdown returns a detailed breakdown of scores.
func (score *PRISMScore) GetScoreBreakdown() *ScoreBreakdown {
	breakdown := &ScoreBreakdown{
		DomainBreakdown: make(map[string]DomainScoreBreakdown),
		StageBreakdown:  make(map[string]StageScoreBreakdown),
	}

	// Aggregate by domain
	domainScores := make(map[string]float64)
	domainWeights := make(map[string]float64)
	domainCounts := make(map[string]int)

	// Aggregate by stage
	stageScores := make(map[string]float64)
	stageWeights := make(map[string]float64)
	stageCounts := make(map[string]int)

	for _, cs := range score.CellScores {
		domainScores[cs.Domain] += cs.CellScore * cs.Weight
		domainWeights[cs.Domain] += cs.Weight
		domainCounts[cs.Domain]++

		stageScores[cs.Stage] += cs.CellScore * cs.Weight
		stageWeights[cs.Stage] += cs.Weight
		stageCounts[cs.Stage]++
	}

	for domain := range domainScores {
		if domainWeights[domain] > 0 {
			breakdown.DomainBreakdown[domain] = DomainScoreBreakdown{
				Score:       domainScores[domain] / domainWeights[domain],
				Weight:      domainWeights[domain],
				MetricCount: domainCounts[domain],
			}
		}
	}

	for stage := range stageScores {
		if stageWeights[stage] > 0 {
			breakdown.StageBreakdown[stage] = StageScoreBreakdown{
				Score:       stageScores[stage] / stageWeights[stage],
				Weight:      stageWeights[stage],
				MetricCount: stageCounts[stage],
			}
		}
	}

	return breakdown
}

// HealthStatus represents overall health based on score.
type HealthStatus struct {
	Level       string  `json:"level"` // Elite, Strong, Medium, Weak, Critical
	Score       float64 `json:"score"` // 0.0-1.0
	Color       string  `json:"color"` // Green, Yellow, Red
	Description string  `json:"description"`
}

// GetHealthStatus returns the health status based on the PRISM score.
func (score *PRISMScore) GetHealthStatus() *HealthStatus {
	hs := &HealthStatus{
		Level: score.Interpretation,
		Score: score.Overall,
	}

	switch {
	case score.Overall >= 0.75:
		hs.Color = StatusGreen
		hs.Description = "System health is good with strong performance across domains"
	case score.Overall >= 0.5:
		hs.Color = StatusYellow
		hs.Description = "System health needs attention in some areas"
	default:
		hs.Color = StatusRed
		hs.Description = "System health requires immediate attention"
	}

	return hs
}
