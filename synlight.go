//
// Copyright (c) 2018,2019 Ted Unangst <tedu@tedunangst.com>
// Copyright (c) 2024 theorytoe <theorytoe@cock.li>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package synlight

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
	"path"
	"regexp"
	"strconv"
	"strings"
)

const (
	None OutputFormat = iota
	HTML
	TTY

	TTYFgBlack   = "\x1b[30m"
	TTYFgRed     = "\x1b[31m"
	TTYFgGreen   = "\x1b[32m"
	TTYFgYellow  = "\x1b[33m"
	TTYFgBlue    = "\x1b[34m"
	TTYFgMagenta = "\x1b[35m"
	TTYFgCyan    = "\x1b[36m"
	TTYFgWhite   = "\x1b[37m"
	TTYBgBlack   = "\x1b[40m"
	TTYBgRed     = "\x1b[41m"
	TTYBgGreen   = "\x1b[42m"
	TTYBgYellow  = "\x1b[43m"
	TTYBgBlue    = "\x1b[44m"
	TTYBgMagenta = "\x1b[45m"
	TTYBgCyan    = "\x1b[46m"
	TTYBgWhite   = "\x1b[47m"
)

var (
	htmlmarkers = make(map[string]*marker)
	ttymarkers  = make(map[string]*marker)
)

type (
	token struct {
		name      string
		re        *regexp.Regexp
		state     int
		nextstate int
	}

	// Contains all the highliging primitives
	Lighter struct {
		markers map[string]*marker
		aliases map[string]string
		lexers  map[string][]*token
		write   func(io.Writer, []byte)
	}

	marker struct {
		before []byte
		after  []byte
	}

	// Options for creating a new highlighter.
	// HTML or TTY output are supported.
	Options struct {
		Format OutputFormat
	}

	OutputFormat int
)

func init() {
	htmlmarkers["keyword"] = newmarker("<span class=keyword>", "</span>")
	htmlmarkers["builtin"] = newmarker("<span class=builtin>", "</span>")
	htmlmarkers["string"] = newmarker("<span class=string>", "</span>")
	htmlmarkers["number"] = newmarker("<span class=number>", "</span>")
	htmlmarkers["type"] = newmarker("<span class=type>", "</span>")
	htmlmarkers["operator"] = newmarker("<span class=operator>", "</span>")
	htmlmarkers["comment"] = newmarker("<span class=comment>", "</span>")
	htmlmarkers["addline"] = newmarker("<span class=addline>", "</span>")
	htmlmarkers["delline"] = newmarker("<span class=delline>", "</span>")

	ttymarkers["keyword"] = newmarker("\x1b[33m", "\x1b[0m")
	ttymarkers["builtin"] = newmarker("\x1b[32m", "\x1b[0m")
	ttymarkers["string"] = newmarker("\x1b[31m", "\x1b[0m")
	ttymarkers["number"] = newmarker("\x1b[31m", "\x1b[0m")
	ttymarkers["type"] = newmarker("\x1b[32m", "\x1b[0m")
	ttymarkers["comment"] = newmarker("\x1b[34m", "\x1b[0m")
	ttymarkers["addline"] = newmarker("\x1b[32m", "\x1b[0m")
	ttymarkers["delline"] = newmarker("\x1b[31m", "\x1b[0m")
}

func newmarker(before, after string) *marker {
	return &marker{before: []byte(before), after: []byte(after)}
}

func plainwrite(w io.Writer, data []byte) {
	w.Write(data)
}

func newtoken(name string, regex string) *token {
	m := strings.Split(name, ":")
	state := 0
	nextstate := 0
	if len(m) == 3 {
		name = m[0]
		state, _ = strconv.Atoi(m[1])
		nextstate, _ = strconv.Atoi(m[2])
	}
	return &token{
		name,
		regexp.MustCompile("^" + regex),
		state,
		nextstate,
	}
}

// Add a new lexer to this highlighter.
func (hl *Lighter) AddLexer(lang string, r io.Reader) {
	var tokens []*token
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		m := strings.SplitN(line, " ", 2)
		tokens = append(tokens, newtoken(m[0], m[1]))
	}
	tokens = append(tokens, newtoken("unknown", "(?s:.)"))
	hl.lexers[lang] = tokens
}

