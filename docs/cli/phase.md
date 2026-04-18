# prism phase

Manage and inspect phases.

## Synopsis

```bash
prism phase <command> [options]
```

## Commands

### prism phase list

List all phases in the document.

```bash
prism phase list [file]
```

**Arguments:**

- `file` - Path to PRISM document (default: `prism.json`)

**Output:**

```
ID              Name        Quarter  Year  Status       Start        End
phase-q1-2026   Q1 2026     Q1       2026  in_progress  2026-01-01   2026-03-31
phase-q2-2026   Q2 2026     Q2       2026  planning     2026-04-01   2026-06-30
```

### prism phase show

Show detailed information about a specific phase.

```bash
prism phase show <phase-id> [file]
```

**Arguments:**

- `phase-id` - The ID of the phase to show
- `file` - Path to PRISM document (default: `prism.json`)

**Output:**

```
Phase: Q1 2026 (phase-q1-2026)
Period: 2026-01-01 to 2026-03-31
Status: in_progress

Goal Targets:
  goal-reliability: Level 2 → Level 3
  goal-velocity: Level 2 → Level 3

Swimlanes:
  Platform Initiatives (operations):
    - init-ci-cd: CI/CD Pipeline Enhancement [completed]
    - init-monitoring: Observability Platform [in_progress]
```

### prism phase metrics

Show enter/exit metrics for a phase.

```bash
prism phase metrics <phase-id> [file]
```

**Arguments:**

- `phase-id` - The ID of the phase
- `file` - Path to PRISM document (default: `prism.json`)

**Output:**

```
Phase: Q1 2026 (phase-q1-2026)

Goal Progress:
┌─────────────────────────┬───────┬─────────┬────────┬─────────────┬──────────┐
│ Goal                    │ Enter │ Current │ Target │ Initiatives │ SLOs Met │
├─────────────────────────┼───────┼─────────┼────────┼─────────────┼──────────┤
│ High Reliability        │ 2     │ 3       │ 3      │ 1/2 (50%)   │ 2/3      │
│ Delivery Velocity       │ 2     │ 3       │ 3      │ 1/2 (50%)   │ 2/3      │
└─────────────────────────┴───────┴─────────┴────────┴─────────────┴──────────┘

Initiative Summary:
  Total: 4
  Dev Complete: 2 (50%)
  Fully Deployed: 1 (25%)
  Avg Adoption: 95%

SLO Compliance:
  ✓ ops-availability: 99.95% (target: >=99.9%)
  ✓ ops-deploy-frequency: 5/day (target: >=5/day)
  ✓ ops-mttr: 2h (target: <=4h)
  ✓ ops-p99-latency: 150ms (target: <=200ms)
```

## Examples

List all phases:

```bash
prism phase list examples/goal-roadmap.json
```

Show Q1 2026 phase details:

```bash
prism phase show phase-q1-2026 examples/goal-roadmap.json
```

Get phase metrics in JSON format:

```bash
prism phase metrics phase-q1-2026 --json
```

## See Also

- [prism goal](goal.md) - Goal management commands
- [prism roadmap](roadmap.md) - Roadmap overview commands
- [Phases Concept](../concepts/phases.md) - Understanding phases

## Domain Extensions

For security-specific phase examples, see [prism-security](https://github.com/grokify/prism-security).
