# Ecosystem API Reference

The `ecosystem` package provides unified loading and querying across all PRISM modules. This page documents the Go API for loading documents and performing cross-module queries.

## Configuration

### Config Structure

The `Config` struct defines all document sources for an ecosystem:

```go
type Config struct {
    Name string `json:"name" yaml:"name"`

    Capability CapabilityConfig `json:"capability" yaml:"capability"`
    Maturity   MaturityConfig   `json:"maturity" yaml:"maturity"`
    Roadmap    RoadmapConfig    `json:"roadmap" yaml:"roadmap"`
}

type CapabilityConfig struct {
    Files []string `json:"files" yaml:"files"`
}

type MaturityConfig struct {
    Files []string `json:"files" yaml:"files"`
}

type RoadmapConfig struct {
    OKRs     []string `json:"okrs" yaml:"okrs"`
    Roadmaps []string `json:"roadmaps" yaml:"roadmaps"`
}
```

Each config section lists file paths to JSON documents of that type:

- `Capability.Files` — capability stack documents (`prism-capability`)
- `Maturity.Files` — PRISM maturity documents (`prism-maturity`)
- `Roadmap.OKRs` — OKR set documents (`prism-roadmap`)
- `Roadmap.Roadmaps` — roadmap documents (`prism-roadmap`)

## Loading

### Load

Load an ecosystem from an in-memory config:

```go
config := ecosystem.Config{
    Name: "platform-team",
    Capability: ecosystem.CapabilityConfig{
        Files: []string{"capabilities/platform.json"},
    },
    Maturity: ecosystem.MaturityConfig{
        Files: []string{"maturity/model.json"},
    },
    Roadmap: ecosystem.RoadmapConfig{
        OKRs:     []string{"plans/okrs/q2-2026.json"},
        Roadmaps: []string{"plans/roadmaps/2026.json"},
    },
}

eco, err := ecosystem.Load(config)
if err != nil {
    log.Fatalf("failed to load ecosystem: %v", err)
}
```

Loading fails with a wrapped error naming the offending file if any document cannot be read or parsed.

### LoadFromFile

Load an ecosystem from a JSON config file:

```go
eco, err := ecosystem.LoadFromFile("prism.json")
if err != nil {
    log.Fatalf("failed to load ecosystem: %v", err)
}
```

### LoadFromDirectory

Load an ecosystem from a conventional directory structure:

```go
eco, err := ecosystem.LoadFromDirectory("examples/platform-team")
if err != nil {
    log.Fatalf("failed to load ecosystem: %v", err)
}
```

Expected directory layout:

```
ecosystem-dir/
├── capability/
│   └── *.json          # capability stack files
├── maturity/
│   └── *.json          # PRISM maturity documents
└── roadmap/
    ├── okrs/
    │   └── *.json      # OKR set files
    └── roadmaps/
        └── *.json      # roadmap files
```

## Capability Queries

### AllCapabilities

Returns all capabilities from all loaded stacks:

```go
caps := eco.AllCapabilities()
for _, cap := range caps {
    fmt.Printf("%s: %s (%s)\n", cap.ID, cap.Name, cap.Status)
}
```

### GetCapabilityByID

Find a capability by ID across all stacks:

```go
cap := eco.GetCapabilityByID("internal-dev-portal")
if cap != nil {
    fmt.Printf("Found: %s\n", cap.Name)
}
```

### CapabilitiesByStatus

Filter capabilities by status:

```go
operational := eco.CapabilitiesByStatus("operational")
inProgress := eco.CapabilitiesByStatus("in-progress")
```

### CapabilitiesByDomain

Filter capabilities by domain (from stack metadata):

```go
securityCaps := eco.CapabilitiesByDomain("security")
```

## Maturity Queries

### AllMetrics

Returns all metrics from all PRISM documents:

```go
metrics := eco.AllMetrics()
```

### GetMetricByID

Find a metric by ID:

```go
metric := eco.GetMetricByID("sli-availability")
```

### AllServices

Returns all services from all PRISM documents:

```go
services := eco.AllServices()
```

### GetServiceByID

Find a service by ID:

```go
svc := eco.GetServiceByID("api-gateway")
```

### AllInitiatives

Returns all initiatives from all PRISM documents:

```go
initiatives := eco.AllInitiatives()
```

### GetInitiativeByID

Find an initiative by ID:

```go
init := eco.GetInitiativeByID("auth-modernization")
```

## Roadmap Queries

### AllObjectives

Returns all objectives from all OKR sets:

```go
objectives := eco.AllObjectives()
```

### GetObjectiveByID

Find an objective by ID:

```go
obj := eco.GetObjectiveByID("improve-developer-velocity")
```

### AllPhases

Returns all phases from all roadmaps:

```go
phases := eco.AllPhases()
```

