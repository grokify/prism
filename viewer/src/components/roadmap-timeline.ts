import { LitElement, html, css, PropertyValues } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';
import type { JourneyRoadmap, CapabilityJourney, Period, TargetState } from '../schema/roadmap/index.js';

type ColorScheme = 'default' | 'heatmap' | 'monochrome';
type FilterMode = 'all' | 'current' | 'future';

interface PeriodState {
  maturityLevel: string;
  confidence?: number;
  commitment?: string;
  isCurrent?: boolean;
}

/**
 * RoadmapTimeline displays capability journeys over time periods as an interactive timeline.
 *
 * @example
 * ```html
 * <roadmap-timeline theme="dark" show-confidence>
 *   <script type="application/json">
 *     {"id": "roadmap-1", "name": "Platform Roadmap", ...}
 *   </script>
 * </roadmap-timeline>
 * ```
 */
@customElement('roadmap-timeline')
export class RoadmapTimeline extends LitElement {
  static override styles = css`
    :host {
      --rt-bg: #ffffff;
      --rt-text: #1f2937;
      --rt-border: #e5e7eb;
      --rt-header-bg: #f8fafc;
      --rt-row-hover: #f1f5f9;
      --rt-current-bg: rgba(59, 130, 246, 0.1);
      display: block;
      font-family: system-ui, -apple-system, sans-serif;
    }

    :host([theme='dark']) {
      --rt-bg: #0f172a;
      --rt-text: #f1f5f9;
      --rt-border: #334155;
      --rt-header-bg: #1e293b;
      --rt-row-hover: #1e293b;
      --rt-current-bg: rgba(59, 130, 246, 0.2);
    }

    .container {
      background: var(--rt-bg);
      color: var(--rt-text);
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
      letter-spacing: -0.025em;
    }

    .vision {
      font-size: 0.875rem;
      opacity: 0.7;
      margin: 0;
      line-height: 1.5;
    }

    .controls {
      display: flex;
      gap: 16px;
      margin-bottom: 20px;
      flex-wrap: wrap;
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
      border: 1px solid var(--rt-border);
      background: transparent;
      color: var(--rt-text);
      transition: all 0.15s ease;
    }

    .btn:hover {
      background: var(--rt-row-hover);
    }

    .btn.active {
      background: var(--rt-text);
      color: var(--rt-bg);
    }

    .timeline-container {
      overflow-x: auto;
    }

    .timeline-table {
      width: 100%;
      border-collapse: collapse;
      font-size: 0.875rem;
    }

    .timeline-table th,
    .timeline-table td {
      padding: 12px;
      text-align: center;
      border: 1px solid var(--rt-border);
    }

    .timeline-table th {
      background: var(--rt-header-bg);
      font-weight: 600;
      font-size: 0.75rem;
      text-transform: uppercase;
      letter-spacing: 0.05em;
    }

    .timeline-table th.current {
      background: var(--rt-current-bg);
    }

    .capability-cell {
      text-align: left;
      min-width: 150px;
    }

    .capability-name {
      font-weight: 600;
      margin-bottom: 4px;
    }

    .capability-owner {
      font-size: 0.75rem;
      opacity: 0.7;
    }

    .timeline-table tbody tr:hover {
      background: var(--rt-row-hover);
    }

    .period-cell {
      min-width: 100px;
      vertical-align: middle;
    }

    .period-cell.current {
      background: var(--rt-current-bg);
    }

    .period-cell.empty {
      opacity: 0.4;
    }

    .maturity-badge {
      display: inline-block;
      padding: 4px 8px;
      border-radius: 4px;
      font-size: 0.75rem;
      font-weight: 600;
      color: #ffffff;
    }

    /* Default maturity colors */
    .maturity-0 { background: #6b7280; }
    .maturity-1 { background: #ef4444; }
    .maturity-2 { background: #f59e0b; }
    .maturity-3 { background: #eab308; }
    .maturity-4 { background: #22c55e; }
    .maturity-5 { background: #10b981; }

    /* Heatmap colors */
    .maturity-heat-0 { background: #fee2e2; color: #991b1b; }
    .maturity-heat-1 { background: #fecaca; color: #991b1b; }
    .maturity-heat-2 { background: #fef08a; color: #854d0e; }
    .maturity-heat-3 { background: #d9f99d; color: #3f6212; }
    .maturity-heat-4 { background: #86efac; color: #166534; }
    .maturity-heat-5 { background: #4ade80; color: #166534; }

    .confidence {
      display: block;
      font-size: 0.625rem;
      margin-top: 4px;
      opacity: 0.8;
    }

    .confidence-high { color: #22c55e; }
    .confidence-medium { color: #f59e0b; }
    .confidence-low { color: #ef4444; }

    .commitment {
      margin-left: 4px;
    }

    .legend {
      margin-top: 20px;
      padding: 16px;
      background: var(--rt-header-bg);
      border-radius: 8px;
      display: flex;
      flex-wrap: wrap;
      gap: 16px;
      font-size: 0.75rem;
    }

    .legend-group {
      display: flex;
      gap: 8px;
      align-items: center;
    }

    .legend-item {
      display: flex;
      gap: 4px;
      align-items: center;
    }

    .legend-label {
      opacity: 0.7;
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

  @property({ type: Boolean, attribute: 'show-confidence' })
  showConfidence = true;

  @property({ type: Boolean, attribute: 'show-commitment' })
  showCommitment = true;

  @property({ type: Boolean, attribute: 'show-legend' })
  showLegend = true;

  @property({ type: Boolean, attribute: 'show-controls' })
  showControls = true;

  @state()
  private data: JourneyRoadmap | null = null;

  @state()
  private colorScheme: ColorScheme = 'default';

  @state()
  private filterMode: FilterMode = 'all';

  @state()
  private selectedCapability: string | null = null;

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

  private getFilteredPeriods(): Period[] {
    if (!this.data?.timeModel?.periods) return [];

    const periods = this.data.timeModel.periods;
    const currentIndex = periods.findIndex((p) => p.isCurrent);

    switch (this.filterMode) {
      case 'current':
        return currentIndex >= 0 ? [periods[currentIndex]] : [];
      case 'future':
        return currentIndex >= 0 ? periods.slice(currentIndex) : periods;
      default:
        return periods;
    }
  }

  private getStateForPeriod(
    journey: CapabilityJourney,
    periodId: string
  ): PeriodState | null {
    if (journey.currentState?.periodId === periodId) {
      return {
        maturityLevel: journey.currentState.maturityLevel,
        isCurrent: true,
      };
    }

    const target = journey.targetStates.find((t: TargetState) => t.periodId === periodId);
    if (target) {
      return {
        maturityLevel: target.maturityLevel,
        confidence: target.confidence,
        commitment: target.commitment,
      };
    }

    return null;
  }

  private parseMaturityLevel(level: string): number {
    const match = level.match(/M(\d+)/i);
    return match ? parseInt(match[1], 10) : 0;
  }

  private getMaturityClass(level: number): string {
    const clampedLevel = Math.min(level, 5);
    if (this.colorScheme === 'heatmap') {
      return `maturity-heat-${clampedLevel}`;
    }
    return `maturity-${clampedLevel}`;
  }

  private getConfidenceClass(confidence: number): string {
    if (confidence >= 0.8) return 'confidence-high';
    if (confidence >= 0.5) return 'confidence-medium';
    return 'confidence-low';
  }

  private getCommitmentIcon(commitment: string): string {
    const icons: Record<string, string> = {
      committed: '🎯',
      planned: '📋',
      targeted: '🎯',
      aspirant: '⭐',
    };
    return icons[commitment] || '';
  }

  private handleCapabilityClick(capId: string) {
    this.selectedCapability = this.selectedCapability === capId ? null : capId;
    this.dispatchEvent(
      new CustomEvent('capability-select', {
        detail: { capabilityId: capId, selected: this.selectedCapability === capId },
      })
    );
  }

  private renderControls() {
    if (!this.showControls) return null;

    return html`
      <div class="controls">
        <div class="control-group">
          <span class="control-label">View:</span>
          <button
            class="btn ${this.filterMode === 'all' ? 'active' : ''}"
            @click=${() => (this.filterMode = 'all')}
          >
            All
          </button>
          <button
            class="btn ${this.filterMode === 'current' ? 'active' : ''}"
            @click=${() => (this.filterMode = 'current')}
          >
            Current
          </button>
          <button
            class="btn ${this.filterMode === 'future' ? 'active' : ''}"
            @click=${() => (this.filterMode = 'future')}
          >
            Future
          </button>
        </div>

