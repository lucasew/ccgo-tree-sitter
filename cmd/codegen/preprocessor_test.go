package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"
)

func TestPreprocessorCmdDefaultClangLinux(t *testing.T) {
	t.Setenv("CC", "clang")
	os.Unsetenv("CFLAGS")
	os.Unsetenv("MINGW_SYSROOT")

	cmd, err := preprocessorCmd("linux", "amd64", "-E", "-o", "-", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"clang", "-E", "-o", "-", "x.c"}
	if !slices.Equal(cmd.Args, want) {
		t.Fatalf("args=%v, want %v", cmd.Args, want)
	}
}

func TestResolveCCPrefersMingwGCCOverClang(t *testing.T) {
	dir := t.TempDir()
	// LookPath on Windows only matches *.exe; isolate PATH so a host MinGW
	// install cannot shadow the stub (as on windows-latest runners).
	name := "x86_64-w64-mingw32-gcc"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	gcc := filepath.Join(dir, name)
	if err := os.WriteFile(gcc, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)
	t.Setenv("CC", "clang")

	want, err := exec.LookPath("x86_64-w64-mingw32-gcc")
	if err != nil {
		t.Fatal(err)
	}
	got := resolveCC("windows", "amd64")
	if got != want {
		t.Fatalf("resolveCC=%q, want %q", got, want)
	}
}

