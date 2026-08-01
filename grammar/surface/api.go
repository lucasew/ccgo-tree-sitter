package grammar_surface

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for surface
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_surface(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("surface", Language())
}
