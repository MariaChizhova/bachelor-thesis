package evaluator

import (
	"bachelor-thesis/parser"
	"bachelor-thesis/parser/ast"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type evaluatorTest struct {
	input    string
	expected interface{}
}

var evaluatorTests = []evaluatorTest{{"true", true},
	{"false", false},
	{"nil", nil},
	{"1", int64(1)},
	{"-1", int64(-1)},
	{"+1", int64(1)},
	{"+1.1", 1.1},
	{"2.4", 2.4},
	{"-2.4", -2.4},
	{"1 + 2", int64(3)},
	{"1 + 2.1", 3.1},
	{"1 - 2", int64(-1)},
	{"1.1 - 2.2", -1.1},
	{"1 * 2", int64(2)},
	{"1 * 2.1", 2.1},
	{"2.1 * 1", 2.1},
	{"4 / 2", int64(2)},
	{"4.1 / 2", 2.05},
	{"4 / 2.5", 1.6},
	{"2 ^ 2", int64(4)},
	{"2 ^ 2.0", 4.0},
	{"2.1 ^ 2", 4.41},
	{"2.0 ^ 2.0", 4.0},
	{"5 % 2", int64(1)},
	{"1 < 2", true},
	{"1 < 1", false},
	{"2 <= 2", true},
	{"2.1 <= 2", false},
	{"2.1 <= 1.2", false},
	{"2 <= 1.2", false},
	{"1 > 0.2", true},
	{"1.1 > 0", true},
	{"1.1 >= 2.1", false},
	{"1 >= 2", false},
	{"1.1 > 2.2", false},
	{"1.2 >= 0.2", true},
	{"1.1 == 1.1", true},
	{"1 == 2", false},
	{"1 == 1", true},
	{"1 != 2", true},
	{"1.1 != 1.2", true},
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
	{`[1, "a", 3 + 5, "c"][1]`, "a"},
	{"not true", false},
	{"not false", true},
	{"true or false", true},
	{"false or false", false},
	{"true and false", false},
	{"false and false", false},
	{`false or true`, true},
	{"(5 > 2) and (2 <= 3) == true", true},
	{"(1 > 2) or (2 >= 3) == false", true},
	{`"abc" < "bcd"`, true},
	{`"abc" <= "bcd"`, true},
	{`"abcd" > "abc"`, true},
	{`"abc" >= "abcd"`, false},
	{`("rv" == "t") and ("dntxr" > "c") or ("ssjy" == "l") or ("snso" < "uox") and ("qym" < "qyi") and ("tvzew" < "i") or ("bv" <= "xw")`, true},
	{`false and false or true`, true},
	{`("qym" < "qyi") and ("tvzew" < "i") or ("bv" <= "xw")`, true},
	{`("kt" >= "cwcg") and ("pppvp" > "xqqew") or ("geh" <= "wst") and ("je" != "wvvkr") or ("oejgc" < "obsjo") and ("r" != "ml") or ("bkyay" >= "hqdnn")`, true},
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
	{`a + b`,
		map[string]interface{}{"a": 1.2, "b": 2.3},
		3.5,
	},
	{`a and b and not c`,
		map[string]interface{}{"a": false, "b": true, "c": false},
		false,
	},
	{`a + b + add(1, 2)`,
		map[string]interface{}{"a": 1.2, "b": 2.3, "add": func(a, b int64) int64 { return a + b }},
		6.5,
	},
}

func TestEvaluatorWithEnvironment(t *testing.T) {
	for _, test := range evaluatorTestsWithEnvironment {
		tree := parser.Parse(test.input)
		evaluated, err := Eval(tree, test.env)
		if err != nil {
			fmt.Println(ast.Print(tree))
		}
		require.NoError(t, err, test.input)
		assert.Equal(t, test.expected, evaluated)
	}
}
