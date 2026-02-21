package grammar_go

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for go
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_go(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("go", Language())
}
