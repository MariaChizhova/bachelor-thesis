package vm3

import (
	"bachelor-thesis/vm/code"
	"encoding/binary"
	"fmt"
	"reflect"
)

const StackSize = 2048

type VM struct {
	constants    []interface{}
	instructions code.Instructions
	stack        []reflect.Value
	sp           int
}

func New(instructions code.Instructions, constants []interface{}) *VM {
	return &VM{
		instructions: instructions,
		constants:    constants,
		stack:        make([]reflect.Value, StackSize),
		sp:           0,
	}
}

func (vm *VM) StackTop() interface{} {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1].Interface()
}

func (vm *VM) Run(env interface{}) error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		switch code.Opcode(vm.instructions[ip]) {
		case code.OpConstant:
			constIndex := binary.BigEndian.Uint16(vm.instructions[ip+1:])
			ip += 2
			err := vm.push(reflect.ValueOf(vm.constants[constIndex]))
			if err != nil {
				return err
			}
		case code.OpPop:
			vm.pop()
		case code.OpTrue:
			err := vm.push(reflect.ValueOf(true))
			if err != nil {
				return err
			}
		case code.OpFalse:
			err := vm.push(reflect.ValueOf(false))
			if err != nil {
				return err
			}
		case code.OpNil:
			var v any
			err := vm.push(reflect.ValueOf(&v).Elem())
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
			a := vm.pop()
			b := vm.pop()
			vm.push(vm.executeComparisonOperation(b, a, code.Opcode(vm.instructions[ip])))
		case code.OpArray:
			numElements := int(binary.BigEndian.Uint16(vm.instructions[ip+1:]))
			ip += 2
			array := make([]interface{}, numElements)
			for i := numElements - 1; i >= 0; i-- {
				array[i] = vm.pop().Interface()
			}
			err := vm.push(reflect.ValueOf(array))
			if err != nil {
				return err
			}
		case code.OpIndex:
			index := vm.pop()
			array := vm.pop()
			vm.executeIndexOperation(array, index)
		case code.OpNot:
			v := vm.pop()
			vm.push(reflect.ValueOf(!v.Bool()))
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
		case code.OpCall:
			fn := vm.pop()
			size := int(binary.BigEndian.Uint16(vm.instructions[ip+1:]))
			ip += 2
			in := make([]reflect.Value, size)
			for i := int(size) - 1; i >= 0; i-- {
				in[i] = vm.pop()
			}
			out := fn.Call(in)
			vm.push(out[0])
		case code.OpLoadConst:
			constIndex := binary.BigEndian.Uint16(vm.instructions[ip+1:])
			ip += 2
			v := reflect.ValueOf(env)
			kind := v.Kind()
			if kind == reflect.Invalid {
				panic(fmt.Sprintf("cannot fetch %v from %T", vm.constants[constIndex], env))
			}

			if kind == reflect.Ptr {
				v = reflect.Indirect(v)
				kind = v.Kind()
			}

			switch kind {
			case reflect.Map:
				value := v.MapIndex(reflect.ValueOf(vm.constants[constIndex]))
				if value.IsValid() {
					vm.push(value.Elem())
				} else {
					elem := reflect.TypeOf(env)
					vm.push(reflect.Zero(elem).Elem())
				}
			}
		default:
			return fmt.Errorf("unsupported opcode: %d", code.Opcode(vm.instructions[ip]))
		}
	}
	return nil
}

func (vm *VM) push(value reflect.Value) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = value
	vm.sp++
	return nil
}

func (vm *VM) pop() reflect.Value {
	value := vm.stack[vm.sp-1]
	vm.sp--
	return value
}
