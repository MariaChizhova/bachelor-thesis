package parser

import (
	ast "bachelor-thesis/parser/ast"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Parser struct {
	input     string
	tokens    []Token
	currToken Token
	pos       int
}

var unaryOperators = map[string]bool{
	"-": true,
	"+": true,
}

var binaryOperators = map[string]bool{
	"or":  true,
	"and": true,
	"+":   true,
	"-":   true,
	"*":   true,
	"/":   true,
	"%":   true,
	"^":   true,
	"<":   true,
	"<=":  true,
	">":   true,
	">=":  true,
	"==":  true,
	"!=":  true,
}

func (parser *Parser) next() {
	if parser.pos+1 >= len(parser.tokens) {
		return
	}
	parser.pos++
	parser.currToken = parser.tokens[parser.pos]
}

func (parser *Parser) parsePrimaryExpression() ast.Node {
	token := parser.currToken
	parser.next()
	switch token.tokenType {
	case itemNumber:
		if strings.ContainsAny(token.val, ".eE") {
			number, _ := strconv.ParseFloat(token.val, 64)
			return &ast.NumberNode{Value: token.val, Float64: number, IsFloat: true, IsInt: false, NodeType: ast.NodeNumber}
		} else {
			number, _ := strconv.ParseInt(token.val, 10, 64)
			return &ast.NumberNode{Value: token.val, Int64: number, IsInt: true, IsFloat: false, NodeType: ast.NodeNumber}
		}
	case itemBool:
		return &ast.BoolNode{Value: token.val == "true", NodeType: ast.NodeBool}
	case itemNil:
		return &ast.NilNode{NodeType: ast.NodeNil}
	case itemString:
		return &ast.StringNode{Value: token.val, NodeType: ast.NodeString}
	case itemIdentifier:
		if parser.currToken.val == "(" {
			return parser.parseFunctionCall(token)
		}
		return &ast.IdentifierNode{Value: token.val, NodeType: ast.NodeIdentifier}
	default:
		var node ast.Node
		return node
	}
}

func (parser *Parser) parsePrimary() ast.Node {
	token := parser.currToken
	switch token.tokenType {
	case itemOperator:
		if unaryOperators[token.val] {
			parser.next()
			expr := parser.parseExpression()
			node := &ast.UnaryNode{Operator: token.val, Node: expr}
			return node
		}
	case itemBracket:
		if token.val == "(" {
			parser.next()
			expr := parser.parseExpression()
			if parser.currToken.val == ")" {
				parser.next()
			} else {
				parser.errorf("')' is expected")
			}
			return expr
		} else if token.val == "[" {
			return parser.parseArray()
		}
	}
	return parser.parsePrimaryExpression()
}

func (parser *Parser) parseExpression() ast.Node {
	left := parser.parsePrimary()
	token := parser.currToken
	if token.tokenType == itemOperator {
		if binaryOperators[token.val] {
			parser.next()
			right := parser.parseExpression()
			left = &ast.BinaryNode{
				Operator: token.val,
				Left:     left,
				Right:    right,
			}
		}
	}
	return left
}

func (parser *Parser) parseFunctionCall(token Token) ast.Node {
	arguments := make([]ast.Node, 0)
	for parser.currToken.val != ")" {
		if len(arguments) > 0 {
			if parser.currToken.val == "," {
				parser.next()
			} else {
				break
			}
		} else {
			parser.next()
		}
		node := parser.parseExpression()
		if node != nil && !reflect.ValueOf(node).IsNil() {
			arguments = append(arguments, node)
		} else {
			break
		}
	}
	if parser.currToken.val == ")" {
		parser.next()
	}
	return &ast.FunctionNode{
		Function:  &ast.IdentifierNode{Value: token.val, NodeType: ast.NodeIdentifier},
		Arguments: arguments,
		NodeType:  ast.NodeFunction,
	}
}

func (parser *Parser) parseArray() ast.Node {
	nodes := make([]ast.Node, 0)
	for parser.currToken.val != "]" {
		if len(nodes) > 0 {
			if parser.currToken.val == "," {
				parser.next()
			} else {
				parser.errorf("',' or ']` are expected")
				break
			}
		} else {
			parser.next()
		}
		node := parser.parseExpression()
		if node != nil && !reflect.ValueOf(node).IsNil() {
			nodes = append(nodes, node)
		} else {
			break
		}
	}
	if parser.currToken.val == "]" {
		parser.next()
	}
	return &ast.ArrayNode{
		Nodes:    nodes,
		NodeType: ast.NodeArray,
	}
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

func (parser *Parser) errorf(format string, args ...any) {
	format = fmt.Sprintf("Parse error: %s", format)
	panic(fmt.Errorf(format, args...))
}
