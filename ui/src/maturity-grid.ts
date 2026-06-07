import { LitElement, html, css, PropertyValues } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';
import {
  MaturityGridData,
  Capability,
  Layer,
  Category,
  CapabilityStatus,
  STATUS_COLORS,
  STATUS_TEXT_COLORS,
  MATURITY_COLORS,
} from './types.js';

type ViewMode = 'by-layer' | 'by-category';

interface CapabilityGroup {
  id: string;
  name: string;
  order: number;
  capabilities: Capability[];
}

/**
 * MaturityGrid displays capabilities organized by layer or category,
 * with status colors and optional maturity level badges.
 *
 * @example
 * ```html
 * <maturity-grid view="by-layer" theme="dark">
 *   <script type="application/json">
 *     {"layers": [...], "capabilities": [...]}
 *   </script>
 * </maturity-grid>
 * ```
 */
@customElement('maturity-grid')
export class MaturityGrid extends LitElement {
  static override styles = css`
    :host {
      --mg-bg: #ffffff;
      --mg-text: #1f2937;
      --mg-border: #e5e7eb;
      --mg-layer-bg: #f8fafc;
      --mg-inactive-bg: #f3f4f6;
      --mg-inactive-text: #9ca3af;
      display: block;
      font-family: system-ui, -apple-system, sans-serif;
    }

    :host([theme='dark']) {
      --mg-bg: #0f172a;
      --mg-text: #f1f5f9;
      --mg-border: #334155;
      --mg-layer-bg: #1e293b;
      --mg-inactive-bg: #334155;
      --mg-inactive-text: #94a3b8;
    }

    .container {
      background: var(--mg-bg);
      color: var(--mg-text);
      padding: 24px;
      min-height: 100%;
      box-sizing: border-box;
    }

    .title {
      font-size: 1.75rem;
      font-weight: 700;
      margin: 0 0 24px 0;
      text-align: center;
      letter-spacing: -0.025em;
    }

    .filters {
      background: var(--mg-layer-bg);
      border: 1px solid var(--mg-border);
      border-radius: 12px;
      padding: 20px;
      margin-bottom: 24px;
    }

    .filter-group {
      margin-bottom: 16px;
    }

    .filter-group:last-of-type {
      margin-bottom: 0;
    }

    .filter-label {
      font-size: 0.75rem;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 0.05em;
      opacity: 0.7;
      margin-bottom: 10px;
    }

    .filter-options {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
    }

    .filter-btn {
      display: flex;
      align-items: center;
      gap: 8px;
      padding: 8px 14px;
      border-radius: 8px;
      font-size: 0.8125rem;
      font-weight: 500;
      cursor: pointer;
      transition: all 0.15s ease;
      border: 1px solid var(--mg-border);
      background: transparent;
      color: var(--mg-text);
      user-select: none;
    }

    .filter-btn:hover {
      background: rgba(255, 255, 255, 0.05);
    }

    .filter-btn.inactive {
      opacity: 0.4;
    }

    .filter-color {
      width: 14px;
      height: 14px;
      border-radius: 4px;
      flex-shrink: 0;
    }

    .filter-actions {
      display: flex;
      gap: 8px;
      margin-top: 16px;
      padding-top: 16px;
      border-top: 1px solid var(--mg-border);
    }

    .btn {
      padding: 6px 12px;
      border-radius: 6px;
      font-size: 0.75rem;
      font-weight: 500;
      cursor: pointer;
      border: 1px solid var(--mg-border);
      background: transparent;
      color: var(--mg-text);
      transition: all 0.15s ease;
    }

    .btn:hover {
      background: rgba(255, 255, 255, 0.1);
    }

    .stack {
      display: flex;
      flex-direction: column;
      gap: 20px;
    }

    .layer {
      background: var(--mg-layer-bg);
      border: 1px solid var(--mg-border);
      border-radius: 12px;
      padding: 20px;
    }

    .layer-header {
      font-size: 0.875rem;
      font-weight: 600;
      margin-bottom: 12px;
      opacity: 0.8;
    }

    .capabilities {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
      gap: 10px;
    }

    .capability {
      padding: 14px 16px;
      border-radius: 8px;
      font-size: 0.875rem;
      font-weight: 500;
      text-align: center;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 6px;
      transition: all 0.2s ease;
      cursor: default;
    }

    .capability.filtered-out {
      background-color: var(--mg-inactive-bg) !important;
      color: var(--mg-inactive-text) !important;
      box-shadow: none;
      border: 1px solid var(--mg-border);
    }

    .capability.filtered-out .badge {
      background-color: transparent !important;
      color: var(--mg-inactive-text) !important;
      border: 1px solid var(--mg-border);
    }

    .cap-name {
      display: block;
      line-height: 1.3;
    }

    .badge {
      font-size: 0.625rem;
      font-weight: 700;
      padding: 3px 8px;
      border-radius: 4px;
      text-transform: uppercase;
      letter-spacing: 0.5px;
    }

    .view-toggle {
      display: flex;
      gap: 8px;
      margin-bottom: 16px;
      justify-content: center;
    }

    .view-btn {
      padding: 8px 16px;
      border-radius: 8px;
      font-size: 0.875rem;
      font-weight: 500;
      cursor: pointer;
      border: 1px solid var(--mg-border);
      background: transparent;
      color: var(--mg-text);
      transition: all 0.15s ease;
    }

    .view-btn.active {
      background: var(--mg-text);
      color: var(--mg-bg);
    }

    .view-btn:hover:not(.active) {
      background: rgba(255, 255, 255, 0.1);
    }
  `;

