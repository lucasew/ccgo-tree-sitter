package main

import (
	"bufio"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

//go:embed freestanding/*
var freestandingFS embed.FS

// GrammarKind classifies original tree-sitter grammar units for codegen.
type GrammarKind int

const (
	// GrammarPure is parser.c only — freestanding stubs, no host libc headers.
	GrammarPure GrammarKind = iota
	// GrammarScanner is parser.c + scanner.c — freestanding stubs covering
	// malloc/string/wctype used by external scanners.
	GrammarScanner
)

func (k GrammarKind) String() string {
	switch k {
	case GrammarPure:
		return "pure"
	case GrammarScanner:
		return "scanner"
	default:
		return fmt.Sprintf("GrammarKind(%d)", int(k))
	}
}

// classifyGrammarUnit inspects original files under grammarPath and fills Kind,
// ScannerC, ExtraIncs, InputHash on a GrammarUnit (see discover.go).
func classifyGrammarUnit(grammarPath, name string) (GrammarUnit, error) {
	u := GrammarUnit{
		Name:    name,
		Path:    grammarPath,
		Kind:    GrammarPure,
		ParserC: filepath.Join(grammarPath, "src", "parser.c"),
	}
	if _, err := os.Stat(u.ParserC); err != nil {
		return u, fmt.Errorf("parser.c: %w", err)
	}
	scanner := filepath.Join(grammarPath, "src", "scanner.c")
	if st, err := os.Stat(scanner); err == nil && !st.IsDir() {
		u.Kind = GrammarScanner
		u.ScannerC = scanner
	} else if err != nil && !os.IsNotExist(err) {
		return u, err
	}

	// Monorepo: sibling common/ or shared headers next to the unit.
	for _, rel := range []string{
		filepath.Join("..", "common"),
		"common",
		filepath.Join("..", "src"),
	} {
		p := filepath.Join(grammarPath, rel)
		if st, err := os.Stat(p); err == nil && st.IsDir() {
			u.ExtraIncs = append(u.ExtraIncs, p)
		}
	}
	// Parent of unit (e.g. tree-sitter-typescript/) for monorepo includes.
	u.ExtraIncs = append(u.ExtraIncs, filepath.Dir(grammarPath))

	h, err := hashGrammarInputs(u)
	if err != nil {
		return u, err
	}
	u.InputHash = h
	return u, nil
}

func hashGrammarInputs(u GrammarUnit) (string, error) {
	h := sha256.New()
	_, _ = fmt.Fprintf(h, "kind=%s\n", u.Kind)
	_, _ = fmt.Fprintf(h, "freestanding=1\n")
	// Bust skip when freestanding stubs change.
	entries, err := freestandingFS.ReadDir("freestanding")
	if err != nil {
		return "", err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		data, err := freestandingFS.ReadFile("freestanding/" + e.Name())
		if err != nil {
			return "", err
		}
		sum := sha256.Sum256(data)
		_, _ = fmt.Fprintf(h, "fs %s %x\n", e.Name(), sum)
	}
	paths := []string{u.ParserC}
	if u.ScannerC != "" {
		paths = append(paths, u.ScannerC)
	}
	tsDir := filepath.Join(u.Path, "src", "tree_sitter")
	_ = filepath.WalkDir(tsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if strings.HasSuffix(path, ".h") {
			paths = append(paths, path)
		}
		return nil
	})
	sort.Strings(paths)
	for _, p := range paths {
		sum, err := fileSHA256(p)
		if err != nil {
			return "", err
		}
		_, _ = fmt.Fprintf(h, "%s %s\n", sum, filepath.ToSlash(p))
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func fileSHA256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// grammarOutputUpToDate reports whether the classic output already matches inputs.
func grammarOutputUpToDate(outputFile, wantHash string) bool {
	data, err := os.ReadFile(outputFile)
	if err != nil {
		return false
	}
	const prefix = "// ts-grammar-hash: "
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, prefix) {
			return strings.TrimSpace(strings.TrimPrefix(line, prefix)) == wantHash
		}
	}
	return false
}

