package parser

import (
	"crank/ast"
	"crank/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foo = 83;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram returned nothing")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("statement shouldn't have more than 3 statements. Got %d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		if !testLetStatement(t, program.Statements[i], tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Expected let. Got %q", s.TokenLiteral())
		return false
	}
	letstmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not Let statement. got %T", s)
		return false
	}

	if letstmt.Name.Value != name {
		t.Errorf("letstmt.Name.Value not '%s'. got=%s", name, letstmt.Name.Value)
		return false
	}

	if letstmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letstmt.Name.Value)
		return false
	}

	return true
}
