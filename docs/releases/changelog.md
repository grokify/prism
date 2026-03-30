# Changelog

All notable changes to PRISM will be documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html),
and commits follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

## [Unreleased]

## [v0.1.0] - 2026-03-30

### Highlights

- PRISM (Proactive Reliability & Security Maturity Model) - a unified framework for B2B SaaS health metrics combining SLOs, DMAIC, OKRs, and maturity modeling
- CLI tool for creating, validating, and scoring PRISM documents
- Auto-generated JSON Schema from Go types for editor validation

### Added

- Core types: PRISMDocument, Metric, SLI, SLO, Thresholds, DataPoint
- Domain constants: security, operations
- Lifecycle stage constants: design, build, test, runtime, response
- Category constants: prevention, detection, response, reliability, efficiency, quality
- Metric types: coverage, rate, latency, ratio, count, distribution, score
- Framework mapping support: NIST_CSF, MITRE_ATTACK, DORA, SRE
- Validation functions for all constants and document structure
- ValidationError and ValidationErrors types for detailed error reporting
- 5-level maturity model: Reactive, Basic, Defined, Managed, Optimizing
- MaturityModel, MaturityLevelDef, MaturityCell types
- NewMaturityModelForDomains() for domain-filtered initialization
- Customer awareness model with four states: unaware, aware_not_remediating, aware_remediating, aware_remediated
- AwarenessScore() using mutually exclusive state weights (0.0 → 0.25 → 0.5 → 1.0)
- PRISM score calculation combining maturity and performance
- Configurable weights: maturity 40%, performance 60%, stage weights, domain weights
- Score interpretation: Elite (≥0.9), Strong (≥0.75), Medium (≥0.5), Weak (≥0.25), Critical
- SLO.Operator and SLO.Value fields for machine-evaluable targets
- Metric.MeetsSLO() method for programmatic SLO checking
- Metric.CalculateStatus() for threshold-based status calculation
- Metric.ProgressToTarget() for progress tracking
- `prism init` command with domain filtering (-d) and output path (-o)
- `prism validate` command for document validation
- `prism score` command with --detailed and --json flags
- `prism catalog` command to list available constants
- JSON Schema auto-generation from Go types using invopop/jsonschema
- Embedded schema access via schema.PRISMSchemaJSON()

### Dependencies

- github.com/spf13/cobra v1.10.2 for CLI
- github.com/invopop/jsonschema v0.13.0 for schema generation

### Documentation

- README with installation, CLI usage, schema overview, and examples
- Security metrics example (5 metrics: vuln coverage, MTTR, SAST, threat modeling, pentest)
- Operations metrics example (8 DORA-aligned metrics: availability, latency, error rate, deployment frequency, lead time, MTTR, change failure rate, IaC coverage)
- Score weight normalization behavior documented in ScoreConfig

### Tests

- Unit tests for validation, maturity, awareness, and scoring
- Integration tests for example JSON files
- DataPoint timestamp JSON round-trip tests
- MeetsSLO() tests for all operators (gte, lte, gt, lt, eq)

### Infrastructure

- GitHub Actions workflows for CI, linting, and CodeQL
- Dependabot configuration for automated dependency updates
- golangci-lint configuration

[unreleased]: https://github.com/grokify/prism/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/grokify/prism/releases/tag/v0.1.0
