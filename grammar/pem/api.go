package grammar_pem

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for pem
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_pem(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("pem", Language())
}
