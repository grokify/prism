# prism ecosystem

Cross-module ecosystem operations for loading and querying across PRISM modules.

## Commands

### load

Load an ecosystem from a configuration file or directory:

```bash
prism ecosystem load [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--config`, `-c` | Configuration file (JSON/YAML) |
| `--dir`, `-d` | Directory to load from |
| `--json` | Output statistics as JSON |

**Examples:**

```bash
# Load from config file
prism ecosystem load --config prism.yaml

# Load from directory structure
prism ecosystem load --dir ./ecosystem

# Output as JSON
prism ecosystem load --config prism.yaml --json
```

## Directory Structure

When using `--dir`, the expected structure is:

```
ecosystem/
├── capability/
│   ├── security-stack.json
│   └── platform-stack.json
├── maturity/
│   ├── model.json
│   └── state.json
└── roadmap/
    ├── okrs/
    │   └── q1-2026.json
    └── roadmaps/
        └── 2026-roadmap.json
```

## Configuration File

```yaml
# prism.yaml
ecosystem:
  name: my-organization

  capability:
    files:
      - capabilities/security.json
      - capabilities/platform.json

  maturity:
    files:
      - maturity/model.json
      - maturity/state.json

  roadmap:
    okrs:
      - plans/okrs/q1-2026.json
    roadmaps:
      - plans/roadmaps/2026.json
```

## Output

### Text Output (default)

```
Ecosystem: my-organization

Capability:
  Stacks:       2
  Capabilities: 45

Maturity:
  Documents:    2
  Metrics:      28
  Services:     12
  Initiatives:  8

Roadmap:
  OKR Sets:     1
  Objectives:   5
  Roadmaps:     1
  Phases:       4

By Domain:
  security: 20
  platform: 25

By Status:
  active: 35
  planned: 10
```

### JSON Output

```bash
prism ecosystem load --config prism.yaml --json
```

```json
{
  "capabilityStacks": 2,
  "totalCapabilities": 45,
  "prismDocuments": 2,
  "totalMetrics": 28,
  "totalServices": 12,
  "totalInitiatives": 8,
  "totalOkrSets": 1,
  "totalObjectives": 5,
  "totalRoadmaps": 1,
  "totalPhases": 4,
  "byStatus": {
    "active": 35,
    "planned": 10
  },
  "byDomain": {
    "security": 20,
    "platform": 25
  }
}
```

## Cross-Module Queries

The ecosystem loader enables cross-module queries:

- **Capability Context**: Get a capability with its linked SLIs/metrics
- **Validation**: Validate cross-references (e.g., capability → SLI links)
- **Statistics**: Aggregate counts across all modules

## Validation

When loading an ecosystem, cross-references are validated:

```
Validation errors:
  - capability.slo-framework.prismRef.sliIds: references non-existent metric/SLI "sli-missing"
```

Validation checks:

- Capability `PRISMRef.SLIIDs` reference existing metrics
- Each capability stack is internally valid
- Each PRISM document is internally valid
