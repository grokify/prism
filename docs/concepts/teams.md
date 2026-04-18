# Teams

PRISM supports team topology modeling based on the Team Topologies framework. This enables clear ownership of metrics, services, and standards across the organization.

## Team Types

| Type | Constant | Description |
|------|----------|-------------|
| Stream-Aligned | `stream_aligned` | Teams that build and run services end-to-end |
| Platform | `platform` | Teams that provide infrastructure as a product |
| Enabling | `enabling` | Teams that help other teams adopt practices |
| Overlay | `overlay` | Teams that define standards across the organization |

## Stream-Aligned Teams

Stream-aligned teams are responsible for a flow of work aligned to a single, valuable stream of work.

- Own specific services end-to-end
- Responsible for metrics in their services
- Primary implementers of standards

```json
{
  "id": "payments-team",
  "name": "Payments Team",
  "type": "stream_aligned",
  "serviceIds": ["payments-api", "payments-worker"],
  "layerAccountability": ["code", "runtime"]
}
```

## Platform Teams

Platform teams provide internal services and tooling that accelerate delivery.

- Own infrastructure layer metrics
- Provide self-service capabilities
- Enable stream-aligned teams

```json
{
  "id": "platform-team",
  "name": "Platform Engineering",
  "type": "platform",
  "layerAccountability": ["infra"],
  "serviceIds": ["kubernetes-platform", "observability-stack"]
}
```

## Enabling Teams

Enabling teams help other teams adopt new practices and technologies.

- Focus on capability building
- Work across teams temporarily
- Help improve practices organization-wide

```json
{
  "id": "devex-team",
  "name": "Developer Experience",
  "type": "enabling",
  "domain": "operations"
}
```

## Overlay Teams

Overlay teams define and maintain standards across the organization. Common examples:

- Security team (AppSec, CloudSec)
- Quality Engineering team
- SRE/Reliability team
- Accessibility team

```json
{
  "id": "security-team",
  "name": "Security Team",
  "type": "overlay",
  "domain": "security"
}
```

## Domain Accountability

Overlay teams typically own a domain and define standards that other teams implement:

| Team | Domain | Role |
|------|--------|------|
| Security Team | security | Define security standards, review compliance |
| QE Team | quality | Define quality standards, testing practices |
| SRE Team | operations | Define reliability standards, SLOs |

## RACI Simplification

PRISM keeps ownership simple by using the Team Topologies model instead of complex RACI matrices:

| Role | Team Type |
|------|-----------|
| Accountable (Standards) | Overlay/Enabling |
| Responsible (Implementation) | Stream-Aligned/Platform |

## Example: Full Team Structure

```json
{
  "teams": [
    {
      "id": "security-team",
      "name": "Security Team",
      "type": "overlay",
      "domain": "security",
      "owner": "Jane Smith",
      "slack": "#security"
    },
    {
      "id": "platform-team",
      "name": "Platform Engineering",
      "type": "platform",
      "layerAccountability": ["infra"],
      "owner": "Bob Johnson",
      "slack": "#platform"
    },
    {
      "id": "payments-team",
      "name": "Payments Team",
      "type": "stream_aligned",
      "serviceIds": ["payments-api"],
      "layerAccountability": ["code", "runtime"],
      "owner": "Alice Chen"
    }
  ]
}
```

## CLI Commands

List teams in a document:

```bash
prism team list prism.json
```

Show team details with services:

```bash
prism team show prism.json security-team
```

## Best Practices

1. **Use standard team types** - Align with Team Topologies patterns
2. **Assign domain to overlay teams** - Clarifies standards ownership
3. **Link stream-aligned teams to services** - Establishes service ownership
4. **Define layer accountability** - Specifies which layers the team owns
5. **Keep ownership simple** - Avoid complex RACI, use team types instead
