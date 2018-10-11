package ast

import (
	"bytes"
	"moka/token"
)

type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Type  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {}

func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}


func (vs *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString((vs.TokenLiteral()+" "))
	out.WriteString(vs.Name.String() + " ")
	out.WriteString(vs.Type.String())
	out.WriteString(" = ")

	if vs.Value != nil {
		out.WriteString(vs.Value.String())
	}

	out.WriteString(";")
	return out.String()
}