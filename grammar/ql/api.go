package grammar_ql

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for ql
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_ql(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("ql", Language())
}
