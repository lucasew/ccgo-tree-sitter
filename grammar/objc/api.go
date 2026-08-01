package grammar_objc

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for objc
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_objc(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("objc", Language())
}
