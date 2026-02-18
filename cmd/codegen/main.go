package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	TREE_SITTER_PATH = "./third-party/tree-sitter"
	OUTPUT_DIR       = "."
	targetGOOS       string
	targetGOARCH     string
	keepTemp         bool
)

var rootCmd = &cobra.Command{
	Use:   "tree-sitter-go",
	Short: "Transpile tree-sitter C code to Go using ccgo",
	Long: `A tool to transpile tree-sitter core library and grammars from C to Go.

This tool uses ccgo to convert tree-sitter's C implementation into Go code,
allowing you to use tree-sitter parsers natively in Go without CGO.`,
	RunE: run,
}

func init() {
	rootCmd.Flags().StringVar(&targetGOOS, "goos", runtime.GOOS, "Target GOOS for generated code")
	rootCmd.Flags().StringVar(&targetGOARCH, "goarch", runtime.GOARCH, "Target GOARCH for generated code")
	rootCmd.Flags().BoolVarP(&keepTemp, "keep-temp", "k", false, "Keep temporary files for debugging")

	rootCmd.MarkFlagRequired("tree-sitter")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	// Create transpiler
	transpiler := &Transpiler{
		TreeSitterPath: TREE_SITTER_PATH,
		GOOS:           targetGOOS,
		GOARCH:         targetGOARCH,
		KeepTemp:       keepTemp,
	}

	// Transpile core
	slog.Info("transpiling tree-sitter core", "path", TREE_SITTER_PATH)
	coreOutput := OUTPUT_DIR
	if OUTPUT_DIR != "" {
		coreOutput = filepath.Join(OUTPUT_DIR, "core")
	}
	if err := transpiler.TranspileCore(coreOutput); err != nil {
		return fmt.Errorf("failed to transpile core: %w", err)
	}

	grammars, err := filepath.Glob("third-party/tree-sitter-*")
	if err != nil {
		return err
	}
	// Transpile grammars
	for i, grammarPath := range grammars {
		slog.Info("transpiling grammar", "index", i+1, "total", len(grammars), "path", grammarPath)
		if err := transpiler.TranspileGrammar(grammarPath, OUTPUT_DIR); err != nil {
			return fmt.Errorf("failed to transpile grammar %s: %w", grammarPath, err)
		}
	}

	// Generate project structure if grammars were transpiled
	if OUTPUT_DIR != "" && len(grammars) > 0 {
		slog.Info("generating project structure")
		pg := &ProjectGenerator{
			OutputDir: OUTPUT_DIR,
			Grammars:  grammars,
		}
		if err := pg.Generate(); err != nil {
			return fmt.Errorf("failed to generate project: %w", err)
		}
		slog.Info("project generated successfully", "output", OUTPUT_DIR)
	}

	return nil
}
