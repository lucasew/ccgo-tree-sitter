package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

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

func env(key, defaultValue string) string {
	if ret := os.Getenv(key); ret != "" {
		return ret
	}
	return defaultValue
}

func init() {
	rootCmd.Flags().StringVar(&targetGOOS, "goos", env("TARGET_GOOS", runtime.GOOS), "Target GOOS for generated code")
	rootCmd.Flags().StringVar(&targetGOARCH, "goarch", env("TARGET_GOARCH", runtime.GOARCH), "Target GOARCH for generated code")
	rootCmd.Flags().BoolVarP(&keepTemp, "keep-temp", "k", false, "Keep temporary files for debugging")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	slog.Info("compiling for target", "GOOS", targetGOOS, "GOARCH", targetGOARCH)
	// Create transpiler
	transpiler := &Transpiler{
		TreeSitterPath: TREE_SITTER_PATH,
		GOOS:           targetGOOS,
		GOARCH:         targetGOARCH,
		KeepTemp:       keepTemp,
	}

	// Transpile core
	slog.Info("transpiling tree-sitter core", "path", TREE_SITTER_PATH)
	coreOutput := filepath.Join(OUTPUT_DIR, "grammar")
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
		if err := transpiler.TranspileGrammar(grammarPath, OUTPUT_DIR+"/grammar"); err != nil {
			return fmt.Errorf("failed to transpile grammar %s: %w", grammarPath, err)
		}
	}

	slog.Info("updating languages registry in cmd/parse/languages.go")
	if err := updateLanguagesGo(OUTPUT_DIR); err != nil {
		return fmt.Errorf("failed to update languages registry: %w", err)
	}

	return nil
}

func updateLanguagesGo(outputDir string) error {
	grammarDir := filepath.Join(outputDir, "grammar")
	entries, err := os.ReadDir(grammarDir)
	if err != nil {
		return err
	}

	var languages []string
	for _, entry := range entries {
		if entry.IsDir() {
			languages = append(languages, entry.Name())
		}
	}
	sort.Strings(languages)

	moduleName := "github.com/lucasew/ccgo-tree-sitter"

	var sb strings.Builder
	sb.WriteString("package main\n\n")
	sb.WriteString("import (\n")
	for _, lang := range languages {
		sb.WriteString(fmt.Sprintf("\t_ \"%s/grammar/%s\"\n", moduleName, lang))
	}
	sb.WriteString(")\n")

	targetFile := filepath.Join(outputDir, "cmd", "parse", "languages.go")
	return os.WriteFile(targetFile, []byte(sb.String()), 0644)
}
