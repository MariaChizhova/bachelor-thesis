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
		`nil`,
		Program{
			Instructions: concatInstructions([]code.Instructions{code.Make(code.OpNil)}),
		},
	},
	{
		`true + false`,
		Program{
			Instructions: concatInstructions([]code.Instructions{code.Make(code.OpTrue),
				code.Make(code.OpFalse),
				code.Make(code.OpAdd),
				/*code.Make(code.OpPop)*/}),
		},
	},
	{
		`true < false`,
		Program{
			Instructions: concatInstructions([]code.Instructions{code.Make(code.OpTrue),
				code.Make(code.OpFalse),
				code.Make(code.OpLessThan),
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
		`1 - 2`,
		Program{
			Constants: []interface{}{
				int64(1), int64(2),
			},
			Instructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpSub),
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
}

// TODO: change the location of this function
func CompileMain(input string) (*Program, error) {
	tree := parser.Parse(input)
	program, err := Compile(tree)
	if err != nil {
		return nil, err
	}
	return program, nil
}

func TestCompiler(t *testing.T) {
	for _, test := range compilerTests {
		program, err := CompileMain(test.input)
		require.NoError(t, err, test.input)
		assert.Equal(t, test.program.Instructions, program.Instructions)
		assert.Equal(t, test.program.Constants, program.Constants)
	}
}
