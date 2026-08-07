package grammar_gstlaunch

import (
	"unsafe"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for gstlaunch (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_gstlaunch()))
}

func init() {
	grammar.Register("gstlaunch", Language())
}
