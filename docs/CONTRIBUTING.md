# Contributing to PRISM

Thank you for your interest in contributing to PRISM! This document covers how to contribute code, documentation, and bug reports.

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js 20+ and pnpm (for viewer/ TypeScript development)
- golangci-lint (for linting)

### Clone and Build

```bash
git clone https://github.com/grokify/prism.git
cd prism

# Build CLI
go build -o prism ./cmd/prism

# Install viewer dependencies
cd viewer && pnpm install
```

### Run Tests

```bash
# Go tests
go test ./...

# TypeScript tests
cd viewer && pnpm test
```

## Development Workflow

### 1. Find or Create an Issue

Before starting work, check if an issue exists. For new features, open an issue to discuss the approach.

For roadmap items tracked in [prism-control](https://github.com/ProductBuildersHQ/prism-control), use:
```bash
prismctl work ready --repo github.com/grokify/prism
```

### 2. Create a Branch

```bash
git checkout -b feat/your-feature-name
# or
git checkout -b fix/issue-description
```

### 3. Make Changes

Follow the code style and patterns documented in [DEVELOPMENT.md](DEVELOPMENT.md).

### 4. Test Your Changes

```bash
# Go
go test ./...
golangci-lint run

# TypeScript (if applicable)
cd viewer && pnpm test && pnpm typecheck
```

### 5. Commit with Conventional Commits

Use [Conventional Commits](https://www.conventionalcommits.org/) format:

```bash
git commit -m "feat(sitegen): add roadmap page generation"
git commit -m "fix(viewer): correct theme toggle behavior"
git commit -m "docs: update CLI reference for new flags"
```

For roadmap items, include the trailer:
```bash
git commit -m "feat(ecosystem): add V2MOM loading

Refs: RMI-PRISM-004"
```

### 6. Push and Create PR

```bash
git push origin feat/your-feature-name
```

Then create a pull request on GitHub.

## Code Style

### Go

- Run `gofmt` and `golangci-lint run` before committing
- Follow standard Go conventions
- Keep functions focused and small
- Prefer returning errors over panicking

### TypeScript

- Use TypeScript strict mode
- Define types with Zod schemas, infer TypeScript types from them
- Lit components follow the pattern in `viewer/src/components/`
- Run `pnpm typecheck` before committing

### Commit Messages

Format: `<type>(<scope>): <description>`

Types:
- `feat` — New feature
- `fix` — Bug fix
- `docs` — Documentation changes
- `test` — Test additions or fixes
- `refactor` — Code restructuring
- `chore` — Maintenance tasks

Scopes: `sitegen`, `viewer`, `ecosystem`, `cli`, `deps`

## What to Contribute

### Good First Issues

Look for issues labeled `good first issue` on GitHub.

### Documentation

- Improve existing docs in `docs/`
- Add examples to `examples/`
- Fix typos or unclear explanations

### Bug Fixes

- Check the issue tracker for reported bugs
- Include a test case that reproduces the bug
- Keep the fix minimal and focused

### New Features

For significant features:
1. Open an issue to discuss the approach
2. Reference the relevant spec in `docs/specs/` if applicable
3. Update documentation alongside the code
4. Add tests for new functionality

## Module Boundaries

This repo orchestrates the PRISM ecosystem. Schema changes belong in the module repos:

| Change Type | Repository |
|-------------|------------|
| Capability stack schema | prism-capability |
| SLI/SLO types, maturity models | prism-maturity |
| OKR, V2MOM, roadmap types | prism-roadmap |
| Shared primitives | prism-core |
| Cross-module queries, site generation | **this repo** |

If your feature requires a schema change:
1. Contribute the schema change to the module repo first
2. Wait for it to be tagged and released
3. Update the dependency in this repo
4. Build the feature here

## Pull Request Guidelines

- Keep PRs focused on a single change
- Include tests for new functionality
- Update documentation if behavior changes
- Ensure CI passes before requesting review
- Respond to review feedback promptly

## Reporting Issues

When reporting bugs:
- Include Go/Node version
- Provide minimal reproduction steps
- Include error messages and stack traces
- Attach sample input files if relevant

## Code of Conduct

Be respectful and constructive. We're all here to build something useful.

## Questions?

Open a discussion on GitHub or reach out to the maintainers.
