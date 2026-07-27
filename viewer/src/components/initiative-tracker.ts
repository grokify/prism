import { LitElement, html, css, PropertyValues } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';
import type { JourneyRoadmap, Initiative, InitiativeStatus, Team } from '../schema/roadmap/index.js';

type GroupBy = 'status' | 'team' | 'period';
type SortBy = 'name' | 'status' | 'team';

const STATUS_ORDER: Record<InitiativeStatus, number> = {
  proposed: 0,
  planned: 1,
  in_progress: 2,
  completed: 3,
  on_hold: 4,
  cancelled: 5,
};

const STATUS_COLORS: Record<InitiativeStatus, { bg: string; text: string }> = {
  proposed: { bg: '#e0e7ff', text: '#3730a3' },
  planned: { bg: '#dbeafe', text: '#1e40af' },
  in_progress: { bg: '#fef3c7', text: '#92400e' },
  completed: { bg: '#d1fae5', text: '#065f46' },
  on_hold: { bg: '#fce7f3', text: '#9d174d' },
  cancelled: { bg: '#f3f4f6', text: '#6b7280' },
};

/**
 * InitiativeTracker displays initiatives with filtering and grouping options.
 *
 * @example
 * ```html
 * <initiative-tracker group-by="team" theme="dark">
 *   <script type="application/json">
 *     {"initiatives": [...], "teams": [...]}
 *   </script>
 * </initiative-tracker>
 * ```
 */
@customElement('initiative-tracker')
export class InitiativeTracker extends LitElement {
  static override styles = css`
    :host {
      --it-bg: #ffffff;
      --it-text: #1f2937;
      --it-border: #e5e7eb;
      --it-card-bg: #f8fafc;
      --it-card-hover: #f1f5f9;
      display: block;
      font-family: system-ui, -apple-system, sans-serif;
    }

    :host([theme='dark']) {
      --it-bg: #0f172a;
      --it-text: #f1f5f9;
      --it-border: #334155;
      --it-card-bg: #1e293b;
      --it-card-hover: #334155;
    }

    .container {
      background: var(--it-bg);
      color: var(--it-text);
      padding: 24px;
      min-height: 100%;
      box-sizing: border-box;
    }

    .header {
      margin-bottom: 24px;
    }

    .title {
      font-size: 1.75rem;
      font-weight: 700;
      margin: 0 0 8px 0;
    }

    .subtitle {
      font-size: 0.875rem;
      opacity: 0.7;
      margin: 0;
    }

    .controls {
      display: flex;
      gap: 16px;
      margin-bottom: 20px;
      flex-wrap: wrap;
      align-items: center;
    }

    .control-group {
      display: flex;
      gap: 8px;
      align-items: center;
    }

    .control-label {
      font-size: 0.75rem;
      font-weight: 600;
      text-transform: uppercase;
      opacity: 0.7;
    }

    .btn {
      padding: 6px 12px;
      border-radius: 6px;
      font-size: 0.75rem;
      font-weight: 500;
      cursor: pointer;
      border: 1px solid var(--it-border);
      background: transparent;
      color: var(--it-text);
      transition: all 0.15s ease;
    }

    .btn:hover {
      background: var(--it-card-hover);
    }

    .btn.active {
      background: var(--it-text);
      color: var(--it-bg);
    }

    .search-input {
      padding: 8px 12px;
      border-radius: 6px;
      border: 1px solid var(--it-border);
      background: var(--it-bg);
      color: var(--it-text);
      font-size: 0.875rem;
      min-width: 200px;
    }

    .search-input::placeholder {
      opacity: 0.5;
    }

    .stats {
      display: flex;
      gap: 16px;
      margin-bottom: 20px;
      flex-wrap: wrap;
    }

    .stat-card {
      background: var(--it-card-bg);
      border: 1px solid var(--it-border);
      border-radius: 8px;
      padding: 12px 16px;
      min-width: 100px;
    }

    .stat-value {
      font-size: 1.5rem;
      font-weight: 700;
    }

    .stat-label {
      font-size: 0.75rem;
      opacity: 0.7;
      text-transform: uppercase;
    }

    .groups {
      display: flex;
      flex-direction: column;
      gap: 24px;
    }

    .group {
      background: var(--it-card-bg);
      border: 1px solid var(--it-border);
      border-radius: 12px;
      overflow: hidden;
    }

    .group-header {
      padding: 16px 20px;
      border-bottom: 1px solid var(--it-border);
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .group-title {
      font-size: 1rem;
      font-weight: 600;
      margin: 0;
    }

    .group-count {
      font-size: 0.75rem;
      opacity: 0.7;
      background: var(--it-border);
      padding: 4px 8px;
      border-radius: 12px;
    }

    .initiatives {
      padding: 12px;
      display: grid;
      gap: 12px;
    }

    .initiative-card {
      background: var(--it-bg);
      border: 1px solid var(--it-border);
      border-radius: 8px;
      padding: 16px;
      cursor: pointer;
      transition: all 0.15s ease;
    }

    .initiative-card:hover {
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    }

    .initiative-header {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      margin-bottom: 8px;
    }

    .initiative-name {
      font-weight: 600;
      margin: 0;
      flex: 1;
    }

    .status-badge {
      padding: 4px 8px;
      border-radius: 4px;
      font-size: 0.625rem;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 0.5px;
      white-space: nowrap;
    }

    .initiative-meta {
      display: flex;
      gap: 16px;
      font-size: 0.75rem;
      opacity: 0.7;
      margin-top: 8px;
    }

    .meta-item {
      display: flex;
      gap: 4px;
      align-items: center;
    }

    .initiative-description {
      font-size: 0.875rem;
      margin: 12px 0 0;
      line-height: 1.5;
      opacity: 0.85;
    }

    .advances {
      margin-top: 12px;
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
    }

    .advance-chip {
      background: var(--it-card-bg);
      padding: 4px 8px;
      border-radius: 4px;
      font-size: 0.75rem;
    }

    .advance-arrow {
      opacity: 0.5;
      margin: 0 4px;
    }

    .empty-state {
      text-align: center;
      padding: 48px;
      opacity: 0.6;
    }
  `;

