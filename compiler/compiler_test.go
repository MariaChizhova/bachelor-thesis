package compiler

import (
	"bachelor-thesis/code"
	"bachelor-thesis/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type compilerTest struct {
	input   string
	program Program
}

// TODO: implement more tests
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
				code.Make(code.OpConstant),
				/*code.Make(code.OpPop)*/}),
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
				code.Make(code.OpMinus),
				/*code.Make(code.OpPop)*/}),
		},
	},
	{
		`1.2`,
		Program{
			Constants: []interface{}{
				1.2,
			},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant),
				/*code.Make(code.OpPop)*/}),
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
				code.Make(code.OpEqual),
				/*code.Make(code.OpPop)*/}),
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
				code.Make(code.OpLessThan),
				/*code.Make(code.OpPop)*/}),
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
				code.Make(code.OpAdd),
				/*code.Make(code.OpPop)*/}),
		},
	},
}

func TestCompiler(t *testing.T) {
	for _, test := range compilerTests {
		tree := parser.Parse(test.input)
		program, err := Compile(tree)
		require.NoError(t, err, test.input)
		// print(program.Instructions.String())
		assert.Equal(t, test.program.Instructions, program.Instructions)
		assert.Equal(t, test.program.Constants, program.Constants)
	}
}
