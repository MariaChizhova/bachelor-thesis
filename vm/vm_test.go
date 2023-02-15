package vm

import (
	"bachelor-thesis/compiler"
	"bachelor-thesis/parser"
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
	{"1 + 2", 3},
	{"1 - 2", -1},
	{"1 * 2", 2},
	{"4 / 2", 2},
	{"2 ^ 2", 4},
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
	{"5 * 2 + 10", 20},
	{"5 + 2 * 10", 25},
	{"10 / 2 * 2 + 10 - 5", 15},
	{"5 * (2 + 10)", int64(60)},
	{"-5 + 10 + -5", 0},
	{"2 * 2 * 2 * 2 * 2", int64(32)},
	{"1 - 2 + 3 * 4 / 2 ^ 2 % 3", -1},
	{"true == true", true},
	{"true == false", false},
	{"true != false", true},
	{"(1 < 2) == true", true},
	{"(1 < 2) == false", false},
	{`"hello"`, "hello"},
	{`"hello " + "world!"`, "hello world!"},
	{"[]", []interface{}{}},
	{"[1, 2, 3.1]", []interface{}{int64(1), int64(2), 3.1}},
	{"[1 + 2, 2 * 3]", []interface{}{3, 6}},
	// TODO: implement more tests
}

func TestVM(t *testing.T) {
	for _, test := range vmTests {
		tree := parser.Parse(test.input)
		program, err := compiler.Compile(tree)
		// print(program.Instructions.String())
		vm := New(program.Instructions, program.Constants)
		err = vm.Run()
		require.NoError(t, err, test.input)
		stackElem := vm.StackTop() // vm.LastPoppedStackElem()
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
