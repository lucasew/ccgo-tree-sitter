package grammar_ini

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for ini
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_ini(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("ini", Language())
}
