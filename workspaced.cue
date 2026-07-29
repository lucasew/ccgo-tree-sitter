package workspaced

// Grammar C sources only (src/ + monorepo units). Materialised by core:place.
// Add a language: one field under #grammar, then:
//   workspaced mod lock && workspaced codebase apply
#grammar: {}

#grammar: [string]: {
	from:    string
	version: string | *"HEAD"
	paths:   [...string] | *["src"]
	// Dest basename under third-party/ (usually tree-sitter-<key>).
	repo:    string
}

// Test corpora (not codegen). Re-add with narrow origin when tests need them.
#fixture: {}

#fixture: [string]: {
	from:    string
	version: string | *"HEAD"
	// Path inside the input; "." places the whole tree at dest.
	origin:  string | *"."
	dest:    string
}

workspaced: {
	inputs: {
		for name, g in #grammar {
			"grammar_\(name)": {
				from:    g.from
				version: g.version
			}
		}
		for name, f in #fixture {
			"fixture_\(name)": {
				from:    f.from
				version: f.version
			}
		}
	}
	modules: {
		// One core:place module per grammar entry.
		for name, g in #grammar {
			"grammar_\(name)": {
				from: "core:place"
				config: {
					// Skip grammars that do not ship pregenerated src/parser.c.
					ignore_missing: true
					items: {
						for p in g.paths {
							"third-party/\(g.repo)/\(p)": "grammar_\(name):\(p)"
						}
					}
				}
			}
		}
		// One core:place module per fixture entry.
		for name, f in #fixture {
			"fixture_\(name)": {
				from: "core:place"
				config: {
					items: {
						"\(f.dest)": "fixture_\(name):\(f.origin)"
					}
				}
			}
		}
	}
}