func TestResolveCCKeepsExplicitMingwGCC(t *testing.T) {
	dir := t.TempDir()
	name := "x86_64-w64-mingw32-gcc"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	shadow := filepath.Join(dir, name)
	if err := os.WriteFile(shadow, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)

	explicit := filepath.Join(dir, "explicit-gcc")
	if runtime.GOOS == "windows" {
		explicit += ".exe"
	}
	if err := os.WriteFile(explicit, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("CC", explicit)

	got := resolveCC("windows", "amd64")
	if got != explicit {
		t.Fatalf("resolveCC=%q, want explicit CC %q", got, explicit)
	}
}

func TestPreprocessorCmdWindowsMingwGccFlags(t *testing.T) {
	t.Setenv("CC", "/usr/bin/x86_64-w64-mingw32-gcc")
	os.Unsetenv("CFLAGS")
	os.Unsetenv("MINGW_SYSROOT")

	cmd, err := preprocessorCmd("windows", "amd64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range cmd.Args {
		if strings.HasPrefix(a, "--target=") {
			t.Fatalf("unexpected clang target flag %q in mingw gcc args %v", a, cmd.Args)
		}
	}
	for _, flag := range []string{
		"-D__extension__=",
		"-D__attribute__(...)=",
		"-D__declspec(x)=",
		"-D__stdcall=",
		"-D__cdecl=",
		"-D_X86INTRIN_H_INCLUDED",
		"-D_IMMINTRIN_H_INCLUDED",
		"-D_EMMINTRIN_H_INCLUDED",
		"-D_XMMINTRIN_H_INCLUDED",
		"-D_MMINTRIN_H_INCLUDED",
	} {
		if !slices.Contains(cmd.Args, flag) {
			t.Fatalf("missing %q in %v", flag, cmd.Args)
		}
	}
}

func TestSanitizePreprocessedCStripsMingwJunk(t *testing.T) {
	in := "int x __attribute__((dllimport));\n" +
		"void f(void) __asm__(\"foo\");\n" +
		"void g(void) __asm__ __volatile__ (\"lock bts{q %[Offset],%[Base] | %[Base],%[Offset]}\" : [old] \"=@ccc\" (old), [Base] \"+m\" (*Base) : [Offset] \"J\" \"r\" (Offset) : \"memory\" );\n" +
		"__declspec(dllexport) int y;\n" +
		"typedef float _Float32 f32;\n" +
		`# 42 "D:\a\ccgo-tree-sitter\third-party\tree-sitter\lib\include\tree_sitter\api.h"` + "\n"
	out := sanitizePreprocessedC(in)
	for _, junk := range []string{"__attribute__", "__asm__", "__declspec", "_Float32", "__volatile__"} {
		if strings.Contains(out, junk) {
			t.Fatalf("sanitize left %q in %q", junk, out)
		}
	}
	if !strings.Contains(out, "int x") || !strings.Contains(out, "void f(void)") || !strings.Contains(out, "void g(void)") || !strings.Contains(out, "int y") {
		t.Fatalf("sanitize removed declarations: %q", out)
	}
	if strings.Contains(out, `\a\`) || strings.Contains(out, `D:\`) {
		t.Fatalf("line directive still has backslashes: %q", out)
	}
	if !strings.Contains(out, `D:/a/ccgo-tree-sitter/`) {
		t.Fatalf("expected forward-slash line path, got %q", out)
	}
}

func TestSanitizeFixesWindowsLineDirectiveEscapes(t *testing.T) {
	// \a is bell in C strings; must not remain inside #line paths.
	in := "# 1 \"D:\\a\\work\\api.h\"\nint x;\n"
	out := fixLineDirectivePaths(in)
	if strings.Contains(out, `\`) {
		t.Fatalf("backslash remains: %q", out)
	}
	if !strings.Contains(out, `D:/a/work/api.h`) {
		t.Fatalf("got %q", out)
	}
}

func TestSplitCompilerEnvWindowsPathKeepsBackslashes(t *testing.T) {
	const cc = `C:\mingw64\bin\x86_64-w64-mingw32-gcc.exe`
	got, err := splitCompilerEnv(cc)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{cc}
	if !slices.Equal(got, want) {
		t.Fatalf("splitCompilerEnv(%q)=%q, want %q", cc, got, want)
	}
}

func TestSplitCompilerEnvWindowsPathWithArgs(t *testing.T) {
	const cc = `C:\mingw64\bin\gcc.exe -DFOO=1`
	got, err := splitCompilerEnv(cc)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{`C:\mingw64\bin\gcc.exe`, "-DFOO=1"}
	if !slices.Equal(got, want) {
		t.Fatalf("splitCompilerEnv(%q)=%q, want %q", cc, got, want)
	}
}

func TestPreprocessorCmdWindowsBackslashCC(t *testing.T) {
	const cc = `C:\mingw64\bin\x86_64-w64-mingw32-gcc.exe`
	t.Setenv("CC", cc)
	t.Setenv("PATH", "/nonexistent")
	os.Unsetenv("CFLAGS")
	os.Unsetenv("MINGW_SYSROOT")

	cmd, err := preprocessorCmd("windows", "amd64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	if cmd.Path != cc && cmd.Args[0] != cc {
		t.Fatalf("Path=%q Args[0]=%q, want %q", cmd.Path, cmd.Args[0], cc)
	}
	if strings.Contains(cmd.Args[0], `C:mingw64`) {
		t.Fatalf("backslashes were eaten: %q", cmd.Args[0])
	}
}

func TestPreprocessorCmdWindowsClangKeepsTarget(t *testing.T) {
	// Force clang by pointing CC at a clang path with no mingw gcc on PATH.
	t.Setenv("CC", "clang")
	t.Setenv("PATH", "/nonexistent")
	t.Setenv("MINGW_SYSROOT", `C:\mingw64`)
	os.Unsetenv("CFLAGS")

	cmd, err := preprocessorCmd("windows", "amd64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Contains(cmd.Args, "--target=x86_64-w64-mingw32") {
		t.Fatalf("args=%v", cmd.Args)
	}
	if !slices.Contains(cmd.Args, "-D__attribute__(...)=") {
		t.Fatalf("missing attribute erase in clang-mingw args %v", cmd.Args)
	}
}

func TestPreprocessorCmdDarwinFlags(t *testing.T) {
	t.Setenv("CC", "clang")
	os.Unsetenv("CFLAGS")

	for _, goarch := range []string{"amd64", "arm64"} {
		cmd, err := preprocessorCmd("darwin", goarch, "-E", "x.c")
		if err != nil {
			t.Fatal(err)
		}
		for _, flag := range []string{"-fno-blocks", "-D__attribute__(...)=", "-DAPI_AVAILABLE(...)=", "-D_Nonnull="} {
			if !slices.Contains(cmd.Args, flag) {
				t.Fatalf("goarch=%s args=%v, missing %q", goarch, cmd.Args, flag)
			}
		}
	}
}

func TestMingwTriple(t *testing.T) {
	cases := map[string]string{
		"amd64": "x86_64-w64-mingw32",
		"arm64": "aarch64-w64-mingw32",
		"386":   "i686-w64-mingw32",
		"":      "x86_64-w64-mingw32",
	}
	for goarch, want := range cases {
		if got := mingwTriple(goarch); got != want {
			t.Fatalf("mingwTriple(%q)=%q, want %q", goarch, got, want)
		}
	}
}

func TestPreprocessorCmdWindowsArm64ClangTarget(t *testing.T) {
	t.Setenv("CC", "clang")
	t.Setenv("PATH", "/nonexistent")
	t.Setenv("MINGW_SYSROOT", `C:\llvm-mingw`)
	os.Unsetenv("CFLAGS")

	cmd, err := preprocessorCmd("windows", "arm64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Contains(cmd.Args, "--target=aarch64-w64-mingw32") {
		t.Fatalf("args=%v", cmd.Args)
	}
}

func TestCcgoExtraArgs(t *testing.T) {
	if !slices.Contains(ccgoExtraArgs("darwin"), "-ignore-unsupported-alignment") {
		t.Fatal(ccgoExtraArgs("darwin"))
	}
	if !slices.Contains(ccgoExtraArgs("windows"), "--winapi-no-errno") {
		t.Fatal(ccgoExtraArgs("windows"))
	}
	if slices.Contains(ccgoExtraArgs("windows"), "-winapi-no-errno") {
		t.Fatal("single-dash -winapi-no-errno is rejected by ccgo opt parser")
	}
	if !slices.Contains(ccgoExtraArgs("windows"), "-ignore-link-errors") {
		t.Fatal(ccgoExtraArgs("windows"))
	}
	if slices.Contains(ccgoExtraArgs("linux"), "-ignore-link-errors") {
		t.Fatal("ignore-link-errors is windows-only")
	}
}

func TestPostProcessWindowsLibcCalls(t *testing.T) {
	in := "//go:build windows && amd64\n\npackage grammar\n\n" +
		"const InterlockedIncrement = \"_InterlockedIncrement\"\n" +
		"const InterlockedDecrement = \"_InterlockedDecrement\"\n" +
		"func InterlockedBitTestAndSet(tls *libc.TLS) {}\n" +
		"func f(tls *libc.TLS) {\n" +
		"\tiqlibc.__builtin_vsnprintf(tls, 0, 0, 0, 0)\n" +
		"\tv1 = strnlen(tls, _src, _count)\n" +
		"\t_ = iswctype(tls, uint16(1), uint16(8))\n" +
		"\t_ = _sync_fetch_and_add(tls, p, v)\n" +
		"\t_ = towupper(tls, uint16('a'))\n" +
		"\tlibc.X__builtin_ia32_sfence(tls)\n" +
		"\t_ = InterlockedIncrement(tls, p)\n" +
		"\t_ = InterlockedDecrement(tls, p)\n" +
		"\t_ = _InterlockedIncrement(tls, p)\n" +
		"\tInterlockedBitTestAndSet(tls)\n" +
		"}\n"
	out := postProcessWindowsLibcCalls(in)
	for _, want := range []string{
		"libc.X__builtin_vsnprintf(",
		"libc.Xstrnlen(",
		"libc.Xiswctype(",
		"libc.X__sync_fetch_and_add(",
		"libc.Xtowupper(",
		"libc.X__sync_synchronize(",
		"_InterlockedIncrement(tls, p)",
		"_InterlockedDecrement(tls, p)",
		"InterlockedBitTestAndSet(tls)",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("missing %q in\n%s", want, out)
		}
	}
	for _, junk := range []string{"iqlibc.", "\tstrnlen(", " iswctype(", " _sync_fetch_and_add(", " towupper(", "__builtin_ia32_sfence", "\tInterlockedIncrement(", "\tInterlockedDecrement("} {
		if strings.Contains(out, junk) {
			t.Fatalf("left junk %q in\n%s", junk, out)
		}
	}
	// Already-underscore form must not gain a double underscore.
	if strings.Contains(out, "__InterlockedIncrement(") {
		t.Fatalf("double-underscore Interlocked in\n%s", out)
	}
	// Real funcs (no string-const alias) must keep their bare name.
	if strings.Contains(out, "_InterlockedBitTestAndSet(") {
		t.Fatalf("rewrote real InterlockedBitTestAndSet in\n%s", out)
	}
	// Missing _Interlocked* bodies must be injected after rewrite.
	if !strings.Contains(out, "func _InterlockedIncrement(tls *libc.TLS, Addend uintptr)") {
		t.Fatalf("missing injected _InterlockedIncrement helper in\n%s", out)
	}
	if !strings.Contains(out, "func _InterlockedDecrement(tls *libc.TLS, Addend uintptr)") {
		t.Fatalf("missing injected _InterlockedDecrement helper in\n%s", out)
	}
	// Non-windows sources must be untouched.
	linux := "package grammar\nfunc f() { strnlen(tls, 0, 0) }\n"
	if got := postProcessWindowsLibcCalls(linux); got != linux {
		t.Fatalf("rewrote non-windows source: %q", got)
	}
}

func TestEnsureMissingInterlockedHelpersSkipsDefined(t *testing.T) {
	in := "//go:build windows && amd64\n\npackage grammar\n\n" +
		"func _InterlockedIncrement(tls *libc.TLS, Addend uintptr) (r int32) { return 0 }\n" +
		"func f(tls *libc.TLS) { _ = _InterlockedIncrement(tls, p) }\n"
	out := postProcessWindowsLibcCalls(in)
	// Exactly one definition — do not inject a second body.
	if n := strings.Count(out, "func _InterlockedIncrement("); n != 1 {
		t.Fatalf("want 1 _InterlockedIncrement def, got %d in\n%s", n, out)
	}
	if strings.Contains(out, "Injected by postProcessWindowsLibcCalls") {
		t.Fatalf("injected helpers despite existing def:\n%s", out)
	}
}
