package lexer

import (
	"moka/token"
	"testing"
)

func TestNext(t *testing.T) {
	input := `var counter int = 5;
			  var j,i float = 5.4, 3;
			  var add = fn(x int, y int) {
					return x+y;
              }
			  10 == 10;
			  10 != 11;
			  `


	expectations := []token.Token{
		{Type: token.VAR, Literal: "var"},
		{Type: token.IDENTIFIER, Literal: "counter"},
		{Type: token.IDENTIFIER, Literal: "int"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.VAL_INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.VAR, Literal: "var"},
		{Type: token.IDENTIFIER, Literal: "j"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "i"},
		{Type: token.IDENTIFIER, Literal: "float"},
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
		{Type: token.IDENTIFIER, Literal: "int"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.IDENTIFIER, Literal: "int"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},

		{Type: token.VAL_INT, Literal: "10"},
		{Type: token.EQUAL, Literal: "="},
		{Type: token.VAL_INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.VAL_INT, Literal: "10"},
		{Type: token.NOT_EQUAL, Literal: "!="},
		{Type: token.VAL_INT, Literal: "11"},
		{Type: token.SEMICOLON, Literal: ";"},

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

func TestConsumeWhiteSpaces(t *testing.T) {
	input := " \t\n12345"; //one space, a tab, a new line
	l := NewLexer(input)
	l.consumeWhitespace()
	if l.position != 3 {
		t.Fatalf("test consumeWhitespace - should have consumed two characters and position should be 3 insted of %q", l.position);
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

func TestPeekChar(t *testing.T) {
	input := "12345";
	l := NewLexer(input);

	if (l.peekChar() != '2') {
		t.Fatalf("PeekChar should have returned 2")
	}
}

func TestPeekCharOverflow(t *testing.T) {
	input := "";
	l := NewLexer(input);

	if (l.peekChar() != 0) {
		t.Fatalf("PeekChar should have returned an empty rune")
	}
}