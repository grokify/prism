# Configuration

This guide covers how to configure PRISM ecosystem documents and the umbrella CLI.

## Ecosystem Configuration

### YAML Format

```yaml
# prism.yaml
ecosystem:
  name: my-organization

  capability:
    files:
      - capabilities/security-stack.json
      - capabilities/platform-stack.json

  maturity:
    files:
      - maturity/operations-model.json
      - maturity/operations-state.json

  roadmap:
    okrs:
      - plans/okrs/q1-2026.json
      - plans/okrs/q2-2026.json
    roadmaps:
      - plans/roadmaps/2026-roadmap.json
```

### JSON Format

```json
{
  "name": "my-organization",
  "capability": {
    "files": [
      "capabilities/security-stack.json",
      "capabilities/platform-stack.json"
    ]
  },
  "maturity": {
    "files": [
      "maturity/operations-model.json",
      "maturity/operations-state.json"
    ]
  },
  "roadmap": {
    "okrs": ["plans/okrs/q1-2026.json"],
    "roadmaps": ["plans/roadmaps/2026-roadmap.json"]
  }
}
```

## Document Schemas

All PRISM documents support JSON Schema validation. Reference the schema in your documents:

### Capability Stack

```json
{
  "$schema": "https://github.com/grokify/prism-capability/schema/capstack.schema.json",
  "metadata": {
    "domain": "security",
    "version": "1.0.0"
  }
}
```

### Maturity Model

```json
{
  "$schema": "https://github.com/grokify/prism-maturity/schema/maturity-model.schema.json",
  "domains": {}
}
```

### PRD

```json
{
  "$schema": "https://github.com/grokify/prism-roadmap/schema/prd.schema.json",
  "metadata": {}
}
```

## Linking Documents

### Capability → SLI Links

Capabilities link to maturity SLIs via `PRISMRef`:

```json
{
  "id": "slo-framework",
  "name": "SLO Framework",
  "prismRef": {
    "sliIds": ["sli-slo-coverage", "sli-slo-attainment"]
  }
}
```

### Roadmap Item → Requirements

Roadmap items reference requirement documents:

```json
{
  "id": "rmi-auth-redesign",
  "title": "Authentication Redesign",
  "requirements": {
    "mrd": "requirements/auth-mrd.json",
    "prd": "requirements/auth-prd.json",
    "trd": "requirements/auth-trd.json"
  }
}
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PRISM_CONFIG` | Default config file path | `prism.yaml` |
| `PRISM_OUTPUT_DIR` | Default output directory | `.` |

## Editor Support

### VS Code

Add JSON Schema associations in `.vscode/settings.json`:

```json
{
  "json.schemas": [
    {
      "fileMatch": ["**/capabilities/*.json"],
      "url": "https://github.com/grokify/prism-capability/schema/capstack.schema.json"
    },
    {
      "fileMatch": ["**/maturity/model*.json"],
      "url": "https://github.com/grokify/prism-maturity/schema/maturity-model.schema.json"
    }
  ]
}
```

## Recommended Directory Structure

```
project/
├── prism.yaml                    # Ecosystem config
├── capabilities/
│   ├── security-stack.json
│   └── platform-stack.json
├── maturity/
│   ├── model.json                # Maturity model definition
│   └── state-q1-2026.json        # Current state snapshot
├── plans/
│   ├── okrs/
│   │   ├── q1-2026.json
│   │   └── q2-2026.json
│   ├── roadmaps/
│   │   └── 2026-roadmap.json
│   └── requirements/
│       ├── feature-a-prd.json
│       └── feature-b-prd.json
└── output/
    └── dashboard.html
```
