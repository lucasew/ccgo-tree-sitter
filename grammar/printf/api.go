package grammar_printf

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for printf
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_printf(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("printf", Language())
}
