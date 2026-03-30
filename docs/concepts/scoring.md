# PRISM Score

The PRISM score is a composite health metric (0.0-1.0) that combines maturity levels, metric performance, and customer awareness into a single actionable number.

## Score Formula

```
CellScore = (MaturityWeight × MaturityScore) + (PerformanceWeight × PerformanceScore)
BaseScore = Σ(CellScore × Weight) / Σ(Weight)
Overall = BaseScore × AwarenessScore
```

## Score Components

### 1. Maturity Score (per cell)

The maturity score for each domain/stage cell:

```
MaturityScore = CurrentLevel / 5
```

| Level | Score |
|-------|-------|
| 1 (Reactive) | 0.2 |
| 2 (Basic) | 0.4 |
| 3 (Defined) | 0.6 |
| 4 (Managed) | 0.8 |
| 5 (Optimizing) | 1.0 |

### 2. Performance Score (per cell)

Average `ProgressToTarget()` for metrics in each cell:

```
PerformanceScore = Σ(ProgressToTarget) / MetricCount
```

Where `ProgressToTarget()` returns:

- For `higher_better`: `min(current/target, 1.0)`
- For `lower_better`: `min(target/current, 1.0)` (if current > 0)

### 3. Cell Score

Combined maturity and performance:

```
CellScore = (0.4 × MaturityScore) + (0.6 × PerformanceScore)
```

### 4. Base Score

Weighted average across all cells:

```
BaseScore = Σ(CellScore × Weight) / Σ(Weight)
```

Where `Weight = DomainWeight × StageWeight`

### 5. Awareness Multiplier

Final score adjusted by customer awareness:

```
Overall = BaseScore × AwarenessScore
```

## Default Weights

### Component Weights

| Component | Default Weight |
|-----------|----------------|
| Maturity | 40% |
| Performance | 60% |

### Stage Weights

| Stage | Weight | Rationale |
|-------|--------|-----------|
| Design | 15% | Foundation, less frequent |
| Build | 20% | High automation potential |
| Test | 15% | Validation before production |
| Runtime | 30% | Live system health (highest) |
| Response | 20% | Recovery capability |

### Domain Weights

| Domain | Weight |
|--------|--------|
| Security | 50% |
| Operations | 50% |

## Score Interpretation

| Score Range | Level | Description |
|-------------|-------|-------------|
| ≥ 0.90 | Elite | Industry-leading practices |
| ≥ 0.75 | Strong | Well-managed, proactive approach |
| ≥ 0.50 | Medium | Adequate, room for improvement |
| ≥ 0.25 | Weak | Significant gaps to address |
| < 0.25 | Critical | Immediate attention required |

## Example Calculation

### Setup

- 2 domains × 5 stages = 10 cells
- Security/Build: Maturity L4 (0.8), Performance 0.85
- Operations/Runtime: Maturity L4 (0.8), Performance 0.95

### Cell Score: Security/Build

```
CellScore = (0.4 × 0.8) + (0.6 × 0.85)
          = 0.32 + 0.51
          = 0.83
Weight = 0.5 × 0.20 = 0.10
```

### Cell Score: Operations/Runtime

```
CellScore = (0.4 × 0.8) + (0.6 × 0.95)
          = 0.32 + 0.57
          = 0.89
Weight = 0.5 × 0.30 = 0.15
```

### Base Score (simplified)

```
BaseScore = (0.83 × 0.10) + (0.89 × 0.15) + ...
          / (0.10 + 0.15 + ...)
```

## Customizing Weights

```go
config := &prism.ScoreConfig{
    // Emphasize performance over maturity
    MaturityWeight:    0.3,
    PerformanceWeight: 0.7,

    // Emphasize runtime and response
    StageWeights: map[string]float64{
        "design":   0.10,
        "build":    0.15,
        "test":     0.15,
        "runtime":  0.35,
        "response": 0.25,
    },

    // Emphasize security
    DomainWeights: map[string]float64{
        "security":   0.6,
        "operations": 0.4,
    },
}

score := doc.CalculatePRISMScore(config, nil)
```

## Weight Normalization

Cell weights are calculated as:

```
CellWeight = DomainWeight × StageWeight
```

The final score divides by total weight, which normalizes the result. This means:

- Domain weights only affect score when domains have different coverage
- To make domain weights meaningful, have different numbers of metrics per domain

## Using the Score

### Basic Usage

```go
score := doc.CalculatePRISMScore(nil, nil)
fmt.Printf("PRISM Score: %.1f%% (%s)\n",
    score.Overall*100, score.Interpretation)
```

### Detailed Breakdown

```go
score := doc.CalculatePRISMScore(nil, nil)

fmt.Printf("Overall: %.1f%%\n", score.Overall*100)
fmt.Printf("Base Score: %.1f%%\n", score.BaseScore*100)
fmt.Printf("Security: %.1f%%\n", score.SecurityScore*100)
fmt.Printf("Operations: %.1f%%\n", score.OperationsScore*100)

for _, cs := range score.CellScores {
    fmt.Printf("  %s/%s: %.1f%% (weight: %.2f)\n",
        cs.Domain, cs.Stage, cs.CellScore*100, cs.Weight)
}
```

### Health Status

```go
health := score.GetHealthStatus()
fmt.Printf("Level: %s\n", health.Level)
fmt.Printf("Color: %s\n", health.Color)
fmt.Printf("Description: %s\n", health.Description)
```

## CLI Usage

```bash
# Basic score
prism score prism.json

# Detailed breakdown
prism score prism.json --detailed

# JSON output for automation
prism score prism.json --json
```

## Score Trends

Track PRISM score over time to measure improvement:

| Period | Score | Level |
|--------|-------|-------|
| Q1 2024 | 0.52 | Medium |
| Q2 2024 | 0.61 | Medium |
| Q3 2024 | 0.72 | Medium |
| Q4 2024 | 0.78 | Strong |
