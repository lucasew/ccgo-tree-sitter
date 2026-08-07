// Package grammar binds tree-sitter for pure Go (leaven, no CGO, no modernc).
//
// # Generated files
//
// core.go and per-language grammar/<lang>/grammar.go are raw leaven output
// (mise run codegen: clang 14 → LLVM IR → go tool leaven). Do not edit by hand.
//
// # Hand-written API
//
// api.go, query.go, parse.go, and registry.go wrap leaven symbols with a
// thread-safe Go API (no modernc, no CGO): Parser/Query serialize native work
// under a mutex; Tree and Node are immutable pure-Go snapshots safe for
// concurrent reads.
//
// Live parse still depends on leaven correctly modeling tree-sitter's tagged
// Subtree union. LiveParseReady(lang) probes once (sync.Once) and caches.
//
// # Registry
//
// Blank-import a language package so its init Registers the language:
//
//	import _ "github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/json"
//
// # Line index
//
// LineIndex converts tree-sitter byte offsets to 1-based lines and 0-based
// columns (byte offset within the line).
package grammar
