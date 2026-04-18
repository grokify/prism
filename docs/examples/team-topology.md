# Team Topology Example

This example demonstrates using team topology patterns with services, layers, and cross-domain metrics.

## Document Overview

The example includes:

- Teams of different types (stream-aligned, platform, overlay)
- Services owned by teams
- Metrics across operations, security, and quality domains
- Layer accountability assignments

## Key Concepts

### Team Types

The example includes all four team topology types:

| Type | Example Team | Purpose |
|------|--------------|---------|
| Overlay | Security Team, QE Team | Define domain standards |
| Platform | Platform Engineering | Provide infrastructure |
| Stream-Aligned | Payments Team, Users Team | Build and run services |

### Service Ownership

Stream-aligned teams own services:

```json
{
  "id": "payments-team",
  "name": "Payments Team",
  "type": "stream_aligned",
  "serviceIds": ["payments-api", "payments-worker"]
}
```

### Layer Accountability

Teams declare which layers they're accountable for:

```json
{
  "id": "platform-team",
  "layerAccountability": ["infra"]
}
```

## Services

| Service | Owner | Layer | Tier |
|---------|-------|-------|------|
| kubernetes-platform | Platform Engineering | infra | tier1 |
| observability-stack | Platform Engineering | infra | tier1 |
| payments-api | Payments Team | runtime | tier1 |
| payments-worker | Payments Team | runtime | tier2 |
| users-api | User Management Team | runtime | tier1 |

## Download

[team-topology.json](https://github.com/grokify/prism/blob/main/examples/team-topology.json)

## See Also

- [Teams Concept](../concepts/teams.md) - Understanding team topology
- [Services Concept](../concepts/services.md) - Understanding the service model
- [prism team](../cli/team.md) - CLI commands for teams
- [prism service](../cli/service.md) - CLI commands for services
