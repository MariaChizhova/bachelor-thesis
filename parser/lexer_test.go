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
			{tokenType: itemNumber, val: "1", pos: 0},
			{tokenType: itemNumber, val: "1.2", pos: 2},
			{tokenType: itemNumber, val: "1e2", pos: 4},
			{tokenType: itemNumber, val: "1.2e-3", pos: 6},
			{tokenType: itemNumber, val: "01", pos: 8},
			{tokenType: itemNumber, val: "0.", pos: 10},
			{tokenType: itemNumber, val: "1.2e+33", pos: 12},
			{tokenType: itemNumber, val: ".1", pos: 14},
			{tokenType: itemEOF, pos: 15},
		},
	},
	{
		"1 + 2 - 3 % 4 / 5",
		[]Token{
			{tokenType: itemNumber, val: "1", pos: 0},
			{tokenType: itemOperator, val: "+", pos: 2},
			{tokenType: itemNumber, val: "2", pos: 4},
			{tokenType: itemOperator, val: "-", pos: 6},
			{tokenType: itemNumber, val: "3", pos: 8},
			{tokenType: itemOperator, val: "%", pos: 10},
			{tokenType: itemNumber, val: "4", pos: 12},
			{tokenType: itemOperator, val: "/", pos: 14},
			{tokenType: itemNumber, val: "5", pos: 16},
			{tokenType: itemEOF, pos: 17},
		},
	},
	{
		"(1 + 2)",
		[]Token{
			{tokenType: itemBracket, val: "(", pos: 0},
			{tokenType: itemNumber, val: "1", pos: 1},
			{tokenType: itemOperator, val: "+", pos: 3},
			{tokenType: itemNumber, val: "2", pos: 5},
			{tokenType: itemBracket, val: ")", pos: 6},
			{tokenType: itemEOF, pos: 7},
		},
	},
	{
		"true + false",
		[]Token{
			{tokenType: itemBool, val: "true", pos: 0},
			{tokenType: itemOperator, val: "+", pos: 5},
			{tokenType: itemBool, val: "false", pos: 7},
			{tokenType: itemEOF, pos: 8},
		},
	},
	{
		"a and b or c",
		[]Token{
			{tokenType: itemIdentifier, val: "a", pos: 0},
			{tokenType: itemOperator, val: "and", pos: 2},
			{tokenType: itemIdentifier, val: "b", pos: 6},
			{tokenType: itemOperator, val: "or", pos: 8},
			{tokenType: itemIdentifier, val: "c", pos: 10},
			{tokenType: itemEOF, pos: 11},
		},
	},
	{
		`a = "abc"`,
		[]Token{
			{tokenType: itemIdentifier, val: "a", pos: 0},
			{tokenType: itemOperator, val: "=", pos: 2},
			{tokenType: itemString, val: "abc", pos: 4},
			{tokenType: itemEOF},
		},
	},
	{
		`'abc'`,
		[]Token{
			{tokenType: itemString, val: "abc"},
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
