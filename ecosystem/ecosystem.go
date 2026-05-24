// Package ecosystem provides unified loading and querying across PRISM modules.
package ecosystem

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	capability "github.com/grokify/prism-capability"
	maturity "github.com/grokify/prism-maturity"
	"github.com/grokify/prism-roadmap/goals/okr"
	"github.com/grokify/prism-roadmap/roadmap"
)

// Config defines the ecosystem configuration.
type Config struct {
	Name string `json:"name" yaml:"name"`

	Capability CapabilityConfig `json:"capability" yaml:"capability"`
	Maturity   MaturityConfig   `json:"maturity" yaml:"maturity"`
	Roadmap    RoadmapConfig    `json:"roadmap" yaml:"roadmap"`
}

// CapabilityConfig defines capability stack sources.
type CapabilityConfig struct {
	Files []string `json:"files" yaml:"files"`
}

// MaturityConfig defines maturity document sources.
type MaturityConfig struct {
	Files []string `json:"files" yaml:"files"`
}

// RoadmapConfig defines roadmap document sources.
type RoadmapConfig struct {
	OKRs     []string `json:"okrs" yaml:"okrs"`
	Roadmaps []string `json:"roadmaps" yaml:"roadmaps"`
}

// Ecosystem holds loaded documents from all PRISM modules.
type Ecosystem struct {
	Config Config

	// Loaded documents
	CapabilityStacks []*capability.CapabilityStack
	PRISMDocuments   []*maturity.PRISMDocument
	OKRSets          []*okr.OKRSet
	Roadmaps         []*roadmap.Roadmap
}

// Load creates an Ecosystem from a configuration.
func Load(config Config) (*Ecosystem, error) {
	eco := &Ecosystem{
		Config:           config,
		CapabilityStacks: make([]*capability.CapabilityStack, 0),
		PRISMDocuments:   make([]*maturity.PRISMDocument, 0),
		OKRSets:          make([]*okr.OKRSet, 0),
		Roadmaps:         make([]*roadmap.Roadmap, 0),
	}

	// Load capability stacks
	for _, file := range config.Capability.Files {
		stack, err := capability.LoadFromFile(file)
		if err != nil {
			return nil, fmt.Errorf("loading capability stack %s: %w", file, err)
		}
		eco.CapabilityStacks = append(eco.CapabilityStacks, stack)
	}

	// Load PRISM documents (maturity)
	for _, file := range config.Maturity.Files {
		doc, err := loadPRISMDocument(file)
		if err != nil {
			return nil, fmt.Errorf("loading PRISM document %s: %w", file, err)
		}
		eco.PRISMDocuments = append(eco.PRISMDocuments, doc)
	}

	// Load OKRs
	for _, file := range config.Roadmap.OKRs {
		okrSet, err := loadOKRSet(file)
		if err != nil {
			return nil, fmt.Errorf("loading OKR set %s: %w", file, err)
		}
		eco.OKRSets = append(eco.OKRSets, okrSet)
	}

	// Load Roadmaps
	for _, file := range config.Roadmap.Roadmaps {
		rm, err := loadRoadmap(file)
		if err != nil {
			return nil, fmt.Errorf("loading roadmap %s: %w", file, err)
		}
		eco.Roadmaps = append(eco.Roadmaps, rm)
	}

	return eco, nil
}

// LoadFromFile loads an Ecosystem from a JSON config file.
func LoadFromFile(path string) (*Ecosystem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	return Load(config)
}

