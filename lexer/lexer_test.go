package lexer

import (
	"moka/token"
	"testing"
)

func TestNext(t *testing.T) {
	input := `var counter int = 5; var j,i float = 5.4, 3;
			  var add = fn(x int, y int) {
					return x+y;
              }
			  if (5<6) {
					return true;
			  else {
					return false;
			  }`

	expectations := []token.Token{
		{Type: token.VAR, Literal: "var"},
		{Type: token.IDENTIFIER, Literal: "counter"},
		{Type: token.TYPE_INT, Literal: "int"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.VAL_INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.VAR, Literal: "var"},
		{Type: token.IDENTIFIER, Literal: "j"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "i"},
		{Type: token.TYPE_FLOAT, Literal: "float"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.VAL_FLOAT, Literal: "5.4"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.VAL_INT, Literal: "3"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.VAR, Literal: "var"},
		{Type: token.IDENTIFIER, Literal: "add"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.FUNCTION, Literal: "fn"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.TYPE_INT, Literal: "int"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.TYPE_INT, Literal: "int"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
	}
	l := NewLexer(input)

	for i, expected := range expectations {
		token := l.NextToken()

		if token.Type != expected.Type {
			t.Fatalf("test [%d] - wrokng token type, expected %d, received %d, Literal received: %q",
				i, expected.Type, token.Type, token.Literal)
		}

		if token.Type != expected.Type {
			t.Fatalf("test [%d] - wrokng token type, expected %q, received %q",
				i, expected.Literal, token.Literal)
		}

	}

}

func TestReadNumberInt(t *testing.T) {
	input := "12345"
	l := NewLexer(input)
	tok, tokType := l.readNumber()
	if tokType != token.VAL_INT {
		t.Fatalf("test ReadNumberInt - wrokng token type, expected %q, received %q", token.VAL_INT, tokType)
	}
	if string(tok) != input {
		t.Fatalf("test ReadNumberInt - wrokng literal %q, received %q", input, tok)
	}
}

func TestReadNumberFloat(t *testing.T) {
	input := "1.2345"
	l := NewLexer(input)
	tok, tokType := l.readNumber()
	if tokType != token.VAL_FLOAT {
		t.Fatalf("test ReadNumberFloat - wrokng token type, expected %q, received %q", token.VAL_FLOAT, tokType)
	}
	if string(tok) != input {
		t.Fatalf("test ReadNumberFloat - wrokng literal %q, received %q", input, tok)
	}
}

func TestReadNumberIllegal(t *testing.T) {
	input := "1.234.5"
	l := NewLexer(input)
	tok, tokType := l.readNumber()
	if tokType != token.ILLEGAL {
		t.Fatalf("test ReadNumberIllegal - wrokng token type, expected %q, received %q", token.ILLEGAL, tokType)
	}
	if string(tok) != input {
		t.Fatalf("test ReadNumberIllegal - wrokng literal %q, received %q", input, tok)
	}
}