// Create a new highlighter.
// It should be reused if possible.
func New(options Options) *Lighter {
	hl := new(Lighter)
	switch options.Format {
	case HTML:
		hl.markers = htmlmarkers
		hl.write = template.HTMLEscape
	case TTY:
		hl.markers = ttymarkers
		hl.write = plainwrite
	default:
		panic("invalid output format")
	}

	// setup hl
	hl.lexers = make(map[string][]*token)
	hl.aliases = make(map[string]string)

	// c
	hl.AddLexer("c", strings.NewReader(lexer_c))
	hl.aliases["h"] = "c"

	// c++
	hl.AddLexer("cpp", strings.NewReader(lexer_cpp))
	hl.aliases["cpp"] = "cpp"
	hl.aliases["hpp"] = "cpp"
	hl.aliases["hxx"] = "cpp"
	hl.aliases["cc"] = "cpp"
	hl.aliases["cxx"] = "cpp"
	hl.aliases["hh"] = "cpp"
	hl.aliases["c++"] = "cpp"

	// go
	hl.AddLexer("go", strings.NewReader(lexer_go))

	// python
	hl.AddLexer("py", strings.NewReader(lexer_py))
	hl.aliases["python"] = "py"
	hl.aliases["pyw"] = "py"
	hl.aliases["pyi"] = "py"

	// rust
	hl.AddLexer("rs", strings.NewReader(lexer_rs))
	hl.aliases["rust"] = "rs"

	// html+xml+svg
	hl.AddLexer("html", strings.NewReader(lexer_html))
	hl.aliases["htm"] = "html"
	hl.aliases["xhtml"] = "html"
	hl.aliases["xml"] = "html"
	hl.aliases["svg"] = "html"

	// javascript
	hl.AddLexer("js", strings.NewReader(lexer_js))
	hl.aliases["javascript"] = "js"
	hl.aliases["jsx"] = "js"
	hl.aliases["mjs"] = "js"
	hl.aliases["cjs"] = "js"

	// typescript
	hl.AddLexer("ts", strings.NewReader(lexer_ts))
	hl.aliases["typescript"] = "ts"
	hl.aliases["tsx"] = "ts"
	hl.aliases["mts"] = "ts"
	hl.aliases["cts"] = "ts"

	// lua
	hl.AddLexer("lua", strings.NewReader(lexer_lua))

	// holyc
	hl.AddLexer("hc", strings.NewReader(lexer_hc))
	hl.aliases["holyc"] = "hc"

	// diff+patch
	hl.AddLexer("diff", strings.NewReader(lexer_diff))
	hl.aliases["patch"] = "diff"

	// sql
	hl.AddLexer("sql", strings.NewReader(lexer_sql))

	// json
	hl.AddLexer("json", strings.NewReader(lexer_json))
	hl.aliases["jsonc"] = "json"

	// yaml
	hl.AddLexer("yaml", strings.NewReader(lexer_yaml))
	hl.aliases["yml"] = "yaml"

	// toml
	hl.AddLexer("toml", strings.NewReader(lexer_toml))

	// ini and friends
	hl.AddLexer("ini", strings.NewReader(lexer_ini))
	hl.aliases["conf"] = "ini"
	hl.aliases["cfg"] = "ini"
	hl.aliases["desktop"] = "ini"
	hl.aliases["service"] = "ini"
	hl.aliases["env"] = "ini"

	// markdown
	hl.AddLexer("markdown", strings.NewReader(lexer_markdown))
	hl.aliases["md"] = "markdown"
	hl.aliases["mkd"] = "markdown"

	// shell
	hl.AddLexer("shell", strings.NewReader(lexer_shell))
	hl.aliases["sh"] = "shell"
	hl.aliases["bash"] = "shell"
	hl.aliases["zsh"] = "shell"
	hl.aliases["ksh"] = "shell"

	// css
	hl.AddLexer("css", strings.NewReader(lexer_css))

	// make
	hl.AddLexer("make", strings.NewReader(lexer_make))
	hl.aliases["mk"] = "make"
	hl.aliases["mak"] = "make"
	hl.aliases["makefile"] = "make"

	// dockerfile
	hl.AddLexer("dockerfile", strings.NewReader(lexer_dockerfile))
	hl.aliases["docker"] = "dockerfile"

	// java
	hl.AddLexer("java", strings.NewReader(lexer_java))

	// c#
	hl.AddLexer("cs", strings.NewReader(lexer_cs))
	hl.aliases["csharp"] = "cs"

	// ruby
	hl.AddLexer("rb", strings.NewReader(lexer_rb))
	hl.aliases["ruby"] = "rb"
	hl.aliases["gemspec"] = "rb"
	hl.aliases["rake"] = "rb"

	// php
	hl.AddLexer("php", strings.NewReader(lexer_php))
	hl.aliases["phtml"] = "php"

	// kotlin
	hl.AddLexer("kt", strings.NewReader(lexer_kt))
	hl.aliases["kts"] = "kt"
	hl.aliases["kotlin"] = "kt"

	// swift
	hl.AddLexer("swift", strings.NewReader(lexer_swift))

	// zig
	hl.AddLexer("zig", strings.NewReader(lexer_zig))

	// perl
	hl.AddLexer("pl", strings.NewReader(lexer_pl))
	hl.aliases["pm"] = "pl"
	hl.aliases["perl"] = "pl"

	// haskell
	hl.AddLexer("hs", strings.NewReader(lexer_hs))
	hl.aliases["haskell"] = "hs"
	hl.aliases["lhs"] = "hs"

	// elixir
	hl.AddLexer("ex", strings.NewReader(lexer_ex))
	hl.aliases["exs"] = "ex"
	hl.aliases["elixir"] = "ex"

	// myrddin
	hl.AddLexer("myr", strings.NewReader(lexer_myr))
	hl.aliases["myrddin"] = "myr"

	return hl
}

