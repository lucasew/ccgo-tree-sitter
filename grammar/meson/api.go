package grammar_meson

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for meson
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_meson(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("meson", Language())
}
