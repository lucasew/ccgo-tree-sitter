package grammar_chatito

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for chatito
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_chatito(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("chatito", Language())
}
