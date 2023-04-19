package vm5

import (
	"fmt"
	"math"
)

func (vm *VM) executeAddOperation(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int:
		switch y := b.(type) {
		case int:
			return x + y
		case int64:
			return int64(x) + y
		case float64:
			return float64(x) + y
		}
	case int64:
		switch y := b.(type) {
		case int:
			return x + int64(y)
		case int64:
			return x + y
		case float64:
			return float64(x) + y
		}
	case float64:
		switch y := b.(type) {
		case int:
			return x + float64(y)
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
	case int:
		switch y := b.(type) {
		case int:
			return x - y
		case int64:
			return int64(x) - y
		case float64:
			return float64(x) - y
		}
	case int64:
		switch y := b.(type) {
		case int:
			return x - int64(y)
		case int64:
			return x - y
		case float64:
			return float64(x) - y
		}
	case float64:
		switch y := b.(type) {
		case int:
			return x - float64(y)
		case int64:
			return x - float64(y)
		case float64:
			return x - y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T - %T", a, b))
}

func (vm *VM) executeMultiplyOperation(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int:
		switch y := b.(type) {
		case int:
			return x * y
		case int64:
			return int64(x) * y
		case float64:
			return float64(x) * y
		}
	case int64:
		switch y := b.(type) {
		case int:
			return x * int64(y)
		case int64:
			return x * y
		case float64:
			return float64(x) * y
		}
	case float64:
		switch y := b.(type) {
		case int:
			return x * float64(y)
		case int64:
			return x * float64(y)
		case float64:
			return x * y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T * %T", a, b))
}

func (vm *VM) executeDivideOperation(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int:
		switch y := b.(type) {
		case int:
			return x / y
		case int64:
			return int64(x) / y
		case float64:
			return float64(x) / y
		}
	case int64:
		switch y := b.(type) {
		case int:
			return x / int64(y)
		case int64:
			return x / y
		case float64:
			return float64(x) / y
		}
	case float64:
		switch y := b.(type) {
		case int:
			return x / float64(y)
		case int64:
			return x / float64(y)
		case float64:
			return x / y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T / %T", a, b))
}

func (vm *VM) executeRemainderOperation(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int:
		switch y := b.(type) {
		case int:
			return x % y
		case int64:
			return int64(x) % y
		}
	case int64:
		switch y := b.(type) {
		case int:
			return x % int64(y)
		case int64:
			return x % y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T mod %T", a, b))
}
func (vm *VM) executeExponentiationOperation(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int:
		switch y := b.(type) {
		case int:
			return int(math.Pow(float64(x), float64(y)))
		case int64:
			return int(math.Pow(float64(x), float64(y)))
		case float64:
			return math.Pow(float64(x), y)
		}
	case int64:
		switch y := b.(type) {
		case int:
			return int64(math.Pow(float64(x), float64(y)))
		case int64:
			return int64(math.Pow(float64(x), float64(y)))
		case float64:
			return math.Pow(float64(x), y)
		}
	case float64:
		switch y := b.(type) {
		case int:
			return int64(math.Pow(x, float64(y)))
		case int64:
			return math.Pow(x, float64(y))
		case float64:
			return math.Pow(x, y)
		}
	}
	panic(fmt.Sprintf("invalid operation: %T ^ %T", a, b))
}

func (vm *VM) executeMinusOperator(operand interface{}) interface{} {
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
