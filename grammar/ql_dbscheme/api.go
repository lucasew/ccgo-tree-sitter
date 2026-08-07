package grammar_ql_dbscheme

import (
	"unsafe"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for ql_dbscheme (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_ql_dbscheme()))
}

func init() {
	grammar.Register("ql_dbscheme", Language())
}
