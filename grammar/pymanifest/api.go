package grammar_pymanifest

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for pymanifest
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_pymanifest(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("pymanifest", Language())
}
