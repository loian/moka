package lexer

import (
	"moka/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `+=(){},`

	expectations := []token.Token{
		{
			token.PLUS,
			"+",
		},
		{
			token.ASSIGN,
			"=",
		},
		{
			token.LPAREN,
			"(",
		},
		{
			token.RPAREN,
			")",
		},
		{token.LBRACE,
			"{",
		},
		{token.RBRACE,
			"}",
		},
	}

	l := NewLexer(input)

	for i, expected := range expectations {
		token := l.NextToken()

		if token.Type != expected.Type {
			t.Fatalf("test [%d] - wrokng token type, expected %d, received %d",
				i, expected.Type, token.Type)
		}

		if token.Type != expected.Type {
			t.Fatalf("test [%d] - wrokng token type, expected %q, received %q",
				i, expected.Literal, token.Literal)
		}

	}
}
