import { LitElement, html, css } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';

export interface FilterState {
  pillar: string | null;
  status: string | null;
  domain: string | null;
}

/**
 * DashboardFilters provides interactive filtering for the portfolio dashboard.
 * Emits 'filter-change' events when filters are updated.
 *
 * @fires filter-change - Fired when any filter changes, with detail: FilterState
 *
 * @example
 * ```html
 * <prism-dashboard-filters
 *   pillars='["CSAT", "SAM-SOM", "TAM"]'
 *   statuses='["active", "planned", "proposed"]'
 *   domains='["security", "reliability", "platform"]'>
 * </prism-dashboard-filters>
 * ```
 */
@customElement('prism-dashboard-filters')
export class DashboardFilters extends LitElement {
  static override styles = css`
    :host {
      --df-bg: var(--bg-secondary, #f8fafc);
      --df-text: var(--text-primary, #1f2937);
      --df-text-muted: var(--text-secondary, #6b7280);
      --df-border: var(--border-color, #e5e7eb);
      --df-accent: var(--accent-primary, #3b82f6);
      --df-accent-hover: var(--accent-hover, #2563eb);
      display: block;
    }

    :host([theme='dark']) {
      --df-bg: #1e293b;
      --df-text: #f1f5f9;
      --df-text-muted: #94a3b8;
      --df-border: #334155;
    }

    .filters {
      display: flex;
      flex-wrap: wrap;
      gap: 16px;
      padding: 16px;
      background: var(--df-bg);
      border: 1px solid var(--df-border);
      border-radius: 12px;
    }

    .filter-group {
      display: flex;
      flex-direction: column;
      gap: 6px;
    }

    .filter-label {
      font-size: 0.75rem;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 0.05em;
      color: var(--df-text-muted);
    }

    .filter-options {
      display: flex;
      gap: 8px;
      flex-wrap: wrap;
    }

    .filter-btn {
      padding: 6px 12px;
      border: 1px solid var(--df-border);
      border-radius: 6px;
      background: transparent;
      color: var(--df-text);
      font-size: 0.875rem;
      cursor: pointer;
      transition: all 0.15s ease;
    }

    .filter-btn:hover {
      border-color: var(--df-accent);
      color: var(--df-accent);
    }

    .filter-btn.active {
      background: var(--df-accent);
      border-color: var(--df-accent);
      color: white;
    }

    .filter-btn.csat.active { background: #10b981; border-color: #10b981; }
    .filter-btn.sam-som.active { background: #3b82f6; border-color: #3b82f6; }
    .filter-btn.tam.active { background: #8b5cf6; border-color: #8b5cf6; }

    .clear-btn {
      margin-left: auto;
      padding: 6px 12px;
      border: none;
      border-radius: 6px;
      background: transparent;
      color: var(--df-text-muted);
      font-size: 0.875rem;
      cursor: pointer;
      transition: color 0.15s ease;
    }

    .clear-btn:hover {
      color: var(--df-text);
    }
  `;

  @property({ type: Array }) pillars: string[] = ['CSAT', 'SAM-SOM', 'TAM'];
  @property({ type: Array }) statuses: string[] = ['active', 'planned', 'proposed'];
  @property({ type: Array }) domains: string[] = [];
  @property({ type: String, reflect: true }) theme: 'light' | 'dark' = 'light';

  @state() private selectedPillar: string | null = null;
  @state() private selectedStatus: string | null = null;
  @state() private selectedDomain: string | null = null;

  private emitFilterChange() {
    const detail: FilterState = {
      pillar: this.selectedPillar,
      status: this.selectedStatus,
      domain: this.selectedDomain,
    };
    this.dispatchEvent(new CustomEvent('filter-change', { detail, bubbles: true, composed: true }));
  }

  private togglePillar(pillar: string) {
    this.selectedPillar = this.selectedPillar === pillar ? null : pillar;
    this.emitFilterChange();
  }

  private toggleStatus(status: string) {
    this.selectedStatus = this.selectedStatus === status ? null : status;
    this.emitFilterChange();
  }

  private toggleDomain(domain: string) {
    this.selectedDomain = this.selectedDomain === domain ? null : domain;
    this.emitFilterChange();
  }

  private clearFilters() {
    this.selectedPillar = null;
    this.selectedStatus = null;
    this.selectedDomain = null;
    this.emitFilterChange();
  }

  private hasActiveFilters(): boolean {
    return this.selectedPillar !== null || this.selectedStatus !== null || this.selectedDomain !== null;
  }

  override render() {
    return html`
      <div class="filters">
        <div class="filter-group">
          <span class="filter-label">Pillar</span>
          <div class="filter-options">
            ${this.pillars.map(pillar => html`
              <button
                class="filter-btn ${pillar.toLowerCase().replace(' ', '-')} ${this.selectedPillar === pillar ? 'active' : ''}"
                @click=${() => this.togglePillar(pillar)}
              >
                ${pillar}
              </button>
            `)}
          </div>
        </div>

        <div class="filter-group">
          <span class="filter-label">Status</span>
          <div class="filter-options">
            ${this.statuses.map(status => html`
              <button
                class="filter-btn ${this.selectedStatus === status ? 'active' : ''}"
                @click=${() => this.toggleStatus(status)}
              >
                ${status}
              </button>
            `)}
          </div>
        </div>

        ${this.domains.length > 0 ? html`
          <div class="filter-group">
            <span class="filter-label">Domain</span>
            <div class="filter-options">
              ${this.domains.map(domain => html`
                <button
                  class="filter-btn ${this.selectedDomain === domain ? 'active' : ''}"
                  @click=${() => this.toggleDomain(domain)}
                >
                  ${domain}
                </button>
              `)}
            </div>
          </div>
        ` : ''}

        ${this.hasActiveFilters() ? html`
          <button class="clear-btn" @click=${this.clearFilters}>Clear filters</button>
        ` : ''}
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'prism-dashboard-filters': DashboardFilters;
  }
}
