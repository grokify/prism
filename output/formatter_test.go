package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestNewFormatter(t *testing.T) {
	tests := []struct {
		format string
		want   Format
	}{
		{"text", FormatText},
		{"json", FormatJSON},
		{"markdown", FormatMarkdown},
		{"toon", FormatTOON},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			f := NewFormatter(tt.format)
			if f.Format != tt.want {
				t.Errorf("NewFormatter(%q).Format = %v, want %v", tt.format, f.Format, tt.want)
			}
		})
	}
}

func TestValidFormats(t *testing.T) {
	formats := ValidFormats()
	if len(formats) != 4 {
		t.Errorf("ValidFormats() returned %d formats, want 4", len(formats))
	}

	expected := []string{"text", "json", "markdown", "toon"}
	for i, f := range expected {
		if formats[i] != f {
			t.Errorf("ValidFormats()[%d] = %q, want %q", i, formats[i], f)
		}
	}
}

func TestIsValidFormat(t *testing.T) {
	tests := []struct {
		format string
		valid  bool
	}{
		{"text", true},
		{"json", true},
		{"markdown", true},
		{"toon", true},
		{"xml", false},
		{"csv", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			if got := IsValidFormat(tt.format); got != tt.valid {
				t.Errorf("IsValidFormat(%q) = %v, want %v", tt.format, got, tt.valid)
			}
		})
	}
}

func TestWriteJSON(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("json", &buf)

	data := map[string]string{"key": "value"}
	if err := f.WriteJSON(data); err != nil {
		t.Fatalf("WriteJSON() error = %v", err)
	}

	var result map[string]string
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	if result["key"] != "value" {
		t.Errorf("WriteJSON() output key = %q, want %q", result["key"], "value")
	}
}

func TestWriteTableText(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("text", &buf)

	data := &TableData{
		Title:   "Test Table",
		Headers: []string{"Name", "Value"},
		Rows: [][]string{
			{"foo", "1"},
			{"bar", "2"},
		},
		Summary: "Total: 2 items",
	}

	if err := f.WriteTable(data); err != nil {
		t.Fatalf("WriteTable() error = %v", err)
	}

	output := buf.String()

	// Check title
	if !strings.Contains(output, "Test Table") {
		t.Error("WriteTable() output missing title")
	}

	// Check headers
	if !strings.Contains(output, "Name") || !strings.Contains(output, "Value") {
		t.Error("WriteTable() output missing headers")
	}

	// Check data
	if !strings.Contains(output, "foo") || !strings.Contains(output, "bar") {
		t.Error("WriteTable() output missing data rows")
	}

	// Check summary
	if !strings.Contains(output, "Total: 2 items") {
		t.Error("WriteTable() output missing summary")
	}
}

func TestWriteTableMarkdown(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("markdown", &buf)

	data := &TableData{
		Title:   "Test Table",
		Headers: []string{"Name", "Value"},
		Rows: [][]string{
			{"foo", "1"},
			{"bar", "2"},
		},
	}

	if err := f.WriteTable(data); err != nil {
		t.Fatalf("WriteTable() error = %v", err)
	}

	output := buf.String()

	// Check markdown title
	if !strings.Contains(output, "## Test Table") {
		t.Error("WriteTable() markdown output missing ## title")
	}

	// Check markdown table format
	if !strings.Contains(output, "| Name |") {
		t.Error("WriteTable() markdown output missing pipe-delimited headers")
	}

	// Check separator
	if !strings.Contains(output, "| --- |") {
		t.Error("WriteTable() markdown output missing separator")
	}
}

func TestWriteTableTOON(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("toon", &buf)

	data := &TableData{
		Title:   "Goals",
		Headers: []string{"Name", "Level"},
		Rows: [][]string{
			{"Goal A", "M3"},
			{"Goal B", "M4"},
		},
	}

	if err := f.WriteTable(data); err != nil {
		t.Fatalf("WriteTable() error = %v", err)
	}

	output := strings.TrimSpace(buf.String())

	// TOON format: title;headers;row1;row2
	expected := "Goals;Name,Level;Goal A,M3;Goal B,M4"
	if output != expected {
		t.Errorf("WriteTable() TOON output = %q, want %q", output, expected)
	}
}

