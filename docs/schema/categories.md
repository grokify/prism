# Categories

Categories classify metrics by their functional purpose, helping to ensure balanced coverage across different types of controls and capabilities.

## Available Categories

| Category | Constant | Description |
|----------|----------|-------------|
| Prevention | `prevention` | Proactive controls that prevent issues |
| Detection | `detection` | Monitoring and alerting capabilities |
| Response | `response` | Incident handling and remediation |
| Reliability | `reliability` | Availability and durability |
| Efficiency | `efficiency` | Performance and resource utilization |
| Quality | `quality` | Code and process quality |

## Prevention

Prevention metrics measure proactive controls that stop issues before they occur.

### Examples

| Metric | Domain | Stage |
|--------|--------|-------|
| SAST Coverage | Security | Build |
| Threat Modeling Coverage | Security | Design |
| Dependency Scanning | Security | Build |
| Input Validation Coverage | Security | Build |

### JSON Example

```json
{
  "id": "sec-sast-coverage",
  "name": "SAST Coverage",
  "domain": "security",
  "stage": "build",
  "category": "prevention",
  "metricType": "coverage",
  "description": "Percentage of code analyzed by static analysis"
}
```

## Detection

Detection metrics measure the ability to identify issues when they occur.

### Examples

| Metric | Domain | Stage |
|--------|--------|-------|
| Security Monitoring Coverage | Security | Runtime |
| Anomaly Detection Rate | Security | Runtime |
| Log Coverage | Operations | Runtime |
| Alert Coverage | Operations | Runtime |

### JSON Example

```json
{
  "id": "sec-monitoring-coverage",
  "name": "Security Monitoring Coverage",
  "domain": "security",
  "stage": "runtime",
  "category": "detection",
  "metricType": "coverage",
  "description": "Percentage of systems with security monitoring"
}
```

## Response

Response metrics measure the ability to handle and remediate incidents.

### Examples

| Metric | Domain | Stage |
|--------|--------|-------|
| Vulnerability MTTR | Security | Response |
| Incident Response Time | Security | Response |
| Mean Time to Recovery | Operations | Response |
| Post-Mortem Completion | Operations | Response |

### JSON Example

```json
{
  "id": "sec-vuln-mttr",
  "name": "Vulnerability MTTR",
  "domain": "security",
  "stage": "response",
  "category": "response",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "days"
}
```

## Reliability

Reliability metrics measure system availability and durability.

### Examples

| Metric | Domain | Stage |
|--------|--------|-------|
| Service Availability | Operations | Runtime |
| Data Durability | Operations | Runtime |
| Backup Success Rate | Operations | Runtime |
| Failover Success Rate | Operations | Response |

### JSON Example

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
    "value": 99.99
  }
}
```

## Efficiency

Efficiency metrics measure performance and resource utilization.

### Examples

| Metric | Domain | Stage |
|--------|--------|-------|
| Deployment Frequency | Operations | Build |
| Lead Time for Changes | Operations | Build |
| Resource Utilization | Operations | Runtime |
| Cost per Transaction | Operations | Runtime |

### JSON Example

```json
{
  "id": "ops-deploy-frequency",
  "name": "Deployment Frequency",
  "domain": "operations",
  "stage": "build",
  "category": "efficiency",
  "metricType": "rate",
  "unit": "deploys/day",
  "trendDirection": "higher_better"
}
```

## Quality

Quality metrics measure code and process quality.

### Examples

| Metric | Domain | Stage |
|--------|--------|-------|
| Change Failure Rate | Operations | Build |
| Code Review Coverage | Operations | Build |
| Test Coverage | Operations | Test |
| Documentation Coverage | Operations | Design |

### JSON Example

```json
{
  "id": "ops-change-failure-rate",
  "name": "Change Failure Rate",
  "domain": "operations",
  "stage": "build",
  "category": "quality",
  "metricType": "rate",
  "unit": "%",
  "trendDirection": "lower_better"
}
```

## Category by Domain

### Security Domain

| Category | Common Use |
|----------|------------|
| Prevention | SAST, SCA, threat modeling |
| Detection | Monitoring, alerting, DAST |
| Response | Remediation, incident handling |

### Operations Domain

| Category | Common Use |
|----------|------------|
| Reliability | Availability, durability |
| Efficiency | DORA metrics, performance |
| Quality | Testing, change management |
| Response | MTTR, recovery |
