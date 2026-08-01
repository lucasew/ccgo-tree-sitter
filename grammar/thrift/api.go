package grammar_thrift

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for thrift
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_thrift(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("thrift", Language())
}
