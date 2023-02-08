package code

import (
	"encoding/binary"
)

type Opcode byte

type Instructions []byte

const (
	OpConstant Opcode = iota
	OpPop

	OpTrue
	OpFalse
	OpNil

	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod
	OpExp

	OpEqual
	OpNotEqual
	OpLessThan
	OpGreaterThan
	OpLessOrEqual
	OpGreaterOrEqual

	OpMinus
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{

	OpPop:      {"OpPop", []int{}},
	OpConstant: {"OpConstant", []int{2}},

	OpTrue:  {"OpTrue", []int{}},
	OpFalse: {"OpFalse", []int{}},
	OpNil:   {"OpNil", []int{}},

	OpAdd: {"OpAdd", []int{}},
	OpSub: {"OpSub", []int{}},
	OpMul: {"OpMul", []int{}},
	OpDiv: {"OpDiv", []int{}},
	OpMod: {"OpMod", []int{}},
	OpExp: {"OpExp", []int{}},

	OpEqual:          {"OpEqual", []int{}},
	OpNotEqual:       {"OpNotEqual", []int{}},
	OpLessThan:       {"OpLessThan", []int{}},
	OpGreaterThan:    {"OpGreaterThan", []int{}},
	OpLessOrEqual:    {"OpLessOrEqual", []int{}},
	OpGreaterOrEqual: {"OpGreaterOrEqual", []int{}},

	OpMinus: {"OpMinus", []int{}},
}

func Make(op Opcode, operands ...int) Instructions {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}
	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)
	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}
