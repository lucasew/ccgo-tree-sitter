package grammar_json5

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for json5
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_json5(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("json5", Language())
}
