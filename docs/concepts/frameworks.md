# Framework Mappings

PRISM metrics can be mapped to industry frameworks to provide context and enable compliance reporting.

## Supported Frameworks

| Framework | Constant | Description |
|-----------|----------|-------------|
| NIST CSF | `NIST_CSF` | NIST Cybersecurity Framework |
| MITRE ATT&CK | `MITRE_ATTACK` | MITRE ATT&CK Framework |
| DORA | `DORA` | DevOps Research and Assessment |
| SRE | `SRE` | Site Reliability Engineering |

## Framework Mapping Structure

```json
{
  "frameworkMappings": [
    {"framework": "NIST_CSF", "reference": "PR.IP-1"},
    {"framework": "DORA", "reference": "deployment-frequency"}
  ]
}
```

## NIST Cybersecurity Framework

The NIST CSF provides a framework for managing cybersecurity risk.

### Functions

| Function | Description | PRISM Categories |
|----------|-------------|------------------|
| Identify (ID) | Asset management, risk assessment | - |
| Protect (PR) | Access control, training, maintenance | Prevention |
| Detect (DE) | Anomalies, monitoring, detection | Detection |
| Respond (RS) | Response planning, communications | Response |
| Recover (RC) | Recovery planning, improvements | Response |

### Example Mappings

```json
[
  {
    "id": "sec-sast-coverage",
    "name": "SAST Coverage",
    "frameworkMappings": [
      {"framework": "NIST_CSF", "reference": "PR.DS-6"}
    ]
  },
  {
    "id": "sec-monitoring",
    "name": "Security Monitoring Coverage",
    "frameworkMappings": [
      {"framework": "NIST_CSF", "reference": "DE.CM-1"},
      {"framework": "NIST_CSF", "reference": "DE.CM-7"}
    ]
  }
]
```

### Common NIST CSF References

| Reference | Category | Description |
|-----------|----------|-------------|
| PR.DS-6 | Protect | Integrity checking mechanisms |
| PR.IP-1 | Protect | Configuration management |
| DE.CM-1 | Detect | Network monitoring |
| DE.CM-7 | Detect | Monitoring for unauthorized activity |
| RS.AN-1 | Respond | Notifications from detection systems |

## MITRE ATT&CK

MITRE ATT&CK is a knowledge base of adversary tactics and techniques.

### Mapping Approach

Map detection and prevention metrics to ATT&CK techniques:

```json
{
  "id": "sec-endpoint-detection",
  "name": "Endpoint Detection Coverage",
  "frameworkMappings": [
    {"framework": "MITRE_ATTACK", "reference": "T1059"},
    {"framework": "MITRE_ATTACK", "reference": "T1047"}
  ]
}
```

### Example Mappings

| Metric | ATT&CK Technique | Description |
|--------|------------------|-------------|
| Endpoint Detection | T1059 | Command and Scripting Interpreter |
| Network Monitoring | T1071 | Application Layer Protocol |
| File Integrity | T1565 | Data Manipulation |

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
