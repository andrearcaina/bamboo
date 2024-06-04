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
	INT8  = "INT8"
	INT16 = "INT16"
	INT32 = "INT32"
	INT64 = "INT64"

	UINT   = "UINT"
	UINT8  = "UINT8"
	UINT16 = "UINT16"
	UINT32 = "UINT32"
	UINT64 = "UINT64"

	FLOAT   = "FLOAT"
	FLOAT32 = "FLOAT32"
	FLOAT64 = "FLOAT64"

	STRING = "STRING"
	CHAR   = "CHAR"

	BOOL  = "BOOL"
	TRUE  = "TRUE"
	FALSE = "FALSE"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MOD      = "%"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	PP  = "++"
	MM  = "--"
	EXP = "**"
	PE  = "+="
	ME  = "-="
	MU  = "*="
	DE  = "/="
	MO  = "%="
	AND = "&"
	OR  = "|"
	XOR = "^"

	EQ = "=="
	NE = "!="
	LT = "<"
	LE = "<="
	GT = ">"
	GE = ">="

	COMMENT   = "#"
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNC"
	LET      = "LET"
	CONST    = "CONST"
	RETURN   = "RETURN"

	IF   = "IF"
	ELSE = "ELSE"
	ELIF = "ELSE IF"
)

var keywords = map[string]TokenType{
	"func":    FUNCTION,
	"let":     LET,
	"const":   CONST,
	"return":  RETURN,
	"if":      IF,
	"else":    ELSE,
	"else if": ELIF,
	"int":     INT64,
	"int8":    INT8,
	"int16":   INT16,
	"int32":   INT32,
	"int64":   INT64,
	"uint":    UINT64,
	"uint8":   UINT8,
	"uint16":  UINT16,
	"uint32":  UINT32,
	"uint64":  UINT64,
	"float":   FLOAT64,
	"float32": FLOAT32,
	"float64": FLOAT64,
	"string":  STRING,
	"char":    CHAR,
	"bool":    BOOL,
	"true":    TRUE,
	"false":   FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
