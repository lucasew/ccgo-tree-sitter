package grammar_cst

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for cst
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_cst(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("cst", Language())
}
