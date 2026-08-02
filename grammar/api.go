package grammar

import (
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"modernc.org/libc"
)

// Language wraps a TSLanguage pointer.
// Languages are immutable after load and may be shared across parsers and goroutines.
type Language = *TSLanguage

// Parser wraps a tree-sitter parser.
//
// Ownership is GC-managed: a runtime cleanup frees the native parser and TLS
// when the *Parser becomes unreachable. Explicit Delete is optional (eager free).
//
// Parse methods are safe for concurrent use: an internal mutex serializes the
// native parse and the copy into a pure-Go Tree. After Parse returns, the
// native tree is already freed; Tree/Node methods do not touch the parser.
// For throughput under heavy parallel load, prefer one Parser per goroutine.
type Parser struct {
	mu      sync.Mutex
	ptr     uintptr
	tls     *libc.TLS
	lang    Language
	cleanup runtime.Cleanup
}

// Tree is an immutable pure-Go snapshot of a parse tree.
//
// Parse copies node metadata out of tree-sitter and frees the native tree
// before returning. Concurrent use is safe without external locking (the
// snapshot is immutable). Delete is optional and only drops the Go root for
// earlier GC of a large tree.
type Tree struct {
	root *nodeData
	lang Language
}

// Node is a handle into a Tree snapshot.
//
// Keep the *Tree reachable while using Nodes from it (the Node holds a
// pointer to the Tree). Concurrent use is safe: methods only read immutable
// snapshot data. Null/empty nodes have a nil data pointer.
type Node struct {
	data *nodeData
	tree *Tree
}

// nodeData is one syntax node in the pure-Go snapshot.
type nodeData struct {
	typ        string
	start, end uint32
	named      bool
	extra      bool
	isError    bool
	hasError   bool
	hasChanges bool
	// path is child indices from the tree root (empty for root). Used to
	// re-locate the node after a temporary reparse for queries.
	path     []uint32
	fields   []string    // field name per child index (may be "")
	children []*nodeData // all children (named and anonymous)
}

type parserRes struct {
	ptr uintptr
	tls *libc.TLS
}

// NewParser creates a parser. Callers need not Delete; the GC will free it.
func NewParser() *Parser {
	tls := libc.NewTLS()
	ptr := ts_parser_new(tls)
	p := &Parser{ptr: ptr, tls: tls}
	p.cleanup = runtime.AddCleanup(p, freeParser, parserRes{ptr: ptr, tls: tls})
	return p
}

// SetLanguage sets the language for parsing.
func (p *Parser) SetLanguage(lang Language) bool {
	if p == nil {
		return false
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ptr == 0 || p.tls == nil {
		return false
	}
	langPtr := uintptr(unsafe.Pointer(lang))
	ok := ts_parser_set_language(p.tls, p.ptr, langPtr) != 0
	if ok {
		p.lang = lang
	}
	return ok
}

// ParseString parses a string.
//
// Source is copied into a libc CString for the call and freed with defer.
// Tree-sitter only borrows the buffer during parse. The returned Tree is a
// pure-Go snapshot; the native tree is freed before return.
func (p *Parser) ParseString(source string) *Tree {
	if p == nil {
		return &Tree{}
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ptr == 0 || p.tls == nil {
		return &Tree{lang: p.lang}
	}
	cstr, err := libc.CString(source)
	if err != nil {
		return &Tree{lang: p.lang}
	}
	defer libc.Xfree(nil, cstr)

	ptr := ts_parser_parse_string(p.tls, p.ptr, 0, cstr, uint32(len(source)))
	return snapshotAndFree(p, ptr)
}

// ParseBytes parses a contiguous UTF-8 source buffer.
//
// Source is copied into a NUL-terminated libc buffer for the call and freed
// with defer (no intermediate string allocation). Tree-sitter only borrows
// the buffer during parse. The returned Tree is a pure-Go snapshot; the
// native tree is freed before return.
func (p *Parser) ParseBytes(source []byte) *Tree {
	if p == nil {
		return &Tree{}
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ptr == 0 || p.tls == nil {
		return &Tree{lang: p.lang}
	}
	n := len(source)
	cstr := libc.Xmalloc(nil, libc.Tsize_t(n)+1)
	if cstr == 0 {
		return &Tree{lang: p.lang}
	}
	defer libc.Xfree(nil, cstr)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(cstr)), n), source)
	*(*byte)(unsafe.Pointer(cstr + uintptr(n))) = 0

	ptr := ts_parser_parse_string(p.tls, p.ptr, 0, cstr, uint32(n))
	return snapshotAndFree(p, ptr)
}

