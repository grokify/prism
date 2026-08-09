import { LitElement, html, css } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';

export interface Initiative {
  id: string;
  name: string;
  score: number;
  pillars: string[];
  evidenceCount: number;
  dimensionScores?: DimensionScore[];
}

export interface DimensionScore {
  dimension: string;
  rawValue: number;
  weight: number;
  weightedScore: number;
}

type SortColumn = 'score' | 'pillar' | 'evidence' | 'name';
type SortDirection = 'asc' | 'desc';

/**
 * InitiativeTable displays a sortable, expandable table of initiatives.
 * Responds to filter-change events from prism-dashboard-filters.
 *
 * @example
 * ```html
 * <prism-initiative-table>
 *   <script type="application/json">
 *     [{"id": "init-1", "name": "Auth Improvements", "score": 85, "pillars": ["CSAT"], "evidenceCount": 12}]
 *   </script>
 * </prism-initiative-table>
 * ```
 */
@customElement('prism-initiative-table')
export class InitiativeTable extends LitElement {
  static override styles = css`
    :host {
      --it-bg: var(--bg-card, #ffffff);
      --it-text: var(--text-primary, #1f2937);
      --it-text-muted: var(--text-secondary, #6b7280);
      --it-border: var(--border-color, #e5e7eb);
      --it-hover: var(--bg-secondary, #f8fafc);
      --it-success: #22c55e;
      --it-warning: #f59e0b;
      --it-danger: #ef4444;
      display: block;
    }

    :host([theme='dark']) {
      --it-bg: #1e293b;
      --it-text: #f1f5f9;
      --it-text-muted: #94a3b8;
      --it-border: #334155;
      --it-hover: #334155;
    }

    .table-container {
      overflow-x: auto;
    }

    table {
      width: 100%;
      border-collapse: collapse;
      font-size: 0.875rem;
    }

    th, td {
      padding: 12px 16px;
      text-align: left;
      border-bottom: 1px solid var(--it-border);
    }

    th {
      font-size: 0.75rem;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 0.05em;
      color: var(--it-text-muted);
      cursor: pointer;
      user-select: none;
      white-space: nowrap;
    }

    th:hover {
      color: var(--it-text);
    }

    th.sorted {
      color: var(--it-text);
    }

    th .sort-icon {
      margin-left: 4px;
      opacity: 0.5;
    }

    th.sorted .sort-icon {
      opacity: 1;
    }

    tbody tr {
      transition: background 0.15s ease;
    }

    tbody tr:hover {
      background: var(--it-hover);
    }

    tbody tr.expanded {
      background: var(--it-hover);
    }

    .expand-btn {
      background: none;
      border: none;
      cursor: pointer;
      padding: 4px;
      color: var(--it-text-muted);
      transition: transform 0.2s ease;
    }

    .expand-btn.expanded {
      transform: rotate(90deg);
    }

    .rank {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      width: 28px;
      height: 28px;
      border-radius: 50%;
      background: var(--it-hover);
      font-weight: 600;
    }

    .score {
      display: inline-block;
      padding: 4px 12px;
      border-radius: 9999px;
      font-weight: 600;
    }

    .score-high { background: rgba(34, 197, 94, 0.1); color: var(--it-success); }
    .score-medium { background: rgba(245, 158, 11, 0.1); color: var(--it-warning); }
    .score-low { background: rgba(239, 68, 68, 0.1); color: var(--it-danger); }

    .pillars {
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

    .pillar-chip.csat { background: rgba(16, 185, 129, 0.1); color: #10b981; }
    .pillar-chip.sam-som { background: rgba(59, 130, 246, 0.1); color: #3b82f6; }
    .pillar-chip.tam { background: rgba(139, 92, 246, 0.1); color: #8b5cf6; }

    .expansion-row td {
      padding: 0;
      border-bottom: 1px solid var(--it-border);
    }

    .expansion-content {
      padding: 16px 24px;
      background: var(--it-hover);
    }

    .dimension-grid {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
      gap: 12px;
    }

    .dimension-item {
      display: flex;
      flex-direction: column;
      gap: 4px;
    }

    .dimension-header {
      display: flex;
      justify-content: space-between;
      font-size: 0.75rem;
    }

    .dimension-name {
      color: var(--it-text-muted);
      text-transform: capitalize;
    }

    .dimension-value {
      font-weight: 600;
    }

    .dimension-bar {
      height: 6px;
      background: var(--it-border);
      border-radius: 3px;
      overflow: hidden;
    }

    .dimension-fill {
      height: 100%;
      background: var(--it-success);
      border-radius: 3px;
    }

    .empty-state {
      text-align: center;
      padding: 48px 24px;
      color: var(--it-text-muted);
    }
  `;

  @property({ type: Array }) initiatives: Initiative[] = [];
  @property({ type: String, reflect: true }) theme: 'light' | 'dark' = 'light';

  @state() private sortColumn: SortColumn = 'score';
  @state() private sortDirection: SortDirection = 'desc';
  @state() private expandedRows: Set<string> = new Set();
  @state() private filterPillar: string | null = null;

  override connectedCallback() {
    super.connectedCallback();
    this.loadFromScript();
    this.addEventListener('filter-change', this.handleFilterChange as EventListener);
    document.addEventListener('filter-change', this.handleFilterChange as EventListener);
  }

