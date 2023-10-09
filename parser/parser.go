package parser

import (
	"github.com/tuantran1810/go-interpreter/token"

	"github.com/tuantran1810/go-interpreter/ast"
	"github.com/tuantran1810/go-interpreter/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
