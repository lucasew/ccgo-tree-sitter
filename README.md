# ccgo-tree-sitter

Tree-sitter in pure Go via **clang → LLVM IR → [leaven](https://github.com/lewtec/leaven)** (no CGO, no ccgo).

Extremely experimental mid-migration: leaven output is raw and often not yet valid Go; the hand-written `grammar` API still needs rewiring onto leaven symbols.

## Tools

```bash
mise install                 # go, workspaced, conda:clang 14
export CC="$(mise which clang)"
```

## Grammar sources

```bash
mise run grammars:sync       # place vendored third-party C
mise run tree-sitter:path    # core tree-sitter cache
```

## Codegen (leaven)

```bash
mise run codegen             # core (best-effort) + all grammars
go run ./cmd/codegen --only=json,python
```

Outputs:

- `grammar/core.go` (when leaven accepts core IR)
- `grammar/<lang>/grammar.go` — raw leaven
- `grammar/<lang>/api.go` — stub package

`leaven` is a module tool: `go tool leaven` (declared in `go.mod`).

## Test

```bash
mise run test
```

Expect failures until leaven output and the API layer converge.
