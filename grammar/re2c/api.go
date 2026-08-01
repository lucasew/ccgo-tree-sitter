package grammar_re2c

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for re2c
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_re2c(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("re2c", Language())
}
