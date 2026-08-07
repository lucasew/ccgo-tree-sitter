package grammar_test

import (
	"sync"
	"testing"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
	gogrammar "github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/go"
	jsongrammar "github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/json"
)

func TestParserConcurrentShare(t *testing.T) {
	// Use Language() directly: registry tests may overwrite Get("go") with stubs.
	lang := gogrammar.Language()
	if lang == nil {
		t.Fatal("go Language() is nil")
	}
	if !grammar.LiveParseReady(lang) {
		t.Skip("leaven Subtree tagged-pointer crash in core parse")
	}
	p := grammar.NewParser()
	if !p.SetLanguage(lang) {
		t.Fatal("SetLanguage failed")
	}
	const src = "package main\n\nfunc main() {}\n"
	var wg sync.WaitGroup
	errCh := make(chan string, 32)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				tree := p.ParseString(src)
				root := tree.RootNode()
				if root.IsNull() {
					errCh <- "null root"
					return
				}
				if root.Type() != "source_file" {
					errCh <- "bad type " + root.Type()
					return
				}
				_ = root.ChildCount()
				if root.ChildCount() > 0 {
					_ = root.Child(0).Type()
				}
			}
		}()
	}
	wg.Wait()
	close(errCh)
	for e := range errCh {
		t.Fatal(e)
	}
}

func TestQueryConcurrentExecute(t *testing.T) {
	lang := jsongrammar.Language()
	if lang == nil {
		t.Fatal("json Language() is nil")
	}
	if !grammar.LiveParseReady(lang) {
		t.Skip("leaven Subtree tagged-pointer crash in core parse")
	}
	p := grammar.NewParser()
	if !p.SetLanguage(lang) {
		t.Fatal("SetLanguage failed")
	}
	src := []byte(`{"a": 1, "b": 2}`)
	tree := p.ParseBytes(src)
	root := tree.RootNode()
	if root.IsNull() {
		t.Fatal("null root")
	}
	q, err := grammar.NewQuery(lang, "(pair key: (string) @k)")
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	errCh := make(chan string, 32)
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 15; j++ {
				// Mix concurrent walks and query on shared parser tree + query.
				_ = root.Type()
				matches := q.ExecuteMatches(root, src)
				if len(matches) == 0 {
					errCh <- "expected matches"
					return
				}
			}
		}()
	}
	wg.Wait()
	close(errCh)
	for e := range errCh {
		t.Fatal(e)
	}
}
