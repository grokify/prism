package sitegen

// pillarDashboardTemplate is the HTML template for the pillar investment dashboard.
const pillarDashboardTemplate = `<!DOCTYPE html>
<html lang="en" theme="{{.Theme}}">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Portfolio Pillars - {{.Title}}</title>
  <link rel="stylesheet" href="{{.BaseURL}}/assets/site.css">
  {{if .HasSiteNavJS}}<script type="module" src="{{.BaseURL}}/assets/site-nav.es.js"></script>{{end}}
  <style>
    .pillar-dashboard {
      max-width: 1200px;
      margin: 0 auto;
    }

    .pillar-summary {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 24px;
      margin-bottom: 40px;
    }

    .pillar-card {
      background: var(--bg-secondary);
      border: 1px solid var(--border);
      border-radius: var(--radius-lg);
      padding: 24px;
      transition: transform 0.2s, box-shadow 0.2s;
    }

    .pillar-card:hover {
      transform: translateY(-2px);
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    }

    .pillar-header {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      margin-bottom: 16px;
    }

    .pillar-name {
      font-size: 1.25rem;
      font-weight: 600;
      color: var(--text-primary);
      margin: 0;
    }

    .pillar-percentage {
      font-size: 1.5rem;
      font-weight: 700;
      color: var(--accent);
    }

    .pillar-description {
      font-size: 0.875rem;
      color: var(--text-secondary);
      margin-bottom: 16px;
    }

    .pillar-stats {
      display: flex;
      gap: 24px;
      padding-top: 16px;
      border-top: 1px solid var(--border);
    }

    .pillar-stat {
      text-align: center;
    }

    .pillar-stat-value {
      font-size: 1.25rem;
      font-weight: 600;
      color: var(--text-primary);
    }

    .pillar-stat-label {
      font-size: 0.75rem;
      color: var(--text-muted);
      text-transform: uppercase;
    }

    /* Investment Mix Chart */
    .investment-mix {
      background: var(--bg-secondary);
      border: 1px solid var(--border);
      border-radius: var(--radius-lg);
      padding: 24px;
      margin-bottom: 40px;
    }

    .investment-mix h2 {
      font-size: 1.125rem;
      font-weight: 600;
      margin-bottom: 20px;
      color: var(--text-primary);
    }

    .mix-bar-container {
      height: 48px;
      display: flex;
      border-radius: var(--radius-md);
      overflow: hidden;
      margin-bottom: 16px;
    }

    .mix-bar-segment {
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      font-weight: 600;
      font-size: 0.875rem;
      transition: flex-grow 0.3s;
    }

    .mix-bar-segment.csat {
      background: linear-gradient(135deg, #22c55e, #16a34a);
    }

    .mix-bar-segment.sam-som {
      background: linear-gradient(135deg, #3b82f6, #2563eb);
    }

    .mix-bar-segment.tam {
      background: linear-gradient(135deg, #8b5cf6, #7c3aed);
    }

    .mix-legend {
      display: flex;
      justify-content: center;
      gap: 32px;
    }

    .mix-legend-item {
      display: flex;
      align-items: center;
      gap: 8px;
      font-size: 0.875rem;
      color: var(--text-secondary);
    }

    .mix-legend-color {
      width: 12px;
      height: 12px;
      border-radius: 2px;
    }

    .mix-legend-color.csat { background: #22c55e; }
    .mix-legend-color.sam-som { background: #3b82f6; }
    .mix-legend-color.tam { background: #8b5cf6; }

    /* Top Initiatives */
    .top-initiatives {
      background: var(--bg-secondary);
      border: 1px solid var(--border);
      border-radius: var(--radius-lg);
      padding: 24px;
    }

    .top-initiatives h2 {
      font-size: 1.125rem;
      font-weight: 600;
      margin-bottom: 20px;
      color: var(--text-primary);
    }

    .initiatives-by-pillar {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 24px;
    }

    .pillar-initiatives {
      background: var(--bg-tertiary);
      border-radius: var(--radius-md);
      padding: 16px;
    }

    .pillar-initiatives h3 {
      font-size: 1rem;
      font-weight: 600;
      margin-bottom: 12px;
      color: var(--text-primary);
    }

    .initiative-list {
      list-style: none;
      padding: 0;
      margin: 0;
    }

    .initiative-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 10px 0;
      border-bottom: 1px solid var(--border);
    }

    .initiative-item:last-child {
      border-bottom: none;
    }

    .initiative-name {
      font-size: 0.875rem;
      color: var(--text-primary);
    }

    .initiative-level {
      font-size: 0.75rem;
      font-weight: 500;
      padding: 2px 8px;
      border-radius: var(--radius-sm);
    }

    .initiative-level.high {
      background: #dcfce7;
      color: #166534;
    }

    .initiative-level.medium {
      background: #dbeafe;
      color: #1e40af;
    }

    .initiative-level.low {
      background: #f3f4f6;
      color: #6b7280;
    }

    /* Portfolio Summary */
    .portfolio-summary {
      display: flex;
      justify-content: center;
      gap: 48px;
      padding: 24px;
      background: var(--bg-secondary);
      border: 1px solid var(--border);
      border-radius: var(--radius-lg);
      margin-bottom: 40px;
    }

    .summary-stat {
      text-align: center;
    }

    .summary-stat-value {
      font-size: 2rem;
      font-weight: 700;
      color: var(--accent);
    }

    .summary-stat-label {
      font-size: 0.875rem;
      color: var(--text-muted);
    }
  </style>
</head>
<body>
  {{if .HasSiteNavJS}}
  <wt-navbar id="navbar" theme="{{.Theme}}"></wt-navbar>
  {{end}}

  <div class="site-container">
    {{if not .HasSiteNavJS}}
    <header class="fallback-header">
      <h1>Portfolio Pillars</h1>
      <p>Investment distribution across strategic pillars</p>
    </header>
    {{end}}

    <main class="site-main pillar-dashboard">
      <!-- Portfolio Summary -->
      <section class="portfolio-summary">
        <div class="summary-stat">
          <div class="summary-stat-value">{{.Rollup.TotalInitiatives}}</div>
          <div class="summary-stat-label">Total Initiatives</div>
        </div>
        <div class="summary-stat">
          <div class="summary-stat-value">{{.Rollup.TotalWeight}}</div>
          <div class="summary-stat-label">Weighted Score</div>
        </div>
      </section>

      <!-- Investment Mix Bar -->
      <section class="investment-mix">
        <h2>Investment Mix</h2>
        <div class="mix-bar-container">
          {{range .Rollup.ByPillar}}
          {{if gt .Percentage 0}}
          <div class="mix-bar-segment {{.Pillar | pillarClass}}" style="flex-grow: {{.Percentage}};">
            {{printf "%.0f" .Percentage}}%
          </div>
          {{end}}
          {{end}}
        </div>
        <div class="mix-legend">
          <div class="mix-legend-item">
            <div class="mix-legend-color csat"></div>
            <span>CSAT</span>
          </div>
          <div class="mix-legend-item">
            <div class="mix-legend-color sam-som"></div>
            <span>SAM-SOM</span>
          </div>
          <div class="mix-legend-item">
            <div class="mix-legend-color tam"></div>
            <span>TAM</span>
          </div>
        </div>
      </section>

      <!-- Pillar Cards -->
      <section class="pillar-summary">
        {{range .Details}}
        <div class="pillar-card">
          <div class="pillar-header">
            <h3 class="pillar-name">{{.Pillar}}</h3>
            <span class="pillar-percentage">{{printf "%.1f" .Percentage}}%</span>
          </div>
          <p class="pillar-description">{{.Description}}</p>
          <div class="pillar-stats">
            <div class="pillar-stat">
              <div class="pillar-stat-value">{{.Count}}</div>
              <div class="pillar-stat-label">Initiatives</div>
            </div>
          </div>
        </div>
        {{end}}
      </section>

      <!-- Top Initiatives by Pillar -->
      <section class="top-initiatives">
        <h2>Top Initiatives by Pillar</h2>
        <div class="initiatives-by-pillar">
          {{range .Details}}
          {{if .TopInitiatives}}
          <div class="pillar-initiatives">
            <h3>{{.Pillar}}</h3>
            <ul class="initiative-list">
              {{range .TopInitiatives}}
              <li class="initiative-item">
                <span class="initiative-name">{{.Name}}</span>
                <span class="initiative-level {{.Level}}">{{.Level}}</span>
              </li>
              {{end}}
            </ul>
          </div>
          {{end}}
          {{end}}
        </div>
      </section>
    </main>

    {{if not .HideGeneratedBy}}
    <footer class="site-footer">
      <p>Generated by <a href="https://github.com/grokify/prism">PRISM</a></p>
    </footer>
    {{end}}
  </div>
</body>
</html>
`
