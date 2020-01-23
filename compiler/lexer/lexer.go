package lexer

import "mitto/fsm"

// Lexer structure used to convert programs into a stream of tokens.
//
// Input is an array of runes, and processing is done assuming the
// text is encoded in UTF8. There is a FSM, which is used to do some
// context-sensitive lexing.
type Lexer struct {
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
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.advance()

	l.fsm = fsm.New(
		initial,
		fsm.States{},
	)
	return l
}

func (l *Lexer) advance() {}
