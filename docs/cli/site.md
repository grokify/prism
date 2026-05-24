# prism site

Generate static websites from PRISM data.

## Commands

### generate

Generate a multi-page static site from capability stacks, maturity models, and state documents.

```bash
prism site generate [flags]
```

**Generated Site Structure:**

- Homepage with cards linking to all capability stacks
- Stack pages with filterable capability grids
- Capability detail pages with SLI metrics

## Standard PRISM Directory Structure

PRISM supports a standard directory structure that enables auto-discovery of related files:

```
{stack-name}/
├── stack.json           # Capability stack definition
├── model.json           # Maturity model (SLIs, levels, criteria)
├── state.json           # Current state (SLI values)
└── roadmap.json         # OKRs, initiatives (optional)
```

When using the standard structure, all files are auto-discovered:

```bash
prism site generate --stack=./security/ --stack=./reliability/
```

Or point to a parent directory containing multiple stack directories:

```
stacks/
├── security/
│   ├── stack.json
│   ├── model.json
│   └── state.json
└── reliability/
    ├── stack.json
    ├── model.json
    └── state.json
```

```bash
prism site generate --stack=./stacks/ --output=./dist
```

## Legacy Mode

For non-standard layouts, use `--models` and `--states` directories:

```bash
prism site generate \
  --stack=./stacks/security.json \
  --models=./models/ \
  --states=./states/ \
  --output=./dist
```

Model/state files are matched by naming convention:

- `{basename}-model.json` or `{basename}/model.json`
- `{basename}-state.json` or `{basename}/state.json`

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--stack` | (required) | Capability stack file(s) or directory (can be repeated) |
| `--output`, `-o` | `./dist` | Output directory for generated site |
| `--title` | `PRISM Dashboard` | Site title |
| `--description` | `Capability maturity dashboard` | Site description |
| `--theme` | `dark` | Theme: `light` or `dark` |
| `--base-url` | | Base URL for the site (for absolute links) |
| `--models` | | Directory containing maturity model JSON files (legacy mode) |
| `--states` | | Directory containing state JSON files (legacy mode) |
| `--site-nav-js` | | Path to site-nav JavaScript bundle |

## Examples

### Single standard stack directory

```bash
prism site generate --stack=./security/ --output=./dist
```

### Multiple standard stack directories

```bash
prism site generate \
  --stack=./security/ \
  --stack=./reliability/ \
  --output=./dist \
  --title="PRISM Dashboard"
```

### Parent directory with multiple stacks

```bash
prism site generate \
  --stack=./stacks/ \
  --output=./dist/ \
  --theme=dark
```

### Legacy: separate files with models/states directories

```bash
prism site generate \
  --stack=./security-stack.json \
  --models=./models/ \
  --states=./states/ \
  --output=./dist
```

### With site-nav web components

```bash
prism site generate \
  --stack=./stacks/ \
  --output=./dist \
  --site-nav-js=./node_modules/@grokify/site-nav/dist/site-nav.es.js
```

## Output Structure

The generated site follows this structure:

```
dist/
├── index.html              # Homepage with stack cards
├── assets/
│   └── styles.css          # Theme-aware CSS
├── {stack-basename}/
│   ├── index.html          # Stack overview with capability grid
│   └── {capability-id}.html  # Capability detail pages
└── site-nav.es.js          # (optional) Web components bundle
```

## Theming

The generated site supports both light and dark themes. The default theme is `dark`.

When using the `site-nav` JavaScript bundle, users can toggle between themes using the theme toggle control in the navigation.

Without the JavaScript bundle, the site uses CSS-only theming with graceful fallbacks for badges and status indicators.

### CSS Custom Properties

The site uses CSS custom properties for theming. Override these in your own CSS:

```css
:root {
  --ds-bg-primary: #ffffff;
  --ds-text-primary: #1e293b;
  --ds-accent: #06b6d4;
}

[theme="dark"] {
  --ds-bg-primary: #0f172a;
  --ds-text-primary: #f1f5f9;
  --ds-accent: #22d3ee;
}
```