  @property({ type: String })
  view: ViewMode = 'by-layer';

  @property({ type: String, reflect: true })
  theme: 'light' | 'dark' = 'light';

  @property({ type: String })
  src?: string;

  @property({ type: Boolean, attribute: 'show-legend' })
  showLegend = true;

  @property({ type: Boolean, attribute: 'show-view-toggle' })
  showViewToggle = true;

  @state()
  private data: MaturityGridData | null = null;

  @state()
  private statusFilters: Set<CapabilityStatus> = new Set([
    'operational',
    'implemented',
    'in-progress',
    'planned',
    'deprecated',
  ]);

  @state()
  private maturityFilters: Set<number> = new Set([1, 2, 3, 4, 5]);

  override async connectedCallback() {
    super.connectedCallback();
    // Use requestAnimationFrame to ensure DOM children are parsed
    requestAnimationFrame(() => this.loadData());
  }

  protected override updated(changedProperties: PropertyValues) {
    if (changedProperties.has('src')) {
      this.loadData();
    }
  }

  private async loadData() {
    // Try loading from src attribute
    if (this.src) {
      try {
        const response = await fetch(this.src);
        this.data = await response.json();
        return;
      } catch (e) {
        console.error('Failed to load data from src:', e);
      }
    }

    // Try loading from inline script
    const script = this.querySelector('script[type="application/json"]');
    if (script?.textContent) {
      try {
        this.data = JSON.parse(script.textContent);
      } catch (e) {
        console.error('Failed to parse inline JSON:', e);
      }
    }
  }

  private getGroups(): CapabilityGroup[] {
    if (!this.data) return [];

    const groups: CapabilityGroup[] = [];
    const groupMap = new Map<string, Capability[]>();
    const orderMap = new Map<string, number>();
    const nameMap = new Map<string, string>();

    if (this.view === 'by-layer') {
      for (let i = 0; i < this.data.layers.length; i++) {
        const layer = this.data.layers[i];
        orderMap.set(layer.id, layer.order ?? i);
        nameMap.set(layer.id, layer.name);
      }

      for (const cap of this.data.capabilities) {
        const layerId = cap.layerId || 'other';
        if (!groupMap.has(layerId)) {
          groupMap.set(layerId, []);
        }
        groupMap.get(layerId)!.push(cap);
        if (!orderMap.has(layerId)) {
          orderMap.set(layerId, 100);
          nameMap.set(layerId, layerId);
        }
      }
    } else {
      for (let i = 0; i < this.data.categories.length; i++) {
        const cat = this.data.categories[i];
        orderMap.set(cat.id, cat.order ?? i);
        nameMap.set(cat.id, cat.name);
      }

      for (const cap of this.data.capabilities) {
        const catId = cap.categoryId || 'uncategorized';
        if (!groupMap.has(catId)) {
          groupMap.set(catId, []);
        }
        groupMap.get(catId)!.push(cap);
        if (!orderMap.has(catId)) {
          orderMap.set(catId, 100);
          nameMap.set(catId, catId.charAt(0).toUpperCase() + catId.slice(1));
        }
      }
    }

    for (const [id, caps] of groupMap) {
      groups.push({
        id,
        name: nameMap.get(id) || id,
        order: orderMap.get(id) || 100,
        capabilities: caps,
      });
    }

    groups.sort((a, b) => a.order - b.order);
    return groups;
  }

  private hasMaturityData(): boolean {
    return !!this.data?.maturity && Object.keys(this.data.maturity).length > 0;
  }

  private getMaturityLevel(capId: string): number | null {
    if (!this.data?.maturity?.[capId]) return null;
    return this.data.maturity[capId].level;
  }

  private isFiltered(cap: Capability): boolean {
    if (!this.statusFilters.has(cap.status)) {
      return true;
    }

    if (this.hasMaturityData()) {
      const level = this.getMaturityLevel(cap.id);
      if (level !== null && !this.maturityFilters.has(level)) {
        return true;
      }
    }

    return false;
  }

  private toggleStatusFilter(status: CapabilityStatus) {
    const newFilters = new Set(this.statusFilters);
    if (newFilters.has(status)) {
      newFilters.delete(status);
    } else {
      newFilters.add(status);
    }
    this.statusFilters = newFilters;
  }