func withGrammarHashComment(goCode, hash string) string {
	const marker = "// ts-grammar-hash: "
	if strings.Contains(goCode, marker) {
		re := regexp.MustCompile(`(?m)^// ts-grammar-hash: .*\n`)
		return re.ReplaceAllString(goCode, marker+hash+"\n")
	}
	pkg := regexp.MustCompile(`(?m)^package \w+\n`)
	loc := pkg.FindStringIndex(goCode)
	if loc == nil {
		return marker + hash + "\n" + goCode
	}
	return goCode[:loc[1]] + marker + hash + "\n" + goCode[loc[1]:]
}

var systemIncludeRe = regexp.MustCompile(`^(\s*#\s*include\s*)<([A-Za-z0-9_./-]+)>`)

var freestandingHeaders = map[string]struct{}{
	"stdint.h":   {},
	"stdbool.h":  {},
	"stddef.h":   {},
	"stdlib.h":   {},
	"string.h":   {},
	"assert.h":   {},
	"ctype.h":    {},
	"wctype.h":   {},
	"stdio.h":    {},
	"limits.h":   {},
	"wchar.h":    {},
	"inttypes.h": {},
}

// rewriteSystemIncludes redirects known angle-bracket includes to freestanding quotes.
func rewriteSystemIncludes(src string) string {
	var b strings.Builder
	b.Grow(len(src) + 64)
	for _, line := range strings.Split(src, "\n") {
		if m := systemIncludeRe.FindStringSubmatch(line); m != nil {
			hdr := m[2]
			base := filepath.Base(hdr)
			if _, ok := freestandingHeaders[base]; ok {
				b.WriteString(m[1])
				b.WriteByte('"')
				b.WriteString(base)
				b.WriteByte('"')
				b.WriteByte('\n')
				continue
			}
		}
		b.WriteString(line)
		b.WriteByte('\n')
	}
	return b.String()
}

// writeFreestandingHeaders copies embedded stubs into destDir.
func writeFreestandingHeaders(destDir string) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}
	entries, err := freestandingFS.ReadDir("freestanding")
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		data, err := freestandingFS.ReadFile("freestanding/" + e.Name())
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(destDir, e.Name()), data, 0644); err != nil {
			return err
		}
	}
	return nil
}

// stageGrammarSources writes rewritten C (+ tree_sitter headers) into workDir
// so system includes resolve to freestanding stubs.
// Returns path to the main .c file for ccgo.
func stageGrammarSources(u GrammarUnit, workDir string) (mainC string, err error) {
	fsDir := filepath.Join(workDir, "fs")
	if err := writeFreestandingHeaders(fsDir); err != nil {
		return "", err
	}

	// Stage tree_sitter headers with rewritten includes.
	srcTS := filepath.Join(u.Path, "src", "tree_sitter")
	dstTS := filepath.Join(workDir, "tree_sitter")
	if err := copyDirRewritten(srcTS, dstTS); err != nil && !os.IsNotExist(err) {
		return "", fmt.Errorf("stage tree_sitter headers: %w", err)
	}

	// Stage other headers under src/ that scanners may include (array.h lives in tree_sitter).
	// Also rewrite any loose .h in src/ next to parser.
	srcDir := filepath.Join(u.Path, "src")
	if err := filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			// skip tree_sitter (already staged) and node_modules-like
			if d.Name() == "tree_sitter" && path != srcDir {
				return fs.SkipDir
			}
			return nil
		}
		rel, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		// Only headers and the TU sources we need.
		switch {
		case strings.HasSuffix(path, ".h"), strings.HasSuffix(path, ".c"), strings.HasSuffix(path, ".inc"):
		default:
			return nil
		}
		// Skip the main parser/scanner — handled below as combined unit.
		base := filepath.Base(path)
		if base == "parser.c" || base == "scanner.c" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		out := filepath.Join(workDir, rel)
		if err := os.MkdirAll(filepath.Dir(out), 0755); err != nil {
			return err
		}
		return os.WriteFile(out, []byte(rewriteSystemIncludes(string(data))), 0644)
	}); err != nil {
		return "", err
	}

	// Combined TU: scanner first (if any), then parser — same as classic path.
	var parts []string
	if u.ScannerC != "" {
		parts = append(parts, u.ScannerC)
	}
	parts = append(parts, u.ParserC)

	mainC = filepath.Join(workDir, "unit.c")
	var combined strings.Builder
	combined.WriteString("/* staged freestanding grammar unit — generated by codegen, not checked in */\n")
	for _, p := range parts {
		data, err := os.ReadFile(p)
		if err != nil {
			return "", err
		}
		combined.WriteString(rewriteSystemIncludes(string(data)))
		combined.WriteString("\n\n")
	}
	if err := os.WriteFile(mainC, []byte(combined.String()), 0644); err != nil {
		return "", err
	}
	return mainC, nil
}

