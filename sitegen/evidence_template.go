package sitegen

const evidenceTemplate = `<!DOCTYPE html>
<html lang="en" data-theme="{{.Theme}}">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.CapabilityName}} - Evidence | {{.Title}}</title>
    <style>
        :root {
            --bg-primary: #ffffff;
            --bg-secondary: #f8f9fa;
            --bg-card: #ffffff;
            --text-primary: #212529;
            --text-secondary: #6c757d;
            --border-color: #dee2e6;
            --accent-csat: #28a745;
            --accent-signal: #007bff;
            --accent-customer: #6f42c1;
            --accent-market: #fd7e14;
            --accent-idea: #20c997;
        }

        [data-theme="dark"] {
            --bg-primary: #1a1a2e;
            --bg-secondary: #16213e;
            --bg-card: #1f2937;
            --text-primary: #f8f9fa;
            --text-secondary: #adb5bd;
            --border-color: #495057;
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: var(--bg-primary);
            color: var(--text-primary);
            line-height: 1.6;
            padding: 2rem;
        }

        .container {
            max-width: 1200px;
            margin: 0 auto;
        }

        h1 {
            font-size: 1.75rem;
            margin-bottom: 0.5rem;
        }

        .subtitle {
            color: var(--text-secondary);
            margin-bottom: 2rem;
        }

        .summary-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 1rem;
            margin-bottom: 2rem;
        }

        .summary-card {
            background: var(--bg-card);
            border: 1px solid var(--border-color);
            border-radius: 8px;
            padding: 1.25rem;
            text-align: center;
        }

        .summary-card .value {
            font-size: 2rem;
            font-weight: 700;
            margin-bottom: 0.25rem;
        }

        .summary-card .label {
            color: var(--text-secondary);
            font-size: 0.875rem;
        }

        .summary-card.signal .value { color: var(--accent-signal); }
        .summary-card.customer .value { color: var(--accent-customer); }
        .summary-card.frustration .value { color: #dc3545; }
        .summary-card.arr .value { color: var(--accent-csat); }

        .section {
            margin-bottom: 2rem;
        }

        .section-title {
            font-size: 1.25rem;
            margin-bottom: 1rem;
            padding-bottom: 0.5rem;
            border-bottom: 2px solid var(--border-color);
        }

        .evidence-list {
            list-style: none;
        }

        .evidence-item {
            display: flex;
            align-items: center;
            gap: 0.75rem;
            padding: 0.75rem 1rem;
            background: var(--bg-card);
            border: 1px solid var(--border-color);
            border-radius: 6px;
            margin-bottom: 0.5rem;
        }

        .evidence-type {
            display: inline-block;
            padding: 0.25rem 0.5rem;
            border-radius: 4px;
            font-size: 0.75rem;
            font-weight: 600;
            text-transform: uppercase;
        }

        .evidence-type.signal { background: var(--accent-signal); color: white; }
        .evidence-type.customer { background: var(--accent-customer); color: white; }
        .evidence-type.market { background: var(--accent-market); color: white; }
        .evidence-type.idea { background: var(--accent-idea); color: white; }

        .evidence-id {
            font-family: monospace;
            color: var(--text-secondary);
        }

        .customer-list {
            display: flex;
            flex-wrap: wrap;
            gap: 0.5rem;
        }

        .customer-tag {
            background: var(--accent-customer);
            color: white;
            padding: 0.25rem 0.75rem;
            border-radius: 999px;
            font-size: 0.875rem;
        }

        .no-evidence {
            text-align: center;
            padding: 3rem;
            color: var(--text-secondary);
        }

        {{if not .HideGeneratedBy}}
        .footer {
            margin-top: 3rem;
            padding-top: 1rem;
            border-top: 1px solid var(--border-color);
            text-align: center;
            color: var(--text-secondary);
            font-size: 0.875rem;
        }
        {{end}}
    </style>
</head>
<body>
    <div class="container">
        <h1>{{.CapabilityName}}</h1>
        <p class="subtitle">Evidence Summary for {{.CapabilityID}}</p>

        {{if .Summary}}
        {{if .Summary.HasEvidence}}
        <div class="summary-grid">
            <div class="summary-card signal">
                <div class="value">{{.Summary.SignalCount}}</div>
                <div class="label">Signals</div>
            </div>
            <div class="summary-card customer">
                <div class="value">{{.Summary.UniqueCustomers}}</div>
                <div class="label">Customers</div>
            </div>
            <div class="summary-card frustration">
                <div class="value">{{printf "%.0f" .Summary.AvgFrustrationScore}}</div>
                <div class="label">Avg Frustration</div>
            </div>
            <div class="summary-card arr">
                <div class="value">${{printf "%.0f" .Summary.ARRInDollars}}</div>
                <div class="label">ARR Affected</div>
            </div>
            <div class="summary-card">
                <div class="value">{{.Summary.TotalVotes}}</div>
                <div class="label">Total Votes</div>
            </div>
        </div>

        {{if .Summary.CustomerNames}}
        <div class="section">
            <h2 class="section-title">Affected Customers</h2>
            <div class="customer-list">
                {{range .Summary.CustomerNames}}
                <span class="customer-tag">{{.}}</span>
                {{end}}
            </div>
        </div>
        {{end}}
        {{end}}
        {{end}}

        {{if .EvidenceLinks}}
        <div class="section">
            <h2 class="section-title">Evidence Links ({{len .EvidenceLinks}})</h2>
            <ul class="evidence-list">
                {{range .EvidenceLinks}}
                <li class="evidence-item">
                    <span class="evidence-type {{.Type}}">{{.Type}}</span>
                    <span class="evidence-id">{{.ID}}</span>
                </li>
                {{end}}
            </ul>
        </div>
        {{else}}
        <div class="no-evidence">
            <p>No evidence links configured for this capability.</p>
        </div>
        {{end}}

        {{if not .HideGeneratedBy}}
        <footer class="footer">
            Generated by PRISM
        </footer>
        {{end}}
    </div>
</body>
</html>
`
