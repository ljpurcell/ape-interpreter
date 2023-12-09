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

	// Arithmetic operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	// Logical operators
	EQ  = "=="
	NOT_EQ = "=="
	NOT = "!"
	LT  = "<"
	GT  = ">"

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
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
    "true": TRUE,
    "false": FALSE,
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
