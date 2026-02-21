package grammar_regex

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for regex
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_regex(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("regex", Language())
}
