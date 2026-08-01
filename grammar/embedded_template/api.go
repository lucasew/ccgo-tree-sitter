package grammar_embedded_template

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for embedded_template
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_embedded_template(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("embedded_template", Language())
}
