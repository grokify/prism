package output

// TruncateString truncates a string to maxLen characters, adding "..." if truncated.
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// OperatorSymbol returns the display symbol for an SLO operator.
func OperatorSymbol(op string) string {
	switch op {
	case "gte":
		return ">="
	case "lte":
		return "<="
	case "gt":
		return ">"
	case "lt":
		return "<"
	case "eq":
		return "="
	default:
		return op
	}
}

// SafePercent calculates percentage safely, returning 0 if total is 0.
func SafePercent(value, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(value) / float64(total) * 100
}

// GoalStatus returns a status string based on current vs target maturity level.
func GoalStatus(current, target int) string {
	if current >= target {
		return "Achieved"
	}
	if current == target-1 {
		return "On Track"
	}
	return "Behind"
}

// MaturityLevelName returns the name for a maturity level (1-5).
func MaturityLevelName(level int) string {
	names := map[int]string{
		1: "Reactive",
		2: "Basic",
		3: "Defined",
		4: "Managed",
		5: "Optimizing",
	}
	if name, ok := names[level]; ok {
		return name
	}
	return "Unknown"
}

// StatusSymbol returns a visual symbol for a status.
func StatusSymbol(status string) string {
	switch status {
	case "achieved", "completed", "done":
		return "[x]"
	case "in_progress", "active":
		return "[~]"
	case "not_started", "planned", "pending":
		return "[ ]"
	case "blocked":
		return "[!]"
	default:
		return "[-]"
	}
}

// FormatInitiativeStatus formats an initiative status for display.
func FormatInitiativeStatus(status string) string {
	switch status {
	case "completed":
		return "Completed"
	case "in_progress":
		return "In Progress"
	case "planned":
		return "Planned"
	case "not_started":
		return "Not Started"
	case "cancelled":
		return "Cancelled"
	default:
		return status
	}
}
