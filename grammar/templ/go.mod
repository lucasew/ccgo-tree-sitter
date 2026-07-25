module github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/templ

go 1.25.0

require (
	github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar v0.0.0
	modernc.org/libc v1.67.6
)

replace github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar => ../

replace modernc.org/libc => ../../third-party/libc
