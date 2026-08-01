package grammar_vhdl

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for vhdl
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_vhdl(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("vhdl", Language())
}
