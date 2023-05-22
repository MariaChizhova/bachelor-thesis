package vm3

import (
	"bachelor-thesis/vm/code"
	"fmt"
	"math"
	"reflect"
)

func (vm *VM) executeAddOperation(a, b reflect.Value) reflect.Value {
	switch a.Kind() {
	case reflect.Int:
		switch b.Kind() {
		case reflect.Int:
			return reflect.ValueOf(a.Int() + b.Int())
		case reflect.Int64:
			return reflect.ValueOf(a.Int() + b.Int())
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) + b.Float())
		}
	case reflect.Int64:
		switch b.Kind() {
		case reflect.Int:
			return reflect.ValueOf(a.Int() + b.Int())
		case reflect.Int64:
			return reflect.ValueOf(a.Int() + b.Int())
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) + b.Float())
		}
	case reflect.Float64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() + float64(b.Int()))
		case reflect.Float64:
			return reflect.ValueOf(a.Float() + b.Float())
		}
	case reflect.String:
		switch b.Kind() {
		case reflect.String:
			return reflect.ValueOf(a.String() + b.String())
		}
	}
	panic(fmt.Sprintf("invalid operation: %T + %T", a, b))
}

func (vm *VM) executeSubtractOperation(a, b reflect.Value) reflect.Value {
	switch a.Kind() {
	case reflect.Int64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() - b.Int())
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) - b.Float())
		}
	case reflect.Float64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() - float64(b.Int()))
		case reflect.Float64:
			return reflect.ValueOf(a.Float() - b.Float())
		}
	}
	panic(fmt.Sprintf("invalid operation: %T - %T", a, b))
}

func (vm *VM) executeMultiplyOperation(a, b reflect.Value) reflect.Value {
	switch a.Kind() {
	case reflect.Int64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() * b.Int())
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) * b.Float())
		}
	case reflect.Float64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() * float64(b.Int()))
		case reflect.Float64:
			return reflect.ValueOf(a.Float() * b.Float())
		}
	}
	panic(fmt.Sprintf("invalid operation: %T * %T", a, b))
}

func (vm *VM) executeDivideOperation(a, b reflect.Value) reflect.Value {
	switch a.Kind() {
	case reflect.Int64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() / b.Int())
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) / b.Float())
		}
	case reflect.Float64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() / float64(b.Int()))
		case reflect.Float64:
			return reflect.ValueOf(a.Float() / b.Float())
		}
	}
	panic(fmt.Sprintf("invalid operation: %T / %T", a, b))
}

func (vm *VM) executeRemainderOperation(a, b reflect.Value) reflect.Value {
	switch a.Kind() {
	case reflect.Int64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() % b.Int())
		}
	}
	panic(fmt.Sprintf("invalid operation: %T mod %T", a, b))
}

func (vm *VM) executeExponentiationOperation(a, b reflect.Value) reflect.Value {
	switch a.Kind() {
	case reflect.Int64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(int64(math.Pow(float64(a.Int()), float64(b.Int()))))
		case reflect.Float64:
			return reflect.ValueOf(math.Pow(a.Float(), b.Float()))
		}
	case reflect.Float64:
		switch b.Kind() {
		case reflect.Int64:
			return reflect.ValueOf(int64(math.Pow(float64(a.Int()), float64(b.Int()))))
		case reflect.Float64:
			return reflect.ValueOf(math.Pow(a.Float(), b.Float()))
		}
	}
	panic(fmt.Sprintf("invalid operation: %T ^ %T", a, b))
}

func (vm *VM) executeMinusOperator() reflect.Value {
	operand := vm.pop()
	switch operand.Kind() {
	case reflect.Int64:
		return reflect.ValueOf(-operand.Int())
	case reflect.Float64:
		return reflect.ValueOf(-operand.Float())
	}
	panic(fmt.Errorf("unsupported type for negation: %s", operand))
}

