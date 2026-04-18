# Framework Mappings

PRISM metrics can be mapped to industry frameworks to provide context and enable compliance reporting.

## Supported Frameworks

| Framework | Constant | Description |
|-----------|----------|-------------|
| DORA | `DORA` | DevOps Research and Assessment |
| SRE | `SRE` | Site Reliability Engineering |
| NIST CSF | `NIST_CSF` | NIST Cybersecurity Framework (see prism-security) |
| MITRE ATT&CK | `MITRE_ATTACK` | MITRE ATT&CK Framework (see prism-security) |

## Framework Mapping Structure

```json
{
  "frameworkMappings": [
    {"framework": "NIST_CSF", "reference": "PR.IP-1"},
    {"framework": "DORA", "reference": "deployment-frequency"}
  ]
}
```

## DORA Metrics

DORA (DevOps Research and Assessment) defines four key metrics for software delivery performance.

### DORA Metrics Mapping

| DORA Metric | PRISM Domain | Stage | Category |
|-------------|--------------|-------|----------|
| Deployment Frequency | Operations | Build | Efficiency |
| Lead Time for Changes | Operations | Build | Efficiency |
| Mean Time to Recovery | Operations | Response | Reliability |
| Change Failure Rate | Operations | Build | Quality |

### Example Document

```json
{
  "metrics": [
    {
      "id": "dora-deploy-frequency",
      "name": "Deployment Frequency",
      "domain": "operations",
      "stage": "build",
      "category": "efficiency",
      "metricType": "rate",
      "unit": "deploys/day",
      "current": 5,
      "target": 10,
      "frameworkMappings": [
        {"framework": "DORA", "reference": "deployment-frequency"}
      ]
    },
    {
      "id": "dora-lead-time",
      "name": "Lead Time for Changes",
      "domain": "operations",
      "stage": "build",
      "category": "efficiency",
      "metricType": "latency",
      "unit": "hours",
      "current": 24,
      "target": 1,
      "frameworkMappings": [
        {"framework": "DORA", "reference": "lead-time"}
      ]
    },
    {
      "id": "dora-mttr",
      "name": "Mean Time to Recovery",
      "domain": "operations",
      "stage": "response",
      "category": "reliability",
      "metricType": "latency",
      "unit": "hours",
      "current": 4,
      "target": 1,
      "frameworkMappings": [
        {"framework": "DORA", "reference": "mttr"}
      ]
    },
    {
      "id": "dora-change-failure",
      "name": "Change Failure Rate",
      "domain": "operations",
      "stage": "build",
      "category": "quality",
      "metricType": "rate",
      "unit": "%",
      "current": 10,
      "target": 5,
      "frameworkMappings": [
        {"framework": "DORA", "reference": "change-failure-rate"}
      ]
    }
  ]
}
```

### DORA Performance Levels

| Metric | Elite | High | Medium | Low |
|--------|-------|------|--------|-----|
| Deploy Frequency | On-demand | 1/day-1/week | 1/week-1/month | 1/month-6/month |
| Lead Time | <1 hour | 1 day-1 week | 1 week-1 month | 1-6 months |
| MTTR | <1 hour | <1 day | <1 day | 1 week-1 month |
| Change Failure | 0-15% | 16-30% | 16-30% | 16-30% |

## SRE (Site Reliability Engineering)

SRE practices focus on reliability through SLIs, SLOs, and error budgets.

### SRE Concepts in PRISM

| SRE Concept | PRISM Field | Description |
|-------------|-------------|-------------|
| SLI | `sli` | Service Level Indicator |
| SLO | `slo` | Service Level Objective |
| Error Budget | `thresholds` | Acceptable variance |

### Example SRE Mapping

```json
{
  "id": "sre-availability",
  "name": "Service Availability",
  "domain": "operations",
  "stage": "runtime",
  "category": "reliability",
  "sli": {
    "name": "Availability",
    "formula": "successful_requests / total_requests"
  },
  "slo": {
    "target": ">=99.99%",
    "operator": "gte",
    "value": 99.99,
    "window": "30d"
  },
  "frameworkMappings": [
    {"framework": "SRE", "reference": "availability-slo"}
  ]
}
```

## Using Framework Mappings

### Query by Framework

```go
// Find all metrics mapped to DORA
for _, metric := range doc.Metrics {
    for _, mapping := range metric.FrameworkMappings {
        if mapping.Framework == prism.FrameworkDORA {
            fmt.Printf("%s → %s\n", metric.Name, mapping.Reference)
        }
    }
}
```

### Generate Compliance Reports

Framework mappings enable automated compliance reporting:

1. Extract metrics by framework
2. Calculate coverage per framework category
3. Identify gaps in framework coverage
4. Generate framework-specific reports

## Security Frameworks

For security-specific framework mappings (NIST CSF, MITRE ATT&CK), see [prism-security](https://github.com/grokify/prism-security).
