package grammar_test

import (
	"testing"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
	jsongrammar "github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/json"
)

func TestBuildParseNodeNilAndNull(t *testing.T) {
	// Nil node and zero nodes must not panic; BuildParseNode returns nil.
	if got := grammar.BuildParseNode(nil, nil, ""); got != nil {
		t.Fatalf("nil node: got %#v want nil", got)
	}
	if got := grammar.BuildParseNode(&grammar.Node{}, []byte("{}"), ""); got != nil {
		t.Fatalf("zero node: got %#v want nil", got)
	}
}

func TestBuildParseNodeJSONStructure(t *testing.T) {
	lang := jsongrammar.Language()
	if lang == nil {
		t.Fatal("json Language() is nil")
	}
	if !grammar.LiveParseReady(lang) {
		t.Skip("leaven Subtree tagged-pointer crash in core parse")
	}

	src := []byte(`{"a": 1}`)
	p := grammar.NewParser()
	if !p.SetLanguage(lang) {
		t.Fatal("SetLanguage failed")
	}
	tree := p.ParseBytes(src)
	root := tree.RootNode()
	if root.IsNull() {
		t.Fatal("null root")
	}
	if root.HasError() {
		t.Fatalf("parse has errors:\n%s", root.PrintTree())
	}

	doc := grammar.BuildParseNode(root, src, "")
	if doc == nil {
		t.Fatal("BuildParseNode returned nil")
	}
	if doc.Type != "document" {
		t.Fatalf("root type: got %q want document", doc.Type)
	}
	if doc.Field != "" {
		t.Fatalf("root field: got %q want empty", doc.Field)
	}
	if doc.StartByte != 0 || doc.EndByte != uint32(len(src)) {
		t.Fatalf("root span: [%d,%d) want [0,%d)", doc.StartByte, doc.EndByte, len(src))
	}
	// Non-leaves omit Text.
	if doc.Text != "" {
		t.Fatalf("root Text: got %q want empty", doc.Text)
	}
	if len(doc.Children) != 1 {
		t.Fatalf("document children: got %d want 1", len(doc.Children))
	}

	obj := doc.Children[0]
	if obj.Type != "object" {
		t.Fatalf("first child type: got %q want object", obj.Type)
	}
	// object: "{" pair "}"
	if len(obj.Children) != 3 {
		t.Fatalf("object children: got %d want 3", len(obj.Children))
	}
	if obj.Children[0].Type != "{" || obj.Children[0].Text != "{" {
		t.Fatalf("object open: type=%q text=%q", obj.Children[0].Type, obj.Children[0].Text)
	}
	if obj.Children[2].Type != "}" || obj.Children[2].Text != "}" {
		t.Fatalf("object close: type=%q text=%q", obj.Children[2].Type, obj.Children[2].Text)
	}

	pair := obj.Children[1]
	if pair.Type != "pair" {
		t.Fatalf("pair type: got %q want pair", pair.Type)
	}
	// pair: key:string, ":", value:number
	if len(pair.Children) != 3 {
		t.Fatalf("pair children: got %d want 3", len(pair.Children))
	}

	key := pair.Children[0]
	if key.Type != "string" {
		t.Fatalf("key type: got %q want string", key.Type)
	}
	if key.Field != "key" {
		t.Fatalf("key field: got %q want key", key.Field)
	}
	// string has children (quotes + content); Text only on leaves.
	if key.Text != "" {
		t.Fatalf("string Text: got %q want empty (non-leaf)", key.Text)
	}
	if len(key.Children) != 3 {
		t.Fatalf("string children: got %d want 3", len(key.Children))
	}
	content := key.Children[1]
	if content.Type != "string_content" || content.Text != "a" {
		t.Fatalf("string_content: type=%q text=%q want string_content / a", content.Type, content.Text)
	}

	colon := pair.Children[1]
	if colon.Type != ":" || colon.Text != ":" || colon.Field != "" {
		t.Fatalf("colon: type=%q text=%q field=%q", colon.Type, colon.Text, colon.Field)
	}

	val := pair.Children[2]
	if val.Type != "number" {
		t.Fatalf("value type: got %q want number", val.Type)
	}
	if val.Field != "value" {
		t.Fatalf("value field: got %q want value", val.Field)
	}
	if val.Text != "1" {
		t.Fatalf("value text: got %q want 1", val.Text)
	}
	if len(val.Children) != 0 {
		t.Fatalf("number should be leaf, got %d children", len(val.Children))
	}
	if val.StartByte != 6 || val.EndByte != 7 {
		t.Fatalf("number span: [%d,%d) want [6,7)", val.StartByte, val.EndByte)
	}
}

func TestBuildParseNodeRootFieldName(t *testing.T) {
	lang := jsongrammar.Language()
	if lang == nil {
		t.Fatal("json Language() is nil")
	}
	if !grammar.LiveParseReady(lang) {
		t.Skip("leaven Subtree tagged-pointer crash in core parse")
	}
	src := []byte(`true`)
	p := grammar.NewParser()
	if !p.SetLanguage(lang) {
		t.Fatal("SetLanguage failed")
	}
	tree := p.ParseBytes(src)
	root := tree.RootNode()
	if root.IsNull() {
		t.Fatal("null root")
	}

	// fieldName is passed through on the root of this call.
	n := grammar.BuildParseNode(root, src, "from_caller")
	if n == nil {
		t.Fatal("BuildParseNode returned nil")
	}
	if n.Field != "from_caller" {
		t.Fatalf("root field: got %q want from_caller", n.Field)
	}
	if n.Type != "document" {
		t.Fatalf("type: got %q want document", n.Type)
	}
}
