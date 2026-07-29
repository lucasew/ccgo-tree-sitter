#!/usr/bin/env bash
# Codegen is CI-only. Local hosts must not transpile.
set -euo pipefail
echo "error: do not run grammar codegen/transpile locally." >&2
echo "Grammar C sources: mise run grammars:sync (workspaced core:place)." >&2
echo "Go bindings: CI matrix on main (mise run codegen:\$GOOS-\$GOARCH on runners)." >&2
exit 1
