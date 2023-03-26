package vm2

import (
	"bachelor-thesis/vm/code"
	"encoding/binary"
	"fmt"
)

type VM struct {
	constants    []interface{}
	instructions code.Instructions
	stack        []interface{}
	stackString  []string
	sp           int
}

func New(instructions code.Instructions, constants []interface{}) *VM {
	return &VM{
		instructions: instructions,
		constants:    constants,
		stack:        make([]interface{}, 0),
		stackString:  make([]string, 0),
		sp:           0,
	}
}

func (vm *VM) StackTop() interface{} {
	fmt.Println("stack:", vm.stack)
	fmt.Println("stackString:", vm.stackString)
	fmt.Println("sp:", vm.sp)
	if vm.sp == 0 {
		return nil
	}
	if len(vm.stackString) > 0 {
		return vm.stackString[len(vm.stackString)-1]
	}
	return vm.stack[len(vm.stack)-1]
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		switch code.Opcode(vm.instructions[ip]) {
		case code.OpConstant:
			constIndex := binary.BigEndian.Uint16(vm.instructions[ip+1:])
			ip += 2
			if int(constIndex) >= len(vm.constants) {
				return fmt.Errorf("constant index out of range: %d", constIndex)
			}
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		case code.OpPop:
			vm.pop()
		case code.OpTrue:
			err := vm.push(true)
			if err != nil {
				return err
			}
		case code.OpFalse:
			err := vm.push(false)
			if err != nil {
				return err
			}
		case code.OpNil:
			err := vm.push(nil)
			if err != nil {
				return err
			}
		case code.OpAdd:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.executeAddOperation(b, a))
		case code.OpSub:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.executeSubtractOperation(b, a))
		case code.OpMul:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.executeMultiplyOperation(a, b))
		case code.OpDiv:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.executeDivideOperation(b, a))
		case code.OpMod:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.executeRemainderOperation(b, a))
		case code.OpExp:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.executeExponentiationOperation(b, a))
		case code.OpMinus:
			err := vm.push(vm.executeMinusOperator())
			if err != nil {
				return err
			}
		case code.OpEqual, code.OpNotEqual, code.OpLessThan, code.OpGreaterThan, code.OpLessOrEqual, code.OpGreaterOrEqual:
			vm.push(vm.executeComparisonOperation(code.Opcode(vm.instructions[ip])))
		case code.OpArray:
			numElements := int(binary.BigEndian.Uint16(vm.instructions[ip+1:]))
			ip += 2
			array := make([]interface{}, numElements)
			for i := numElements - 1; i >= 0; i-- {
				array[i] = vm.pop()
			}
			err := vm.push(array)
			if err != nil {
				return err
			}
		case code.OpIndex:
			index := vm.pop()
			array := vm.pop()
			vm.executeIndexOperation(array, index)
		case code.OpNot:
			v := vm.pop().(bool)
			vm.push(!v)
		case code.OpJumpIfTrue:
			pos := int(binary.BigEndian.Uint16(vm.instructions[ip+1:]))
			if vm.StackTop().(bool) {
				ip = pos - 1
			}
		case code.OpJumpIfFalse:
			pos := int(binary.BigEndian.Uint16(vm.instructions[ip+1:]))
			if !vm.StackTop().(bool) {
				ip = pos - 1
			}
		default:
			return fmt.Errorf("unsupported opcode: %d", code.Opcode(vm.instructions[ip]))
		}
	}
	return nil
}

func (vm *VM) push(value interface{}) error {

	fmt.Println("stack:", vm.stack)
	fmt.Println("stackString:", vm.stackString)
	fmt.Println("sp:", vm.sp)
	switch value.(type) {
	case int64, float64, nil, bool, []interface{}:
		vm.stack = append(vm.stack, value)
	case string:
		vm.stackString = append(vm.stackString, value.(string))
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	vm.sp++
	return nil
}

func (vm *VM) pop() interface{} {
	if vm.sp == 0 {
		return nil
	}
	vm.sp--
	if len(vm.stackString) > 0 && vm.sp < len(vm.stackString) {
		value := vm.stackString[len(vm.stackString)-1]
		vm.stackString = vm.stackString[:len(vm.stackString)-1]
		return value
	}
	value := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return value
}
