package grammar_capnp

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for capnp
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_capnp(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("capnp", Language())
}
