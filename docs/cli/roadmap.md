# prism roadmap

View roadmap overview and progress across all phases and goals.

## Synopsis

```bash
prism roadmap <command> [options]
```

## Commands

### prism roadmap show

Show the roadmap overview with all phases and goals.

```bash
prism roadmap show [file]
```

**Arguments:**

- `file` - Path to PRISM document (default: `prism.json`)

**Output:**

```
Roadmap: Security & Operations Excellence Roadmap
FY Start: January

Goals:
┌─────────────────────────────┬────────┬──────────┬─────────┬────────┐
│ Goal                        │ Owner  │ Priority │ Current │ Target │
├─────────────────────────────┼────────┼──────────┼─────────┼────────┤
│ Strengthen Security Posture │ CISO   │ 1        │ 3       │ 4      │
│ Operational Excellence      │ VP Eng │ 2        │ 3       │ 4      │
└─────────────────────────────┴────────┴──────────┴─────────┴────────┘

Phases:
┌────────────┬────────────────────┬───────────────┬─────────────┐
│ Phase      │ Period             │ Status        │ Initiatives │
├────────────┼────────────────────┼───────────────┼─────────────┤
│ Q1 2026    │ Jan 1 - Mar 31     │ in_progress   │ 4           │
│ Q2 2026    │ Apr 1 - Jun 30     │ planning      │ 4           │
└────────────┴────────────────────┴───────────────┴─────────────┘
```

### prism roadmap progress

Show progress across all phases and goals.

```bash
prism roadmap progress [file]
```

**Arguments:**

- `file` - Path to PRISM document (default: `prism.json`)

**Output:**

```
Roadmap Progress

Goal: Strengthen Security Posture
Current: Level 3 (Defined) → Target: Level 4 (Managed)

Phase Progress:
┌──────────┬───────┬─────────┬────────┬─────────────┬──────────┐
│ Phase    │ Enter │ Current │ Target │ Initiatives │ SLOs Met │
├──────────┼───────┼─────────┼────────┼─────────────┼──────────┤
│ Q1 2026  │ 2     │ 3       │ 3      │ 2/2 (100%)  │ 2/2 ✓    │
│ Q2 2026  │ 3     │ -       │ 4      │ 0/2 (0%)    │ 2/3      │
└──────────┴───────┴─────────┴────────┴─────────────┴──────────┘

Goal: Achieve Operational Excellence
Current: Level 3 (Defined) → Target: Level 4 (Managed)

Phase Progress:
┌──────────┬───────┬─────────┬────────┬─────────────┬──────────┐
│ Phase    │ Enter │ Current │ Target │ Initiatives │ SLOs Met │
├──────────┼───────┼─────────┼────────┼─────────────┼──────────┤
│ Q1 2026  │ 2     │ 3       │ 3      │ 2/2 (100%)  │ 2/2 ✓    │
│ Q2 2026  │ 3     │ -       │ 4      │ 0/2 (0%)    │ 2/2 ✓    │
└──────────┴───────┴─────────┴────────┴─────────────┴──────────┘

Overall:
  Total Initiatives: 8
  Completed: 4 (50%)
  On Track: 2 goals meeting phase targets
```

## Options

Common options for roadmap commands:

| Option | Description |
|--------|-------------|
| `--json` | Output in JSON format |
| `--goal <id>` | Filter to specific goal |
| `--phase <id>` | Filter to specific phase |

## Examples

Show roadmap overview:

```bash
prism roadmap show examples/goal-roadmap.json
```

Show detailed progress:

```bash
prism roadmap progress examples/goal-roadmap.json
```

Get progress as JSON:

```bash
prism roadmap progress --json
```

Filter to security goal:

```bash
prism roadmap progress --goal goal-security-posture
```

## See Also

- [prism goal](goal.md) - Goal management commands
- [prism phase](phase.md) - Phase management commands
- [Goals Concept](../concepts/goals.md) - Understanding goals
- [Phases Concept](../concepts/phases.md) - Understanding phases
