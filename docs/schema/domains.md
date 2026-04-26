# Domains

PRISM (Platform for Reliability, Improvement, and Strategic Maturity) is an extensible framework that organizes metrics into domains. The core framework provides base domain support, with domain-specific content available through extension modules.

## Core Domains

| Domain | Constant | Description |
|--------|----------|-------------|
| Operations | `operations` | Reliability, performance, and efficiency metrics |
| Security | `security` | Security metrics |
| Quality | `quality` | Code quality, testing, and defect management metrics |

## Operations Domain

The operations domain covers metrics related to reliability, performance, and efficiency. This is the primary domain included in prism core.

### Common Metric Categories

| Category | Example Metrics |
|----------|-----------------|
| Reliability | Availability, durability, error rate |
| Efficiency | Deployment frequency, lead time, resource utilization |
| Quality | Change failure rate, code coverage |
| Response | Mean time to recovery, incident response |

### DORA Alignment

Operations metrics often align with DORA (DevOps Research and Assessment) metrics:

| DORA Metric | PRISM Stage | Category |
|-------------|-------------|----------|
| Deployment Frequency | Build | Efficiency |
| Lead Time for Changes | Build | Efficiency |
| Mean Time to Recovery | Response | Response |
| Change Failure Rate | Runtime | Quality |

### Example Metrics

```json
[
  {
    "id": "ops-availability",
    "name": "Service Availability",
    "domain": "operations",
    "stage": "runtime",
    "category": "reliability",
    "metricType": "rate",
    "unit": "%",
    "current": 99.95,
    "target": 99.99
  },
  {
    "id": "ops-deploy-frequency",
    "name": "Deployment Frequency",
    "domain": "operations",
    "stage": "build",
    "category": "efficiency",
    "metricType": "count",
    "unit": "/day",
    "current": 5,
    "target": 10
  },
  {
    "id": "ops-mttr",
    "name": "Mean Time to Recovery",
    "domain": "operations",
    "stage": "response",
    "category": "response",
    "metricType": "latency",
    "unit": "hours",
    "current": 2,
    "target": 1
  }
]
```

## Security Domain

The security domain covers application security, infrastructure security, and compliance metrics.

## Quality Domain

The quality domain covers code quality, testing effectiveness, and defect management. Quality Engineering (QE) teams typically own standards in this domain.

See [Quality Domain](quality.md) for detailed documentation including ISO 25010 quality verticals.

## Domain Extensibility

PRISM is designed to be extensible. Domain modules can provide:

- **Metrics** - Domain-specific metric definitions
- **Goals** - Strategic objectives with maturity models
- **Dashboards** - Visualization configurations
- **Framework Mappings** - Industry framework alignment

### Using Domain Modules

```bash
# Create operations-focused document
prism init -d operations -o ops.json

# Create security-focused document
prism init -d security -o security.json
```

## Domain Weights

In PRISM score calculation, domains have configurable weights:

| Domain | Default Weight |
|--------|----------------|
| Operations | 40% |
| Security | 40% |
| Quality | 20% |

Weights can be customized in the `ScoreConfig`:

```go
config := &prism.ScoreConfig{
    DomainWeights: map[string]float64{
        "operations": 0.5, // 50% weight
        "security":   0.3, // 30% weight
        "quality":    0.2, // 20% weight
    },
}
```

## Best Practices

1. **Start with Operations** - Use ops metrics as a foundation
2. **Add Domain Modules** - Extend with security or other domain modules as needed
3. **Cover All Stages** - Each domain should have metrics across the lifecycle
4. **Align with Frameworks** - Map metrics to industry frameworks (DORA, SRE, NIST)
5. **Set Appropriate Weights** - Adjust domain weights based on organizational priorities
