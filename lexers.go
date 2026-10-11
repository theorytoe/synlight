package synlight

var lexer_cpp = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+)\n" +
	"number '(\\\\.|[^'])'\n" +
	"keyword (if|else|for|while|do|continue|break|switch|case|default|return|typedef|extern|goto)\\b\n" +
	"keyword (new|delete|this|friend|using|public|protected|private|inline|virtual|explicit|export|operator|typeid|mutable)\\b\n" +
	"keyword (throw|try|catch|noexcept|static_assert|concept|consteval|constinit|requires|co_await|co_return|co_yield)\\b\n" +
	"keyword (class|typename|template|namespace|struct|union|enum)\\b\n" +
	"keyword (constexpr|decltype|thread_local|alignas|alignof|static_cast|dynamic_cast|reinterpret_cast|const_cast)\\b\n" +
	"builtin (NULL|va_list|true|false)\\b\n" +
	"builtin #(\\\\\\n|[^\\n])*\n" +
	"type (static|const|volatile|register|unsigned|signed|char|int|long|short|float|double|byte|bool|void|auto|nullptr|nullptr_t|wchar_t|char8_t|char16_t|char32_t|size_t|ssize_t|int8_t|uint8_t|int16_t|uint16_t|int32_t|uint32_t|int64_t|uint64_t)\\b\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"operator (and|bitor|or|xor|compl|bitand|and_eq|or_eq|xor_eq|not|not_eq)\\b\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_css = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"number (#[0-9a-fA-F]{3,8}|[0-9]+(\\.[0-9]+)?[a-zA-Z%]*)\n" +
	"keyword @[a-zA-Z-]+\n" +
	"keyword [a-zA-Z-]+[ \\t]*:\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"builtin [.#][a-zA-Z0-9_-]+\n" +
	"builtin :[a-zA-Z-]+\n" +
	"operator [{};:>+~()\\[\\]=]\n" +
	"word [a-zA-Z_][a-zA-Z0-9_-]*\n" +
	""
var lexer_cs = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"builtin #(\\\\\\n|[^\\n])*\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?[fFdDmMuUlL]?)\n" +
	"number '(\\\\.|[^'])*'\n" +
	"keyword (abstract|as|async|await|base|break|case|catch|checked|class|const|continue|default|delegate|do|else|enum|event|explicit|extern|finally|fixed|for|foreach|goto|if|implicit|in|interface|internal|is|lock|namespace|new|operator|out|override|params|partial|private|protected|public|readonly|record|ref|return|sealed|sizeof|stackalloc|static|struct|switch|this|throw|try|typeof|unchecked|unsafe|using|virtual|volatile|while|yield|var|dynamic|get|set|init|where|when)\\b\n" +
	"builtin (true|false|null|value)\\b\n" +
	"type (bool|byte|sbyte|char|decimal|double|float|int|uint|long|ulong|short|ushort|object|string|void|String|Object|List|Dictionary|IEnumerable|Task|Exception)\\b\n" +
	"string @\"(?:[^\"]|\"\")*\"\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\.|[^'])*'\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_c = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+)\n" +
	"number '(\\\\.|[^'])'\n" +
	"keyword (if|else|for|while|do|continue|break|switch|case|default|return|typedef|extern|goto|sizeof|struct|union|enum)\\b\n" +
	"builtin (NULL|va_list|true|false)\\b\n" +
	"builtin #(\\\\\\n|[^\\n])*\n" +
	"type (static|const|volatile|register|inline|unsigned|signed|char|int|long|short|float|double|byte|bool|void|size_t|ssize_t|int8_t|uint8_t|int16_t|uint16_t|int32_t|uint32_t|int64_t|uint64_t|FILE)\\b\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_diff = "" +
	"keyword diff[^\\n]*\n" +
	"builtin @@[^\\n]*\n" +
	"delline -[^\\n]*\n" +
	"addline \\+[^\\n]*\n" +
	"text [^\\n]+\n" +
	"newline \\n\n" +
	""
var lexer_dockerfile = "" +
	"whitespace [ \\t]+\n" +
	"newline \\n\n" +
	"comment #[^\\n]*\n" +
	"keyword ^(?i:from|run|cmd|label|maintainer|expose|env|add|copy|entrypoint|volume|user|workdir|arg|onbuild|stopsignal|healthcheck|shell)\\b\n" +
	"builtin \\$[a-zA-Z_][a-zA-Z0-9_]*\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"word [a-zA-Z_][a-zA-Z0-9_.-]*\n" +
	"operator [=:,]\n" +
	""
var lexer_ex = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment #[^\\n]*\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?)\n" +
	"keyword (after|alias|and|case|catch|cond|def|defp|defmodule|defmacro|defstruct|defprotocol|defimpl|defdelegate|defguard|do|else|end|fn|for|if|import|in|not|or|quote|raise|receive|require|rescue|super|throw|try|unless|unquote|use|when|with)\\b\n" +
	"builtin (true|false|nil|self|__MODULE__|__ENV__|__DIR__|__CALLER__)\\b\n" +
	"builtin (Enum|String|List|Map|MapSet|Keyword|Tuple|IO|File|Path|Agent|Task|GenServer|Supervisor|Process|Kernel|Float|Integer|Stream|Range|Regex|Date|Time|DateTime|NaiveDateTime|URI|System|Node|Application|Base)\\b\n" +
	"string \"\"\"(?s:.*?)\"\"\"\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"string ~[rRwWsScC][a-zA-Z]*/([^/\\n]|\\\\.)*/\n" +
	"string ~[rRwWsScC][a-zA-Z]*\\(([^)\\n]|\\\\.)*\\)\n" +
	"builtin :[a-zA-Z_][a-zA-Z0-9_]*[?!]?\n" +
	"builtin @[a-zA-Z_][a-zA-Z0-9_]*\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*[?!]?\n" +
	"operator [-+/*=:<>,\\.!&|^~%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_go = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?([eE][-+]?[0-9]+)?)\n" +
	"number '(\\\\.|[^'])'\n" +
	"keyword (break|case|chan|const|continue|default|defer|else|fallthrough|for|func|go|goto|if|import|interface|map|package|range|return|select|struct|switch|type|var)\\b\n" +
	"builtin (append|cap|clear|close|complex|copy|delete|imag|len|make|max|min|new|panic|print|println|real|recover)\\b\n" +
	"type (any|bool|byte|comparable|complex128|complex64|error|float32|float64|int|int16|int32|int64|int8|rune|string|uint|uint16|uint32|uint64|uint8|uintptr)\\b\n" +
	"string `[^`]*`\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"nop (\\[]|{})\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_hc = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+)\n" +
	"number '(\\\\.|[^'])'\n" +
	"keyword (if|else|for|while|do|continue|break|switch|case|default|return|typedef|extern|goto|sizeof|class|union|struct|enum|public|private|protected)\\b\n" +
	"builtin (NULL|va_list|true|false|argc|argv)\\b\n" +
	"builtin #(\\\\\\n|[^\\n])*\n" +
	"type (static|const|unsigned|signed|char|int|long|short|float|double|byte|bool|void|U0|U8|U16|U32|U64|I8|I16|I32|I64|F64|CDate|CTask)\\b\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_hs = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?)\n" +
	"keyword (case|class|data|default|deriving|do|else|foreign|if|import|in|infix|infixl|infixr|instance|let|module|newtype|of|then|type|where|qualified|as|hiding|forall|mdo|proc|rec)\\b\n" +
	"builtin (map|filter|foldl|foldr|foldl1|foldr1|head|tail|init|last|length|reverse|concat|concatMap|zip|zipWith|unzip|take|drop|takeWhile|dropWhile|span|break|elem|notElem|lookup|sum|product|maximum|minimum|replicate|show|read|print|putStrLn|putStr|getLine|return|pure|fmap|mapM|mapM_|sequence|id|const|flip|curry|uncurry|error|undefined|otherwise|True|False|Just|Nothing)\\b\n" +
	"type (Int|Integer|Float|Double|Char|String|Bool|Maybe|Either|Ordering|IO|Show|Read|Eq|Ord|Num|Enum|Bounded|Functor|Applicative|Monad|Foldable|Traversable|Semigroup|Monoid)\\b\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_][a-zA-Z0-9_']*\n" +
	"comment --[^\\n]*\n" +
	"comment \\{-(?s:.*?)-\\}\n" +
	"operator [-+/*=:<>,\\.!&|^~?%$]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_html = "" +
	"comment:0:0 <!--(?s:.*?)-->\n" +
	"builtin:0:1 <[/a-zA-Z]*\n" +
	"keyword:1:1 [^=/>\"]+\n" +
	"string:1:1 \"[^\"]*\"\n" +
	"text:1:1 [\\s=]+\n" +
	"builtin:1:0 /?>\n" +
	"comment:0:0 {{(?s:.*?)}}\n" +
	"text:0:0 [^<{]+\n" +
	""
