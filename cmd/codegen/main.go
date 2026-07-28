package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	treeSitterPath   = "./third-party/tree-sitter"
	defaultOutputDir = "."
	targetGOOS       string
	targetGOARCH     string
	keepTemp         bool
)

var rootCmd = &cobra.Command{
	Use:   "codegen",
	Short: "Transpile tree-sitter C code to Go using ccgo",
	Long: `A tool to transpile tree-sitter core library and grammars from C to Go.

This tool uses ccgo to convert tree-sitter's C implementation into Go code,
allowing you to use tree-sitter parsers natively in Go without CGO.

ccgo preprocesses and compiles C sources itself (via modernc.org/cc and host
CC from NewConfig). CC selects the probe compiler (clang by default; on
Windows, MinGW gcc on PATH is preferred). Host-specific -D flags soften
MinGW/Darwin headers. Optional CFLAGS still affect the CC probe when set.
Unix-style multi-word CC values are parsed with mvdan.cc/sh/v3/shell.Fields.`,
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

	modulesCmd := &cobra.Command{
		Use:   "modules",
		Short: "Write nested go.mod files and go.work for grammar packages",
		Long: `Write grammar/go.mod, grammar/<lang>/go.mod (with local replace
directives), and go.work. Optionally run go work sync / go mod tidy.

Does not transpile C sources.`,
		RunE: runModules,
	}
	modulesCmd.Flags().Bool("tidy", false, "Run go work sync and go mod tidy for each module")
	rootCmd.AddCommand(modulesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runModules(cmd *cobra.Command, args []string) error {
	tidy, err := cmd.Flags().GetBool("tidy")
	if err != nil {
		return fmt.Errorf("failed to get --tidy flag: %w", err)
	}
	slog.Info("writing grammar modules", "dir", defaultOutputDir, "tidy", tidy)
	if err := ensureGrammarModules(defaultOutputDir); err != nil {
		return err
	}
	if tidy {
		slog.Info("tidying workspace modules")
		if err := tidyGrammarModules(defaultOutputDir); err != nil {
			return err
		}
	}
	return nil
}

func run(cmd *cobra.Command, args []string) error {
	slog.Info("compiling for target", "GOOS", targetGOOS, "GOARCH", targetGOARCH, "CC", preprocessorIdentity(targetGOOS, targetGOARCH))
	// Create transpiler
	transpiler := &Transpiler{
		TreeSitterPath: treeSitterPath,
		GOOS:           targetGOOS,
		GOARCH:         targetGOARCH,
		KeepTemp:       keepTemp,
	}

	// Transpile core
	slog.Info("transpiling tree-sitter core", "path", treeSitterPath)
	coreOutput := filepath.Join(defaultOutputDir, "grammar")
	if err := transpiler.TranspileCore(coreOutput); err != nil {
		return fmt.Errorf("failed to transpile core: %w", err)
	}
	if err := writeCoreGoMod(coreOutput); err != nil {
		return fmt.Errorf("failed to write grammar go.mod: %w", err)
	}

	units, err := discoverGrammarUnits("third-party/tree-sitter-*")
	if err != nil {
		return err
	}
	// ONLY_GRAMMAR=name limits the loop (local one-off; CI leaves unset).
	if only := strings.TrimSpace(os.Getenv("ONLY_GRAMMAR")); only != "" {
		var filtered []GrammarUnit
		for _, u := range units {
			if u.Name == only {
				filtered = append(filtered, u)
			}
		}
		if len(filtered) == 0 {
			return fmt.Errorf("ONLY_GRAMMAR=%q matched no discovered unit", only)
		}
		units = filtered
	}
	slog.Info("discovered grammar units", "count", len(units))

	var summaryWriter io.Writer

	if summaryPath := os.Getenv("GITHUB_STEP_SUMMARY"); summaryPath != "" {
		summaryFile, err := os.OpenFile(summaryPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			slog.Warn("failed to open GITHUB_STEP_SUMMARY", "path", summaryPath, "error", err)
		} else {
			summaryWriter = summaryFile
			defer summaryFile.Close()
		}
	} else {
		var buf bytes.Buffer
		defer fmt.Println(buf.String())
		summaryWriter = &buf
	}

	fmt.Fprintf(summaryWriter, "## Grammar Codegen Summary\n\n")
	// Transpile grammars (one unit per language name; monorepos may contribute multiple)
	for i, unit := range units {
		slog.Info("transpiling grammar", "index", i+1, "total", len(units), "grammar", unit.Name, "path", unit.Path, "priority", unit.Priority)
		if err := transpiler.TranspileGrammar(unit.Path, unit.Name, filepath.Join(defaultOutputDir, "grammar")); err != nil {
			slog.Warn("failed to transpile grammar, skipping", "grammar", unit.Name, "path", unit.Path, "error", err)
			fmt.Fprintf(summaryWriter, "- `%s/%s` `%s`: ❌\n", targetGOOS, targetGOARCH, unit.Name)
			continue
		}
		fmt.Fprintf(summaryWriter, "- `%s/%s` `%s`: ✅\n", targetGOOS, targetGOARCH, unit.Name)
	}

	slog.Info("updating languages registry in cmd/parse/languages.go")
	if err := updateLanguagesGo(defaultOutputDir); err != nil {
		return fmt.Errorf("failed to update languages registry: %w", err)
	}

	slog.Info("writing per-grammar go.mod and go.work")
	if err := ensureGrammarModules(defaultOutputDir); err != nil {
		return fmt.Errorf("failed to write grammar modules: %w", err)
	}

	return nil
}

func updateLanguagesGo(outputDir string) error {
	grammarDir := filepath.Join(outputDir, "grammar")
	languages, err := listGrammarLangs(grammarDir)
	if err != nil {
		return err
	}

	moduleName := "github.com/modernc-tree-sitter/ccgo-tree-sitter"

	var sb strings.Builder
	sb.WriteString("// Code generated by codegen (updateLanguagesGo); DO NOT EDIT.\n\n")
	sb.WriteString("package main\n\n")
	sb.WriteString("import (\n")
	for _, lang := range languages {
		fmt.Fprintf(&sb, "\t_ \"%s/grammar/%s\"\n", moduleName, lang)
	}
	sb.WriteString(")\n")

	targetFile := filepath.Join(outputDir, "cmd", "parse", "languages.go")
	return os.WriteFile(targetFile, []byte(sb.String()), 0644)
}
