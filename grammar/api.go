package grammar

import (
	"unsafe"
	"modernc.org/libc"
)

// Language wraps a TSLanguage pointer
type Language = *TSLanguage

// Parser wraps parser functions
type Parser struct {
	ptr uintptr
	tls *libc.TLS
}

// Tree wraps a parse tree
type Tree struct {
	ptr uintptr
	tls *libc.TLS
}

// Node wraps a TSNode
type Node struct {
	node TSNode
	tls *libc.TLS
}

// NewParser creates a new parser
func NewParser() *Parser {
	tls := libc.NewTLS()
	ptr := ts_parser_new(tls)
	return &Parser{ptr: ptr, tls: tls}
}

// SetLanguage sets the language for parsing
func (p *Parser) SetLanguage(lang Language) bool {
	langPtr := uintptr(unsafe.Pointer(lang))
	return ts_parser_set_language(p.tls, p.ptr, langPtr) != 0
}

// ParseString parses a string
func (p *Parser) ParseString(source string) *Tree {
	sourceBytes := []byte(source)
	var sourcePtr uintptr
	if len(sourceBytes) > 0 {
		sourcePtr = uintptr(unsafe.Pointer(&sourceBytes[0]))
	}
	ptr := ts_parser_parse_string(p.tls, p.ptr, 0, sourcePtr, uint32(len(source)))
	return &Tree{ptr: ptr, tls: p.tls}
}

// Delete frees the parser
func (p *Parser) Delete() {
	ts_parser_delete(p.tls, p.ptr)
	p.tls.Close()
}

// Delete frees the tree
func (t *Tree) Delete() {
	ts_tree_delete(t.tls, t.ptr)
}

// RootNode returns the root node of the tree
func (t *Tree) RootNode() *Node {
	node := ts_tree_root_node(t.tls, t.ptr)
	return &Node{node: node, tls: t.tls}
}

// Type returns the node type as a string
func (n *Node) Type() string {
	ptr := ts_node_type(n.tls, n.node)
	if ptr == 0 {
		return ""
	}
	return libc.GoString(ptr)
}

// ChildCount returns the number of children
func (n *Node) ChildCount() uint32 {
	return ts_node_child_count(n.tls, n.node)
}

// Child returns the child at the given index
func (n *Node) Child(index uint32) *Node {
	node := ts_node_child(n.tls, n.node, index)
	return &Node{node: node, tls: n.tls}
}

// FieldNameForChild returns the field name for the child at the given index
func (n *Node) FieldNameForChild(index uint32) string {
	ptr := ts_node_field_name_for_child(n.tls, n.node, index)
	if ptr == 0 {
		return ""
	}
	return libc.GoString(ptr)
}

// NamedChildCount returns the number of named children
func (n *Node) NamedChildCount() uint32 {
	return ts_node_named_child_count(n.tls, n.node)
}

// NamedChild returns the named child at the given index
func (n *Node) NamedChild(index uint32) *Node {
	node := ts_node_named_child(n.tls, n.node, index)
	return &Node{node: node, tls: n.tls}
}

// StartByte returns the start byte offset
func (n *Node) StartByte() uint32 {
	return ts_node_start_byte(n.tls, n.node)
}

// EndByte returns the end byte offset
func (n *Node) EndByte() uint32 {
	return ts_node_end_byte(n.tls, n.node)
}

// String returns the S-expression representation of the node
func (n *Node) String() string {
	ptr := ts_node_string(n.tls, n.node)
	if ptr == 0 {
		return ""
	}
	str := libc.GoString(ptr)
	libc.Xfree(n.tls, ptr)
	return str
}

// IsNull returns true if the node is null
func (n *Node) IsNull() bool {
	return ts_node_is_null(n.tls, n.node) != 0
}

// IsNamed returns true if the node is named
func (n *Node) IsNamed() bool {
	return ts_node_is_named(n.tls, n.node) != 0
}

// IsExtra returns true if the node is extra
func (n *Node) IsExtra() bool {
	return ts_node_is_extra(n.tls, n.node) != 0
}

// IsError returns true if the node is an error
func (n *Node) IsError() bool {
	return ts_node_is_error(n.tls, n.node) != 0
}

// HasError returns true if the node or any descendant has an error
func (n *Node) HasError() bool {
	return ts_node_has_error(n.tls, n.node) != 0
}

// HasChanges returns true if the node has changed
func (n *Node) HasChanges() bool {
	return ts_node_has_changes(n.tls, n.node) != 0
}

// PrintTree prints the node tree in S-expression format
func (n *Node) PrintTree() string {
	if n.IsNull() {
		return "(null)"
	}
	return n.String()
}
