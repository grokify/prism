# Type Reference

This page documents the main types in the PRISM Go library.

## Document Types

### PRISMDocument

The root document type.

```go
type PRISMDocument struct {
    Schema      string          `json:"$schema,omitempty"`
    Metadata    *Metadata       `json:"metadata,omitempty"`
    Metrics     []Metric        `json:"metrics"`
    Maturity    *MaturityModel  `json:"maturity,omitempty"`
    OKRs        []OKRMapping    `json:"okrs,omitempty"`
    Initiatives []Initiative    `json:"initiatives,omitempty"`
}
```

**Methods:**

- `Validate() *ValidationErrors` - Validate the document
- `CalculatePRISMScore(config *ScoreConfig, awareness *CustomerAwarenessData) *PRISMScore` - Calculate composite score

### Metadata

Document metadata.

```go
type Metadata struct {
    Name        string `json:"name,omitempty"`
    Version     string `json:"version,omitempty"`
    Description string `json:"description,omitempty"`
    Owner       string `json:"owner,omitempty"`
    LastUpdated string `json:"lastUpdated,omitempty"`
}
```

## Metric Types

### Metric

A single metric definition.

```go
type Metric struct {
    // Identity
    ID          string `json:"id,omitempty"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`

    // Classification
    Domain   string `json:"domain"`
    Stage    string `json:"stage"`
    Category string `json:"category"`

    // Measurement
    MetricType     string  `json:"metricType"`
    TrendDirection string  `json:"trendDirection,omitempty"`
    Unit           string  `json:"unit,omitempty"`
    Baseline       float64 `json:"baseline,omitempty"`
    Current        float64 `json:"current"`
    Target         float64 `json:"target"`

    // SLI/SLO
    SLI *SLI `json:"sli,omitempty"`
    SLO *SLO `json:"slo,omitempty"`

    // Thresholds
    Thresholds *Thresholds `json:"thresholds,omitempty"`
    Status     string      `json:"status,omitempty"`

    // Mappings
    MaturityMapping   *MaturityMapping   `json:"maturityMapping,omitempty"`
    DMAIC             *DMAICMapping      `json:"dmaic,omitempty"`
    FrameworkMappings []FrameworkMapping `json:"frameworkMappings,omitempty"`

    // Ownership
    Owner      string `json:"owner,omitempty"`
    DataSource string `json:"dataSource,omitempty"`

    // History
    DataPoints []DataPoint `json:"dataPoints,omitempty"`
}
```

**Methods:**

- `CalculateStatus() string` - Calculate status based on thresholds
- `MeetsSLO() bool` - Check if current value meets SLO
- `ProgressToTarget() float64` - Calculate progress toward target (0.0-1.0)

### SLI

Service Level Indicator.

```go
type SLI struct {
    Name        string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Formula     string `json:"formula,omitempty"`
}
```

### SLO

Service Level Objective.

```go
type SLO struct {
    Target     string      `json:"target"`
    Operator   string      `json:"operator,omitempty"`
    Value      float64     `json:"value,omitempty"`
    Window     string      `json:"window,omitempty"`
    Thresholds *Thresholds `json:"thresholds,omitempty"`
}
```

### Thresholds

Status thresholds.

```go
type Thresholds struct {
    Green  float64 `json:"green"`
    Yellow float64 `json:"yellow"`
    Red    float64 `json:"red"`
}
```

### DataPoint

Historical data point.

```go
type DataPoint struct {
    Timestamp time.Time `json:"timestamp"`
    Value     float64   `json:"value"`
    Label     string    `json:"label,omitempty"`
}
```

### FrameworkMapping

Framework reference mapping.

```go
type FrameworkMapping struct {
    Framework string `json:"framework"`
    Reference string `json:"reference"`
}
```

## Maturity Types

### MaturityModel

Maturity model definition.

```go
type MaturityModel struct {
    Levels []MaturityLevelDef `json:"levels"`
    Cells  []MaturityCell     `json:"cells"`
}
```

**Methods:**

- `GetCell(domain, stage string) *MaturityCell` - Get cell by domain/stage

### MaturityLevelDef

Maturity level definition.

```go
type MaturityLevelDef struct {
    Level       int    `json:"level"`
    Name        string `json:"name"`
    Description string `json:"description"`
}
```

### MaturityCell

Maturity assessment for a domain/stage.

```go
type MaturityCell struct {
    Domain       string  `json:"domain"`
    Stage        string  `json:"stage"`
    CurrentLevel int     `json:"currentLevel"`
    TargetLevel  int     `json:"targetLevel,omitempty"`
    PrimaryKPI   string  `json:"primaryKPI,omitempty"`
    KPITarget    string  `json:"kpiTarget,omitempty"`
}
```

**Methods:**

- `CalculateMaturityScore() float64` - Calculate score (0.0-1.0)

## Awareness Types

### CustomerAwarenessData

Customer awareness distribution.

```go
type CustomerAwarenessData struct {
    Period       string                  `json:"period"`
    Distribution []AwarenessDistribution `json:"distribution"`
}
```

**Methods:**

- `UnawareRate() float64` - Percentage unaware
- `AwareNotActingRate() float64` - Percentage aware but not acting
- `RemediationInProgressRate() float64` - Percentage remediating
- `ProactiveResolutionRate() float64` - Percentage remediated
- `ProactiveDetectionRate() float64` - 1 - unaware rate
- `AwarenessScore() float64` - Weighted awareness score

### AwarenessDistribution

Single awareness state distribution.

```go
type AwarenessDistribution struct {
    State   string  `json:"state"`
    Count   int     `json:"count"`
    Percent float64 `json:"percent"`
}
```

## Score Types

### PRISMScore

Composite PRISM score.

```go
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
```

**Methods:**

- `GetScoreBreakdown() *ScoreBreakdown` - Get detailed breakdown
- `GetHealthStatus() *HealthStatus` - Get health status

### CellScore

Score for a domain/stage cell.

```go
type CellScore struct {
    Domain           string  `json:"domain"`
    Stage            string  `json:"stage"`
    MaturityScore    float64 `json:"maturityScore"`
    PerformanceScore float64 `json:"performanceScore"`
    CellScore        float64 `json:"cellScore"`
    Weight           float64 `json:"weight"`
}
```

### ScoreConfig

Score calculation configuration.

```go
type ScoreConfig struct {
    MaturityWeight    float64            `json:"maturityWeight"`
    PerformanceWeight float64            `json:"performanceWeight"`
    StageWeights      map[string]float64 `json:"stageWeights"`
    DomainWeights     map[string]float64 `json:"domainWeights"`
}
```

**Methods:**

- `GetStageWeight(stage string) float64` - Get weight for stage
- `GetDomainWeight(domain string) float64` - Get weight for domain

### HealthStatus

Health status based on score.

```go
type HealthStatus struct {
    Level       string  `json:"level"`
    Score       float64 `json:"score"`
    Color       string  `json:"color"`
    Description string  `json:"description"`
}
```

## Validation Types

### ValidationError

Single validation error.

```go
type ValidationError struct {
    Field   string
    Message string
}
```

### ValidationErrors

Collection of validation errors.

```go
type ValidationErrors struct {
    Errors []ValidationError
}
```

**Methods:**

- `HasErrors() bool` - Check if any errors exist
- `Error() string` - Get error string
- `Add(field, message string)` - Add an error
