package vm4

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type vmTest struct {
	input    []int64
	expected interface{}
}

var vmTests = []vmTest{
	//{ // true
	//	[]int64{
	//		OpBool, R0, 1,
	//		OpPrint, R0,
	//		OpHalt,
	//	}, true}, // TODO: change to bool
	//{ // false
	//	[]int64{
	//		OpBool, R0, 0,
	//		OpPrint, R0,
	//		OpHalt,
	//	}, false}, // TODO: change to bool
	//{ // nil
	//	[]int64{
	//		OpNil, R0,
	//		OpPrint, R0,
	//		OpHalt,
	//	}, false}, // TODO: change to nil
	{ // 1
		[]int64{
			OpConstant, R0, 1,
			OpPrint, R0,
			OpHalt,
		}, int64(1)},
	{ // -1
		[]int64{
			OpConstant, R0, 1,
			OpMinus, R0,
			OpPrint, R0,
			OpHalt,
		}, int64(-1)},
	{ // 1 + 2
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 2,
			OpAdd, R0, R1,
			OpPrint, R0,
			OpHalt,
		}, int64(3)},
	{ // 1 - 2
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 2,
			OpSub, R0, R1,
			OpPrint, R0,
			OpHalt,
		}, int64(-1)},
	{ // 1 * 2
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 2,
			OpMul, R0, R1,
			OpPrint, R0,
			OpHalt,
		}, int64(2)},
	{ // 4 / 2
		[]int64{
			OpConstant, R0, 4,
			OpConstant, R1, 2,
			OpDiv, R0, R1,
			OpPrint, R0,
			OpHalt,
		}, int64(2)},
	{ // 2 ^ 2
		[]int64{
			OpConstant, R0, 2,
			OpConstant, R1, 2,
			OpExp, R0, R1,
			OpPrint, R0,
			OpHalt,
		}, int64(4)},
	{ // 5 % 2
		[]int64{
			OpConstant, R0, 5,
			OpConstant, R1, 2,
			OpMod, R0, R1,
			OpPrint, R0,
			OpHalt,
		}, int64(1)},
	{ // 1 + 2 + (3 - 4)
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 2,
			OpAdd, R0, R1,
			OpConstant, R2, 3,
			OpConstant, R3, 4,
			OpSub, R2, R3,
			OpAdd, R0, R2,
			OpPrint, R0,
			OpHalt,
		}, int64(2)},
	{ // (1 + 2) * 4
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 2,
			OpAdd, R0, R1,
			OpConstant, R2, 4,
			OpMul, R0, R2,
			OpPrint, R0,
			OpHalt}, int64(12),
	},
	{ // 5 * 2 + 10
		[]int64{
			OpConstant, R0, 5,
			OpConstant, R1, 2,
			OpMul, R0, R1,
			OpConstant, R1, 10,
			OpAdd, R0, R1,
			OpPrint, R0,
			OpHalt,
		}, int64(20),
	},
	{ // 5 + 2 * 10
		[]int64{OpConstant, R0, 5,
			OpConstant, R1, 2,
			OpConstant, R2, 10,
			OpMul, R1, R2,
			OpAdd, R0, R1,
			OpPrint, R0,
			OpHalt}, int64(25),
	},
	{ // 10 / 2 * 2 + 10 - 5
		[]int64{
			OpConstant, R0, 10,
			OpConstant, R1, 2,
			OpDiv, R0, R1,
			OpConstant, R1, 2,
			OpMul, R0, R1,
			OpConstant, R1, 10,
			OpAdd, R0, R1,
			OpConstant, R1, 5,
			OpSub, R0, R1,
			OpPrint, R0,
			OpHalt,
		}, int64(15),
	},
	{ // 5 * (2 + 10)
		[]int64{
			OpConstant, R0, 5,
			OpConstant, R1, 2,
			OpConstant, R2, 10,
			OpAdd, R1, R2,
			OpMul, R0, R1,
			OpPrint, R0,
			OpHalt,
		}, int64(60),
	},
	////{"-5 + 10 + -5", int64(0)},
	{ // 2 * 2 * 2 * 2 * 2
		[]int64{
			OpConstant, R0, 2,
			OpConstant, R1, 2,
			OpMul, R0, R1,
			OpConstant, R1, 2,
			OpMul, R0, R1,
			OpConstant, R1, 2,
			OpMul, R0, R1,
			OpConstant, R1, 2,
			OpMul, R0, R1,
			OpPrint, R0,
			OpHalt}, int64(32),
	},
	{ // 1 < 2
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 2,
			OpLessThan, R2, R0, R1,
			OpPrint, R2,
			OpHalt}, int64(1), // TODO: change to bool
	},
	{ // 1 > 1
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 1,
			OpGreaterThan, R2, R0, R1,
			OpPrint, R2,
			OpHalt}, int64(0), // TODO: change to bool
	},
	{ // 2 <= 2
		[]int64{
			OpConstant, R0, 2,
			OpConstant, R1, 2,
			OpLessOrEqual, R2, R0, R1,
			OpPrint, R2,
			OpHalt}, int64(1), // TODO: change to bool
	},
	{ // 1 >= 2
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 2,
			OpGreaterOrEqual, R2, R0, R1,
			OpPrint, R2,
			OpHalt}, int64(0), // TODO: change to bool
	},
	{ // 1 == 2
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 2,
			OpEqual, R2, R0, R1,
			OpPrint, R2,
			OpHalt}, int64(0), // TODO: change to bool
	},
	{ // 1 == 1
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 1,
			OpEqual, R2, R0, R1,
			OpPrint, R2,
			OpHalt}, int64(1), // TODO: change to bool
	},
	{ // 1 != 2
		[]int64{
			OpConstant, R0, 1,
			OpConstant, R1, 2,
			OpNotEqual, R2, R0, R1,
			OpPrint, R2,
			OpHalt}, int64(1), // TODO: change to bool
	},
}

func TestVM(t *testing.T) {
	for _, test := range vmTests {
		vm := New(test.input)
		err := vm.Run()
		require.NoError(t, err, test.input)
		testExpectedObject(t, test.expected, vm.GetResult())
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
