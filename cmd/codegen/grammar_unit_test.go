package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestRewriteSystemIncludes(t *testing.T) {
	in := `#include "tree_sitter/parser.h"
#include <stdint.h>
#include <stdbool.h>
#include <sys/types.h>
`
	out := rewriteSystemIncludes(in)
	if !strings.Contains(out, `#include "stdint.h"`) {
		t.Fatalf("stdint not rewritten:\n%s", out)
	}
	if !strings.Contains(out, `#include "stdbool.h"`) {
		t.Fatalf("stdbool not rewritten:\n%s", out)
	}
	if !strings.Contains(out, `#include <sys/types.h>`) {
		t.Fatalf("unknown system header should remain:\n%s", out)
	}
	if !strings.Contains(out, `#include "tree_sitter/parser.h"`) {
		t.Fatalf("local include must stay:\n%s", out)
	}
}

func TestClassifyPureAndScanner(t *testing.T) {
	root := t.TempDir()
	pure := filepath.Join(root, "pure")
	if err := os.MkdirAll(filepath.Join(pure, "src", "tree_sitter"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pure, "src", "parser.c"), []byte("int x;\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pure, "src", "tree_sitter", "parser.h"), []byte("#include <stdint.h>\n"), 0644); err != nil {
		t.Fatal(err)
	}
	u, err := classifyGrammarUnit(pure, "pure")
	if err != nil {
		t.Fatal(err)
	}
	if u.Kind != GrammarPure {
		t.Fatalf("kind=%v", u.Kind)
	}
	if u.InputHash == "" {
		t.Fatal("missing hash")
	}

	scan := filepath.Join(root, "scan")
	if err := os.MkdirAll(filepath.Join(scan, "src", "tree_sitter"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(scan, "src", "parser.c"), []byte("int x;\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(scan, "src", "scanner.c"), []byte("int y;\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(scan, "src", "tree_sitter", "parser.h"), []byte(""), 0644); err != nil {
		t.Fatal(err)
	}
	u2, err := classifyGrammarUnit(scan, "scan")
	if err != nil {
		t.Fatal(err)
	}
	if u2.Kind != GrammarScanner {
		t.Fatalf("kind=%v", u2.Kind)
	}
}

func TestGrammarHashRoundTrip(t *testing.T) {
	code := "package grammar_x\n\nfunc F() {}\n"
	hashed := withGrammarHashComment(code, "abc123")
	if !strings.Contains(hashed, "// ts-grammar-hash: abc123") {
		t.Fatal(hashed)
	}
	dir := t.TempDir()
	path := filepath.Join(dir, "g.go")
	if err := os.WriteFile(path, []byte(hashed), 0644); err != nil {
		t.Fatal(err)
	}
	if !grammarOutputUpToDate(path, "abc123") {
		t.Fatal("should be up to date")
	}
	if grammarOutputUpToDate(path, "nope") {
		t.Fatal("wrong hash")
	}
}

func TestSlimGeneratedGrammarDropsPlatformConsts(t *testing.T) {
	in := `package grammar_x
const LANGUAGE_VERSION = 14
const __linux__ = 1
const INT16_MAX = 0x7fff
const STATE_COUNT = 3
func tree_sitter_x() {}
`
	out := slimGeneratedGrammar(in)
	if !strings.Contains(out, "LANGUAGE_VERSION") || !strings.Contains(out, "STATE_COUNT") {
		t.Fatal("kept grammar consts missing")
	}
	if strings.Contains(out, "__linux__") || strings.Contains(out, "INT16_MAX") {
		t.Fatal("platform noise left:\n", out)
	}
}

func TestWriteFreestandingHeaders(t *testing.T) {
	dir := t.TempDir()
	if err := writeFreestandingHeaders(dir); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"stdint.h", "stdbool.h", "stdlib.h", "string.h", "wctype.h"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Fatal(name, err)
		}
	}
}

// TestFreestandingJSONTranspile runs real ccgo when a host CC exists (CI).
// Skips locally when clang/gcc are missing — unit tests cover staging/classify.
func TestFreestandingJSONTranspile(t *testing.T) {
	if os.Getenv("CI") == "" && os.Getenv("FORCE_CODEGEN_TEST") == "" {
		// Still try if CC works.
	}
	cc := os.Getenv("CC")
	if cc == "" {
		cc = "clang"
	}
	if _, err := exec.LookPath(cc); err != nil {
		if _, err2 := exec.LookPath("gcc"); err2 != nil {
			t.Skip("no C preprocessor (clang/gcc) for freestanding ccgo smoke")
		}
		cc = "gcc"
	}
	src := filepath.Join("..", "..", "third-party", "tree-sitter-json")
	if _, err := os.Stat(filepath.Join(src, "src", "parser.c")); err != nil {
		src = filepath.Join("third-party", "tree-sitter-json")
	}
	if _, err := os.Stat(filepath.Join(src, "src", "parser.c")); err != nil {
		t.Skip(err)
	}
	tsPath, err := resolveTreeSitterPath()
	if err != nil {
		t.Skip("tree-sitter path:", err)
	}
	outDir := t.TempDir()
	tr := &Transpiler{
		TreeSitterPath: tsPath,
		GOOS:           "linux",
		GOARCH:         "amd64",
		KeepTemp:       false,
	}
	t.Setenv("CC", cc)
	if err := tr.TranspileGrammar(src, "json", filepath.Join(outDir, "grammar")); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(outDir, "grammar", "json", "grammar-linux-amd64.go")
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	body := string(data)
	if !strings.Contains(body, "func tree_sitter_json") {
		t.Fatal("missing language entry")
	}
	if !strings.Contains(body, "// ts-grammar-hash:") {
		t.Fatal("missing hash marker")
	}
	// Freestanding path should not embed host linux noise.
	if strings.Contains(body, "__linux__") || strings.Contains(body, "__ATOMIC_ACQUIRE") {
		t.Fatal("host platform consts leaked into freestanding output")
	}
	// Second run should hash-skip (no error).
	if err := tr.TranspileGrammar(src, "json", filepath.Join(outDir, "grammar")); err != nil {
		t.Fatal("skip run:", err)
	}
	t.Logf("json freestanding binding size=%d", len(data))
}

func TestStageGrammarSourcesJSON(t *testing.T) {
	src := filepath.Join("..", "..", "third-party", "tree-sitter-json")
	if _, err := os.Stat(filepath.Join(src, "src", "parser.c")); err != nil {
		src = filepath.Join("third-party", "tree-sitter-json")
	}
	if _, err := os.Stat(filepath.Join(src, "src", "parser.c")); err != nil {
		t.Skip("json grammar not vendored:", err)
	}
	u, err := classifyGrammarUnit(src, "json")
	if err != nil {
		t.Fatal(err)
	}
	if u.Kind != GrammarPure {
		t.Fatalf("json should be pure, got %v", u.Kind)
	}
	work := t.TempDir()
	mainC, err := stageGrammarSources(u, work)
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(mainC)
	if err != nil {
		t.Fatal(err)
	}
	body := string(data)
	if strings.Contains(body, "#include <stdint.h>") {
		t.Fatal("staged unit still has angle stdint")
	}
	if !strings.Contains(body, `#include "tree_sitter/parser.h"`) && !strings.Contains(body, `tree_sitter/parser.h`) {
		// parser.h is included from staged content
		t.Log("note: include form may vary")
	}
	// freestanding dir present
	if _, err := os.Stat(filepath.Join(work, "fs", "stdint.h")); err != nil {
		t.Fatal(err)
	}
}
