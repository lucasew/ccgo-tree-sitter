package grammar_xcompose

import (
	"unsafe"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for xcompose (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_xcompose()))
}

func init() {
	grammar.Register("xcompose", Language())
}
