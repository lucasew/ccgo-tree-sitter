package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"

	"modernc.org/ccgo/v4/lib"
	"mvdan.cc/sh/v3/shell"
)

// Transpiler handles C to Go transpilation using ccgo.
// It orchestrates the process of transpiling tree-sitter C code into Go,
// setting up isolated execution environments to prevent ccgo state pollution.
type Transpiler struct {
	// TreeSitterPath is the root directory of the tree-sitter source.
	TreeSitterPath string
	// GOOS is the target operating system for generated Go code.
	GOOS string
	// GOARCH is the target architecture for generated Go code.
	GOARCH string
	// KeepTemp specifies whether to leave intermediate C and go files intact for debugging.
	KeepTemp bool
}

const (
	deterministicTempRoot = "ccgo-tree-sitter-gen"
	defaultPreprocessor   = "clang"
)

// extensionFloatTypedefRe matches glibc interchange-type typedefs that clang
// emits from bits/floatn*.h. GCC omits them (builtins); ccgo cannot type-check
// the clang forms (__float128, duplicate specifiers, __mode__ (__TC__), …).
var extensionFloatTypedefRe = regexp.MustCompile(
	`(?m)^typedef\s+[^;]*\b(_Float(16|32|32x|64|64x|128)|__float128|__cfloat128)\b[^;]*;\s*\n?`,
)

var (
	mingwAsmStartRe  = regexp.MustCompile(`__asm(?:__)?(?:\s+__(?:volatile)__|\s+volatile)?\s*\(`)
	mingwAttrStartRe = regexp.MustCompile(`__(?:attribute|declspec)(?:__)?\s*\(`)
	// # 123 "C:\path\file.h"  — backslashes are C string escapes and corrupt the parser.
	lineDirectiveRe = regexp.MustCompile(`(?m)^(#\s*\d+\s*")([^"]*)(".*)`)

	libcVersionOnce sync.Once
	libcVersion     string
	libcVersionErr  error
)

// stripBalancedCalls removes each match of startRe through its balanced "(...)".
func stripBalancedCalls(src string, startRe *regexp.Regexp) string {
	var b strings.Builder
	b.Grow(len(src))
	i := 0
	for i < len(src) {
		loc := startRe.FindStringIndex(src[i:])
		if loc == nil {
			b.WriteString(src[i:])
			break
		}
		start := i + loc[0]
		open := i + loc[1] - 1
		b.WriteString(src[i:start])
		depth := 0
		j := open
		for j < len(src) {
			switch src[j] {
			case '(':
				depth++
			case ')':
				depth--
				if depth == 0 {
					j++
					i = j
					goto next
				}
			}
			j++
		}
		b.WriteString(src[start:])
		return b.String()
	next:
	}
	return b.String()
}

// preprocessorCmd builds the C preprocessor command from CC and optional CFLAGS.
// On Windows, MinGW gcc is preferred over clang when present on PATH.
func preprocessorCmd(goos, goarch string, args ...string) (*exec.Cmd, error) {
	ccFields, err := splitCompilerEnv(resolveCC(goos, goarch))
	if err != nil {
		return nil, fmt.Errorf("parse CC: %w", err)
	}
	if len(ccFields) == 0 {
		return nil, fmt.Errorf("CC is empty")
	}
	cflags, err := splitCompilerEnv(os.Getenv("CFLAGS"))
	if err != nil {
		return nil, fmt.Errorf("parse CFLAGS: %w", err)
	}
	cmdArgs := append([]string{}, ccFields[1:]...)
	cmdArgs = append(cmdArgs, defaultPreprocessorFlags(goos, goarch, ccFields[0])...)
	cmdArgs = append(cmdArgs, cflags...)
	cmdArgs = append(cmdArgs, args...)
	return exec.Command(ccFields[0], cmdArgs...), nil
}

// splitCompilerEnv splits CC/CFLAGS. mvdan.cc/sh treats \ as an escape, which
// turns Windows paths like C:\mingw64\bin\gcc.exe into C:mingw64bingcc.exe.
// Keep drive/UNC paths intact (split on ".exe" when args follow).
func splitCompilerEnv(val string) ([]string, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil, nil
	}
	if isWindowsCmdPath(val) {
		return splitWindowsCompilerEnv(val)
	}
	return shell.Fields(val, nil)
}

func isWindowsCmdPath(s string) bool {
	if strings.HasPrefix(s, `\\`) {
		return true
	}
	return len(s) >= 3 && s[1] == ':' && (s[2] == '\\' || s[2] == '/')
}

func splitWindowsCompilerEnv(val string) ([]string, error) {
	lower := strings.ToLower(val)
	if i := strings.Index(lower, ".exe"); i >= 0 {
		exe := val[:i+len(".exe")]
		rest := strings.TrimSpace(val[i+len(".exe"):])
		fields := []string{exe}
		if rest == "" {
			return fields, nil
		}
		extra, err := shell.Fields(rest, nil)
		if err != nil {
			return nil, err
		}
		return append(fields, extra...), nil
	}
	if !strings.ContainsAny(val, " \t") {
		return []string{val}, nil
	}
	return shell.Fields(filepath.ToSlash(val), nil)
}

