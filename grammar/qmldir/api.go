package grammar_qmldir

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for qmldir
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_qmldir(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("qmldir", Language())
}