func copyDirRewritten(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		out := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(out, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(out, []byte(rewriteSystemIncludes(string(data))), 0644)
	})
}

// --- output slimming (drop host preprocessor consts from ccgo dumps) ---

// Grammar const names that come from tree-sitter generated parser.c / parser.h.
var grammarConstKeep = map[string]struct{}{
	"ALIAS_COUNT":                           {},
	"EXTERNAL_TOKEN_COUNT":                  {},
	"FIELD_COUNT":                           {},
	"LANGUAGE_VERSION":                      {},
	"LARGE_STATE_COUNT":                     {},
	"MAX_ALIAS_SEQUENCE_LENGTH":             {},
	"PRODUCTION_ID_COUNT":                   {},
	"STATE_COUNT":                           {},
	"SYMBOL_COUNT":                          {},
	"TOKEN_COUNT":                           {},
	"TREE_SITTER_SERIALIZATION_BUFFER_SIZE": {},
	"TSParseActionTypeShift":                {},
	"TSParseActionTypeReduce":               {},
	"TSParseActionTypeAccept":               {},
	"TSParseActionTypeRecover":              {},
}

var (
	constDeclRe = regexp.MustCompile(`^const\s+(\w+)\s*=`)
	// Host / ABI macros and CRT limits — not part of the language tables.
	platformConstNameRe = regexp.MustCompile(
		`^(__|_|INT|UINT|PTRDIFF|SIZE_MAX|SIG_|WINT_|WNOHANG|WUNTRACED|RAND_MAX|EXIT_|BYTE_ORDER|BIG_ENDIAN|LITTLE_ENDIAN|NULL$)`,
	)
)

// keepGrammarConst reports whether a top-level const belongs in the grammar binding.
func keepGrammarConst(name string) bool {
	if _, ok := grammarConstKeep[name]; ok {
		return true
	}
	// Symbol / field enums from parser.c.
	if strings.HasPrefix(name, "anon_sym_") ||
		strings.HasPrefix(name, "sym_") ||
		strings.HasPrefix(name, "aux_sym_") ||
		strings.HasPrefix(name, "field_") ||
		strings.HasPrefix(name, "alias_") {
		return true
	}
	if platformConstNameRe.MatchString(name) {
		return false
	}
	// Unknown ALL_CAPS from headers — drop.
	if name == strings.ToUpper(name) && len(name) > 1 {
		return false
	}
	return true
}

// slimGeneratedGrammar drops host preprocessor consts from ccgo output while
// keeping grammar tables, types, and functions.
func slimGeneratedGrammar(goCode string) string {
	var b strings.Builder
	b.Grow(len(goCode))
	sc := bufio.NewScanner(strings.NewReader(goCode))
	sc.Buffer(make([]byte, 0, 64*1024), 16*1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if m := constDeclRe.FindStringSubmatch(line); m != nil {
			if !keepGrammarConst(m[1]) {
				continue
			}
		}
		b.WriteString(line)
		b.WriteByte('\n')
	}
	if err := sc.Err(); err != nil {
		return goCode
	}
	return b.String()
}
