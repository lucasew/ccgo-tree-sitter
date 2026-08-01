package grammar_wat

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for wat
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_wat(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("wat", Language())
}
