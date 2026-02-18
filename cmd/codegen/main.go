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
	treeSitterPath string
	grammars       []string
	targetGOOS     string
	targetGOARCH   string
	outputDir      string
	keepTemp       bool
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
	rootCmd.Flags().StringVarP(&treeSitterPath, "tree-sitter", "t", "", "Path to tree-sitter source directory (required)")
	rootCmd.Flags().StringArrayVarP(&grammars, "grammar", "g", []string{}, "Path to grammar source directory (can be specified multiple times)")
	rootCmd.Flags().StringVar(&targetGOOS, "goos", runtime.GOOS, "Target GOOS for generated code")
	rootCmd.Flags().StringVar(&targetGOARCH, "goarch", runtime.GOARCH, "Target GOARCH for generated code")
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory (default: stdout)")
	rootCmd.Flags().BoolVarP(&keepTemp, "keep-temp", "k", false, "Keep temporary files for debugging")

	rootCmd.MarkFlagRequired("tree-sitter")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	if treeSitterPath == "" {
		return fmt.Errorf("--tree-sitter flag is required")
	}

	// Create transpiler
	transpiler := &Transpiler{
		TreeSitterPath: treeSitterPath,
		GOOS:           targetGOOS,
		GOARCH:         targetGOARCH,
		KeepTemp:       keepTemp,
	}

	// Transpile core
	slog.Info("transpiling tree-sitter core", "path", treeSitterPath)
	coreOutput := outputDir
	if outputDir != "" {
		coreOutput = filepath.Join(outputDir, "core")
	}
	if err := transpiler.TranspileCore(coreOutput); err != nil {
		return fmt.Errorf("failed to transpile core: %w", err)
	}

	// Transpile grammars
	for i, grammarPath := range grammars {
		slog.Info("transpiling grammar", "index", i+1, "total", len(grammars), "path", grammarPath)
		if err := transpiler.TranspileGrammar(grammarPath, outputDir); err != nil {
			return fmt.Errorf("failed to transpile grammar %s: %w", grammarPath, err)
		}
	}

	// Generate project structure if grammars were transpiled
	if outputDir != "" && len(grammars) > 0 {
		slog.Info("generating project structure")
		pg := &ProjectGenerator{
			OutputDir: outputDir,
			Grammars:  grammars,
		}
		if err := pg.Generate(); err != nil {
			return fmt.Errorf("failed to generate project: %w", err)
		}
		slog.Info("project generated successfully", "output", outputDir)
		slog.Info("build with: cd %s && go build ./cmd/parse", outputDir)
	}

	return nil
}
