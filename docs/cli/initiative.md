# prism initiative

Work with PRISM initiatives - improvement projects linked to goals and phases.

## Synopsis

```bash
prism initiative <subcommand> <prism-file> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `list` | List all initiatives |
| `show` | Show details of a specific initiative |

## prism initiative list

List all initiatives in a PRISM document, grouped by status, phase, or goal.

### Usage

```bash
prism initiative list <prism-file> [flags]
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--format` | `-f` | Output format: `text`, `json` (default: `text`) |
| `--by-phase` | | Group initiatives by phase |
| `--by-goal` | | Group initiatives by goal |

### Examples

```bash
# List initiatives by status (default)
prism initiative list prism.json

# List initiatives grouped by phase
prism initiative list prism.json --by-phase

# List initiatives grouped by goal
prism initiative list prism.json --by-goal

# Output as JSON
prism initiative list prism.json -f json
```

### Output (by status)

```
Initiatives by Status
=====================

In Progress (1):
  Observability Platform [in_progress] (90% dev, 90% deployed)
    Owner: SRE Team

Planned (2):
  SLO Dashboard [planned]
    Owner: Platform Team
  Deployment Automation [planned]
    Owner: Platform Team

Completed (1):
  CI/CD Pipeline Enhancement [completed] (100% dev, 100% deployed)
    Owner: Platform Team

Total: 4 initiatives
```

### Output (by phase)

```
Initiatives by Phase
====================

Q1 2026 (2 initiatives):
  CI/CD Pipeline Enhancement [completed] (100% dev, 100% deployed)
    Owner: Platform Team
  Observability Platform [in_progress] (90% dev, 90% deployed)
    Owner: SRE Team

Q2 2026 (2 initiatives):
  SLO Dashboard [planned]
    Owner: Platform Team
  Deployment Automation [planned]
    Owner: Platform Team

Total: 4 initiatives
```

## prism initiative show

Show detailed information about a specific initiative.

### Usage

```bash
prism initiative show <prism-file> <initiative-id> [flags]
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--format` | `-f` | Output format: `text`, `json` (default: `text`) |

### Examples

```bash
# Show initiative details
prism initiative show prism.json init-monitoring

# Output as JSON
prism initiative show prism.json init-monitoring -f json
```

### Output

```
Initiative: Observability Platform
ID: init-monitoring
Description: Deploy comprehensive monitoring, logging, and alerting
Status: In Progress
Priority: 2
Owner: SRE Team

Progress:
  Dev Completion: 90%

Deployment:
  Status: In Progress
  Customers: 45/50 (90%)

Linked Goals (1):
  - Achieve High Reliability (M3 → M5)

Phase: Q1 2026

Linked Metrics (3):
  - Service Availability [SLO Met]
  - Mean Time to Recovery [SLO Not Met]
  - P99 Latency [SLO Met]
```

## Initiative Properties

Initiatives track improvement projects with the following properties:

| Property | Description |
|----------|-------------|
| `id` | Unique identifier |
| `name` | Display name |
| `description` | Detailed description |
| `status` | Current status (planned, not_started, in_progress, completed, cancelled) |
| `priority` | Priority level (1 = highest) |
| `owner` | Responsible person or team |
| `team` | Owning team |
| `goalIds` | Linked goal IDs |
| `phaseId` | Phase this initiative belongs to |
| `metricIds` | Metrics this initiative impacts |
| `devCompletionPercent` | Development completion (0-100%) |
| `deploymentStatus` | B2B deployment tracking |

## Deployment Status

For B2B SaaS, initiatives track customer deployment:

| Property | Description |
|----------|-------------|
| `status` | Deployment status |
| `totalCustomers` | Total customers to deploy to |
| `deployedCustomers` | Customers already deployed |
| `adoptionPercent` | Adoption percentage |

## Related Commands

- [`prism goal`](goal.md) - View goals that initiatives support
- [`prism phase`](phase.md) - View phases that contain initiatives
- [`prism roadmap`](roadmap.md) - View roadmap with initiative progress
