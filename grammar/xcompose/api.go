package grammar_xcompose

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for xcompose
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_xcompose(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("xcompose", Language())
}
