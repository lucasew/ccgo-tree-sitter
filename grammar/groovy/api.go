package grammar_groovy

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for groovy
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_groovy(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("groovy", Language())
}
