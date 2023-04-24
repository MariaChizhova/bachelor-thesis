package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	OpJumpIfTrue
	OpJumpIfFalse

	OpMinus

	OpArray
	OpIndex

	OpNot

	OpCall
	OpLoadConst
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
	OpJumpIfTrue:     {"OpJumpIfTrue", []int{2}},
	OpJumpIfFalse:    {"OpJumpIfFalse", []int{2}},

	OpMinus: {"OpMinus", []int{}},

	OpArray: {"OpArray", []int{2}},
	OpIndex: {"OpIndex", []int{}},

	OpNot: {"OpNot", []int{}},

	OpCall:      {"OpCall", []int{2}},
	OpLoadConst: {"OpConstant", []int{2}},
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

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func (ins Instructions) String() string {
	var out bytes.Buffer
	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}
		operands, read := ReadOperands(def, ins[i+1:])
		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))
		i += 1 + read
	}
	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n",
			len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	case 2:
		return fmt.Sprintf("%s %d %d", def.Name, operands[0], operands[1])
	}

	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(binary.BigEndian.Uint16(ins[offset:]))
		case 1:
			operands[i] = int(ins[offset:][0])
		}

		offset += width
	}
	return operands, offset
}

func UpdateOperands(ins Instructions, newOperands []int) (Instructions, error) {
	def, err := Lookup(ins[0])
	if err != nil {
		return nil, err
	}

	if len(newOperands) != len(def.OperandWidths) {
		return nil, fmt.Errorf("operand count mismatch: expected %d, got %d", len(def.OperandWidths), len(newOperands))
	}

	newIns := make(Instructions, len(ins))
	copy(newIns, ins)

	offset := 1
	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			binary.BigEndian.PutUint16(newIns[offset:], uint16(newOperands[i]))
		case 1:
			newIns[offset] = byte(newOperands[i])
		default:
			return nil, fmt.Errorf("unsupported operand width: %d", width)
		}

		offset += width
	}

	return newIns, nil
}
