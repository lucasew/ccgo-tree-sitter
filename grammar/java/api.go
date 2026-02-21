package grammar_java

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for java
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_java(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("java", Language())
}
