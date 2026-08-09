// Package ecosystem provides unified loading and querying across PRISM modules.
package ecosystem

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	capability "github.com/grokify/prism-capability"
	maturity "github.com/grokify/prism-maturity"
	"github.com/grokify/prism-roadmap/canvas"
	"github.com/grokify/prism-roadmap/goals/okr"
	"github.com/grokify/prism-roadmap/goals/v2mom"
	"github.com/grokify/prism-roadmap/roadmap"
)

// Config defines the ecosystem configuration.
type Config struct {
	Name string `json:"name" yaml:"name"`

	Capability CapabilityConfig `json:"capability" yaml:"capability"`
	Maturity   MaturityConfig   `json:"maturity" yaml:"maturity"`
	Roadmap    RoadmapConfig    `json:"roadmap" yaml:"roadmap"`
	Canvas     CanvasConfig     `json:"canvas" yaml:"canvas"`
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
	OKRs             []string `json:"okrs" yaml:"okrs"`
	V2MOMs           []string `json:"v2moms" yaml:"v2moms"`
	Roadmaps         []string `json:"roadmaps" yaml:"roadmaps"`
	OpportunitySpecs []string `json:"opportunitySpecs" yaml:"opportunitySpecs"`
}

// CanvasConfig defines strategic canvas document sources.
type CanvasConfig struct {
	BMCs []string `json:"bmcs" yaml:"bmcs"`
}

// Ecosystem holds loaded documents from all PRISM modules.
type Ecosystem struct {
	Config Config

	// Loaded documents
	CapabilityStacks []*capability.CapabilityStack
	PRISMDocuments   []*maturity.PRISMDocument
	OKRSets          []*okr.OKRSet
	V2MOMs           []*v2mom.V2MOM
	Roadmaps         []*roadmap.Roadmap
	OpportunitySpecs []*canvas.OpportunitySpec
	BMCs             []*canvas.BusinessModelCanvas

	// Evidence resolution (optional, nil if OmniSignal not available)
	EvidenceResolver EvidenceResolver

	// Evidence links by capability ID (loaded from external config or set programmatically)
	CapabilityEvidence map[string][]EvidenceLink
}

