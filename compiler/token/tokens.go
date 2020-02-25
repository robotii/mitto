package token

// Type is used to determine what type of token we have.
// It is aliased here for convenience.
type Type string

// Token is the structure to represent a single token.
// This is used by the lexer and parser.
type Token struct {
	Type    Type   // the type of the token
	Literal string // the literal string value of the token as lexed
	Line    int    // the line number the token appeared on
}

// Literals
const (
	Illegal          = "ILLEGAL"
	EOF              = "EOF"
	Constant         = "CONSTANT"
	Ident            = "IDENT"
	InstanceVariable = "INSTANCE_VAR"
	Int              = "INT"
	Float            = "FLOAT"
	String           = "STRING"
	Comment          = "COMMENT"

	Assign   = "="
	Plus     = "+"
	PlusEq   = "+="
	Minus    = "-"
	MinusEq  = "-="
	Bang     = "!"
	Asterisk = "*"
	Pow      = "**"
	Slash    = "/"
	Dot      = "."
	And      = "&&"
	Or       = "||"
	OrEq     = "||="
	Modulo   = "%"

	Match = "=~"
	LT    = "<"
	LTE   = "<="
	GT    = ">"
	GTE   = ">="
	COMP  = "<=>"

	Comma     = ","
	Semicolon = ";"
	Colon     = ":"
	Bar       = "|"
	Amp       = "&"

	LParen   = "("
	RParen   = ")"
	LBrace   = "{"
	RBrace   = "}"
	LBracket = "["
	RBracket = "]"

	Eq        = "=="
	NotEq     = "!="
	IsSame    = "==="
	IsNotSame = "!=="
	Range     = ".."
	RangeExcl = "..."

	True     = "TRUE"
	False    = "FALSE"
	Null     = "Null"
	If       = "IF"
	ElsIf    = "ELSIF"
	Else     = "ELSE"
	Default  = "DEFAULT"
	Switch   = "SWITCH"
	Case     = "CASE"
	Return   = "RETURN"
	Next     = "NEXT"
	Break    = "BREAK"
	Def      = "DEF"
	Self     = "SELF"
	Super    = "SUPER"
	While    = "WHILE"
	Yield    = "YIELD"
	GetBlock = "BLOCK"
	HasBlock = "HASBLOCK"
	Class    = "CLASS"
	Module   = "MODULE"
	Catch    = "CATCH"
	Finally  = "FINALLY"

	ResolutionOperator = "::"

	RightArrow = "->"
	LeftArrow  = "<-"
)

var keywords = map[string]Type{
	"def":     Def,
	"true":    True,
	"false":   False,
	"nil":     Null,
	"if":      If,
	"elsif":   ElsIf,
	"else":    Else,
	"switch":  Switch,
	"case":    Case,
	"default": Default,
	"return":  Return,
	"self":    Self,
	"super":   Super,
	"while":   While,
	"yield":   Yield,
	"next":    Next,
	"class":   Class,
	"module":  Module,
	"break":   Break,
	"block!":  GetBlock,
	"block?":  HasBlock,
	"catch":   Catch,
	"finally": Finally,
}

var operators = map[string]Type{
	"=":   Assign,
	"+":   Plus,
	"+=":  PlusEq,
	"-":   Minus,
	"-=":  MinusEq,
	"!":   Bang,
	"*":   Asterisk,
	"**":  Pow,
	"/":   Slash,
	".":   Dot,
	"&&":  And,
	"||":  Or,
	"||=": OrEq,
	"%":   Modulo,
	"&":   Amp,

	"=~":  Match,
	"<":   LT,
	"<=":  LTE,
	">":   GT,
	">=":  GTE,
	"<=>": COMP,

	"==":  Eq,
	"!=":  NotEq,
	"===": IsSame,
	"!==": IsNotSame,
	"..":  Range,
	"...": RangeExcl,

	"::": ResolutionOperator,
	"->": RightArrow,
	"<-": LeftArrow,
}

var separators = map[string]Type{
	",": Comma,
	";": Semicolon,
	":": Colon,
	"|": Bar,

	"(": LParen,
	")": RParen,
	"{": LBrace,
	"}": RBrace,
	"[": LBracket,
	"]": RBracket,
}

// Create will create a token
func Create(t Type, literal string, line int) Token {
	return Token{Type: t, Literal: literal, Line: line}
}
