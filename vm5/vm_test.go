package vm5

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type vmTest struct {
	input    []byte
	expected interface{}
}

var vmTests = []vmTest{
	{
		// 10 + 20
		[]byte{byte(OpStoreInt), 1, 10, // 0,
			byte(OpStoreInt), 2, 20, // 0,
			byte(OpAdd), 3, 1, 2,
			byte(OpExit)}, 30,
	},
	{ // 10 + 20 + 5
		[]byte{byte(OpStoreInt), 1, 10, // 0,
			byte(OpStoreInt), 2, 20, //0,
			byte(OpAdd), 3, 1, 2,
			byte(OpStoreInt), 1, 5, //0,
			byte(OpAdd), 3, 1, 3,
			byte(OpExit)}, 35,
	},
	{ // 1 + 2 + 3 + 4
		[]byte{byte(OpStoreInt), 1, byte(1), // 0,
			byte(OpStoreInt), 2, byte(2), //, 0,
			byte(OpAdd), 3, 1, 2,
			byte(OpStoreInt), 1, byte(3), // 0,
			byte(OpAdd), 3, 1, 3,
			byte(OpStoreInt), 1, byte(4), // 0,
			byte(OpAdd), 3, 1, 3,
			byte(OpExit)}, 10,
	},
	{ // 20 - 10
		[]byte{byte(OpStoreInt), 1, 10, // 0,
			byte(OpStoreInt), 2, 20, //0,
			byte(OpSub), 3, 2, 1,
			byte(OpExit)}, 10,
	},
	{ // 1 - 2 + 3 - 4
		[]byte{byte(OpStoreInt), 1, byte(1), // 0,
			byte(OpStoreInt), 2, byte(2), //, 0,
			byte(OpSub), 3, 1, 2,
			byte(OpStoreInt), 1, byte(3), // 0,
			byte(OpAdd), 3, 1, 3,
			byte(OpStoreInt), 1, byte(4), // 0,
			byte(OpSub), 3, 3, 1,
			byte(OpExit)}, -2,
	},
	{
		[]byte{
			byte(OpStoreString), 03,
			05,        //00, // length of the string
			byte('h'), // "hello"
			byte('e'),
			byte('l'),
			byte('l'),
			byte('o'),
			byte(OpExit),
		}, "hello",
	},
	{
		[]byte{
			byte(OpStoreBool), 03, 1,
			byte(OpExit),
		}, true,
	},
	{
		[]byte{
			byte(OpStoreBool), 03, 0,
			byte(OpExit),
		}, false,
	},
	{
		[]byte{
			byte(OpStoreString), 01, 01, byte('a'),
			byte(OpStoreString), 02, 01, byte('b'),
			byte(OpEQ), 03, 01, 02,
			byte(OpExit),
		}, false,
	},
}

func TestVM(t *testing.T) {
	for _, test := range vmTests {
		vm := New(test.input)
		err := vm.Run()
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
