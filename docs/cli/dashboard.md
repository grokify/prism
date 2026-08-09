# prism dashboard

Generate a portfolio dashboard from the ecosystem: scored and ranked
initiatives, strategic pillar rollups, gap analysis, and a maturity heatmap.

## Persistent Flags

These flags apply to all `dashboard` subcommands:

| Flag | Description |
|------|-------------|
| `--dir` | Ecosystem directory to load |
| `--config` | Ecosystem config file (JSON) |
| `--theme` | Theme: `light` (default) or `dark` |
| `--gaps` | Include gap analysis (default `true`) |
| `--heatmap` | Include maturity heatmap (default `true`) |

## Commands

### generate

Generate the dashboard HTML into a site output directory:

```bash
prism dashboard generate --dir ./ecosystem --output ./dist
```

| Flag | Description |
|------|-------------|
| `--output` | Output directory for the site (default `./dist`) |

### export

Export a self-contained dashboard HTML file:

```bash
prism dashboard export --dir ./ecosystem --export dashboard.html --theme dark
```

| Flag | Description |
|------|-------------|
| `--export` | Output file path (default `./dashboard.html`) |

### data

Output the dashboard data as JSON (for custom rendering or tooling):

```bash
prism dashboard data --dir ./ecosystem
```

| Flag | Description |
|------|-------------|
| `--format` | Output format: `json` (default) |

## See Also

- [`prism prioritize`](prioritize.md) - The scoring model behind ranked initiatives
- [`prism gaps`](gaps.md) - Gap analysis included in the dashboard
- [`prism site`](site.md) - Full static site generation
