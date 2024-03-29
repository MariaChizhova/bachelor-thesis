package vm5

import "fmt"

var (
	OpExit           = 0x00
	OpStoreInt       = 0x01
	OpAdd            = 0x02
	OpSub            = 0x03
	OpMul            = 0x04
	OpDiv            = 0x05
	OpMod            = 0x06
	OpExp            = 0x07
	OpStoreString    = 0x08
	OpStringConcat   = 0x09
	OpStoreBool      = 0x10
	OpEqual          = 0x11
	OpNotEqual       = 0x12
	OpLessThan       = 0x13
	OpGreaterThan    = 0x14
	OpLessOrEqual    = 0x15
	OpGreaterOrEqual = 0x16
	OpCall           = 0x17
	OpJumpIfTrue     = 0x18
	OpJumpIfFalse    = 0x19
	OpLoadConst      = 0x20
	OpStoreFloat     = 0x21
)

type Opcode struct {
	instruction byte
}

func NewOpcode(instruction byte) *Opcode {
	o := &Opcode{}
	o.instruction = instruction
	return o
}

func (o *Opcode) String() string {
	switch int(o.instruction) {
	case OpExit:
		return "OpExit"
	case OpStoreInt:
		return "OpStoreInt"
	case OpAdd:
		return "OpAdd"
	case OpSub:
		return "OpSub"
	case OpMul:
		return "OpMul"
	case OpDiv:
		return "OpDiv"
	case OpMod:
		return "OpMod"
	case OpExp:
		return "OpExp"
	case OpStoreString:
		return "OpStoreString"
	case OpStringConcat:
		return "OpStringConcat"
	case OpStoreBool:
		return "OpStoreBool"
	case OpEqual:
		return "OpEqual"
	case OpNotEqual:
		return "OpNotEqual"
	case OpLessThan:
		return "OpLessThan"
	case OpGreaterThan:
		return "OpGreaterThan"
	case OpLessOrEqual:
		return "OpLessOrEqual"
	case OpGreaterOrEqual:
		return "OpGreaterOrEqual"
	case OpCall:
		return "OpCall"
	case OpJumpIfFalse:
		return "OpJumpIfFalse"
	case OpJumpIfTrue:
		return "OpJumpIfTrue"
	case OpLoadConst:
		return "OpLoadConst"
	case OpStoreFloat:
		return "OpStoreFloat"
	}
	return "unknown opcode .."
}

func (o *Opcode) Value() byte {
	return o.instruction
}

func (p *Program) PrintBytecode() {
	for i := 0; i < len(p.Instructions); {
		opcode := NewOpcode(p.Instructions[i])
		argsCount := opcodeArgsCount(opcode)
		args := p.Instructions[i+1 : i+1+argsCount]

		fmt.Printf("%s", opcode.String())

		for _, arg := range args {
			fmt.Printf(" %v", arg)
		}

		fmt.Println()
		i += 1 + argsCount
	}
}

func opcodeArgsCount(o *Opcode) int {
	switch int(o.instruction) {
	case OpExit:
		return 0
	case OpAdd, OpSub, OpMul, OpDiv, OpMod, OpExp, OpStringConcat,
		OpEqual, OpNotEqual, OpLessThan, OpGreaterThan, OpLessOrEqual,
		OpGreaterOrEqual:
		return 3
	case OpStoreInt, OpStoreString, OpStoreBool, OpJumpIfTrue, OpJumpIfFalse:
		return 2
	default:
		return 0
	}
}
