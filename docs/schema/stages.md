# Lifecycle Stages

PRISM maps metrics to software delivery lifecycle stages, enabling measurement across the entire development and operations process.

## Available Stages

| Stage | Constant | Weight | Description |
|-------|----------|--------|-------------|
| Design | `design` | 15% | Architecture, threat modeling, requirements |
| Build | `build` | 20% | CI/CD, SAST, dependency scanning |
| Test | `test` | 15% | Testing coverage, penetration testing |
| Runtime | `runtime` | 30% | Production monitoring, availability, detection |
| Response | `response` | 20% | Incident response, remediation, recovery |

## Design Stage

The design stage covers early-phase security and architecture decisions.

### Security Metrics

- Threat modeling coverage
- Security requirements documentation
- Architecture review completion

### Operations Metrics

- Capacity planning accuracy
- Architecture documentation
- Design review velocity

### Example

```json
{
  "id": "sec-threat-modeling",
  "name": "Threat Modeling Coverage",
  "domain": "security",
  "stage": "design",
  "category": "prevention",
  "description": "Percentage of new features with completed threat models"
}
```

## Build Stage

The build stage covers CI/CD pipeline and pre-deployment activities.

### Security Metrics

- SAST coverage and findings
- Dependency scanning coverage
- Secrets scanning
- Container image scanning

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
  "frameworkMappings": [
    {"framework": "DORA", "reference": "deployment-frequency"}
  ]
}
```

## Test Stage

The test stage covers testing and validation activities.

### Security Metrics

- Penetration testing coverage
- Security test pass rate
- DAST coverage

### Operations Metrics

- Test coverage percentage
- Integration test pass rate
- Performance test completion

### Example

```json
{
  "id": "sec-pentest-coverage",
  "name": "Penetration Testing Coverage",
  "domain": "security",
  "stage": "test",
  "category": "detection",
  "unit": "%"
}
```

## Runtime Stage

The runtime stage covers production monitoring and detection. This stage typically has the highest weight (30%) as it represents live system health.

### Security Metrics

- Runtime security monitoring
- Anomaly detection rate
- Intrusion detection coverage

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
  "slo": {
    "target": ">=99.99%",
    "operator": "gte",
    "value": 99.99,
    "window": "30d"
  }
}
```

## Response Stage

The response stage covers incident handling and remediation.

### Security Metrics

- Vulnerability MTTR
- Incident response time
- Remediation completion rate

### Operations Metrics

- Mean time to recovery (DORA)
- Incident resolution time
- Post-mortem completion rate

### Example

```json
{
  "id": "ops-mttr",
  "name": "Mean Time to Recovery",
  "domain": "operations",
  "stage": "response",
  "category": "reliability",
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
