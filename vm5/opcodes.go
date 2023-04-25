package vm5

var (
	OpExit     = 0x00
	OpStoreInt = 0x01
	OpAdd      = 0x02
	OpSub      = 0x03
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
	}
	return "unknown opcode .."
}

func (o *Opcode) Value() byte {
	return o.instruction
}
