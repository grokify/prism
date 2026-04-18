# PRISM

**Proactive Reliability & Security Maturity Model**

PRISM is a unified framework for B2B SaaS health metrics that combines SLOs, DMAIC, OKRs, and maturity modeling into a single coherent system. It provides structured schemas for defining metrics, calculating composite health scores, and tracking organizational maturity across security and operations domains.

## Key Features

- **Unified Metrics Framework** - Combine security and operations metrics in a single document
- **5-Level Maturity Model** - Track organizational capability from Reactive to Optimizing
- **Composite Scoring** - Calculate weighted PRISM scores across domains and lifecycle stages
- **Customer Awareness Tracking** - Model customer awareness states for proactive communication
- **Framework Mappings** - Map metrics to DORA, SRE, NIST CSF, and MITRE ATT&CK frameworks
- **Machine-Evaluable SLOs** - Define SLOs with operators for programmatic checking
- **CLI Tool** - Initialize, validate, and score PRISM documents from the command line
- **JSON Schema** - Auto-generated schema for editor validation and IDE support

## Enterprise Features

- **Dashforge Integration** - Embed PRISM dashboards in dashforge sites or standalone pages
- **Marp Presentations** - Generate executive presentations from PRISM documents
- **Excel Export** - Export metrics and scores to XLSX for stakeholder reporting

## Quick Example

```json
{
  "$schema": "https://github.com/grokify/prism/schema/prism.schema.json",
  "metadata": {
    "name": "Acme Corp PRISM",
    "version": "1.0.0"
  },
  "metrics": [
    {
      "id": "ops-availability",
      "name": "Service Availability",
      "domain": "operations",
      "stage": "runtime",
      "category": "reliability",
      "metricType": "rate",
      "current": 99.95,
      "target": 99.99,
      "slo": {
        "target": ">=99.99%",
        "operator": "gte",
        "value": 99.99,
        "window": "30d"
      }
    }
  ]
}
```


## PRISM Score

The PRISM score provides a single composite metric (0.0-1.0) representing organizational health:

| Score | Level | Description |
|-------|-------|-------------|
| ≥0.90 | Elite | Industry-leading practices |
| ≥0.75 | Strong | Well-managed, proactive |
| ≥0.50 | Medium | Adequate, room for improvement |
| ≥0.25 | Weak | Significant gaps |
| <0.25 | Critical | Immediate attention required |

## Getting Started

```bash
# Install the CLI
go install github.com/grokify/prism/cmd/prism@latest

# Initialize a new document
prism init -o prism.json

# Validate the document
prism validate prism.json

# Calculate the PRISM score
prism score prism.json --detailed
```

## License

MIT
