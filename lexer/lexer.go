package lexer

import (
	"strings"

	"github.com/andrearcaina/bamboo/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// NewLexer creates a new lexer with an input string
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character in the input and advances the position and readPosition pointers
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken reads the next token in the input and returns it as a token with a type and a literal value (if applicable)
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.consumeWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.EQ
			tok.Literal = "=="
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.NE
			tok.Literal = "!="
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.LE
			tok.Literal = "<="
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.GE
			tok.Literal = ">="
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '+':
		if l.peekChar() == '+' {
			l.readChar()
			tok.Type = token.PP
			tok.Literal = "++"
		} else {
			tok = newToken(token.PLUS, l.ch)
		}
	case '-':
		if l.peekChar() == '-' {
			l.readChar()
			tok.Type = token.MM
			tok.Literal = "--"
		} else {
			tok = newToken(token.MINUS, l.ch)
		}
	case '*':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.MU
			tok.Literal = "*="
		} else if l.peekChar() == '*' {
			l.readChar()
			tok.Type = token.EXP
			tok.Literal = "**"
		} else {
			tok = newToken(token.ASTERISK, l.ch)
		}
	case '/':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.DE
			tok.Literal = "/="
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
	case '%':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.MO
			tok.Literal = "%="
		} else {
			tok = newToken(token.MOD, l.ch)
		}
	case '&':
		tok = newToken(token.AND, l.ch)
	case '|':
		tok = newToken(token.OR, l.ch)
	case '^':
		tok = newToken(token.XOR, l.ch)
	case '#':
		tok = newToken(token.COMMENT, l.ch)
		l.readChar()
		for l.ch != '\n' {
			l.readChar()
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			if strings.Contains(tok.Literal, ".") {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// readIdentifier reads the next identifier in the input and advances the position and readPosition pointers
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads the next number in the input and advances the position and readPosition pointers
func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// consumeWhitespace consumes all whitespace characters
func (l *Lexer) consumeWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
