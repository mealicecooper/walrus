package parser

import (
	"fmt"
	"walrus/ast"
	"walrus/errors"
	"walrus/lexer"
)

// parseExpr parses an expression with the given binding power.
// It first parses the NUD (Null Denotation) of the expression,
// then continues to parse the LED (Left Denotation) of the expression
// until the binding power of the current token is less than or equal to the given binding power.
// The parsed expression is returned as an ast.Node.
func parseExpr(p *Parser, bp BINDING_POWER) ast.Node {

	// Fist parse the NUD

	token := p.currentToken()

	tokenKind := token.Kind

	nudFunction, exists := NUDLookup[tokenKind]

	if !exists {

		var msg string
		if lexer.IsKeyword(string(tokenKind)) {
			msg = fmt.Sprintf("parser:nud:unexpected keyword '%s'\n", tokenKind)
		} else {
			msg = fmt.Sprintf("parser:nud:unexpected token '%s'\n", tokenKind)
		}
		errors.MakeError(p.FilePath, p.currentToken().Start.Line, p.currentToken().Start.Column, p.currentToken().End.Column, msg).Display()
	}

	left := nudFunction(p)

	for GetBP(p.currentTokenKind()) > bp {

		tokenKind = p.currentTokenKind()

		ledFunction, exists := LEDLookup[tokenKind]

		if !exists {
			msg := fmt.Sprintf("parser:led:unexpected token %s\n", tokenKind)
			errors.MakeError(p.FilePath, p.currentToken().Start.Line, p.currentToken().Start.Column, p.currentToken().End.Column, msg).Display()
		}

		left = ledFunction(p, left, GetBP(p.currentTokenKind()))
	}

	return left
}

// parsePrimaryExpr parses a primary expression in the input stream.
// It handles numeric literals, string literals, identifiers, boolean literals, and null literals.
// If the current token does not match any of these types, it panics with an error message.
func parsePrimaryExpr(p *Parser) ast.Node {

	startpos := p.currentToken().Start

	endpos := p.currentToken().End

	primaryToken := p.advance()

	rawValue := primaryToken.Value

	loc := ast.Location{
		Start: startpos,
		End:   endpos,
	}

	switch primaryToken.Kind {
	case lexer.INT:
		return ast.IntegerLiteralExpr{
			Value:    rawValue,
			Location: loc,
		}
	case lexer.FLOAT:

		return ast.FloatLiteralExpr{
			Value:    rawValue,
			Location: loc,
		}

	case lexer.STR:
		return ast.StringLiteralExpr{
			Value:    rawValue,
			Location: loc,
		}
	case lexer.CHR:
		return ast.CharLiteralExpr{
			Value:    rawValue,
			Location: loc,
		}
	case lexer.BOOL:
		return ast.BooleanLiteralExpr{
			Value:    rawValue,
			Location: loc,
		}
	case lexer.IDENTIFIER_TOKEN:
		return ast.IdentifierExpr{
			Name:     rawValue,
			Location: loc,
		}
	default:
		msg := fmt.Sprintf("Cannot create primary expression from %s\n", primaryToken.Value)
		errors.MakeError(p.FilePath, p.currentToken().Start.Line, p.currentToken().Start.Column, p.currentToken().End.Column, msg).Display()
	}

	return nil
}

func parseVarAssignmentExpr(p *Parser, left ast.Node, bp BINDING_POWER) ast.Node {

	start := p.currentToken().Start

	switch left.(type) {
	case ast.IdentifierExpr:
		break
	case ast.ArrayIndexAccess:
		break
	default:
		errMsg := "Cannot assign to a non-identifier\n"
		errors.MakeError(p.FilePath, left.StartPos().Line, left.StartPos().Column, left.EndPos().Column, errMsg).Display()
	}

	operator := p.advance()

	right := parseExpr(p, bp)

	endPos := right.EndPos()

	return ast.VarAssignmentExpr{
		Assignee: left,
		Value:    right,
		Operator: operator,
		Location: ast.Location{
			Start: start,
			End:   endPos,
		},
	}
}

func parseArrayExpr(p *Parser) ast.Node {

	start := p.advance().Start //eat the [ token

	var values []ast.Node

	for p.currentTokenKind() != lexer.CLOSE_BRACKET {
		value := parseExpr(p, PRIMARY_BP)
		values = append(values, value)
		if p.currentTokenKind() != lexer.CLOSE_BRACKET {
			p.expect(lexer.COMMA)
		}
	}

	end := p.expect(lexer.CLOSE_BRACKET).End

	return ast.ArrayExpr{
		Values: values,
		Location: ast.Location{
			Start: start,
			End:   end,
		},
	}
}

func parseArrayAccess(p *Parser, left ast.Node, bp BINDING_POWER) ast.Node {
	start := p.expect(lexer.OPEN_BRACKET).Start
	index := parseExpr(p, bp)
	end := p.expect(lexer.CLOSE_BRACKET).End
	return ast.ArrayIndexAccess{
		Arrayvalue: left,
		Index:      index,
		Location: ast.Location{
			Start: start,
			End:   end,
		},
	}
}