# synlight

A fork of synlight from [webs](https://humungus.tedunangst.com/r/webs) with
more language support and some planned
improvements.

supported langs:

- c (.c, .h)
- cpp (.cpp, .hpp, .hxx, .cc, .cxx, .hh, .c++)
- holyc (.hc, .holyc)
- go (.go)
- python (.py, .pyw, .pyi)
- rust (.rs)
- javascript (.js, .jsx, .mjs, .cjs)
- typescript (.ts, .tsx, .mts, .cts)
- java (.java)
- c# (.cs)
- kotlin (.kt, .kts)
- swift (.swift)
- zig (.zig)
- ruby (.rb, .gemspec, .rake)
- php (.php, .phtml)
- perl (.pl, .pm)
- lua (.lua)
- haskell (.hs, .lhs)
- elixir (.ex, .exs)
- myrddin (.myr)
- sql (.sql)
- json (.json, .jsonc)
- yaml (.yaml, .yml)
- toml (.toml)
- ini (.ini, .conf, .cfg, .desktop, .service, .env)
- markdown (.md, .mkd)
- shell (.sh, .bash, .zsh, .ksh)
- css (.css)
- make (.mk, .mak, Makefile)
- dockerfile (Dockerfile)
- diff (.diff, .patch)
- html (.html, .htm, .xhtml, .xml, .svg)

Lexers live in `lexers/*.txt`, one token per line as `name regex`. The
`lexers.go` file is generated from them:

	make

Files with no extension, such as `Makefile` and `Dockerfile`, are matched on
their base name. Extensions are matched case insensitively.
