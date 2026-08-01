package grammar_prisma

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for prisma
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_prisma(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("prisma", Language())
}
