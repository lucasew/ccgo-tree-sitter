package grammar_sparql

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for sparql
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_sparql(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("sparql", Language())
}
