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
			{tokenType: number, val: "1"},
			{tokenType: number, val: "1.2"},
			{tokenType: number, val: "1e2"},
			{tokenType: number, val: "1.2e-3"},
			{tokenType: number, val: "01"},
			{tokenType: number, val: "0."},
			{tokenType: number, val: "1.2e+33"},
			{tokenType: number, val: ".1"},
			{tokenType: eof},
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
