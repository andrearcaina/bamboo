package lexer

import "github.com/andrearcaina/bamboo/token"

// isNumber checks if a character is a number
func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

// isLetter checks if a character is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// newToken creates a new token with a type and a literal value
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
