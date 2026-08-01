package grammar_rasi

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for rasi
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_rasi(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("rasi", Language())
}
