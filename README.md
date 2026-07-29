# ccgo-tree-sitter
Tree sitter in Go using a patched modernc.org ccgo

Extremely experimental, but it can parse some stuff already.

## Test

```bash
mise run test                     # go test ./... (CGO_ENABLED=0)
```

CI runs tests on every push and pull request across the full native target matrix.
Platforms without checked-in `grammar/core-*.go` bindings codegen on the runner
first, then run `mise run test`.

| `GOOS`/`GOARCH` | CI runner | Checked in | Notes |
| --- | --- | --- | --- |
| linux/amd64 | `ubuntu-latest` | yes | |
| linux/arm64 | `ubuntu-24.04-arm` | yes | |
| linux/386 | `ubuntu-latest` | partial | experimental |
| darwin/amd64 | `macos-15-intel` | no | experimental |
| darwin/arm64 | `macos-latest` | no | erase availability attrs; ccgo ignores align-16 |
| windows/amd64 | `windows-latest` | no | **MinGW `gcc -E`** (not clang/MSVC) |
| windows/arm64 | `windows-11-arm` | no | experimental; llvm-mingw `aarch64-w64-mingw32` |

## Grammar sources (local + CI)

Grammar C inputs are **not** git submodules. They are declared in `workspaced.cue`
(`#grammar`) and placed with workspaced `core:place` (only `src/` / monorepo units).
**mise** installs both **go** and **workspaced**:

```bash
mise install                      # go + workspaced from mise.toml
mise run grammars:sync            # workspaced codebase apply
./setup-grammars owner/tree-sitter-foo   # add a #grammar entry
mise run grammars:lock            # refresh workspaced.lock.json
mise run test                     # pure Go tests (no transpile)
```

Grammar trees under `third-party/tree-sitter-*/` are **vendored** (committed);
pins stay in `workspaced.lock.json`. Re-sync after changing `#grammar` entries.
Do not vendor full application checkouts (`tensorflow`, `django`) — those are
optional local corpora only and remain gitignored.

Core tree-sitter is a workspaced **source** (not placed into the repo):

```bash
mise run tree-sitter:path   # print/ensure ~/.cache/workspaced/sources/github/…
# or: TREE_SITTER_PATH=/path/to/tree-sitter  (override for codegen)
```

Parse fixtures (astdump goldens) — one minimal sample per generated language:

```text
testdata/<language>/<file.ext>
testdata/<language>/<file.ext>.golden.json
```

Sources are declared in `cmd/parse/fixture_sources.go`. Goldens are `grammar.ParseOutput` JSON.

```bash
UPDATE_GOLDENS=1 go test ./cmd/parse/ -run TestLanguageFixtures   # rewrite sources + goldens
go test ./cmd/parse/ -run TestLanguageFixtures                       # compare goldens
```

Each language is its own `#grammar: name: { … }` in `workspaced.cue` (optional `astdump` metadata).

## Codegen (CI only)

**Do not transpile locally.** Each `codegen:*` mise task depends on
`grammars:sync` (workspaced) and is intended for CI runners with the right
preprocessor. Multi-platform regen is the GitHub Actions matrix on `main`
(merge opens a PR with regenerated bindings).

### Target matrix

| Task | `GOOS`/`GOARCH` | CI runner | Notes |
| --- | --- | --- | --- |
| `codegen:linux-amd64` | linux/amd64 | `ubuntu-latest` | checked in |
| `codegen:linux-arm64` | linux/arm64 | `ubuntu-24.04-arm` | checked in |
| `codegen:linux-386` | linux/386 | `ubuntu-latest` | experimental |
| `codegen:darwin-amd64` | darwin/amd64 | `macos-15-intel` | experimental |
| `codegen:darwin-arm64` | darwin/arm64 | `macos-latest` | erase availability attrs; ccgo ignores align-16 |
| `codegen:windows-amd64` | windows/amd64 | `windows-latest` | **MinGW `gcc -E`** (not clang/MSVC) |
| `codegen:windows-arm64` | windows/arm64 | `windows-11-arm` | experimental; llvm-mingw |

Outputs: `grammar/core-{GOOS}-{GOARCH}.go`, `grammar/<lang>/grammar-{GOOS}-{GOARCH}.go`.

On Windows, codegen prefers `x86_64-w64-mingw32-gcc` / `aarch64-w64-mingw32-gcc` /
`gcc` from MinGW on `PATH` even if `CC=clang`, and passes
`-ignore-unsupported-alignment` (and friends) to ccgo. On Darwin, availability
macros are defined empty and the same ccgo ignores apply. Linux still drops
`_Float*` typedefs after `-E`.
