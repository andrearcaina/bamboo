package parser

import (
	"github.com/andrearcaina/bamboo/ast"
	"github.com/andrearcaina/bamboo/lexer"
	"testing"
)

func TestStatements(t *testing.T) {
	input := `
		let x: int64 = 5
		let y: int32 = 10
		let foobar: int = 838383
		const bar: string = "hello"
		const foo: bool = true
		const baz: bool = false
		const qux: float64 = 3.14
		const t: char = 'a'
`
	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 8 {
		t.Fatalf("program.Statements does not contain 8 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
		expectedType       string
	}{
		{"x", "int64"},
		{"y", "int32"},
		{"foobar", "int"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testStatement(t, stmt, tt.expectedIdentifier, tt.expectedType) {
			return
		}
	}
}

func testStatement(t *testing.T, s ast.Statement, name string, expectedType string) bool {
	if s.TokenLiteral() == "let" {
		return testLetStmt(t, s, name, expectedType)
	} else if s.TokenLiteral() == "const" {
		return testConstStmt(t, s, name, expectedType)
	} else {
		t.Errorf("s not a let or const statement. got=%T", s)
		return false
	}
}

func testConstStmt(t *testing.T, s ast.Statement, name string, expectedType string) bool {
	if s.TokenLiteral() != "const" {
		t.Errorf("s.TokenLiteral not 'const'. got=%q", s.TokenLiteral())
		return false
	}

	stmt, ok := s.(*ast.ConstStatement)
	if !ok {
		t.Errorf("s not *ast.ConstStatement. got=%T", s)
		return false
	}

	if stmt.Name.Value != name {
		t.Errorf("constStmt.Name.Value not '%s'. got=%s", name, stmt.Name.Value)
		return false
	}

	if stmt.Type.Value != expectedType {
		t.Errorf("constStmt.Type.Value not '%s'. got=%s", expectedType, stmt.Type.Value)
		return false

	}

	if stmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, stmt.Name)
		return false
	}

	return true
}

func testLetStmt(t *testing.T, s ast.Statement, name string, expectedType string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	stmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if stmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, stmt.Name.Value)
		return false
	}

	if stmt.Type.Value != expectedType {
		t.Errorf("letStmt.Type.Value not '%s'. got=%s", expectedType, stmt.Type.Value)
		return false

	}

	if stmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, stmt.Name)
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
