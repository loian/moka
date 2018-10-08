package token

const (
	ILLEGAL = iota
	EOF

	// Literals and identifiers
	// 2
	IDENTIFIER
	BOOLEAN
	VAL_INT
	VAL_FLOAT
	VAL_DECIMAL
	VAL_STRING

	//Operator
	//8
	ASSIGN
	PLUS

	//Delimiters
	//10
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACKET
	RBRACKET
	LBRACE
	RBRACE

	//Keywords
	//18
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
	"bool":	TYPE_BOOLEAN,
	"int":	TYPE_INT,
	"float": TYPE_FLOAT,
	"decimal": TYPE_DECIMAL,
	"string": TYPE_STRING,
	"struct": TYPE_STRUCT,
}


func LookupIdentifier(i string) TokenType {
	if tokenType, ok := keywords[i]; ok {
		return tokenType
	}

	return IDENTIFIER
}
