// Package grammar binds tree-sitter for pure Go (leaven, no CGO / no ccgo).
//
// # Generated files
//
// core.go and per-language grammar/<lang>/grammar.go are raw leaven output
// (mise run codegen: clang 14 → LLVM IR → go tool leaven). Do not edit by hand.
//
// The hand-written API layer (api.go, parse.go, …) still needs rewiring onto
// leaven symbols; treat this package as mid-migration.
//
// # Registry
//
// Blank-import a language package so its init can Register once wrappers match.
//
// # Line index
//
// LineIndex converts tree-sitter byte offsets to 1-based lines and 0-based
// columns (byte offset within the line).
package grammar
