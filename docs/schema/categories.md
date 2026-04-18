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
| Input Validation Coverage | Operations | Build |
| Configuration Validation | Operations | Build |
| Pre-deploy Checks | Operations | Build |
| Schema Validation | Operations | Design |

### JSON Example

```json
{
  "id": "ops-predeploy-checks",
  "name": "Pre-deploy Check Coverage",
  "domain": "operations",
  "stage": "build",
  "category": "prevention",
  "metricType": "coverage",
  "description": "Percentage of deployments with automated pre-deploy validation"
}
```

## Detection

Detection metrics measure the ability to identify issues when they occur.

### Examples

| Metric | Domain | Stage |
|--------|--------|-------|
| Alert Coverage | Operations | Runtime |
| Log Coverage | Operations | Runtime |
| Anomaly Detection | Operations | Runtime |
| Error Rate Monitoring | Operations | Runtime |

### JSON Example

```json
{
  "id": "ops-alert-coverage",
  "name": "Alert Coverage",
  "domain": "operations",
  "stage": "runtime",
  "category": "detection",
  "metricType": "coverage",
  "description": "Percentage of services with alerting configured"
}
```

## Response

Response metrics measure the ability to handle and remediate incidents.

### Examples

| Metric | Domain | Stage |
|--------|--------|-------|
| Mean Time to Recovery | Operations | Response |
| Incident Response Time | Operations | Response |
| Post-Mortem Completion | Operations | Response |
| Rollback Success Rate | Operations | Response |

### JSON Example

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
| P99 Latency | Operations | Runtime |
| Resource Utilization | Operations | Runtime |

### JSON Example

```json
{
  "id": "ops-deploy-frequency",
  "name": "Deployment Frequency",
  "domain": "operations",
  "stage": "build",
  "category": "efficiency",
  "metricType": "count",
  "unit": "/day",
  "trendDirection": "higher_better"
}
```

## Quality

Quality metrics measure code and process quality.

### Examples

| Metric | Domain | Stage |
|--------|--------|-------|
| Change Failure Rate | Operations | Runtime |
| Code Review Coverage | Operations | Build |
| Test Coverage | Operations | Test |
| Documentation Coverage | Operations | Design |

### JSON Example

```json
{
  "id": "ops-change-failure-rate",
  "name": "Change Failure Rate",
  "domain": "operations",
  "stage": "runtime",
  "category": "quality",
  "metricType": "rate",
  "unit": "%",
  "trendDirection": "lower_better"
}
```

## Category by Domain

### Operations Domain (Core)

| Category | Common Use |
|----------|------------|
| Reliability | Availability, durability |
| Efficiency | DORA metrics, performance |
| Quality | Testing, change management |
| Response | MTTR, recovery |
| Prevention | Pre-deploy validation |
| Detection | Monitoring, alerting |

