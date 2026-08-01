package grammar_firrtl

import (
	"unsafe"
	"reflect"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for firrtl with external scanner properly connected
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_firrtl(nil)
	lang := (*grammar.TSLanguage)(unsafe.Pointer(ptr))

	// WORKAROUND: ccgo doesn't properly initialize function pointers in struct literals
	// Manually connect external scanner functions
	if lang.Fexternal_scanner.Fcreate == 0 {
		lang.Fexternal_scanner.Fcreate = reflect.ValueOf(tree_sitter_firrtl_external_scanner_create).Pointer()
		lang.Fexternal_scanner.Fdestroy = reflect.ValueOf(tree_sitter_firrtl_external_scanner_destroy).Pointer()
		lang.Fexternal_scanner.Fscan = reflect.ValueOf(tree_sitter_firrtl_external_scanner_scan).Pointer()
		lang.Fexternal_scanner.Fserialize = reflect.ValueOf(tree_sitter_firrtl_external_scanner_serialize).Pointer()
		lang.Fexternal_scanner.Fdeserialize = reflect.ValueOf(tree_sitter_firrtl_external_scanner_deserialize).Pointer()
	}

	return lang
}

func init() {
	grammar.Register("firrtl", Language())
}
