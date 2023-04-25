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
	//{
	//	[]byte{byte(OpStoreInt), 1, 10, 0,
	//		byte(OpStoreInt), 2, 20, 0,
	//		byte(OpAdd), 3, 1, 2,
	//		byte(OpExit)}, 30,
	//},
	//{
	//	[]byte{byte(OpStoreInt), 1, 10, // 0,
	//		byte(OpStoreInt), 2, 20, //0,
	//		byte(OpAdd), 3, 1, 2,
	//		byte(OpStoreInt), 1, 5, //0,
	//		byte(OpAdd), 3, 1, 3,
	//		byte(OpExit)}, 35,
	//},
	//{
	//	[]byte{byte(OpStoreInt), 1, byte(1), // 0,
	//		byte(OpStoreInt), 2, byte(2), //, 0,
	//		byte(OpAdd), 3, 1, 2,
	//		byte(OpStoreInt), 1, byte(3), // 0,
	//		byte(OpAdd), 3, 1, 3,
	//		byte(OpStoreInt), 1, byte(4), // 0,
	//		byte(OpAdd), 3, 1, 3,
	//		byte(OpExit)}, 10,
	//},
	//{
	//	[]byte{byte(OpStoreInt), 1, 10, // 0,
	//		byte(OpStoreInt), 2, 20, //0,
	//		byte(OpSub), 3, 2, 1,
	//		byte(OpExit)}, 10,
	//},
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