  @property({ type: String, reflect: true })
  theme: 'light' | 'dark' = 'light';

  @property({ type: String })
  src?: string;

  @property({ type: String, attribute: 'group-by' })
  groupBy: GroupBy = 'status';

  @property({ type: Boolean, attribute: 'show-stats' })
  showStats = true;

  @property({ type: Boolean, attribute: 'show-search' })
  showSearch = true;

  @state()
  private data: JourneyRoadmap | null = null;

  @state()
  private searchQuery = '';

  @state()
  private sortBy: SortBy = 'status';

  @state()
  private statusFilters: Set<InitiativeStatus> = new Set([
    'proposed',
    'planned',
    'in_progress',
  ]);

  @state()
  private expandedInitiative: string | null = null;

  override async connectedCallback() {
    super.connectedCallback();
    requestAnimationFrame(() => this.loadData());
  }

  protected override updated(changedProperties: PropertyValues) {
    if (changedProperties.has('src')) {
      this.loadData();
    }
  }

  private async loadData() {
    if (this.src) {
      try {
        const response = await fetch(this.src);
        this.data = await response.json();
        return;
      } catch (e) {
        console.error('Failed to load data from src:', e);
      }
    }

    const script = this.querySelector('script[type="application/json"]');
    if (script?.textContent) {
      try {
        this.data = JSON.parse(script.textContent);
      } catch (e) {
        console.error('Failed to parse inline JSON:', e);
      }
    }
  }