        <div class="control-group">
          <span class="control-label">Colors:</span>
          <button
            class="btn ${this.colorScheme === 'default' ? 'active' : ''}"
            @click=${() => (this.colorScheme = 'default')}
          >
            Default
          </button>
          <button
            class="btn ${this.colorScheme === 'heatmap' ? 'active' : ''}"
            @click=${() => (this.colorScheme = 'heatmap')}
          >
            Heatmap
          </button>
        </div>
      </div>
    `;
  }

  private renderLegend() {
    if (!this.showLegend) return null;

    const levels = [
      { code: 'M0', label: 'Ad-hoc' },
      { code: 'M1', label: 'Developing' },
      { code: 'M2', label: 'Defined' },
      { code: 'M3', label: 'Managed' },
      { code: 'M4', label: 'Optimizing' },
      { code: 'M5', label: 'Leading' },
    ];

    return html`
      <div class="legend">
        <div class="legend-group">
          <span class="control-label">Maturity:</span>
          ${levels.map(
            (l, i) => html`
              <span class="legend-item">
                <span class="maturity-badge ${this.getMaturityClass(i)}">${l.code}</span>
                <span class="legend-label">${l.label}</span>
              </span>
            `
          )}
        </div>

        ${this.showCommitment
          ? html`
              <div class="legend-group">
                <span class="control-label">Commitment:</span>
                <span class="legend-item">🎯 Committed</span>
                <span class="legend-item">📋 Planned</span>
                <span class="legend-item">⭐ Aspirant</span>
              </div>
            `
          : null}
      </div>
    `;
  }

  private renderPeriodCell(state: PeriodState | null, period: Period) {
    if (!state) {
      return html`<td class="period-cell empty">-</td>`;
    }

    const level = this.parseMaturityLevel(state.maturityLevel);
    const maturityClass = this.getMaturityClass(level);
    const currentClass = period.isCurrent ? 'current' : '';

    return html`
      <td class="period-cell ${currentClass}">
        <span class="maturity-badge ${maturityClass}">${state.maturityLevel}</span>
        ${this.showConfidence && state.confidence !== undefined
          ? html`
              <span class="confidence ${this.getConfidenceClass(state.confidence)}">
                ${Math.round(state.confidence * 100)}%
              </span>
            `
          : null}
        ${this.showCommitment && state.commitment
          ? html`<span class="commitment">${this.getCommitmentIcon(state.commitment)}</span>`
          : null}
      </td>
    `;
  }

  private renderCapabilityRow(journey: CapabilityJourney, periods: Period[]) {
    const isSelected = this.selectedCapability === journey.id;

    return html`
      <tr
        class="${isSelected ? 'selected' : ''}"
        @click=${() => this.handleCapabilityClick(journey.id)}
      >
        <td class="capability-cell">
          <div class="capability-name">${journey.name}</div>
          ${journey.owner ? html`<div class="capability-owner">${journey.owner}</div>` : null}
        </td>
        ${periods.map((period) => {
          const state = this.getStateForPeriod(journey, period.id);
          return this.renderPeriodCell(state, period);
        })}
      </tr>
    `;
  }

  override render() {
    if (!this.data) {
      return html`<div class="container empty-state">Loading...</div>`;
    }

    const periods = this.getFilteredPeriods();
    const journeys = this.data.capabilityJourneys || [];

    if (periods.length === 0 || journeys.length === 0) {
      return html`
        <div class="container empty-state">
          No capability journeys or time periods defined
        </div>
      `;
    }

    return html`
      <div class="container">
        <div class="header">
          ${this.data.name ? html`<h1 class="title">${this.data.name}</h1>` : null}
          ${this.data.vision ? html`<p class="vision">${this.data.vision}</p>` : null}
        </div>

        ${this.renderControls()}

        <div class="timeline-container">
          <table class="timeline-table">
            <thead>
              <tr>
                <th class="capability-cell">Capability</th>
                ${periods.map(
                  (p) => html`<th class="${p.isCurrent ? 'current' : ''}">${p.label}</th>`
                )}
              </tr>
            </thead>
            <tbody>
              ${journeys.map((j) => this.renderCapabilityRow(j, periods))}
            </tbody>
          </table>
        </div>

        ${this.renderLegend()}
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'roadmap-timeline': RoadmapTimeline;
  }
}
