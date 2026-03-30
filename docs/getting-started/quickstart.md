# Quick Start

This guide walks you through creating your first PRISM document and calculating a health score.

## 1. Initialize a Document

Create a new PRISM document with default settings:

```bash
prism init -o prism.json
```

This creates a document with both security and operations domains.

### Domain-Specific Documents

Create a security-focused document:

```bash
prism init -d security -o security-metrics.json
```

Create an operations-focused document:

```bash
prism init -d operations -o ops-metrics.json
```

## 2. Edit Your Metrics

Open `prism.json` and customize the metrics. Each metric requires:

- **id** - Unique identifier
- **name** - Human-readable name
- **domain** - `security` or `operations`
- **stage** - `design`, `build`, `test`, `runtime`, or `response`
- **category** - `prevention`, `detection`, `response`, `reliability`, `efficiency`, or `quality`
- **metricType** - `coverage`, `rate`, `latency`, `ratio`, `count`, `distribution`, or `score`
- **current** - Current value
- **target** - Target value

Example metric:

```json
{
  "id": "ops-availability",
  "name": "Service Availability",
  "domain": "operations",
  "stage": "runtime",
  "category": "reliability",
  "metricType": "rate",
  "unit": "%",
  "current": 99.95,
  "target": 99.99,
  "slo": {
    "target": ">=99.99%",
    "operator": "gte",
    "value": 99.99,
    "window": "30d"
  }
}
```

## 3. Validate Your Document

Check that your document is valid:

```bash
prism validate prism.json
```

If there are errors, the CLI will display them with field paths.

## 4. Calculate Your Score

Get your PRISM score:

```bash
prism score prism.json
```

For a detailed breakdown:

```bash
prism score prism.json --detailed
```

For JSON output (useful for automation):

```bash
prism score prism.json --json
```

## 5. View Available Constants

List all valid domains, stages, categories, and metric types:

```bash
prism catalog
```

## Next Steps

- [Create Your First Document](first-document.md) - Detailed walkthrough
- [Schema Reference](../schema/index.md) - Complete schema documentation
- [PRISM Score](../concepts/scoring.md) - Understanding the scoring formula
