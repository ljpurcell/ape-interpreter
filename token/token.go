package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	FLOAT = "FLOAT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func NewToken(tt TokenType, char byte) Token {
	return Token{
		Type:    tt,
		Literal: string(char),
	}
}

func GetTypeFrom(ident string) TokenType {
	if tt, exists := keywords[ident]; exists {
		return tt
	}
	return IDENT
}

func GetNumericType(number string) TokenType {
	if strings.Contains(number, ".") {
		return FLOAT
	}

	return INT
}
