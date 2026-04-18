# prism slo-report

Generate SLO compliance reports from a PRISM document.

## Synopsis

```bash
prism slo-report <file> [options]
```

## Description

Generate a report showing SLO compliance status for all metrics with defined SLOs. The report includes compliance rate, metric status breakdown, and detailed compliance information.

## Options

| Option | Description |
|--------|-------------|
| `-o`, `--output <file>` | Output file (default: stdout) |
| `--json` | Output in JSON format |
| `--goal <id>` | Filter to metrics for a specific goal |
| `--phase <id>` | Filter to metrics relevant to a specific phase |

## Examples

Generate SLO report to stdout:

```bash
prism slo-report prism.json
```

Generate SLO report to file:

```bash
prism slo-report prism.json -o slo-report.md
```

Generate JSON report:

```bash
prism slo-report prism.json --json
```

Filter to a specific goal:

```bash
prism slo-report prism.json --goal goal-reliability
```

## Output Format

### Text Output

```
SLO Compliance Report
Generated: 2026-04-18

Overall Compliance: 6/8 (75%)

By Status:
  Meeting SLO:    6 metrics
  Below Target:   2 metrics
  No SLO Defined: 0 metrics

Detailed Status:
  [PASS] ops-availability: 99.95% (target: >=99.9%)
  [PASS] ops-p99-latency: 150ms (target: <=200ms)
  [FAIL] ops-mttr: 2h (target: <=1h)
  [PASS] ops-deploy-frequency: 5/day (target: >=5/day)
  ...
```

### JSON Output

```json
{
  "generatedAt": "2026-04-18T10:00:00Z",
  "complianceRate": 0.75,
  "summary": {
    "total": 8,
    "meeting": 6,
    "below": 2,
    "noSLO": 0
  },
  "metrics": [
    {
      "id": "ops-availability",
      "name": "Service Availability",
      "current": 99.95,
      "target": 99.9,
      "operator": "gte",
      "status": "meeting"
    }
  ]
}
```

## See Also

- [prism report](report.md) - Roadmap reports
- [prism score](score.md) - PRISM score calculation
- [SLIs & SLOs](../schema/slos.md) - SLO schema reference
