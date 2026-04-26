// Package output provides formatting utilities for PRISM data output.
package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Format represents supported output formats.
type Format string

const (
	FormatText     Format = "text"
	FormatJSON     Format = "json"
	FormatMarkdown Format = "markdown"
	FormatTOON     Format = "toon"
)

// ValidFormats returns all valid format strings.
func ValidFormats() []string {
	return []string{string(FormatText), string(FormatJSON), string(FormatMarkdown), string(FormatTOON)}
}

// IsValidFormat checks if a format string is valid.
func IsValidFormat(f string) bool {
	for _, valid := range ValidFormats() {
		if f == valid {
			return true
		}
	}
	return false
}

// Formatter handles output formatting.
type Formatter struct {
	Format Format
	Writer io.Writer
}

// NewFormatter creates a formatter for the given format string.
func NewFormatter(format string) *Formatter {
	return &Formatter{
		Format: Format(format),
		Writer: os.Stdout,
	}
}

// NewFormatterWithWriter creates a formatter with a custom writer.
func NewFormatterWithWriter(format string, w io.Writer) *Formatter {
	return &Formatter{
		Format: Format(format),
		Writer: w,
	}
}

// WriteJSON outputs data as JSON.
func (f *Formatter) WriteJSON(data interface{}) error {
	enc := json.NewEncoder(f.Writer)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

// TableData represents tabular data for formatting.
type TableData struct {
	Title   string
	Headers []string
	Rows    [][]string
	Summary string
}

// WriteTable outputs tabular data in the appropriate format.
func (f *Formatter) WriteTable(data *TableData) error {
	switch f.Format {
	case FormatJSON:
		return f.writeTableJSON(data)
	case FormatMarkdown:
		return f.writeTableMarkdown(data)
	case FormatTOON:
		return f.writeTableTOON(data)
	default:
		return f.writeTableText(data)
	}
}

func (f *Formatter) writeTableText(data *TableData) error {
	if data.Title != "" {
		fmt.Fprintln(f.Writer, data.Title)
		fmt.Fprintln(f.Writer, strings.Repeat("=", len(data.Title)))
		fmt.Fprintln(f.Writer)
	}

	if len(data.Headers) == 0 || len(data.Rows) == 0 {
		return nil
	}

	// Calculate column widths
	widths := make([]int, len(data.Headers))
	for i, h := range data.Headers {
		widths[i] = len(h)
	}
	for _, row := range data.Rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	// Print headers
	for i, h := range data.Headers {
		fmt.Fprintf(f.Writer, "%-*s ", widths[i], h)
	}
	fmt.Fprintln(f.Writer)

	// Print separator
	for i := range data.Headers {
		fmt.Fprintf(f.Writer, "%-*s ", widths[i], strings.Repeat("-", widths[i]))
	}
	fmt.Fprintln(f.Writer)

	// Print rows
	for _, row := range data.Rows {
		for i, cell := range row {
			if i < len(widths) {
				fmt.Fprintf(f.Writer, "%-*s ", widths[i], cell)
			}
		}
		fmt.Fprintln(f.Writer)
	}

	if data.Summary != "" {
		fmt.Fprintln(f.Writer)
		fmt.Fprintln(f.Writer, data.Summary)
	}

	return nil
}

func (f *Formatter) writeTableMarkdown(data *TableData) error {
	if data.Title != "" {
		fmt.Fprintf(f.Writer, "## %s\n\n", data.Title)
	}

	if len(data.Headers) == 0 {
		return nil
	}

	// Print headers
	fmt.Fprint(f.Writer, "|")
	for _, h := range data.Headers {
		fmt.Fprintf(f.Writer, " %s |", h)
	}
	fmt.Fprintln(f.Writer)

	// Print separator
	fmt.Fprint(f.Writer, "|")
	for range data.Headers {
		fmt.Fprint(f.Writer, " --- |")
	}
	fmt.Fprintln(f.Writer)

	// Print rows
	for _, row := range data.Rows {
		fmt.Fprint(f.Writer, "|")
		for _, cell := range row {
			fmt.Fprintf(f.Writer, " %s |", cell)
		}
		fmt.Fprintln(f.Writer)
	}

	if data.Summary != "" {
		fmt.Fprintln(f.Writer)
		fmt.Fprintf(f.Writer, "*%s*\n", data.Summary)
	}

	return nil
}

func (f *Formatter) writeTableTOON(data *TableData) error {
	// TOON: Token-Optimized Object Notation
	// Compact format optimized for LLM token efficiency
	// Format: title;h1,h2,h3;r1c1,r1c2,r1c3;r2c1,r2c2,r2c3

	if data.Title != "" {
		fmt.Fprintf(f.Writer, "%s;", data.Title)
	}

	// Headers
	fmt.Fprintf(f.Writer, "%s", strings.Join(data.Headers, ","))

	// Rows
	for _, row := range data.Rows {
		fmt.Fprintf(f.Writer, ";%s", strings.Join(row, ","))
	}

	fmt.Fprintln(f.Writer)

	return nil
}

func (f *Formatter) writeTableJSON(data *TableData) error {
	// Convert to JSON-friendly structure
	result := map[string]interface{}{
		"title":   data.Title,
		"headers": data.Headers,
		"rows":    data.Rows,
	}
	if data.Summary != "" {
		result["summary"] = data.Summary
	}
	return f.WriteJSON(result)
}

// DetailData represents detailed view data (key-value pairs with sections).
type DetailData struct {
	Title    string
	Fields   []DetailField
	Sections []DetailSection
}

// DetailField is a key-value pair.
type DetailField struct {
	Key   string
	Value string
}

// DetailSection is a named section with content.
type DetailSection struct {
	Title   string
	Content string
	Items   []string
	Table   *TableData
}

// WriteDetail outputs detailed data in the appropriate format.
func (f *Formatter) WriteDetail(data *DetailData) error {
	switch f.Format {
	case FormatJSON:
		return f.writeDetailJSON(data)
	case FormatMarkdown:
		return f.writeDetailMarkdown(data)
	case FormatTOON:
		return f.writeDetailTOON(data)
	default:
		return f.writeDetailText(data)
	}
}

func (f *Formatter) writeDetailText(data *DetailData) error {
	if data.Title != "" {
		fmt.Fprintln(f.Writer, data.Title)
	}

	for _, field := range data.Fields {
		fmt.Fprintf(f.Writer, "%s: %s\n", field.Key, field.Value)
	}

	for _, section := range data.Sections {
		fmt.Fprintln(f.Writer)
		if section.Title != "" {
			fmt.Fprintf(f.Writer, "%s:\n", section.Title)
		}
		if section.Content != "" {
			fmt.Fprintf(f.Writer, "  %s\n", section.Content)
		}
		for _, item := range section.Items {
			fmt.Fprintf(f.Writer, "  - %s\n", item)
		}
		if section.Table != nil {
			// Indent table output
			for _, row := range section.Table.Rows {
				fmt.Fprintf(f.Writer, "  %s\n", strings.Join(row, " | "))
			}
		}
	}

	return nil
}

func (f *Formatter) writeDetailMarkdown(data *DetailData) error {
	if data.Title != "" {
		fmt.Fprintf(f.Writer, "# %s\n\n", data.Title)
	}

	for _, field := range data.Fields {
		fmt.Fprintf(f.Writer, "**%s:** %s\n", field.Key, field.Value)
	}

	for _, section := range data.Sections {
		fmt.Fprintln(f.Writer)
		if section.Title != "" {
			fmt.Fprintf(f.Writer, "### %s\n\n", section.Title)
		}
		if section.Content != "" {
			fmt.Fprintf(f.Writer, "%s\n", section.Content)
		}
		if len(section.Items) > 0 {
			fmt.Fprintln(f.Writer)
			for _, item := range section.Items {
				fmt.Fprintf(f.Writer, "- %s\n", item)
			}
		}
		if section.Table != nil {
			fmt.Fprintln(f.Writer)
			_ = f.writeTableMarkdown(section.Table)
		}
	}

	return nil
}

func (f *Formatter) writeDetailTOON(data *DetailData) error {
	// TOON format for details: title|k1=v1|k2=v2|section:item1,item2
	parts := []string{}

	if data.Title != "" {
		parts = append(parts, data.Title)
	}

	for _, field := range data.Fields {
		parts = append(parts, fmt.Sprintf("%s=%s", field.Key, field.Value))
	}

	for _, section := range data.Sections {
		if len(section.Items) > 0 {
			parts = append(parts, fmt.Sprintf("%s:%s", section.Title, strings.Join(section.Items, ",")))
		}
	}

	fmt.Fprintln(f.Writer, strings.Join(parts, "|"))
	return nil
}

func (f *Formatter) writeDetailJSON(data *DetailData) error {
	result := map[string]interface{}{
		"title": data.Title,
	}

	fields := make(map[string]string)
	for _, field := range data.Fields {
		fields[field.Key] = field.Value
	}
	result["fields"] = fields

	sections := make([]map[string]interface{}, 0, len(data.Sections))
	for _, section := range data.Sections {
		s := map[string]interface{}{
			"title": section.Title,
		}
		if section.Content != "" {
			s["content"] = section.Content
		}
		if len(section.Items) > 0 {
			s["items"] = section.Items
		}
		sections = append(sections, s)
	}
	result["sections"] = sections

	return f.WriteJSON(result)
}
