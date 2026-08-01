package grammar_ada

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for ada
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_ada(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("ada", Language())
}
