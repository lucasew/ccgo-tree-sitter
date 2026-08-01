package grammar_pascal

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for pascal
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_pascal(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("pascal", Language())
}