func resolveCC(goos, goarch string) string {
	cc := strings.TrimSpace(os.Getenv("CC"))
	if goos == "windows" {
		if mingw := lookupMingwGCC(goarch); mingw != "" {
			if cc == "" || strings.Contains(filepath.Base(cc), "clang") {
				return mingw
			}
		}
	}
	if cc != "" {
		return cc
	}
	return defaultPreprocessor
}

// lookupMingwGCC returns the MinGW GCC driver on PATH so -E uses GNU headers
// without clang rewriting __declspec through __attribute__.
func lookupMingwGCC(goarch string) string {
	for _, cand := range []string{mingwTriple(goarch) + "-gcc", "gcc"} {
		if p, err := exec.LookPath(cand); err == nil {
			return p
		}
	}
	return ""
}

// defaultPreprocessorFlags returns compiler flags so host -E output stays
// within what ccgo can parse. Prefer triples and macros over rewriting sources.
func defaultPreprocessorFlags(goos, goarch, ccBinary string) []string {
	base := filepath.Base(ccBinary)
	isClang := strings.Contains(base, "clang")
	switch goos {
	case "windows":
		var flags []string
		if isClang {
			flags = append(flags, "--target="+mingwTriple(goarch))
			if sysroot := mingwSysroot(goarch); sysroot != "" {
				flags = append(flags, "--sysroot="+sysroot)
			}
		}
		// Erase GNU/MinGW extensions ccgo cannot parse. Empty __attribute__
		// on the command line so headers' `#define __declspec(a) __attribute__((a))`
		// expands to nothing (not stray `(dllimport)` tokens). Calling-convention
		// keywords similarly survive -E as junk identifiers.
		//
		// winnt.h pulls <x86intrin.h> and <emmintrin.h> (SSE/AVX vector junk
		// with __builtin_ia32_* that ccgo cannot type-check). tree-sitter only
		// needs Interlocked* from windows.h (psdk_inc/intrin-impl.h, earlier).
		// Pre-define include guards so those SIMD headers stay empty.
		flags = append(flags,
			"-D__extension__=",
			"-D__forceinline=static inline",
			"-D__attribute__(...)=",
			"-D__declspec(x)=",
			"-D__cdecl=",
			"-D__stdcall=",
			"-D__fastcall=",
			"-D__thiscall=",
			"-D_cdecl=",
			"-D__restrict=",
			"-D__restrict__=",
			"-D__MINGW_EXTENSION=",
			"-D_X86INTRIN_H_INCLUDED",
			"-D_X86GPRINTRIN_H_INCLUDED",
			"-D_IMMINTRIN_H_INCLUDED",
			"-D_MMINTRIN_H_INCLUDED",
			"-D_XMMINTRIN_H_INCLUDED",
			"-D_EMMINTRIN_H_INCLUDED",
			"-D_PMMINTRIN_H_INCLUDED",
			"-D_MM3DNOW_H_INCLUDED",
		)
		return flags
	case "darwin":
		// Blocks + availability(macos,introduced=10.13.4) leave tokens like
		// 10.13.4 that ccgo parses as invalid floats. Neutralize via macros.
		// Alignment-16 neon structs are ignored later via ccgo flags.
		return []string{
			"-fno-blocks",
			"-D__attribute__(...)=",
			"-D__extension__=",
			"-D_Nonnull=",
			"-D_Nullable=",
			"-D_Null_unspecified=",
			"-DAPI_AVAILABLE(...)=",
			"-DAPI_UNAVAILABLE(...)=",
			"-DAPI_DEPRECATED(...)=",
			"-DAPI_DEPRECATED_WITH_REPLACEMENT(...)=",
			"-D__API_AVAILABLE(...)=",
			"-D__API_UNAVAILABLE(...)=",
			"-D__API_DEPRECATED(...)=",
			"-D__API_DEPRECATED_WITH_REPLACEMENT(...)=",
		}
	default:
		return nil
	}
}

func ccgoExtraArgs(goos string) []string {
	args := []string{
		"-ignore-unsupported-alignment",
		"-ignore-unsupported-atomic-sizes",
		"-ignore-vector-functions",
	}
	if goos == "windows" {
		// Registered in ccgo as Opt("-winapi-no-errno"); modernc.org/opt
		// strips one leading '-', so callers must pass the double-dash form.
		args = append(args, "--winapi-no-errno")
		// MinGW windows.h pulls many CRT/API decls tree-sitter never calls.
		// Without this, link fails on symbols modernc.org/libc does not export
		// (e.g. MapViewOfFileNuma2) even though used paths only need Interlocked*
		// and a few CRT helpers that libc does provide. postProcessWindows
		// rewrites remaining bare/iqlibc names to libc.X*.
		args = append(args, "-ignore-link-errors")
	}
	return args
}

func mingwTriple(goarch string) string {
	switch goarch {
	case "arm64":
		return "aarch64-w64-mingw32"
	case "386":
		return "i686-w64-mingw32"
	default:
		return "x86_64-w64-mingw32"
	}
}

