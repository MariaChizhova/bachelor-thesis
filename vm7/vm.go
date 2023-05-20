package vm7

import (
	"bachelor-thesis/vm/code"
	"encoding/binary"
	"fmt"
	"math"
)

type VM struct {
	constants    []interface{}
	instructions code.Instructions
	stack        []interface{}
	stackString  []string
	stackInt     []int64
	adds         []int // to track from what stack we need to pop and push,
	// 0 - interface, 1 - string, 2 - int
	sp int
}

func New(instructions code.Instructions, constants []interface{}) *VM {
	return &VM{
		instructions: instructions,
		constants:    constants,
		stack:        make([]interface{}, 0),
		stackString:  make([]string, 0),
		stackInt:     make([]int64, 0),
		adds:         make([]int, 0),
		sp:           0,
	}
}

func (vm *VM) StackTop() interface{} {
	switch vm.adds[len(vm.adds)-1] {
	case 0:
		return vm.stack[len(vm.stack)-1]
	case 1:
		return vm.stackString[len(vm.stackString)-1]
	case 2:
		return vm.stackInt[len(vm.stackInt)-1]
	default:
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
			a, as, ai := vm.pop()
			b, bs, bi := vm.pop()
			if a == nil && b == nil && as == "" && bs == "" {
				vm.push(bi + ai)
			} else if a == nil && b == nil {
				vm.push(bs + as)
			} else if a != nil && b != nil {
				vm.push(vm.executeAddOperation(b, a))
			} else if b != nil {
				vm.push(vm.executeAddOperation(b, ai))
			} else {
				vm.push(vm.executeAddOperation(bi, a))
			}
		case code.OpSub:
			a, _, ai := vm.pop()
			b, _, bi := vm.pop()
			if a == nil && b == nil {
				vm.push(bi - ai)
			} else if a != nil && b != nil {
				vm.push(vm.executeSubtractOperation(b, a))
			} else if b != nil {
				vm.push(vm.executeSubtractOperation(b, ai))
			} else {
				vm.push(vm.executeSubtractOperation(bi, a))
			}
		case code.OpMul:
			a, _, ai := vm.pop()
			b, _, bi := vm.pop()
			if a == nil && b == nil {
				vm.push(bi * ai)
			} else if a != nil && b != nil {
				vm.push(vm.executeMultiplyOperation(b, a))
			} else if b != nil {
				vm.push(vm.executeMultiplyOperation(b, ai))
			} else {
				vm.push(vm.executeMultiplyOperation(bi, a))
			}
		case code.OpDiv:
			a, _, ai := vm.pop()
			b, _, bi := vm.pop()
			if a == nil && b == nil {
				vm.push(bi / ai)
			} else if a != nil && b != nil {
				vm.push(vm.executeDivideOperation(b, a))
			} else if b != nil {
				vm.push(vm.executeDivideOperation(b, ai))
			} else {
				vm.push(vm.executeDivideOperation(bi, a))
			}
		case code.OpMod:
			a, _, ai := vm.pop()
			b, _, bi := vm.pop()
			if a == nil && b == nil {
				vm.push(bi % ai)
			} else {
				panic(fmt.Sprintf("invalid operation: %T mod %T", a, b))
			}
		case code.OpExp:
			a, _, ai := vm.pop()
			b, _, bi := vm.pop()
			if a == nil && b == nil {
				vm.push(int64(math.Pow(float64(bi), float64(ai))))
			} else if a != nil && b != nil {
				vm.push(vm.executeExponentiationOperation(b, a))
			} else if b != nil {
				vm.push(vm.executeExponentiationOperation(b, ai))
			} else {
				vm.push(vm.executeExponentiationOperation(bi, a))
			}
		case code.OpMinus:
			vm.push(vm.executeMinusOperator())
		case code.OpEqual, code.OpNotEqual, code.OpLessThan, code.OpGreaterThan, code.OpLessOrEqual, code.OpGreaterOrEqual:
			a, as, ai := vm.pop()
			b, bs, bi := vm.pop()
			if a == nil && b == nil && as == "" && bs == "" {
				switch code.Opcode(vm.instructions[vm.sp]) {
				case code.OpLessThan:
					vm.push(bi < ai)
				case code.OpGreaterThan:
					vm.push(bi > ai)
				case code.OpLessOrEqual:
					vm.push(bi <= ai)
				case code.OpGreaterOrEqual:
					vm.push(bi >= ai)
				case code.OpNotEqual:
					vm.push(bi != ai)
				case code.OpEqual:
					vm.push(bi == ai)
				}
			} else if a == nil && b == nil && as != "" && bs != "" {
				switch code.Opcode(vm.instructions[vm.sp]) {
				case code.OpLessThan:
					vm.push(bs < as)
				case code.OpGreaterThan:
					vm.push(bs > as)
				case code.OpLessOrEqual:
					vm.push(bs <= as)
				case code.OpGreaterOrEqual:
					vm.push(bs >= as)
				case code.OpNotEqual:
					vm.push(bs != as)
				case code.OpEqual:
					vm.push(bs == as)
				}
			} else if a != nil && b != nil {
				vm.push(vm.executeComparisonOperation(a, b, code.Opcode(vm.instructions[vm.sp])))
			} else if b != nil {
				vm.push(vm.executeComparisonOperation(ai, b, code.Opcode(vm.instructions[vm.sp])))
			} else {
				vm.push(vm.executeComparisonOperation(a, bi, code.Opcode(vm.instructions[vm.sp])))
			}
		case code.OpArray:
			numElements := int(binary.BigEndian.Uint16(vm.instructions[vm.sp+1:]))
			vm.sp += 2
			array := make([]interface{}, numElements)
			for i := numElements - 1; i >= 0; i-- {
				a, as, ai := vm.pop()
				if a == nil && as == "" {
					array[i] = ai
				} else if a == nil {
					array[i] = as
				} else {
					array[i] = a
				}
			}
			vm.push(array)
		case code.OpIndex:
			_, _, index := vm.pop()
			array, _, _ := vm.pop()
			vm.executeIndexOperation(array, index)
		case code.OpNot:
			v, _, _ := vm.pop()
			vm.push(!v.(bool))
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
		default:
			return fmt.Errorf("unsupported opcode: %d", code.Opcode(vm.instructions[vm.sp]))
		}
		vm.sp++
	}
	return nil
}

func (vm *VM) push(value interface{}) {
	switch v := value.(type) {
	case string:
		vm.stackString = append(vm.stackString, v)
		vm.adds = append(vm.adds, 1)
	case int64:
		vm.stackInt = append(vm.stackInt, v)
		vm.adds = append(vm.adds, 2)
	default:
		vm.stack = append(vm.stack, value)
		vm.adds = append(vm.adds, 0)
	}
}

func (vm *VM) pop() (interface{}, string, int64) {
	switch vm.adds[len(vm.adds)-1] {
	case 0:
		value := vm.stack[len(vm.stack)-1]
		vm.stack = vm.stack[:len(vm.stack)-1]
		vm.adds = vm.adds[:len(vm.adds)-1]
		return value, "", 0
	case 1:
		valueString := vm.stackString[len(vm.stackString)-1]
		vm.stackString = vm.stackString[:len(vm.stackString)-1]
		vm.adds = vm.adds[:len(vm.adds)-1]
		return nil, valueString, 0
	case 2:
		valueInt := vm.stackInt[len(vm.stackInt)-1]
		vm.stackInt = vm.stackInt[:len(vm.stackInt)-1]
		vm.adds = vm.adds[:len(vm.adds)-1]
		return nil, "", valueInt
	default:
		return nil, "", 0
	}
}
