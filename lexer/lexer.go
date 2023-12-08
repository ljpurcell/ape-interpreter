// The lexer is responsible for taking source code and turning it into tokens, which are easier to work than plain text.
package lexer

import (
	"github.com/ljpurcell/monkey-interpreter/token"
	"github.com/ljpurcell/monkey-interpreter/utils"
)

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

	l.skipWhitespace()

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
		tok.Type = token.EOF
	default:
		if utils.IsLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.GetTypeFrom(tok.Literal)
			return tok
		} else if utils.IsDigit(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.GetNumericType(tok.Literal)
            return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	start := l.position
	for utils.IsDigit(l.char) || l.char == '.' {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for utils.IsLetter(l.char) || l.char == '_' {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\n' || l.char == '\t' || l.char == '\r' {
		l.readChar()
	}
}