func mingwSysroot(goarch string) string {
	if s := strings.TrimSpace(os.Getenv("MINGW_SYSROOT")); s != "" {
		// Trust an explicit amd64 sysroot; other arches must match the triple layout.
		if goarch == "amd64" || goarch == "" {
			return s
		}
		if mingwSysrootMatches(s, goarch) {
			return s
		}
		return ""
	}
	var candidates []string
	switch goarch {
	case "arm64":
		candidates = []string{
			`C:\llvm-mingw`,
			`C:\msys64\clangarm64`,
			"/usr/aarch64-w64-mingw32",
		}
	case "386":
		candidates = []string{
			`C:\msys64\mingw32`,
			`C:\mingw32`,
			"/usr/i686-w64-mingw32",
		}
	default:
		candidates = []string{
			`C:\msys64\mingw64`,
			`C:\mingw64`,
			`C:\ProgramData\mingw64\mingw64`,
			`C:\ProgramData\chocolatey\lib\mingw\tools\install\mingw64`,
			"/mingw64",
			"/usr/x86_64-w64-mingw32",
		}
	}
	for _, candidate := range candidates {
		if mingwSysrootMatches(candidate, goarch) {
			return candidate
		}
	}
	return ""
}

func mingwSysrootMatches(root, goarch string) bool {
	if st, err := os.Stat(filepath.Join(root, "include")); err != nil || !st.IsDir() {
		return false
	}
	triple := mingwTriple(goarch)
	// Accept plain mingw64 roots for amd64; require triple-specific layout otherwise.
	if goarch == "amd64" || goarch == "" {
		return true
	}
	if st, err := os.Stat(filepath.Join(root, triple)); err == nil && st.IsDir() {
		return true
	}
	if st, err := os.Stat(filepath.Join(root, "lib", triple)); err == nil && st.IsDir() {
		return true
	}
	// llvm-mingw uses <root>/include and <root>/aarch64-w64-mingw32
	if st, err := os.Stat(filepath.Join(root, triple, "include")); err == nil && st.IsDir() {
		return true
	}
	return false
}

// sanitizePreprocessedC removes constructs that survive -E but are rejected by ccgo.
func sanitizePreprocessedC(src string) string {
	src = fixLineDirectivePaths(src)
	src = extensionFloatTypedefRe.ReplaceAllString(src, "")
	src = stripBalancedCalls(src, mingwAsmStartRe)
	src = stripBalancedCalls(src, mingwAttrStartRe)
	return src
}

