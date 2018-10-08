package lexer

import (
	"moka/token"
	"testing"
)

func TestNextTokenSingleChars(t *testing.T) {
	input := `var counter int = 5`

	expectations := []token.Token{
		{
			Type: token.VAR,
			Literal: "var",
		},
		{
			Type: token.IDENTIFIER,
			Literal: "counter",
		},
		{
			Type: token.TYPE_INT,
			Literal: "int",
		},
		{
			Type: token.ASSIGN,
			Literal: "=",
		},
		{
			Type: token.VAL_INT,
			Literal: "5",
		},
		{
			Type: token.SEMICOLON,
			Literal: ";",
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

func TestNextToken(t *testing.T) {
}
