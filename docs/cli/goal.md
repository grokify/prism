# prism goal

Manage and inspect goals.

## Synopsis

```bash
prism goal <command> [options]
```

## Commands

### prism goal list

List all goals in the document.

```bash
prism goal list [file]
```

**Arguments:**

- `file` - Path to PRISM document (default: `prism.json`)

**Output:**

```
ID                      Name                           Status    Current  Target
goal-reliability        Achieve High Reliability       active    3        5
goal-velocity           Accelerate Delivery Velocity   active    3        4
```

### prism goal show

Show detailed information about a specific goal.

```bash
prism goal show <goal-id> [file]
```

**Arguments:**

- `goal-id` - The ID of the goal to show
- `file` - Path to PRISM document (default: `prism.json`)

**Output:**

```
Goal: Achieve High Reliability (goal-reliability)
Owner: VP Engineering
Priority: 1
Status: active
Target Date: 2026-12-31

Maturity:
  Current Level: 3 (Defined)
  Target Level: 5 (Optimizing)

Maturity Model:
  Level 1 (Reactive): Frequent outages, no SLOs
  Level 2 (Basic): Basic monitoring [ACHIEVED]
  Level 3 (Defined): SLOs tracked [ACHIEVED]
  Level 4 (Managed): Proactive monitoring [TARGET]
  Level 5 (Optimizing): Self-healing systems

Initiatives: 4 total, 2 completed
```

### prism goal progress

Show progress details for a goal including initiatives and SLO compliance.

```bash
prism goal progress <goal-id> [file]
```

**Arguments:**

- `goal-id` - The ID of the goal
- `file` - Path to PRISM document (default: `prism.json`)

**Output:**

```
Goal: Achieve High Reliability
Current Level: 3 → Target: 5

Initiative Progress:
  Completed: 2 / 4 (50%)
  - [✓] CI/CD Pipeline Enhancement (100% deployed)
  - [ ] Observability Platform (90% dev, 90% deployed)
  - [ ] SLO Dashboard (planned)
  - [ ] Deployment Automation (planned)

SLO Compliance for Level 4:
  - [✓] ops-availability: 99.95% (target: >=99.9%)
  - [ ] ops-mttr: 2h (target: <=1h)
  - [✓] ops-p99-latency: 150ms (target: <=200ms)

SLOs Met: 2 / 3 (67%)
```

## Examples

List all goals:

```bash
prism goal list prism.json
```

Show reliability goal details:

```bash
prism goal show goal-reliability examples/goal-roadmap.json
```

Check progress toward target level:

```bash
prism goal progress goal-reliability --json
```

