package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

var structPaddingPattern = regexp.MustCompile(`&struct \{ _ \[`)

const goldenSuffix = ".golden.json"

// Layout:
//
//	testdata/<language>/<file.ext>
//	testdata/<language>/<file.ext>.golden.json
//
// Golden is grammar.ParseOutput JSON (language, file, root ParseNode tree).
// Set UPDATE_GOLDENS=1 to rewrite missing or outdated goldens.

func TestGeneratedCoreHasNoStructPadding(t *testing.T) {
	root := repoRoot(t)
	grammarDir := filepath.Join(root, "grammar")

	entries, err := os.ReadDir(grammarDir)
	if err != nil {
		t.Fatalf("failed to read grammar directory: %v", err)
	}

	var coreFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, "core-") && strings.HasSuffix(name, ".go") {
			coreFiles = append(coreFiles, filepath.Join(grammarDir, name))
		}
	}
	if len(coreFiles) == 0 {
		t.Fatalf("no generated core files found in %s", grammarDir)
	}

	for _, coreFile := range coreFiles {
		content, err := os.ReadFile(coreFile)
		if err != nil {
			t.Fatalf("failed to read %s: %v", coreFile, err)
		}
		if structPaddingPattern.Match(content) {
			t.Fatalf("found struct padding pattern in %s", coreFile)
		}
	}
}

func TestLanguageFixtures(t *testing.T) {
	root := repoRoot(t)
	testdataRoot := filepath.Join(root, "testdata")
	update := os.Getenv("UPDATE_GOLDENS") == "1"

	langDirs, err := os.ReadDir(testdataRoot)
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}

	fixtureCount := 0
	for _, langEnt := range langDirs {
		if !langEnt.IsDir() || strings.HasPrefix(langEnt.Name(), ".") {
			continue
		}
		language := langEnt.Name()
		langDir := filepath.Join(testdataRoot, language)

		lang, ok := grammar.Get(language)
		if !ok {
			t.Fatalf("language %q has testdata/ but is not registered; supported: %s", language, grammar.SupportedLanguages())
		}

		entries, err := os.ReadDir(langDir)
		if err != nil {
			t.Fatalf("read %s: %v", langDir, err)
		}
		var sources []string
		for _, e := range entries {
			if e.IsDir() || strings.HasPrefix(e.Name(), ".") {
				continue
			}
			name := e.Name()
			if strings.HasSuffix(name, goldenSuffix) {
				continue
			}
			sources = append(sources, name)
		}
		if len(sources) == 0 {
			t.Fatalf("no source fixtures under %s (expected file.ext + file.ext%s)", langDir, goldenSuffix)
		}

		for _, base := range sources {
			fixtureCount++
			srcPath := filepath.Join(langDir, base)
			goldenPath := srcPath + goldenSuffix
			t.Run(language+"/"+base, func(t *testing.T) {
				source, err := os.ReadFile(srcPath)
				if err != nil {
					t.Fatalf("read fixture: %v", err)
				}

				parser := grammar.NewParser()
				if !parser.SetLanguage(lang) {
					t.Fatalf("SetLanguage %q failed", language)
				}
				tree := parser.ParseBytes(source)
				rootNode := tree.RootNode()
				if rootNode.IsNull() {
					t.Fatal("null root")
				}
				if rootNode.HasError() {
					t.Fatalf("parse errors:\n%s", rootNode.PrintTree())
				}

				parseRoot := grammar.BuildParseNode(rootNode, source, "")
				if parseRoot == nil {
					t.Fatal("BuildParseNode returned nil")
				}

				gotOut := grammar.ParseOutput{
					Language: language,
					File:     base,
					Root:     parseRoot,
				}
				gotJSON, err := json.MarshalIndent(gotOut, "", "  ")
				if err != nil {
					t.Fatalf("marshal golden candidate: %v", err)
				}
				gotJSON = append(gotJSON, '\n')

				want, err := os.ReadFile(goldenPath)
				if err != nil {
					if os.IsNotExist(err) && update {
						if err := os.WriteFile(goldenPath, gotJSON, 0o644); err != nil {
							t.Fatalf("write golden %s: %v", goldenPath, err)
						}
						t.Logf("wrote %s", goldenPath)
						return
					}
					t.Fatalf("missing golden %s (set UPDATE_GOLDENS=1 to create)", goldenPath)
				}

				if update && !bytes.Equal(want, gotJSON) {
					if err := os.WriteFile(goldenPath, gotJSON, 0o644); err != nil {
						t.Fatalf("update golden %s: %v", goldenPath, err)
					}
					t.Logf("updated %s", goldenPath)
					return
				}

				if !bytes.Equal(want, gotJSON) {
					t.Fatalf("golden mismatch for %s\n--- want ---\n%s\n--- got ---\n%s\n(set UPDATE_GOLDENS=1 to rewrite)",
						goldenPath, want, gotJSON)
				}
			})
		}
	}

	if fixtureCount == 0 {
		t.Fatalf("no fixtures under %s/<language>/<file.ext>", testdataRoot)
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatalf("could not locate repository root from %s", dir)
		}
		dir = parent
	}
}