var lexer_ini = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment [;#][^\\n]*\n" +
	"keyword ^\\[[^\\]]*\\]\n" +
	"builtin [^=\\n]+=\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"number -?(0[xX][0-9a-fA-F]+|[0-9]+)\n" +
	"word [a-zA-Z_][a-zA-Z0-9_.-]*\n" +
	"operator [=:,]\n" +
	""
var lexer_java = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?[fFdDlL]?)\n" +
	"number '(\\\\.|[^'])*'\n" +
	"keyword (abstract|assert|break|case|catch|class|const|continue|default|do|else|enum|extends|final|finally|for|goto|if|implements|import|instanceof|interface|native|new|package|permits|private|protected|public|record|return|sealed|static|strictfp|super|switch|synchronized|this|throw|throws|transient|try|var|volatile|while|yield)\\b\n" +
	"builtin (true|false|null)\\b\n" +
	"type (boolean|byte|char|double|float|int|long|short|void|String|Object|Integer|Long|Double|Float|Boolean|Character|Byte|Short|List|Map|Set|ArrayList|HashMap|HashSet|Optional|Stream|Exception|RuntimeException|Override)\\b\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\.|[^'])*'\n" +
	"word [a-zA-Z_$][a-zA-Z0-9_$]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_json = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"number -?(0|[1-9][0-9]*)(\\.[0-9]+)?([eE][-+]?[0-9]+)?\n" +
	"keyword (true|false|null)\\b\n" +
	"operator [\\[\\]{},:]+\n" +
	""
