package grammar_wast

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for wast
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_wast(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("wast", Language())
}
