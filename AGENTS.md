# Guidelines
- Don't edit `grammar/core.go` or `grammar/<lang>/grammar.go` by hand — codegen only.
- Codegen is **only** clang 14 → LLVM IR → `go tool leaven`. Pin clang via mise `conda:clang@14.0.6`.
- Output is single-file pure Go: `grammar/core.go` and `grammar/<lang>/grammar.go` (no platform splits).
- Grammar **C sources** under `third-party/tree-sitter-*/` are **vendored** (committed). Declared in `workspaced.cue` (`#grammar`), pinned in `workspaced.lock.json`, refreshed with `mise run grammars:lock` / `mise run grammars:sync`.
- **Core tree-sitter** is a workspaced *source* only (`inputs.tree_sitter` / `#tree_sitter`) — not placed into the repo. Pins in `workspaced.lock.json`; on-disk path is the workspaced github cache (`mise run tree-sitter:path` or `TREE_SITTER_PATH`).
- Run codegen: `mise run codegen` (or `go run ./cmd/codegen` / `--only=…`).
- `leaven` is a Go module tool (`go get -tool` / `go tool leaven`), not a global install.
  Source of truth: **`github.com/lewtec/leaven`** (replace of `github.com/andybalholm/leaven` module path).
- Do not depend on `modernc.org/ccgo`, `modernc.org/libc`, or any CGO/platform-split transpiler path.
