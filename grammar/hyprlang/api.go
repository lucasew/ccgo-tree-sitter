package grammar_hyprlang

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for hyprlang
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_hyprlang(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("hyprlang", Language())
}
