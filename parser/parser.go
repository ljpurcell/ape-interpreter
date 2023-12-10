package parser

import (
	"github.com/ljpurcell/ape-interpreter/ast"
	"github.com/ljpurcell/ape-interpreter/lexer"
	"github.com/ljpurcell/ape-interpreter/token"
)

type Parser struct {
    lex lexer.Lexer

    curToken token.Token
    peekToken token.Token
}

func NewParser(l lexer.Lexer) *Parser {
    p := &Parser{lex: l}

    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.lex.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    return nil
}
