package grammar_ssh_config

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for ssh_config
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_ssh_config(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("ssh_config", Language())
}
