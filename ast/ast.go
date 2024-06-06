package ast

import (
	"github.com/andrearcaina/bamboo/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Type  *Identifier
	Name  *Identifier
	Value Expression
}

type ConstStatement struct {
	Token token.Token // the token.CONST token
	Name  *Identifier
	Type  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (cs *ConstStatement) statementNode()       {}
func (cs *ConstStatement) TokenLiteral() string { return cs.Token.Literal }

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