// Load creates an Ecosystem from a configuration.
func Load(config Config) (*Ecosystem, error) {
	eco := &Ecosystem{
		Config:           config,
		CapabilityStacks: make([]*capability.CapabilityStack, 0),
		PRISMDocuments:   make([]*maturity.PRISMDocument, 0),
		OKRSets:          make([]*okr.OKRSet, 0),
		V2MOMs:           make([]*v2mom.V2MOM, 0),
		Roadmaps:         make([]*roadmap.Roadmap, 0),
		OpportunitySpecs: make([]*canvas.OpportunitySpec, 0),
		BMCs:             make([]*canvas.BusinessModelCanvas, 0),
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

	// Load V2MOMs
	for _, file := range config.Roadmap.V2MOMs {
		v, err := v2mom.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("loading V2MOM %s: %w", file, err)
		}
		eco.V2MOMs = append(eco.V2MOMs, v)
	}

	// Load Roadmaps
	for _, file := range config.Roadmap.Roadmaps {
		rm, err := loadRoadmap(file)
		if err != nil {
			return nil, fmt.Errorf("loading roadmap %s: %w", file, err)
		}
		eco.Roadmaps = append(eco.Roadmaps, rm)
	}

	// Load OpportunitySpecs
	for _, file := range config.Roadmap.OpportunitySpecs {
		spec, err := loadOpportunitySpec(file)
		if err != nil {
			return nil, fmt.Errorf("loading opportunity spec %s: %w", file, err)
		}
		eco.OpportunitySpecs = append(eco.OpportunitySpecs, spec)
	}

	// Load BMCs
	for _, file := range config.Canvas.BMCs {
		bmc, err := loadBMC(file)
		if err != nil {
			return nil, fmt.Errorf("loading BMC %s: %w", file, err)
		}
		eco.BMCs = append(eco.BMCs, bmc)
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

// loadOpportunitySpec loads an OpportunitySpec from a JSON file.
func loadOpportunitySpec(path string) (*canvas.OpportunitySpec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var spec canvas.OpportunitySpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
}

// loadBMC loads a BusinessModelCanvas from a JSON file.
func loadBMC(path string) (*canvas.BusinessModelCanvas, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var bmc canvas.BusinessModelCanvas
	if err := json.Unmarshal(data, &bmc); err != nil {
		return nil, err
	}
	return &bmc, nil
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
// V2MOM Queries
// =============================================================================

// AllV2MOMs returns all loaded V2MOM documents.
func (e *Ecosystem) AllV2MOMs() []*v2mom.V2MOM {
	return e.V2MOMs
}

// GetV2MOMByID finds a V2MOM by its metadata ID.
func (e *Ecosystem) GetV2MOMByID(id string) *v2mom.V2MOM {
	for _, v := range e.V2MOMs {
		if v.Metadata != nil && v.Metadata.ID == id {
			return v
		}
	}
	return nil
}

// =============================================================================
// OpportunitySpec Queries
// =============================================================================

// AllOpportunitySpecs returns all loaded OpportunitySpec documents.
func (e *Ecosystem) AllOpportunitySpecs() []*canvas.OpportunitySpec {
	return e.OpportunitySpecs
}

// GetOpportunitySpecByID finds an OpportunitySpec by its metadata ID.
func (e *Ecosystem) GetOpportunitySpecByID(id string) *canvas.OpportunitySpec {
	for _, spec := range e.OpportunitySpecs {
		if spec.Metadata.ID == id {
			return spec
		}
	}
	return nil
}

// OpportunitySpecCapabilityRefs returns all capability IDs referenced by OpportunitySpecs.
// These come from UniqueCapabilities and MustHaveCapabilities fields.
func (e *Ecosystem) OpportunitySpecCapabilityRefs() []string {
	seen := make(map[string]bool)
	var refs []string
	for _, spec := range e.OpportunitySpecs {
		for _, capID := range spec.CompetitiveEdge.UniqueCapabilities {
			if !seen[capID] {
				seen[capID] = true
				refs = append(refs, capID)
			}
		}
		for _, capID := range spec.CriticalRequirements.MustHaveCapabilities {
			if !seen[capID] {
				seen[capID] = true
				refs = append(refs, capID)
			}
		}
	}
	return refs
}

// =============================================================================
// BMC Queries
// =============================================================================

// AllBMCs returns all loaded Business Model Canvas documents.
func (e *Ecosystem) AllBMCs() []*canvas.BusinessModelCanvas {
	return e.BMCs
}

// GetBMCByID finds a Business Model Canvas by its metadata ID.
func (e *Ecosystem) GetBMCByID(id string) *canvas.BusinessModelCanvas {
	for _, bmc := range e.BMCs {
		if bmc.Metadata.ID == id {
			return bmc
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

	// Evidence links and summary (RMI-013)
	EvidenceLinks   []EvidenceLink   `json:"evidenceLinks,omitempty"`
	EvidenceSummary *EvidenceSummary `json:"evidenceSummary,omitempty"`
}

// GetCapabilityContext returns full context for a capability ID.
// Links capabilities to metrics via PRISMRef.SLIIDs and resolves evidence if available.
func (e *Ecosystem) GetCapabilityContext(capabilityID string) *CapabilityContext {
	cap := e.GetCapabilityByID(capabilityID)
	if cap == nil {
		return nil
	}

	ctx := &CapabilityContext{
		Capability:    cap,
		Metrics:       make([]maturity.Metric, 0),
		EvidenceLinks: make([]EvidenceLink, 0),
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

	// Add evidence links if available
	if e.CapabilityEvidence != nil {
		if links, ok := e.CapabilityEvidence[capabilityID]; ok {
			ctx.EvidenceLinks = links

			// Resolve evidence if resolver is available
			if e.EvidenceResolver != nil {
				resolved, err := e.EvidenceResolver.ResolveAll(links)
				if err == nil {
					ctx.EvidenceSummary = ComputeEvidenceSummary(links, resolved)
				}
			}
		}
	}

	return ctx
}

// SetCapabilityEvidence sets evidence links for a capability.
func (e *Ecosystem) SetCapabilityEvidence(capabilityID string, links []EvidenceLink) {
	if e.CapabilityEvidence == nil {
		e.CapabilityEvidence = make(map[string][]EvidenceLink)
	}
	e.CapabilityEvidence[capabilityID] = links
}

// AddCapabilityEvidence adds evidence links to a capability.
func (e *Ecosystem) AddCapabilityEvidence(capabilityID string, links ...EvidenceLink) {
	if e.CapabilityEvidence == nil {
		e.CapabilityEvidence = make(map[string][]EvidenceLink)
	}
	e.CapabilityEvidence[capabilityID] = append(e.CapabilityEvidence[capabilityID], links...)
}

// AllCapabilityEvidence returns all capability evidence links.
func (e *Ecosystem) AllCapabilityEvidence() map[string][]EvidenceLink {
	if e.CapabilityEvidence == nil {
		return make(map[string][]EvidenceLink)
	}
	return e.CapabilityEvidence
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

	// Cross-reference validation: OpportunitySpec → capability references
	for _, spec := range e.OpportunitySpecs {
		specID := spec.Metadata.ID
		if specID == "" {
			specID = spec.Metadata.Title
		}

		// Validate UniqueCapabilities
		for _, capID := range spec.CompetitiveEdge.UniqueCapabilities {
			if e.GetCapabilityByID(capID) == nil {
				errs = append(errs, ValidationError{
					Module:  "canvas",
					Type:    "opportunitySpec",
					ID:      specID,
					Field:   "competitiveEdge.uniqueCapabilities",
					RefID:   capID,
					Message: "references non-existent capability",
				})
			}
		}

		// Validate MustHaveCapabilities
		for _, capID := range spec.CriticalRequirements.MustHaveCapabilities {
			if e.GetCapabilityByID(capID) == nil {
				errs = append(errs, ValidationError{
					Module:  "canvas",
					Type:    "opportunitySpec",
					ID:      specID,
					Field:   "criticalRequirements.mustHaveCapabilities",
					RefID:   capID,
					Message: "references non-existent capability",
				})
			}
		}
	}

	// Cross-reference validation: BMC internal references
	for _, bmc := range e.BMCs {
		bmcID := bmc.Metadata.ID
		if bmcID == "" {
			bmcID = bmc.Metadata.Title
		}

		// Build lookup sets for internal IDs
		segmentIDs := make(map[string]bool)
		for _, seg := range bmc.CustomerSegments {
			segmentIDs[seg.ID] = true
		}

		valuePropIDs := make(map[string]bool)
		for _, vp := range bmc.ValuePropositions {
			valuePropIDs[vp.ID] = true
		}

		// Validate ValueProposition.SegmentRefs
		for _, vp := range bmc.ValuePropositions {
			for _, segRef := range vp.SegmentRefs {
				if !segmentIDs[segRef] {
					errs = append(errs, ValidationError{
						Module:  "canvas",
						Type:    "bmc",
						ID:      bmcID,
						Field:   "valuePropositions.segmentRefs",
						RefID:   segRef,
						Message: "references non-existent customer segment",
					})
				}
			}
		}

		// Validate Channel.SegmentRefs
		for _, ch := range bmc.Channels {
			for _, segRef := range ch.SegmentRefs {
				if !segmentIDs[segRef] {
					errs = append(errs, ValidationError{
						Module:  "canvas",
						Type:    "bmc",
						ID:      bmcID,
						Field:   "channels.segmentRefs",
						RefID:   segRef,
						Message: "references non-existent customer segment",
					})
				}
			}
		}

		// Validate CustomerRelation.SegmentRefs
		for _, cr := range bmc.CustomerRelationships {
			for _, segRef := range cr.SegmentRefs {
				if !segmentIDs[segRef] {
					errs = append(errs, ValidationError{
						Module:  "canvas",
						Type:    "bmc",
						ID:      bmcID,
						Field:   "customerRelationships.segmentRefs",
						RefID:   segRef,
						Message: "references non-existent customer segment",
					})
				}
			}
		}

		// Validate RevenueStream.SegmentRefs and ValuePropRefs
		for _, rs := range bmc.RevenueStreams {
			for _, segRef := range rs.SegmentRefs {
				if !segmentIDs[segRef] {
					errs = append(errs, ValidationError{
						Module:  "canvas",
						Type:    "bmc",
						ID:      bmcID,
						Field:   "revenueStreams.segmentRefs",
						RefID:   segRef,
						Message: "references non-existent customer segment",
					})
				}
			}
			for _, vpRef := range rs.ValuePropRefs {
				if !valuePropIDs[vpRef] {
					errs = append(errs, ValidationError{
						Module:  "canvas",
						Type:    "bmc",
						ID:      bmcID,
						Field:   "revenueStreams.valuePropRefs",
						RefID:   vpRef,
						Message: "references non-existent value proposition",
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
	CapabilityStacks     int            `json:"capabilityStacks"`
	TotalCapabilities    int            `json:"totalCapabilities"`
	PRISMDocuments       int            `json:"prismDocuments"`
	TotalMetrics         int            `json:"totalMetrics"`
	TotalServices        int            `json:"totalServices"`
	TotalInitiatives     int            `json:"totalInitiatives"`
	TotalOKRSets         int            `json:"totalOkrSets"`
	TotalV2MOMs          int            `json:"totalV2moms"`
	TotalObjectives      int            `json:"totalObjectives"`
	TotalRoadmaps        int            `json:"totalRoadmaps"`
	TotalPhases          int            `json:"totalPhases"`
	TotalOpportunitySpec int            `json:"totalOpportunitySpecs"`
	TotalBMCs            int            `json:"totalBmcs"`
	ByStatus             map[string]int `json:"byStatus"`
	ByDomain             map[string]int `json:"byDomain"`
}

// Stats returns ecosystem statistics.
func (e *Ecosystem) Stats() Stats {
	stats := Stats{
		CapabilityStacks:     len(e.CapabilityStacks),
		TotalCapabilities:    0,
		PRISMDocuments:       len(e.PRISMDocuments),
		TotalMetrics:         len(e.AllMetrics()),
		TotalServices:        len(e.AllServices()),
		TotalInitiatives:     len(e.AllInitiatives()),
		TotalOKRSets:         len(e.OKRSets),
		TotalV2MOMs:          len(e.V2MOMs),
		TotalObjectives:      len(e.AllObjectives()),
		TotalRoadmaps:        len(e.Roadmaps),
		TotalPhases:          len(e.AllPhases()),
		TotalOpportunitySpec: len(e.OpportunitySpecs),
		TotalBMCs:            len(e.BMCs),
		ByStatus:             make(map[string]int),
		ByDomain:             make(map[string]int),
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
//	    v2moms/*.json
//	    roadmaps/*.json
//	    opportunity-specs/*.json
//	  canvas/
//	    bmcs/*.json
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
	if files, err := filepath.Glob(filepath.Join(roadmapDir, "v2moms", "*.json")); err == nil {
		config.Roadmap.V2MOMs = files
	}
	if files, err := filepath.Glob(filepath.Join(roadmapDir, "roadmaps", "*.json")); err == nil {
		config.Roadmap.Roadmaps = files
	}
	if files, err := filepath.Glob(filepath.Join(roadmapDir, "opportunity-specs", "*.json")); err == nil {
		config.Roadmap.OpportunitySpecs = files
	}

	// Scan canvas files
	canvasDir := filepath.Join(dir, "canvas")
	if files, err := filepath.Glob(filepath.Join(canvasDir, "bmcs", "*.json")); err == nil {
		config.Canvas.BMCs = files
	}

	return Load(config)
}
