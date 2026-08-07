package grammar_rescript

import (
	"unsafe"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for rescript (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_rescript()))
}

func init() {
	grammar.Register("rescript", Language())
}
