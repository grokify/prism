package sitegen

// dashboardTemplate is the HTML template for the unified portfolio dashboard.
const dashboardTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.Title}} - Portfolio Dashboard</title>
  <style>
    :root {
      --bg-primary: #ffffff;
      --bg-secondary: #f8fafc;
      --bg-card: #ffffff;
      --text-primary: #1f2937;
      --text-secondary: #6b7280;
      --border-color: #e5e7eb;
      --accent-csat: #10b981;
      --accent-sam: #3b82f6;
      --accent-tam: #8b5cf6;
      --accent-success: #22c55e;
      --accent-warning: #f59e0b;
      --accent-danger: #ef4444;
    }

    {{if eq .Theme "dark"}}
    :root {
      --bg-primary: #0f172a;
      --bg-secondary: #1e293b;
      --bg-card: #1e293b;
      --text-primary: #f1f5f9;
      --text-secondary: #94a3b8;
      --border-color: #334155;
    }
    {{end}}

    * {
      box-sizing: border-box;
      margin: 0;
      padding: 0;
    }

    body {
      font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      background: var(--bg-primary);
      color: var(--text-primary);
      line-height: 1.5;
    }

    .dashboard {
      max-width: 1400px;
      margin: 0 auto;
      padding: 24px;
    }

    .dashboard-header {
      text-align: center;
      margin-bottom: 32px;
    }

    .dashboard-header h1 {
      font-size: 2rem;
      font-weight: 700;
      margin-bottom: 8px;
    }

    .dashboard-header p {
      color: var(--text-secondary);
    }

    .stats-row {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
      gap: 16px;
      margin-bottom: 32px;
    }

    .stat-card {
      background: var(--bg-card);
      border: 1px solid var(--border-color);
      border-radius: 12px;
      padding: 20px;
      text-align: center;
    }

    .stat-value {
      font-size: 2.5rem;
      font-weight: 700;
      color: var(--text-primary);
    }

    .stat-label {
      font-size: 0.875rem;
      color: var(--text-secondary);
      text-transform: uppercase;
      letter-spacing: 0.05em;
    }

    .dashboard-grid {
      display: grid;
      grid-template-columns: repeat(2, 1fr);
      gap: 24px;
    }

    @media (max-width: 1024px) {
      .dashboard-grid {
        grid-template-columns: 1fr;
      }
    }

    .card {
      background: var(--bg-card);
      border: 1px solid var(--border-color);
      border-radius: 12px;
      padding: 24px;
    }

    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20px;
    }

    .card-title {
      font-size: 1.25rem;
      font-weight: 600;
    }

    .card-full {
      grid-column: 1 / -1;
    }

    /* Pillar Investment Mix */
    .pillar-bars {
      display: flex;
      flex-direction: column;
      gap: 12px;
    }

    .pillar-bar {
      display: flex;
      align-items: center;
      gap: 12px;
    }

    .pillar-label {
      width: 80px;
      font-weight: 500;
      font-size: 0.875rem;
    }

    .pillar-track {
      flex: 1;
      height: 24px;
      background: var(--bg-secondary);
      border-radius: 4px;
      overflow: hidden;
    }

    .pillar-fill {
      height: 100%;
      border-radius: 4px;
      transition: width 0.3s ease;
    }

    .pillar-fill.csat { background: var(--accent-csat); }
    .pillar-fill.sam-som { background: var(--accent-sam); }
    .pillar-fill.tam { background: var(--accent-tam); }

    .pillar-value {
      width: 50px;
      text-align: right;
      font-weight: 600;
      font-size: 0.875rem;
    }

    /* Initiative Table */
    .initiative-table {
      width: 100%;
      border-collapse: collapse;
    }

    .initiative-table th,
    .initiative-table td {
      padding: 12px;
      text-align: left;
      border-bottom: 1px solid var(--border-color);
    }

    .initiative-table th {
      font-size: 0.75rem;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 0.05em;
      color: var(--text-secondary);
    }

    .initiative-table tr:hover {
      background: var(--bg-secondary);
    }

    .rank-badge {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      width: 28px;
      height: 28px;
      border-radius: 50%;
      background: var(--bg-secondary);
      font-weight: 600;
      font-size: 0.875rem;
    }

    .score-badge {
      display: inline-block;
      padding: 4px 12px;
      border-radius: 9999px;
      font-weight: 600;
      font-size: 0.875rem;
    }

    .score-high { background: rgba(34, 197, 94, 0.1); color: var(--accent-success); }
    .score-medium { background: rgba(245, 158, 11, 0.1); color: var(--accent-warning); }
    .score-low { background: rgba(239, 68, 68, 0.1); color: var(--accent-danger); }

    .pillar-chips {
      display: flex;
      gap: 4px;
      flex-wrap: wrap;
    }

    .pillar-chip {
      padding: 2px 8px;
      border-radius: 4px;
      font-size: 0.75rem;
      font-weight: 500;
    }

    .pillar-chip.csat { background: rgba(16, 185, 129, 0.1); color: var(--accent-csat); }
    .pillar-chip.sam-som { background: rgba(59, 130, 246, 0.1); color: var(--accent-sam); }
    .pillar-chip.tam { background: rgba(139, 92, 246, 0.1); color: var(--accent-tam); }

    /* Maturity Heatmap */
    .heatmap {
      display: grid;
      gap: 8px;
    }

    .heatmap-row {
      display: flex;
      align-items: center;
      gap: 8px;
    }

    .heatmap-label {
      width: 120px;
      font-size: 0.875rem;
      font-weight: 500;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .heatmap-cells {
      display: flex;
      gap: 4px;
      flex: 1;
    }

    .heatmap-cell {
      flex: 1;
      height: 32px;
      border-radius: 4px;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 0.75rem;
      font-weight: 600;
    }

    .heatmap-cell.active { background: var(--accent-success); color: white; }
    .heatmap-cell.planned { background: var(--accent-warning); color: white; }
    .heatmap-cell.proposed { background: var(--accent-danger); color: white; }
    .heatmap-cell.unknown { background: var(--bg-secondary); color: var(--text-secondary); }

    /* Gap Summary */
    .gap-list {
      display: flex;
      flex-direction: column;
      gap: 12px;
    }

    .gap-item {
      display: flex;
      align-items: flex-start;
      gap: 12px;
      padding: 12px;
      background: var(--bg-secondary);
      border-radius: 8px;
    }

    .gap-severity {
      width: 8px;
      height: 8px;
      border-radius: 50%;
      margin-top: 6px;
    }

    .gap-severity.high { background: var(--accent-danger); }
    .gap-severity.medium { background: var(--accent-warning); }
    .gap-severity.low { background: var(--accent-success); }

    .gap-content {
      flex: 1;
    }

    .gap-title {
      font-weight: 500;
    }

    .gap-type {
      font-size: 0.75rem;
      color: var(--text-secondary);
    }

    /* Filter Slot */
    .filter-section {
      margin-bottom: 24px;
    }

    /* Footer */
    .dashboard-footer {
      text-align: center;
      margin-top: 48px;
      padding-top: 24px;
      border-top: 1px solid var(--border-color);
      color: var(--text-secondary);
      font-size: 0.875rem;
    }
  </style>
</head>
<body>
  <div class="dashboard">
    <header class="dashboard-header">
      <h1>{{.Title}}</h1>
      <p>Portfolio Investment Dashboard</p>
    </header>

    <!-- Stats Row -->
    <div class="stats-row">
      <div class="stat-card">
        <div class="stat-value">{{.Dashboard.TotalCapabilities}}</div>
        <div class="stat-label">Capabilities</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{.Dashboard.TotalInitiatives}}</div>
        <div class="stat-label">Initiatives</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{.Dashboard.TotalMetrics}}</div>
        <div class="stat-label">Metrics</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{.Dashboard.PillarRollup.TotalWeight}}</div>
        <div class="stat-label">Investment Score</div>
      </div>
    </div>

    <!-- Filter Section (slot for prism-dashboard-filters) -->
    <div class="filter-section">
      <prism-dashboard-filters></prism-dashboard-filters>
    </div>

    <div class="dashboard-grid">
      <!-- Pillar Investment Mix -->
      <div class="card">
        <div class="card-header">
          <h2 class="card-title">Pillar Investment Mix</h2>
        </div>
        <div class="pillar-bars">
          {{range .Dashboard.PillarRollup.ByPillar}}
          <div class="pillar-bar">
            <div class="pillar-label">{{.Pillar}}</div>
            <div class="pillar-track">
              <div class="pillar-fill {{if eq .Pillar "CSAT"}}csat{{else if eq .Pillar "SAM-SOM"}}sam-som{{else}}tam{{end}}" style="width: {{printf "%.1f" .Percentage}}%"></div>
            </div>
            <div class="pillar-value">{{printf "%.0f" .Percentage}}%</div>
          </div>
          {{end}}
        </div>
      </div>

      <!-- Gap Summary -->
      {{if .Dashboard.GapSummary}}
      <div class="card">
        <div class="card-header">
          <h2 class="card-title">Gap Analysis</h2>
          <span>{{.Dashboard.GapSummary.TotalGaps}} gaps</span>
        </div>
        <div class="gap-list">
          {{range .Dashboard.GapSummary.TopGaps}}
          <div class="gap-item">
            <div class="gap-severity {{.Severity}}"></div>
            <div class="gap-content">
              <div class="gap-title">{{.ID}}</div>
              <div class="gap-type">{{.Type}} - {{.Severity}}</div>
            </div>
          </div>
          {{end}}
        </div>
      </div>
      {{end}}

      <!-- Top Initiatives -->
      <div class="card card-full">
        <div class="card-header">
          <h2 class="card-title">Top Initiatives</h2>
        </div>
        <prism-initiative-table>
          <table class="initiative-table">
            <thead>
              <tr>
                <th>Rank</th>
                <th>Initiative</th>
                <th>Score</th>
                <th>Pillars</th>
                <th>Evidence</th>
              </tr>
            </thead>
            <tbody>
              {{range .Dashboard.TopInitiatives}}
              <tr data-id="{{.ID}}">
                <td><span class="rank-badge">{{.Rank}}</span></td>
                <td>{{.Name}}</td>
                <td>
                  <span class="score-badge {{if ge .Score 70.0}}score-high{{else if ge .Score 40.0}}score-medium{{else}}score-low{{end}}">
                    {{printf "%.0f" .Score}}
                  </span>
                </td>
                <td>
                  <div class="pillar-chips">
                    {{range .Pillars}}
                    <span class="pillar-chip {{if eq .Pillar "CSAT"}}csat{{else if eq .Pillar "SAM-SOM"}}sam-som{{else}}tam{{end}}">{{.Pillar}}</span>
                    {{end}}
                  </div>
                </td>
                <td>{{.EvidenceCount}}</td>
              </tr>
              {{end}}
            </tbody>
          </table>
        </prism-initiative-table>
      </div>

      <!-- Maturity Heatmap -->
      {{if .Dashboard.MaturityHeatmap}}
      <div class="card card-full">
        <div class="card-header">
          <h2 class="card-title">Capability Maturity by Domain</h2>
        </div>
        <div class="heatmap">
          {{range .Dashboard.MaturityHeatmap}}
          <div class="heatmap-row">
            <div class="heatmap-label" title="{{.Domain}}">{{.Domain}}</div>
            <div class="heatmap-cells">
              {{range $status, $count := .StatusCount}}
              <div class="heatmap-cell {{$status}}" title="{{$status}}: {{$count}}">{{$count}}</div>
              {{end}}
            </div>
          </div>
          {{end}}
        </div>
      </div>
      {{end}}
    </div>

    {{if not .HideGeneratedBy}}
    <footer class="dashboard-footer">
      Generated by PRISM
    </footer>
    {{end}}
  </div>

  {{if .HasSiteNavJS}}
  <script src="{{.BaseURL}}/site-nav.js"></script>
  {{end}}