  private getFilteredInitiatives(): Initiative[] {
    if (!this.data?.initiatives) return [];

    let initiatives = [...this.data.initiatives];

    // Apply status filter
    initiatives = initiatives.filter(
      (i) => i.status && this.statusFilters.has(i.status)
    );

    // Apply search filter
    if (this.searchQuery) {
      const query = this.searchQuery.toLowerCase();
      initiatives = initiatives.filter(
        (i) =>
          i.name.toLowerCase().includes(query) ||
          i.description?.toLowerCase().includes(query) ||
          i.ownerTeam?.toLowerCase().includes(query)
      );
    }

    // Sort
    initiatives.sort((a, b) => {
      switch (this.sortBy) {
        case 'name':
          return a.name.localeCompare(b.name);
        case 'team':
          return (a.ownerTeam || '').localeCompare(b.ownerTeam || '');
        case 'status':
        default:
          const aOrder = STATUS_ORDER[a.status || 'proposed'];
          const bOrder = STATUS_ORDER[b.status || 'proposed'];
          return aOrder - bOrder;
      }
    });

    return initiatives;
  }

  private getGroupedInitiatives(): Map<string, Initiative[]> {
    const initiatives = this.getFilteredInitiatives();
    const groups = new Map<string, Initiative[]>();

    for (const initiative of initiatives) {
      let groupKey: string;

      switch (this.groupBy) {
        case 'team':
          groupKey = initiative.ownerTeam || 'Unassigned';
          break;
        case 'period':
          groupKey = initiative.periods?.[0] || 'No Period';
          break;
        case 'status':
        default:
          groupKey = initiative.status || 'proposed';
          break;
      }

      if (!groups.has(groupKey)) {
        groups.set(groupKey, []);
      }
      groups.get(groupKey)!.push(initiative);
    }

    return groups;
  }

  private getTeamName(teamId: string): string {
    const team = this.data?.teams?.find((t: Team) => t.id === teamId);
    return team?.name || teamId;
  }

  private formatStatus(status: string): string {
    return status
      .split('_')
      .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ');
  }

  private getStatusStyle(status: InitiativeStatus): string {
    const colors = STATUS_COLORS[status];
    return `background-color: ${colors.bg}; color: ${colors.text}`;
  }

  private toggleStatusFilter(status: InitiativeStatus) {
    const newFilters = new Set(this.statusFilters);
    if (newFilters.has(status)) {
      newFilters.delete(status);
    } else {
      newFilters.add(status);
    }
    this.statusFilters = newFilters;
  }

  private handleInitiativeClick(initiativeId: string) {
    this.expandedInitiative =
      this.expandedInitiative === initiativeId ? null : initiativeId;
    this.dispatchEvent(
      new CustomEvent('initiative-select', {
        detail: {
          initiativeId,
          expanded: this.expandedInitiative === initiativeId,
        },
      })
    );
  }

  private renderStats() {
    if (!this.showStats || !this.data?.initiatives) return null;

    const total = this.data.initiatives.length;
    const byStatus = this.data.initiatives.reduce(
      (acc, i) => {
        const status = i.status || 'proposed';
        acc[status] = (acc[status] || 0) + 1;
        return acc;
      },
      {} as Record<string, number>
    );

    return html`
      <div class="stats">
        <div class="stat-card">
          <div class="stat-value">${total}</div>
          <div class="stat-label">Total</div>
        </div>
        <div class="stat-card">
          <div class="stat-value">${byStatus['in_progress'] || 0}</div>
          <div class="stat-label">In Progress</div>
        </div>
        <div class="stat-card">
          <div class="stat-value">${byStatus['completed'] || 0}</div>
          <div class="stat-label">Completed</div>
        </div>
        <div class="stat-card">
          <div class="stat-value">${byStatus['planned'] || 0}</div>
          <div class="stat-label">Planned</div>
        </div>
      </div>
    `;
  }

