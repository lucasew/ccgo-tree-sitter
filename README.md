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

## Grammar sources

Grammar C inputs are **not** git submodules. They are declared in `workspaced.cue`
(`#grammar`) and placed with workspaced `core:place` (only `src/` / monorepo units):

```bash
mise run grammars:sync            # workspaced codebase apply
# or: workspaced mod lock && workspaced codebase apply
./setup-grammars owner/tree-sitter-foo   # add a #grammar entry
```

Placed trees under `third-party/tree-sitter-*/` are gitignored; pins are in
`workspaced.lock.json`. Core tree-sitter stays a submodule at `third-party/tree-sitter`.

## Codegen

```bash
mise run grammars:sync            # first time / after lock changes
mise run codegen                  # host GOOS/GOARCH only
mise run codegen:darwin-arm64     # explicit triple
```

Prefer the **CI matrix** (each runner owns its native triple; merge opens a PR).

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
