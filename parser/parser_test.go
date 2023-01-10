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
		ast.NumberNode{Value: fmt.Sprint(1)},
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
