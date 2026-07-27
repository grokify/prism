import { describe, it, expect, afterEach } from 'vitest';

// Import components to register custom elements
import './maturity-grid.js';
import './roadmap-timeline.js';
import './initiative-tracker.js';

describe('Lit Components', () => {
  afterEach(() => {
    // Clean up any elements created during tests
    document.body.innerHTML = '';
  });

  describe('MaturityGrid', () => {
    it('is registered as a custom element', () => {
      expect(customElements.get('maturity-grid')).toBeDefined();
    });

    it('can be instantiated', () => {
      const el = document.createElement('maturity-grid');
      expect(el).toBeInstanceOf(HTMLElement);
      expect(el.tagName.toLowerCase()).toBe('maturity-grid');
    });

    it('has default property values', () => {
      const el = document.createElement('maturity-grid') as any;
      expect(el.view).toBe('by-layer');
      expect(el.theme).toBe('light');
      expect(el.showLegend).toBe(true);
      expect(el.showViewToggle).toBe(true);
    });

    it('reflects theme attribute', () => {
      const el = document.createElement('maturity-grid');
      el.setAttribute('theme', 'dark');
      expect(el.getAttribute('theme')).toBe('dark');
    });

    it('renders shadow DOM', async () => {
      const el = document.createElement('maturity-grid');
      document.body.appendChild(el);
      await el.updateComplete;
      expect(el.shadowRoot).not.toBeNull();
      expect(el.shadowRoot?.querySelector('.container')).not.toBeNull();
    });
  });

  describe('RoadmapTimeline', () => {
    it('is registered as a custom element', () => {
      expect(customElements.get('roadmap-timeline')).toBeDefined();
    });

    it('can be instantiated', () => {
      const el = document.createElement('roadmap-timeline');
      expect(el).toBeInstanceOf(HTMLElement);
      expect(el.tagName.toLowerCase()).toBe('roadmap-timeline');
    });

    it('has default property values', () => {
      const el = document.createElement('roadmap-timeline') as any;
      expect(el.theme).toBe('light');
      expect(el.showConfidence).toBe(true);
      expect(el.showCommitment).toBe(true);
      expect(el.showLegend).toBe(true);
      expect(el.showControls).toBe(true);
    });

    it('reflects theme attribute', () => {
      const el = document.createElement('roadmap-timeline');
      el.setAttribute('theme', 'dark');
      expect(el.getAttribute('theme')).toBe('dark');
    });

    it('renders shadow DOM', async () => {
      const el = document.createElement('roadmap-timeline');
      document.body.appendChild(el);
      await el.updateComplete;
      expect(el.shadowRoot).not.toBeNull();
      expect(el.shadowRoot?.querySelector('.container')).not.toBeNull();
    });

    it('dispatches capability-select event on click', async () => {
      const roadmapData = {
        id: 'test',
        name: 'Test Roadmap',
        timeModel: {
          type: 'quarterly',
          periods: [{ id: 'q1', label: 'Q1', isCurrent: true }],
        },
        capabilityJourneys: [
          {
            id: 'cj-1',
            capabilityId: 'cap-1',
            name: 'Test Cap',
            targetStates: [{ periodId: 'q1', maturityLevel: 'M2' }],
          },
        ],
      };

      const el = document.createElement('roadmap-timeline') as any;
      const script = document.createElement('script');
      script.type = 'application/json';
      script.textContent = JSON.stringify(roadmapData);
      el.appendChild(script);
      document.body.appendChild(el);

      // Wait for component to load data
      await new Promise((r) => setTimeout(r, 50));
      await el.updateComplete;

      let eventFired = false;
      el.addEventListener('capability-select', () => {
        eventFired = true;
      });

      // Simulate click on capability row
      const row = el.shadowRoot?.querySelector('tr:not(:first-child)');
      if (row) {
        row.click();
        expect(eventFired).toBe(true);
      }
    });
  });

  describe('InitiativeTracker', () => {
    it('is registered as a custom element', () => {
      expect(customElements.get('initiative-tracker')).toBeDefined();
    });

    it('can be instantiated', () => {
      const el = document.createElement('initiative-tracker');
      expect(el).toBeInstanceOf(HTMLElement);
      expect(el.tagName.toLowerCase()).toBe('initiative-tracker');
    });

    it('has default property values', () => {
      const el = document.createElement('initiative-tracker') as any;
      expect(el.theme).toBe('light');
      expect(el.groupBy).toBe('status');
      expect(el.showSearch).toBe(true);
      expect(el.showStats).toBe(true);
    });

    it('reflects theme attribute', () => {
      const el = document.createElement('initiative-tracker');
      el.setAttribute('theme', 'dark');
      expect(el.getAttribute('theme')).toBe('dark');
    });

    it('renders shadow DOM', async () => {
      const el = document.createElement('initiative-tracker');
      document.body.appendChild(el);
      await el.updateComplete;
      expect(el.shadowRoot).not.toBeNull();
      expect(el.shadowRoot?.querySelector('.container')).not.toBeNull();
    });

    it('supports group-by attribute', () => {
      const el = document.createElement('initiative-tracker') as any;
      el.setAttribute('group-by', 'team');
      expect(el.getAttribute('group-by')).toBe('team');
    });
  });

  describe('Component exports', () => {
    it('exports all components from index', async () => {
      const { MaturityGrid, RoadmapTimeline, InitiativeTracker } = await import('./index.js');
      expect(MaturityGrid).toBeDefined();
      expect(RoadmapTimeline).toBeDefined();
      expect(InitiativeTracker).toBeDefined();
    });
  });
});
