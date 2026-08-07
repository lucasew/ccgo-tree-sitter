module github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/firrtl

go 1.25.0

require (
	github.com/andybalholm/leaven v0.0.0-20260807161919-d7e0c93ee95b
	github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar v0.0.0
)

replace github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar => ../

replace github.com/andybalholm/leaven => github.com/lewtec/leaven v0.0.0-20260807161919-d7e0c93ee95b
