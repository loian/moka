package ast


type Node interface {
	TokenLiteral() string
	Token
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}


type Program struct {
	Statements []Statement
}