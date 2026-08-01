package grammar_poe_filter

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for poe_filter
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_poe_filter(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("poe_filter", Language())
}
