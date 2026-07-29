package workspaced

// Grammar C sources only (src/ + monorepo units). Materialised by core:place.
// Add a language: one field under #grammar, then:
//   workspaced mod lock && workspaced codebase apply
#grammar: {
	ada: {
		from: "github:briot/tree-sitter-ada"
		repo: "tree-sitter-ada"
	}
	agda: {
		from: "github:tree-sitter/tree-sitter-agda"
		repo: "tree-sitter-agda"
	}
	angular: {
		from: "github:dlvandenberg/tree-sitter-angular"
		repo: "tree-sitter-angular"
	}
	arduino: {
		from: "github:tree-sitter-grammars/tree-sitter-arduino"
		repo: "tree-sitter-arduino"
	}
	asm: {
		from: "github:RubixDev/tree-sitter-asm"
		repo: "tree-sitter-asm"
	}
	astro: {
		from: "github:virchau13/tree-sitter-astro"
		repo: "tree-sitter-astro"
	}
	bash: {
		from: "github:tree-sitter/tree-sitter-bash"
		repo: "tree-sitter-bash"
	}
	beancount: {
		from: "github:polarmutex/tree-sitter-beancount"
		repo: "tree-sitter-beancount"
	}
	bibtex: {
		from: "github:latex-lsp/tree-sitter-bibtex"
		repo: "tree-sitter-bibtex"
	}
	bicep: {
		from: "github:tree-sitter-grammars/tree-sitter-bicep"
		repo: "tree-sitter-bicep"
	}
	bitbake: {
		from: "github:tree-sitter-grammars/tree-sitter-bitbake"
		repo: "tree-sitter-bitbake"
	}
	blade: {
		from: "github:EmranMR/tree-sitter-blade"
		repo: "tree-sitter-blade"
	}
	c: {
		from: "github:tree-sitter/tree-sitter-c"
		repo: "tree-sitter-c"
	}
	"c-sharp": {
		from: "github:tree-sitter/tree-sitter-c-sharp"
		repo: "tree-sitter-c-sharp"
	}
	cairo: {
		from: "github:tree-sitter-grammars/tree-sitter-cairo"
		repo: "tree-sitter-cairo"
	}
	capnp: {
		from: "github:tree-sitter-grammars/tree-sitter-capnp"
		repo: "tree-sitter-capnp"
	}
	chatito: {
		from: "github:tree-sitter-grammars/tree-sitter-chatito"
		repo: "tree-sitter-chatito"
	}
	clojure: {
		from: "github:sogaiu/tree-sitter-clojure"
		repo: "tree-sitter-clojure"
	}
	cmake: {
		from: "github:uyha/tree-sitter-cmake"
		repo: "tree-sitter-cmake"
	}
	cobol: {
		from: "github:yutaro-sakamoto/tree-sitter-cobol"
		repo: "tree-sitter-cobol"
	}
	comment: {
		from: "github:stsewd/tree-sitter-comment"
		repo: "tree-sitter-comment"
	}
	commonlisp: {
		from: "github:theHamsta/tree-sitter-commonlisp"
		repo: "tree-sitter-commonlisp"
	}
	cpon: {
		from: "github:tree-sitter-grammars/tree-sitter-cpon"
		repo: "tree-sitter-cpon"
	}
	cpp: {
		from: "github:tree-sitter/tree-sitter-cpp"
		repo: "tree-sitter-cpp"
	}
	crystal: {
		from: "github:crystal-lang-tools/tree-sitter-crystal"
		repo: "tree-sitter-crystal"
	}
	css: {
		from: "github:tree-sitter/tree-sitter-css"
		repo: "tree-sitter-css"
	}
	cst: {
		from: "github:tree-sitter-grammars/tree-sitter-cst"
		repo: "tree-sitter-cst"
	}
	csv: {
		from: "github:tree-sitter-grammars/tree-sitter-csv"
		repo: "tree-sitter-csv"
		paths: ["csv/src", "tsv/src", "psv/src", "common"]
	}
	cuda: {
		from: "github:tree-sitter-grammars/tree-sitter-cuda"
		repo: "tree-sitter-cuda"
	}
	cue: {
		from: "github:eonpatapon/tree-sitter-cue"
		repo: "tree-sitter-cue"
	}
	cyberchef: {
		from: "github:tree-sitter-grammars/tree-sitter-cyberchef"
		repo: "tree-sitter-cyberchef"
	}
	dart: {
		from: "github:UserNobody14/tree-sitter-dart"
		repo: "tree-sitter-dart"
	}
	diff: {
		from: "github:tree-sitter-grammars/tree-sitter-diff"
		repo: "tree-sitter-diff"
	}
	dockerfile: {
		from: "github:camdencheek/tree-sitter-dockerfile"
		repo: "tree-sitter-dockerfile"
	}
	doxygen: {
		from: "github:tree-sitter-grammars/tree-sitter-doxygen"
		repo: "tree-sitter-doxygen"
	}
	eex: {
		from: "github:connorlay/tree-sitter-eex"
		repo: "tree-sitter-eex"
	}
	elixir: {
		from: "github:elixir-lang/tree-sitter-elixir"
		repo: "tree-sitter-elixir"
	}
	elm: {
		from: "github:elm-tooling/tree-sitter-elm"
		repo: "tree-sitter-elm"
	}
	"embedded-template": {
		from: "github:tree-sitter/tree-sitter-embedded-template"
		repo: "tree-sitter-embedded-template"
	}
	erlang: {
		from: "github:WhatsApp/tree-sitter-erlang"
		repo: "tree-sitter-erlang"
	}
	firrtl: {
		from: "github:tree-sitter-grammars/tree-sitter-firrtl"
		repo: "tree-sitter-firrtl"
	}
	fish: {
		from: "github:ram02z/tree-sitter-fish"
		repo: "tree-sitter-fish"
	}
	fluent: {
		from: "github:tree-sitter/tree-sitter-fluent"
		repo: "tree-sitter-fluent"
	}
	fortran: {
		from: "github:stadelmanma/tree-sitter-fortran"
		repo: "tree-sitter-fortran"
	}
	fsharp: {
		from: "github:ionide/tree-sitter-fsharp"
		repo: "tree-sitter-fsharp"
		paths: ["fsharp/src", "fsharp_signature/src", "common"]
	}
	func: {
		from: "github:tree-sitter-grammars/tree-sitter-func"
		repo: "tree-sitter-func"
	}
	gdscript: {
		from: "github:PrestonKnopp/tree-sitter-gdscript"
		repo: "tree-sitter-gdscript"
	}
	"git-config": {
		from: "github:the-mikedavis/tree-sitter-git-config"
		repo: "tree-sitter-git-config"
	}
	"git-rebase": {
		from: "github:the-mikedavis/tree-sitter-git-rebase"
		repo: "tree-sitter-git-rebase"
	}
	gitattributes: {
		from: "github:tree-sitter-grammars/tree-sitter-gitattributes"
		repo: "tree-sitter-gitattributes"
	}
	gitcommit: {
		from: "github:gbprod/tree-sitter-gitcommit"
		repo: "tree-sitter-gitcommit"
	}
	gitignore: {
		from: "github:shunsambongi/tree-sitter-gitignore"
		repo: "tree-sitter-gitignore"
	}
	gleam: {
		from: "github:gleam-lang/tree-sitter-gleam"
		repo: "tree-sitter-gleam"
	}
	glsl: {
		from: "github:tree-sitter-grammars/tree-sitter-glsl"
		repo: "tree-sitter-glsl"
	}
	gn: {
		from: "github:tree-sitter-grammars/tree-sitter-gn"
		repo: "tree-sitter-gn"
	}
	go: {
		from: "github:tree-sitter/tree-sitter-go"
		repo: "tree-sitter-go"
	}
	"go-sum": {
		from: "github:tree-sitter-grammars/tree-sitter-go-sum"
		repo: "tree-sitter-go-sum"
	}
	"godot-resource": {
		from: "github:PrestonKnopp/tree-sitter-godot-resource"
		repo: "tree-sitter-godot-resource"
	}
	"gpg-config": {
		from: "github:tree-sitter-grammars/tree-sitter-gpg-config"
		repo: "tree-sitter-gpg-config"
	}
	graph: {
		from: "github:tree-sitter/tree-sitter-graph"
		repo: "tree-sitter-graph"
	}
	graphql: {
		from: "github:bkegley/tree-sitter-graphql"
		repo: "tree-sitter-graphql"
	}
	groovy: {
		from: "github:murtaza64/tree-sitter-groovy"
		repo: "tree-sitter-groovy"
	}
	gstlaunch: {
		from: "github:tree-sitter-grammars/tree-sitter-gstlaunch"
		repo: "tree-sitter-gstlaunch"
	}
	hare: {
		from: "github:tree-sitter-grammars/tree-sitter-hare"
		repo: "tree-sitter-hare"
	}
	haskell: {
		from: "github:tree-sitter/tree-sitter-haskell"
		repo: "tree-sitter-haskell"
	}
	hcl: {
		from: "github:tree-sitter-grammars/tree-sitter-hcl"
		repo: "tree-sitter-hcl"
		paths: ["src", "dialects/terraform/src"]
	}
	heex: {
		from: "github:phoenixframework/tree-sitter-heex"
		repo: "tree-sitter-heex"
	}
	hlsl: {
		from: "github:tree-sitter-grammars/tree-sitter-hlsl"
		repo: "tree-sitter-hlsl"
	}
	html: {
		from: "github:tree-sitter/tree-sitter-html"
		repo: "tree-sitter-html"
	}
	http: {
		from: "github:rest-nvim/tree-sitter-http"
		repo: "tree-sitter-http"
	}
	hyprlang: {
		from: "github:tree-sitter-grammars/tree-sitter-hyprlang"
		repo: "tree-sitter-hyprlang"
	}
	idris: {
		from: "github:kayhide/tree-sitter-idris"
		repo: "tree-sitter-idris"
	}
	ini: {
		from: "github:justinmk/tree-sitter-ini"
		repo: "tree-sitter-ini"
	}
	ispc: {
		from: "github:tree-sitter-grammars/tree-sitter-ispc"
		repo: "tree-sitter-ispc"
	}
	java: {
		from: "github:tree-sitter/tree-sitter-java"
		repo: "tree-sitter-java"
	}
	javascript: {
		from: "github:tree-sitter/tree-sitter-javascript"
		repo: "tree-sitter-javascript"
	}
	jsdoc: {
		from: "github:tree-sitter/tree-sitter-jsdoc"
		repo: "tree-sitter-jsdoc"
	}
	json: {
		from: "github:tree-sitter/tree-sitter-json"
		repo: "tree-sitter-json"
	}
	json5: {
		from: "github:Joakker/tree-sitter-json5"
		repo: "tree-sitter-json5"
	}
	julia: {
		from: "github:tree-sitter/tree-sitter-julia"
		repo: "tree-sitter-julia"
	}
	just: {
		from: "github:casey/tree-sitter-just"
		repo: "tree-sitter-just"
	}
	kconfig: {
		from: "github:tree-sitter-grammars/tree-sitter-kconfig"
		repo: "tree-sitter-kconfig"
	}
	kdl: {
		from: "github:tree-sitter-grammars/tree-sitter-kdl"
		repo: "tree-sitter-kdl"
	}
	kotlin: {
		from: "github:fwcd/tree-sitter-kotlin"
		repo: "tree-sitter-kotlin"
	}
	latex: {
		from: "github:latex-lsp/tree-sitter-latex"
		repo: "tree-sitter-latex"
	}
	linkerscript: {
		from: "github:tree-sitter-grammars/tree-sitter-linkerscript"
		repo: "tree-sitter-linkerscript"
	}
	llvm: {
		from: "github:benwilliamgraham/tree-sitter-llvm"
		repo: "tree-sitter-llvm"
	}
	lua: {
		from: "github:tree-sitter-grammars/tree-sitter-lua"
		repo: "tree-sitter-lua"
	}
	luadoc: {
		from: "github:tree-sitter-grammars/tree-sitter-luadoc"
		repo: "tree-sitter-luadoc"
	}
	luap: {
		from: "github:tree-sitter-grammars/tree-sitter-luap"
		repo: "tree-sitter-luap"
	}
	luau: {
		from: "github:tree-sitter-grammars/tree-sitter-luau"
		repo: "tree-sitter-luau"
	}
	make: {
		from: "github:tree-sitter-grammars/tree-sitter-make"
		repo: "tree-sitter-make"
	}
	markdown: {
		from: "github:tree-sitter-grammars/tree-sitter-markdown"
		repo: "tree-sitter-markdown"
		paths: ["tree-sitter-markdown/src", "tree-sitter-markdown-inline/src", "common"]
	}
	matlab: {
		from: "github:acristoffers/tree-sitter-matlab"
		repo: "tree-sitter-matlab"
	}
	meson: {
		from: "github:tree-sitter-grammars/tree-sitter-meson"
		repo: "tree-sitter-meson"
	}
	move: {
		from: "github:tree-sitter-grammars/tree-sitter-move"
		repo: "tree-sitter-move"
	}
	nasm: {
		from: "github:naclsn/tree-sitter-nasm"
		repo: "tree-sitter-nasm"
	}
	nginx: {
		from: "github:opa-oz/tree-sitter-nginx"
		repo: "tree-sitter-nginx"
	}
	nickel: {
		from: "github:nickel-lang/tree-sitter-nickel"
		repo: "tree-sitter-nickel"
	}
	nim: {
		from: "github:alaviss/tree-sitter-nim"
		repo: "tree-sitter-nim"
	}
	nix: {
		from: "github:nix-community/tree-sitter-nix"
		repo: "tree-sitter-nix"
	}
	nqc: {
		from: "github:tree-sitter-grammars/tree-sitter-nqc"
		repo: "tree-sitter-nqc"
	}
	objc: {
		from: "github:tree-sitter-grammars/tree-sitter-objc"
		repo: "tree-sitter-objc"
	}
	objdump: {
		from: "github:ColinKennedy/tree-sitter-objdump"
		repo: "tree-sitter-objdump"
	}
	ocaml: {
		from: "github:tree-sitter/tree-sitter-ocaml"
		repo: "tree-sitter-ocaml"
		paths: ["grammars/ocaml/src", "grammars/interface/src", "grammars/type/src", "common"]
	}
	odin: {
		from: "github:tree-sitter-grammars/tree-sitter-odin"
		repo: "tree-sitter-odin"
	}
	pascal: {
		from: "github:Isopod/tree-sitter-pascal"
		repo: "tree-sitter-pascal"
	}
	pem: {
		from: "github:tree-sitter-grammars/tree-sitter-pem"
		repo: "tree-sitter-pem"
	}
	perl: {
		from: "github:tree-sitter-perl/tree-sitter-perl"
		repo: "tree-sitter-perl"
	}
	php: {
		from: "github:tree-sitter/tree-sitter-php"
		repo: "tree-sitter-php"
		paths: ["php/src", "php_only/src", "common"]
	}
	pkl: {
		from: "github:apple/tree-sitter-pkl"
		repo: "tree-sitter-pkl"
	}
	po: {
		from: "github:tree-sitter-grammars/tree-sitter-po"
		repo: "tree-sitter-po"
	}
	"poe-filter": {
		from: "github:tree-sitter-grammars/tree-sitter-poe-filter"
		repo: "tree-sitter-poe-filter"
	}
	pony: {
		from: "github:tree-sitter-grammars/tree-sitter-pony"
		repo: "tree-sitter-pony"
	}
	powershell: {
		from: "github:airbus-cert/tree-sitter-powershell"
		repo: "tree-sitter-powershell"
	}
	printf: {
		from: "github:tree-sitter-grammars/tree-sitter-printf"
		repo: "tree-sitter-printf"
	}
	prisma: {
		from: "github:victorhqc/tree-sitter-prisma"
		repo: "tree-sitter-prisma"
	}
	properties: {
		from: "github:tree-sitter-grammars/tree-sitter-properties"
		repo: "tree-sitter-properties"
	}
	proto: {
		from: "github:mitchellh/tree-sitter-proto"
		repo: "tree-sitter-proto"
	}
	puppet: {
		from: "github:tree-sitter-grammars/tree-sitter-puppet"
		repo: "tree-sitter-puppet"
	}
	purescript: {
		from: "github:postsolar/tree-sitter-purescript"
		repo: "tree-sitter-purescript"
	}
	pymanifest: {
		from: "github:tree-sitter-grammars/tree-sitter-pymanifest"
		repo: "tree-sitter-pymanifest"
	}
	python: {
		from: "github:tree-sitter/tree-sitter-python"
		repo: "tree-sitter-python"
	}
	ql: {
		from: "github:tree-sitter/tree-sitter-ql"
		repo: "tree-sitter-ql"
	}
	"ql-dbscheme": {
		from: "github:tree-sitter/tree-sitter-ql-dbscheme"
		repo: "tree-sitter-ql-dbscheme"
	}
	qmldir: {
		from: "github:tree-sitter-grammars/tree-sitter-qmldir"
		repo: "tree-sitter-qmldir"
	}
	query: {
		from: "github:tree-sitter-grammars/tree-sitter-query"
		repo: "tree-sitter-query"
	}
	racket: {
		from: "github:6cdh/tree-sitter-racket"
		repo: "tree-sitter-racket"
	}
	rasi: {
		from: "github:Fymyte/tree-sitter-rasi"
		repo: "tree-sitter-rasi"
	}
	re2c: {
		from: "github:tree-sitter-grammars/tree-sitter-re2c"
		repo: "tree-sitter-re2c"
	}
	readline: {
		from: "github:tree-sitter-grammars/tree-sitter-readline"
		repo: "tree-sitter-readline"
	}
	regex: {
		from: "github:tree-sitter/tree-sitter-regex"
		repo: "tree-sitter-regex"
	}
	rego: {
		from: "github:FallenAngel97/tree-sitter-rego"
		repo: "tree-sitter-rego"
	}
	requirements: {
		from: "github:tree-sitter-grammars/tree-sitter-requirements"
		repo: "tree-sitter-requirements"
	}
	rescript: {
		from: "github:rescript-lang/tree-sitter-rescript"
		repo: "tree-sitter-rescript"
	}
	roc: {
		from: "github:faldor20/tree-sitter-roc"
		repo: "tree-sitter-roc"
	}
	ron: {
		from: "github:tree-sitter-grammars/tree-sitter-ron"
		repo: "tree-sitter-ron"
	}
	ruby: {
		from: "github:tree-sitter/tree-sitter-ruby"
		repo: "tree-sitter-ruby"
	}
	rust: {
		from: "github:tree-sitter/tree-sitter-rust"
		repo: "tree-sitter-rust"
	}
	scala: {
		from: "github:tree-sitter/tree-sitter-scala"
		repo: "tree-sitter-scala"
	}
	scheme: {
		from: "github:6cdh/tree-sitter-scheme"
		repo: "tree-sitter-scheme"
	}
	scss: {
		from: "github:tree-sitter-grammars/tree-sitter-scss"
		repo: "tree-sitter-scss"
	}
	slang: {
		from: "github:tree-sitter-grammars/tree-sitter-slang"
		repo: "tree-sitter-slang"
	}
	smali: {
		from: "github:tree-sitter-grammars/tree-sitter-smali"
		repo: "tree-sitter-smali"
	}
	smithy: {
		from: "github:indoorvivants/tree-sitter-smithy"
		repo: "tree-sitter-smithy"
	}
	sml: {
		from: "github:MatthewFluet/tree-sitter-sml"
		repo: "tree-sitter-sml"
	}
	solidity: {
		from: "github:JoranHonig/tree-sitter-solidity"
		repo: "tree-sitter-solidity"
	}
	sparql: {
		from: "github:GordianDziwis/tree-sitter-sparql"
		repo: "tree-sitter-sparql"
	}
	sql: {
		from: "github:DerekStride/tree-sitter-sql"
		repo: "tree-sitter-sql"
	}
	squirrel: {
		from: "github:tree-sitter-grammars/tree-sitter-squirrel"
		repo: "tree-sitter-squirrel"
	}
	"ssh-config": {
		from: "github:tree-sitter-grammars/tree-sitter-ssh-config"
		repo: "tree-sitter-ssh-config"
	}
	starlark: {
		from: "github:tree-sitter-grammars/tree-sitter-starlark"
		repo: "tree-sitter-starlark"
	}
	surface: {
		from: "github:connorlay/tree-sitter-surface"
		repo: "tree-sitter-surface"
	}
	svelte: {
		from: "github:tree-sitter-grammars/tree-sitter-svelte"
		repo: "tree-sitter-svelte"
	}
	swift: {
		from: "github:alex-pinkus/tree-sitter-swift"
		repo: "tree-sitter-swift"
	}
	tablegen: {
		from: "github:tree-sitter-grammars/tree-sitter-tablegen"
		repo: "tree-sitter-tablegen"
	}
	tcl: {
		from: "github:tree-sitter-grammars/tree-sitter-tcl"
		repo: "tree-sitter-tcl"
	}
	templ: {
		from: "github:vrischmann/tree-sitter-templ"
		repo: "tree-sitter-templ"
	}
	test: {
		from: "github:tree-sitter-grammars/tree-sitter-test"
		repo: "tree-sitter-test"
	}
	thrift: {
		from: "github:tree-sitter-grammars/tree-sitter-thrift"
		repo: "tree-sitter-thrift"
	}
	tmux: {
		from: "github:Freed-Wu/tree-sitter-tmux"
		repo: "tree-sitter-tmux"
	}
	toml: {
		from: "github:tree-sitter/tree-sitter-toml"
		repo: "tree-sitter-toml"
	}
	turtle: {
		from: "github:GordianDziwis/tree-sitter-turtle"
		repo: "tree-sitter-turtle"
	}
	twig: {
		from: "github:gbprod/tree-sitter-twig"
		repo: "tree-sitter-twig"
	}
	typescript: {
		from: "github:tree-sitter/tree-sitter-typescript"
		repo: "tree-sitter-typescript"
		paths: ["typescript/src", "tsx/src", "common"]
	}
	typst: {
		from: "github:uben0/tree-sitter-typst"
		repo: "tree-sitter-typst"
	}
	udev: {
		from: "github:tree-sitter-grammars/tree-sitter-udev"
		repo: "tree-sitter-udev"
	}
	ungrammar: {
		from: "github:tree-sitter-grammars/tree-sitter-ungrammar"
		repo: "tree-sitter-ungrammar"
	}
	uxntal: {
		from: "github:tree-sitter-grammars/tree-sitter-uxntal"
		repo: "tree-sitter-uxntal"
	}
	verilog: {
		from: "github:tree-sitter/tree-sitter-verilog"
		repo: "tree-sitter-verilog"
	}
	vhdl: {
		from: "github:alemuller/tree-sitter-vhdl"
		repo: "tree-sitter-vhdl"
	}
	vim: {
		from: "github:tree-sitter-grammars/tree-sitter-vim"
		repo: "tree-sitter-vim"
	}
	vimdoc: {
		from: "github:neovim/tree-sitter-vimdoc"
		repo: "tree-sitter-vimdoc"
	}
	vue: {
		from: "github:tree-sitter-grammars/tree-sitter-vue"
		repo: "tree-sitter-vue"
	}
	wasm: {
		from: "github:wasm-lsp/tree-sitter-wasm"
		repo: "tree-sitter-wasm"
		paths: ["wast/src", "wat/src"]
	}
	wgsl: {
		from: "github:szebniok/tree-sitter-wgsl"
		repo: "tree-sitter-wgsl"
	}
	"wgsl-bevy": {
		from: "github:tree-sitter-grammars/tree-sitter-wgsl-bevy"
		repo: "tree-sitter-wgsl-bevy"
	}
	wit: {
		from: "github:bytecodealliance/tree-sitter-wit"
		repo: "tree-sitter-wit"
	}
	xcompose: {
		from: "github:tree-sitter-grammars/tree-sitter-xcompose"
		repo: "tree-sitter-xcompose"
	}
	xml: {
		from: "github:tree-sitter-grammars/tree-sitter-xml"
		repo: "tree-sitter-xml"
		paths: ["xml/src", "dtd/src", "common"]
	}
	yaml: {
		from: "github:tree-sitter-grammars/tree-sitter-yaml"
		repo: "tree-sitter-yaml"
		paths: ["src", "schema/json/src", "schema/legacy/src", "schema/core/src"]
	}
	yuck: {
		from: "github:tree-sitter-grammars/tree-sitter-yuck"
		repo: "tree-sitter-yuck"
	}
	zig: {
		from: "github:tree-sitter-grammars/tree-sitter-zig"
		repo: "tree-sitter-zig"
	}
}

#grammar: [string]: {
	from:    string
	version: string | *"HEAD"
	paths:   [...string] | *["src"]
	repo:    string
}

// Test corpora (not codegen). Re-add with narrow origin when tests need them.
#fixture: {}

#fixture: [string]: {
	from:    string
	version: string | *"HEAD"
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

