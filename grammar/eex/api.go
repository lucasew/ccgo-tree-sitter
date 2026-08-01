package grammar_eex

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for eex
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_eex(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("eex", Language())
}
