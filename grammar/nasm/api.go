package grammar_nasm

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for nasm
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_nasm(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("nasm", Language())
}
