# prism analyze

Analyze a PRISM document and generate initiative recommendations for achieving maturity targets.

## Synopsis

```bash
prism analyze <prism-file> [flags]
```

## Description

The `analyze` command examines a PRISM document to identify:

- Current maturity levels vs targets for each goal
- SLO compliance status at each maturity level
- Gaps that need to be addressed
- Phase-by-phase progress tracking

The output can be used directly for planning or fed to an LLM for detailed initiative recommendations.

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--format` | `-f` | Output format: `text`, `json`, or `prompt` (default: `text`) |
| `--recommend` | | Generate detailed recommendations (placeholder for LLM integration) |

## Output Formats

### Text (default)

Human-readable summary with goals, phases, and identified gaps:

```bash
prism analyze prism.json
```

Output:

```
PRISM Analysis
==============

Summary
-------
Goals: 2 | Phases: 2 | SLOs: 5/6 met (83%)
Average Maturity Gap: 1.5 levels

Goal Analysis
-------------

Achieve High Reliability (M3 → M5) [Behind]
  Gap: 2 levels | SLOs: 2/3 met
  Required SLOs:
    - [M4] ops-availability: >=99.9% [Met]
    - [M4] ops-mttr: <=1h [Not Met]
    - [M5] ops-p99-latency: <=100ms [Met]

Phase Analysis
--------------

Q1 2026 [in_progress] - 50% complete
  Initiatives: 2
  Goal Targets:
    - Achieve High Reliability: M2 → M3 (2 SLOs needed)

Identified Gaps
---------------
  [HIGH] maturity: Goal 'Achieve High Reliability' has 2-level gap to target
  [MEDIUM] slo: SLO 'ops-mttr' not met for maturity level M4
```

### JSON

Structured output for programmatic use:

```bash
prism analyze prism.json -f json
```

Output:

```json
{
  "summary": {
    "totalGoals": 2,
    "totalPhases": 2,
    "totalSLOs": 6,
    "slosMet": 5,
    "sloCompliance": 83.33,
    "avgMaturityGap": 1.5
  },
  "goals": [...],
  "phases": [...],
  "gaps": [...],
  "recommendations": [...]
}
```

### Prompt

Generates a structured prompt for LLM-based initiative planning:

```bash
prism analyze prism.json -f prompt
```

Output:

```markdown
# PRISM Analysis Prompt

You are an operational planning assistant. Analyze the following PRISM document
and recommend initiatives to achieve the maturity targets.

## Current State

### Goals

- **Achieve High Reliability**: Currently at M3, targeting M5 (2 level gap)
  - Required SLOs:
    - [M4] ops-availability: target >=99.9%, current 99.95 (MET)
    - [M4] ops-mttr: target <=1h, current 2.00 (NOT MET)

### Phases

- **Q1 2026** [in_progress]
  - Achieve High Reliability: M2 → M3
  - Current initiatives: 2

## Request

Based on the above analysis, please recommend initiatives that will:

1. Enable achievement of SLOs required for each maturity level progression
2. Be appropriately sequenced across phases (dependencies considered)
3. Address the identified gaps

For each initiative, provide:

- **Title**: Clear, actionable name
- **Description**: What will be delivered
- **Phase**: Which phase to execute in
- **Goals**: Which goals this supports
- **SLOs Enabled**: Which SLOs this helps achieve
- **Priority**: High/Medium/Low
- **Dependencies**: Other initiatives this depends on

Output as JSON array of recommendations.
```

## Examples

### Basic Analysis

```bash
prism analyze prism.json
```

### JSON Output for Automation

```bash
prism analyze prism.json -f json | jq '.gaps[] | select(.severity == "high")'
```

### Generate LLM Prompt

```bash
prism analyze prism.json -f prompt > analysis-prompt.md
```

## Use with LLM Workflows

The prompt format is designed for use with LLMs to generate initiative recommendations:

```bash
# Generate prompt and pipe to LLM CLI
prism analyze prism.json -f prompt | llm -m gpt-4

# Or save and use with API
prism analyze prism.json -f prompt > prompt.md
curl -X POST https://api.openai.com/v1/chat/completions \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -d @- <<EOF
{
  "model": "gpt-4",
  "messages": [{"role": "user", "content": "$(cat prompt.md)"}]
}
EOF
```

## Related Commands

- [`prism export`](export.md) - Export to OKR/V2MOM formats
- [`prism goal progress`](goal.md) - View goal progress
- [`prism roadmap progress`](roadmap.md) - View roadmap progress