### GetPhaseByID

Find a phase by ID:

```go
phase := eco.GetPhaseByID("q2-2026")
```

## Cross-Module Queries

### GetCapabilityContext

Returns full context for a capability, joining data across modules:

```go
type CapabilityContext struct {
    Capability *capability.Capability
    Metrics    []maturity.Metric
}

ctx := eco.GetCapabilityContext("internal-dev-portal")
if ctx != nil {
    fmt.Printf("Capability: %s\n", ctx.Capability.Name)
    fmt.Printf("Linked metrics: %d\n", len(ctx.Metrics))
    for _, m := range ctx.Metrics {
        fmt.Printf("  - %s\n", m.ID)
    }
}
```

The context links capabilities to metrics via the capability's `PRISMRef.SLIIDs` field.

## Validation

### Validate

Validates all loaded documents and cross-references:

```go
errs := eco.Validate()
if errs.HasErrors() {
    for _, e := range errs {
        fmt.Printf("[%s/%s] %s: %s\n", e.Module, e.Type, e.ID, e.Message)
    }
}
```

Validation checks:

- Each capability stack's internal consistency
- Each PRISM document's internal consistency
- Cross-references: capability `prismRef.sliIds` must reference existing metrics

### ValidationError

```go
type ValidationError struct {
    Module  string `json:"module"`  // "capability", "maturity"
    Type    string `json:"type"`    // "stack", "document", "capability"
    ID      string `json:"id"`
    Field   string `json:"field"`
    RefID   string `json:"refId,omitempty"`
    Message string `json:"message"`
}
```

## Statistics

### Stats

Returns summary statistics about the loaded ecosystem:

```go
stats := eco.Stats()
fmt.Printf("Capability stacks: %d\n", stats.CapabilityStacks)
fmt.Printf("Total capabilities: %d\n", stats.TotalCapabilities)
fmt.Printf("PRISM documents: %d\n", stats.PRISMDocuments)
fmt.Printf("Total metrics: %d\n", stats.TotalMetrics)
fmt.Printf("Total services: %d\n", stats.TotalServices)
fmt.Printf("Total initiatives: %d\n", stats.TotalInitiatives)
fmt.Printf("OKR sets: %d\n", stats.TotalOKRSets)
fmt.Printf("Roadmaps: %d\n", stats.TotalRoadmaps)

fmt.Println("By status:")
for status, count := range stats.ByStatus {
    fmt.Printf("  %s: %d\n", status, count)
}

fmt.Println("By domain:")
for domain, count := range stats.ByDomain {
    fmt.Printf("  %s: %d\n", domain, count)
}
```

### Stats Structure

```go
type Stats struct {
    CapabilityStacks  int            `json:"capabilityStacks"`
    TotalCapabilities int            `json:"totalCapabilities"`
    PRISMDocuments    int            `json:"prismDocuments"`
    TotalMetrics      int            `json:"totalMetrics"`
    TotalServices     int            `json:"totalServices"`
    TotalInitiatives  int            `json:"totalInitiatives"`
    TotalOKRSets      int            `json:"totalOkrSets"`
    TotalObjectives   int            `json:"totalObjectives"`
    TotalRoadmaps     int            `json:"totalRoadmaps"`
    TotalPhases       int            `json:"totalPhases"`
    ByStatus          map[string]int `json:"byStatus"`
    ByDomain          map[string]int `json:"byDomain"`
}
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/grokify/prism/ecosystem"
)

func main() {
    // Load from directory structure
    eco, err := ecosystem.LoadFromDirectory("examples/platform-team")
    if err != nil {
        log.Fatalf("load failed: %v", err)
    }

    // Print statistics
    stats := eco.Stats()
    fmt.Printf("Loaded %d capabilities across %d stacks\n",
        stats.TotalCapabilities, stats.CapabilityStacks)

    // Validate cross-references
    if errs := eco.Validate(); errs.HasErrors() {
        fmt.Printf("Validation errors: %d\n", len(errs))
        for _, e := range errs {
            fmt.Printf("  [%s] %s: %s\n", e.Module, e.ID, e.Message)
        }
    }

    // Query capabilities
    for _, cap := range eco.CapabilitiesByStatus("in-progress") {
        ctx := eco.GetCapabilityContext(cap.ID)
        fmt.Printf("\n%s (%s)\n", cap.Name, cap.Status)
        if ctx != nil && len(ctx.Metrics) > 0 {
            fmt.Printf("  Linked SLIs: %d\n", len(ctx.Metrics))
        }
    }
}
```

## See Also

- [Configuration Guide](configuration.md) — YAML/JSON config file format
- [PRISM Ecosystem Design](../ECOSYSTEM.md) — architecture and module relationships
- [examples/platform-team](https://github.com/grokify/prism/tree/main/examples/platform-team) — working example
