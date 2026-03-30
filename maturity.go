package prism

import "fmt"

// MaturityModel defines the maturity model configuration and cell scores.
type MaturityModel struct {
	Levels []MaturityLevelDef `json:"levels,omitempty"`
	Cells  []MaturityCell     `json:"cells,omitempty"`
}

// MaturityLevelDef defines a maturity level.
type MaturityLevelDef struct {
	Level       int    `json:"level"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// MaturityCell represents maturity state for a domain/stage combination.
type MaturityCell struct {
	Domain        string  `json:"domain"`
	Stage         string  `json:"stage"`
	CurrentLevel  int     `json:"currentLevel"`
	TargetLevel   int     `json:"targetLevel,omitempty"`
	PrimaryKPI    string  `json:"primaryKPI,omitempty"`
	KPITarget     string  `json:"kpiTarget,omitempty"`
	MaturityScore float64 `json:"maturityScore,omitempty"`
}

// DefaultMaturityLevels returns the standard 5-level maturity definitions.
func DefaultMaturityLevels() []MaturityLevelDef {
	return []MaturityLevelDef{
		{
			Level:       MaturityLevel1,
			Name:        "Reactive",
			Description: "Ad-hoc processes, minimal documentation, reactive response",
		},
		{
			Level:       MaturityLevel2,
			Name:        "Basic",
			Description: "Basic processes in place, some documentation, limited automation",
		},
		{
			Level:       MaturityLevel3,
			Name:        "Defined",
			Description: "Standardized processes, documented procedures, consistent execution",
		},
		{
			Level:       MaturityLevel4,
			Name:        "Managed",
			Description: "Measured and controlled, data-driven decisions, proactive management",
		},
		{
			Level:       MaturityLevel5,
			Name:        "Optimizing",
			Description: "Continuous improvement, automated optimization, industry-leading practices",
		},
	}
}

// NewMaturityModel creates a new maturity model with default levels.
func NewMaturityModel() *MaturityModel {
	return &MaturityModel{
		Levels: DefaultMaturityLevels(),
		Cells:  make([]MaturityCell, 0),
	}
}

// NewMaturityModelWithCells creates a maturity model with cells for all domain/stage combinations.
func NewMaturityModelWithCells() *MaturityModel {
	return NewMaturityModelForDomains(AllDomains())
}

// NewMaturityModelForDomains creates a maturity model with cells for specified domains only.
func NewMaturityModelForDomains(domains []string) *MaturityModel {
	model := NewMaturityModel()

	// Create cells for specified domain/stage combinations
	for _, domain := range domains {
		for _, stage := range AllStages() {
			model.Cells = append(model.Cells, MaturityCell{
				Domain:       domain,
				Stage:        stage,
				CurrentLevel: MaturityLevel1,
				TargetLevel:  MaturityLevel3,
			})
		}
	}

	return model
}

// GetCell returns the maturity cell for a specific domain/stage combination.
func (m *MaturityModel) GetCell(domain, stage string) *MaturityCell {
	for i := range m.Cells {
		if m.Cells[i].Domain == domain && m.Cells[i].Stage == stage {
			return &m.Cells[i]
		}
	}
	return nil
}

// GetCellsByDomain returns all maturity cells for a specific domain.
func (m *MaturityModel) GetCellsByDomain(domain string) []MaturityCell {
	var result []MaturityCell
	for _, cell := range m.Cells {
		if cell.Domain == domain {
			result = append(result, cell)
		}
	}
	return result
}

// GetCellsByStage returns all maturity cells for a specific stage.
func (m *MaturityModel) GetCellsByStage(stage string) []MaturityCell {
	var result []MaturityCell
	for _, cell := range m.Cells {
		if cell.Stage == stage {
			result = append(result, cell)
		}
	}
	return result
}

// CalculateMaturityScore calculates the normalized maturity score (0.0-1.0) for a cell.
func (c *MaturityCell) CalculateMaturityScore() float64 {
	return float64(c.CurrentLevel) / float64(MaturityLevel5)
}

// AverageMaturityLevel returns the average maturity level across all cells.
func (m *MaturityModel) AverageMaturityLevel() float64 {
	if len(m.Cells) == 0 {
		return 0
	}

	var sum float64
	for _, cell := range m.Cells {
		sum += float64(cell.CurrentLevel)
	}
	return sum / float64(len(m.Cells))
}

// AverageMaturityScore returns the average normalized maturity score (0.0-1.0).
func (m *MaturityModel) AverageMaturityScore() float64 {
	return m.AverageMaturityLevel() / float64(MaturityLevel5)
}

// DomainMaturityLevel returns the average maturity level for a specific domain.
func (m *MaturityModel) DomainMaturityLevel(domain string) float64 {
	cells := m.GetCellsByDomain(domain)
	if len(cells) == 0 {
		return 0
	}

	var sum float64
	for _, cell := range cells {
		sum += float64(cell.CurrentLevel)
	}
	return sum / float64(len(cells))
}

// StageMaturityLevel returns the average maturity level for a specific stage.
func (m *MaturityModel) StageMaturityLevel(stage string) float64 {
	cells := m.GetCellsByStage(stage)
	if len(cells) == 0 {
		return 0
	}

	var sum float64
	for _, cell := range cells {
		sum += float64(cell.CurrentLevel)
	}
	return sum / float64(len(cells))
}

// Validate validates the maturity model.
func (m *MaturityModel) Validate() ValidationErrors {
	var errs ValidationErrors

	// Validate levels
	for i, level := range m.Levels {
		if err := ValidateMaturityLevel(level.Level); err != nil {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("levels[%d].level", i),
				Value:   fmt.Sprintf("%d", level.Level),
				Message: err.Error(),
			})
		}
		if level.Name == "" {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("levels[%d].name", i),
				Message: "is required",
			})
		}
	}

	// Validate cells
	for i, cell := range m.Cells {
		if err := ValidateDomain(cell.Domain); err != nil {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("cells[%d].domain", i),
				Value:   cell.Domain,
				Message: err.Error(),
			})
		}
		if err := ValidateStage(cell.Stage); err != nil {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("cells[%d].stage", i),
				Value:   cell.Stage,
				Message: err.Error(),
			})
		}
		if err := ValidateMaturityLevel(cell.CurrentLevel); err != nil {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("cells[%d].currentLevel", i),
				Value:   fmt.Sprintf("%d", cell.CurrentLevel),
				Message: err.Error(),
			})
		}
		if cell.TargetLevel != 0 {
			if err := ValidateMaturityLevel(cell.TargetLevel); err != nil {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("cells[%d].targetLevel", i),
					Value:   fmt.Sprintf("%d", cell.TargetLevel),
					Message: err.Error(),
				})
			}
		}
	}

	// Check for duplicate domain/stage combinations
	seen := make(map[string]int)
	for i, cell := range m.Cells {
		key := cell.Domain + "/" + cell.Stage
		if prevIdx, exists := seen[key]; exists {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("cells[%d]", i),
				Value:   key,
				Message: fmt.Sprintf("duplicate domain/stage, also defined at cells[%d]", prevIdx),
			})
		}
		seen[key] = i
	}

	return errs
}

// SetCellLevel sets the current maturity level for a domain/stage combination.
// Creates the cell if it doesn't exist.
func (m *MaturityModel) SetCellLevel(domain, stage string, level int) error {
	if err := ValidateDomain(domain); err != nil {
		return err
	}
	if err := ValidateStage(stage); err != nil {
		return err
	}
	if err := ValidateMaturityLevel(level); err != nil {
		return err
	}

	cell := m.GetCell(domain, stage)
	if cell == nil {
		m.Cells = append(m.Cells, MaturityCell{
			Domain:       domain,
			Stage:        stage,
			CurrentLevel: level,
		})
	} else {
		cell.CurrentLevel = level
	}
	return nil
}

// UpdateMaturityScores calculates and updates maturity scores for all cells.
func (m *MaturityModel) UpdateMaturityScores() {
	for i := range m.Cells {
		m.Cells[i].MaturityScore = m.Cells[i].CalculateMaturityScore()
	}
}
