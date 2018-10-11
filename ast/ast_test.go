package ast

import (
	"moka/token"
	"testing"
)

func TestString(t *testing.T) {

	programString := "var myVar int = anotherVar;"
	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.VAR, Literal: "var"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
				},
				Type: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "int"},
					Value: "int",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != programString {
		t.Errorf("program.String() wrong. expected %q, got %q", programString, program.String())
	}
}

