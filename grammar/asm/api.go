package grammar_asm

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for asm
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_asm(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("asm", Language())
}
