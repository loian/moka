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
	MINUS
	BANG
	ASTERISK
	SLASH

	LT
	GT

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
	ELSE
	FUNCTION
	FALSE
	IF
	RETURN
	TRUE
	TYPE_BOOLEAN
	TYPE_INT
	TYPE_FLOAT
	TYPE_DECIMAL
	TYPE_STRING
	TYPE_STRUCT
	VAR
)

type TokenType uint8

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":      FUNCTION,
	"var":     VAR,
	"if":      IF,
	"else":    ELSE,
	"true":    TRUE,
	"false":   FALSE,
	"return":  RETURN,
	"bool":    TYPE_BOOLEAN,
	"int":     TYPE_INT,
	"float":   TYPE_FLOAT,
	"decimal": TYPE_DECIMAL,
	"string":  TYPE_STRING,
	"struct":  TYPE_STRUCT,
}

func LookupIdentifier(i string) TokenType {
	if tokenType, ok := keywords[i]; ok {
		return tokenType
	}

	return IDENTIFIER
}
