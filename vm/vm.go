package vm

import (
	"bachelor-thesis/vm/code"
	"encoding/binary"
	"fmt"
	"reflect"
)

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
		stack:        make([]interface{}, 0),
		sp:           0,
	}
}

func (vm *VM) StackTop() interface{} {
	return vm.stack[len(vm.stack)-1]
}

func (vm *VM) Run(env interface{}) error {
	if vm.stack == nil {
		vm.stack = make([]interface{}, 0, 2)
	} else {
		vm.stack = vm.stack[0:0]
	}
	vm.sp = 0
	for vm.sp < len(vm.instructions) {
		switch code.Opcode(vm.instructions[vm.sp]) {
		case code.OpConstant:
			constIndex := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			vm.push(vm.constants[constIndex])
			vm.sp += 2
		case code.OpPop:
			vm.pop()
		case code.OpTrue:
			vm.push(true)
		case code.OpFalse:
			vm.push(false)
		case code.OpNil:
			vm.push(nil)
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
			vm.push(vm.executeMinusOperator())
		case code.OpEqual, code.OpNotEqual, code.OpLessThan, code.OpGreaterThan, code.OpLessOrEqual, code.OpGreaterOrEqual:
			vm.push(vm.executeComparisonOperation(code.Opcode(vm.instructions[vm.sp])))
		case code.OpArray:
			numElements := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			vm.sp += 2
			array := make([]interface{}, numElements)
			for i := numElements - 1; i >= 0; i-- {
				array[i] = vm.pop()
			}
			vm.push(array)
		case code.OpIndex:
			index := vm.pop()
			array := vm.pop()
			vm.executeIndexOperation(array, index)
		case code.OpNot:
			v := vm.pop().(bool)
			vm.push(!v)
		case code.OpJumpIfTrue:
			pos := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			vm.sp += 2
			if vm.StackTop().(bool) {
				vm.sp += pos
			}
		case code.OpJumpIfFalse:
			pos := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			vm.sp += 2
			if !vm.StackTop().(bool) {
				vm.sp += pos
			}
		case code.OpCall:
			fn := reflect.ValueOf(vm.pop())
			size := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			vm.sp += 2
			in := make([]reflect.Value, size)
			for i := int(size) - 1; i >= 0; i-- {
				param := vm.pop()
				if param == nil && reflect.TypeOf(param) == nil {
					in[i] = reflect.ValueOf(&param).Elem()
				} else {
					in[i] = reflect.ValueOf(param)
				}
			}
			out := fn.Call(in)
			if len(out) == 2 && out[1].Type() == reflect.TypeOf((*error)(nil)).Elem() && !out[1].IsNil() {
				panic(out[1].Interface().(error))
			}
			vm.push(out[0].Interface())
		case code.OpLoadConst:
			constIndex := binary.BigEndian.Uint16(vm.instructions[vm.sp+1:])
			vm.sp += 2
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
					vm.push(value.Interface())
				} else {
					elem := reflect.TypeOf(env).Elem()
					vm.push(reflect.Zero(elem).Interface())
				}
			}
		default:
			return fmt.Errorf("unsupported opcode: %d", code.Opcode(vm.instructions[vm.sp]))
		}
		vm.sp++
	}
	return nil
}

func (vm *VM) push(value interface{}) {
	vm.stack = append(vm.stack, value)
}

func (vm *VM) pop() interface{} {
	value := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return value
}
