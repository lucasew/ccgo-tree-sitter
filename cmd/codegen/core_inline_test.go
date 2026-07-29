package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Verifies address-taken header static inlines (ts_decode_*) get out-of-line
// Go funcs so __ccgo_fp(name) resolves, without requiring -fno-inline.
func TestCoreEmitsUTFDecodeFuncs(t *testing.T) {
	if testing.Short() {
		t.Skip("core transpile is slow")
	}
	if _, err := exec.LookPath("clang"); err != nil {
		if _, err2 := exec.LookPath("gcc"); err2 != nil {
			t.Skip("no C compiler on PATH")
		}
	}
	dir := t.TempDir()
	tsPath, err := resolveTreeSitterPath()
	if err != nil {
		t.Skip("tree-sitter source not available:", err)
	}
	tr := &Transpiler{
		TreeSitterPath: tsPath,
		GOOS:           runtime.GOOS,
		GOARCH:         runtime.GOARCH,
		KeepTemp:       false,
	}
	if err := tr.TranspileCore(dir); err != nil {
		t.Fatalf("TranspileCore: %v", err)
	}
	core := filepath.Join(dir, "core-"+tr.GOOS+"-"+tr.GOARCH+".go")
	data, err := os.ReadFile(core)
	if err != nil {
		t.Fatal(err)
	}
	s := string(data)
	for _, name := range []string{"ts_decode_utf8", "ts_decode_utf16_le", "ts_decode_utf16_be"} {
		if !strings.Contains(s, "func "+name+"(") {
			t.Errorf("missing func %s in %s", name, core)
		}
		if !strings.Contains(s, "__ccgo_fp("+name+")") {
			t.Errorf("missing __ccgo_fp(%s) use site in %s", name, core)
		}
	}
}
