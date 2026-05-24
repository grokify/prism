# Quick Start

This guide walks through the PRISM document flow from capability definition to execution planning.

## Document Flow Overview

```
Capability Stack  →  Maturity Model + State  →  OKR/V2MOM + Roadmap  →  MRD/PRD/TRD
```

## Step 1: Define Capabilities

Create a capability stack defining what your organization needs:

```bash
# Initialize a capability stack
prism capability init --domain security --output capabilities.json

# Validate the stack
prism capability validate capabilities.json

# List capabilities
prism capability list capabilities.json
```

## Step 2: Define Maturity Model

Create a maturity model with SLIs and level criteria:

```bash
# Generate a maturity dashboard
prism maturity model dashboard model.json --output dashboard.html

# Validate the model
prism maturity model validate model.json

# Lint for common issues
prism maturity model lint model.json
```

## Step 3: Track Current State

Create a state document with current measurements:

```bash
# Generate dashboard with state overlay
prism maturity model dashboard model.json \
  --state state.json \
  --output dashboard.html
```

## Step 4: Link Capabilities to SLIs

Add `PRISMRef` to capabilities to link them to maturity SLIs:

```json
{
  "id": "slo-framework",
  "name": "SLO Framework",
  "prismRef": {
    "sliIds": ["sli-slo-coverage", "sli-slo-attainment"]
  }
}
```

Then generate a dashboard with capability layer views:

```bash
prism maturity model dashboard model.json \
  --state state.json \
  --capstack capabilities.json \
  --aggregation min \
  --output dashboard.html
```

## Step 5: Define Goals

Create OKRs or V2MOMs to define success:

```bash
# Render OKRs to markdown
prism roadmap goals okr render okrs.json --output okrs.md

# Render V2MOM
prism roadmap goals v2mom render v2mom.json --output v2mom.md
```

## Step 6: Create Requirements

For each roadmap item, create detailed requirements:

```bash
# Render PRD to markdown
prism roadmap requirements prd render feature.json --output feature-prd.md

# Generate Marp slides
prism roadmap requirements prd render feature.json --format marp --output feature.md
```

## Step 7: Generate Static Site

Generate a multi-page static website from your PRISM data:

```bash
# Using standard PRISM directory structure
prism site generate --stack=./security/ --output=./dist

# Multiple stacks
prism site generate \
  --stack=./security/ \
  --stack=./reliability/ \
  --output=./dist \
  --title="PRISM Dashboard"
```

The generated site includes:

- Homepage with cards linking to all capability stacks
- Stack pages with filterable capability grids
- Capability detail pages with SLI metrics

## Standard PRISM Directory Structure

PRISM supports a standard directory structure that enables auto-discovery:

```
{stack-name}/
├── stack.json           # Capability stack definition
├── model.json           # Maturity model (SLIs, levels, criteria)
├── state.json           # Current state (SLI values)
└── roadmap.json         # OKRs, initiatives (optional)
```

When using this structure, files are auto-discovered:

```bash
# Single stack
prism site generate --stack=./security/

# Parent directory with multiple stacks
prism site generate --stack=./stacks/ --output=./dist
```

## Loading an Ecosystem

Once you have all documents, load them as an ecosystem:

```bash
# From config file
prism ecosystem load --config prism.yaml

# From directory structure
prism ecosystem load --dir ./ecosystem
```

Expected ecosystem directory structure:

```
ecosystem/
├── capability/
│   └── *.json
├── maturity/
│   └── *.json
└── roadmap/
    ├── okrs/*.json
    └── roadmaps/*.json
```

## Next Steps

- [CLI Reference](../cli/index.md) - Full command documentation
- [Configuration](../guide/configuration.md) - Ecosystem configuration options
- [Integration Guide](../guide/integration.md) - Cross-module workflows
