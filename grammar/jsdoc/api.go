package grammar_jsdoc

import (
	"unsafe"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for jsdoc (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_jsdoc()))
}

func init() {
	grammar.Register("jsdoc", Language())
}
