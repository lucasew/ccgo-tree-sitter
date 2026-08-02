// Package grammar binds tree-sitter for pure Go (modernc.org/ccgo, no CGO).
//
// # Ownership
//
// Parser and Query free native state via runtime cleanup when the Go value
// becomes unreachable. Delete is optional (eager free).
//
// ParseString/ParseBytes copy the syntax tree into an immutable pure-Go
// snapshot (*Tree / *Node) and free the native tree before returning. Keep
// the *Tree reachable while using Nodes from it. There is no native tree to
// pin after Parse returns.
//
// # Concurrency
//
// Language values are immutable after load and may be shared across parsers
// and goroutines. Parser methods that touch the native parser serialize on an
// internal mutex. Tree and Node methods only read the snapshot and need no
// locking. Query methods serialize on the query's mutex; ExecuteMatches
// re-parses temporarily to run the native query engine.
// For high parallel throughput, prefer one Parser per goroutine.
//
// # Registry
//
// Blank-import a language package (for example
// _ "github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/go") so its init
// calls Register. Look up languages with Get (by name) or GetByExtension.
// Call Register only for languages you load outside those packages.
//
// # Generated files
//
// core-*.go and per-language grammar-*.go are codegen output
// (mise run codegen). Do not edit them by hand; change sources or the
// generator and regenerate.
//
// # Line index
//
// LineIndex converts tree-sitter byte offsets to 1-based lines and 0-based
// columns (byte offset within the line). Build once per source with
// NewLineIndex or NewLineIndexBytes; lookups are O(log lines).
package grammar
