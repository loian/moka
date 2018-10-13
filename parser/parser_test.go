package parser

import (
	"fmt"
	"moka/ast"
	"moka/lexer"
	"testing"
)

func TestVarStatementErrors(t *testing.T) {
	input := `
		var x int 5;
		var y  = 10;
		var foobar int;
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	p.ParseProgram()

	if len(p.Errors()) <= 3 {
		t.Errorf("expected at least 3 errors, got %d", len(p.Errors()))
	}
}

func TestVarStatement(t *testing.T) {
	input := `
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
		t.Errorf("wrong number of statements. They should be 3, got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
		expectedType       string
	}{
		{"x", "int"},
		{"y", "int"},
		{"foobar", "int"},
	}

	for i, tt := range tests {
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

func TestReturnStatements(t *testing.T) {
	input := `return 5;
			return 10;
			return 912233;`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Errorf("wrong number of statements. They should be 3, got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {

		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt not a return statement, got %T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returmStmt.tokenLiteral not 'return', got '%s", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "cat;"

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Errorf("Expecred 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Errorf("program.Statements[0] is not an ExpressionStatement, got %T ", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Errorf("expression is not an *Identifier, got %T", stmt.Expression)
	}

	if ident.Value != "cat" {
		t.Errorf("ident.Value is not 'cat', got %q", ident.Value)
	}

	if ident.TokenLiteral() != "cat" {
		t.Errorf("tokenLiteral is not 'cat', got %q", ident.TokenLiteral())
	}

}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not an ast.ExpressionStatement, got %T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expression not a *ast.IntegerLiteral, got %T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Fatalf("literal.Value not 5, got %d", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Fatalf("literal.TokenLiteral() not \"5\", got %q", literal.TokenLiteral())
	}

}

func TestPrefixExpression(t *testing.T) {
	testCases := []struct {
		input    string
		operator string
		intValue int64
	}{
		{input: "!5", operator: "!", intValue: 5},
		{input: "-15", operator: "-", intValue: 15},
	}

	for i, tc := range testCases {
		l := lexer.NewLexer(tc.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("wrong number of statement in testCase[%d], expected 1, got %d", i, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf(
				"program.Statements[0] is not an *ast.ExpressionStatement in testCase[%d], got %T",
				i,
				stmt)
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf(
				"stmt.Expression is not an ast.Expression in testCase[%d], got %T",
				i,
				stmt.Expression)

		}

		if exp.Operator != tc.operator {
			t.Fatalf("operator is not %q in testCases[%d], got %q", tc.operator, i, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tc.intValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il not an integer literal, got %T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d, got %d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteran() != %q", fmt.Sprintf("%d", value))
		return false
	}

	return true
}
