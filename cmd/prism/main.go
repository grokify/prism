// Package main provides the prism CLI entry point.
package main

import (
	"os"

	"github.com/grokify/prism/cmd/prism/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
