package tokens

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
