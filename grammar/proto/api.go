package grammar_proto

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for proto
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_proto(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("proto", Language())
}
