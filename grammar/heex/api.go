package grammar_heex

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for heex
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_heex(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("heex", Language())
}