func TestWriteTableJSON(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("json", &buf)

	data := &TableData{
		Title:   "Test",
		Headers: []string{"A", "B"},
		Rows: [][]string{
			{"1", "2"},
		},
		Summary: "Summary text",
	}

	if err := f.WriteTable(data); err != nil {
		t.Fatalf("WriteTable() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if result["title"] != "Test" {
		t.Errorf("WriteTable() JSON title = %v, want %q", result["title"], "Test")
	}
	if result["summary"] != "Summary text" {
		t.Errorf("WriteTable() JSON summary = %v, want %q", result["summary"], "Summary text")
	}
}

func TestWriteTableEmptyData(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("text", &buf)

	// Empty headers and rows should not error
	data := &TableData{
		Title:   "Empty",
		Headers: []string{},
		Rows:    [][]string{},
	}

	if err := f.WriteTable(data); err != nil {
		t.Fatalf("WriteTable() error = %v", err)
	}

	// Only title should be present
	output := buf.String()
	if !strings.Contains(output, "Empty") {
		t.Error("WriteTable() empty data should still show title")
	}
}

func TestWriteDetailText(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("text", &buf)

	data := &DetailData{
		Title: "Goal Details",
		Fields: []DetailField{
			{Key: "Name", Value: "Reliability"},
			{Key: "Owner", Value: "SRE Team"},
		},
		Sections: []DetailSection{
			{
				Title: "Requirements",
				Items: []string{"SLO >= 99.9%", "MTTR < 1h"},
			},
		},
	}

	if err := f.WriteDetail(data); err != nil {
		t.Fatalf("WriteDetail() error = %v", err)
	}

	output := buf.String()

	if !strings.Contains(output, "Goal Details") {
		t.Error("WriteDetail() output missing title")
	}
	if !strings.Contains(output, "Name: Reliability") {
		t.Error("WriteDetail() output missing field")
	}
	if !strings.Contains(output, "Requirements:") {
		t.Error("WriteDetail() output missing section title")
	}
	if !strings.Contains(output, "- SLO >= 99.9%") {
		t.Error("WriteDetail() output missing section item")
	}
}

func TestWriteDetailMarkdown(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("markdown", &buf)

	data := &DetailData{
		Title: "Goal Details",
		Fields: []DetailField{
			{Key: "Owner", Value: "Team A"},
		},
		Sections: []DetailSection{
			{
				Title:   "Notes",
				Content: "Some content here",
			},
		},
	}

	if err := f.WriteDetail(data); err != nil {
		t.Fatalf("WriteDetail() error = %v", err)
	}

	output := buf.String()

	if !strings.Contains(output, "# Goal Details") {
		t.Error("WriteDetail() markdown output missing # title")
	}
	if !strings.Contains(output, "**Owner:** Team A") {
		t.Error("WriteDetail() markdown output missing bold field")
	}
	if !strings.Contains(output, "### Notes") {
		t.Error("WriteDetail() markdown output missing ### section")
	}
}

func TestWriteDetailTOON(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("toon", &buf)

	data := &DetailData{
		Title: "Goal",
		Fields: []DetailField{
			{Key: "level", Value: "M3"},
		},
		Sections: []DetailSection{
			{
				Title: "items",
				Items: []string{"a", "b"},
			},
		},
	}

	if err := f.WriteDetail(data); err != nil {
		t.Fatalf("WriteDetail() error = %v", err)
	}

	output := strings.TrimSpace(buf.String())

	// TOON detail format: title|key=value|section:items
	if !strings.Contains(output, "Goal|level=M3|items:a,b") {
		t.Errorf("WriteDetail() TOON output = %q, expected Goal|level=M3|items:a,b", output)
	}
}

func TestWriteDetailJSON(t *testing.T) {
	var buf bytes.Buffer
	f := NewFormatterWithWriter("json", &buf)

	data := &DetailData{
		Title: "Details",
		Fields: []DetailField{
			{Key: "key1", Value: "value1"},
		},
		Sections: []DetailSection{
			{
				Title: "Section1",
				Items: []string{"item1"},
			},
		},
	}

	if err := f.WriteDetail(data); err != nil {
		t.Fatalf("WriteDetail() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if result["title"] != "Details" {
		t.Errorf("WriteDetail() JSON title = %v, want %q", result["title"], "Details")
	}

	fields, ok := result["fields"].(map[string]interface{})
	if !ok {
		t.Fatal("WriteDetail() JSON fields not a map")
	}
	if fields["key1"] != "value1" {
		t.Errorf("WriteDetail() JSON fields[key1] = %v, want %q", fields["key1"], "value1")
	}
}
