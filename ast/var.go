package ast

import (
	"moka/token"
)

type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Type  *Identifier
	Value *Expression
}

func (ls *VarStatement) statementNode() {}

func (ls *VarStatement) TokenLiteral() string {
	return ls.Token.Literal
}
