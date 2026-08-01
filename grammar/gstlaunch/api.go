package grammar_gstlaunch

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for gstlaunch
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_gstlaunch(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("gstlaunch", Language())
}
