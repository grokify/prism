# Lifecycle Stages

PRISM maps metrics to software delivery lifecycle stages, enabling measurement across the entire development and operations process.

## Available Stages

| Stage | Constant | Weight | Description |
|-------|----------|--------|-------------|
| Design | `design` | 15% | Architecture, requirements, planning |
| Build | `build` | 20% | CI/CD, code quality, automation |
| Test | `test` | 15% | Testing coverage, quality assurance |
| Runtime | `runtime` | 30% | Production monitoring, availability |
| Response | `response` | 20% | Incident response, recovery |

## Design Stage

The design stage covers early-phase architecture and planning decisions.

### Operations Metrics

- Capacity planning accuracy
- Architecture documentation
- Design review velocity
- Requirements coverage

### Example

```json
{
  "id": "ops-design-review",
  "name": "Design Review Coverage",
  "domain": "operations",
  "stage": "design",
  "category": "quality",
  "description": "Percentage of features with completed design reviews"
}
```

## Build Stage

The build stage covers CI/CD pipeline and pre-deployment activities.

### Operations Metrics

- Deployment frequency (DORA)
- Lead time for changes (DORA)
- Build success rate
- Infrastructure as Code coverage

### Example

```json
{
  "id": "ops-deploy-frequency",
  "name": "Deployment Frequency",
  "domain": "operations",
  "stage": "build",
  "category": "efficiency",
  "metricType": "count",
  "unit": "/day",
  "frameworkMappings": [
    {"framework": "DORA", "reference": "deployment-frequency"}
  ]
}
```

## Test Stage

The test stage covers testing and validation activities.

### Operations Metrics

- Test coverage percentage
- Integration test pass rate
- Performance test completion
- End-to-end test coverage

### Example

```json
{
  "id": "ops-test-coverage",
  "name": "Test Coverage",
  "domain": "operations",
  "stage": "test",
  "category": "quality",
  "metricType": "coverage",
  "unit": "%"
}
```

## Runtime Stage

The runtime stage covers production monitoring and detection. This stage typically has the highest weight (30%) as it represents live system health.

### Operations Metrics

- Service availability (SRE)
- P99 latency
- Error rate
- Resource utilization

### Example

```json
{
  "id": "ops-availability",
  "name": "Service Availability",
  "domain": "operations",
  "stage": "runtime",
  "category": "reliability",
  "metricType": "rate",
  "unit": "%",
  "slo": {
    "target": ">=99.99%",
    "operator": "gte",
    "value": 99.99,
    "window": "30d"
  }
}
```

## Response Stage

The response stage covers incident handling and recovery.

### Operations Metrics

- Mean time to recovery (DORA)
- Incident resolution time
- Post-mortem completion rate
- Rollback success rate

### Example

```json
{
  "id": "ops-mttr",
  "name": "Mean Time to Recovery",
  "domain": "operations",
  "stage": "response",
  "category": "response",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "hours"
}
```

## Stage Weights

Stages have default weights in PRISM score calculation:

| Stage | Default Weight | Rationale |
|-------|----------------|-----------|
| Design | 15% | Foundation, but less frequent |
| Build | 20% | High automation opportunity |
| Test | 15% | Validation before production |
| Runtime | 30% | Live system health (highest) |
| Response | 20% | Recovery capability |

Weights can be customized:

```go
config := &prism.ScoreConfig{
    StageWeights: map[string]float64{
        "design":   0.10,
        "build":    0.25,
        "test":     0.15,
        "runtime":  0.30,
        "response": 0.20,
    },
}
```

## Security Stage Examples

For security-specific stage examples (SAST, threat modeling, penetration testing, etc.), see [prism-security](https://github.com/grokify/prism-security).
