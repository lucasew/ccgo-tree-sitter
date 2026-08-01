# Freestanding grammar codegen

## Goal

Convert original tree-sitter grammar C (`parser.c` / optional `scanner.c`) to Go
via **ccgo**, with less host-header noise:

- **Smaller** bindings (no host libc const dump per language)
- **Fewer failures** (pure parsers never touch MinGW/Darwin headers)
- **Faster CI** (skip when input hash matches)

Still **ccgo** end-to-end. Same output names as before:
`grammar/<lang>/grammar-$GOOS-$GOARCH.go`.

Platform runtime stays separate: `grammar/core-$GOOS-$GOARCH.go` from core `lib.c`.

## Flow

```text
classify unit
  pure     → parser.c only
  scanner  → parser.c + scanner.c

hash inputs (parser, scanner, tree_sitter/*.h, freestanding stubs)
  → if grammar-$GOOS-$GOARCH.go has matching // ts-grammar-hash:  → skip

stage freestanding unit
  rewrite #include <stdint.h> → "stdint.h" (and friends)
  -I cmd/codegen/freestanding stubs
  -I staged grammar headers

ccgo (freestanding)
  on failure → host-header fallback (classic path)

postProcess + slim host consts + write grammar-$GOOS-$GOARCH.go
```

```bash
mise run codegen:linux-amd64   # go run ./cmd/codegen
```

## Cases

| Case | Detection | Path |
| --- | --- | --- |
| Pure parser | no `scanner.c` | freestanding stubs |
| External scanner | `scanner.c` | freestanding + malloc/string/wctype stubs |
| Monorepo includes | sibling `common/`, parent dir | extra `-I` |
| Odd system headers | freestanding fails | **host fallback** |
| Huge / known-bad | size / skip list | skip (existing `fsharp` etc.) |

## Verify

```bash
go test ./cmd/codegen/ -count=1
# with a preprocessor on PATH (optional smoke):
FORCE_CODEGEN_TEST=1 go test ./cmd/codegen/ -run TestFreestandingJSONTranspile -v
```
