package token

const (
	EOF = iota
	ILLEGAL

	// Literals and identifiers
	IDENTIFIER
	VAL_DECIMAL
	VAL_FLOAT
	VAL_INT
	VAL_STRING

	//Operator
	ASSIGN
	ASTERISK
	BANG
	EQUAL
	GT
	LT
	MINUS
	NOT_EQUAL
	PLUS
	SLASH

	//10
	COMMA
	LBRACE
	LBRACKET
	LPAREN
	RBRACE
	RBRACKET
	RPAREN
	SEMICOLON

	//Keywords
	//18
	ELSE
	FUNCTION
	FALSE
	IF
	RETURN
	TRUE
	//TYPE_BOOLEAN
	//TYPE_DECIMAL
	//TYPE_FLOAT
	//TYPE_INT
	//TYPE_STRING
	//TYPE_STRUCT
	VAR
)

type TokenType uint8

type Token struct {
	Type    TokenType
	Literal string
}

var tokenTypeNames = map[TokenType]string{
	EOF:         "EOF",
	ILLEGAL:     "ILLEGAL",
	IDENTIFIER:  "IDENTIFIER",
	VAL_DECIMAL: "VAL_DECIMAL",
	VAL_FLOAT:   "VAL_FLOAT",
	VAL_INT:     "VAL_INT",
	VAL_STRING:  "VAL_STRING",
	ASSIGN:      "ASSIGN",
	ASTERISK:    "ASTERISK",
	BANG:        "BANG",
	EQUAL:       "EQUAL",
	GT:          "GT",
	LT:          "LT",
	MINUS:       "MINUS",
	NOT_EQUAL:   "NOT_EQUAL",
	PLUS:        "PLUS",
	SLASH:       "SLASH",
	COMMA:       "COMMA",
	LBRACE:      "LBRACE",
	LBRACKET:    "LBRACKET",
	LPAREN:      "LPAREN",
	RBRACE:      "RBRACE",
	RBRACKET:    "RBRACKET",
	RPAREN:      "RPAREN",
	SEMICOLON:   "SEMICOLON",
	ELSE:        "ELSE",
	FUNCTION:    "FUNCTION",
	FALSE:       "FALSE",
	IF:          "IF",
	RETURN:      "RETURN",
	TRUE:        "TRUE",
	//TYPE_BOOLEAN :"TYPE_BOOLEAN",
	//TYPE_DECIMAL :"TYPE_DECIMAL",
	//TYPE_FLOAT :"TYPE_FLOAT",
	//TYPE_INT :"TYPE_INT",
	//TYPE_STRING :"TYPE_STRING",
	//TYPE_STRUCT :"TYPE_STRUCT",
	VAR: "VAR",
}

var keywords = map[string]TokenType{
	"fn":      FUNCTION,
	"var":     VAR,
	"if":      IF,
	"else":    ELSE,
	"true":    TRUE,
	"false":   FALSE,
	"return":  RETURN,
	"bool":    IDENTIFIER,
	"int":     IDENTIFIER,
	"float":   IDENTIFIER,
	"decimal": IDENTIFIER,
	"string":  IDENTIFIER,
	"struct":  IDENTIFIER,
}

func LookupIdentifier(i string) TokenType {
	if tokenType, ok := keywords[i]; ok {
		return tokenType
	}

	return IDENTIFIER
}

func LookupTokenTypeName(tokenType TokenType) string {
	if t, ok := tokenTypeNames[tokenType]; ok {
		return t
	}

	return "Unknown"

}