var lexer_js = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?([eE][-+]?[0-9]+)?)\n" +
	"keyword (as|async|await|break|case|catch|class|const|continue|debugger|default|delete|do|else|export|extends|finally|for|from|function|get|if|import|in|instanceof|let|new|of|return|set|static|super|switch|this|throw|try|typeof|var|void|while|with|yield)\\b\n" +
	"builtin (true|false|null|undefined|NaN|Infinity|arguments|eval|escape)\\b\n" +
	"builtin ((console|Math|JSON|Object|Array|String|Number|Boolean|Date|RegExp|Promise|Map|Set|WeakMap|WeakSet|Symbol|Error|TypeError|RangeError|document|window|globalThis|history|location|navigator)\\.[[:alpha:]]+)\n" +
	"type (console|Math|JSON|Object|Array|String|Number|Boolean|Date|RegExp|Promise|Map|Set|WeakMap|WeakSet|Symbol|Error|TypeError|RangeError|document|window|globalThis|history|location|navigator)\\b\n" +
	"string `[^`]*`\n" +
	"string '(\\\\'|[^'])*'\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_$][a-zA-Z0-9_$]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_kt = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?[fFdDlL]?)\n" +
	"keyword (as|break|class|continue|do|else|false|for|fun|if|in|interface|is|null|object|package|return|super|this|throw|true|try|typealias|typeof|val|var|when|while|by|catch|constructor|delegate|dynamic|field|file|finally|get|import|init|param|property|receiver|set|setparam|where|actual|abstract|annotation|companion|const|crossinline|data|enum|expect|external|final|infix|inline|inner|internal|lateinit|noinline|open|operator|out|override|private|protected|public|reified|sealed|suspend|tailrec|vararg)\\b\n" +
	"type (Boolean|Byte|Char|Double|Float|Int|Long|Short|String|Unit|Any|Nothing|List|Map|Set|Array|MutableList|MutableMap|MutableSet|Pair|Triple)\\b\n" +
	"string \"\"\"(?s:.*?)\"\"\"\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"string `[^`]*`\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_lua = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?)\n" +
	"keyword (and|break|do|else|elseif|end|for|function|goto|if|in|local|not|or|repeat|return|then|until|while)\\b\n" +
	"builtin (require|print|pairs|ipairs|next|pcall|xpcall|error|assert|type|tostring|tonumber|setmetatable|getmetatable|rawget|rawset|select|unpack|true|false|nil|self)\\b\n" +
	"builtin ((string|table|math|io|os|coroutine|debug)\\.[a-zA-Z_]+)\n" +
	"string \\[=*\\[(?s:.*?)\\]=*\\]\n" +
	"string '(\\\\'|[^'])*'\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment --\\[=*\\[(?s:.*?)\\]=*\\]\n" +
	"comment --[^\\n]*\n" +
	"operator [-+/*=:<>,\\.!&|^~#%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_make = "" +
	"whitespace [ \\t]+\n" +
	"newline \\n\n" +
	"comment #[^\\n]*\n" +
	"builtin ^\\.[A-Z]+:\n" +
	"keyword ^[a-zA-Z0-9_./-]+[ \\t]*:\n" +
	"builtin ^(include|ifdef|ifndef|ifeq|ifneq|else|endif|define|endef|export|unexport|override|vpath|undefine)\\b\n" +
	"builtin \\$[({]?[a-zA-Z_][a-zA-Z0-9_]*[)}]?\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"operator [:=?+!@|]\n" +
	""
var lexer_markdown = "" +
	"whitespace [ \\t]+\n" +
	"newline \\n\n" +
	"comment <!--(?s:.*?)-->\n" +
	"keyword ^#{1,6}[^\\n]*\n" +
	"keyword ^>[^\\n]*\n" +
	"keyword ^\\*\\*\\*[^\\n]*\n" +
	"keyword ^---[^\\n]*\n" +
	"keyword ^___[^\\n]*\n" +
	"number ^[ \\t]*\\[[ xX]\\]\n" +
	"string `[^`]*`\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"builtin !?\\[[^\\]]*\\]\\([^)]*\\)\n" +
	"builtin \\*\\*[^\\n]*?\\*\\*\n" +
	"builtin __[^_\\n]*?__\n" +
	"builtin ~~[^\\n]*?~~\n" +
	"builtin \\*[^\\n]*?\\*\n" +
	"builtin _[^_\\n]*?_\n" +
	"builtin ^[ \\t]*([-*+]|[0-9]+\\.)[ \\t]\n" +
	""
var lexer_myr = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"number (0[xX][0-9a-fA-F_]+|0[bB][01_]+|0[oO][0-7_]+|[0-9][0-9_]*(\\.[0-9][0-9_]*)?([eE][-+]?[0-9]+)?)([uUlLiI][0-9]*)?\n" +
	"number '(\\\\.|[^'])'\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"keyword \\$[a-zA-Z_][a-zA-Z0-9_]*\n" +
	"keyword (auto|break|const|continue|defer|elif|else|extern|for|generic|goto|if|impl|match|pkg|pkglocal|sizeof|struct|trait|type|union|use|var|while)\\b\n" +
	"builtin (true|false|void)\\b\n" +
	"type (bool|char|byte|int8|int16|int32|int64|int|uint8|uint16|uint32|uint64|uint|flt32|flt64)\\b\n" +
	"builtin @[a-zA-Z_][a-zA-Z0-9_]*\n" +
	"builtin \\.\\.\\.\n" +
	"keyword ;;\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"operator (<<=|>>=|->|::|&&|\\|\\||==|!=|<=|>=|<<|>>|\\+\\+|--|\\+=|-=|\\*=|/=|%=|\\|=|&=|\\^=)\n" +
	"operator [-+*/%#&|^~!=<>:.,?@`;]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_php = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment //[^\\n]*\n" +
	"comment #[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"keyword (<\\?php|<\\?=|\\?>)\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?)\n" +
	"keyword (abstract|and|as|break|callable|case|catch|class|clone|const|continue|declare|default|do|echo|else|elseif|empty|enddeclare|endfor|endforeach|endif|endswitch|endwhile|enum|eval|exit|extends|final|finally|fn|for|foreach|function|global|goto|if|implements|include|include_once|instanceof|insteadof|interface|isset|list|match|namespace|new|or|print|private|protected|public|readonly|require|require_once|return|static|switch|throw|trait|try|unset|use|var|while|xor|yield|from)\\b\n" +
	"builtin (true|false|null|TRUE|FALSE|NULL|self|parent|__construct|__destruct|__toString|__get|__set|__call)\n" +
	"type (int|float|string|bool|void|iterable|object|mixed|never)\\b\n" +
	"builtin \\$[a-zA-Z_][a-zA-Z0-9_]*\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"operator [-+/*=:<>,\\.!&|^~?%@]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_pl = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment #[^\\n]*\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?)\n" +
	"keyword (and|or|not|xor|cmp|eq|ne|lt|gt|le|ge|if|elsif|else|unless|while|until|for|foreach|continue|do|given|when|default|last|next|redo|return|sub|my|our|local|state|package|use|require|no|BEGIN|END|CHECK|INIT|UNITCHECK|DESTROY|eval|die|warn|print|printf|say|open|close|read|write|chomp|chop|split|join|push|pop|shift|unshift|splice|keys|values|each|exists|delete|defined|ref|bless|scalar|sort|reverse|map|grep|wantarray|caller|tie|tied|untie|lock|sleep|exit)\\b\n" +
	"builtin (\\$[a-zA-Z_][a-zA-Z0-9_]*|@[a-zA-Z_][a-zA-Z0-9_]*|%[a-zA-Z_][a-zA-Z0-9_]*|&[a-zA-Z_][a-zA-Z0-9_]*)\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_py = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?([eE][-+]?[0-9]+)?)\n" +
	"keyword (and|as|assert|async|await|break|class|continue|def|del|elif|else|except|finally|for|from|global|if|import|in|is|lambda|nonlocal|not|or|pass|raise|return|try|while|with|yield|match|case)\\b\n" +
	"builtin (abs|all|any|bin|bool|bytearray|bytes|callable|chr|classmethod|compile|complex|delattr|dict|dir|divmod|enumerate|eval|exec|filter|float|format|frozenset|getattr|globals|hasattr|hash|help|hex|id|input|int|isinstance|issubclass|iter|len|list|locals|map|max|memoryview|min|next|object|oct|open|ord|pow|print|property|range|repr|reversed|round|set|setattr|slice|sorted|staticmethod|str|sum|super|tuple|type|vars|zip|Exception|ValueError|TypeError|KeyError|IndexError|RuntimeError|StopIteration|AttributeError|ImportError|None|True|False|self|cls)\\b\n" +
	"string [rRbBuUfF]{0,2}\"\"\"(?s:.*?)\"\"\"\n" +
	"string [rRbBuUfF]{0,2}'''(?s:.*?)'''\n" +
	"string [rRbBuUfF]{0,2}\"(\\\\\"|[^\"])*\"\n" +
	"string [rRbBuUfF]{0,2}'(\\\\'|[^'])*'\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment #[^\\n]*\n" +
	"operator [-+/*=:<>,\\.!&|^~%@]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_rb = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment #[^\\n]*\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?)\n" +
	"keyword (alias|and|begin|break|case|class|def|defined|do|else|elsif|end|ensure|false|for|if|in|module|next|nil|not|or|redo|rescue|retry|return|self|super|then|true|undef|unless|until|when|while|yield|require|require_relative|include|extend|attr_accessor|attr_reader|attr_writer|private|public|protected|raise|lambda|proc|loop)\\b\n" +
	"builtin (puts|print|gets|chomp|to_s|to_i|to_f|to_a|to_h|each|map|select|reject|inject|reduce|freeze|new)\\b\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"string :[a-zA-Z_][a-zA-Z0-9_]*[?!]?\n" +
	"string /[^/\\n]*/\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*[?!]?\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_rs = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?([eE][-+]?[0-9]+)?)\n" +
	"number '(\\\\.|[^'])'\n" +
	"builtin ^#!?\\[[^\\]]*\\]\n" +
	"keyword (as|async|await|break|const|continue|crate|dyn|else|enum|extern|fn|for|if|impl|in|let|loop|match|mod|move|mut|pub|ref|return|self|Self|static|struct|super|trait|type|unsafe|use|where|while|union|macro_rules)\\b\n" +
	"builtin (Err|Ok|Result|Some|None|Box|Rc|Arc|Vec|String|Option|HashMap|HashSet|Iterator|Clone|Copy|Debug|Default|Display|Eq|PartialEq|Ord|PartialOrd|Hash|Send|Sync|Drop|Fn|FnMut|FnOnce)\\b\n" +
	"type (bool|char|f32|f64|i8|i16|i32|i64|i128|isize|u8|u16|u32|u64|u128|usize|str)\\b\n" +
	"builtin [a-zA-Z_][a-zA-Z0-9_]*!\n" +
	"string b?\"(\\\\\"|[^\"])*\n" +
	"string b?'(\\\\.|[^'])*'\n" +
	"string `[^`]*`\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_shell = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment #[^\\n]*\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '(\\\\'|[^'])*'\n" +
	"string `[^`]*`\n" +
	"keyword (if|then|else|elif|fi|for|while|until|do|done|case|esac|in|function|select|time|return|local|export|readonly|declare|typeset|shift|break|continue|exit|trap|set|unset|exec|eval|source|alias|unalias|let|test|read|cd|pwd|echo|printf|true|false)\\b\n" +
	"builtin \\$[({]?[a-zA-Z_][a-zA-Z0-9_]*[)}]?\n" +
	"builtin [0-9]?>&?[0-9]*\n" +
	"word [a-zA-Z_][a-zA-Z0-9_-]*\n" +
	"operator [-+/*=:<>,\\\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_sql = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment --[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?)\n" +
	"keyword ((?i)add|all|alter|and|any|as|asc|begin|between|by|case|check|column|commit|constraint|create|database|default|delete|desc|distinct|drop|else|end|exists|foreign|from|full|group|having|in|index|inner|insert|into|is|join|left|like|limit|not|null|offset|on|or|order|outer|primary|references|right|rollback|select|set|table|then|top|truncate|union|unique|update|values|view|when|where)\\b\n" +
	"builtin ((?i)bigint|blob|boolean|char|date|datetime|decimal|double|float|int|integer|json|numeric|real|serial|smallint|text|time|timestamp|uuid|varchar)\\b\n" +
	"string '(''|[^'])*'\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"operator [-+/*=:<>,\\.!|%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_swift = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?([eE][-+]?[0-9]+)?)\n" +
	"keyword (associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|precedencegroup|protocol|public|rethrows|static|struct|subscript|typealias|var|break|case|catch|continue|default|defer|do|else|fallthrough|for|guard|if|in|repeat|return|throw|switch|where|while|as|any|await|false|is|nil|self|Self|super|throws|true|try|async|actor|some|convenience|dynamic|final|indirect|infix|lazy|mutating|nonmutating|optional|override|postfix|prefix|required|unowned|weak|willSet|didSet)\\b\n" +
	"type (Any|AnyObject|Array|Bool|Character|Dictionary|Double|Float|Int|Int8|Int16|Int32|Int64|Optional|Never|Set|String|UInt|UInt8|UInt16|UInt32|UInt64|Void)\\b\n" +
	"builtin @[a-zA-Z_][a-zA-Z0-9_]*\n" +
	"builtin ^#[a-zA-Z]+\n" +
	"string \"\"\"(?s:.*?)\"\"\"\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_toml = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment #[^\\n]*\n" +
	"keyword ^\\[\\[?[^\\]]*\\]\\]?\n" +
	"keyword (true|false)\\b\n" +
	"number -?(0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?([eE][-+]?[0-9]+)?)\n" +
	"string \"\"\"(?s:.*?)\"\"\"\n" +
	"string '''(?s:.*?)'''\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '((''|[^'])*)'\n" +
	"keyword [a-zA-Z_][a-zA-Z0-9_.-]*[ \\t]*=\n" +
	"operator [\\[\\]{},.=]\n" +
	""
var lexer_ts = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?([eE][-+]?[0-9]+)?)\n" +
	"keyword (abstract|any|as|asserts|async|await|break|case|catch|class|const|constructor|continue|declare|default|delete|do|else|enum|export|extends|finally|for|from|function|get|if|implements|import|in|infer|instanceof|interface|is|keyof|let|module|namespace|new|of|private|protected|public|readonly|return|satisfies|set|static|super|switch|this|throw|try|type|typeof|unique|var|while|with|yield)\\b\n" +
	"builtin (true|false|null|undefined|NaN|Infinity)\\b\n" +
	"type (string|number|boolean|object|symbol|bigint|unknown|never|any|void|Array|Promise|Record|Partial|Required|Readonly|Map|Set|Date|RegExp|Error)\\b\n" +
	"builtin ((console|Math|JSON|Object|Array|String|Number|document|window|globalThis)\\.[[:alpha:]]+)\n" +
	"string `[^`]*`\n" +
	"string '(\\\\'|[^'])*'\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_$][a-zA-Z0-9_$]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
var lexer_yaml = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"comment #[^\\n]*\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"string '((''|[^'])*)'\n" +
	"number -?(0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?)\n" +
	"keyword (true|True|false|False|yes|Yes|no|No|on|On|off|Off|null|Null|NULL)\\b\n" +
	"keyword ^[ \\t]*-[ \\t]\n" +
	"keyword [^ \\t\\n#:]+:\n" +
	"operator [|>&*!%@`\\[\\]{},]\n" +
	""
var lexer_zig = "" +
	"whitespace [ \\t\\r\\n]+\n" +
	"number (0[xX][0-9a-fA-F]+|[0-9]+(\\.[0-9]+)?([eE][-+]?[0-9]+)?)\n" +
	"number '(\\\\.|[^'])*'\n" +
	"keyword (addrspace|align|allowzero|and|anyframe|anytype|asm|async|await|break|callconv|catch|comptime|const|continue|defer|else|enum|errdefer|error|export|extern|fn|for|if|inline|linksection|noalias|noinline|nosuspend|opaque|or|orelse|packed|pub|resume|return|struct|suspend|switch|test|threadlocal|try|union|unreachable|usingnamespace|var|volatile|while)\\b\n" +
	"type (anyerror|anyframe|anyopaque|bool|comptime_float|comptime_int|f16|f32|f64|f80|f128|isize|noreturn|type|usize|void|i8|i16|i32|i64|i128|u8|u16|u32|u64|u128)\\b\n" +
	"builtin @[a-zA-Z_][a-zA-Z0-9_]*\n" +
	"string \"(\\\\\"|[^\"])*\"\n" +
	"word [a-zA-Z_][a-zA-Z0-9_]*\n" +
	"comment //[^\\n]*\n" +
	"comment /\\*(?s:.*?)\\*/\n" +
	"operator [-+/*=:<>,\\.!&|^~?%]+\n" +
	"pair [([{]\n" +
	"unpair [)\\]}]\n" +
	""
