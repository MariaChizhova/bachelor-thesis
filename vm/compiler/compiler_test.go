package compiler

import (
	"bachelor-thesis/parser"
	"bachelor-thesis/vm/code"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type compilerTest struct {
	input   string
	program Program
}

var compilerTests = []compilerTest{
	{
		`true`,
		Program{
			Instructions: concatInstructions([]code.Instructions{code.Make(code.OpTrue)}),
		},
	},
	{
		`false`,
		Program{
			Instructions: concatInstructions([]code.Instructions{code.Make(code.OpFalse)}),
		},
	},
	{
		`nil`,
		Program{
			Instructions: concatInstructions([]code.Instructions{code.Make(code.OpNil)}),
		},
	},
	{
		`10`,
		Program{
			Constants: []interface{}{
				int64(10),
			},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant)}),
		},
	},
	{
		`-1`,
		Program{
			Constants: []interface{}{
				int64(1),
			},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant),
				code.Make(code.OpMinus)}),
		},
	},
	{
		`1.2`,
		Program{
			Constants: []interface{}{
				1.2,
			},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant)}),
		},
	},
	{
		`"text"`,
		Program{
			Constants: []interface{}{
				"text",
			},
			Instructions: code.Make(code.OpConstant),
		},
	},
	{
		`1 == 3`,
		Program{
			Constants: []interface{}{
				int64(1), int64(3),
			},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpEqual)}),
		},
	},
	{
		`1 < 2`,
		Program{
			Constants: []interface{}{
				int64(1), int64(2),
			},
			Instructions: concatInstructions([]code.Instructions{code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpLessThan)}),
		},
	},
	{
		`1 + 2`,
		Program{
			Constants: []interface{}{
				int64(1), int64(2),
			},
			Instructions: concatInstructions([]code.Instructions{code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpAdd)}),
		},
	},
	{
		`1 - 2 + 3 * 4 / 2 ^ 2 % 3`,
		Program{
			Constants: []interface{}{
				int64(1), int64(2), int64(3), int64(4), int64(2), int64(2), int64(3),
			},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpSub),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpConstant, 3),
				code.Make(code.OpMul),
				code.Make(code.OpConstant, 4),
				code.Make(code.OpDiv),
				code.Make(code.OpConstant, 5),
				code.Make(code.OpExp),
				code.Make(code.OpConstant, 6),
				code.Make(code.OpMod),
				code.Make(code.OpAdd)}),
		},
	},
	{
		`"hello " + "world!"`,
		Program{
			Constants: []interface{}{
				"hello ", "world!",
			},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpAdd)}),
		},
	},
	{
		"[]",
		Program{
			// Constants: []interface{}{},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpArray, 0)}),
		},
	},
	{
		"[1, 2]",
		Program{
			Constants: []interface{}{int64(1), int64(2)},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpArray, 2)}),
		},
	},
	{
		`[1 + 2, 3 * 4]`,
		Program{
			Constants: []interface{}{int64(1), int64(2), int64(3), int64(4)},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpAdd),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpConstant, 3),
				code.Make(code.OpMul),
				code.Make(code.OpArray, 2)}),
		},
	},
	{
		`not true`,
		Program{
			Instructions: concatInstructions([]code.Instructions{code.Make(code.OpTrue), code.Make(code.OpNot)}),
		},
	},
	{
		`true or false`,
		Program{
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpTrue),
				code.Make(code.OpJumpIfTrue, 4),
				code.Make(code.OpPop),
				code.Make(code.OpFalse)}),
		},
	},
	{
		`false and false`,
		Program{
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpFalse),
				code.Make(code.OpJumpIfFalse, 4),
				code.Make(code.OpPop),
				code.Make(code.OpFalse)}),
		},
	},
	{
		"[1, 2, 3][1 + 1]",
		Program{
			Constants: []interface{}{int64(1), int64(2), int64(3), int64(1), int64(1)},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpArray, 3),
				code.Make(code.OpConstant, 3),
				code.Make(code.OpConstant, 4),
				code.Make(code.OpAdd),
				code.Make(code.OpIndex)}),
		},
	},
	{
		`["a", "b"][1]`,
		Program{
			Constants: []interface{}{"a", "b", int64(1)},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpArray, 2),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpIndex)}),
		},
	},
	{
		`foo()`,
		Program{
			Constants: []interface{}{"foo"},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpLoadConst, 0),
				code.Make(code.OpCall, 0),
			}),
		},
	},
	{
		`foo(bar())`,
		Program{
			Constants: []interface{}{"bar", "foo"},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpLoadConst, 0),
				code.Make(code.OpCall, 0),
				code.Make(code.OpLoadConst, 1),
				code.Make(code.OpCall, 1),
			}),
		},
	},
	{
		`foo("arg1", 2, true)`,
		Program{
			Constants: []interface{}{"arg1", int64(2), "foo"},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpTrue),
				code.Make(code.OpLoadConst, 2),
				code.Make(code.OpCall, 3),
			}),
		},
	},
}

func TestCompiler(t *testing.T) {
	for _, test := range compilerTests {
		tree := parser.Parse(test.input)
		program, err := Compile(tree)
		// print(program.Instructions.String())
		require.NoError(t, err, test.input)
		assert.Equal(t, test.program.Instructions, program.Instructions)
		assert.Equal(t, test.program.Constants, program.Constants)
	}
}
