package prism

// Domain constants represent the three primary domains in PRISM.
const (
	DomainSecurity   = "security"
	DomainOperations = "operations"
	DomainQuality    = "quality"
)

// AllDomains returns all valid domain values.
func AllDomains() []string {
	return []string{DomainSecurity, DomainOperations, DomainQuality}
}

// Layer constants represent ownership boundaries in the stack.
const (
	LayerCode    = "code"
	LayerInfra   = "infra"
	LayerRuntime = "runtime"
)

// AllLayers returns all valid layer values.
func AllLayers() []string {
	return []string{LayerCode, LayerInfra, LayerRuntime}
}

// QualityVertical constants based on ISO 25010 quality characteristics.
const (
	QualityVerticalFunctional      = "functional"
	QualityVerticalReliability     = "reliability"
	QualityVerticalPerformance     = "performance"
	QualityVerticalSecurity        = "security"
	QualityVerticalUsability       = "usability"
	QualityVerticalMaintainability = "maintainability"
)

// AllQualityVerticals returns all valid ISO 25010 quality vertical values.
func AllQualityVerticals() []string {
	return []string{
		QualityVerticalFunctional,
		QualityVerticalReliability,
		QualityVerticalPerformance,
		QualityVerticalSecurity,
		QualityVerticalUsability,
		QualityVerticalMaintainability,
	}
}

// Lifecycle stage constants represent stages in the software delivery lifecycle.
const (
	StageDesign   = "design"
	StageBuild    = "build"
	StageTest     = "test"
	StageRuntime  = "runtime"
	StageResponse = "response"
)

// AllStages returns all valid stage values.
func AllStages() []string {
	return []string{StageDesign, StageBuild, StageTest, StageRuntime, StageResponse}
}

// Category constants represent metric categories.
const (
	CategoryPrevention  = "prevention"
	CategoryDetection   = "detection"
	CategoryResponse    = "response"
	CategoryReliability = "reliability"
	CategoryEfficiency  = "efficiency"
	CategoryQuality     = "quality"
)

// AllCategories returns all valid category values.
func AllCategories() []string {
	return []string{
		CategoryPrevention,
		CategoryDetection,
		CategoryResponse,
		CategoryReliability,
		CategoryEfficiency,
		CategoryQuality,
	}
}

// Maturity level constants represent the 5-level maturity model.
const (
	MaturityLevel1 = 1 // Reactive
	MaturityLevel2 = 2 // Basic
	MaturityLevel3 = 3 // Defined
	MaturityLevel4 = 4 // Managed
	MaturityLevel5 = 5 // Optimizing
)

// MaturityLevelName returns the name for a maturity level.
func MaturityLevelName(level int) string {
	switch level {
	case MaturityLevel1:
		return "Reactive"
	case MaturityLevel2:
		return "Basic"
	case MaturityLevel3:
		return "Defined"
	case MaturityLevel4:
		return "Managed"
	case MaturityLevel5:
		return "Optimizing"
	default:
		return ""
	}
}

// Customer awareness state constants.
const (
	AwarenessUnaware          = "unaware"
	AwarenessAwareNotActing   = "aware_not_remediating"
	AwarenessAwareRemediating = "aware_remediating"
	AwarenessAwareRemediated  = "aware_remediated"
)

// AllAwarenessStates returns all valid awareness state values.
func AllAwarenessStates() []string {
	return []string{
		AwarenessUnaware,
		AwarenessAwareNotActing,
		AwarenessAwareRemediating,
		AwarenessAwareRemediated,
	}
}

// Framework constants for external framework mappings.
const (
	FrameworkNISTCSF     = "NIST_CSF"
	FrameworkNIST80053   = "NIST_800_53"
	FrameworkMITREATTACK = "MITRE_ATTACK"
	FrameworkDORA        = "DORA"
	FrameworkSRE         = "SRE"
	FrameworkFEDRAMP     = "FEDRAMP"
)

// AllFrameworks returns all valid framework values.
func AllFrameworks() []string {
	return []string{
		FrameworkNISTCSF,
		FrameworkNIST80053,
		FrameworkMITREATTACK,
		FrameworkDORA,
		FrameworkSRE,
		FrameworkFEDRAMP,
	}
}

// Metric type constants.
const (
	MetricTypeCoverage     = "coverage"
	MetricTypeRate         = "rate"
	MetricTypeLatency      = "latency"
	MetricTypeRatio        = "ratio"
	MetricTypeCount        = "count"
	MetricTypeDistribution = "distribution"
	MetricTypeScore        = "score"
)

// AllMetricTypes returns all valid metric type values.
func AllMetricTypes() []string {
	return []string{
		MetricTypeCoverage,
		MetricTypeRate,
		MetricTypeLatency,
		MetricTypeRatio,
		MetricTypeCount,
		MetricTypeDistribution,
		MetricTypeScore,
	}
}

// Trend direction constants.
const (
	TrendHigherBetter = "higher_better"
	TrendLowerBetter  = "lower_better"
	TrendTargetValue  = "target_value"
)

// AllTrendDirections returns all valid trend direction values.
func AllTrendDirections() []string {
	return []string{TrendHigherBetter, TrendLowerBetter, TrendTargetValue}
}

// Status constants for metric health.
const (
	StatusGreen  = "Green"
	StatusYellow = "Yellow"
	StatusRed    = "Red"
)

// AllStatuses returns all valid status values.
func AllStatuses() []string {
	return []string{StatusGreen, StatusYellow, StatusRed}
}

// SLO window constants.
const (
	Window7Days  = "7d"
	Window30Days = "30d"
	Window90Days = "90d"
)

// AllWindows returns all valid SLO window values.
func AllWindows() []string {
	return []string{Window7Days, Window30Days, Window90Days}
}
