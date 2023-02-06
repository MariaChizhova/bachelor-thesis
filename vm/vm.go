package vm

import (
	"bachelor-thesis/code"
	"encoding/binary"
	"fmt"
	"math"
)

const StackSize = 2048

type VM struct {
	constants    []interface{}
	instructions code.Instructions
	stack        []interface{}
	sp           int
}

func New(instructions code.Instructions, constants []interface{}) *VM {
	return &VM{
		instructions: instructions,
		constants:    constants,
		stack:        make([]interface{}, StackSize),
		sp:           0,
	}
}

func (vm *VM) StackTop() interface{} {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) LastPoppedStackElem() interface{} {
	return vm.stack[vm.sp]
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		switch code.Opcode(vm.instructions[ip]) {
		case code.OpConstant:
			constIndex := binary.BigEndian.Uint16(vm.instructions[ip+1:])
			ip += 2
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
		case code.OpAdd:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.add(a, b))
		case code.OpSub:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.substract(b, a))
		case code.OpMul:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.multiply(a, b))
		case code.OpDiv:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.divide(b, a))
		case code.OpMod:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.remainder(b, a))
		case code.OpExp:
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.exponentiation(a, b))
		case code.OpMinus:
			err := vm.push(vm.minusOperator())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (vm *VM) push(value interface{}) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = value
	vm.sp++
	return nil
}

func (vm *VM) pop() interface{} {
	value := vm.stack[vm.sp-1]
	vm.sp--
	return value
}

// TODO: change these functions + move somewhere
func (vm *VM) add(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return int(x) + int(y)
		case float64:
			return float64(x) + y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return int(x) + int(y)
		case float64:
			return x + y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T * %T", a, b))
}

func (vm *VM) substract(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return int(x) - int(y)
		case float64:
			return float64(x) - y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return int(x) - int(y)
		case float64:
			return x - y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T * %T", a, b))
}
func (vm *VM) multiply(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return int(x) * int(y)
		case float64:
			return float64(x) * y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return int(x) * int(y)
		case float64:
			return x * y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T * %T", a, b))
}
func (vm *VM) divide(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return int(x) / int(y)
		case float64:
			return float64(x) / y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return int(x) / int(y)
		case float64:
			return x / y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T * %T", a, b))
}
func (vm *VM) remainder(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return int(x) % int(y)
		}
	}
	panic(fmt.Sprintf("invalid operation: %T * %T", a, b))
}
func (vm *VM) exponentiation(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return int(math.Pow(float64(x), float64(y)))
		case float64:
			return math.Pow(float64(x), y)
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return math.Pow(x, float64(y))
		case float64:
			return math.Pow(x, y)
		}
	}
	panic(fmt.Sprintf("invalid operation: %T * %T", a, b))
}

func (vm *VM) minusOperator() interface{} {
	operand := vm.pop()
	switch x := operand.(type) {
	case int:
		return -x
	case int64:
		return -x
	case float64:
		return -x
	}
	panic(fmt.Errorf("unsupported type for negation: %s", operand))
}
