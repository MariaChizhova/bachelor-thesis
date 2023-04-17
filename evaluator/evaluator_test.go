package evaluator

import (
	"bachelor-thesis/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type evaluatorTest struct {
	input    string
	expected interface{}
}

var evaluatorTests = []evaluatorTest{
	{"true", true},
	{"false", false},
	{"nil", nil},
	{"1", int64(1)},
	{"-1", int64(-1)},
	{"2.4", 2.4},
	{"-2.4", -2.4},
	{"1 + 2", int64(3)},
	{"1 - 2", int64(-1)},
	{"1 * 2", int64(2)},
	{"4 / 2", int64(2)},
	{"2 ^ 2", int64(4)},
	{"5 % 2", int64(1)},
	{"1 < 2", true},
	{"1 < 1", false},
	{"2 <= 2", true},
	{"1 > 0.2", true},
	{"1 >= 2", false},
	{"1 == 2", false},
	{"1 == 1", true},
	{"1 != 2", true},
	{"(1.1 + 2.1) * 4.1", 13.12},
	{"1.2 + 3  < 4", false},
	{"(1 + 2) * 4", int64(12)},
	{"5 * 2 + 10", int64(20)},
	{"5 + 2 * 10", int64(25)},
	{"10 / 2 * 2 + 10 - 5", int64(15)},
	{"5 * (2 + 10)", int64(60)},
	{"-5 + 10 + -5", int64(0)},
	{"2 * 2 * 2 * 2 * 2", int64(32)},
	{"1 - 2 + 3 * 4 / 2 ^ 2 % 3", int64(-1)},
	{"true == true", true},
	{"true == false", false},
	{"true != false", true},
	{"(1 < 2) == true", true},
	{"(1 < 2) == false", false},
	{`"aaa" == "aaa"`, true},
	{`"aaa" != "aab"`, true},
	{`"aaa" == "aab"`, false},
	{`"hello"`, "hello"},
	{`"hello " + "world!"`, "hello world!"},
	{"[]", []interface{}{}},
	{"[1, 2, 3.1]", []interface{}{int64(1), int64(2), 3.1}},
	{"[1 + 2, 2 * 3]", []interface{}{int64(3), int64(6)}},
	{"[1, 2, 3][1]", int64(2)},
	{"[1, 2, 3][1 + 1]", int64(3)},
	{`["a", 2][0]`, "a"},
	{"not true", false},
	{"not false", true},
	{"true or false", true},
	{"false or false", false},
	{"true and false", false},
	{"false and false", false},
	{"(5 > 2) and (2 <= 3) == true", true},
	{"(1 > 2) or (2 >= 3) == false", true},
}

func TestEvaluator(t *testing.T) {
	for _, test := range evaluatorTests {
		tree := parser.Parse(test.input)
		evaluated, err := Eval(tree, nil)
		require.NoError(t, err, test.input)
		assert.Equal(t, test.expected, evaluated)
	}
}

type evaluatorTestWithEnvironment struct {
	input    string
	env      interface{}
	expected interface{}
}

var evaluatorTestsWithEnvironment = []evaluatorTestWithEnvironment{
	{`foo("world")`,
		map[string]interface{}{"foo": func(input string) string { return "hello " + input }},
		"hello world",
	},
	{`add(1, 2)`,
		map[string]interface{}{"add": func(a, b int64) int64 { return a + b }},
		int64(3),
	},
}

func TestEvaluatorWithEnvironment(t *testing.T) {
	for _, test := range evaluatorTestsWithEnvironment {
		tree := parser.Parse(test.input)
		evaluated, err := Eval(tree, test.env)
		require.NoError(t, err, test.input)
		assert.Equal(t, test.expected, evaluated)
	}
}
