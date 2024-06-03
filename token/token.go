package token

// TokenType is a string that represents a token type
type TokenType string

// Token is a struct that represents a token with a type and a literal value
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	EQ = "=="
	NE = "!="
	LT = "<"
	LE = "<="
	GT = ">"
	GE = ">="

	COMMENT   = "#"
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNC"
	LET      = "LET"
	CONST    = "CONST"
)
