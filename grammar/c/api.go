package grammar_c

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for c
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_c(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("c", Language())
}