</body>
</html>`

// selfContainedDashboardTemplate is a fully self-contained HTML export.
const selfContainedDashboardTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.Title}}</title>
  <style>
    :root {
      --bg-primary: #ffffff;
      --bg-secondary: #f8fafc;
      --bg-card: #ffffff;
      --text-primary: #1f2937;
      --text-secondary: #6b7280;
      --border-color: #e5e7eb;
      --accent-csat: #10b981;
      --accent-sam: #3b82f6;
      --accent-tam: #8b5cf6;
      --accent-success: #22c55e;
      --accent-warning: #f59e0b;
      --accent-danger: #ef4444;
    }

    @media (prefers-color-scheme: dark) {
      :root {
        --bg-primary: #0f172a;
        --bg-secondary: #1e293b;
        --bg-card: #1e293b;
        --text-primary: #f1f5f9;
        --text-secondary: #94a3b8;
        --border-color: #334155;
      }
    }

    * { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: system-ui, -apple-system, sans-serif;
      background: var(--bg-primary);
      color: var(--text-primary);
      line-height: 1.5;
      padding: 24px;
    }

    .dashboard { max-width: 1200px; margin: 0 auto; }
    .header { text-align: center; margin-bottom: 32px; }
    .header h1 { font-size: 2rem; font-weight: 700; }

    .stats { display: grid; grid-template-columns: repeat(auto-fit, minmax(150px, 1fr)); gap: 16px; margin-bottom: 32px; }
    .stat { background: var(--bg-card); border: 1px solid var(--border-color); border-radius: 12px; padding: 20px; text-align: center; }
    .stat-value { font-size: 2rem; font-weight: 700; }
    .stat-label { font-size: 0.75rem; color: var(--text-secondary); text-transform: uppercase; }

    .grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: 24px; }
    @media (max-width: 768px) { .grid { grid-template-columns: 1fr; } }

    .card { background: var(--bg-card); border: 1px solid var(--border-color); border-radius: 12px; padding: 24px; }
    .card-title { font-size: 1.125rem; font-weight: 600; margin-bottom: 16px; }
    .card-full { grid-column: 1 / -1; }

    .pillar-bar { display: flex; align-items: center; gap: 12px; margin-bottom: 8px; }
    .pillar-label { width: 80px; font-size: 0.875rem; font-weight: 500; }
    .pillar-track { flex: 1; height: 20px; background: var(--bg-secondary); border-radius: 4px; overflow: hidden; }
    .pillar-fill { height: 100%; border-radius: 4px; }
    .pillar-fill.csat { background: var(--accent-csat); }
    .pillar-fill.sam-som { background: var(--accent-sam); }
    .pillar-fill.tam { background: var(--accent-tam); }
    .pillar-pct { width: 50px; text-align: right; font-weight: 600; }

    table { width: 100%; border-collapse: collapse; }
    th, td { padding: 12px; text-align: left; border-bottom: 1px solid var(--border-color); }
    th { font-size: 0.75rem; text-transform: uppercase; color: var(--text-secondary); }

    .score { display: inline-block; padding: 4px 12px; border-radius: 9999px; font-weight: 600; }
    .score-high { background: rgba(34, 197, 94, 0.1); color: var(--accent-success); }
    .score-med { background: rgba(245, 158, 11, 0.1); color: var(--accent-warning); }
    .score-low { background: rgba(239, 68, 68, 0.1); color: var(--accent-danger); }

    .footer { text-align: center; margin-top: 48px; color: var(--text-secondary); font-size: 0.875rem; }
  </style>
</head>
<body>
  <div class="dashboard">
    <header class="header">
      <h1>{{.Title}}</h1>
    </header>

    <div class="stats">
      <div class="stat">
        <div class="stat-value">{{.Dashboard.TotalCapabilities}}</div>
        <div class="stat-label">Capabilities</div>
      </div>
      <div class="stat">
        <div class="stat-value">{{.Dashboard.TotalInitiatives}}</div>
        <div class="stat-label">Initiatives</div>
      </div>
      <div class="stat">
        <div class="stat-value">{{.Dashboard.TotalMetrics}}</div>
        <div class="stat-label">Metrics</div>
      </div>
    </div>

    <div class="grid">
      <div class="card">
        <h2 class="card-title">Pillar Investment</h2>
        {{range .Dashboard.PillarRollup.ByPillar}}
        <div class="pillar-bar">
          <div class="pillar-label">{{.Pillar}}</div>
          <div class="pillar-track">
            <div class="pillar-fill {{if eq .Pillar "CSAT"}}csat{{else if eq .Pillar "SAM-SOM"}}sam-som{{else}}tam{{end}}" style="width: {{printf "%.1f" .Percentage}}%"></div>
          </div>
          <div class="pillar-pct">{{printf "%.0f" .Percentage}}%</div>
        </div>
        {{end}}
      </div>

      {{if .Dashboard.GapSummary}}
      <div class="card">
        <h2 class="card-title">Gaps: {{.Dashboard.GapSummary.TotalGaps}}</h2>
        <p style="color: var(--text-secondary);">High: {{index .Dashboard.GapSummary.BySeverity "high"}}, Medium: {{index .Dashboard.GapSummary.BySeverity "medium"}}, Low: {{index .Dashboard.GapSummary.BySeverity "low"}}</p>
      </div>
      {{end}}

      <div class="card card-full">
        <h2 class="card-title">Top Initiatives</h2>
        <table>
          <thead>
            <tr><th>#</th><th>Initiative</th><th>Score</th><th>Evidence</th></tr>
          </thead>
          <tbody>
            {{range .Dashboard.TopInitiatives}}
            <tr>
              <td>{{.Rank}}</td>
              <td>{{.Name}}</td>
              <td><span class="score {{if ge .Score 70.0}}score-high{{else if ge .Score 40.0}}score-med{{else}}score-low{{end}}">{{printf "%.0f" .Score}}</span></td>
              <td>{{.EvidenceCount}}</td>
            </tr>
            {{end}}
          </tbody>
        </table>
      </div>
    </div>

    <footer class="footer">Exported from PRISM</footer>
  </div>

  {{if .UIBundle}}
  <script type="module">
{{.UIBundle}}
  </script>
  {{end}}
</body>
</html>`
