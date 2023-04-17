package vm2

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
	stack        []interface{}
	stackString  []string
	sp           int
}

func New(instructions code.Instructions, constants []interface{}) *VM {
	return &VM{
		instructions: instructions,
		constants:    constants,
		stack:        make([]interface{}, StackSize),
		stackString:  make([]string, StackSize),
		sp:           0,
	}
}

func (vm *VM) StackTop() interface{} {
	if vm.sp == 0 {
		return nil
	}
	if vm.stack[vm.sp-1] != nil {
		return vm.stack[vm.sp-1]
	} else if vm.stackString[vm.sp-1] != "" {
		return vm.stackString[vm.sp-1]
	} else {
		return nil
	}
}

func (vm *VM) Run(env interface{}) error {
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
			err := vm.push(vm.executeMinusOperator())
			if err != nil {
				return err
			}
		case code.OpEqual, code.OpNotEqual, code.OpLessThan, code.OpGreaterThan, code.OpLessOrEqual, code.OpGreaterOrEqual:
			a, as := vm.pop()
			b, bs := vm.pop()
			if a == nil && b == nil {
				switch code.Opcode(vm.instructions[ip]) {
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
				vm.push(vm.executeComparisonOperation(a, b, code.Opcode(vm.instructions[ip])))
			}
		case code.OpArray:
			numElements := int(binary.BigEndian.Uint16(vm.instructions[ip+1:]))
			ip += 2
			array := make([]interface{}, numElements)
			for i := numElements - 1; i >= 0; i-- {
				a, as := vm.pop()
				if a == nil {
					array[i] = as
				} else {
					array[i] = a
				}
			}
			err := vm.push(array)
			if err != nil {
				return err
			}
		case code.OpIndex:
			index, _ := vm.pop()
			array, _ := vm.pop()
			vm.executeIndexOperation(array, index)
		case code.OpNot:
			v, _ := vm.pop()
			vm.push(!v.(bool))
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
			elem, _ := vm.pop()
			fn := reflect.ValueOf(elem)
			size := int(binary.BigEndian.Uint16(vm.instructions[ip+1:]))
			ip += 2
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
					vm.push(value.Interface())
				} else {
					elem := reflect.TypeOf(env).Elem()
					vm.push(reflect.Zero(elem).Interface())
				}
			}
		default:
			return fmt.Errorf("unsupported opcode: %d", code.Opcode(vm.instructions[ip]))
		}
	}
	return nil
}

func (vm *VM) push(value interface{}) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	switch value.(type) {
	case string:
		vm.stackString[vm.sp] = value.(string)
		vm.sp++
	default:
		vm.stack[vm.sp] = value
		vm.sp++
	}
	return nil
}

func (vm *VM) pop() (interface{}, string) {
	value := vm.stack[vm.sp-1]
	valueString := vm.stackString[vm.sp-1]
	vm.sp--
	if value == nil {
		vm.stackString[vm.sp] = ""
		return nil, valueString
	}
	vm.stack[vm.sp] = nil
	return value, ""
}
