package ast

import "walrus/lexer"

// Any word which is not a keyword or literal
type IdentifierExpr struct {
	Name string
	Location
}

func (a IdentifierExpr) INode() {
	//empty method implements Node interface
}
func (a IdentifierExpr) StartPos() lexer.Position {
	return a.Location.Start
}
func (a IdentifierExpr) EndPos() lexer.Position {
	return a.Location.End
}

//Literals or Raw values like: 1,2,3,4.6, "hello world", 'a' ...etc
type IntegerLiteralExpr struct {
	Value string
	Location
}

func (a IntegerLiteralExpr) INode() {
	//empty method implements Node interface
}
func (a IntegerLiteralExpr) StartPos() lexer.Position {
	return a.Location.Start
}
func (a IntegerLiteralExpr) EndPos() lexer.Position {
	return a.Location.End
}

type FloatLiteralExpr struct {
	Value string
	Location
}

func (a FloatLiteralExpr) INode() {
	//empty method implements Node interface
}
func (a FloatLiteralExpr) StartPos() lexer.Position {
	return a.Location.Start
}
func (a FloatLiteralExpr) EndPos() lexer.Position {
	return a.Location.End
}

type StringLiteralExpr struct {
	Value string
	Location
}

func (a StringLiteralExpr) INode() {
	//empty method implements Node interface
}
func (a StringLiteralExpr) StartPos() lexer.Position {
	return a.Location.Start
}
func (a StringLiteralExpr) EndPos() lexer.Position {
	return a.Location.End
}

type CharLiteralExpr struct {
	Value string
	Location
}

func (a CharLiteralExpr) INode() {
	//empty method implements Node interface
}
func (a CharLiteralExpr) StartPos() lexer.Position {
	return a.Location.Start
}
func (a CharLiteralExpr) EndPos() lexer.Position {
	return a.Location.End
}

type BooleanLiteralExpr struct {
	Value string
	Location
}

func (a BooleanLiteralExpr) INode() {
	//empty method implements Node interface
}
func (a BooleanLiteralExpr) StartPos() lexer.Position {
	return a.Location.Start
}
func (a BooleanLiteralExpr) EndPos() lexer.Position {
	return a.Location.End
}

type NullLiteralExpr struct {
	Value string
	Location
}

func (a NullLiteralExpr) INode() {
	//empty method implements Node interface
}
func (a NullLiteralExpr) StartPos() lexer.Position {
	return a.Location.Start
}
func (a NullLiteralExpr) EndPos() lexer.Position {
	return a.Location.End
}

type VarAssignmentExpr struct {
	Assignee Node // Check later if we should use IdentifierExpr instead
	Value    Node
	Operator lexer.Token // Looks odd right? Well, we know the operator must be '='. But what about +=, -=, *= and so on?😀
	Location
}

func (a VarAssignmentExpr) INode() {
	//empty method implements Node interface
}
func (a VarAssignmentExpr) StartPos() lexer.Position {
	return a.Location.Start
}
func (a VarAssignmentExpr) EndPos() lexer.Position {
	return a.Location.End
}

type ArrayExpr struct {
	Values []Node
	Location
}

func (a ArrayExpr) INode() {
	//empty method implements Node interface
}
func (a ArrayExpr) StartPos() lexer.Position {
	return a.Location.Start
}
func (a ArrayExpr) EndPos() lexer.Position {
	return a.Location.End
}

type ArrayIndexAccess struct {
	Index      Node
	Arrayvalue Node
	Location
}

func (a ArrayIndexAccess) INode() {
	//empty method implements Node interface
}
func (a ArrayIndexAccess) StartPos() lexer.Position {
	return a.Location.Start
}
func (a ArrayIndexAccess) EndPos() lexer.Position {
	return a.Location.End
}