// fixLineDirectivePaths rewrites backslashes in #line file paths to forward
// slashes. On Windows, gcc -E emits # 12 "D:\a\foo.h"; the C parser then
// treats \a as an escape and reports bogus syntax errors far from the cause.
func fixLineDirectivePaths(src string) string {
	return lineDirectiveRe.ReplaceAllStringFunc(src, func(m string) string {
		parts := lineDirectiveRe.FindStringSubmatch(m)
		if len(parts) != 4 {
			return m
		}
		return parts[1] + strings.ReplaceAll(parts[2], `\`, `/`) + parts[3]
	})
}

func preprocessorIdentity(goos, goarch string) string {
	ccFields, err := splitCompilerEnv(resolveCC(goos, goarch))
	if err != nil || len(ccFields) == 0 {
		ccFields = []string{resolveCC(goos, goarch)}
	}
	parts := append(append([]string{}, ccFields...), defaultPreprocessorFlags(goos, goarch, ccFields[0])...)
	if cflags := strings.TrimSpace(os.Getenv("CFLAGS")); cflags != "" {
		parts = append(parts, cflags)
	}
	return strings.Join(parts, " ")
}

// TranspileCore transpiles the tree-sitter core library into the target directory.
// Sources are fed to ccgo directly (ccgo's own preprocessor + NewConfig via CC);
// host-specific -D flags keep MinGW/Darwin headers within what ccgo can parse.
func (t *Transpiler) TranspileCore(outputDir string) error {
	tmpDir, err := t.prepareWorkDir("tree-sitter-core", "")
	if err != nil {
		return err
	}
	if !t.KeepTemp {
		defer os.RemoveAll(tmpDir)
	} else {
		slog.Info("keeping temp dir", "path", tmpDir)
	}

	// Setup go.mod
	if err := setupGoMod(tmpDir); err != nil {
		return fmt.Errorf("go.mod setup failed: %w", err)
	}

	libSrc := filepath.Join(t.TreeSitterPath, "lib/src")
	libC, err := filepath.Abs(filepath.Join(libSrc, "lib.c"))
	if err != nil {
		return err
	}
	stubsH, err := writeAtomicStubsHeader(tmpDir)
	if err != nil {
		return err
	}

	// Transpile: ccgo resolves #includes (including system headers) itself.
	coreGo := filepath.Join(tmpDir, "core.go")
	includes := []string{
		filepath.Join(t.TreeSitterPath, "lib/include"),
		libSrc,
	}
	if err := t.runCcgo(tmpDir, []string{libC}, coreGo, includes, nil, stubsH); err != nil {
		return fmt.Errorf("ccgo failed: %w", err)
	}

	// Post-process and write
	if outputDir != "" {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return err
		}

		// Read and modify to be a proper package
		data, err := os.ReadFile(coreGo)
		if err != nil {
			return err
		}

		goCode := postProcess(string(data), tmpDir)
		// Change package from main to grammar
		goCode = strings.Replace(goCode, "package main", "package grammar", 1)

		outputFile := filepath.Join(outputDir, fmt.Sprintf("core-%s-%s.go", t.GOOS, t.GOARCH))
		return os.WriteFile(outputFile, []byte(goCode), 0644)
	}

	// Output to stdout
	data, err := os.ReadFile(coreGo)
	if err != nil {
		return err
	}
	fmt.Print(postProcess(string(data), tmpDir))
	return nil
}

// TranspileGrammar transpiles a tree-sitter grammar unit into the target directory.
// grammarPath is the unit root (parent of src/); grammarName is the package/registry id
// (typically from tree_sitter_<name> in parser.c). If grammarName is empty, it is derived
// from the unit path.
func (t *Transpiler) TranspileGrammar(grammarPath, grammarName, outputDir string) error {
	if grammarName == "" {
		grammarName = normalizeGrammarName(grammarPath)
	}

	// Check parser.c exists
	parserC := filepath.Join(grammarPath, "src", "parser.c")
	if _, err := os.Stat(parserC); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("parser.c not found at %s", parserC)
		}
		return fmt.Errorf("cannot stat parser.c at %s: %w", parserC, err)
	}

	// Check for scanner.c
	scannerC := filepath.Join(grammarPath, "src", "scanner.c")
	hasScanner := true
	if _, err := os.Stat(scannerC); err != nil {
		if os.IsNotExist(err) {
			hasScanner = false
		} else {
			return fmt.Errorf("cannot stat scanner.c at %s: %w", scannerC, err)
		}
	}

	tmpDir, err := t.prepareWorkDir("grammar", grammarName)
	if err != nil {
		return err
	}
	if !t.KeepTemp {
		defer os.RemoveAll(tmpDir)
	} else {
		slog.Info("keeping temp dir", "path", tmpDir)
	}

	// If has scanner, combine parser and scanner into one file
	var combinedC string
	if hasScanner {
		combinedC = filepath.Join(tmpDir, "combined.c")
		if err := combineFiles([]string{scannerC, parserC}, combinedC); err != nil {
			return fmt.Errorf("failed to combine files: %w", err)
		}
	} else {
		combinedC = parserC
	}

	// Transpile combined file
	grammarGo := filepath.Join(tmpDir, "grammar.go")
	if err := t.transpileGrammarFile(grammarPath, combinedC, grammarGo, tmpDir); err != nil {
		return fmt.Errorf("transpilation failed: %w", err)
	}

	// Write output as a proper Go package
	if outputDir != "" {
		grammarOutDir := filepath.Join(outputDir, grammarName)
		if err := os.MkdirAll(grammarOutDir, 0755); err != nil {
			return err
		}

		// Read and modify package name
		data, err := os.ReadFile(grammarGo)
		if err != nil {
			return err
		}

		goCode := postProcess(string(data), tmpDir)
		// Change package from main to grammar name (preserve comments)
		goCode = strings.Replace(goCode, "package main", "package grammar_"+grammarName, 1)

		outputFile := filepath.Join(grammarOutDir, fmt.Sprintf("grammar-%s-%s.go", t.GOOS, t.GOARCH))
		if err := os.WriteFile(outputFile, []byte(goCode), 0644); err != nil {
			return err
		}

		// Generate API wrapper
		var apiErr error
		if hasScanner {
			apiErr = GenerateAPIWrapperWithScanner(outputDir, grammarName)
		} else {
			apiErr = GenerateAPIWrapper(outputDir, grammarName)
		}
		if apiErr != nil {
			return fmt.Errorf("failed to generate API: %w", apiErr)
		}
		if err := writeLangGoMod(outputDir, grammarName); err != nil {
			return fmt.Errorf("failed to write grammar go.mod: %w", err)
		}
	}

	return nil
}

// combineFiles sequentially concatenates multiple C input source paths
// into a single output file to supply ccgo with unified compilation units.
func combineFiles(inputs []string, output string) (err error) {
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := out.Close(); err == nil {
			err = cerr
		}
	}()

	for _, input := range inputs {
		var data []byte
		data, err = os.ReadFile(input)
		if err != nil {
			return err
		}
		if _, err = out.Write(data); err != nil {
			return err
		}
		if _, err = out.WriteString("\n\n"); err != nil {
			return err
		}
	}

	return nil
}

// goKeywordDefines renames C identifiers that collide with Go keywords in
// grammar sources (tree-sitter token names).
func goKeywordDefines() []string {
	return []string{
		"-Dfunc=func_token",
		"-Dinterface=interface_token",
		"-Dselect=select_token",
		"-Dchan=chan_token",
		"-Dgo=go_token",
		"-Dmap=map_token",
		"-Dpackage=package_token",
		"-Dtype=type_token",
		"-Dvar=var_token",
		"-Dimport=import_token",
		"-Ddefer=defer_token",
		"-Dfallthrough=fallthrough_token",
		"-Drange=range_token",
	}
}

// hostDefinesForCcgo returns -D flags for the target OS so ccgo's preprocessor
// sees the same header-softening macros we used to inject via external -E.
func hostDefinesForCcgo(goos, goarch string) []string {
	cc := resolveCC(goos, goarch)
	var out []string
	for _, f := range defaultPreprocessorFlags(goos, goarch, cc) {
		if strings.HasPrefix(f, "-D") {
			out = append(out, f)
		}
	}
	return out
}

//go:embed atomic_stubs.h
var atomicStubsHeader string

// writeAtomicStubsHeader writes static inline __atomic_* helpers forced into
// every TU via ccgo -include (separate .c static inlines are not visible).
// Content is the embedded atomic_stubs.h (single source of truth).
func writeAtomicStubsHeader(dir string) (string, error) {
	path := filepath.Join(dir, "atomic_stubs.h")
	if err := os.WriteFile(path, []byte(atomicStubsHeader), 0644); err != nil {
		return "", err
	}
	return path, nil
}

func (t *Transpiler) transpileGrammarFile(grammarPath, inputC, outputGo, workDir string) error {
	if err := setupGoMod(workDir); err != nil {
		return err
	}

	inputAbs, err := filepath.Abs(inputC)
	if err != nil {
		return err
	}
	includes := []string{
		filepath.Join(grammarPath, "src"),
		// Monorepo grammars (typescript, markdown) often include headers from a sibling common/
		grammarPath,
		filepath.Dir(grammarPath),
		filepath.Join(t.TreeSitterPath, "lib/include"),
		filepath.Join(t.TreeSitterPath, "lib/src"),
	}
	return t.runCcgo(workDir, []string{inputAbs}, outputGo, includes, goKeywordDefines(), "")
}

// runCcgo invokes ccgo on raw C sources. ccgo runs its own preprocessor using
// CC (via NewConfig) for system includes/predefines; we pass -I/-D/-std.
// forceInclude, if non-empty, is passed as -include (forced header).
func (t *Transpiler) runCcgo(workDir string, sources []string, outputPath string, includes, extraDefines []string, forceInclude string) (err error) {
	outputArg := outputPath
	if rel, err := filepath.Rel(workDir, outputPath); err == nil {
		outputArg = rel
	}

	ccgoArgs := append([]string{"ccgo"}, ccgoExtraArgs(t.GOOS)...)
	// gnu11: GCC 15+ defaults to C23 where bool is a keyword; ccgo uses _Bool.
	// Address-taken header static inlines (e.g. ts_decode_utf* in unicode.h) get
	// out-of-line bodies from ccgo itself; do not pass -fno-inline (that forces
	// emission of MinGW asm intrinsics and breaks Windows).
	ccgoArgs = append(ccgoArgs, "-std=gnu11", "-O0")
	ccgoArgs = append(ccgoArgs, hostDefinesForCcgo(t.GOOS, t.GOARCH)...)
	ccgoArgs = append(ccgoArgs, extraDefines...)
	if forceInclude != "" {
		incArg := forceInclude
		if rel, err := filepath.Rel(workDir, forceInclude); err == nil {
			incArg = rel
		}
		ccgoArgs = append(ccgoArgs, "-include", incArg)
	}
	for _, inc := range includes {
		abs, err := filepath.Abs(inc)
		if err != nil {
			return err
		}
		ccgoArgs = append(ccgoArgs, "-I", abs)
	}
	for _, src := range sources {
		arg := src
		if rel, err := filepath.Rel(workDir, src); err == nil && !strings.HasPrefix(rel, "..") {
			arg = rel
		}
		ccgoArgs = append(ccgoArgs, arg)
	}
	ccgoArgs = append(ccgoArgs, "-o", outputArg)

	origDir, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.Chdir(workDir); err != nil {
		return err
	}
	// Failed restore leaves later codegen steps in the temp workdir.
	defer func() {
		if restoreErr := os.Chdir(origDir); restoreErr != nil && err == nil {
			err = fmt.Errorf("restore working directory: %w", restoreErr)
		}
	}()

	// NewConfig probes CC for predefines and system include paths.
	cc := resolveCC(t.GOOS, t.GOARCH)
	// NewConfig expects a single executable name/path (not multi-word).
	if fields, err := splitCompilerEnv(cc); err == nil && len(fields) > 0 {
		cc = fields[0]
	}
	prevCC, hadCC := os.LookupEnv("CC")
	if err := os.Setenv("CC", cc); err != nil {
		return err
	}
	defer func() {
		if hadCC {
			_ = os.Setenv("CC", prevCC)
		} else {
			_ = os.Unsetenv("CC")
		}
	}()

	err = runCcgoIsolated(t.GOOS, t.GOARCH, ccgoArgs)
	if err != nil {
		if _, statErr := os.Stat(outputPath); statErr == nil {
			// File exists: ccgo often still writes output with -ignore-link-errors.
			return nil
		}
		return err
	}
	return nil
}

// runCcgoIsolated runs ccgo with completely isolated state.
// Because ccgo pollutes the global flag.CommandLine and os.Args during processing,
// running multiple transpilation tasks within the same process leads to panics
// and argument bleed. This function safely restores the global state upon exit.
func runCcgoIsolated(goos, goarch string, args []string) error {
	// Save state
	oldArgs := os.Args
	oldCmdLine := flag.CommandLine
	oldUsage := flag.Usage
	oldStderr := os.Stderr

	defer func() {
		os.Args = oldArgs
		flag.CommandLine = oldCmdLine
		flag.Usage = oldUsage
		os.Stderr = oldStderr
	}()

	// Fresh state
	os.Args = args
	flag.CommandLine = flag.NewFlagSet(args[0], flag.ContinueOnError)
	flag.CommandLine.SetOutput(io.Discard)
	flag.Usage = func() {}

	// Redirect stderr to capture output
	r, w, err := os.Pipe()
	if err != nil {
		return fmt.Errorf("create stderr pipe: %w", err)
	}
	os.Stderr = w

	var stderrBuf bytes.Buffer
	done := make(chan error, 1)
	go func() {
		_, copyErr := io.Copy(&stderrBuf, r)
		done <- copyErr
	}()

	// Run ccgo
	t := ccgo.NewTask(goos, goarch, args, io.Discard, w, nil)
	err = t.Main()

	if closeErr := w.Close(); closeErr != nil && err == nil {
		err = fmt.Errorf("close stderr pipe: %w", closeErr)
	}
	if copyErr := <-done; copyErr != nil {
		// Prefer combining so a truncated drain does not hide a real ccgo failure.
		if err != nil {
			err = fmt.Errorf("%w; drain stderr: %v", err, copyErr)
		} else {
			err = fmt.Errorf("drain stderr: %w", copyErr)
		}
	}

	if err != nil {
		return fmt.Errorf("%w\nstderr: %s", err, stderrBuf.String())
	}

	return nil
}

// Helper functions

func setupGoMod(dir string) error {
	libcVer, err := currentLibcVersion()
	if err != nil {
		return err
	}

	goModPath := filepath.Join(dir, "go.mod")
	goModContent := "module treesitter\n\ngo " + moduleGoVersion + "\n\nrequire modernc.org/libc " + libcVer + "\n"
	if err := os.WriteFile(goModPath, []byte(goModContent), 0644); err != nil {
		return err
	}

	getCmd := exec.Command("go", "mod", "download", "modernc.org/libc@"+libcVer)
	getCmd.Dir = dir
	getCmd.Stderr = os.Stderr
	getCmd.Stdout = io.Discard
	return getCmd.Run()
}

func currentLibcVersion() (string, error) {
	libcVersionOnce.Do(func() {
		cmd := exec.Command("go", "list", "-m", "-f", "{{.Version}}", "modernc.org/libc")
		out, err := cmd.Output()
		if err != nil {
			libcVersionErr = fmt.Errorf("resolve modernc.org/libc version: %w", err)
			return
		}
		libcVersion = strings.TrimSpace(string(out))
		if libcVersion == "" {
			libcVersionErr = fmt.Errorf("resolve modernc.org/libc version: empty version from go list")
		}
	})
	if libcVersionErr != nil {
		return "", libcVersionErr
	}
	return libcVersion, nil
}

// postProcess normalizes generated paths for deterministic cross-run outputs,
// and applies structural patches to pointer dereferencing in the tree-sitter AST
// to prevent panics when traversing null child relationships natively in Go.
func postProcess(goCode, workDir string) string {
	// Normalize temp workdir path so regenerated code is deterministic across runs.
	workDirSlash := filepath.ToSlash(workDir)
	goCode = strings.ReplaceAll(goCode, workDirSlash+"/", "./")
	goCode = strings.ReplaceAll(goCode, workDirSlash, ".")
	goCode = regexp.MustCompile(`/tmp/(grammar-gen|tree-sitter-gen)-\d+`).ReplaceAllString(goCode, "/tmp/$1")

	// Windows + -ignore-link-errors leaves unresolved C names as bare Go
	// identifiers, and builtins as iqlibc.* (import-qualifier tag without a
	// matching import). Map them to modernc.org/libc exports.
	goCode = postProcessWindowsLibcCalls(goCode)

	// Fix assert types
	goCode = regexp.MustCompile(`libc\.X__assert_fail\(tls, ([^,]*), ([^,]*), uint32\((\d+)\)`).
		ReplaceAllString(goCode, `libc.X__assert_fail(tls, $1, $2, int32($3)`)

	// Remove negative padding (any number of tabs)
	goCode = regexp.MustCompile(`\t+_ \[-\d+\]byte\n`).ReplaceAllString(goCode, "")

	// Remove debug println statements from ccgo
	goCode = regexp.MustCompile(`println\(__ccgo_ts \+ \d+\)[\s;]*`).ReplaceAllString(goCode, "")
	goCode = regexp.MustCompile(`println\(__ccgo_ts\+\d+, __ccgo_ts\+\d+\)[\s;]*`).ReplaceAllString(goCode, "")

	// Patch ts_subtree_child_count to check for NULL pointer before dereferencing
	goCode = regexp.MustCompile(`func ts_subtree_child_count\(tls \*libc\.TLS, _self Subtree\) \(r uint32_t\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 uint32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = uint32\(0\)
	\} else \{
		v1 = \(\*SubtreeHeapData\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\)\)\)\.Fchild_count
	\}
	return v1
\}`).ReplaceAllString(goCode, `func ts_subtree_child_count(tls *libc.TLS, _self Subtree) (r uint32_t) {
	// NULL check - if the subtree pointer is 0, return 0
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 uint32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = uint32(0)
	} else {
		v1 = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count
	}
	return v1
}`)

	// Patch ts_subtree_extra to check for NULL pointer
	goCode = regexp.MustCompile(`func ts_subtree_extra\(tls \*libc\.TLS, _self Subtree\) \(r uint8\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 int32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\) & 0x8 >> 3\)
	\} else \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\) \+ 44\)\) & 0x4 >> 2\)
	\}
	return libc\.Uint8FromInt32\(libc\.BoolInt32\(v1 != 0\)\)
