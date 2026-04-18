# Quality Metrics Example

This example demonstrates the quality domain with ISO 25010 quality verticals.

## Document Overview

The example includes:

- Quality domain metrics
- ISO 25010 quality vertical classifications
- Testing effectiveness and code quality metrics
- Maturity mappings for test coverage

## Key Concepts

### Quality Domain

The quality domain focuses on code quality, testing, and defect management:

```json
{
  "domains": [
    {
      "name": "quality",
      "description": "Code quality, testing, and defect management metrics"
    }
  ]
}
```

### ISO 25010 Quality Verticals

Metrics are classified by quality characteristic:

| Vertical | Example Metrics |
|----------|-----------------|
| Functional | Test coverage, mutation score, test effectiveness |
| Reliability | Defect density, regression rate, bug fix time |
| Maintainability | Tech debt ratio, complexity, code duplication |
| Usability | Accessibility compliance |
| Performance | API response time budget |

### Quality Vertical Assignment

```json
{
  "id": "qa-unit-coverage",
  "name": "Unit Test Coverage",
  "domain": "quality",
  "qualityVertical": "functional"
}
```

## Metrics Summary

| Metric | Vertical | Type | Trend |
|--------|----------|------|-------|
| Unit Test Coverage | functional | coverage | higher_better |
| Mutation Testing Score | functional | score | higher_better |
| Test Effectiveness (DDP) | functional | rate | higher_better |
| Defect Density | reliability | ratio | lower_better |
| Technical Debt Ratio | maintainability | ratio | lower_better |
| Cyclomatic Complexity | maintainability | score | lower_better |
| Bug Fix Lead Time | reliability | latency | lower_better |
| Accessibility Compliance | usability | rate | higher_better |

## Download

[quality-metrics.json](https://github.com/grokify/prism/blob/main/examples/quality-metrics.json)

## See Also

- [Quality Domain](../schema/quality.md) - Quality domain documentation
- [prism catalog](../cli/catalog.md) - List quality verticals
