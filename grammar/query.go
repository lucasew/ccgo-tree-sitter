package grammar

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	"modernc.org/libc"
)

// Query wraps a compiled tree-sitter query.
//
// Ownership is GC-managed via runtime cleanup. Explicit Delete is optional.
//
// Query methods are safe for concurrent use. An internal mutex serializes
// access to the native query and TLS; callers do not need external locking.
// ExecuteMatches re-parses source temporarily (Trees are pure-Go snapshots),
// runs the query, and frees the native tree before return.
type Query struct {
	mu      sync.Mutex
	ptr     uintptr
	tls     *libc.TLS
	lang    Language
	cleanup runtime.Cleanup
}

// QueryCursor wraps a query cursor. Cleanup pins the parent *Query.
// Methods are safe for concurrent use; they lock the parent Query.
type QueryCursor struct {
	ptr     uintptr
	query   *Query
	cleanup runtime.Cleanup
}

type QueryCompileError struct {
	Offset uint32
	Type   TSQueryError
}

func (e *QueryCompileError) Error() string {
	return fmt.Sprintf("query compile error at offset %d: %s", e.Offset, queryErrorName(e.Type))
}

type QueryCapture struct {
	Index     uint32 `json:"index"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	StartByte uint32 `json:"start_byte"`
	EndByte   uint32 `json:"end_byte"`
	Text      string `json:"text,omitempty"`
}

type QueryMatch struct {
	ID           uint32         `json:"id"`
	PatternIndex uint16         `json:"pattern_index"`
	Captures     []QueryCapture `json:"captures"`
}

type queryRes struct {
	ptr uintptr
	tls *libc.TLS
}

type cursorRes struct {
	ptr   uintptr
	query *Query // pins query until cursor cleanup runs
}

// NewQuery compiles a query. Callers need not Delete; the GC will free it.
func NewQuery(lang Language, source string) (*Query, error) {
	tls := libc.NewTLS()

	cstr, err := libc.CString(source)
	if err != nil {
		tls.Close()
		return nil, err
	}
	defer libc.Xfree(nil, cstr)

	var errOffset uint32
	var errType TSQueryError1
	ptr := ts_query_new(
		tls,
		uintptr(unsafe.Pointer(lang)),
		cstr,
		uint32(len(source)),
		uintptr(unsafe.Pointer(&errOffset)),
		uintptr(unsafe.Pointer(&errType)),
	)
	if ptr == 0 {
		tls.Close()
		return nil, &QueryCompileError{
			Offset: errOffset,
			Type:   TSQueryError(errType),
		}
	}

	q := &Query{ptr: ptr, tls: tls, lang: lang}
	q.cleanup = runtime.AddCleanup(q, freeQuery, queryRes{ptr: ptr, tls: tls})
	return q, nil
}

func freeQuery(r queryRes) {
	if r.ptr != 0 && r.tls != nil {
		ts_query_delete(r.tls, r.ptr)
	}
	if r.tls != nil {
		r.tls.Close()
	}
}

func freeCursor(r cursorRes) {
	if r.ptr == 0 || r.query == nil {
		return
	}
	r.query.mu.Lock()
	defer r.query.mu.Unlock()
	freeCursorUnlocked(r)
}

// freeCursorUnlocked deletes the native cursor. Caller must hold r.query.mu.
func freeCursorUnlocked(r cursorRes) {
	if r.ptr == 0 || r.query == nil || r.query.tls == nil {
		return
	}
	ts_query_cursor_delete(r.query.tls, r.ptr)
}

// Delete eagerly frees the query. Optional: the GC will free it if omitted.
func (q *Query) Delete() {
	if q == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.ptr == 0 {
		return
	}
	q.cleanup.Stop()
	freeQuery(queryRes{ptr: q.ptr, tls: q.tls})
	q.ptr = 0
	q.tls = nil
	q.lang = nil
}

// NewCursor creates a cursor for this query.
func (q *Query) NewCursor() *QueryCursor {
	if q == nil {
		return &QueryCursor{}
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.newCursorUnlocked()
}

func (q *Query) newCursorUnlocked() *QueryCursor {
	if q.ptr == 0 || q.tls == nil {
		return &QueryCursor{}
	}
	ptr := ts_query_cursor_new(q.tls)
	c := &QueryCursor{ptr: ptr, query: q}
	if ptr != 0 {
		c.cleanup = runtime.AddCleanup(c, freeCursor, cursorRes{ptr: ptr, query: q})
	}
	return c
}

// Delete eagerly frees the cursor. Optional: the GC will free it if omitted.
func (c *QueryCursor) Delete() {
	if c == nil || c.ptr == 0 || c.query == nil {
		return
	}
	c.cleanup.Stop()
	freeCursor(cursorRes{ptr: c.ptr, query: c.query})
	c.ptr = 0
	c.query = nil
}

// ExecuteMatches runs the query over root and returns all matches.
//
// Trees are pure-Go snapshots, so this re-parses source with a temporary
// parser, walks root's child-index path to the matching native node, runs the
// query, then frees all native state before return.
// Returns nil if the query is unusable, root is nil/null, or reparse fails.
// Safe for concurrent use.
func (q *Query) ExecuteMatches(root *Node, source []byte) []QueryMatch {
	if q == nil || q.ptr == 0 {
		return nil
	}
	if root == nil || root.data == nil || root.IsNull() {
		return nil
	}

	var lang Language
	if root.tree != nil {
		lang = root.tree.lang
	}
	if lang == nil {
		lang = q.lang
	}
	if lang == nil {
		return nil
	}

	q.mu.Lock()
	defer q.mu.Unlock()
	if q.tls == nil || q.ptr == 0 {
		return nil
	}

	// Temporary parser for a native tree (snapshot path cannot feed queries).
	p := NewParser()
	defer p.Delete()
	if !p.SetLanguage(lang) {
		return nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	treePtr := p.parseNative(source)
	if treePtr == 0 {
		return nil
	}
	defer ts_tree_delete(p.tls, treePtr)

	nativeRoot := ts_tree_root_node(p.tls, treePtr)
	nativeNode := walkChildPath(p.tls, nativeRoot, root.data.path)
	if ts_node_is_null(p.tls, nativeNode) != 0 {
		return nil
	}

	cursor := q.newCursorUnlocked()
	if cursor.ptr == 0 {
		return nil
	}
	defer func() {
		cursor.cleanup.Stop()
		freeCursorUnlocked(cursorRes{ptr: cursor.ptr, query: q})
		cursor.ptr = 0
		cursor.query = nil
	}()

	ts_query_cursor_exec(q.tls, cursor.ptr, q.ptr, nativeNode)

	matches := make([]QueryMatch, 0)
	var rawMatch TSQueryMatch
	for ts_query_cursor_next_match(q.tls, cursor.ptr, uintptr(unsafe.Pointer(&rawMatch))) != 0 {
		captures := make([]QueryCapture, 0, rawMatch.Fcapture_count)
		for i := uint16(0); i < rawMatch.Fcapture_count; i++ {
			rawCapture := (*TSQueryCapture)(unsafe.Pointer(rawMatch.Fcaptures + uintptr(i)*unsafe.Sizeof(TSQueryCapture{})))
			name := q.captureNameUnlocked(rawCapture.Findex)
			start := ts_node_start_byte(q.tls, rawCapture.Fnode)
			end := ts_node_end_byte(q.tls, rawCapture.Fnode)
			typePtr := ts_node_type(q.tls, rawCapture.Fnode)
			typeName := ""
			if typePtr != 0 {
				typeName = libc.GoString(typePtr)
			}
			capture := QueryCapture{
				Index:     rawCapture.Findex,
				Name:      name,
				Type:      typeName,
				StartByte: start,
				EndByte:   end,
			}
			if int(end) <= len(source) && start <= end {
				capture.Text = string(source[start:end])
			}
			captures = append(captures, capture)
		}

		matches = append(matches, QueryMatch{
			ID:           rawMatch.Fid,
			PatternIndex: rawMatch.Fpattern_index,
			Captures:     captures,
		})
	}
	return matches
}

// walkChildPath follows child indices from root (empty path = root).
func walkChildPath(tls *libc.TLS, root TSNode, path []uint32) TSNode {
	n := root
	for _, i := range path {
		if ts_node_is_null(tls, n) != 0 {
			return n
		}
		n = ts_node_child(tls, n, i)
	}
	return n
}

func (q *Query) captureNameUnlocked(captureIndex uint32) string {
	var length uint32
	ptr := ts_query_capture_name_for_id(q.tls, q.ptr, captureIndex, uintptr(unsafe.Pointer(&length)))
	if ptr == 0 || length == 0 {
		return ""
	}

	data := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), int(length))
	return string(data)
}

func queryErrorName(errType TSQueryError) string {
	switch errType {
	case TSQueryErrorSyntax:
		return "syntax"
	case TSQueryErrorNodeType:
		return "node_type"
	case TSQueryErrorField:
		return "field"
	case TSQueryErrorCapture:
		return "capture"
	case TSQueryErrorStructure:
		return "structure"
	case TSQueryErrorLanguage:
		return "language"
	default:
		return "unknown"
	}
}
