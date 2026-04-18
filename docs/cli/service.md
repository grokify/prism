# prism service

Work with PRISM services in a document.

## Synopsis

```bash
prism service [command]
```

## Commands

### prism service list

List all services defined in a PRISM document, grouped by layer.

```bash
prism service list <prism-file> [flags]
```

#### Flags

| Flag | Description |
|------|-------------|
| `--json` | Output in JSON format |

#### Examples

```bash
# List services in a document
prism service list prism.json

# Output as JSON
prism service list prism.json --json
```

#### Sample Output

```
Services:
=========

infra layer:
  Kubernetes Platform (kubernetes-platform)
    Owner: Platform Engineering
    Tier: tier1
  Observability Stack (observability-stack)
    Owner: Platform Engineering
    Tier: tier1

runtime layer:
  Payments API (payments-api)
    Owner: Payments Team
    Tier: tier1
  Payments Worker (payments-worker)
    Owner: Payments Team
    Tier: tier2
  Users API (users-api)
    Owner: User Management Team
    Tier: tier1
```

### prism service show

Show details of a specific service including metrics and ownership.

```bash
prism service show <prism-file> <service-id> [flags]
```

#### Flags

| Flag | Description |
|------|-------------|
| `--json` | Output in JSON format |

#### Examples

```bash
# Show service details
prism service show prism.json payments-api

# Output as JSON
prism service show prism.json payments-api --json
```

#### Sample Output

```
Service: Payments API
ID: payments-api
Description: Core payments processing service
Layer: runtime
Tier: tier1
Repository: https://github.com/example/payments-api

Owner Team: Payments Team (payments-team)

Metrics (2):
  - Payments API Availability [Green]
  - Payments API P99 Latency [Green]

Linked Metrics (2):
  - Payments API Availability [Green]
  - Payments API P99 Latency [Green]
```

## Service Tiers

| Tier | Description | Typical SLO |
|------|-------------|-------------|
| tier1 | Business-critical services | 99.99% |
| tier2 | Important services | 99.9% |
| tier3 | Internal/supporting services | 99% |

## See Also

- [Services Concept](../concepts/services.md) - Understanding the service model
- [prism team](team.md) - Work with teams
- [prism layer](layer.md) - Work with layers
