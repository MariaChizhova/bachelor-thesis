package vm4

import (
	"fmt"
	"math"
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
	registers    []int64
	ip           int
	instructions []int64
}

func New(instructions []int64) *VM {
	return &VM{
		registers:    make([]int64, 4),
		ip:           -1,
		instructions: instructions,
	}
}

func (vm *VM) NextCode() int64 {
	vm.ip++
	return vm.instructions[vm.ip]
}

func (vm *VM) GetResult() interface{} {
	//return math.Float64frombits(vm.registers[vm.instructions[vm.ip-1]])
	return vm.registers[vm.instructions[vm.ip-1]]
}

func (vm *VM) PrintRuntimeInfo() {
	fmt.Printf("IP: %d \n", vm.ip)
	for k, v := range vm.registers {
		fmt.Printf("reg: %d, value: %d\n", k, v)
	}
}

func (vm *VM) Run() error {
	for {
		instr := vm.NextCode()
		switch instr {
		case OpConstant:
			reg := vm.NextCode()
			value := vm.NextCode()
			vm.registers[uint(reg)] = value
		case OpBool:
			reg := vm.NextCode()
			value := vm.NextCode()
			vm.registers[uint(reg)] = value
		case OpNil:
			reg := vm.NextCode()
			vm.registers[uint(reg)] = -1
		case OpAdd:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			vm.registers[reg1] += vm.registers[reg2]
		case OpSub:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			vm.registers[reg1] -= vm.registers[reg2]
		case OpMul:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			vm.registers[reg1] *= vm.registers[reg2]
		case OpDiv:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			vm.registers[reg1] /= vm.registers[reg2]
		case OpMod:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			vm.registers[reg1] %= vm.registers[reg2]
		case OpExp:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			vm.registers[reg1] = int64(math.Pow(float64(vm.registers[reg1]), float64(vm.registers[reg2])))
		case OpMinus:
			reg := uint(vm.NextCode())
			vm.registers[reg] = -vm.registers[reg]
		case OpEqual:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			reg3 := uint(vm.NextCode())
			if vm.registers[reg2] == vm.registers[reg3] {
				vm.registers[reg1] = 1
			} else {
				vm.registers[reg1] = 0
			}
		case OpNotEqual:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			reg3 := uint(vm.NextCode())
			if vm.registers[reg2] != vm.registers[reg3] {
				vm.registers[reg1] = 1
			} else {
				vm.registers[reg1] = 0
			}
		case OpLessThan:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			reg3 := uint(vm.NextCode())
			if vm.registers[reg2] < vm.registers[reg3] {
				vm.registers[reg1] = 1
			} else {
				vm.registers[reg1] = 0
			}
		case OpLessOrEqual:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			reg3 := uint(vm.NextCode())
			if vm.registers[reg2] <= vm.registers[reg3] {
				vm.registers[reg1] = 1
			} else {
				vm.registers[reg1] = 0
			}
		case OpGreaterThan:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			reg3 := uint(vm.NextCode())
			if vm.registers[reg2] > vm.registers[reg3] {
				vm.registers[reg1] = 1
			} else {
				vm.registers[reg1] = 0
			}
		case OpGreaterOrEqual:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			reg3 := uint(vm.NextCode())
			if vm.registers[reg2] >= vm.registers[reg3] {
				vm.registers[reg1] = 1
			} else {
				vm.registers[reg1] = 0
			}
		case OpPrint:
			reg := uint(vm.NextCode())
			fmt.Println("Result: ", vm.registers[reg])
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