  override disconnectedCallback() {
    super.disconnectedCallback();
    document.removeEventListener('filter-change', this.handleFilterChange as EventListener);
  }

  private loadFromScript() {
    const script = this.querySelector('script[type="application/json"]');
    if (script?.textContent) {
      try {
        this.initiatives = JSON.parse(script.textContent);
      } catch (e) {
        console.error('Failed to parse initiative data:', e);
      }
    }
  }

  private handleFilterChange = (e: CustomEvent) => {
    this.filterPillar = e.detail?.pillar || null;
  };

  private toggleSort(column: SortColumn) {
    if (this.sortColumn === column) {
      this.sortDirection = this.sortDirection === 'asc' ? 'desc' : 'asc';
    } else {
      this.sortColumn = column;
      this.sortDirection = column === 'name' ? 'asc' : 'desc';
    }
  }

  private toggleExpand(id: string) {
    if (this.expandedRows.has(id)) {
      this.expandedRows.delete(id);
    } else {
      this.expandedRows.add(id);
    }
    this.requestUpdate();
  }

  private getSortedInitiatives(): Initiative[] {
    let filtered = this.initiatives;

    if (this.filterPillar) {
      filtered = filtered.filter(init =>
        init.pillars.some(p => p === this.filterPillar)
      );
    }

    return [...filtered].sort((a, b) => {
      let comparison = 0;
      switch (this.sortColumn) {
        case 'score':
          comparison = a.score - b.score;
          break;
        case 'evidence':
          comparison = a.evidenceCount - b.evidenceCount;
          break;
        case 'name':
          comparison = a.name.localeCompare(b.name);
          break;
        case 'pillar':
          comparison = (a.pillars[0] || '').localeCompare(b.pillars[0] || '');
          break;
      }
      return this.sortDirection === 'asc' ? comparison : -comparison;
    });
  }

  private getScoreClass(score: number): string {
    if (score >= 70) return 'score-high';
    if (score >= 40) return 'score-medium';
    return 'score-low';
  }

  private getPillarClass(pillar: string): string {
    return pillar.toLowerCase().replace(' ', '-');
  }

  private renderSortIcon(column: SortColumn): string {
    if (this.sortColumn !== column) return '↕';
    return this.sortDirection === 'asc' ? '↑' : '↓';
  }

  override render() {
    const initiatives = this.getSortedInitiatives();

    if (initiatives.length === 0) {
      return html`
        <div class="empty-state">
          <p>No initiatives found${this.filterPillar ? ` for ${this.filterPillar}` : ''}</p>
        </div>
      `;
    }

    return html`
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th style="width: 40px;"></th>
              <th>#</th>
              <th
                class="${this.sortColumn === 'name' ? 'sorted' : ''}"
                @click=${() => this.toggleSort('name')}
              >
                Initiative <span class="sort-icon">${this.renderSortIcon('name')}</span>
              </th>
              <th
                class="${this.sortColumn === 'score' ? 'sorted' : ''}"
                @click=${() => this.toggleSort('score')}
              >
                Score <span class="sort-icon">${this.renderSortIcon('score')}</span>
              </th>
              <th
                class="${this.sortColumn === 'pillar' ? 'sorted' : ''}"
                @click=${() => this.toggleSort('pillar')}
              >
                Pillars <span class="sort-icon">${this.renderSortIcon('pillar')}</span>
              </th>
              <th
                class="${this.sortColumn === 'evidence' ? 'sorted' : ''}"
                @click=${() => this.toggleSort('evidence')}
              >
                Evidence <span class="sort-icon">${this.renderSortIcon('evidence')}</span>
              </th>
            </tr>
          </thead>
          <tbody>
            ${initiatives.map((init, index) => html`
              <tr class="${this.expandedRows.has(init.id) ? 'expanded' : ''}">
                <td>
                  <button
                    class="expand-btn ${this.expandedRows.has(init.id) ? 'expanded' : ''}"
                    @click=${() => this.toggleExpand(init.id)}
                  >
                    ▶
                  </button>
                </td>
                <td><span class="rank">${index + 1}</span></td>
                <td>${init.name}</td>
                <td>
                  <span class="score ${this.getScoreClass(init.score)}">
                    ${init.score.toFixed(0)}
                  </span>
                </td>
                <td>
                  <div class="pillars">
                    ${init.pillars.map(pillar => html`
                      <span class="pillar-chip ${this.getPillarClass(pillar)}">${pillar}</span>
                    `)}
                  </div>
                </td>
                <td>${init.evidenceCount}</td>
              </tr>
              ${this.expandedRows.has(init.id) && init.dimensionScores ? html`
                <tr class="expansion-row">
                  <td colspan="6">
                    <div class="expansion-content">
                      <div class="dimension-grid">
                        ${init.dimensionScores.map(dim => html`
                          <div class="dimension-item">
                            <div class="dimension-header">
                              <span class="dimension-name">${dim.dimension.replace('_', ' ')}</span>
                              <span class="dimension-value">${(dim.rawValue * 100).toFixed(0)}%</span>
                            </div>
                            <div class="dimension-bar">
                              <div class="dimension-fill" style="width: ${dim.rawValue * 100}%"></div>
                            </div>
                          </div>
                        `)}
                      </div>
                    </div>
                  </td>
                </tr>
              ` : ''}
            `)}
          </tbody>
        </table>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'prism-initiative-table': InitiativeTable;
  }
}
