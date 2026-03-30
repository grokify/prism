// Package schema provides embedded JSON Schema for PRISM types.
package schema

import (
	_ "embed"
	"encoding/json"
)

//go:embed prism.schema.json
var prismSchemaJSON []byte

// PRISMSchemaJSON returns the raw JSON Schema bytes.
func PRISMSchemaJSON() []byte {
	return prismSchemaJSON
}

// PRISMSchemaMap returns the schema as a map for programmatic access.
func PRISMSchemaMap() (map[string]any, error) {
	var m map[string]any
	if err := json.Unmarshal(prismSchemaJSON, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// SchemaID returns the schema ID.
func SchemaID() string {
	return "https://github.com/grokify/prism/schema/prism.schema.json"
}