// loadPRISMDocument loads a PRISMDocument from a JSON file.
func loadPRISMDocument(path string) (*maturity.PRISMDocument, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var doc maturity.PRISMDocument
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

// loadOKRSet loads an OKRSet from a JSON file.
func loadOKRSet(path string) (*okr.OKRSet, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var okrSet okr.OKRSet
	if err := json.Unmarshal(data, &okrSet); err != nil {
		return nil, err
	}
	return &okrSet, nil
}

// loadRoadmap loads a Roadmap from a JSON file.
func loadRoadmap(path string) (*roadmap.Roadmap, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var rm roadmap.Roadmap
	if err := json.Unmarshal(data, &rm); err != nil {
		return nil, err
	}
	return &rm, nil
}

// =============================================================================
// Capability Queries
// =============================================================================

// AllCapabilities returns all capabilities from all stacks.
func (e *Ecosystem) AllCapabilities() []capability.Capability {
	var caps []capability.Capability
	for _, stack := range e.CapabilityStacks {
		caps = append(caps, stack.AllCapabilities()...)
	}
	return caps
}

// GetCapabilityByID finds a capability by ID across all stacks.
func (e *Ecosystem) GetCapabilityByID(id string) *capability.Capability {
	for _, stack := range e.CapabilityStacks {
		if cap := stack.GetCapabilityByID(id); cap != nil {
			return cap
		}
	}
	return nil
}

// CapabilitiesByStatus returns capabilities matching a status across all stacks.
func (e *Ecosystem) CapabilitiesByStatus(status string) []capability.Capability {
	var caps []capability.Capability
	for _, stack := range e.CapabilityStacks {
		caps = append(caps, stack.CapabilitiesByStatus(status)...)
	}
	return caps
}

// CapabilitiesByDomain returns capabilities from stacks matching a domain.
func (e *Ecosystem) CapabilitiesByDomain(domain string) []capability.Capability {
	var caps []capability.Capability
	for _, stack := range e.CapabilityStacks {
		if stack.Metadata.Domain == domain {
			caps = append(caps, stack.AllCapabilities()...)
		}
	}
	return caps
}

// =============================================================================
// Maturity Queries
// =============================================================================

// AllMetrics returns all metrics from all PRISM documents.
func (e *Ecosystem) AllMetrics() []maturity.Metric {
	var metrics []maturity.Metric
	for _, doc := range e.PRISMDocuments {
		metrics = append(metrics, doc.Metrics...)
	}
	return metrics
}

// GetMetricByID finds a metric by ID across all PRISM documents.
func (e *Ecosystem) GetMetricByID(id string) *maturity.Metric {
	for _, doc := range e.PRISMDocuments {
		if m := doc.GetMetricByID(id); m != nil {
			return m
		}
	}
	return nil
}

// AllServices returns all services from all PRISM documents.
func (e *Ecosystem) AllServices() []maturity.Service {
	var services []maturity.Service
	for _, doc := range e.PRISMDocuments {
		services = append(services, doc.Services...)
	}
	return services
}

// GetServiceByID finds a service by ID across all PRISM documents.
func (e *Ecosystem) GetServiceByID(id string) *maturity.Service {
	for _, doc := range e.PRISMDocuments {
		if s := doc.GetServiceByID(id); s != nil {
			return s
		}
	}
	return nil
}

// AllInitiatives returns all initiatives from all PRISM documents.
func (e *Ecosystem) AllInitiatives() []maturity.Initiative {
	var initiatives []maturity.Initiative
	for _, doc := range e.PRISMDocuments {
		initiatives = append(initiatives, doc.Initiatives...)
	}
	return initiatives
}

// GetInitiativeByID finds an initiative by ID across all PRISM documents.
func (e *Ecosystem) GetInitiativeByID(id string) *maturity.Initiative {
	for _, doc := range e.PRISMDocuments {
		if init := doc.GetInitiativeByID(id); init != nil {
			return init
		}
	}
	return nil
}

// =============================================================================
// Roadmap Queries
// =============================================================================

// AllObjectives returns all objectives from all OKR sets.
func (e *Ecosystem) AllObjectives() []okr.Objective {
	var objectives []okr.Objective
	for _, okrSet := range e.OKRSets {
		objectives = append(objectives, okrSet.ToObjectives()...)
	}
	return objectives
}

// GetObjectiveByID finds an objective by ID across all OKR sets.
func (e *Ecosystem) GetObjectiveByID(id string) *okr.Objective {
	for _, obj := range e.AllObjectives() {
		if obj.ID == id {
			return &obj
		}
	}
	return nil
}

// AllPhases returns all roadmap phases from all roadmaps.
func (e *Ecosystem) AllPhases() []roadmap.Phase {
	var phases []roadmap.Phase
	for _, rm := range e.Roadmaps {
		phases = append(phases, rm.Phases...)
	}
	return phases
}

// GetPhaseByID finds a phase by ID across all roadmaps.
func (e *Ecosystem) GetPhaseByID(id string) *roadmap.Phase {
	for _, rm := range e.Roadmaps {
		for _, phase := range rm.Phases {
			if phase.ID == id {
				return &phase
			}
		}
	}
	return nil
}

// =============================================================================
// Cross-Module Queries
// =============================================================================

// CapabilityContext provides full context for a capability across all modules.
type CapabilityContext struct {
	Capability *capability.Capability
	Metrics    []maturity.Metric
}

// GetCapabilityContext returns full context for a capability ID.
// Currently links capabilities to metrics via the capability's PRISMRef.SLIIDs.
func (e *Ecosystem) GetCapabilityContext(capabilityID string) *CapabilityContext {
	cap := e.GetCapabilityByID(capabilityID)
	if cap == nil {
		return nil
	}

	ctx := &CapabilityContext{
		Capability: cap,
		Metrics:    make([]maturity.Metric, 0),
	}

	// Find metrics linked via PRISMRef.SLIIDs
	if cap.PRISMRef != nil {
		sliIDs := make(map[string]bool)
		for _, sliID := range cap.PRISMRef.SLIIDs {
			sliIDs[sliID] = true
		}
		for _, m := range e.AllMetrics() {
			if sliIDs[m.ID] {
				ctx.Metrics = append(ctx.Metrics, m)
			}
		}
	}

	return ctx
}

// =============================================================================
// Validation
// =============================================================================

// ValidationError represents a cross-module validation error.
type ValidationError struct {
	Module  string `json:"module"`
	Type    string `json:"type"`
	ID      string `json:"id"`
	Field   string `json:"field"`
	RefID   string `json:"refId,omitempty"`
	Message string `json:"message"`
}

// ValidationErrors is a collection of validation errors.
type ValidationErrors []ValidationError

// HasErrors returns true if there are validation errors.
func (ve ValidationErrors) HasErrors() bool {
	return len(ve) > 0
}

// Error implements the error interface.
func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return ""
	}
	return fmt.Sprintf("%d validation errors", len(ve))
}