  private toggleMaturityFilter(level: number) {
    const newFilters = new Set(this.maturityFilters);
    if (newFilters.has(level)) {
      newFilters.delete(level);
    } else {
      newFilters.add(level);
    }
    this.maturityFilters = newFilters;
  }

  private selectAll() {
    this.statusFilters = new Set([
      'operational',
      'implemented',
      'in-progress',
      'planned',
      'deprecated',
    ]);
    this.maturityFilters = new Set([1, 2, 3, 4, 5]);
  }

  private clearAll() {
    this.statusFilters = new Set();
    this.maturityFilters = new Set();
  }

  private setView(view: ViewMode) {
    this.view = view;
    this.dispatchEvent(new CustomEvent('view-change', { detail: { view } }));
  }

  private renderFilters() {
    if (!this.showLegend) return null;

    const statuses: CapabilityStatus[] = [
      'operational',
      'implemented',
      'in-progress',
      'planned',
      'deprecated',
    ];

    return html`
      <div class="filters">
        <div class="filter-group">
          <div class="filter-label">Status</div>
          <div class="filter-options">
            ${statuses.map(
              (status) => html`
                <button
                  class="filter-btn ${this.statusFilters.has(status)
                    ? ''
                    : 'inactive'}"
                  @click=${() => this.toggleStatusFilter(status)}
                >
                  <span
                    class="filter-color"
                    style="background-color: ${STATUS_COLORS[status]}"
                  ></span>
                  <span>${this.formatStatus(status)}</span>
                </button>
              `
            )}
          </div>
        </div>

        ${this.hasMaturityData()
          ? html`
              <div class="filter-group">
                <div class="filter-label">Maturity Level</div>
                <div class="filter-options">
                  ${[1, 2, 3, 4, 5].map(
                    (level) => html`
                      <button
                        class="filter-btn ${this.maturityFilters.has(level)
                          ? ''
                          : 'inactive'}"
                        @click=${() => this.toggleMaturityFilter(level)}
                      >
                        <span
                          class="filter-color"
                          style="background-color: ${MATURITY_COLORS[level]}"
                        ></span>
                        <span>M${level}</span>
                      </button>
                    `
                  )}
                </div>
              </div>
            `
          : null}

        <div class="filter-actions">
          <button class="btn" @click=${this.selectAll}>Select All</button>
          <button class="btn" @click=${this.clearAll}>Clear All</button>
        </div>
      </div>
    `;
  }

  private formatStatus(status: CapabilityStatus): string {
    return status
      .split('-')
      .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ');
  }

  private renderCapability(cap: Capability) {
    const filtered = this.isFiltered(cap);
    const bgColor = STATUS_COLORS[cap.status] || '#e5e7eb';
    const textColor = STATUS_TEXT_COLORS[cap.status] || '#000000';
    const maturityLevel = this.getMaturityLevel(cap.id);

    const tooltip = [
      cap.fullName,
      cap.description,
      cap.owner ? `Owner: ${cap.owner}` : null,
      `Status: ${cap.status}`,
      maturityLevel !== null ? `Maturity: M${maturityLevel}` : null,
    ]
      .filter(Boolean)
      .join(' | ');

    return html`
      <div
        class="capability ${filtered ? 'filtered-out' : ''}"
        style="background-color: ${bgColor}; color: ${textColor}"
        title=${tooltip}
      >
        <span class="cap-name">${cap.name}</span>
        ${maturityLevel !== null
          ? html`
              <span
                class="badge"
                style="background-color: ${MATURITY_COLORS[maturityLevel]}; color: #ffffff"
              >
                M${maturityLevel}
              </span>
            `
          : null}
      </div>
    `;
  }

  override render() {
    if (!this.data) {
      return html`<div class="container">Loading...</div>`;
    }

    const groups = this.getGroups();

    return html`
      <div class="container">
        ${this.data.title ? html`<h1 class="title">${this.data.title}</h1>` : null}

        ${this.showViewToggle
          ? html`
              <div class="view-toggle">
                <button
                  class="view-btn ${this.view === 'by-layer' ? 'active' : ''}"
                  @click=${() => this.setView('by-layer')}
                >
                  By Layer
                </button>
                <button
                  class="view-btn ${this.view === 'by-category' ? 'active' : ''}"
                  @click=${() => this.setView('by-category')}
                >
                  By Category
                </button>
              </div>
            `
          : null}

        ${this.renderFilters()}

        <div class="stack">
          ${groups.map(
            (group) => html`
              <div class="layer">
                <div class="layer-header">${group.name}</div>
                <div class="capabilities">
                  ${group.capabilities.map((cap) => this.renderCapability(cap))}
                </div>
              </div>
            `
          )}
        </div>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'maturity-grid': MaturityGrid;
  }
}
