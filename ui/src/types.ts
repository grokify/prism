/**
 * Types for PRISM UI components.
 * These mirror the Go structs from prism-capability and prism-maturity.
 */

export interface Capability {
  id: string;
  name: string;
  fullName?: string;
  description?: string;
  layerId: string;
  categoryId?: string;
  status: CapabilityStatus;
  priority?: string;
  importance?: string;
  order?: number;
  owner?: string;
  tooling?: Tool[];
  tags?: string[];
}

export type CapabilityStatus =
  | 'planned'
  | 'in-progress'
  | 'implemented'
  | 'operational'
  | 'deprecated';

export interface Tool {
  name: string;
  type?: string;
  status?: string;
  url?: string;
}

export interface Layer {
  id: string;
  name: string;
  description?: string;
  order?: number;
}

export interface Category {
  id: string;
  name: string;
  description?: string;
  order?: number;
}

export interface MaturityData {
  capabilityId: string;
  level: number;
  sliCount?: number;
}

export interface MaturityGridData {
  title?: string;
  layers: Layer[];
  categories: Category[];
  capabilities: Capability[];
  maturity?: Record<string, MaturityData>;
}

export interface SLIThreshold {
  level: number;
  value: number;
  valueStr: string;
  operator: string;
}

export interface SLIBulletData {
  id: string;
  title: string;
  subtitle?: string;
  measures: number[];
  markers?: number[];
  ranges: number[];
  thresholds: SLIThreshold[];
  currentValue: number;
  unit?: string;
}

export interface SLIGroup {
  id: string;
  name: string;
  order: number;
  slis: SLIBulletData[];
}

export interface MetricsTableRow {
  category: string;
  name: string;
  unit: string;
  tags: string;
  frameworks: string;
  currentLevel: number;
  currentValue: string;
  targetQ1: number;
  targetQ2: number;
  m1: string;
  m2: string;
  m3: string;
  m4: string;
  m5: string;
}

// Status colors matching Go render/html.go
export const STATUS_COLORS: Record<CapabilityStatus, string> = {
  operational: '#10b981',
  implemented: '#3b82f6',
  'in-progress': '#f59e0b',
  planned: '#9ca3af',
  deprecated: '#ef4444',
};

export const STATUS_TEXT_COLORS: Record<CapabilityStatus, string> = {
  operational: '#ffffff',
  implemented: '#ffffff',
  'in-progress': '#000000',
  planned: '#000000',
  deprecated: '#ffffff',
};

// Maturity level colors (M1-M5)
export const MATURITY_COLORS: Record<number, string> = {
  1: '#ef4444', // red
  2: '#f59e0b', // amber
  3: '#eab308', // yellow
  4: '#22c55e', // green
  5: '#3b82f6', // blue
};
