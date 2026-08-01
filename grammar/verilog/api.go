package grammar_verilog

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for verilog
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_verilog(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("verilog", Language())
}
