package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	ccgo "modernc.org/ccgo/v4/lib"
)

// Transpiler handles C to Go transpilation using ccgo
type Transpiler struct {
	TreeSitterPath string
	GOOS           string
	GOARCH         string
	KeepTemp       bool
}

// TranspileCore transpiles the tree-sitter core library
func (t *Transpiler) TranspileCore(outputDir string) error {
	libC, err := filepath.Abs(filepath.Join(t.TreeSitterPath, "lib/src/lib.c"))
	if err != nil {
		return err
	}
	tsInclude, _ := filepath.Abs(filepath.Join(t.TreeSitterPath, "lib/include"))
	tsSrc, _ := filepath.Abs(filepath.Join(t.TreeSitterPath, "lib/src"))

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	outputFile, _ := filepath.Abs(filepath.Join(outputDir, fmt.Sprintf("core-%s-%s.go", t.GOOS, t.GOARCH)))

	includeFlags := []string{
		"-I", tsInclude,
		"-I", tsSrc,
		"-include", "cmd/codegen/atomic_stubs.h",
		"--package-name", "grammar",
		"--prefix-external", "X",
	}
	if err := t.runCcgo(includeFlags, libC, outputFile); err != nil {
		return fmt.Errorf("ccgo failed: %w", err)
	}

	return postProcessFile(outputFile, "package main", "package grammar")
}

// TranspileGrammar transpiles a tree-sitter grammar
func (t *Transpiler) TranspileGrammar(grammarPath, outputDir string) error {
	grammarName := extractGrammarName(grammarPath)

	parserC := filepath.Join(grammarPath, "src", "parser.c")
	if _, err := os.Stat(parserC); os.IsNotExist(err) {
		return fmt.Errorf("parser.c not found at %s", parserC)
	}

	scannerC := filepath.Join(grammarPath, "src", "scanner.c")
	hasScanner := true
	if _, err := os.Stat(scannerC); os.IsNotExist(err) {
		hasScanner = false
	}

	// Input: parser.c alone, or scanner.c+parser.c combined into a temp C file
	var inputC string
	if hasScanner {
		tmp, err := os.CreateTemp("", "combined_*.c")
		if err != nil {
			return err
		}
		tmp.Close()
		if !t.KeepTemp {
			defer os.Remove(tmp.Name())
		}
		if err := combineFiles([]string{scannerC, parserC}, tmp.Name()); err != nil {
			return fmt.Errorf("failed to combine files: %w", err)
		}
		inputC = tmp.Name()
	} else {
		var err error
		inputC, err = filepath.Abs(parserC)
		if err != nil {
			return err
		}
	}

	grammarOutDir := filepath.Join(outputDir, grammarName)
	if err := os.MkdirAll(grammarOutDir, 0755); err != nil {
		return err
	}
	outputFile, _ := filepath.Abs(filepath.Join(grammarOutDir, fmt.Sprintf("grammar-%s-%s.go", t.GOOS, t.GOARCH)))

	grammarSrc, _ := filepath.Abs(filepath.Join(grammarPath, "src"))
	tsInclude, _ := filepath.Abs(filepath.Join(t.TreeSitterPath, "lib/include"))
	tsSrc, _ := filepath.Abs(filepath.Join(t.TreeSitterPath, "lib/src"))
	includeFlags := []string{
		"-I", grammarSrc,
		"-I", tsInclude,
		"-I", tsSrc,
		"--package-name", grammarName,
		"--prefix-external", "X",
	}
	if err := t.runCcgo(includeFlags, inputC, outputFile); err != nil {
		return fmt.Errorf("transpilation failed: %w", err)
	}

	if hasScanner {
		return GenerateAPIWrapperWithScanner(outputDir, grammarName)
	}
	return GenerateAPIWrapper(outputDir, grammarName)
}

func postProcessFile(path, oldPkg, newPkg string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	goCode := postProcess(string(data))
	goCode = strings.Replace(goCode, oldPkg, newPkg, 1)
	return os.WriteFile(path, []byte(goCode), 0644)
}

func combineFiles(inputs []string, output string) error {
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, input := range inputs {
		data, err := os.ReadFile(input)
		if err != nil {
			return err
		}
		if _, err := out.Write(data); err != nil {
			return err
		}
		out.WriteString("\n\n")
	}

	return nil
}

func (t *Transpiler) runCcgo(extraFlags []string, inputPath, outputPath string) error {
	ccgoArgs := append([]string{"ccgo"}, extraFlags...)
	ccgoArgs = append(ccgoArgs, inputPath, "-o", outputPath)
	task := ccgo.NewTask(t.GOOS, t.GOARCH, ccgoArgs, os.Stdout, os.Stderr, nil)
	return task.Main()
}

func postProcess(goCode string) string {
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

func extractGrammarName(path string) string {
	name := filepath.Base(path)
	prefix := "tree-sitter-"
	name = strings.TrimPrefix(name, prefix)
	name = strings.ReplaceAll(name, "-", "_")
	return name
}
