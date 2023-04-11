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
		`a + b * c`,
		&ast.BinaryNode{Operator: "+",
			Left: &ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
			Right: &ast.BinaryNode{Operator: "*",
				Left:  &ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier},
				Right: &ast.IdentifierNode{Value: "c", NodeType: ast.NodeIdentifier}}},
	},
	{
		`a * b + c`,
		&ast.BinaryNode{Operator: "+",
			Left: &ast.BinaryNode{Operator: "*",
				Left:  &ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
				Right: &ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier}},
			Right: &ast.IdentifierNode{Value: "c", NodeType: ast.NodeIdentifier}},
	},
	{
		"a and b or c",
		&ast.BinaryNode{Operator: "and",
			Left: &ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
			Right: &ast.BinaryNode{Operator: "or",
				Left:  &ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier},
				Right: &ast.IdentifierNode{Value: "c", NodeType: ast.NodeIdentifier}}},
	},
	{
		"(a + b)",
		&ast.BinaryNode{Operator: "+",
			Left:  &ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
			Right: &ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier}},
	},
	{
		"(a + b) * c",
		&ast.BinaryNode{Operator: "*",
			Left: &ast.BinaryNode{Operator: "+",
				Left:  &ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
				Right: &ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier}},
			Right: &ast.IdentifierNode{Value: "c", NodeType: ast.NodeIdentifier}},
	},
	{
		"(a != b) and (c >= b)",
		&ast.BinaryNode{
			Operator: "and",
			Left: &ast.BinaryNode{Operator: "!=",
				Left:  &ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
				Right: &ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier}},
			Right: &ast.BinaryNode{Operator: ">=",
				Left:  &ast.IdentifierNode{Value: "c", NodeType: ast.NodeIdentifier},
				Right: &ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier}},
		},
	},
	{
		"foo()",
		&ast.CallNode{Callee: &ast.IdentifierNode{Value: "foo", NodeType: ast.NodeIdentifier},
			Arguments: []ast.Node{},
			NodeType:  ast.NodeCall,
		},
	},
	{
		"foo(a)",
		&ast.CallNode{Callee: &ast.IdentifierNode{Value: "foo", NodeType: ast.NodeIdentifier},
			Arguments: []ast.Node{&ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier}},
			NodeType:  ast.NodeCall,
		},
	},
	{
		"foo(a, b)",
		&ast.CallNode{Callee: &ast.IdentifierNode{Value: "foo", NodeType: ast.NodeIdentifier},
			Arguments: []ast.Node{&ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
				&ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier}},
			NodeType: ast.NodeCall,
		},
	},
	{
		`foo("arg1", 2, true)`,
		&ast.CallNode{Callee: &ast.IdentifierNode{Value: "foo", NodeType: ast.NodeIdentifier},
			Arguments: []ast.Node{&ast.StringNode{Value: "arg1", NodeType: ast.NodeString},
				&ast.NumberNode{Value: fmt.Sprint(2), Int64: 2, IsInt: true, IsFloat: false, NodeType: ast.NodeNumber},
				&ast.BoolNode{Value: true, NodeType: ast.NodeBool}},
			NodeType: ast.NodeCall,
		},
	},
	{
		"foo(baz())",
		&ast.CallNode{Callee: &ast.IdentifierNode{Value: "foo", NodeType: ast.NodeIdentifier},
			Arguments: []ast.Node{&ast.CallNode{Callee: &ast.IdentifierNode{Value: "baz", NodeType: ast.NodeIdentifier},
				Arguments: []ast.Node{},
				NodeType:  ast.NodeCall}},
			NodeType: ast.NodeCall,
		},
	},
	{
		"[]",
		&ast.ArrayNode{
			Nodes:    []ast.Node{},
			NodeType: ast.NodeArray,
		},
	},
	{
		"[a, b]",
		&ast.ArrayNode{
			Nodes: []ast.Node{&ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
				&ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier}},
			NodeType: ast.NodeArray,
		},
	},
	{
		"[a, b][0]",
		&ast.MemberNode{
			NodeType: ast.NodeMember,
			Node: &ast.ArrayNode{
				Nodes: []ast.Node{&ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
					&ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier}},
				NodeType: ast.NodeArray,
			},
			Property: &ast.NumberNode{Value: fmt.Sprint(0), Int64: 0, IsInt: true, IsFloat: false, NodeType: ast.NodeNumber}},
	},
	{
		"[a, [b, c]]",
		&ast.ArrayNode{
			Nodes: []ast.Node{&ast.IdentifierNode{Value: "a", NodeType: ast.NodeIdentifier},
				&ast.ArrayNode{
					Nodes: []ast.Node{&ast.IdentifierNode{Value: "b", NodeType: ast.NodeIdentifier},
						&ast.IdentifierNode{Value: "c", NodeType: ast.NodeIdentifier}},
					NodeType: ast.NodeArray,
				}},
			NodeType: ast.NodeArray,
		},
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
