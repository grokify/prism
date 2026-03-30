# Security Metrics Example

This example demonstrates security-focused metrics covering the software delivery lifecycle.

## Overview

The security metrics example includes 5 metrics:

| Metric | Stage | Category | Type |
|--------|-------|----------|------|
| Vulnerability Scan Coverage | Build | Prevention | Coverage |
| Critical Vulnerability MTTR | Response | Response | Latency |
| SAST Coverage | Build | Prevention | Coverage |
| Threat Modeling Coverage | Design | Prevention | Coverage |
| Penetration Testing Coverage | Test | Detection | Coverage |

## File Location

```
examples/security-metrics.json
```

## Metric Details

### 1. Vulnerability Scan Coverage

Measures the percentage of repositories with vulnerability scanning enabled.

```json
{
  "id": "sec-vuln-coverage",
  "name": "Vulnerability Scan Coverage",
  "description": "Percentage of repositories with vulnerability scanning enabled",
  "domain": "security",
  "stage": "build",
  "category": "prevention",
  "metricType": "coverage",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 60,
  "current": 92,
  "target": 100,
  "thresholds": {
    "green": 95,
    "yellow": 80,
    "red": 60
  },
  "slo": {
    "target": ">=95%",
    "operator": "gte",
    "value": 95,
    "window": "30d"
  }
}
```

**Key Points:**

- Stage: Build (CI/CD integration)
- Category: Prevention (proactive control)
- Target: 100% coverage
- SLO: At least 95% coverage

### 2. Critical Vulnerability MTTR

Measures the mean time to remediate critical vulnerabilities.

```json
{
  "id": "sec-vuln-mttr",
  "name": "Critical Vulnerability MTTR",
  "description": "Mean time to remediate critical vulnerabilities",
  "domain": "security",
  "stage": "response",
  "category": "response",
  "metricType": "latency",
  "trendDirection": "lower_better",
  "unit": "days",
  "baseline": 30,
  "current": 7,
  "target": 3,
  "thresholds": {
    "green": 7,
    "yellow": 14,
    "red": 30
  },
  "slo": {
    "target": "<=7 days",
    "operator": "lte",
    "value": 7,
    "window": "30d"
  }
}
```

**Key Points:**

- Stage: Response (remediation)
- Trend: Lower is better
- Target: 3 days (aggressive)
- SLO: At most 7 days

### 3. SAST Coverage

Measures static application security testing coverage.

```json
{
  "id": "sec-sast-coverage",
  "name": "SAST Coverage",
  "description": "Percentage of code analyzed by static analysis",
  "domain": "security",
  "stage": "build",
  "category": "prevention",
  "metricType": "coverage",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 40,
  "current": 95,
  "target": 100,
  "thresholds": {
    "green": 95,
    "yellow": 80,
    "red": 60
  },
  "slo": {
    "target": ">=95%",
    "operator": "gte",
    "value": 95
  },
  "frameworkMappings": [
    {"framework": "NIST_CSF", "reference": "PR.DS-6"}
  ]
}
```

**Key Points:**

- Mapped to NIST CSF PR.DS-6 (integrity checking)
- High current coverage (95%)
- Part of CI/CD pipeline

### 4. Threat Modeling Coverage

Measures the percentage of features with threat models.

```json
{
  "id": "sec-threat-modeling",
  "name": "Threat Modeling Coverage",
  "description": "Percentage of new features with completed threat models",
  "domain": "security",
  "stage": "design",
  "category": "prevention",
  "metricType": "coverage",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 10,
  "current": 75,
  "target": 100,
  "thresholds": {
    "green": 90,
    "yellow": 70,
    "red": 50
  }
}
```

**Key Points:**

- Stage: Design (shift-left)
- Significant improvement from baseline
- No SLO yet (emerging practice)

### 5. Penetration Testing Coverage

Measures the percentage of applications with recent penetration tests.

```json
{
  "id": "sec-pentest-coverage",
  "name": "Penetration Testing Coverage",
  "description": "Percentage of applications with penetration tests in last 12 months",
  "domain": "security",
  "stage": "test",
  "category": "detection",
  "metricType": "coverage",
  "trendDirection": "higher_better",
  "unit": "%",
  "baseline": 30,
  "current": 80,
  "target": 100
}
```

**Key Points:**

- Stage: Test
- Category: Detection (finding vulnerabilities)
- Annual testing cadence

## Usage

### Validate

```bash
prism validate examples/security-metrics.json
```

### Score

```bash
prism score examples/security-metrics.json --detailed
```

### Expected Output

```
PRISM Score: 72.5% (Strong)

Security: 72.5%

By Stage:
  Design:   75.0%
  Build:    93.5%
  Test:     80.0%
  Response: 42.9%
```

## Extending the Example

Add additional security metrics:

- Dependency scanning coverage
- Container image scanning
- Runtime application self-protection (RASP)
- Security training completion
- Incident response time
