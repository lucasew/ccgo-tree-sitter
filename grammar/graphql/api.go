package grammar_graphql

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for graphql
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_graphql(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("graphql", Language())
}