\}`).ReplaceAllString(goCode, `func ts_subtree_extra(tls *libc.TLS, _self Subtree) (r uint8) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x8 >> 3)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x4 >> 2)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}`)

	// Patch ts_subtree_symbol to check for NULL pointer
	goCode = regexp.MustCompile(`func ts_subtree_symbol\(tls \*libc\.TLS, _self Subtree\) \(r TSSymbol\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 int32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = libc\.Int32FromUint8\(\(\*\(\*SubtreeInlineData\)\(unsafe\.Pointer\(bp\)\)\)\.Fsymbol\)
	\} else \{
		v1 = libc\.Int32FromUint16\(\(\*SubtreeHeapData\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\)\)\)\.Fsymbol\)
	\}
	return libc\.Uint16FromInt32\(v1\)
\}`).ReplaceAllString(goCode, `func ts_subtree_symbol(tls *libc.TLS, _self Subtree) (r TSSymbol) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = libc.Int32FromUint8((*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fsymbol)
	} else {
		v1 = libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol)
	}
	return libc.Uint16FromInt32(v1)
}`)

	// Patch ts_subtree_visible to check for NULL pointer
	goCode = regexp.MustCompile(`func ts_subtree_visible\(tls \*libc\.TLS, _self Subtree\) \(r uint8\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 int32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\) & 0x2 >> 1\)
	\} else \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\) \+ 44\)\) & 0x1 >> 0\)
	\}
	return libc\.Uint8FromInt32\(libc\.BoolInt32\(v1 != 0\)\)
