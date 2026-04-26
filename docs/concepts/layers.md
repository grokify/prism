# Layers

PRISM uses a value stream model to organize metrics across the full product lifecycle. This clarifies ownership from ideation through customer support.

## The Value Stream

```
Requirements → Code → Infra → Runtime → Adoption → Support
```

| Layer | Description | Typical Owner |
|-------|-------------|---------------|
| `requirements` | Product ideation, specs, design | Product/Design |
| `code` | Application code, libraries, dependencies | Dev teams |
| `infra` | Cloud resources, networking, platform | Platform team |
| `runtime` | Running services, production workloads | Stream-aligned + SRE |
| `adoption` | Product analytics, user engagement | Product/Growth |
| `support` | Customer support, incident management | Support/CS |

## Layer Definitions

### Requirements Layer

The requirements layer covers product ideation and specification:

- Product requirements documents (PRDs)
- Design specifications
- User research and validation
- Feature prioritization

**Example metrics:**

- Requirements clarity score
- Spec completion rate
- Design review coverage
- User research coverage

### Code Layer

The code layer encompasses application development:

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

The infrastructure layer covers platform and resources:

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

The runtime layer focuses on production systems:

- Service availability and performance
- Runtime security monitoring
- Observability completeness
- Incident response

**Example metrics:**

- Service availability (SLO)
- P99 latency
- Error rate
- Mean time to recovery

### Adoption Layer

The adoption layer tracks product usage and engagement:

- Feature adoption rates
- User activation and retention
- Self-service success rates
- Product analytics (Pendo, Amplitude, etc.)

**Example metrics:**

- Feature adoption rate
- User activation percentage
- Self-service completion rate
- DAU/MAU ratio
- Time to value

### Support Layer

The support layer covers customer assistance:

- Support ticket volume and resolution
- Escalation patterns
- Customer satisfaction
- Knowledge base effectiveness

**Example metrics:**

- Ticket resolution time
- First contact resolution rate
- Customer satisfaction (CSAT)
- Escalation rate
- Knowledge base deflection rate

## Golden Signals

Each layer can define golden signals based on Google SRE's four golden signals:

| Signal | Description | Example |
|--------|-------------|---------|
| Latency | Time to complete | Response time, resolution time |
| Traffic | Throughput | Requests/sec, tickets/day |
| Errors | Error rate | 5xx errors, failed tickets |
| Saturation | Resource utilization | CPU usage, agent capacity |

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
    },
    {
      "id": "support",
      "name": "Support",
      "signals": {
        "latency": "metric-resolution-time",
        "traffic": "metric-tickets-per-day",
        "errors": "metric-escalation-rate",
        "saturation": "metric-agent-utilization"
      }
    }
  ]
}
```

## Layers vs Stages

Layers and stages are orthogonal dimensions:

| Dimension | Purpose | Values |
|-----------|---------|--------|
| **Layer** | Where in value stream | requirements, code, infra, runtime, adoption, support |
| **Stage** | When in delivery cycle | design, build, test, runtime, response |

A metric belongs to one layer AND one stage. For example:

- Support layer + Design stage = Designing support processes
- Support layer + Runtime stage = Handling live tickets
- Support layer + Response stage = Escalation management

## Layer Accountability

Teams can declare layer accountability:

```json
{
  "teams": [
    {
      "id": "product-team",
      "name": "Product Team",
      "type": "stream_aligned",
      "layerAccountability": ["requirements", "adoption"]
    },
    {
      "id": "platform-team",
      "name": "Platform Team",
      "type": "platform",
      "layerAccountability": ["infra"]
    },
    {
      "id": "app-team",
      "name": "Application Team",
      "type": "stream_aligned",
      "layerAccountability": ["code", "runtime"]
    },
    {
      "id": "support-team",
      "name": "Support Team",
      "type": "stream_aligned",
      "layerAccountability": ["support"]
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
prism layer show prism.json adoption
```

## Best Practices

1. **Cover the full value stream** - Ensure metrics across all layers
2. **Assign every metric to a layer** - Clarifies ownership
3. **Define golden signals** - Standardizes observability per layer
4. **Map teams to layers** - Establishes accountability
5. **Track handoffs** - Measure transitions between layers
6. **Use layers for reporting** - Generate layer-specific dashboards
