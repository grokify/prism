# Layers

PRISM uses a three-layer model to organize metrics by ownership boundaries in the technology stack. This helps clarify accountability and enables targeted improvements at each level.

## The Three Layers

| Layer | Constant | Description |
|-------|----------|-------------|
| Code | `code` | Application code, libraries, and dependencies |
| Infrastructure | `infra` | Cloud resources, networking, and platform services |
| Runtime | `runtime` | Running services, containers, and workloads |

## Layer Definitions

### Code Layer

The code layer encompasses everything related to the application source code:

- Source code quality and coverage
- Dependency management and vulnerabilities
- Static analysis results
- Build artifacts

**Example metrics:**

- Code coverage percentage
- Dependency vulnerability count
- Static analysis findings
- Code complexity scores

### Infrastructure Layer

The infrastructure layer covers the platform and resources that support applications:

- Cloud resource configuration
- Network security controls
- Infrastructure as Code compliance
- Resource utilization

**Example metrics:**

- Infrastructure compliance rate
- Misconfigurations detected
- Resource optimization score
- Infrastructure drift

### Runtime Layer

The runtime layer focuses on deployed services in production:

- Service availability and performance
- Runtime security monitoring
- Observability completeness
- Incident response

**Example metrics:**

- Service availability (SLO)
- P99 latency
- Error rate
- Mean time to recovery

## Golden Signals

Each layer can define golden signals based on Google SRE's four golden signals:

| Signal | Description | Example |
|--------|-------------|---------|
| Latency | Time to serve requests | P99 response time |
| Traffic | Request throughput | Requests per second |
| Errors | Error rate | 5xx error percentage |
| Saturation | Resource utilization | CPU/memory usage |

### Defining Golden Signals

```json
{
  "layers": [
    {
      "id": "runtime",
      "name": "Runtime",
      "signals": {
        "latency": "metric-p99-latency",
        "traffic": "metric-rps",
        "errors": "metric-error-rate",
        "saturation": "metric-cpu-usage"
      }
    }
  ]
}
```

## Layer and Domain Relationship

Layers are orthogonal to domains:

| Dimension | Purpose | Values |
|-----------|---------|--------|
| Domain | Functional area | operations, security, quality |
| Layer | Ownership boundary | code, infra, runtime |

A metric belongs to one domain AND one layer:

```json
{
  "id": "sec-runtime-threats",
  "name": "Runtime Threat Detection",
  "domain": "security",
  "layer": "runtime",
  "stage": "runtime",
  "category": "detection"
}
```

## Layer Accountability

Teams can declare layer accountability:

```json
{
  "teams": [
    {
      "id": "platform-team",
      "name": "Platform Team",
      "type": "platform",
      "layerAccountability": ["infra", "runtime"]
    },
    {
      "id": "app-team",
      "name": "Application Team",
      "type": "stream_aligned",
      "layerAccountability": ["code"]
    }
  ]
}
```

## CLI Commands

List layers in a document:

```bash
prism layer list prism.json
```

Show layer details with associated metrics:

```bash
prism layer show prism.json runtime
```

## Best Practices

1. **Assign every metric to a layer** - Clarifies ownership
2. **Define golden signals** - Standardizes observability per layer
3. **Map teams to layers** - Establishes accountability
4. **Cover all layers** - Ensure metrics across the full stack
5. **Use layers for filtering** - Generate layer-specific reports
