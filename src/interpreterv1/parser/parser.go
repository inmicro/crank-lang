package parser

import (
	"crank/ast"
	"crank/lexer"
	"crank/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // curToken => current Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	// program := newProgramASTNode()

	return nil
}
