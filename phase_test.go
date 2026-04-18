package prism

import "testing"

func TestPhaseGetGoalTarget(t *testing.T) {
	phase := &Phase{
		ID:   "phase-q1",
		Name: "Q1 2026",
		GoalTargets: []PhaseGoalTarget{
			{GoalID: "goal-a", EnterLevel: 2, ExitLevel: 3},
			{GoalID: "goal-b", EnterLevel: 3, ExitLevel: 4},
		},
	}

	tests := []struct {
		goalID    string
		wantNil   bool
		wantEnter int
		wantExit  int
	}{
		{"goal-a", false, 2, 3},
		{"goal-b", false, 3, 4},
		{"goal-c", true, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.goalID, func(t *testing.T) {
			got := phase.GetGoalTarget(tt.goalID)
			if tt.wantNil {
				if got != nil {
					t.Errorf("GetGoalTarget(%q) = %v, want nil", tt.goalID, got)
				}
			} else {
				if got == nil {
					t.Errorf("GetGoalTarget(%q) = nil, want target", tt.goalID)
				} else {
					if got.EnterLevel != tt.wantEnter {
						t.Errorf("GetGoalTarget(%q).EnterLevel = %d, want %d", tt.goalID, got.EnterLevel, tt.wantEnter)
					}
					if got.ExitLevel != tt.wantExit {
						t.Errorf("GetGoalTarget(%q).ExitLevel = %d, want %d", tt.goalID, got.ExitLevel, tt.wantExit)
					}
				}
			}
		})
	}
}

func TestPhaseGetSwimlane(t *testing.T) {
	phase := &Phase{
		ID:   "phase-q1",
		Name: "Q1 2026",
		Swimlanes: []Swimlane{
			{ID: "sw-security", Name: "Security", Domain: DomainSecurity},
			{ID: "sw-ops", Name: "Operations", Domain: DomainOperations},
		},
	}

	tests := []struct {
		swimlaneID string
		wantNil    bool
		wantName   string
	}{
		{"sw-security", false, "Security"},
		{"sw-ops", false, "Operations"},
		{"sw-unknown", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.swimlaneID, func(t *testing.T) {
			got := phase.GetSwimlane(tt.swimlaneID)
			if tt.wantNil {
				if got != nil {
					t.Errorf("GetSwimlane(%q) = %v, want nil", tt.swimlaneID, got)
				}
			} else {
				if got == nil {
					t.Errorf("GetSwimlane(%q) = nil, want swimlane", tt.swimlaneID)
				} else if got.Name != tt.wantName {
					t.Errorf("GetSwimlane(%q).Name = %q, want %q", tt.swimlaneID, got.Name, tt.wantName)
				}
			}
		})
	}
}

func TestPhaseAllInitiativeIDs(t *testing.T) {
	phase := &Phase{
		ID:   "phase-q1",
		Name: "Q1 2026",
		Swimlanes: []Swimlane{
			{
				ID:            "sw-security",
				Name:          "Security",
				InitiativeIDs: []string{"init-a", "init-b"},
			},
			{
				ID:            "sw-ops",
				Name:          "Operations",
				InitiativeIDs: []string{"init-c", "init-b"}, // init-b is duplicated
			},
		},
	}

	got := phase.AllInitiativeIDs()

	// Should have 3 unique IDs
	if len(got) != 3 {
		t.Errorf("AllInitiativeIDs() returned %d IDs, want 3", len(got))
	}

	// Check all expected IDs are present
	expected := map[string]bool{"init-a": true, "init-b": true, "init-c": true}
	for _, id := range got {
		if !expected[id] {
			t.Errorf("AllInitiativeIDs() contains unexpected ID %q", id)
		}
		delete(expected, id)
	}
	if len(expected) > 0 {
		t.Errorf("AllInitiativeIDs() missing IDs: %v", expected)
	}
}

func TestAllQuarters(t *testing.T) {
	quarters := AllQuarters()
	if len(quarters) != 4 {
		t.Errorf("AllQuarters() returned %d quarters, want 4", len(quarters))
	}

	expected := []string{QuarterQ1, QuarterQ2, QuarterQ3, QuarterQ4}
	for i, q := range expected {
		if quarters[i] != q {
			t.Errorf("AllQuarters()[%d] = %q, want %q", i, quarters[i], q)
		}
	}
}

func TestAllPhaseStatuses(t *testing.T) {
	statuses := AllPhaseStatuses()
	if len(statuses) != 3 {
		t.Errorf("AllPhaseStatuses() returned %d statuses, want 3", len(statuses))
	}

	expected := map[string]bool{
		PhaseStatusPlanning:   true,
		PhaseStatusInProgress: true,
		PhaseStatusCompleted:  true,
	}

	for _, s := range statuses {
		if !expected[s] {
			t.Errorf("AllPhaseStatuses() contains unexpected status %q", s)
		}
	}
}