\}`).ReplaceAllString(goCode, `func ts_subtree_visible(tls *libc.TLS, _self Subtree) (r uint8) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x2 >> 1)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x1 >> 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}`)

	// Patch ts_subtree_named to check for NULL pointer
	goCode = regexp.MustCompile(`func ts_subtree_named\(tls \*libc\.TLS, _self Subtree\) \(r uint8\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 int32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\) & 0x4 >> 2\)
	\} else \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\) \+ 44\)\) & 0x2 >> 1\)
	\}
	return libc\.Uint8FromInt32\(libc\.BoolInt32\(v1 != 0\)\)
\}`).ReplaceAllString(goCode, `func ts_subtree_named(tls *libc.TLS, _self Subtree) (r uint8) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x4 >> 2)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x2 >> 1)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}`)

	return goCode
}

// postProcessWindowsLibcCalls rewrites names left unresolved by ccgo's
// -ignore-link-errors path into modernc.org/libc calls.
//
// Patterns:
//   - iqlibc.FOO  → libc.XFOO  (builtin path uses import-qualifier tag "iq")
//   - bare CRT names used by MinGW headers / tree-sitter → libc.X*
func postProcessWindowsLibcCalls(goCode string) string {
	if !strings.Contains(goCode, "//go:build windows") && !strings.Contains(goCode, "WIN32 = 1") {
		return goCode
	}
	// Builtins and other symbols that were attributed to the libc import
	// qualifier without an actual import alias.
	goCode = strings.ReplaceAll(goCode, "iqlibc.", "libc.X")
	// Accidental double package prefix from a bad --prefix-undefined experiment.
	goCode = strings.ReplaceAll(goCode, "libc.Xlibc.X", "libc.X")

	// Bare C identifiers that should call into libc. Word-boundary via
	// negative lookbehind for letters/digits/underscore and selector dots.
	// Only rewrite call sites (name followed by '(').
	bareToLibc := []string{
		"__mingw_vswprintf",
		"__mingw_strtof",
		"__mingw_wcstod",
		"__mingw_wcstof",
		"_mingw_vswprintf",
		"_mingw_strtof",
		"_mingw_wcstod",
		"_mingw_wcstof",
		"strnlen",
		"wcsnlen",
		"iswctype",
		"iswspace",
		"iswalnum",
		"fdopen",
		"_fdopen",
		"_get_osfhandle",
		"_open_osfhandle",
		// ccgo rawName of __sync_* often drops one leading '_'.
		"__sync_fetch_and_and",
		"__sync_fetch_and_or",
		"__sync_fetch_and_xor",
		"__sync_fetch_and_add",
		"__sync_lock_test_and_set",
		"_sync_fetch_and_and",
		"_sync_fetch_and_or",
		"_sync_fetch_and_xor",
		"_sync_fetch_and_add",
		"_sync_lock_test_and_set",
		"_sync_val_compare_and_swapUintptr",
		"__sync_val_compare_and_swapUintptr",
		"MapViewOfFileNuma2",
		"open_osfhandle",
		"get_osfhandle",
		"towupper",
		"towlower",
	}
	for _, name := range bareToLibc {
		// Avoid rewriting already-qualified libc.Xname or definitions.
		re := regexp.MustCompile(`([^.\w])` + regexp.QuoteMeta(name) + `\(`)
		goCode = re.ReplaceAllString(goCode, `${1}libc.X`+name+`(`)
		// Start-of-line calls (rare).
		re2 := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(name) + `\(`)
		goCode = re2.ReplaceAllString(goCode, `libc.X`+name+`(`)
	}
	// Normalize single-underscore forms to the double-underscore libc exports.
	goCode = strings.ReplaceAll(goCode, "libc.X_mingw_", "libc.X__mingw_")
	goCode = strings.ReplaceAll(goCode, "libc.X_sync_", "libc.X__sync_")
	// MinGW MemoryBarrier/__faststorefence → portable fence.
	goCode = strings.ReplaceAll(goCode, "libc.X__builtin_ia32_sfence(", "libc.X__sync_synchronize(")

	// winnt.h maps Interlocked* → _Interlocked* via #define. ccgo emits both
	// `const InterlockedIncrement = "_InterlockedIncrement"` and call sites
	// `InterlockedIncrement(tls, p)`, which is a string not a function. Prefer
	// the underscore form that ccgo also emits as a real func.
	//
	// Only rewrite names whose const value is the underscore form; leave real
	// funcs like InterlockedBitTestAndSet alone.
	goCode = rewriteInterlockedStringAliasCalls(goCode)
	// MinGW x86_64 often leaves Interlocked* as string consts + bare call sites
	// without emitting the _Interlocked* bodies that arm64 headers produce.
	// Inject portable wrappers for any remaining undefined call targets.
	goCode = ensureMissingInterlockedHelpers(goCode)
	return goCode
}

// rewriteInterlockedStringAliasCalls rewrites call sites of Interlocked* names
// that ccgo materialized as string consts pointing at _Interlocked*.
func rewriteInterlockedStringAliasCalls(goCode string) string {
	// const InterlockedFoo = "_InterlockedFoo"
	constRe := regexp.MustCompile(`(?m)^const (Interlocked[A-Za-z0-9_]+) = "(_Interlocked[A-Za-z0-9_]+)"\s*$`)
	aliases := map[string]string{}
	for _, m := range constRe.FindAllStringSubmatch(goCode, -1) {
		aliases[m[1]] = m[2]
	}
	if len(aliases) == 0 {
		return goCode
	}
	for bare, underscored := range aliases {
		re := regexp.MustCompile(`\b` + regexp.QuoteMeta(bare) + `\(`)
		goCode = re.ReplaceAllString(goCode, underscored+"(")
	}
	return goCode
}

// interlockedHelperBodies are portable implementations matching MinGW arm64
// output (libc.X__sync_*). Only the helpers tree-sitter actually calls on
// windows/amd64 today; extend if new call sites appear.
var interlockedHelperBodies = map[string]string{
	"_InterlockedIncrement": `func _InterlockedIncrement(tls *libc.TLS, Addend uintptr) (r int32) {
	return int32(libc.X__sync_add_and_fetch(tls, Addend, int32(1)))
}
`,
	"_InterlockedDecrement": `func _InterlockedDecrement(tls *libc.TLS, Addend uintptr) (r int32) {
	return int32(libc.X__sync_sub_and_fetch(tls, Addend, int32(1)))
}
`,
	"_InterlockedIncrement64": `func _InterlockedIncrement64(tls *libc.TLS, Addend uintptr) (r int64) {
	return int64(libc.X__sync_add_and_fetch(tls, Addend, int64(1)))
}
`,
	"_InterlockedDecrement64": `func _InterlockedDecrement64(tls *libc.TLS, Addend uintptr) (r int64) {
	return int64(libc.X__sync_sub_and_fetch(tls, Addend, int64(1)))
}
`,
	"_InterlockedExchangeAdd": `func _InterlockedExchangeAdd(tls *libc.TLS, Addend uintptr, Value int32) (r int32) {
	return int32(libc.X__sync_fetch_and_add(tls, Addend, Value))
}
`,
	"_InterlockedCompareExchange": `func _InterlockedCompareExchange(tls *libc.TLS, Destination uintptr, ExChange int32, Comperand int32) (r int32) {
	return libc.X__sync_val_compare_and_swapInt32(tls, Destination, Comperand, ExChange)
}
`,
}

// ensureMissingInterlockedHelpers appends wrappers for _Interlocked* symbols
// that are called but never defined (common on MinGW x86_64 after rewriting
// Interlocked* string-const aliases).
func ensureMissingInterlockedHelpers(goCode string) string {
	defined := map[string]bool{}
	for _, m := range regexp.MustCompile(`(?m)^func (_Interlocked[A-Za-z0-9_]+)\(`).FindAllStringSubmatch(goCode, -1) {
		defined[m[1]] = true
	}
	called := map[string]bool{}
	for _, m := range regexp.MustCompile(`\b(_Interlocked[A-Za-z0-9_]+)\(`).FindAllStringSubmatch(goCode, -1) {
		called[m[1]] = true
	}
	var missing []string
	for name := range called {
		if defined[name] {
			continue
		}
		if _, ok := interlockedHelperBodies[name]; ok {
			missing = append(missing, name)
		}
	}
	if len(missing) == 0 {
		return goCode
	}
	sort.Strings(missing)
	var b strings.Builder
	b.WriteString(goCode)
	if !strings.HasSuffix(goCode, "\n") {
		b.WriteByte('\n')
	}
	b.WriteString("\n// Injected by postProcessWindowsLibcCalls: MinGW left Interlocked*\n")
	b.WriteString("// call sites without emitting _Interlocked* bodies.\n")
	for _, name := range missing {
		b.WriteString(interlockedHelperBodies[name])
		b.WriteByte('\n')
	}
	return b.String()
}

func (t *Transpiler) prepareWorkDir(kind, suffix string) (string, error) {
	parts := []string{
		os.TempDir(),
		deterministicTempRoot,
		fmt.Sprintf("%s-%s", t.GOOS, t.GOARCH),
		kind,
	}
	if suffix != "" {
		parts = append(parts, suffix)
	}
	workDir := filepath.Join(parts...)
	if err := os.RemoveAll(workDir); err != nil {
		return "", err
	}
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return "", err
	}
	return workDir, nil
}
