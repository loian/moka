package ast




type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
	acceptedTypes()
}

type Expression interface {
	Node
	expressionNode()
	expressionTypeAnnotation()
}

