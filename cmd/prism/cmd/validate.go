package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/prism/ecosystem"
	"github.com/spf13/cobra"
)

var (
	validateConfigFile string
	validateDirectory  string
	validateJSON       bool
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate ecosystem documents",
	Long: `Validate all ecosystem documents and cross-references.

Examples:
  prism validate --config prism.yaml
  prism validate --dir ./ecosystem
  prism validate --all`,
	RunE: runValidate,
}

func init() {
	validateCmd.Flags().StringVarP(&validateConfigFile, "config", "c", "", "Configuration file (JSON)")
	validateCmd.Flags().StringVarP(&validateDirectory, "dir", "d", "", "Directory to load from")
	validateCmd.Flags().BoolVar(&validateJSON, "json", false, "Output as JSON")
	validateCmd.Flags().Bool("all", false, "Validate all documents (alias for --dir .)")
}

func runValidate(cmd *cobra.Command, args []string) error {
	var eco *ecosystem.Ecosystem
	var err error

	// Handle --all flag
	all, _ := cmd.Flags().GetBool("all")
	if all && validateDirectory == "" && validateConfigFile == "" {
		validateDirectory = "."
	}

	switch {
	case validateConfigFile != "":
		eco, err = ecosystem.LoadFromFile(validateConfigFile)
	case validateDirectory != "":
		eco, err = ecosystem.LoadFromDirectory(validateDirectory)
	default:
		return fmt.Errorf("either --config, --dir, or --all is required")
	}

	if err != nil {
		return fmt.Errorf("loading ecosystem: %w", err)
	}

	errs := eco.Validate()

	if validateJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		result := map[string]interface{}{
			"valid":  !errs.HasErrors(),
			"errors": errs,
		}
		return enc.Encode(result)
	}

	// Text output
	if !errs.HasErrors() {
		fmt.Println("✓ All documents valid")
		return nil
	}

	fmt.Printf("✗ Found %d validation error(s):\n\n", len(errs))
	for i, e := range errs {
		fmt.Printf("%d. [%s/%s] %s\n", i+1, e.Module, e.Type, e.ID)
		fmt.Printf("   Field: %s\n", e.Field)
		if e.RefID != "" {
			fmt.Printf("   RefID: %s\n", e.RefID)
		}
		fmt.Printf("   %s\n\n", e.Message)
	}

	return fmt.Errorf("validation failed with %d errors", len(errs))
}
