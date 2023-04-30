package vm5

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
		return "JumpIfFalse"
	case OpJumpIfTrue:
		return "JumpIfTrue"
	}
	return "unknown opcode .."
}

func (o *Opcode) Value() byte {
	return o.instruction
}
