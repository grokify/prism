# Customer Awareness

PRISM includes a customer awareness model to track how well customers are informed about and responding to issues.

## Awareness States

PRISM defines four mutually exclusive awareness states:

| State | Constant | Weight | Description |
|-------|----------|--------|-------------|
| Unaware | `unaware` | 0.0 | Customer doesn't know about the issue |
| Aware (not acting) | `aware_not_remediating` | 0.25 | Customer knows but isn't taking action |
| Remediating | `aware_remediating` | 0.5 | Customer is actively working on remediation |
| Remediated | `aware_remediated` | 1.0 | Customer has resolved the issue |

## Why Track Awareness?

Customer awareness is crucial for B2B SaaS health because:

1. **Proactive Communication** - Shows you're communicating issues to customers
2. **Risk Management** - Unaware customers can't protect themselves
3. **Trust Building** - Transparency builds customer trust
4. **Compliance** - Many regulations require customer notification
5. **Support Planning** - Helps predict support ticket volume

## Awareness Score Calculation

The awareness score uses mutually exclusive state weights:

```
AwarenessScore = (unaware × 0.0) + (aware_not_acting × 0.25) +
                 (remediating × 0.5) + (remediated × 1.0)
```

Where each rate is a percentage (0.0-1.0) and all rates sum to 1.0.

### Example Calculation

| State | Customers | Percentage |
|-------|-----------|------------|
| Unaware | 100 | 10% |
| Aware (not acting) | 200 | 20% |
| Remediating | 300 | 30% |
| Remediated | 400 | 40% |
| **Total** | **1000** | **100%** |

```
AwarenessScore = (0.10 × 0.0) + (0.20 × 0.25) + (0.30 × 0.5) + (0.40 × 1.0)
               = 0 + 0.05 + 0.15 + 0.40
               = 0.60
```

## Impact on PRISM Score

The awareness score acts as a multiplier on the base score:

```
OverallScore = BaseScore × AwarenessScore
```

This means:

- Perfect awareness (all remediated): Score unchanged
- Poor awareness (all unaware): Score reduced to 0
- Mixed awareness: Score proportionally reduced

### Example Impact

| Base Score | Awareness | Overall |
|------------|-----------|---------|
| 0.80 | 1.0 | 0.80 |
| 0.80 | 0.60 | 0.48 |
| 0.80 | 0.25 | 0.20 |

## Data Structure

### CustomerAwarenessConfig

Enable awareness tracking for a metric:

```json
{
  "id": "sec-critical-vuln",
  "name": "Critical Vulnerabilities",
  "customerAwareness": {
    "enabled": true,
    "states": ["unaware", "aware_not_remediating", "aware_remediating", "aware_remediated"]
  }
}
```

### CustomerAwarenessData

Track the distribution across states:

```json
{
  "period": "2024-01",
  "distribution": [
    {"state": "unaware", "count": 100, "percent": 0.10},
    {"state": "aware_not_remediating", "count": 200, "percent": 0.20},
    {"state": "aware_remediating", "count": 300, "percent": 0.30},
    {"state": "aware_remediated", "count": 400, "percent": 0.40}
  ]
}
```

## Using Awareness in Go

```go
// Create awareness data
awareness := &prism.CustomerAwarenessData{
    Period: "2024-01",
    Distribution: []prism.AwarenessDistribution{
        {State: prism.AwarenessUnaware, Count: 100, Percent: 0.10},
        {State: prism.AwarenessAwareNotActing, Count: 200, Percent: 0.20},
        {State: prism.AwarenessAwareRemediating, Count: 300, Percent: 0.30},
        {State: prism.AwarenessAwareRemediated, Count: 400, Percent: 0.40},
    },
}

// Calculate scores
fmt.Printf("Unaware Rate: %.1f%%\n", awareness.UnawareRate()*100)
fmt.Printf("Awareness Score: %.2f\n", awareness.AwarenessScore())

// Use in PRISM score calculation
score := doc.CalculatePRISMScore(nil, awareness)
fmt.Printf("Overall: %.1f%% (with awareness)\n", score.Overall*100)
```

## Key Metrics

### Unaware Rate

Percentage of customers who don't know about the issue:

```go
rate := awareness.UnawareRate()
```

### Proactive Detection Rate

Percentage of customers who were proactively notified:

```go
rate := awareness.ProactiveDetectionRate() // 1 - unaware rate
```

### Proactive Resolution Rate

Percentage of customers who have remediated:

```go
rate := awareness.ProactiveResolutionRate()
```

## Best Practices

1. **Track All States** - Don't just track "aware" vs "unaware"
2. **Update Regularly** - Awareness changes over time
3. **Set Targets** - Aim for high remediation rates
4. **Automate Collection** - Integrate with CRM/support systems
5. **Segment by Severity** - Track awareness by issue severity

## State Transitions

```
Unaware → Aware (not acting) → Remediating → Remediated
  │                │                │
  └────────────────┴────────────────┴── (Customer may skip states)
```

Customers can skip states (e.g., go directly from Unaware to Remediating), but should never move backwards.
