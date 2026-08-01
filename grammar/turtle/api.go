package grammar_turtle

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for turtle
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_turtle(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("turtle", Language())
}
