package vm4

import (
	"fmt"
)

func (vm *VM) executeAddOperation(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case float64:
			return float64(x) + y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x + float64(y)
		case float64:
			return x + y
		}
	case string:
		switch y := b.(type) {
		case string:
			return x + y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T + %T", a, b))
}

func (vm *VM) executeSubtractOperation(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case float64:
			return float64(x) - y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x - float64(y)
		case float64:
			return x - y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T - %T", a, b))
}
func (vm *VM) executeMinusOperator() interface{} {
	operand, operandInt := vm.pop()
	if operand == nil {
		return -operandInt
	}
	switch x := operand.(type) {
	case int64:
		return -x
	case float64:
		return -x
	}
	panic(fmt.Errorf("unsupported type for negation: %s", operand))
}
