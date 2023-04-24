package vm5

import (
	"fmt"
)

const (
	R0 = iota
	R1
	R2
	R3
)

const (
	OpConstant = iota
	OpBool
	OpNil
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod
	OpExp
	OpMinus
	OpEqual
	OpNotEqual
	OpLessThan
	OpGreaterThan
	OpLessOrEqual
	OpGreaterOrEqual
	OpPrint
	OpHalt
)

type VM struct {
	registers    []interface{}
	ip           int
	instructions []interface{}
}

func New(instructions []interface{}) *VM {
	return &VM{
		registers:    make([]interface{}, 4),
		ip:           -1,
		instructions: instructions,
	}
}

func (vm *VM) NextCode() interface{} {
	vm.ip++
	return vm.instructions[vm.ip]
}

func (vm *VM) GetResult() interface{} {
	return vm.registers[vm.instructions[vm.ip-1].(int)]
}

func (vm *VM) PrintRuntimeInfo() {
	fmt.Printf("IP: %d \n", vm.ip)
	for k, v := range vm.registers {
		fmt.Printf("reg: %d, value: %d\n", k, v)
	}
}

func (vm *VM) Run() error {
	vm.ip = -1
	for {
		instr := vm.NextCode()
		switch instr {
		case OpConstant:
			reg := vm.NextCode()
			value := vm.NextCode()
			vm.registers[reg.(int)] = value
		case OpBool:
			reg := vm.NextCode()
			value := vm.NextCode()
			vm.registers[reg.(int)] = value
		case OpNil:
			reg := vm.NextCode()
			vm.registers[reg.(int)] = nil
		case OpAdd:
			reg1 := vm.NextCode().(int)
			reg2 := vm.NextCode().(int)
			vm.registers[reg1] = vm.executeAddOperation(vm.registers[reg1], vm.registers[reg2])
		case OpSub:
			reg1 := vm.NextCode().(int)
			reg2 := vm.NextCode().(int)
			vm.registers[reg1] = vm.executeSubtractOperation(vm.registers[reg1], vm.registers[reg2])
		case OpMul:
			reg1 := vm.NextCode().(int)
			reg2 := vm.NextCode().(int)
			vm.registers[reg1] = vm.executeMultiplyOperation(vm.registers[reg1], vm.registers[reg2])
		case OpDiv:
			reg1 := vm.NextCode().(int)
			reg2 := vm.NextCode().(int)
			vm.registers[reg1] = vm.executeDivideOperation(vm.registers[reg1], vm.registers[reg2])
		case OpMod:
			reg1 := vm.NextCode().(int)
			reg2 := vm.NextCode().(int)
			vm.registers[reg1] = vm.executeRemainderOperation(vm.registers[reg1], vm.registers[reg2])
		case OpExp:
			reg1 := vm.NextCode().(int)
			reg2 := vm.NextCode().(int)
			vm.registers[reg1] = vm.executeExponentiationOperation(vm.registers[reg1], vm.registers[reg2])
		case OpMinus:
			reg := vm.NextCode()
			vm.registers[reg.(int)] = vm.executeMinusOperator(vm.registers[reg.(int)])
		//case OpEqual:
		//	reg1 := uint(vm.NextCode())
		//	reg2 := uint(vm.NextCode())
		//	reg3 := uint(vm.NextCode())
		//	if vm.registers[reg2] == vm.registers[reg3] {
		//		vm.registers[reg1] = 1
		//	} else {
		//		vm.registers[reg1] = 0
		//	}
		//case OpNotEqual:
		//	reg1 := uint(vm.NextCode())
		//	reg2 := uint(vm.NextCode())
		//	reg3 := uint(vm.NextCode())
		//	if vm.registers[reg2] != vm.registers[reg3] {
		//		vm.registers[reg1] = 1
		//	} else {
		//		vm.registers[reg1] = 0
		//	}
		//case OpLessThan:
		//	reg1 := uint(vm.NextCode())
		//	reg2 := uint(vm.NextCode())
		//	reg3 := uint(vm.NextCode())
		//	if vm.registers[reg2] < vm.registers[reg3] {
		//		vm.registers[reg1] = 1
		//	} else {
		//		vm.registers[reg1] = 0
		//	}
		//case OpLessOrEqual:
		//	reg1 := uint(vm.NextCode())
		//	reg2 := uint(vm.NextCode())
		//	reg3 := uint(vm.NextCode())
		//	if vm.registers[reg2] <= vm.registers[reg3] {
		//		vm.registers[reg1] = 1
		//	} else {
		//		vm.registers[reg1] = 0
		//	}
		//case OpGreaterThan:
		//	reg1 := uint(vm.NextCode())
		//	reg2 := uint(vm.NextCode())
		//	reg3 := uint(vm.NextCode())
		//	if vm.registers[reg2] > vm.registers[reg3] {
		//		vm.registers[reg1] = 1
		//	} else {
		//		vm.registers[reg1] = 0
		//	}
		//case OpGreaterOrEqual:
		//	reg1 := uint(vm.NextCode())
		//	reg2 := uint(vm.NextCode())
		//	reg3 := uint(vm.NextCode())
		//	if vm.registers[reg2] >= vm.registers[reg3] {
		//		vm.registers[reg1] = 1
		//	} else {
		//		vm.registers[reg1] = 0
		//	}
		case OpPrint:
			vm.NextCode()
			//reg := uint(vm.NextCode())
			//fmt.Println("Result: ", vm.registers[reg])
		case OpHalt:
			return nil
		default:
			if vm.ip >= len(vm.instructions) {
				return fmt.Errorf("read all programs")
			}
			return fmt.Errorf("unsupported opcode: %d", vm.instructions[vm.ip])
		}
	}
	return nil
}
