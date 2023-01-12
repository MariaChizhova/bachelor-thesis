package parser

import (
	"bachelor-thesis/parser/ast"
	"fmt"
	"reflect"
	"testing"
)

type parseTest struct {
	input    string
	expected ast.Node
}

var parseTests = []parseTest{
	{
		"1",
		&ast.NumberNode{Value: fmt.Sprint(1), Int64: 1, IsInt: true, IsFloat: false, NodeType: ast.NodeNumber},
	},
	{
		"1.2",
		&ast.NumberNode{Value: fmt.Sprint(1.2), Float64: 1.2, IsFloat: true, IsInt: false, NodeType: ast.NodeNumber},
	},
	{
		"true",
		&ast.BoolNode{Value: true, NodeType: ast.NodeBool},
	},
	{
		"nil",
		&ast.NilNode{NodeType: ast.NodeNil},
	},
	{
		`"abc"`,
		&ast.StringNode{Value: "abc", NodeType: ast.NodeString},
	},
	{
		`var`,
		&ast.IdentifierNode{Value: "var", NodeType: ast.NodeIdentifier},
	},
	{
		`+1`,
		&ast.UnaryNode{Operator: "+", Node: &ast.NumberNode{Value: "1", Int64: 1, IsInt: true, IsFloat: false, NodeType: ast.NodeNumber}},
	},
	{
		`a + b`,
		&ast.BinaryNode{Operator: "+",
			Left:  &ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
			Right: &ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier}},
	},
	{
		"a and b or c",
		&ast.BinaryNode{Operator: "and",
			Left: &ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
			Right: &ast.BinaryNode{Operator: "or",
				Left:  &ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier},
				Right: &ast.IdentifierNode{Value: "c", NodeType: ast.NodeIdentifier}}},
	},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		parseResult := Parse(test.input)
		if !reflect.DeepEqual(parseResult, test.expected) {
			t.Errorf("%s:\ngot\n\t%#v\nexpected\n\t%#v", test.input, parseResult, test.expected)
		}
	}
}
