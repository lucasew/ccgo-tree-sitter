package grammar_udev

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for udev
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_udev(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("udev", Language())
}
