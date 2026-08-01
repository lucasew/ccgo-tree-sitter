package grammar_nqc

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for nqc
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_nqc(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("nqc", Language())
}
