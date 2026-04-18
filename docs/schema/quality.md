# Quality Domain

The quality domain covers code quality, testing effectiveness, and defect management metrics. Quality Engineering (QE) teams typically own standards in this domain.

## Quality Domain Constant

```go
const DomainQuality = "quality"
```

## Quality Metrics Categories

| Category | Description | Example Metrics |
|----------|-------------|-----------------|
| Prevention | Defect prevention | Code review coverage, linting compliance |
| Detection | Defect detection | Test coverage, mutation score |
| Quality | Code quality | Technical debt, complexity |
| Response | Issue response | Bug fix time, regression rate |

## ISO 25010 Quality Verticals

PRISM supports ISO 25010 quality characteristics for fine-grained classification:

| Vertical | Constant | Description |
|----------|----------|-------------|
| Functional | `functional` | Correctness, completeness |
| Reliability | `reliability` | Availability, fault tolerance |
| Performance | `performance` | Time behavior, resource utilization |
| Security | `security` | Confidentiality, integrity |
| Usability | `usability` | Learnability, accessibility |
| Maintainability | `maintainability` | Modularity, reusability |

## Example Quality Metrics

### Test Effectiveness

```json
{
  "id": "qa-test-effectiveness",
  "name": "Test Effectiveness",
  "domain": "quality",
  "stage": "test",
  "category": "detection",
  "layer": "code",
  "qualityVertical": "functional",
  "metricType": "rate",
  "description": "Defect Detection Percentage (DDP)",
  "baseline": 60,
  "current": 75,
  "target": 90
}
```

### Code Coverage

```json
{
  "id": "qa-code-coverage",
  "name": "Unit Test Coverage",
  "domain": "quality",
  "stage": "test",
  "category": "detection",
  "layer": "code",
  "qualityVertical": "functional",
  "metricType": "coverage",
  "baseline": 60,
  "current": 82,
  "target": 80,
  "slo": {
    "target": ">=80%",
    "operator": "gte",
    "value": 80
  }
}
```

### Mutation Score

```json
{
  "id": "qa-mutation-score",
  "name": "Mutation Testing Score",
  "domain": "quality",
  "stage": "test",
  "category": "detection",
  "layer": "code",
  "qualityVertical": "functional",
  "metricType": "score",
  "description": "Percentage of mutants killed by tests",
  "current": 65,
  "target": 80
}
```

### Technical Debt

```json
{
  "id": "qa-tech-debt",
  "name": "Technical Debt Ratio",
  "domain": "quality",
  "stage": "build",
  "category": "quality",
  "layer": "code",
  "qualityVertical": "maintainability",
  "metricType": "ratio",
  "trendDirection": "lower_better",
  "current": 8.5,
  "target": 5.0
}
```

### Defect Density

```json
{
  "id": "qa-defect-density",
  "name": "Defect Density",
  "domain": "quality",
  "stage": "runtime",
  "category": "quality",
  "layer": "code",
  "qualityVertical": "reliability",
  "metricType": "ratio",
  "unit": "defects/KLOC",
  "trendDirection": "lower_better",
  "current": 2.3,
  "target": 1.0
}
```

## Quality vs Operations

| Aspect | Operations | Quality |
|--------|------------|---------|
| Focus | Running systems | Code quality |
| Key metrics | Availability, latency | Coverage, defects |
| Stage emphasis | Runtime, Response | Build, Test |
| Teams | SRE, Platform | QE, Development |

## Quality Team as Overlay

The QE team typically operates as an overlay team:

```json
{
  "id": "qe-team",
  "name": "Quality Engineering",
  "type": "overlay",
  "domain": "quality",
  "description": "Defines quality standards and testing practices"
}
```

## Best Practices

1. **Use quality verticals** - Align with ISO 25010 for standard taxonomy
2. **Cover all test types** - Unit, integration, E2E, performance
3. **Track trends** - Defect rates, coverage changes over time
4. **Link to code layer** - Most quality metrics belong in code layer
5. **Define SLOs** - Set targets for coverage, mutation scores
6. **Connect to goals** - Quality improvement as maturity goals