var pairnames = []string{"string", "keyword", "comment", "builtin"}

// normalizeExt lowercases an extension and drops a leading dot.
func normalizeExt(ext string) string {
	return strings.ToLower(strings.TrimPrefix(ext, "."))
}

// extOf picks the lexer extension for filename. Files without an extension,
// such as Makefile or Dockerfile, fall back to the bare base name.
func extOf(filename string) string {
	ext := path.Ext(filename)
	if ext == "" {
		return strings.ToLower(path.Base(filename))
	}
	return normalizeExt(ext)
}

// Highlight code, writing it to w.
func (hl *Lighter) Highlight(data []byte, filename string, w io.Writer) {
	ext := extOf(filename)

	alt, ok := hl.aliases[ext]
	if ok {
		ext = alt
	}

	markers := hl.markers

	tokens := hl.lexers[ext]
	if tokens == nil {
		hl.write(w, data)
		return
	}

	dataloc := 0
	state := 0
	pairidx := uint(0)
	for dataloc < len(data) {
	restart:
		for _, tok := range tokens {
			if tok.state != state {
				continue
			}
			m := tok.re.Find(data[dataloc:])
			if m != nil {
				state = tok.nextstate
				name := tok.name
				if name == "pair" {
					name = pairnames[pairidx%uint(len(pairnames))]
					pairidx += 1
				} else if name == "unpair" {
					pairidx -= 1
					name = pairnames[pairidx%uint(len(pairnames))]
				}
				mk := markers[name]
				if mk != nil {
					w.Write(mk.before)
				}
				hl.write(w, m)
				if mk != nil {
					w.Write(mk.after)
				}
				dataloc += len(m)
				goto restart
			}
		}
		state = 0
	}
}

// Highlight code, returning a string
func (hl *Lighter) HighlightString(data string, filename string) string {
	var buf bytes.Buffer
	hl.Highlight([]byte(data), filename, &buf)
	return buf.String()
}

// Highlight code based on provided extension
func (hl *Lighter) HighlightExt(data []byte, extension string, w io.Writer) {
	ext := normalizeExt(extension)

	alt, ok := hl.aliases[ext]
	if ok {
		ext = alt
	}

	markers := hl.markers

	tokens := hl.lexers[ext]
	if tokens == nil {
		hl.write(w, data)
		return
	}

	dataloc := 0
	state := 0
	pairidx := uint(0)
	for dataloc < len(data) {
	restart:
		for _, tok := range tokens {
			if tok.state != state {
				continue
			}
			m := tok.re.Find(data[dataloc:])
			if m != nil {
				state = tok.nextstate
				name := tok.name
				if name == "pair" {
					name = pairnames[pairidx%uint(len(pairnames))]
					pairidx += 1
				} else if name == "unpair" {
					pairidx -= 1
					name = pairnames[pairidx%uint(len(pairnames))]
				}
				mk := markers[name]
				if mk != nil {
					w.Write(mk.before)
				}
				hl.write(w, m)
				if mk != nil {
					w.Write(mk.after)
				}
				dataloc += len(m)
				goto restart
			}
		}
		state = 0
	}
}

// Hihglight code based on provided extension, returning a string
func (hl *Lighter) HighlightExtString(data string, extenstion string) string {
	var buf bytes.Buffer
	hl.HighlightExt([]byte(data), extenstion, &buf)
	return buf.String()
}
