package grammar_vimdoc

import (
	"unsafe"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for vimdoc (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_vimdoc()))
}

func init() {
	grammar.Register("vimdoc", Language())
}
