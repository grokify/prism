package output

import "testing"

func TestTruncateString(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		maxLen int
		want   string
	}{
		{"short string", "hello", 10, "hello"},
		{"exact length", "hello", 5, "hello"},
		{"truncate with ellipsis", "hello world", 8, "hello..."},
		{"very short max", "hello", 3, "hel"},
		{"truncate at 2", "hello", 2, "he"},
		{"empty string", "", 5, ""},
		{"single char", "a", 5, "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TruncateString(tt.input, tt.maxLen)
			if got != tt.want {
				t.Errorf("TruncateString(%q, %d) = %q, want %q", tt.input, tt.maxLen, got, tt.want)
			}
		})
	}
}

func TestOperatorSymbol(t *testing.T) {
	tests := []struct {
		op   string
		want string
	}{
		{"gte", ">="},
		{"lte", "<="},
		{"gt", ">"},
		{"lt", "<"},
		{"eq", "="},
		{"unknown", "unknown"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.op, func(t *testing.T) {
			got := OperatorSymbol(tt.op)
			if got != tt.want {
				t.Errorf("OperatorSymbol(%q) = %q, want %q", tt.op, got, tt.want)
			}
		})
	}
}

func TestSafePercent(t *testing.T) {
	tests := []struct {
		name  string
		value int
		total int
		want  float64
	}{
		{"zero total", 5, 0, 0},
		{"zero value", 0, 10, 0},
		{"50 percent", 5, 10, 50},
		{"100 percent", 10, 10, 100},
		{"25 percent", 1, 4, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SafePercent(tt.value, tt.total)
			if got != tt.want {
				t.Errorf("SafePercent(%d, %d) = %f, want %f", tt.value, tt.total, got, tt.want)
			}
		})
	}
}

func TestGoalStatus(t *testing.T) {
	tests := []struct {
		name    string
		current int
		target  int
		want    string
	}{
		{"achieved exact", 5, 5, "Achieved"},
		{"achieved exceeded", 5, 4, "Achieved"},
		{"on track", 4, 5, "On Track"},
		{"behind", 3, 5, "Behind"},
		{"far behind", 1, 5, "Behind"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GoalStatus(tt.current, tt.target)
			if got != tt.want {
				t.Errorf("GoalStatus(%d, %d) = %q, want %q", tt.current, tt.target, got, tt.want)
			}
		})
	}
}

func TestMaturityLevelName(t *testing.T) {
	tests := []struct {
		level int
		want  string
	}{
		{1, "Reactive"},
		{2, "Basic"},
		{3, "Defined"},
		{4, "Managed"},
		{5, "Optimizing"},
		{0, "Unknown"},
		{6, "Unknown"},
		{-1, "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := MaturityLevelName(tt.level)
			if got != tt.want {
				t.Errorf("MaturityLevelName(%d) = %q, want %q", tt.level, got, tt.want)
			}
		})
	}
}

func TestStatusSymbol(t *testing.T) {
	tests := []struct {
		status string
		want   string
	}{
		{"achieved", "[x]"},
		{"completed", "[x]"},
		{"done", "[x]"},
		{"in_progress", "[~]"},
		{"active", "[~]"},
		{"not_started", "[ ]"},
		{"planned", "[ ]"},
		{"pending", "[ ]"},
		{"blocked", "[!]"},
		{"unknown", "[-]"},
		{"", "[-]"},
	}

	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			got := StatusSymbol(tt.status)
			if got != tt.want {
				t.Errorf("StatusSymbol(%q) = %q, want %q", tt.status, got, tt.want)
			}
		})
	}
}

func TestFormatInitiativeStatus(t *testing.T) {
	tests := []struct {
		status string
		want   string
	}{
		{"completed", "Completed"},
		{"in_progress", "In Progress"},
		{"planned", "Planned"},
		{"not_started", "Not Started"},
		{"cancelled", "Cancelled"},
		{"custom", "custom"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			got := FormatInitiativeStatus(tt.status)
			if got != tt.want {
				t.Errorf("FormatInitiativeStatus(%q) = %q, want %q", tt.status, got, tt.want)
			}
		})
	}
}
