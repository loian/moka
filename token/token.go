package token

const (
	ILLEGAL = iota
	EOF

	// Literals and identifiers
	IDENTIFIER
	BOOLEAN
	VAL_INT
	VAL_FLOAT
	VAL_DECIMAL
	VAL_STRING

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
	VAR
	TYPE_BOOLEAN
	TYPE_INT
	TYPE_FLOAT
	TYPE_DECIMAL
	TYPE_STRING
	TYPE_STRUCT
)

type TokenType uint8

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string] TokenType {
	"fn": FUNCTION,
	"var": VAR,
}


func LookupIdentifier(i string) TokenType {
	if tokenType, ok := keywords[i]; ok {
		return tokenType
	}

	return IDENTIFIER
}
