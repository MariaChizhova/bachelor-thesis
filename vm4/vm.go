package vm4

import (
	"bachelor-thesis/vm/code"
	"encoding/binary"
	"fmt"
)

type VM struct {
	constants    []interface{}
	instructions code.Instructions
	stack        []interface{}
	stackInt     []int64
	sp           int
}

func New(instructions code.Instructions, constants []interface{}) *VM {
	return &VM{
		instructions: instructions,
		constants:    constants,
		stack:        make([]interface{}, 0),
		stackInt:     make([]int64, 0),
		sp:           0,
	}
}

func (vm *VM) StackTop() interface{} {
	if vm.stack[len(vm.stack)-1] != nil {
		return vm.stack[len(vm.stack)-1]
	} else {
		return vm.stackInt[len(vm.stackInt)-1]
	}
}

func (vm *VM) Run(env interface{}) error {
	if vm.stack == nil {
		vm.stack = make([]interface{}, 0, 2)
	} else {
		vm.stack = vm.stack[0:0]
	}
	if vm.stackInt == nil {
		vm.stackInt = make([]int64, 0, 2)
	} else {
		vm.stackInt = vm.stackInt[0:0]
	}
	vm.sp = 0
	for vm.sp < len(vm.instructions) {
		switch code.Opcode(vm.instructions[vm.sp]) {
		case code.OpConstant:
			constIndex := binary.BigEndian.Uint16(vm.instructions[vm.sp+1:])
			vm.sp += 2
			if int(constIndex) >= len(vm.constants) {
				return fmt.Errorf("constant index out of range: %d", constIndex)
			}
			vm.push(vm.constants[constIndex])
		case code.OpPop:
			vm.pop()
		case code.OpTrue:
			vm.push(true)
		case code.OpFalse:
			vm.push(false)
		case code.OpNil:
			vm.push(nil)
		case code.OpAdd:
			a, ai := vm.pop()
			b, bi := vm.pop()
			if a == nil && b == nil {
				vm.push(bi + ai)
			} else {
				vm.push(vm.executeAddOperation(b, a))
			}
		case code.OpSub:
			a, ai := vm.pop()
			b, bi := vm.pop()
			if a == nil && b == nil {
				vm.push(bi - ai)
			} else {
				vm.push(vm.executeAddOperation(b, a))
			}
		case code.OpMinus:
			vm.push(vm.executeMinusOperator())
		default:
			return fmt.Errorf("unsupported opcode: %d", code.Opcode(vm.instructions[vm.sp]))
		}
		vm.sp++
	}
	return nil
}

func (vm *VM) push(value interface{}) {
	switch v := value.(type) {
	case int64:
		vm.stackInt = append(vm.stackInt, v)
		vm.stack = append(vm.stack, nil)
	default:
		vm.stack = append(vm.stack, value)
		vm.stackInt = append(vm.stackInt, 0)
	}
}

func (vm *VM) pop() (interface{}, int64) {
	value := vm.stack[len(vm.stack)-1]
	valueInt := vm.stackInt[len(vm.stackInt)-1]
	vm.stackInt = vm.stackInt[:len(vm.stackInt)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return value, valueInt
}
