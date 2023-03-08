package vm

import (
	"bachelor-thesis/vm/code"
	"fmt"
	"math"
)

// TODO: check and change these functions
// TODO: check when it's int64 and when it's int

func (vm *VM) executeAddOperation(a, b interface{}) interface{} {
	switch x := a.(type) {
	case int:
		switch y := b.(type) {
		case int:
			return x + y
		case int64:
			return x + int(y)
		case float64:
			return float64(x) + y
		}
	case int64:
		switch y := b.(type) {
		case int:
			return int(x) + y
		case int64:
			return int(x) + int(y)
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
			return x - int(y)
		case float64:
			return float64(x) - y
		}
	case int64:
		switch y := b.(type) {
		case int:
			return int(x) - y
		case int64:
			return int(x) - int(y)
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
			return int(x) * int(y)
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
			return int(x) / int(y)
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
	panic(fmt.Sprintf("invalid operation: %T ^ %T", a, b))
}

func (vm *VM) executeMinusOperator() interface{} {
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

func (vm *VM) executeComparisonOperation(opcode code.Opcode) interface{} {
	a := vm.pop()
	b := vm.pop()

	switch opcode {
	case code.OpLessThan:
		switch y := a.(type) {
		case int64:
			switch x := b.(type) {
			case int64:
				return int(x) < int(y)
			case float64:
				return x < float64(y)
			}
		case float64:
			switch x := b.(type) {
			case int64:
				return float64(x) < y
			case float64:
				return x < y
			}
		}
	case code.OpLessOrEqual:
		switch y := a.(type) {
		case int64:
			switch x := b.(type) {
			case int64:
				return int(x) <= int(y)
			case float64:
				return x <= float64(y)
			}
		case float64:
			switch x := b.(type) {
			case int64:
				return float64(x) <= y
			case float64:
				return x <= y
			}
		}
	case code.OpGreaterThan:
		switch y := a.(type) {
		case int64:
			switch x := b.(type) {
			case int64:
				return int(x) > int(y)
			case float64:
				return x > float64(y)
			}
		case float64:
			switch x := b.(type) {
			case int64:
				return float64(x) > y
			case float64:
				return x > y
			}
		}
	case code.OpGreaterOrEqual:
		switch y := a.(type) {
		case int64:
			switch x := b.(type) {
			case int64:
				return int(x) >= int(y)
			case float64:
				return x >= float64(y)
			}
		case float64:
			switch x := b.(type) {
			case int64:
				return float64(x) >= y
			case float64:
				return x >= y
			}
		}
	case code.OpEqual:
		switch y := a.(type) {
		case bool:
			switch x := b.(type) {
			case bool:
				return x == y
			}
		case int64:
			switch x := b.(type) {
			case int64:
				return int(x) == int(y)
			case float64:
				return x == float64(y)
			}
		case float64:
			switch x := b.(type) {
			case int64:
				return float64(x) == y
			case float64:
				return x == y
			}
		case string:
			switch x := b.(type) {
			case string:
				return x == y
			}
		}
	case code.OpNotEqual:
		switch y := a.(type) {
		case bool:
			switch x := b.(type) {
			case bool:
				return x != y
			}
		case int64:
			switch x := b.(type) {
			case int64:
				return int(x) != int(y)
			case float64:
				return x != float64(y)
			}
		case float64:
			switch x := b.(type) {
			case int64:
				return float64(x) != y
			case float64:
				return x != y
			}
		case string:
			switch x := b.(type) {
			case string:
				return x != y
			}
		}
	}
	panic(fmt.Sprintf("invalid operation: %T comparison %T", a, b))
}