// snapshotAndFree copies the native tree into a pure-Go Tree and deletes the
// native tree. Caller must hold p.mu.
func snapshotAndFree(p *Parser, treePtr uintptr) *Tree {
	t := &Tree{lang: p.lang}
	if treePtr == 0 || p.tls == nil {
		return t
	}
	rootNative := ts_tree_root_node(p.tls, treePtr)
	t.root = captureNode(p.tls, rootNative, nil)
	ts_tree_delete(p.tls, treePtr)
	return t
}

// captureNode recursively copies one native node and its descendants.
func captureNode(tls *libc.TLS, node TSNode, path []uint32) *nodeData {
	if ts_node_is_null(tls, node) != 0 {
		return nil
	}
	typPtr := ts_node_type(tls, node)
	typ := ""
	if typPtr != 0 {
		typ = libc.GoString(typPtr)
	}
	count := ts_node_child_count(tls, node)
	d := &nodeData{
		typ:        typ,
		start:      ts_node_start_byte(tls, node),
		end:        ts_node_end_byte(tls, node),
		named:      ts_node_is_named(tls, node) != 0,
		extra:      ts_node_is_extra(tls, node) != 0,
		isError:    ts_node_is_error(tls, node) != 0,
		hasError:   ts_node_has_error(tls, node) != 0,
		hasChanges: ts_node_has_changes(tls, node) != 0,
		path:       path,
	}
	if count == 0 {
		return d
	}
	d.children = make([]*nodeData, 0, count)
	d.fields = make([]string, 0, count)
	for i := uint32(0); i < count; i++ {
		fieldPtr := ts_node_field_name_for_child(tls, node, i)
		field := ""
		if fieldPtr != 0 {
			field = libc.GoString(fieldPtr)
		}
		childPath := append(append([]uint32(nil), path...), i)
		child := captureNode(tls, ts_node_child(tls, node, i), childPath)
		d.fields = append(d.fields, field)
		d.children = append(d.children, child) // nil if native child is null
	}
	return d
}

// parseNative parses source and returns the native tree pointer without
// snapshotting. Caller must hold p.mu and must ts_tree_delete the result
// (when non-zero) with p.tls.
func (p *Parser) parseNative(source []byte) uintptr {
	if p.ptr == 0 || p.tls == nil {
		return 0
	}
	n := len(source)
	cstr := libc.Xmalloc(nil, libc.Tsize_t(n)+1)
	if cstr == 0 {
		return 0
	}
	defer libc.Xfree(nil, cstr)
	if n > 0 {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(cstr)), n), source)
	}
	*(*byte)(unsafe.Pointer(cstr + uintptr(n))) = 0
	return ts_parser_parse_string(p.tls, p.ptr, 0, cstr, uint32(n))
}

func freeParser(r parserRes) {
	if r.ptr != 0 && r.tls != nil {
		ts_parser_delete(r.tls, r.ptr)
	}
	if r.tls != nil {
		r.tls.Close()
	}
}

