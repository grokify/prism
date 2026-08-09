package sitegen

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/grokify/prism/ecosystem"
)

// EvidenceConfig holds configuration for evidence display generation.
type EvidenceConfig struct {
	// CapabilityID is the capability to show evidence for.
	CapabilityID string

	// CapabilityName is the display name.
	CapabilityName string
}

// GenerateEvidencePage generates an evidence summary page for a capability.
func (g *Generator) GenerateEvidencePage(eco *ecosystem.Ecosystem, capabilityID string) error {
	ctx := eco.GetCapabilityContext(capabilityID)
	if ctx == nil {
		return nil // No capability found
	}

	tmpl, err := template.New("evidence").Parse(evidenceTemplate)
	if err != nil {
		return err
	}

	data := struct {
		Title           string
		Theme           string
		BaseURL         string
		HasSiteNavJS    bool
		HideGeneratedBy bool
		CapabilityID    string
		CapabilityName  string
		EvidenceLinks   []ecosystem.EvidenceLink
		Summary         *ecosystem.EvidenceSummary
	}{
		Title:           g.config.Title,
		Theme:           g.config.Theme,
		BaseURL:         g.config.BaseURL,
		HasSiteNavJS:    g.config.SiteNavJS != "",
		HideGeneratedBy: g.config.HideGeneratedBy,
		CapabilityID:    capabilityID,
		CapabilityName:  ctx.Capability.Name,
		EvidenceLinks:   ctx.EvidenceLinks,
		Summary:         ctx.EvidenceSummary,
	}

	// Ensure output directory exists
	evidenceDir := filepath.Join(g.config.OutputDir, "evidence")
	if err := os.MkdirAll(evidenceDir, 0755); err != nil {
		return err
	}

	// Write evidence page
	outPath := filepath.Join(evidenceDir, capabilityID+".html")
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}

// GenerateAllEvidencePages generates evidence pages for all capabilities with evidence.
func (g *Generator) GenerateAllEvidencePages(eco *ecosystem.Ecosystem) error {
	for capID := range eco.AllCapabilityEvidence() {
		if err := g.GenerateEvidencePage(eco, capID); err != nil {
			return err
		}
	}
	return nil
}
