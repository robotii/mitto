package lexer

import (
	"mitto/compiler/token"
	"mitto/fsm"
)

// Lexer the interface between the parser and the lexer
type Lexer interface {
	NextToken() token.Token
}

// mLexer structure used to convert programs into a stream of token.
//
// Input is an array of runes, and processing is done assuming the
// text is encoded in UTF8. There is a FSM, which is used to do some
// context-sensitive lexing.
type mLexer struct {
	input        []rune   // the current input buffer being lexed
	position     int      // the index of the character we are currently lexing
	readPosition int      // the index of the next character to be read
	ch           rune     // the current character we are processing
	line         int      // the line number we are on
	fsm          *fsm.FSM // the finite state machine used to perform context-sensitive lexing
}

const (
	initial = "initial" // the initial state of the lexer
)

// New initializes a new lexer with input string
func New(input string) *mLexer {
	l := &mLexer{input: []rune(input)}
	l.advance()

	l.fsm = fsm.New(
		initial,
		fsm.States{},
	)
	return l
}

// NextToken makes lexer tokenize next character(s)
func (l *mLexer) NextToken() token.Token {

	var tok token.Token

	switch l.ch {
	default:
		tok = token.Create(token.Illegal, string(l.ch), l.line)
	}

	l.advance()
	return tok
}

// advance sets the current character and advances the read position.
// If the position exceeds the bounds of the input, the current
// character is set to null.
func (l *mLexer) advance() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
