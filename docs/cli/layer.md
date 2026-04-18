# prism layer

Work with PRISM layers in a document.

## Synopsis

```bash
prism layer [command]
```

## Commands

### prism layer list

List all layers defined in a PRISM document.

```bash
prism layer list <prism-file> [flags]
```

#### Flags

| Flag | Description |
|------|-------------|
| `--json` | Output in JSON format |

#### Examples

```bash
# List layers in a document
prism layer list prism.json

# Output as JSON
prism layer list prism.json --json
```

#### Sample Output

```
Layers:
=======

Code (code)
  Description: Application code, libraries, and dependencies

Infrastructure (infra)
  Description: Cloud resources, networking, and platform services
  Golden Signals:
    - Saturation: metric-infra-cpu

Runtime (runtime)
  Description: Running services, containers, and workloads
  Golden Signals:
    - Latency: metric-runtime-p99
    - Errors: metric-runtime-errors
```

### prism layer show

Show details of a specific layer including associated metrics.

```bash
prism layer show <prism-file> <layer-id> [flags]
```

#### Flags

| Flag | Description |
|------|-------------|
| `--json` | Output in JSON format |

#### Examples

```bash
# Show layer details
prism layer show prism.json runtime

# Output as JSON
prism layer show prism.json runtime --json
```

#### Sample Output

```
Layer: Runtime
ID: runtime
Description: Running services, containers, and workloads

Golden Signals:
  Latency: metric-runtime-p99
  Traffic: metric-runtime-rps
  Errors: metric-runtime-errors
  Saturation: metric-runtime-cpu

Metrics (4):
  - P99 Response Latency [Green]
  - Requests Per Second
  - Error Rate [Yellow]
  - Service CPU Utilization
```

## Default Layers

If no layers are defined in a document, PRISM uses these defaults:

| ID | Name | Description |
|----|------|-------------|
| `code` | Code | Application code, libraries, and dependencies |
| `infra` | Infrastructure | Cloud resources, networking, and platform services |
| `runtime` | Runtime | Running services, containers, and workloads |

## See Also

- [Layers Concept](../concepts/layers.md) - Understanding the three-layer model
- [prism catalog](catalog.md) - List all available constants including layers
