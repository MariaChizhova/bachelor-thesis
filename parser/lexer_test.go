package parser

import (
	"testing"
)

type lexTest struct {
	input  string
	tokens []Token
}

var lexTests = []lexTest{
	{
		"1 1.2 1e2 1.2e-3 01 0. 1.2e+33 .1",
		[]Token{
			{tokenType: itemNumber, val: "1"},
			{tokenType: itemNumber, val: "1.2"},
			{tokenType: itemNumber, val: "1e2"},
			{tokenType: itemNumber, val: "1.2e-3"},
			{tokenType: itemNumber, val: "01"},
			{tokenType: itemNumber, val: "0."},
			{tokenType: itemNumber, val: "1.2e+33"},
			{tokenType: itemNumber, val: ".1"},
			{tokenType: itemEOF},
		},
	},
	{
		"1 + 2 - 3 % 4 / 5",
		[]Token{
			{tokenType: itemNumber, val: "1"},
			{tokenType: itemOperator, val: "+"},
			{tokenType: itemNumber, val: "2"},
			{tokenType: itemOperator, val: "-"},
			{tokenType: itemNumber, val: "3"},
			{tokenType: itemOperator, val: "%"},
			{tokenType: itemNumber, val: "4"},
			{tokenType: itemOperator, val: "/"},
			{tokenType: itemNumber, val: "5"},
			{tokenType: itemEOF},
		},
	},
	{
		"(1 + 2)",
		[]Token{
			{tokenType: itemBracket, val: "("},
			{tokenType: itemNumber, val: "1"},
			{tokenType: itemOperator, val: "+"},
			{tokenType: itemNumber, val: "2"},
			{tokenType: itemBracket, val: ")"},
			{tokenType: itemEOF},
		},
	},
	{
		"true + false",
		[]Token{
			{tokenType: itemBool, val: "true"},
			{tokenType: itemOperator, val: "+"},
			{tokenType: itemBool, val: "false"},
			{tokenType: itemEOF},
		},
	},
	{
		"a and b or c",
		[]Token{
			{tokenType: itemIdentifier, val: "a"},
			{tokenType: itemOperator, val: "and"},
			{tokenType: itemIdentifier, val: "b"},
			{tokenType: itemOperator, val: "or"},
			{tokenType: itemIdentifier, val: "c"},
			{tokenType: itemEOF},
		},
	},
	{
		"a = 1",
		[]Token{
			{tokenType: itemIdentifier, val: "a"},
			{tokenType: itemOperator, val: "="},
			{tokenType: itemNumber, val: "1"},
			{tokenType: itemEOF},
		},
	},
}

func compareTokens(token1, token2 []Token) bool {
	if len(token1) != len(token2) {
		return false
	}
	for k := range token1 {
		if token1[k].tokenType != token2[k].tokenType || token1[k].val != token2[k].val {
			return false
		}
	}
	return true
}

func TestLexer(t *testing.T) {
	for _, test := range lexTests {
		tokens := lex(test.input)
		if !compareTokens(tokens, test.tokens) {
			t.Errorf("%s:\ngot\n\t%+v\nexpected\n\t%v", test.input, tokens, test.tokens)
		}
	}
}
