# prism capability

Capability stack operations for managing capability definitions.

## Commands

### list

List capabilities in a stack:

```bash
prism capability list <stack-file>
prism capability list stack.json --format json
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--format`, `-f` | Output format: `table` (default), `json` |

### show

Show details for a specific capability:

```bash
prism capability show <stack-file> <capability-id>
prism capability show stack.json slo-framework
```

### validate

Validate a capability stack:

```bash
prism capability validate <stack-file>
prism capability validate stack.json
```

Checks:

- Valid JSON structure
- Required fields present
- Layer references resolve
- No duplicate IDs
- Dependency cycles detected

### init

Initialize a new capability stack:

```bash
prism capability init --domain <domain> --output <file>
prism capability init --domain security --output security-stack.json
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--domain`, `-d` | Domain name for the stack |
| `--output`, `-o` | Output file path |

### render

Render a capability stack to different formats:

```bash
prism capability render <stack-file> --format <format>
prism capability render stack.json --format markdown -o capabilities.md
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--format`, `-f` | Output format: `markdown`, `marp`, `html` |
| `--output`, `-o` | Output file path |

## Examples

```bash
# List all capabilities with their layers
prism capability list ops-stack.json

# Show a specific capability with its SLI links
prism capability show ops-stack.json slo-framework

# Validate before committing
prism capability validate ops-stack.json

# Create a new stack
prism capability init --domain platform --output platform-stack.json
```

## Capability Stack Format

```json
{
  "$schema": "https://github.com/grokify/prism-capability/schema/capstack.schema.json",
  "metadata": {
    "domain": "operations",
    "version": "1.0.0"
  },
  "layers": [
    {
      "id": "observe",
      "name": "Observe",
      "order": 1
    }
  ],
  "capabilities": [
    {
      "id": "slo-framework",
      "name": "SLO Framework",
      "layerId": "observe",
      "status": "active",
      "prismRef": {
        "sliIds": ["sli-slo-coverage", "sli-slo-attainment"]
      }
    }
  ]
}
```
