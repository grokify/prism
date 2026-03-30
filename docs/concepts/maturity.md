# Maturity Model

PRISM uses a 5-level maturity model to assess organizational capability across domains and lifecycle stages.

## Maturity Levels

| Level | Name | Score | Description |
|-------|------|-------|-------------|
| 1 | Reactive | 0.2 | Ad-hoc processes, firefighting mode |
| 2 | Basic | 0.4 | Basic controls, some documentation |
| 3 | Defined | 0.6 | Standardized processes, consistent execution |
| 4 | Managed | 0.8 | Data-driven, measured and controlled |
| 5 | Optimizing | 1.0 | Continuous improvement, automated optimization |

## Level Descriptions

### Level 1: Reactive

- Processes are ad-hoc and chaotic
- Success depends on individual heroics
- No formal documentation
- Firefighting mode is common
- Results are unpredictable

**Security Example**: Vulnerabilities addressed only when exploited
**Operations Example**: Incidents handled without runbooks

### Level 2: Basic

- Basic processes are documented
- Some repeatability exists
- Policies are defined but inconsistently applied
- Manual processes predominate
- Limited metrics collection

**Security Example**: Vulnerability scanning exists but coverage is partial
**Operations Example**: Basic monitoring with manual alerting

### Level 3: Defined

- Standardized processes across the organization
- Consistent execution of practices
- Documentation is maintained
- Roles and responsibilities are clear
- Metrics are collected systematically

**Security Example**: SAST/DAST integrated into all pipelines
**Operations Example**: Standardized incident response procedures

### Level 4: Managed

- Data-driven decision making
- Processes are measured and controlled
- Quantitative quality goals
- Variation is understood and addressed
- Predictable outcomes

**Security Example**: Vulnerability SLOs with automated tracking
**Operations Example**: SLOs with error budgets and automated alerting

### Level 5: Optimizing

- Continuous process improvement
- Innovation and optimization
- Automated remediation where possible
- Proactive risk management
- Industry-leading practices

**Security Example**: Automated vulnerability remediation, zero-day response playbooks
**Operations Example**: Self-healing systems, automated capacity management

## Maturity Matrix

PRISM assesses maturity for each domain/stage combination:

|  | Design | Build | Test | Runtime | Response |
|--|--------|-------|------|---------|----------|
| **Security** | L3 | L4 | L3 | L4 | L3 |
| **Operations** | L3 | L4 | L3 | L4 | L4 |

## Maturity Model Structure

### MaturityModel

```json
{
  "maturity": {
    "levels": [
      {"level": 1, "name": "Reactive", "description": "..."},
      {"level": 2, "name": "Basic", "description": "..."},
      {"level": 3, "name": "Defined", "description": "..."},
      {"level": 4, "name": "Managed", "description": "..."},
      {"level": 5, "name": "Optimizing", "description": "..."}
    ],
    "cells": [...]
  }
}
```

### MaturityCell

Each cell represents a domain/stage intersection:

```json
{
  "domain": "security",
  "stage": "build",
  "currentLevel": 4,
  "targetLevel": 5,
  "primaryKPI": "sec-sast-coverage",
  "kpiTarget": ">=95%"
}
```

## Maturity Score Calculation

The maturity score for a cell is:

```
MaturityScore = CurrentLevel / 5
```

| Level | Score |
|-------|-------|
| 1 | 0.2 |
| 2 | 0.4 |
| 3 | 0.6 |
| 4 | 0.8 |
| 5 | 1.0 |

## Using Maturity in Go

```go
// Create a maturity model
model := prism.NewMaturityModel()

// Get a specific cell
cell := model.GetCell("security", "build")
cell.CurrentLevel = 4

// Calculate maturity score
score := cell.CalculateMaturityScore()
fmt.Printf("Maturity: %.1f%%\n", score*100) // 80%

// Create domain-filtered model
securityOnly := prism.NewMaturityModelForDomains([]string{"security"})
```

## Assessment Guidelines

### Level 1 → 2 (Basic)

- Document existing processes
- Establish basic policies
- Implement foundational tools

### Level 2 → 3 (Defined)

- Standardize across teams
- Create formal procedures
- Establish consistent metrics

### Level 3 → 4 (Managed)

- Define quantitative goals
- Implement continuous monitoring
- Establish feedback loops

### Level 4 → 5 (Optimizing)

- Automate optimization
- Implement predictive capabilities
- Drive continuous improvement

## Maturity Weight in PRISM Score

By default, maturity contributes 40% to the PRISM score:

```go
config := &prism.ScoreConfig{
    MaturityWeight:    0.4, // 40%
    PerformanceWeight: 0.6, // 60%
}
```

The cell score formula is:

```
CellScore = (0.4 × MaturityScore) + (0.6 × PerformanceScore)
```
