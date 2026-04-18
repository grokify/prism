# prism team

Work with PRISM teams in a document.

## Synopsis

```bash
prism team [command]
```

## Commands

### prism team list

List all teams defined in a PRISM document, grouped by team type.

```bash
prism team list <prism-file> [flags]
```

#### Flags

| Flag | Description |
|------|-------------|
| `--json` | Output in JSON format |

#### Examples

```bash
# List teams in a document
prism team list prism.json

# Output as JSON
prism team list prism.json --json
```

#### Sample Output

```
Teams:
======

Stream-Aligned:
  Payments Team (payments-team)
    Services: 2
  User Management Team (users-team)
    Services: 1

Platform:
  Platform Engineering (platform-team)
    Services: 2

Overlay:
  Security Team (security-team)
    Domain: security
  Quality Engineering (qe-team)
    Domain: quality
```

### prism team show

Show details of a specific team including services and accountability.

```bash
prism team show <prism-file> <team-id> [flags]
```

#### Flags

| Flag | Description |
|------|-------------|
| `--json` | Output in JSON format |

#### Examples

```bash
# Show team details
prism team show prism.json payments-team

# Output as JSON
prism team show prism.json payments-team --json
```

#### Sample Output

```
Team: Payments Team
ID: payments-team
Type: Stream-Aligned
Description: Owns the payments domain end-to-end
Owner: Charlie Brown

Contact:
  Slack: #payments

Layer Accountability:
  - code
  - runtime

Services (2):
  - Payments API (payments-api)
  - Payments Worker (payments-worker)
```

## Team Types

| Type | Description |
|------|-------------|
| Stream-Aligned | Teams that build and run services end-to-end |
| Platform | Teams that provide infrastructure as a product |
| Enabling | Teams that help other teams adopt practices |
| Overlay | Teams that define standards across the organization |

## See Also

- [Teams Concept](../concepts/teams.md) - Understanding team topology
- [prism service](service.md) - Work with services
- [prism catalog](catalog.md) - List all available constants including team types
