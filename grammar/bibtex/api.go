package grammar_bibtex

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for bibtex
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_bibtex(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("bibtex", Language())
}
