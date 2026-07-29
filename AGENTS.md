# Guidelines
- Don't edit ./grammar directly. Those files are all code generated.
- **Never transpile/codegen locally.** Bindings are produced only on the CI
  matrix (`mise run codegen:$GOOS-$GOARCH` on runners after workspaced place).
- Grammar **C sources** under `third-party/tree-sitter-*/` are materialised by
  workspaced (`core:place`), not git submodules. Declare them in `workspaced.cue`
  (`#grammar: <name>: { from, repo, paths? }`), then `mise run grammars:lock` /
  `mise run grammars:sync` (mise installs go + workspaced).
- **Core tree-sitter** is a workspaced *source* only (`inputs.tree_sitter` /
  `#tree_sitter`) — not placed into the repo. Pins in `workspaced.lock.json`;
  on-disk path is the workspaced github cache (`mise run tree-sitter:path` or
  `TREE_SITTER_PATH`). Codegen resolves that path automatically.
- Don't commit placed trees; they are gitignored. Pins live in `workspaced.lock.json`.
- Codegen preprocesses with clang by default; on Windows prefer MinGW gcc on PATH.
- Do not add MSVC-header regex sanitizers; use MinGW gcc -E and ccgo ignore flags.
- `ccgo` / `libc` are normal Go module deps (forks via `replace` to
  `github.com/modernc-tree-sitter/...`), not third-party submodules.