func (vm *VM) executeComparisonOperation(a, b reflect.Value, opcode code.Opcode) reflect.Value {
	switch opcode {
	case code.OpLessThan:
		switch a.Kind() {
		case reflect.Int64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Int() < b.Int())
			case reflect.Float64:
				return reflect.ValueOf(float64(a.Int()) < b.Float())
			}
		case reflect.Float64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Float() < float64(b.Int()))
			case reflect.Float64:
				return reflect.ValueOf(a.Float() < b.Float())
			}
		case reflect.String:
			switch b.Kind() {
			case reflect.String:
				return reflect.ValueOf(a.String() < b.String())
			}
		}
	case code.OpLessOrEqual:
		switch a.Kind() {
		case reflect.Int64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Int() <= b.Int())
			case reflect.Float64:
				return reflect.ValueOf(float64(a.Int()) <= b.Float())
			}
		case reflect.Float64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Float() <= float64(b.Int()))
			case reflect.Float64:
				return reflect.ValueOf(a.Float() <= b.Float())
			}
		case reflect.String:
			switch b.Kind() {
			case reflect.String:
				return reflect.ValueOf(a.String() <= b.String())
			}
		}
	case code.OpGreaterThan:
		switch a.Kind() {
		case reflect.Int64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Int() > b.Int())
			case reflect.Float64:
				return reflect.ValueOf(float64(a.Int()) > b.Float())
			}
		case reflect.Float64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Float() > float64(b.Int()))
			case reflect.Float64:
				return reflect.ValueOf(a.Float() > b.Float())
			}
		case reflect.String:
			switch b.Kind() {
			case reflect.String:
				return reflect.ValueOf(a.String() > b.String())
			}
		}
	case code.OpGreaterOrEqual:
		switch a.Kind() {
		case reflect.Int64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Int() >= b.Int())
			case reflect.Float64:
				return reflect.ValueOf(float64(a.Int()) >= b.Float())
			}
		case reflect.Float64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Float() >= float64(b.Int()))
			case reflect.Float64:
				return reflect.ValueOf(a.Float() >= b.Float())
			}
		case reflect.String:
			switch b.Kind() {
			case reflect.String:
				return reflect.ValueOf(a.String() >= b.String())
			}
		}
	case code.OpEqual:
		switch a.Kind() {
		case reflect.Bool:
			switch b.Kind() {
			case reflect.Bool:
				return reflect.ValueOf(a.Bool() == b.Bool())
			}
		case reflect.Int64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Int() == b.Int())
			case reflect.Float64:
				return reflect.ValueOf(float64(a.Int()) == b.Float())
			}
		case reflect.Float64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Float() == float64(b.Int()))
			case reflect.Float64:
				return reflect.ValueOf(a.Float() == b.Float())
			}
		case reflect.String:
			switch b.Kind() {
			case reflect.String:
				return reflect.ValueOf(a.String() == b.String())
			}
		}
	case code.OpNotEqual:
		switch a.Kind() {
		case reflect.Bool:
			switch b.Kind() {
			case reflect.Bool:
				return reflect.ValueOf(a.Bool() != b.Bool())
			}
		case reflect.Int64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Int() != b.Int())
			case reflect.Float64:
				return reflect.ValueOf(float64(a.Int()) != b.Float())
			}
		case reflect.Float64:
			switch b.Kind() {
			case reflect.Int64:
				return reflect.ValueOf(a.Float() != float64(b.Int()))
			case reflect.Float64:
				return reflect.ValueOf(a.Float() != b.Float())
			}
		case reflect.String:
			switch b.Kind() {
			case reflect.String:
				return reflect.ValueOf(a.String() != b.String())
			}
		}
	}
	panic(fmt.Sprintf("invalid operation: %T comparison %T", a, b))
}

func (vm *VM) executeIndexOperation(array reflect.Value, index reflect.Value) {
	arrayObject := array.Interface().([]interface{})
	i := index.Int()
	max := int64(len(arrayObject) - 1)
	if i < 0 || i > max {
		vm.push(reflect.ValueOf(nil))
		return
	}
	switch arrayObject[i].(type) {
	case int, int64, float64, bool, string:
		vm.push(reflect.ValueOf(arrayObject[i]))
		return
	}
}
