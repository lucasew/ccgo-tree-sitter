package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasew/ccgo-tree-sitter/grammar"

	// Register all grammars here (anonymous imports)
	_ "github.com/lucasew/ccgo-tree-sitter/grammar/json"
	_ "github.com/lucasew/ccgo-tree-sitter/grammar/lua"
	_ "github.com/lucasew/ccgo-tree-sitter/grammar/svelte"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	filename := os.Args[1]

	// Read file
	source, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Detect language
	lang, ok := grammar.GetByExtension(filename)
	if !ok {
		// Try by extension
		ext := strings.TrimPrefix(filepath.Ext(filename), ".")
		fmt.Fprintf(os.Stderr, "Unsupported file extension: .%s\n", ext)
		fmt.Fprintf(os.Stderr, "Supported languages: %s\n", grammar.SupportedLanguages())
		os.Exit(1)
	}
	fmt.Printf("Using grammar: %p\n", lang)

	// Create parser
	parser := grammar.NewParser()
	defer parser.Delete()
	if !parser.SetLanguage(lang) {
		fmt.Fprintf(os.Stderr, "Failed to set language\n")
		os.Exit(1)
	}

	// Parse
	tree := parser.ParseString(string(source))
	defer tree.Delete()

	// Print tree
	root := tree.RootNode()
	printNode(root, source, "", "")
}

func printNode(n *grammar.Node, source []byte, indent string, fieldName string) {
	if n.IsNull() {
		return
	}

	typeStr := n.Type()
	start := n.StartByte()
	end := n.EndByte()

	prefix := ""
	if fieldName != "" {
		prefix = fieldName + ": "
	}

	fmt.Printf("%s%s%s [%d-%d]", indent, prefix, typeStr, start, end)

	if n.ChildCount() == 0 {
		fmt.Printf(" %q", string(source[start:end]))
	}
	fmt.Println()

	count := n.ChildCount()
	for i := uint32(0); i < count; i++ {
		child := n.Child(i)
		field := n.FieldNameForChild(i)
		printNode(child, source, indent+"  ", field)
	}
}

func printUsage() {
	fmt.Println("Usage: parse <file>")
	fmt.Println("\nSupported languages:")
	fmt.Println(grammar.SupportedLanguages())
}
