// Package ecosystem provides unified loading and querying across PRISM modules.
package ecosystem

import (
	"encoding/json"
	"fmt"
	"os"

	capability "github.com/grokify/prism-capability"
)

// Config defines the ecosystem configuration.
type Config struct {
	Name string `json:"name" yaml:"name"`

	Capability CapabilityConfig `json:"capability" yaml:"capability"`
	// Intelligence IntelligenceConfig `json:"intelligence" yaml:"intelligence"`
	// Execution    ExecutionConfig    `json:"execution" yaml:"execution"`
}

// CapabilityConfig defines capability stack sources.
type CapabilityConfig struct {
	Files []string `json:"files" yaml:"files"`
}

// Ecosystem holds loaded documents from all PRISM modules.
type Ecosystem struct {
	Config Config

	// Loaded documents
	CapabilityStacks []*capability.CapabilityStack
	// MaturityModel    *intelligence.MaturityModel
	// MaturityState    *intelligence.MaturityState
	// Roadmaps         []*execution.Roadmap
	// OKRs             []*execution.OKR
}

// Load creates an Ecosystem from a configuration.
func Load(config Config) (*Ecosystem, error) {
	eco := &Ecosystem{
		Config:           config,
		CapabilityStacks: make([]*capability.CapabilityStack, 0),
	}

	// Load capability stacks
	for _, file := range config.Capability.Files {
		stack, err := capability.LoadFromFile(file)
		if err != nil {
			return nil, fmt.Errorf("loading capability stack %s: %w", file, err)
		}
		eco.CapabilityStacks = append(eco.CapabilityStacks, stack)
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

// Validate validates all loaded documents and cross-references.
func (e *Ecosystem) Validate() error {
	// Validate each capability stack
	for i, stack := range e.CapabilityStacks {
		if errs := stack.Validate(); errs.HasErrors() {
			return fmt.Errorf("capability stack %d: %w", i, errs)
		}
	}

	// TODO: Cross-reference validation
	// - Check prismRef.sliIds exist in MaturityModel
	// - Check initiative.capabilityId exist in CapabilityStacks

	return nil
}

// Stats returns summary statistics about the ecosystem.
type Stats struct {
	CapabilityStacks int            `json:"capabilityStacks"`
	TotalCapabilities int           `json:"totalCapabilities"`
	ByStatus         map[string]int `json:"byStatus"`
	ByDomain         map[string]int `json:"byDomain"`
}

// Stats returns ecosystem statistics.
func (e *Ecosystem) Stats() Stats {
	stats := Stats{
		CapabilityStacks:  len(e.CapabilityStacks),
		TotalCapabilities: 0,
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
