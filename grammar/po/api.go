package grammar_po

import (
	"unsafe"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for po (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_po()))
}

func init() {
	grammar.Register("po", Language())
}
