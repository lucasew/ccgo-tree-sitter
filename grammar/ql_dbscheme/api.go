package grammar_ql_dbscheme

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for ql_dbscheme
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_ql_dbscheme(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("ql_dbscheme", Language())
}
