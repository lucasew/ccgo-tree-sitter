package grammar_git_config

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for git_config
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_git_config(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("git_config", Language())
}
