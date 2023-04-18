package vm3

import (
	"bachelor-thesis/parser"
	"bachelor-thesis/vm/compiler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type vmTest struct {
	input    string
	expected interface{}
}

var vmTests = []vmTest{
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
	{`"hello"`, "hello"}, {`"hello " + "world!"`, "hello world!"},
	{"[]", []interface{}{}},
	{"[1, 2, 3.1]", []interface{}{int64(1), int64(2), 3.1}},
	{"[1 + 2, 2 * 3]", []interface{}{int64(3), int64(6)}},
	{"[1, 2, 3][1]", int64(2)},
	{"[1, 2, 3][1 + 1]", int64(3)},
	{`["a", 2][0]`, "a"},
	{`[2, "a"][1]`, "a"},
	{`[1, "a", 3 + 5, "c"][0]`, int64(1)},
	{`[1, "a", 3 + 5, "c"][1]`, "a"},
	{`[1, "a", 3 + 5, "c"][2]`, int64(8)},
	{`[1, "a", 3 + 5, "c"][3]`, "c"},
	{"not true", false},
	{"not false", true},
	{"true or false", true},
	{"false or false", false},
	{"true and false", false},
	{"false and false", false},
	{"(5 > 2) and (2 <= 3) == true", true},
	{"(1 > 2) or (2 >= 3) == false", true},
}

func TestVM(t *testing.T) {
	for _, test := range vmTests {
		tree := parser.Parse(test.input)
		program, err := compiler.Compile(tree)
		//print(program.Instructions.String())
		vm := New(program.Instructions, program.Constants)
		err = vm.Run(nil)
		require.NoError(t, err, test.input)
		stackElem := vm.StackTop()
		testExpectedObject(t, test.expected, stackElem)
	}
}

func testExpectedObject(
	t *testing.T,
	expected interface{},
	actual interface{},
) {
	switch expected := expected.(type) {
	case bool, int, int64, float64, nil, string, []interface{}:
		assert.Equal(t, expected, actual)
	}
}

type vmTestWithEnvironment struct {
	input    string
	env      interface{}
	expected interface{}
}

var vmTestsWithEnvironment = []vmTestWithEnvironment{
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

func TestVMWithEnvironment(t *testing.T) {
	for _, test := range vmTestsWithEnvironment {
		tree := parser.Parse(test.input)
		program, err := compiler.Compile(tree)
		//print(program.Instructions.String())
		vm := New(program.Instructions, program.Constants)
		err = vm.Run(test.env)
		require.NoError(t, err, test.input)
		stackElem := vm.StackTop()
		testExpectedObject(t, test.expected, stackElem)
	}
}
