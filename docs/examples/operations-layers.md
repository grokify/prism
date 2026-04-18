# Operations with Layers Example

This example demonstrates using the three-layer model (code, infra, runtime) to organize operations metrics.

## Document Overview

The example includes:

- Layer definitions with golden signals
- Metrics tagged with layer assignments
- Coverage across code, infrastructure, and runtime layers

## Key Concepts

### Layer Definitions

Each layer defines golden signals pointing to specific metrics:

```json
{
  "layers": [
    {
      "id": "runtime",
      "name": "Runtime",
      "signals": {
        "latency": "metric-runtime-p99",
        "traffic": "metric-runtime-rps",
        "errors": "metric-runtime-errors",
        "saturation": "metric-runtime-cpu"
      }
    }
  ]
}
```

### Layer-Tagged Metrics

Metrics include a `layer` field for classification:

```json
{
  "id": "metric-code-lint-errors",
  "name": "Linting Errors",
  "domain": "operations",
  "stage": "build",
  "layer": "code"
}
```

## Metrics by Layer

| Layer | Metrics |
|-------|---------|
| Code | Linting errors, outdated dependencies |
| Infra | Compliance rate, CPU utilization |
| Runtime | P99 latency, RPS, error rate, CPU |

## Download

[operations-layers.json](https://github.com/grokify/prism/blob/main/examples/operations-layers.json)

## See Also

- [Layers Concept](../concepts/layers.md) - Understanding the three-layer model
- [prism layer](../cli/layer.md) - CLI commands for layers
