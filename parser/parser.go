package parser

import "bachelor-thesis/parser/ast"

type Parser struct {
	input     string
	tokens    []Token
	currToken Token
	currPos   int
}

func (parser *Parser) parsePrimaryExpression() ast.Node {
	token := parser.currToken
	switch token.tokenType {
	case number:
		return ast.NumberNode{Value: token.val}
	default:
		return nil
	}
}

func (parser *Parser) parsePrimary() ast.Node {
	return parser.parsePrimaryExpression()
}

func (parser *Parser) parseExpression() ast.Node {
	node := parser.parsePrimary()
	return node
}

func Parse(input string) ast.Node {
	tokens := lex(input)

	parser := &Parser{
		input:     input,
		tokens:    tokens,
		currToken: tokens[0],
	}

	node := parser.parseExpression()
	return node
}
