package grammar_cyberchef

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for cyberchef
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_cyberchef(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("cyberchef", Language())
}
