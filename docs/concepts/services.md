# Services

PRISM models services as first-class entities that link teams, layers, and metrics together. Services represent deployable units that teams own and operate.

## Service Definition

| Field | Description |
|-------|-------------|
| `id` | Unique identifier |
| `name` | Display name |
| `description` | Service description |
| `ownerTeamId` | Team responsible for this service |
| `layerId` | Primary layer (code, infra, runtime) |
| `metricIds` | Metrics associated with this service |
| `repository` | Git repository URL |
| `tier` | Service tier (tier1, tier2, tier3) |

## Basic Example

```json
{
  "services": [
    {
      "id": "payments-api",
      "name": "Payments API",
      "description": "Core payments processing service",
      "ownerTeamId": "payments-team",
      "layerId": "runtime",
      "tier": "tier1",
      "repository": "https://github.com/example/payments-api"
    }
  ]
}
```

## Service Tiers

Service tiers indicate criticality and SLO expectations:

| Tier | Description | Typical SLO |
|------|-------------|-------------|
| tier1 | Business-critical services | 99.99% |
| tier2 | Important services | 99.9% |
| tier3 | Internal/supporting services | 99% |

## Linking Services to Metrics

Services can be linked to metrics in two ways:

### 1. Service MetricIDs Array

```json
{
  "id": "payments-api",
  "metricIds": ["payments-availability", "payments-latency"]
}
```

### 2. Metric ServiceID Field

```json
{
  "id": "payments-availability",
  "name": "Payments Availability",
  "serviceId": "payments-api",
  "domain": "operations",
  "stage": "runtime"
}
```

## Service Ownership

Services connect to teams through the `ownerTeamId` field:

```json
{
  "teams": [
    {
      "id": "payments-team",
      "name": "Payments Team",
      "type": "stream_aligned",
      "serviceIds": ["payments-api", "payments-worker"]
    }
  ],
  "services": [
    {
      "id": "payments-api",
      "ownerTeamId": "payments-team"
    }
  ]
}
```

## Full Example

```json
{
  "teams": [
    {
      "id": "payments-team",
      "name": "Payments Team",
      "type": "stream_aligned"
    }
  ],
  "services": [
    {
      "id": "payments-api",
      "name": "Payments API",
      "description": "Handles payment processing",
      "ownerTeamId": "payments-team",
      "layerId": "runtime",
      "tier": "tier1",
      "metricIds": ["slo-payments-availability"]
    }
  ],
  "metrics": [
    {
      "id": "slo-payments-availability",
      "name": "Payments Availability",
      "domain": "operations",
      "stage": "runtime",
      "category": "reliability",
      "layer": "runtime",
      "serviceId": "payments-api",
      "metricType": "rate",
      "current": 99.95,
      "target": 99.99,
      "slo": {
        "target": ">=99.99%",
        "operator": "gte",
        "value": 99.99,
        "window": "30d"
      }
    }
  ]
}
```

## CLI Commands

List services in a document:

```bash
prism service list prism.json
```

Show service details with metrics:

```bash
prism service show prism.json payments-api
```

## Best Practices

1. **Define all production services** - Complete service catalog
2. **Assign owner teams** - Clear accountability
3. **Set service tiers** - Differentiate criticality
4. **Link to metrics** - Connect services to their SLOs
5. **Include repository URLs** - Enable code navigation
6. **Use consistent naming** - Service IDs should match deployment names
