package grammar_scheme

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for scheme
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_scheme(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("scheme", Language())
}
