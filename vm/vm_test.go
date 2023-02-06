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
	{"1", int64(1)},
	{"-1", int64(-1)},
	{"2.4", 2.4},
	{"1 + 2", 3},
	{"1 - 2", -1},
	{"1 * 2", 2},
	{"4 / 2", 2},
	{"2 ^ 2", 4},
	{"5 % 2", 1},
}

// TODO: change the location of this function
func CompileMain(input string) (*compiler.Program, error) {
	tree := parser.Parse(input)
	program, err := compiler.Compile(tree)
	if err != nil {
		return nil, err
	}
	return program, nil
}

func TestVM(t *testing.T) {
	for _, test := range vmTests {
		program, err := CompileMain(test.input)
		require.NoError(t, err, test.input)
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
	case bool:
		assert.Equal(t, expected, actual)
	case int:
		assert.Equal(t, expected, actual)
	case int64:
		assert.Equal(t, expected, actual)
	case float64:
		assert.Equal(t, expected, actual)
	}
}
