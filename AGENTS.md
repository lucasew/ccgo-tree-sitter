# Guidelines
- Don't edit ./grammar directly. Those files are all code generated.
- Grammar **C sources** under `third-party/tree-sitter-*/` are materialised by
  workspaced (`core:place`), not git submodules. Declare them in `workspaced.cue`
  (`#grammar: <name>: { from, repo, paths? }`), then
  `workspaced mod lock` and `mise run grammars:sync` (or `workspaced codebase apply`).
- Don't commit placed trees; they are gitignored. Pins live in `workspaced.lock.json`.
- Codegen preprocesses with clang by default; on Windows prefer MinGW gcc on PATH.
- `mise run codegen` is host-native only. Multi-platform regen is CI matrix + merge PR.
- Do not add MSVC-header regex sanitizers; use MinGW gcc -E and ccgo ignore flags.
- `ccgo` / `libc` are normal Go module deps (forks via `replace` to
  `github.com/modernc-tree-sitter/...`), not third-party submodules.
