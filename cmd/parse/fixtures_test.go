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
// Sources for every generated language live in fixtureSources.
// Set UPDATE_GOLDENS=1 to rewrite sources (from fixtureSources) and goldens.

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

	langs := grammar.List()
	if len(langs) == 0 {
		t.Fatal("no languages registered (import grammar packages in languages.go)")
	}

	// Every registered language must have a curated source snippet.
	var missingSnippets []string
	for _, language := range langs {
		if _, ok := fixtureSources[language]; !ok {
			missingSnippets = append(missingSnippets, language)
		}
	}
	if len(missingSnippets) > 0 {
		t.Fatalf("fixtureSources missing languages: %s", strings.Join(missingSnippets, ", "))
	}
	// No stale snippets for unregistered languages.
	for language := range fixtureSources {
		if _, ok := grammar.Get(language); !ok {
			t.Fatalf("fixtureSources has %q but language is not registered", language)
		}
	}

	fixtureCount := 0
	var parseErrors []string

	for _, language := range langs {
		snip := fixtureSources[language]
		langDir := filepath.Join(testdataRoot, language)
		srcPath := filepath.Join(langDir, snip.File)
		goldenPath := srcPath + goldenSuffix

		lang, ok := grammar.Get(language)
		if !ok {
			t.Fatalf("language %q not registered", language)
		}

		if update {
			if err := os.MkdirAll(langDir, 0o755); err != nil {
				t.Fatalf("mkdir %s: %v", langDir, err)
			}
			if err := os.WriteFile(srcPath, []byte(snip.Source), 0o644); err != nil {
				t.Fatalf("write source %s: %v", srcPath, err)
			}
		} else if _, err := os.Stat(srcPath); err != nil {
			t.Fatalf("missing source %s (set UPDATE_GOLDENS=1 to materialise from fixtureSources)", srcPath)
		}

		fixtureCount++
		t.Run(language+"/"+snip.File, func(t *testing.T) {
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
				msg := language + "/" + snip.File + ":\n" + rootNode.PrintTree()
				if update {
					// Still record failure when updating — goldens for broken trees are useless.
					parseErrors = append(parseErrors, msg)
				}
				t.Fatalf("parse errors:\n%s", rootNode.PrintTree())
			}

			parseRoot := grammar.BuildParseNode(rootNode, source, "")
			if parseRoot == nil {
				t.Fatal("BuildParseNode returned nil")
			}

			gotOut := grammar.ParseOutput{
				Language: language,
				File:     snip.File,
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

	if fixtureCount == 0 {
		t.Fatal("no language fixtures exercised")
	}
	if len(parseErrors) > 0 && update {
		t.Fatalf("%d fixtures still parse with errors — fix fixtureSources:\n%s",
			len(parseErrors), strings.Join(parseErrors, "\n---\n"))
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
