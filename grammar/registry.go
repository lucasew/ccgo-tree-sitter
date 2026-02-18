package grammar

import (
	"path/filepath"
	"strings"
	"sync"
)

var (
	registry = make(map[string]Language)
	mu       sync.RWMutex
)

// Register adds a new grammar to the registry
func Register(name string, lang Language) {
	mu.Lock()
	defer mu.Unlock()
	registry[name] = lang
}

// Get retrieves a grammar by name
func Get(name string) (Language, bool) {
	mu.RLock()
	defer mu.RUnlock()
	lang, ok := registry[name]
	return lang, ok
}

// List returns all registered grammar names
func List() []string {
	mu.RLock()
	defer mu.RUnlock()
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	return names
}

// GetByExtension tries to find a grammar based on file extension
func GetByExtension(filename string) (Language, bool) {
	ext := strings.TrimPrefix(filepath.Ext(filename), ".")

	// Direct match
	if lang, ok := Get(ext); ok {
		return lang, ok
	}

	// Common mappings
	mapping := map[string]string{
		"js":   "javascript",
		"ts":   "typescript",
		"py":   "python",
		"rb":   "ruby",
		"rs":   "rust",
		"go":   "go",
		"c":    "c",
		"cpp":  "cpp",
		"h":    "c",
		"hpp":  "cpp",
		"lua":  "lua",
		"json": "json",
	}

	if name, ok := mapping[ext]; ok {
		return Get(name)
	}

	return nil, false
}

func SupportedLanguages() string {
	langs := List()
	if len(langs) == 0 {
		return "none"
	}
	return strings.Join(langs, ", ")
}
