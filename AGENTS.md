# Guidelines
- Don't edit `grammar/core.go` or `grammar/<lang>/grammar.go` by hand — codegen (leaven).
- Codegen is **clang 14 → LLVM IR → go tool leaven** (not ccgo). Pin clang via mise `conda:clang@14.0.6`.
- Grammar **C sources** under `third-party/tree-sitter-*/` are **vendored** (committed). Declared in `workspaced.cue` (`#grammar`), pinned in `workspaced.lock.json`, refreshed with `mise run grammars:lock` / `mise run grammars:sync`.
- **Core tree-sitter** is a workspaced *source* only (`inputs.tree_sitter` / `#tree_sitter`) — not placed into the repo. Pins in `workspaced.lock.json`; on-disk path is the workspaced github cache (`mise run tree-sitter:path` or `TREE_SITTER_PATH`).
- Run codegen: `mise run codegen` (or `go run ./cmd/codegen [lang...]`). No multi-GOOS matrix.
- `leaven` is a Go module tool (`go get -tool` / `go tool leaven`), not a global install.
- Do not reintroduce modernc.org/ccgo or platform-split `grammar-*-*.go` / `core-*-*.go`.
