# Go Library

PRISM provides a Go library for programmatic access to all functionality.

## Installation

```bash
go get github.com/grokify/prism
```

## Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/grokify/prism"
)

func main() {
    // Load document
    data, _ := os.ReadFile("prism.json")
    var doc prism.PRISMDocument
    json.Unmarshal(data, &doc)

    // Validate
    if errs := doc.Validate(); errs.HasErrors() {
        fmt.Println("Validation errors:", errs)
        return
    }

    // Calculate score
    score := doc.CalculatePRISMScore(nil, nil)
    fmt.Printf("PRISM Score: %.1f%% (%s)\n",
        score.Overall*100, score.Interpretation)

    // Check individual metrics
    for _, m := range doc.Metrics {
        status := m.CalculateStatus()
        meetsSLO := m.MeetsSLO()
        fmt.Printf("  %s: %s (SLO met: %v)\n",
            m.Name, status, meetsSLO)
    }
}
```

## Loading Documents

### From File

```go
data, err := os.ReadFile("prism.json")
if err != nil {
    log.Fatal(err)
}

var doc prism.PRISMDocument
if err := json.Unmarshal(data, &doc); err != nil {
    log.Fatal(err)
}
```

### From HTTP

```go
resp, err := http.Get("https://example.com/prism.json")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var doc prism.PRISMDocument
if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
    log.Fatal(err)
}
```

## Validation

### Basic Validation

```go
errs := doc.Validate()
if errs.HasErrors() {
    for _, e := range errs.Errors {
        fmt.Printf("Error at %s: %s\n", e.Field, e.Message)
    }
}
```

### Validation Error Types

```go
type ValidationError struct {
    Field   string // e.g., "metrics[0].domain"
    Message string // e.g., "invalid domain 'sec'"
}

type ValidationErrors struct {
    Errors []ValidationError
}

func (ve *ValidationErrors) HasErrors() bool
func (ve *ValidationErrors) Error() string
```

## Score Calculation

### Basic Score

```go
score := doc.CalculatePRISMScore(nil, nil)

fmt.Printf("Overall: %.1f%%\n", score.Overall*100)
fmt.Printf("Base Score: %.1f%%\n", score.BaseScore*100)
fmt.Printf("Security: %.1f%%\n", score.SecurityScore*100)
fmt.Printf("Operations: %.1f%%\n", score.OperationsScore*100)
```

### Custom Weights

```go
config := &prism.ScoreConfig{
    MaturityWeight:    0.3,
    PerformanceWeight: 0.7,
    StageWeights: map[string]float64{
        prism.StageDesign:   0.10,
        prism.StageBuild:    0.20,
        prism.StageTest:     0.15,
        prism.StageRuntime:  0.35,
        prism.StageResponse: 0.20,
    },
    DomainWeights: map[string]float64{
        prism.DomainSecurity:   0.6,
        prism.DomainOperations: 0.4,
    },
}

score := doc.CalculatePRISMScore(config, nil)
```

### With Awareness Data

```go
awareness := &prism.CustomerAwarenessData{
    Period: "2024-01",
    Distribution: []prism.AwarenessDistribution{
        {State: prism.AwarenessUnaware, Percent: 0.10},
        {State: prism.AwarenessAwareNotActing, Percent: 0.20},
        {State: prism.AwarenessAwareRemediating, Percent: 0.30},
        {State: prism.AwarenessAwareRemediated, Percent: 0.40},
    },
}

score := doc.CalculatePRISMScore(nil, awareness)
fmt.Printf("Awareness Score: %.2f\n", score.AwarenessScore)
```

## Metric Operations

### Calculate Status

```go
for _, m := range doc.Metrics {
    status := m.CalculateStatus()
    fmt.Printf("%s: %s\n", m.Name, status)
}
```

### Check SLO

```go
for _, m := range doc.Metrics {
    if m.MeetsSLO() {
        fmt.Printf("✓ %s meets SLO\n", m.Name)
    } else {
        fmt.Printf("✗ %s does not meet SLO\n", m.Name)
    }
}
```

### Progress to Target

```go
for _, m := range doc.Metrics {
    progress := m.ProgressToTarget()
    fmt.Printf("%s: %.1f%% of target\n", m.Name, progress*100)
}
```

## Maturity Model

### Create Model

```go
// Full model (all domains)
model := prism.NewMaturityModel()

// Domain-filtered model
securityOnly := prism.NewMaturityModelForDomains(
    []string{prism.DomainSecurity},
)
```

### Access Cells

```go
cell := model.GetCell(prism.DomainSecurity, prism.StageBuild)
if cell != nil {
    fmt.Printf("Level: %d\n", cell.CurrentLevel)
    fmt.Printf("Score: %.1f%%\n", cell.CalculateMaturityScore()*100)
}
```

## Constants

### Available Constants

```go
// Domains
prism.DomainSecurity   // "security"
prism.DomainOperations // "operations"
prism.AllDomains()     // []string

// Stages
prism.StageDesign   // "design"
prism.StageBuild    // "build"
prism.StageTest     // "test"
prism.StageRuntime  // "runtime"
prism.StageResponse // "response"
prism.AllStages()   // []string

// Categories
prism.CategoryPrevention  // "prevention"
prism.CategoryDetection   // "detection"
prism.CategoryResponse    // "response"
prism.CategoryReliability // "reliability"
prism.CategoryEfficiency  // "efficiency"
prism.CategoryQuality     // "quality"
prism.AllCategories()     // []string

// Metric Types
prism.MetricTypeCoverage     // "coverage"
prism.MetricTypeRate         // "rate"
prism.MetricTypeLatency      // "latency"
prism.MetricTypeRatio        // "ratio"
prism.MetricTypeCount        // "count"
prism.MetricTypeDistribution // "distribution"
prism.MetricTypeScore        // "score"

// SLO Operators
prism.SLOOperatorGTE // "gte"
prism.SLOOperatorLTE // "lte"
prism.SLOOperatorGT  // "gt"
prism.SLOOperatorLT  // "lt"
prism.SLOOperatorEQ  // "eq"
```

### Validation Functions

```go
if prism.IsValidDomain("security") {
    // valid
}

if prism.IsValidStage("build") {
    // valid
}

if prism.IsValidCategory("prevention") {
    // valid
}
```

## JSON Schema

### Access Embedded Schema

```go
import "github.com/grokify/prism/schema"

schemaJSON := schema.PRISMSchemaJSON()
fmt.Println(string(schemaJSON))
```

## Error Handling

```go
data, err := os.ReadFile("prism.json")
if err != nil {
    log.Fatalf("Failed to read file: %v", err)
}

var doc prism.PRISMDocument
if err := json.Unmarshal(data, &doc); err != nil {
    log.Fatalf("Failed to parse JSON: %v", err)
}

if errs := doc.Validate(); errs.HasErrors() {
    log.Fatalf("Validation failed: %v", errs)
}

score := doc.CalculatePRISMScore(nil, nil)
// Use score...
```
