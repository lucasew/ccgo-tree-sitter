package grammar_smithy

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for smithy
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_smithy(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("smithy", Language())
}