  private renderControls() {
    const statuses: InitiativeStatus[] = [
      'proposed',
      'planned',
      'in_progress',
      'completed',
      'on_hold',
      'cancelled',
    ];

    return html`
      <div class="controls">
        <div class="control-group">
          <span class="control-label">Group by:</span>
          <button
            class="btn ${this.groupBy === 'status' ? 'active' : ''}"
            @click=${() => (this.groupBy = 'status')}
          >
            Status
          </button>
          <button
            class="btn ${this.groupBy === 'team' ? 'active' : ''}"
            @click=${() => (this.groupBy = 'team')}
          >
            Team
          </button>
          <button
            class="btn ${this.groupBy === 'period' ? 'active' : ''}"
            @click=${() => (this.groupBy = 'period')}
          >
            Period
          </button>
        </div>

        <div class="control-group">
          <span class="control-label">Status:</span>
          ${statuses.map(
            (status) => html`
              <button
                class="btn ${this.statusFilters.has(status) ? 'active' : ''}"
                @click=${() => this.toggleStatusFilter(status)}
              >
                ${this.formatStatus(status)}
              </button>
            `
          )}
        </div>

        ${this.showSearch
          ? html`
              <input
                type="text"
                class="search-input"
                placeholder="Search initiatives..."
                .value=${this.searchQuery}
                @input=${(e: Event) =>
                  (this.searchQuery = (e.target as HTMLInputElement).value)}
              />
            `
          : null}
      </div>
    `;
  }

  private renderInitiativeCard(initiative: Initiative) {
    const isExpanded = this.expandedInitiative === initiative.id;

    return html`
      <div
        class="initiative-card"
        @click=${() => this.handleInitiativeClick(initiative.id)}
      >
        <div class="initiative-header">
          <h4 class="initiative-name">${initiative.name}</h4>
          ${initiative.status
            ? html`
                <span
                  class="status-badge"
                  style=${this.getStatusStyle(initiative.status)}
                >
                  ${this.formatStatus(initiative.status)}
                </span>
              `
            : null}
        </div>

        <div class="initiative-meta">
          ${initiative.ownerTeam
            ? html`<span class="meta-item">👥 ${this.getTeamName(initiative.ownerTeam)}</span>`
            : null}
          ${initiative.periods && initiative.periods.length > 0
            ? html`<span class="meta-item">📅 ${initiative.periods.join(', ')}</span>`
            : null}
          ${initiative.advances && initiative.advances.length > 0
            ? html`<span class="meta-item">📈 ${initiative.advances.length} capabilities</span>`
            : null}
        </div>

        ${isExpanded && initiative.description
          ? html`<p class="initiative-description">${initiative.description}</p>`
          : null}

        ${isExpanded && initiative.advances && initiative.advances.length > 0
          ? html`
              <div class="advances">
                ${initiative.advances.map(
                  (adv) => html`
                    <span class="advance-chip">
                      ${adv.capabilityName || adv.capabilityId}
                      <span class="advance-arrow">→</span>
                      ${adv.to}
                    </span>
                  `
                )}
              </div>
            `
          : null}
      </div>
    `;
  }

  private renderGroup(groupKey: string, initiatives: Initiative[]) {
    let groupTitle = groupKey;

    if (this.groupBy === 'status') {
      groupTitle = this.formatStatus(groupKey);
    } else if (this.groupBy === 'team') {
      groupTitle = this.getTeamName(groupKey);
    }

    return html`
      <div class="group">
        <div class="group-header">
          <h3 class="group-title">${groupTitle}</h3>
          <span class="group-count">${initiatives.length}</span>
        </div>
        <div class="initiatives">
          ${initiatives.map((i) => this.renderInitiativeCard(i))}
        </div>
      </div>
    `;
  }

  override render() {
    if (!this.data) {
      return html`<div class="container empty-state">Loading...</div>`;
    }

    const groups = this.getGroupedInitiatives();

    if (groups.size === 0) {
      return html`
        <div class="container">
          <div class="header">
            <h1 class="title">Initiatives</h1>
          </div>
          ${this.renderControls()}
          <div class="empty-state">No initiatives match the current filters</div>
        </div>
      `;
    }

    return html`
      <div class="container">
        <div class="header">
          <h1 class="title">Initiatives</h1>
          <p class="subtitle">
            ${this.data.initiatives?.length || 0} initiatives across ${this.data.teams?.length || 0} teams
          </p>
        </div>

        ${this.renderStats()} ${this.renderControls()}

        <div class="groups">
          ${Array.from(groups.entries()).map(([key, initiatives]) =>
            this.renderGroup(key, initiatives)
          )}
        </div>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'initiative-tracker': InitiativeTracker;
  }
}
