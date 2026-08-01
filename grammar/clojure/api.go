package grammar_clojure

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for clojure
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_clojure(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("clojure", Language())
}
