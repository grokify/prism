# Dashforge Integration

!!! note "Coming Soon"
    Dashforge integration is planned for a future release.

## Overview

Dashforge integration enables PRISM dashboards as:

- **Standalone pages** - Full-page PRISM dashboards
- **Embedded widgets** - PRISM components in dashforge sites
- **Trend visualization** - Historical score tracking

## Planned Features

### Standalone Dashboard

A full-page dashboard showing:

- Overall PRISM score with interpretation
- Domain scores (Security, Operations)
- Stage breakdown heatmap
- Metric status table
- Score trends over time

### Dashboard Widgets

Embeddable components for dashforge sites:

| Widget | Description |
|--------|-------------|
| Score Gauge | PRISM score with color indicator |
| Domain Cards | Security/Operations score cards |
| Heatmap | Domain × Stage matrix |
| Trend Chart | Score history line chart |
| Metric Table | Filterable metric list |

### Heatmap Visualization

A domain × stage heatmap showing cell scores:

```
           │ Design │ Build │ Test │ Runtime │ Response │
───────────┼────────┼───────┼──────┼─────────┼──────────┤
Security   │  🟡 65 │ 🟢 92 │ 🟡 70│  🟢 88  │   🟡 72  │
───────────┼────────┼───────┼──────┼─────────┼──────────┤
Operations │  🟡 68 │ 🟢 85 │ 🟡 75│  🟢 95  │   🟢 82  │
```

## Configuration

### Document Configuration

```json
{
  "integrations": {
    "dashforge": {
      "enabled": true,
      "theme": "default",
      "refreshInterval": "1h",
      "widgets": ["score", "heatmap", "trends"]
    }
  }
}
```

### Dashforge Site Configuration

```yaml
# dashforge.yml
pages:
  - name: PRISM Dashboard
    type: prism
    source: prism.json
    layout: full
    widgets:
      - type: prism-score
        position: top
      - type: prism-heatmap
        position: center
      - type: prism-trends
        position: bottom
```

## Planned CLI Commands

```bash
# Generate standalone dashboard
prism dashboard prism.json -o dashboard.html

# Generate dashforge widget data
prism dashboard prism.json --format dashforge -o prism-widget.json

# Start live dashboard server
prism serve prism.json --port 8080
```

## Integration with Dashforge

### Embedding in Dashforge Site

```markdown
<!-- In a dashforge page -->
# Platform Health

{{< prism-score source="prism.json" >}}

## Domain Breakdown

{{< prism-heatmap source="prism.json" >}}
```

### API Access

```go
import (
    "github.com/grokify/prism"
    "github.com/grokify/dashforge"
)

// Load PRISM document
doc, _ := prism.LoadDocument("prism.json")

// Generate dashforge widget
widget := dashforge.NewPRISMWidget(doc)
widget.Render(w)
```

## Roadmap

1. **Phase 1**: Static dashboard generation
2. **Phase 2**: Dashforge widget integration
3. **Phase 3**: Live data refresh
4. **Phase 4**: Historical trend storage
