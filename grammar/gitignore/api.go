package grammar_gitignore

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for gitignore
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_gitignore(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("gitignore", Language())
}