// Delete eagerly frees the parser. Optional: the GC will free it if omitted.
func (p *Parser) Delete() {
	if p == nil {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ptr == 0 {
		return
	}
	p.cleanup.Stop()
	freeParser(parserRes{ptr: p.ptr, tls: p.tls})
	p.ptr = 0
	p.tls = nil
	p.lang = nil
}

// Delete drops the snapshot root so a large tree can be collected earlier.
// Optional: the GC will collect the Tree if omitted. There is no native state.
func (t *Tree) Delete() {
	if t == nil {
		return
	}
	t.root = nil
}

// RootNode returns the root node of the tree.
func (t *Tree) RootNode() *Node {
	if t == nil || t.root == nil {
		return &Node{}
	}
	return &Node{data: t.root, tree: t}
}

// Type returns the node type as a string.
func (n *Node) Type() string {
	if n == nil || n.data == nil {
		return ""
	}
	return n.data.typ
}

// ChildCount returns the number of children.
func (n *Node) ChildCount() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	return uint32(len(n.data.children))
}

// Child returns the child at the given index.
func (n *Node) Child(index uint32) *Node {
	if n == nil || n.data == nil || int(index) >= len(n.data.children) {
		return &Node{}
	}
	return &Node{data: n.data.children[index], tree: n.tree}
}

// FieldNameForChild returns the field name for the child at the given index.
func (n *Node) FieldNameForChild(index uint32) string {
	if n == nil || n.data == nil || int(index) >= len(n.data.fields) {
		return ""
	}
	return n.data.fields[index]
}

// NamedChildCount returns the number of named children.
func (n *Node) NamedChildCount() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	var c uint32
	for _, ch := range n.data.children {
		if ch != nil && ch.named {
			c++
		}
	}
	return c
}

// NamedChild returns the named child at the given index.
func (n *Node) NamedChild(index uint32) *Node {
	if n == nil || n.data == nil {
		return &Node{}
	}
	var seen uint32
	for _, ch := range n.data.children {
		if ch == nil || !ch.named {
			continue
		}
		if seen == index {
			return &Node{data: ch, tree: n.tree}
		}
		seen++
	}
	return &Node{}
}

// StartByte returns the start byte offset.
func (n *Node) StartByte() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	return n.data.start
}

// EndByte returns the end byte offset.
func (n *Node) EndByte() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	return n.data.end
}

// String returns the S-expression representation of the node.
func (n *Node) String() string {
	if n == nil || n.data == nil {
		return ""
	}
	var b strings.Builder
	writeSexpr(&b, n.data)
	return b.String()
}

func writeSexpr(b *strings.Builder, d *nodeData) {
	if d == nil {
		return
	}
	// Match tree-sitter's ts_node_string shape closely enough for debugging
	// and error dumps: named nodes are (type …); anonymous leaves are quoted.
	if !d.named && len(d.children) == 0 {
		b.WriteString(quoteSexprAtom(d.typ))
		return
	}
	b.WriteByte('(')
	b.WriteString(d.typ)
	for i, ch := range d.children {
		if ch == nil {
			continue
		}
		b.WriteByte(' ')
		if i < len(d.fields) && d.fields[i] != "" {
			b.WriteString(d.fields[i])
			b.WriteString(": ")
		}
		writeSexpr(b, ch)
	}
	b.WriteByte(')')
}

func quoteSexprAtom(s string) string {
	// Anonymous tokens in tree-sitter sexprs are typically double-quoted.
	return `"` + strings.ReplaceAll(s, `"`, `\"`) + `"`
}

// IsNull returns true if the node is null.
func (n *Node) IsNull() bool {
	return n == nil || n.data == nil
}

// IsNamed returns true if the node is named.
func (n *Node) IsNamed() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.named
}

// IsExtra returns true if the node is extra.
func (n *Node) IsExtra() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.extra
}

// IsError returns true if the node is an error.
func (n *Node) IsError() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.isError
}

// HasError returns true if the node or any descendant has an error.
func (n *Node) HasError() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.hasError
}

// HasChanges returns true if the node has changed.
func (n *Node) HasChanges() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.hasChanges
}

// PrintTree returns the node tree in S-expression format, or "(null)" if the node is null.
func (n *Node) PrintTree() string {
	if n.IsNull() {
		return "(null)"
	}
	return n.String()
}
