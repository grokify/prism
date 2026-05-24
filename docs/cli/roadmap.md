# prism roadmap

Goals and requirements operations for planning and execution.

## Command Groups

### goals

Work with goal documents (OKR, V2MOM):

```bash
prism roadmap goals <type> <command> [flags]
```

### requirements

Work with requirements documents (PRD, MRD, TRD):

```bash
prism roadmap requirements <type> <command> [flags]
```

## Goals Commands

### okr render

Render OKRs to different formats:

```bash
prism roadmap goals okr render <file> [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--format`, `-f` | Output format: `markdown` (default), `marp`, `json` |
| `--output`, `-o` | Output file path |

**Examples:**

```bash
# Render to markdown
prism roadmap goals okr render q1-okrs.json -o okrs.md

# Generate Marp slides
prism roadmap goals okr render q1-okrs.json -f marp -o okrs-slides.md
```

### v2mom render

Render V2MOM to different formats:

```bash
prism roadmap goals v2mom render <file> [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--format`, `-f` | Output format: `markdown` (default), `marp`, `json` |
| `--output`, `-o` | Output file path |

## Requirements Commands

### prd render

Render a PRD to different formats:

```bash
prism roadmap requirements prd render <file> [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--format`, `-f` | Output format: `markdown` (default), `marp`, `pandoc` |
| `--output`, `-o` | Output file path |

**Examples:**

```bash
# Render to markdown with YAML frontmatter
prism roadmap requirements prd render feature.json -o feature.md

# Generate Marp presentation
prism roadmap requirements prd render feature.json -f marp -o feature-slides.md
```

### mrd render

Render an MRD (Market Requirements Document):

```bash
prism roadmap requirements mrd render <file> [flags]
```

### trd render

Render a TRD (Technical Requirements Document):

```bash
prism roadmap requirements trd render <file> [flags]
```

## Schema Commands

### schema generate

Generate JSON Schema for document types:

```bash
prism roadmap schema generate <type> [flags]
```

**Types:** `prd`, `mrd`, `trd`, `okr`, `v2mom`

### schema validate

Validate a document against its schema:

```bash
prism roadmap schema validate <file>
```

## Document Formats

### OKR Format

```json
{
  "$schema": "https://github.com/grokify/prism-roadmap/schema/okr.schema.json",
  "metadata": {
    "period": "Q1 2026",
    "team": "Platform"
  },
  "objectives": [
    {
      "id": "obj-reliability",
      "title": "Improve platform reliability",
      "keyResults": [
        {
          "id": "kr-availability",
          "description": "Achieve 99.9% availability",
          "target": 99.9,
          "current": 99.5,
          "unit": "%"
        }
      ]
    }
  ]
}
```

### PRD Format

```json
{
  "$schema": "https://github.com/grokify/prism-roadmap/schema/prd.schema.json",
  "metadata": {
    "title": "Feature Name",
    "author": "Product Manager",
    "status": "draft"
  },
  "problem": {
    "statement": "Description of the problem",
    "impact": "Business impact of the problem"
  },
  "solution": {
    "overview": "Proposed solution",
    "requirements": [
      {
        "id": "req-1",
        "description": "Requirement description",
        "priority": "P0"
      }
    ]
  }
}
```

## Workflow

The typical document flow:

1. **Goals** (OKR/V2MOM) define what the roadmap achieves
2. **Roadmap** sequences work items (RMIs)
3. **Requirements** (MRD/PRD/TRD) detail how to execute each RMI

```
OKR/V2MOM → Roadmap → RMI → MRD → PRD → TRD
```
