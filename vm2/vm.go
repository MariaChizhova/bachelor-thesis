package vm2

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
	if vm.stack[len(vm.stack)-1] != nil {
		return vm.stack[len(vm.stack)-1]
	} else if vm.stackString[len(vm.stackString)-1] != "" {
		return vm.stackString[len(vm.stackString)-1]
	} else {
		return nil
	}
}

func (vm *VM) Run(env interface{}) error {
	if vm.stack == nil {
		vm.stack = make([]interface{}, 0, 2)
	} else {
		vm.stack = vm.stack[0:0]
	}
	if vm.stackString == nil {
		vm.stackString = make([]string, 0, 2)
	} else {
		vm.stackString = vm.stackString[0:0]
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
			a, as := vm.pop()
			b, bs := vm.pop()
			if a == nil && b == nil {
				vm.push(bs + as)
			} else {
				vm.push(vm.executeAddOperation(b, a))
			}
		case code.OpSub:
			a, _ := vm.pop()
			b, _ := vm.pop()
			vm.push(vm.executeSubtractOperation(b, a))
		case code.OpMul:
			a, _ := vm.pop()
			b, _ := vm.pop()
			vm.push(vm.executeMultiplyOperation(a, b))
		case code.OpDiv:
			a, _ := vm.pop()
			b, _ := vm.pop()
			vm.push(vm.executeDivideOperation(b, a))
		case code.OpMod:
			a, _ := vm.pop()
			b, _ := vm.pop()
			vm.push(vm.executeRemainderOperation(b, a))
		case code.OpExp:
			a, _ := vm.pop()
			b, _ := vm.pop()
			vm.push(vm.executeExponentiationOperation(b, a))
		case code.OpMinus:
			vm.push(vm.executeMinusOperator())
		case code.OpEqual, code.OpNotEqual, code.OpLessThan, code.OpGreaterThan, code.OpLessOrEqual, code.OpGreaterOrEqual:
			a, as := vm.pop()
			b, bs := vm.pop()
			if a == nil && b == nil {
				switch code.Opcode(vm.instructions[vm.sp]) {
				case code.OpEqual:
					vm.push(as == bs)
				case code.OpNotEqual:
					vm.push(as != bs)
				case code.OpLessThan:
					vm.push(as < bs)
				case code.OpGreaterThan:
					vm.push(as > bs)
				case code.OpLessOrEqual:
					vm.push(as <= bs)
				case code.OpGreaterOrEqual:
					vm.push(as >= bs)
				}
			} else {
				vm.push(vm.executeComparisonOperation(a, b, code.Opcode(vm.instructions[vm.sp])))
			}
		case code.OpArray:
			numElements := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			vm.sp += 2
			array := make([]interface{}, numElements)
			for i := numElements - 1; i >= 0; i-- {
				a, as := vm.pop()
				if a == nil {
					array[i] = as
				} else {
					array[i] = a
				}
			}
			vm.push(array)
		case code.OpIndex:
			index, _ := vm.pop()
			array, _ := vm.pop()
			vm.executeIndexOperation(array, index)
		case code.OpNot:
			v, _ := vm.pop()
			vm.push(!v.(bool))
		case code.OpJumpIfTrue:
			pos := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			if vm.StackTop().(bool) {
				vm.sp = pos - 1
			}
		case code.OpJumpIfFalse:
			pos := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			if !vm.StackTop().(bool) {
				vm.sp = pos - 1
			}
		case code.OpCall:
			elem, _ := vm.pop()
			fn := reflect.ValueOf(elem)
			size := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			vm.sp += 2
			in := make([]reflect.Value, size)
			for i := int(size) - 1; i >= 0; i-- {
				param, s := vm.pop()
				if param == nil {
					param = s
				}
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
	switch value.(type) {
	case string:
		vm.stackString = append(vm.stackString, value.(string))
		vm.stack = append(vm.stack, nil)
	default:
		vm.stack = append(vm.stack, value)
		vm.stackString = append(vm.stackString, "")
	}
}

func (vm *VM) pop() (interface{}, string) {
	value := vm.stack[len(vm.stack)-1]
	valueString := vm.stackString[len(vm.stackString)-1]
	vm.stackString = vm.stackString[:len(vm.stackString)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	if value == nil {
		return nil, valueString
	}
	return value, ""
}
