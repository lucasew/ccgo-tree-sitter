package grammar_c_sharp

import (
	"unsafe"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for c_sharp (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_c_sharp()))
}

func init() {
	grammar.Register("c_sharp", Language())
}
