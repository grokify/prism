# Marp Presentations

!!! note "Coming Soon"
    Marp integration is planned for a future release.

## Overview

Generate executive presentations from PRISM documents using [Marp](https://marp.app/).

## Planned Features

### Slide Templates

| Template | Description | Slides |
|----------|-------------|--------|
| Executive Summary | High-level overview | 3-5 |
| Full Report | Detailed breakdown | 10-15 |
| Domain Focus | Single domain deep-dive | 5-7 |
| Trend Report | Historical comparison | 5-8 |

### Generated Slides

**Executive Summary Template:**

1. **Title Slide** - PRISM report title, date, period
2. **Score Overview** - Overall score, interpretation, health status
3. **Domain Summary** - Security vs Operations scores
4. **Key Metrics** - Top performing and at-risk metrics
5. **Recommendations** - Action items based on scores

### Example Output

```markdown
---
marp: true
theme: corporate
---

# PRISM Report
## Acme Corporation
### Q1 2024

---

# Overall Score

## 78.2% - Strong

- Security: 72.5%
- Operations: 84.0%

![bg right:40%](score-gauge.svg)

---

# Domain × Stage Heatmap

|            | Design | Build | Test | Runtime | Response |
|------------|--------|-------|------|---------|----------|
| Security   | 🟡 65  | 🟢 92 | 🟡 70| 🟢 88   | 🟡 72    |
| Operations | 🟡 68  | 🟢 85 | 🟡 75| 🟢 95   | 🟢 82    |

---

# Top Performing Metrics

1. **Service Availability** - 99.95% (Target: 99.99%)
2. **SAST Coverage** - 95% (Target: 100%)
3. **Deployment Frequency** - 5/day (Target: 10/day)

---

# Areas Needing Attention

1. **Threat Modeling** - 75% (Target: 100%)
2. **Lead Time** - 24 hours (Target: 1 hour)
3. **Error Rate** - 0.15% (Target: 0.1%)
```

## Configuration

### Document Configuration

```json
{
  "integrations": {
    "marp": {
      "enabled": true,
      "template": "executive",
      "theme": "corporate",
      "includeCharts": true
    }
  }
}
```

### Template Options

| Option | Values | Description |
|--------|--------|-------------|
| `template` | executive, full, domain, trend | Slide template |
| `theme` | default, corporate, dark | Marp theme |
| `includeCharts` | true/false | Include generated charts |
| `format` | pdf, pptx, html | Output format |

## Planned CLI Commands

```bash
# Generate Marp markdown
prism slides prism.json -o slides.md

# Generate PDF directly
prism slides prism.json --format pdf -o report.pdf

# Generate with custom template
prism slides prism.json --template executive -o slides.md

# Generate PowerPoint
prism slides prism.json --format pptx -o report.pptx
```

## Chart Generation

Charts will be generated as SVG for embedding:

- **Score Gauge** - Circular gauge with score
- **Trend Line** - Historical score trend
- **Heatmap** - Domain × Stage matrix
- **Bar Charts** - Metric comparisons

## Use Cases

### Monthly Executive Report

```bash
prism slides prism.json \
  --template executive \
  --title "Monthly PRISM Report" \
  --period "January 2024" \
  --format pdf \
  -o monthly-report.pdf
```

### Quarterly Review

```bash
prism slides prism.json \
  --template trend \
  --compare prism-q3.json prism-q4.json \
  --format pptx \
  -o quarterly-review.pptx
```

## Roadmap

1. **Phase 1**: Basic Marp markdown generation
2. **Phase 2**: PDF/PPTX export
3. **Phase 3**: Chart embedding
4. **Phase 4**: Trend comparison slides
