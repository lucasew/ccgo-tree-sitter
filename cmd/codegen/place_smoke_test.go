package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func TestTranspilePlacedJSON(t *testing.T) {
	root := findRepoRoot(t)
	parser := filepath.Join(root, "third-party/tree-sitter-json/src/parser.c")
	if _, err := os.Stat(parser); err != nil {
		t.Skip("placed json not available:", err)
	}
	if _, err := os.Stat(filepath.Join(root, "third-party/tree-sitter-json/bindings")); !os.IsNotExist(err) {
		t.Fatalf("expected clean place without bindings, err=%v", err)
	}
	// ccgo needs a host C compiler probe (clang/gcc).
	if _, err := exec.LookPath("clang"); err != nil {
		if _, err := exec.LookPath("gcc"); err != nil {
			t.Skip("no clang/gcc on PATH for ccgo probe")
		}
	}
	tsPath, err := resolveTreeSitterPath()
	if err != nil {
		t.Skip("tree-sitter source not available:", err)
	}
	out := t.TempDir()
	tr := &Transpiler{
		TreeSitterPath: tsPath,
		GOOS:           runtime.GOOS,
		GOARCH:         runtime.GOARCH,
	}
	if err := tr.TranspileGrammar(filepath.Join(root, "third-party/tree-sitter-json"), "json", filepath.Join(out, "grammar")); err != nil {
		t.Fatal(err)
	}
	got := filepath.Join(out, "grammar", "json", "grammar-"+runtime.GOOS+"-"+runtime.GOARCH+".go")
	if _, err := os.Stat(got); err != nil {
		t.Fatal("missing output", got, err)
	}
}
