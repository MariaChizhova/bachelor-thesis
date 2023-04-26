package vm4

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
	{"1", int64(1)},
	{"-1", int64(-1)},
	{"2.4", 2.4},
	{"-2.4", -2.4},
	{"1 + 2", int64(3)},
	{"1 - 2", int64(-1)},
	{"-5 + 10 + -5", int64(0)},
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
	case bool, int64, float64, nil, string, []interface{}:
		assert.Equal(t, expected, actual)
	}
}
