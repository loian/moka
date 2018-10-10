package parser

import (
	"moka/ast"
	"moka/lexer"
	"testing"
)

func TestVarStatementErrors (t *testing.T) {
	input:=`
		var x int 5;
		var y  = 10;
		var foobar int ;
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	p.ParseProgram()

	if len(p.Errors()) != 3 {
		t.Errorf("expected 3 errors, got %d", len(p.Errors()))
	}
}

func TestVarStatement (t *testing.T) {
	input:=`
		var x int = 5;
		var y int = 10;
		var foobar int = 44;
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if program == nil {
		t.Errorf("ParseProgram returned nil")
		return
	}

	if len(program.Statements) != 3 {
		t.Errorf("Wrong number of statements. They should be 3, got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
		expectedType string
	}{
		{"x", "int"},
		{"y", "int"},
		{"foobar", "int"},
	}

	for i,tt := range tests {
		stmt := program.Statements[i]
		if !testVarStatement(t, stmt, tt.expectedIdentifier, tt.expectedType) {
			return
		}
	}
}

func testVarStatement(t *testing.T, s ast.Statement, name string, typeident string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("token literal not var, got %s", s.TokenLiteral())
		return false
	}

	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("not a VarStatement, got %s ", s)
		return false
	}

	if varStmt.Name.Value != name {
		t.Errorf("varStatement.Name.Value not %s, got %s ", name, varStmt.Name.Value)
		return false
	}

	if varStmt.Type.Value != typeident {
		t.Errorf("varStatement.Type.Value not %s, got %s ", typeident, varStmt.Type.Value)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

