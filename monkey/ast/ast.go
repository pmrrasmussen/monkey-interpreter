package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String() + " = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral())
	if rs.ReturnValue != nil {
		out.WriteString(" " + rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, statement := range bs.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fn *FunctionLiteral) TokenLiteral() string { return fn.Token.Literal }
func (fn *FunctionLiteral) expressionNode()      {}
func (fn *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, param := range fn.Parameters {
		params = append(params, param.String())
	}

	out.WriteString(fn.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")

	out.WriteString(fn.Body.String())

	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ifExp *IfExpression) TokenLiteral() string { return ifExp.Token.Literal }
func (ifExp *IfExpression) expressionNode()      {}
func (ifExp *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ifExp.Condition.String())
	out.WriteString(" ")
	out.WriteString(ifExp.Consequence.String())

	if ifExp.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ifExp.Alternative.String())
	}

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (in *InfixExpression) TokenLiteral() string { return in.Token.Literal }
func (in *InfixExpression) expressionNode()      {}
func (in *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(in.Left.String())
	out.WriteString(" " + in.Operator + " ")
	out.WriteString(in.Right.String())
	out.WriteString(")")

	return out.String()
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) TokenLiteral() string { return p.Token.Literal }
func (p *PrefixExpression) expressionNode()      {}
func (p *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")

	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) expressionNode()      {}
func (i *Identifier) String() string {
	return i.Value
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (i *Boolean) TokenLiteral() string { return i.Token.Literal }
func (i *Boolean) expressionNode()      {}
func (i *Boolean) String() string       { return i.Token.Literal }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) String() string       { return i.Token.Literal }
