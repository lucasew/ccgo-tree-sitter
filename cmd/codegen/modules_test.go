package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteLangGoMod(t *testing.T) {
	libcVer, err := currentLibcVersion()
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	langDir := filepath.Join(dir, "python")
	if err := os.MkdirAll(langDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(langDir, "api.go"), []byte("package grammar_python\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := writeLangGoMod(dir, "python"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(langDir, "go.mod"))
	if err != nil {
		t.Fatal(err)
	}
	s := string(data)
	for _, want := range []string{
		"module github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/python",
		"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar v0.0.0",
		"modernc.org/libc " + libcVer,
		"replace github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar => ../",
		"replace modernc.org/libc => " + libcReplacePath + " " + libcReplaceVer,
	} {
		if !strings.Contains(s, want) {
			t.Errorf("go.mod missing %q\n%s", want, s)
		}
	}
}

func TestWriteCoreGoMod(t *testing.T) {
	libcVer, err := currentLibcVersion()
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	if err := writeCoreGoMod(dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		t.Fatal(err)
	}
	s := string(data)
	for _, want := range []string{
		"module github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar",
		"require modernc.org/libc " + libcVer,
		"replace modernc.org/libc => " + libcReplacePath + " " + libcReplaceVer,
	} {
		if !strings.Contains(s, want) {
			t.Errorf("go.mod missing %q\n%s", want, s)
		}
	}
}

func TestEnsureGrammarModules(t *testing.T) {
	root := t.TempDir()
	grammarDir := filepath.Join(root, "grammar")
	langDir := filepath.Join(grammarDir, "c")
	if err := os.MkdirAll(langDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(langDir, "api.go"), []byte("package grammar_c\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := ensureGrammarModules(root); err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{
		filepath.Join(grammarDir, "go.mod"),
		filepath.Join(langDir, "go.mod"),
		filepath.Join(root, "go.work"),
	} {
		if _, err := os.Stat(p); err != nil {
			t.Errorf("missing %s: %v", p, err)
		}
	}
	work, _ := os.ReadFile(filepath.Join(root, "go.work"))
	if !strings.Contains(string(work), "./grammar/c") {
		t.Fatalf("go.work: %s", work)
	}
}
