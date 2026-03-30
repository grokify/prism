# Excel Export

!!! note "Coming Soon"
    Excel export is planned for a future release.

## Overview

Export PRISM documents to XLSX format for stakeholder reporting and offline analysis.

## Planned Features

### Worksheets

| Worksheet | Description |
|-----------|-------------|
| Summary | Overall scores and health status |
| Metrics | All metrics with current values |
| Security | Security domain metrics |
| Operations | Operations domain metrics |
| Maturity | Maturity model assessment |
| Trends | Historical data (if available) |
| Raw Data | Full JSON data for reference |

### Summary Worksheet

| Field | Value |
|-------|-------|
| Document Name | Acme PRISM |
| Report Date | 2024-01-15 |
| Overall Score | 78.2% |
| Interpretation | Strong |
| Security Score | 72.5% |
| Operations Score | 84.0% |
| Maturity Average | 3.6 |
| Performance Average | 82.3% |

### Metrics Worksheet

| ID | Name | Domain | Stage | Current | Target | Status | SLO Met |
|----|------|--------|-------|---------|--------|--------|---------|
| ops-availability | Service Availability | operations | runtime | 99.95% | 99.99% | Yellow | No |
| sec-sast | SAST Coverage | security | build | 95% | 100% | Green | Yes |

### Conditional Formatting

- **Green** cells for status = Green
- **Yellow** cells for status = Yellow
- **Red** cells for status = Red
- **Bold** for metrics not meeting SLO

## Configuration

### Document Configuration

```json
{
  "integrations": {
    "excel": {
      "enabled": true,
      "includeCharts": true,
      "worksheets": ["summary", "metrics", "maturity"],
      "conditionalFormatting": true
    }
  }
}
```

### Export Options

| Option | Values | Description |
|--------|--------|-------------|
| `worksheets` | array | Worksheets to include |
| `includeCharts` | true/false | Embed charts |
| `conditionalFormatting` | true/false | Color-code cells |
| `includeRawData` | true/false | Include JSON worksheet |

## Planned CLI Commands

```bash
# Export to Excel
prism export prism.json -o report.xlsx

# Export specific worksheets
prism export prism.json --worksheets summary,metrics -o report.xlsx

# Export with charts
prism export prism.json --include-charts -o report.xlsx

# Export filtered by domain
prism export prism.json --domain security -o security-report.xlsx
```

## Chart Types

Charts embedded in Excel:

| Chart | Worksheet | Description |
|-------|-----------|-------------|
| Score Gauge | Summary | Circular gauge |
| Domain Bar | Summary | Security vs Operations |
| Heatmap | Summary | Domain × Stage matrix |
| Metric Status | Metrics | Green/Yellow/Red distribution |
| Trend Line | Trends | Historical score |

## Use Cases

### Stakeholder Report

```bash
prism export prism.json \
  --worksheets summary,metrics \
  --include-charts \
  -o stakeholder-report.xlsx
```

### Compliance Audit

```bash
prism export prism.json \
  --worksheets metrics \
  --filter-framework NIST_CSF \
  -o nist-compliance.xlsx
```

### Data Analysis

```bash
prism export prism.json \
  --worksheets raw \
  --format csv \
  -o metrics.csv
```

## Library Usage

```go
import (
    "github.com/grokify/prism"
    "github.com/grokify/prism/export"
)

doc, _ := prism.LoadDocument("prism.json")

exporter := export.NewExcelExporter(doc)
exporter.IncludeCharts = true
exporter.Worksheets = []string{"summary", "metrics"}

err := exporter.Export("report.xlsx")
```

## Roadmap

1. **Phase 1**: Basic XLSX generation with metrics
2. **Phase 2**: Conditional formatting
3. **Phase 3**: Embedded charts
4. **Phase 4**: Trend data support
