package vm5

import "fmt"

type VM struct {
	Registers    [4]interface{}
	ip           int
	instructions []byte
}

func New(instructions []byte) *VM {
	return &VM{
		instructions: instructions,
	}
}

func (vm *VM) read2Val() int {
	l := int(vm.instructions[vm.ip])
	vm.ip++
	h := int(vm.instructions[vm.ip])
	vm.ip++

	val := l + h*256
	return val
}

func (vm *VM) Run() error {
	vm.ip = 0
	run := true
	for run {
		op := NewOpcode(vm.instructions[vm.ip])
		switch int(op.Value()) {
		case OpExit:
			run = false
		case OpStoreInt:
			vm.ip++
			reg := int(vm.instructions[vm.ip])
			if reg >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", reg)
			}
			vm.ip++
			val := int(vm.instructions[vm.ip]) //vm.read2Val()
			vm.ip++
			vm.Registers[reg] = val
		case OpAdd:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			a := vm.instructions[vm.ip]
			vm.ip++
			b := vm.instructions[vm.ip]
			vm.ip++

			if int(a) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", a)
			}
			if int(b) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", b)
			}
			if int(res) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", res)
			}
			aVal := vm.Registers[a].(int)
			bVal := vm.Registers[b].(int)
			vm.Registers[res] = aVal + bVal
			//fmt.Println("aVal: ", aVal, "bVal: ", bVal, "resVal: ", vm.Registers[res].(int))
		case OpSub:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			a := vm.instructions[vm.ip]
			vm.ip++
			b := vm.instructions[vm.ip]
			vm.ip++

			if int(a) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", a)
			}
			if int(b) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", b)
			}
			if int(res) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", res)
			}
			aVal := vm.Registers[a].(int)
			bVal := vm.Registers[b].(int)
			vm.Registers[res] = aVal - bVal
			//fmt.Println("aVal: ", aVal, "bVal: ", bVal, "resVal: ", vm.Registers[res].(int))
		}
	}
	return nil
}
