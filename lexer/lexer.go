// The lexer is responsible for taking source code and turning it into tokens, which are easier to work than plain text.
package lexer

import "github.com/ljpurcell/monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int  // position after 'position'
	char         byte // the character pointed to by 'position'
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.char {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.char)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.char)
	case '(':
		tok = token.NewToken(token.LPAREN, l.char)
	case ')':
		tok = token.NewToken(token.RPAREN, l.char)
	case ',':
		tok = token.NewToken(token.COMMA, l.char)
	case '+':
		tok = token.NewToken(token.PLUS, l.char)
	case '{':
		tok = token.NewToken(token.LBRACE, l.char)
	case '}':
		tok = token.NewToken(token.RBRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type  = token.EOF

	}
	l.readChar()
	return tok
}

