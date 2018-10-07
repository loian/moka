package token

const (
	ILLEGAL = iota
	EOF

	// Literals and identifiers
	IDENTIFIER
	BOOLEAN
	INT
	FLOAT
	DECIMAL
	STRING

	//Operator
	ASSIGN
	PLUS

	//Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACKET
	RBRACKET
	LBRACE
	RBRACE

	//Keywords
	FUNCTION
	LET
	STRUCT
)

type TokenType uint8

type Token struct {
	Type    TokenType
	Literal string
}
