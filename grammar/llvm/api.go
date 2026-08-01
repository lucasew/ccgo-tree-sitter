package grammar_llvm

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for llvm
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_llvm(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("llvm", Language())
}
