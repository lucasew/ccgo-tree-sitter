package grammar_git_rebase

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for git_rebase
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_git_rebase(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("git_rebase", Language())
}
