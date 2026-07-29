// This go.mod exists only to prevent the Go toolchain from treating
// any subdirectory under third-party/ as belonging to the parent module
// (github.com/modernc-tree-sitter/ccgo-tree-sitter).
//
// Grammar trees are placed C sources (often without their own go.mod).
// Without this module boundary, `go test ./...` could walk them.
module ignore

go 1.25
