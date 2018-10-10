package parser


import (
	"moka/ast"
	"moka/token"
	"moka/lexer"
)

type Parser struct {
	l *lexer.Lexer
	currentToken token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p:=&Parser{l: l}

	//avance the pointers to fill currentToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken () {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram () *ast.Program{
	return nil
}