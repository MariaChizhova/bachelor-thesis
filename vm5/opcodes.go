package vm5

var (
	OpExit         = 0x00
	OpStoreInt     = 0x01
	OpAdd          = 0x02
	OpSub          = 0x03
	OpStoreString  = 0x04
	OpStringConcat = 0x05
	OpStoreBool    = 0x06
	OpEQ           = 0x07
	OpCall         = 0x08
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
	case OpStoreString:
		return "OpStoreString"
	case OpStringConcat:
		return "OpStringConcat"
	case OpStoreBool:
		return "OpStoreBool"
	case OpEQ:
		return "OpEQ"
	case OpCall:
		return "OpCall"
	}
	return "unknown opcode .."
}

func (o *Opcode) Value() byte {
	return o.instruction
}