// Validate validates all loaded documents and cross-references.
func (e *Ecosystem) Validate() ValidationErrors {
	var errs ValidationErrors

	// Validate each capability stack
	for i, stack := range e.CapabilityStacks {
		if stackErrs := stack.Validate(); stackErrs.HasErrors() {
			for _, err := range stackErrs {
				errs = append(errs, ValidationError{
					Module:  "capability",
					Type:    "stack",
					ID:      fmt.Sprintf("stack[%d]", i),
					Field:   err.Field,
					Message: err.Message,
				})
			}
		}
	}

	// Validate each PRISM document
	for i, doc := range e.PRISMDocuments {
		if docErrs := doc.Validate(); docErrs.HasErrors() {
			for _, err := range docErrs {
				errs = append(errs, ValidationError{
					Module:  "maturity",
					Type:    "document",
					ID:      fmt.Sprintf("doc[%d]", i),
					Field:   err.Field,
					Message: err.Message,
				})
			}
		}
	}

	// Cross-reference validation: capability → SLI references
	for _, cap := range e.AllCapabilities() {
		if cap.PRISMRef != nil {
			for _, sliID := range cap.PRISMRef.SLIIDs {
				if e.GetMetricByID(sliID) == nil {
					errs = append(errs, ValidationError{
						Module:  "capability",
						Type:    "capability",
						ID:      cap.ID,
						Field:   "prismRef.sliIds",
						RefID:   sliID,
						Message: "references non-existent metric/SLI",
					})
				}
			}
		}
	}

	return errs
}

// =============================================================================
// Statistics
// =============================================================================

// Stats returns summary statistics about the ecosystem.
type Stats struct {
	CapabilityStacks  int            `json:"capabilityStacks"`
	TotalCapabilities int            `json:"totalCapabilities"`
	PRISMDocuments    int            `json:"prismDocuments"`
	TotalMetrics      int            `json:"totalMetrics"`
	TotalServices     int            `json:"totalServices"`
	TotalInitiatives  int            `json:"totalInitiatives"`
	TotalOKRSets      int            `json:"totalOkrSets"`
	TotalObjectives   int            `json:"totalObjectives"`
	TotalRoadmaps     int            `json:"totalRoadmaps"`
	TotalPhases       int            `json:"totalPhases"`
	ByStatus          map[string]int `json:"byStatus"`
	ByDomain          map[string]int `json:"byDomain"`
}

// Stats returns ecosystem statistics.
func (e *Ecosystem) Stats() Stats {
	stats := Stats{
		CapabilityStacks:  len(e.CapabilityStacks),
		TotalCapabilities: 0,
		PRISMDocuments:    len(e.PRISMDocuments),
		TotalMetrics:      len(e.AllMetrics()),
		TotalServices:     len(e.AllServices()),
		TotalInitiatives:  len(e.AllInitiatives()),
		TotalOKRSets:      len(e.OKRSets),
		TotalObjectives:   len(e.AllObjectives()),
		TotalRoadmaps:     len(e.Roadmaps),
		TotalPhases:       len(e.AllPhases()),
		ByStatus:          make(map[string]int),
		ByDomain:          make(map[string]int),
	}

	for _, stack := range e.CapabilityStacks {
		caps := stack.AllCapabilities()
		stats.TotalCapabilities += len(caps)

		domain := stack.Metadata.Domain
		if domain == "" {
			domain = "unspecified"
		}
		stats.ByDomain[domain] += len(caps)

		for _, cap := range caps {
			status := cap.Status
			if status == "" {
				status = "unspecified"
			}
			stats.ByStatus[status]++
		}
	}

	return stats
}

// =============================================================================
// File Loading Helpers
// =============================================================================

// LoadFromDirectory loads an ecosystem from a directory structure.
// Expected structure:
//
//	ecosystem/
//	  capability/
//	    *.json
//	  maturity/
//	    *.json
//	  roadmap/
//	    okrs/*.json
//	    roadmaps/*.json
func LoadFromDirectory(dir string) (*Ecosystem, error) {
	config := Config{
		Name: filepath.Base(dir),
	}

	// Scan capability files
	capDir := filepath.Join(dir, "capability")
	if files, err := filepath.Glob(filepath.Join(capDir, "*.json")); err == nil {
		config.Capability.Files = files
	}

	// Scan maturity files
	matDir := filepath.Join(dir, "maturity")
	if files, err := filepath.Glob(filepath.Join(matDir, "*.json")); err == nil {
		config.Maturity.Files = files
	}

	// Scan roadmap files
	roadmapDir := filepath.Join(dir, "roadmap")
	if files, err := filepath.Glob(filepath.Join(roadmapDir, "okrs", "*.json")); err == nil {
		config.Roadmap.OKRs = files
	}
	if files, err := filepath.Glob(filepath.Join(roadmapDir, "roadmaps", "*.json")); err == nil {
		config.Roadmap.Roadmaps = files
	}

	return Load(config)
}
