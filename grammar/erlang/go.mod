module github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/erlang

go 1.25.0

require (
	github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar v0.0.0
	modernc.org/libc v1.67.6
)

replace github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar => ../

replace modernc.org/libc => github.com/modernc-tree-sitter/libc v0.0.0-20260707203921-3c7a53d19f3f
