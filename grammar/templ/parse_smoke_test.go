package grammar_templ_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
	_ "github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/templ"
)

func dump(n *grammar.Node, src string, ind int, b *strings.Builder) {
	if n == nil || n.IsNull() {
		return
	}
	leaf := ""
	if n.ChildCount() == 0 {
		leaf = fmt.Sprintf(" %q", src[n.StartByte():n.EndByte()])
	}
	fmt.Fprintf(b, "%s%s%s [%d:%d]\n", strings.Repeat("  ", ind), n.Type(), leaf, n.StartByte(), n.EndByte())
	for i := uint32(0); i < n.ChildCount(); i++ {
		dump(n.Child(i), src, ind+1, b)
	}
}

func TestTemplParseSmoke(t *testing.T) {
	src := `package views

import (
	"fmt"
	annotate "github.com/lucasew/refactree/pkg/web/annotate"
)

templ CodeSegments(segments []annotate.Segment) {
	<pre class="code">
		for _, s := range segments {
			@codeSegment(s)
		}
	</pre>
}

templ codeSegment(s annotate.Segment) {
	if s.IsLink {
		<a href={ templ.URL(s.Href) }>{ s.Text }</a>
	} else {
		{ s.Text }
	}
}

func SpanFragmentID(start, end uint32) string {
	return fmt.Sprintf("%d-%d", start, end)
}
`
	lang, ok := grammar.Get("templ")
	if !ok {
		t.Fatal("templ not registered")
	}
	p := grammar.NewParser()
	defer p.Delete()
	if !p.SetLanguage(lang) {
		t.Fatal("SetLanguage")
	}
	tree := p.ParseString(src)
	defer tree.Delete()
	var b strings.Builder
	dump(tree.RootNode(), src, 0, &b)
	t.Log("\n" + b.String())
	// always pass - dump only
}
