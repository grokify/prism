# Domains

PRISM organizes metrics into two primary domains representing the main areas of B2B SaaS operational health.

## Available Domains

| Domain | Constant | Description |
|--------|----------|-------------|
| Security | `security` | Application and infrastructure security metrics |
| Operations | `operations` | Reliability, performance, and efficiency metrics |

## Security Domain

The security domain covers metrics related to protecting the application and infrastructure:

### Common Metric Categories

| Category | Example Metrics |
|----------|-----------------|
| Prevention | SAST coverage, dependency scanning, threat modeling |
| Detection | Runtime security monitoring, anomaly detection |
| Response | Vulnerability MTTR, incident response time |

### Lifecycle Coverage

| Stage | Focus Areas |
|-------|-------------|
| Design | Threat modeling, security requirements |
| Build | SAST, SCA, secrets scanning |
| Test | Penetration testing, security test coverage |
| Runtime | Runtime protection, monitoring |
| Response | Vulnerability remediation, incident response |

### Example Metrics

```json
[
  {
    "id": "sec-sast-coverage",
    "name": "SAST Coverage",
    "domain": "security",
    "stage": "build",
    "category": "prevention"
  },
  {
    "id": "sec-vuln-mttr",
    "name": "Critical Vulnerability MTTR",
    "domain": "security",
    "stage": "response",
    "category": "response"
  }
]
```

## Operations Domain

The operations domain covers metrics related to reliability, performance, and efficiency:

### Common Metric Categories

| Category | Example Metrics |
|----------|-----------------|
| Reliability | Availability, durability, error rate |
| Efficiency | Deployment frequency, lead time, resource utilization |
| Quality | Change failure rate, code coverage |

### DORA Alignment

Operations metrics often align with DORA (DevOps Research and Assessment) metrics:

| DORA Metric | PRISM Stage | Category |
|-------------|-------------|----------|
| Deployment Frequency | Build | Efficiency |
| Lead Time for Changes | Build | Efficiency |
| Mean Time to Recovery | Response | Reliability |
| Change Failure Rate | Build | Quality |

### Example Metrics

```json
[
  {
    "id": "ops-availability",
    "name": "Service Availability",
    "domain": "operations",
    "stage": "runtime",
    "category": "reliability"
  },
  {
    "id": "ops-deploy-frequency",
    "name": "Deployment Frequency",
    "domain": "operations",
    "stage": "build",
    "category": "efficiency"
  }
]
```

## Domain Weights

In PRISM score calculation, domains have configurable weights:

| Domain | Default Weight |
|--------|----------------|
| Security | 50% |
| Operations | 50% |

Weights can be customized in the `ScoreConfig`:

```go
config := &prism.ScoreConfig{
    DomainWeights: map[string]float64{
        "security":   0.6, // 60% weight
        "operations": 0.4, // 40% weight
    },
}
```

## Best Practices

1. **Balance Coverage** - Include metrics from both domains
2. **Cover All Stages** - Each domain should have metrics across the lifecycle
3. **Align with Frameworks** - Map metrics to industry frameworks (NIST, DORA)
4. **Set Appropriate Weights** - Adjust domain weights based on organizational priorities
