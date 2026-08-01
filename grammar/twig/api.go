package grammar_twig

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for twig
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_twig(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("twig", Language())
}
