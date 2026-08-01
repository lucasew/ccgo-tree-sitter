package grammar_gpg

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for gpg
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_gpg(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("gpg", Language())
}
