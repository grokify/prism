# prism catalog

List all available PRISM constants.

## Synopsis

```bash
prism catalog
```

## Description

Displays all valid constants that can be used in PRISM documents, including:

- Domains
- Lifecycle stages
- Categories
- Metric types
- Trend directions
- SLO operators
- Status values
- Maturity levels
- Awareness states
- Framework identifiers

## Examples

### List All Constants

```bash
prism catalog
```

Output:

```
PRISM Catalog
=============

Domains:
  security     - Application and infrastructure security
  operations   - Reliability, performance, and efficiency

Stages:
  design       - Architecture, requirements, planning
  build        - CI/CD, code analysis, dependency management
  test         - Testing coverage, integration testing
  runtime      - Production monitoring, availability
  response     - Incident response, remediation

Categories:
  prevention   - Proactive controls
  detection    - Monitoring and alerting
  response     - Incident handling
  reliability  - Availability and durability
  efficiency   - Performance and utilization
  quality      - Code and process quality

Metric Types:
  coverage     - Percentage coverage (e.g., test coverage)
  rate         - Frequency or percentage (e.g., error rate)
  latency      - Time duration (e.g., P99 latency)
  ratio        - Proportion (e.g., success ratio)
  count        - Absolute count (e.g., incidents)
  distribution - Statistical distribution (e.g., percentiles)
  score        - Composite score (e.g., health score)

Trend Directions:
  higher_better - Higher values are better
  lower_better  - Lower values are better
  target_value  - Target a specific value

SLO Operators:
  gte - Greater than or equal (>=)
  lte - Less than or equal (<=)
  gt  - Greater than (>)
  lt  - Less than (<)
  eq  - Equal (=)

Maturity Levels:
  1 - Reactive   (Ad-hoc processes)
  2 - Basic      (Basic controls)
  3 - Defined    (Standardized processes)
  4 - Managed    (Measured and controlled)
  5 - Optimizing (Continuous improvement)

Awareness States:
  unaware              - Customer not aware of issue
  aware_not_remediating - Customer aware but not acting
  aware_remediating    - Customer actively remediating
  aware_remediated     - Customer has remediated

Frameworks:
  DORA         - DevOps Research and Assessment
  SRE          - Site Reliability Engineering
  NIST_CSF     - NIST Cybersecurity Framework (see prism-security)
  MITRE_ATTACK - MITRE ATT&CK Framework (see prism-security)
```

## Use Cases

- Reference valid values when writing PRISM documents
- Understand the taxonomy of metrics and categories
- Identify framework mapping options
