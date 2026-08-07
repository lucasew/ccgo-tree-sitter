package grammar_kotlin

import (
	"unsafe"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for kotlin (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_kotlin()))
}

func init() {
	grammar.Register("kotlin", Language())
}
