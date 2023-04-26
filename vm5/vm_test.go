package vm5

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type vmTest struct {
	input    Program
	expected interface{}
}

var vmTests = []vmTest{
	{ // 10 + 20
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpAdd), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{10, 20},
		}, 30,
	},
	{ // 10 + 20 + 5
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpAdd), 03, 01, 02,
				byte(OpStoreInt), 01, 2,
				byte(OpAdd), 03, 01, 03,
				byte(OpExit)},
			Constants: []interface{}{10, 20, 5},
		}, 35,
	},
	{ // 20 - 10
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpSub), 03, 02, 01,
				byte(OpExit)},
			Constants: []interface{}{10, 20},
		}, 10,
	},
	{ // 1 - 2 + 3 - 4
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpSub), 03, 01, 02,
				byte(OpStoreInt), 01, 2,
				byte(OpAdd), 03, 01, 03,
				byte(OpStoreInt), 01, 3,
				byte(OpSub), 03, 03, 01,
				byte(OpExit)},
			Constants: []interface{}{1, 2, 3, 4},
		}, -2,
	},
	{ // "hello"
		Program{
			Instructions: []byte{byte(OpStoreString), 03, 0,
				byte(OpExit)},
			Constants: []interface{}{"hello"},
		}, "hello",
	},
	{Program{
		Instructions: []byte{
			byte(OpStoreBool), 03, 1,
			byte(OpExit),
		}}, true,
	},
	{Program{
		Instructions: []byte{
			byte(OpStoreBool), 03, 0,
			byte(OpExit),
		}}, false,
	},
	{ // "a" == "b"
		Program{
			Instructions: []byte{byte(OpStoreString), 01, 0,
				byte(OpStoreString), 02, 1,
				byte(OpEQ), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{"a", "b"},
		}, false,
	},
	{ // "a" + "b"
		Program{
			Instructions: []byte{byte(OpStoreString), 01, 0,
				byte(OpStoreString), 02, 1,
				byte(OpStringConcat), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{"a", "b"},
		}, "ab",
	},
	{ // foo(1, 2)
		Program{
			Instructions: []byte{
				byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpCall), 03, 2, 2, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{1, 2, "foo"}}, 3,
	},
	{ // bar(1, 2, 3)
		Program{
			Instructions: []byte{
				byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpStoreInt), 03, 2,
				byte(OpCall), 03, 4, 3, 01, 02, 03,
				byte(OpExit)},
			Constants: []interface{}{1, 2, 3, "bar"}}, 6,
	},
}

func TestVM(t *testing.T) {
	for _, test := range vmTests {
		vm := New(test.input)
		err := vm.Run(map[string]interface{}{
			"foo": func(a, b int) int { return a + b },
			"bar": func(a, b, c int) int { return a + b + c }})
		require.NoError(t, err, test.input)
		testExpectedObject(t, test.expected, vm.Registers[3])
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
