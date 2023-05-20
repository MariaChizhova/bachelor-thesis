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
	{ // 20 * 10
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpMul), 03, 02, 01,
				byte(OpExit)},
			Constants: []interface{}{20, 10},
		}, 200,
	},
	{ // 20 / 10
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpDiv), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{20, 10},
		}, 2,
	},
	{ // 20 % 19
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpMod), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{20, 19},
		}, 1,
	},
	{ // 2 ^ 2
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpExp), 03, 02, 01,
				byte(OpExit)},
			Constants: []interface{}{2, 2},
		}, 4,
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
				byte(OpEqual), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{"a", "b"},
		}, false,
	},
	{ // 1 < 2
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpLessThan), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{1, 2},
		}, true,
	},
	{ // 1 > 2
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpGreaterThan), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{1, 2},
		}, false,
	},
	{ // 1 == 2
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpEqual), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{1, 2},
		}, false,
	},
	{ // 1 != 2
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpNotEqual), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{1, 2},
		}, true,
	},
	{ // true != false
		Program{
			Instructions: []byte{byte(OpStoreInt), 01, 0,
				byte(OpStoreInt), 02, 1,
				byte(OpNotEqual), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{1, 2},
		}, true,
	},
	{ // "a" <= "ab"
		Program{
			Instructions: []byte{byte(OpStoreString), 01, 0,
				byte(OpStoreString), 02, 1,
				byte(OpLessOrEqual), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{"a", "ab"},
		}, true,
	},
	{ // "a" >= "ab"
		Program{
			Instructions: []byte{byte(OpStoreString), 01, 0,
				byte(OpStoreString), 02, 1,
				byte(OpGreaterOrEqual), 03, 01, 02,
				byte(OpExit)},
			Constants: []interface{}{"a", "ab"},
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
				byte(OpCall), 03, 3, 3, 01, 02, 03,
				byte(OpExit)},
			Constants: []interface{}{1, 2, 3, "bar"}}, 6,
	},
	{ // true or false
		Program{
			Instructions: []byte{
				byte(OpStoreBool), 01, 1,
				byte(OpStoreBool), 03, 1,
				byte(OpJumpIfTrue), 01, 4,
				byte(OpStoreBool), 03, 0,
				byte(OpExit)}}, true,
	},
	{ // false or false
		Program{
			Instructions: []byte{
				byte(OpStoreBool), 01, 0,
				byte(OpStoreBool), 03, 1,
				byte(OpJumpIfTrue), 01, 4,
				byte(OpStoreBool), 03, 0,
				byte(OpExit)}}, false,
	},
	{ // false and true
		Program{
			Instructions: []byte{
				byte(OpStoreBool), 01, 0,
				byte(OpStoreBool), 03, 0,
				byte(OpJumpIfFalse), 01, 4,
				byte(OpStoreBool), 03, 1,
				byte(OpExit)}}, false,
	},
	{ // true and true
		Program{
			Instructions: []byte{
				byte(OpStoreBool), 01, 1,
				byte(OpStoreBool), 03, 0,
				byte(OpJumpIfFalse), 01, 4,
				byte(OpStoreBool), 03, 1,
				byte(OpExit)}}, true,
	},
	{ // true and false or true
		Program{
			Instructions: []byte{
				byte(OpStoreBool), 01, 1,
				byte(OpJumpIfFalse), 01, 4,
				byte(OpStoreBool), 03, 1,
				byte(OpJumpIfTrue), 03, 4,
				byte(OpStoreBool), 03, 0,
				byte(OpExit),
			}}, true,
	},
	{ // ("a" == "a") and ("a" == "b")
		Program{
			Instructions: []byte{
				byte(OpStoreString), 01, 0, // ("a" == "a")
				byte(OpStoreString), 02, 1,
				byte(OpEqual), 03, 01, 02,
				byte(OpJumpIfFalse), 03, 11,
				byte(OpStoreString), 01, 2, // ("a" == "b")
				byte(OpStoreString), 02, 3,
				byte(OpEqual), 03, 01, 02,
				byte(OpExit),
			},
			Constants: []interface{}{"a", "a", "a", "b"}}, false,
	},
	{ // ("a" != "a") and ("a" != "b")
		Program{
			Instructions: []byte{
				byte(OpStoreString), 01, 0, // ("a" != "a")
				byte(OpStoreString), 02, 1,
				byte(OpEqual), 03, 01, 02,
				byte(OpJumpIfFalse), 03, 11,
				byte(OpStoreString), 01, 2, // ("a" != "b")
				byte(OpStoreString), 02, 3,
				byte(OpEqual), 03, 01, 02,
				byte(OpExit),
			},
			Constants: []interface{}{"a", "a", "a", "b"}}, false,
	},
	{ // ("a" != "a") or ("a" != "b")
		Program{
			Instructions: []byte{
				byte(OpStoreString), 01, 0, // ("a" != "a")
				byte(OpStoreString), 02, 1,
				byte(OpEqual), 03, 01, 02,
				byte(OpJumpIfTrue), 03, 11,
				byte(OpStoreString), 01, 2, // ("a" != "b")
				byte(OpStoreString), 02, 3,
				byte(OpEqual), 03, 01, 02,
				byte(OpExit),
			},
			Constants: []interface{}{"a", "a", "a", "b"}}, true,
	},
	{ // ("a" == "a") or ("a" != "b")
		Program{
			Instructions: []byte{
				byte(OpStoreString), 01, 0, // ("a" == "a")
				byte(OpStoreString), 02, 1,
				byte(OpEqual), 03, 01, 02,
				byte(OpJumpIfTrue), 03, 11,
				byte(OpStoreString), 01, 2, // ("a" != "b")
				byte(OpStoreString), 02, 3,
				byte(OpEqual), 03, 01, 02,
				byte(OpExit),
			},
			Constants: []interface{}{"a", "a", "a", "b"}}, true,
	},
	{ // ("bhhl" > "yp") and ("yvi" < "bwpj") and ("w" > "fa")
		Program{
			Instructions: []byte{
				byte(OpStoreString), 01, 0,
				byte(OpStoreString), 02, 1,
				byte(OpGreaterThan), 03, 01, 02,
				byte(OpJumpIfFalse), 03, 11,
				byte(OpStoreString), 01, 2,
				byte(OpStoreString), 02, 3,
				byte(OpLessThan), 03, 01, 02,
				byte(OpJumpIfFalse), 03, 11,
				byte(OpStoreString), 01, 4,
				byte(OpStoreString), 02, 5,
				byte(OpGreaterThan), 03, 01, 02,
				byte(OpExit),
			},
			Constants: []interface{}{"bhhl", "yp", "yvi", "bwpj", "w", "fa"}}, false,
	},
	{ // ("u" != "fmlpg") or ("wosxx" < "o") or ("m" <= "cwrq")
		Program{
			Instructions: []byte{
				byte(OpStoreString), 01, 0,
				byte(OpStoreString), 02, 1,
				byte(OpNotEqual), 03, 01, 02,
				byte(OpJumpIfTrue), 03, 11,
				byte(OpStoreString), 01, 2,
				byte(OpStoreString), 02, 3,
				byte(OpLessThan), 03, 01, 02,
				byte(OpJumpIfTrue), 03, 11,
				byte(OpStoreString), 01, 4,
				byte(OpStoreString), 02, 5,
				byte(OpLessOrEqual), 03, 01, 02,
				byte(OpExit),
			},
			Constants: []interface{}{"u", "fmlpg", "wosxx", "o", "m", "cwrq"}}, true,
	},
	{ // ("qg" >= "mbqo") or ("ymehe" >= "lh") or ("gnr" == "d")
		Program{
			Instructions: []byte{
				byte(OpStoreString), 01, 0,
				byte(OpStoreString), 02, 1,
				byte(OpGreaterOrEqual), 03, 01, 02,
				byte(OpJumpIfTrue), 03, 11,
				byte(OpStoreString), 01, 2,
				byte(OpStoreString), 02, 3,
				byte(OpGreaterOrEqual), 03, 01, 02,
				byte(OpJumpIfTrue), 03, 11,
				byte(OpStoreString), 01, 4,
				byte(OpStoreString), 02, 5,
				byte(OpEqual), 03, 01, 02,
				byte(OpExit),
			},
			Constants: []interface{}{"qg", "mbqo", "ymehe", "lh", "gnr", "d"}}, true,
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
